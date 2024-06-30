package core

import "github.com/unhanded/flownet/pkg/flownet"

type EnzoItem struct {
	Attr  flownet.Attributes `json:"attributes"`
	Nodes []string           `json:"nodes"`
}

func (it *EnzoItem) Attributes() flownet.Attributes {
	return nil
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
