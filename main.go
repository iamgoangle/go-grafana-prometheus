package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	procSubmitData = MakeSubmitCounter()
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(procSubmitData)
}

func Index(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
