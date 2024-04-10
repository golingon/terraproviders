// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cdnfrontdoororigin

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PrivateLink struct {
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PrivateLinkTargetId: string, required
	PrivateLinkTargetId terra.StringValue `hcl:"private_link_target_id,attr" validate:"required"`
	// RequestMessage: string, optional
	RequestMessage terra.StringValue `hcl:"request_message,attr"`
	// TargetType: string, optional
	TargetType terra.StringValue `hcl:"target_type,attr"`
}

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

type PrivateLinkAttributes struct {
	ref terra.Reference
}

func (pl PrivateLinkAttributes) InternalRef() (terra.Reference, error) {
	return pl.ref, nil
}

func (pl PrivateLinkAttributes) InternalWithRef(ref terra.Reference) PrivateLinkAttributes {
	return PrivateLinkAttributes{ref: ref}
}

func (pl PrivateLinkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pl.ref.InternalTokens()
}

func (pl PrivateLinkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pl.ref.Append("location"))
}

func (pl PrivateLinkAttributes) PrivateLinkTargetId() terra.StringValue {
	return terra.ReferenceAsString(pl.ref.Append("private_link_target_id"))
}

func (pl PrivateLinkAttributes) RequestMessage() terra.StringValue {
	return terra.ReferenceAsString(pl.ref.Append("request_message"))
}

func (pl PrivateLinkAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(pl.ref.Append("target_type"))
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

type PrivateLinkState struct {
	Location            string `json:"location"`
	PrivateLinkTargetId string `json:"private_link_target_id"`
	RequestMessage      string `json:"request_message"`
	TargetType          string `json:"target_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
