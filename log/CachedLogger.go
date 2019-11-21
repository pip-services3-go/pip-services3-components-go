package log

import (
	"sync"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

/*
Abstract logger that caches captured log messages in memory and periodically dumps them. Child classes implement saving cached messages to their specified destinations.

Configuration parameters
level: maximum log level to capture
source: source (context) name
options:
interval: interval in milliseconds to save log messages (default: 10 seconds)
max_cache_size: maximum number of messages stored in this cache (default: 100)
References
*:context-info:*:*:1.0 (optional) ContextInfo to detect the context id and specify counters source
*/
type ICachedLogSaver interface {
	Save(messages []*LogMessage) error
}

type CachedLogger struct {
	Logger
	cache        []*LogMessage
	updated      bool
	lastDumpTime time.Time
	maxCacheSize int
	interval     int
	lock         *sync.Mutex
	saver        ICachedLogSaver
}

// Creates a new instance of the logger from ICachedLogSaver
// Parameters:
// 			- saver ICachedLogSaver
// Returns CachedLogger
func InheritCachedLogger(saver ICachedLogSaver) *CachedLogger {
	c := &CachedLogger{
		cache:        []*LogMessage{},
		updated:      false,
		lastDumpTime: time.Now(),
		maxCacheSize: 100,
		interval:     10000,
		lock:         &sync.Mutex{},
		saver:        saver,
	}
	c.Logger = *InheritLogger(c)
	return c
}

// Writes a log message to the logger destination.
// Parameters:
// 		- level LogLevel
// 		a log level.
// 		- correlationId string
// 		transaction id to trace execution through call chain.
// 		- err error
// 		an error object associated with this message.
// 		- message string
// 		a human-readable message to log.
func (c *CachedLogger) Write(level int, correlationId string, err error, message string) {
	logMessage := &LogMessage{
		Time:          time.Now().UTC(),
		Level:         level,
		Source:        c.source,
		Message:       message,
		CorrelationId: correlationId,
	}

	if err != nil {
		errorDescription := errors.NewErrorDescription(err)
		logMessage.Error = *errorDescription
	}

	c.lock.Lock()
	c.cache = append(c.cache, logMessage)
	c.lock.Unlock()

	c.Update()
}

// Configures component by passing configuration parameters.
// Parameters:
// 		- config *config.ConfigParams
// 		configuration parameters to be set.
func (c *CachedLogger) Configure(cfg *config.ConfigParams) {
	c.Logger.Configure(cfg)

	c.interval = cfg.GetAsIntegerWithDefault("options.interval", c.interval)
	c.maxCacheSize = cfg.GetAsIntegerWithDefault("options.max_cache_size", c.maxCacheSize)
}

// Clears (removes) all cached log messages.
func (c *CachedLogger) Clear() {
	c.lock.Lock()
	c.cache = []*LogMessage{}
	c.updated = false
	c.lock.Unlock()
}

// Dumps (writes) the currently cached log messages.
func (c *CachedLogger) Dump() error {
	if c.updated {
		if !c.updated {
			return nil
		}

		var messages []*LogMessage
		c.lock.Lock()

		messages = c.cache
		c.cache = []*LogMessage{}

		c.lock.Unlock()

		err := c.saver.Save(messages)
		if err != nil {
			c.lock.Lock()

			// Put failed messages back to cache
			c.cache = append(messages, c.cache...)

			// Truncate cache to max size
			if len(c.cache) > c.maxCacheSize {
				c.cache = c.cache[len(c.cache)-c.maxCacheSize:]
			}

			c.lock.Unlock()
		}

		c.updated = false
		c.lastDumpTime = time.Now()
		return err
	}
	return nil
}

// Makes message cache as updated and dumps it when timeout expires.
func (c *CachedLogger) Update() {
	c.updated = true

	elapsed := int(time.Since(c.lastDumpTime).Seconds() * 1000)

	if elapsed > c.interval {
		// Todo: Decide what to do with the error
		c.Dump()
	}
}
