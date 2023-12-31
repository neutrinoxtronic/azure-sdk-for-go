//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Create_SystemAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate_assignmentWithSystemAssignedManagedIdentityAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().CreateOrUpdate(ctx, "managementGroups/ContosoOnlineGroup", "assignSimpleBlueprint", armblueprint.Assignment{
		Location: to.Ptr("eastus"),
		Identity: &armblueprint.ManagedServiceIdentity{
			Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armblueprint.AssignmentProperties{
			Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
			BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
			Parameters: map[string]*armblueprint.ParameterValue{
				"costCenter": {
					Value: "Contoso/Online/Shopping/Production",
				},
				"owners": {
					Value: []any{
						"johnDoe@contoso.com",
						"johnsteam@contoso.com",
					},
				},
				"storageAccountType": {
					Value: "Standard_LRS",
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
				"storageRG": {
					Name:     to.Ptr("defaultRG"),
					Location: to.Ptr("eastus"),
				},
			},
			Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Create_SystemAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate_assignmentWithSystemAssignedManagedIdentityAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().CreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", armblueprint.Assignment{
		Location: to.Ptr("eastus"),
		Identity: &armblueprint.ManagedServiceIdentity{
			Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armblueprint.AssignmentProperties{
			Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
			BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
			Parameters: map[string]*armblueprint.ParameterValue{
				"costCenter": {
					Value: "Contoso/Online/Shopping/Production",
				},
				"owners": {
					Value: []any{
						"johnDoe@contoso.com",
						"johnsteam@contoso.com",
					},
				},
				"storageAccountType": {
					Value: "Standard_LRS",
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
				"storageRG": {
					Name:     to.Ptr("defaultRG"),
					Location: to.Ptr("eastus"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Create_UserAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate_assignmentWithUserAssignedManagedIdentityAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().CreateOrUpdate(ctx, "managementGroups/ContosoOnlineGroup", "assignSimpleBlueprint", armblueprint.Assignment{
		Location: to.Ptr("eastus"),
		Identity: &armblueprint.ManagedServiceIdentity{
			Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armblueprint.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": {},
			},
		},
		Properties: &armblueprint.AssignmentProperties{
			Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
			BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
			Parameters: map[string]*armblueprint.ParameterValue{
				"costCenter": {
					Value: "Contoso/Online/Shopping/Production",
				},
				"owners": {
					Value: []any{
						"johnDoe@contoso.com",
						"johnsteam@contoso.com",
					},
				},
				"storageAccountType": {
					Value: "Standard_LRS",
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
				"storageRG": {
					Name:     to.Ptr("defaultRG"),
					Location: to.Ptr("eastus"),
				},
			},
			Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Create_UserAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate_assignmentWithUserAssignedManagedIdentityAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().CreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", armblueprint.Assignment{
		Location: to.Ptr("eastus"),
		Identity: &armblueprint.ManagedServiceIdentity{
			Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armblueprint.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": {},
			},
		},
		Properties: &armblueprint.AssignmentProperties{
			Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
			BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
			Parameters: map[string]*armblueprint.ParameterValue{
				"costCenter": {
					Value: "Contoso/Online/Shopping/Production",
				},
				"owners": {
					Value: []any{
						"johnDoe@contoso.com",
						"johnsteam@contoso.com",
					},
				},
				"storageAccountType": {
					Value: "Standard_LRS",
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
				"storageRG": {
					Name:     to.Ptr("defaultRG"),
					Location: to.Ptr("eastus"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Get.json
func ExampleAssignmentsClient_Get_assignmentAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().Get(ctx, "managementGroups/ContosoOnlineGroup", "assignSimpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assignment = armblueprint.Assignment{
	// 	Name: to.Ptr("assignSimpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/Assignment"),
	// 	ID: to.Ptr("/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"),
	// 	Location: to.Ptr("eastus"),
	// 	Identity: &armblueprint.ManagedServiceIdentity{
	// 		Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// 	Properties: &armblueprint.AssignmentProperties{
	// 		Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
	// 		BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 		Parameters: map[string]*armblueprint.ParameterValue{
	// 			"costCenter": &armblueprint.ParameterValue{
	// 				Value: "Contoso/Online/Shopping/Production",
	// 			},
	// 			"owners": &armblueprint.ParameterValue{
	// 				Value: []any{
	// 					"johnDoe@contoso.com",
	// 					"johnsteam@contoso.com",
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterValue{
	// 				Value: "Standard_LRS",
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armblueprint.AssignmentProvisioningState("Succeeded")),
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
	// 			"storageRG": &armblueprint.ResourceGroupValue{
	// 				Name: to.Ptr("defaultRG"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 		},
	// 		Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Get.json
func ExampleAssignmentsClient_Get_assignmentAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assignment = armblueprint.Assignment{
	// 	Name: to.Ptr("assignSimpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/Assignment"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"),
	// 	Location: to.Ptr("eastus"),
	// 	Identity: &armblueprint.ManagedServiceIdentity{
	// 		Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// 	Properties: &armblueprint.AssignmentProperties{
	// 		Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
	// 		BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 		Parameters: map[string]*armblueprint.ParameterValue{
	// 			"costCenter": &armblueprint.ParameterValue{
	// 				Value: "Contoso/Online/Shopping/Production",
	// 			},
	// 			"owners": &armblueprint.ParameterValue{
	// 				Value: []any{
	// 					"johnDoe@contoso.com",
	// 					"johnsteam@contoso.com",
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterValue{
	// 				Value: "Standard_LRS",
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armblueprint.AssignmentProvisioningState("Succeeded")),
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
	// 			"storageRG": &armblueprint.ResourceGroupValue{
	// 				Name: to.Ptr("defaultRG"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Delete.json
func ExampleAssignmentsClient_Delete_assignmentDeleteAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Delete(ctx, "managementGroups/ContosoOnlineGroup", "assignSimpleBlueprint", &armblueprint.AssignmentsClientDeleteOptions{DeleteBehavior: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Delete_AndDeleteChildren.json
func ExampleAssignmentsClient_Delete_assignmentDeleteAtManagementGroupScopeAndDeleteTheResourcesCreatedByTheAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Delete(ctx, "managementGroups/ContosoOnlineGroup", "assignSimpleBlueprint", &armblueprint.AssignmentsClientDeleteOptions{DeleteBehavior: to.Ptr(armblueprint.AssignmentDeleteBehaviorAll)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Delete.json
func ExampleAssignmentsClient_Delete_assignmentDeleteAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Delete(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", &armblueprint.AssignmentsClientDeleteOptions{DeleteBehavior: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Delete_AndDeleteChildren.json
func ExampleAssignmentsClient_Delete_assignmentDeleteAtSubscriptionScopeAndDeleteTheResourcesCreatedByTheAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Delete(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", &armblueprint.AssignmentsClientDeleteOptions{DeleteBehavior: to.Ptr(armblueprint.AssignmentDeleteBehaviorAll)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/WhoIsBlueprint_Action.json
func ExampleAssignmentsClient_WhoIsBlueprint_whoIsBlueprintActionAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().WhoIsBlueprint(ctx, "managementGroups/ContosoOnlineGroup", "assignSimpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WhoIsBlueprintContract = armblueprint.WhoIsBlueprintContract{
	// 	ObjectID: to.Ptr("00000000-1111-0000-1111-000000000000"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/WhoIsBlueprint_Action.json
func ExampleAssignmentsClient_WhoIsBlueprint_whoIsBlueprintActionAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().WhoIsBlueprint(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WhoIsBlueprintContract = armblueprint.WhoIsBlueprintContract{
	// 	ObjectID: to.Ptr("00000000-1111-0000-1111-000000000000"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_List.json
func ExampleAssignmentsClient_NewListPager_assignmentAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentsClient().NewListPager("managementGroups/ContosoOnlineGroup", nil)
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
		// page.AssignmentList = armblueprint.AssignmentList{
		// 	Value: []*armblueprint.Assignment{
		// 		{
		// 			Name: to.Ptr("assignSimpleBlueprint"),
		// 			Type: to.Ptr("Microsoft.Blueprint/Assignment"),
		// 			ID: to.Ptr("/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"),
		// 			Location: to.Ptr("eastus"),
		// 			Identity: &armblueprint.ManagedServiceIdentity{
		// 				Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 			Properties: &armblueprint.AssignmentProperties{
		// 				Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
		// 				BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
		// 				Parameters: map[string]*armblueprint.ParameterValue{
		// 					"costCenter": &armblueprint.ParameterValue{
		// 						Value: "Contoso/Online/Shopping/Production",
		// 					},
		// 					"owners": &armblueprint.ParameterValue{
		// 						Value: []any{
		// 							"johnDoe@contoso.com",
		// 							"johnsteam@contoso.com",
		// 						},
		// 					},
		// 					"storageAccountType": &armblueprint.ParameterValue{
		// 						Value: "Standard_LRS",
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armblueprint.AssignmentProvisioningState("Succeeded")),
		// 				ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
		// 					"storageRG": &armblueprint.ResourceGroupValue{
		// 						Name: to.Ptr("defaultRG"),
		// 						Location: to.Ptr("eastus"),
		// 					},
		// 				},
		// 				Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_List.json
func ExampleAssignmentsClient_NewListPager_assignmentAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", nil)
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
		// page.AssignmentList = armblueprint.AssignmentList{
		// 	Value: []*armblueprint.Assignment{
		// 		{
		// 			Name: to.Ptr("assignSimpleBlueprint"),
		// 			Type: to.Ptr("Microsoft.Blueprint/Assignment"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"),
		// 			Location: to.Ptr("eastus"),
		// 			Identity: &armblueprint.ManagedServiceIdentity{
		// 				Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 			Properties: &armblueprint.AssignmentProperties{
		// 				Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
		// 				BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
		// 				Parameters: map[string]*armblueprint.ParameterValue{
		// 					"costCenter": &armblueprint.ParameterValue{
		// 						Value: "Contoso/Online/Shopping/Production",
		// 					},
		// 					"owners": &armblueprint.ParameterValue{
		// 						Value: []any{
		// 							"johnDoe@contoso.com",
		// 							"johnsteam@contoso.com",
		// 						},
		// 					},
		// 					"storageAccountType": &armblueprint.ParameterValue{
		// 						Value: "Standard_LRS",
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armblueprint.AssignmentProvisioningState("Succeeded")),
		// 				ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
		// 					"storageRG": &armblueprint.ResourceGroupValue{
		// 						Name: to.Ptr("defaultRG"),
		// 						Location: to.Ptr("eastus"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
