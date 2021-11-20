package main

import (
	"log"
	"sync"
)

const (
	evtSuccess = 1
	evtError = 2
	evtTimeout = 3
	evtCount = evtTimeout
)

type Bucket struct {
	mut sync.RWMutex
	counts [evtCount+1]int
}

func NewBucket() *Bucket {
	return &Bucket{}
}

func (b *Bucket) Incr(evt int) {
	b.mut.Lock()
	defer b.mut.Unlock()

	if evt > 0 && evt <= evtCount {
		b.counts[evt]++
	} else {
		log.Printf("invalid event %d\n", evt)
	}
}

func (b *Bucket) SuccessCount() int {
	b.mut.RLock()
	defer b.mut.RUnlock()

	return b.counts[evtSuccess]
}

func (b *Bucket) ErrorCount() int {
	b.mut.RLock()
	defer b.mut.RUnlock()

	return b.counts[evtError]
}

func (b *Bucket) TimeoutCount() int {
	b.mut.RLock()
	defer b.mut.RUnlock()

	return b.counts[evtTimeout]
}

func (b *Bucket) SuccessRate() float64 {
	b.mut.RLock()
	defer b.mut.RUnlock()

	var totalCount int
	for _, count := range b.counts {
		totalCount += count
	}
	if totalCount == 0 {
		return 0
	}
	return float64(b.counts[evtSuccess])/float64(totalCount)
}