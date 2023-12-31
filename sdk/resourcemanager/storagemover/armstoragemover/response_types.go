//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover

// AgentsClientCreateOrUpdateResponse contains the response from method AgentsClient.CreateOrUpdate.
type AgentsClientCreateOrUpdateResponse struct {
	Agent
}

// AgentsClientDeleteResponse contains the response from method AgentsClient.BeginDelete.
type AgentsClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentsClientGetResponse contains the response from method AgentsClient.Get.
type AgentsClientGetResponse struct {
	Agent
}

// AgentsClientListResponse contains the response from method AgentsClient.NewListPager.
type AgentsClientListResponse struct {
	AgentList
}

// AgentsClientUpdateResponse contains the response from method AgentsClient.Update.
type AgentsClientUpdateResponse struct {
	Agent
}

// EndpointsClientCreateOrUpdateResponse contains the response from method EndpointsClient.CreateOrUpdate.
type EndpointsClientCreateOrUpdateResponse struct {
	Endpoint
}

// EndpointsClientDeleteResponse contains the response from method EndpointsClient.BeginDelete.
type EndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// EndpointsClientGetResponse contains the response from method EndpointsClient.Get.
type EndpointsClientGetResponse struct {
	Endpoint
}

// EndpointsClientListResponse contains the response from method EndpointsClient.NewListPager.
type EndpointsClientListResponse struct {
	EndpointList
}

// EndpointsClientUpdateResponse contains the response from method EndpointsClient.Update.
type EndpointsClientUpdateResponse struct {
	Endpoint
}

// JobDefinitionsClientCreateOrUpdateResponse contains the response from method JobDefinitionsClient.CreateOrUpdate.
type JobDefinitionsClientCreateOrUpdateResponse struct {
	JobDefinition
}

// JobDefinitionsClientDeleteResponse contains the response from method JobDefinitionsClient.BeginDelete.
type JobDefinitionsClientDeleteResponse struct {
	// placeholder for future response values
}

// JobDefinitionsClientGetResponse contains the response from method JobDefinitionsClient.Get.
type JobDefinitionsClientGetResponse struct {
	JobDefinition
}

// JobDefinitionsClientListResponse contains the response from method JobDefinitionsClient.NewListPager.
type JobDefinitionsClientListResponse struct {
	JobDefinitionList
}

// JobDefinitionsClientStartJobResponse contains the response from method JobDefinitionsClient.StartJob.
type JobDefinitionsClientStartJobResponse struct {
	JobRunResourceID
}

// JobDefinitionsClientStopJobResponse contains the response from method JobDefinitionsClient.StopJob.
type JobDefinitionsClientStopJobResponse struct {
	JobRunResourceID
}

// JobDefinitionsClientUpdateResponse contains the response from method JobDefinitionsClient.Update.
type JobDefinitionsClientUpdateResponse struct {
	JobDefinition
}

// JobRunsClientGetResponse contains the response from method JobRunsClient.Get.
type JobRunsClientGetResponse struct {
	JobRun
}

// JobRunsClientListResponse contains the response from method JobRunsClient.NewListPager.
type JobRunsClientListResponse struct {
	JobRunList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// ProjectsClientCreateOrUpdateResponse contains the response from method ProjectsClient.CreateOrUpdate.
type ProjectsClientCreateOrUpdateResponse struct {
	Project
}

// ProjectsClientDeleteResponse contains the response from method ProjectsClient.BeginDelete.
type ProjectsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProjectsClientGetResponse contains the response from method ProjectsClient.Get.
type ProjectsClientGetResponse struct {
	Project
}

// ProjectsClientListResponse contains the response from method ProjectsClient.NewListPager.
type ProjectsClientListResponse struct {
	ProjectList
}

// ProjectsClientUpdateResponse contains the response from method ProjectsClient.Update.
type ProjectsClientUpdateResponse struct {
	Project
}

// StorageMoversClientCreateOrUpdateResponse contains the response from method StorageMoversClient.CreateOrUpdate.
type StorageMoversClientCreateOrUpdateResponse struct {
	StorageMover
}

// StorageMoversClientDeleteResponse contains the response from method StorageMoversClient.BeginDelete.
type StorageMoversClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageMoversClientGetResponse contains the response from method StorageMoversClient.Get.
type StorageMoversClientGetResponse struct {
	StorageMover
}

// StorageMoversClientListBySubscriptionResponse contains the response from method StorageMoversClient.NewListBySubscriptionPager.
type StorageMoversClientListBySubscriptionResponse struct {
	List
}

// StorageMoversClientListResponse contains the response from method StorageMoversClient.NewListPager.
type StorageMoversClientListResponse struct {
	List
}

// StorageMoversClientUpdateResponse contains the response from method StorageMoversClient.Update.
type StorageMoversClientUpdateResponse struct {
	StorageMover
}
