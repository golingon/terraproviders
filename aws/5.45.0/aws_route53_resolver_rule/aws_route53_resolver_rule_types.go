// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53_resolver_rule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type TargetIp struct {
	// Ip: string, required
	Ip terra.StringValue `hcl:"ip,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TargetIpAttributes struct {
	ref terra.Reference
}

func (ti TargetIpAttributes) InternalRef() (terra.Reference, error) {
	return ti.ref, nil
}

func (ti TargetIpAttributes) InternalWithRef(ref terra.Reference) TargetIpAttributes {
	return TargetIpAttributes{ref: ref}
}

func (ti TargetIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ti.ref.InternalTokens()
}

func (ti TargetIpAttributes) Ip() terra.StringValue {
	return terra.ReferenceAsString(ti.ref.Append("ip"))
}

func (ti TargetIpAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ti.ref.Append("port"))
}

func (ti TargetIpAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ti.ref.Append("protocol"))
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

type TargetIpState struct {
	Ip       string  `json:"ip"`
	Port     float64 `json:"port"`
	Protocol string  `json:"protocol"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
