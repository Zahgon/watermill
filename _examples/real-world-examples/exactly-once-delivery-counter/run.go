package main

import (
	stdSQL "database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/google/uuid"
)

const messagesCount = 5000

var restartMySQLAt = map[int]struct{}{
	50:   {},
	1000: {},
	1500: {},
	3000: {},
}

var restartWorkerAt = map[int]struct{}{
	100:  {},
	1500: {},
	1600: {},
	3000: {},
}

const senderGoroutines = 5

func main() {
	db := createDB()
	counterUUID := uuid.New().String()

	wg := &sync.WaitGroup{}
	wg.Add(messagesCount)

	bar := pb.StartNew(messagesCount)

	sendCounter := make(chan struct{}, 0)
	go func() {
		for i := 0; i < messagesCount; i++ {
			sendCounter <- struct{}{}

			if _, ok := restartMySQLAt[i]; ok {
				restartMySQL()
			}
			if _, ok := restartWorkerAt[i]; ok {
				restartWorker()
			}
		}
		close(sendCounter)
	}()

	for i := 0; i < senderGoroutines; i++ {
		go func() {
			for range sendCounter {
				sendCountRequest(counterUUID)
				wg.Done()
				bar.Increment()
			}
		}()
	}

	wg.Wait()
	bar.Finish()

	timeout := time.Now().Add(time.Second * 30)

	fmt.Println("checking counter with DB, expected count:", messagesCount)

	matchedOnce := true

	for {
		if time.Now().After(timeout) {
			fmt.Println("timeout")
			break
		}

		dbCounterValue, err := getDbCounterValue(db, counterUUID)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		fmt.Println("db counter value", dbCounterValue)
		if dbCounterValue == messagesCount {
			if !matchedOnce {

				matchedOnce = true
				time.Sleep(time.Second * 2)
				continue
			} else {
				fmt.Println("expected counter value is matching DB value")
				break
			}
		}

		time.Sleep(time.Second)
	}
}

func getDbCounterValue(db *stdSQL.DB, counterUUID string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func restartWorker() { _ = "STUB: not implemented"; return }

func restartMySQL() { _ = "STUB: not implemented"; return }

func sendCountRequest(counterUUID string) { _ = "STUB: not implemented"; return }

func createDB() *stdSQL.DB { _ = "STUB: not implemented"; return nil }
