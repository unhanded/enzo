package routing

type EnzoDynamicStep interface {
	Options() []string
	IsCompleted() bool
	MarkAsComplete() error
}
