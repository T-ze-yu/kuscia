// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lite

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestConfigFile(t *testing.T) string {
	root := t.TempDir()
	path := filepath.Join(root, "config.yaml")
	content := fmt.Sprintf(`
rootDir: %s
domainID: alice
caKeyFile: etc/certs/ca.key
caFile: etc/certs/ca.crt
domainKeyFile: etc/certs/domain.key
agent:
  node:
    enableNodeReuse: true
  capacity:
    cpu: 4
    memory: 4Gi
  provider:
    runtime: runk
    k8s:
      kubeconfigFile: /home/kuscia/etc/certs/k3s.yaml
      namespace: default
  allowPrivileged: true
  plugins:
    - name: config-render
`, root)
	assert.NoError(t, os.WriteFile(path, []byte(content), 0644))

	transportDir := filepath.Join(root, "etc/conf/transport/")
	assert.NoError(t, os.MkdirAll(transportDir, 0755))
	transportFile := filepath.Join(transportDir, "transport.yaml")
	transportFileContent := `
msqConfig:
  SessionExpireSeconds: 600
  DeadSessionIDExpireSeconds: 1800
  TotalByteSizeLimit: 17179869184 # 16GB
  PerSessionByteSizeLimit: 62914560 # 60MB
httpConfig:
  port: 8081
  ReadTimeout: 300 # seconds
  WriteTimeout: 300 # seconds
  IdleTimeout: 60 # seconds
  ReqBodyMaxSize: 134217728 # 128MB
`
	assert.NoError(t, os.WriteFile(transportFile, []byte(transportFileContent), 0644))

	return path
}

func TestGetInitConfig(t *testing.T) {
	configFile := newTestConfigFile(t)

	d := getInitConfig(configFile, "alice")
	assert.Equal(t, "runk", d.Agent.Provider.Runtime)
}
