package main

import (
	"sync"

	"github.com/ThreeDotsLabs/watermill-routing-example/server/common"
)

type storage struct {
	lock             *sync.Mutex
	receivedMessages map[string][]common.MessageReceived
}

func (s *storage) Append(message common.MessageReceived) { _ = "STUB: not implemented"; return }

func (s *storage) PopAll(key string) []common.MessageReceived {
	_ = "STUB: not implemented"
	return nil
}
