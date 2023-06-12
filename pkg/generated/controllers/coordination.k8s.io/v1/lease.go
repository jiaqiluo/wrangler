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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

// LeaseController interface for managing Lease resources.
type LeaseController interface {
	generic.ControllerMeta
	LeaseClient

	// OnChange runs the given handler when the controller detects a resource was changed.
	OnChange(ctx context.Context, name string, sync LeaseHandler)

	// OnRemove runs the given handler when the controller detects a resource was changed.
	OnRemove(ctx context.Context, name string, sync LeaseHandler)

	// Enqueue adds the resource with the given name to the worker queue of the controller.
	Enqueue(namespace, name string)

	// EnqueueAfter runs Enqueue after the provided duration.
	EnqueueAfter(namespace, name string, duration time.Duration)

	// Cache returns a cache for the resource type T.
	Cache() LeaseCache
}

// LeaseClient interface for managing Lease resources in Kubernetes.
type LeaseClient interface {
	// Create creates a new object and return the newly created Object or an error.
	Create(*v1.Lease) (*v1.Lease, error)

	// Update updates the object and return the newly updated Object or an error.
	Update(*v1.Lease) (*v1.Lease, error)

	// Delete deletes the Object in the given name.
	Delete(namespace, name string, options *metav1.DeleteOptions) error

	// Get will attempt to retrieve the resource with the specified name.
	Get(namespace, name string, options metav1.GetOptions) (*v1.Lease, error)

	// List will attempt to find multiple resources.
	List(namespace string, opts metav1.ListOptions) (*v1.LeaseList, error)

	// Watch will start watching resources.
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)

	// Patch will patch the resource with the matching name.
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Lease, err error)

	// WithImpersonation returns a new client that will use the provided impersonation config for new request.
	WithImpersonation(impersonate rest.ImpersonationConfig) (generic.ClientInterface[*v1.Lease, *v1.LeaseList], error)
}

// LeaseCache interface for retrieving Lease resources in memory.
type LeaseCache interface {
	// Get returns the resources with the specified name from the cache.
	Get(namespace, name string) (*v1.Lease, error)

	// List will attempt to find resources from the Cache.
	List(namespace string, selector labels.Selector) ([]*v1.Lease, error)

	// AddIndexer adds  a new Indexer to the cache with the provided name.
	// If you call this after you already have data in the store, the results are undefined.
	AddIndexer(indexName string, indexer LeaseIndexer)

	// GetByIndex returns the stored objects whose set of indexed values
	// for the named index includes the given indexed value.
	GetByIndex(indexName, key string) ([]*v1.Lease, error)
}

// LeaseHandler is function for performing any potential modifications to a Lease resource.
type LeaseHandler func(string, *v1.Lease) (*v1.Lease, error)

// LeaseIndexer computes a set of indexed values for the provided object.
type LeaseIndexer func(obj *v1.Lease) ([]string, error)

// LeaseGenericController wraps wrangler/pkg/generic.Controller so that the function definitions adhere to LeaseController interface.
type LeaseGenericController struct {
	generic.ControllerInterface[*v1.Lease, *v1.LeaseList]
}

// OnChange runs the given resource handler when the controller detects a resource was changed.
func (c *LeaseGenericController) OnChange(ctx context.Context, name string, sync LeaseHandler) {
	c.ControllerInterface.OnChange(ctx, name, generic.ObjectHandler[*v1.Lease](sync))
}

// OnRemove runs the given object handler when the controller detects a resource was changed.
func (c *LeaseGenericController) OnRemove(ctx context.Context, name string, sync LeaseHandler) {
	c.ControllerInterface.OnRemove(ctx, name, generic.ObjectHandler[*v1.Lease](sync))
}

// Cache returns a cache of resources in memory.
func (c *LeaseGenericController) Cache() LeaseCache {
	return &LeaseGenericCache{
		c.ControllerInterface.Cache(),
	}
}

// LeaseGenericCache wraps wrangler/pkg/generic.Cache so the function definitions adhere to LeaseCache interface.
type LeaseGenericCache struct {
	generic.CacheInterface[*v1.Lease]
}

// AddIndexer adds  a new Indexer to the cache with the provided name.
// If you call this after you already have data in the store, the results are undefined.
func (c LeaseGenericCache) AddIndexer(indexName string, indexer LeaseIndexer) {
	c.CacheInterface.AddIndexer(indexName, generic.Indexer[*v1.Lease](indexer))
}
