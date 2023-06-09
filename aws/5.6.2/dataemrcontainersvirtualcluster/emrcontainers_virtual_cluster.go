// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataemrcontainersvirtualcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ContainerProvider struct {
	// Info: min=0
	Info []Info `hcl:"info,block" validate:"min=0"`
}

type Info struct {
	// EksInfo: min=0
	EksInfo []EksInfo `hcl:"eks_info,block" validate:"min=0"`
}

type EksInfo struct{}

type ContainerProviderAttributes struct {
	ref terra.Reference
}

func (cp ContainerProviderAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ContainerProviderAttributes) InternalWithRef(ref terra.Reference) ContainerProviderAttributes {
	return ContainerProviderAttributes{ref: ref}
}

func (cp ContainerProviderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp ContainerProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("id"))
}

func (cp ContainerProviderAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("type"))
}

func (cp ContainerProviderAttributes) Info() terra.ListValue[InfoAttributes] {
	return terra.ReferenceAsList[InfoAttributes](cp.ref.Append("info"))
}

type InfoAttributes struct {
	ref terra.Reference
}

func (i InfoAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i InfoAttributes) InternalWithRef(ref terra.Reference) InfoAttributes {
	return InfoAttributes{ref: ref}
}

func (i InfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i InfoAttributes) EksInfo() terra.ListValue[EksInfoAttributes] {
	return terra.ReferenceAsList[EksInfoAttributes](i.ref.Append("eks_info"))
}

type EksInfoAttributes struct {
	ref terra.Reference
}

func (ei EksInfoAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei EksInfoAttributes) InternalWithRef(ref terra.Reference) EksInfoAttributes {
	return EksInfoAttributes{ref: ref}
}

func (ei EksInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei EksInfoAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("namespace"))
}

type ContainerProviderState struct {
	Id   string      `json:"id"`
	Type string      `json:"type"`
	Info []InfoState `json:"info"`
}

type InfoState struct {
	EksInfo []EksInfoState `json:"eks_info"`
}

type EksInfoState struct {
	Namespace string `json:"namespace"`
}
