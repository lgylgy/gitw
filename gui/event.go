package gui

import "github.com/lgylgy/gitw/git"

type EventType int32

const (
	Draw   EventType = 1
	Add    EventType = 2
	Update EventType = 3
	Remove EventType = 4
)

type Event struct {
	T    EventType
	View string
	Repo *git.Repository
}
