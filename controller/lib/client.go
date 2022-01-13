package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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

func SetMed(pathToConfig string) {
	var conf pb.SetMedRequest
	raw, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		log.Fatal(err)
	}
	print(string(raw))
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		log.Fatal(err)
	}
	address := "[::1]:50051"
	conn, err := grpc.Dial(
		address,

		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                3 * time.Second,
			Timeout:             3 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	client := pb.NewRouteMapServiceClient(conn)

	// setMedRequest := pb.SetMedRequest{RouteMap: routeMap, SequenceNumber: sequenceNumber, Type: permitDeny, Med: med}

	res, err := client.SetMed(ctx, &conf)

	if err != nil {
		return
	}

	fmt.Println(res)
}

func ConfigInterface(pathToConfig string) {
	var conf pb.ConfigInterfaceRequest
	raw, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		log.Fatal(err)
	}
	print(string(raw))
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		log.Fatal(err)
	}
	address := "[::1]:50051"
	conn, err := grpc.Dial(
		address,

		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                3 * time.Second,
			Timeout:             3 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	client := pb.NewInterfaceServiceClient(conn)

	// setMedRequest := pb.SetMedRequest{RouteMap: routeMap, SequenceNumber: sequenceNumber, Type: permitDeny, Med: med}

	res, err := client.ConfigInterface(ctx, &conf)

	if err != nil {
		return
	}

	fmt.Println(res)
}
