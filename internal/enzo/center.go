package enzo

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/unhanded/enzo-vsm/pkg/vsm"
)

func NewWorkcenter(wCfg vsm.WorkcenterConfig) vsm.EnzoWorkcenter {
	return &workCenter{
		id:                  wCfg.Id,
		label:               wCfg.Label,
		baselineProcessTime: wCfg.BaselineProcessTime,
		queue:               make([]vsm.WorkItem, 0),
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
	queue []vsm.WorkItem // FIFO
	// Current work item being processed
	cell cell
	// Mesh peer to peer networking for direct routing
	parent vsm.MeshNetwork
}

func (wc *workCenter) Init() {
	fmt.Printf("Workcenter %s initializing\n", wc.Id())

	wc.cell = cell{clock: wc.parent.Clock(), parent: wc, processTime: wc.baselineProcessTime}
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(int64(wc.parent.Clock().GetTickInterval()/2)-1))
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

func (wc *workCenter) Queue(workItem vsm.WorkItem) {
	wc.queue = append(wc.queue, workItem)
}

func (wc *workCenter) Queued() int {
	return len(wc.queue)
}

func (wc *workCenter) Recieve(workItem vsm.WorkItem) error {
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

func (wc *workCenter) Update(cfg vsm.WorkcenterConfig) error {
	wc.baselineProcessTime = cfg.BaselineProcessTime
	wc.label = cfg.Label
	return nil
}

func (wc *workCenter) Parent() vsm.MeshNetwork {
	return wc.parent
}

func (wc *workCenter) SetParent(p vsm.MeshNetwork) error {
	wc.parent = p
	return nil
}

type cell struct {
	parent        *workCenter
	item          vsm.WorkItem
	lastStartTime int64
	processTime   int64
	clock         vsm.EnzoClock
}

func (c cell) ItemExtractable() bool {
	return c.PassedFinishTime() && c.HasItem()
}

func (c cell) IsProcessing() bool {
	return c.lastStartTime+c.processTime > c.clock.Now() && c.HasItem()
}

func (c cell) ItemInsertable() bool {
	return !c.HasItem()
}

func (c cell) PassedFinishTime() bool {
	return c.lastStartTime+c.processTime < c.clock.Now()
}

func (c cell) HasItem() bool {
	return c.item != nil
}

func (c *cell) Insert(w vsm.WorkItem) {
	fmt.Printf("WORKCENTER: %s beginning work on item %s, ETA: %d ticks\n", c.parent.Id(), w.Id(), c.processTime)
	c.lastStartTime = c.clock.Now()
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
