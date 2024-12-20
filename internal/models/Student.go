package models

import "time"

type Student struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Company       string    `json:"company"`
	Band          string    `json:"band"`
	EnrrolledDate time.Time `json:"enrrolledDate"`
}
