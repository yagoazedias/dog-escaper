package model

import "time"

type Port struct {
	IsOpen    bool   `json:"is_open"`
	Timestamp time.Time `json:"timestamp"`
}

func (Port) TableName() string {
	return "port"
}