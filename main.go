package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
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
	submitPolicyTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "submit_policy",
			Help: "Number of pushes to submit policy etl.",
		},
		[]string{"success"},
	)
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

func MetricSubmitInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("submit info metric")
	passed := "false"
	if rand.Intn(100) > 50 {
		passed = "true"
	}

	submitPolicyTotal.With(prometheus.Labels{"success": passed}).Inc()
	fmt.Fprintln(w, "MetricSubmitInfo")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/submit", MetricSubmitInfo)
	router.Handle("/metrics", prometheus.Handler())

	// http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", router))
}
