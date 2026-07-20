package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "feeds"

type FeedsStorage struct {
	collection *mongo.Collection
}

func NewFeedsStorage() FeedsStorage { _ = "STUB: not implemented"; return *new(FeedsStorage) }

func (s FeedsStorage) Add(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FeedsStorage) All(ctx context.Context) ([]Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s FeedsStorage) ByName(ctx context.Context, name string) (Feed, error) {
	_ = "STUB: not implemented"
	return *new(Feed), nil
}

func (s FeedsStorage) AppendPost(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FeedsStorage) UpdatePost(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FeedsStorage) updatePostIfPresent(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FeedsStorage) appendPostIfNotPresent(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FeedsStorage) removePostIfNotInFeed(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

func isDuplicateError(err error) bool { _ = "STUB: not implemented"; return false }
