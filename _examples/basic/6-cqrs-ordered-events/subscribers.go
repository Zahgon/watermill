package main

import (
	"context"
	"sync"
)

type SubscriberReadModel struct {
	subscribers map[string]string
	lock        sync.RWMutex
}

func NewSubscriberReadModel() *SubscriberReadModel { _ = "STUB: not implemented"; return nil }

func (m *SubscriberReadModel) OnSubscribed(ctx context.Context, event *SubscriberSubscribed) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SubscriberReadModel) OnUnsubscribed(ctx context.Context, event *SubscriberUnsubscribed) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SubscriberReadModel) OnEmailUpdated(ctx context.Context, event *SubscriberEmailUpdated) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SubscriberReadModel) GetSubscriberCount() int { _ = "STUB: not implemented"; return 0 }
