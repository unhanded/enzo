package enzo

type RouteConfig struct {
	Steps []StepConfig `json:"steps"`
}

type EnzoDynamicRoute interface {
	All() []EnzoDynamicStep
	Current() (EnzoDynamicStep, error)
	Sign(workcenterId string) error
	IsFinished() bool
}

type StepConfig struct {
	Options []string `json:"options"`
}

type EnzoDynamicStep interface {
	Options() []string
	IsCompleted() bool
	MarkAsComplete() error
}
