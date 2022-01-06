package router

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
	"github.com/mattn/go-pipeline"
)

type Router struct{}

func execShowCommand(obj interface{}, vtyshArg string, jqArgs ...string) error {
	var commands = [][]string{
		{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c 'show " + vtyshArg + " json'"},
	}
	for _, jqArg := range jqArgs {
		builtArg := append([]string{"jq"}, jqArg)
		commands = append(commands, builtArg)
	}
	res, cmd_err := pipeline.Output(commands...)
	if cmd_err != nil {
		fmt.Println(cmd_err)
	}
	print(string(res))
	err := json.Unmarshal(res, &obj)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (r *Router) ShowBgpIpv6Summary(ctx context.Context, req *pb.ShowBgpIpv6SummaryRequest) (*pb.ShowBgpIpv6SummaryResult, error) {
	var obj pb.ShowBgpIpv6SummaryResult
	err := execShowCommand(&obj, "bgp ipv6 summary", ".ipv6Unicast")
	return &obj, err
}

func (r *Router) ShowRouteMap(ctx context.Context, req *pb.ShowRouteMapRequest) (*pb.ShowRouteMapResult, error) {
	var obj pb.ShowRouteMapResult
	err := execShowCommand(&obj, "route-map", ".BGP", "select(type != \"null\")", "."+req.GetName())
	return &obj, err
}
