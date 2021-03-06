/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RqliteClusterSpec defines the desired state of RqliteCluster
type RqliteClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name          string            `json:"name,omitempty"`
	ClusterID     string            `json:"clusterid,omitempty"`
	ClusterSize   int8              `json:"clustersize,omitempty"`
	InternalID    string            `json:"ii,omitempty"`
	CustomerInfos map[string]string `json:"ci,omitempty"`
	//	OthersInfos    []interface{}
}

// RqliteClusterStatus defines the observed state of RqliteCluster
type RqliteClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	CurrentStatus string `json:"cs,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.name", name=Name, type=string
// +kubebuilder:printcolumn:JSONPath=".status.cs", name=CurrentStatus, type=string

// RqliteCluster is the Schema for the rqliteclusters API
type RqliteCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RqliteClusterSpec   `json:"spec,omitempty"`
	Status RqliteClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RqliteClusterList contains a list of RqliteCluster
type RqliteClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RqliteCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RqliteCluster{}, &RqliteClusterList{})
}
