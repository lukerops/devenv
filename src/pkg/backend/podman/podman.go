package podman

import (
	"bytes"
	"encoding/json"

	"github.com/lucasscvvieira/devenv/pkg/model"
	"github.com/lucasscvvieira/devenv/pkg/shell"
)

type Podman struct {
}

func New() *Podman {
	return &Podman{}
}

func (podman *Podman) GetImages() ([]model.Image, error) {
	var stdout bytes.Buffer

	args := []string{"images", "--format", "json"}
	if err := shell.Run("podman", nil, &stdout, nil, args...); err != nil {
		return nil, err
	}

	output := stdout.Bytes()
	var images []model.Image

	if err := json.Unmarshal(output, &images); err != nil {
		return nil, err
	}

	return images, nil
}
