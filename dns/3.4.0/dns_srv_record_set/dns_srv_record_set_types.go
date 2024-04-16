// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dns_srv_record_set

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Srv struct {
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
	// Weight: number, required
	Weight terra.NumberValue `hcl:"weight,attr" validate:"required"`
}

type SrvAttributes struct {
	ref terra.Reference
}

func (s SrvAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SrvAttributes) InternalWithRef(ref terra.Reference) SrvAttributes {
	return SrvAttributes{ref: ref}
}

func (s SrvAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SrvAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("port"))
}

func (s SrvAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("priority"))
}

func (s SrvAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("target"))
}

func (s SrvAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("weight"))
}

type SrvState struct {
	Port     float64 `json:"port"`
	Priority float64 `json:"priority"`
	Target   string  `json:"target"`
	Weight   float64 `json:"weight"`
}
