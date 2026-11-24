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
	"testing"
)

func TestErrorConnectToServer_Error(t *testing.T) {
	type fields struct {
		Msg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "a1",
			fields: fields{Msg: "a1"},
			want:   "a1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorConnectToServer{
				Msg: tt.fields.Msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
