package main

import (
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("error response")
	w.WriteHeader(http.StatusNotFound)
}
