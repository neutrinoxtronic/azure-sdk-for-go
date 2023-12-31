//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

import "time"

// Actor - The agent that initiated the event. For most situations, this could be from the authorization context of the request.
type Actor struct {
	// The subject or username associated with the request context that generated the event.
	Name *string `json:"name,omitempty"`
}

// CallbackConfig - The configuration of service URI and custom headers for the webhook.
type CallbackConfig struct {
	// REQUIRED; The service URI for the webhook to post notifications.
	ServiceURI *string `json:"serviceUri,omitempty"`

	// Custom headers that will be added to the webhook notifications.
	CustomHeaders map[string]*string `json:"customHeaders,omitempty"`
}

// Event - The event for a webhook.
type Event struct {
	// The event request message sent to the service URI.
	EventRequestMessage *EventRequestMessage `json:"eventRequestMessage,omitempty"`

	// The event response message received from the service URI.
	EventResponseMessage *EventResponseMessage `json:"eventResponseMessage,omitempty"`

	// The event ID.
	ID *string `json:"id,omitempty"`
}

// EventContent - The content of the event request message.
type EventContent struct {
	// The action that encompasses the provided event.
	Action *string `json:"action,omitempty"`

	// The agent that initiated the event. For most situations, this could be from the authorization context of the request.
	Actor *Actor `json:"actor,omitempty"`

	// The event ID.
	ID *string `json:"id,omitempty"`

	// The request that generated the event.
	Request *Request `json:"request,omitempty"`

	// The registry node that generated the event. Put differently, while the actor initiates the event, the source generates
	// it.
	Source *Source `json:"source,omitempty"`

	// The target of the event.
	Target *Target `json:"target,omitempty"`

	// The time at which the event occurred.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// EventInfo - The basic information of an event.
type EventInfo struct {
	// The event ID.
	ID *string `json:"id,omitempty"`
}

// EventListResult - The result of a request to list events for a webhook.
type EventListResult struct {
	// The URI that can be used to request the next list of events.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of events. Since this list may be incomplete, the nextLink field should be used to request the next list of events.
	Value []*Event `json:"value,omitempty"`
}

// EventRequestMessage - The event request message sent to the service URI.
type EventRequestMessage struct {
	// The content of the event request message.
	Content *EventContent `json:"content,omitempty"`

	// The headers of the event request message.
	Headers map[string]*string `json:"headers,omitempty"`

	// The HTTP method used to send the event request message.
	Method *string `json:"method,omitempty"`

	// The URI used to send the event request message.
	RequestURI *string `json:"requestUri,omitempty"`

	// The HTTP message version.
	Version *string `json:"version,omitempty"`
}

// EventResponseMessage - The event response message received from the service URI.
type EventResponseMessage struct {
	// The content of the event response message.
	Content *string `json:"content,omitempty"`

	// The headers of the event response message.
	Headers map[string]*string `json:"headers,omitempty"`

	// The reason phrase of the event response message.
	ReasonPhrase *string `json:"reasonPhrase,omitempty"`

	// The status code of the event response message.
	StatusCode *string `json:"statusCode,omitempty"`

	// The HTTP message version.
	Version *string `json:"version,omitempty"`
}

// IPRule - IP rule with specific IP or IP range in CIDR format.
type IPRule struct {
	// REQUIRED; Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	IPAddressOrRange *string `json:"value,omitempty"`

	// The action of IP ACL rule.
	Action *Action `json:"action,omitempty"`
}

type ImportImageParameters struct {
	// REQUIRED; The source of the image.
	Source *ImportSource `json:"source,omitempty"`

	// When Force, any existing target tags will be overwritten. When NoForce, any existing target tags will fail the operation
	// before any copying begins.
	Mode *ImportMode `json:"mode,omitempty"`

	// List of strings of the form repo[:tag]. When tag is omitted the source will be used (or 'latest' if source tag is also
	// omitted).
	TargetTags []*string `json:"targetTags,omitempty"`

	// List of strings of repository names to do a manifest only copy. No tag will be created.
	UntaggedTargetRepositories []*string `json:"untaggedTargetRepositories,omitempty"`
}

type ImportSource struct {
	// REQUIRED; Repository name of the source image. Specify an image by repository ('hello-world'). This will use the 'latest'
	// tag. Specify an image by tag ('hello-world:latest'). Specify an image by sha256-based
	// manifest digest ('hello-world@sha256:abc123').
	SourceImage *string `json:"sourceImage,omitempty"`

	// Credentials used when importing from a registry uri.
	Credentials *ImportSourceCredentials `json:"credentials,omitempty"`

	// The address of the source registry (e.g. 'mcr.microsoft.com').
	RegistryURI *string `json:"registryUri,omitempty"`

	// The resource identifier of the source Azure Container Registry.
	ResourceID *string `json:"resourceId,omitempty"`
}

type ImportSourceCredentials struct {
	// REQUIRED; The password used to authenticate with the source registry.
	Password *string `json:"password,omitempty"`

	// The username to authenticate with the source registry.
	Username *string `json:"username,omitempty"`
}

// NetworkRuleSet - The network rule set for a container registry.
type NetworkRuleSet struct {
	// REQUIRED; The default action of allow or deny when no other rules match.
	DefaultAction *DefaultAction `json:"defaultAction,omitempty"`

	// The IP ACL rules.
	IPRules []*IPRule `json:"ipRules,omitempty"`

	// The virtual network rules.
	VirtualNetworkRules []*VirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
}

// OperationDefinition - The definition of a container registry operation.
type OperationDefinition struct {
	// The display information for the container registry operation.
	Display *OperationDisplayDefinition `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`

	// The origin information of the container registry operation.
	Origin *string `json:"origin,omitempty"`

	// The properties information for the container registry operation.
	Properties *OperationPropertiesDefinition `json:"properties,omitempty"`
}

// OperationDisplayDefinition - The display information for a container registry operation.
type OperationDisplayDefinition struct {
	// The description for the operation.
	Description *string `json:"description,omitempty"`

	// The operation that users can perform.
	Operation *string `json:"operation,omitempty"`

	// The resource provider name: Microsoft.ContainerRegistry.
	Provider *string `json:"provider,omitempty"`

	// The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - The result of a request to list container registry operations.
type OperationListResult struct {
	// The URI that can be used to request the next list of container registry operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of container registry operations. Since this list may be incomplete, the nextLink field should be used to request
	// the next list of operations.
	Value []*OperationDefinition `json:"value,omitempty"`
}

// OperationMetricSpecificationDefinition - The definition of Azure Monitoring metric.
type OperationMetricSpecificationDefinition struct {
	// Metric aggregation type.
	AggregationType *string `json:"aggregationType,omitempty"`

	// Metric description.
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// Metric display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Internal metric name.
	InternalMetricName *string `json:"internalMetricName,omitempty"`

	// Metric name.
	Name *string `json:"name,omitempty"`

	// Metric unit.
	Unit *string `json:"unit,omitempty"`
}

// OperationPropertiesDefinition - The definition of Azure Monitoring properties.
type OperationPropertiesDefinition struct {
	// The definition of Azure Monitoring service.
	ServiceSpecification *OperationServiceSpecificationDefinition `json:"serviceSpecification,omitempty"`
}

// OperationServiceSpecificationDefinition - The definition of Azure Monitoring list.
type OperationServiceSpecificationDefinition struct {
	// A list of Azure Monitoring metrics definition.
	MetricSpecifications []*OperationMetricSpecificationDefinition `json:"metricSpecifications,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Policies - The policies for a container registry.
type Policies struct {
	// The quarantine policy for a container registry.
	QuarantinePolicy *QuarantinePolicy `json:"quarantinePolicy,omitempty"`

	// The retention policy for a container registry.
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`

	// The content trust policy for a container registry.
	TrustPolicy *TrustPolicy `json:"trustPolicy,omitempty"`
}

// QuarantinePolicy - The quarantine policy for a container registry.
type QuarantinePolicy struct {
	// The value that indicates whether the policy is enabled or not.
	Status *PolicyStatus `json:"status,omitempty"`
}

// RegenerateCredentialParameters - The parameters used to regenerate the login credential.
type RegenerateCredentialParameters struct {
	// REQUIRED; Specifies name of the password which should be regenerated -- password or password2.
	Name *PasswordName `json:"name,omitempty"`
}

// RegistriesClientBeginCreateOptions contains the optional parameters for the RegistriesClient.BeginCreate method.
type RegistriesClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RegistriesClientBeginDeleteOptions contains the optional parameters for the RegistriesClient.BeginDelete method.
type RegistriesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RegistriesClientBeginImportImageOptions contains the optional parameters for the RegistriesClient.BeginImportImage method.
type RegistriesClientBeginImportImageOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RegistriesClientBeginUpdateOptions contains the optional parameters for the RegistriesClient.BeginUpdate method.
type RegistriesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RegistriesClientCheckNameAvailabilityOptions contains the optional parameters for the RegistriesClient.CheckNameAvailability
// method.
type RegistriesClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// RegistriesClientGetOptions contains the optional parameters for the RegistriesClient.Get method.
type RegistriesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RegistriesClientListByResourceGroupOptions contains the optional parameters for the RegistriesClient.NewListByResourceGroupPager
// method.
type RegistriesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// RegistriesClientListCredentialsOptions contains the optional parameters for the RegistriesClient.ListCredentials method.
type RegistriesClientListCredentialsOptions struct {
	// placeholder for future optional parameters
}

// RegistriesClientListOptions contains the optional parameters for the RegistriesClient.NewListPager method.
type RegistriesClientListOptions struct {
	// placeholder for future optional parameters
}

// RegistriesClientListUsagesOptions contains the optional parameters for the RegistriesClient.ListUsages method.
type RegistriesClientListUsagesOptions struct {
	// placeholder for future optional parameters
}

// RegistriesClientRegenerateCredentialOptions contains the optional parameters for the RegistriesClient.RegenerateCredential
// method.
type RegistriesClientRegenerateCredentialOptions struct {
	// placeholder for future optional parameters
}

// Registry - An object that represents a container registry.
type Registry struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// REQUIRED; The SKU of the container registry.
	SKU *SKU `json:"sku,omitempty"`

	// The properties of the container registry.
	Properties *RegistryProperties `json:"properties,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RegistryListCredentialsResult - The response from the ListCredentials operation.
type RegistryListCredentialsResult struct {
	// The list of passwords for a container registry.
	Passwords []*RegistryPassword `json:"passwords,omitempty"`

	// The username for a container registry.
	Username *string `json:"username,omitempty"`
}

// RegistryListResult - The result of a request to list container registries.
type RegistryListResult struct {
	// The URI that can be used to request the next list of container registries.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of container registries. Since this list may be incomplete, the nextLink field should be used to request the next
	// list of container registries.
	Value []*Registry `json:"value,omitempty"`
}

// RegistryNameCheckRequest - A request to check whether a container registry name is available.
type RegistryNameCheckRequest struct {
	// REQUIRED; The name of the container registry.
	Name *string `json:"name,omitempty"`

	// CONSTANT; The resource type of the container registry. This field must be set to 'Microsoft.ContainerRegistry/registries'.
	// Field has constant value "Microsoft.ContainerRegistry/registries", any specified value is ignored.
	Type *string `json:"type,omitempty"`
}

// RegistryNameStatus - The result of a request to check the availability of a container registry name.
type RegistryNameStatus struct {
	// If any, the error message that provides more detail for the reason that the name is not available.
	Message *string `json:"message,omitempty"`

	// The value that indicates whether the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// If any, the reason that the name is not available.
	Reason *string `json:"reason,omitempty"`
}

// RegistryPassword - The login password for the container registry.
type RegistryPassword struct {
	// The password name.
	Name *PasswordName `json:"name,omitempty"`

	// The password value.
	Value *string `json:"value,omitempty"`
}

// RegistryProperties - The properties of a container registry.
type RegistryProperties struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet `json:"networkRuleSet,omitempty"`

	// The policies for a container registry.
	Policies *Policies `json:"policies,omitempty"`

	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount *StorageAccountProperties `json:"storageAccount,omitempty"`

	// READ-ONLY; The creation date of the container registry in ISO8601 format.
	CreationDate *time.Time `json:"creationDate,omitempty" azure:"ro"`

	// READ-ONLY; The URL that can be used to log into the container registry.
	LoginServer *string `json:"loginServer,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state of the container registry at the time the operation was called.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The status of the container registry at the time the operation was called.
	Status *Status `json:"status,omitempty" azure:"ro"`
}

// RegistryPropertiesUpdateParameters - The parameters for updating the properties of a container registry.
type RegistryPropertiesUpdateParameters struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`

	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet `json:"networkRuleSet,omitempty"`

	// The policies for a container registry.
	Policies *Policies `json:"policies,omitempty"`
}

// RegistryUpdateParameters - The parameters for updating a container registry.
type RegistryUpdateParameters struct {
	// The properties that the container registry will be updated with.
	Properties *RegistryPropertiesUpdateParameters `json:"properties,omitempty"`

	// The SKU of the container registry.
	SKU *SKU `json:"sku,omitempty"`

	// The tags for the container registry.
	Tags map[string]*string `json:"tags,omitempty"`
}

// RegistryUsage - The quota usage for a container registry.
type RegistryUsage struct {
	// The current value of the usage.
	CurrentValue *int64 `json:"currentValue,omitempty"`

	// The limit of the usage.
	Limit *int64 `json:"limit,omitempty"`

	// The name of the usage.
	Name *string `json:"name,omitempty"`

	// The unit of measurement.
	Unit *RegistryUsageUnit `json:"unit,omitempty"`
}

// RegistryUsageListResult - The result of a request to get container registry quota usages.
type RegistryUsageListResult struct {
	// The list of container registry quota usages.
	Value []*RegistryUsage `json:"value,omitempty"`
}

// Replication - An object that represents a replication for a container registry.
type Replication struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The properties of the replication.
	Properties *ReplicationProperties `json:"properties,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ReplicationListResult - The result of a request to list replications for a container registry.
type ReplicationListResult struct {
	// The URI that can be used to request the next list of replications.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of replications. Since this list may be incomplete, the nextLink field should be used to request the next list
	// of replications.
	Value []*Replication `json:"value,omitempty"`
}

// ReplicationProperties - The properties of a replication.
type ReplicationProperties struct {
	// READ-ONLY; The provisioning state of the replication at the time the operation was called.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The status of the replication at the time the operation was called.
	Status *Status `json:"status,omitempty" azure:"ro"`
}

// ReplicationUpdateParameters - The parameters for updating a replication.
type ReplicationUpdateParameters struct {
	// The tags for the replication.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ReplicationsClientBeginCreateOptions contains the optional parameters for the ReplicationsClient.BeginCreate method.
type ReplicationsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReplicationsClientBeginDeleteOptions contains the optional parameters for the ReplicationsClient.BeginDelete method.
type ReplicationsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReplicationsClientBeginUpdateOptions contains the optional parameters for the ReplicationsClient.BeginUpdate method.
type ReplicationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReplicationsClientGetOptions contains the optional parameters for the ReplicationsClient.Get method.
type ReplicationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ReplicationsClientListOptions contains the optional parameters for the ReplicationsClient.NewListPager method.
type ReplicationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Request - The request that generated the event.
type Request struct {
	// The IP or hostname and possibly port of the client connection that initiated the event. This is the RemoteAddr from the
	// standard http request.
	Addr *string `json:"addr,omitempty"`

	// The externally accessible hostname of the registry instance, as specified by the http host header on incoming requests.
	Host *string `json:"host,omitempty"`

	// The ID of the request that initiated the event.
	ID *string `json:"id,omitempty"`

	// The request method that generated the event.
	Method *string `json:"method,omitempty"`

	// The user agent header of the request.
	Useragent *string `json:"useragent,omitempty"`
}

// Resource - An Azure resource.
type Resource struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RetentionPolicy - The retention policy for a container registry.
type RetentionPolicy struct {
	// The number of days to retain an untagged manifest after which it gets purged.
	Days *int32 `json:"days,omitempty"`

	// The value that indicates whether the policy is enabled or not.
	Status *PolicyStatus `json:"status,omitempty"`

	// READ-ONLY; The timestamp when the policy was last updated.
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty" azure:"ro"`
}

// SKU - The SKU of a container registry.
type SKU struct {
	// REQUIRED; The SKU name of the container registry. Required for registry creation.
	Name *SKUName `json:"name,omitempty"`

	// READ-ONLY; The SKU tier based on the SKU name.
	Tier *SKUTier `json:"tier,omitempty" azure:"ro"`
}

// Source - The registry node that generated the event. Put differently, while the actor initiates the event, the source generates
// it.
type Source struct {
	// The IP or hostname and the port of the registry node that generated the event. Generally, this will be resolved by os.Hostname()
	// along with the running port.
	Addr *string `json:"addr,omitempty"`

	// The running instance of an application. Changes after each restart.
	InstanceID *string `json:"instanceID,omitempty"`
}

// Status - The status of an Azure resource at the time the operation was called.
type Status struct {
	// READ-ONLY; The short label for the status.
	DisplayStatus *string `json:"displayStatus,omitempty" azure:"ro"`

	// READ-ONLY; The detailed message for the status, including alerts and error messages.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp when the status was changed to the current value.
	Timestamp *time.Time `json:"timestamp,omitempty" azure:"ro"`
}

// StorageAccountProperties - The properties of a storage account for a container registry. Only applicable to Classic SKU.
type StorageAccountProperties struct {
	// REQUIRED; The resource ID of the storage account.
	ID *string `json:"id,omitempty"`
}

// Target - The target of the event.
type Target struct {
	// The digest of the content, as defined by the Registry V2 HTTP API Specification.
	Digest *string `json:"digest,omitempty"`

	// The number of bytes of the content. Same as Size field.
	Length *int64 `json:"length,omitempty"`

	// The MIME type of the referenced object.
	MediaType *string `json:"mediaType,omitempty"`

	// The name of the artifact.
	Name *string `json:"name,omitempty"`

	// The repository name.
	Repository *string `json:"repository,omitempty"`

	// The number of bytes of the content. Same as Length field.
	Size *int64 `json:"size,omitempty"`

	// The tag name.
	Tag *string `json:"tag,omitempty"`

	// The direct URL to the content.
	URL *string `json:"url,omitempty"`

	// The version of the artifact.
	Version *string `json:"version,omitempty"`
}

// TrustPolicy - The content trust policy for a container registry.
type TrustPolicy struct {
	// The value that indicates whether the policy is enabled or not.
	Status *PolicyStatus `json:"status,omitempty"`

	// The type of trust policy.
	Type *TrustPolicyType `json:"type,omitempty"`
}

// VirtualNetworkRule - Virtual network rule.
type VirtualNetworkRule struct {
	// REQUIRED; Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	VirtualNetworkResourceID *string `json:"id,omitempty"`

	// The action of virtual network rule.
	Action *Action `json:"action,omitempty"`
}

// Webhook - An object that represents a webhook for a container registry.
type Webhook struct {
	// REQUIRED; The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The properties of the webhook.
	Properties *WebhookProperties `json:"properties,omitempty"`

	// The tags of the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// WebhookCreateParameters - The parameters for creating a webhook.
type WebhookCreateParameters struct {
	// REQUIRED; The location of the webhook. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`

	// The properties that the webhook will be created with.
	Properties *WebhookPropertiesCreateParameters `json:"properties,omitempty"`

	// The tags for the webhook.
	Tags map[string]*string `json:"tags,omitempty"`
}

// WebhookListResult - The result of a request to list webhooks for a container registry.
type WebhookListResult struct {
	// The URI that can be used to request the next list of webhooks.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of webhooks. Since this list may be incomplete, the nextLink field should be used to request the next list of
	// webhooks.
	Value []*Webhook `json:"value,omitempty"`
}

// WebhookProperties - The properties of a webhook.
type WebhookProperties struct {
	// REQUIRED; The list of actions that trigger the webhook to post notifications.
	Actions []*WebhookAction `json:"actions,omitempty"`

	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository
	// 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to
	// 'foo:latest'. Empty means all events.
	Scope *string `json:"scope,omitempty"`

	// The status of the webhook at the time the operation was called.
	Status *WebhookStatus `json:"status,omitempty"`

	// READ-ONLY; The provisioning state of the webhook at the time the operation was called.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// WebhookPropertiesCreateParameters - The parameters for creating the properties of a webhook.
type WebhookPropertiesCreateParameters struct {
	// REQUIRED; The list of actions that trigger the webhook to post notifications.
	Actions []*WebhookAction `json:"actions,omitempty"`

	// REQUIRED; The service URI for the webhook to post notifications.
	ServiceURI *string `json:"serviceUri,omitempty"`

	// Custom headers that will be added to the webhook notifications.
	CustomHeaders map[string]*string `json:"customHeaders,omitempty"`

	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository
	// 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to
	// 'foo:latest'. Empty means all events.
	Scope *string `json:"scope,omitempty"`

	// The status of the webhook at the time the operation was called.
	Status *WebhookStatus `json:"status,omitempty"`
}

// WebhookPropertiesUpdateParameters - The parameters for updating the properties of a webhook.
type WebhookPropertiesUpdateParameters struct {
	// The list of actions that trigger the webhook to post notifications.
	Actions []*WebhookAction `json:"actions,omitempty"`

	// Custom headers that will be added to the webhook notifications.
	CustomHeaders map[string]*string `json:"customHeaders,omitempty"`

	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository
	// 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to
	// 'foo:latest'. Empty means all events.
	Scope *string `json:"scope,omitempty"`

	// The service URI for the webhook to post notifications.
	ServiceURI *string `json:"serviceUri,omitempty"`

	// The status of the webhook at the time the operation was called.
	Status *WebhookStatus `json:"status,omitempty"`
}

// WebhookUpdateParameters - The parameters for updating a webhook.
type WebhookUpdateParameters struct {
	// The properties that the webhook will be updated with.
	Properties *WebhookPropertiesUpdateParameters `json:"properties,omitempty"`

	// The tags for the webhook.
	Tags map[string]*string `json:"tags,omitempty"`
}

// WebhooksClientBeginCreateOptions contains the optional parameters for the WebhooksClient.BeginCreate method.
type WebhooksClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WebhooksClientBeginDeleteOptions contains the optional parameters for the WebhooksClient.BeginDelete method.
type WebhooksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WebhooksClientBeginUpdateOptions contains the optional parameters for the WebhooksClient.BeginUpdate method.
type WebhooksClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WebhooksClientGetCallbackConfigOptions contains the optional parameters for the WebhooksClient.GetCallbackConfig method.
type WebhooksClientGetCallbackConfigOptions struct {
	// placeholder for future optional parameters
}

// WebhooksClientGetOptions contains the optional parameters for the WebhooksClient.Get method.
type WebhooksClientGetOptions struct {
	// placeholder for future optional parameters
}

// WebhooksClientListEventsOptions contains the optional parameters for the WebhooksClient.NewListEventsPager method.
type WebhooksClientListEventsOptions struct {
	// placeholder for future optional parameters
}

// WebhooksClientListOptions contains the optional parameters for the WebhooksClient.NewListPager method.
type WebhooksClientListOptions struct {
	// placeholder for future optional parameters
}

// WebhooksClientPingOptions contains the optional parameters for the WebhooksClient.Ping method.
type WebhooksClientPingOptions struct {
	// placeholder for future optional parameters
}
