package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	pb "github.com/chansuke/go-translator-grpc/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToTranslationServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("cound not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speach, error) {
	return nil, fmt.Errorf("not implemented")
}
