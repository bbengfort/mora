// Listener simply echos any requests it gets for now.
package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/bbengfort/mora"
	pb "github.com/bbengfort/mora/echo"
	"golang.org/x/net/context"
)

func main() {
	// Set up the server
	server := &mora.Node{Name: "Obi Wan Kenobi", Address: "192.168.1.11:3265"}
	local := &mora.Node{Name: "Luke Skywalker", Address: "localhost:3265"}
	deadline := time.Duration(20) * time.Second

	// Connect to the server
	conn, err := grpc.Dial(server.Address, grpc.WithInsecure(), grpc.WithTimeout(deadline))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client := pb.NewEchoClient(conn)

	// Contact the echo server and print out response
	r, err := client.Bounce(context.Background(), &pb.EchoRequest{
		Source:  local.ToEchoNode(),
		Target:  server.ToEchoNode(),
		Sent:    &pb.Time{Nanoseconds: time.Now().UnixNano()},
		Payload: []byte("This is just a test"),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.String())

}
