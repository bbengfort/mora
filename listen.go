package mora

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/bbengfort/mora/echo"
	"golang.org/x/net/context"
)

// Bounce implements the echo.EchoServer interface on the Sonar struct
func (s *Sonar) Bounce(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {

	// Print the incoming message
	fmt.Println(in.String())

	// Return the echo reply
	return &pb.EchoReply{
		Receiver: s.Local.ToEchoNode(),
		Received: &pb.Time{
			Seconds:     0,
			Nanoseconds: time.Now().UnixNano(),
		},
		Echo: in,
	}, nil

}

// Listen runs the echo server (typically as a Go routine) to respond to echos.
func (s *Sonar) Listen(addr string) error {
	// Create the socket to listen on
	sock, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Log the fact that we are listening on the address we are.
	log.Printf("Listening for Echo Requests on %s\n", addr)

	// Create the grpc server, handler, and listen
	server := grpc.NewServer()
	pb.RegisterEchoServer(server, s)
	server.Serve(sock)

	// Serve until finished
	return nil
}
