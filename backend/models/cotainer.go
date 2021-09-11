package models

type Container struct {
	Hash   string `json:"hash"`
	Image  string `json:"image"`
	Status string `json:"status"`
	Port   int    `json:"port"`
}
