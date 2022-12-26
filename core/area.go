package core

import "github.com/romeq/birdnest-reaktor/api"

func VerifyDrone(drone api.Drone, deviceInformation api.DeviceInformation) (valid bool, distance float64) {
	noFlyZone := NoFlyZone{origon: [2]float64{250_000, 250_000}, radius: 100_000}
	return noFlyZone.IsLegal(drone)
}
