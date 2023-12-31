//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsqlvirtualmachine

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TroubleshootClient contains the methods for the SQLVirtualMachineTroubleshoot group.
// Don't use this type directly, use NewTroubleshootClient() instead.
type TroubleshootClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTroubleshootClient creates a new instance of TroubleshootClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTroubleshootClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TroubleshootClient, error) {
	cl, err := arm.NewClient(moduleName+".TroubleshootClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TroubleshootClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginTroubleshoot - Starts SQL virtual machine troubleshooting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineName - Name of the SQL virtual machine.
//   - parameters - The SQL virtual machine troubleshooting entity.
//   - options - TroubleshootClientBeginTroubleshootOptions contains the optional parameters for the TroubleshootClient.BeginTroubleshoot
//     method.
func (client *TroubleshootClient) BeginTroubleshoot(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVMTroubleshooting, options *TroubleshootClientBeginTroubleshootOptions) (*runtime.Poller[TroubleshootClientTroubleshootResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.troubleshoot(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TroubleshootClientTroubleshootResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[TroubleshootClientTroubleshootResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Troubleshoot - Starts SQL virtual machine troubleshooting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
func (client *TroubleshootClient) troubleshoot(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVMTroubleshooting, options *TroubleshootClientBeginTroubleshootOptions) (*http.Response, error) {
	req, err := client.troubleshootCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// troubleshootCreateRequest creates the Troubleshoot request.
func (client *TroubleshootClient) troubleshootCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVMTroubleshooting, options *TroubleshootClientBeginTroubleshootOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}/troubleshoot"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
