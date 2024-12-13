package core

import (
	"github.com/unhanded/flownet/pkg/core"
	"github.com/unhanded/flownet/pkg/flownet"
)

func NewNetwork() flownet.FNet[AuxData] {
	return core.New[AuxData]()
}

type AuxData struct {
	coordinateX float64
	coordinateY float64
}

type NetNode struct {
	NodeName  string  `json:"name"`
	NodeId    string  `json:"id"`
	NodeValue float64 `json:"value"`
	AuxData   AuxData
}

func (n *NetNode) GetResistance(r flownet.Probe) float64 {
	return n.NodeValue
}

func (n *NetNode) Id() string {
	return n.NodeId
}

func (n *NetNode) Name() string {
	return n.NodeName
}

func (n *NetNode) Data() AuxData {
	return n.AuxData
}
