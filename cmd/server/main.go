package main

import (
	"flag"
	"log"

	"github.com/longbei65/reptilia/protobuf/register"
	"github.com/longbei65/reptilia/version"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	isVersion *bool = flag.Bool("-version", false, "show build info")
)

func main() {
	flag.Parse()
	if *isVersion {
		version.PrintCliVersion()
	}

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := register.NewRegistrationClient(conn)
	resp, err := client.Register(context.Background(), &register.RegisterRequest{Secret: "Test"})

	log.Println(resp, err)
}
