//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroup.json
func ExampleSQLPoolWorkloadGroupClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolWorkloadGroupClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "smallrc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkloadGroup = armsynapse.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/smallrc"),
	// 	Properties: &armsynapse.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadGroupMax.json
func ExampleSQLPoolWorkloadGroupClient_BeginCreateOrUpdate_createAWorkloadGroupWithAllPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolWorkloadGroupClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "smallrc", armsynapse.WorkloadGroup{
		Properties: &armsynapse.WorkloadGroupProperties{
			Importance:                   to.Ptr("normal"),
			MaxResourcePercent:           to.Ptr[int32](100),
			MaxResourcePercentPerRequest: to.Ptr[float64](3),
			MinResourcePercent:           to.Ptr[int32](0),
			MinResourcePercentPerRequest: to.Ptr[float64](3),
			QueryExecutionTimeout:        to.Ptr[int32](0),
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
	// res.WorkloadGroup = armsynapse.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/workloadGroups/smallrc"),
	// 	Properties: &armsynapse.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadGroupMin.json
func ExampleSQLPoolWorkloadGroupClient_BeginCreateOrUpdate_createAWorkloadGroupWithTheRequiredPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolWorkloadGroupClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "smallrc", armsynapse.WorkloadGroup{
		Properties: &armsynapse.WorkloadGroupProperties{
			MaxResourcePercent:           to.Ptr[int32](100),
			MinResourcePercent:           to.Ptr[int32](0),
			MinResourcePercentPerRequest: to.Ptr[float64](3),
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
	// res.WorkloadGroup = armsynapse.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/workloadGroups/smallrc"),
	// 	Properties: &armsynapse.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolWorkloadGroup.json
func ExampleSQLPoolWorkloadGroupClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolWorkloadGroupClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "wlm_workloadgroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupList.json
func ExampleSQLPoolWorkloadGroupClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolWorkloadGroupClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", nil)
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
		// page.WorkloadGroupListResult = armsynapse.WorkloadGroupListResult{
		// 	Value: []*armsynapse.WorkloadGroup{
		// 		{
		// 			Name: to.Ptr("smallrc"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/smallrc"),
		// 			Properties: &armsynapse.WorkloadGroupProperties{
		// 				Importance: to.Ptr("normal"),
		// 				MaxResourcePercent: to.Ptr[int32](100),
		// 				MaxResourcePercentPerRequest: to.Ptr[float64](5),
		// 				MinResourcePercent: to.Ptr[int32](0),
		// 				MinResourcePercentPerRequest: to.Ptr[float64](5),
		// 				QueryExecutionTimeout: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mediumrc"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/mediumrc"),
		// 			Properties: &armsynapse.WorkloadGroupProperties{
		// 				Importance: to.Ptr("normal"),
		// 				MaxResourcePercent: to.Ptr[int32](100),
		// 				MaxResourcePercentPerRequest: to.Ptr[float64](10),
		// 				MinResourcePercent: to.Ptr[int32](0),
		// 				MinResourcePercentPerRequest: to.Ptr[float64](10),
		// 				QueryExecutionTimeout: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("largerc"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/workloadGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-2080/sqlPools/sqlcrudtest-9187/workloadGroups/largerc"),
		// 			Properties: &armsynapse.WorkloadGroupProperties{
		// 				Importance: to.Ptr("high"),
		// 				MaxResourcePercent: to.Ptr[int32](100),
		// 				MaxResourcePercentPerRequest: to.Ptr[float64](20),
		// 				MinResourcePercent: to.Ptr[int32](0),
		// 				MinResourcePercentPerRequest: to.Ptr[float64](20),
		// 				QueryExecutionTimeout: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
