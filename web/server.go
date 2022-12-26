package web

import (
	"encoding/json"
	"net/http"

	"github.com/romeq/birdnest-reaktor/api"
)

var sourceData []api.Pilot

func serveList(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "public/index.html")
}

func serveJson(wr http.ResponseWriter, _ *http.Request) {
	wr.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(sourceData)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	wr.Write(json)
}

func SetSourceData(pilots []api.Pilot) {
	sourceData = pilots
}

func Serve(address string) {
	http.HandleFunc("/", serveList)
	http.HandleFunc("/json", serveJson)
	http.ListenAndServe(address, nil)
}
