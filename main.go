package main

import (
	"flag"

	"github.com/micro/go-micro"
	"github.com/skratchdot/open-golang/open"

	"log"

	proto "github.com/mier85/minicontrol/proto"
	"golang.org/x/net/context"
)

type (
	Handler struct {
	}
)

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Open(ctx context.Context, in *proto.OpenRequest, out *proto.OpenResponse) error {
	log.Printf("gonna open: %s", in.Url)
	err := open.Run(in.Url)
	if nil != err {
		out.Error = err.Error()
	}
	return nil
}

var sname = flag.String("name", "com.github.mier85.minicontrol", "name the service operates on")

func main() {
	flag.Parse()
	if *sname == "" {
		log.Fatalf("service name may not be empty")
	}
	service := micro.NewService(
		micro.Name(*sname),
		micro.Version("0.1"),
	)
	service.Init()

	proto.RegisterControlHandler(service.Server(), NewHandler())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
