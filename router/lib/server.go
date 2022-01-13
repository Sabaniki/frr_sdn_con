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
	fmt.Println(commands)
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
		vtyshArgs[i] = "-c '" + vtyshArg + "'"
	}
	commands := []string{"docker", "exec", "nfv-kit_frr_1", "bash", "-c", "vtysh -c 'conf t' " + strings.Join(vtyshArgs, " ")}

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
			for _, setClause := range rule.SetClauses {
				if strings.Contains(setClause, "metric") {
					fmt.Println(setClause)
					med, err = toInt64(setClause)
				}
			}
		}
	}
	println("end get med: " + strconv.Itoa(int(med)))
	return int32(med), err
}

func (r *Router) SetMed(ctx context.Context, req *pb.SetMedRequest) (*pb.SetMedResult, error) {
	oldMed, err := getMed(req)
	if err != nil {
		println("error occurred!")
		return nil, err
	}
	var obj pb.SetMedResult
	err = execCommand(
		"route-map "+req.RouteMap+" "+req.Type+" "+strconv.Itoa(int(req.SequenceNumber)),
		"set metric "+strconv.Itoa(int(req.Med)),
	)
	if err != nil {
		println("error occurred!")
		return nil, err
	}
	currentMed, err := getMed(req)
	if err != nil {
		println("error occurred!")
		return nil, err
	}
	obj.OldMed = oldMed
	obj.CurrentMed = currentMed
	return &obj, err
}

func (r *Router) ShowOneInterface(ctx context.Context, req *pb.ShowOneInterfaceRequest) (*pb.ShowOneInterfaceResult, error) {
	var obj pb.ShowOneInterfaceResult
	err := execShowCommand(&obj, "interface", "."+req.GetName(), ".evpnMh = \"no\"")
	println("requested!: ")
	fmt.Println(&obj)
	return &obj, err
}

func (r *Router) ShowAllInterface(ctx context.Context, req *pb.ShowAllInterfaceRequest) (*pb.ShowAllInterfaceResult, error) {
	// TODO: impl
	return nil, nil
}

func (r *Router) ConfigInterface(ctx context.Context, req *pb.ConfigInterfaceRequest) (*pb.ConfigInterfaceResult, error) {
	// TODO: impl
	return nil, nil
}
