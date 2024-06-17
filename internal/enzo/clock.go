package enzo

import (
	"time"

	"github.com/unhanded/enzo-vsm/pkg/enzo"
)

func NewClock(initalTime int64, intervalMillisecond int64) enzo.EnzoClock {
	return &clock{t: initalTime, tickIntervalMs: intervalMillisecond}
}

type clock struct {
	t              int64
	tickIntervalMs int64
	paused         bool
}

func (c *clock) Init() {
	c.paused = false
	go c.continuousTick()
}

func (c *clock) Pause() {
	c.paused = true
}

func (c *clock) Unpause() {
	c.paused = false
}

func (c *clock) Now() int64 {
	return c.t
}

func (c *clock) GetTickInterval() int64 {
	return c.tickIntervalMs
}

func (c *clock) SetTickInterval(tickInterval int64) error {
	c.tickIntervalMs = int64(tickInterval)
	return nil
}

func (c *clock) continuousTick() {
	for {
		time.Sleep(time.Duration(c.tickIntervalMs) * time.Millisecond)
		if !c.paused {
			c.t++
		}
	}
}
