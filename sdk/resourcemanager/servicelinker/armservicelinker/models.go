//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicelinker

import "time"

// AuthInfoBaseClassification provides polymorphic access to related types.
// Call the interface's GetAuthInfoBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AuthInfoBase, *SecretAuthInfo, *ServicePrincipalCertificateAuthInfo, *ServicePrincipalSecretAuthInfo, *SystemAssignedIdentityAuthInfo,
// - *UserAssignedIdentityAuthInfo
type AuthInfoBaseClassification interface {
	// GetAuthInfoBase returns the AuthInfoBase content of the underlying type.
	GetAuthInfoBase() *AuthInfoBase
}

// AuthInfoBase - The authentication info
type AuthInfoBase struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type AuthInfoBase.
func (a *AuthInfoBase) GetAuthInfoBase() *AuthInfoBase { return a }

// AzureKeyVaultProperties - The resource properties when type is Azure Key Vault
type AzureKeyVaultProperties struct {
	// REQUIRED; The azure resource type.
	Type *AzureResourceType

	// True if connect via Kubernetes CSI Driver.
	ConnectAsKubernetesCsiDriver *bool
}

// GetAzureResourcePropertiesBase implements the AzureResourcePropertiesBaseClassification interface for type AzureKeyVaultProperties.
func (a *AzureKeyVaultProperties) GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase {
	return &AzureResourcePropertiesBase{
		Type: a.Type,
	}
}

// AzureResource - The azure resource info when target service type is AzureResource
type AzureResource struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType

	// The Id of azure resource.
	ID *string

	// The azure resource connection related properties.
	ResourceProperties AzureResourcePropertiesBaseClassification
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type AzureResource.
func (a *AzureResource) GetTargetServiceBase() *TargetServiceBase {
	return &TargetServiceBase{
		Type: a.Type,
	}
}

// AzureResourcePropertiesBaseClassification provides polymorphic access to related types.
// Call the interface's GetAzureResourcePropertiesBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureKeyVaultProperties, *AzureResourcePropertiesBase
type AzureResourcePropertiesBaseClassification interface {
	// GetAzureResourcePropertiesBase returns the AzureResourcePropertiesBase content of the underlying type.
	GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase
}

// AzureResourcePropertiesBase - The azure resource properties
type AzureResourcePropertiesBase struct {
	// REQUIRED; The azure resource type.
	Type *AzureResourceType
}

// GetAzureResourcePropertiesBase implements the AzureResourcePropertiesBaseClassification interface for type AzureResourcePropertiesBase.
func (a *AzureResourcePropertiesBase) GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase {
	return a
}

// ConfluentBootstrapServer - The service properties when target service type is ConfluentBootstrapServer
type ConfluentBootstrapServer struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType

	// The endpoint of service.
	Endpoint *string
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type ConfluentBootstrapServer.
func (c *ConfluentBootstrapServer) GetTargetServiceBase() *TargetServiceBase {
	return &TargetServiceBase{
		Type: c.Type,
	}
}

// ConfluentSchemaRegistry - The service properties when target service type is ConfluentSchemaRegistry
type ConfluentSchemaRegistry struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType

	// The endpoint of service.
	Endpoint *string
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type ConfluentSchemaRegistry.
func (c *ConfluentSchemaRegistry) GetTargetServiceBase() *TargetServiceBase {
	return &TargetServiceBase{
		Type: c.Type,
	}
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// KeyVaultSecretReferenceSecretInfo - The secret info when type is keyVaultSecretReference. It's for scenario that user provides
// a secret stored in user's keyvault and source is Azure Kubernetes. The key Vault's resource id is linked to
// secretStore.keyVaultId.
type KeyVaultSecretReferenceSecretInfo struct {
	// REQUIRED; The secret type.
	SecretType *SecretType

	// Name of the Key Vault secret.
	Name *string

	// Version of the Key Vault secret.
	Version *string
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type KeyVaultSecretReferenceSecretInfo.
func (k *KeyVaultSecretReferenceSecretInfo) GetSecretInfoBase() *SecretInfoBase {
	return &SecretInfoBase{
		SecretType: k.SecretType,
	}
}

// KeyVaultSecretURISecretInfo - The secret info when type is keyVaultSecretUri. It's for scenario that user provides a secret
// stored in user's keyvault and source is Web App, Spring Cloud or Container App.
type KeyVaultSecretURISecretInfo struct {
	// REQUIRED; The secret type.
	SecretType *SecretType

	// URI to the keyvault secret
	Value *string
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type KeyVaultSecretURISecretInfo.
func (k *KeyVaultSecretURISecretInfo) GetSecretInfoBase() *SecretInfoBase {
	return &SecretInfoBase{
		SecretType: k.SecretType,
	}
}

// LinkerClientBeginCreateOrUpdateOptions contains the optional parameters for the LinkerClient.BeginCreateOrUpdate method.
type LinkerClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientBeginDeleteOptions contains the optional parameters for the LinkerClient.BeginDelete method.
type LinkerClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientBeginUpdateOptions contains the optional parameters for the LinkerClient.BeginUpdate method.
type LinkerClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientBeginValidateOptions contains the optional parameters for the LinkerClient.BeginValidate method.
type LinkerClientBeginValidateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LinkerClientGetOptions contains the optional parameters for the LinkerClient.Get method.
type LinkerClientGetOptions struct {
	// placeholder for future optional parameters
}

// LinkerClientListConfigurationsOptions contains the optional parameters for the LinkerClient.ListConfigurations method.
type LinkerClientListConfigurationsOptions struct {
	// placeholder for future optional parameters
}

// LinkerClientListOptions contains the optional parameters for the LinkerClient.NewListPager method.
type LinkerClientListOptions struct {
	// placeholder for future optional parameters
}

// LinkerList - The list of Linker.
type LinkerList struct {
	// The link used to get the next page of Linker list.
	NextLink *string

	// The list of Linkers.
	Value []*LinkerResource
}

// LinkerPatch - A linker to be updated.
type LinkerPatch struct {
	// Linker properties
	Properties *LinkerProperties
}

// LinkerProperties - The properties of the linker.
type LinkerProperties struct {
	// The authentication type.
	AuthInfo AuthInfoBaseClassification

	// The application client type
	ClientType *ClientType

	// connection scope in source service.
	Scope *string

	// An option to store secret value in secure place
	SecretStore *SecretStore

	// The target service properties
	TargetService TargetServiceBaseClassification

	// The VNet solution.
	VNetSolution *VNetSolution

	// READ-ONLY; The provisioning state.
	ProvisioningState *string
}

// LinkerResource - Linker of source and target resource
type LinkerResource struct {
	// REQUIRED; The properties of the linker.
	Properties *LinkerProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system data.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SecretAuthInfo - The authentication info when authType is secret
type SecretAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType

	// Username or account name for secret auth.
	Name *string

	// Password or key vault secret for secret auth.
	SecretInfo SecretInfoBaseClassification
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type SecretAuthInfo.
func (s *SecretAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// SecretInfoBaseClassification provides polymorphic access to related types.
// Call the interface's GetSecretInfoBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *KeyVaultSecretReferenceSecretInfo, *KeyVaultSecretURISecretInfo, *SecretInfoBase, *ValueSecretInfo
type SecretInfoBaseClassification interface {
	// GetSecretInfoBase returns the SecretInfoBase content of the underlying type.
	GetSecretInfoBase() *SecretInfoBase
}

// SecretInfoBase - The secret info
type SecretInfoBase struct {
	// REQUIRED; The secret type.
	SecretType *SecretType
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type SecretInfoBase.
func (s *SecretInfoBase) GetSecretInfoBase() *SecretInfoBase { return s }

// SecretStore - An option to store secret value in secure place
type SecretStore struct {
	// The key vault id to store secret
	KeyVaultID *string
}

// ServicePrincipalCertificateAuthInfo - The authentication info when authType is servicePrincipal certificate
type ServicePrincipalCertificateAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType

	// REQUIRED; ServicePrincipal certificate for servicePrincipal auth.
	Certificate *string

	// REQUIRED; Application clientId for servicePrincipal auth.
	ClientID *string

	// REQUIRED; Principal Id for servicePrincipal auth.
	PrincipalID *string
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type ServicePrincipalCertificateAuthInfo.
func (s *ServicePrincipalCertificateAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// ServicePrincipalSecretAuthInfo - The authentication info when authType is servicePrincipal secret
type ServicePrincipalSecretAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType

	// REQUIRED; ServicePrincipal application clientId for servicePrincipal auth.
	ClientID *string

	// REQUIRED; Principal Id for servicePrincipal auth.
	PrincipalID *string

	// REQUIRED; Secret for servicePrincipal auth.
	Secret *string
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type ServicePrincipalSecretAuthInfo.
func (s *ServicePrincipalSecretAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// SourceConfiguration - A configuration item for source resource
type SourceConfiguration struct {
	// The name of setting.
	Name *string

	// The value of setting
	Value *string
}

// SourceConfigurationResult - Configurations for source resource, include appSettings, connectionString and serviceBindings
type SourceConfigurationResult struct {
	// The configuration properties for source resource.
	Configurations []*SourceConfiguration
}

// SystemAssignedIdentityAuthInfo - The authentication info when authType is systemAssignedIdentity
type SystemAssignedIdentityAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type SystemAssignedIdentityAuthInfo.
func (s *SystemAssignedIdentityAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: s.AuthType,
	}
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TargetServiceBaseClassification provides polymorphic access to related types.
// Call the interface's GetTargetServiceBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureResource, *ConfluentBootstrapServer, *ConfluentSchemaRegistry, *TargetServiceBase
type TargetServiceBaseClassification interface {
	// GetTargetServiceBase returns the TargetServiceBase content of the underlying type.
	GetTargetServiceBase() *TargetServiceBase
}

// TargetServiceBase - The target service properties
type TargetServiceBase struct {
	// REQUIRED; The target service type.
	Type *TargetServiceType
}

// GetTargetServiceBase implements the TargetServiceBaseClassification interface for type TargetServiceBase.
func (t *TargetServiceBase) GetTargetServiceBase() *TargetServiceBase { return t }

// UserAssignedIdentityAuthInfo - The authentication info when authType is userAssignedIdentity
type UserAssignedIdentityAuthInfo struct {
	// REQUIRED; The authentication type.
	AuthType *AuthType

	// Client Id for userAssignedIdentity.
	ClientID *string

	// Subscription id for userAssignedIdentity.
	SubscriptionID *string
}

// GetAuthInfoBase implements the AuthInfoBaseClassification interface for type UserAssignedIdentityAuthInfo.
func (u *UserAssignedIdentityAuthInfo) GetAuthInfoBase() *AuthInfoBase {
	return &AuthInfoBase{
		AuthType: u.AuthType,
	}
}

// VNetSolution - The VNet solution for linker
type VNetSolution struct {
	// Type of VNet solution.
	Type *VNetSolutionType
}

// ValidateOperationResult - The validation operation result for a linker.
type ValidateOperationResult struct {
	// The validation result detail.
	Properties *ValidateResult

	// Validated linker id.
	ResourceID *string

	// Validation operation status.
	Status *string
}

// ValidateResult - The validation result for a linker.
type ValidateResult struct {
	// The authentication type.
	AuthType *AuthType

	// A boolean value indicating whether the connection is available or not
	IsConnectionAvailable *bool

	// The linker name.
	LinkerName *string

	// The end time of the validation report.
	ReportEndTimeUTC *time.Time

	// The start time of the validation report.
	ReportStartTimeUTC *time.Time

	// The resource id of the linker source application.
	SourceID *string

	// The resource Id of target service.
	TargetID *string

	// The detail of validation result
	ValidationDetail []*ValidationResultItem
}

// ValidationResultItem - The validation item for a linker.
type ValidationResultItem struct {
	// The display name of validation item
	Description *string

	// The error code of validation result
	ErrorCode *string

	// The error message of validation result
	ErrorMessage *string

	// The validation item name.
	Name *string

	// The result of validation
	Result *ValidationResultStatus
}

// ValueSecretInfo - The secret info when type is rawValue. It's for scenarios that user input the secret.
type ValueSecretInfo struct {
	// REQUIRED; The secret type.
	SecretType *SecretType

	// The actual value of the secret.
	Value *string
}

// GetSecretInfoBase implements the SecretInfoBaseClassification interface for type ValueSecretInfo.
func (v *ValueSecretInfo) GetSecretInfoBase() *SecretInfoBase {
	return &SecretInfoBase{
		SecretType: v.SecretType,
	}
}
