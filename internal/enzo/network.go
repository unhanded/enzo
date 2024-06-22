package enzo

import (
	"github.com/unhanded/flownet/pkg/fnet"
	"github.com/unhanded/flownet/pkg/ifnet"
)

func NewNetwork() ifnet.FNet {
	return fnet.New()
}

type NetNode struct {
	NodeName  string  `json:"name" yaml:"name"`
	NodeId    string  `json:"id" yaml:"id"`
	NodeValue float64 `json:"value" yaml:"value"`
}

func (n *NetNode) GetTimeoutDuration(r ifnet.Route) float64 {
	return n.NodeValue
}

func (n *NetNode) Id() string {
	return n.NodeId
}

func (n *NetNode) Name() string {
	return n.NodeName
}
