/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AfxdpOperatorConfigSpec defines the desired state of AfxdpOperatorConfig
type AfxdpOperatorConfigSpec struct {
	// NodeSelector selects the nodes to be configured
	ConfigDaemonNodeSelector map[string]string `json:"configDaemonNodeSelector,omitempty"`
	// Flag to control whether the network resource injector webhook shall be deployed
	EnableInjector *bool `json:"enableInjector,omitempty"`
	// Flag to control whether the operator admission controller webhook shall be deployed
	EnableOperatorWebhook *bool `json:"enableOperatorWebhook,omitempty"`
	// Flag to control the log verbose level of the operator. Set to '0' to show only the basic logs. And set to '2' to show all the available logs.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2
	LogLevel int `json:"logLevel,omitempty"`
	// Flag to disable nodes drain during debugging
	DisableDrain bool `json:"disableDrain,omitempty"`

	//KindCluster  bool    `json:"kindCluster"`
}

// AfxdpOperatorConfigStatus defines the observed state of AfxdpOperatorConfig
type AfxdpOperatorConfigStatus struct {
	// Show the runtime status of the network resource injector webhook
	Injector string `json:"injector,omitempty"`
	// Show the runtime status of the operator admission controller webhook
	OperatorWebhook string `json:"operatorWebhook,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AfxdpOperatorConfig is the Schema for the afxdpoperatorconfigs API
type AfxdpOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AfxdpOperatorConfigSpec   `json:"spec,omitempty"`
	Status AfxdpOperatorConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AfxdpOperatorConfigList contains a list of AfxdpOperatorConfig
type AfxdpOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfxdpOperatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AfxdpOperatorConfig{}, &AfxdpOperatorConfigList{})
}
