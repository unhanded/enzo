package enzo

import (
	"fmt"

	"github.com/unhanded/enzo-vsm/pkg/enzo"
)

func NewWorkItem(id string, label string, characteristics []byte, route enzo.EnzoDynamicRoute) enzo.WorkItem {
	return &workItem{
		id:              id,
		label:           label,
		characteristics: characteristics,
		route:           route,
	}
}

type workItem struct {
	id              string
	label           string
	characteristics []byte
	route           enzo.EnzoDynamicRoute
}

func (wi *workItem) Id() string {
	return wi.id
}

func (wi *workItem) Label() string {
	return wi.label
}

func (wi *workItem) Characteristic() []byte {
	return wi.characteristics
}

func (wi workItem) Route() enzo.EnzoDynamicRoute {
	return wi.route
}

func (wi workItem) Destinations() ([]string, error) {
	dst, dstErr := wi.route.Current()
	if dstErr != nil {
		return []string{}, fmt.Errorf("no destination")
	}
	return dst.Options(), nil
}
