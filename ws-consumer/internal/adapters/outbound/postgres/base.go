package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	BASE_USER_QUERY = `
		SELECT id, 
			email, 
			password_hash,
			phone,
			EXTRACT(EPOCH FROM created_at) as created_at,
			EXTRACT(EPOCH FROM updated_at) as updated_at
		FROM users`

	BASE_EVENT_QUERY = `
		INSERT INTO events (id, user_id, data)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, data, EXTRACT(EPOCH FROM created_at) as created_at, EXTRACT(EPOCH FROM updated_at) as updated_at
	`
)

type QueryBuilder struct {
	query      string
	where      []string
	orderBy    string
	pagination string
	args       []any
	paramCount int
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		where: make([]string, 0, 4),
		args:  make([]any, 0, 4),
	}
}

func (qb *QueryBuilder) Query(query string) *QueryBuilder {
	qb.query = query
	return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.where = append(qb.where, condition)
	return qb
}

func (qb *QueryBuilder) OrderBy(orderBy string) *QueryBuilder {
	qb.orderBy = orderBy
	return qb
}

func (qb *QueryBuilder) Paginate(pagination string) *QueryBuilder {
	qb.pagination = pagination
	return qb
}

func (qb *QueryBuilder) AddArgs(args ...any) *QueryBuilder {
	for _, arg := range args {
		qb.paramCount++
		qb.args = append(qb.args, arg)
	}
	return qb
}

func (qb *QueryBuilder) GetArgs() []any {
	return qb.args
}

func (qb *QueryBuilder) Build() string {
	var query strings.Builder
	query.WriteString(qb.query)

	if len(qb.where) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(qb.where, " AND "))
	}

	if qb.orderBy != "" {
		query.WriteString(" ORDER BY ")
		query.WriteString(qb.orderBy)
	}

	if qb.pagination != "" {
		query.WriteString(" ")
		query.WriteString(qb.pagination)
	}

	return query.String()
}

func MultipleQuery[T any](ctx context.Context, db *sqlx.DB, query string, args ...interface{}) ([]*T, error) {
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanRows[T](rows)
}

func ScanRows[T any](rows *sqlx.Rows) ([]*T, error) {
	var results []*T
	for rows.Next() {
		var item T
		if err := rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, nil
}
