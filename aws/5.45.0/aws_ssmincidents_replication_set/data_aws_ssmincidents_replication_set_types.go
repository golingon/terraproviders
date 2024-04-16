// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssmincidents_replication_set

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataRegionAttributes struct {
	ref terra.Reference
}

func (r DataRegionAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r DataRegionAttributes) InternalWithRef(ref terra.Reference) DataRegionAttributes {
	return DataRegionAttributes{ref: ref}
}

func (r DataRegionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r DataRegionAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("kms_key_arn"))
}

func (r DataRegionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r DataRegionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status"))
}

func (r DataRegionAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status_message"))
}

type DataRegionState struct {
	KmsKeyArn     string `json:"kms_key_arn"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
}
