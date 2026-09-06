package tests

import (
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

var defaultTimeout = time.Second * 15

func TestPubSub(
	t *testing.T,
	features Features,
	pubSubConstructor PubSubConstructor,
	consumerGroupPubSubConstructor ConsumerGroupPubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

type Features struct {
	ConsumerGroups bool

	ExactlyOnceDelivery bool

	GuaranteedOrder bool

	GuaranteedOrderWithSingleSubscriber bool

	Persistent bool

	RestartServiceCommand []string

	RequireSingleInstance bool

	NewSubscriberReceivesOldMessages bool

	GenerateTopicFunc func(tctx TestContext) string

	GenerateIDFunc func() TestID

	ForceShort bool

	ContextPreserved bool
}

func RunOnlyFastTests() bool { _ = "STUB: not implemented"; return false }

type PubSubConstructor func(t *testing.T) (message.Publisher, message.Subscriber)

type ConsumerGroupPubSubConstructor func(t *testing.T, consumerGroup string) (message.Publisher, message.Subscriber)

type SimpleMessage struct {
	Num int `json:"num"`
}

func getTestName(testFunc interface{}) string { _ = "STUB: not implemented"; return "" }

type TestID string

func (t TestID) String() string { _ = "STUB: not implemented"; return "" }

func NewTestID() TestID { _ = "STUB: not implemented"; return *new(TestID) }

func NewTestULID() TestID { _ = "STUB: not implemented"; return *new(TestID) }

type TestContext struct {
	TestID TestID

	Features Features
}

func runTest(
	t *testing.T,
	name string,
	fn func(t *testing.T, testCtx TestContext),
	features Features,
	parallel bool,
) {
	_ = "STUB: not implemented"
	return
}

const defaultStressTestTestsCount = 10

func TestPubSubStressTest(
	t *testing.T,
	features Features,
	pubSubConstructor PubSubConstructor,
	consumerGroupPubSubConstructor ConsumerGroupPubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestPublishSubscribe(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestConcurrentSubscribe(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestConcurrentSubscribeMultipleTopics(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestPublishSubscribeInOrder(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestResendOnError(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestNoAck(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestContinueAfterSubscribeClose(
	t *testing.T,
	tCtx TestContext,
	createPubSub PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestConcurrentClose(
	t *testing.T,
	tCtx TestContext,
	createPubSub PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestContinueAfterErrors(
	t *testing.T,
	tCtx TestContext,
	createPubSub PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestConsumerGroups(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor ConsumerGroupPubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestPublisherClose(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestTopic(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestMessageCtx(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

type contextKey string

func TestSubscribeCtx(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestReconnect(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func TestNewSubscriberReceivesOldMessages(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor PubSubConstructor,
) {
	_ = "STUB: not implemented"
	return
}

func restartServer(t *testing.T, features Features) { _ = "STUB: not implemented"; return }

func assertConsumerGroupReceivedMessages(
	t *testing.T,
	tCtx TestContext,
	pubSubConstructor ConsumerGroupPubSubConstructor,
	consumerGroup string,
	topicName string,
	expectedMessages []*message.Message,
) {
	_ = "STUB: not implemented"
	return
}

func testTopicName(tCtx TestContext) string { _ = "STUB: not implemented"; return "" }

func newTestID(tCtx TestContext) TestID { _ = "STUB: not implemented"; return *new(TestID) }

func closePubSub(t *testing.T, pub message.Publisher, sub message.Subscriber) {
	_ = "STUB: not implemented"
	return
}

func generateConsumerGroup(t *testing.T, pubSubConstructor ConsumerGroupPubSubConstructor, topicName string, tCtx TestContext) string {
	_ = "STUB: not implemented"
	return ""
}

func PublishSimpleMessages(t *testing.T, messagesCount int, publisher message.Publisher, topicName string) message.Messages {
	_ = "STUB: not implemented"
	return *new(message.Messages)
}

func PublishSimpleMessagesWithContext(t *testing.T, messagesCount int, contextKeyString string, publisher message.Publisher, topicName string) message.Messages {
	_ = "STUB: not implemented"
	return *new(message.Messages)
}

func AddSimpleMessagesParallel(t *testing.T, messagesCount int, publisher message.Publisher, topicName string, publishers int) message.Messages {
	_ = "STUB: not implemented"
	return *new(message.Messages)
}

func assertMessagesChannelClosed(t *testing.T, messages <-chan *message.Message) bool {
	_ = "STUB: not implemented"
	return false
}

func publishWithRetry(publisher message.Publisher, topic string, messages ...*message.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func bulkRead(testCtx TestContext, messagesCh <-chan *message.Message, limit int, timeout time.Duration) (receivedMessages message.Messages, all bool) {
	_ = "STUB: not implemented"
	return *new(message.Messages), false
}

func createMultipliedSubscriber(t *testing.T, pubSubConstructor PubSubConstructor, subscribersCount int) message.Subscriber {
	_ = "STUB: not implemented"
	return *new(message.Subscriber)
}
