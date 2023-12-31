//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysqlflexibleservers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureADAdministratorsClient contains the methods for the AzureADAdministrators group.
// Don't use this type directly, use NewAzureADAdministratorsClient() instead.
type AzureADAdministratorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureADAdministratorsClient creates a new instance of AzureADAdministratorsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureADAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureADAdministratorsClient, error) {
	cl, err := arm.NewClient(moduleName+".AzureADAdministratorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureADAdministratorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an existing Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - administratorName - The name of the Azure AD Administrator.
//   - parameters - The required parameters for creating or updating an aad administrator.
//   - options - AzureADAdministratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureADAdministratorsClient.BeginCreateOrUpdate
//     method.
func (client *AzureADAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters AzureADAdministrator, options *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureADAdministratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, administratorName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AzureADAdministratorsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AzureADAdministratorsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an existing Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
func (client *AzureADAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters AzureADAdministrator, options *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, administratorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AzureADAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters AzureADAdministrator, options *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an Azure AD Administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - administratorName - The name of the Azure AD Administrator.
//   - options - AzureADAdministratorsClientBeginDeleteOptions contains the optional parameters for the AzureADAdministratorsClient.BeginDelete
//     method.
func (client *AzureADAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *AzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[AzureADAdministratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, administratorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureADAdministratorsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AzureADAdministratorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an Azure AD Administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
func (client *AzureADAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *AzureADAdministratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, administratorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AzureADAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *AzureADAdministratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about an azure ad administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - administratorName - The name of the Azure AD Administrator.
//   - options - AzureADAdministratorsClientGetOptions contains the optional parameters for the AzureADAdministratorsClient.Get
//     method.
func (client *AzureADAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *AzureADAdministratorsClientGetOptions) (AzureADAdministratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, administratorName, options)
	if err != nil {
		return AzureADAdministratorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureADAdministratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureADAdministratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureADAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *AzureADAdministratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureADAdministratorsClient) getHandleResponse(resp *http.Response) (AzureADAdministratorsClientGetResponse, error) {
	result := AzureADAdministratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureADAdministrator); err != nil {
		return AzureADAdministratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - List all the AAD administrators in a given server.
//
// Generated from API version 2021-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - AzureADAdministratorsClientListByServerOptions contains the optional parameters for the AzureADAdministratorsClient.NewListByServerPager
//     method.
func (client *AzureADAdministratorsClient) NewListByServerPager(resourceGroupName string, serverName string, options *AzureADAdministratorsClientListByServerOptions) *runtime.Pager[AzureADAdministratorsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureADAdministratorsClientListByServerResponse]{
		More: func(page AzureADAdministratorsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureADAdministratorsClientListByServerResponse) (AzureADAdministratorsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureADAdministratorsClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AzureADAdministratorsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureADAdministratorsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *AzureADAdministratorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *AzureADAdministratorsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *AzureADAdministratorsClient) listByServerHandleResponse(resp *http.Response) (AzureADAdministratorsClientListByServerResponse, error) {
	result := AzureADAdministratorsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdministratorListResult); err != nil {
		return AzureADAdministratorsClientListByServerResponse{}, err
	}
	return result, nil
}
