/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/oam-dev/kubevela/pkg/controller/v1alpha1/metrics"
	"github.com/oam-dev/kubevela/pkg/controller/v1alpha1/podspecworkload"
	"github.com/oam-dev/kubevela/pkg/controller/v1alpha1/routes"
)

// Setup workload controllers.
func Setup(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		metrics.Setup, podspecworkload.Setup, routes.Setup,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
