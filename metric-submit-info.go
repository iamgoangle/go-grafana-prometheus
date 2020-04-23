package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type SubmitInfo struct {
	Code   string `json:"code"`
	Method string `json:"method"`
	Status string `json:"status"`
}

var LABELS = []string{"code", "method", "status"}

func MakeSubmitCounter() *prometheus.CounterVec {
	prom := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "proc_submit_data_total",
			Help: "Number of pushes to submit data total.",
		},
		LABELS,
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
		Code:   "200",
		Method: "GET",
		Status: passed,
	}

	procSubmitData.With(prometheus.Labels{"code": "200", "method": "GET", "status": passed}).Inc()
	procSubmitData.WithLabelValues("200", "GET", passed)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(submit)
}
