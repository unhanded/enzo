package core

import "github.com/unhanded/flownet/pkg/flownet"

type EnzoItem struct {
	Attr  flownet.Attributes `json:"attributes"`
	Nodes []string           `json:"nodes"`
}

func (it *EnzoItem) Attributes() flownet.Attributes {
	if it.Attr == nil {
		return flownet.Attributes{}
	} else {
		return it.Attr
	}
}

func (it *EnzoItem) NodeIds() []string {
	return it.Nodes
}

func (it *EnzoItem) Validate() bool {
	if len(it.Nodes) < 1 {
		return false
	}
	return true
}
