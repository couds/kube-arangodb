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
// Author Adam Janikowski
//

package policy

import (
	"fmt"
	"reflect"
	"time"

	"github.com/arangodb/kube-arangodb/pkg/backup/operator/event"

	"github.com/arangodb/kube-arangodb/pkg/backup/operator/operation"

	"k8s.io/client-go/kubernetes"

	database "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1alpha"
	arangoClientSet "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	"github.com/robfig/cron"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	BackupCreated = "ArangoBackupCreated"
	PolicyError   = "Error"
	Rescheduled   = "Rescheduled"
)

type handler struct {
	client        arangoClientSet.Interface
	kubeClient    kubernetes.Interface
	eventRecorder event.EventRecorderInstance
}

func (*handler) Name() string {
	return database.ArangoBackupPolicyResourceKind
}

func (h *handler) Handle(item operation.Item) error {
	// Do not act on delete event, finalizers are used
	if item.Operation == operation.OperationDelete {
		return nil
	}

	// Get Backup object. It also cover NotFound case
	policy, err := h.client.DatabaseV1alpha().ArangoBackupPolicies(item.Namespace).Get(item.Name, meta.GetOptions{})
	if err != nil {
		return err
	}

	status, err := h.processBackupPolicy(policy.DeepCopy())
	if err != nil {
		return err
	}

	// Nothing to update, objects are equal
	if reflect.DeepEqual(policy.Status, status) {
		return nil
	}

	policy.Status = status

	// Update status on object
	if _, err = h.client.DatabaseV1alpha().ArangoBackupPolicies(item.Namespace).UpdateStatus(policy); err != nil {
		return err
	}

	return nil
}

func (h *handler) processBackupPolicy(policy *database.ArangoBackupPolicy) (database.ArangoBackupPolicyStatus, error) {
	if err := policy.Validate(); err != nil {
		h.eventRecorder.Warning(policy, PolicyError, "Policy Error: %s", err.Error())

		return database.ArangoBackupPolicyStatus{
			Message: fmt.Sprintf("Validation error: %s", err.Error()),
		}, nil
	}

	if policy.Status.Scheduled.IsZero() {
		expr, err := cron.Parse(policy.Spec.Schedule)
		if err != nil {
			h.eventRecorder.Warning(policy, PolicyError, "Policy Error: %s", err.Error())

			return database.ArangoBackupPolicyStatus{
				Message: fmt.Sprintf("error while parsing expr: %s", err.Error()),
			}, nil
		}

		next := expr.Next(time.Now())

		return database.ArangoBackupPolicyStatus{
			Scheduled: meta.Time{
				Time: next,
			},
		}, nil
	}

	// Check if schedule is required
	if policy.Status.Scheduled.Unix() > time.Now().Unix() {
		return policy.Status, nil
	}

	// Schedule new deployments
	listOptions := meta.ListOptions{}

	if policy.Spec.DeploymentSelector != nil &&
		(policy.Spec.DeploymentSelector.MatchLabels != nil &&
			len(policy.Spec.DeploymentSelector.MatchLabels) > 0 ||
			policy.Spec.DeploymentSelector.MatchExpressions != nil) {
		listOptions.LabelSelector = meta.FormatLabelSelector(policy.Spec.DeploymentSelector)
	}

	deployments, err := h.client.DatabaseV1alpha().ArangoDeployments(policy.Namespace).List(listOptions)

	if err != nil {
		h.eventRecorder.Warning(policy, PolicyError, "Policy Error: %s", err.Error())

		return database.ArangoBackupPolicyStatus{
			Scheduled: policy.Status.Scheduled,
			Message:   fmt.Sprintf("deployments listing failed: %s", err.Error()),
		}, nil
	}

	for _, deployment := range deployments.Items {
		backup := policy.NewBackup(deployment.DeepCopy())

		if _, err := h.client.DatabaseV1alpha().ArangoBackups(backup.Namespace).Create(backup); err != nil {
			h.eventRecorder.Warning(policy, PolicyError, "Policy Error: %s", err.Error())

			return database.ArangoBackupPolicyStatus{
				Scheduled: policy.Status.Scheduled,
				Message:   fmt.Sprintf("backup creation failed: %s", err.Error()),
			}, nil
		}

		h.eventRecorder.Normal(policy, BackupCreated, "Created ArangoBackup: %s/%s", backup.Namespace, backup.Name)
	}

	expr, err := cron.Parse(policy.Spec.Schedule)
	if err != nil {
		return database.ArangoBackupPolicyStatus{
			Message: fmt.Sprintf("error while parsing expr: %s", err.Error()),
		}, nil
	}

	next := expr.Next(time.Now())

	h.eventRecorder.Normal(policy, Rescheduled, "Rescheduled for: %s", next.String())

	return database.ArangoBackupPolicyStatus{
		Scheduled: meta.Time{
			Time: next,
		},
	}, nil
}

func (*handler) CanBeHandled(item operation.Item) bool {
	return item.Group == database.SchemeGroupVersion.Group &&
		item.Version == database.SchemeGroupVersion.Version &&
		item.Kind == database.ArangoBackupPolicyResourceKind
}