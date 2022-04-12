//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_ListByRuleSet.json
func ExampleRulesClient_ListByRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByRuleSet("<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Get.json
func ExampleRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Create.json
func ExampleRulesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		armcdn.Rule{
			Properties: &armcdn.RuleProperties{
				Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
					&armcdn.DeliveryRuleResponseHeaderAction{
						Name: to.Ptr(armcdn.DeliveryRuleActionModifyResponseHeader),
						Parameters: &armcdn.HeaderActionParameters{
							HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
							HeaderName:   to.Ptr("<header-name>"),
							TypeName:     to.Ptr(armcdn.HeaderActionParametersTypeNameDeliveryRuleHeaderActionParameters),
							Value:        to.Ptr("<value>"),
						},
					}},
				Conditions: []armcdn.DeliveryRuleConditionClassification{
					&armcdn.DeliveryRuleRequestMethodCondition{
						Name: to.Ptr(armcdn.MatchVariableRequestMethod),
						Parameters: &armcdn.RequestMethodMatchConditionParameters{
							MatchValues: []*armcdn.RequestMethodMatchConditionParametersMatchValuesItem{
								to.Ptr(armcdn.RequestMethodMatchConditionParametersMatchValuesItemGET)},
							NegateCondition: to.Ptr(false),
							Operator:        to.Ptr(armcdn.RequestMethodOperatorEqual),
							TypeName:        to.Ptr(armcdn.RequestMethodMatchConditionParametersTypeNameDeliveryRuleRequestMethodConditionParameters),
						},
					}},
				Order: to.Ptr[int32](1),
			},
		},
		&armcdn.RulesClientBeginCreateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Update.json
func ExampleRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		armcdn.RuleUpdateParameters{
			Properties: &armcdn.RuleUpdatePropertiesParameters{
				Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
					&armcdn.DeliveryRuleResponseHeaderAction{
						Name: to.Ptr(armcdn.DeliveryRuleActionModifyResponseHeader),
						Parameters: &armcdn.HeaderActionParameters{
							HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
							HeaderName:   to.Ptr("<header-name>"),
							TypeName:     to.Ptr(armcdn.HeaderActionParametersTypeNameDeliveryRuleHeaderActionParameters),
							Value:        to.Ptr("<value>"),
						},
					}},
				Order: to.Ptr[int32](1),
			},
		},
		&armcdn.RulesClientBeginUpdateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Delete.json
func ExampleRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		&armcdn.RulesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
