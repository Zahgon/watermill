package main

import (
	"context"
	"sync"
	"time"
)

type ActivityEntry struct {
	Timestamp    time.Time
	SubscriberID string
	ActivityType string
	Details      string
}

type ActivityTimelineReadModel struct {
	activities []ActivityEntry
	lock       sync.RWMutex
}

func NewActivityTimelineModel() *ActivityTimelineReadModel { _ = "STUB: not implemented"; return nil }

func (m *ActivityTimelineReadModel) OnSubscribed(ctx context.Context, event *SubscriberSubscribed) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ActivityTimelineReadModel) OnUnsubscribed(ctx context.Context, event *SubscriberUnsubscribed) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ActivityTimelineReadModel) OnEmailUpdated(ctx context.Context, event *SubscriberEmailUpdated) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ActivityTimelineReadModel) logActivity(entry ActivityEntry) {
	_ = "STUB: not implemented"
	return
}
