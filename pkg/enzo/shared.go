package enzo

import "github.com/prometheus/client_golang/prometheus"

type Configurable[T any] interface {
	ApplyConfig(cfg T) error
	GetConfig() T
}

type LifetimeManagable interface {
	Init() error
	DeInit() error
}

type Observable interface {
	Busy() bool
	Queued() int
	Label() string
}

type Parentable[T any] interface {
	Parent() T
	SetParent(T) error
}

type Vsm interface {
	LifetimeManagable
	RegisterCollector(c prometheus.Collector) error
}
