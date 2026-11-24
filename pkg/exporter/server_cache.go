// Copyright (c) 2021 Huawei Technologies Co., Ltd.
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
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type cachedMetrics struct {
	metrics        []prometheus.Metric
	lastScrape     time.Time
	nonFatalErrors []error
	err            error
	name           string
	collect        bool
}

// IsValid true is cache valid
func (c *cachedMetrics) IsValid(ttl float64) bool {
	if ttl == 0 {
		return false
	}
	return !(time.Now().Sub(c.lastScrape).Seconds() >= ttl)
}

func (c *cachedMetrics) IsCollect() bool {
	return c.collect
}
