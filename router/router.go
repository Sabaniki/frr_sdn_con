package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
)

func main() {
	res := `
	{

		  "routerId":"100.100.30.1",
		  "as":65000,
		  "vrfId":0,
		  "vrfName":"default",
		  "tableVersion":40,
		  "ribCount":3,
		  "ribMemory":552,
		  "peerCount":2,
		  "peerMemory":1480592,
		  "peerGroupCount":1,
		  "peerGroupMemory":64,
		  "peers":{
			"2001:200:e00:b0:4690:0:100:2":{
			  "hostname":"bb02-blue.fujisawa.vsix.wide.ad.jp",
			  "remoteAs":4690,
			  "localAs":65000,
			  "version":4,
			  "msgRcvd":54831,
			  "msgSent":54822,
			  "tableVersion":0,
			  "outq":0,
			  "inq":0,
			  "peerUptime":"04w1d21h",
			  "peerUptimeMsec":2584036000,
			  "peerUptimeEstablishedEpoch":1638307255,
			  "pfxRcd":2,
			  "pfxSnt":0,
			  "state":"Established",
			  "peerState":"OK",
			  "connectionsEstablished":8,
			  "connectionsDropped":7,
			  "desc":"ve1602.bb02-blue.fujisawa.vsix.wide.ad.jp",
			  "idType":"ipv6"
			},
			"2001:200:e00:b0:4690:0:100:3":{
			  "hostname":"bb02-green.fujisawa.vsix.wide.ad.jp",
			  "remoteAs":4690,
			  "localAs":65000,
			  "version":4,
			  "msgRcvd":30939,
			  "msgSent":30573,
			  "tableVersion":0,
			  "outq":0,
			  "inq":0,
			  "peerUptime":"01w4d13h",
			  "peerUptimeMsec":997811000,
			  "peerUptimeEstablishedEpoch":1639893480,
			  "pfxRcd":2,
			  "pfxSnt":0,
			  "state":"Established",
			  "peerState":"OK",
			  "connectionsEstablished":412,
			  "connectionsDropped":411,
			  "desc":"ve1602.bb02-green.fujisawa.vsix.wide.ad.jp",
			  "idType":"ipv6"
			}
		  },
		  "failedPeers":0,
		  "displayedPeers":2,
		  "totalPeers":2,
		  "dynamicPeers":0,
		  "bestPath":{
			"multiPathRelax":"false"
		  }
		}`
	var obj pb.ShowBgpIpv6SummaryResult
	err := json.Unmarshal([]byte(res), &obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(obj.Peers["2001:200:e00:b0:4690:0:100:3"].PeerState)

	// port := 50051
	// listenport, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// // gRPCサーバーの生成
	// server := grpc.NewServer()
	// rt := router.Router{}
	// pb.RegisterShowBgpIpv6SummaryServiceServer(server, &rt)

	// reflection.Register(server)
	// server.Serve(listenport)
}
