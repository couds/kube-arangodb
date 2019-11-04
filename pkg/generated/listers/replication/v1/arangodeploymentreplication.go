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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/apis/replication/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoDeploymentReplicationLister helps list ArangoDeploymentReplications.
type ArangoDeploymentReplicationLister interface {
	// List lists all ArangoDeploymentReplications in the indexer.
	List(selector labels.Selector) (ret []*v1.ArangoDeploymentReplication, err error)
	// ArangoDeploymentReplications returns an object that can list and get ArangoDeploymentReplications.
	ArangoDeploymentReplications(namespace string) ArangoDeploymentReplicationNamespaceLister
	ArangoDeploymentReplicationListerExpansion
}

// arangoDeploymentReplicationLister implements the ArangoDeploymentReplicationLister interface.
type arangoDeploymentReplicationLister struct {
	indexer cache.Indexer
}

// NewArangoDeploymentReplicationLister returns a new ArangoDeploymentReplicationLister.
func NewArangoDeploymentReplicationLister(indexer cache.Indexer) ArangoDeploymentReplicationLister {
	return &arangoDeploymentReplicationLister{indexer: indexer}
}

// List lists all ArangoDeploymentReplications in the indexer.
func (s *arangoDeploymentReplicationLister) List(selector labels.Selector) (ret []*v1.ArangoDeploymentReplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoDeploymentReplication))
	})
	return ret, err
}

// ArangoDeploymentReplications returns an object that can list and get ArangoDeploymentReplications.
func (s *arangoDeploymentReplicationLister) ArangoDeploymentReplications(namespace string) ArangoDeploymentReplicationNamespaceLister {
	return arangoDeploymentReplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoDeploymentReplicationNamespaceLister helps list and get ArangoDeploymentReplications.
type ArangoDeploymentReplicationNamespaceLister interface {
	// List lists all ArangoDeploymentReplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ArangoDeploymentReplication, err error)
	// Get retrieves the ArangoDeploymentReplication from the indexer for a given namespace and name.
	Get(name string) (*v1.ArangoDeploymentReplication, error)
	ArangoDeploymentReplicationNamespaceListerExpansion
}

// arangoDeploymentReplicationNamespaceLister implements the ArangoDeploymentReplicationNamespaceLister
// interface.
type arangoDeploymentReplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoDeploymentReplications in the indexer for a given namespace.
func (s arangoDeploymentReplicationNamespaceLister) List(selector labels.Selector) (ret []*v1.ArangoDeploymentReplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoDeploymentReplication))
	})
	return ret, err
}

// Get retrieves the ArangoDeploymentReplication from the indexer for a given namespace and name.
func (s arangoDeploymentReplicationNamespaceLister) Get(name string) (*v1.ArangoDeploymentReplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("arangodeploymentreplication"), name)
	}
	return obj.(*v1.ArangoDeploymentReplication), nil
}
