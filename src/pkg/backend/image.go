package backend

import "github.com/lucasscvvieira/devenv/pkg/model"

type ImageBackend interface {
	Filter(*model.Filter) (model.Images, error)
}
