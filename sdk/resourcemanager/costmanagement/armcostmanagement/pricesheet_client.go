//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement

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

// PriceSheetClient contains the methods for the PriceSheet group.
// Don't use this type directly, use NewPriceSheetClient() instead.
type PriceSheetClient struct {
	internal *arm.Client
}

// NewPriceSheetClient creates a new instance of PriceSheetClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPriceSheetClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PriceSheetClient, error) {
	cl, err := arm.NewClient(moduleName+".PriceSheetClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PriceSheetClient{
		internal: cl,
	}
	return client, nil
}

// BeginDownload - Gets a URL to download the pricesheet for an invoice. The operation is supported for billing accounts with
// agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceName - The ID that uniquely identifies an invoice.
//   - options - PriceSheetClientBeginDownloadOptions contains the optional parameters for the PriceSheetClient.BeginDownload
//     method.
func (client *PriceSheetClient) BeginDownload(ctx context.Context, billingAccountName string, billingProfileName string, invoiceName string, options *PriceSheetClientBeginDownloadOptions) (*runtime.Poller[PriceSheetClientDownloadResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.download(ctx, billingAccountName, billingProfileName, invoiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PriceSheetClientDownloadResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PriceSheetClientDownloadResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Download - Gets a URL to download the pricesheet for an invoice. The operation is supported for billing accounts with agreement
// type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
func (client *PriceSheetClient) download(ctx context.Context, billingAccountName string, billingProfileName string, invoiceName string, options *PriceSheetClientBeginDownloadOptions) (*http.Response, error) {
	req, err := client.downloadCreateRequest(ctx, billingAccountName, billingProfileName, invoiceName, options)
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

// downloadCreateRequest creates the Download request.
func (client *PriceSheetClient) downloadCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceName string, options *PriceSheetClientBeginDownloadOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices/{invoiceName}/providers/Microsoft.CostManagement/pricesheets/default/download"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDownloadByBillingProfile - Gets a URL to download the current month's pricesheet for a billing profile. The operation
// is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer
// Agreement.Due to Azure product growth, the Azure price sheet download experience in this preview version will be updated
// from a single csv file to a Zip file containing multiple csv files, each with
// max 200k records.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - PriceSheetClientBeginDownloadByBillingProfileOptions contains the optional parameters for the PriceSheetClient.BeginDownloadByBillingProfile
//     method.
func (client *PriceSheetClient) BeginDownloadByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, options *PriceSheetClientBeginDownloadByBillingProfileOptions) (*runtime.Poller[PriceSheetClientDownloadByBillingProfileResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.downloadByBillingProfile(ctx, billingAccountName, billingProfileName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PriceSheetClientDownloadByBillingProfileResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PriceSheetClientDownloadByBillingProfileResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DownloadByBillingProfile - Gets a URL to download the current month's pricesheet for a billing profile. The operation is
// supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer
// Agreement.Due to Azure product growth, the Azure price sheet download experience in this preview version will be updated
// from a single csv file to a Zip file containing multiple csv files, each with
// max 200k records.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
func (client *PriceSheetClient) downloadByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, options *PriceSheetClientBeginDownloadByBillingProfileOptions) (*http.Response, error) {
	req, err := client.downloadByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
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

// downloadByBillingProfileCreateRequest creates the DownloadByBillingProfile request.
func (client *PriceSheetClient) downloadByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *PriceSheetClientBeginDownloadByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.CostManagement/pricesheets/default/download"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
