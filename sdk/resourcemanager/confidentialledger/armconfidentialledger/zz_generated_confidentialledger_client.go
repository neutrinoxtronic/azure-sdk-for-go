//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfidentialledger

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConfidentialLedgerClient contains the methods for the ConfidentialLedger group.
// Don't use this type directly, use NewConfidentialLedgerClient() instead.
type ConfidentialLedgerClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConfidentialLedgerClient creates a new instance of ConfidentialLedgerClient with the specified values.
func NewConfidentialLedgerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfidentialLedgerClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ConfidentialLedgerClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - To check whether a resource name is available.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfidentialLedgerClient) CheckNameAvailability(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *ConfidentialLedgerCheckNameAvailabilityOptions) (ConfidentialLedgerCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return ConfidentialLedgerCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfidentialLedgerCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfidentialLedgerCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ConfidentialLedgerClient) checkNameAvailabilityCreateRequest(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *ConfidentialLedgerCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConfidentialLedger/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, nameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ConfidentialLedgerClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ConfidentialLedgerCheckNameAvailabilityResponse, error) {
	result := ConfidentialLedgerCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ConfidentialLedgerCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *ConfidentialLedgerClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}