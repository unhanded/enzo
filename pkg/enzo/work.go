package enzo

type WorkcenterConfig struct {
	Id                  string `yaml:"id"`
	Label               string `yaml:"label"`
	BaselineProcessTime int64  `yaml:"baseline_process_time"`
}

type Workcenter interface {
	Parentable[MeshNetwork]
	Observable
	LifetimeManagable

	Id() string
	Recieve(item WorkItem) error
	Queue(item WorkItem)
	ApplyConfig(WorkcenterConfig) error
}

type WorkItemConfig struct {
	Label          string      `json:"label"`
	Characteristic []byte      `json:"characteristic"`
	Route          RouteConfig `json:"route"`
}

type WorkItem interface {
	RoutableItem

	Id() string
	Label() string
	Characteristic() []byte
	Destinations() ([]string, error)
	Route() EnzoDynamicRoute
}
