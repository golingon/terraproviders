// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkmanagersite

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Location struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// Latitude: string, optional
	Latitude terra.StringValue `hcl:"latitude,attr"`
	// Longitude: string, optional
	Longitude terra.StringValue `hcl:"longitude,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type LocationAttributes struct {
	ref terra.Reference
}

func (l LocationAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LocationAttributes) InternalWithRef(ref terra.Reference) LocationAttributes {
	return LocationAttributes{ref: ref}
}

func (l LocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LocationAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("address"))
}

func (l LocationAttributes) Latitude() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("latitude"))
}

func (l LocationAttributes) Longitude() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("longitude"))
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

type LocationState struct {
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}