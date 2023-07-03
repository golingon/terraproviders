// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package healthcareworkspace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PrivateEndpointConnection struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PrivateEndpointConnectionAttributes struct {
	ref terra.Reference
}

func (pec PrivateEndpointConnectionAttributes) InternalRef() (terra.Reference, error) {
	return pec.ref, nil
}

func (pec PrivateEndpointConnectionAttributes) InternalWithRef(ref terra.Reference) PrivateEndpointConnectionAttributes {
	return PrivateEndpointConnectionAttributes{ref: ref}
}

func (pec PrivateEndpointConnectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pec.ref.InternalTokens()
}

func (pec PrivateEndpointConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("id"))
}

func (pec PrivateEndpointConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("name"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type PrivateEndpointConnectionState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
