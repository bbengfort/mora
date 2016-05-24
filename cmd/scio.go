// This command implements the Mora library as a background daemon that can
// be managed with LaunchAgent or Upstart on OS X and Ubuntu machines (Windows
// builds to follow once we have a requirement for that).
package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/bbengfort/mora"
	"github.com/codegangsta/cli"
	"github.com/joho/godotenv"

	pb "github.com/bbengfort/mora/echo"
)

func main() {

	// Load the .env file if it exists
	godotenv.Load()

	// Instantiate the command line application.
	app := cli.NewApp()
	app.Name = "scio"
	app.Usage = "run the scio experiment in the background"
	app.Version = mora.Version
	app.Author = "Benjamin Bengfort"
	app.Email = "benjamin@bengfort.com"
	app.EnableBashCompletion = true
	app.Action = beginSonar

	// Run the command line application
	app.Run(os.Args)
}

// Begins the listening and pinging threads
func beginSonar(ctx *cli.Context) error {
	// Set up the server
	server := &mora.Node{Name: "Obi Wan Kenobi", Address: "localhost:3265"}
	deadline := time.Duration(20) * time.Second

	sonar, err := mora.New()

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// Connect to the server
	conn, err := grpc.Dial(server.Address, grpc.WithInsecure(), grpc.WithTimeout(deadline))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	defer conn.Close()
	client := pb.NewEchoClient(conn)

	// Contact the echo server and print out response
	r, err := client.Bounce(context.Background(), &pb.EchoRequest{
		Source:  sonar.Local.ToEchoNode(),
		Target:  server.ToEchoNode(),
		Sent:    &pb.Time{Nanoseconds: time.Now().UnixNano()},
		Payload: []byte("This is just a test"),
	})

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(r.String())

	return nil
}
