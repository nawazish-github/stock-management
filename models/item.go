package models

import "time"

type Item struct {
	ID       int64
	Name     string
	Category string
	Expiry   time.Time
}
