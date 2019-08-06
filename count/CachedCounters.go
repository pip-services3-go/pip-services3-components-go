package count

import (
	"math"
	"sync"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type CachedCounters struct {
	cache         map[string]*Counter
	updated       bool
	lastDumpTime  time.Time
	lastResetTime time.Time
	mux           sync.Mutex
	interval      int64
	resetTimeout  int64
	saver         ICountersSaver
}

type ICountersSaver interface {
	Save(counters []*Counter) error
}

func InheritCacheCounters(saver ICountersSaver) *CachedCounters {
	return &CachedCounters{
		cache:         map[string]*Counter{},
		updated:       false,
		lastDumpTime:  time.Now(),
		lastResetTime: time.Now(),
		interval:      300000,
		resetTimeout:  0,
		saver:         saver,
	}
}

func (c *CachedCounters) Configure(config *config.ConfigParams) {
	c.interval = config.GetAsLongWithDefault("interval", c.interval)
	c.resetTimeout = config.GetAsLongWithDefault("reset_timeout", c.resetTimeout)
}

func (c *CachedCounters) Clear(name string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.cache, name)
}

func (c *CachedCounters) ClearAll() {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache = map[string]*Counter{}
}

func (c *CachedCounters) Dump() error {
	if !c.updated {
		return nil
	}

	counters := c.GetAll()
	err := c.saver.Save(counters)
	if err != nil {
		return err
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.updated = false
	c.lastDumpTime = time.Now()

	return nil
}

func (c *CachedCounters) update() error {
	c.updated = true
	newDumpTime := c.lastDumpTime.Add(time.Duration(c.interval) * time.Millisecond)
	if time.Now().After(newDumpTime) {
		return c.Dump()
	}
	return nil
}

func (c *CachedCounters) resetIfNeeded() {
	if c.resetTimeout == 0 {
		return
	}

	newResetTime := c.lastResetTime.Add(time.Duration(c.resetTimeout) * time.Millisecond)
	if time.Now().After(newResetTime) {
		c.cache = map[string]*Counter{}
		c.updated = false
		c.lastDumpTime = time.Now()
	}
}

func (c *CachedCounters) GetAll() []*Counter {
	c.mux.Lock()
	defer c.mux.Unlock()

	result := []*Counter{}
	for _, v := range c.cache {
		result = append(result, v)
	}

	return result
}

func (c *CachedCounters) Get(name string, typ int) *Counter {
	if name == "" {
		panic("Counter name cannot be null")
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.resetIfNeeded()

	counter, ok := c.cache[name]
	if !ok || counter.Type != typ {
		counter = NewCounter(name, typ)
		c.cache[name] = counter
	}

	return counter
}

func (c *CachedCounters) calculateStats(counter *Counter, value float32) {
	if counter == nil {
		panic("Counter cannot be nil")
	}

	counter.Last = value
	counter.Count++
	counter.Max = float32(math.Max(float64(counter.Max), float64(value)))
	counter.Min = float32(math.Min(float64(counter.Min), float64(value)))
	counter.Average = ((counter.Average * float32(counter.Count-1)) + value) / float32(counter.Count)
}

func (c *CachedCounters) BeginTiming(name string) *Timing {
	return NewTiming(name, c)
}

func (c *CachedCounters) EndTiming(name string, elapsed float32) {
	counter := c.Get(name, Interval)
	c.calculateStats(counter, elapsed)
	c.update()
}

func (c *CachedCounters) Stats(name string, value float32) {
	counter := c.Get(name, Statistics)
	c.calculateStats(counter, value)
	c.update()
}

func (c *CachedCounters) Last(name string, value float32) {
	counter := c.Get(name, LastValue)
	counter.Last = value
	c.update()
}

func (c *CachedCounters) TimestampNow(name string) {
	c.Timestamp(name, time.Now())
}

func (c *CachedCounters) Timestamp(name string, value time.Time) {
	counter := c.Get(name, Timestamp)
	counter.Time = value
	c.update()
}

func (c *CachedCounters) IncrementOne(name string) {
	c.Increment(name, 1)
}

func (c *CachedCounters) Increment(name string, value int) {
	counter := c.Get(name, Increment)
	counter.Count = counter.Count + value
	c.update()
}
