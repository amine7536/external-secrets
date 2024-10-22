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

// FakePasswords implements PasswordInterface
type FakePasswords struct {
	Fake *FakeGeneratorsV1alpha1
	ns   string
}

var passwordsResource = v1alpha1.SchemeGroupVersion.WithResource("passwords")

var passwordsKind = v1alpha1.SchemeGroupVersion.WithKind("Password")

// Get takes name of the password, and returns the corresponding password object, and an error if there is any.
func (c *FakePasswords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Password, err error) {
	emptyResult := &v1alpha1.Password{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(passwordsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Password), err
}

// List takes label and field selectors, and returns the list of Passwords that match those selectors.
func (c *FakePasswords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PasswordList, err error) {
	emptyResult := &v1alpha1.PasswordList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(passwordsResource, passwordsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PasswordList{ListMeta: obj.(*v1alpha1.PasswordList).ListMeta}
	for _, item := range obj.(*v1alpha1.PasswordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested passwords.
func (c *FakePasswords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(passwordsResource, c.ns, opts))

}

// Create takes the representation of a password and creates it.  Returns the server's representation of the password, and an error, if there is any.
func (c *FakePasswords) Create(ctx context.Context, password *v1alpha1.Password, opts v1.CreateOptions) (result *v1alpha1.Password, err error) {
	emptyResult := &v1alpha1.Password{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(passwordsResource, c.ns, password, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Password), err
}

// Update takes the representation of a password and updates it. Returns the server's representation of the password, and an error, if there is any.
func (c *FakePasswords) Update(ctx context.Context, password *v1alpha1.Password, opts v1.UpdateOptions) (result *v1alpha1.Password, err error) {
	emptyResult := &v1alpha1.Password{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(passwordsResource, c.ns, password, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Password), err
}

// Delete takes name of the password and deletes it. Returns an error if one occurs.
func (c *FakePasswords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(passwordsResource, c.ns, name, opts), &v1alpha1.Password{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePasswords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(passwordsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PasswordList{})
	return err
}

// Patch applies the patch and returns the patched password.
func (c *FakePasswords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Password, err error) {
	emptyResult := &v1alpha1.Password{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(passwordsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Password), err
}