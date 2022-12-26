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
	BASEURL = "https://assignments.reaktor.com/birdnest"
	pilots  = []api.Pilot{}
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
		pilot.Drone = drone
		pilot.ViolationDate = time.Now()

		if !lo.ContainsBy(pilots, func(a api.Pilot) bool {
			return a.PilotId == pilot.PilotId
		}) {
			pilots = append(pilots, pilot)
		} else {
			pilot.Drone = drone
		}
	}

	pilots = lo.Filter(pilots, func(item api.Pilot, _ int) bool {
		return item.ViolationDate.After(time.Now().Add(-time.Minute * 10))
	})

	web.SetSourceData(pilots)
	return report
}
