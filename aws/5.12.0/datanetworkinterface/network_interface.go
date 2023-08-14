// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datanetworkinterface

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Association struct{}

type Attachment struct{}

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

type AssociationAttributes struct {
	ref terra.Reference
}

func (a AssociationAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AssociationAttributes) InternalWithRef(ref terra.Reference) AssociationAttributes {
	return AssociationAttributes{ref: ref}
}

func (a AssociationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AssociationAttributes) AllocationId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("allocation_id"))
}

func (a AssociationAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("association_id"))
}

func (a AssociationAttributes) CarrierIp() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("carrier_ip"))
}

func (a AssociationAttributes) CustomerOwnedIp() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("customer_owned_ip"))
}

func (a AssociationAttributes) IpOwnerId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("ip_owner_id"))
}

func (a AssociationAttributes) PublicDnsName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("public_dns_name"))
}

func (a AssociationAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("public_ip"))
}

type AttachmentAttributes struct {
	ref terra.Reference
}

func (a AttachmentAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AttachmentAttributes) InternalWithRef(ref terra.Reference) AttachmentAttributes {
	return AttachmentAttributes{ref: ref}
}

func (a AttachmentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AttachmentAttributes) AttachmentId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("attachment_id"))
}

func (a AttachmentAttributes) DeviceIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("device_index"))
}

func (a AttachmentAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("instance_id"))
}

func (a AttachmentAttributes) InstanceOwnerId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("instance_owner_id"))
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

type AssociationState struct {
	AllocationId    string `json:"allocation_id"`
	AssociationId   string `json:"association_id"`
	CarrierIp       string `json:"carrier_ip"`
	CustomerOwnedIp string `json:"customer_owned_ip"`
	IpOwnerId       string `json:"ip_owner_id"`
	PublicDnsName   string `json:"public_dns_name"`
	PublicIp        string `json:"public_ip"`
}

type AttachmentState struct {
	AttachmentId    string  `json:"attachment_id"`
	DeviceIndex     float64 `json:"device_index"`
	InstanceId      string  `json:"instance_id"`
	InstanceOwnerId string  `json:"instance_owner_id"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}