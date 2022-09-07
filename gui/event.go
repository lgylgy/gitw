package gui

import "github.com/lgylgy/gitw/git"

type EventType int32

const (
	Add    EventType = 1
	Update EventType = 2
	Remove EventType = 3
)

type Event struct {
	T     EventType
	Views []string
	Repo  *git.Repository
}
