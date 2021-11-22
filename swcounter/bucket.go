package main

import (
	"log"
	"sync/atomic"
)

const (
	EvtSuccess = 1
	EvtError   = 2
	EvtTimeout = 3
	EvtCount   = EvtTimeout
)

type Bucket struct {
	counts [EvtCount +1]int32
}

func NewBucket() *Bucket {
	return &Bucket{}
}

func (b *Bucket) Incr(evt int) {
	if evt > 0 && evt <= EvtCount {
		atomic.AddInt32(&b.counts[evt], 1)
	} else {
		log.Printf("invalid event %d\n", evt)
	}
}

func (b *Bucket) SuccessCount() int32 {
	return atomic.LoadInt32(&b.counts[EvtSuccess])
}

func (b *Bucket) ErrorCount() int32 {
	return atomic.LoadInt32(&b.counts[EvtError])
}

func (b *Bucket) TimeoutCount() int32 {
	return atomic.LoadInt32(&b.counts[EvtTimeout])
}

func (b *Bucket) SuccessRate() float64 {
	var totalCount int32
	for i := 1; i <=EvtCount; i++ {
		totalCount += atomic.LoadInt32(&b.counts[i])
	}
	if totalCount == 0 {
		return 0
	}
	return float64(atomic.LoadInt32(&b.counts[EvtSuccess]))/float64(totalCount)
}