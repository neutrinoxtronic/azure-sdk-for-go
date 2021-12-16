//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ComputeCreateOrUpdatePollerResponse contains the response from method Compute.CreateOrUpdate.
type ComputeCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ComputeCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ComputeCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ComputeCreateOrUpdateResponse, error) {
	respType := ComputeCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ComputeResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ComputeCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ComputeCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ComputeClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ComputeClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &ComputeCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ComputeCreateOrUpdateResponse contains the response from method Compute.CreateOrUpdate.
type ComputeCreateOrUpdateResponse struct {
	ComputeCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeCreateOrUpdateResult contains the result from method Compute.CreateOrUpdate.
type ComputeCreateOrUpdateResult struct {
	ComputeResource
}

// ComputeDeletePollerResponse contains the response from method Compute.Delete.
type ComputeDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ComputeDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ComputeDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ComputeDeleteResponse, error) {
	respType := ComputeDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ComputeDeletePollerResponse from the provided client and resume token.
func (l *ComputeDeletePollerResponse) Resume(ctx context.Context, client *ComputeClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ComputeClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ComputeDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ComputeDeleteResponse contains the response from method Compute.Delete.
type ComputeDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeGetResponse contains the response from method Compute.Get.
type ComputeGetResponse struct {
	ComputeGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeGetResult contains the result from method Compute.Get.
type ComputeGetResult struct {
	ComputeResource
}

// ComputeListKeysResponse contains the response from method Compute.ListKeys.
type ComputeListKeysResponse struct {
	ComputeListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeListKeysResult contains the result from method Compute.ListKeys.
type ComputeListKeysResult struct {
	ComputeSecretsClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComputeListKeysResult.
func (c *ComputeListKeysResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalComputeSecretsClassification(data)
	if err != nil {
		return err
	}
	c.ComputeSecretsClassification = res
	return nil
}

// ComputeListNodesResponse contains the response from method Compute.ListNodes.
type ComputeListNodesResponse struct {
	ComputeListNodesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeListNodesResult contains the result from method Compute.ListNodes.
type ComputeListNodesResult struct {
	AmlComputeNodesInformation
}

// ComputeListResponse contains the response from method Compute.List.
type ComputeListResponse struct {
	ComputeListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeListResult contains the result from method Compute.List.
type ComputeListResult struct {
	PaginatedComputeResourcesList
}

// ComputeRestartPollerResponse contains the response from method Compute.Restart.
type ComputeRestartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ComputeRestartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ComputeRestartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ComputeRestartResponse, error) {
	respType := ComputeRestartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ComputeRestartPollerResponse from the provided client and resume token.
func (l *ComputeRestartPollerResponse) Resume(ctx context.Context, client *ComputeClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ComputeClient.Restart", token, client.pl, client.restartHandleError)
	if err != nil {
		return err
	}
	poller := &ComputeRestartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ComputeRestartResponse contains the response from method Compute.Restart.
type ComputeRestartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeStartPollerResponse contains the response from method Compute.Start.
type ComputeStartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ComputeStartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ComputeStartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ComputeStartResponse, error) {
	respType := ComputeStartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ComputeStartPollerResponse from the provided client and resume token.
func (l *ComputeStartPollerResponse) Resume(ctx context.Context, client *ComputeClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ComputeClient.Start", token, client.pl, client.startHandleError)
	if err != nil {
		return err
	}
	poller := &ComputeStartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ComputeStartResponse contains the response from method Compute.Start.
type ComputeStartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeStopPollerResponse contains the response from method Compute.Stop.
type ComputeStopPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ComputeStopPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ComputeStopPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ComputeStopResponse, error) {
	respType := ComputeStopResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ComputeStopPollerResponse from the provided client and resume token.
func (l *ComputeStopPollerResponse) Resume(ctx context.Context, client *ComputeClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ComputeClient.Stop", token, client.pl, client.stopHandleError)
	if err != nil {
		return err
	}
	poller := &ComputeStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ComputeStopResponse contains the response from method Compute.Stop.
type ComputeStopResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeUpdatePollerResponse contains the response from method Compute.Update.
type ComputeUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ComputeUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ComputeUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ComputeUpdateResponse, error) {
	respType := ComputeUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ComputeResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ComputeUpdatePollerResponse from the provided client and resume token.
func (l *ComputeUpdatePollerResponse) Resume(ctx context.Context, client *ComputeClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ComputeClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ComputeUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ComputeUpdateResponse contains the response from method Compute.Update.
type ComputeUpdateResponse struct {
	ComputeUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputeUpdateResult contains the result from method Compute.Update.
type ComputeUpdateResult struct {
	ComputeResource
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionsCreateOrUpdateResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResult contains the result from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListResponse contains the response from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResponse struct {
	PrivateEndpointConnectionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListResult contains the result from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesListResponse contains the response from method PrivateLinkResources.List.
type PrivateLinkResourcesListResponse struct {
	PrivateLinkResourcesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesListResult contains the result from method PrivateLinkResources.List.
type PrivateLinkResourcesListResult struct {
	PrivateLinkResourceListResult
}

// QuotasListResponse contains the response from method Quotas.List.
type QuotasListResponse struct {
	QuotasListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QuotasListResult contains the result from method Quotas.List.
type QuotasListResult struct {
	ListWorkspaceQuotas
}

// QuotasUpdateResponse contains the response from method Quotas.Update.
type QuotasUpdateResponse struct {
	QuotasUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QuotasUpdateResult contains the result from method Quotas.Update.
type QuotasUpdateResult struct {
	UpdateWorkspaceQuotasResult
}

// UsagesListResponse contains the response from method Usages.List.
type UsagesListResponse struct {
	UsagesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UsagesListResult contains the result from method Usages.List.
type UsagesListResult struct {
	ListUsagesResult
}

// VirtualMachineSizesListResponse contains the response from method VirtualMachineSizes.List.
type VirtualMachineSizesListResponse struct {
	VirtualMachineSizesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachineSizesListResult contains the result from method VirtualMachineSizes.List.
type VirtualMachineSizesListResult struct {
	VirtualMachineSizeListResult
}

// WorkspaceConnectionsCreateResponse contains the response from method WorkspaceConnections.Create.
type WorkspaceConnectionsCreateResponse struct {
	WorkspaceConnectionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceConnectionsCreateResult contains the result from method WorkspaceConnections.Create.
type WorkspaceConnectionsCreateResult struct {
	WorkspaceConnection
}

// WorkspaceConnectionsDeleteResponse contains the response from method WorkspaceConnections.Delete.
type WorkspaceConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceConnectionsGetResponse contains the response from method WorkspaceConnections.Get.
type WorkspaceConnectionsGetResponse struct {
	WorkspaceConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceConnectionsGetResult contains the result from method WorkspaceConnections.Get.
type WorkspaceConnectionsGetResult struct {
	WorkspaceConnection
}

// WorkspaceConnectionsListResponse contains the response from method WorkspaceConnections.List.
type WorkspaceConnectionsListResponse struct {
	WorkspaceConnectionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceConnectionsListResult contains the result from method WorkspaceConnections.List.
type WorkspaceConnectionsListResult struct {
	PaginatedWorkspaceConnectionsList
}

// WorkspaceFeaturesListResponse contains the response from method WorkspaceFeatures.List.
type WorkspaceFeaturesListResponse struct {
	WorkspaceFeaturesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceFeaturesListResult contains the result from method WorkspaceFeatures.List.
type WorkspaceFeaturesListResult struct {
	ListAmlUserFeatureResult
}

// WorkspaceSKUsListResponse contains the response from method WorkspaceSKUs.List.
type WorkspaceSKUsListResponse struct {
	WorkspaceSKUsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspaceSKUsListResult contains the result from method WorkspaceSKUs.List.
type WorkspaceSKUsListResult struct {
	SKUListResult
}

// WorkspacesCreateOrUpdatePollerResponse contains the response from method Workspaces.CreateOrUpdate.
type WorkspacesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspacesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesCreateOrUpdateResponse, error) {
	respType := WorkspacesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Workspace)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *WorkspacesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesCreateOrUpdateResponse contains the response from method Workspaces.CreateOrUpdate.
type WorkspacesCreateOrUpdateResponse struct {
	WorkspacesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesCreateOrUpdateResult contains the result from method Workspaces.CreateOrUpdate.
type WorkspacesCreateOrUpdateResult struct {
	Workspace
}

// WorkspacesDeletePollerResponse contains the response from method Workspaces.Delete.
type WorkspacesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspacesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesDeleteResponse, error) {
	respType := WorkspacesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesDeletePollerResponse from the provided client and resume token.
func (l *WorkspacesDeletePollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesDeleteResponse contains the response from method Workspaces.Delete.
type WorkspacesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesDiagnosePollerResponse contains the response from method Workspaces.Diagnose.
type WorkspacesDiagnosePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesDiagnosePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspacesDiagnosePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesDiagnoseResponse, error) {
	respType := WorkspacesDiagnoseResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DiagnoseResponseResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesDiagnosePollerResponse from the provided client and resume token.
func (l *WorkspacesDiagnosePollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.Diagnose", token, client.pl, client.diagnoseHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesDiagnosePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesDiagnoseResponse contains the response from method Workspaces.Diagnose.
type WorkspacesDiagnoseResponse struct {
	WorkspacesDiagnoseResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesDiagnoseResult contains the result from method Workspaces.Diagnose.
type WorkspacesDiagnoseResult struct {
	DiagnoseResponseResult
}

// WorkspacesGetResponse contains the response from method Workspaces.Get.
type WorkspacesGetResponse struct {
	WorkspacesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesGetResult contains the result from method Workspaces.Get.
type WorkspacesGetResult struct {
	Workspace
}

// WorkspacesListByResourceGroupResponse contains the response from method Workspaces.ListByResourceGroup.
type WorkspacesListByResourceGroupResponse struct {
	WorkspacesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListByResourceGroupResult contains the result from method Workspaces.ListByResourceGroup.
type WorkspacesListByResourceGroupResult struct {
	WorkspaceListResult
}

// WorkspacesListBySubscriptionResponse contains the response from method Workspaces.ListBySubscription.
type WorkspacesListBySubscriptionResponse struct {
	WorkspacesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListBySubscriptionResult contains the result from method Workspaces.ListBySubscription.
type WorkspacesListBySubscriptionResult struct {
	WorkspaceListResult
}

// WorkspacesListKeysResponse contains the response from method Workspaces.ListKeys.
type WorkspacesListKeysResponse struct {
	WorkspacesListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListKeysResult contains the result from method Workspaces.ListKeys.
type WorkspacesListKeysResult struct {
	ListWorkspaceKeysResult
}

// WorkspacesListNotebookAccessTokenResponse contains the response from method Workspaces.ListNotebookAccessToken.
type WorkspacesListNotebookAccessTokenResponse struct {
	WorkspacesListNotebookAccessTokenResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListNotebookAccessTokenResult contains the result from method Workspaces.ListNotebookAccessToken.
type WorkspacesListNotebookAccessTokenResult struct {
	NotebookAccessTokenResult
}

// WorkspacesListNotebookKeysResponse contains the response from method Workspaces.ListNotebookKeys.
type WorkspacesListNotebookKeysResponse struct {
	WorkspacesListNotebookKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListNotebookKeysResult contains the result from method Workspaces.ListNotebookKeys.
type WorkspacesListNotebookKeysResult struct {
	ListNotebookKeysResult
}

// WorkspacesListOutboundNetworkDependenciesEndpointsResponse contains the response from method Workspaces.ListOutboundNetworkDependenciesEndpoints.
type WorkspacesListOutboundNetworkDependenciesEndpointsResponse struct {
	WorkspacesListOutboundNetworkDependenciesEndpointsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListOutboundNetworkDependenciesEndpointsResult contains the result from method Workspaces.ListOutboundNetworkDependenciesEndpoints.
type WorkspacesListOutboundNetworkDependenciesEndpointsResult struct {
	ExternalFQDNResponse
}

// WorkspacesListStorageAccountKeysResponse contains the response from method Workspaces.ListStorageAccountKeys.
type WorkspacesListStorageAccountKeysResponse struct {
	WorkspacesListStorageAccountKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesListStorageAccountKeysResult contains the result from method Workspaces.ListStorageAccountKeys.
type WorkspacesListStorageAccountKeysResult struct {
	ListStorageAccountKeysResult
}

// WorkspacesPrepareNotebookPollerResponse contains the response from method Workspaces.PrepareNotebook.
type WorkspacesPrepareNotebookPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesPrepareNotebookPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspacesPrepareNotebookPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesPrepareNotebookResponse, error) {
	respType := WorkspacesPrepareNotebookResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.NotebookResourceInfo)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesPrepareNotebookPollerResponse from the provided client and resume token.
func (l *WorkspacesPrepareNotebookPollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.PrepareNotebook", token, client.pl, client.prepareNotebookHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesPrepareNotebookPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesPrepareNotebookResponse contains the response from method Workspaces.PrepareNotebook.
type WorkspacesPrepareNotebookResponse struct {
	WorkspacesPrepareNotebookResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesPrepareNotebookResult contains the result from method Workspaces.PrepareNotebook.
type WorkspacesPrepareNotebookResult struct {
	NotebookResourceInfo
}

// WorkspacesResyncKeysPollerResponse contains the response from method Workspaces.ResyncKeys.
type WorkspacesResyncKeysPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *WorkspacesResyncKeysPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l WorkspacesResyncKeysPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (WorkspacesResyncKeysResponse, error) {
	respType := WorkspacesResyncKeysResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a WorkspacesResyncKeysPollerResponse from the provided client and resume token.
func (l *WorkspacesResyncKeysPollerResponse) Resume(ctx context.Context, client *WorkspacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("WorkspacesClient.ResyncKeys", token, client.pl, client.resyncKeysHandleError)
	if err != nil {
		return err
	}
	poller := &WorkspacesResyncKeysPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// WorkspacesResyncKeysResponse contains the response from method Workspaces.ResyncKeys.
type WorkspacesResyncKeysResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesUpdateResponse contains the response from method Workspaces.Update.
type WorkspacesUpdateResponse struct {
	WorkspacesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkspacesUpdateResult contains the result from method Workspaces.Update.
type WorkspacesUpdateResult struct {
	Workspace
}