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
	"strings"
)

// Opt ExporterOpt configures Exporter
type Opt func(*Exporter)

// WithDNS add config dsn to Exporter
func WithDNS(dsn []string) Opt {
	return func(e *Exporter) {
		e.dsn = dsn
	}
}

// WithConfig add config path to Exporter
func WithConfig(configPath string) Opt {
	return func(e *Exporter) {
		e.configPath = configPath
	}
}

// WithConstLabels add const label to exporter. 0 length label returns nil
func WithConstLabels(s string) Opt {
	return func(e *Exporter) {
		e.constantLabels = parseConstLabels(s)
	}
}

// WithCacheDisabled set cache param to exporter
func WithCacheDisabled(disableCache bool) Opt {
	return func(e *Exporter) {
		e.disableCache = disableCache
	}
}

// WithDisableSettingsMetrics set cache param to exporter
func WithDisableSettingsMetrics(b bool) Opt {
	return func(e *Exporter) {
		e.disableSettingsMetrics = b
	}
}

// WithFailFast marks exporter fail instead of waiting during start-up
func WithFailFast(failFast bool) Opt {
	return func(e *Exporter) {
		e.failFast = failFast
	}
}

// WithNamespace will specify metric namespace, by default is pg or pgbouncer
func WithNamespace(namespace string) Opt {
	return func(e *Exporter) {
		e.namespace = namespace
	}
}

// WithTags will register given tags to Exporter and all belonged servers
func WithTags(tags string) Opt {
	return func(e *Exporter) {
		e.tags = parseCSV(tags)
	}
}

// WithTimeToString WithTags will register given tags to Exporter and all belonged servers
func WithTimeToString(b bool) Opt {
	return func(e *Exporter) {
		e.timeToString = b
	}
}
func WithParallel(i int) Opt {
	return func(e *Exporter) {
		e.parallel = i
	}
}

// WithAutoDiscovery configures exporter with excluded database
func WithAutoDiscovery(flag bool) Opt {
	return func(e *Exporter) {
		e.autoDiscovery = flag
	}
}

// WithExcludeDatabases configures exporter with excluded database
func WithExcludeDatabases(excludeStr string) Opt {
	return func(e *Exporter) {
		if excludeStr == "" {
			return
		}
		e.excludedDatabases = strings.Split(excludeStr, ",")
	}
}

// WithIncludeDatabases configures exporter with excluded database
func WithIncludeDatabases(includeStr string) Opt {
	return func(e *Exporter) {
		if includeStr == "" {
			return
		}
		e.includeDatabases = strings.Split(includeStr, ",")
	}
}

type autoDiscoverOption struct {
	autoDiscovery     bool     // discovery other database on primary server
	excludedDatabases []string // excluded database for auto discovery
	includeDatabases  []string // include database for auto discovery
}

type metricMap struct {
	allMetricMap map[string]*QueryInstance // 全部采集指标 不判断Public为true
	priMetricMap map[string]*QueryInstance // 私有采集指标 autoDiscover下公用指标,只采集一次
}
