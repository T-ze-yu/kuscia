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

package handler

import (
	kusciaapisv1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
)

// RunningHandler will handle kuscia job in Running phase.
type RunningHandler struct {
	jobScheduler *JobScheduler
}

// NewRunningHandler return RunningHandler to handle Running kuscia job.
func NewRunningHandler(deps *Dependencies) *RunningHandler {
	return &RunningHandler{
		jobScheduler: NewJobScheduler(deps.KusciaClient, deps.NamespaceLister, deps.KusciaTaskLister),
	}
}

// HandlePhase implements the KusciaJobPhaseHandler interface.
// It will do scheduling when the job phase is Running.
func (h *RunningHandler) HandlePhase(kusciaJob *kusciaapisv1alpha1.KusciaJob) (needUpdate bool, retErr error) {
	return h.jobScheduler.push(kusciaJob)
}
