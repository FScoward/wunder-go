package main

import "time"

type List struct {
	ID        int
	Title     string
	OwnerType string
	OwnerID   int
	ListType  string
	Public    bool
	Revision  int
	CreatedAt time.Time
	Type      string
}
