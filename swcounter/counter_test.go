package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCounterNormal(t *testing.T) {
	start := time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local)
	c := NewSwCounter(10, time.Second, start)
	c.IncrEvent(evtSuccess, start)
	c.IncrEvent(evtSuccess, start.Add(100*time.Millisecond))
	c.IncrEvent(evtError, start.Add(200*time.Millisecond))
	c.IncrEvent(evtTimeout, start.Add(300*time.Millisecond))

	b := c.FirstBucket()
	assert.Equal(t, 2, b.SuccessCount())
	assert.Equal(t, 1, b.ErrorCount())
	assert.Equal(t, 1, b.TimeoutCount())
	assert.Equal(t, 0.5, b.SuccessRate())
}

func TestCounterSlide(t *testing.T) {
	start := time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local)
	c := NewSwCounter(10, time.Second, start)
	c.IncrEvent(evtSuccess, start)
	c.IncrEvent(evtSuccess, start.Add(100*time.Millisecond))
	c.IncrEvent(evtError, start.Add(200*time.Millisecond))
	c.IncrEvent(evtTimeout, start.Add(300*time.Millisecond))

	c.IncrEvent(evtSuccess, start.Add(time.Second))

	c.IncrEvent(evtError, start.Add(10*time.Second))

	b := c.FirstBucket()
	assert.Equal(t, 1, b.SuccessCount())
	assert.Equal(t, 1.0, b.SuccessRate())
}