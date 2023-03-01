//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3035214dea933cd02b1ecfa982c185a572f84b8a/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/PolicyCRUD/ListBackupPolicy.json
func ExampleBackupPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupPoliciesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("000pikumar", "PrivatePreviewVault", nil)
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
		// page.BaseBackupPolicyResourceList = armdataprotection.BaseBackupPolicyResourceList{
		// 	Value: []*armdataprotection.BaseBackupPolicyResource{
		// 		{
		// 			Name: to.Ptr("OSSDBPolicy"),
		// 			Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies"),
		// 			ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PrivatePreviewVault/backupPolicies/OSSDBPolicy"),
		// 			Properties: &armdataprotection.BackupPolicy{
		// 				DatasourceTypes: []*string{
		// 					to.Ptr("OssDB")},
		// 					ObjectType: to.Ptr("BackupPolicy"),
		// 					PolicyRules: []armdataprotection.BasePolicyRuleClassification{
		// 						&armdataprotection.AzureBackupRule{
		// 							Name: to.Ptr("BackupWeekly"),
		// 							ObjectType: to.Ptr("AzureBackupRule"),
		// 							BackupParameters: &armdataprotection.AzureBackupParams{
		// 								ObjectType: to.Ptr("AzureBackupParams"),
		// 								BackupType: to.Ptr("Full"),
		// 							},
		// 							DataStore: &armdataprotection.DataStoreInfoBase{
		// 								DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
		// 								ObjectType: to.Ptr("DataStoreInfoBase"),
		// 							},
		// 							Trigger: &armdataprotection.ScheduleBasedTriggerContext{
		// 								ObjectType: to.Ptr("ScheduleBasedTriggerContext"),
		// 								Schedule: &armdataprotection.BackupSchedule{
		// 									RepeatingTimeIntervals: []*string{
		// 										to.Ptr("R/2019-11-20T08:00:00-08:00/P1W")},
		// 									},
		// 									TaggingCriteria: []*armdataprotection.TaggingCriteria{
		// 										{
		// 											IsDefault: to.Ptr(true),
		// 											TagInfo: &armdataprotection.RetentionTag{
		// 												ID: to.Ptr("Default_"),
		// 												TagName: to.Ptr("Default"),
		// 											},
		// 											TaggingPriority: to.Ptr[int64](99),
		// 										},
		// 										{
		// 											Criteria: []armdataprotection.BackupCriteriaClassification{
		// 												&armdataprotection.ScheduleBasedBackupCriteria{
		// 													ObjectType: to.Ptr("ScheduleBasedBackupCriteria"),
		// 													DaysOfTheWeek: []*armdataprotection.DayOfWeek{
		// 														to.Ptr(armdataprotection.DayOfWeekSunday)},
		// 														ScheduleTimes: []*time.Time{
		// 															to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t}())},
		// 													}},
		// 													IsDefault: to.Ptr(false),
		// 													TagInfo: &armdataprotection.RetentionTag{
		// 														ID: to.Ptr("Weekly_"),
		// 														TagName: to.Ptr("Weekly"),
		// 													},
		// 													TaggingPriority: to.Ptr[int64](20),
		// 											}},
		// 										},
		// 									},
		// 									&armdataprotection.AzureRetentionRule{
		// 										Name: to.Ptr("Default"),
		// 										ObjectType: to.Ptr("AzureRetentionRule"),
		// 										IsDefault: to.Ptr(true),
		// 										Lifecycles: []*armdataprotection.SourceLifeCycle{
		// 											{
		// 												DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
		// 													Duration: to.Ptr("P1W"),
		// 													ObjectType: to.Ptr("AbsoluteDeleteOption"),
		// 												},
		// 												SourceDataStore: &armdataprotection.DataStoreInfoBase{
		// 													DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
		// 													ObjectType: to.Ptr("DataStoreInfoBase"),
		// 												},
		// 										}},
		// 									},
		// 									&armdataprotection.AzureRetentionRule{
		// 										Name: to.Ptr("Weekly"),
		// 										ObjectType: to.Ptr("AzureRetentionRule"),
		// 										IsDefault: to.Ptr(false),
		// 										Lifecycles: []*armdataprotection.SourceLifeCycle{
		// 											{
		// 												DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
		// 													Duration: to.Ptr("P12W"),
		// 													ObjectType: to.Ptr("AbsoluteDeleteOption"),
		// 												},
		// 												SourceDataStore: &armdataprotection.DataStoreInfoBase{
		// 													DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
		// 													ObjectType: to.Ptr("DataStoreInfoBase"),
		// 												},
		// 										}},
		// 								}},
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3035214dea933cd02b1ecfa982c185a572f84b8a/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/PolicyCRUD/GetBackupPolicy.json
func ExampleBackupPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupPoliciesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "000pikumar", "PrivatePreviewVault", "OSSDBPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BaseBackupPolicyResource = armdataprotection.BaseBackupPolicyResource{
	// 	Name: to.Ptr("OSSDBPolicy"),
	// 	Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PrivatePreviewVault/backupPolicies/OSSDBPolicy"),
	// 	Properties: &armdataprotection.BackupPolicy{
	// 		DatasourceTypes: []*string{
	// 			to.Ptr("OssDB")},
	// 			ObjectType: to.Ptr("BackupPolicy"),
	// 			PolicyRules: []armdataprotection.BasePolicyRuleClassification{
	// 				&armdataprotection.AzureBackupRule{
	// 					Name: to.Ptr("BackupWeekly"),
	// 					ObjectType: to.Ptr("AzureBackupRule"),
	// 					BackupParameters: &armdataprotection.AzureBackupParams{
	// 						ObjectType: to.Ptr("AzureBackupParams"),
	// 						BackupType: to.Ptr("Full"),
	// 					},
	// 					DataStore: &armdataprotection.DataStoreInfoBase{
	// 						DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
	// 						ObjectType: to.Ptr("DataStoreInfoBase"),
	// 					},
	// 					Trigger: &armdataprotection.ScheduleBasedTriggerContext{
	// 						ObjectType: to.Ptr("ScheduleBasedTriggerContext"),
	// 						Schedule: &armdataprotection.BackupSchedule{
	// 							RepeatingTimeIntervals: []*string{
	// 								to.Ptr("R/2019-11-20T08:00:00-08:00/P1W")},
	// 							},
	// 							TaggingCriteria: []*armdataprotection.TaggingCriteria{
	// 								{
	// 									IsDefault: to.Ptr(true),
	// 									TagInfo: &armdataprotection.RetentionTag{
	// 										ID: to.Ptr("Default_"),
	// 										TagName: to.Ptr("Default"),
	// 									},
	// 									TaggingPriority: to.Ptr[int64](99),
	// 								},
	// 								{
	// 									Criteria: []armdataprotection.BackupCriteriaClassification{
	// 										&armdataprotection.ScheduleBasedBackupCriteria{
	// 											ObjectType: to.Ptr("ScheduleBasedBackupCriteria"),
	// 											DaysOfTheWeek: []*armdataprotection.DayOfWeek{
	// 												to.Ptr(armdataprotection.DayOfWeekSunday)},
	// 												ScheduleTimes: []*time.Time{
	// 													to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t}())},
	// 											}},
	// 											IsDefault: to.Ptr(false),
	// 											TagInfo: &armdataprotection.RetentionTag{
	// 												ID: to.Ptr("Weekly_"),
	// 												TagName: to.Ptr("Weekly"),
	// 											},
	// 											TaggingPriority: to.Ptr[int64](20),
	// 									}},
	// 								},
	// 							},
	// 							&armdataprotection.AzureRetentionRule{
	// 								Name: to.Ptr("Default"),
	// 								ObjectType: to.Ptr("AzureRetentionRule"),
	// 								IsDefault: to.Ptr(true),
	// 								Lifecycles: []*armdataprotection.SourceLifeCycle{
	// 									{
	// 										DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
	// 											Duration: to.Ptr("P1W"),
	// 											ObjectType: to.Ptr("AbsoluteDeleteOption"),
	// 										},
	// 										SourceDataStore: &armdataprotection.DataStoreInfoBase{
	// 											DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
	// 											ObjectType: to.Ptr("DataStoreInfoBase"),
	// 										},
	// 								}},
	// 							},
	// 							&armdataprotection.AzureRetentionRule{
	// 								Name: to.Ptr("Weekly"),
	// 								ObjectType: to.Ptr("AzureRetentionRule"),
	// 								IsDefault: to.Ptr(false),
	// 								Lifecycles: []*armdataprotection.SourceLifeCycle{
	// 									{
	// 										DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
	// 											Duration: to.Ptr("P12W"),
	// 											ObjectType: to.Ptr("AbsoluteDeleteOption"),
	// 										},
	// 										SourceDataStore: &armdataprotection.DataStoreInfoBase{
	// 											DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
	// 											ObjectType: to.Ptr("DataStoreInfoBase"),
	// 										},
	// 								}},
	// 						}},
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3035214dea933cd02b1ecfa982c185a572f84b8a/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/PolicyCRUD/CreateOrUpdateBackupPolicy.json
func ExampleBackupPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupPoliciesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "000pikumar", "PrivatePreviewVault", "OSSDBPolicy", armdataprotection.BaseBackupPolicyResource{
		Properties: &armdataprotection.BackupPolicy{
			DatasourceTypes: []*string{
				to.Ptr("OssDB")},
			ObjectType: to.Ptr("BackupPolicy"),
			PolicyRules: []armdataprotection.BasePolicyRuleClassification{
				&armdataprotection.AzureBackupRule{
					Name:       to.Ptr("BackupWeekly"),
					ObjectType: to.Ptr("AzureBackupRule"),
					BackupParameters: &armdataprotection.AzureBackupParams{
						ObjectType: to.Ptr("AzureBackupParams"),
						BackupType: to.Ptr("Full"),
					},
					DataStore: &armdataprotection.DataStoreInfoBase{
						DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
						ObjectType:    to.Ptr("DataStoreInfoBase"),
					},
					Trigger: &armdataprotection.ScheduleBasedTriggerContext{
						ObjectType: to.Ptr("ScheduleBasedTriggerContext"),
						Schedule: &armdataprotection.BackupSchedule{
							RepeatingTimeIntervals: []*string{
								to.Ptr("R/2019-11-20T08:00:00-08:00/P1W")},
						},
						TaggingCriteria: []*armdataprotection.TaggingCriteria{
							{
								IsDefault: to.Ptr(true),
								TagInfo: &armdataprotection.RetentionTag{
									TagName: to.Ptr("Default"),
								},
								TaggingPriority: to.Ptr[int64](99),
							},
							{
								Criteria: []armdataprotection.BackupCriteriaClassification{
									&armdataprotection.ScheduleBasedBackupCriteria{
										ObjectType: to.Ptr("ScheduleBasedBackupCriteria"),
										DaysOfTheWeek: []*armdataprotection.DayOfWeek{
											to.Ptr(armdataprotection.DayOfWeekSunday)},
										ScheduleTimes: []*time.Time{
											to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t }())},
									}},
								IsDefault: to.Ptr(false),
								TagInfo: &armdataprotection.RetentionTag{
									TagName: to.Ptr("Weekly"),
								},
								TaggingPriority: to.Ptr[int64](20),
							}},
					},
				},
				&armdataprotection.AzureRetentionRule{
					Name:       to.Ptr("Default"),
					ObjectType: to.Ptr("AzureRetentionRule"),
					IsDefault:  to.Ptr(true),
					Lifecycles: []*armdataprotection.SourceLifeCycle{
						{
							DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
								Duration:   to.Ptr("P1W"),
								ObjectType: to.Ptr("AbsoluteDeleteOption"),
							},
							SourceDataStore: &armdataprotection.DataStoreInfoBase{
								DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
								ObjectType:    to.Ptr("DataStoreInfoBase"),
							},
						}},
				},
				&armdataprotection.AzureRetentionRule{
					Name:       to.Ptr("Weekly"),
					ObjectType: to.Ptr("AzureRetentionRule"),
					IsDefault:  to.Ptr(false),
					Lifecycles: []*armdataprotection.SourceLifeCycle{
						{
							DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
								Duration:   to.Ptr("P12W"),
								ObjectType: to.Ptr("AbsoluteDeleteOption"),
							},
							SourceDataStore: &armdataprotection.DataStoreInfoBase{
								DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
								ObjectType:    to.Ptr("DataStoreInfoBase"),
							},
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BaseBackupPolicyResource = armdataprotection.BaseBackupPolicyResource{
	// 	Name: to.Ptr("OSSDBPolicy"),
	// 	Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PrivatePreviewVault/backupPolicies/OSSDBPolicy"),
	// 	Properties: &armdataprotection.BackupPolicy{
	// 		DatasourceTypes: []*string{
	// 			to.Ptr("OssDB")},
	// 			ObjectType: to.Ptr("BackupPolicy"),
	// 			PolicyRules: []armdataprotection.BasePolicyRuleClassification{
	// 				&armdataprotection.AzureBackupRule{
	// 					Name: to.Ptr("BackupWeekly"),
	// 					ObjectType: to.Ptr("AzureBackupRule"),
	// 					BackupParameters: &armdataprotection.AzureBackupParams{
	// 						ObjectType: to.Ptr("AzureBackupParams"),
	// 						BackupType: to.Ptr("Full"),
	// 					},
	// 					DataStore: &armdataprotection.DataStoreInfoBase{
	// 						DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
	// 						ObjectType: to.Ptr("DataStoreInfoBase"),
	// 					},
	// 					Trigger: &armdataprotection.ScheduleBasedTriggerContext{
	// 						ObjectType: to.Ptr("ScheduleBasedTriggerContext"),
	// 						Schedule: &armdataprotection.BackupSchedule{
	// 							RepeatingTimeIntervals: []*string{
	// 								to.Ptr("R/2019-11-20T08:00:00-08:00/P1W")},
	// 							},
	// 							TaggingCriteria: []*armdataprotection.TaggingCriteria{
	// 								{
	// 									IsDefault: to.Ptr(true),
	// 									TagInfo: &armdataprotection.RetentionTag{
	// 										ID: to.Ptr("Default_"),
	// 										TagName: to.Ptr("Default"),
	// 									},
	// 									TaggingPriority: to.Ptr[int64](99),
	// 								},
	// 								{
	// 									Criteria: []armdataprotection.BackupCriteriaClassification{
	// 										&armdataprotection.ScheduleBasedBackupCriteria{
	// 											ObjectType: to.Ptr("ScheduleBasedBackupCriteria"),
	// 											DaysOfTheWeek: []*armdataprotection.DayOfWeek{
	// 												to.Ptr(armdataprotection.DayOfWeekSunday)},
	// 												ScheduleTimes: []*time.Time{
	// 													to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00Z"); return t}())},
	// 											}},
	// 											IsDefault: to.Ptr(false),
	// 											TagInfo: &armdataprotection.RetentionTag{
	// 												ID: to.Ptr("Weekly_"),
	// 												TagName: to.Ptr("Weekly"),
	// 											},
	// 											TaggingPriority: to.Ptr[int64](20),
	// 									}},
	// 								},
	// 							},
	// 							&armdataprotection.AzureRetentionRule{
	// 								Name: to.Ptr("Default"),
	// 								ObjectType: to.Ptr("AzureRetentionRule"),
	// 								IsDefault: to.Ptr(true),
	// 								Lifecycles: []*armdataprotection.SourceLifeCycle{
	// 									{
	// 										DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
	// 											Duration: to.Ptr("P1W"),
	// 											ObjectType: to.Ptr("AbsoluteDeleteOption"),
	// 										},
	// 										SourceDataStore: &armdataprotection.DataStoreInfoBase{
	// 											DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
	// 											ObjectType: to.Ptr("DataStoreInfoBase"),
	// 										},
	// 								}},
	// 							},
	// 							&armdataprotection.AzureRetentionRule{
	// 								Name: to.Ptr("Weekly"),
	// 								ObjectType: to.Ptr("AzureRetentionRule"),
	// 								IsDefault: to.Ptr(false),
	// 								Lifecycles: []*armdataprotection.SourceLifeCycle{
	// 									{
	// 										DeleteAfter: &armdataprotection.AbsoluteDeleteOption{
	// 											Duration: to.Ptr("P12W"),
	// 											ObjectType: to.Ptr("AbsoluteDeleteOption"),
	// 										},
	// 										SourceDataStore: &armdataprotection.DataStoreInfoBase{
	// 											DataStoreType: to.Ptr(armdataprotection.DataStoreTypesVaultStore),
	// 											ObjectType: to.Ptr("DataStoreInfoBase"),
	// 										},
	// 								}},
	// 						}},
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3035214dea933cd02b1ecfa982c185a572f84b8a/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/PolicyCRUD/DeleteBackupPolicy.json
func ExampleBackupPoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupPoliciesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "000pikumar", "PrivatePreviewVault", "OSSDBPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
