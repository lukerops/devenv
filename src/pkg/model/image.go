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

type Images []*Image

func (imgs Images) FilterByNames(names []string) Images {
	var filtered Images

	for _, name := range names {
		filtered = append(filtered, imgs.FilterByName(name)...)
	}

	return filtered
}

func (imgs Images) FilterByName(name string) Images {
	var filtered Images

	for _, image := range imgs {
		for _, imageName := range image.Names {
			if imageName == name {
				filtered = append(filtered, image)
				break
			}
		}
	}

	return filtered
}

func (imgs Images) FilterByLabels(labels map[string]string) Images {
	var filtered Images

	for key, value := range labels {
		filtered = append(filtered, imgs.FilterByLabel(key, value)...)
	}

	return filtered
}

func (imgs Images) FilterByLabel(key, value string) Images {
	var filtered Images

	for _, image := range imgs {
		if label, ok := image.Labels[key]; ok {
			if label == value {
				filtered = append(filtered, image)
			}
		}
	}

	return filtered
}
