//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SupportedOperatingSystemsClient contains the methods for the SupportedOperatingSystems group.
// Don't use this type directly, use NewSupportedOperatingSystemsClient() instead.
type SupportedOperatingSystemsClient struct {
	ep                string
	pl                runtime.Pipeline
	resourceName      string
	resourceGroupName string
	subscriptionID    string
}

// NewSupportedOperatingSystemsClient creates a new instance of SupportedOperatingSystemsClient with the specified values.
func NewSupportedOperatingSystemsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SupportedOperatingSystemsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SupportedOperatingSystemsClient{resourceName: resourceName, resourceGroupName: resourceGroupName, subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets the data of supported operating systems by SRS.
// If the operation fails it returns a generic error.
func (client *SupportedOperatingSystemsClient) Get(ctx context.Context, options *SupportedOperatingSystemsGetOptions) (SupportedOperatingSystemsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return SupportedOperatingSystemsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SupportedOperatingSystemsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SupportedOperatingSystemsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SupportedOperatingSystemsClient) getCreateRequest(ctx context.Context, options *SupportedOperatingSystemsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationSupportedOperatingSystems"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	if options != nil && options.InstanceType != nil {
		reqQP.Set("instanceType", *options.InstanceType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SupportedOperatingSystemsClient) getHandleResponse(resp *http.Response) (SupportedOperatingSystemsGetResponse, error) {
	result := SupportedOperatingSystemsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SupportedOperatingSystems); err != nil {
		return SupportedOperatingSystemsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SupportedOperatingSystemsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}