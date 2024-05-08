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

//nolint:dulp
package autonomy

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/secretflow/kuscia/cmd/kuscia/confloader"
	"github.com/secretflow/kuscia/cmd/kuscia/modules"
	"github.com/secretflow/kuscia/cmd/kuscia/utils"
	"github.com/secretflow/kuscia/pkg/common"
	"github.com/secretflow/kuscia/pkg/utils/nlog"
)

func NewAutonomyCommand(ctx context.Context) *cobra.Command {
	configFile := ""
	onlyControllers := false
	cmd := &cobra.Command{
		Use:          "autonomy",
		Short:        "Autonomy contains all modules",
		Long:         `Autonomy contains all modules, such as: k3s, controllers, scheduler, agent, domainroute, envoy and so on`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(ctx, configFile, onlyControllers)
		},
	}
	cmd.Flags().StringVarP(&configFile, "conf", "c", "etc/conf/kuscia.yaml", "config path")
	cmd.Flags().BoolVar(&onlyControllers, "controllers", false, "only run controllers and scheduler, will remove later")
	return cmd
}

func Run(ctx context.Context, configFile string, onlyControllers bool) error {
	kusciaConf := confloader.ReadConfig(configFile, common.RunModeAutonomy)
	conf := modules.InitDependencies(ctx, kusciaConf)
	defer conf.Close()

	if onlyControllers {
		conf.MakeClients()
		modules.RunOperatorsAllinOneWithDestroy(conf)

		utils.SetupPprof(conf.Debug, conf.CtrDebugPort, true)
		nlog.Info("Scheduler and controllers are all started")
		// wait any controller failed
	} else {
		coreDnsModule := modules.RunCoreDNSWithDestroy(conf)
		modules.RunK3sWithDestroy(conf)
		// make clients after k3s start
		conf.MakeClients()

		cdsModule, ok := coreDnsModule.(*modules.CorednsModule)
		if !ok {
			return errors.New("coredns module type is invalid")
		}
		cdsModule.StartControllers(ctx, conf.Clients.KubeClient)

		if err := modules.CreateDefaultDomain(ctx, conf); err != nil {
			nlog.Error(err)
			return err
		}

		if err := modules.CreateCrossNamespace(ctx, conf); err != nil {
			nlog.Error(err)
			return err
		}

		if conf.EnableContainerd {
			modules.RunContainerdWithDestroy(conf)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			modules.RunOperatorsInSubProcessWithDestroy(conf)
		}()
		go func() {
			defer wg.Done()
			modules.RunEnvoyWithDestroy(conf)
		}()
		wg.Wait()
		modules.RunKusciaAPIWithDestroy(conf)
		modules.RunAgentWithDestroy(conf)
		modules.RunConfManagerWithDestroy(conf)
		modules.RunDataMeshWithDestroy(conf)
		modules.RunTransportWithDestroy(conf)
		modules.RunNodeExporterWithDestroy(conf)
		modules.RunSsExporterWithDestroy(conf)
		modules.RunMetricExporterWithDestroy(conf)
		utils.SetupPprof(conf.Debug, conf.DebugPort, false)

		modules.SetKusciaOOMScore()
	}
	conf.WaitAllModulesDone(ctx.Done())
	return nil
}
