package mora

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bbengfort/mora/echo"
)

// Node is a model that represents a participant in the network
type Node struct {
	ID      int64  `json:"id,omitempty"` // Unique ID of the node
	Name    string `json:"name"`         // Name/DNS of the node
	Address string `json:"address"`      // IP Address of the node
	DNS     string `json:"dns"`          // DNS Lookup for the node
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

// ToEchoNode returns a protocol buffer message ready struct from a node.
func (node Node) ToEchoNode() *echo.Node {
	return &echo.Node{
		Id:      node.ID,
		Name:    node.Name,
		Address: node.Address,
		Dns:     node.DNS,
	}
}

// Sync updates information about a node to the server and fills in missing
// details, particularly the primary key for determining endpoints. Sync will
// create the node with a POST if necessary. Note that information on the
// local node is considered more up to date than on the server Node, only
// missing or zero values will be filled in by Scribo. The primary intent for
// this function is to synchronize the Local node (and act as a heartbeat).
func (scribo *ScriboClient) Sync(node *Node) error {

	// If we don't have an ID, perform a lookup
	if node.ID == 0 {
		lookup, err := scribo.LookupNode(node.Name)
		if err != nil {
			// TODO: Add a POST here to create the node, for now just give up.
			return err
		}

		// Save ID from lookup onto the node.
		node.ID = lookup.ID
	}

	// Ok now PUT the node to the server.
	_, err := scribo.Put(node, NODES, strconv.FormatInt(node.ID, 10))
	if err != nil {
		return err
	}

	return nil
}

// GetNodes uses the Scribo client to get the most recent list of nodes.
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

// LookupNode uses the Scribo client to lookup a node by its name.
// Note this is a heavy weight function because it's doing all the work that
// Scribo.Get does but also has to deal with query and lookup failures.
func (scribo *ScriboClient) LookupNode(name string) (*Node, error) {
	// Create the Nodes endpoint
	endpoint, err := scribo.Endpoint(NODES)
	if err != nil {
		return nil, err
	}

	// Construct an HTTP request
	request, err := http.NewRequest(MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Add the lookup parameter
	query := request.URL.Query()
	query.Add("lookup", name)
	request.URL.RawQuery = query.Encode()

	// Execute the HTTP request
	response, err := scribo.Do(request)
	if err != nil {
		return nil, err
	}

	// Was it a 404?
	if response.StatusCode == 404 {
		return nil, fmt.Errorf("Could not find a node named %s", name)
	}

	// Parse the response body into a list of nodes.
	var nodes Nodes
	if err = json.NewDecoder(response.Body).Decode(&nodes); err != nil {
		return nil, err
	}

	if len(nodes) > 0 {
		return &nodes[0], nil
	}

	return nil, fmt.Errorf("Something went very wrong in LookupNode!")
}
