//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineAssessPatches.json
func ExampleVirtualMachinesClient_BeginAssessPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginAssessPatches(ctx, "myResourceGroupName", "myMachineName", nil)
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
	// res.VirtualMachineAssessPatchesResult = armconnectedvmware.VirtualMachineAssessPatchesResult{
	// 	AssessmentActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 	AvailablePatchCountByClassification: &armconnectedvmware.AvailablePatchCountByClassification{
	// 		Critical: to.Ptr[int32](0),
	// 		Definition: to.Ptr[int32](0),
	// 		FeaturePack: to.Ptr[int32](0),
	// 		Security: to.Ptr[int32](0),
	// 		ServicePack: to.Ptr[int32](0),
	// 		Tools: to.Ptr[int32](0),
	// 		UpdateRollup: to.Ptr[int32](1),
	// 		Updates: to.Ptr[int32](1),
	// 	},
	// 	LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:16:06.9740000Z"); return t}()),
	// 	OSType: to.Ptr(armconnectedvmware.OsTypeUMWindows),
	// 	RebootPending: to.Ptr(true),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:15:20.9340000Z"); return t}()),
	// 	StartedBy: to.Ptr(armconnectedvmware.PatchOperationStartedByUser),
	// 	Status: to.Ptr(armconnectedvmware.PatchOperationStatusSucceeded),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineInstallPatches.json
func ExampleVirtualMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginInstallPatches(ctx, "myResourceGroupName", "myMachineName", armconnectedvmware.VirtualMachineInstallPatchesParameters{
		MaximumDuration: to.Ptr("PT3H"),
		RebootSetting:   to.Ptr(armconnectedvmware.VMGuestPatchRebootSettingIfRequired),
		WindowsParameters: &armconnectedvmware.WindowsParameters{
			ClassificationsToInclude: []*armconnectedvmware.VMGuestPatchClassificationWindows{
				to.Ptr(armconnectedvmware.VMGuestPatchClassificationWindowsCritical),
				to.Ptr(armconnectedvmware.VMGuestPatchClassificationWindowsSecurity)},
			MaxPatchPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:36:43.0539904+00:00"); return t }()),
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
	// res.VirtualMachineInstallPatchesResult = armconnectedvmware.VirtualMachineInstallPatchesResult{
	// 	ExcludedPatchCount: to.Ptr[int32](0),
	// 	FailedPatchCount: to.Ptr[int32](0),
	// 	InstallationActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 	InstalledPatchCount: to.Ptr[int32](3),
	// 	LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-15T02:16:06.9740000Z"); return t}()),
	// 	MaintenanceWindowExceeded: to.Ptr(false),
	// 	NotSelectedPatchCount: to.Ptr[int32](0),
	// 	OSType: to.Ptr(armconnectedvmware.OsTypeUMWindows),
	// 	PendingPatchCount: to.Ptr[int32](2),
	// 	RebootStatus: to.Ptr(armconnectedvmware.VMGuestPatchRebootStatusCompleted),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:15:06.9740000Z"); return t}()),
	// 	StartedBy: to.Ptr(armconnectedvmware.PatchOperationStartedByUser),
	// 	Status: to.Ptr(armconnectedvmware.PatchOperationStatusSucceeded),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreate(ctx, "testrg", "DemoVM", armconnectedvmware.VirtualMachine{
		ExtendedLocation: &armconnectedvmware.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
			Type: to.Ptr("customLocation"),
		},
		Location: to.Ptr("East US"),
		Properties: &armconnectedvmware.VirtualMachineProperties{
			HardwareProfile: &armconnectedvmware.HardwareProfile{
				MemorySizeMB: to.Ptr[int32](4196),
				NumCPUs:      to.Ptr[int32](4),
			},
			ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
			TemplateID:     to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
			VCenterID:      to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
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
	// res.VirtualMachine = armconnectedvmware.VirtualMachine{
	// 	Name: to.Ptr("DemoVM"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines"),
	// 	ExtendedLocation: &armconnectedvmware.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachines/DemoVM"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armconnectedvmware.VirtualMachineProperties{
	// 		HardwareProfile: &armconnectedvmware.HardwareProfile{
	// 			MemorySizeMB: to.Ptr[int32](4196),
	// 			NumCPUs: to.Ptr[int32](4),
	// 		},
	// 		MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		OSProfile: &armconnectedvmware.OsProfile{
	// 			ComputerName: to.Ptr("DemoVM"),
	// 			OSType: to.Ptr(armconnectedvmware.OsTypeWindows),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
	// 		TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
	// 		VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetVirtualMachine.json
func ExampleVirtualMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "testrg", "DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachine = armconnectedvmware.VirtualMachine{
	// 	Name: to.Ptr("DemoVM"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines"),
	// 	ExtendedLocation: &armconnectedvmware.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachines/DemoVM"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armconnectedvmware.VirtualMachineProperties{
	// 		HardwareProfile: &armconnectedvmware.HardwareProfile{
	// 			MemorySizeMB: to.Ptr[int32](4196),
	// 			NumCPUs: to.Ptr[int32](4),
	// 		},
	// 		MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		OSProfile: &armconnectedvmware.OsProfile{
	// 			ComputerName: to.Ptr("DemoVM"),
	// 			OSType: to.Ptr(armconnectedvmware.OsTypeWindows),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
	// 		TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
	// 		VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/UpdateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginUpdate(ctx, "testrg", "DemoVM", armconnectedvmware.VirtualMachineUpdate{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.VirtualMachine = armconnectedvmware.VirtualMachine{
	// 	Name: to.Ptr("DemoVM"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines"),
	// 	ExtendedLocation: &armconnectedvmware.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachines/DemoVM"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armconnectedvmware.VirtualMachineProperties{
	// 		HardwareProfile: &armconnectedvmware.HardwareProfile{
	// 			MemorySizeMB: to.Ptr[int32](4196),
	// 			NumCPUs: to.Ptr[int32](4),
	// 		},
	// 		MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		OSProfile: &armconnectedvmware.OsProfile{
	// 			ComputerName: to.Ptr("DemoVM"),
	// 			OSType: to.Ptr(armconnectedvmware.OsTypeWindows),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
	// 		TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
	// 		VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteVirtualMachine.json
func ExampleVirtualMachinesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginDelete(ctx, "testrg", "DemoVM", &armconnectedvmware.VirtualMachinesClientBeginDeleteOptions{Force: nil,
		Retain: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/StopVirtualMachine.json
func ExampleVirtualMachinesClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginStop(ctx, "testrg", "DemoVM", &armconnectedvmware.VirtualMachinesClientBeginStopOptions{Body: &armconnectedvmware.StopVirtualMachineOptions{
		SkipShutdown: to.Ptr(true),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/StartVirtualMachine.json
func ExampleVirtualMachinesClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginStart(ctx, "testrg", "DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/RestartVirtualMachine.json
func ExampleVirtualMachinesClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginRestart(ctx, "testrg", "DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualMachines.json
func ExampleVirtualMachinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListPager(nil)
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
		// page.VirtualMachinesList = armconnectedvmware.VirtualMachinesList{
		// 	Value: []*armconnectedvmware.VirtualMachine{
		// 		{
		// 			Name: to.Ptr("DemoVM"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines"),
		// 			ExtendedLocation: &armconnectedvmware.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachines/DemoVM"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armconnectedvmware.VirtualMachineProperties{
		// 				HardwareProfile: &armconnectedvmware.HardwareProfile{
		// 					MemorySizeMB: to.Ptr[int32](4196),
		// 					NumCPUs: to.Ptr[int32](4),
		// 				},
		// 				MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 				OSProfile: &armconnectedvmware.OsProfile{
		// 					ComputerName: to.Ptr("DemoVM"),
		// 					OSType: to.Ptr(armconnectedvmware.OsTypeWindows),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
		// 				TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
		// 				VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualMachinesByResourceGroup.json
func ExampleVirtualMachinesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListByResourceGroupPager("testrg", nil)
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
		// page.VirtualMachinesList = armconnectedvmware.VirtualMachinesList{
		// 	Value: []*armconnectedvmware.VirtualMachine{
		// 		{
		// 			Name: to.Ptr("DemoVM"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines"),
		// 			ExtendedLocation: &armconnectedvmware.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachines/DemoVM"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armconnectedvmware.VirtualMachineProperties{
		// 				HardwareProfile: &armconnectedvmware.HardwareProfile{
		// 					MemorySizeMB: to.Ptr[int32](4196),
		// 					NumCPUs: to.Ptr[int32](4),
		// 				},
		// 				MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 				OSProfile: &armconnectedvmware.OsProfile{
		// 					ComputerName: to.Ptr("DemoVM"),
		// 					OSType: to.Ptr(armconnectedvmware.OsTypeWindows),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
		// 				TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
		// 				VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
		// 			},
		// 	}},
		// }
	}
}
