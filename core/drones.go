package core

import (
	"log"

	"github.com/romeq/birdnest-reaktor/api"
)

func GetViolatingDrones(drones []api.Drone, dv api.DeviceInformation) (violating []api.Drone) {
	for _, drone := range drones {
		valid, dist := VerifyDrone(drone, dv)
		drone.LatestDistance = dist
		if drone.SmallestDistance == 0 || drone.SmallestDistance > dist {
			drone.SmallestDistance = dist
		}

		if !valid {
			violating = append(violating, drone)
		}
	}

	log.Println("found", len(violating), "violating drones")
	return violating
}
