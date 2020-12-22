package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Demo is an CRD specification.
type Demo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec DemoSpec `json:"spec"`
}

// DemoSpec is a specification for a DemoSpec resource.
type DemoSpec struct {
	Name string `json:"name"`
}
