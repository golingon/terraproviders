// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iotprovisioningtemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PreProvisioningHook struct {
	// PayloadVersion: string, optional
	PayloadVersion terra.StringValue `hcl:"payload_version,attr"`
	// TargetArn: string, required
	TargetArn terra.StringValue `hcl:"target_arn,attr" validate:"required"`
}

type PreProvisioningHookAttributes struct {
	ref terra.Reference
}

func (pph PreProvisioningHookAttributes) InternalRef() terra.Reference {
	return pph.ref
}

func (pph PreProvisioningHookAttributes) InternalWithRef(ref terra.Reference) PreProvisioningHookAttributes {
	return PreProvisioningHookAttributes{ref: ref}
}

func (pph PreProvisioningHookAttributes) InternalTokens() hclwrite.Tokens {
	return pph.ref.InternalTokens()
}

func (pph PreProvisioningHookAttributes) PayloadVersion() terra.StringValue {
	return terra.ReferenceAsString(pph.ref.Append("payload_version"))
}

func (pph PreProvisioningHookAttributes) TargetArn() terra.StringValue {
	return terra.ReferenceAsString(pph.ref.Append("target_arn"))
}

type PreProvisioningHookState struct {
	PayloadVersion string `json:"payload_version"`
	TargetArn      string `json:"target_arn"`
}