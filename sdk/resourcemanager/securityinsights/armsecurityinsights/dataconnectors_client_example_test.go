//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetDataConnectors.json
func ExampleDataConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataConnectorsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.DataConnectorList = armsecurityinsights.DataConnectorList{
		// 	Value: []armsecurityinsights.DataConnectorClassification{
		// 		&armsecurityinsights.ASCDataConnector{
		// 			Name: to.Ptr("763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureSecurityCenter),
		// 			Properties: &armsecurityinsights.ASCDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				SubscriptionID: to.Ptr("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0"),
		// 			},
		// 		},
		// 		&armsecurityinsights.TIDataConnector{
		// 			Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
		// 			Properties: &armsecurityinsights.TIDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
		// 					Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.AADDataConnector{
		// 			Name: to.Ptr("f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureActiveDirectory),
		// 			Properties: &armsecurityinsights.AADDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.OfficeDataConnector{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindOffice365),
		// 			Properties: &armsecurityinsights.OfficeDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.OfficeDataConnectorDataTypes{
		// 					Exchange: &armsecurityinsights.OfficeDataConnectorDataTypesExchange{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 					SharePoint: &armsecurityinsights.OfficeDataConnectorDataTypesSharePoint{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.MCASDataConnector{
		// 			Name: to.Ptr("b96d014d-b5c2-4a01-9aba-a8058f629d42"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/b96d014d-b5c2-4a01-9aba-a8058f629d42"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftCloudAppSecurity),
		// 			Properties: &armsecurityinsights.MCASDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.MCASDataConnectorDataTypes{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 					DiscoveryLogs: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.AATPDataConnector{
		// 			Name: to.Ptr("07e42cb3-e658-4e90-801c-efa0f29d3d44"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/07e42cb3-e658-4e90-801c-efa0f29d3d44"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureAdvancedThreatProtection),
		// 			Properties: &armsecurityinsights.AATPDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.AwsCloudTrailDataConnector{
		// 			Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAmazonWebServicesCloudTrail),
		// 			Properties: &armsecurityinsights.AwsCloudTrailDataConnectorProperties{
		// 				AwsRoleArn: to.Ptr("myAwsRoleArn"),
		// 				DataTypes: &armsecurityinsights.AwsCloudTrailDataConnectorDataTypes{
		// 					Logs: &armsecurityinsights.AwsCloudTrailDataConnectorDataTypesLogs{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		&armsecurityinsights.MDATPDataConnector{
		// 			Name: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
		// 			Properties: &armsecurityinsights.MDATPDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetAzureSecurityCenterById.json
func ExampleDataConnectorsClient_Get_getAAscDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "763f9fa1-c2d3-4fa2-93e9-bccd4899aa12", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.ASCDataConnector{
	// 		Name: to.Ptr("763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureSecurityCenter),
	// 		Properties: &armsecurityinsights.ASCDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
	// 				Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			SubscriptionID: to.Ptr("c0688291-89d7-4bed-87a2-a7b1bff43f4c"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetMicrosoftCloudAppSecurityById.json
func ExampleDataConnectorsClient_Get_getAMcasDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "b96d014d-b5c2-4a01-9aba-a8058f629d42", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.MCASDataConnector{
	// 		Name: to.Ptr("b96d014d-b5c2-4a01-9aba-a8058f629d42"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/b96d014d-b5c2-4a01-9aba-a8058f629d42"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftCloudAppSecurity),
	// 		Properties: &armsecurityinsights.MCASDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.MCASDataConnectorDataTypes{
	// 				Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 				DiscoveryLogs: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetMicrosoftDefenderAdvancedThreatProtectionById.json
func ExampleDataConnectorsClient_Get_getAMdatpDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "06b3ccb8-1384-4bcc-aec7-852f6d57161b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.MDATPDataConnector{
	// 		Name: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
	// 		Properties: &armsecurityinsights.MDATPDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
	// 				Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetThreatIntelligenceById.json
func ExampleDataConnectorsClient_Get_getATiDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.TIDataConnector{
	// 		Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
	// 		Properties: &armsecurityinsights.TIDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
	// 				Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 			TipLookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T13:00:30.123Z"); return t}()),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetAzureActiveDirectoryById.json
func ExampleDataConnectorsClient_Get_getAnAadDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.AADDataConnector{
	// 		Name: to.Ptr("f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureActiveDirectory),
	// 		Properties: &armsecurityinsights.AADDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
	// 				Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetAzureAdvancedThreatProtectionById.json
func ExampleDataConnectorsClient_Get_getAnAatpDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "07e42cb3-e658-4e90-801c-efa0f29d3d44", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.AATPDataConnector{
	// 		Name: to.Ptr("07e42cb3-e658-4e90-801c-efa0f29d3d44"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/07e42cb3-e658-4e90-801c-efa0f29d3d44"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureAdvancedThreatProtection),
	// 		Properties: &armsecurityinsights.AATPDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
	// 				Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetAmazonWebServicesCloudTrailById.json
func ExampleDataConnectorsClient_Get_getAnAwsCloudTrailDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.AwsCloudTrailDataConnector{
	// 		Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAmazonWebServicesCloudTrail),
	// 		Properties: &armsecurityinsights.AwsCloudTrailDataConnectorProperties{
	// 			AwsRoleArn: to.Ptr("myAwsRoleArn"),
	// 			DataTypes: &armsecurityinsights.AwsCloudTrailDataConnectorDataTypes{
	// 				Logs: &armsecurityinsights.AwsCloudTrailDataConnectorDataTypesLogs{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetOfficeDataConnetorById.json
func ExampleDataConnectorsClient_Get_getAnOffice365DataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.OfficeDataConnector{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindOffice365),
	// 		Properties: &armsecurityinsights.OfficeDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.OfficeDataConnectorDataTypes{
	// 				Exchange: &armsecurityinsights.OfficeDataConnectorDataTypesExchange{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 				SharePoint: &armsecurityinsights.OfficeDataConnectorDataTypesSharePoint{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 				Teams: &armsecurityinsights.OfficeDataConnectorDataTypesTeams{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/CreateOfficeDataConnetor.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAnOffice365DataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.OfficeDataConnector{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindOffice365),
		Properties: &armsecurityinsights.OfficeDataConnectorProperties{
			DataTypes: &armsecurityinsights.OfficeDataConnectorDataTypes{
				Exchange: &armsecurityinsights.OfficeDataConnectorDataTypesExchange{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
				SharePoint: &armsecurityinsights.OfficeDataConnectorDataTypesSharePoint{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
				Teams: &armsecurityinsights.OfficeDataConnectorDataTypesTeams{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.OfficeDataConnector{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindOffice365),
	// 		Properties: &armsecurityinsights.OfficeDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.OfficeDataConnectorDataTypes{
	// 				Exchange: &armsecurityinsights.OfficeDataConnectorDataTypesExchange{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 				SharePoint: &armsecurityinsights.OfficeDataConnectorDataTypesSharePoint{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 				Teams: &armsecurityinsights.OfficeDataConnectorDataTypesTeams{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/CreateThreatIntelligenceDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAnThreatIntelligencePlatformDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.TIDataConnector{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
		Properties: &armsecurityinsights.TIDataConnectorProperties{
			DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
				Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			TenantID:          to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
			TipLookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T13:00:30.123Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.TIDataConnector{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
	// 		Properties: &armsecurityinsights.TIDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
	// 				Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
	// 			TipLookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T13:00:30.123Z"); return t}()),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/DeleteOfficeDataConnetor.json
func ExampleDataConnectorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataConnectorsClient().Delete(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
