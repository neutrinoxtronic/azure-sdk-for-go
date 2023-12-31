//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansInOrderList.json
func ExampleSavingsPlanClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlanClient().NewListPager("20000000-0000-0000-0000-000000000000", nil)
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
		// page.SavingsPlanModelList = armbillingbenefits.SavingsPlanModelList{
		// 	Value: []*armbillingbenefits.SavingsPlanModel{
		// 		{
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000000/30000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbillingbenefits.SavingsPlanModelProperties{
		// 				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
		// 					ResourceGroupID: to.Ptr("/subscriptions/eef82110-c91b-4395-9420-fcfcbefc5a47/resourcegroups/3ppRG"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 				BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000009"),
		// 				Commitment: &armbillingbenefits.Commitment{
		// 					Amount: to.Ptr[float64](0.002),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("Compute_SavingsPlan_10-24-2022_15-44"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-24T22:47:04.8539493Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-24T22:47:04.6196635Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-24T22:45:13.6202486Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbillingbenefits.TermP1Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("ResourceGroup"),
		// 				Utilization: &armbillingbenefits.Utilization{
		// 					Aggregates: []*armbillingbenefits.UtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr(""),
		// 				},
		// 			},
		// 			SKU: &armbillingbenefits.SKU{
		// 				Name: to.Ptr("Compute_Savings_Plan"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansList.json
func ExampleSavingsPlanClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlanClient().NewListAllPager(&armbillingbenefits.SavingsPlanClientListAllOptions{Filter: to.Ptr("(properties%2farchived+eq+false)"),
		Orderby:        to.Ptr("properties/displayName asc"),
		RefreshSummary: nil,
		Skiptoken:      to.Ptr[float32](50),
		SelectedState:  nil,
		Take:           to.Ptr[float32](1),
	})
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
		// page.SavingsPlanModelListResult = armbillingbenefits.SavingsPlanModelListResult{
		// 	AdditionalProperties: []*armbillingbenefits.SavingsPlanSummary{
		// 		{
		// 			Name: to.Ptr("summary"),
		// 			Value: &armbillingbenefits.SavingsPlanSummaryCount{
		// 				CancelledCount: to.Ptr[float32](0),
		// 				ExpiredCount: to.Ptr[float32](0),
		// 				ExpiringCount: to.Ptr[float32](0),
		// 				FailedCount: to.Ptr[float32](0),
		// 				NoBenefitCount: to.Ptr[float32](0),
		// 				PendingCount: to.Ptr[float32](0),
		// 				ProcessingCount: to.Ptr[float32](0),
		// 				SucceededCount: to.Ptr[float32](1),
		// 				WarningCount: to.Ptr[float32](0),
		// 			},
		// 	}},
		// 	Value: []*armbillingbenefits.SavingsPlanModel{
		// 		{
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000000/30000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbillingbenefits.SavingsPlanModelProperties{
		// 				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
		// 					DisplayName: to.Ptr("Azure subscription 1"),
		// 					SubscriptionID: to.Ptr("/subscriptions/20000000-0000-0000-0000-000000000005"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 				BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000009"),
		// 				Commitment: &armbillingbenefits.Commitment{
		// 					Amount: to.Ptr[float64](0.001),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("Compute_SavingsPlan_10-19-2022_11-03"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-19T18:05:37.1034288Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T18:05:36.5252231Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-19T18:03:56.4032132Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbillingbenefits.TermP1Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Single"),
		// 				Utilization: &armbillingbenefits.Utilization{
		// 					Aggregates: []*armbillingbenefits.UtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](100),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](78),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](78.12),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr(""),
		// 				},
		// 			},
		// 			SKU: &armbillingbenefits.SKU{
		// 				Name: to.Ptr("Compute_Savings_Plan"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanItemGet.json
func ExampleSavingsPlanClient_Get_savingsPlanItemGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().Get(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", &armbillingbenefits.SavingsPlanClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanModel = armbillingbenefits.SavingsPlanModel{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000/30000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 	Properties: &armbillingbenefits.SavingsPlanModelProperties{
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeShared),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 		BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000009"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("Compute_SavingsPlan_patch_rename2"),
	// 		DisplayProvisioningState: to.Ptr("Succeeded"),
	// 		EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-21T18:15:42.4098626Z"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-18T21:16:13.1854959Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-18T21:14:26.8279361Z"); return t}()),
	// 		Renew: to.Ptr(true),
	// 		Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 		UserFriendlyAppliedScopeType: to.Ptr("Shared"),
	// 		Utilization: &armbillingbenefits.Utilization{
	// 			Aggregates: []*armbillingbenefits.UtilizationAggregates{
	// 				{
	// 					Grain: to.Ptr[float32](1),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](100),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](7),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](84),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](30),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](83.87),
	// 					ValueUnit: to.Ptr("percentage"),
	// 			}},
	// 			Trend: to.Ptr(""),
	// 		},
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanItemExpandedGet.json
func ExampleSavingsPlanClient_Get_savingsPlanItemWithExpandedRenewPropertiesGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().Get(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", &armbillingbenefits.SavingsPlanClientGetOptions{Expand: to.Ptr("renewProperties")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanModel = armbillingbenefits.SavingsPlanModel{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000/30000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 	Properties: &armbillingbenefits.SavingsPlanModelProperties{
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeShared),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 		BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000009"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("Compute_SavingsPlan_patch_rename2"),
	// 		DisplayProvisioningState: to.Ptr("Succeeded"),
	// 		EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-21T18:15:42.4098626Z"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-18T21:16:13.1854959Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-18T21:14:26.8279361Z"); return t}()),
	// 		Renew: to.Ptr(true),
	// 		RenewProperties: &armbillingbenefits.RenewProperties{
	// 			PurchaseProperties: &armbillingbenefits.PurchaseRequest{
	// 				Properties: &armbillingbenefits.PurchaseRequestProperties{
	// 					AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeShared),
	// 					BillingPlan: to.Ptr(armbillingbenefits.BillingPlan("Upfront")),
	// 					BillingScopeID: to.Ptr("/subscriptions/eef82110-c91b-4395-9420-fcfcbefc5a47"),
	// 					Commitment: &armbillingbenefits.Commitment{
	// 						Amount: to.Ptr[float64](0.001),
	// 						CurrencyCode: to.Ptr("USD"),
	// 						Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 					},
	// 					DisplayName: to.Ptr("Compute_SavingsPlan_patch_rename2_renewed"),
	// 					Term: to.Ptr(armbillingbenefits.TermP1Y),
	// 				},
	// 				SKU: &armbillingbenefits.SKU{
	// 					Name: to.Ptr("Compute_Savings_Plan"),
	// 				},
	// 			},
	// 		},
	// 		Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 		UserFriendlyAppliedScopeType: to.Ptr("Shared"),
	// 		Utilization: &armbillingbenefits.Utilization{
	// 			Aggregates: []*armbillingbenefits.UtilizationAggregates{
	// 				{
	// 					Grain: to.Ptr[float32](1),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](100),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](7),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](84),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](30),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](83.87),
	// 					ValueUnit: to.Ptr("percentage"),
	// 			}},
	// 			Trend: to.Ptr(""),
	// 		},
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanUpdate.json
func ExampleSavingsPlanClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().Update(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbillingbenefits.SavingsPlanUpdateRequest{
		Properties: &armbillingbenefits.SavingsPlanUpdateRequestProperties{
			AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
				ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
			},
			AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
			DisplayName:      to.Ptr("TestDisplayName"),
			Renew:            to.Ptr(true),
			RenewProperties: &armbillingbenefits.RenewProperties{
				PurchaseProperties: &armbillingbenefits.PurchaseRequest{
					Properties: &armbillingbenefits.PurchaseRequestProperties{
						AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
							ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
						},
						AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
						BillingPlan:      to.Ptr(armbillingbenefits.BillingPlanP1M),
						BillingScopeID:   to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
						Commitment: &armbillingbenefits.Commitment{
							Amount:       to.Ptr[float64](15.23),
							CurrencyCode: to.Ptr("USD"),
							Grain:        to.Ptr(armbillingbenefits.CommitmentGrainHourly),
						},
						DisplayName: to.Ptr("TestDisplayName_renewed"),
						Renew:       to.Ptr(false),
						Term:        to.Ptr(armbillingbenefits.TermP1Y),
					},
					SKU: &armbillingbenefits.SKU{
						Name: to.Ptr("Compute_Savings_Plan"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanModel = armbillingbenefits.SavingsPlanModel{
	// 	Name: to.Ptr("30000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("microsoft.billingbenefits/savingsPlanOrders/savingsPlans"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 	Properties: &armbillingbenefits.SavingsPlanModelProperties{
	// 		AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
	// 			DisplayName: to.Ptr("Azure subscription 1"),
	// 			SubscriptionID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 		},
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-27T00:34:33.6697600Z"); return t}()),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 		BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("riName"),
	// 		DisplayProvisioningState: to.Ptr("Succeeded"),
	// 		EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T00:12:54.549373Z"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-27T00:34:33.6697600Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-27T00:32:45.5823486Z"); return t}()),
	// 		Renew: to.Ptr(true),
	// 		Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 		UserFriendlyAppliedScopeType: to.Ptr("Single"),
	// 		Utilization: &armbillingbenefits.Utilization{
	// 			Aggregates: []*armbillingbenefits.UtilizationAggregates{
	// 				{
	// 					Grain: to.Ptr[float32](1),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](100),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](7),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](37),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](30),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](53.85),
	// 					ValueUnit: to.Ptr("percentage"),
	// 			}},
	// 			Trend: to.Ptr("DOWN"),
	// 		},
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanValidateUpdate.json
func ExampleSavingsPlanClient_ValidateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().ValidateUpdate(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbillingbenefits.SavingsPlanUpdateValidateRequest{
		Benefits: []*armbillingbenefits.SavingsPlanUpdateRequestProperties{
			{
				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/30000000-0000-0000-0000-000000000100"),
					TenantID:          to.Ptr("30000000-0000-0000-0000-000000000100"),
				},
				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeManagementGroup),
			},
			{
				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/MockMG"),
					TenantID:          to.Ptr("30000000-0000-0000-0000-000000000100"),
				},
				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeManagementGroup),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanValidateResponse = armbillingbenefits.SavingsPlanValidateResponse{
	// 	Benefits: []*armbillingbenefits.SavingsPlanValidResponseProperty{
	// 		{
	// 			Valid: to.Ptr(true),
	// 		},
	// 		{
	// 			Valid: to.Ptr(true),
	// 	}},
	// }
}
