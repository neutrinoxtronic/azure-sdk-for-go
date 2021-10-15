//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package logz

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/logz/mgmt/2020-10-01/logz"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type LiftrResourceCategories = original.LiftrResourceCategories

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = original.LiftrResourceCategoriesMonitorLogs
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = original.LiftrResourceCategoriesUnknown
)

type ManagedIdentityTypes = original.ManagedIdentityTypes

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = original.ManagedIdentityTypesSystemAssigned
	ManagedIdentityTypesUserAssigned   ManagedIdentityTypes = original.ManagedIdentityTypesUserAssigned
)

type MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatus

const (
	MarketplaceSubscriptionStatusActive    MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatusActive
	MarketplaceSubscriptionStatusSuspended MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatusSuspended
)

type MonitoringStatus = original.MonitoringStatus

const (
	MonitoringStatusDisabled MonitoringStatus = original.MonitoringStatusDisabled
	MonitoringStatusEnabled  MonitoringStatus = original.MonitoringStatusEnabled
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateNotSpecified ProvisioningState = original.ProvisioningStateNotSpecified
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type SingleSignOnStates = original.SingleSignOnStates

const (
	SingleSignOnStatesDisable  SingleSignOnStates = original.SingleSignOnStatesDisable
	SingleSignOnStatesEnable   SingleSignOnStates = original.SingleSignOnStatesEnable
	SingleSignOnStatesExisting SingleSignOnStates = original.SingleSignOnStatesExisting
	SingleSignOnStatesInitial  SingleSignOnStates = original.SingleSignOnStatesInitial
)

type TagAction = original.TagAction

const (
	TagActionExclude TagAction = original.TagActionExclude
	TagActionInclude TagAction = original.TagActionInclude
)

type UserRole = original.UserRole

const (
	UserRoleAdmin UserRole = original.UserRoleAdmin
	UserRoleNone  UserRole = original.UserRoleNone
	UserRoleUser  UserRole = original.UserRoleUser
)

type VMHostUpdateStates = original.VMHostUpdateStates

const (
	VMHostUpdateStatesDelete  VMHostUpdateStates = original.VMHostUpdateStatesDelete
	VMHostUpdateStatesInstall VMHostUpdateStates = original.VMHostUpdateStatesInstall
)

type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type FilteringTag = original.FilteringTag
type IdentityProperties = original.IdentityProperties
type LogRules = original.LogRules
type MonitorClient = original.MonitorClient
type MonitorProperties = original.MonitorProperties
type MonitorResource = original.MonitorResource
type MonitorResourceListResponse = original.MonitorResourceListResponse
type MonitorResourceListResponseIterator = original.MonitorResourceListResponseIterator
type MonitorResourceListResponsePage = original.MonitorResourceListResponsePage
type MonitorResourceUpdateParameters = original.MonitorResourceUpdateParameters
type MonitorUpdateProperties = original.MonitorUpdateProperties
type MonitoredResource = original.MonitoredResource
type MonitoredResourceListResponse = original.MonitoredResourceListResponse
type MonitoredResourceListResponseIterator = original.MonitoredResourceListResponseIterator
type MonitoredResourceListResponsePage = original.MonitoredResourceListResponsePage
type MonitoringTagRules = original.MonitoringTagRules
type MonitoringTagRulesListResponse = original.MonitoringTagRulesListResponse
type MonitoringTagRulesListResponseIterator = original.MonitoringTagRulesListResponseIterator
type MonitoringTagRulesListResponsePage = original.MonitoringTagRulesListResponsePage
type MonitoringTagRulesProperties = original.MonitoringTagRulesProperties
type MonitorsClient = original.MonitorsClient
type MonitorsCreateFuture = original.MonitorsCreateFuture
type MonitorsDeleteFuture = original.MonitorsDeleteFuture
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type OrganizationProperties = original.OrganizationProperties
type PlanData = original.PlanData
type SingleSignOnClient = original.SingleSignOnClient
type SingleSignOnCreateOrUpdateFuture = original.SingleSignOnCreateOrUpdateFuture
type SingleSignOnProperties = original.SingleSignOnProperties
type SingleSignOnResource = original.SingleSignOnResource
type SingleSignOnResourceListResponse = original.SingleSignOnResourceListResponse
type SingleSignOnResourceListResponseIterator = original.SingleSignOnResourceListResponseIterator
type SingleSignOnResourceListResponsePage = original.SingleSignOnResourceListResponsePage
type SubAccountClient = original.SubAccountClient
type SubAccountCreateFuture = original.SubAccountCreateFuture
type SubAccountDeleteFuture = original.SubAccountDeleteFuture
type SubAccountTagRulesClient = original.SubAccountTagRulesClient
type SystemData = original.SystemData
type TagRulesClient = original.TagRulesClient
type UserInfo = original.UserInfo
type UserRoleListResponse = original.UserRoleListResponse
type UserRoleListResponseIterator = original.UserRoleListResponseIterator
type UserRoleListResponsePage = original.UserRoleListResponsePage
type UserRoleRequest = original.UserRoleRequest
type UserRoleResponse = original.UserRoleResponse
type VMExtensionPayload = original.VMExtensionPayload
type VMHostUpdateRequest = original.VMHostUpdateRequest
type VMResources = original.VMResources
type VMResourcesListResponse = original.VMResourcesListResponse
type VMResourcesListResponseIterator = original.VMResourcesListResponseIterator
type VMResourcesListResponsePage = original.VMResourcesListResponsePage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewMonitorClient(subscriptionID string) MonitorClient {
	return original.NewMonitorClient(subscriptionID)
}
func NewMonitorClientWithBaseURI(baseURI string, subscriptionID string) MonitorClient {
	return original.NewMonitorClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorResourceListResponseIterator(page MonitorResourceListResponsePage) MonitorResourceListResponseIterator {
	return original.NewMonitorResourceListResponseIterator(page)
}
func NewMonitorResourceListResponsePage(cur MonitorResourceListResponse, getNextPage func(context.Context, MonitorResourceListResponse) (MonitorResourceListResponse, error)) MonitorResourceListResponsePage {
	return original.NewMonitorResourceListResponsePage(cur, getNextPage)
}
func NewMonitoredResourceListResponseIterator(page MonitoredResourceListResponsePage) MonitoredResourceListResponseIterator {
	return original.NewMonitoredResourceListResponseIterator(page)
}
func NewMonitoredResourceListResponsePage(cur MonitoredResourceListResponse, getNextPage func(context.Context, MonitoredResourceListResponse) (MonitoredResourceListResponse, error)) MonitoredResourceListResponsePage {
	return original.NewMonitoredResourceListResponsePage(cur, getNextPage)
}
func NewMonitoringTagRulesListResponseIterator(page MonitoringTagRulesListResponsePage) MonitoringTagRulesListResponseIterator {
	return original.NewMonitoringTagRulesListResponseIterator(page)
}
func NewMonitoringTagRulesListResponsePage(cur MonitoringTagRulesListResponse, getNextPage func(context.Context, MonitoringTagRulesListResponse) (MonitoringTagRulesListResponse, error)) MonitoringTagRulesListResponsePage {
	return original.NewMonitoringTagRulesListResponsePage(cur, getNextPage)
}
func NewMonitorsClient(subscriptionID string) MonitorsClient {
	return original.NewMonitorsClient(subscriptionID)
}
func NewMonitorsClientWithBaseURI(baseURI string, subscriptionID string) MonitorsClient {
	return original.NewMonitorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSingleSignOnClient(subscriptionID string) SingleSignOnClient {
	return original.NewSingleSignOnClient(subscriptionID)
}
func NewSingleSignOnClientWithBaseURI(baseURI string, subscriptionID string) SingleSignOnClient {
	return original.NewSingleSignOnClientWithBaseURI(baseURI, subscriptionID)
}
func NewSingleSignOnResourceListResponseIterator(page SingleSignOnResourceListResponsePage) SingleSignOnResourceListResponseIterator {
	return original.NewSingleSignOnResourceListResponseIterator(page)
}
func NewSingleSignOnResourceListResponsePage(cur SingleSignOnResourceListResponse, getNextPage func(context.Context, SingleSignOnResourceListResponse) (SingleSignOnResourceListResponse, error)) SingleSignOnResourceListResponsePage {
	return original.NewSingleSignOnResourceListResponsePage(cur, getNextPage)
}
func NewSubAccountClient(subscriptionID string) SubAccountClient {
	return original.NewSubAccountClient(subscriptionID)
}
func NewSubAccountClientWithBaseURI(baseURI string, subscriptionID string) SubAccountClient {
	return original.NewSubAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubAccountTagRulesClient(subscriptionID string) SubAccountTagRulesClient {
	return original.NewSubAccountTagRulesClient(subscriptionID)
}
func NewSubAccountTagRulesClientWithBaseURI(baseURI string, subscriptionID string) SubAccountTagRulesClient {
	return original.NewSubAccountTagRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagRulesClient(subscriptionID string) TagRulesClient {
	return original.NewTagRulesClient(subscriptionID)
}
func NewTagRulesClientWithBaseURI(baseURI string, subscriptionID string) TagRulesClient {
	return original.NewTagRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserRoleListResponseIterator(page UserRoleListResponsePage) UserRoleListResponseIterator {
	return original.NewUserRoleListResponseIterator(page)
}
func NewUserRoleListResponsePage(cur UserRoleListResponse, getNextPage func(context.Context, UserRoleListResponse) (UserRoleListResponse, error)) UserRoleListResponsePage {
	return original.NewUserRoleListResponsePage(cur, getNextPage)
}
func NewVMResourcesListResponseIterator(page VMResourcesListResponsePage) VMResourcesListResponseIterator {
	return original.NewVMResourcesListResponseIterator(page)
}
func NewVMResourcesListResponsePage(cur VMResourcesListResponse, getNextPage func(context.Context, VMResourcesListResponse) (VMResourcesListResponse, error)) VMResourcesListResponsePage {
	return original.NewVMResourcesListResponsePage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return original.PossibleLiftrResourceCategoriesValues()
}
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return original.PossibleManagedIdentityTypesValues()
}
func PossibleMarketplaceSubscriptionStatusValues() []MarketplaceSubscriptionStatus {
	return original.PossibleMarketplaceSubscriptionStatusValues()
}
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return original.PossibleMonitoringStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSingleSignOnStatesValues() []SingleSignOnStates {
	return original.PossibleSingleSignOnStatesValues()
}
func PossibleTagActionValues() []TagAction {
	return original.PossibleTagActionValues()
}
func PossibleUserRoleValues() []UserRole {
	return original.PossibleUserRoleValues()
}
func PossibleVMHostUpdateStatesValues() []VMHostUpdateStates {
	return original.PossibleVMHostUpdateStatesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}