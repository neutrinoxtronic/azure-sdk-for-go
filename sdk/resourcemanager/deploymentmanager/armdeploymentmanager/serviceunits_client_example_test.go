//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate_noartifactsource.json
func ExampleServiceUnitsClient_BeginCreateOrUpdate_createServiceUnitUsingSasUrIs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceUnitsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myTopology", "myService", "myServiceUnit", armdeploymentmanager.ServiceUnitResource{
		Location: to.Ptr("centralus"),
		Tags:     map[string]*string{},
		Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
			Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
				ParametersURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myartifactsource/parameter/myTopologyUnit.parameters.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"),
				TemplateURI:   to.Ptr("https://mystorageaccount.blob.core.windows.net/myartifactsource/templates/myTopologyUnit.template.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"),
			},
			DeploymentMode:      to.Ptr(armdeploymentmanager.DeploymentModeIncremental),
			TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate.json
func ExampleServiceUnitsClient_BeginCreateOrUpdate_createServiceUnitUsingRelativePathsIntoTheArtifactSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceUnitsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myTopology", "myService", "myServiceUnit", armdeploymentmanager.ServiceUnitResource{
		Location: to.Ptr("centralus"),
		Tags:     map[string]*string{},
		Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
			Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
				ParametersArtifactSourceRelativePath: to.Ptr("parameter/myTopologyUnit.parameters.json"),
				TemplateArtifactSourceRelativePath:   to.Ptr("templates/myTopologyUnit.template.json"),
			},
			DeploymentMode:      to.Ptr(armdeploymentmanager.DeploymentModeIncremental),
			TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_get.json
func ExampleServiceUnitsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceUnitsClient().Get(ctx, "myResourceGroup", "myTopology", "myService", "myServiceUnit", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceUnitResource = armdeploymentmanager.ServiceUnitResource{
	// 	Name: to.Ptr("myServiceUnit"),
	// 	Type: to.Ptr("Microsoft.DeploymentManager/servicetopologies/services/serviceunits"),
	// 	Location: to.Ptr("centralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
	// 		Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
	// 			ParametersArtifactSourceRelativePath: to.Ptr("parameter/myTopologyUnit.parameters.json"),
	// 			TemplateArtifactSourceRelativePath: to.Ptr("templates/myTopologyUnit.template.json"),
	// 		},
	// 		DeploymentMode: to.Ptr(armdeploymentmanager.DeploymentModeIncremental),
	// 		TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_delete.json
func ExampleServiceUnitsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServiceUnitsClient().Delete(ctx, "myResourceGroup", "myTopology", "myService", "myServiceUnit", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunits_list.json
func ExampleServiceUnitsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceUnitsClient().List(ctx, "myResourceGroup", "myTopology", "myService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceUnitResourceArray = []*armdeploymentmanager.ServiceUnitResource{
	// 	{
	// 		Name: to.Ptr("BackEndServiceUnit"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/servicetopologies/services/serviceunits"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
	// 			Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
	// 				ParametersArtifactSourceRelativePath: to.Ptr("parameter/backend.parameters.json"),
	// 				TemplateArtifactSourceRelativePath: to.Ptr("templates/backend.template.json"),
	// 			},
	// 			DeploymentMode: to.Ptr(armdeploymentmanager.DeploymentModeIncremental),
	// 			TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("FrontEndServiceUnit"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/servicetopologies/services/serviceunits"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
	// 			Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
	// 				ParametersArtifactSourceRelativePath: to.Ptr("parameter/frontend.parameters.json"),
	// 				TemplateArtifactSourceRelativePath: to.Ptr("templates/frontend.template.json"),
	// 			},
	// 			DeploymentMode: to.Ptr(armdeploymentmanager.DeploymentModeComplete),
	// 			TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
	// 		},
	// }}
}
