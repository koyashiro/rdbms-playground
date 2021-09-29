package model

type Workspace struct {
	ID        string     `json:"id"`
	Container *Container `json:"container"`
}
