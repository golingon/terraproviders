// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package defaultsecuritygroup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Egress struct {
	// CidrBlocks: list of string, optional
	CidrBlocks terra.ListValue[terra.StringValue] `hcl:"cidr_blocks,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// Ipv6CidrBlocks: list of string, optional
	Ipv6CidrBlocks terra.ListValue[terra.StringValue] `hcl:"ipv6_cidr_blocks,attr"`
	// PrefixListIds: list of string, optional
	PrefixListIds terra.ListValue[terra.StringValue] `hcl:"prefix_list_ids,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// Self: bool, optional
	Self terra.BoolValue `hcl:"self,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type Ingress struct {
	// CidrBlocks: list of string, optional
	CidrBlocks terra.ListValue[terra.StringValue] `hcl:"cidr_blocks,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// Ipv6CidrBlocks: list of string, optional
	Ipv6CidrBlocks terra.ListValue[terra.StringValue] `hcl:"ipv6_cidr_blocks,attr"`
	// PrefixListIds: list of string, optional
	PrefixListIds terra.ListValue[terra.StringValue] `hcl:"prefix_list_ids,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// Self: bool, optional
	Self terra.BoolValue `hcl:"self,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type EgressAttributes struct {
	ref terra.Reference
}

func (e EgressAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EgressAttributes) InternalWithRef(ref terra.Reference) EgressAttributes {
	return EgressAttributes{ref: ref}
}

func (e EgressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EgressAttributes) CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("cidr_blocks"))
}

func (e EgressAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("description"))
}

func (e EgressAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("from_port"))
}

func (e EgressAttributes) Ipv6CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("ipv6_cidr_blocks"))
}

func (e EgressAttributes) PrefixListIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("prefix_list_ids"))
}

func (e EgressAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("protocol"))
}

func (e EgressAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("security_groups"))
}

func (e EgressAttributes) Self() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("self"))
}

func (e EgressAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("to_port"))
}

type IngressAttributes struct {
	ref terra.Reference
}

func (i IngressAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IngressAttributes) InternalWithRef(ref terra.Reference) IngressAttributes {
	return IngressAttributes{ref: ref}
}

func (i IngressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IngressAttributes) CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("cidr_blocks"))
}

func (i IngressAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("description"))
}

func (i IngressAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("from_port"))
}

func (i IngressAttributes) Ipv6CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("ipv6_cidr_blocks"))
}

func (i IngressAttributes) PrefixListIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("prefix_list_ids"))
}

func (i IngressAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("protocol"))
}

func (i IngressAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("security_groups"))
}

func (i IngressAttributes) Self() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("self"))
}

func (i IngressAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("to_port"))
}

type EgressState struct {
	CidrBlocks     []string `json:"cidr_blocks"`
	Description    string   `json:"description"`
	FromPort       float64  `json:"from_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	Protocol       string   `json:"protocol"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	ToPort         float64  `json:"to_port"`
}

type IngressState struct {
	CidrBlocks     []string `json:"cidr_blocks"`
	Description    string   `json:"description"`
	FromPort       float64  `json:"from_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	Protocol       string   `json:"protocol"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	ToPort         float64  `json:"to_port"`
}
