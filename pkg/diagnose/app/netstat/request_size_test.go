// Copyright 2024 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netstat

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/secretflow/kuscia/pkg/diagnose/app/client"
	"github.com/secretflow/kuscia/pkg/diagnose/common"
	"github.com/secretflow/kuscia/proto/api/v1alpha1/diagnose"
	"golang.org/x/net/context"
	"gotest.tools/v3/assert"
)

func TestRequestSizeTask(t *testing.T) {
	tests := []struct {
		name   string
		size   int
		result string
	}{
		{"request_size_1MB_Pass", 1, common.Pass},
		{"request_size_3MB_Fail", 3, common.Fail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client.NewDiagnoseClient("")
			threshold := 2
			task := NewReqSizeTask(cli, threshold).(*ReqSizeTask)
			patches := gomonkey.ApplyMethod(reflect.TypeOf(cli), "Mock", func(_ *client.Client, ctx context.Context, request *diagnose.MockRequest) (response *diagnose.MockResponse, err error) {
				if tt.size > threshold {
					return nil, fmt.Errorf("server timeout")
				}
				return nil, nil
			})
			defer patches.Reset()
			task.Run(context.Background())
			assert.Equal(t, task.output.Result, tt.result)
		})
	}
}
