package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpcService "github.com/longbei65/reptilia/grpc-services"
	"github.com/longbei65/reptilia/protobuf/register"
	"github.com/longbei65/reptilia/version"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	register.RegisterRegistrationServer(grpcServer, &grpcService.RegistrationService{})
	log.Println("grpc server start on :50001")
	if err = grpcServer.Serve(lis); err != nil {
		log.Panicln(err)
		return
	}
}
