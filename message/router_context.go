package message

import (
	"context"
)

type ctxKey string

const (
	handlerNameKey    ctxKey = "handler_name"
	publisherNameKey  ctxKey = "publisher_name"
	subscriberNameKey ctxKey = "subscriber_name"
	subscribeTopicKey ctxKey = "subscribe_topic"
	publishTopicKey   ctxKey = "publish_topic"
)

func valFromCtx(ctx context.Context, key ctxKey) string { _ = "STUB: not implemented"; return "" }

func HandlerNameFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func PublisherNameFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func SubscriberNameFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func SubscribeTopicFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func PublishTopicFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
