package enzo

import (
	"fmt"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/unhanded/enzo-vsm/pkg/vsm"
)

func NewMesh() vsm.MeshNetwork {
	return &mesh{
		clk:         NewClock(0, 100), // Default 10 TPS
		workcenters: map[string]vsm.EnzoWorkcenter{},
	}
}

type mesh struct {
	parent        vsm.Vsm
	clk           vsm.EnzoClock
	workcenters   map[string]vsm.EnzoWorkcenter
	finishedItems int
}

func (m *mesh) Init() error {
	fmt.Println("MESH: Initializing")

	if m.parent == nil {
		return fmt.Errorf("MESH: Parent not set (expected VSM)")
	}

	m.parent.RegisterCollector(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "mesh_nodes", Help: "Number of nodes in the mesh network"},
		func() float64 { return float64(len(m.workcenters)) },
	))

	m.parent.RegisterCollector(prometheus.NewCounterFunc(
		prometheus.CounterOpts{Name: "finished_items", Help: "Number of finished work items"},
		func() float64 { return float64(m.finishedItems) },
	))

	m.clk.Init()

	return nil
}

func (m *mesh) Clock() vsm.EnzoClock {
	return m.clk
}

func (m *mesh) Enroll(center vsm.EnzoWorkcenter) error {
	m.workcenters[center.Id()] = center
	m.workcenters[center.Id()].SetParent(m)
	m.workcenters[center.Id()].Init()

	return nil
}

func (m *mesh) Unenroll(id string) error {
	delete(m.workcenters, id)
	return nil
}

func (m *mesh) Transfer(item vsm.WorkItem) error {
	if item.Route().IsFinished() {
		fmt.Printf("MESH: Workitem %s leaving network (finished)\n", item.Id())
		m.finishedItems++
		return nil
	}
	destinations, dstErr := item.Destinations()
	if dstErr != nil {
		fmt.Printf("MESH: Workitem %s dropped (no destinations)\n", item.Id())
		return nil
	}

	attemptLimit := 10

	for i := 0; i < attemptLimit; i++ {
		sel := rand.Int63n(int64(len(destinations)))
		targetId := destinations[sel]
		if m.ValidDestination(targetId) {
			m.workcenters[targetId].Recieve(item)
			fmt.Printf("MESH: Transferring %s(workitem) to %s(workcenter)\n", item.Id(), targetId)
			return nil
		}
	}
	fmt.Printf("MESH: Workitem %s dropped (no valid destinations)\n", item.Id())

	return nil
}

func (m *mesh) ValidDestination(id string) bool {
	_, ok := m.workcenters[id]
	return ok
}

func (m *mesh) Nodes() []vsm.EnzoWorkcenter {
	nodes := make([]vsm.EnzoWorkcenter, 0)
	for _, node := range m.workcenters {
		nodes = append(nodes, node)
	}
	return nodes
}

func (m *mesh) Parent() vsm.Vsm {
	return m.parent
}

func (m *mesh) SetParent(p vsm.Vsm) error {
	m.parent = p
	return nil
}
