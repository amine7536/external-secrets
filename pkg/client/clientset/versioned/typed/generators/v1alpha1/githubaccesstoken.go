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

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/external-secrets/external-secrets/apis/generators/v1alpha1"
	scheme "github.com/external-secrets/external-secrets/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// GithubAccessTokensGetter has a method to return a GithubAccessTokenInterface.
// A group's client should implement this interface.
type GithubAccessTokensGetter interface {
	GithubAccessTokens(namespace string) GithubAccessTokenInterface
}

// GithubAccessTokenInterface has methods to work with GithubAccessToken resources.
type GithubAccessTokenInterface interface {
	Create(ctx context.Context, githubAccessToken *v1alpha1.GithubAccessToken, opts v1.CreateOptions) (*v1alpha1.GithubAccessToken, error)
	Update(ctx context.Context, githubAccessToken *v1alpha1.GithubAccessToken, opts v1.UpdateOptions) (*v1alpha1.GithubAccessToken, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GithubAccessToken, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GithubAccessTokenList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GithubAccessToken, err error)
	GithubAccessTokenExpansion
}

// githubAccessTokens implements GithubAccessTokenInterface
type githubAccessTokens struct {
	*gentype.ClientWithList[*v1alpha1.GithubAccessToken, *v1alpha1.GithubAccessTokenList]
}

// newGithubAccessTokens returns a GithubAccessTokens
func newGithubAccessTokens(c *GeneratorsV1alpha1Client, namespace string) *githubAccessTokens {
	return &githubAccessTokens{
		gentype.NewClientWithList[*v1alpha1.GithubAccessToken, *v1alpha1.GithubAccessTokenList](
			"githubaccesstokens",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.GithubAccessToken { return &v1alpha1.GithubAccessToken{} },
			func() *v1alpha1.GithubAccessTokenList { return &v1alpha1.GithubAccessTokenList{} }),
	}
}