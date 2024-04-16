// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_datasync_location_nfs

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MountOptions struct {
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type OnPremConfig struct {
	// AgentArns: set of string, required
	AgentArns terra.SetValue[terra.StringValue] `hcl:"agent_arns,attr" validate:"required"`
}

type MountOptionsAttributes struct {
	ref terra.Reference
}

func (mo MountOptionsAttributes) InternalRef() (terra.Reference, error) {
	return mo.ref, nil
}

func (mo MountOptionsAttributes) InternalWithRef(ref terra.Reference) MountOptionsAttributes {
	return MountOptionsAttributes{ref: ref}
}

func (mo MountOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mo.ref.InternalTokens()
}

func (mo MountOptionsAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("version"))
}

type OnPremConfigAttributes struct {
	ref terra.Reference
}

func (opc OnPremConfigAttributes) InternalRef() (terra.Reference, error) {
	return opc.ref, nil
}

func (opc OnPremConfigAttributes) InternalWithRef(ref terra.Reference) OnPremConfigAttributes {
	return OnPremConfigAttributes{ref: ref}
}

func (opc OnPremConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return opc.ref.InternalTokens()
}

func (opc OnPremConfigAttributes) AgentArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](opc.ref.Append("agent_arns"))
}

type MountOptionsState struct {
	Version string `json:"version"`
}

type OnPremConfigState struct {
	AgentArns []string `json:"agent_arns"`
}
