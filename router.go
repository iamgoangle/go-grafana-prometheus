package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/submit", MetricSubmitInfo)
	router.HandleFunc("error", ErrorResponse)
	router.Handle("/metrics", prometheus.Handler())

	return router
}
