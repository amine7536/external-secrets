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

// FakeVaultDynamicSecrets implements VaultDynamicSecretInterface
type FakeVaultDynamicSecrets struct {
	Fake *FakeGeneratorsV1alpha1
	ns   string
}

var vaultdynamicsecretsResource = v1alpha1.SchemeGroupVersion.WithResource("vaultdynamicsecrets")

var vaultdynamicsecretsKind = v1alpha1.SchemeGroupVersion.WithKind("VaultDynamicSecret")

// Get takes name of the vaultDynamicSecret, and returns the corresponding vaultDynamicSecret object, and an error if there is any.
func (c *FakeVaultDynamicSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VaultDynamicSecret, err error) {
	emptyResult := &v1alpha1.VaultDynamicSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(vaultdynamicsecretsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VaultDynamicSecret), err
}

// List takes label and field selectors, and returns the list of VaultDynamicSecrets that match those selectors.
func (c *FakeVaultDynamicSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VaultDynamicSecretList, err error) {
	emptyResult := &v1alpha1.VaultDynamicSecretList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(vaultdynamicsecretsResource, vaultdynamicsecretsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultDynamicSecretList{ListMeta: obj.(*v1alpha1.VaultDynamicSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.VaultDynamicSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vaultDynamicSecrets.
func (c *FakeVaultDynamicSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(vaultdynamicsecretsResource, c.ns, opts))

}

// Create takes the representation of a vaultDynamicSecret and creates it.  Returns the server's representation of the vaultDynamicSecret, and an error, if there is any.
func (c *FakeVaultDynamicSecrets) Create(ctx context.Context, vaultDynamicSecret *v1alpha1.VaultDynamicSecret, opts v1.CreateOptions) (result *v1alpha1.VaultDynamicSecret, err error) {
	emptyResult := &v1alpha1.VaultDynamicSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(vaultdynamicsecretsResource, c.ns, vaultDynamicSecret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VaultDynamicSecret), err
}

// Update takes the representation of a vaultDynamicSecret and updates it. Returns the server's representation of the vaultDynamicSecret, and an error, if there is any.
func (c *FakeVaultDynamicSecrets) Update(ctx context.Context, vaultDynamicSecret *v1alpha1.VaultDynamicSecret, opts v1.UpdateOptions) (result *v1alpha1.VaultDynamicSecret, err error) {
	emptyResult := &v1alpha1.VaultDynamicSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(vaultdynamicsecretsResource, c.ns, vaultDynamicSecret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VaultDynamicSecret), err
}

// Delete takes name of the vaultDynamicSecret and deletes it. Returns an error if one occurs.
func (c *FakeVaultDynamicSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vaultdynamicsecretsResource, c.ns, name, opts), &v1alpha1.VaultDynamicSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultDynamicSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(vaultdynamicsecretsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultDynamicSecretList{})
	return err
}

// Patch applies the patch and returns the patched vaultDynamicSecret.
func (c *FakeVaultDynamicSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VaultDynamicSecret, err error) {
	emptyResult := &v1alpha1.VaultDynamicSecret{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(vaultdynamicsecretsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VaultDynamicSecret), err
}