//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedRuleSetsClient contains the methods for the ManagedRuleSets group.
// Don't use this type directly, use NewManagedRuleSetsClient() instead.
type ManagedRuleSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedRuleSetsClient creates a new instance of ManagedRuleSetsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedRuleSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedRuleSetsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedRuleSetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Lists all available managed rule sets.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ManagedRuleSetsClientListOptions contains the optional parameters for the ManagedRuleSetsClient.List method.
func (client *ManagedRuleSetsClient) List(options *ManagedRuleSetsClientListOptions) *runtime.Pager[ManagedRuleSetsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ManagedRuleSetsClientListResponse]{
		More: func(page ManagedRuleSetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedRuleSetsClientListResponse) (ManagedRuleSetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedRuleSetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedRuleSetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedRuleSetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ManagedRuleSetsClient) listCreateRequest(ctx context.Context, options *ManagedRuleSetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedRuleSetsClient) listHandleResponse(resp *http.Response) (ManagedRuleSetsClientListResponse, error) {
	result := ManagedRuleSetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedRuleSetDefinitionList); err != nil {
		return ManagedRuleSetsClientListResponse{}, err
	}
	return result, nil
}
