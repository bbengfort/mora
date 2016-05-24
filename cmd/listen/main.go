// Listener simply echos any requests it gets for now.
package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/bbengfort/mora"
	pb "github.com/bbengfort/mora/echo"
	"golang.org/x/net/context"
)

const (
	name = "onest"
	addr = "192.168.1.13:3265"
	key  = "ExKspLt9qo5HC59QerE_Squ2iCxSo_TjxonXhxGAQ8Q="
)

// Server implements the echo.EchoServer interface
type server struct {
	node *mora.Node
}

// Echo implements echo.EchoServer
func (s *server) Bounce(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	fmt.Println(in.String())

	return &pb.EchoReply{
		Receiver: s.node.ToEchoNode(),
		Received: &pb.Time{
			Seconds:     0,
			Nanoseconds: time.Now().UnixNano(),
		},
		Echo: in,
	}, nil
}

func main() {
	node := &mora.Node{
		Name:    name,
		Address: addr,
	}

	lis, err := net.Listen("tcp", ":3265")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{node})
	s.Serve(lis)
}
