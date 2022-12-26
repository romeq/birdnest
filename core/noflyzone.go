package core

import (
	"math"

	"github.com/romeq/birdnest-reaktor/api"
)

type NoFlyZone struct {
	origon [2]float64
	radius float64
}

// IsLegal returns bool indicating whether given drone is within the no fly zone
// and also the distance from the drone to the origon in meters
func (n *NoFlyZone) IsLegal(drone api.Drone) (status bool, distance float64) {
	xbmxa := n.origon[0] - drone.PositionX
	ybmya := n.origon[1] - drone.PositionY
	st := math.Pow(xbmxa, 2) + math.Pow(ybmya, 2)
	distance = math.Sqrt(st)
	return distance > n.radius, distance / 1000
}
