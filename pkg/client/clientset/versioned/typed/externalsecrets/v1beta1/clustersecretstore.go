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

// ClusterSecretStoresGetter has a method to return a ClusterSecretStoreInterface.
// A group's client should implement this interface.
type ClusterSecretStoresGetter interface {
	ClusterSecretStores(namespace string) ClusterSecretStoreInterface
}

// ClusterSecretStoreInterface has methods to work with ClusterSecretStore resources.
type ClusterSecretStoreInterface interface {
	Create(ctx context.Context, clusterSecretStore *v1beta1.ClusterSecretStore, opts v1.CreateOptions) (*v1beta1.ClusterSecretStore, error)
	Update(ctx context.Context, clusterSecretStore *v1beta1.ClusterSecretStore, opts v1.UpdateOptions) (*v1beta1.ClusterSecretStore, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterSecretStore *v1beta1.ClusterSecretStore, opts v1.UpdateOptions) (*v1beta1.ClusterSecretStore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ClusterSecretStore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterSecretStoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterSecretStore, err error)
	ClusterSecretStoreExpansion
}

// clusterSecretStores implements ClusterSecretStoreInterface
type clusterSecretStores struct {
	*gentype.ClientWithList[*v1beta1.ClusterSecretStore, *v1beta1.ClusterSecretStoreList]
}

// newClusterSecretStores returns a ClusterSecretStores
func newClusterSecretStores(c *ExternalSecretsV1beta1Client, namespace string) *clusterSecretStores {
	return &clusterSecretStores{
		gentype.NewClientWithList[*v1beta1.ClusterSecretStore, *v1beta1.ClusterSecretStoreList](
			"clustersecretstores",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.ClusterSecretStore { return &v1beta1.ClusterSecretStore{} },
			func() *v1beta1.ClusterSecretStoreList { return &v1beta1.ClusterSecretStoreList{} }),
	}
}
