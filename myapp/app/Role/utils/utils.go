package utils

import "time"

type Role_mag struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Status   bool      `json:"status"`
	Sort     int       `json:"sort"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
