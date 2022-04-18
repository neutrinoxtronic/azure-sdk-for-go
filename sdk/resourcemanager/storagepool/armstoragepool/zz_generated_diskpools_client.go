//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

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

// DiskPoolsClient contains the methods for the DiskPools group.
// Don't use this type directly, use NewDiskPoolsClient() instead.
type DiskPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiskPoolsClient creates a new instance of DiskPoolsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiskPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiskPoolsClient, error) {
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
	client := &DiskPoolsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update Disk pool. This create or update operation can take 15 minutes to complete. This
// is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// diskPoolCreatePayload - Request payload for Disk Pool create operation
// options - DiskPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the DiskPoolsClient.BeginCreateOrUpdate
// method.
func (client *DiskPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolCreatePayload DiskPoolCreate, options *DiskPoolsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DiskPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, diskPoolName, diskPoolCreatePayload, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskPoolsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskPoolsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or Update Disk pool. This create or update operation can take 15 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolCreatePayload DiskPoolCreate, options *DiskPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskPoolName, diskPoolCreatePayload, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiskPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolCreatePayload DiskPoolCreate, options *DiskPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskPoolCreatePayload)
}

// BeginDeallocate - Shuts down the Disk Pool and releases the compute resources. You are not billed for the compute resources
// that this Disk Pool uses. This operation can take 10 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - DiskPoolsClientBeginDeallocateOptions contains the optional parameters for the DiskPoolsClient.BeginDeallocate
// method.
func (client *DiskPoolsClient) BeginDeallocate(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeallocateOptions) (*armruntime.Poller[DiskPoolsClientDeallocateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deallocate(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskPoolsClientDeallocateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskPoolsClientDeallocateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Deallocate - Shuts down the Disk Pool and releases the compute resources. You are not billed for the compute resources
// that this Disk Pool uses. This operation can take 10 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskPoolsClient) deallocate(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeallocateOptions) (*http.Response, error) {
	req, err := client.deallocateCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deallocateCreateRequest creates the Deallocate request.
func (client *DiskPoolsClient) deallocateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeallocateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/deallocate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDelete - Delete a Disk pool; attached disks are not affected. This delete operation can take 10 minutes to complete.
// This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - DiskPoolsClientBeginDeleteOptions contains the optional parameters for the DiskPoolsClient.BeginDelete method.
func (client *DiskPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeleteOptions) (*armruntime.Poller[DiskPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskPoolsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskPoolsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Disk pool; attached disks are not affected. This delete operation can take 10 minutes to complete. This
// is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiskPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - DiskPoolsClientGetOptions contains the optional parameters for the DiskPoolsClient.Get method.
func (client *DiskPoolsClient) Get(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientGetOptions) (DiskPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return DiskPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskPoolsClient) getHandleResponse(resp *http.Response) (DiskPoolsClientGetResponse, error) {
	result := DiskPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPool); err != nil {
		return DiskPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of DiskPools in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DiskPoolsClientListByResourceGroupOptions contains the optional parameters for the DiskPoolsClient.ListByResourceGroup
// method.
func (client *DiskPoolsClient) NewListByResourceGroupPager(resourceGroupName string, options *DiskPoolsClientListByResourceGroupOptions) *runtime.Pager[DiskPoolsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiskPoolsClientListByResourceGroupResponse]{
		More: func(page DiskPoolsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskPoolsClientListByResourceGroupResponse) (DiskPoolsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskPoolsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskPoolsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskPoolsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskPoolsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (DiskPoolsClientListByResourceGroupResponse, error) {
	result := DiskPoolsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPoolListResult); err != nil {
		return DiskPoolsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of Disk Pools in a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - DiskPoolsClientListBySubscriptionOptions contains the optional parameters for the DiskPoolsClient.ListBySubscription
// method.
func (client *DiskPoolsClient) NewListBySubscriptionPager(options *DiskPoolsClientListBySubscriptionOptions) *runtime.Pager[DiskPoolsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiskPoolsClientListBySubscriptionResponse]{
		More: func(page DiskPoolsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskPoolsClientListBySubscriptionResponse) (DiskPoolsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskPoolsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskPoolsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskPoolsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DiskPoolsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DiskPoolsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StoragePool/diskPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DiskPoolsClient) listBySubscriptionHandleResponse(resp *http.Response) (DiskPoolsClientListBySubscriptionResponse, error) {
	result := DiskPoolsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPoolListResult); err != nil {
		return DiskPoolsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListOutboundNetworkDependenciesEndpointsPager - Gets the network endpoints of all outbound dependencies of a Disk Pool
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DiskPoolsClient.ListOutboundNetworkDependenciesEndpoints
// method.
func (client *DiskPoolsClient) NewListOutboundNetworkDependenciesEndpointsPager(resourceGroupName string, diskPoolName string, options *DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions) *runtime.Pager[DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse]{
		More: func(page DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse) (DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOutboundNetworkDependenciesEndpointsCreateRequest(ctx, resourceGroupName, diskPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
		},
	})
}

// listOutboundNetworkDependenciesEndpointsCreateRequest creates the ListOutboundNetworkDependenciesEndpoints request.
func (client *DiskPoolsClient) listOutboundNetworkDependenciesEndpointsCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/outboundNetworkDependenciesEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOutboundNetworkDependenciesEndpointsHandleResponse handles the ListOutboundNetworkDependenciesEndpoints response.
func (client *DiskPoolsClient) listOutboundNetworkDependenciesEndpointsHandleResponse(resp *http.Response) (DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse, error) {
	result := DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointList); err != nil {
		return DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	return result, nil
}

// BeginStart - The operation to start a Disk Pool. This start operation can take 10 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - DiskPoolsClientBeginStartOptions contains the optional parameters for the DiskPoolsClient.BeginStart method.
func (client *DiskPoolsClient) BeginStart(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginStartOptions) (*armruntime.Poller[DiskPoolsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskPoolsClientStartResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskPoolsClientStartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Start - The operation to start a Disk Pool. This start operation can take 10 minutes to complete. This is expected service
// behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskPoolsClient) start(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *DiskPoolsClient) startCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Update a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// diskPoolUpdatePayload - Request payload for Disk Pool update operation.
// options - DiskPoolsClientBeginUpdateOptions contains the optional parameters for the DiskPoolsClient.BeginUpdate method.
func (client *DiskPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolUpdatePayload DiskPoolUpdate, options *DiskPoolsClientBeginUpdateOptions) (*armruntime.Poller[DiskPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, diskPoolName, diskPoolUpdatePayload, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskPoolsClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskPoolsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskPoolsClient) update(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolUpdatePayload DiskPoolUpdate, options *DiskPoolsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskPoolName, diskPoolUpdatePayload, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DiskPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolUpdatePayload DiskPoolUpdate, options *DiskPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskPoolUpdatePayload)
}

// BeginUpgrade - Upgrade replaces the underlying virtual machine hosts one at a time. This operation can take 10-15 minutes
// to complete. This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// diskPoolName - The name of the Disk Pool.
// options - DiskPoolsClientBeginUpgradeOptions contains the optional parameters for the DiskPoolsClient.BeginUpgrade method.
func (client *DiskPoolsClient) BeginUpgrade(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginUpgradeOptions) (*armruntime.Poller[DiskPoolsClientUpgradeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgrade(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[DiskPoolsClientUpgradeResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[DiskPoolsClientUpgradeResponse](options.ResumeToken, client.pl, nil)
	}
}

// Upgrade - Upgrade replaces the underlying virtual machine hosts one at a time. This operation can take 10-15 minutes to
// complete. This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DiskPoolsClient) upgrade(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginUpgradeOptions) (*http.Response, error) {
	req, err := client.upgradeCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// upgradeCreateRequest creates the Upgrade request.
func (client *DiskPoolsClient) upgradeCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/upgrade"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
