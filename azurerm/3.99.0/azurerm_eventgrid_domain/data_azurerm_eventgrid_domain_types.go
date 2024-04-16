// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_eventgrid_domain

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataIdentityAttributes struct {
	ref terra.Reference
}

func (i DataIdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataIdentityAttributes) InternalWithRef(ref terra.Reference) DataIdentityAttributes {
	return DataIdentityAttributes{ref: ref}
}

func (i DataIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataIdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i DataIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i DataIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i DataIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type DataInboundIpRuleAttributes struct {
	ref terra.Reference
}

func (iir DataInboundIpRuleAttributes) InternalRef() (terra.Reference, error) {
	return iir.ref, nil
}

func (iir DataInboundIpRuleAttributes) InternalWithRef(ref terra.Reference) DataInboundIpRuleAttributes {
	return DataInboundIpRuleAttributes{ref: ref}
}

func (iir DataInboundIpRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iir.ref.InternalTokens()
}

func (iir DataInboundIpRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("action"))
}

func (iir DataInboundIpRuleAttributes) IpMask() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("ip_mask"))
}

type DataInputMappingDefaultValuesAttributes struct {
	ref terra.Reference
}

func (imdv DataInputMappingDefaultValuesAttributes) InternalRef() (terra.Reference, error) {
	return imdv.ref, nil
}

func (imdv DataInputMappingDefaultValuesAttributes) InternalWithRef(ref terra.Reference) DataInputMappingDefaultValuesAttributes {
	return DataInputMappingDefaultValuesAttributes{ref: ref}
}

func (imdv DataInputMappingDefaultValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return imdv.ref.InternalTokens()
}

func (imdv DataInputMappingDefaultValuesAttributes) DataVersion() terra.StringValue {
	return terra.ReferenceAsString(imdv.ref.Append("data_version"))
}

func (imdv DataInputMappingDefaultValuesAttributes) EventType() terra.StringValue {
	return terra.ReferenceAsString(imdv.ref.Append("event_type"))
}

func (imdv DataInputMappingDefaultValuesAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(imdv.ref.Append("subject"))
}

type DataInputMappingFieldsAttributes struct {
	ref terra.Reference
}

func (imf DataInputMappingFieldsAttributes) InternalRef() (terra.Reference, error) {
	return imf.ref, nil
}

func (imf DataInputMappingFieldsAttributes) InternalWithRef(ref terra.Reference) DataInputMappingFieldsAttributes {
	return DataInputMappingFieldsAttributes{ref: ref}
}

func (imf DataInputMappingFieldsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return imf.ref.InternalTokens()
}

func (imf DataInputMappingFieldsAttributes) DataVersion() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("data_version"))
}

func (imf DataInputMappingFieldsAttributes) EventTime() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("event_time"))
}

func (imf DataInputMappingFieldsAttributes) EventType() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("event_type"))
}

func (imf DataInputMappingFieldsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("id"))
}

func (imf DataInputMappingFieldsAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("subject"))
}

func (imf DataInputMappingFieldsAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("topic"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataIdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type DataInboundIpRuleState struct {
	Action string `json:"action"`
	IpMask string `json:"ip_mask"`
}

type DataInputMappingDefaultValuesState struct {
	DataVersion string `json:"data_version"`
	EventType   string `json:"event_type"`
	Subject     string `json:"subject"`
}

type DataInputMappingFieldsState struct {
	DataVersion string `json:"data_version"`
	EventTime   string `json:"event_time"`
	EventType   string `json:"event_type"`
	Id          string `json:"id"`
	Subject     string `json:"subject"`
	Topic       string `json:"topic"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
