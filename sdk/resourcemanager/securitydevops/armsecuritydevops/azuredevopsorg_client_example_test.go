//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgList.json
func ExampleAzureDevOpsOrgClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureDevOpsOrgClient().NewListPager("westusrg", "testconnector", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AzureDevOpsOrgListResponse = armsecuritydevops.AzureDevOpsOrgListResponse{
		// 	Value: []*armsecuritydevops.AzureDevOpsOrg{
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors/orgs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg1"),
		// 			Properties: &armsecuritydevops.AzureDevOpsOrgProperties{
		// 				AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors/orgs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg2"),
		// 			Properties: &armsecuritydevops.AzureDevOpsOrgProperties{
		// 				AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgGet.json
func ExampleAzureDevOpsOrgClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureDevOpsOrgClient().Get(ctx, "westusrg", "testconnector", "myOrg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsOrg = armsecuritydevops.AzureDevOpsOrg{
	// 	Type: to.Ptr("Microsoft.SecurityDevOps/azureDevOpsConnectos/orgs"),
	// 	ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg"),
	// 	Properties: &armsecuritydevops.AzureDevOpsOrgProperties{
	// 		AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgCreateOrUpdate.json
func ExampleAzureDevOpsOrgClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsOrgClient().BeginCreateOrUpdate(ctx, "westusrg", "testconnector", "myOrg", armsecuritydevops.AzureDevOpsOrg{
		Properties: &armsecuritydevops.AzureDevOpsOrgProperties{
			AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsOrg = armsecuritydevops.AzureDevOpsOrg{
	// 	Name: to.Ptr("myOrg"),
	// 	Type: to.Ptr("Microsoft.SecurityDevOps/azureDevOpsConnectos/orgs"),
	// 	ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg"),
	// 	Properties: &armsecuritydevops.AzureDevOpsOrgProperties{
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgUpdate.json
func ExampleAzureDevOpsOrgClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsOrgClient().BeginUpdate(ctx, "westusrg", "testconnector", "myOrg", armsecuritydevops.AzureDevOpsOrg{
		Properties: &armsecuritydevops.AzureDevOpsOrgProperties{
			AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
