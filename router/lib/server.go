package router

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func execCommand(vtyshArgs ...string) error {
	for i, vtyshArg := range vtyshArgs {
		vtyshArgs[i] = " -c '" + vtyshArg + "'"
	}
	commands := append([]string{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c 'conf t'"}, vtyshArgs...)
	fmt.Println(commands)
	_, err := pipeline.Output(commands)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func toInt64(strVal string) (int64, error) {
	rex := regexp.MustCompile("[0-9]+")
	strVal = rex.FindString(strVal)
	return strconv.ParseInt(strVal, 10, 64)
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

func getMed(req *pb.SetMedRequest) (int32, error) {
	var obj pb.ShowRouteMapResult
	err := execShowCommand(&obj, "route-map", ".BGP", "select(type != \"null\")", "."+req.RouteMap)
	if err != nil {
		return -1, err
	}
	var med int64
	for _, rule := range obj.Rules {
		if rule.SequenceNumber == req.SequenceNumber {
			for _, matchClause := range rule.MatchClauses {
				if strings.Contains(matchClause, "metric") {
					med, err = toInt64(matchClause)
				}
			}
		}
	}
	return int32(med), err
}

func (r *Router) SetMed(ctx context.Context, req *pb.SetMedRequest) (*pb.SetMedResult, error) {
	oldMed, err := getMed(req)
	if err != nil {
		return nil, nil
	}
	var obj pb.SetMedResult
	err = execCommand(
		"route-map "+req.RouteMap+" "+req.Type+" "+string(req.SequenceNumber),
		"set metric "+string(req.Med),
	)
	currentMed, err := getMed(req)
	if err != nil {
		return nil, nil
	}
	obj.OldMed = oldMed
	obj.CurrentMed = currentMed
	return &obj, err
}
