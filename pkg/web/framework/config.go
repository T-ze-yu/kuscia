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

package framework

import (
	"github.com/spf13/pflag"

	"github.com/secretflow/kuscia/pkg/web/errorcode"
)

type Config interface {
	// Validate options.
	Validate(errs *errorcode.Errs)
}

// ConfigLoader provides behavioral abstraction for different Config sources.
type ConfigLoader interface {
	// RegisterFlags registers flags.
	RegisterFlags(conf Config, name string, fs *pflag.FlagSet)
	// Process config obtained by Flag.
	Process(conf Config, name string, errs *errorcode.Errs)
}
