//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package generated

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start:  (*timeRFC3339)(a.Start),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type DequeuedMessage.
func (d DequeuedMessage) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias DequeuedMessage
	aux := &struct {
		*alias
		ExpirationTime  *timeRFC1123 `xml:"ExpirationTime"`
		InsertionTime   *timeRFC1123 `xml:"InsertionTime"`
		TimeNextVisible *timeRFC1123 `xml:"TimeNextVisible"`
	}{
		alias:           (*alias)(&d),
		ExpirationTime:  (*timeRFC1123)(d.ExpirationTime),
		InsertionTime:   (*timeRFC1123)(d.InsertionTime),
		TimeNextVisible: (*timeRFC1123)(d.TimeNextVisible),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type DequeuedMessage.
func (d *DequeuedMessage) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias DequeuedMessage
	aux := &struct {
		*alias
		ExpirationTime  *timeRFC1123 `xml:"ExpirationTime"`
		InsertionTime   *timeRFC1123 `xml:"InsertionTime"`
		TimeNextVisible *timeRFC1123 `xml:"TimeNextVisible"`
	}{
		alias: (*alias)(d),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	d.ExpirationTime = (*time.Time)(aux.ExpirationTime)
	d.InsertionTime = (*time.Time)(aux.InsertionTime)
	d.TimeNextVisible = (*time.Time)(aux.TimeNextVisible)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type EnqueuedMessage.
func (e EnqueuedMessage) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias EnqueuedMessage
	aux := &struct {
		*alias
		ExpirationTime  *timeRFC1123 `xml:"ExpirationTime"`
		InsertionTime   *timeRFC1123 `xml:"InsertionTime"`
		TimeNextVisible *timeRFC1123 `xml:"TimeNextVisible"`
	}{
		alias:           (*alias)(&e),
		ExpirationTime:  (*timeRFC1123)(e.ExpirationTime),
		InsertionTime:   (*timeRFC1123)(e.InsertionTime),
		TimeNextVisible: (*timeRFC1123)(e.TimeNextVisible),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type EnqueuedMessage.
func (e *EnqueuedMessage) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias EnqueuedMessage
	aux := &struct {
		*alias
		ExpirationTime  *timeRFC1123 `xml:"ExpirationTime"`
		InsertionTime   *timeRFC1123 `xml:"InsertionTime"`
		TimeNextVisible *timeRFC1123 `xml:"TimeNextVisible"`
	}{
		alias: (*alias)(e),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	e.ExpirationTime = (*time.Time)(aux.ExpirationTime)
	e.InsertionTime = (*time.Time)(aux.InsertionTime)
	e.TimeNextVisible = (*time.Time)(aux.TimeNextVisible)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type GeoReplication.
func (g GeoReplication) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias:        (*alias)(&g),
		LastSyncTime: (*timeRFC1123)(g.LastSyncTime),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type GeoReplication.
func (g *GeoReplication) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias: (*alias)(g),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	g.LastSyncTime = (*time.Time)(aux.LastSyncTime)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ListQueuesSegmentResponse.
func (l ListQueuesSegmentResponse) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ListQueuesSegmentResponse
	aux := &struct {
		*alias
		Queues *[]*Queue `xml:"Queues>Queue"`
	}{
		alias: (*alias)(&l),
	}
	if l.Queues != nil {
		aux.Queues = &l.Queues
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type PeekedMessage.
func (p PeekedMessage) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias PeekedMessage
	aux := &struct {
		*alias
		ExpirationTime *timeRFC1123 `xml:"ExpirationTime"`
		InsertionTime  *timeRFC1123 `xml:"InsertionTime"`
	}{
		alias:          (*alias)(&p),
		ExpirationTime: (*timeRFC1123)(p.ExpirationTime),
		InsertionTime:  (*timeRFC1123)(p.InsertionTime),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type PeekedMessage.
func (p *PeekedMessage) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias PeekedMessage
	aux := &struct {
		*alias
		ExpirationTime *timeRFC1123 `xml:"ExpirationTime"`
		InsertionTime  *timeRFC1123 `xml:"InsertionTime"`
	}{
		alias: (*alias)(p),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	p.ExpirationTime = (*time.Time)(aux.ExpirationTime)
	p.InsertionTime = (*time.Time)(aux.InsertionTime)
	return nil
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Queue.
func (q *Queue) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias Queue
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(q),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	q.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageError.
func (s StorageError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "Message", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageError.
func (s *StorageError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "Message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type StorageServiceProperties.
func (s StorageServiceProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias StorageServiceProperties
	aux := &struct {
		*alias
		CORS *[]*CORSRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&s),
	}
	if s.CORS != nil {
		aux.CORS = &s.CORS
	}
	return enc.EncodeElement(aux, start)
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
