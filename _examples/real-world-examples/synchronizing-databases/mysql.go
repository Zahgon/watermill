package main

import (
	"time"

	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
)

type mysqlUser struct {
	ID        int64
	User      string
	FirstName string
	LastName  string
	CreatedAt time.Time
}

type mysqlSchemaAdapter struct {
	sql.DefaultMySQLSchema
}

func (m mysqlSchemaAdapter) SchemaInitializingQueries(params sql.SchemaInitializingQueriesParams) ([]sql.Query, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m mysqlSchemaAdapter) InsertQuery(params sql.InsertQueryParams) (sql.Query, error) {
	_ = "STUB: not implemented"
	return *new(sql.Query), nil
}

func (m mysqlSchemaAdapter) SelectQuery(params sql.SelectQueryParams) (sql.Query, error) {
	_ = "STUB: not implemented"
	return *new(sql.Query), nil
}

func (m mysqlSchemaAdapter) UnmarshalMessage(params sql.UnmarshalMessageParams) (_ sql.Row, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Row), nil
}
