// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataec2transitgatewaymulticastdomain

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Associations struct{}

type Members struct{}

type Sources struct{}

type Filter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AssociationsAttributes struct {
	ref terra.Reference
}

func (a AssociationsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AssociationsAttributes) InternalWithRef(ref terra.Reference) AssociationsAttributes {
	return AssociationsAttributes{ref: ref}
}

func (a AssociationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AssociationsAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("subnet_id"))
}

func (a AssociationsAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("transit_gateway_attachment_id"))
}

type MembersAttributes struct {
	ref terra.Reference
}

func (m MembersAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MembersAttributes) InternalWithRef(ref terra.Reference) MembersAttributes {
	return MembersAttributes{ref: ref}
}

func (m MembersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MembersAttributes) GroupIpAddress() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("group_ip_address"))
}

func (m MembersAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("network_interface_id"))
}

type SourcesAttributes struct {
	ref terra.Reference
}

func (s SourcesAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourcesAttributes) InternalWithRef(ref terra.Reference) SourcesAttributes {
	return SourcesAttributes{ref: ref}
}

func (s SourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourcesAttributes) GroupIpAddress() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("group_ip_address"))
}

func (s SourcesAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("network_interface_id"))
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

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f FilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("values"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type AssociationsState struct {
	SubnetId                   string `json:"subnet_id"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
}

type MembersState struct {
	GroupIpAddress     string `json:"group_ip_address"`
	NetworkInterfaceId string `json:"network_interface_id"`
}

type SourcesState struct {
	GroupIpAddress     string `json:"group_ip_address"`
	NetworkInterfaceId string `json:"network_interface_id"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
