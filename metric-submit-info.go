package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type SubmitInfo struct {
	Name    string `json:"name"`
	Success string `json:"success"`
}

func MakeSubmitCounter() *prometheus.CounterVec {
	prom := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "submit_policy",
			Help: "Number of pushes to submit policy etl.",
		},
		[]string{"success"},
	)

	return prom
}

func MetricSubmitInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("submit info metric")

	passed := "false"
	if rand.Intn(100) > 50 {
		passed = "true"
	}

	submit := SubmitInfo{
		Name:    "Submit new policy",
		Success: passed,
	}

	submitPolicyTotal.With(prometheus.Labels{"success": passed}).Inc()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(submit)
}
