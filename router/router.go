package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	router "github.com/Sabaniki/frr_sdn_con/router/lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//TODO:  (cd /home/vsix/nfv-kit/ && sudo docker-compose exec frr bash -c "vtysh -c 'show bgp ipv6 sum json'") | jq ".ipv6Unicast" を exec
	port := 50051
	listenport, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーの生成
	server := grpc.NewServer()
	rt := router.Router{}
	pb.RegisterShowBgpIpv6SummaryServiceServer(server, &rt)

	reflection.Register(server)
	server.Serve(listenport)
}
