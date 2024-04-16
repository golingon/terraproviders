// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_kinesis_stream

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type StreamModeDetails struct {
	// StreamMode: string, required
	StreamMode terra.StringValue `hcl:"stream_mode,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type StreamModeDetailsAttributes struct {
	ref terra.Reference
}

func (smd StreamModeDetailsAttributes) InternalRef() (terra.Reference, error) {
	return smd.ref, nil
}

func (smd StreamModeDetailsAttributes) InternalWithRef(ref terra.Reference) StreamModeDetailsAttributes {
	return StreamModeDetailsAttributes{ref: ref}
}

func (smd StreamModeDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return smd.ref.InternalTokens()
}

func (smd StreamModeDetailsAttributes) StreamMode() terra.StringValue {
	return terra.ReferenceAsString(smd.ref.Append("stream_mode"))
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

type StreamModeDetailsState struct {
	StreamMode string `json:"stream_mode"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
