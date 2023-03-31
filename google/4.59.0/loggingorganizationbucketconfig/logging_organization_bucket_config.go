// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package loggingorganizationbucketconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CmekSettings struct {
	// KmsKeyName: string, required
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr" validate:"required"`
}

type CmekSettingsAttributes struct {
	ref terra.Reference
}

func (cs CmekSettingsAttributes) InternalRef() terra.Reference {
	return cs.ref
}

func (cs CmekSettingsAttributes) InternalWithRef(ref terra.Reference) CmekSettingsAttributes {
	return CmekSettingsAttributes{ref: ref}
}

func (cs CmekSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return cs.ref.InternalTokens()
}

func (cs CmekSettingsAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("kms_key_name"))
}

func (cs CmekSettingsAttributes) KmsKeyVersionName() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("kms_key_version_name"))
}

func (cs CmekSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("name"))
}

func (cs CmekSettingsAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("service_account_id"))
}

type CmekSettingsState struct {
	KmsKeyName        string `json:"kms_key_name"`
	KmsKeyVersionName string `json:"kms_key_version_name"`
	Name              string `json:"name"`
	ServiceAccountId  string `json:"service_account_id"`
}
