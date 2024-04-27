package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScalingSpec defines the desired state of Scaling
type ScalingSpec struct {
	// targetNamespace is the name of the namespace that target exists
	TargetNamespace string `json:"targetNamespace"`
	// targetDeployment is the name of the deployment to scale
	TargetDeployment string `json:"targetDeployment"`
	// targetReplicaCount is the desired number of replicas for scaling
	TargetReplicaCount int `json:"targetReplicaCount"`
}

// ScalingStatus defines the observed state of Scaling
type ScalingStatus struct {
	// lastScaled is the timestamp when the last scaling action was performed
	LastScaled string `json:"lastScaled,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Scaling is the Schema for the scaling API
type Scaling struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScalingSpec   `json:"spec,omitempty"`
	Status ScalingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScalingList contains a list of Scaling
type ScalingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scaling `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Scaling{}, &ScalingList{})
}
