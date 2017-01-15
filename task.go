package main

import "time"

type Task struct {
	ID          int
	AssigneeID  int
	CreatedAt   time.Time
	CreatedByID int
	DueDate     time.Time
	ListID      int
	Revision    int
	Starred     bool
	Title       string
}
