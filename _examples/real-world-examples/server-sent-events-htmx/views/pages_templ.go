package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"github.com/a-h/templ"
)

type Post struct {
	ID      string
	Content string
	Author  string
	Date    string
}

type Reaction struct {
	ID          string
	Label       string
	Count       string
	JustChanged bool
}

type PostStats struct {
	PostID    string
	Views     PostViews
	Reactions []Reaction
}

type PostViews struct {
	Count       string
	JustChanged bool
}

func Index(posts []Post) templ.Component { _ = "STUB: not implemented"; return *new(templ.Component) }

func Posts(posts []Post) templ.Component { _ = "STUB: not implemented"; return *new(templ.Component) }

func Idle() templ.Component { _ = "STUB: not implemented"; return *new(templ.Component) }

func postView(post Post) templ.Component { _ = "STUB: not implemented"; return *new(templ.Component) }

func PostStatsView(stats PostStats) templ.Component {
	_ = "STUB: not implemented"
	return *new(templ.Component)
}

func reactionButton(postID string, reaction Reaction) templ.Component {
	_ = "STUB: not implemented"
	return *new(templ.Component)
}

func UpdatedButton(label string) templ.Component {
	_ = "STUB: not implemented"
	return *new(templ.Component)
}
