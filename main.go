package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)
	submitPolicyTotal = MakeSubmitCounter()
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
	prometheus.MustRegister(submitPolicyTotal)
}

func Index(w http.ResponseWriter, r *http.Request) {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
}

func main() {
	router := CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
