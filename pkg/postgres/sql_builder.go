package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type builder struct {
	*goqu.DialectWrapper
}

func newBuilder() *builder {
	d := goqu.Dialect("postgres")
	return &builder{DialectWrapper: &d}
}

func (b *builder) Gt(value interface{}) exp.BooleanExpression {
	return exp.Default().Gt(value)
}

func (b *builder) Gte(value interface{}) exp.BooleanExpression {
	return exp.Default().Gte(value)
}

func (b *builder) Lt(value interface{}) exp.BooleanExpression {
	return exp.Default().Lt(value)
}

func (b *builder) Lte(value interface{}) exp.BooleanExpression {
	return exp.Default().Lte(value)
}

func (b *builder) Eq(value interface{}) exp.BooleanExpression {
	return exp.Default().Eq(value)
}

func (b *builder) In(value ...interface{}) exp.BooleanExpression {
	return exp.Default().In(value...)
}

func (b *builder) NotEq(value interface{}) exp.BooleanExpression {
	return exp.Default().Neq(value)
}

func (b *builder) Asc(value interface{}) exp.OrderedExpression {
	return exp.Default().Asc()
}

func (b *builder) Desc(value interface{}) exp.OrderedExpression {
	return exp.Default().Desc()
}

func (b *builder) Like(value interface{}) exp.BooleanExpression {
	return exp.Default().Like(value)
}
