/*
Copyright 2019 Banzai Cloud.

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

package fake

import (
	"context"

	v1beta1 "github.com/banzaicloud/istio-operator/pkg/apis/istio/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMeshGateways implements MeshGatewayInterface
type FakeMeshGateways struct {
	Fake *FakeIstioV1beta1
	ns   string
}

var meshgatewaysResource = schema.GroupVersionResource{Group: "istio.banzaicloud.io", Version: "v1beta1", Resource: "meshgateways"}

var meshgatewaysKind = schema.GroupVersionKind{Group: "istio.banzaicloud.io", Version: "v1beta1", Kind: "MeshGateway"}

// Get takes name of the meshGateway, and returns the corresponding meshGateway object, and an error if there is any.
func (c *FakeMeshGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MeshGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(meshgatewaysResource, c.ns, name), &v1beta1.MeshGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MeshGateway), err
}

// List takes label and field selectors, and returns the list of MeshGateways that match those selectors.
func (c *FakeMeshGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MeshGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(meshgatewaysResource, meshgatewaysKind, c.ns, opts), &v1beta1.MeshGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MeshGatewayList{ListMeta: obj.(*v1beta1.MeshGatewayList).ListMeta}
	for _, item := range obj.(*v1beta1.MeshGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested meshGateways.
func (c *FakeMeshGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(meshgatewaysResource, c.ns, opts))

}

// Create takes the representation of a meshGateway and creates it.  Returns the server's representation of the meshGateway, and an error, if there is any.
func (c *FakeMeshGateways) Create(ctx context.Context, meshGateway *v1beta1.MeshGateway, opts v1.CreateOptions) (result *v1beta1.MeshGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(meshgatewaysResource, c.ns, meshGateway), &v1beta1.MeshGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MeshGateway), err
}

// Update takes the representation of a meshGateway and updates it. Returns the server's representation of the meshGateway, and an error, if there is any.
func (c *FakeMeshGateways) Update(ctx context.Context, meshGateway *v1beta1.MeshGateway, opts v1.UpdateOptions) (result *v1beta1.MeshGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(meshgatewaysResource, c.ns, meshGateway), &v1beta1.MeshGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MeshGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMeshGateways) UpdateStatus(ctx context.Context, meshGateway *v1beta1.MeshGateway, opts v1.UpdateOptions) (*v1beta1.MeshGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(meshgatewaysResource, "status", c.ns, meshGateway), &v1beta1.MeshGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MeshGateway), err
}

// Delete takes name of the meshGateway and deletes it. Returns an error if one occurs.
func (c *FakeMeshGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(meshgatewaysResource, c.ns, name), &v1beta1.MeshGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMeshGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(meshgatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MeshGatewayList{})
	return err
}

// Patch applies the patch and returns the patched meshGateway.
func (c *FakeMeshGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MeshGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(meshgatewaysResource, c.ns, name, pt, data, subresources...), &v1beta1.MeshGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MeshGateway), err
}
