// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package transferserver

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EndpointDetails struct {
	// AddressAllocationIds: set of string, optional
	AddressAllocationIds terra.SetValue[terra.StringValue] `hcl:"address_allocation_ids,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// VpcEndpointId: string, optional
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
}

type ProtocolDetails struct {
	// As2Transports: set of string, optional
	As2Transports terra.SetValue[terra.StringValue] `hcl:"as2_transports,attr"`
	// PassiveIp: string, optional
	PassiveIp terra.StringValue `hcl:"passive_ip,attr"`
	// SetStatOption: string, optional
	SetStatOption terra.StringValue `hcl:"set_stat_option,attr"`
	// TlsSessionResumptionMode: string, optional
	TlsSessionResumptionMode terra.StringValue `hcl:"tls_session_resumption_mode,attr"`
}

type WorkflowDetails struct {
	// OnPartialUpload: optional
	OnPartialUpload *OnPartialUpload `hcl:"on_partial_upload,block"`
	// OnUpload: optional
	OnUpload *OnUpload `hcl:"on_upload,block"`
}

type OnPartialUpload struct {
	// ExecutionRole: string, required
	ExecutionRole terra.StringValue `hcl:"execution_role,attr" validate:"required"`
	// WorkflowId: string, required
	WorkflowId terra.StringValue `hcl:"workflow_id,attr" validate:"required"`
}

type OnUpload struct {
	// ExecutionRole: string, required
	ExecutionRole terra.StringValue `hcl:"execution_role,attr" validate:"required"`
	// WorkflowId: string, required
	WorkflowId terra.StringValue `hcl:"workflow_id,attr" validate:"required"`
}

type EndpointDetailsAttributes struct {
	ref terra.Reference
}

func (ed EndpointDetailsAttributes) InternalRef() (terra.Reference, error) {
	return ed.ref, nil
}

func (ed EndpointDetailsAttributes) InternalWithRef(ref terra.Reference) EndpointDetailsAttributes {
	return EndpointDetailsAttributes{ref: ref}
}

func (ed EndpointDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ed.ref.InternalTokens()
}

func (ed EndpointDetailsAttributes) AddressAllocationIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ed.ref.Append("address_allocation_ids"))
}

func (ed EndpointDetailsAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ed.ref.Append("security_group_ids"))
}

func (ed EndpointDetailsAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ed.ref.Append("subnet_ids"))
}

func (ed EndpointDetailsAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("vpc_endpoint_id"))
}

func (ed EndpointDetailsAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("vpc_id"))
}

type ProtocolDetailsAttributes struct {
	ref terra.Reference
}

func (pd ProtocolDetailsAttributes) InternalRef() (terra.Reference, error) {
	return pd.ref, nil
}

func (pd ProtocolDetailsAttributes) InternalWithRef(ref terra.Reference) ProtocolDetailsAttributes {
	return ProtocolDetailsAttributes{ref: ref}
}

func (pd ProtocolDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pd.ref.InternalTokens()
}

func (pd ProtocolDetailsAttributes) As2Transports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pd.ref.Append("as2_transports"))
}

func (pd ProtocolDetailsAttributes) PassiveIp() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("passive_ip"))
}

func (pd ProtocolDetailsAttributes) SetStatOption() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("set_stat_option"))
}

func (pd ProtocolDetailsAttributes) TlsSessionResumptionMode() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("tls_session_resumption_mode"))
}

type WorkflowDetailsAttributes struct {
	ref terra.Reference
}

func (wd WorkflowDetailsAttributes) InternalRef() (terra.Reference, error) {
	return wd.ref, nil
}

func (wd WorkflowDetailsAttributes) InternalWithRef(ref terra.Reference) WorkflowDetailsAttributes {
	return WorkflowDetailsAttributes{ref: ref}
}

func (wd WorkflowDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wd.ref.InternalTokens()
}

func (wd WorkflowDetailsAttributes) OnPartialUpload() terra.ListValue[OnPartialUploadAttributes] {
	return terra.ReferenceAsList[OnPartialUploadAttributes](wd.ref.Append("on_partial_upload"))
}

func (wd WorkflowDetailsAttributes) OnUpload() terra.ListValue[OnUploadAttributes] {
	return terra.ReferenceAsList[OnUploadAttributes](wd.ref.Append("on_upload"))
}

type OnPartialUploadAttributes struct {
	ref terra.Reference
}

func (opu OnPartialUploadAttributes) InternalRef() (terra.Reference, error) {
	return opu.ref, nil
}

func (opu OnPartialUploadAttributes) InternalWithRef(ref terra.Reference) OnPartialUploadAttributes {
	return OnPartialUploadAttributes{ref: ref}
}

func (opu OnPartialUploadAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return opu.ref.InternalTokens()
}

func (opu OnPartialUploadAttributes) ExecutionRole() terra.StringValue {
	return terra.ReferenceAsString(opu.ref.Append("execution_role"))
}

func (opu OnPartialUploadAttributes) WorkflowId() terra.StringValue {
	return terra.ReferenceAsString(opu.ref.Append("workflow_id"))
}

type OnUploadAttributes struct {
	ref terra.Reference
}

func (ou OnUploadAttributes) InternalRef() (terra.Reference, error) {
	return ou.ref, nil
}

func (ou OnUploadAttributes) InternalWithRef(ref terra.Reference) OnUploadAttributes {
	return OnUploadAttributes{ref: ref}
}

func (ou OnUploadAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ou.ref.InternalTokens()
}

func (ou OnUploadAttributes) ExecutionRole() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("execution_role"))
}

func (ou OnUploadAttributes) WorkflowId() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("workflow_id"))
}

type EndpointDetailsState struct {
	AddressAllocationIds []string `json:"address_allocation_ids"`
	SecurityGroupIds     []string `json:"security_group_ids"`
	SubnetIds            []string `json:"subnet_ids"`
	VpcEndpointId        string   `json:"vpc_endpoint_id"`
	VpcId                string   `json:"vpc_id"`
}

type ProtocolDetailsState struct {
	As2Transports            []string `json:"as2_transports"`
	PassiveIp                string   `json:"passive_ip"`
	SetStatOption            string   `json:"set_stat_option"`
	TlsSessionResumptionMode string   `json:"tls_session_resumption_mode"`
}

type WorkflowDetailsState struct {
	OnPartialUpload []OnPartialUploadState `json:"on_partial_upload"`
	OnUpload        []OnUploadState        `json:"on_upload"`
}

type OnPartialUploadState struct {
	ExecutionRole string `json:"execution_role"`
	WorkflowId    string `json:"workflow_id"`
}

type OnUploadState struct {
	ExecutionRole string `json:"execution_role"`
	WorkflowId    string `json:"workflow_id"`
}
