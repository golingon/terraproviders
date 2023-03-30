// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalb

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccessLogs struct{}

type SubnetMapping struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AccessLogsAttributes struct {
	ref terra.Reference
}

func (al AccessLogsAttributes) InternalRef() terra.Reference {
	return al.ref
}

func (al AccessLogsAttributes) InternalWithRef(ref terra.Reference) AccessLogsAttributes {
	return AccessLogsAttributes{ref: ref}
}

func (al AccessLogsAttributes) InternalTokens() hclwrite.Tokens {
	return al.ref.InternalTokens()
}

func (al AccessLogsAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(al.ref.Append("bucket"))
}

func (al AccessLogsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(al.ref.Append("enabled"))
}

func (al AccessLogsAttributes) Prefix() terra.StringValue {
	return terra.ReferenceString(al.ref.Append("prefix"))
}

type SubnetMappingAttributes struct {
	ref terra.Reference
}

func (sm SubnetMappingAttributes) InternalRef() terra.Reference {
	return sm.ref
}

func (sm SubnetMappingAttributes) InternalWithRef(ref terra.Reference) SubnetMappingAttributes {
	return SubnetMappingAttributes{ref: ref}
}

func (sm SubnetMappingAttributes) InternalTokens() hclwrite.Tokens {
	return sm.ref.InternalTokens()
}

func (sm SubnetMappingAttributes) AllocationId() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("allocation_id"))
}

func (sm SubnetMappingAttributes) Ipv6Address() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("ipv6_address"))
}

func (sm SubnetMappingAttributes) OutpostId() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("outpost_id"))
}

func (sm SubnetMappingAttributes) PrivateIpv4Address() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("private_ipv4_address"))
}

func (sm SubnetMappingAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("subnet_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type AccessLogsState struct {
	Bucket  string `json:"bucket"`
	Enabled bool   `json:"enabled"`
	Prefix  string `json:"prefix"`
}

type SubnetMappingState struct {
	AllocationId       string `json:"allocation_id"`
	Ipv6Address        string `json:"ipv6_address"`
	OutpostId          string `json:"outpost_id"`
	PrivateIpv4Address string `json:"private_ipv4_address"`
	SubnetId           string `json:"subnet_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
