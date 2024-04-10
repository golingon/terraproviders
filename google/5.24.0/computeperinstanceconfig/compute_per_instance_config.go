// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package computeperinstanceconfig

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PreservedState struct {
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Disk: min=0
	Disk []Disk `hcl:"disk,block" validate:"min=0"`
	// ExternalIp: min=0
	ExternalIp []ExternalIp `hcl:"external_ip,block" validate:"min=0"`
	// InternalIp: min=0
	InternalIp []InternalIp `hcl:"internal_ip,block" validate:"min=0"`
}

type Disk struct {
	// DeleteRule: string, optional
	DeleteRule terra.StringValue `hcl:"delete_rule,attr"`
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
}

type ExternalIp struct {
	// AutoDelete: string, optional
	AutoDelete terra.StringValue `hcl:"auto_delete,attr"`
	// InterfaceName: string, required
	InterfaceName terra.StringValue `hcl:"interface_name,attr" validate:"required"`
	// ExternalIpIpAddress: optional
	IpAddress *ExternalIpIpAddress `hcl:"ip_address,block"`
}

type ExternalIpIpAddress struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
}

type InternalIp struct {
	// AutoDelete: string, optional
	AutoDelete terra.StringValue `hcl:"auto_delete,attr"`
	// InterfaceName: string, required
	InterfaceName terra.StringValue `hcl:"interface_name,attr" validate:"required"`
	// InternalIpIpAddress: optional
	IpAddress *InternalIpIpAddress `hcl:"ip_address,block"`
}

type InternalIpIpAddress struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PreservedStateAttributes struct {
	ref terra.Reference
}

func (ps PreservedStateAttributes) InternalRef() (terra.Reference, error) {
	return ps.ref, nil
}

func (ps PreservedStateAttributes) InternalWithRef(ref terra.Reference) PreservedStateAttributes {
	return PreservedStateAttributes{ref: ref}
}

func (ps PreservedStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ps.ref.InternalTokens()
}

func (ps PreservedStateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ps.ref.Append("metadata"))
}

func (ps PreservedStateAttributes) Disk() terra.SetValue[DiskAttributes] {
	return terra.ReferenceAsSet[DiskAttributes](ps.ref.Append("disk"))
}

func (ps PreservedStateAttributes) ExternalIp() terra.SetValue[ExternalIpAttributes] {
	return terra.ReferenceAsSet[ExternalIpAttributes](ps.ref.Append("external_ip"))
}

func (ps PreservedStateAttributes) InternalIp() terra.SetValue[InternalIpAttributes] {
	return terra.ReferenceAsSet[InternalIpAttributes](ps.ref.Append("internal_ip"))
}

type DiskAttributes struct {
	ref terra.Reference
}

func (d DiskAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DiskAttributes) InternalWithRef(ref terra.Reference) DiskAttributes {
	return DiskAttributes{ref: ref}
}

func (d DiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DiskAttributes) DeleteRule() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("delete_rule"))
}

func (d DiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("device_name"))
}

func (d DiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("mode"))
}

func (d DiskAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source"))
}

type ExternalIpAttributes struct {
	ref terra.Reference
}

func (ei ExternalIpAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei ExternalIpAttributes) InternalWithRef(ref terra.Reference) ExternalIpAttributes {
	return ExternalIpAttributes{ref: ref}
}

func (ei ExternalIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei ExternalIpAttributes) AutoDelete() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("auto_delete"))
}

func (ei ExternalIpAttributes) InterfaceName() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("interface_name"))
}

func (ei ExternalIpAttributes) IpAddress() terra.ListValue[ExternalIpIpAddressAttributes] {
	return terra.ReferenceAsList[ExternalIpIpAddressAttributes](ei.ref.Append("ip_address"))
}

type ExternalIpIpAddressAttributes struct {
	ref terra.Reference
}

func (ia ExternalIpIpAddressAttributes) InternalRef() (terra.Reference, error) {
	return ia.ref, nil
}

func (ia ExternalIpIpAddressAttributes) InternalWithRef(ref terra.Reference) ExternalIpIpAddressAttributes {
	return ExternalIpIpAddressAttributes{ref: ref}
}

func (ia ExternalIpIpAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ia.ref.InternalTokens()
}

func (ia ExternalIpIpAddressAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("address"))
}

type InternalIpAttributes struct {
	ref terra.Reference
}

func (ii InternalIpAttributes) InternalRef() (terra.Reference, error) {
	return ii.ref, nil
}

func (ii InternalIpAttributes) InternalWithRef(ref terra.Reference) InternalIpAttributes {
	return InternalIpAttributes{ref: ref}
}

func (ii InternalIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ii.ref.InternalTokens()
}

func (ii InternalIpAttributes) AutoDelete() terra.StringValue {
	return terra.ReferenceAsString(ii.ref.Append("auto_delete"))
}

func (ii InternalIpAttributes) InterfaceName() terra.StringValue {
	return terra.ReferenceAsString(ii.ref.Append("interface_name"))
}

func (ii InternalIpAttributes) IpAddress() terra.ListValue[InternalIpIpAddressAttributes] {
	return terra.ReferenceAsList[InternalIpIpAddressAttributes](ii.ref.Append("ip_address"))
}

type InternalIpIpAddressAttributes struct {
	ref terra.Reference
}

func (ia InternalIpIpAddressAttributes) InternalRef() (terra.Reference, error) {
	return ia.ref, nil
}

func (ia InternalIpIpAddressAttributes) InternalWithRef(ref terra.Reference) InternalIpIpAddressAttributes {
	return InternalIpIpAddressAttributes{ref: ref}
}

func (ia InternalIpIpAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ia.ref.InternalTokens()
}

func (ia InternalIpIpAddressAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("address"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type PreservedStateState struct {
	Metadata   map[string]string `json:"metadata"`
	Disk       []DiskState       `json:"disk"`
	ExternalIp []ExternalIpState `json:"external_ip"`
	InternalIp []InternalIpState `json:"internal_ip"`
}

type DiskState struct {
	DeleteRule string `json:"delete_rule"`
	DeviceName string `json:"device_name"`
	Mode       string `json:"mode"`
	Source     string `json:"source"`
}

type ExternalIpState struct {
	AutoDelete    string                     `json:"auto_delete"`
	InterfaceName string                     `json:"interface_name"`
	IpAddress     []ExternalIpIpAddressState `json:"ip_address"`
}

type ExternalIpIpAddressState struct {
	Address string `json:"address"`
}

type InternalIpState struct {
	AutoDelete    string                     `json:"auto_delete"`
	InterfaceName string                     `json:"interface_name"`
	IpAddress     []InternalIpIpAddressState `json:"ip_address"`
}

type InternalIpIpAddressState struct {
	Address string `json:"address"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
