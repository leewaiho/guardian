package model

import "time"

type Network struct {
	ID          int64
	Name        string
	CreatedTime time.Time
	UpdatedTime time.Time
}
