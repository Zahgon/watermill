package main

import (
	"context"
	"database/sql"
)

const migration = `
CREATE TABLE IF NOT EXISTS posts (
	id serial PRIMARY KEY,
	author VARCHAR NOT NULL,
	content TEXT NOT NULL,
	views INT NOT NULL DEFAULT 0,
	reactions JSONB NOT NULL DEFAULT '{}',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO posts (id, author, content) VALUES
	(1, 'Miłosz', 'Oh, I remember the days when we used to write code in PHP!'),
	(2, 'Robert', 'Back in my days, we used to write code in assembly!')
ON CONFLICT (id) DO NOTHING;
`

func MigrateDB(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository { _ = "STUB: not implemented"; return nil }

func (s *Repository) PostByID(ctx context.Context, id int) (Post, error) {
	_ = "STUB: not implemented"
	return *new(Post), nil
}

func (s *Repository) AllPosts(ctx context.Context) ([]Post, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Repository) UpdatePost(ctx context.Context, id int, updateFn func(post *Post)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanPost(s scanner) (Post, error) { _ = "STUB: not implemented"; return *new(Post), nil }
