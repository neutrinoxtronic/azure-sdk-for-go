//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

// ClientCheckExistenceByIDResponse contains the response from method Client.CheckExistenceByID.
type ClientCheckExistenceByIDResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// ClientCheckExistenceResponse contains the response from method Client.CheckExistence.
type ClientCheckExistenceResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// ClientCreateOrUpdateByIDResponse contains the response from method Client.BeginCreateOrUpdateByID.
type ClientCreateOrUpdateByIDResponse struct {
	GenericResource
}

// ClientCreateOrUpdateResponse contains the response from method Client.BeginCreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	GenericResource
}

// ClientDeleteByIDResponse contains the response from method Client.BeginDeleteByID.
type ClientDeleteByIDResponse struct {
	// placeholder for future response values
}

// ClientDeleteResponse contains the response from method Client.BeginDelete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetByIDResponse contains the response from method Client.GetByID.
type ClientGetByIDResponse struct {
	GenericResource
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	GenericResource
}

// ClientListByResourceGroupResponse contains the response from method Client.NewListByResourceGroupPager.
type ClientListByResourceGroupResponse struct {
	ResourceListResult
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	ResourceListResult
}

// ClientMoveResourcesResponse contains the response from method Client.BeginMoveResources.
type ClientMoveResourcesResponse struct {
	// placeholder for future response values
}

// ClientUpdateByIDResponse contains the response from method Client.BeginUpdateByID.
type ClientUpdateByIDResponse struct {
	GenericResource
}

// ClientUpdateResponse contains the response from method Client.BeginUpdate.
type ClientUpdateResponse struct {
	GenericResource
}

// ClientValidateMoveResourcesResponse contains the response from method Client.BeginValidateMoveResources.
type ClientValidateMoveResourcesResponse struct {
	// placeholder for future response values
}

// DeploymentOperationsClientGetAtManagementGroupScopeResponse contains the response from method DeploymentOperationsClient.GetAtManagementGroupScope.
type DeploymentOperationsClientGetAtManagementGroupScopeResponse struct {
	DeploymentOperation
}

// DeploymentOperationsClientGetAtScopeResponse contains the response from method DeploymentOperationsClient.GetAtScope.
type DeploymentOperationsClientGetAtScopeResponse struct {
	DeploymentOperation
}

// DeploymentOperationsClientGetAtSubscriptionScopeResponse contains the response from method DeploymentOperationsClient.GetAtSubscriptionScope.
type DeploymentOperationsClientGetAtSubscriptionScopeResponse struct {
	DeploymentOperation
}

// DeploymentOperationsClientGetAtTenantScopeResponse contains the response from method DeploymentOperationsClient.GetAtTenantScope.
type DeploymentOperationsClientGetAtTenantScopeResponse struct {
	DeploymentOperation
}

// DeploymentOperationsClientGetResponse contains the response from method DeploymentOperationsClient.Get.
type DeploymentOperationsClientGetResponse struct {
	DeploymentOperation
}

// DeploymentOperationsClientListAtManagementGroupScopeResponse contains the response from method DeploymentOperationsClient.NewListAtManagementGroupScopePager.
type DeploymentOperationsClientListAtManagementGroupScopeResponse struct {
	DeploymentOperationsListResult
}

// DeploymentOperationsClientListAtScopeResponse contains the response from method DeploymentOperationsClient.NewListAtScopePager.
type DeploymentOperationsClientListAtScopeResponse struct {
	DeploymentOperationsListResult
}

// DeploymentOperationsClientListAtSubscriptionScopeResponse contains the response from method DeploymentOperationsClient.NewListAtSubscriptionScopePager.
type DeploymentOperationsClientListAtSubscriptionScopeResponse struct {
	DeploymentOperationsListResult
}

// DeploymentOperationsClientListAtTenantScopeResponse contains the response from method DeploymentOperationsClient.NewListAtTenantScopePager.
type DeploymentOperationsClientListAtTenantScopeResponse struct {
	DeploymentOperationsListResult
}

// DeploymentOperationsClientListResponse contains the response from method DeploymentOperationsClient.NewListPager.
type DeploymentOperationsClientListResponse struct {
	DeploymentOperationsListResult
}

// DeploymentsClientCalculateTemplateHashResponse contains the response from method DeploymentsClient.CalculateTemplateHash.
type DeploymentsClientCalculateTemplateHashResponse struct {
	TemplateHashResult
}

// DeploymentsClientCancelAtManagementGroupScopeResponse contains the response from method DeploymentsClient.CancelAtManagementGroupScope.
type DeploymentsClientCancelAtManagementGroupScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientCancelAtScopeResponse contains the response from method DeploymentsClient.CancelAtScope.
type DeploymentsClientCancelAtScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientCancelAtSubscriptionScopeResponse contains the response from method DeploymentsClient.CancelAtSubscriptionScope.
type DeploymentsClientCancelAtSubscriptionScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientCancelAtTenantScopeResponse contains the response from method DeploymentsClient.CancelAtTenantScope.
type DeploymentsClientCancelAtTenantScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientCancelResponse contains the response from method DeploymentsClient.Cancel.
type DeploymentsClientCancelResponse struct {
	// placeholder for future response values
}

// DeploymentsClientCheckExistenceAtManagementGroupScopeResponse contains the response from method DeploymentsClient.CheckExistenceAtManagementGroupScope.
type DeploymentsClientCheckExistenceAtManagementGroupScopeResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// DeploymentsClientCheckExistenceAtScopeResponse contains the response from method DeploymentsClient.CheckExistenceAtScope.
type DeploymentsClientCheckExistenceAtScopeResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// DeploymentsClientCheckExistenceAtSubscriptionScopeResponse contains the response from method DeploymentsClient.CheckExistenceAtSubscriptionScope.
type DeploymentsClientCheckExistenceAtSubscriptionScopeResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// DeploymentsClientCheckExistenceAtTenantScopeResponse contains the response from method DeploymentsClient.CheckExistenceAtTenantScope.
type DeploymentsClientCheckExistenceAtTenantScopeResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// DeploymentsClientCheckExistenceResponse contains the response from method DeploymentsClient.CheckExistence.
type DeploymentsClientCheckExistenceResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// DeploymentsClientCreateOrUpdateAtManagementGroupScopeResponse contains the response from method DeploymentsClient.BeginCreateOrUpdateAtManagementGroupScope.
type DeploymentsClientCreateOrUpdateAtManagementGroupScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientCreateOrUpdateAtScopeResponse contains the response from method DeploymentsClient.BeginCreateOrUpdateAtScope.
type DeploymentsClientCreateOrUpdateAtScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientCreateOrUpdateAtSubscriptionScopeResponse contains the response from method DeploymentsClient.BeginCreateOrUpdateAtSubscriptionScope.
type DeploymentsClientCreateOrUpdateAtSubscriptionScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientCreateOrUpdateAtTenantScopeResponse contains the response from method DeploymentsClient.BeginCreateOrUpdateAtTenantScope.
type DeploymentsClientCreateOrUpdateAtTenantScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientCreateOrUpdateResponse contains the response from method DeploymentsClient.BeginCreateOrUpdate.
type DeploymentsClientCreateOrUpdateResponse struct {
	DeploymentExtended
}

// DeploymentsClientDeleteAtManagementGroupScopeResponse contains the response from method DeploymentsClient.BeginDeleteAtManagementGroupScope.
type DeploymentsClientDeleteAtManagementGroupScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientDeleteAtScopeResponse contains the response from method DeploymentsClient.BeginDeleteAtScope.
type DeploymentsClientDeleteAtScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientDeleteAtSubscriptionScopeResponse contains the response from method DeploymentsClient.BeginDeleteAtSubscriptionScope.
type DeploymentsClientDeleteAtSubscriptionScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientDeleteAtTenantScopeResponse contains the response from method DeploymentsClient.BeginDeleteAtTenantScope.
type DeploymentsClientDeleteAtTenantScopeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientDeleteResponse contains the response from method DeploymentsClient.BeginDelete.
type DeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeploymentsClientExportTemplateAtManagementGroupScopeResponse contains the response from method DeploymentsClient.ExportTemplateAtManagementGroupScope.
type DeploymentsClientExportTemplateAtManagementGroupScopeResponse struct {
	DeploymentExportResult
}

// DeploymentsClientExportTemplateAtScopeResponse contains the response from method DeploymentsClient.ExportTemplateAtScope.
type DeploymentsClientExportTemplateAtScopeResponse struct {
	DeploymentExportResult
}

// DeploymentsClientExportTemplateAtSubscriptionScopeResponse contains the response from method DeploymentsClient.ExportTemplateAtSubscriptionScope.
type DeploymentsClientExportTemplateAtSubscriptionScopeResponse struct {
	DeploymentExportResult
}

// DeploymentsClientExportTemplateAtTenantScopeResponse contains the response from method DeploymentsClient.ExportTemplateAtTenantScope.
type DeploymentsClientExportTemplateAtTenantScopeResponse struct {
	DeploymentExportResult
}

// DeploymentsClientExportTemplateResponse contains the response from method DeploymentsClient.ExportTemplate.
type DeploymentsClientExportTemplateResponse struct {
	DeploymentExportResult
}

// DeploymentsClientGetAtManagementGroupScopeResponse contains the response from method DeploymentsClient.GetAtManagementGroupScope.
type DeploymentsClientGetAtManagementGroupScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientGetAtScopeResponse contains the response from method DeploymentsClient.GetAtScope.
type DeploymentsClientGetAtScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientGetAtSubscriptionScopeResponse contains the response from method DeploymentsClient.GetAtSubscriptionScope.
type DeploymentsClientGetAtSubscriptionScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientGetAtTenantScopeResponse contains the response from method DeploymentsClient.GetAtTenantScope.
type DeploymentsClientGetAtTenantScopeResponse struct {
	DeploymentExtended
}

// DeploymentsClientGetResponse contains the response from method DeploymentsClient.Get.
type DeploymentsClientGetResponse struct {
	DeploymentExtended
}

// DeploymentsClientListAtManagementGroupScopeResponse contains the response from method DeploymentsClient.NewListAtManagementGroupScopePager.
type DeploymentsClientListAtManagementGroupScopeResponse struct {
	DeploymentListResult
}

// DeploymentsClientListAtScopeResponse contains the response from method DeploymentsClient.NewListAtScopePager.
type DeploymentsClientListAtScopeResponse struct {
	DeploymentListResult
}

// DeploymentsClientListAtSubscriptionScopeResponse contains the response from method DeploymentsClient.NewListAtSubscriptionScopePager.
type DeploymentsClientListAtSubscriptionScopeResponse struct {
	DeploymentListResult
}

// DeploymentsClientListAtTenantScopeResponse contains the response from method DeploymentsClient.NewListAtTenantScopePager.
type DeploymentsClientListAtTenantScopeResponse struct {
	DeploymentListResult
}

// DeploymentsClientListByResourceGroupResponse contains the response from method DeploymentsClient.NewListByResourceGroupPager.
type DeploymentsClientListByResourceGroupResponse struct {
	DeploymentListResult
}

// DeploymentsClientValidateAtManagementGroupScopeResponse contains the response from method DeploymentsClient.BeginValidateAtManagementGroupScope.
type DeploymentsClientValidateAtManagementGroupScopeResponse struct {
	DeploymentValidateResult
}

// DeploymentsClientValidateAtScopeResponse contains the response from method DeploymentsClient.BeginValidateAtScope.
type DeploymentsClientValidateAtScopeResponse struct {
	DeploymentValidateResult
}

// DeploymentsClientValidateAtSubscriptionScopeResponse contains the response from method DeploymentsClient.BeginValidateAtSubscriptionScope.
type DeploymentsClientValidateAtSubscriptionScopeResponse struct {
	DeploymentValidateResult
}

// DeploymentsClientValidateAtTenantScopeResponse contains the response from method DeploymentsClient.BeginValidateAtTenantScope.
type DeploymentsClientValidateAtTenantScopeResponse struct {
	DeploymentValidateResult
}

// DeploymentsClientValidateResponse contains the response from method DeploymentsClient.BeginValidate.
type DeploymentsClientValidateResponse struct {
	DeploymentValidateResult
}

// DeploymentsClientWhatIfAtManagementGroupScopeResponse contains the response from method DeploymentsClient.BeginWhatIfAtManagementGroupScope.
type DeploymentsClientWhatIfAtManagementGroupScopeResponse struct {
	WhatIfOperationResult
}

// DeploymentsClientWhatIfAtSubscriptionScopeResponse contains the response from method DeploymentsClient.BeginWhatIfAtSubscriptionScope.
type DeploymentsClientWhatIfAtSubscriptionScopeResponse struct {
	WhatIfOperationResult
}

// DeploymentsClientWhatIfAtTenantScopeResponse contains the response from method DeploymentsClient.BeginWhatIfAtTenantScope.
type DeploymentsClientWhatIfAtTenantScopeResponse struct {
	WhatIfOperationResult
}

// DeploymentsClientWhatIfResponse contains the response from method DeploymentsClient.BeginWhatIf.
type DeploymentsClientWhatIfResponse struct {
	WhatIfOperationResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// ProviderResourceTypesClientListResponse contains the response from method ProviderResourceTypesClient.List.
type ProviderResourceTypesClientListResponse struct {
	ProviderResourceTypeListResult
}

// ProvidersClientGetAtTenantScopeResponse contains the response from method ProvidersClient.GetAtTenantScope.
type ProvidersClientGetAtTenantScopeResponse struct {
	Provider
}

// ProvidersClientGetResponse contains the response from method ProvidersClient.Get.
type ProvidersClientGetResponse struct {
	Provider
}

// ProvidersClientListAtTenantScopeResponse contains the response from method ProvidersClient.NewListAtTenantScopePager.
type ProvidersClientListAtTenantScopeResponse struct {
	ProviderListResult
}

// ProvidersClientListResponse contains the response from method ProvidersClient.NewListPager.
type ProvidersClientListResponse struct {
	ProviderListResult
}

// ProvidersClientProviderPermissionsResponse contains the response from method ProvidersClient.ProviderPermissions.
type ProvidersClientProviderPermissionsResponse struct {
	ProviderPermissionListResult
}

// ProvidersClientRegisterAtManagementGroupScopeResponse contains the response from method ProvidersClient.RegisterAtManagementGroupScope.
type ProvidersClientRegisterAtManagementGroupScopeResponse struct {
	// placeholder for future response values
}

// ProvidersClientRegisterResponse contains the response from method ProvidersClient.Register.
type ProvidersClientRegisterResponse struct {
	Provider
}

// ProvidersClientUnregisterResponse contains the response from method ProvidersClient.Unregister.
type ProvidersClientUnregisterResponse struct {
	Provider
}

// ResourceGroupsClientCheckExistenceResponse contains the response from method ResourceGroupsClient.CheckExistence.
type ResourceGroupsClientCheckExistenceResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// ResourceGroupsClientCreateOrUpdateResponse contains the response from method ResourceGroupsClient.CreateOrUpdate.
type ResourceGroupsClientCreateOrUpdateResponse struct {
	ResourceGroup
}

// ResourceGroupsClientDeleteResponse contains the response from method ResourceGroupsClient.BeginDelete.
type ResourceGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceGroupsClientExportTemplateResponse contains the response from method ResourceGroupsClient.BeginExportTemplate.
type ResourceGroupsClientExportTemplateResponse struct {
	ResourceGroupExportResult
}

// ResourceGroupsClientGetResponse contains the response from method ResourceGroupsClient.Get.
type ResourceGroupsClientGetResponse struct {
	ResourceGroup
}

// ResourceGroupsClientListResponse contains the response from method ResourceGroupsClient.NewListPager.
type ResourceGroupsClientListResponse struct {
	ResourceGroupListResult
}

// ResourceGroupsClientUpdateResponse contains the response from method ResourceGroupsClient.Update.
type ResourceGroupsClientUpdateResponse struct {
	ResourceGroup
}

// TagsClientCreateOrUpdateAtScopeResponse contains the response from method TagsClient.CreateOrUpdateAtScope.
type TagsClientCreateOrUpdateAtScopeResponse struct {
	TagsResource
}

// TagsClientCreateOrUpdateResponse contains the response from method TagsClient.CreateOrUpdate.
type TagsClientCreateOrUpdateResponse struct {
	TagDetails
}

// TagsClientCreateOrUpdateValueResponse contains the response from method TagsClient.CreateOrUpdateValue.
type TagsClientCreateOrUpdateValueResponse struct {
	TagValue
}

// TagsClientDeleteAtScopeResponse contains the response from method TagsClient.DeleteAtScope.
type TagsClientDeleteAtScopeResponse struct {
	// placeholder for future response values
}

// TagsClientDeleteResponse contains the response from method TagsClient.Delete.
type TagsClientDeleteResponse struct {
	// placeholder for future response values
}

// TagsClientDeleteValueResponse contains the response from method TagsClient.DeleteValue.
type TagsClientDeleteValueResponse struct {
	// placeholder for future response values
}

// TagsClientGetAtScopeResponse contains the response from method TagsClient.GetAtScope.
type TagsClientGetAtScopeResponse struct {
	TagsResource
}

// TagsClientListResponse contains the response from method TagsClient.NewListPager.
type TagsClientListResponse struct {
	TagsListResult
}

// TagsClientUpdateAtScopeResponse contains the response from method TagsClient.UpdateAtScope.
type TagsClientUpdateAtScopeResponse struct {
	TagsResource
}
