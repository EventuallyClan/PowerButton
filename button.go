package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func onHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Power button on request.")
	//target := os.Getenv("TARGET")
	resp, err := http.NewRequest("POST",
		fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/silent-space-421/zones/us-central1-a/instances/%s/start", target), nil)
	log.Print(resp, err)
}

func offHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Power button off request.")
	target := os.Getenv("TARGET")
	resp, err := http.NewRequest("POST",
		fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/silent-space-421/zones/us-central1-a/instances/%s/stop", target), nil)
	log.Print(resp, err)
}

func main() {
	log.Print("Power Button Pushed")

	http.HandleFunc("/on", onHandler)
	http.HandleFunc("/off", offHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
