package enzo

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/unhanded/enzo-vsm/pkg/vsm"
	"gopkg.in/yaml.v3"
)

type Vsm struct {
	// Fields
	Network vsm.MeshNetwork
	Prm     *prometheus.Registry
}

func (v *Vsm) Init() error {
	v.Network.SetParent(v)
	v.Network.Init()
	return nil
}

func (v *Vsm) RegisterCollector(c prometheus.Collector) error {
	return v.Prm.Register(c)
}

// Apply is a function that applies the configuration to the runtime
// A Kubernetes-like interaction style
func (v *Vsm) Apply(data []byte) (string, error) {
	cfg := vsm.WorkcenterConfig{}
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return "", err
	}
	if !existsInNetwork(v.Network, cfg.Id) {
		ctr := NewWorkcenter(cfg)
		ctr.SetParent(v.Network)
		v.Network.Enroll(ctr)
		return fmt.Sprintf("workcenter/%s created", cfg.Id), nil
	} else {
		ctr := findInNetwork(v.Network, cfg.Id)
		if ctr != nil {
			ctr.Update(cfg)
			return fmt.Sprintf("updated workcenter/%s", cfg.Id), nil
		}
	}
	return "", nil
}

func (v *Vsm) Submit(data []byte) error {
	cfg := vsm.WorkItemConfig{}
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}
	id, idErr := uuid.NewV7()
	if idErr != nil {
		return idErr
	}
	steps := []vsm.EnzoDynamicStep{}
	for _, s := range cfg.Route.Steps {
		stp := NewDynamicStep(s.Options...)
		steps = append(steps, stp)
	}

	route := NewDynamicRoute(steps...)
	item := NewWorkItem(id.String(), cfg.Label, cfg.Characteristic, route)
	v.Network.Transfer(item)

	return nil
}

func findInNetwork(n vsm.MeshNetwork, id string) vsm.EnzoWorkcenter {
	for _, node := range n.Nodes() {
		if node.Id() == id {
			return node
		}
	}
	return nil
}

func existsInNetwork(n vsm.MeshNetwork, id string) bool {
	for _, node := range n.Nodes() {
		if node.Id() == id {
			return true
		}
	}
	return false
}
