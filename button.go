package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var project string = os.Getenv("PROJECT")
var zone string = os.Getenv("ZONE")
var instance string = os.Getenv("INSTANCE")

func onHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Start(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp, resp.Status)
}

func offHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Stop(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp, resp.Status)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Get(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(resp.Status, resp.StatusMessage)
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
