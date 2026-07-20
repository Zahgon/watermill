package main

import (
	"context"
	"database/sql"
)

type PostsStorage struct {
	db *sql.DB
}

func NewPostsStorage() PostsStorage { _ = "STUB: not implemented"; return *new(PostsStorage) }

func (s PostsStorage) ByID(ctx context.Context, id string) (Post, error) {
	_ = "STUB: not implemented"
	return *new(Post), nil
}

func (s PostsStorage) Add(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PostsStorage) Update(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}
