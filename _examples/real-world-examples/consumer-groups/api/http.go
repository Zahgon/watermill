package main

import (
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	storage *storage

	subscriber message.Subscriber
	publisher  message.Publisher
	logger     watermill.LoggerAdapter

	lastIDs map[string]int
}

func (h Handler) Mux() *chi.Mux { _ = "STUB: not implemented"; return nil }

func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type messagesStream struct {
	storage *storage
	logger  watermill.LoggerAdapter
}

func (p messagesStream) GetResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p messagesStream) Validate(r *http.Request, msg *message.Message) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func fileServer(r chi.Router, path string, root http.FileSystem) { _ = "STUB: not implemented"; return }

func logAndWriteError(logger watermill.LoggerAdapter, w http.ResponseWriter, err error) {
	_ = "STUB: not implemented"
	return
}

func requestIDMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
