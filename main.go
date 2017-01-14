package main

import (
	proto "github.com/mier85/minicontrol/proto"

	"flag"
	"log"
	"net"
	"strconv"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type (
	Handler struct {
	}
)

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Open(ctx context.Context, in *proto.OpenRequest) (*proto.OpenResponse, error) {
	log.Printf("i am going to open this url: %s", in.Url)
	err := open.Run(in.Url)
	if nil != err {
		return &proto.OpenResponse{Error: err.Error()}, nil
	}
	return &proto.OpenResponse{Error: ""}, nil
}

var sPort = flag.Int("port", 15555, "port")

func main() {
	flag.Parse()
	if 0 == *sPort {
		log.Fatalf("please specify port")
	}
	l, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(*sPort))
	if nil != err {
		log.Fatalf("failed listening: %s", err.Error())
	}
	defer l.Close()
	server := grpc.NewServer()

	proto.RegisterControlServer(server, NewHandler())
	err = server.Serve(l)
	if nil != err {
		log.Fatalf("failed serving: %s", err.Error())
	}
}
