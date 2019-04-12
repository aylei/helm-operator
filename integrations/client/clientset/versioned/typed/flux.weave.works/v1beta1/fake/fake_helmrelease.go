/*
Copyright 2018 Weaveworks Ltd.

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
package fake

import (
	v1beta1 "github.com/weaveworks/flux/integrations/apis/flux.weave.works/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHelmReleases implements HelmReleaseInterface
type FakeHelmReleases struct {
	Fake *FakeFluxV1beta1
	ns   string
}

var helmreleasesResource = schema.GroupVersionResource{Group: "flux.weave.works", Version: "v1beta1", Resource: "helmreleases"}

var helmreleasesKind = schema.GroupVersionKind{Group: "flux.weave.works", Version: "v1beta1", Kind: "HelmRelease"}

// Get takes name of the helmRelease, and returns the corresponding helmRelease object, and an error if there is any.
func (c *FakeHelmReleases) Get(name string, options v1.GetOptions) (result *v1beta1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(helmreleasesResource, c.ns, name), &v1beta1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HelmRelease), err
}

// List takes label and field selectors, and returns the list of HelmReleases that match those selectors.
func (c *FakeHelmReleases) List(opts v1.ListOptions) (result *v1beta1.HelmReleaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(helmreleasesResource, helmreleasesKind, c.ns, opts), &v1beta1.HelmReleaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.HelmReleaseList{ListMeta: obj.(*v1beta1.HelmReleaseList).ListMeta}
	for _, item := range obj.(*v1beta1.HelmReleaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helmReleases.
func (c *FakeHelmReleases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(helmreleasesResource, c.ns, opts))

}

// Create takes the representation of a helmRelease and creates it.  Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *FakeHelmReleases) Create(helmRelease *v1beta1.HelmRelease) (result *v1beta1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(helmreleasesResource, c.ns, helmRelease), &v1beta1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HelmRelease), err
}

// Update takes the representation of a helmRelease and updates it. Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *FakeHelmReleases) Update(helmRelease *v1beta1.HelmRelease) (result *v1beta1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(helmreleasesResource, c.ns, helmRelease), &v1beta1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HelmRelease), err
}

// Delete takes name of the helmRelease and deletes it. Returns an error if one occurs.
func (c *FakeHelmReleases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(helmreleasesResource, c.ns, name), &v1beta1.HelmRelease{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelmReleases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(helmreleasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.HelmReleaseList{})
	return err
}

// Patch applies the patch and returns the patched helmRelease.
func (c *FakeHelmReleases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(helmreleasesResource, c.ns, name, types.ApplyPatchType, data, subresources...), &v1beta1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.HelmRelease), err
}
