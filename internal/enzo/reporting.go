package enzo

import "github.com/unhanded/flownet/pkg/flownet"

type NodeCollection struct {
	Nodes []flownet.Node[AuxData] `json:"nodes"`
}
