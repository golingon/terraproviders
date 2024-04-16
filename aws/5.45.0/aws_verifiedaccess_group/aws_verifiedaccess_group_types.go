// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_verifiedaccess_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SseConfiguration struct {
	// CustomerManagedKeyEnabled: bool, optional
	CustomerManagedKeyEnabled terra.BoolValue `hcl:"customer_managed_key_enabled,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
}

type SseConfigurationAttributes struct {
	ref terra.Reference
}

func (sc SseConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SseConfigurationAttributes) InternalWithRef(ref terra.Reference) SseConfigurationAttributes {
	return SseConfigurationAttributes{ref: ref}
}

func (sc SseConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SseConfigurationAttributes) CustomerManagedKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("customer_managed_key_enabled"))
}

func (sc SseConfigurationAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("kms_key_arn"))
}

type SseConfigurationState struct {
	CustomerManagedKeyEnabled bool   `json:"customer_managed_key_enabled"`
	KmsKeyArn                 string `json:"kms_key_arn"`
}
