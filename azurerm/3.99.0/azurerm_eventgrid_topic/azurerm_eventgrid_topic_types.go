// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_eventgrid_topic

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type InboundIpRule struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// IpMask: string, optional
	IpMask terra.StringValue `hcl:"ip_mask,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type InputMappingDefaultValues struct {
	// DataVersion: string, optional
	DataVersion terra.StringValue `hcl:"data_version,attr"`
	// EventType: string, optional
	EventType terra.StringValue `hcl:"event_type,attr"`
	// Subject: string, optional
	Subject terra.StringValue `hcl:"subject,attr"`
}

type InputMappingFields struct {
	// DataVersion: string, optional
	DataVersion terra.StringValue `hcl:"data_version,attr"`
	// EventTime: string, optional
	EventTime terra.StringValue `hcl:"event_time,attr"`
	// EventType: string, optional
	EventType terra.StringValue `hcl:"event_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Subject: string, optional
	Subject terra.StringValue `hcl:"subject,attr"`
	// Topic: string, optional
	Topic terra.StringValue `hcl:"topic,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type InboundIpRuleAttributes struct {
	ref terra.Reference
}

func (iir InboundIpRuleAttributes) InternalRef() (terra.Reference, error) {
	return iir.ref, nil
}

func (iir InboundIpRuleAttributes) InternalWithRef(ref terra.Reference) InboundIpRuleAttributes {
	return InboundIpRuleAttributes{ref: ref}
}

func (iir InboundIpRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iir.ref.InternalTokens()
}

func (iir InboundIpRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("action"))
}

func (iir InboundIpRuleAttributes) IpMask() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("ip_mask"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type InputMappingDefaultValuesAttributes struct {
	ref terra.Reference
}

func (imdv InputMappingDefaultValuesAttributes) InternalRef() (terra.Reference, error) {
	return imdv.ref, nil
}

func (imdv InputMappingDefaultValuesAttributes) InternalWithRef(ref terra.Reference) InputMappingDefaultValuesAttributes {
	return InputMappingDefaultValuesAttributes{ref: ref}
}

func (imdv InputMappingDefaultValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return imdv.ref.InternalTokens()
}

func (imdv InputMappingDefaultValuesAttributes) DataVersion() terra.StringValue {
	return terra.ReferenceAsString(imdv.ref.Append("data_version"))
}

func (imdv InputMappingDefaultValuesAttributes) EventType() terra.StringValue {
	return terra.ReferenceAsString(imdv.ref.Append("event_type"))
}

func (imdv InputMappingDefaultValuesAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(imdv.ref.Append("subject"))
}

type InputMappingFieldsAttributes struct {
	ref terra.Reference
}

func (imf InputMappingFieldsAttributes) InternalRef() (terra.Reference, error) {
	return imf.ref, nil
}

func (imf InputMappingFieldsAttributes) InternalWithRef(ref terra.Reference) InputMappingFieldsAttributes {
	return InputMappingFieldsAttributes{ref: ref}
}

func (imf InputMappingFieldsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return imf.ref.InternalTokens()
}

func (imf InputMappingFieldsAttributes) DataVersion() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("data_version"))
}

func (imf InputMappingFieldsAttributes) EventTime() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("event_time"))
}

func (imf InputMappingFieldsAttributes) EventType() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("event_type"))
}

func (imf InputMappingFieldsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("id"))
}

func (imf InputMappingFieldsAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("subject"))
}

func (imf InputMappingFieldsAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(imf.ref.Append("topic"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type InboundIpRuleState struct {
	Action string `json:"action"`
	IpMask string `json:"ip_mask"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type InputMappingDefaultValuesState struct {
	DataVersion string `json:"data_version"`
	EventType   string `json:"event_type"`
	Subject     string `json:"subject"`
}

type InputMappingFieldsState struct {
	DataVersion string `json:"data_version"`
	EventTime   string `json:"event_time"`
	EventType   string `json:"event_type"`
	Id          string `json:"id"`
	Subject     string `json:"subject"`
	Topic       string `json:"topic"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
