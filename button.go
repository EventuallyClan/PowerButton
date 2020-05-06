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

var ctx context.Context = context.Background()

func onHandler(w http.ResponseWriter, r *http.Request) {

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
	fmt.Fprintf(w, "%s, %s", resp.Status, resp.StatusMessage)
}

func v2onHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "GET unsupported on v2")
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	instanceName := r.FormValue("instanceName")

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Start(project, zone, instanceName).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp, resp.Status)
	fmt.Fprintf(w, "%s, %s", resp.Status, resp.StatusMessage)
}

func v2offHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "GET unsupported on v2")
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	instanceName := r.FormValue("instanceName")

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Stop(project, zone, instanceName).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp, resp.Status)
	fmt.Fprintf(w, "%s, %s", resp.Status, resp.StatusMessage)
}

func main() {
	log.Print("Power Button Pushed")

	http.HandleFunc("/on", onHandler)
	http.HandleFunc("/off", offHandler)
	http.HandleFunc("/status", statusHandler)

	http.HandleFunc("v2/on", v2onHandler)
	http.HandleFunc("v2/off", v2offHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
