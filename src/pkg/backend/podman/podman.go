package podman

import (
	"github.com/lucasscvvieira/devenv/pkg/backend"
)

func NewPodman() *backend.Backend {
	return &backend.Backend{
		Image: &Image{},
	}
}
