package main

import (
	"log"
	"sync"
	"time"
)

type SwCounter struct {
	mut sync.Mutex
	buckets []*Bucket
	numBuckets int
	bucketSize time.Duration
	wndOpen time.Time
	wndClose time.Time
}

func NewSwCounter(numBuckets int, bucketSize time.Duration, start time.Time) *SwCounter {
	c := &SwCounter{
		buckets: make([]*Bucket, numBuckets),
		numBuckets: numBuckets,
		bucketSize: bucketSize,
		wndOpen: start,
		wndClose: start.Add(time.Duration(numBuckets)*bucketSize),
	}
	for i := 0; i < numBuckets; i++ {
		c.buckets[i] = NewBucket()
	}
	return c
}

// 疑问：怎么使用这个计数器？用最久的桶buckets[0]的统计数据？
func (c *SwCounter) FirstBucket() *Bucket {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.buckets[0]
}

func (c *SwCounter) IncrEvent(evt int, tm time.Time) {
	c.mut.Lock()
	defer c.mut.Unlock()

	if tm.Before(c.wndOpen){
		log.Printf("event stale\n")
		return
	}

	var idx int
	for {
		idx = c.bucketIndex(tm)
		if idx >= 0 {
			break
		}
		c.slideWindow()
	}

	c.buckets[idx].Incr(evt)
}

func (c *SwCounter) bucketIndex(tm time.Time) int {
	if tm.UnixNano() >= c.wndClose.UnixNano() {
		return -1
	}

	return int(tm.Sub(c.wndOpen)/c.bucketSize)
}

func (c *SwCounter) slideWindow() {
	c.buckets = c.buckets[1:]
	c.buckets = append(c.buckets, NewBucket())
	c.wndOpen = c.wndOpen.Add(c.bucketSize)
	c.wndClose = c.wndClose.Add(c.bucketSize)
}