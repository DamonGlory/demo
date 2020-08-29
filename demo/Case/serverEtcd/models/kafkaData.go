package models

import "time"

type Kafkamsg struct {
	Filename string `json:"filename"`
	Data string `json:"data"`
	Time time.Time `json:"time"`
	Username string `json:"username"`
}
