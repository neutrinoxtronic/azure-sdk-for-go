//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// CalculateExchangePostPoller provides polling facilities until the operation reaches a terminal state.
type CalculateExchangePostPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *CalculateExchangePostPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *CalculateExchangePostPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final CalculateExchangePostResponse will be returned.
func (p *CalculateExchangePostPoller) FinalResponse(ctx context.Context) (CalculateExchangePostResponse, error) {
	respType := CalculateExchangePostResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.CalculateExchangeOperationResultResponse)
	if err != nil {
		return CalculateExchangePostResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *CalculateExchangePostPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ExchangePostPoller provides polling facilities until the operation reaches a terminal state.
type ExchangePostPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ExchangePostPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ExchangePostPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ExchangePostResponse will be returned.
func (p *ExchangePostPoller) FinalResponse(ctx context.Context) (ExchangePostResponse, error) {
	respType := ExchangePostResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ExchangeOperationResultResponse)
	if err != nil {
		return ExchangePostResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ExchangePostPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// QuotaCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type QuotaCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *QuotaCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *QuotaCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final QuotaCreateOrUpdateResponse will be returned.
func (p *QuotaCreateOrUpdatePoller) FinalResponse(ctx context.Context) (QuotaCreateOrUpdateResponse, error) {
	respType := QuotaCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.QuotaRequestOneResourceSubmitResponse)
	if err != nil {
		return QuotaCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *QuotaCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// QuotaUpdatePoller provides polling facilities until the operation reaches a terminal state.
type QuotaUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *QuotaUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *QuotaUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final QuotaUpdateResponse will be returned.
func (p *QuotaUpdatePoller) FinalResponse(ctx context.Context) (QuotaUpdateResponse, error) {
	respType := QuotaUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.QuotaRequestOneResourceSubmitResponse)
	if err != nil {
		return QuotaUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *QuotaUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationAvailableScopesPoller provides polling facilities until the operation reaches a terminal state.
type ReservationAvailableScopesPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationAvailableScopesPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationAvailableScopesPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationAvailableScopesResponse will be returned.
func (p *ReservationAvailableScopesPoller) FinalResponse(ctx context.Context) (ReservationAvailableScopesResponse, error) {
	respType := ReservationAvailableScopesResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.AvailableScopeProperties)
	if err != nil {
		return ReservationAvailableScopesResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationAvailableScopesPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationMergePoller provides polling facilities until the operation reaches a terminal state.
type ReservationMergePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationMergePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationMergePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationMergeResponse will be returned.
func (p *ReservationMergePoller) FinalResponse(ctx context.Context) (ReservationMergeResponse, error) {
	respType := ReservationMergeResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationResponseArray)
	if err != nil {
		return ReservationMergeResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationMergePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationOrderPurchasePoller provides polling facilities until the operation reaches a terminal state.
type ReservationOrderPurchasePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationOrderPurchasePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationOrderPurchasePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationOrderPurchaseResponse will be returned.
func (p *ReservationOrderPurchasePoller) FinalResponse(ctx context.Context) (ReservationOrderPurchaseResponse, error) {
	respType := ReservationOrderPurchaseResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationOrderResponse)
	if err != nil {
		return ReservationOrderPurchaseResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationOrderPurchasePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationSplitPoller provides polling facilities until the operation reaches a terminal state.
type ReservationSplitPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationSplitPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationSplitPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationSplitResponse will be returned.
func (p *ReservationSplitPoller) FinalResponse(ctx context.Context) (ReservationSplitResponse, error) {
	respType := ReservationSplitResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationResponseArray)
	if err != nil {
		return ReservationSplitResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationSplitPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ReservationUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ReservationUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ReservationUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ReservationUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ReservationUpdateResponse will be returned.
func (p *ReservationUpdatePoller) FinalResponse(ctx context.Context) (ReservationUpdateResponse, error) {
	respType := ReservationUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ReservationResponse)
	if err != nil {
		return ReservationUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ReservationUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}