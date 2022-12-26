package api

import (
	"time"
)

type Pilot struct {
	PilotId       string    `json:"pilotId"`
	Firstname     string    `json:"firstName"`
	Lastname      string    `json:"lastName"`
	PhoneNumber   string    `json:"phoneNumber"`
	CreatedDt     time.Time `json:"createdDt"`
	Email         string    `json:"email"`
	ViolationDate time.Time `json:"violationDate"`
	Drone         Drone
}
