//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspacesInResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspace.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.WorkspacePatchInfo{
			Identity: &armsynapse.ManagedIdentity{
				Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armsynapse.WorkspacePatchProperties{
				Encryption: &armsynapse.EncryptionDetails{
					Cmk: &armsynapse.CustomerManagedKeyDetails{
						Key: &armsynapse.WorkspaceKeyDetails{
							Name:        to.Ptr("<name>"),
							KeyVaultURL: to.Ptr("<key-vault-url>"),
						},
					},
				},
				ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
					AllowedAADTenantIDsForLinking: []*string{
						to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
					LinkedAccessCheckOnTargetResource: to.Ptr(false),
					PreventDataExfiltration:           to.Ptr(false),
				},
				PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
				PurviewConfiguration: &armsynapse.PurviewConfiguration{
					PurviewResourceID: to.Ptr("<purview-resource-id>"),
				},
				SQLAdministratorLoginPassword: to.Ptr("<sqladministrator-login-password>"),
				WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
					Type:                to.Ptr("<type>"),
					AccountName:         to.Ptr("<account-name>"),
					CollaborationBranch: to.Ptr("<collaboration-branch>"),
					HostName:            to.Ptr("<host-name>"),
					ProjectName:         to.Ptr("<project-name>"),
					RepositoryName:      to.Ptr("<repository-name>"),
					RootFolder:          to.Ptr("<root-folder>"),
				},
			},
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
		},
		&armsynapse.WorkspacesClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.Workspace{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Identity: &armsynapse.ManagedIdentity{
				Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armsynapse.UserAssignedManagedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": {},
				},
			},
			Properties: &armsynapse.WorkspaceProperties{
				CspWorkspaceAdminProperties: &armsynapse.CspWorkspaceAdminProperties{
					InitialWorkspaceAdminObjectID: to.Ptr("<initial-workspace-admin-object-id>"),
				},
				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
					AccountURL: to.Ptr("<account-url>"),
					Filesystem: to.Ptr("<filesystem>"),
				},
				Encryption: &armsynapse.EncryptionDetails{
					Cmk: &armsynapse.CustomerManagedKeyDetails{
						KekIdentity: &armsynapse.KekIdentityProperties{
							UseSystemAssignedIdentity: false,
							UserAssignedIdentity:      to.Ptr("<user-assigned-identity>"),
						},
						Key: &armsynapse.WorkspaceKeyDetails{
							Name:        to.Ptr("<name>"),
							KeyVaultURL: to.Ptr("<key-vault-url>"),
						},
					},
				},
				ManagedResourceGroupName: to.Ptr("<managed-resource-group-name>"),
				ManagedVirtualNetwork:    to.Ptr("<managed-virtual-network>"),
				ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
					AllowedAADTenantIDsForLinking: []*string{
						to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
					LinkedAccessCheckOnTargetResource: to.Ptr(false),
					PreventDataExfiltration:           to.Ptr(false),
				},
				PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
				PurviewConfiguration: &armsynapse.PurviewConfiguration{
					PurviewResourceID: to.Ptr("<purview-resource-id>"),
				},
				SQLAdministratorLogin:         to.Ptr("<sqladministrator-login>"),
				SQLAdministratorLoginPassword: to.Ptr("<sqladministrator-login-password>"),
				WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
					Type:                to.Ptr("<type>"),
					AccountName:         to.Ptr("<account-name>"),
					CollaborationBranch: to.Ptr("<collaboration-branch>"),
					HostName:            to.Ptr("<host-name>"),
					ProjectName:         to.Ptr("<project-name>"),
					RepositoryName:      to.Ptr("<repository-name>"),
					RootFolder:          to.Ptr("<root-folder>"),
				},
			},
		},
		&armsynapse.WorkspacesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspace.json
func ExampleWorkspacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		&armsynapse.WorkspacesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspacesInSubscription.json
func ExampleWorkspacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
