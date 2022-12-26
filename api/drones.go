package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

type Drone struct {
	SerialNumber     string  `xml:"serialNumber"`
	Model            string  `xml:"model"`
	Manufacturer     string  `xml:"manufacturer"`
	Mac              string  `xml:"mac"`
	Ipv4             string  `xml:"ipv4"`
	Ipv6             string  `xml:"ipv6"`
	Firmware         string  `xml:"firmware"`
	PositionY        float64 `xml:"positionY"`
	PositionX        float64 `xml:"positionX"`
	Altitude         float64 `xml:"altitude"`
	LatestDistance   float64
	SmallestDistance float64
}

func (a *Drone) GetPilot(source string) (pilot Pilot) {
	pilotResponse := request(fmt.Sprintf("%s/%s", source, a.SerialNumber))
	json.Unmarshal(pilotResponse, &pilot)
	return pilot
}

type DeviceInformation struct {
	ListenRange      int       `xml:"listenRange"`
	DeviceStarted    time.Time `xml:"deviceStarted"`
	UptimeSeconds    int       `xml:"uptimeSeconds"`
	UpdateIntervalMs int32     `xml:"updateIntervalMs"`
}

type DroneReport struct {
	DeviceInformation DeviceInformation `xml:"deviceInformation"`
	Capture           []Drone           `xml:"capture>drone"`
}

func GetReport(source string) (drones DroneReport) {
	res := request(source)
	if err := xml.Unmarshal(res, &drones); err != nil {
		log.Fatalln(err)
	}

	return drones
}
