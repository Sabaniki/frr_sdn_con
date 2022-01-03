package router

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	"github.com/mattn/go-pipeline"
)

type Router struct{}

func execCommand(vtyshArg string, jqArg string, obj interface{}) error {
	res, cmd_err := pipeline.Output(
		[]string{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c '" + vtyshArg + "'"},
		[]string{"jq", jqArg},
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

func (r *Router) ShowBgpIpv6Summary(ctx context.Context, req *pb.ShowBgpIpv6SummaryRequest) (*pb.ShowBgpIpv6SummaryResult, error) {
	// res, cmd_err := pipeline.Output(
	// 	[]string{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c 'show bgp ipv6 sum json'"},
	// 	[]string{"jq", ".ipv6Unicast"},
	// )
	// if cmd_err != nil {
	// 	fmt.Println(cmd_err)
	// }

	// var obj pb.ShowBgpIpv6SummaryResult
	// json_err := json.Unmarshal(res, &obj)
	// if json_err != nil {
	// 	fmt.Println(json_err)
	// }
	var obj pb.ShowBgpIpv6SummaryResult
	err := execCommand("show bgp ipv6 sum json", ".ipv6Unicast", &obj)
	return &obj, err
}

func (r *Router) ShowRouteMap(ctx context.Context, req *pb.ShowRouteMapRequest) (*pb.ShowRouteMapRequest, error) {
	// res, cmd_err := pipeline.Output(
	// 	[]string{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c 'show bgp ipv6 sum json'"},
	// 	[]string{"jq", ".ipv6Unicast"},
	// )
	// if cmd_err != nil {
	// 	fmt.Println(cmd_err)
	// }

	// var obj pb.ShowBgpIpv6SummaryResult
	// json_err := json.Unmarshal(res, &obj)
	// if json_err != nil {
	// 	fmt.Println(json_err)
	// }
	var obj pb.ShowRouteMapRequest
	err := execCommand("show route-map json", ".BGP", &obj)
	return &obj, err
}
