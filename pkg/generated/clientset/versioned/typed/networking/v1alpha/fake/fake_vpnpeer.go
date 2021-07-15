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

package fake

import (
	"context"

	v1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/networking/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVPNPeers implements VPNPeerInterface
type FakeVPNPeers struct {
	Fake *FakeNetworkingV1alpha
}

var vpnpeersResource = schema.GroupVersionResource{Group: "networking.edgenet.io", Version: "v1alpha", Resource: "vpnpeers"}

var vpnpeersKind = schema.GroupVersionKind{Group: "networking.edgenet.io", Version: "v1alpha", Kind: "VPNPeer"}

// Get takes name of the vPNPeer, and returns the corresponding vPNPeer object, and an error if there is any.
func (c *FakeVPNPeers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha.VPNPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(vpnpeersResource, name), &v1alpha.VPNPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.VPNPeer), err
}

// List takes label and field selectors, and returns the list of VPNPeers that match those selectors.
func (c *FakeVPNPeers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha.VPNPeerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(vpnpeersResource, vpnpeersKind, opts), &v1alpha.VPNPeerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.VPNPeerList{ListMeta: obj.(*v1alpha.VPNPeerList).ListMeta}
	for _, item := range obj.(*v1alpha.VPNPeerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vPNPeers.
func (c *FakeVPNPeers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(vpnpeersResource, opts))
}

// Create takes the representation of a vPNPeer and creates it.  Returns the server's representation of the vPNPeer, and an error, if there is any.
func (c *FakeVPNPeers) Create(ctx context.Context, vPNPeer *v1alpha.VPNPeer, opts v1.CreateOptions) (result *v1alpha.VPNPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(vpnpeersResource, vPNPeer), &v1alpha.VPNPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.VPNPeer), err
}

// Update takes the representation of a vPNPeer and updates it. Returns the server's representation of the vPNPeer, and an error, if there is any.
func (c *FakeVPNPeers) Update(ctx context.Context, vPNPeer *v1alpha.VPNPeer, opts v1.UpdateOptions) (result *v1alpha.VPNPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(vpnpeersResource, vPNPeer), &v1alpha.VPNPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.VPNPeer), err
}

// Delete takes name of the vPNPeer and deletes it. Returns an error if one occurs.
func (c *FakeVPNPeers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(vpnpeersResource, name), &v1alpha.VPNPeer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVPNPeers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(vpnpeersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha.VPNPeerList{})
	return err
}

// Patch applies the patch and returns the patched vPNPeer.
func (c *FakeVPNPeers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.VPNPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(vpnpeersResource, name, pt, data, subresources...), &v1alpha.VPNPeer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.VPNPeer), err
}
