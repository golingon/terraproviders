// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_queue

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataOutboundCallerConfigAttributes struct {
	ref terra.Reference
}

func (occ DataOutboundCallerConfigAttributes) InternalRef() (terra.Reference, error) {
	return occ.ref, nil
}

func (occ DataOutboundCallerConfigAttributes) InternalWithRef(ref terra.Reference) DataOutboundCallerConfigAttributes {
	return DataOutboundCallerConfigAttributes{ref: ref}
}

func (occ DataOutboundCallerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return occ.ref.InternalTokens()
}

func (occ DataOutboundCallerConfigAttributes) OutboundCallerIdName() terra.StringValue {
	return terra.ReferenceAsString(occ.ref.Append("outbound_caller_id_name"))
}

func (occ DataOutboundCallerConfigAttributes) OutboundCallerIdNumberId() terra.StringValue {
	return terra.ReferenceAsString(occ.ref.Append("outbound_caller_id_number_id"))
}

func (occ DataOutboundCallerConfigAttributes) OutboundFlowId() terra.StringValue {
	return terra.ReferenceAsString(occ.ref.Append("outbound_flow_id"))
}

type DataOutboundCallerConfigState struct {
	OutboundCallerIdName     string `json:"outbound_caller_id_name"`
	OutboundCallerIdNumberId string `json:"outbound_caller_id_number_id"`
	OutboundFlowId           string `json:"outbound_flow_id"`
}
