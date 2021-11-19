//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

const (
	module  = "armbatch"
	version = "v0.1.0"
)

// AccountKeyType - The type of account key to regenerate.
type AccountKeyType string

const (
	// AccountKeyTypePrimary - The primary account key.
	AccountKeyTypePrimary AccountKeyType = "Primary"
	// AccountKeyTypeSecondary - The secondary account key.
	AccountKeyTypeSecondary AccountKeyType = "Secondary"
)

// PossibleAccountKeyTypeValues returns the possible values for the AccountKeyType const type.
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return []AccountKeyType{
		AccountKeyTypePrimary,
		AccountKeyTypeSecondary,
	}
}

// ToPtr returns a *AccountKeyType pointing to the current value.
func (c AccountKeyType) ToPtr() *AccountKeyType {
	return &c
}

// AllocationState - Whether the pool is resizing.
type AllocationState string

const (
	// AllocationStateSteady - The pool is not resizing. There are no changes to the number of nodes in the pool in progress. A pool enters this state when
	// it is created and when no operations are being performed on the pool to change the number of nodes.
	AllocationStateSteady AllocationState = "Steady"
	// AllocationStateResizing - The pool is resizing; that is, compute nodes are being added to or removed from the pool.
	AllocationStateResizing AllocationState = "Resizing"
	// AllocationStateStopping - The pool was resizing, but the user has requested that the resize be stopped, but the stop request has not yet been completed.
	AllocationStateStopping AllocationState = "Stopping"
)

// PossibleAllocationStateValues returns the possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{
		AllocationStateSteady,
		AllocationStateResizing,
		AllocationStateStopping,
	}
}

// ToPtr returns a *AllocationState pointing to the current value.
func (c AllocationState) ToPtr() *AllocationState {
	return &c
}

// AuthenticationMode - The authentication mode for the Batch account.
type AuthenticationMode string

const (
	// AuthenticationModeSharedKey - The authentication mode using shared keys.
	AuthenticationModeSharedKey AuthenticationMode = "SharedKey"
	// AuthenticationModeAAD - The authentication mode using Azure Active Directory.
	AuthenticationModeAAD AuthenticationMode = "AAD"
	// AuthenticationModeTaskAuthenticationToken - The authentication mode using task authentication tokens.
	AuthenticationModeTaskAuthenticationToken AuthenticationMode = "TaskAuthenticationToken"
)

// PossibleAuthenticationModeValues returns the possible values for the AuthenticationMode const type.
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return []AuthenticationMode{
		AuthenticationModeSharedKey,
		AuthenticationModeAAD,
		AuthenticationModeTaskAuthenticationToken,
	}
}

// ToPtr returns a *AuthenticationMode pointing to the current value.
func (c AuthenticationMode) ToPtr() *AuthenticationMode {
	return &c
}

// AutoStorageAuthenticationMode - The authentication mode which the Batch service will use to manage the auto-storage account.
type AutoStorageAuthenticationMode string

const (
	// AutoStorageAuthenticationModeStorageKeys - The Batch service will authenticate requests to auto-storage using storage account keys.
	AutoStorageAuthenticationModeStorageKeys AutoStorageAuthenticationMode = "StorageKeys"
	// AutoStorageAuthenticationModeBatchAccountManagedIdentity - The Batch service will authenticate requests to auto-storage using the managed identity assigned
	// to the Batch account.
	AutoStorageAuthenticationModeBatchAccountManagedIdentity AutoStorageAuthenticationMode = "BatchAccountManagedIdentity"
)

// PossibleAutoStorageAuthenticationModeValues returns the possible values for the AutoStorageAuthenticationMode const type.
func PossibleAutoStorageAuthenticationModeValues() []AutoStorageAuthenticationMode {
	return []AutoStorageAuthenticationMode{
		AutoStorageAuthenticationModeStorageKeys,
		AutoStorageAuthenticationModeBatchAccountManagedIdentity,
	}
}

// ToPtr returns a *AutoStorageAuthenticationMode pointing to the current value.
func (c AutoStorageAuthenticationMode) ToPtr() *AutoStorageAuthenticationMode {
	return &c
}

// AutoUserScope - The default value is Pool. If the pool is running Windows a value of Task should be specified if stricter isolation between tasks is
// required. For example, if the task mutates the registry in a way
// which could impact other tasks, or if certificates have been specified on the pool which should not be accessible by normal tasks but should be accessible
// by start tasks.
type AutoUserScope string

const (
	// AutoUserScopeTask - Specifies that the service should create a new user for the task.
	AutoUserScopeTask AutoUserScope = "Task"
	// AutoUserScopePool - Specifies that the task runs as the common auto user account which is created on every node in a pool.
	AutoUserScopePool AutoUserScope = "Pool"
)

// PossibleAutoUserScopeValues returns the possible values for the AutoUserScope const type.
func PossibleAutoUserScopeValues() []AutoUserScope {
	return []AutoUserScope{
		AutoUserScopeTask,
		AutoUserScopePool,
	}
}

// ToPtr returns a *AutoUserScope pointing to the current value.
func (c AutoUserScope) ToPtr() *AutoUserScope {
	return &c
}

// CachingType - The type of caching to enable for the disk.
type CachingType string

const (
	// CachingTypeNone - The caching mode for the disk is not enabled.
	CachingTypeNone CachingType = "None"
	// CachingTypeReadOnly - The caching mode for the disk is read only.
	CachingTypeReadOnly CachingType = "ReadOnly"
	// CachingTypeReadWrite - The caching mode for the disk is read and write.
	CachingTypeReadWrite CachingType = "ReadWrite"
)

// PossibleCachingTypeValues returns the possible values for the CachingType const type.
func PossibleCachingTypeValues() []CachingType {
	return []CachingType{
		CachingTypeNone,
		CachingTypeReadOnly,
		CachingTypeReadWrite,
	}
}

// ToPtr returns a *CachingType pointing to the current value.
func (c CachingType) ToPtr() *CachingType {
	return &c
}

// CertificateFormat - The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
type CertificateFormat string

const (
	// CertificateFormatPfx - The certificate is a PFX (PKCS#12) formatted certificate or certificate chain.
	CertificateFormatPfx CertificateFormat = "Pfx"
	// CertificateFormatCer - The certificate is a base64-encoded X.509 certificate.
	CertificateFormatCer CertificateFormat = "Cer"
)

// PossibleCertificateFormatValues returns the possible values for the CertificateFormat const type.
func PossibleCertificateFormatValues() []CertificateFormat {
	return []CertificateFormat{
		CertificateFormatPfx,
		CertificateFormatCer,
	}
}

// ToPtr returns a *CertificateFormat pointing to the current value.
func (c CertificateFormat) ToPtr() *CertificateFormat {
	return &c
}

type CertificateProvisioningState string

const (
	// CertificateProvisioningStateSucceeded - The certificate is available for use in pools.
	CertificateProvisioningStateSucceeded CertificateProvisioningState = "Succeeded"
	// CertificateProvisioningStateDeleting - The user has requested that the certificate be deleted, but the delete operation has not yet completed. You may
	// not reference the certificate when creating or updating pools.
	CertificateProvisioningStateDeleting CertificateProvisioningState = "Deleting"
	// CertificateProvisioningStateFailed - The user requested that the certificate be deleted, but there are pools that still have references to the certificate,
	// or it is still installed on one or more compute nodes. (The latter can occur if the certificate has been removed from the pool, but the node has not
	// yet restarted. Nodes refresh their certificates only when they restart.) You may use the cancel certificate delete operation to cancel the delete, or
	// the delete certificate operation to retry the delete.
	CertificateProvisioningStateFailed CertificateProvisioningState = "Failed"
)

// PossibleCertificateProvisioningStateValues returns the possible values for the CertificateProvisioningState const type.
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return []CertificateProvisioningState{
		CertificateProvisioningStateSucceeded,
		CertificateProvisioningStateDeleting,
		CertificateProvisioningStateFailed,
	}
}

// ToPtr returns a *CertificateProvisioningState pointing to the current value.
func (c CertificateProvisioningState) ToPtr() *CertificateProvisioningState {
	return &c
}

// CertificateStoreLocation - The default value is currentUser. This property is applicable only for pools configured with Windows nodes (that is, created
// with cloudServiceConfiguration, or with virtualMachineConfiguration using a
// Windows image reference). For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable
// AZBATCHCERTIFICATES_DIR is supplied to the
// task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g.,
// /home/{user-name}/certs) and certificates are
// placed in that directory.
type CertificateStoreLocation string

const (
	// CertificateStoreLocationCurrentUser - Certificates should be installed to the CurrentUser certificate store.
	CertificateStoreLocationCurrentUser CertificateStoreLocation = "CurrentUser"
	// CertificateStoreLocationLocalMachine - Certificates should be installed to the LocalMachine certificate store.
	CertificateStoreLocationLocalMachine CertificateStoreLocation = "LocalMachine"
)

// PossibleCertificateStoreLocationValues returns the possible values for the CertificateStoreLocation const type.
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return []CertificateStoreLocation{
		CertificateStoreLocationCurrentUser,
		CertificateStoreLocationLocalMachine,
	}
}

// ToPtr returns a *CertificateStoreLocation pointing to the current value.
func (c CertificateStoreLocation) ToPtr() *CertificateStoreLocation {
	return &c
}

type CertificateVisibility string

const (
	// CertificateVisibilityStartTask - The certificate should be visible to the user account under which the start task is run. Note that if AutoUser Scope
	// is Pool for both the StartTask and a Task, this certificate will be visible to the Task as well.
	CertificateVisibilityStartTask CertificateVisibility = "StartTask"
	// CertificateVisibilityTask - The certificate should be visible to the user accounts under which job tasks are run.
	CertificateVisibilityTask CertificateVisibility = "Task"
	// CertificateVisibilityRemoteUser - The certificate should be visible to the user accounts under which users remotely access the node.
	CertificateVisibilityRemoteUser CertificateVisibility = "RemoteUser"
)

// PossibleCertificateVisibilityValues returns the possible values for the CertificateVisibility const type.
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return []CertificateVisibility{
		CertificateVisibilityStartTask,
		CertificateVisibilityTask,
		CertificateVisibilityRemoteUser,
	}
}

// ToPtr returns a *CertificateVisibility pointing to the current value.
func (c CertificateVisibility) ToPtr() *CertificateVisibility {
	return &c
}

// ComputeNodeDeallocationOption - Determines what to do with a node and its running task(s) after it has been selected for deallocation.
type ComputeNodeDeallocationOption string

const (
	// ComputeNodeDeallocationOptionRequeue - Terminate running task processes and requeue the tasks. The tasks will run again when a node is available. Remove
	// nodes as soon as tasks have been terminated.
	ComputeNodeDeallocationOptionRequeue ComputeNodeDeallocationOption = "Requeue"
	// ComputeNodeDeallocationOptionTerminate - Terminate running tasks. The tasks will be completed with failureInfo indicating that they were terminated,
	// and will not run again. Remove nodes as soon as tasks have been terminated.
	ComputeNodeDeallocationOptionTerminate ComputeNodeDeallocationOption = "Terminate"
	// ComputeNodeDeallocationOptionTaskCompletion - Allow currently running tasks to complete. Schedule no new tasks while waiting. Remove nodes when all tasks
	// have completed.
	ComputeNodeDeallocationOptionTaskCompletion ComputeNodeDeallocationOption = "TaskCompletion"
	// ComputeNodeDeallocationOptionRetainedData - Allow currently running tasks to complete, then wait for all task data retention periods to expire. Schedule
	// no new tasks while waiting. Remove nodes when all task retention periods have expired.
	ComputeNodeDeallocationOptionRetainedData ComputeNodeDeallocationOption = "RetainedData"
)

// PossibleComputeNodeDeallocationOptionValues returns the possible values for the ComputeNodeDeallocationOption const type.
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return []ComputeNodeDeallocationOption{
		ComputeNodeDeallocationOptionRequeue,
		ComputeNodeDeallocationOptionTerminate,
		ComputeNodeDeallocationOptionTaskCompletion,
		ComputeNodeDeallocationOptionRetainedData,
	}
}

// ToPtr returns a *ComputeNodeDeallocationOption pointing to the current value.
func (c ComputeNodeDeallocationOption) ToPtr() *ComputeNodeDeallocationOption {
	return &c
}

// ComputeNodeFillType - How tasks should be distributed across compute nodes.
type ComputeNodeFillType string

const (
	// ComputeNodeFillTypeSpread - Tasks should be assigned evenly across all nodes in the pool.
	ComputeNodeFillTypeSpread ComputeNodeFillType = "Spread"
	// ComputeNodeFillTypePack - As many tasks as possible (taskSlotsPerNode) should be assigned to each node in the pool before any tasks are assigned to the
	// next node in the pool.
	ComputeNodeFillTypePack ComputeNodeFillType = "Pack"
)

// PossibleComputeNodeFillTypeValues returns the possible values for the ComputeNodeFillType const type.
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return []ComputeNodeFillType{
		ComputeNodeFillTypeSpread,
		ComputeNodeFillTypePack,
	}
}

// ToPtr returns a *ComputeNodeFillType pointing to the current value.
func (c ComputeNodeFillType) ToPtr() *ComputeNodeFillType {
	return &c
}

// ContainerWorkingDirectory - A flag to indicate where the container task working directory is. The default is 'taskWorkingDirectory'.
type ContainerWorkingDirectory string

const (
	// ContainerWorkingDirectoryTaskWorkingDirectory - Use the standard Batch service task working directory, which will contain the Task resource files populated
	// by Batch.
	ContainerWorkingDirectoryTaskWorkingDirectory ContainerWorkingDirectory = "TaskWorkingDirectory"
	// ContainerWorkingDirectoryContainerImageDefault - Using container image defined working directory. Beware that this directory will not contain the resource
	// files downloaded by Batch.
	ContainerWorkingDirectoryContainerImageDefault ContainerWorkingDirectory = "ContainerImageDefault"
)

// PossibleContainerWorkingDirectoryValues returns the possible values for the ContainerWorkingDirectory const type.
func PossibleContainerWorkingDirectoryValues() []ContainerWorkingDirectory {
	return []ContainerWorkingDirectory{
		ContainerWorkingDirectoryTaskWorkingDirectory,
		ContainerWorkingDirectoryContainerImageDefault,
	}
}

// ToPtr returns a *ContainerWorkingDirectory pointing to the current value.
func (c ContainerWorkingDirectory) ToPtr() *ContainerWorkingDirectory {
	return &c
}

// DiskEncryptionTarget - If omitted, no disks on the compute nodes in the pool will be encrypted.
type DiskEncryptionTarget string

const (
	// DiskEncryptionTargetOsDisk - The OS Disk on the compute node is encrypted.
	DiskEncryptionTargetOsDisk DiskEncryptionTarget = "OsDisk"
	// DiskEncryptionTargetTemporaryDisk - The temporary disk on the compute node is encrypted. On Linux this encryption applies to other partitions (such as
	// those on mounted data disks) when encryption occurs at boot time.
	DiskEncryptionTargetTemporaryDisk DiskEncryptionTarget = "TemporaryDisk"
)

// PossibleDiskEncryptionTargetValues returns the possible values for the DiskEncryptionTarget const type.
func PossibleDiskEncryptionTargetValues() []DiskEncryptionTarget {
	return []DiskEncryptionTarget{
		DiskEncryptionTargetOsDisk,
		DiskEncryptionTargetTemporaryDisk,
	}
}

// ToPtr returns a *DiskEncryptionTarget pointing to the current value.
func (c DiskEncryptionTarget) ToPtr() *DiskEncryptionTarget {
	return &c
}

// ElevationLevel - The elevation level of the user.
type ElevationLevel string

const (
	// ElevationLevelNonAdmin - The user is a standard user without elevated access.
	ElevationLevelNonAdmin ElevationLevel = "NonAdmin"
	// ElevationLevelAdmin - The user is a user with elevated access and operates with full Administrator permissions.
	ElevationLevelAdmin ElevationLevel = "Admin"
)

// PossibleElevationLevelValues returns the possible values for the ElevationLevel const type.
func PossibleElevationLevelValues() []ElevationLevel {
	return []ElevationLevel{
		ElevationLevelNonAdmin,
		ElevationLevelAdmin,
	}
}

// ToPtr returns a *ElevationLevel pointing to the current value.
func (c ElevationLevel) ToPtr() *ElevationLevel {
	return &c
}

// IPAddressProvisioningType - The provisioning type for Public IP Addresses for the Batch Pool.
type IPAddressProvisioningType string

const (
	// IPAddressProvisioningTypeBatchManaged - A public IP will be created and managed by Batch. There may be multiple public IPs depending on the size of the
	// Pool.
	IPAddressProvisioningTypeBatchManaged IPAddressProvisioningType = "BatchManaged"
	// IPAddressProvisioningTypeUserManaged - Public IPs are provided by the user and will be used to provision the Compute Nodes.
	IPAddressProvisioningTypeUserManaged IPAddressProvisioningType = "UserManaged"
	// IPAddressProvisioningTypeNoPublicIPAddresses - No public IP Address will be created for the Compute Nodes in the Pool.
	IPAddressProvisioningTypeNoPublicIPAddresses IPAddressProvisioningType = "NoPublicIPAddresses"
)

// PossibleIPAddressProvisioningTypeValues returns the possible values for the IPAddressProvisioningType const type.
func PossibleIPAddressProvisioningTypeValues() []IPAddressProvisioningType {
	return []IPAddressProvisioningType{
		IPAddressProvisioningTypeBatchManaged,
		IPAddressProvisioningTypeUserManaged,
		IPAddressProvisioningTypeNoPublicIPAddresses,
	}
}

// ToPtr returns a *IPAddressProvisioningType pointing to the current value.
func (c IPAddressProvisioningType) ToPtr() *IPAddressProvisioningType {
	return &c
}

// InboundEndpointProtocol - The protocol of the endpoint.
type InboundEndpointProtocol string

const (
	// InboundEndpointProtocolTCP - Use TCP for the endpoint.
	InboundEndpointProtocolTCP InboundEndpointProtocol = "TCP"
	// InboundEndpointProtocolUDP - Use UDP for the endpoint.
	InboundEndpointProtocolUDP InboundEndpointProtocol = "UDP"
)

// PossibleInboundEndpointProtocolValues returns the possible values for the InboundEndpointProtocol const type.
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return []InboundEndpointProtocol{
		InboundEndpointProtocolTCP,
		InboundEndpointProtocolUDP,
	}
}

// ToPtr returns a *InboundEndpointProtocol pointing to the current value.
func (c InboundEndpointProtocol) ToPtr() *InboundEndpointProtocol {
	return &c
}

// InterNodeCommunicationState - This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the
// requested number of nodes to be allocated in the pool. If not specified, this
// value defaults to 'Disabled'.
type InterNodeCommunicationState string

const (
	// InterNodeCommunicationStateEnabled - Enable network communication between virtual machines.
	InterNodeCommunicationStateEnabled InterNodeCommunicationState = "Enabled"
	// InterNodeCommunicationStateDisabled - Disable network communication between virtual machines.
	InterNodeCommunicationStateDisabled InterNodeCommunicationState = "Disabled"
)

// PossibleInterNodeCommunicationStateValues returns the possible values for the InterNodeCommunicationState const type.
func PossibleInterNodeCommunicationStateValues() []InterNodeCommunicationState {
	return []InterNodeCommunicationState{
		InterNodeCommunicationStateEnabled,
		InterNodeCommunicationStateDisabled,
	}
}

// ToPtr returns a *InterNodeCommunicationState pointing to the current value.
func (c InterNodeCommunicationState) ToPtr() *InterNodeCommunicationState {
	return &c
}

// KeySource - Type of the key source.
type KeySource string

const (
	// KeySourceMicrosoftBatch - Batch creates and manages the encryption keys used to protect the account data.
	KeySourceMicrosoftBatch KeySource = "Microsoft.Batch"
	// KeySourceMicrosoftKeyVault - The encryption keys used to protect the account data are stored in an external key vault. If this is set then the Batch
	// Account identity must be set to `SystemAssigned` and a valid Key Identifier must also be supplied under the keyVaultProperties.
	KeySourceMicrosoftKeyVault KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftBatch,
		KeySourceMicrosoftKeyVault,
	}
}

// ToPtr returns a *KeySource pointing to the current value.
func (c KeySource) ToPtr() *KeySource {
	return &c
}

// LoginMode - Specifies login mode for the user. The default value for VirtualMachineConfiguration pools is interactive mode and for CloudServiceConfiguration
// pools is batch mode.
type LoginMode string

const (
	// LoginModeBatch - The LOGON32_LOGON_BATCH Win32 login mode. The batch login mode is recommended for long running parallel processes.
	LoginModeBatch LoginMode = "Batch"
	// LoginModeInteractive - The LOGON32_LOGON_INTERACTIVE Win32 login mode. Some applications require having permissions associated with the interactive login
	// mode. If this is the case for an application used in your task, then this option is recommended.
	LoginModeInteractive LoginMode = "Interactive"
)

// PossibleLoginModeValues returns the possible values for the LoginMode const type.
func PossibleLoginModeValues() []LoginMode {
	return []LoginMode{
		LoginModeBatch,
		LoginModeInteractive,
	}
}

// ToPtr returns a *LoginMode pointing to the current value.
func (c LoginMode) ToPtr() *LoginMode {
	return &c
}

// NameAvailabilityReason - Gets the reason that a Batch account name could not be used. The Reason element is only returned if NameAvailable is false.
type NameAvailabilityReason string

const (
	// NameAvailabilityReasonInvalid - The requested name is invalid.
	NameAvailabilityReasonInvalid NameAvailabilityReason = "Invalid"
	// NameAvailabilityReasonAlreadyExists - The requested name is already in use.
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = "AlreadyExists"
)

// PossibleNameAvailabilityReasonValues returns the possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{
		NameAvailabilityReasonInvalid,
		NameAvailabilityReasonAlreadyExists,
	}
}

// ToPtr returns a *NameAvailabilityReason pointing to the current value.
func (c NameAvailabilityReason) ToPtr() *NameAvailabilityReason {
	return &c
}

// NetworkSecurityGroupRuleAccess - The action that should be taken for a specified IP address, subnet range or tag.
type NetworkSecurityGroupRuleAccess string

const (
	// NetworkSecurityGroupRuleAccessAllow - Allow access.
	NetworkSecurityGroupRuleAccessAllow NetworkSecurityGroupRuleAccess = "Allow"
	// NetworkSecurityGroupRuleAccessDeny - Deny access.
	NetworkSecurityGroupRuleAccessDeny NetworkSecurityGroupRuleAccess = "Deny"
)

// PossibleNetworkSecurityGroupRuleAccessValues returns the possible values for the NetworkSecurityGroupRuleAccess const type.
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return []NetworkSecurityGroupRuleAccess{
		NetworkSecurityGroupRuleAccessAllow,
		NetworkSecurityGroupRuleAccessDeny,
	}
}

// ToPtr returns a *NetworkSecurityGroupRuleAccess pointing to the current value.
func (c NetworkSecurityGroupRuleAccess) ToPtr() *NetworkSecurityGroupRuleAccess {
	return &c
}

// NodePlacementPolicyType - The default value is regional.
type NodePlacementPolicyType string

const (
	// NodePlacementPolicyTypeRegional - All nodes in the pool will be allocated in the same region.
	NodePlacementPolicyTypeRegional NodePlacementPolicyType = "Regional"
	// NodePlacementPolicyTypeZonal - Nodes in the pool will be spread across different zones with best effort balancing.
	NodePlacementPolicyTypeZonal NodePlacementPolicyType = "Zonal"
)

// PossibleNodePlacementPolicyTypeValues returns the possible values for the NodePlacementPolicyType const type.
func PossibleNodePlacementPolicyTypeValues() []NodePlacementPolicyType {
	return []NodePlacementPolicyType{
		NodePlacementPolicyTypeRegional,
		NodePlacementPolicyTypeZonal,
	}
}

// ToPtr returns a *NodePlacementPolicyType pointing to the current value.
func (c NodePlacementPolicyType) ToPtr() *NodePlacementPolicyType {
	return &c
}

// PackageState - The current state of the application package.
type PackageState string

const (
	// PackageStatePending - The application package has been created but has not yet been activated.
	PackageStatePending PackageState = "Pending"
	// PackageStateActive - The application package is ready for use.
	PackageStateActive PackageState = "Active"
)

// PossiblePackageStateValues returns the possible values for the PackageState const type.
func PossiblePackageStateValues() []PackageState {
	return []PackageState{
		PackageStatePending,
		PackageStateActive,
	}
}

// ToPtr returns a *PackageState pointing to the current value.
func (c PackageState) ToPtr() *PackageState {
	return &c
}

// PoolAllocationMode - The allocation mode for creating pools in the Batch account.
type PoolAllocationMode string

const (
	// PoolAllocationModeBatchService - Pools will be allocated in subscriptions owned by the Batch service.
	PoolAllocationModeBatchService PoolAllocationMode = "BatchService"
	// PoolAllocationModeUserSubscription - Pools will be allocated in a subscription owned by the user.
	PoolAllocationModeUserSubscription PoolAllocationMode = "UserSubscription"
)

// PossiblePoolAllocationModeValues returns the possible values for the PoolAllocationMode const type.
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return []PoolAllocationMode{
		PoolAllocationModeBatchService,
		PoolAllocationModeUserSubscription,
	}
}

// ToPtr returns a *PoolAllocationMode pointing to the current value.
func (c PoolAllocationMode) ToPtr() *PoolAllocationMode {
	return &c
}

// PoolIdentityType - The type of identity used for the Batch Pool.
type PoolIdentityType string

const (
	// PoolIdentityTypeUserAssigned - Batch pool has user assigned identities with it.
	PoolIdentityTypeUserAssigned PoolIdentityType = "UserAssigned"
	// PoolIdentityTypeNone - Batch pool has no identity associated with it. Setting `None` in update pool will remove existing identities.
	PoolIdentityTypeNone PoolIdentityType = "None"
)

// PossiblePoolIdentityTypeValues returns the possible values for the PoolIdentityType const type.
func PossiblePoolIdentityTypeValues() []PoolIdentityType {
	return []PoolIdentityType{
		PoolIdentityTypeUserAssigned,
		PoolIdentityTypeNone,
	}
}

// ToPtr returns a *PoolIdentityType pointing to the current value.
func (c PoolIdentityType) ToPtr() *PoolIdentityType {
	return &c
}

// PoolProvisioningState - The current state of the pool.
type PoolProvisioningState string

const (
	// PoolProvisioningStateSucceeded - The pool is available to run tasks subject to the availability of compute nodes.
	PoolProvisioningStateSucceeded PoolProvisioningState = "Succeeded"
	// PoolProvisioningStateDeleting - The user has requested that the pool be deleted, but the delete operation has not yet completed.
	PoolProvisioningStateDeleting PoolProvisioningState = "Deleting"
)

// PossiblePoolProvisioningStateValues returns the possible values for the PoolProvisioningState const type.
func PossiblePoolProvisioningStateValues() []PoolProvisioningState {
	return []PoolProvisioningState{
		PoolProvisioningStateSucceeded,
		PoolProvisioningStateDeleting,
	}
}

// ToPtr returns a *PoolProvisioningState pointing to the current value.
func (c PoolProvisioningState) ToPtr() *PoolProvisioningState {
	return &c
}

// PrivateEndpointConnectionProvisioningState - The provisioning state of the private endpoint connection.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateSucceeded - The connection status is final and is ready for use if Status is Approved.
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
	// PrivateEndpointConnectionProvisioningStateUpdating - The user has requested that the connection status be updated, but the update operation has not yet
	// completed. You may not reference the connection when connecting the Batch account.
	PrivateEndpointConnectionProvisioningStateUpdating PrivateEndpointConnectionProvisioningState = "Updating"
	// PrivateEndpointConnectionProvisioningStateFailed - The user requested that the connection be updated and it failed. You may retry the update operation.
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
		PrivateEndpointConnectionProvisioningStateFailed,
	}
}

// ToPtr returns a *PrivateEndpointConnectionProvisioningState pointing to the current value.
func (c PrivateEndpointConnectionProvisioningState) ToPtr() *PrivateEndpointConnectionProvisioningState {
	return &c
}

// PrivateLinkServiceConnectionStatus - The status of the Batch private endpoint connection
type PrivateLinkServiceConnectionStatus string

const (
	// PrivateLinkServiceConnectionStatusApproved - The private endpoint connection is approved and can be used to access Batch account
	PrivateLinkServiceConnectionStatusApproved PrivateLinkServiceConnectionStatus = "Approved"
	// PrivateLinkServiceConnectionStatusPending - The private endpoint connection is pending and cannot be used to access Batch account
	PrivateLinkServiceConnectionStatusPending PrivateLinkServiceConnectionStatus = "Pending"
	// PrivateLinkServiceConnectionStatusRejected - The private endpoint connection is rejected and cannot be used to access Batch account
	PrivateLinkServiceConnectionStatusRejected PrivateLinkServiceConnectionStatus = "Rejected"
	// PrivateLinkServiceConnectionStatusDisconnected - The private endpoint connection is disconnected and cannot be used to access Batch account
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
		PrivateLinkServiceConnectionStatusDisconnected,
	}
}

// ToPtr returns a *PrivateLinkServiceConnectionStatus pointing to the current value.
func (c PrivateLinkServiceConnectionStatus) ToPtr() *PrivateLinkServiceConnectionStatus {
	return &c
}

// ProvisioningState - The provisioned state of the resource
type ProvisioningState string

const (
	// ProvisioningStateInvalid - The account is in an invalid state.
	ProvisioningStateInvalid ProvisioningState = "Invalid"
	// ProvisioningStateCreating - The account is being created.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting - The account is being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateSucceeded - The account has been created and is ready for use.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateFailed - The last operation for the account is failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateCancelled - The last operation for the account is cancelled.
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateInvalid,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateSucceeded,
		ProvisioningStateFailed,
		ProvisioningStateCancelled,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// PublicNetworkAccessType - The network access type for operating on the resources in the Batch account.
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeEnabled - Enables connectivity to Azure Batch through public DNS.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
	// PublicNetworkAccessTypeDisabled - Disables public connectivity and enables private connectivity to Azure Batch Service through private endpoint resource.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
)

// PossiblePublicNetworkAccessTypeValues returns the possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{
		PublicNetworkAccessTypeEnabled,
		PublicNetworkAccessTypeDisabled,
	}
}

// ToPtr returns a *PublicNetworkAccessType pointing to the current value.
func (c PublicNetworkAccessType) ToPtr() *PublicNetworkAccessType {
	return &c
}

// ResourceIdentityType - The type of identity used for the Batch account.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeSystemAssigned - Batch account has a system assigned identity with it.
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeUserAssigned - Batch account has user assigned identities with it.
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
	// ResourceIdentityTypeNone - Batch account has no identity associated with it. Setting `None` in update account will remove existing identities.
	ResourceIdentityTypeNone ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// ToPtr returns a *ResourceIdentityType pointing to the current value.
func (c ResourceIdentityType) ToPtr() *ResourceIdentityType {
	return &c
}

// StorageAccountType - The storage account type for use in creating data disks.
type StorageAccountType string

const (
	// StorageAccountTypeStandardLRS - The data disk should use standard locally redundant storage.
	StorageAccountTypeStandardLRS StorageAccountType = "Standard_LRS"
	// StorageAccountTypePremiumLRS - The data disk should use premium locally redundant storage.
	StorageAccountTypePremiumLRS StorageAccountType = "Premium_LRS"
)

// PossibleStorageAccountTypeValues returns the possible values for the StorageAccountType const type.
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return []StorageAccountType{
		StorageAccountTypeStandardLRS,
		StorageAccountTypePremiumLRS,
	}
}

// ToPtr returns a *StorageAccountType pointing to the current value.
func (c StorageAccountType) ToPtr() *StorageAccountType {
	return &c
}