// Copyright 2023 Ant Group Co., Ltd.
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

package node

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/secretflow/kuscia/pkg/agent/config"
)

func TestGenericNode_ConfigureNode(t *testing.T) {
	agentConfig := config.DefaultStaticAgentConfig()
	agentConfig.RootDir = t.TempDir()
	capacityManager, err := NewCapacityManager(&agentConfig.Capacity, ".", true)
	assert.NoError(t, err)
	dep := &GenericNodeDependence{
		BaseNodeDependence: BaseNodeDependence{
			Runtime:         config.ContainerRuntime,
			Namespace:       "test-namespace",
			Address:         "1.1.1.1",
			CapacityManager: capacityManager,
		},
		RootDir: agentConfig.RootDir,
	}

	n := NewGenericNodeProvider(dep)
	node := n.ConfigureNode(context.Background(), "test-name")
	assert.Equal(t, "test-name", node.Name)
	assert.Equal(t, n.runtime, node.Labels[labelRuntime])
	assert.Equal(t, 5, len(node.Status.Conditions))

}
