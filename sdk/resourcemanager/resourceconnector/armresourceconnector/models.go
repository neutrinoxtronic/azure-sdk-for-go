//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourceconnector

import "time"

// Appliances definition.
type Appliance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Identity for the resource.
	Identity *Identity

	// The set of properties specific to an Appliance
	Properties *ApplianceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ApplianceCredentialKubeconfig - Cluster User Credential appliance.
type ApplianceCredentialKubeconfig struct {
	// READ-ONLY; Name which contains the role of the kubeconfig.
	Name *AccessProfileType

	// READ-ONLY; Contains the kubeconfig value.
	Value *string
}

// ApplianceGetTelemetryConfigResult - The Get Telemetry Config Result appliance.
type ApplianceGetTelemetryConfigResult struct {
	// READ-ONLY; Telemetry instrumentation key.
	TelemetryInstrumentationKey *string
}

// ApplianceListCredentialResults - The List Cluster User Credential appliance.
type ApplianceListCredentialResults struct {
	// READ-ONLY; Contains the REP (rendezvous endpoint) and “Listener” access token from notification service (NS).
	HybridConnectionConfig *HybridConnectionConfig

	// READ-ONLY; The list of appliance kubeconfigs.
	Kubeconfigs []*ApplianceCredentialKubeconfig
}

// ApplianceListKeysResults - The List Cluster Keys Results appliance.
type ApplianceListKeysResults struct {
	// READ-ONLY; Map of artifacts that contains a list of ArtifactProfile used to upload artifacts such as logs.
	ArtifactProfiles map[string]*ArtifactProfile

	// READ-ONLY; The list of appliance kubeconfigs.
	Kubeconfigs []*ApplianceCredentialKubeconfig

	// READ-ONLY; Map of Customer User Public, Private SSH Keys and Certificate when available.
	SSHKeys map[string]*SSHKey
}

// ApplianceListResult - The List Appliances operation response.
type ApplianceListResult struct {
	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string

	// READ-ONLY; The list of Appliances.
	Value []*Appliance
}

// ApplianceOperation - Appliances operation.
type ApplianceOperation struct {
	// Describes the properties of an Appliances Operation Value Display.
	Display *ApplianceOperationValueDisplay

	// READ-ONLY; Is this Operation a data plane operation
	IsDataAction *bool

	// READ-ONLY; The name of the compute operation.
	Name *string

	// READ-ONLY; The origin of the compute operation.
	Origin *string
}

// ApplianceOperationValueDisplay - Describes the properties of an Appliances Operation Value Display.
type ApplianceOperationValueDisplay struct {
	// READ-ONLY; The description of the operation.
	Description *string

	// READ-ONLY; The display name of the compute operation.
	Operation *string

	// READ-ONLY; The resource provider for the operation.
	Provider *string

	// READ-ONLY; The display name of the resource the operation applies to.
	Resource *string
}

// ApplianceOperationsList - Lists of Appliances operations.
type ApplianceOperationsList struct {
	// REQUIRED; Array of applianceOperation
	Value []*ApplianceOperation

	// Next page of operations.
	NextLink *string
}

// ApplianceProperties - Properties for an appliance.
type ApplianceProperties struct {
	// Represents a supported Fabric/Infra. (AKSEdge etc…).
	Distro *Distro

	// Contains infrastructure information about the Appliance
	InfrastructureConfig *AppliancePropertiesInfrastructureConfig

	// Certificates pair used to download MSI certificate from HIS. Can only be set once.
	PublicKey *string

	// Version of the Appliance
	Version *string

	// READ-ONLY; The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string

	// READ-ONLY; Appliance’s health and state of connection to on-prem
	Status *Status
}

// AppliancePropertiesInfrastructureConfig - Contains infrastructure information about the Appliance
type AppliancePropertiesInfrastructureConfig struct {
	// Information about the connected appliance.
	Provider *Provider
}

// AppliancesClientBeginCreateOrUpdateOptions contains the optional parameters for the AppliancesClient.BeginCreateOrUpdate
// method.
type AppliancesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AppliancesClientBeginDeleteOptions contains the optional parameters for the AppliancesClient.BeginDelete method.
type AppliancesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AppliancesClientGetOptions contains the optional parameters for the AppliancesClient.Get method.
type AppliancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientGetTelemetryConfigOptions contains the optional parameters for the AppliancesClient.GetTelemetryConfig
// method.
type AppliancesClientGetTelemetryConfigOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientGetUpgradeGraphOptions contains the optional parameters for the AppliancesClient.GetUpgradeGraph method.
type AppliancesClientGetUpgradeGraphOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientListByResourceGroupOptions contains the optional parameters for the AppliancesClient.NewListByResourceGroupPager
// method.
type AppliancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientListBySubscriptionOptions contains the optional parameters for the AppliancesClient.NewListBySubscriptionPager
// method.
type AppliancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientListClusterUserCredentialOptions contains the optional parameters for the AppliancesClient.ListClusterUserCredential
// method.
type AppliancesClientListClusterUserCredentialOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientListKeysOptions contains the optional parameters for the AppliancesClient.ListKeys method.
type AppliancesClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientListOperationsOptions contains the optional parameters for the AppliancesClient.NewListOperationsPager
// method.
type AppliancesClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// AppliancesClientUpdateOptions contains the optional parameters for the AppliancesClient.Update method.
type AppliancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ArtifactProfile - Appliance ArtifactProfile definition.
type ArtifactProfile struct {
	// READ-ONLY; Endpoint is the URL to upload artifacts to.
	Endpoint *string
}

// HybridConnectionConfig - Contains the REP (rendezvous endpoint) and “Listener” access token from notification service (NS).
type HybridConnectionConfig struct {
	// READ-ONLY; Timestamp when this token will be expired.
	ExpirationTime *int64

	// READ-ONLY; Name of the connection
	HybridConnectionName *string

	// READ-ONLY; Name of the notification service.
	Relay *string

	// READ-ONLY; Listener access token
	Token *string
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *ResourceIdentityType

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// PatchableAppliance - The Appliances patchable resource definition.
type PatchableAppliance struct {
	// Resource tags
	Tags map[string]*string
}

// SSHKey - Appliance SSHKey definition.
type SSHKey struct {
	// READ-ONLY; Certificate associated with the public key if the key is signed.
	Certificate *string

	// READ-ONLY; Certificate creation timestamp (Unix).
	CreationTimeStamp *int64

	// READ-ONLY; Certificate expiration timestamp (Unix).
	ExpirationTimeStamp *int64

	// READ-ONLY; Private Key.
	PrivateKey *string

	// READ-ONLY; Public Key.
	PublicKey *string
}

// SupportedVersion - The SupportedVersion object for appliance.
type SupportedVersion struct {
	// READ-ONLY; This is the metadata of the supported newer version.
	Metadata *SupportedVersionMetadata

	// READ-ONLY; The newer version available for upgrade.
	Version *string
}

// SupportedVersionCatalogVersion - The SupportedVersionCatalogVersion object for appliance.
type SupportedVersionCatalogVersion struct {
	// READ-ONLY; The newer supported version catalog version data.
	Data *SupportedVersionCatalogVersionData

	// READ-ONLY; The catalog version name for the version available for upgrade.
	Name *string

	// READ-ONLY; The catalog version namespace for the version available for upgrade.
	Namespace *string
}

// SupportedVersionCatalogVersionData - The SupportedVersionCatalogVersionData object for appliance.
type SupportedVersionCatalogVersionData struct {
	// READ-ONLY; The image audience name for the version available for upgrade.
	Audience *string

	// READ-ONLY; The image catalog name for the version available for upgrade.
	Catalog *string

	// READ-ONLY; The image offer name for the version available for upgrade.
	Offer *string

	// READ-ONLY; The image version for the version available for upgrade.
	Version *string
}

// SupportedVersionMetadata - The SupportedVersionMetadata object for appliance.
type SupportedVersionMetadata struct {
	// READ-ONLY; The newer supported version catalog version.
	CatalogVersion *SupportedVersionCatalogVersion
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

// UpgradeGraph - The Upgrade Graph for appliance.
type UpgradeGraph struct {
	// The properties of supported version
	Properties *UpgradeGraphProperties

	// READ-ONLY; The appliance resource path
	ID *string

	// READ-ONLY; The release train name.
	Name *string
}

// UpgradeGraphProperties - The Upgrade Graph Properties for appliance.
type UpgradeGraphProperties struct {
	// READ-ONLY; The current appliance version
	ApplianceVersion *string

	// READ-ONLY; This contains the current version and supported upgrade versions.
	SupportedVersions []*SupportedVersion
}
