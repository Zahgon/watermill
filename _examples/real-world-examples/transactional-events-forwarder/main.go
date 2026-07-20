package main

import (
	"context"
	stdSQL "database/sql"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/v2/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/forwarder"
)

const (
	projectID             = "transactional-events"
	forwarderSQLTopic     = "eventsToForward"
	googleCloudEventTopic = "lottery-concluded"

	simulatedErrorProbability = 0.5
)

var (
	logger = watermill.NewStdLogger(false, false)
	db     = createDB()
)

type LotteryConcludedEvent struct {
	LotteryID int `json:"lottery_id"`
}

func main() {

	sqlSubscriber, err := sql.NewSubscriber(
		sql.BeginnerFromStdSQL(db),
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultMySQLSchema{},
			OffsetsAdapter:   sql.DefaultMySQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	expectNoErr(err)

	gcpPublisher, err := googlecloud.NewPublisher(
		googlecloud.PublisherConfig{
			ProjectID: projectID,
		},
		logger,
	)
	expectNoErr(err)

	fwd, err := forwarder.NewForwarder(sqlSubscriber, gcpPublisher, logger, forwarder.Config{
		ForwarderTopic: forwarderSQLTopic,
	})
	expectNoErr(err)

	go func() {
		err := fwd.Run(context.Background())
		expectNoErr(err)
	}()

	go runLotteryService(logger)
	go runPrizeSenderService(logger)

	time.Sleep(time.Second * 60)
}

func runLotteryService(logger watermill.LoggerAdapter) { _ = "STUB: not implemented"; return }

func publishEventAndPersistData(lotteryID int, pickedUser string, logger watermill.LoggerAdapter) error {
	_ = "STUB: not implemented"
	return nil
}

func persistDataAndPublishEvent(lotteryID int, pickedUser string, logger watermill.LoggerAdapter) error {
	_ = "STUB: not implemented"
	return nil
}

func persistDataAndPublishEventInTransaction(lotteryID int, pickedUser string, logger watermill.LoggerAdapter) error {
	_ = "STUB: not implemented"
	return nil
}

func runPrizeSenderService(logger watermill.LoggerAdapter) { _ = "STUB: not implemented"; return }

func expectNoErr(err error) { _ = "STUB: not implemented"; return }

func simulateError() error { _ = "STUB: not implemented"; return nil }

func createDB() *stdSQL.DB { _ = "STUB: not implemented"; return nil }
