package main

import (
	"context"
	"fmt"
	pb "github.com/mikedutuandu/shippy-vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	defaultHost = "mongodb+srv://admin:AAC0w4Q6jvNv4r8Z@bookingcluster-tgpah.gcp.mongodb.net/test?retryWrites=true&w=majority"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.vessel.service"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{
		vesselCollection,
	}


	// Register our implementation with
	pb.RegisterVesselServiceHandler(service.Server(), &handler{repository})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}