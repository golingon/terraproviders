// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package chimesdkvoiceglobalsettings

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type VoiceConnector struct {
	// CdrBucket: string, optional
	CdrBucket terra.StringValue `hcl:"cdr_bucket,attr"`
}

type VoiceConnectorAttributes struct {
	ref terra.Reference
}

func (vc VoiceConnectorAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc VoiceConnectorAttributes) InternalWithRef(ref terra.Reference) VoiceConnectorAttributes {
	return VoiceConnectorAttributes{ref: ref}
}

func (vc VoiceConnectorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc VoiceConnectorAttributes) CdrBucket() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("cdr_bucket"))
}

type VoiceConnectorState struct {
	CdrBucket string `json:"cdr_bucket"`
}
