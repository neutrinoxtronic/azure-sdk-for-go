//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createTestJob.json
func ExampleTestJobClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTestJobClient().Create(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", armautomation.TestJobCreateParameters{
		Parameters: map[string]*string{
			"key01": to.Ptr("value01"),
			"key02": to.Ptr("value02"),
		},
		RunOn: to.Ptr(""),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getTestJob.json
func ExampleTestJobClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTestJobClient().Get(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestJob = armautomation.TestJob{
	// 	CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T05:52:00.9344172+00:00"); return t}()),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T05:55:16.7827254+00:00"); return t}()),
	// 	LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T05:55:16.7827254+00:00"); return t}()),
	// 	LastStatusModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T05:55:16.7827254+00:00"); return t}()),
	// 	Parameters: map[string]*string{
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T05:55:10.8344743+00:00"); return t}()),
	// 	Status: to.Ptr("Completed"),
	// 	StatusDetails: to.Ptr("None"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/resumeTestJob.json
func ExampleTestJobClient_Resume() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTestJobClient().Resume(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/stopTestJob.json
func ExampleTestJobClient_Stop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTestJobClient().Stop(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/suspendTestJob.json
func ExampleTestJobClient_Suspend() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTestJobClient().Suspend(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
