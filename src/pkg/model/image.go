package model

import "time"

type Image struct {
	ID           string
	RepoTags     []string
	Architecture string
	Created      time.Time
	OS           string
	Size         int64
	Labels       map[string]string
}
