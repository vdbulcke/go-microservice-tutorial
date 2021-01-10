package main


import (
	"github.com/hashicorp/go-hclog"
	licenseprotos "go-microservice-tutorial/license-service/protos/license"
	"go-microservice-tutorial/license-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"fmt"
)	


func main() {
	// new logger
	log := hclog.Default()

	// new GRPC Server
	gs := grpc.NewServer()


	// Create new LicenseServer
	licenseServer := server.NewLicenseServer(log)

	// register our LicenseServer to grpc server
	licenseprotos.RegisterLicenseServiceServer(gs, licenseServer)

	// for debug purposes only 
	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)


}
