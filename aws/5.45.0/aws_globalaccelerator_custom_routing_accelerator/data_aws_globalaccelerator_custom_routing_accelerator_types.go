// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_globalaccelerator_custom_routing_accelerator

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAttributesAttributes struct {
	ref terra.Reference
}

func (a DataAttributesAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a DataAttributesAttributes) InternalWithRef(ref terra.Reference) DataAttributesAttributes {
	return DataAttributesAttributes{ref: ref}
}

func (a DataAttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a DataAttributesAttributes) FlowLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("flow_logs_enabled"))
}

func (a DataAttributesAttributes) FlowLogsS3Bucket() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("flow_logs_s3_bucket"))
}

func (a DataAttributesAttributes) FlowLogsS3Prefix() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("flow_logs_s3_prefix"))
}

type DataIpSetsAttributes struct {
	ref terra.Reference
}

func (is DataIpSetsAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is DataIpSetsAttributes) InternalWithRef(ref terra.Reference) DataIpSetsAttributes {
	return DataIpSetsAttributes{ref: ref}
}

func (is DataIpSetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is DataIpSetsAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](is.ref.Append("ip_addresses"))
}

func (is DataIpSetsAttributes) IpFamily() terra.StringValue {
	return terra.ReferenceAsString(is.ref.Append("ip_family"))
}

type DataAttributesState struct {
	FlowLogsEnabled  bool   `json:"flow_logs_enabled"`
	FlowLogsS3Bucket string `json:"flow_logs_s3_bucket"`
	FlowLogsS3Prefix string `json:"flow_logs_s3_prefix"`
}

type DataIpSetsState struct {
	IpAddresses []string `json:"ip_addresses"`
	IpFamily    string   `json:"ip_family"`
}
