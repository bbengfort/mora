package main

import (
	"fmt"
	"log"

	"github.com/bbengfort/mora"
	"github.com/joho/godotenv"
)

var sonar *mora.Sonar
var serr error

func action() error {
	nodes, err := sonar.Scribo.GetNodes()
	if err != nil {
		return err
	}

	for idx, node := range nodes {
		fmt.Printf("%d: %s %s\n", idx, node.Name, node.Address)
	}

	return nil
}

func main() {
	// Load the .env file if it exists
	godotenv.Load()

	sonar, serr = mora.New()
	if serr != nil {
		fmt.Println(serr)
		return
	}

	log.Println("Main started")
	// worker := mora.NewJitteredInterval(action, 5*time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println(sonar.Local)
		err := sonar.Scribo.Sync(sonar.Local)
		if err != nil {
			log.Fatal(err)
		}
	}

	// go worker.Run()
	// time.Sleep(30 * time.Second)
	// log.Println("Main out!")
}
