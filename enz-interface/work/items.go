package work

import (
	"github.com/unhanded/enzo-vsm/enz-interface/networking"
	"github.com/unhanded/enzo-vsm/enz-interface/routing"
)

type EnzoWorkItem interface {
	networking.RoutableItem

	Id() string
	Label() string
	Characteristic() []byte
	Route() routing.EnzoDynamicRoute
}
