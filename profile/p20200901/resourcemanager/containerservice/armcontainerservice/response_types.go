//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

// AgentPoolsClientCreateOrUpdateResponse contains the response from method AgentPoolsClient.BeginCreateOrUpdate.
type AgentPoolsClientCreateOrUpdateResponse struct {
	AgentPool
}

// AgentPoolsClientDeleteResponse contains the response from method AgentPoolsClient.BeginDelete.
type AgentPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientGetAvailableAgentPoolVersionsResponse contains the response from method AgentPoolsClient.GetAvailableAgentPoolVersions.
type AgentPoolsClientGetAvailableAgentPoolVersionsResponse struct {
	AgentPoolAvailableVersions
}

// AgentPoolsClientGetResponse contains the response from method AgentPoolsClient.Get.
type AgentPoolsClientGetResponse struct {
	AgentPool
}

// AgentPoolsClientGetUpgradeProfileResponse contains the response from method AgentPoolsClient.GetUpgradeProfile.
type AgentPoolsClientGetUpgradeProfileResponse struct {
	AgentPoolUpgradeProfile
}

// AgentPoolsClientListResponse contains the response from method AgentPoolsClient.NewListPager.
type AgentPoolsClientListResponse struct {
	AgentPoolListResult
}

// AgentPoolsClientUpgradeNodeImageVersionResponse contains the response from method AgentPoolsClient.BeginUpgradeNodeImageVersion.
type AgentPoolsClientUpgradeNodeImageVersionResponse struct {
	AgentPool
}

// ContainerServicesClientCreateOrUpdateResponse contains the response from method ContainerServicesClient.BeginCreateOrUpdate.
type ContainerServicesClientCreateOrUpdateResponse struct {
	ContainerService
}

// ContainerServicesClientDeleteResponse contains the response from method ContainerServicesClient.BeginDelete.
type ContainerServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerServicesClientGetResponse contains the response from method ContainerServicesClient.Get.
type ContainerServicesClientGetResponse struct {
	ContainerService
}

// ContainerServicesClientListByResourceGroupResponse contains the response from method ContainerServicesClient.NewListByResourceGroupPager.
type ContainerServicesClientListByResourceGroupResponse struct {
	ListResult
}

// ContainerServicesClientListOrchestratorsResponse contains the response from method ContainerServicesClient.ListOrchestrators.
type ContainerServicesClientListOrchestratorsResponse struct {
	OrchestratorVersionProfileListResult
}

// ContainerServicesClientListResponse contains the response from method ContainerServicesClient.NewListPager.
type ContainerServicesClientListResponse struct {
	ListResult
}

// ManagedClustersClientCreateOrUpdateResponse contains the response from method ManagedClustersClient.BeginCreateOrUpdate.
type ManagedClustersClientCreateOrUpdateResponse struct {
	ManagedCluster
}

// ManagedClustersClientDeleteResponse contains the response from method ManagedClustersClient.BeginDelete.
type ManagedClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientGetAccessProfileResponse contains the response from method ManagedClustersClient.GetAccessProfile.
type ManagedClustersClientGetAccessProfileResponse struct {
	ManagedClusterAccessProfile
}

// ManagedClustersClientGetResponse contains the response from method ManagedClustersClient.Get.
type ManagedClustersClientGetResponse struct {
	ManagedCluster
}

// ManagedClustersClientGetUpgradeProfileResponse contains the response from method ManagedClustersClient.GetUpgradeProfile.
type ManagedClustersClientGetUpgradeProfileResponse struct {
	ManagedClusterUpgradeProfile
}

// ManagedClustersClientListByResourceGroupResponse contains the response from method ManagedClustersClient.NewListByResourceGroupPager.
type ManagedClustersClientListByResourceGroupResponse struct {
	ManagedClusterListResult
}

// ManagedClustersClientListClusterAdminCredentialsResponse contains the response from method ManagedClustersClient.ListClusterAdminCredentials.
type ManagedClustersClientListClusterAdminCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListClusterMonitoringUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterMonitoringUserCredentials.
type ManagedClustersClientListClusterMonitoringUserCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListClusterUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterUserCredentials.
type ManagedClustersClientListClusterUserCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListResponse contains the response from method ManagedClustersClient.NewListPager.
type ManagedClustersClientListResponse struct {
	ManagedClusterListResult
}

// ManagedClustersClientResetAADProfileResponse contains the response from method ManagedClustersClient.BeginResetAADProfile.
type ManagedClustersClientResetAADProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientResetServicePrincipalProfileResponse contains the response from method ManagedClustersClient.BeginResetServicePrincipalProfile.
type ManagedClustersClientResetServicePrincipalProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRotateClusterCertificatesResponse contains the response from method ManagedClustersClient.BeginRotateClusterCertificates.
type ManagedClustersClientRotateClusterCertificatesResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientStartResponse contains the response from method ManagedClustersClient.BeginStart.
type ManagedClustersClientStartResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientStopResponse contains the response from method ManagedClustersClient.BeginStop.
type ManagedClustersClientStopResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientUpdateTagsResponse contains the response from method ManagedClustersClient.BeginUpdateTags.
type ManagedClustersClientUpdateTagsResponse struct {
	ManagedCluster
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesListResult
}

// ResolvePrivateLinkServiceIDClientPOSTResponse contains the response from method ResolvePrivateLinkServiceIDClient.POST.
type ResolvePrivateLinkServiceIDClientPOSTResponse struct {
	PrivateLinkResource
}
