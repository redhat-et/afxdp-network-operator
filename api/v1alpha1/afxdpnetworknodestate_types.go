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

// AfxdpNetworkNodeStateSpec defines the desired state of AfxdpNetworkNodeState
type AfxdpNetworkNodeStateSpec struct {
	DpConfigVersion string     `json:"dpConfigVersion,omitempty"`
	Interfaces      Interfaces `json:"interfaces,omitempty"`
}

type Interfaces []Interface

type Interface struct {
	Name       string `json:"name,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Driver     string `json:"driver,omitempty"`
	PciAddress string `json:"pciAddress"`
	Mac        string `json:"mac,omitempty"`
	// Primary     *Interface
	SubFunctions []SubFunction
	NumSfs       int `json:"numSfs,omitempty"`
	// SfGroups   []SfGroup `json:"sfGroups,omitempty"`
}

// type SfGroup struct {
// 	DeviceType string `json:"deviceType,omitempty"`
// 	SfRange    string `json:"sfRange,omitempty"`
// 	PolicyName string `json:"policyName,omitempty"`
// 	Mtu        int    `json:"mtu,omitempty"`
// }

type InterfaceExt struct {
	Name        string        `json:"name,omitempty"`
	Mode        string        `json:"mode,omitempty"`
	Mac         string        `json:"mac,omitempty"`
	Driver      string        `json:"driver,omitempty"`
	PciAddress  string        `json:"pciAddress"`
	Vendor      string        `json:"vendor,omitempty"`
	DeviceID    string        `json:"deviceID,omitempty"`
	NetFilter   string        `json:"netFilter,omitempty"`
	Mtu         int           `json:"mtu,omitempty"`
	NumSfs      int           `json:"numSfs,omitempty"`
	LinkSpeed   string        `json:"linkSpeed,omitempty"`
	LinkType    string        `json:"linkType,omitempty"`
	EswitchMode string        `json:"eSwitchMode,omitempty"`
	SFs         []SubFunction `json:"Sfs,omitempty"`
}

type InterfaceExts []InterfaceExt

type SubFunction struct {
	Name       string `json:"name,omitempty"`
	Mac        string `json:"mac,omitempty"`
	Assigned   string `json:"assigned,omitempty"`
	Driver     string `json:"driver,omitempty"`
	PciAddress string `json:"pciAddress"`
	Mtu        int    `json:"mtu,omitempty"`
	SfID       int    `json:"sfID"`
}

// AfxdpNetworkNodeStateStatus defines the observed state of AfxdpNetworkNodeState
type AfxdpNetworkNodeStateStatus struct {
	Interfaces    InterfaceExts `json:"interfaces,omitempty"`
	SyncStatus    string        `json:"syncStatus,omitempty"`
	LastSyncError string        `json:"lastSyncError,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AfxdpNetworkNodeState is the Schema for the afxdpnetworknodestates API
type AfxdpNetworkNodeState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AfxdpNetworkNodeStateSpec   `json:"spec,omitempty"`
	Status AfxdpNetworkNodeStateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AfxdpNetworkNodeStateList contains a list of AfxdpNetworkNodeState
type AfxdpNetworkNodeStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfxdpNetworkNodeState `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AfxdpNetworkNodeState{}, &AfxdpNetworkNodeStateList{})
}
