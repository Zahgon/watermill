package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

var generatedTags = []string{"watermill", "golang", "pubsub", "unicorn", "HelloWorld", "example", "ThreeDotsLabs"}

type Router struct {
	Subscriber   message.Subscriber
	Publisher    Publisher
	PostsStorage PostsStorage
	FeedsStorage FeedsStorage
	Logger       watermill.LoggerAdapter
}

func (router Router) Mux() *chi.Mux { _ = "STUB: not implemented"; return nil }

type feedSummary struct {
	Name  string `json:"name"`
	Posts int    `json:"posts"`
}

type AllFeedsResponse struct {
	Feeds []feedSummary `json:"feeds"`
}

type allFeedsStreamAdapter struct {
	storage FeedsStorage
	logger  watermill.LoggerAdapter
}

func (f allFeedsStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f allFeedsStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f allFeedsStreamAdapter) getResponse(r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (router Router) CreatePost(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (router Router) GeneratePost(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (router Router) addPost(ctx context.Context, post Post) error {
	_ = "STUB: not implemented"
	return nil
}

type UpdatePostRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (router Router) UpdatePost(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type feedStreamAdapter struct {
	storage FeedsStorage
	logger  watermill.LoggerAdapter
}

func (f feedStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f feedStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f feedStreamAdapter) getResponse(r *http.Request) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type postStreamAdapter struct {
	storage PostsStorage
	logger  watermill.LoggerAdapter
}

func (p postStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p postStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p postStreamAdapter) getResponse(r *http.Request) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FileServer(r chi.Router, path string, root http.FileSystem) { _ = "STUB: not implemented"; return }

func logAndWriteError(logger watermill.LoggerAdapter, w http.ResponseWriter, err error) {
	_ = "STUB: not implemented"
	return
}
