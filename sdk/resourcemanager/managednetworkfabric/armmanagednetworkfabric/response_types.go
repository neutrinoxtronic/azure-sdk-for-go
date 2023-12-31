//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

import "encoding/json"

// AccessControlListsClientCreateResponse contains the response from method AccessControlListsClient.Create.
type AccessControlListsClientCreateResponse struct {
	AccessControlList
}

// AccessControlListsClientDeleteResponse contains the response from method AccessControlListsClient.Delete.
type AccessControlListsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccessControlListsClientGetResponse contains the response from method AccessControlListsClient.Get.
type AccessControlListsClientGetResponse struct {
	AccessControlList
}

// AccessControlListsClientListByResourceGroupResponse contains the response from method AccessControlListsClient.NewListByResourceGroupPager.
type AccessControlListsClientListByResourceGroupResponse struct {
	AccessControlListsListResult
}

// AccessControlListsClientListBySubscriptionResponse contains the response from method AccessControlListsClient.NewListBySubscriptionPager.
type AccessControlListsClientListBySubscriptionResponse struct {
	AccessControlListsListResult
}

// AccessControlListsClientUpdateResponse contains the response from method AccessControlListsClient.Update.
type AccessControlListsClientUpdateResponse struct {
	AccessControlList
}

// ExternalNetworksClientClearArpEntriesResponse contains the response from method ExternalNetworksClient.BeginClearArpEntries.
type ExternalNetworksClientClearArpEntriesResponse struct {
	// placeholder for future response values
}

// ExternalNetworksClientClearIPv6NeighborsResponse contains the response from method ExternalNetworksClient.BeginClearIPv6Neighbors.
type ExternalNetworksClientClearIPv6NeighborsResponse struct {
	// placeholder for future response values
}

// ExternalNetworksClientCreateResponse contains the response from method ExternalNetworksClient.BeginCreate.
type ExternalNetworksClientCreateResponse struct {
	ExternalNetwork
}

// ExternalNetworksClientDeleteResponse contains the response from method ExternalNetworksClient.BeginDelete.
type ExternalNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// ExternalNetworksClientGetResponse contains the response from method ExternalNetworksClient.Get.
type ExternalNetworksClientGetResponse struct {
	ExternalNetwork
}

// ExternalNetworksClientListResponse contains the response from method ExternalNetworksClient.NewListPager.
type ExternalNetworksClientListResponse struct {
	ExternalNetworksList
}

// ExternalNetworksClientUpdateAdministrativeStateResponse contains the response from method ExternalNetworksClient.BeginUpdateAdministrativeState.
type ExternalNetworksClientUpdateAdministrativeStateResponse struct {
	// placeholder for future response values
}

// ExternalNetworksClientUpdateBfdForBgpAdministrativeStateResponse contains the response from method ExternalNetworksClient.BeginUpdateBfdForBgpAdministrativeState.
type ExternalNetworksClientUpdateBfdForBgpAdministrativeStateResponse struct {
	// placeholder for future response values
}

// ExternalNetworksClientUpdateBgpAdministrativeStateResponse contains the response from method ExternalNetworksClient.BeginUpdateBgpAdministrativeState.
type ExternalNetworksClientUpdateBgpAdministrativeStateResponse struct {
	// placeholder for future response values
}

// ExternalNetworksClientUpdateResponse contains the response from method ExternalNetworksClient.BeginUpdate.
type ExternalNetworksClientUpdateResponse struct {
	ExternalNetwork
}

// IPCommunitiesClientCreateResponse contains the response from method IPCommunitiesClient.BeginCreate.
type IPCommunitiesClientCreateResponse struct {
	IPCommunity
}

// IPCommunitiesClientDeleteResponse contains the response from method IPCommunitiesClient.BeginDelete.
type IPCommunitiesClientDeleteResponse struct {
	// placeholder for future response values
}

// IPCommunitiesClientGetResponse contains the response from method IPCommunitiesClient.Get.
type IPCommunitiesClientGetResponse struct {
	IPCommunity
}

// IPCommunitiesClientListByResourceGroupResponse contains the response from method IPCommunitiesClient.NewListByResourceGroupPager.
type IPCommunitiesClientListByResourceGroupResponse struct {
	IPCommunitiesListResult
}

// IPCommunitiesClientListBySubscriptionResponse contains the response from method IPCommunitiesClient.NewListBySubscriptionPager.
type IPCommunitiesClientListBySubscriptionResponse struct {
	IPCommunitiesListResult
}

// IPCommunitiesClientUpdateResponse contains the response from method IPCommunitiesClient.BeginUpdate.
type IPCommunitiesClientUpdateResponse struct {
	IPCommunity
}

// IPExtendedCommunitiesClientCreateResponse contains the response from method IPExtendedCommunitiesClient.BeginCreate.
type IPExtendedCommunitiesClientCreateResponse struct {
	IPExtendedCommunity
}

// IPExtendedCommunitiesClientDeleteResponse contains the response from method IPExtendedCommunitiesClient.BeginDelete.
type IPExtendedCommunitiesClientDeleteResponse struct {
	// placeholder for future response values
}

// IPExtendedCommunitiesClientGetResponse contains the response from method IPExtendedCommunitiesClient.Get.
type IPExtendedCommunitiesClientGetResponse struct {
	IPExtendedCommunity
}

// IPExtendedCommunitiesClientListByResourceGroupResponse contains the response from method IPExtendedCommunitiesClient.NewListByResourceGroupPager.
type IPExtendedCommunitiesClientListByResourceGroupResponse struct {
	IPExtendedCommunityListResult
}

// IPExtendedCommunitiesClientListBySubscriptionResponse contains the response from method IPExtendedCommunitiesClient.NewListBySubscriptionPager.
type IPExtendedCommunitiesClientListBySubscriptionResponse struct {
	IPExtendedCommunityListResult
}

// IPExtendedCommunitiesClientUpdateResponse contains the response from method IPExtendedCommunitiesClient.BeginUpdate.
type IPExtendedCommunitiesClientUpdateResponse struct {
	IPExtendedCommunity
}

// IPPrefixesClientCreateResponse contains the response from method IPPrefixesClient.BeginCreate.
type IPPrefixesClientCreateResponse struct {
	IPPrefix
}

// IPPrefixesClientDeleteResponse contains the response from method IPPrefixesClient.BeginDelete.
type IPPrefixesClientDeleteResponse struct {
	// placeholder for future response values
}

// IPPrefixesClientGetResponse contains the response from method IPPrefixesClient.Get.
type IPPrefixesClientGetResponse struct {
	IPPrefix
}

// IPPrefixesClientListByResourceGroupResponse contains the response from method IPPrefixesClient.NewListByResourceGroupPager.
type IPPrefixesClientListByResourceGroupResponse struct {
	IPPrefixesListResult
}

// IPPrefixesClientListBySubscriptionResponse contains the response from method IPPrefixesClient.NewListBySubscriptionPager.
type IPPrefixesClientListBySubscriptionResponse struct {
	IPPrefixesListResult
}

// IPPrefixesClientUpdateResponse contains the response from method IPPrefixesClient.BeginUpdate.
type IPPrefixesClientUpdateResponse struct {
	IPPrefix
}

// InternalNetworksClientClearArpEntriesResponse contains the response from method InternalNetworksClient.BeginClearArpEntries.
type InternalNetworksClientClearArpEntriesResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientClearIPv6NeighborsResponse contains the response from method InternalNetworksClient.BeginClearIPv6Neighbors.
type InternalNetworksClientClearIPv6NeighborsResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientCreateResponse contains the response from method InternalNetworksClient.BeginCreate.
type InternalNetworksClientCreateResponse struct {
	InternalNetwork
}

// InternalNetworksClientDeleteResponse contains the response from method InternalNetworksClient.BeginDelete.
type InternalNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientGetResponse contains the response from method InternalNetworksClient.Get.
type InternalNetworksClientGetResponse struct {
	InternalNetwork
}

// InternalNetworksClientListResponse contains the response from method InternalNetworksClient.NewListPager.
type InternalNetworksClientListResponse struct {
	InternalNetworksList
}

// InternalNetworksClientUpdateAdministrativeStateResponse contains the response from method InternalNetworksClient.BeginUpdateAdministrativeState.
type InternalNetworksClientUpdateAdministrativeStateResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientUpdateBfdForBgpAdministrativeStateResponse contains the response from method InternalNetworksClient.BeginUpdateBfdForBgpAdministrativeState.
type InternalNetworksClientUpdateBfdForBgpAdministrativeStateResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientUpdateBfdForStaticRouteAdministrativeStateResponse contains the response from method InternalNetworksClient.BeginUpdateBfdForStaticRouteAdministrativeState.
type InternalNetworksClientUpdateBfdForStaticRouteAdministrativeStateResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientUpdateBgpAdministrativeStateResponse contains the response from method InternalNetworksClient.BeginUpdateBgpAdministrativeState.
type InternalNetworksClientUpdateBgpAdministrativeStateResponse struct {
	// placeholder for future response values
}

// InternalNetworksClientUpdateResponse contains the response from method InternalNetworksClient.BeginUpdate.
type InternalNetworksClientUpdateResponse struct {
	InternalNetwork
}

// L2IsolationDomainsClientClearArpTableResponse contains the response from method L2IsolationDomainsClient.BeginClearArpTable.
type L2IsolationDomainsClientClearArpTableResponse struct {
	// placeholder for future response values
}

// L2IsolationDomainsClientClearNeighborTableResponse contains the response from method L2IsolationDomainsClient.BeginClearNeighborTable.
type L2IsolationDomainsClientClearNeighborTableResponse struct {
	// placeholder for future response values
}

// L2IsolationDomainsClientCreateResponse contains the response from method L2IsolationDomainsClient.BeginCreate.
type L2IsolationDomainsClientCreateResponse struct {
	L2IsolationDomain
}

// L2IsolationDomainsClientDeleteResponse contains the response from method L2IsolationDomainsClient.BeginDelete.
type L2IsolationDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// L2IsolationDomainsClientGetArpEntriesResponse contains the response from method L2IsolationDomainsClient.BeginGetArpEntries.
type L2IsolationDomainsClientGetArpEntriesResponse struct {
	// Show ARP entries response per network device
	Value map[string]*ARPProperties
}

// UnmarshalJSON implements the json.Unmarshaller interface for type L2IsolationDomainsClientGetArpEntriesResponse.
func (l *L2IsolationDomainsClientGetArpEntriesResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.Value)
}

// L2IsolationDomainsClientGetResponse contains the response from method L2IsolationDomainsClient.Get.
type L2IsolationDomainsClientGetResponse struct {
	L2IsolationDomain
}

// L2IsolationDomainsClientListByResourceGroupResponse contains the response from method L2IsolationDomainsClient.NewListByResourceGroupPager.
type L2IsolationDomainsClientListByResourceGroupResponse struct {
	L2IsolationDomainsListResult
}

// L2IsolationDomainsClientListBySubscriptionResponse contains the response from method L2IsolationDomainsClient.NewListBySubscriptionPager.
type L2IsolationDomainsClientListBySubscriptionResponse struct {
	L2IsolationDomainsListResult
}

// L2IsolationDomainsClientUpdateAdministrativeStateResponse contains the response from method L2IsolationDomainsClient.BeginUpdateAdministrativeState.
type L2IsolationDomainsClientUpdateAdministrativeStateResponse struct {
	// placeholder for future response values
}

// L2IsolationDomainsClientUpdateResponse contains the response from method L2IsolationDomainsClient.BeginUpdate.
type L2IsolationDomainsClientUpdateResponse struct {
	L2IsolationDomain
}

// L3IsolationDomainsClientClearArpTableResponse contains the response from method L3IsolationDomainsClient.BeginClearArpTable.
type L3IsolationDomainsClientClearArpTableResponse struct {
	// placeholder for future response values
}

// L3IsolationDomainsClientClearNeighborTableResponse contains the response from method L3IsolationDomainsClient.BeginClearNeighborTable.
type L3IsolationDomainsClientClearNeighborTableResponse struct {
	// placeholder for future response values
}

// L3IsolationDomainsClientCreateResponse contains the response from method L3IsolationDomainsClient.BeginCreate.
type L3IsolationDomainsClientCreateResponse struct {
	L3IsolationDomain
}

// L3IsolationDomainsClientDeleteResponse contains the response from method L3IsolationDomainsClient.BeginDelete.
type L3IsolationDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// L3IsolationDomainsClientGetResponse contains the response from method L3IsolationDomainsClient.Get.
type L3IsolationDomainsClientGetResponse struct {
	L3IsolationDomain
}

// L3IsolationDomainsClientListByResourceGroupResponse contains the response from method L3IsolationDomainsClient.NewListByResourceGroupPager.
type L3IsolationDomainsClientListByResourceGroupResponse struct {
	L3IsolationDomainsListResult
}

// L3IsolationDomainsClientListBySubscriptionResponse contains the response from method L3IsolationDomainsClient.NewListBySubscriptionPager.
type L3IsolationDomainsClientListBySubscriptionResponse struct {
	L3IsolationDomainsListResult
}

// L3IsolationDomainsClientUpdateAdministrativeStateResponse contains the response from method L3IsolationDomainsClient.BeginUpdateAdministrativeState.
type L3IsolationDomainsClientUpdateAdministrativeStateResponse struct {
	// placeholder for future response values
}

// L3IsolationDomainsClientUpdateOptionBAdministrativeStateResponse contains the response from method L3IsolationDomainsClient.BeginUpdateOptionBAdministrativeState.
type L3IsolationDomainsClientUpdateOptionBAdministrativeStateResponse struct {
	// placeholder for future response values
}

// L3IsolationDomainsClientUpdateResponse contains the response from method L3IsolationDomainsClient.BeginUpdate.
type L3IsolationDomainsClientUpdateResponse struct {
	L3IsolationDomain
}

// NetworkDeviceSKUsClientGetResponse contains the response from method NetworkDeviceSKUsClient.Get.
type NetworkDeviceSKUsClientGetResponse struct {
	NetworkDeviceSKU
}

// NetworkDeviceSKUsClientListBySubscriptionResponse contains the response from method NetworkDeviceSKUsClient.NewListBySubscriptionPager.
type NetworkDeviceSKUsClientListBySubscriptionResponse struct {
	NetworkDeviceSKUsListResult
}

// NetworkDevicesClientCreateResponse contains the response from method NetworkDevicesClient.BeginCreate.
type NetworkDevicesClientCreateResponse struct {
	NetworkDevice
}

// NetworkDevicesClientDeleteResponse contains the response from method NetworkDevicesClient.BeginDelete.
type NetworkDevicesClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkDevicesClientGenerateSupportPackageResponse contains the response from method NetworkDevicesClient.BeginGenerateSupportPackage.
type NetworkDevicesClientGenerateSupportPackageResponse struct {
	SupportPackageProperties
}

// NetworkDevicesClientGetDynamicInterfaceMapsResponse contains the response from method NetworkDevicesClient.BeginGetDynamicInterfaceMaps.
type NetworkDevicesClientGetDynamicInterfaceMapsResponse struct {
	// Layer 2 Network interfaces status
	GetDynamicInterfaceMapsPropertiesItemArray []*GetDynamicInterfaceMapsPropertiesItem
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkDevicesClientGetDynamicInterfaceMapsResponse.
func (n *NetworkDevicesClientGetDynamicInterfaceMapsResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.GetDynamicInterfaceMapsPropertiesItemArray)
}

// NetworkDevicesClientGetResponse contains the response from method NetworkDevicesClient.Get.
type NetworkDevicesClientGetResponse struct {
	NetworkDevice
}

// NetworkDevicesClientGetStaticInterfaceMapsResponse contains the response from method NetworkDevicesClient.BeginGetStaticInterfaceMaps.
type NetworkDevicesClientGetStaticInterfaceMapsResponse struct {
	// Layer 2 Network interfaces status
	GetStaticInterfaceMapsPropertiesItemArray []*GetStaticInterfaceMapsPropertiesItem
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkDevicesClientGetStaticInterfaceMapsResponse.
func (n *NetworkDevicesClientGetStaticInterfaceMapsResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.GetStaticInterfaceMapsPropertiesItemArray)
}

// NetworkDevicesClientGetStatusResponse contains the response from method NetworkDevicesClient.BeginGetStatus.
type NetworkDevicesClientGetStatusResponse struct {
	GetDeviceStatusProperties
}

// NetworkDevicesClientListByResourceGroupResponse contains the response from method NetworkDevicesClient.NewListByResourceGroupPager.
type NetworkDevicesClientListByResourceGroupResponse struct {
	NetworkDevicesListResult
}

// NetworkDevicesClientListBySubscriptionResponse contains the response from method NetworkDevicesClient.NewListBySubscriptionPager.
type NetworkDevicesClientListBySubscriptionResponse struct {
	NetworkDevicesListResult
}

// NetworkDevicesClientRebootResponse contains the response from method NetworkDevicesClient.BeginReboot.
type NetworkDevicesClientRebootResponse struct {
	// placeholder for future response values
}

// NetworkDevicesClientRestoreConfigResponse contains the response from method NetworkDevicesClient.BeginRestoreConfig.
type NetworkDevicesClientRestoreConfigResponse struct {
	// placeholder for future response values
}

// NetworkDevicesClientUpdatePowerCycleResponse contains the response from method NetworkDevicesClient.BeginUpdatePowerCycle.
type NetworkDevicesClientUpdatePowerCycleResponse struct {
	// placeholder for future response values
}

// NetworkDevicesClientUpdateResponse contains the response from method NetworkDevicesClient.BeginUpdate.
type NetworkDevicesClientUpdateResponse struct {
	NetworkDevice
}

// NetworkDevicesClientUpdateVersionResponse contains the response from method NetworkDevicesClient.BeginUpdateVersion.
type NetworkDevicesClientUpdateVersionResponse struct {
	// placeholder for future response values
}

// NetworkFabricControllersClientCreateResponse contains the response from method NetworkFabricControllersClient.BeginCreate.
type NetworkFabricControllersClientCreateResponse struct {
	NetworkFabricController
}

// NetworkFabricControllersClientDeleteResponse contains the response from method NetworkFabricControllersClient.BeginDelete.
type NetworkFabricControllersClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkFabricControllersClientDisableWorkloadManagementNetworkResponse contains the response from method NetworkFabricControllersClient.BeginDisableWorkloadManagementNetwork.
type NetworkFabricControllersClientDisableWorkloadManagementNetworkResponse struct {
	// placeholder for future response values
}

// NetworkFabricControllersClientEnableWorkloadManagementNetworkResponse contains the response from method NetworkFabricControllersClient.BeginEnableWorkloadManagementNetwork.
type NetworkFabricControllersClientEnableWorkloadManagementNetworkResponse struct {
	// placeholder for future response values
}

// NetworkFabricControllersClientGetResponse contains the response from method NetworkFabricControllersClient.Get.
type NetworkFabricControllersClientGetResponse struct {
	NetworkFabricController
}

// NetworkFabricControllersClientListByResourceGroupResponse contains the response from method NetworkFabricControllersClient.NewListByResourceGroupPager.
type NetworkFabricControllersClientListByResourceGroupResponse struct {
	NetworkFabricControllersListResult
}

// NetworkFabricControllersClientListBySubscriptionResponse contains the response from method NetworkFabricControllersClient.NewListBySubscriptionPager.
type NetworkFabricControllersClientListBySubscriptionResponse struct {
	NetworkFabricControllersListResult
}

// NetworkFabricControllersClientUpdateResponse contains the response from method NetworkFabricControllersClient.BeginUpdate.
type NetworkFabricControllersClientUpdateResponse struct {
	NetworkFabricController
}

// NetworkFabricSKUsClientGetResponse contains the response from method NetworkFabricSKUsClient.Get.
type NetworkFabricSKUsClientGetResponse struct {
	NetworkFabricSKU
}

// NetworkFabricSKUsClientListBySubscriptionResponse contains the response from method NetworkFabricSKUsClient.NewListBySubscriptionPager.
type NetworkFabricSKUsClientListBySubscriptionResponse struct {
	NetworkFabricSKUsListResult
}

// NetworkFabricsClientCreateResponse contains the response from method NetworkFabricsClient.BeginCreate.
type NetworkFabricsClientCreateResponse struct {
	NetworkFabric
}

// NetworkFabricsClientDeleteResponse contains the response from method NetworkFabricsClient.BeginDelete.
type NetworkFabricsClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkFabricsClientDeprovisionResponse contains the response from method NetworkFabricsClient.BeginDeprovision.
type NetworkFabricsClientDeprovisionResponse struct {
	// placeholder for future response values
}

// NetworkFabricsClientGetResponse contains the response from method NetworkFabricsClient.Get.
type NetworkFabricsClientGetResponse struct {
	NetworkFabric
}

// NetworkFabricsClientListByResourceGroupResponse contains the response from method NetworkFabricsClient.NewListByResourceGroupPager.
type NetworkFabricsClientListByResourceGroupResponse struct {
	NetworkFabricsListResult
}

// NetworkFabricsClientListBySubscriptionResponse contains the response from method NetworkFabricsClient.NewListBySubscriptionPager.
type NetworkFabricsClientListBySubscriptionResponse struct {
	NetworkFabricsListResult
}

// NetworkFabricsClientProvisionResponse contains the response from method NetworkFabricsClient.BeginProvision.
type NetworkFabricsClientProvisionResponse struct {
	// placeholder for future response values
}

// NetworkFabricsClientUpdateResponse contains the response from method NetworkFabricsClient.BeginUpdate.
type NetworkFabricsClientUpdateResponse struct {
	NetworkFabric
}

// NetworkInterfacesClientCreateResponse contains the response from method NetworkInterfacesClient.BeginCreate.
type NetworkInterfacesClientCreateResponse struct {
	NetworkInterface
}

// NetworkInterfacesClientDeleteResponse contains the response from method NetworkInterfacesClient.BeginDelete.
type NetworkInterfacesClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkInterfacesClientGetResponse contains the response from method NetworkInterfacesClient.Get.
type NetworkInterfacesClientGetResponse struct {
	NetworkInterface
}

// NetworkInterfacesClientGetStatusResponse contains the response from method NetworkInterfacesClient.BeginGetStatus.
type NetworkInterfacesClientGetStatusResponse struct {
	InterfaceStatus
}

// NetworkInterfacesClientListResponse contains the response from method NetworkInterfacesClient.NewListPager.
type NetworkInterfacesClientListResponse struct {
	NetworkInterfacesList
}

// NetworkInterfacesClientUpdateAdministrativeStateResponse contains the response from method NetworkInterfacesClient.BeginUpdateAdministrativeState.
type NetworkInterfacesClientUpdateAdministrativeStateResponse struct {
	// placeholder for future response values
}

// NetworkInterfacesClientUpdateResponse contains the response from method NetworkInterfacesClient.BeginUpdate.
type NetworkInterfacesClientUpdateResponse struct {
	NetworkInterface
}

// NetworkRackSKUsClientGetResponse contains the response from method NetworkRackSKUsClient.Get.
type NetworkRackSKUsClientGetResponse struct {
	NetworkRackSKU
}

// NetworkRackSKUsClientListBySubscriptionResponse contains the response from method NetworkRackSKUsClient.NewListBySubscriptionPager.
type NetworkRackSKUsClientListBySubscriptionResponse struct {
	NetworkRackSKUsListResult
}

// NetworkRacksClientCreateResponse contains the response from method NetworkRacksClient.BeginCreate.
type NetworkRacksClientCreateResponse struct {
	NetworkRack
}

// NetworkRacksClientDeleteResponse contains the response from method NetworkRacksClient.BeginDelete.
type NetworkRacksClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkRacksClientGetResponse contains the response from method NetworkRacksClient.Get.
type NetworkRacksClientGetResponse struct {
	NetworkRack
}

// NetworkRacksClientListByResourceGroupResponse contains the response from method NetworkRacksClient.NewListByResourceGroupPager.
type NetworkRacksClientListByResourceGroupResponse struct {
	NetworkRacksListResult
}

// NetworkRacksClientListBySubscriptionResponse contains the response from method NetworkRacksClient.NewListBySubscriptionPager.
type NetworkRacksClientListBySubscriptionResponse struct {
	NetworkRacksListResult
}

// NetworkRacksClientUpdateResponse contains the response from method NetworkRacksClient.BeginUpdate.
type NetworkRacksClientUpdateResponse struct {
	NetworkRack
}

// NetworkToNetworkInterconnectsClientCreateResponse contains the response from method NetworkToNetworkInterconnectsClient.BeginCreate.
type NetworkToNetworkInterconnectsClientCreateResponse struct {
	NetworkToNetworkInterconnect
}

// NetworkToNetworkInterconnectsClientDeleteResponse contains the response from method NetworkToNetworkInterconnectsClient.BeginDelete.
type NetworkToNetworkInterconnectsClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkToNetworkInterconnectsClientGetResponse contains the response from method NetworkToNetworkInterconnectsClient.Get.
type NetworkToNetworkInterconnectsClientGetResponse struct {
	NetworkToNetworkInterconnect
}

// NetworkToNetworkInterconnectsClientListResponse contains the response from method NetworkToNetworkInterconnectsClient.NewListPager.
type NetworkToNetworkInterconnectsClientListResponse struct {
	NetworkToNetworkInterconnectsList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// RoutePoliciesClientCreateResponse contains the response from method RoutePoliciesClient.BeginCreate.
type RoutePoliciesClientCreateResponse struct {
	RoutePolicy
}

// RoutePoliciesClientDeleteResponse contains the response from method RoutePoliciesClient.BeginDelete.
type RoutePoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// RoutePoliciesClientGetResponse contains the response from method RoutePoliciesClient.Get.
type RoutePoliciesClientGetResponse struct {
	RoutePolicy
}

// RoutePoliciesClientListByResourceGroupResponse contains the response from method RoutePoliciesClient.NewListByResourceGroupPager.
type RoutePoliciesClientListByResourceGroupResponse struct {
	RoutePoliciesListResult
}

// RoutePoliciesClientListBySubscriptionResponse contains the response from method RoutePoliciesClient.NewListBySubscriptionPager.
type RoutePoliciesClientListBySubscriptionResponse struct {
	RoutePoliciesListResult
}

// RoutePoliciesClientUpdateResponse contains the response from method RoutePoliciesClient.BeginUpdate.
type RoutePoliciesClientUpdateResponse struct {
	RoutePolicy
}
