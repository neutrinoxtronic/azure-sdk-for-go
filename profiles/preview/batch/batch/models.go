// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package batch

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/batch/2020-03-01.11.0/batch"
)

type AccessScope = original.AccessScope

const (
	Job AccessScope = original.Job
)

type AllocationState = original.AllocationState

const (
	Resizing AllocationState = original.Resizing
	Steady   AllocationState = original.Steady
	Stopping AllocationState = original.Stopping
)

type AutoUserScope = original.AutoUserScope

const (
	Pool AutoUserScope = original.Pool
	Task AutoUserScope = original.Task
)

type CachingType = original.CachingType

const (
	None      CachingType = original.None
	ReadOnly  CachingType = original.ReadOnly
	ReadWrite CachingType = original.ReadWrite
)

type CertificateFormat = original.CertificateFormat

const (
	Cer CertificateFormat = original.Cer
	Pfx CertificateFormat = original.Pfx
)

type CertificateState = original.CertificateState

const (
	Active       CertificateState = original.Active
	DeleteFailed CertificateState = original.DeleteFailed
	Deleting     CertificateState = original.Deleting
)

type CertificateStoreLocation = original.CertificateStoreLocation

const (
	CurrentUser  CertificateStoreLocation = original.CurrentUser
	LocalMachine CertificateStoreLocation = original.LocalMachine
)

type CertificateVisibility = original.CertificateVisibility

const (
	CertificateVisibilityRemoteUser CertificateVisibility = original.CertificateVisibilityRemoteUser
	CertificateVisibilityStartTask  CertificateVisibility = original.CertificateVisibilityStartTask
	CertificateVisibilityTask       CertificateVisibility = original.CertificateVisibilityTask
)

type ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOption

const (
	Requeue        ComputeNodeDeallocationOption = original.Requeue
	RetainedData   ComputeNodeDeallocationOption = original.RetainedData
	TaskCompletion ComputeNodeDeallocationOption = original.TaskCompletion
	Terminate      ComputeNodeDeallocationOption = original.Terminate
)

type ComputeNodeFillType = original.ComputeNodeFillType

const (
	Pack   ComputeNodeFillType = original.Pack
	Spread ComputeNodeFillType = original.Spread
)

type ComputeNodeRebootOption = original.ComputeNodeRebootOption

const (
	ComputeNodeRebootOptionRequeue        ComputeNodeRebootOption = original.ComputeNodeRebootOptionRequeue
	ComputeNodeRebootOptionRetainedData   ComputeNodeRebootOption = original.ComputeNodeRebootOptionRetainedData
	ComputeNodeRebootOptionTaskCompletion ComputeNodeRebootOption = original.ComputeNodeRebootOptionTaskCompletion
	ComputeNodeRebootOptionTerminate      ComputeNodeRebootOption = original.ComputeNodeRebootOptionTerminate
)

type ComputeNodeReimageOption = original.ComputeNodeReimageOption

const (
	ComputeNodeReimageOptionRequeue        ComputeNodeReimageOption = original.ComputeNodeReimageOptionRequeue
	ComputeNodeReimageOptionRetainedData   ComputeNodeReimageOption = original.ComputeNodeReimageOptionRetainedData
	ComputeNodeReimageOptionTaskCompletion ComputeNodeReimageOption = original.ComputeNodeReimageOptionTaskCompletion
	ComputeNodeReimageOptionTerminate      ComputeNodeReimageOption = original.ComputeNodeReimageOptionTerminate
)

type ComputeNodeState = original.ComputeNodeState

const (
	Creating            ComputeNodeState = original.Creating
	Idle                ComputeNodeState = original.Idle
	LeavingPool         ComputeNodeState = original.LeavingPool
	Offline             ComputeNodeState = original.Offline
	Preempted           ComputeNodeState = original.Preempted
	Rebooting           ComputeNodeState = original.Rebooting
	Reimaging           ComputeNodeState = original.Reimaging
	Running             ComputeNodeState = original.Running
	Starting            ComputeNodeState = original.Starting
	StartTaskFailed     ComputeNodeState = original.StartTaskFailed
	Unknown             ComputeNodeState = original.Unknown
	Unusable            ComputeNodeState = original.Unusable
	WaitingForStartTask ComputeNodeState = original.WaitingForStartTask
)

type ContainerWorkingDirectory = original.ContainerWorkingDirectory

const (
	ContainerImageDefault ContainerWorkingDirectory = original.ContainerImageDefault
	TaskWorkingDirectory  ContainerWorkingDirectory = original.TaskWorkingDirectory
)

type DependencyAction = original.DependencyAction

const (
	Block   DependencyAction = original.Block
	Satisfy DependencyAction = original.Satisfy
)

type DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOption

const (
	DisableComputeNodeSchedulingOptionRequeue        DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOptionRequeue
	DisableComputeNodeSchedulingOptionTaskCompletion DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOptionTaskCompletion
	DisableComputeNodeSchedulingOptionTerminate      DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOptionTerminate
)

type DisableJobOption = original.DisableJobOption

const (
	DisableJobOptionRequeue   DisableJobOption = original.DisableJobOptionRequeue
	DisableJobOptionTerminate DisableJobOption = original.DisableJobOptionTerminate
	DisableJobOptionWait      DisableJobOption = original.DisableJobOptionWait
)

type DiskEncryptionTarget = original.DiskEncryptionTarget

const (
	OsDisk        DiskEncryptionTarget = original.OsDisk
	TemporaryDisk DiskEncryptionTarget = original.TemporaryDisk
)

type DynamicVNetAssignmentScope = original.DynamicVNetAssignmentScope

const (
	DynamicVNetAssignmentScopeJob  DynamicVNetAssignmentScope = original.DynamicVNetAssignmentScopeJob
	DynamicVNetAssignmentScopeNone DynamicVNetAssignmentScope = original.DynamicVNetAssignmentScopeNone
)

type ElevationLevel = original.ElevationLevel

const (
	Admin    ElevationLevel = original.Admin
	NonAdmin ElevationLevel = original.NonAdmin
)

type ErrorCategory = original.ErrorCategory

const (
	ServerError ErrorCategory = original.ServerError
	UserError   ErrorCategory = original.UserError
)

type IPAddressProvisioningType = original.IPAddressProvisioningType

const (
	BatchManaged        IPAddressProvisioningType = original.BatchManaged
	NoPublicIPAddresses IPAddressProvisioningType = original.NoPublicIPAddresses
	UserManaged         IPAddressProvisioningType = original.UserManaged
)

type InboundEndpointProtocol = original.InboundEndpointProtocol

const (
	TCP InboundEndpointProtocol = original.TCP
	UDP InboundEndpointProtocol = original.UDP
)

type JobAction = original.JobAction

const (
	JobActionDisable   JobAction = original.JobActionDisable
	JobActionNone      JobAction = original.JobActionNone
	JobActionTerminate JobAction = original.JobActionTerminate
)

type JobPreparationTaskState = original.JobPreparationTaskState

const (
	JobPreparationTaskStateCompleted JobPreparationTaskState = original.JobPreparationTaskStateCompleted
	JobPreparationTaskStateRunning   JobPreparationTaskState = original.JobPreparationTaskStateRunning
)

type JobReleaseTaskState = original.JobReleaseTaskState

const (
	JobReleaseTaskStateCompleted JobReleaseTaskState = original.JobReleaseTaskStateCompleted
	JobReleaseTaskStateRunning   JobReleaseTaskState = original.JobReleaseTaskStateRunning
)

type JobScheduleState = original.JobScheduleState

const (
	JobScheduleStateActive      JobScheduleState = original.JobScheduleStateActive
	JobScheduleStateCompleted   JobScheduleState = original.JobScheduleStateCompleted
	JobScheduleStateDeleting    JobScheduleState = original.JobScheduleStateDeleting
	JobScheduleStateDisabled    JobScheduleState = original.JobScheduleStateDisabled
	JobScheduleStateTerminating JobScheduleState = original.JobScheduleStateTerminating
)

type JobState = original.JobState

const (
	JobStateActive      JobState = original.JobStateActive
	JobStateCompleted   JobState = original.JobStateCompleted
	JobStateDeleting    JobState = original.JobStateDeleting
	JobStateDisabled    JobState = original.JobStateDisabled
	JobStateDisabling   JobState = original.JobStateDisabling
	JobStateEnabling    JobState = original.JobStateEnabling
	JobStateTerminating JobState = original.JobStateTerminating
)

type LoginMode = original.LoginMode

const (
	Batch       LoginMode = original.Batch
	Interactive LoginMode = original.Interactive
)

type NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccess

const (
	Allow NetworkSecurityGroupRuleAccess = original.Allow
	Deny  NetworkSecurityGroupRuleAccess = original.Deny
)

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

type OnAllTasksComplete = original.OnAllTasksComplete

const (
	NoAction     OnAllTasksComplete = original.NoAction
	TerminateJob OnAllTasksComplete = original.TerminateJob
)

type OnTaskFailure = original.OnTaskFailure

const (
	OnTaskFailureNoAction                    OnTaskFailure = original.OnTaskFailureNoAction
	OnTaskFailurePerformExitOptionsJobAction OnTaskFailure = original.OnTaskFailurePerformExitOptionsJobAction
)

type OutputFileUploadCondition = original.OutputFileUploadCondition

const (
	OutputFileUploadConditionTaskCompletion OutputFileUploadCondition = original.OutputFileUploadConditionTaskCompletion
	OutputFileUploadConditionTaskFailure    OutputFileUploadCondition = original.OutputFileUploadConditionTaskFailure
	OutputFileUploadConditionTaskSuccess    OutputFileUploadCondition = original.OutputFileUploadConditionTaskSuccess
)

type PoolLifetimeOption = original.PoolLifetimeOption

const (
	PoolLifetimeOptionJob         PoolLifetimeOption = original.PoolLifetimeOptionJob
	PoolLifetimeOptionJobSchedule PoolLifetimeOption = original.PoolLifetimeOptionJobSchedule
)

type PoolState = original.PoolState

const (
	PoolStateActive   PoolState = original.PoolStateActive
	PoolStateDeleting PoolState = original.PoolStateDeleting
)

type SchedulingState = original.SchedulingState

const (
	Disabled SchedulingState = original.Disabled
	Enabled  SchedulingState = original.Enabled
)

type StartTaskState = original.StartTaskState

const (
	StartTaskStateCompleted StartTaskState = original.StartTaskStateCompleted
	StartTaskStateRunning   StartTaskState = original.StartTaskStateRunning
)

type StorageAccountType = original.StorageAccountType

const (
	PremiumLRS  StorageAccountType = original.PremiumLRS
	StandardLRS StorageAccountType = original.StandardLRS
)

type SubtaskState = original.SubtaskState

const (
	SubtaskStateCompleted SubtaskState = original.SubtaskStateCompleted
	SubtaskStatePreparing SubtaskState = original.SubtaskStatePreparing
	SubtaskStateRunning   SubtaskState = original.SubtaskStateRunning
)

type TaskAddStatus = original.TaskAddStatus

const (
	TaskAddStatusClientError TaskAddStatus = original.TaskAddStatusClientError
	TaskAddStatusServerError TaskAddStatus = original.TaskAddStatusServerError
	TaskAddStatusSuccess     TaskAddStatus = original.TaskAddStatusSuccess
)

type TaskExecutionResult = original.TaskExecutionResult

const (
	Failure TaskExecutionResult = original.Failure
	Success TaskExecutionResult = original.Success
)

type TaskState = original.TaskState

const (
	TaskStateActive    TaskState = original.TaskStateActive
	TaskStateCompleted TaskState = original.TaskStateCompleted
	TaskStatePreparing TaskState = original.TaskStatePreparing
	TaskStateRunning   TaskState = original.TaskStateRunning
)

type VerificationType = original.VerificationType

const (
	Unverified VerificationType = original.Unverified
	Verified   VerificationType = original.Verified
)

type AccountClient = original.AccountClient
type AccountListSupportedImagesResult = original.AccountListSupportedImagesResult
type AccountListSupportedImagesResultIterator = original.AccountListSupportedImagesResultIterator
type AccountListSupportedImagesResultPage = original.AccountListSupportedImagesResultPage
type AffinityInformation = original.AffinityInformation
type ApplicationClient = original.ApplicationClient
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationPackageReference = original.ApplicationPackageReference
type ApplicationSummary = original.ApplicationSummary
type AuthenticationTokenSettings = original.AuthenticationTokenSettings
type AutoPoolSpecification = original.AutoPoolSpecification
type AutoScaleRun = original.AutoScaleRun
type AutoScaleRunError = original.AutoScaleRunError
type AutoUserSpecification = original.AutoUserSpecification
type AzureBlobFileSystemConfiguration = original.AzureBlobFileSystemConfiguration
type AzureFileShareConfiguration = original.AzureFileShareConfiguration
type BaseClient = original.BaseClient
type CIFSMountConfiguration = original.CIFSMountConfiguration
type Certificate = original.Certificate
type CertificateAddParameter = original.CertificateAddParameter
type CertificateClient = original.CertificateClient
type CertificateListResult = original.CertificateListResult
type CertificateListResultIterator = original.CertificateListResultIterator
type CertificateListResultPage = original.CertificateListResultPage
type CertificateReference = original.CertificateReference
type CloudJob = original.CloudJob
type CloudJobListPreparationAndReleaseTaskStatusResult = original.CloudJobListPreparationAndReleaseTaskStatusResult
type CloudJobListPreparationAndReleaseTaskStatusResultIterator = original.CloudJobListPreparationAndReleaseTaskStatusResultIterator
type CloudJobListPreparationAndReleaseTaskStatusResultPage = original.CloudJobListPreparationAndReleaseTaskStatusResultPage
type CloudJobListResult = original.CloudJobListResult
type CloudJobListResultIterator = original.CloudJobListResultIterator
type CloudJobListResultPage = original.CloudJobListResultPage
type CloudJobSchedule = original.CloudJobSchedule
type CloudJobScheduleListResult = original.CloudJobScheduleListResult
type CloudJobScheduleListResultIterator = original.CloudJobScheduleListResultIterator
type CloudJobScheduleListResultPage = original.CloudJobScheduleListResultPage
type CloudPool = original.CloudPool
type CloudPoolListResult = original.CloudPoolListResult
type CloudPoolListResultIterator = original.CloudPoolListResultIterator
type CloudPoolListResultPage = original.CloudPoolListResultPage
type CloudServiceConfiguration = original.CloudServiceConfiguration
type CloudTask = original.CloudTask
type CloudTaskListResult = original.CloudTaskListResult
type CloudTaskListResultIterator = original.CloudTaskListResultIterator
type CloudTaskListResultPage = original.CloudTaskListResultPage
type CloudTaskListSubtasksResult = original.CloudTaskListSubtasksResult
type ComputeNode = original.ComputeNode
type ComputeNodeClient = original.ComputeNodeClient
type ComputeNodeEndpointConfiguration = original.ComputeNodeEndpointConfiguration
type ComputeNodeError = original.ComputeNodeError
type ComputeNodeGetRemoteLoginSettingsResult = original.ComputeNodeGetRemoteLoginSettingsResult
type ComputeNodeInformation = original.ComputeNodeInformation
type ComputeNodeListResult = original.ComputeNodeListResult
type ComputeNodeListResultIterator = original.ComputeNodeListResultIterator
type ComputeNodeListResultPage = original.ComputeNodeListResultPage
type ComputeNodeUser = original.ComputeNodeUser
type ContainerConfiguration = original.ContainerConfiguration
type ContainerRegistry = original.ContainerRegistry
type DataDisk = original.DataDisk
type DeleteCertificateError = original.DeleteCertificateError
type DiskEncryptionConfiguration = original.DiskEncryptionConfiguration
type EnvironmentSetting = original.EnvironmentSetting
type Error = original.Error
type ErrorDetail = original.ErrorDetail
type ErrorMessage = original.ErrorMessage
type ExitCodeMapping = original.ExitCodeMapping
type ExitCodeRangeMapping = original.ExitCodeRangeMapping
type ExitConditions = original.ExitConditions
type ExitOptions = original.ExitOptions
type FileClient = original.FileClient
type FileProperties = original.FileProperties
type ImageInformation = original.ImageInformation
type ImageReference = original.ImageReference
type InboundEndpoint = original.InboundEndpoint
type InboundNATPool = original.InboundNATPool
type JobAddParameter = original.JobAddParameter
type JobClient = original.JobClient
type JobConstraints = original.JobConstraints
type JobDisableParameter = original.JobDisableParameter
type JobExecutionInformation = original.JobExecutionInformation
type JobManagerTask = original.JobManagerTask
type JobNetworkConfiguration = original.JobNetworkConfiguration
type JobPatchParameter = original.JobPatchParameter
type JobPreparationAndReleaseTaskExecutionInformation = original.JobPreparationAndReleaseTaskExecutionInformation
type JobPreparationTask = original.JobPreparationTask
type JobPreparationTaskExecutionInformation = original.JobPreparationTaskExecutionInformation
type JobReleaseTask = original.JobReleaseTask
type JobReleaseTaskExecutionInformation = original.JobReleaseTaskExecutionInformation
type JobScheduleAddParameter = original.JobScheduleAddParameter
type JobScheduleClient = original.JobScheduleClient
type JobScheduleExecutionInformation = original.JobScheduleExecutionInformation
type JobSchedulePatchParameter = original.JobSchedulePatchParameter
type JobScheduleStatistics = original.JobScheduleStatistics
type JobScheduleUpdateParameter = original.JobScheduleUpdateParameter
type JobSchedulingError = original.JobSchedulingError
type JobSpecification = original.JobSpecification
type JobStatistics = original.JobStatistics
type JobTerminateParameter = original.JobTerminateParameter
type JobUpdateParameter = original.JobUpdateParameter
type LinuxUserConfiguration = original.LinuxUserConfiguration
type MetadataItem = original.MetadataItem
type MountConfiguration = original.MountConfiguration
type MultiInstanceSettings = original.MultiInstanceSettings
type NFSMountConfiguration = original.NFSMountConfiguration
type NameValuePair = original.NameValuePair
type NetworkConfiguration = original.NetworkConfiguration
type NetworkSecurityGroupRule = original.NetworkSecurityGroupRule
type NodeAgentInformation = original.NodeAgentInformation
type NodeCounts = original.NodeCounts
type NodeDisableSchedulingParameter = original.NodeDisableSchedulingParameter
type NodeFile = original.NodeFile
type NodeFileListResult = original.NodeFileListResult
type NodeFileListResultIterator = original.NodeFileListResultIterator
type NodeFileListResultPage = original.NodeFileListResultPage
type NodeRebootParameter = original.NodeRebootParameter
type NodeReimageParameter = original.NodeReimageParameter
type NodeRemoveParameter = original.NodeRemoveParameter
type NodeUpdateUserParameter = original.NodeUpdateUserParameter
type OutputFile = original.OutputFile
type OutputFileBlobContainerDestination = original.OutputFileBlobContainerDestination
type OutputFileDestination = original.OutputFileDestination
type OutputFileUploadOptions = original.OutputFileUploadOptions
type PoolAddParameter = original.PoolAddParameter
type PoolClient = original.PoolClient
type PoolEnableAutoScaleParameter = original.PoolEnableAutoScaleParameter
type PoolEndpointConfiguration = original.PoolEndpointConfiguration
type PoolEvaluateAutoScaleParameter = original.PoolEvaluateAutoScaleParameter
type PoolInformation = original.PoolInformation
type PoolListUsageMetricsResult = original.PoolListUsageMetricsResult
type PoolListUsageMetricsResultIterator = original.PoolListUsageMetricsResultIterator
type PoolListUsageMetricsResultPage = original.PoolListUsageMetricsResultPage
type PoolNodeCounts = original.PoolNodeCounts
type PoolNodeCountsListResult = original.PoolNodeCountsListResult
type PoolNodeCountsListResultIterator = original.PoolNodeCountsListResultIterator
type PoolNodeCountsListResultPage = original.PoolNodeCountsListResultPage
type PoolPatchParameter = original.PoolPatchParameter
type PoolResizeParameter = original.PoolResizeParameter
type PoolSpecification = original.PoolSpecification
type PoolStatistics = original.PoolStatistics
type PoolUpdatePropertiesParameter = original.PoolUpdatePropertiesParameter
type PoolUsageMetrics = original.PoolUsageMetrics
type PublicIPAddressConfiguration = original.PublicIPAddressConfiguration
type ReadCloser = original.ReadCloser
type RecentJob = original.RecentJob
type ResizeError = original.ResizeError
type ResourceFile = original.ResourceFile
type ResourceStatistics = original.ResourceStatistics
type Schedule = original.Schedule
type StartTask = original.StartTask
type StartTaskInformation = original.StartTaskInformation
type SubtaskInformation = original.SubtaskInformation
type TaskAddCollectionParameter = original.TaskAddCollectionParameter
type TaskAddCollectionResult = original.TaskAddCollectionResult
type TaskAddParameter = original.TaskAddParameter
type TaskAddResult = original.TaskAddResult
type TaskClient = original.TaskClient
type TaskConstraints = original.TaskConstraints
type TaskContainerExecutionInformation = original.TaskContainerExecutionInformation
type TaskContainerSettings = original.TaskContainerSettings
type TaskCounts = original.TaskCounts
type TaskDependencies = original.TaskDependencies
type TaskExecutionInformation = original.TaskExecutionInformation
type TaskFailureInformation = original.TaskFailureInformation
type TaskIDRange = original.TaskIDRange
type TaskInformation = original.TaskInformation
type TaskSchedulingPolicy = original.TaskSchedulingPolicy
type TaskStatistics = original.TaskStatistics
type TaskUpdateParameter = original.TaskUpdateParameter
type UploadBatchServiceLogsConfiguration = original.UploadBatchServiceLogsConfiguration
type UploadBatchServiceLogsResult = original.UploadBatchServiceLogsResult
type UsageStatistics = original.UsageStatistics
type UserAccount = original.UserAccount
type UserIdentity = original.UserIdentity
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type WindowsConfiguration = original.WindowsConfiguration
type WindowsUserConfiguration = original.WindowsUserConfiguration

func New(batchURL string) BaseClient {
	return original.New(batchURL)
}
func NewAccountClient(batchURL string) AccountClient {
	return original.NewAccountClient(batchURL)
}
func NewAccountListSupportedImagesResultIterator(page AccountListSupportedImagesResultPage) AccountListSupportedImagesResultIterator {
	return original.NewAccountListSupportedImagesResultIterator(page)
}
func NewAccountListSupportedImagesResultPage(getNextPage func(context.Context, AccountListSupportedImagesResult) (AccountListSupportedImagesResult, error)) AccountListSupportedImagesResultPage {
	return original.NewAccountListSupportedImagesResultPage(getNextPage)
}
func NewApplicationClient(batchURL string) ApplicationClient {
	return original.NewApplicationClient(batchURL)
}
func NewApplicationListResultIterator(page ApplicationListResultPage) ApplicationListResultIterator {
	return original.NewApplicationListResultIterator(page)
}
func NewApplicationListResultPage(getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(getNextPage)
}
func NewCertificateClient(batchURL string) CertificateClient {
	return original.NewCertificateClient(batchURL)
}
func NewCertificateListResultIterator(page CertificateListResultPage) CertificateListResultIterator {
	return original.NewCertificateListResultIterator(page)
}
func NewCertificateListResultPage(getNextPage func(context.Context, CertificateListResult) (CertificateListResult, error)) CertificateListResultPage {
	return original.NewCertificateListResultPage(getNextPage)
}
func NewCloudJobListPreparationAndReleaseTaskStatusResultIterator(page CloudJobListPreparationAndReleaseTaskStatusResultPage) CloudJobListPreparationAndReleaseTaskStatusResultIterator {
	return original.NewCloudJobListPreparationAndReleaseTaskStatusResultIterator(page)
}
func NewCloudJobListPreparationAndReleaseTaskStatusResultPage(getNextPage func(context.Context, CloudJobListPreparationAndReleaseTaskStatusResult) (CloudJobListPreparationAndReleaseTaskStatusResult, error)) CloudJobListPreparationAndReleaseTaskStatusResultPage {
	return original.NewCloudJobListPreparationAndReleaseTaskStatusResultPage(getNextPage)
}
func NewCloudJobListResultIterator(page CloudJobListResultPage) CloudJobListResultIterator {
	return original.NewCloudJobListResultIterator(page)
}
func NewCloudJobListResultPage(getNextPage func(context.Context, CloudJobListResult) (CloudJobListResult, error)) CloudJobListResultPage {
	return original.NewCloudJobListResultPage(getNextPage)
}
func NewCloudJobScheduleListResultIterator(page CloudJobScheduleListResultPage) CloudJobScheduleListResultIterator {
	return original.NewCloudJobScheduleListResultIterator(page)
}
func NewCloudJobScheduleListResultPage(getNextPage func(context.Context, CloudJobScheduleListResult) (CloudJobScheduleListResult, error)) CloudJobScheduleListResultPage {
	return original.NewCloudJobScheduleListResultPage(getNextPage)
}
func NewCloudPoolListResultIterator(page CloudPoolListResultPage) CloudPoolListResultIterator {
	return original.NewCloudPoolListResultIterator(page)
}
func NewCloudPoolListResultPage(getNextPage func(context.Context, CloudPoolListResult) (CloudPoolListResult, error)) CloudPoolListResultPage {
	return original.NewCloudPoolListResultPage(getNextPage)
}
func NewCloudTaskListResultIterator(page CloudTaskListResultPage) CloudTaskListResultIterator {
	return original.NewCloudTaskListResultIterator(page)
}
func NewCloudTaskListResultPage(getNextPage func(context.Context, CloudTaskListResult) (CloudTaskListResult, error)) CloudTaskListResultPage {
	return original.NewCloudTaskListResultPage(getNextPage)
}
func NewComputeNodeClient(batchURL string) ComputeNodeClient {
	return original.NewComputeNodeClient(batchURL)
}
func NewComputeNodeListResultIterator(page ComputeNodeListResultPage) ComputeNodeListResultIterator {
	return original.NewComputeNodeListResultIterator(page)
}
func NewComputeNodeListResultPage(getNextPage func(context.Context, ComputeNodeListResult) (ComputeNodeListResult, error)) ComputeNodeListResultPage {
	return original.NewComputeNodeListResultPage(getNextPage)
}
func NewFileClient(batchURL string) FileClient {
	return original.NewFileClient(batchURL)
}
func NewJobClient(batchURL string) JobClient {
	return original.NewJobClient(batchURL)
}
func NewJobScheduleClient(batchURL string) JobScheduleClient {
	return original.NewJobScheduleClient(batchURL)
}
func NewNodeFileListResultIterator(page NodeFileListResultPage) NodeFileListResultIterator {
	return original.NewNodeFileListResultIterator(page)
}
func NewNodeFileListResultPage(getNextPage func(context.Context, NodeFileListResult) (NodeFileListResult, error)) NodeFileListResultPage {
	return original.NewNodeFileListResultPage(getNextPage)
}
func NewPoolClient(batchURL string) PoolClient {
	return original.NewPoolClient(batchURL)
}
func NewPoolListUsageMetricsResultIterator(page PoolListUsageMetricsResultPage) PoolListUsageMetricsResultIterator {
	return original.NewPoolListUsageMetricsResultIterator(page)
}
func NewPoolListUsageMetricsResultPage(getNextPage func(context.Context, PoolListUsageMetricsResult) (PoolListUsageMetricsResult, error)) PoolListUsageMetricsResultPage {
	return original.NewPoolListUsageMetricsResultPage(getNextPage)
}
func NewPoolNodeCountsListResultIterator(page PoolNodeCountsListResultPage) PoolNodeCountsListResultIterator {
	return original.NewPoolNodeCountsListResultIterator(page)
}
func NewPoolNodeCountsListResultPage(getNextPage func(context.Context, PoolNodeCountsListResult) (PoolNodeCountsListResult, error)) PoolNodeCountsListResultPage {
	return original.NewPoolNodeCountsListResultPage(getNextPage)
}
func NewTaskClient(batchURL string) TaskClient {
	return original.NewTaskClient(batchURL)
}
func NewWithoutDefaults(batchURL string) BaseClient {
	return original.NewWithoutDefaults(batchURL)
}
func PossibleAccessScopeValues() []AccessScope {
	return original.PossibleAccessScopeValues()
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleAutoUserScopeValues() []AutoUserScope {
	return original.PossibleAutoUserScopeValues()
}
func PossibleCachingTypeValues() []CachingType {
	return original.PossibleCachingTypeValues()
}
func PossibleCertificateFormatValues() []CertificateFormat {
	return original.PossibleCertificateFormatValues()
}
func PossibleCertificateStateValues() []CertificateState {
	return original.PossibleCertificateStateValues()
}
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return original.PossibleCertificateStoreLocationValues()
}
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return original.PossibleCertificateVisibilityValues()
}
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return original.PossibleComputeNodeDeallocationOptionValues()
}
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return original.PossibleComputeNodeFillTypeValues()
}
func PossibleComputeNodeRebootOptionValues() []ComputeNodeRebootOption {
	return original.PossibleComputeNodeRebootOptionValues()
}
func PossibleComputeNodeReimageOptionValues() []ComputeNodeReimageOption {
	return original.PossibleComputeNodeReimageOptionValues()
}
func PossibleComputeNodeStateValues() []ComputeNodeState {
	return original.PossibleComputeNodeStateValues()
}
func PossibleContainerWorkingDirectoryValues() []ContainerWorkingDirectory {
	return original.PossibleContainerWorkingDirectoryValues()
}
func PossibleDependencyActionValues() []DependencyAction {
	return original.PossibleDependencyActionValues()
}
func PossibleDisableComputeNodeSchedulingOptionValues() []DisableComputeNodeSchedulingOption {
	return original.PossibleDisableComputeNodeSchedulingOptionValues()
}
func PossibleDisableJobOptionValues() []DisableJobOption {
	return original.PossibleDisableJobOptionValues()
}
func PossibleDiskEncryptionTargetValues() []DiskEncryptionTarget {
	return original.PossibleDiskEncryptionTargetValues()
}
func PossibleDynamicVNetAssignmentScopeValues() []DynamicVNetAssignmentScope {
	return original.PossibleDynamicVNetAssignmentScopeValues()
}
func PossibleElevationLevelValues() []ElevationLevel {
	return original.PossibleElevationLevelValues()
}
func PossibleErrorCategoryValues() []ErrorCategory {
	return original.PossibleErrorCategoryValues()
}
func PossibleIPAddressProvisioningTypeValues() []IPAddressProvisioningType {
	return original.PossibleIPAddressProvisioningTypeValues()
}
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return original.PossibleInboundEndpointProtocolValues()
}
func PossibleJobActionValues() []JobAction {
	return original.PossibleJobActionValues()
}
func PossibleJobPreparationTaskStateValues() []JobPreparationTaskState {
	return original.PossibleJobPreparationTaskStateValues()
}
func PossibleJobReleaseTaskStateValues() []JobReleaseTaskState {
	return original.PossibleJobReleaseTaskStateValues()
}
func PossibleJobScheduleStateValues() []JobScheduleState {
	return original.PossibleJobScheduleStateValues()
}
func PossibleJobStateValues() []JobState {
	return original.PossibleJobStateValues()
}
func PossibleLoginModeValues() []LoginMode {
	return original.PossibleLoginModeValues()
}
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return original.PossibleNetworkSecurityGroupRuleAccessValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleOnAllTasksCompleteValues() []OnAllTasksComplete {
	return original.PossibleOnAllTasksCompleteValues()
}
func PossibleOnTaskFailureValues() []OnTaskFailure {
	return original.PossibleOnTaskFailureValues()
}
func PossibleOutputFileUploadConditionValues() []OutputFileUploadCondition {
	return original.PossibleOutputFileUploadConditionValues()
}
func PossiblePoolLifetimeOptionValues() []PoolLifetimeOption {
	return original.PossiblePoolLifetimeOptionValues()
}
func PossiblePoolStateValues() []PoolState {
	return original.PossiblePoolStateValues()
}
func PossibleSchedulingStateValues() []SchedulingState {
	return original.PossibleSchedulingStateValues()
}
func PossibleStartTaskStateValues() []StartTaskState {
	return original.PossibleStartTaskStateValues()
}
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return original.PossibleStorageAccountTypeValues()
}
func PossibleSubtaskStateValues() []SubtaskState {
	return original.PossibleSubtaskStateValues()
}
func PossibleTaskAddStatusValues() []TaskAddStatus {
	return original.PossibleTaskAddStatusValues()
}
func PossibleTaskExecutionResultValues() []TaskExecutionResult {
	return original.PossibleTaskExecutionResultValues()
}
func PossibleTaskStateValues() []TaskState {
	return original.PossibleTaskStateValues()
}
func PossibleVerificationTypeValues() []VerificationType {
	return original.PossibleVerificationTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
