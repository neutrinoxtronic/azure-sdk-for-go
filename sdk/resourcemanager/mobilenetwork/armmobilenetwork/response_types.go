//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork

// AttachedDataNetworksClientCreateOrUpdateResponse contains the response from method AttachedDataNetworksClient.BeginCreateOrUpdate.
type AttachedDataNetworksClientCreateOrUpdateResponse struct {
	AttachedDataNetwork
}

// AttachedDataNetworksClientDeleteResponse contains the response from method AttachedDataNetworksClient.BeginDelete.
type AttachedDataNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// AttachedDataNetworksClientGetResponse contains the response from method AttachedDataNetworksClient.Get.
type AttachedDataNetworksClientGetResponse struct {
	AttachedDataNetwork
}

// AttachedDataNetworksClientListByPacketCoreDataPlaneResponse contains the response from method AttachedDataNetworksClient.NewListByPacketCoreDataPlanePager.
type AttachedDataNetworksClientListByPacketCoreDataPlaneResponse struct {
	AttachedDataNetworkListResult
}

// AttachedDataNetworksClientUpdateTagsResponse contains the response from method AttachedDataNetworksClient.UpdateTags.
type AttachedDataNetworksClientUpdateTagsResponse struct {
	AttachedDataNetwork
}

// DataNetworksClientCreateOrUpdateResponse contains the response from method DataNetworksClient.BeginCreateOrUpdate.
type DataNetworksClientCreateOrUpdateResponse struct {
	DataNetwork
}

// DataNetworksClientDeleteResponse contains the response from method DataNetworksClient.BeginDelete.
type DataNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// DataNetworksClientGetResponse contains the response from method DataNetworksClient.Get.
type DataNetworksClientGetResponse struct {
	DataNetwork
}

// DataNetworksClientListByMobileNetworkResponse contains the response from method DataNetworksClient.NewListByMobileNetworkPager.
type DataNetworksClientListByMobileNetworkResponse struct {
	DataNetworkListResult
}

// DataNetworksClientUpdateTagsResponse contains the response from method DataNetworksClient.UpdateTags.
type DataNetworksClientUpdateTagsResponse struct {
	DataNetwork
}

// MobileNetworksClientCreateOrUpdateResponse contains the response from method MobileNetworksClient.BeginCreateOrUpdate.
type MobileNetworksClientCreateOrUpdateResponse struct {
	MobileNetwork
}

// MobileNetworksClientDeleteResponse contains the response from method MobileNetworksClient.BeginDelete.
type MobileNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// MobileNetworksClientGetResponse contains the response from method MobileNetworksClient.Get.
type MobileNetworksClientGetResponse struct {
	MobileNetwork
}

// MobileNetworksClientListByResourceGroupResponse contains the response from method MobileNetworksClient.NewListByResourceGroupPager.
type MobileNetworksClientListByResourceGroupResponse struct {
	ListResult
}

// MobileNetworksClientListBySubscriptionResponse contains the response from method MobileNetworksClient.NewListBySubscriptionPager.
type MobileNetworksClientListBySubscriptionResponse struct {
	ListResult
}

// MobileNetworksClientUpdateTagsResponse contains the response from method MobileNetworksClient.UpdateTags.
type MobileNetworksClientUpdateTagsResponse struct {
	MobileNetwork
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationList
}

// PacketCoreControlPlaneVersionsClientGetResponse contains the response from method PacketCoreControlPlaneVersionsClient.Get.
type PacketCoreControlPlaneVersionsClientGetResponse struct {
	PacketCoreControlPlaneVersion
}

// PacketCoreControlPlaneVersionsClientListResponse contains the response from method PacketCoreControlPlaneVersionsClient.NewListPager.
type PacketCoreControlPlaneVersionsClientListResponse struct {
	PacketCoreControlPlaneVersionListResult
}

// PacketCoreControlPlanesClientCollectDiagnosticsPackageResponse contains the response from method PacketCoreControlPlanesClient.BeginCollectDiagnosticsPackage.
type PacketCoreControlPlanesClientCollectDiagnosticsPackageResponse struct {
	AsyncOperationStatus
}

// PacketCoreControlPlanesClientCreateOrUpdateResponse contains the response from method PacketCoreControlPlanesClient.BeginCreateOrUpdate.
type PacketCoreControlPlanesClientCreateOrUpdateResponse struct {
	PacketCoreControlPlane
}

// PacketCoreControlPlanesClientDeleteResponse contains the response from method PacketCoreControlPlanesClient.BeginDelete.
type PacketCoreControlPlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// PacketCoreControlPlanesClientGetResponse contains the response from method PacketCoreControlPlanesClient.Get.
type PacketCoreControlPlanesClientGetResponse struct {
	PacketCoreControlPlane
}

// PacketCoreControlPlanesClientListByResourceGroupResponse contains the response from method PacketCoreControlPlanesClient.NewListByResourceGroupPager.
type PacketCoreControlPlanesClientListByResourceGroupResponse struct {
	PacketCoreControlPlaneListResult
}

// PacketCoreControlPlanesClientListBySubscriptionResponse contains the response from method PacketCoreControlPlanesClient.NewListBySubscriptionPager.
type PacketCoreControlPlanesClientListBySubscriptionResponse struct {
	PacketCoreControlPlaneListResult
}

// PacketCoreControlPlanesClientReinstallResponse contains the response from method PacketCoreControlPlanesClient.BeginReinstall.
type PacketCoreControlPlanesClientReinstallResponse struct {
	AsyncOperationStatus
}

// PacketCoreControlPlanesClientRollbackResponse contains the response from method PacketCoreControlPlanesClient.BeginRollback.
type PacketCoreControlPlanesClientRollbackResponse struct {
	AsyncOperationStatus
}

// PacketCoreControlPlanesClientUpdateTagsResponse contains the response from method PacketCoreControlPlanesClient.UpdateTags.
type PacketCoreControlPlanesClientUpdateTagsResponse struct {
	PacketCoreControlPlane
}

// PacketCoreDataPlanesClientCreateOrUpdateResponse contains the response from method PacketCoreDataPlanesClient.BeginCreateOrUpdate.
type PacketCoreDataPlanesClientCreateOrUpdateResponse struct {
	PacketCoreDataPlane
}

// PacketCoreDataPlanesClientDeleteResponse contains the response from method PacketCoreDataPlanesClient.BeginDelete.
type PacketCoreDataPlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// PacketCoreDataPlanesClientGetResponse contains the response from method PacketCoreDataPlanesClient.Get.
type PacketCoreDataPlanesClientGetResponse struct {
	PacketCoreDataPlane
}

// PacketCoreDataPlanesClientListByPacketCoreControlPlaneResponse contains the response from method PacketCoreDataPlanesClient.NewListByPacketCoreControlPlanePager.
type PacketCoreDataPlanesClientListByPacketCoreControlPlaneResponse struct {
	PacketCoreDataPlaneListResult
}

// PacketCoreDataPlanesClientUpdateTagsResponse contains the response from method PacketCoreDataPlanesClient.UpdateTags.
type PacketCoreDataPlanesClientUpdateTagsResponse struct {
	PacketCoreDataPlane
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	Service
}

// ServicesClientListByMobileNetworkResponse contains the response from method ServicesClient.NewListByMobileNetworkPager.
type ServicesClientListByMobileNetworkResponse struct {
	ServiceListResult
}

// ServicesClientUpdateTagsResponse contains the response from method ServicesClient.UpdateTags.
type ServicesClientUpdateTagsResponse struct {
	Service
}

// SimGroupsClientCreateOrUpdateResponse contains the response from method SimGroupsClient.BeginCreateOrUpdate.
type SimGroupsClientCreateOrUpdateResponse struct {
	SimGroup
}

// SimGroupsClientDeleteResponse contains the response from method SimGroupsClient.BeginDelete.
type SimGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// SimGroupsClientGetResponse contains the response from method SimGroupsClient.Get.
type SimGroupsClientGetResponse struct {
	SimGroup
}

// SimGroupsClientListByResourceGroupResponse contains the response from method SimGroupsClient.NewListByResourceGroupPager.
type SimGroupsClientListByResourceGroupResponse struct {
	SimGroupListResult
}

// SimGroupsClientListBySubscriptionResponse contains the response from method SimGroupsClient.NewListBySubscriptionPager.
type SimGroupsClientListBySubscriptionResponse struct {
	SimGroupListResult
}

// SimGroupsClientUpdateTagsResponse contains the response from method SimGroupsClient.UpdateTags.
type SimGroupsClientUpdateTagsResponse struct {
	SimGroup
}

// SimPoliciesClientCreateOrUpdateResponse contains the response from method SimPoliciesClient.BeginCreateOrUpdate.
type SimPoliciesClientCreateOrUpdateResponse struct {
	SimPolicy
}

// SimPoliciesClientDeleteResponse contains the response from method SimPoliciesClient.BeginDelete.
type SimPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SimPoliciesClientGetResponse contains the response from method SimPoliciesClient.Get.
type SimPoliciesClientGetResponse struct {
	SimPolicy
}

// SimPoliciesClientListByMobileNetworkResponse contains the response from method SimPoliciesClient.NewListByMobileNetworkPager.
type SimPoliciesClientListByMobileNetworkResponse struct {
	SimPolicyListResult
}

// SimPoliciesClientUpdateTagsResponse contains the response from method SimPoliciesClient.UpdateTags.
type SimPoliciesClientUpdateTagsResponse struct {
	SimPolicy
}

// SimsClientBulkDeleteResponse contains the response from method SimsClient.BeginBulkDelete.
type SimsClientBulkDeleteResponse struct {
	AsyncOperationStatus
}

// SimsClientBulkUploadEncryptedResponse contains the response from method SimsClient.BeginBulkUploadEncrypted.
type SimsClientBulkUploadEncryptedResponse struct {
	AsyncOperationStatus
}

// SimsClientBulkUploadResponse contains the response from method SimsClient.BeginBulkUpload.
type SimsClientBulkUploadResponse struct {
	AsyncOperationStatus
}

// SimsClientCreateOrUpdateResponse contains the response from method SimsClient.BeginCreateOrUpdate.
type SimsClientCreateOrUpdateResponse struct {
	Sim
}

// SimsClientDeleteResponse contains the response from method SimsClient.BeginDelete.
type SimsClientDeleteResponse struct {
	// placeholder for future response values
}

// SimsClientGetResponse contains the response from method SimsClient.Get.
type SimsClientGetResponse struct {
	Sim
}

// SimsClientListByGroupResponse contains the response from method SimsClient.NewListByGroupPager.
type SimsClientListByGroupResponse struct {
	SimListResult
}

// SitesClientCreateOrUpdateResponse contains the response from method SitesClient.BeginCreateOrUpdate.
type SitesClientCreateOrUpdateResponse struct {
	Site
}

// SitesClientDeleteResponse contains the response from method SitesClient.BeginDelete.
type SitesClientDeleteResponse struct {
	// placeholder for future response values
}

// SitesClientGetResponse contains the response from method SitesClient.Get.
type SitesClientGetResponse struct {
	Site
}

// SitesClientListByMobileNetworkResponse contains the response from method SitesClient.NewListByMobileNetworkPager.
type SitesClientListByMobileNetworkResponse struct {
	SiteListResult
}

// SitesClientUpdateTagsResponse contains the response from method SitesClient.UpdateTags.
type SitesClientUpdateTagsResponse struct {
	Site
}

// SlicesClientCreateOrUpdateResponse contains the response from method SlicesClient.BeginCreateOrUpdate.
type SlicesClientCreateOrUpdateResponse struct {
	Slice
}

// SlicesClientDeleteResponse contains the response from method SlicesClient.BeginDelete.
type SlicesClientDeleteResponse struct {
	// placeholder for future response values
}

// SlicesClientGetResponse contains the response from method SlicesClient.Get.
type SlicesClientGetResponse struct {
	Slice
}

// SlicesClientListByMobileNetworkResponse contains the response from method SlicesClient.NewListByMobileNetworkPager.
type SlicesClientListByMobileNetworkResponse struct {
	SliceListResult
}

// SlicesClientUpdateTagsResponse contains the response from method SlicesClient.UpdateTags.
type SlicesClientUpdateTagsResponse struct {
	Slice
}
