// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codepipeline

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ArtifactStore struct {
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// EncryptionKey: optional
	EncryptionKey *EncryptionKey `hcl:"encryption_key,block"`
}

type EncryptionKey struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Stage struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Action: min=1
	Action []Action `hcl:"action,block" validate:"min=1"`
}

type Action struct {
	// Category: string, required
	Category terra.StringValue `hcl:"category,attr" validate:"required"`
	// Configuration: map of string, optional
	Configuration terra.MapValue[terra.StringValue] `hcl:"configuration,attr"`
	// InputArtifacts: list of string, optional
	InputArtifacts terra.ListValue[terra.StringValue] `hcl:"input_artifacts,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// OutputArtifacts: list of string, optional
	OutputArtifacts terra.ListValue[terra.StringValue] `hcl:"output_artifacts,attr"`
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
	// Provider: string, required
	Provider terra.StringValue `hcl:"provider,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// RunOrder: number, optional
	RunOrder terra.NumberValue `hcl:"run_order,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type ArtifactStoreAttributes struct {
	ref terra.Reference
}

func (as ArtifactStoreAttributes) InternalRef() terra.Reference {
	return as.ref
}

func (as ArtifactStoreAttributes) InternalWithRef(ref terra.Reference) ArtifactStoreAttributes {
	return ArtifactStoreAttributes{ref: ref}
}

func (as ArtifactStoreAttributes) InternalTokens() hclwrite.Tokens {
	return as.ref.InternalTokens()
}

func (as ArtifactStoreAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("location"))
}

func (as ArtifactStoreAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("region"))
}

func (as ArtifactStoreAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("type"))
}

func (as ArtifactStoreAttributes) EncryptionKey() terra.ListValue[EncryptionKeyAttributes] {
	return terra.ReferenceAsList[EncryptionKeyAttributes](as.ref.Append("encryption_key"))
}

type EncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ek EncryptionKeyAttributes) InternalRef() terra.Reference {
	return ek.ref
}

func (ek EncryptionKeyAttributes) InternalWithRef(ref terra.Reference) EncryptionKeyAttributes {
	return EncryptionKeyAttributes{ref: ref}
}

func (ek EncryptionKeyAttributes) InternalTokens() hclwrite.Tokens {
	return ek.ref.InternalTokens()
}

func (ek EncryptionKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ek.ref.Append("id"))
}

func (ek EncryptionKeyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ek.ref.Append("type"))
}

type StageAttributes struct {
	ref terra.Reference
}

func (s StageAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s StageAttributes) InternalWithRef(ref terra.Reference) StageAttributes {
	return StageAttributes{ref: ref}
}

func (s StageAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s StageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s StageAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](s.ref.Append("action"))
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("category"))
}

func (a ActionAttributes) Configuration() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("configuration"))
}

func (a ActionAttributes) InputArtifacts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("input_artifacts"))
}

func (a ActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a ActionAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("namespace"))
}

func (a ActionAttributes) OutputArtifacts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("output_artifacts"))
}

func (a ActionAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("owner"))
}

func (a ActionAttributes) Provider() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("provider"))
}

func (a ActionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("region"))
}

func (a ActionAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("role_arn"))
}

func (a ActionAttributes) RunOrder() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("run_order"))
}

func (a ActionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("version"))
}

type ArtifactStoreState struct {
	Location      string               `json:"location"`
	Region        string               `json:"region"`
	Type          string               `json:"type"`
	EncryptionKey []EncryptionKeyState `json:"encryption_key"`
}

type EncryptionKeyState struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type StageState struct {
	Name   string        `json:"name"`
	Action []ActionState `json:"action"`
}

type ActionState struct {
	Category        string            `json:"category"`
	Configuration   map[string]string `json:"configuration"`
	InputArtifacts  []string          `json:"input_artifacts"`
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	OutputArtifacts []string          `json:"output_artifacts"`
	Owner           string            `json:"owner"`
	Provider        string            `json:"provider"`
	Region          string            `json:"region"`
	RoleArn         string            `json:"role_arn"`
	RunOrder        float64           `json:"run_order"`
	Version         string            `json:"version"`
}
