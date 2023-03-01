//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

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

// DatabaseClient contains the methods for the Database group.
// Don't use this type directly, use NewDatabaseClient() instead.
type DatabaseClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDatabaseClient creates a new instance of DatabaseClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListMetricDefinitionsPager - Retrieves metric definitions for the given database.
//
// Generated from API version 2022-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - databaseRid - Cosmos DB database rid.
//   - options - DatabaseClientListMetricDefinitionsOptions contains the optional parameters for the DatabaseClient.NewListMetricDefinitionsPager
//     method.
func (client *DatabaseClient) NewListMetricDefinitionsPager(resourceGroupName string, accountName string, databaseRid string, options *DatabaseClientListMetricDefinitionsOptions) *runtime.Pager[DatabaseClientListMetricDefinitionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseClientListMetricDefinitionsResponse]{
		More: func(page DatabaseClientListMetricDefinitionsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DatabaseClientListMetricDefinitionsResponse) (DatabaseClientListMetricDefinitionsResponse, error) {
			req, err := client.listMetricDefinitionsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, options)
			if err != nil {
				return DatabaseClientListMetricDefinitionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DatabaseClientListMetricDefinitionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabaseClientListMetricDefinitionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricDefinitionsHandleResponse(resp)
		},
	})
}

// listMetricDefinitionsCreateRequest creates the ListMetricDefinitions request.
func (client *DatabaseClient) listMetricDefinitionsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, options *DatabaseClientListMetricDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metricDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricDefinitionsHandleResponse handles the ListMetricDefinitions response.
func (client *DatabaseClient) listMetricDefinitionsHandleResponse(resp *http.Response) (DatabaseClientListMetricDefinitionsResponse, error) {
	result := DatabaseClientListMetricDefinitionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionsListResult); err != nil {
		return DatabaseClientListMetricDefinitionsResponse{}, err
	}
	return result, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given database account and database.
//
// Generated from API version 2022-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - databaseRid - Cosmos DB database rid.
//   - filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
//     name.value (name of the metric, can have an or of multiple names), startTime, endTime,
//     and timeGrain. The supported operator is eq.
//   - options - DatabaseClientListMetricsOptions contains the optional parameters for the DatabaseClient.NewListMetricsPager
//     method.
func (client *DatabaseClient) NewListMetricsPager(resourceGroupName string, accountName string, databaseRid string, filter string, options *DatabaseClientListMetricsOptions) *runtime.Pager[DatabaseClientListMetricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseClientListMetricsResponse]{
		More: func(page DatabaseClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DatabaseClientListMetricsResponse) (DatabaseClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, filter, options)
			if err != nil {
				return DatabaseClientListMetricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DatabaseClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabaseClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *DatabaseClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string, options *DatabaseClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *DatabaseClient) listMetricsHandleResponse(resp *http.Response) (DatabaseClientListMetricsResponse, error) {
	result := DatabaseClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return DatabaseClientListMetricsResponse{}, err
	}
	return result, nil
}

// NewListUsagesPager - Retrieves the usages (most recent data) for the given database.
//
// Generated from API version 2022-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - databaseRid - Cosmos DB database rid.
//   - options - DatabaseClientListUsagesOptions contains the optional parameters for the DatabaseClient.NewListUsagesPager method.
func (client *DatabaseClient) NewListUsagesPager(resourceGroupName string, accountName string, databaseRid string, options *DatabaseClientListUsagesOptions) *runtime.Pager[DatabaseClientListUsagesResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseClientListUsagesResponse]{
		More: func(page DatabaseClientListUsagesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DatabaseClientListUsagesResponse) (DatabaseClientListUsagesResponse, error) {
			req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, options)
			if err != nil {
				return DatabaseClientListUsagesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DatabaseClientListUsagesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabaseClientListUsagesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listUsagesHandleResponse(resp)
		},
	})
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *DatabaseClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, options *DatabaseClientListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *DatabaseClient) listUsagesHandleResponse(resp *http.Response) (DatabaseClientListUsagesResponse, error) {
	result := DatabaseClientListUsagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesResult); err != nil {
		return DatabaseClientListUsagesResponse{}, err
	}
	return result, nil
}
