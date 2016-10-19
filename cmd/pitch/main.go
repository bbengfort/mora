// Listener simply echos any requests it gets for now.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bbengfort/mora"
	"github.com/joho/godotenv"

	pb "github.com/bbengfort/mora/echo"
	"github.com/bbengfort/mora/moralog"
	"google.golang.org/grpc/grpclog"
)

var (
	sonar    *mora.Sonar
	remote   *mora.Node
	interval *mora.JitteredInterval
	response *pb.EchoReply
	err      error
)

func ping() error {
	response, err = sonar.Ping(remote)
	if err != nil {
		return err
	}

	fmt.Println(response)
	return nil
}

func init() {
	// Disable the grpc logger - but apparently this is still taking performance
	// http://stackoverflow.com/questions/10571182/go-disable-a-log-logger
	grpclog.SetLogger(&moralog.NoopLogger{})
}

func main() {
	// Load the .env file if it exists
	godotenv.Load()

	// Set up the server
	remote = &mora.Node{Name: "Apollo", Address: "192.168.1.11:3265"}
	interval = mora.NewJitteredInterval(ping, 8*time.Second)

	sonar, err = mora.New()
	if err != nil {
		log.Fatal(err)
	}

	go interval.Run()

	for i := 0; i < 10; i++ {
		err := <-interval.ErrorChannel
		if err != nil {
			log.Println(err)
		}
	}

	interval.Shutdown()
	log.Fatal("More than 10 errors occurred!")
}
