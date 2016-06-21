// This command implements the Mora library as a background daemon that can
// be managed with LaunchAgent or Upstart on OS X and Ubuntu machines (Windows
// builds to follow once we have a requirement for that).
package main

import (
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

	if err = sonar.Run(); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
