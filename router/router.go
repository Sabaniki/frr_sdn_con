package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listenport, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーの生成
	server := grpc.NewServer()
	rt := Router{}
	pb.RegisterShowBgpIpv6SummaryServiceServer(server, &rt)

	reflection.Register(server)
	server.Serve(listenport)
}

type Router struct{}

func (self *Router) ShowBgpIpv6Summary(ctx context.Context, req *pb.ShowBgpIpv6SummaryRequest) (*pb.ShowBgpIpv6SummaryResult, error) {
	return &pb.ShowBgpIpv6SummaryResult{
		RouterId:        "1.1.1.1",
		As:              65000,
		VrfId:           1,
		VrfName:         "default",
		TableVersion:    9118504,
		RibCount:        123456,
		RibMemory:       48170280,
		PeerCount:       10,
		PeerMemory:      1234322,
		PeerGroupCount:  3,
		PeerGroupMemory: 192,
		Peers: map[string]*pb.BgpIpv6SummaryPeerInfo{
			"2001:db8:1000:1000::1": {
				RemoteAs:                   64600,
				LocalAs:                    64601,
				Version:                    4,
				MsgRcvd:                    211234,
				MsgSent:                    9999,
				TableVersion:               0,
				Outq:                       0,
				Inq:                        0,
				PeerUptime:                 "3d10h20m",
				PeerUptimeMsec:             123456,
				PeerUptimeEstablishedEpoch: 123,
				PfxRcd:                     123,
				PfxSnt:                     321,
				State:                      "Established",
				PeerState:                  "OK",
				ConnectionsEstablished:     5,
				ConnectionsDropped:         2,
				Desc:                       "lo0.hoge.fuga.vsix.wide.ad.jp",
				IdType:                     "ipv6",
			},
			"2001:db8:1000:2000::1": {
				RemoteAs:                   64800,
				LocalAs:                    64801,
				Version:                    4,
				MsgRcvd:                    9823479,
				MsgSent:                    734832,
				TableVersion:               0,
				Outq:                       0,
				Inq:                        0,
				PeerUptime:                 "6d11h40m",
				PeerUptimeMsec:             123456,
				PeerUptimeEstablishedEpoch: 1234,
				PfxRcd:                     1232,
				PfxSnt:                     3211,
				State:                      "Established",
				PeerState:                  "OK",
				ConnectionsEstablished:     5,
				ConnectionsDropped:         2,
				Desc:                       "lo0.foo.bar.vsix.wide.ad.jp",
				IdType:                     "ipv6",
			},
		},
	}, nil
}
