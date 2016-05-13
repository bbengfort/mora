// This command implements the Mora library as a background daemon that can
// be managed with LaunchAgent or Upstart on OS X and Ubuntu machines (Windows
// builds to follow once we have a requirement for that).
package main

import (
	"os"

	"github.com/bbengfort/mora"
	"github.com/codegangsta/cli"
)

func main() {

	// Insantiate the comamnd line application
	app := cli.NewApp()
	app.Name = "scio"
	app.Usage = "run the scio experiment in the background"
	app.Version = mora.Version
	app.Author = "Benjamin Bengfort"
	app.Email = "benjamin@bengfort.com"
	app.Action = beginSonar

	// Run the command line application
	app.Run(os.Args)
}

// Begins the listening and pinging threads
func beginSonar(ctx *cli.Context) {

}
