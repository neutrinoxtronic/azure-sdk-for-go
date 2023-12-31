//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogic

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewWorkflowsClient() *WorkflowsClient {
	subClient, _ := NewWorkflowsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowVersionsClient() *WorkflowVersionsClient {
	subClient, _ := NewWorkflowVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowTriggersClient() *WorkflowTriggersClient {
	subClient, _ := NewWorkflowTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowVersionTriggersClient() *WorkflowVersionTriggersClient {
	subClient, _ := NewWorkflowVersionTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowTriggerHistoriesClient() *WorkflowTriggerHistoriesClient {
	subClient, _ := NewWorkflowTriggerHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunsClient() *WorkflowRunsClient {
	subClient, _ := NewWorkflowRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionsClient() *WorkflowRunActionsClient {
	subClient, _ := NewWorkflowRunActionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionRepetitionsClient() *WorkflowRunActionRepetitionsClient {
	subClient, _ := NewWorkflowRunActionRepetitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionRepetitionsRequestHistoriesClient() *WorkflowRunActionRepetitionsRequestHistoriesClient {
	subClient, _ := NewWorkflowRunActionRepetitionsRequestHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionRequestHistoriesClient() *WorkflowRunActionRequestHistoriesClient {
	subClient, _ := NewWorkflowRunActionRequestHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionScopeRepetitionsClient() *WorkflowRunActionScopeRepetitionsClient {
	subClient, _ := NewWorkflowRunActionScopeRepetitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunOperationsClient() *WorkflowRunOperationsClient {
	subClient, _ := NewWorkflowRunOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountsClient() *IntegrationAccountsClient {
	subClient, _ := NewIntegrationAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountAssembliesClient() *IntegrationAccountAssembliesClient {
	subClient, _ := NewIntegrationAccountAssembliesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountBatchConfigurationsClient() *IntegrationAccountBatchConfigurationsClient {
	subClient, _ := NewIntegrationAccountBatchConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountSchemasClient() *IntegrationAccountSchemasClient {
	subClient, _ := NewIntegrationAccountSchemasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountMapsClient() *IntegrationAccountMapsClient {
	subClient, _ := NewIntegrationAccountMapsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountPartnersClient() *IntegrationAccountPartnersClient {
	subClient, _ := NewIntegrationAccountPartnersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountAgreementsClient() *IntegrationAccountAgreementsClient {
	subClient, _ := NewIntegrationAccountAgreementsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountCertificatesClient() *IntegrationAccountCertificatesClient {
	subClient, _ := NewIntegrationAccountCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationAccountSessionsClient() *IntegrationAccountSessionsClient {
	subClient, _ := NewIntegrationAccountSessionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationServiceEnvironmentsClient() *IntegrationServiceEnvironmentsClient {
	subClient, _ := NewIntegrationServiceEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationServiceEnvironmentSKUsClient() *IntegrationServiceEnvironmentSKUsClient {
	subClient, _ := NewIntegrationServiceEnvironmentSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationServiceEnvironmentNetworkHealthClient() *IntegrationServiceEnvironmentNetworkHealthClient {
	subClient, _ := NewIntegrationServiceEnvironmentNetworkHealthClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationServiceEnvironmentManagedApisClient() *IntegrationServiceEnvironmentManagedApisClient {
	subClient, _ := NewIntegrationServiceEnvironmentManagedApisClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationServiceEnvironmentManagedAPIOperationsClient() *IntegrationServiceEnvironmentManagedAPIOperationsClient {
	subClient, _ := NewIntegrationServiceEnvironmentManagedAPIOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}
