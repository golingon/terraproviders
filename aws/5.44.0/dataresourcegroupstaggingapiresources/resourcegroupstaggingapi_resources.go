// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataresourcegroupstaggingapiresources

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ResourceTagMappingList struct {
	// ComplianceDetails: min=0
	ComplianceDetails []ComplianceDetails `hcl:"compliance_details,block" validate:"min=0"`
}

type ComplianceDetails struct{}

type TagFilter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type ResourceTagMappingListAttributes struct {
	ref terra.Reference
}

func (rtml ResourceTagMappingListAttributes) InternalRef() (terra.Reference, error) {
	return rtml.ref, nil
}

func (rtml ResourceTagMappingListAttributes) InternalWithRef(ref terra.Reference) ResourceTagMappingListAttributes {
	return ResourceTagMappingListAttributes{ref: ref}
}

func (rtml ResourceTagMappingListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rtml.ref.InternalTokens()
}

func (rtml ResourceTagMappingListAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(rtml.ref.Append("resource_arn"))
}

func (rtml ResourceTagMappingListAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rtml.ref.Append("tags"))
}

func (rtml ResourceTagMappingListAttributes) ComplianceDetails() terra.ListValue[ComplianceDetailsAttributes] {
	return terra.ReferenceAsList[ComplianceDetailsAttributes](rtml.ref.Append("compliance_details"))
}

type ComplianceDetailsAttributes struct {
	ref terra.Reference
}

func (cd ComplianceDetailsAttributes) InternalRef() (terra.Reference, error) {
	return cd.ref, nil
}

func (cd ComplianceDetailsAttributes) InternalWithRef(ref terra.Reference) ComplianceDetailsAttributes {
	return ComplianceDetailsAttributes{ref: ref}
}

func (cd ComplianceDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cd.ref.InternalTokens()
}

func (cd ComplianceDetailsAttributes) ComplianceStatus() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("compliance_status"))
}

func (cd ComplianceDetailsAttributes) KeysWithNoncompliantValues() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cd.ref.Append("keys_with_noncompliant_values"))
}

func (cd ComplianceDetailsAttributes) NonCompliantKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cd.ref.Append("non_compliant_keys"))
}

type TagFilterAttributes struct {
	ref terra.Reference
}

func (tf TagFilterAttributes) InternalRef() (terra.Reference, error) {
	return tf.ref, nil
}

func (tf TagFilterAttributes) InternalWithRef(ref terra.Reference) TagFilterAttributes {
	return TagFilterAttributes{ref: ref}
}

func (tf TagFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tf.ref.InternalTokens()
}

func (tf TagFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("key"))
}

func (tf TagFilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tf.ref.Append("values"))
}

type ResourceTagMappingListState struct {
	ResourceArn       string                   `json:"resource_arn"`
	Tags              map[string]string        `json:"tags"`
	ComplianceDetails []ComplianceDetailsState `json:"compliance_details"`
}

type ComplianceDetailsState struct {
	ComplianceStatus           bool     `json:"compliance_status"`
	KeysWithNoncompliantValues []string `json:"keys_with_noncompliant_values"`
	NonCompliantKeys           []string `json:"non_compliant_keys"`
}

type TagFilterState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}
