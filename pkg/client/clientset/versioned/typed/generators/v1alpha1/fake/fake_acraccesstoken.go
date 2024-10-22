/*
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

	v1alpha1 "github.com/external-secrets/external-secrets/apis/generators/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeACRAccessTokens implements ACRAccessTokenInterface
type FakeACRAccessTokens struct {
	Fake *FakeGeneratorsV1alpha1
	ns   string
}

var acraccesstokensResource = v1alpha1.SchemeGroupVersion.WithResource("acraccesstokens")

var acraccesstokensKind = v1alpha1.SchemeGroupVersion.WithKind("ACRAccessToken")

// Get takes name of the aCRAccessToken, and returns the corresponding aCRAccessToken object, and an error if there is any.
func (c *FakeACRAccessTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ACRAccessToken, err error) {
	emptyResult := &v1alpha1.ACRAccessToken{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(acraccesstokensResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ACRAccessToken), err
}

// List takes label and field selectors, and returns the list of ACRAccessTokens that match those selectors.
func (c *FakeACRAccessTokens) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ACRAccessTokenList, err error) {
	emptyResult := &v1alpha1.ACRAccessTokenList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(acraccesstokensResource, acraccesstokensKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ACRAccessTokenList{ListMeta: obj.(*v1alpha1.ACRAccessTokenList).ListMeta}
	for _, item := range obj.(*v1alpha1.ACRAccessTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aCRAccessTokens.
func (c *FakeACRAccessTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(acraccesstokensResource, c.ns, opts))

}

// Create takes the representation of a aCRAccessToken and creates it.  Returns the server's representation of the aCRAccessToken, and an error, if there is any.
func (c *FakeACRAccessTokens) Create(ctx context.Context, aCRAccessToken *v1alpha1.ACRAccessToken, opts v1.CreateOptions) (result *v1alpha1.ACRAccessToken, err error) {
	emptyResult := &v1alpha1.ACRAccessToken{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(acraccesstokensResource, c.ns, aCRAccessToken, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ACRAccessToken), err
}

// Update takes the representation of a aCRAccessToken and updates it. Returns the server's representation of the aCRAccessToken, and an error, if there is any.
func (c *FakeACRAccessTokens) Update(ctx context.Context, aCRAccessToken *v1alpha1.ACRAccessToken, opts v1.UpdateOptions) (result *v1alpha1.ACRAccessToken, err error) {
	emptyResult := &v1alpha1.ACRAccessToken{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(acraccesstokensResource, c.ns, aCRAccessToken, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ACRAccessToken), err
}

// Delete takes name of the aCRAccessToken and deletes it. Returns an error if one occurs.
func (c *FakeACRAccessTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(acraccesstokensResource, c.ns, name, opts), &v1alpha1.ACRAccessToken{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeACRAccessTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(acraccesstokensResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ACRAccessTokenList{})
	return err
}

// Patch applies the patch and returns the patched aCRAccessToken.
func (c *FakeACRAccessTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ACRAccessToken, err error) {
	emptyResult := &v1alpha1.ACRAccessToken{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(acraccesstokensResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ACRAccessToken), err
}