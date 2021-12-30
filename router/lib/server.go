package router

import (
	"context"

	pb "github.com/Sabaniki/frr_sdn_con/pb/api"
)

type Router struct{}

type BgpIpv6Summary struct {
	RouterId        string              `json:"routerId,omitempty"`
	As              int32               `json:"as,omitempty"`
	VrfId           int32               `json:"vrfId,omitempty"`
	VrfName         string              `json:"vrfName,omitempty"`
	TableVersion    int32               `json:"tableVersion,omitempty"`
	RibCount        int32               `json:"ribCount,omitempty"`
	RibMemory       int32               `json:"ribMemory,omitempty"`
	PeerCount       int32               `json:"peerCount,omitempty"`
	PeerMemory      int32               `json:"peerMemory,omitempty"`
	PeerGroupCount  int32               `json:"peerGroupCount,omitempty"`
	PeerGroupMemory int32               `json:"peerGroupMemory,omitempty"`
	Peers           map[string]PeerInfo `json:"peers,omitempty`
}

type PeerInfo struct {
	RemoteAs                   int32  `json:"remoteAs,omitempty"`
	LocalAs                    int32  `json:"localAs,omitempty"`
	Version                    int32  `json:"version,omitempty"`
	MsgRcvd                    int32  `json:"msgRcvd,omitempty"`
	MsgSent                    int32  `json:"msgSent,omitempty"`
	TableVersion               int32  `json:"tableVersion,omitempty"`
	Outq                       int32  `json:"outq,omitempty"`
	Inq                        int32  `json:"inq,omitempty"`
	PeerUptime                 string `json:"peerUptime,omitempty"`
	PeerUptimeMsec             uint32 `json:"peerUptimeMsec,omitempty"` // TODO: proto の方も int32 → uint32 (or int64) にしないとアカン
	PeerUptimeEstablishedEpoch uint32 `json:"peerUptimeEstablishedEpoch,omitempty"`
	PfxRcd                     int32  `json:"pfxRcd,omitempty"`
	PfxSnt                     int32  `json:"pfxSnt,omitempty"`
	State                      string `json:"state,omitempty"`
	PeerState                  string `json:"peerState,omitempty"`
	ConnectionsEstablished     int32  `json:"connectionsEstablished,omitempty"`
	ConnectionsDropped         int32  `json:"connectionsDropped,omitempty"`
	Desc                       string `json:"desc,omitempty"`
	IdType                     string `json:"idType,omitempty"`
}

func (r *Router) ShowBgpIpv6Summary(ctx context.Context, req *pb.ShowBgpIpv6SummaryRequest) (*pb.ShowBgpIpv6SummaryResult, error) {
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
