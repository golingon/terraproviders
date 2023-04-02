// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3controlbucketlifecycleconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// AbortIncompleteMultipartUpload: optional
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUpload `hcl:"abort_incomplete_multipart_upload,block"`
	// Expiration: optional
	Expiration *Expiration `hcl:"expiration,block"`
	// Filter: optional
	Filter *Filter `hcl:"filter,block"`
}

type AbortIncompleteMultipartUpload struct {
	// DaysAfterInitiation: number, required
	DaysAfterInitiation terra.NumberValue `hcl:"days_after_initiation,attr" validate:"required"`
}

type Expiration struct {
	// Date: string, optional
	Date terra.StringValue `hcl:"date,attr"`
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// ExpiredObjectDeleteMarker: bool, optional
	ExpiredObjectDeleteMarker terra.BoolValue `hcl:"expired_object_delete_marker,attr"`
}

type Filter struct {
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

func (r RuleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status"))
}

func (r RuleAttributes) AbortIncompleteMultipartUpload() terra.ListValue[AbortIncompleteMultipartUploadAttributes] {
	return terra.ReferenceAsList[AbortIncompleteMultipartUploadAttributes](r.ref.Append("abort_incomplete_multipart_upload"))
}

func (r RuleAttributes) Expiration() terra.ListValue[ExpirationAttributes] {
	return terra.ReferenceAsList[ExpirationAttributes](r.ref.Append("expiration"))
}

func (r RuleAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](r.ref.Append("filter"))
}

type AbortIncompleteMultipartUploadAttributes struct {
	ref terra.Reference
}

func (aimu AbortIncompleteMultipartUploadAttributes) InternalRef() (terra.Reference, error) {
	return aimu.ref, nil
}

func (aimu AbortIncompleteMultipartUploadAttributes) InternalWithRef(ref terra.Reference) AbortIncompleteMultipartUploadAttributes {
	return AbortIncompleteMultipartUploadAttributes{ref: ref}
}

func (aimu AbortIncompleteMultipartUploadAttributes) InternalTokens() hclwrite.Tokens {
	return aimu.ref.InternalTokens()
}

func (aimu AbortIncompleteMultipartUploadAttributes) DaysAfterInitiation() terra.NumberValue {
	return terra.ReferenceAsNumber(aimu.ref.Append("days_after_initiation"))
}

type ExpirationAttributes struct {
	ref terra.Reference
}

func (e ExpirationAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExpirationAttributes) InternalWithRef(ref terra.Reference) ExpirationAttributes {
	return ExpirationAttributes{ref: ref}
}

func (e ExpirationAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e ExpirationAttributes) Date() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("date"))
}

func (e ExpirationAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("days"))
}

func (e ExpirationAttributes) ExpiredObjectDeleteMarker() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("expired_object_delete_marker"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("prefix"))
}

func (f FilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("tags"))
}

type RuleState struct {
	Id                             string                                `json:"id"`
	Status                         string                                `json:"status"`
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadState `json:"abort_incomplete_multipart_upload"`
	Expiration                     []ExpirationState                     `json:"expiration"`
	Filter                         []FilterState                         `json:"filter"`
}

type AbortIncompleteMultipartUploadState struct {
	DaysAfterInitiation float64 `json:"days_after_initiation"`
}

type ExpirationState struct {
	Date                      string  `json:"date"`
	Days                      float64 `json:"days"`
	ExpiredObjectDeleteMarker bool    `json:"expired_object_delete_marker"`
}

type FilterState struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}
