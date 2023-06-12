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

	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

// NetworkPolicyController interface for managing NetworkPolicy resources.
type NetworkPolicyController interface {
	generic.ControllerMeta
	NetworkPolicyClient

	// OnChange runs the given handler when the controller detects a resource was changed.
	OnChange(ctx context.Context, name string, sync NetworkPolicyHandler)

	// OnRemove runs the given handler when the controller detects a resource was changed.
	OnRemove(ctx context.Context, name string, sync NetworkPolicyHandler)

	// Enqueue adds the resource with the given name to the worker queue of the controller.
	Enqueue(namespace, name string)

	// EnqueueAfter runs Enqueue after the provided duration.
	EnqueueAfter(namespace, name string, duration time.Duration)

	// Cache returns a cache for the resource type T.
	Cache() NetworkPolicyCache
}

// NetworkPolicyClient interface for managing NetworkPolicy resources in Kubernetes.
type NetworkPolicyClient interface {
	// Create creates a new object and return the newly created Object or an error.
	Create(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)

	// Update updates the object and return the newly updated Object or an error.
	Update(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)
	// UpdateStatus updates the Status field of a the object and return the newly updated Object or an error.
	// Will always return an error if the object does not have a status field.
	UpdateStatus(*v1.NetworkPolicy) (*v1.NetworkPolicy, error)

	// Delete deletes the Object in the given name.
	Delete(namespace, name string, options *metav1.DeleteOptions) error

	// Get will attempt to retrieve the resource with the specified name.
	Get(namespace, name string, options metav1.GetOptions) (*v1.NetworkPolicy, error)

	// List will attempt to find multiple resources.
	List(namespace string, opts metav1.ListOptions) (*v1.NetworkPolicyList, error)

	// Watch will start watching resources.
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)

	// Patch will patch the resource with the matching name.
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetworkPolicy, err error)

	// WithImpersonation returns a new client that will use the provided impersonation config for new request.
	WithImpersonation(impersonate rest.ImpersonationConfig) (generic.ClientInterface[*v1.NetworkPolicy, *v1.NetworkPolicyList], error)
}

// NetworkPolicyCache interface for retrieving NetworkPolicy resources in memory.
type NetworkPolicyCache interface {
	// Get returns the resources with the specified name from the cache.
	Get(namespace, name string) (*v1.NetworkPolicy, error)

	// List will attempt to find resources from the Cache.
	List(namespace string, selector labels.Selector) ([]*v1.NetworkPolicy, error)

	// AddIndexer adds  a new Indexer to the cache with the provided name.
	// If you call this after you already have data in the store, the results are undefined.
	AddIndexer(indexName string, indexer NetworkPolicyIndexer)

	// GetByIndex returns the stored objects whose set of indexed values
	// for the named index includes the given indexed value.
	GetByIndex(indexName, key string) ([]*v1.NetworkPolicy, error)
}

// NetworkPolicyHandler is function for performing any potential modifications to a NetworkPolicy resource.
type NetworkPolicyHandler func(string, *v1.NetworkPolicy) (*v1.NetworkPolicy, error)

// NetworkPolicyIndexer computes a set of indexed values for the provided object.
type NetworkPolicyIndexer func(obj *v1.NetworkPolicy) ([]string, error)

// NetworkPolicyGenericController wraps wrangler/pkg/generic.Controller so that the function definitions adhere to NetworkPolicyController interface.
type NetworkPolicyGenericController struct {
	generic.ControllerInterface[*v1.NetworkPolicy, *v1.NetworkPolicyList]
}

// OnChange runs the given resource handler when the controller detects a resource was changed.
func (c *NetworkPolicyGenericController) OnChange(ctx context.Context, name string, sync NetworkPolicyHandler) {
	c.ControllerInterface.OnChange(ctx, name, generic.ObjectHandler[*v1.NetworkPolicy](sync))
}

// OnRemove runs the given object handler when the controller detects a resource was changed.
func (c *NetworkPolicyGenericController) OnRemove(ctx context.Context, name string, sync NetworkPolicyHandler) {
	c.ControllerInterface.OnRemove(ctx, name, generic.ObjectHandler[*v1.NetworkPolicy](sync))
}

// Cache returns a cache of resources in memory.
func (c *NetworkPolicyGenericController) Cache() NetworkPolicyCache {
	return &NetworkPolicyGenericCache{
		c.ControllerInterface.Cache(),
	}
}

// NetworkPolicyGenericCache wraps wrangler/pkg/generic.Cache so the function definitions adhere to NetworkPolicyCache interface.
type NetworkPolicyGenericCache struct {
	generic.CacheInterface[*v1.NetworkPolicy]
}

// AddIndexer adds  a new Indexer to the cache with the provided name.
// If you call this after you already have data in the store, the results are undefined.
func (c NetworkPolicyGenericCache) AddIndexer(indexName string, indexer NetworkPolicyIndexer) {
	c.CacheInterface.AddIndexer(indexName, generic.Indexer[*v1.NetworkPolicy](indexer))
}

type NetworkPolicyStatusHandler func(obj *v1.NetworkPolicy, status v1.NetworkPolicyStatus) (v1.NetworkPolicyStatus, error)

type NetworkPolicyGeneratingHandler func(obj *v1.NetworkPolicy, status v1.NetworkPolicyStatus) ([]runtime.Object, v1.NetworkPolicyStatus, error)

func FromNetworkPolicyHandlerToHandler(sync NetworkPolicyHandler) generic.Handler {
	return generic.FromObjectHandlerToHandler(generic.ObjectHandler[*v1.NetworkPolicy](sync))
}

func RegisterNetworkPolicyStatusHandler(ctx context.Context, controller NetworkPolicyController, condition condition.Cond, name string, handler NetworkPolicyStatusHandler) {
	statusHandler := &networkPolicyStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromNetworkPolicyHandlerToHandler(statusHandler.sync))
}

func RegisterNetworkPolicyGeneratingHandler(ctx context.Context, controller NetworkPolicyController, apply apply.Apply,
	condition condition.Cond, name string, handler NetworkPolicyGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &networkPolicyGeneratingHandler{
		NetworkPolicyGeneratingHandler: handler,
		apply:                          apply,
		name:                           name,
		gvk:                            controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterNetworkPolicyStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type networkPolicyStatusHandler struct {
	client    NetworkPolicyClient
	condition condition.Cond
	handler   NetworkPolicyStatusHandler
}

func (a *networkPolicyStatusHandler) sync(key string, obj *v1.NetworkPolicy) (*v1.NetworkPolicy, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type networkPolicyGeneratingHandler struct {
	NetworkPolicyGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *networkPolicyGeneratingHandler) Remove(key string, obj *v1.NetworkPolicy) (*v1.NetworkPolicy, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.NetworkPolicy{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *networkPolicyGeneratingHandler) Handle(obj *v1.NetworkPolicy, status v1.NetworkPolicyStatus) (v1.NetworkPolicyStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.NetworkPolicyGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
