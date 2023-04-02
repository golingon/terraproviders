// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerdevicefleet

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type OutputConfig struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// S3OutputLocation: string, required
	S3OutputLocation terra.StringValue `hcl:"s3_output_location,attr" validate:"required"`
}

type OutputConfigAttributes struct {
	ref terra.Reference
}

func (oc OutputConfigAttributes) InternalRef() (terra.Reference, error) {
	return oc.ref, nil
}

func (oc OutputConfigAttributes) InternalWithRef(ref terra.Reference) OutputConfigAttributes {
	return OutputConfigAttributes{ref: ref}
}

func (oc OutputConfigAttributes) InternalTokens() hclwrite.Tokens {
	return oc.ref.InternalTokens()
}

func (oc OutputConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("kms_key_id"))
}

func (oc OutputConfigAttributes) S3OutputLocation() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("s3_output_location"))
}

type OutputConfigState struct {
	KmsKeyId         string `json:"kms_key_id"`
	S3OutputLocation string `json:"s3_output_location"`
}
