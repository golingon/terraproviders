// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package beyondcorpappgateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AllocatedConnections struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type AllocatedConnectionsAttributes struct {
	ref terra.Reference
}

func (ac AllocatedConnectionsAttributes) InternalRef() terra.Reference {
	return ac.ref
}

func (ac AllocatedConnectionsAttributes) InternalWithRef(ref terra.Reference) AllocatedConnectionsAttributes {
	return AllocatedConnectionsAttributes{ref: ref}
}

func (ac AllocatedConnectionsAttributes) InternalTokens() hclwrite.Tokens {
	return ac.ref.InternalTokens()
}

func (ac AllocatedConnectionsAttributes) IngressPort() terra.NumberValue {
	return terra.ReferenceNumber(ac.ref.Append("ingress_port"))
}

func (ac AllocatedConnectionsAttributes) PscUri() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("psc_uri"))
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
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

type AllocatedConnectionsState struct {
	IngressPort float64 `json:"ingress_port"`
	PscUri      string  `json:"psc_uri"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
