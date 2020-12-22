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

	mycrdv1 "github.com/wwq-2020/crd-codegen-demo/apis/mycrd/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDemos implements DemoInterface
type FakeDemos struct {
	Fake *FakeMycrdV1
	ns   string
}

var demosResource = schema.GroupVersionResource{Group: "mycrd", Version: "v1", Resource: "demos"}

var demosKind = schema.GroupVersionKind{Group: "mycrd", Version: "v1", Kind: "Demo"}

// Get takes name of the demo, and returns the corresponding demo object, and an error if there is any.
func (c *FakeDemos) Get(ctx context.Context, name string, options v1.GetOptions) (result *mycrdv1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(demosResource, c.ns, name), &mycrdv1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mycrdv1.Demo), err
}

// List takes label and field selectors, and returns the list of Demos that match those selectors.
func (c *FakeDemos) List(ctx context.Context, opts v1.ListOptions) (result *mycrdv1.DemoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(demosResource, demosKind, c.ns, opts), &mycrdv1.DemoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mycrdv1.DemoList{ListMeta: obj.(*mycrdv1.DemoList).ListMeta}
	for _, item := range obj.(*mycrdv1.DemoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested demos.
func (c *FakeDemos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(demosResource, c.ns, opts))

}

// Create takes the representation of a demo and creates it.  Returns the server's representation of the demo, and an error, if there is any.
func (c *FakeDemos) Create(ctx context.Context, demo *mycrdv1.Demo, opts v1.CreateOptions) (result *mycrdv1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(demosResource, c.ns, demo), &mycrdv1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mycrdv1.Demo), err
}

// Update takes the representation of a demo and updates it. Returns the server's representation of the demo, and an error, if there is any.
func (c *FakeDemos) Update(ctx context.Context, demo *mycrdv1.Demo, opts v1.UpdateOptions) (result *mycrdv1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(demosResource, c.ns, demo), &mycrdv1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mycrdv1.Demo), err
}

// Delete takes name of the demo and deletes it. Returns an error if one occurs.
func (c *FakeDemos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(demosResource, c.ns, name), &mycrdv1.Demo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDemos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(demosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &mycrdv1.DemoList{})
	return err
}

// Patch applies the patch and returns the patched demo.
func (c *FakeDemos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *mycrdv1.Demo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(demosResource, c.ns, name, pt, data, subresources...), &mycrdv1.Demo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mycrdv1.Demo), err
}