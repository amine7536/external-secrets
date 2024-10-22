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

	v1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExternalSecrets implements ExternalSecretInterface
type FakeExternalSecrets struct {
	Fake *FakeExternalSecretsV1beta1
	ns   string
}

var externalsecretsResource = v1beta1.SchemeGroupVersion.WithResource("externalsecrets")

var externalsecretsKind = v1beta1.SchemeGroupVersion.WithKind("ExternalSecret")

// Get takes name of the externalSecret, and returns the corresponding externalSecret object, and an error if there is any.
func (c *FakeExternalSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ExternalSecret, err error) {
	emptyResult := &v1beta1.ExternalSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(externalsecretsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ExternalSecret), err
}

// List takes label and field selectors, and returns the list of ExternalSecrets that match those selectors.
func (c *FakeExternalSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ExternalSecretList, err error) {
	emptyResult := &v1beta1.ExternalSecretList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(externalsecretsResource, externalsecretsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ExternalSecretList{ListMeta: obj.(*v1beta1.ExternalSecretList).ListMeta}
	for _, item := range obj.(*v1beta1.ExternalSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalSecrets.
func (c *FakeExternalSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(externalsecretsResource, c.ns, opts))

}

// Create takes the representation of a externalSecret and creates it.  Returns the server's representation of the externalSecret, and an error, if there is any.
func (c *FakeExternalSecrets) Create(ctx context.Context, externalSecret *v1beta1.ExternalSecret, opts v1.CreateOptions) (result *v1beta1.ExternalSecret, err error) {
	emptyResult := &v1beta1.ExternalSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(externalsecretsResource, c.ns, externalSecret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ExternalSecret), err
}

// Update takes the representation of a externalSecret and updates it. Returns the server's representation of the externalSecret, and an error, if there is any.
func (c *FakeExternalSecrets) Update(ctx context.Context, externalSecret *v1beta1.ExternalSecret, opts v1.UpdateOptions) (result *v1beta1.ExternalSecret, err error) {
	emptyResult := &v1beta1.ExternalSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(externalsecretsResource, c.ns, externalSecret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ExternalSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExternalSecrets) UpdateStatus(ctx context.Context, externalSecret *v1beta1.ExternalSecret, opts v1.UpdateOptions) (result *v1beta1.ExternalSecret, err error) {
	emptyResult := &v1beta1.ExternalSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(externalsecretsResource, "status", c.ns, externalSecret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ExternalSecret), err
}

// Delete takes name of the externalSecret and deletes it. Returns an error if one occurs.
func (c *FakeExternalSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(externalsecretsResource, c.ns, name, opts), &v1beta1.ExternalSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(externalsecretsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ExternalSecretList{})
	return err
}

// Patch applies the patch and returns the patched externalSecret.
func (c *FakeExternalSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ExternalSecret, err error) {
	emptyResult := &v1beta1.ExternalSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(externalsecretsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.ExternalSecret), err
}