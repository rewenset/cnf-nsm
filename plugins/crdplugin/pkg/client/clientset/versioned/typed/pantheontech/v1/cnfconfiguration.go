/*
 * Copyright (c) 2019 PANTHEON.tech s.r.o. All rights reserved.
 *
 * CNF-NSM License. For licensing terms please contact sales@pantheon.tech.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "go.cdnf.io/cnf-nsm/plugins/crdplugin/pkg/apis/pantheontech/v1"
	scheme "go.cdnf.io/cnf-nsm/plugins/crdplugin/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CNFConfigurationsGetter has a method to return a CNFConfigurationInterface.
// A group's client should implement this interface.
type CNFConfigurationsGetter interface {
	CNFConfigurations(namespace string) CNFConfigurationInterface
}

// CNFConfigurationInterface has methods to work with CNFConfiguration resources.
type CNFConfigurationInterface interface {
	Create(*v1.CNFConfiguration) (*v1.CNFConfiguration, error)
	Update(*v1.CNFConfiguration) (*v1.CNFConfiguration, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CNFConfiguration, error)
	List(opts metav1.ListOptions) (*v1.CNFConfigurationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CNFConfiguration, err error)
	CNFConfigurationExpansion
}

// cNFConfigurations implements CNFConfigurationInterface
type cNFConfigurations struct {
	client rest.Interface
	ns     string
}

// newCNFConfigurations returns a CNFConfigurations
func newCNFConfigurations(c *PantheonV1Client, namespace string) *cNFConfigurations {
	return &cNFConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cNFConfiguration, and returns the corresponding cNFConfiguration object, and an error if there is any.
func (c *cNFConfigurations) Get(name string, options metav1.GetOptions) (result *v1.CNFConfiguration, err error) {
	result = &v1.CNFConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CNFConfigurations that match those selectors.
func (c *cNFConfigurations) List(opts metav1.ListOptions) (result *v1.CNFConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CNFConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cNFConfigurations.
func (c *cNFConfigurations) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cNFConfiguration and creates it.  Returns the server's representation of the cNFConfiguration, and an error, if there is any.
func (c *cNFConfigurations) Create(cNFConfiguration *v1.CNFConfiguration) (result *v1.CNFConfiguration, err error) {
	result = &v1.CNFConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		Body(cNFConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cNFConfiguration and updates it. Returns the server's representation of the cNFConfiguration, and an error, if there is any.
func (c *cNFConfigurations) Update(cNFConfiguration *v1.CNFConfiguration) (result *v1.CNFConfiguration, err error) {
	result = &v1.CNFConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		Name(cNFConfiguration.Name).
		Body(cNFConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the cNFConfiguration and deletes it. Returns an error if one occurs.
func (c *cNFConfigurations) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cNFConfigurations) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cnfconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cNFConfiguration.
func (c *cNFConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CNFConfiguration, err error) {
	result = &v1.CNFConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cnfconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
