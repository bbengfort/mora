package mora

import (
	"encoding/json"

	"github.com/bbengfort/mora/echo"
)

// Node is a model that represents a participant in the network
type Node struct {
	ID      int64  `json:"id"`      // Unique ID of the node
	Name    string `json:"name"`    // Name/DNS of the node
	Address string `json:"address"` // IP Address of the node
	DNS     string `json:"dns"`     // DNS Lookup for the node
}

// Ping is a model that represents a latency report.
type Ping struct {
	ID      int64   `json:"id"`      // Unique ID of the ping
	Source  int64   `json:"source"`  // The ID of the source node
	Target  int64   `json:"target"`  // The ID of the target node
	Payload int     `json:"payload"` // The size in bytes of the payload
	Latency float64 `json:"latency"` // The time in ms of the round trip
	Timeout bool    `json:"timeout"` // Whether or not the request timed out
}

// Nodes is a collection of node items for use elsewhere.
type Nodes []Node

// Pings is a collection of latency reports for use elsewhere.
type Pings []Ping

// GetNodes uses the scribo client to get the most recent list of nodes.
func (scribo *ScriboClient) GetNodes() (Nodes, error) {

	// Use the client to fetch the list of nodes
	response, err := scribo.Get(NODES)
	if err != nil {
		return nil, err
	}

	// Parse the response body into a list of nodes.
	var nodes Nodes
	if err = json.NewDecoder(response.Body).Decode(&nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

// ToEchoNode returns a protocol buffer message ready struct from a node.
func (node Node) ToEchoNode() *echo.Node {
	return &echo.Node{
		Name:    node.Name,
		Address: node.Address,
		Dns:     node.DNS,
	}
}
