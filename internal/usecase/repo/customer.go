package repo

import (
	"context"
	"github/architecture/internal/entity"
	"github/architecture/pkg/postgres"

	"github.com/doug-martin/goqu/v9"
)

type customer struct {
	db              *postgres.PostgresDB
	tableName       string
	defaultCapacity uint
}

func NewCustomerRepo(db *postgres.PostgresDB) *customer {
	return &customer{
		db:              db,
		tableName:       "customer",
		defaultCapacity: 8,
	}
}

func (c *customer) Create(ctx context.Context, m *entity.Customer) error {
	p := c.paramsQu(m, "create")
	query := c.db.Builder.DialectWrapper.Insert(c.tableName).Rows(p)

	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = c.db.Pool.Exec(ctx, sql, params...)
	if err != nil {
		return c.db.PgError(err)
	}
	return nil
}

func (c *customer) Update(ctx context.Context, m *entity.Customer) error {
	p := c.paramsQu(m, "update")
	query := c.db.Builder.DialectWrapper.Update(c.tableName).Set(p)

	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = c.db.Pool.Exec(ctx, sql, params...)
	if err != nil {
		return c.db.PgError(err)
	}
	return nil
}

func (c *customer) Get(ctx context.Context, m map[string]string) (*entity.Customer, error) {
	// p := c.selectQu()
	query := c.db.Builder.DialectWrapper.From(c.tableName)

	for k, v := range m {
		switch k {
		case "guid":
			query.Where(goqu.C("guid").Eq(v))
		}
	}

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	temp := &entity.Customer{}
	err = c.db.Pool.QueryRow(ctx, sql, params...).Scan(
		&temp.GUID,
		&temp.FirstName,
		&temp.LastName,
		&temp.CreatedAt,
		&temp.UpdatedAt,
	)
	if err != nil {
		return nil, c.db.PgError(err)
	}
	return temp, nil
}

func (c *customer) List(ctx context.Context, m map[string]string) ([]*entity.Customer, error) {
	p := c.selectQu()
	query := c.db.Builder.DialectWrapper.From(c.tableName).Select(p)

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := c.db.Pool.Query(ctx, sql, params...)
	if err != nil {
		return nil, c.db.PgError(err)
	}

	var list = make([]*entity.Customer, 0, c.defaultCapacity)
	for rows.Next() {
		temp := &entity.Customer{}
		err = rows.Scan(
			&temp.GUID,
			&temp.FirstName,
			&temp.LastName,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return nil, c.db.PgError(err)
		}

		list = append(list, temp)
	}

	return list, nil
}

func (c *customer) paramsQu(m *entity.Customer, qType string) map[string]interface{} {
	params := map[string]interface{}{
		"first_name": m.FirstName,
		"last_name":  m.LastName,
		"updated_at": m.UpdatedAt,
	}
	if qType == "create" {
		params["guid"] = m.GUID
		params["created_at"] = m.CreatedAt

	}
	return params
}

func (c *customer) selectQu() []interface{} {
	return []interface{}{
		"guid",
		"first_name",
		"last_name",
		"created_at",
		"updated_at",
	}
}
