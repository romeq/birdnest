package main

import (
	"flag"
	"time"

	"github.com/romeq/birdnest-reaktor/api"
	"github.com/romeq/birdnest-reaktor/core"
	"github.com/romeq/birdnest-reaktor/web"
	"github.com/samber/lo"
)

var (
	BASEURL      = "https://assignments.reaktor.com/birdnest"
	pilotStorage = []api.Pilot{}
)

type DataSources struct {
	Drones string
	Pilots string
}

func main() {
	// parse flags
	var config DataSources
	config.Drones = *flag.String("d", BASEURL+"/drones", "url for retrieving drones")
	config.Pilots = *flag.String("p", BASEURL+"/pilots", "url for retrieving pilots")
	flag.Parse()

	go func() {
		for {
			report := fetchData(config)
			<-time.After(time.Duration(report.DeviceInformation.UpdateIntervalMs) * time.Millisecond)
		}
	}()

	web.Serve(":8080")

}

func fetchData(config DataSources) (report api.DroneReport) {
	report = api.GetReport(config.Drones)
	violating := core.GetViolatingDrones(report.Capture, report.DeviceInformation)

	for _, drone := range violating {
		pilot := drone.GetPilot(config.Pilots)
		pilot.ViolationDate = time.Now()

		if !lo.ContainsBy(pilotStorage, func(a api.Pilot) bool {
			return a.PilotId == pilot.PilotId
		}) {
			pilot.Drone.SmallestDistance = pilot.Drone.LatestDistance
			pilotStorage = append(pilotStorage, pilot)
		} else {
			_, index, _ := lo.FindIndexOf(pilotStorage, func(item api.Pilot) bool {
				return item.PilotId == pilot.PilotId
			})

			pilotStorage[index].Drone.LatestDistance = pilot.Drone.LatestDistance

			sm := pilotStorage[index].Drone.SmallestDistance
			ld := pilotStorage[index].Drone.LatestDistance
			if ld <= sm {
				pilotStorage[index].Drone.SmallestDistance = ld
			}
		}
	}

	pilotStorage = lo.Filter(pilotStorage, func(item api.Pilot, _ int) bool {
		return item.ViolationDate.After(time.Now().Add(-time.Minute * 10))
	})

	web.SetSourceData(pilotStorage)
	return report
}
