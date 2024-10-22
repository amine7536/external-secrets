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

// WebhooksGetter has a method to return a WebhookInterface.
// A group's client should implement this interface.
type WebhooksGetter interface {
	Webhooks(namespace string) WebhookInterface
}

// WebhookInterface has methods to work with Webhook resources.
type WebhookInterface interface {
	Create(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.CreateOptions) (*v1alpha1.Webhook, error)
	Update(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.UpdateOptions) (*v1alpha1.Webhook, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Webhook, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WebhookList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Webhook, err error)
	WebhookExpansion
}

// webhooks implements WebhookInterface
type webhooks struct {
	*gentype.ClientWithList[*v1alpha1.Webhook, *v1alpha1.WebhookList]
}

// newWebhooks returns a Webhooks
func newWebhooks(c *GeneratorsV1alpha1Client, namespace string) *webhooks {
	return &webhooks{
		gentype.NewClientWithList[*v1alpha1.Webhook, *v1alpha1.WebhookList](
			"webhooks",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Webhook { return &v1alpha1.Webhook{} },
			func() *v1alpha1.WebhookList { return &v1alpha1.WebhookList{} }),
	}
}