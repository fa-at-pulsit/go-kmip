package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

import (
	"time"

	"github.com/pkg/errors"
)

// Request is a Request Message Structure
type Request struct {
	Tag `kmip:"REQUEST_MESSAGE"`

	Header     RequestHeader `kmip:"REQUEST_HEADER,required"`
	BatchItems []BatchItem   `kmip:"REQUEST_BATCH_ITEM,required"`
}

// RequestHeader is a Request Header Structure
type RequestHeader struct {
	Tag `kmip:"REQUEST_HEADER"`

	Version                      ProtocolVersion `kmip:"PROTOCOL_VERSION,required"`
	MaxResponseSize              int32           `kmip:"MAXIMUM_RESPONSE_SIZE"`
	ClientCorrelationValue       string          `kmip:"CLIENT_CORRELATION_VALUE"`
	ServerCorrelationValue       string          `kmip:"SERVER_CORRELATION_VALUE"`
	AsynchronousIndicator        bool            `kmip:"ASYNCHRONOUS_INDICATOR"`
	AttestationCapableIndicator  bool            `kmip:"ATTESTATION_CAPABLE_INDICATOR"`
	AttestationType              []Enum          `kmip:"ATTESTATION_TYPE"`
	Authentication               Authentication  `kmip:"AUTHENTICATION"`
	BatchErrorContinuationOption Enum            `kmip:"BATCH_ERROR_CONTINUATION_OPTION"`
	BatchOrderOption             bool            `kmip:"BATCH_ORDER_OPTION"`
	TimeStamp                    time.Time       `kmip:"TIME_STAMP"`
	BatchCount                   int32           `kmip:"BATCH_COUNT,required"`
}

// ProtocolVersion is a Protocol Version structure
type ProtocolVersion struct {
	Tag `kmip:"PROTOCOL_VERSION"`

	Major int32 `kmip:"PROTOCOL_VERSION_MAJOR"`
	Minor int32 `kmip:"PROTOCOL_VERSION_MINOR"`
}

// BatchItem is a Batch Item Structure
type BatchItem struct {
	Tag `kmip:"REQUEST_BATCH_ITEM"`

	Operation        Enum             `kmip:"OPERATION,required"`
	UniqueID         string           `kmip:"UNIQUE_BATCH_ITEM_ID"`
	RequestPayload   interface{}      `kmip:"REQUEST_PAYLOAD,required"`
	MessageExtension MessageExtension `kmip:"MESSAGE_EXTENSION"`
}

// BuildFieldValue builds value for RequestPayload based on Operation
func (bi *BatchItem) BuildFieldValue(name string) (v interface{}, err error) {
	switch bi.Operation {
	case OPERATION_CREATE:
		v = &OperationCreate{}
	case OPERATION_GET:
		v = &OperationGet{}
	default:
		err = errors.Errorf("unsupported operation: %v", bi.Operation)
	}

	return
}

// MessageExtension is a Message Extension structure in a Batch Item
type MessageExtension struct {
	Tag `kmip:"MESSAGE_EXTENSION"`

	VendorIdentification string      `kmip:"VENDOR_IDENTIFICATION,required"`
	CriticalityIndicator bool        `kmip:"CRITICALITY_INDICATOR,required"`
	VendorExtension      interface{} `kmip:"-,skip"`
}

// TemplateAttribute is a Template-Attribute Object Structure
type TemplateAttribute struct {
	Tag `kmip:"TEMPLATE_ATTRIBUTE"`

	Name       Name        `kmip:"NAME"`
	Attributes []Attribute `kmip:"ATTRIBUTE"`
}

// Name is a Name Attribute Structure
type Name struct {
	Tag `kmip:"NAME"`

	Value string `kmip:"NAME_VALUE,required"`
	Type  Enum   `kmip:"NAME_TYPE,required"`
}
