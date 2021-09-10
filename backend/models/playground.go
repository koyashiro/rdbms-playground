package models

type Playground struct {
	ID        string `json:"id"`
	DB        string `json:"db"`
	Version   string `json:"version"`
	Container *Container
}
