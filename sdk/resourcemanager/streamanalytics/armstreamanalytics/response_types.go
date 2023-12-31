//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstreamanalytics

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.BeginCreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	ClusterListResult
}

// ClustersClientListBySubscriptionResponse contains the response from method ClustersClient.NewListBySubscriptionPager.
type ClustersClientListBySubscriptionResponse struct {
	ClusterListResult
}

// ClustersClientListStreamingJobsResponse contains the response from method ClustersClient.NewListStreamingJobsPager.
type ClustersClientListStreamingJobsResponse struct {
	ClusterJobListResult
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.BeginUpdate.
type ClustersClientUpdateResponse struct {
	Cluster
}

// FunctionsClientCreateOrReplaceResponse contains the response from method FunctionsClient.CreateOrReplace.
type FunctionsClientCreateOrReplaceResponse struct {
	Function
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FunctionsClientDeleteResponse contains the response from method FunctionsClient.Delete.
type FunctionsClientDeleteResponse struct {
	// placeholder for future response values
}

// FunctionsClientGetResponse contains the response from method FunctionsClient.Get.
type FunctionsClientGetResponse struct {
	Function
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FunctionsClientListByStreamingJobResponse contains the response from method FunctionsClient.NewListByStreamingJobPager.
type FunctionsClientListByStreamingJobResponse struct {
	FunctionListResult
}

// FunctionsClientRetrieveDefaultDefinitionResponse contains the response from method FunctionsClient.RetrieveDefaultDefinition.
type FunctionsClientRetrieveDefaultDefinitionResponse struct {
	Function
}

// FunctionsClientTestResponse contains the response from method FunctionsClient.BeginTest.
type FunctionsClientTestResponse struct {
	ResourceTestStatus
}

// FunctionsClientUpdateResponse contains the response from method FunctionsClient.Update.
type FunctionsClientUpdateResponse struct {
	Function
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsClientCreateOrReplaceResponse contains the response from method InputsClient.CreateOrReplace.
type InputsClientCreateOrReplaceResponse struct {
	Input
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsClientDeleteResponse contains the response from method InputsClient.Delete.
type InputsClientDeleteResponse struct {
	// placeholder for future response values
}

// InputsClientGetResponse contains the response from method InputsClient.Get.
type InputsClientGetResponse struct {
	Input
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsClientListByStreamingJobResponse contains the response from method InputsClient.NewListByStreamingJobPager.
type InputsClientListByStreamingJobResponse struct {
	InputListResult
}

// InputsClientTestResponse contains the response from method InputsClient.BeginTest.
type InputsClientTestResponse struct {
	ResourceTestStatus
}

// InputsClientUpdateResponse contains the response from method InputsClient.Update.
type InputsClientUpdateResponse struct {
	Input
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// OutputsClientCreateOrReplaceResponse contains the response from method OutputsClient.CreateOrReplace.
type OutputsClientCreateOrReplaceResponse struct {
	Output
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OutputsClientDeleteResponse contains the response from method OutputsClient.Delete.
type OutputsClientDeleteResponse struct {
	// placeholder for future response values
}

// OutputsClientGetResponse contains the response from method OutputsClient.Get.
type OutputsClientGetResponse struct {
	Output
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OutputsClientListByStreamingJobResponse contains the response from method OutputsClient.NewListByStreamingJobPager.
type OutputsClientListByStreamingJobResponse struct {
	OutputListResult
}

// OutputsClientTestResponse contains the response from method OutputsClient.BeginTest.
type OutputsClientTestResponse struct {
	ResourceTestStatus
}

// OutputsClientUpdateResponse contains the response from method OutputsClient.Update.
type OutputsClientUpdateResponse struct {
	Output
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// PrivateEndpointsClientCreateOrUpdateResponse contains the response from method PrivateEndpointsClient.CreateOrUpdate.
type PrivateEndpointsClientCreateOrUpdateResponse struct {
	PrivateEndpoint
}

// PrivateEndpointsClientDeleteResponse contains the response from method PrivateEndpointsClient.BeginDelete.
type PrivateEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointsClientGetResponse contains the response from method PrivateEndpointsClient.Get.
type PrivateEndpointsClientGetResponse struct {
	PrivateEndpoint
}

// PrivateEndpointsClientListByClusterResponse contains the response from method PrivateEndpointsClient.NewListByClusterPager.
type PrivateEndpointsClientListByClusterResponse struct {
	PrivateEndpointListResult
}

// StreamingJobsClientCreateOrReplaceResponse contains the response from method StreamingJobsClient.BeginCreateOrReplace.
type StreamingJobsClientCreateOrReplaceResponse struct {
	StreamingJob
}

// StreamingJobsClientDeleteResponse contains the response from method StreamingJobsClient.BeginDelete.
type StreamingJobsClientDeleteResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientGetResponse contains the response from method StreamingJobsClient.Get.
type StreamingJobsClientGetResponse struct {
	StreamingJob
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// StreamingJobsClientListByResourceGroupResponse contains the response from method StreamingJobsClient.NewListByResourceGroupPager.
type StreamingJobsClientListByResourceGroupResponse struct {
	StreamingJobListResult
}

// StreamingJobsClientListResponse contains the response from method StreamingJobsClient.NewListPager.
type StreamingJobsClientListResponse struct {
	StreamingJobListResult
}

// StreamingJobsClientScaleResponse contains the response from method StreamingJobsClient.BeginScale.
type StreamingJobsClientScaleResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientStartResponse contains the response from method StreamingJobsClient.BeginStart.
type StreamingJobsClientStartResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientStopResponse contains the response from method StreamingJobsClient.BeginStop.
type StreamingJobsClientStopResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientUpdateResponse contains the response from method StreamingJobsClient.Update.
type StreamingJobsClientUpdateResponse struct {
	StreamingJob
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// SubscriptionsClientListQuotasResponse contains the response from method SubscriptionsClient.ListQuotas.
type SubscriptionsClientListQuotasResponse struct {
	SubscriptionQuotasListResult
}

// TransformationsClientCreateOrReplaceResponse contains the response from method TransformationsClient.CreateOrReplace.
type TransformationsClientCreateOrReplaceResponse struct {
	Transformation
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// TransformationsClientGetResponse contains the response from method TransformationsClient.Get.
type TransformationsClientGetResponse struct {
	Transformation
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// TransformationsClientUpdateResponse contains the response from method TransformationsClient.Update.
type TransformationsClientUpdateResponse struct {
	Transformation
	// ETag contains the information returned from the ETag header response.
	ETag *string
}
