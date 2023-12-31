//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armquantum

// OfferingsClientListResponse contains the response from method OfferingsClient.NewListPager.
type OfferingsClientListResponse struct {
	OfferingsListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationsList
}

// WorkspaceClientCheckNameAvailabilityResponse contains the response from method WorkspaceClient.CheckNameAvailability.
type WorkspaceClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.BeginCreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.BeginDelete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.NewListByResourceGroupPager.
type WorkspacesClientListByResourceGroupResponse struct {
	WorkspaceListResult
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.NewListBySubscriptionPager.
type WorkspacesClientListBySubscriptionResponse struct {
	WorkspaceListResult
}

// WorkspacesClientUpdateTagsResponse contains the response from method WorkspacesClient.UpdateTags.
type WorkspacesClientUpdateTagsResponse struct {
	Workspace
}
