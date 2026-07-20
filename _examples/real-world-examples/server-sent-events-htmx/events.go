package main

import (
	"github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
)

type PostViewed struct {
	PostID int `json:"post_id"`
}

type PostReactionAdded struct {
	PostID     int    `json:"post_id"`
	ReactionID string `json:"reaction_id"`
}

type PostStatsUpdated struct {
	PostID          int            `json:"post_id"`
	Views           int            `json:"views"`
	ViewsUpdated    bool           `json:"views_updated"`
	Reactions       map[string]int `json:"reactions"`
	ReactionUpdated *string        `json:"reaction_updated"`
}

type Routers struct {
	EventsRouter *message.Router
	SSERouter    http.SSERouter
	EventBus     *cqrs.EventBus
}

func NewRouters(cfg config, repo *Repository) (Routers, error) {
	_ = "STUB: not implemented"
	return *new(Routers), nil
}
