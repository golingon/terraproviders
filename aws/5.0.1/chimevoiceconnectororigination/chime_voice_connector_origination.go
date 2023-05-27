// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package chimevoiceconnectororigination

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Route struct {
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// Weight: number, required
	Weight terra.NumberValue `hcl:"weight,attr" validate:"required"`
}

type RouteAttributes struct {
	ref terra.Reference
}

func (r RouteAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RouteAttributes) InternalWithRef(ref terra.Reference) RouteAttributes {
	return RouteAttributes{ref: ref}
}

func (r RouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RouteAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("host"))
}

func (r RouteAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("port"))
}

func (r RouteAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("priority"))
}

func (r RouteAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("protocol"))
}

func (r RouteAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("weight"))
}

type RouteState struct {
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
	Priority float64 `json:"priority"`
	Protocol string  `json:"protocol"`
	Weight   float64 `json:"weight"`
}
