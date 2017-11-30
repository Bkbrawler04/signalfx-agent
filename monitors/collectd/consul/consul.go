package consul

//go:generate collectd-template-to-go consul.tmpl

import (
	"github.com/signalfx/neo-agent/core/config"
	"github.com/signalfx/neo-agent/core/services"
	"github.com/signalfx/neo-agent/monitors"
	"github.com/signalfx/neo-agent/monitors/collectd"
)

const monitorType = "collectd/consul"

func init() {
	monitors.Register(monitorType, func() interface{} {
		return &Monitor{
			*collectd.NewServiceMonitorCore(CollectdTemplate),
		}
	}, &Config{})
}

// Config is the monitor-specific config with the generic config embedded
type Config struct {
	config.MonitorConfig
	ACLToken          *string                 `yaml:"aclToken"`
	UseHTTPS          bool                    `yaml:"useHTTPS" default:"false"`
	EnhancedMetrics   bool                    `yaml:"enhancedMetrics" default:"false"`
	CACertificate     *string                 `yaml:"caCertificate"`
	ClientCertificate *string                 `yaml:"clientCertificate"`
	ClientKey         *string                 `yaml:"clientKey"`
	ServiceEndpoints  []services.EndpointCore `yaml:"serviceEndpoints" default:"[]"`
}

// Monitor is the main type that represents the monitor
type Monitor struct {
	collectd.ServiceMonitorCore
}

// Configure configures and runs the plugin in collectd
func (am *Monitor) Configure(conf *Config) bool {
	return am.SetConfigurationAndRun(conf)
}