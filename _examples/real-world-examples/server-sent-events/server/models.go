package main

import (
	"time"
)

type Post struct {
	ID      string   `json:"id" bson:"id"`
	Title   string   `json:"title" bson:"title"`
	Content string   `json:"content" bson:"content"`
	Author  string   `json:"author" bson:"author"`
	Tags    []string `json:"tags" bson:"tags"`
}

func NewPost(id, title, content, author string) Post { _ = "STUB: not implemented"; return *new(Post) }

type Feed struct {
	Name  string `json:"name" bson:"_id"`
	Posts []Post `json:"posts" bson:"posts"`
}

type PostCreated struct {
	Post Post `json:"post"`

	OccurredAt time.Time `json:"occurred_at"`
}

type PostUpdated struct {
	OriginalPost Post `json:"original_post"`
	NewPost      Post `json:"new_post"`

	OccurredAt time.Time `json:"occurred_at"`
}

type FeedUpdated struct {
	Name string `json:"name"`

	OccurredAt time.Time `json:"occurred_at"`
}
