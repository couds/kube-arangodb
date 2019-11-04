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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/arangodb/kube-arangodb/pkg/apis/backup/v1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoBackupPoliciesGetter has a method to return a ArangoBackupPolicyInterface.
// A group's client should implement this interface.
type ArangoBackupPoliciesGetter interface {
	ArangoBackupPolicies(namespace string) ArangoBackupPolicyInterface
}

// ArangoBackupPolicyInterface has methods to work with ArangoBackupPolicy resources.
type ArangoBackupPolicyInterface interface {
	Create(*v1.ArangoBackupPolicy) (*v1.ArangoBackupPolicy, error)
	Update(*v1.ArangoBackupPolicy) (*v1.ArangoBackupPolicy, error)
	UpdateStatus(*v1.ArangoBackupPolicy) (*v1.ArangoBackupPolicy, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ArangoBackupPolicy, error)
	List(opts metav1.ListOptions) (*v1.ArangoBackupPolicyList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ArangoBackupPolicy, err error)
	ArangoBackupPolicyExpansion
}

// arangoBackupPolicies implements ArangoBackupPolicyInterface
type arangoBackupPolicies struct {
	client rest.Interface
	ns     string
}

// newArangoBackupPolicies returns a ArangoBackupPolicies
func newArangoBackupPolicies(c *BackupV1Client, namespace string) *arangoBackupPolicies {
	return &arangoBackupPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoBackupPolicy, and returns the corresponding arangoBackupPolicy object, and an error if there is any.
func (c *arangoBackupPolicies) Get(name string, options metav1.GetOptions) (result *v1.ArangoBackupPolicy, err error) {
	result = &v1.ArangoBackupPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoBackupPolicies that match those selectors.
func (c *arangoBackupPolicies) List(opts metav1.ListOptions) (result *v1.ArangoBackupPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ArangoBackupPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoBackupPolicies.
func (c *arangoBackupPolicies) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a arangoBackupPolicy and creates it.  Returns the server's representation of the arangoBackupPolicy, and an error, if there is any.
func (c *arangoBackupPolicies) Create(arangoBackupPolicy *v1.ArangoBackupPolicy) (result *v1.ArangoBackupPolicy, err error) {
	result = &v1.ArangoBackupPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		Body(arangoBackupPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a arangoBackupPolicy and updates it. Returns the server's representation of the arangoBackupPolicy, and an error, if there is any.
func (c *arangoBackupPolicies) Update(arangoBackupPolicy *v1.ArangoBackupPolicy) (result *v1.ArangoBackupPolicy, err error) {
	result = &v1.ArangoBackupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		Name(arangoBackupPolicy.Name).
		Body(arangoBackupPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *arangoBackupPolicies) UpdateStatus(arangoBackupPolicy *v1.ArangoBackupPolicy) (result *v1.ArangoBackupPolicy, err error) {
	result = &v1.ArangoBackupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		Name(arangoBackupPolicy.Name).
		SubResource("status").
		Body(arangoBackupPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the arangoBackupPolicy and deletes it. Returns an error if one occurs.
func (c *arangoBackupPolicies) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoBackupPolicies) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched arangoBackupPolicy.
func (c *arangoBackupPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ArangoBackupPolicy, err error) {
	result = &v1.ArangoBackupPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangobackuppolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
