/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/convox/rack/pkg/atom/pkg/apis/convox/v2"
	scheme "github.com/convox/rack/pkg/atom/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AtomVersionsGetter has a method to return a AtomVersionInterface.
// A group's client should implement this interface.
type AtomVersionsGetter interface {
	AtomVersions(namespace string) AtomVersionInterface
}

// AtomVersionInterface has methods to work with AtomVersion resources.
type AtomVersionInterface interface {
	Create(*v2.AtomVersion) (*v2.AtomVersion, error)
	Update(*v2.AtomVersion) (*v2.AtomVersion, error)
	UpdateStatus(*v2.AtomVersion) (*v2.AtomVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.AtomVersion, error)
	List(opts v1.ListOptions) (*v2.AtomVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.AtomVersion, err error)
	AtomVersionExpansion
}

// atomVersions implements AtomVersionInterface
type atomVersions struct {
	client rest.Interface
	ns     string
}

// newAtomVersions returns a AtomVersions
func newAtomVersions(c *ConvoxV2Client, namespace string) *atomVersions {
	return &atomVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the atomVersion, and returns the corresponding atomVersion object, and an error if there is any.
func (c *atomVersions) Get(name string, options v1.GetOptions) (result *v2.AtomVersion, err error) {
	result = &v2.AtomVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("atomversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AtomVersions that match those selectors.
func (c *atomVersions) List(opts v1.ListOptions) (result *v2.AtomVersionList, err error) {
	result = &v2.AtomVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("atomversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested atomVersions.
func (c *atomVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("atomversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a atomVersion and creates it.  Returns the server's representation of the atomVersion, and an error, if there is any.
func (c *atomVersions) Create(atomVersion *v2.AtomVersion) (result *v2.AtomVersion, err error) {
	result = &v2.AtomVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("atomversions").
		Body(atomVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a atomVersion and updates it. Returns the server's representation of the atomVersion, and an error, if there is any.
func (c *atomVersions) Update(atomVersion *v2.AtomVersion) (result *v2.AtomVersion, err error) {
	result = &v2.AtomVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("atomversions").
		Name(atomVersion.Name).
		Body(atomVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *atomVersions) UpdateStatus(atomVersion *v2.AtomVersion) (result *v2.AtomVersion, err error) {
	result = &v2.AtomVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("atomversions").
		Name(atomVersion.Name).
		SubResource("status").
		Body(atomVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the atomVersion and deletes it. Returns an error if one occurs.
func (c *atomVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("atomversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *atomVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("atomversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched atomVersion.
func (c *atomVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.AtomVersion, err error) {
	result = &v2.AtomVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("atomversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
