// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Batching struct {
	// EnableBatching: bool, optional
	EnableBatching terra.BoolValue `hcl:"enable_batching,attr"`
	// SendAfter: string, optional
	SendAfter terra.StringValue `hcl:"send_after,attr"`
}

type BatchingAttributes struct {
	ref terra.Reference
}

func (b BatchingAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BatchingAttributes) InternalWithRef(ref terra.Reference) BatchingAttributes {
	return BatchingAttributes{ref: ref}
}

func (b BatchingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BatchingAttributes) EnableBatching() terra.BoolValue {
	return terra.ReferenceAsBool(b.ref.Append("enable_batching"))
}

func (b BatchingAttributes) SendAfter() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("send_after"))
}

type BatchingState struct {
	EnableBatching bool   `json:"enable_batching"`
	SendAfter      string `json:"send_after"`
}
