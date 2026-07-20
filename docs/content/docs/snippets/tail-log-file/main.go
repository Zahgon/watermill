package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-io/pkg/io"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {

	var alertPublisher message.Publisher

	if len(os.Args) < 2 {
		panic(
			fmt.Errorf("usage: %s /path/to/file.log", os.Args[0]),
		)
	}
	logFile, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	sub, err := io.NewSubscriber(logFile, io.SubscriberConfig{
		UnmarshalFunc: io.PayloadUnmarshalFunc,
	}, watermill.NewStdLogger(true, false))
	if err != nil {
		panic(err)
	}

	lines, err := sub.Subscribe(context.Background(), "")
	if err != nil {
		panic(err)
	}

	for line := range lines {
		if criterion(string(line.Payload)) {
			_ = alertPublisher.Publish("alerts", line)
		}
	}
}

func criterion(line string) bool { _ = "STUB: not implemented"; return false }
