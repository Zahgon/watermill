package main

import (
	"context"
	"main/views"
	"net/http"
	"sync/atomic"

	watermillhttp "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo     *Repository
	eventBus *cqrs.EventBus
}

func NewHandler(repo *Repository, eventBus *cqrs.EventBus, sseRouter watermillhttp.SSERouter) *echo.Echo {
	_ = "STUB: not implemented"
	return nil
}

func (h Handler) Index(c echo.Context) error { _ = "STUB: not implemented"; return nil }

func (h Handler) Posts(c echo.Context) error { _ = "STUB: not implemented"; return nil }

func (h Handler) Idle(c echo.Context) error { _ = "STUB: not implemented"; return nil }

func (h Handler) allPosts(c echo.Context) ([]views.Post, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h Handler) AddReaction(c echo.Context) error { _ = "STUB: not implemented"; return nil }

type statsStream struct {
	repo *Repository
}

func (s *statsStream) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *statsStream) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func newPostStatsView(ctx context.Context, stats PostStats) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newPostView(p Post) views.Post { _ = "STUB: not implemented"; return *new(views.Post) }

type sseHandlersCounter struct {
	Count atomic.Int64
}

func (s *sseHandlersCounter) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(echo.HandlerFunc)
}
