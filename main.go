package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		passed := "false"
		if rand.Intn(100) > 50 {
			passed = "true"
		}

		submitPolicyTotal.With(prometheus.Labels{"success": passed}).Inc()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Println(submitPolicyTotal)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
