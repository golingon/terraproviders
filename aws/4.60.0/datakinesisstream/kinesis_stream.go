// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakinesisstream

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type StreamModeDetails struct{}

type StreamModeDetailsAttributes struct {
	ref terra.Reference
}

func (smd StreamModeDetailsAttributes) InternalRef() terra.Reference {
	return smd.ref
}

func (smd StreamModeDetailsAttributes) InternalWithRef(ref terra.Reference) StreamModeDetailsAttributes {
	return StreamModeDetailsAttributes{ref: ref}
}

func (smd StreamModeDetailsAttributes) InternalTokens() hclwrite.Tokens {
	return smd.ref.InternalTokens()
}

func (smd StreamModeDetailsAttributes) StreamMode() terra.StringValue {
	return terra.ReferenceAsString(smd.ref.Append("stream_mode"))
}

type StreamModeDetailsState struct {
	StreamMode string `json:"stream_mode"`
}
