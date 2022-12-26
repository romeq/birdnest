package core

import (
	"log"

	"github.com/romeq/birdnest-reaktor/api"
)

func GetViolatingDrones(drones []api.Drone, dv api.DeviceInformation) (violating []api.Drone) {
	for _, drone := range drones {
		valid, dist := VerifyDrone(drone, dv)
		if !valid {
			drone.LatestDistance = dist
			violating = append(violating, drone)
		}
	}

	log.Println("found", len(violating), "violating drones")
	return violating
}
