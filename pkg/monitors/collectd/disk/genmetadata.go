// Code generated by monitor-code-gen. DO NOT EDIT.

package disk

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/disk"

var groupSet = map[string]bool{}

const (
	diskIoTimeIoTime         = "disk_io_time.io_time"
	diskIoTimeWeightedIoTime = "disk_io_time.weighted_io_time"
	diskMergedRead           = "disk_merged.read"
	diskMergedWrite          = "disk_merged.write"
	diskOctetsRead           = "disk_octets.read"
	diskOctetsWrite          = "disk_octets.write"
	diskOpsRead              = "disk_ops.read"
	diskOpsWrite             = "disk_ops.write"
	diskTimeRead             = "disk_time.read"
	diskTimeWrite            = "disk_time.write"
	pendingOperations        = "pending_operations"
)

var metricSet = map[string]monitors.MetricInfo{
	diskIoTimeIoTime:         {Type: datapoint.Counter},
	diskIoTimeWeightedIoTime: {Type: datapoint.Counter},
	diskMergedRead:           {Type: datapoint.Counter},
	diskMergedWrite:          {Type: datapoint.Counter},
	diskOctetsRead:           {Type: datapoint.Counter},
	diskOctetsWrite:          {Type: datapoint.Counter},
	diskOpsRead:              {Type: datapoint.Counter},
	diskOpsWrite:             {Type: datapoint.Counter},
	diskTimeRead:             {Type: datapoint.Counter},
	diskTimeWrite:            {Type: datapoint.Counter},
	pendingOperations:        {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	diskOpsRead:  true,
	diskOpsWrite: true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "collectd/disk",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
