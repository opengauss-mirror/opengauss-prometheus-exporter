// Copyright (c) 2020 Huawei Technologies Co., Ltd.
// opengauss-prometheus-exporter is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package exporter

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	DISCARD      = "DISCARD" // Ignore this column (when SELECT *)
	LABEL        = "LABEL"   // Use this column as a label
	COUNTER      = "COUNTER" // Use this column as a counter
	GAUGE        = "GAUGE"   // Use this column as a gauge
	HISTOGRAM    = "HISTOGRAM"
	MappedMETRIC = "MAPPEDMETRIC"
	DURATION     = "DURATION"
)

var ColumnUsage = map[string]bool{
	DISCARD:      true,
	LABEL:        true,
	COUNTER:      true,
	GAUGE:        true,
	HISTOGRAM:    true,
	MappedMETRIC: true,
	DURATION:     true,
}

type Column struct {
	CheckUTF8      bool                 `yaml:"checkUTF8"`
	DisCard        bool                 `yaml:"-"`
	Histogram      bool                 `yaml:"-"` // Should metric be treated as a histogram?
	Name           string               `yaml:"name"`
	Desc           string               `yaml:"description,omitempty"`
	Usage          string               `yaml:"usage,omitempty"`
	Rename         string               `yaml:"rename,omitempty"`
	PrometheusDesc *prometheus.Desc     `yaml:"-"`
	PrometheusType prometheus.ValueType `yaml:"-"`
}

func (c *Column) String() string {
	return fmt.Sprintf("%-8s %-30s %s", c.Usage, c.Name, c.Desc)
}
