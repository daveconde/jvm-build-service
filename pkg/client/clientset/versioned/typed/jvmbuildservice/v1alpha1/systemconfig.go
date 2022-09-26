/*
Copyright 2021-2022 Red Hat, Inc.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/redhat-appstudio/jvm-build-service/pkg/apis/jvmbuildservice/v1alpha1"
	scheme "github.com/redhat-appstudio/jvm-build-service/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SystemConfigsGetter has a method to return a SystemConfigInterface.
// A group's client should implement this interface.
type SystemConfigsGetter interface {
	SystemConfigs(namespace string) SystemConfigInterface
}

// SystemConfigInterface has methods to work with SystemConfig resources.
type SystemConfigInterface interface {
	Create(ctx context.Context, systemConfig *v1alpha1.SystemConfig, opts v1.CreateOptions) (*v1alpha1.SystemConfig, error)
	Update(ctx context.Context, systemConfig *v1alpha1.SystemConfig, opts v1.UpdateOptions) (*v1alpha1.SystemConfig, error)
	UpdateStatus(ctx context.Context, systemConfig *v1alpha1.SystemConfig, opts v1.UpdateOptions) (*v1alpha1.SystemConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SystemConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SystemConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SystemConfig, err error)
	SystemConfigExpansion
}

// systemConfigs implements SystemConfigInterface
type systemConfigs struct {
	client rest.Interface
	ns     string
}

// newSystemConfigs returns a SystemConfigs
func newSystemConfigs(c *JvmbuildserviceV1alpha1Client, namespace string) *systemConfigs {
	return &systemConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the systemConfig, and returns the corresponding systemConfig object, and an error if there is any.
func (c *systemConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SystemConfig, err error) {
	result = &v1alpha1.SystemConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("systemconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SystemConfigs that match those selectors.
func (c *systemConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SystemConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SystemConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("systemconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested systemConfigs.
func (c *systemConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("systemconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a systemConfig and creates it.  Returns the server's representation of the systemConfig, and an error, if there is any.
func (c *systemConfigs) Create(ctx context.Context, systemConfig *v1alpha1.SystemConfig, opts v1.CreateOptions) (result *v1alpha1.SystemConfig, err error) {
	result = &v1alpha1.SystemConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("systemconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(systemConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a systemConfig and updates it. Returns the server's representation of the systemConfig, and an error, if there is any.
func (c *systemConfigs) Update(ctx context.Context, systemConfig *v1alpha1.SystemConfig, opts v1.UpdateOptions) (result *v1alpha1.SystemConfig, err error) {
	result = &v1alpha1.SystemConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("systemconfigs").
		Name(systemConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(systemConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *systemConfigs) UpdateStatus(ctx context.Context, systemConfig *v1alpha1.SystemConfig, opts v1.UpdateOptions) (result *v1alpha1.SystemConfig, err error) {
	result = &v1alpha1.SystemConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("systemconfigs").
		Name(systemConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(systemConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the systemConfig and deletes it. Returns an error if one occurs.
func (c *systemConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("systemconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *systemConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("systemconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched systemConfig.
func (c *systemConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SystemConfig, err error) {
	result = &v1alpha1.SystemConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("systemconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}