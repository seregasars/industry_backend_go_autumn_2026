package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	lim := Limiter{}
	lim.clock = clock
	lim.rate = ratePerSec
	lim.burst = burst
	lim.tokens = float64(burst)
	if clock != nil {
		lim.last = clock.Now()
	}
	return &lim
}
func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.clock == nil || l.burst <= 0{
		return false
	}
	if n <= 0 || n > l.burst{
		return false
	}
	curTime := l.clock.Now()
	if curTime.After(l.last) && l.rate > 0{
		tokenAdd := (curTime.Sub(l.last).Seconds()) * l.rate
		l.tokens += tokenAdd
		l.last = curTime
	}
	if l.tokens > float64(l.burst){
		l.tokens = float64(l.burst)
	}
	 if l.tokens < float64(n) {
   	return false
    }
	l.tokens -= float64(n)
	return true
}
