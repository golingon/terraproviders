// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fsxontapfilesystem

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoints struct {
	// Intercluster: min=0
	Intercluster []Intercluster `hcl:"intercluster,block" validate:"min=0"`
	// Management: min=0
	Management []Management `hcl:"management,block" validate:"min=0"`
}

type Intercluster struct{}

type Management struct{}

type DiskIopsConfiguration struct {
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EndpointsAttributes struct {
	ref terra.Reference
}

func (e EndpointsAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EndpointsAttributes) InternalWithRef(ref terra.Reference) EndpointsAttributes {
	return EndpointsAttributes{ref: ref}
}

func (e EndpointsAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EndpointsAttributes) Intercluster() terra.ListValue[InterclusterAttributes] {
	return terra.ReferenceAsList[InterclusterAttributes](e.ref.Append("intercluster"))
}

func (e EndpointsAttributes) Management() terra.ListValue[ManagementAttributes] {
	return terra.ReferenceAsList[ManagementAttributes](e.ref.Append("management"))
}

type InterclusterAttributes struct {
	ref terra.Reference
}

func (i InterclusterAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i InterclusterAttributes) InternalWithRef(ref terra.Reference) InterclusterAttributes {
	return InterclusterAttributes{ref: ref}
}

func (i InterclusterAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i InterclusterAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("dns_name"))
}

func (i InterclusterAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("ip_addresses"))
}

type ManagementAttributes struct {
	ref terra.Reference
}

func (m ManagementAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ManagementAttributes) InternalWithRef(ref terra.Reference) ManagementAttributes {
	return ManagementAttributes{ref: ref}
}

func (m ManagementAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ManagementAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("dns_name"))
}

func (m ManagementAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](m.ref.Append("ip_addresses"))
}

type DiskIopsConfigurationAttributes struct {
	ref terra.Reference
}

func (dic DiskIopsConfigurationAttributes) InternalRef() terra.Reference {
	return dic.ref
}

func (dic DiskIopsConfigurationAttributes) InternalWithRef(ref terra.Reference) DiskIopsConfigurationAttributes {
	return DiskIopsConfigurationAttributes{ref: ref}
}

func (dic DiskIopsConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return dic.ref.InternalTokens()
}

func (dic DiskIopsConfigurationAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(dic.ref.Append("iops"))
}

func (dic DiskIopsConfigurationAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(dic.ref.Append("mode"))
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

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type EndpointsState struct {
	Intercluster []InterclusterState `json:"intercluster"`
	Management   []ManagementState   `json:"management"`
}

type InterclusterState struct {
	DnsName     string   `json:"dns_name"`
	IpAddresses []string `json:"ip_addresses"`
}

type ManagementState struct {
	DnsName     string   `json:"dns_name"`
	IpAddresses []string `json:"ip_addresses"`
}

type DiskIopsConfigurationState struct {
	Iops float64 `json:"iops"`
	Mode string  `json:"mode"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
