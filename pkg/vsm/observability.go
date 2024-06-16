package vsm

type Observable interface {
	Busy() bool
	Queued() int
}
