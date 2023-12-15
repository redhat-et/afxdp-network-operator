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

// AfxdpNetworkNodePolicySpec defines the desired state of AfxdpNetworkNodePolicy
type AfxdpNetworkNodePolicySpec struct {
	// NodeSelector selects the nodes to be configured
	NodeSelector map[string]string `json:"nodeSelector"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=99
	// Priority of the policy, higher priority policies can override lower ones.
	Priority int `json:"priority,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// MTU of PF/SF
	Mtu int `json:"mtu,omitempty"`
	// NumSfs int `json:"numSfs"`
	// Exclude device's NUMA node when advertising this resource by af_xdp device plugin. Default to false.
	ExcludeTopology bool                `json:"excludeTopology,omitempty"`
	Pools           []*AfxdpNetworkPool `json:"pools"`
}

type AfxdpNetworkPoolDevice struct {
	// Name is the netdev name
	Name string `json:"name"`
	// PciAddress is the netdev PCI address
	PciAddress string `json:"pciAddress"`
	// MacAddress is the netdev MAC address
	MacAddress string `json:"macAddress"`
	// Sfs is the number of subfunctions configured on this netdev
	Sfs int `json:"sfs"`
}

type AfxdpNetworkPoolDriver struct {
	// Name is the driver name
	Name string `json:"name"`
	// Sfs is the number of subfunctions configured on this netdev
	Sfs int `json:"sfs"`
	//  Any primary device identified in the ExcludeDevices array will **not** be added to the pool.
	ExcludeDevices []*AfxdpNetworkPoolDevice `json:"excludeDevices"`
	// if ExcludeAddressed == true, does **not** add any device with an IPv4 address to the pool.
	// +kubebuilder:default:= false
	ExcludeAddressed bool `json:"excludeAddressed"`
}

type AfxdpNetworkPool struct {
	Name string `json:"name"`
	// Mode sets the mode of the AF_XDP Device Pool.
	// +kubebuilder:validation:Enum={"primary", "cdq",""}
	// +kubebuilder:default:= "primary"
	Mode string `json:"mode"`
	// NumPrimary defines the maximum number of primary devices this pool will take, per node.
	NumPrimary int `json:"numPrimary"`
	//NumSfs sets the maximum number of secondary devices this pool will create, per primary device.
	NumSfs int `json:"numSfs"`
	// NicSelector selects the NICs to be configured
	NicSelector AfxdpNetworkNicSelector `json:"nicSelector"`
	// udsServer configures UDS handshaking with the AF_XDP DP (enable|disable)
	// +kubebuilder:validation:Enum={"enable","disable"}
	// +kubebuilder:default:= "enable"
	UdsServer bool `json:"udsServer"`
	// BpfMapPinning configures map pinning for the the AF_XDP DP (enable|disable)
	// +kubebuilder:validation:Enum={"enable","disable"}
	// +kubebuilder:default:= "disable"
	BpfMapPinning bool `json:"bpfMapPinning"`
	// UdsTimeout defines the wait time for the AF_XDP on a UDS for a handshake to initiate.
	UdsTimeout int `json:"udsTimeout"`
}

type AfxdpNetworkNicSelector struct {
	Drivers []*AfxdpNetworkPoolDriver `json:"Drivers"`
	Devices []*AfxdpNetworkPoolDevice `json:"Devices"`
}

// AfxdpNetworkNodePolicyStatus defines the observed state of AfxdpNetworkNodePolicy
type AfxdpNetworkNodePolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AfxdpNetworkNodePolicy is the Schema for the afxdpnetworknodepolicies API
type AfxdpNetworkNodePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AfxdpNetworkNodePolicySpec   `json:"spec,omitempty"`
	Status AfxdpNetworkNodePolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AfxdpNetworkNodePolicyList contains a list of AfxdpNetworkNodePolicy
type AfxdpNetworkNodePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfxdpNetworkNodePolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AfxdpNetworkNodePolicy{}, &AfxdpNetworkNodePolicyList{})
}
