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

package v1beta1

import (
	"context"

	v1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	scheme "github.com/external-secrets/external-secrets/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SecretStoresGetter has a method to return a SecretStoreInterface.
// A group's client should implement this interface.
type SecretStoresGetter interface {
	SecretStores(namespace string) SecretStoreInterface
}

// SecretStoreInterface has methods to work with SecretStore resources.
type SecretStoreInterface interface {
	Create(ctx context.Context, secretStore *v1beta1.SecretStore, opts v1.CreateOptions) (*v1beta1.SecretStore, error)
	Update(ctx context.Context, secretStore *v1beta1.SecretStore, opts v1.UpdateOptions) (*v1beta1.SecretStore, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, secretStore *v1beta1.SecretStore, opts v1.UpdateOptions) (*v1beta1.SecretStore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.SecretStore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SecretStoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SecretStore, err error)
	SecretStoreExpansion
}

// secretStores implements SecretStoreInterface
type secretStores struct {
	*gentype.ClientWithList[*v1beta1.SecretStore, *v1beta1.SecretStoreList]
}

// newSecretStores returns a SecretStores
func newSecretStores(c *ExternalSecretsV1beta1Client, namespace string) *secretStores {
	return &secretStores{
		gentype.NewClientWithList[*v1beta1.SecretStore, *v1beta1.SecretStoreList](
			"secretstores",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.SecretStore { return &v1beta1.SecretStore{} },
			func() *v1beta1.SecretStoreList { return &v1beta1.SecretStoreList{} }),
	}
}