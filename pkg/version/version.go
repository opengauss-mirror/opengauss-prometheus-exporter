// Copyright (c) 2020 Huawei Technologies Co., Ltd.
// opengauss-prometheus-exporter is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package version

import (
	"fmt"
)

// Version is the current program version.
var version = "1.1.2"

var (
	// metadata is extra build time data
	metadata = ""
	// BuildTimestamp is the UTC date time when the program is compiled.
	buildTimestamp = "Unknown"
	// GitHash is the git commit hash when the program is compiled.
	gitCommit = "Unknown"
	// GitBranch is the active git branch when the program is compiled.
	gitTagInfo = "Unknown"
	// GoVersion is the Go compiler version used to compile this program.
	// goVersion = "Unknown"
)

// BuildInfo describes the compile time information.
type BuildInfo struct {
	// Version is the current semver.
	Version string `json:"version,omitempty"`
	// GitCommit is the git sha1.
	GitCommit string `json:"git_commit,omitempty"`
	// GitTreeState is the state of the git tree.
	GitTagInfo string `json:"git_tag_info,omitempty"`
	// GoVersion is the version of the Go compiler used.
	GoVersion string `json:"go_version,omitempty"`
}

// GetVersion returns the semver string of the version
func GetVersion() string {
	if metadata == "" {
		return version
	}
	return version + "+" + metadata
}

// Get returns build info
// func Get() BuildInfo {
// 	v := BuildInfo{
// 		Version:    GetVersion(),
// 		GitCommit:  gitCommit,
// 		GitTagInfo: gitTagInfo,
// 		GoVersion:  runtime.Version(),
// 	}
//
// 	// HACK(bacongobbler): strip out GoVersion during a test run for consistent test output
// 	if flag.Lookup("test.v") != nil {
// 		v.GoVersion = ""
// 	}
// 	return v
// }

// LongVersion returns the version information of this program as a string.
func GetLongVersion() string {
	return fmt.Sprintf(
		"Release version: %s\n"+
			"Git Commit hash: %s\n"+
			"Git Tag        : %s\n"+
			"Build timestamp: %sZ\n",
		GetVersion(),
		gitCommit,
		gitTagInfo,
		buildTimestamp,
	)
}
