//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listLocations.json
func ExampleSubscriptionsClient_NewListLocationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListLocationsPager("83aa47df-e3e9-49ff-877b-94304bf3d3ad", nil)
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
		// page.LocationListResult = armsubscription.LocationListResult{
		// 	Value: []*armsubscription.Location{
		// 		{
		// 			Name: to.Ptr("eastasia"),
		// 			DisplayName: to.Ptr("East Asia"),
		// 			ID: to.Ptr("/subscriptions/83aa47df-e3e9-49ff-877b-94304bf3d3ad/locations/eastasia"),
		// 			Latitude: to.Ptr("22.267"),
		// 			Longitude: to.Ptr("114.188"),
		// 		},
		// 		{
		// 			Name: to.Ptr("southeastasia"),
		// 			DisplayName: to.Ptr("Southeast Asia"),
		// 			ID: to.Ptr("/subscriptions/83aa47df-e3e9-49ff-877b-94304bf3d3ad/locations/southeastasia"),
		// 			Latitude: to.Ptr("1.283"),
		// 			Longitude: to.Ptr("103.833"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/getSubscription.json
func ExampleSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().Get(ctx, "83aa47df-e3e9-49ff-877b-94304bf3d3ad", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Subscription = armsubscription.Subscription{
	// 	AuthorizationSource: to.Ptr("Legacy"),
	// 	DisplayName: to.Ptr("Subscription2"),
	// 	ID: to.Ptr("/subscriptions/83aa47df-e3e9-49ff-877b-94304bf3d3ad"),
	// 	State: to.Ptr(armsubscription.SubscriptionStateEnabled),
	// 	SubscriptionID: to.Ptr("83aa47df-e3e9-49ff-877b-94304bf3d3ad"),
	// 	SubscriptionPolicies: &armsubscription.Policies{
	// 		LocationPlacementID: to.Ptr("Internal_2014-09-01"),
	// 		QuotaID: to.Ptr("Internal_2014-09-01"),
	// 		SpendingLimit: to.Ptr(armsubscription.SpendingLimitOff),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listSubscriptions.json
func ExampleSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListPager(nil)
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
		// page.ListResult = armsubscription.ListResult{
		// 	Value: []*armsubscription.Subscription{
		// 		{
		// 			AuthorizationSource: to.Ptr("Legacy"),
		// 			DisplayName: to.Ptr("SubscriptionName"),
		// 			ID: to.Ptr("/subscriptions/83aa47df-e3e9-49ff-877b-94304bf3d3ad"),
		// 			State: to.Ptr(armsubscription.SubscriptionStateEnabled),
		// 			SubscriptionID: to.Ptr("83aa47df-e3e9-49ff-877b-94304bf3d3ad"),
		// 			SubscriptionPolicies: &armsubscription.Policies{
		// 				LocationPlacementID: to.Ptr("Internal_2014-09-01"),
		// 				QuotaID: to.Ptr("Internal_2014-09-01"),
		// 				SpendingLimit: to.Ptr(armsubscription.SpendingLimitOff),
		// 			},
		// 	}},
		// }
	}
}
