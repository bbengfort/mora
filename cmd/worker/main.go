package main

import (
	"log"
	"time"

	"github.com/bbengfort/mora"
)

func action() error {
	time.Sleep(2)
	log.Println("Action complete!")
	return nil
}

func main() {
	log.Println("Main started")
	worker := mora.NewJitteredInterval(action, 5*time.Second)

	go worker.Run()
	time.Sleep(30 * time.Second)
	log.Println("Main out!")
}
