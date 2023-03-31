// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package globalacceleratoraccelerator

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IpSets struct{}

type Attributes struct {
	// FlowLogsEnabled: bool, optional
	FlowLogsEnabled terra.BoolValue `hcl:"flow_logs_enabled,attr"`
	// FlowLogsS3Bucket: string, optional
	FlowLogsS3Bucket terra.StringValue `hcl:"flow_logs_s3_bucket,attr"`
	// FlowLogsS3Prefix: string, optional
	FlowLogsS3Prefix terra.StringValue `hcl:"flow_logs_s3_prefix,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type IpSetsAttributes struct {
	ref terra.Reference
}

func (is IpSetsAttributes) InternalRef() terra.Reference {
	return is.ref
}

func (is IpSetsAttributes) InternalWithRef(ref terra.Reference) IpSetsAttributes {
	return IpSetsAttributes{ref: ref}
}

func (is IpSetsAttributes) InternalTokens() hclwrite.Tokens {
	return is.ref.InternalTokens()
}

func (is IpSetsAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](is.ref.Append("ip_addresses"))
}

func (is IpSetsAttributes) IpFamily() terra.StringValue {
	return terra.ReferenceAsString(is.ref.Append("ip_family"))
}

type AttributesAttributes struct {
	ref terra.Reference
}

func (a AttributesAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AttributesAttributes) InternalWithRef(ref terra.Reference) AttributesAttributes {
	return AttributesAttributes{ref: ref}
}

func (a AttributesAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AttributesAttributes) FlowLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("flow_logs_enabled"))
}

func (a AttributesAttributes) FlowLogsS3Bucket() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("flow_logs_s3_bucket"))
}

func (a AttributesAttributes) FlowLogsS3Prefix() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("flow_logs_s3_prefix"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type IpSetsState struct {
	IpAddresses []string `json:"ip_addresses"`
	IpFamily    string   `json:"ip_family"`
}

type AttributesState struct {
	FlowLogsEnabled  bool   `json:"flow_logs_enabled"`
	FlowLogsS3Bucket string `json:"flow_logs_s3_bucket"`
	FlowLogsS3Prefix string `json:"flow_logs_s3_prefix"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
