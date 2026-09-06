package requestreply

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type PubSubBackend[Result any] struct {
	config    PubSubBackendConfig
	marshaler BackendPubsubMarshaler[Result]
}

func NewPubSubBackend[Result any](
	config PubSubBackendConfig,
	marshaler BackendPubsubMarshaler[Result],
) (*PubSubBackend[Result], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PubSubBackendSubscribeParams struct {
	Command any

	OperationID OperationID
}

type PubSubBackendSubscriberConstructorFn func(PubSubBackendSubscribeParams) (message.Subscriber, error)

type PubSubBackendGenerateSubscribeTopicFn func(PubSubBackendSubscribeParams) (string, error)

type PubSubBackendPublishParams struct {
	Command any

	CommandMessage *message.Message

	OperationID OperationID
}

type PubSubBackendGeneratePublishTopicFn func(PubSubBackendPublishParams) (string, error)

type PubSubBackendOnCommandProcessedParams struct {
	HandleErr error

	PubSubBackendPublishParams
}

type PubSubBackendModifyNotificationMessageFn func(msg *message.Message, params PubSubBackendOnCommandProcessedParams) error

type PubSubBackendOnListenForReplyFinishedFn func(ctx context.Context, params PubSubBackendSubscribeParams)

type ReplyPublishErrorHandler func(replyTopic string, notificationMsg *message.Message, err error) error

type PubSubBackendConfig struct {
	Publisher             message.Publisher
	SubscriberConstructor PubSubBackendSubscriberConstructorFn

	GeneratePublishTopic   PubSubBackendGeneratePublishTopicFn
	GenerateSubscribeTopic PubSubBackendGenerateSubscribeTopicFn

	Logger watermill.LoggerAdapter

	ListenForReplyTimeout *time.Duration

	ModifyNotificationMessage PubSubBackendModifyNotificationMessageFn

	OnListenForReplyFinished PubSubBackendOnListenForReplyFinishedFn

	AckCommandErrors bool

	ReplyPublishErrorHandler ReplyPublishErrorHandler
}

func (p *PubSubBackendConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (p *PubSubBackendConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (p PubSubBackend[Result]) ListenForNotifications(
	ctx context.Context,
	params BackendListenForNotificationsParams,
) (<-chan Reply[Result], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const OperationIDMetadataKey = "_watermill_requestreply_op_id"

func (p PubSubBackend[Result]) OnCommandProcessed(ctx context.Context, params BackendOnCommandProcessedParams[Result]) error {
	_ = "STUB: not implemented"
	return nil
}

func operationIDFromMetadata(msg *message.Message) (OperationID, error) {
	_ = "STUB: not implemented"
	return *new(OperationID), nil
}

func (p PubSubBackend[Result]) handleNotifyMsg(
	msg *message.Message,
	expectedCommandUuid string,
	marshaler BackendPubsubMarshaler[Result],
) (Reply[Result], bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}
