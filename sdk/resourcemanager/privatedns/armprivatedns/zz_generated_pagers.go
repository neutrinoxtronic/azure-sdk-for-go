//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PrivateZonesListByResourceGroupPager provides operations for iterating over paged responses.
type PrivateZonesListByResourceGroupPager struct {
	client    *PrivateZonesClient
	current   PrivateZonesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateZonesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateZonesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateZonesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateZoneListResult.NextLink == nil || len(*p.current.PrivateZoneListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateZonesListByResourceGroupResponse page.
func (p *PrivateZonesListByResourceGroupPager) PageResponse() PrivateZonesListByResourceGroupResponse {
	return p.current
}

// PrivateZonesListPager provides operations for iterating over paged responses.
type PrivateZonesListPager struct {
	client    *PrivateZonesClient
	current   PrivateZonesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateZonesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateZonesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateZonesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateZoneListResult.NextLink == nil || len(*p.current.PrivateZoneListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateZonesListResponse page.
func (p *PrivateZonesListPager) PageResponse() PrivateZonesListResponse {
	return p.current
}

// RecordSetsListByTypePager provides operations for iterating over paged responses.
type RecordSetsListByTypePager struct {
	client    *RecordSetsClient
	current   RecordSetsListByTypeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecordSetsListByTypeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecordSetsListByTypePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecordSetsListByTypePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecordSetListResult.NextLink == nil || len(*p.current.RecordSetListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByTypeHandleError(resp)
		return false
	}
	result, err := p.client.listByTypeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecordSetsListByTypeResponse page.
func (p *RecordSetsListByTypePager) PageResponse() RecordSetsListByTypeResponse {
	return p.current
}

// RecordSetsListPager provides operations for iterating over paged responses.
type RecordSetsListPager struct {
	client    *RecordSetsClient
	current   RecordSetsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecordSetsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecordSetsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecordSetsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecordSetListResult.NextLink == nil || len(*p.current.RecordSetListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecordSetsListResponse page.
func (p *RecordSetsListPager) PageResponse() RecordSetsListResponse {
	return p.current
}

// VirtualNetworkLinksListPager provides operations for iterating over paged responses.
type VirtualNetworkLinksListPager struct {
	client    *VirtualNetworkLinksClient
	current   VirtualNetworkLinksListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworkLinksListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualNetworkLinksListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualNetworkLinksListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkLinkListResult.NextLink == nil || len(*p.current.VirtualNetworkLinkListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current VirtualNetworkLinksListResponse page.
func (p *VirtualNetworkLinksListPager) PageResponse() VirtualNetworkLinksListResponse {
	return p.current
}