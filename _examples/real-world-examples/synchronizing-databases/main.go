package main

import (
	"bytes"
	"context"
	stdSQL "database/sql"
	"encoding/gob"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

var (
	logger        = watermill.NewStdLogger(false, false)
	postgresTable = "users"
	mysqlTable    = "users"
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

	mysqlDB := createMySQLConnection()
	postgresDB := createPostgresConnection()

	subscriber := createSubscriber(mysqlDB)
	publisher := createPublisher(postgresDB)

	go simulateEvents(mysqlDB)

	router.AddHandler(
		"mysql-to-postgres",
		mysqlTable,
		subscriber,
		postgresTable,
		publisher,
		func(msg *message.Message) ([]*message.Message, error) {
			originUser := mysqlUser{}

			decoder := gob.NewDecoder(bytes.NewBuffer(msg.Payload))
			err := decoder.Decode(&originUser)
			if err != nil {
				return nil, err
			}

			log.Printf("received user: %+v", originUser)

			newUser := postgresUser{
				ID:        originUser.ID,
				Username:  originUser.User,
				FullName:  fmt.Sprintf("%s %s", originUser.FirstName, originUser.LastName),
				CreatedAt: originUser.CreatedAt,
			}

			var payload bytes.Buffer
			encoder := gob.NewEncoder(&payload)
			err = encoder.Encode(newUser)
			if err != nil {
				return nil, err
			}

			newMessage := message.NewMessage(watermill.NewULID(), payload.Bytes())

			return []*message.Message{newMessage}, nil
		},
	)

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func createMySQLConnection() *stdSQL.DB { _ = "STUB: not implemented"; return nil }

func createPostgresConnection() *stdSQL.DB { _ = "STUB: not implemented"; return nil }

func createSubscriber(db *stdSQL.DB) message.Subscriber {
	_ = "STUB: not implemented"
	return *new(message.Subscriber)
}

func createPublisher(db *stdSQL.DB) message.Publisher {
	_ = "STUB: not implemented"
	return *new(message.Publisher)
}

func simulateEvents(db *stdSQL.DB) { _ = "STUB: not implemented"; return }
