// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datatemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Postrender struct {
	// BinaryPath: string, required
	BinaryPath terra.StringValue `hcl:"binary_path,attr" validate:"required"`
}

type Set struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type SetSensitive struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type SetString struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type PostrenderAttributes struct {
	ref terra.Reference
}

func (p PostrenderAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PostrenderAttributes) InternalWithRef(ref terra.Reference) PostrenderAttributes {
	return PostrenderAttributes{ref: ref}
}

func (p PostrenderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PostrenderAttributes) BinaryPath() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("binary_path"))
}

type SetAttributes struct {
	ref terra.Reference
}

func (s SetAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SetAttributes) InternalWithRef(ref terra.Reference) SetAttributes {
	return SetAttributes{ref: ref}
}

func (s SetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("type"))
}

func (s SetAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("value"))
}

type SetSensitiveAttributes struct {
	ref terra.Reference
}

func (ss SetSensitiveAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SetSensitiveAttributes) InternalWithRef(ref terra.Reference) SetSensitiveAttributes {
	return SetSensitiveAttributes{ref: ref}
}

func (ss SetSensitiveAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SetSensitiveAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

func (ss SetSensitiveAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("type"))
}

func (ss SetSensitiveAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("value"))
}

type SetStringAttributes struct {
	ref terra.Reference
}

func (ss SetStringAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SetStringAttributes) InternalWithRef(ref terra.Reference) SetStringAttributes {
	return SetStringAttributes{ref: ref}
}

func (ss SetStringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SetStringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

func (ss SetStringAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("value"))
}

type PostrenderState struct {
	BinaryPath string `json:"binary_path"`
}

type SetState struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SetSensitiveState struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SetStringState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
