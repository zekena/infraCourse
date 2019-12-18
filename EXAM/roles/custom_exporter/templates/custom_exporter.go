package main

import (
	"github.com/gorilla/mux"
	"github.com/sparrc/go-ping"
	"net/http"
	"strconv"
)

func getMetrics(w http.ResponseWriter, r *http.Request) {
	pinger, _ := ping.NewPinger("www.google.com")
	pinger.Count = 1
	pinger.Run()
	stats := pinger.Statistics()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("google_ping " + strconv.FormatInt((stats.AvgRtt).Microseconds(), 10)))
}

func main() {
	custom_exporter := mux.NewRouter().StrictSlash(true)
	custom_exporter.HandleFunc("/metrics", getMetrics).Methods("GET")
	http.ListenAndServe(":9105", custom_exporter)
}
