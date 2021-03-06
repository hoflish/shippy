// services/consignment/main.go
package main

import (
	// Import the generated protobuf code
	"fmt"
	"log"
	"os"

	pb "github.com/hoflish/shippy/services/consignment/proto"
	vesselProto "github.com/hoflish/shippy/services/vessel/proto"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "192.168.99.100:27017"
)

func main() {

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselService("vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
