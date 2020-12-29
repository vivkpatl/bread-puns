package main

import (
	"github.com/vivkpatl/bread-puns/controllers/webController"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
Set up basic application parameters. Mostly just the port bc Elastic Beanstalk is poggers
*/
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	// Set up system interrupts channel
	sigClose := make(chan os.Signal, 1)
	signal.Notify(sigClose, syscall.SIGINT, syscall.SIGTERM)

	// Set up WebController
	wc := webController.NewWebController(port, sigClose)

	// Run da ting
	go wc.Start()
	// go snsc.Start()

	<-sigClose
	log.Println("Shutting down service... bye :)")
}
