package enzo

import (
	"fmt"

	"github.com/unhanded/enzo-vsm/pkg/vsm"
)

func NewDynamicRoute(steps ...vsm.EnzoDynamicStep) vsm.EnzoDynamicRoute {
	return &dynRoute{
		steps: steps,
	}
}

type dynRoute struct {
	steps []vsm.EnzoDynamicStep
}

func (dr dynRoute) All() []vsm.EnzoDynamicStep {
	return dr.steps
}

func (dr *dynRoute) Current() (vsm.EnzoDynamicStep, error) {
	for _, s := range dr.steps {
		if !s.IsCompleted() {
			return s, nil
		}
	}
	return dr.steps[len(dr.steps)-1], nil
}

func (dr *dynRoute) IsFinished() bool {
	lastStep := dr.steps[len(dr.steps)-1]
	return lastStep.IsCompleted()
}

func (dr *dynRoute) Sign(workcenterId string) error {
	stp, stpErr := dr.findStepFromId(workcenterId)
	if stpErr != nil {
		return fmt.Errorf("error signing workitem: %s", stpErr.Error())
	}
	return stp.MarkAsComplete()
}

func (dr *dynRoute) findStepFromId(id string) (vsm.EnzoDynamicStep, error) {
	for _, step := range dr.steps {
		for _, idString := range step.Options() {
			if idString == id {
				return step, nil
			}
		}
	}
	return nil, fmt.Errorf("step not found")
}
