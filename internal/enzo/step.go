package enzo

type EnzoDynamicStep interface {
	Options() []string
	IsCompleted() bool
	MarkAsComplete() error
}

type dynamicStep struct {
	opts      []string
	completed bool
}

func (ds *dynamicStep) Options() []string {
	return ds.opts
}

func (ds *dynamicStep) IsCompleted() bool {
	return ds.completed
}

func (ds *dynamicStep) MarkAsComplete() error {
	ds.completed = true
	return nil
}

func NewDynamicStep(opts ...string) EnzoDynamicStep {
	return &dynamicStep{
		opts:      []string(opts),
		completed: false,
	}
}
