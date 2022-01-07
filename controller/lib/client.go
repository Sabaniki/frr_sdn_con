package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	"google.golang.org/grpc"
)

func GetShowBgpIpv6Summary() {
	address := "[::1]:50051"
	conn, err := grpc.Dial(
		address,

		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewShowBgpIpv6SummaryServiceClient(conn)
	showBgpIpv6SummaryRequest := pb.ShowBgpIpv6SummaryRequest{}
	res, err := client.ShowBgpIpv6Summary(ctx, &showBgpIpv6SummaryRequest)

	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	fmt.Println(res)
}

func GetRouteMap(routeMap string) {
	address := "[::1]:50051"
	conn, err := grpc.Dial(
		address,

		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewRouteMapServiceClient(conn)
	showRouteMapRequest := pb.ShowRouteMapRequest{Name: routeMap}
	res, err := client.ShowRouteMap(ctx, &showRouteMapRequest)

	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	fmt.Println(res)
}

func SetMed(routeMap string, med int32) {
	address := "[::1]:50051"
	conn, err := grpc.Dial(
		address,

		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewRouteMapServiceClient(conn)

	setMedRequest := pb.SetMedRequest{RouteMap: routeMap, Med: med}

	res, err := client.SetMed(ctx, &setMedRequest)

	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	fmt.Println(res)
}
