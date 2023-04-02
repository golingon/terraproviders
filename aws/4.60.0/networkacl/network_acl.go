// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkacl

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Egress struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// IcmpCode: number, optional
	IcmpCode terra.NumberValue `hcl:"icmp_code,attr"`
	// IcmpType: number, optional
	IcmpType terra.NumberValue `hcl:"icmp_type,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// RuleNo: number, optional
	RuleNo terra.NumberValue `hcl:"rule_no,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}

type Ingress struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// IcmpCode: number, optional
	IcmpCode terra.NumberValue `hcl:"icmp_code,attr"`
	// IcmpType: number, optional
	IcmpType terra.NumberValue `hcl:"icmp_type,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// RuleNo: number, optional
	RuleNo terra.NumberValue `hcl:"rule_no,attr"`
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

func (e EgressAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EgressAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("action"))
}

func (e EgressAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("cidr_block"))
}

func (e EgressAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("from_port"))
}

func (e EgressAttributes) IcmpCode() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("icmp_code"))
}

func (e EgressAttributes) IcmpType() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("icmp_type"))
}

func (e EgressAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("ipv6_cidr_block"))
}

func (e EgressAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("protocol"))
}

func (e EgressAttributes) RuleNo() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("rule_no"))
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

func (i IngressAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IngressAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("action"))
}

func (i IngressAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("cidr_block"))
}

func (i IngressAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("from_port"))
}

func (i IngressAttributes) IcmpCode() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("icmp_code"))
}

func (i IngressAttributes) IcmpType() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("icmp_type"))
}

func (i IngressAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("ipv6_cidr_block"))
}

func (i IngressAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("protocol"))
}

func (i IngressAttributes) RuleNo() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("rule_no"))
}

func (i IngressAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("to_port"))
}

type EgressState struct {
	Action        string  `json:"action"`
	CidrBlock     string  `json:"cidr_block"`
	FromPort      float64 `json:"from_port"`
	IcmpCode      float64 `json:"icmp_code"`
	IcmpType      float64 `json:"icmp_type"`
	Ipv6CidrBlock string  `json:"ipv6_cidr_block"`
	Protocol      string  `json:"protocol"`
	RuleNo        float64 `json:"rule_no"`
	ToPort        float64 `json:"to_port"`
}

type IngressState struct {
	Action        string  `json:"action"`
	CidrBlock     string  `json:"cidr_block"`
	FromPort      float64 `json:"from_port"`
	IcmpCode      float64 `json:"icmp_code"`
	IcmpType      float64 `json:"icmp_type"`
	Ipv6CidrBlock string  `json:"ipv6_cidr_block"`
	Protocol      string  `json:"protocol"`
	RuleNo        float64 `json:"rule_no"`
	ToPort        float64 `json:"to_port"`
}
