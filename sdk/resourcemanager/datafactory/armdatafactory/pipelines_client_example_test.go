//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_ListByFactory.json
func ExamplePipelinesClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.PipelineListResponse = armdatafactory.PipelineListResponse{
		// 	Value: []*armdatafactory.PipelineResource{
		// 		{
		// 			Name: to.Ptr("examplePipeline"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/pipelines"),
		// 			Etag: to.Ptr("0a006cd4-0000-0000-0000-5b245bd60000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/pipelines/examplePipeline"),
		// 			Properties: &armdatafactory.Pipeline{
		// 				Description: to.Ptr("Example description"),
		// 				Activities: []armdatafactory.ActivityClassification{
		// 					&armdatafactory.ForEachActivity{
		// 						Name: to.Ptr("ExampleForeachActivity"),
		// 						Type: to.Ptr("ForEach"),
		// 						TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
		// 							Activities: []armdatafactory.ActivityClassification{
		// 								&armdatafactory.CopyActivity{
		// 									Name: to.Ptr("ExampleCopyActivity"),
		// 									Type: to.Ptr("Copy"),
		// 									Inputs: []*armdatafactory.DatasetReference{
		// 										{
		// 											Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 											Parameters: map[string]any{
		// 												"MyFileName": "examplecontainer.csv",
		// 												"MyFolderPath": "examplecontainer",
		// 											},
		// 											ReferenceName: to.Ptr("exampleDataset"),
		// 									}},
		// 									Outputs: []*armdatafactory.DatasetReference{
		// 										{
		// 											Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 											Parameters: map[string]any{
		// 												"MyFileName": map[string]any{
		// 													"type": "Expression",
		// 													"value": "@item()",
		// 												},
		// 												"MyFolderPath": "examplecontainer",
		// 											},
		// 											ReferenceName: to.Ptr("exampleDataset"),
		// 									}},
		// 									TypeProperties: &armdatafactory.CopyActivityTypeProperties{
		// 										DataIntegrationUnits: float64(32),
		// 										Sink: &armdatafactory.BlobSink{
		// 											Type: to.Ptr("BlobSink"),
		// 										},
		// 										Source: &armdatafactory.BlobSource{
		// 											Type: to.Ptr("BlobSource"),
		// 										},
		// 									},
		// 							}},
		// 							IsSequential: to.Ptr(true),
		// 							Items: &armdatafactory.Expression{
		// 								Type: to.Ptr(armdatafactory.ExpressionTypeExpression),
		// 								Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
		// 							},
		// 						},
		// 				}},
		// 				Parameters: map[string]*armdatafactory.ParameterSpecification{
		// 					"OutputBlobNameList": &armdatafactory.ParameterSpecification{
		// 						Type: to.Ptr(armdatafactory.ParameterTypeArray),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
func ExamplePipelinesClient_CreateOrUpdate_pipelinesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", armdatafactory.PipelineResource{
		Properties: &armdatafactory.Pipeline{
			Activities: []armdatafactory.ActivityClassification{
				&armdatafactory.ForEachActivity{
					Name: to.Ptr("ExampleForeachActivity"),
					Type: to.Ptr("ForEach"),
					TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
						Activities: []armdatafactory.ActivityClassification{
							&armdatafactory.CopyActivity{
								Name: to.Ptr("ExampleCopyActivity"),
								Type: to.Ptr("Copy"),
								Inputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]any{
											"MyFileName":   "examplecontainer.csv",
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								Outputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]any{
											"MyFileName": map[string]any{
												"type":  "Expression",
												"value": "@item()",
											},
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								TypeProperties: &armdatafactory.CopyActivityTypeProperties{
									DataIntegrationUnits: float64(32),
									Sink: &armdatafactory.BlobSink{
										Type: to.Ptr("BlobSink"),
									},
									Source: &armdatafactory.BlobSource{
										Type: to.Ptr("BlobSource"),
									},
								},
							}},
						IsSequential: to.Ptr(true),
						Items: &armdatafactory.Expression{
							Type:  to.Ptr(armdatafactory.ExpressionTypeExpression),
							Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
						},
					},
				}},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"JobId": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
				"OutputBlobNameList": {
					Type: to.Ptr(armdatafactory.ParameterTypeArray),
				},
			},
			Policy: &armdatafactory.PipelinePolicy{
				ElapsedTimeMetric: &armdatafactory.PipelineElapsedTimeMetricPolicy{
					Duration: "0.00:10:00",
				},
			},
			RunDimensions: map[string]any{
				"JobId": map[string]any{
					"type":  "Expression",
					"value": "@pipeline().parameters.JobId",
				},
			},
			Variables: map[string]*armdatafactory.VariableSpecification{
				"TestVariableArray": {
					Type: to.Ptr(armdatafactory.VariableTypeArray),
				},
			},
		},
	}, &armdatafactory.PipelinesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineResource = armdatafactory.PipelineResource{
	// 	Name: to.Ptr("examplePipeline"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/pipelines"),
	// 	Etag: to.Ptr("0a0069d4-0000-0000-0000-5b245bd50000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/pipelines/examplePipeline"),
	// 	Properties: &armdatafactory.Pipeline{
	// 		Activities: []armdatafactory.ActivityClassification{
	// 			&armdatafactory.ForEachActivity{
	// 				Name: to.Ptr("ExampleForeachActivity"),
	// 				Type: to.Ptr("ForEach"),
	// 				TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
	// 					Activities: []armdatafactory.ActivityClassification{
	// 						&armdatafactory.CopyActivity{
	// 							Name: to.Ptr("ExampleCopyActivity"),
	// 							Type: to.Ptr("Copy"),
	// 							Inputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": "examplecontainer.csv",
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							Outputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": map[string]any{
	// 											"type": "Expression",
	// 											"value": "@item()",
	// 										},
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							TypeProperties: &armdatafactory.CopyActivityTypeProperties{
	// 								DataIntegrationUnits: float64(32),
	// 								Sink: &armdatafactory.BlobSink{
	// 									Type: to.Ptr("BlobSink"),
	// 								},
	// 								Source: &armdatafactory.BlobSource{
	// 									Type: to.Ptr("BlobSource"),
	// 								},
	// 							},
	// 					}},
	// 					IsSequential: to.Ptr(true),
	// 					Items: &armdatafactory.Expression{
	// 						Type: to.Ptr(armdatafactory.ExpressionTypeExpression),
	// 						Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
	// 					},
	// 				},
	// 		}},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"JobId": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeString),
	// 			},
	// 			"OutputBlobNameList": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeArray),
	// 			},
	// 		},
	// 		RunDimensions: map[string]any{
	// 			"JobId": map[string]any{
	// 				"type": "Expression",
	// 				"value": "@pipeline().parameters.JobId",
	// 			},
	// 		},
	// 		Variables: map[string]*armdatafactory.VariableSpecification{
	// 			"TestVariableArray": &armdatafactory.VariableSpecification{
	// 				Type: to.Ptr(armdatafactory.VariableTypeArray),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Update.json
func ExamplePipelinesClient_CreateOrUpdate_pipelinesUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", armdatafactory.PipelineResource{
		Properties: &armdatafactory.Pipeline{
			Description: to.Ptr("Example description"),
			Activities: []armdatafactory.ActivityClassification{
				&armdatafactory.ForEachActivity{
					Name: to.Ptr("ExampleForeachActivity"),
					Type: to.Ptr("ForEach"),
					TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
						Activities: []armdatafactory.ActivityClassification{
							&armdatafactory.CopyActivity{
								Name: to.Ptr("ExampleCopyActivity"),
								Type: to.Ptr("Copy"),
								Inputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]any{
											"MyFileName":   "examplecontainer.csv",
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								Outputs: []*armdatafactory.DatasetReference{
									{
										Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
										Parameters: map[string]any{
											"MyFileName": map[string]any{
												"type":  "Expression",
												"value": "@item()",
											},
											"MyFolderPath": "examplecontainer",
										},
										ReferenceName: to.Ptr("exampleDataset"),
									}},
								TypeProperties: &armdatafactory.CopyActivityTypeProperties{
									DataIntegrationUnits: float64(32),
									Sink: &armdatafactory.BlobSink{
										Type: to.Ptr("BlobSink"),
									},
									Source: &armdatafactory.BlobSource{
										Type: to.Ptr("BlobSource"),
									},
								},
							}},
						IsSequential: to.Ptr(true),
						Items: &armdatafactory.Expression{
							Type:  to.Ptr(armdatafactory.ExpressionTypeExpression),
							Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
						},
					},
				}},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"OutputBlobNameList": {
					Type: to.Ptr(armdatafactory.ParameterTypeArray),
				},
			},
			Policy: &armdatafactory.PipelinePolicy{
				ElapsedTimeMetric: &armdatafactory.PipelineElapsedTimeMetricPolicy{
					Duration: "0.00:10:00",
				},
			},
		},
	}, &armdatafactory.PipelinesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineResource = armdatafactory.PipelineResource{
	// 	Name: to.Ptr("examplePipeline"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/pipelines"),
	// 	Etag: to.Ptr("0a006cd4-0000-0000-0000-5b245bd60000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/pipelines/examplePipeline"),
	// 	Properties: &armdatafactory.Pipeline{
	// 		Description: to.Ptr("Example description"),
	// 		Activities: []armdatafactory.ActivityClassification{
	// 			&armdatafactory.ForEachActivity{
	// 				Name: to.Ptr("ExampleForeachActivity"),
	// 				Type: to.Ptr("ForEach"),
	// 				TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
	// 					Activities: []armdatafactory.ActivityClassification{
	// 						&armdatafactory.CopyActivity{
	// 							Name: to.Ptr("ExampleCopyActivity"),
	// 							Type: to.Ptr("Copy"),
	// 							Inputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": "examplecontainer.csv",
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							Outputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": map[string]any{
	// 											"type": "Expression",
	// 											"value": "@item()",
	// 										},
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							TypeProperties: &armdatafactory.CopyActivityTypeProperties{
	// 								DataIntegrationUnits: float64(32),
	// 								Sink: &armdatafactory.BlobSink{
	// 									Type: to.Ptr("BlobSink"),
	// 								},
	// 								Source: &armdatafactory.BlobSource{
	// 									Type: to.Ptr("BlobSource"),
	// 								},
	// 							},
	// 					}},
	// 					IsSequential: to.Ptr(true),
	// 					Items: &armdatafactory.Expression{
	// 						Type: to.Ptr(armdatafactory.ExpressionTypeExpression),
	// 						Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
	// 					},
	// 				},
	// 		}},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"OutputBlobNameList": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeArray),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Get.json
func ExamplePipelinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", &armdatafactory.PipelinesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PipelineResource = armdatafactory.PipelineResource{
	// 	Name: to.Ptr("examplePipeline"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/pipelines"),
	// 	Etag: to.Ptr("1500504f-0000-0200-0000-5cbe090f0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/pipelines/examplePipeline"),
	// 	Properties: &armdatafactory.Pipeline{
	// 		Description: to.Ptr("Example description"),
	// 		Activities: []armdatafactory.ActivityClassification{
	// 			&armdatafactory.ForEachActivity{
	// 				Name: to.Ptr("ExampleForeachActivity"),
	// 				Type: to.Ptr("ForEach"),
	// 				TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
	// 					Activities: []armdatafactory.ActivityClassification{
	// 						&armdatafactory.CopyActivity{
	// 							Name: to.Ptr("ExampleCopyActivity"),
	// 							Type: to.Ptr("Copy"),
	// 							Inputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": "examplecontainer.csv",
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							Outputs: []*armdatafactory.DatasetReference{
	// 								{
	// 									Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 									Parameters: map[string]any{
	// 										"MyFileName": map[string]any{
	// 											"type": "Expression",
	// 											"value": "@item()",
	// 										},
	// 										"MyFolderPath": "examplecontainer",
	// 									},
	// 									ReferenceName: to.Ptr("exampleDataset"),
	// 							}},
	// 							TypeProperties: &armdatafactory.CopyActivityTypeProperties{
	// 								DataIntegrationUnits: float64(32),
	// 								Sink: &armdatafactory.BlobSink{
	// 									Type: to.Ptr("BlobSink"),
	// 								},
	// 								Source: &armdatafactory.BlobSource{
	// 									Type: to.Ptr("BlobSource"),
	// 								},
	// 							},
	// 					}},
	// 					IsSequential: to.Ptr(true),
	// 					Items: &armdatafactory.Expression{
	// 						Type: to.Ptr(armdatafactory.ExpressionTypeExpression),
	// 						Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
	// 					},
	// 				},
	// 		}},
	// 		Parameters: map[string]*armdatafactory.ParameterSpecification{
	// 			"OutputBlobNameList": &armdatafactory.ParameterSpecification{
	// 				Type: to.Ptr(armdatafactory.ParameterTypeArray),
	// 			},
	// 		},
	// 		Policy: &armdatafactory.PipelinePolicy{
	// 			ElapsedTimeMetric: &armdatafactory.PipelineElapsedTimeMetricPolicy{
	// 				Duration: "0.00:10:00",
	// 			},
	// 		},
	// 		Variables: map[string]*armdatafactory.VariableSpecification{
	// 			"TestVariableArray": &armdatafactory.VariableSpecification{
	// 				Type: to.Ptr(armdatafactory.VariableTypeArray),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Delete.json
func ExamplePipelinesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_CreateRun.json
func ExamplePipelinesClient_CreateRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateRun(ctx, "exampleResourceGroup", "exampleFactoryName", "examplePipeline", &armdatafactory.PipelinesClientCreateRunOptions{ReferencePipelineRunID: nil,
		IsRecovery:        nil,
		StartActivityName: nil,
		StartFromFailure:  nil,
		Parameters: map[string]any{
			"OutputBlobNameList": []any{
				"exampleoutput.csv",
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CreateRunResponse = armdatafactory.CreateRunResponse{
	// 	RunID: to.Ptr("2f7fdb90-5df1-4b8e-ac2f-064cfa58202b"),
	// }
}
