package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	"github.com/mattn/go-pipeline"
)

func execCommand(vtyshArg string, jqArg string, obj interface{}) error {
	res, cmd_err := pipeline.Output(
		[]string{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c '" + vtyshArg + " json'"},
		[]string{"jq", jqArg},
		[]string{"jq", "select(type != \"null\")"},
	)
	if cmd_err != nil {
		fmt.Println(cmd_err)
	}

	err := json.Unmarshal(res, &obj)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func main() {
	//TODO:  (cd /home/vsix/nfv-kit/ && sudo docker-compose exec frr bash -c "vtysh -c 'show bgp ipv6 sum json'") | jq ".ipv6Unicast" を exec
	var obj pb.ShowRouteMapResult
	err := execCommand("show route-map", ".BGP", &obj)
	// var obj pb.ShowBgpIpv6SummaryResult
	// err := execCommand("show bgp ipv6 sum", ".ipv6Unicast", &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
	// port := 50051
	// listenport, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// // gRPCサーバーの生成
	// server := grpc.NewServer()
	// rt := router.Router{}
	// pb.RegisterShowBgpIpv6SummaryServiceServer(server, &rt)
	// pb.RegisterShowRouteMapServiceServer(server, &rt)

	// reflection.Register(server)
	// server.Serve(listenport)
}
