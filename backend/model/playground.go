package model

type Playground struct {
	ID        string     `json:"id"`
	Container *Container `json:"container"`
}
