package info

import (
	"os"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type ContextInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ContextId   string            `json:"context_id"`
	StartTime   time.Time         `json:"start_time"`
	Properties  map[string]string `json:"properties"`
}

func NewContextInfo() *ContextInfo {
	c := &ContextInfo{
		Name:       "unknown",
		StartTime:  time.Now(),
		Properties: map[string]string{},
	}
	c.ContextId, _ = os.Hostname()
	return c
}

func (c *ContextInfo) Uptime() int64 {
	return time.Now().Unix() - c.StartTime.Unix()
}

func (c *ContextInfo) Configure(cfg *config.ConfigParams) {
	c.Name = cfg.GetAsStringWithDefault("name", c.Name)
	c.Name = cfg.GetAsStringWithDefault("info.name", c.Name)

	c.Description = cfg.GetAsStringWithDefault("description", c.Description)
	c.Description = cfg.GetAsStringWithDefault("info.description", c.Description)

	c.Properties = cfg.GetSection("properties").InnerValue().(map[string]string)
}

func NewContextInfoFromConfig(cfg *config.ConfigParams) *ContextInfo {
	result := NewContextInfo()
	result.Configure(cfg)
	return result
}
