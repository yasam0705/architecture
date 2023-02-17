package repo

import (
	"context"
	"github/architecture/internal/entity"
	"github/architecture/pkg/postgres"

	"github.com/doug-martin/goqu/v9"
)

type file struct {
	db              *postgres.PostgresDB
	tableName       string
	defaultCapacity uint
}

func NewFileRepo(db *postgres.PostgresDB) *file {
	return &file{
		db:              db,
		tableName:       "file",
		defaultCapacity: 8,
	}
}

func (f *file) Get(ctx context.Context, guid string) (*entity.File, error) {
	query := f.selectQu().Where(goqu.Ex{"guid": guid})

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	var result entity.File

	err = f.db.Pool.QueryRow(ctx, sql, params...).Scan(
		&result.Guid,
		&result.UserId,
		&result.FileName,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, f.db.PgError(err)
	}

	return &result, nil
}

func (f *file) List(ctx context.Context, limit, offset uint64, filter entity.Parameter) ([]*entity.File, error) {
	query := f.selectQu()

	for k, v := range filter {
		switch k {
		case "user_id":
			query = query.Where(goqu.Ex{k: v})
		}
	}

	if limit != 0 {
		query = query.Limit(uint(limit)).Offset(uint(offset))

	}

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := f.db.Pool.Query(ctx, sql, params...)
	if err != nil {
		return nil, f.db.PgError(err)
	}

	var list = make([]*entity.File, 0, f.defaultCapacity)
	for rows.Next() {
		temp := &entity.File{}
		err = rows.Scan(
			&temp.Guid,
			&temp.UserId,
			&temp.FileName,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return nil, f.db.PgError(err)
		}

		list = append(list, temp)
	}
	return list, nil
}

func (f *file) Create(ctx context.Context, m *entity.File) error {
	p := f.paramsQu(m, "create")
	query := f.db.Builder.DialectWrapper.Insert(f.tableName).Rows(p)

	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = f.db.Pool.Exec(ctx, sql, params...)
	if err != nil {
		return f.db.PgError(err)
	}
	return nil
}

func (f *file) Update(ctx context.Context, m *entity.File) error {
	p := f.paramsQu(m, "update")
	query := f.db.Builder.DialectWrapper.Update(f.tableName).Set(p).Where(goqu.Ex{"guid": m.Guid})

	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = f.db.Pool.Exec(ctx, sql, params...)
	if err != nil {
		return f.db.PgError(err)
	}
	return nil
}

func (f *file) selectQu() *goqu.SelectDataset {
	return f.db.
		Builder.
		DialectWrapper.
		From(f.tableName).
		Select(
			"guid",
			"user_id",
			"file_name",
			"created_at",
			"updated_at",
		)
}

func (c *file) paramsQu(m *entity.File, qType string) map[string]interface{} {
	params := map[string]interface{}{
		"file_name":  m.FileName,
		"updated_at": m.UpdatedAt,
	}
	if qType == "create" {
		params["guid"] = m.Guid
		params["user_id"] = m.UserId
		params["created_at"] = m.CreatedAt

	}
	return params
}
