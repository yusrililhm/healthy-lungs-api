package entity

import "time"

type Disease struct {
	Id          int
	Name        string
	Description string
	// Symtomps    []*Symtomp
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
