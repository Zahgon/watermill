package main

import (
	"context"
	stdSQL "database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
)

const topic = "counter"

func main() {
	db := createDB()
	logger := watermill.NewStdLogger(false, false)

	go runWatermillRouter(db, logger)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}

type messagePayload struct {
	CounterUUID string `json:"counter_uuid"`
}

func runWatermillRouter(db *stdSQL.DB, logger watermill.LoggerAdapter) {
	_ = "STUB: not implemented"
	return
}

func processMessage(msg *message.Message) error { _ = "STUB: not implemented"; return nil }

func updateDbCounter(ctx context.Context, tx sql.Tx, counterUUD string, counterValue int) error {
	_ = "STUB: not implemented"
	return nil
}

func dbCounterValue(ctx context.Context, tx sql.Tx, counterUUID string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func createDB() *stdSQL.DB { _ = "STUB: not implemented"; return nil }
