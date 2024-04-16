// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_queue

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type OutboundCallerConfig struct {
	// OutboundCallerIdName: string, optional
	OutboundCallerIdName terra.StringValue `hcl:"outbound_caller_id_name,attr"`
	// OutboundCallerIdNumberId: string, optional
	OutboundCallerIdNumberId terra.StringValue `hcl:"outbound_caller_id_number_id,attr"`
	// OutboundFlowId: string, optional
	OutboundFlowId terra.StringValue `hcl:"outbound_flow_id,attr"`
}

type OutboundCallerConfigAttributes struct {
	ref terra.Reference
}

func (occ OutboundCallerConfigAttributes) InternalRef() (terra.Reference, error) {
	return occ.ref, nil
}

func (occ OutboundCallerConfigAttributes) InternalWithRef(ref terra.Reference) OutboundCallerConfigAttributes {
	return OutboundCallerConfigAttributes{ref: ref}
}

func (occ OutboundCallerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return occ.ref.InternalTokens()
}

func (occ OutboundCallerConfigAttributes) OutboundCallerIdName() terra.StringValue {
	return terra.ReferenceAsString(occ.ref.Append("outbound_caller_id_name"))
}

func (occ OutboundCallerConfigAttributes) OutboundCallerIdNumberId() terra.StringValue {
	return terra.ReferenceAsString(occ.ref.Append("outbound_caller_id_number_id"))
}

func (occ OutboundCallerConfigAttributes) OutboundFlowId() terra.StringValue {
	return terra.ReferenceAsString(occ.ref.Append("outbound_flow_id"))
}

type OutboundCallerConfigState struct {
	OutboundCallerIdName     string `json:"outbound_caller_id_name"`
	OutboundCallerIdNumberId string `json:"outbound_caller_id_number_id"`
	OutboundFlowId           string `json:"outbound_flow_id"`
}
