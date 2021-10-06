package model

import "time"

type Image struct {
	ID         string
	Names      []string
	Labels     map[string]string
	Containers int8
	Size       int64
	CreatedAt  time.Time
}
