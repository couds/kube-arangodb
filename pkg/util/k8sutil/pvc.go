//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package k8sutil

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreatePersistentVolumeClaimName returns the name of the persistent volume claim for a member with
// a given id in a deployment with a given name.
func CreatePersistentVolumeClaimName(deploymentName, role, id string) string {
	return deploymentName + "-" + role + "-" + stripArangodPrefix(id)
}

// CreatePersistentVolumeClaim creates a persistent volume claim with given name and configuration.
// If the pvc already exists, nil is returned.
// If another error occurs, that error is returned.
func CreatePersistentVolumeClaim(kubecli kubernetes.Interface, pvcName, deploymentName, ns, storageClassName, role string, resources v1.ResourceRequirements, owner metav1.OwnerReference) error {
	labels := LabelsForDeployment(deploymentName, role)
	volumeMode := v1.PersistentVolumeFilesystem
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvcName,
			Labels: labels,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				v1.ReadWriteOnce,
			},
			VolumeMode: &volumeMode,
			Resources:  resources,
		},
	}
	if storageClassName != "" {
		pvc.Spec.StorageClassName = &storageClassName
	}
	addOwnerRefToObject(pvc.GetObjectMeta(), &owner)
	if _, err := kubecli.CoreV1().PersistentVolumeClaims(ns).Create(pvc); err != nil && !IsAlreadyExists(err) {
		return maskAny(err)
	}
	return nil
}
