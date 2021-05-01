package model

type Port struct {
	IsOpen    bool   `json:"is_open"`
	Timestamp string `json:"timestamp"`
}