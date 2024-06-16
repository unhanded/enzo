package clock

type EnzoClock interface {
	Init()
	Pause()
	Unpause()
	Now() int64
	GetTickInterval() int64
	SetTickInterval(tickIntervalMs int64) error
}
