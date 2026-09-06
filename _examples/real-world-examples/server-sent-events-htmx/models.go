package main

import "time"

var allReactions = []Reaction{
	{
		ID:    "fire",
		Label: "🔥",
	},
	{
		ID:    "thinking",
		Label: "🤔",
	},
	{
		ID:    "heart",
		Label: "🩵",
	},
	{
		ID:    "laugh",
		Label: "😂",
	},
	{
		ID:    "sad",
		Label: "😢",
	},
}

func mustReactionByID(id string) Reaction { _ = "STUB: not implemented"; return *new(Reaction) }

type Reaction struct {
	ID    string
	Label string
}

type Post struct {
	ID        int
	Author    string
	Content   string
	CreatedAt time.Time
	Views     int
	Reactions map[string]int
}

type PostStats struct {
	ID              int
	Views           int
	ViewsUpdated    bool
	Reactions       map[string]int
	ReactionUpdated *string
}
