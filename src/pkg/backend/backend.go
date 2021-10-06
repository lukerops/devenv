package backend

import (
	"github.com/lucasscvvieira/devenv/pkg/model"
)

type Backend interface {
	GetImages() []model.Image
}
