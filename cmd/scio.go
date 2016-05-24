// This command implements the Mora library as a background daemon that can
// be managed with LaunchAgent or Upstart on OS X and Ubuntu machines (Windows
// builds to follow once we have a requirement for that).
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/bbengfort/mora"
	"github.com/codegangsta/cli"
	"github.com/joho/godotenv"
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
	sonar, err := mora.New()

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// Gets
	fmt.Println(sonar.Scribo.GetNodes())
	// dumpResponse(sonar.Scribo.Get(mora.NODES))
	// dumpResponse(sonar.Scribo.Get(mora.NODES, "1"))

	// Post a new node
	node := make(map[string]string)
	node["name"] = "burrito"
	node["address"] = "1.2.3.4"
	node["dns"] = "little.donkey.edu"
	// dumpResponse(sonar.Scribo.Post(node, mora.NODES))
	// dumpResponse(sonar.Scribo.Put(node, mora.NODES, "4"))
	// dumpResponse(sonar.Scribo.Delete(mora.NODES, "4"))

	return nil
}

func dumpResponse(response *http.Response, err error) error {
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	fmt.Println("-----------------")
	body, err := httputil.DumpResponse(response, true)
	if err == nil {
		fmt.Println(string(body))
		fmt.Println("-----------------")
	}

	return err
}
