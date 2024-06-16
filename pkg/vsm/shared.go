package vsm

import "github.com/prometheus/client_golang/prometheus"

type Parentable[T any] interface {
	Parent() T
	SetParent(T) error
}

type Initable interface {
	Init() error
}

type Vsm interface {
	Initable
	RegisterCollector(c prometheus.Collector) error
}
