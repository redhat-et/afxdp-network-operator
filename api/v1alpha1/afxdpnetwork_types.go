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

// AfxdpNetworkSpec defines the desired state of AfxdpNetwork
// It is primarily used to generate a NetworkAttachmentDefinition
// CR with an AF_XDP CNI plugin configuration.
type AfxdpNetworkSpec struct {
	// Namespace of the NetworkAttachmentDefinition custom resource
	NetworkNamespace string `json:"networkNamespace,omitempty"`
	//Capabilities to be configured for this AfxdpNetwork network.
	//Capabilities supported: (mappinning|bpfman|udsserver), e.g. '{"mappinning": true}'
	Capabilities string `json:"capabilities,omitempty"`
	//IPAM configuration to be used for this network.
	IPAM string `json:"ipam,omitempty"`
	// PF/SF link state (enable|disable|auto)
	// +kubebuilder:validation:Enum={"auto","enable","disable"}
	LinkState string `json:"linkState,omitempty"`
	// LogLevel sets the log level of the AF_XDP CNI - either of panic, error, warning, info, debug. Defaults
	// to info if left blank.
	// +kubebuilder:validation:Enum={"panic", "error","warning","info","debug",""}
	// +kubebuilder:default:= "info"
	LogLevel string `json:"logLevel,omitempty"`
	// LogFile sets the log file of the AF_XDP CNI logs. If unset (default), this will log to stderr and thus
	// to multus and container runtime logs.
	LogFile string `json:"logFile,omitempty"`
	// Mode sets the mode of the  AF_XDP CNI.
	// +kubebuilder:validation:Enum={"primary", "cdq",""}
	// +kubebuilder:default:= "primary"
	Mode string `json:"mode,omitempty"`
	// EthtoolCmds specifies the ethtool cmds to be issued by the AF_XDP CNI on netdevs in this network.
	EthtoolCmds []string `json:"ethtoolCmds,omitempty"`
	// EthtoolCmds specifies whether or not to use a syncer with the AF_XDP DP.
	DPSyncer bool `json:"dpSyncer,omitempty"`
}

// AfxdpNetworkStatus defines the observed state of AfxdpNetwork
type AfxdpNetworkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AfxdpNetwork is the Schema for the afxdpnetworks API
type AfxdpNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AfxdpNetworkSpec   `json:"spec,omitempty"`
	Status AfxdpNetworkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AfxdpNetworkList contains a list of AfxdpNetwork
type AfxdpNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfxdpNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AfxdpNetwork{}, &AfxdpNetworkList{})
}
