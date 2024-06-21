package enzo

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/unhanded/enzo/pkg/enzo"
)

func NewWorkcenter(wCfg enzo.WorkcenterConfig) enzo.Workcenter {
	return &workCenter{
		id:                  wCfg.Id,
		label:               wCfg.Label,
		baselineProcessTime: wCfg.BaselineProcessTime,
		queue:               make([]enzo.WorkItem, 0),
	}
}

type workCenter struct {
	// Id of the work center
	id string
	// Label for the work center
	label string
	// Baseline process time for the work center
	// This is the time it takes to process a single unit of work
	// without any interruptions
	baselineProcessTime int64
	// Queue of work items to be processed
	queue []enzo.WorkItem // FIFO
	// Current work item being processed
	cell cell
	// Mesh peer to peer networking for direct routing
	parent enzo.MeshNetwork
}

func (wc *workCenter) Init() error {
	fmt.Printf("Workcenter %s initializing\n", wc.Id())

	wc.cell = cell{parent: wc}
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(int64(250)))
			wc.process()
		}
	}()
	wc.Parent().Parent().RegisterCollector(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Namespace: "enzo",
			Subsystem: fmt.Sprintf("workcenter_%s", wc.Id()),
			Name:      "queue",
			Help:      "Number of items in the workcenter queue",
		}, func() float64 {
			return float64(wc.Queued())
		}),
	)
	return nil
}

func (wc *workCenter) Id() string {
	return wc.id
}

func (wc *workCenter) Label() string {
	return wc.label
}

func (wc *workCenter) Busy() bool {
	return wc.cell.IsProcessing()
}

func (wc *workCenter) BaselineProcessTime() int64 {
	return wc.baselineProcessTime
}

func (wc *workCenter) Queue(workItem enzo.WorkItem) {
	wc.queue = append(wc.queue, workItem)
}

func (wc *workCenter) Queued() int {
	return len(wc.queue)
}

func (wc *workCenter) Recieve(workItem enzo.WorkItem) error {
	wc.Queue(workItem)
	return nil
}

func (wc *workCenter) process() {
	if wc.cell.ItemInsertable() && len(wc.queue) > 0 {
		wc.cell.Insert(wc.queue[0])
		wc.queue = wc.queue[1:]

	}
	if wc.cell.ItemExtractable() {
		wc.cell.Output()
	}
}

func (wc *workCenter) ApplyConfig(cfg enzo.WorkcenterConfig) error {
	wc.baselineProcessTime = cfg.BaselineProcessTime
	wc.label = cfg.Label
	return nil
}

func (wc *workCenter) Parent() enzo.MeshNetwork {
	return wc.parent
}

func (wc *workCenter) SetParent(p enzo.MeshNetwork) error {
	wc.parent = p
	return nil
}

func (wc *workCenter) DeInit() error {
	return fmt.Errorf("not implemented")
}

// A cell is a single processing unit within a workcenter, they are theoretically parallelizable
type cell struct {
	parent        *workCenter
	item          enzo.WorkItem
	lastStartTime int64
}

func (c cell) ItemExtractable() bool {
	return c.PassedFinishTime() && c.HasItem()
}

func (c cell) IsProcessing() bool {
	return c.lastStartTime+c.parent.BaselineProcessTime() > Now() && c.HasItem()
}

func (c cell) ItemInsertable() bool {
	return !c.HasItem()
}

func (c cell) PassedFinishTime() bool {
	return c.lastStartTime+c.parent.BaselineProcessTime() < Now()
}

func (c cell) HasItem() bool {
	return c.item != nil
}

func (c *cell) Insert(w enzo.WorkItem) {
	fmt.Printf("WORKCENTER: %s beginning work on item %s, ETA: %d ticks\n", c.parent.Id(), w.Id(), c.parent.BaselineProcessTime())
	c.lastStartTime = Now()
	c.item = w
}

func (c *cell) Output() error {
	fmt.Printf("WORKCENTER: %s transferring %s(workitem) over mesh\n", c.parent.Id(), c.item.Id())
	item := c.item
	c.item = nil
	c.lastStartTime = -1

	signErr := item.Route().Sign(c.parent.Id()) // Signs the step as completed by this workcenter
	if signErr != nil {
		return signErr
	}

	return c.parent.parent.Transfer(item)
}
