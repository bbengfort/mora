package mora

import (
	"time"

	"golang.org/x/net/context"

	pb "github.com/bbengfort/mora/echo"
	"google.golang.org/grpc"
)

// Timeout is the amount of time sonar will wait for a reply
const Timeout = time.Duration(30) * time.Second

// Ping sends an echo request to a node and records the results, reporting
// them to the Scribo web application. Ping works for a single node only.
func (s *Sonar) Ping(node *Node) (*pb.EchoReply, error) {

	// Connect to the remote node
	conn, err := grpc.Dial(node.Address, grpc.WithInsecure(), grpc.WithTimeout(Timeout))
	if err != nil {
		return nil, err
	}

	// Defer closing the connection and create a new Echo client.
	defer conn.Close()
	client := pb.NewEchoClient(conn)

	// Create an Echo/Bounce request to send to the remote
	request := &pb.EchoRequest{
		Source:  s.Local.ToEchoNode(),
		Target:  node.ToEchoNode(),
		Sent:    &pb.Time{Nanoseconds: time.Now().UnixNano()},
		Payload: []byte("Clutter to be replaced with random or actual data."),
	}

	// Send a Bounce request to the remote node and return
	return client.Bounce(context.Background(), request)

}
