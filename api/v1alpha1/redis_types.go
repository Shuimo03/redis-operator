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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisSpec defines the desired state of Redis
type RedisSpec struct {
	RedisReplicas   uint                   `json:"redisReplicas,omitempty"`
	Image           corev1.ContainerImage  `json:"image,omitempty"`
	ImagePullPolicy corev1.PullPolicy      `json:"imagePullPolicy,omitempty"`
	Command         []string               `json:"command,omitempty"`
	Password        string                 `json:"password,omitempty"`
	Storage         RedisStorage           `json:"storage,omitempty"`
	Template        corev1.PodTemplateSpec `json:"template,omitempty"`
}

// RedisStatus defines the observed state of Redis
type RedisStatus struct {
	Creating   string `json:"creating,omitempty"`
	Running    string `json:"running,omitempty"`
	Failed     string `json:"failed,omitempty"`
	Scaling    string `json:"scaling,omitempty"`
	Restarting string `json:"restarting,omitempty"`
	Upgrading  string `json:"upgrading,omitempty"`
	Deleting   string `json:"deleting,omitempty"`
}

type RedisStorage struct {
	EmptyDir              *corev1.EmptyDirVolumeSource `json:"emptyDir,omitempty"`
	PersistentVolumeClaim corev1.PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Redis is the Schema for the redis API
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisSpec   `json:"spec,omitempty"`
	Status RedisStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RedisList contains a list of Redis
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Redis{}, &RedisList{})
}
