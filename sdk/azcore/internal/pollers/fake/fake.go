//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package fake

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pollers"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
)

// Applicable returns true if the LRO is a fake.
func Applicable(resp *http.Response) bool {
	return resp.Header.Get(shared.HeaderFakePollerStatus) != ""
}

// CanResume returns true if the token can rehydrate this poller type.
func CanResume(token map[string]interface{}) bool {
	_, ok := token["fakeURL"]
	return ok
}

// Poller is an LRO poller that uses the Core-Fake-Poller pattern.
type Poller[T any] struct {
	pl exported.Pipeline

	resp *http.Response

	// The API name from CtxAPINameKey
	APIName string `json:"apiName"`

	// The URL from Core-Fake-Poller header.
	FakeURL string `json:"fakeURL"`

	// The LRO's current state.
	FakeStatus string `json:"status"`
}

// New creates a new Poller from the provided initial response.
// Pass nil for response to create an empty Poller for rehydration.
func New[T any](pl exported.Pipeline, resp *http.Response) (*Poller[T], error) {
	if resp == nil {
		log.Write(log.EventLRO, "Resuming Core-Fake-Poller poller.")
		return &Poller[T]{pl: pl}, nil
	}

	log.Write(log.EventLRO, "Using Core-Fake-Poller poller.")
	fakeStatus := resp.Header.Get(shared.HeaderFakePollerStatus)
	if fakeStatus == "" {
		return nil, errors.New("response is missing Fake-Poller-Status header")
	}

	ctxVal := resp.Request.Context().Value(shared.CtxAPINameKey{})
	if ctxVal == nil {
		return nil, errors.New("missing value for CtxAPINameKey")
	}

	apiName, ok := ctxVal.(string)
	if !ok {
		return nil, fmt.Errorf("expected string for CtxAPINameKey, the type was %T", ctxVal)
	}

	p := &Poller[T]{
		pl:         pl,
		resp:       resp,
		APIName:    apiName,
		FakeURL:    fmt.Sprintf("%s://%s%s/get/fake/status", resp.Request.URL.Scheme, resp.Request.URL.Host, resp.Request.URL.Path),
		FakeStatus: fakeStatus,
	}
	return p, nil
}

// Done returns true if the LRO is in a terminal state.
func (p *Poller[T]) Done() bool {
	return pollers.IsTerminalState(p.FakeStatus)
}

// Poll retrieves the current state of the LRO.
func (p *Poller[T]) Poll(ctx context.Context) (*http.Response, error) {
	ctx = context.WithValue(ctx, shared.CtxAPINameKey{}, p.APIName)
	err := pollers.PollHelper(ctx, p.FakeURL, p.pl, func(resp *http.Response) (string, error) {
		fakeStatus := resp.Header.Get(shared.HeaderFakePollerStatus)
		if fakeStatus == "" {
			return "", errors.New("response is missing Fake-Poller-Status header")
		}
		p.resp = resp
		p.FakeStatus = fakeStatus
		return p.FakeStatus, nil
	})
	if err != nil {
		return nil, err
	}
	return p.resp, nil
}

func (p *Poller[T]) Result(ctx context.Context, out *T) error {
	if p.resp.StatusCode == http.StatusNoContent {
		return nil
	} else if pollers.Failed(p.FakeStatus) {
		return exported.NewResponseError(p.resp)
	}

	return pollers.ResultHelper(p.resp, pollers.Failed(p.FakeStatus), out)
}
