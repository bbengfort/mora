// Package mora provides library functionality for a distributed systems
// experiment that measures the latency and uptime of nodes in a heterogeneous,
// partition-prone, user-oriented network. E.g. a network that is not in a
// data center but still has strong requirements for consistency and availability.
package mora

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

// Version specifies the current revision of the Mora library and Scio app.
const Version = "0.1"

// Configuration is loaded from the environment with reasonable defaults but
// will error out if required configuration values are missing.
type Configuration struct {
	Name      string `envconfig:"node_name" required:"true"`
	ScriboURL string `envconfig:"scribo_url" default:"https://mora-scribo.herokuapp.com/"`
	ScriboKey string `envconfig:"scribo_key" required:"true"`
}

// Sonar is the send, respond, and listen application that implements both the
// Scio and Oro interfaces and maintains global configuration information.
type Sonar struct {
	Config Configuration // The configuration loaded from the environment
	Scribo *ScriboClient // A client to connect to the Scribo API
	Local  *Node         // The local node representation
}

// New instantiates a Sonar client and loads the configuration from the
// environment, possibly loading an .env file or a YAML configuration.
func New() (*Sonar, error) {
	sonar := new(Sonar)

	// Process the configuration from the environment
	if err := envconfig.Process("mora", &sonar.Config); err != nil {
		return nil, err
	}

	// Create the ScriboClient and pass a pointer to the configuration
	sonar.Scribo = new(ScriboClient)
	sonar.Scribo.config = &sonar.Config
	sonar.Scribo.client = new(http.Client)

	// Create the local node with IP address discovery
	sonar.Local = &Node{Name: sonar.Config.Name}

	return sonar, nil
}
