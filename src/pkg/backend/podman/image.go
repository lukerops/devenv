package podman

import (
	"bytes"
	"encoding/json"

	"github.com/lucasscvvieira/devenv/pkg/model"
	"github.com/lucasscvvieira/devenv/pkg/shell"
)

type Image struct{}

func (*Image) Filter(filter *model.Filter) (model.Images, error) {
	var stdout bytes.Buffer

	args := []string{"images", "--format", "json"}
	if err := shell.Run("podman", nil, &stdout, nil, args...); err != nil {
		return nil, err
	}

	output := stdout.Bytes()
	var images model.Images

	if err := json.Unmarshal(output, &images); err != nil {
		return nil, err
	}

	if filter == nil {
		return images, nil
	}

	var filtered model.Images

	if len(filter.Names) > 0 {
		filtered = append(filtered, images.FilterByNames(filter.Names)...)
	}

	if len(filter.Labels) > 0 {
		filtered = append(filtered, images.FilterByLabels(filter.Labels)...)
	}

	return filtered, nil
}
