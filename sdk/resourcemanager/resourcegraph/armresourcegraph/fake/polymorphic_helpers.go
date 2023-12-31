//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

func unmarshalFacetClassification(rawMsg json.RawMessage) (armresourcegraph.FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armresourcegraph.FacetClassification
	switch m["resultType"] {
	case "FacetError":
		b = &armresourcegraph.FacetError{}
	case "FacetResult":
		b = &armresourcegraph.FacetResult{}
	default:
		b = &armresourcegraph.Facet{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFacetClassificationArray(rawMsg json.RawMessage) ([]armresourcegraph.FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armresourcegraph.FacetClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFacetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
