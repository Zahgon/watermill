package main

import (
	"time"

	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
)

type postgresUser struct {
	ID        int64
	Username  string
	FullName  string
	CreatedAt time.Time
}

type postgresSchemaAdapter struct {
	sql.DefaultPostgreSQLSchema
}

func (p postgresSchemaAdapter) SchemaInitializingQueries(params sql.SchemaInitializingQueriesParams) ([]sql.Query, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p postgresSchemaAdapter) InsertQuery(params sql.InsertQueryParams) (sql.Query, error) {
	_ = "STUB: not implemented"
	return *new(sql.Query), nil
}

func (p postgresSchemaAdapter) SelectQuery(params sql.SelectQueryParams) (sql.Query, error) {
	_ = "STUB: not implemented"
	return *new(sql.Query), nil
}

func (p postgresSchemaAdapter) UnmarshalMessage(params sql.UnmarshalMessageParams) (sql.Row, error) {
	_ = "STUB: not implemented"
	return *new(sql.Row), nil
}
