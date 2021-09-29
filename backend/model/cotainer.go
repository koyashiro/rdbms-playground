package model

import (
	"github.com/docker/docker/api/types"
)

type Container struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func NewContainerFromContainer(c *types.Container) *Container {
	var name string
	if c.Names[0][0] == '/' {
		name = c.Names[0][1:]
	} else {
		name = c.Names[0]
	}

	return &Container{
		ID:    c.ID,
		Name:  name,
		Image: c.Image,
	}
}

func NewContainerFromContainerJSON(cj *types.ContainerJSON) *Container {
	var name string
	if cj.Name[0] == '/' {
		name = cj.Name[1:]
	} else {
		name = cj.Name
	}

	return &Container{
		ID:    cj.ID,
		Name:  name,
		Image: cj.Config.Image,
	}
}
