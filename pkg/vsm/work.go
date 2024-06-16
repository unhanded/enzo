package vsm

type EnzoWorkcenter interface {
	Parentable[MeshNetwork]
	Observable

	Id() string
	Init()
	Recieve(data WorkItem) error
	Queue(WorkItem)
	Update(WorkcenterConfig) error
}

type WorkcenterConfig struct {
	Id                  string `yaml:"id"`
	Label               string `yaml:"label"`
	BaselineProcessTime int64  `yaml:"baseline_process_time"`
}

type WorkItem interface {
	RoutableItem

	Id() string
	Label() string
	Characteristic() []byte
	Destinations() ([]string, error)
	Route() EnzoDynamicRoute
}

type WorkItemConfig struct {
	Label          string      `json:"label"`
	Characteristic []byte      `json:"characteristic"`
	Route          RouteConfig `json:"route"`
}
