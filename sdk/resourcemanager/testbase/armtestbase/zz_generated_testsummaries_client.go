//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// TestSummariesClient contains the methods for the TestSummaries group.
// Don't use this type directly, use NewTestSummariesClient() instead.
type TestSummariesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTestSummariesClient creates a new instance of TestSummariesClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTestSummariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TestSummariesClient, error) {
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
	client := &TestSummariesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a Test Summary with specific name from all the Test Summaries of all the packages under a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// testSummaryName - The name of the Test Summary.
// options - TestSummariesClientGetOptions contains the optional parameters for the TestSummariesClient.Get method.
func (client *TestSummariesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, testSummaryName string, options *TestSummariesClientGetOptions) (TestSummariesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, testSummaryName, options)
	if err != nil {
		return TestSummariesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestSummariesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestSummariesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TestSummariesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, testSummaryName string, options *TestSummariesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/testSummaries/{testSummaryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if testSummaryName == "" {
		return nil, errors.New("parameter testSummaryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testSummaryName}", url.PathEscape(testSummaryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TestSummariesClient) getHandleResponse(resp *http.Response) (TestSummariesClientGetResponse, error) {
	result := TestSummariesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestSummaryResource); err != nil {
		return TestSummariesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the Test Summaries of all the packages under a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - TestSummariesClientListOptions contains the optional parameters for the TestSummariesClient.List method.
func (client *TestSummariesClient) NewListPager(resourceGroupName string, testBaseAccountName string, options *TestSummariesClientListOptions) *runtime.Pager[TestSummariesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TestSummariesClientListResponse]{
		More: func(page TestSummariesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TestSummariesClientListResponse) (TestSummariesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TestSummariesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TestSummariesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TestSummariesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TestSummariesClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *TestSummariesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/testSummaries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TestSummariesClient) listHandleResponse(resp *http.Response) (TestSummariesClientListResponse, error) {
	result := TestSummariesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestSummaryListResult); err != nil {
		return TestSummariesClientListResponse{}, err
	}
	return result, nil
}
