/*
Copyright 2020 The Kubernetes Authors.

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

package manager

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &ControllerManagerConfig{}

// ControllerManagerConfig scaffolds the config file in config/manager folder.
type ControllerManagerConfig struct {
	machinery.TemplateMixin
	machinery.DomainMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements input.Template
func (f *ControllerManagerConfig) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("config", "manager", "controller_manager_config.yaml")
	}

	f.TemplateBody = controllerManagerConfigTemplate

	f.IfExistsAction = machinery.Error

	return nil
}

const controllerManagerConfigTemplate = `apiVersion: controller-runtime.sigs.k8s.io/v1alpha1
kind: ControllerManagerConfig
health:
  healthProbeBindAddress: :8081
metrics:
  bindAddress: 127.0.0.1:8080
webhook:
  port: 9443
leaderElection:
  leaderElect: true
  resourceName: {{ hashFNV .Repo }}.{{ .Domain }}
# leaderElectionReleaseOnCancel defines if the leader should step down volume
# when the Manager ends. This requires the binary to immediately end when the
# Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
# speeds up voluntary leader transitions as the new leader don't have to wait
# LeaseDuration time first.
# In the default scaffold provided, the program ends immediately after
# the manager stops, so would be fine to enable this option. However,
# if you are doing or is intended to do any operation such as perform cleanups
# after the manager stops then its usage might be unsafe.
# leaderElectionReleaseOnCancel: true
`
