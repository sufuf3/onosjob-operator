package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ONOSJobSpec defines the desired state of ONOSJob
// +k8s:openapi-gen=true
type ONOSJobSpec struct {
	ControllerIp   string        `json:"controller-ip"`
	ControllerPort string        `json:"controller-port"`
	Hosts          []Hosts       `json:"hosts,omitempty" validate:"-"`
	FlowsDevice    []FlowsDevice `json:"flowsDevice,omitempty" validate:"-"`
}

// ONOSJobStatus defines the observed state of ONOSJob
// +k8s:openapi-gen=true
type ONOSJobStatus struct {
}

type Hosts struct {
	Mac         string          `json:"mac"`
	Vlan        string          `json:"vlan"`
	IpAddresses []string        `json:"ipAddresses"`
	Locations   []HostLocations `json:"locations"`
}

type HostLocations struct {
	ElementId string `json:"elementId"`
	Port      string `json:"port"`
}

type FlowsDevice struct {
	Deviceid     string                    `json:"deviceId"`
	Appid        string                    `json:"appId,omitempty"`
	Priority     int64                     `json:"priority"`
	IsPermanent  bool                      `json:"isPermanent"`
	Timeout      int64                     `json:"timeout"`
	Instructions []FlowsDeviceInstructions `json:"instructions"`
	Criteria     []FlowsDeviceCriteria     `json:"criteria"`
}

type FlowsDeviceInstructions struct {
	Type string `json:"type"`
	Port string `json:"port"`
}
type FlowsDeviceCriteria struct {
	Type    string `json:"type"`
	EthType string `json:"ethType"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ONOSJob is the Schema for the onosjobs API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ONOSJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ONOSJobSpec   `json:"spec,omitempty"`
	Status ONOSJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ONOSJobList contains a list of ONOSJob
type ONOSJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ONOSJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ONOSJob{}, &ONOSJobList{})
}
