package routing

type EnzoDynamicRoute interface {
	All() []EnzoDynamicStep
	Current() (EnzoDynamicStep, error)
	Sign(workcenterId string) error
	IsFinished() bool
}
