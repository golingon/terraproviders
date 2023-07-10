// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataimagebuilderimage

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ImageTestsConfiguration struct{}

type OutputResources struct {
	// Amis: min=0
	Amis []Amis `hcl:"amis,block" validate:"min=0"`
	// Containers: min=0
	Containers []Containers `hcl:"containers,block" validate:"min=0"`
}

type Amis struct{}

type Containers struct{}

type ImageTestsConfigurationAttributes struct {
	ref terra.Reference
}

func (itc ImageTestsConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return itc.ref, nil
}

func (itc ImageTestsConfigurationAttributes) InternalWithRef(ref terra.Reference) ImageTestsConfigurationAttributes {
	return ImageTestsConfigurationAttributes{ref: ref}
}

func (itc ImageTestsConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return itc.ref.InternalTokens()
}

func (itc ImageTestsConfigurationAttributes) ImageTestsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(itc.ref.Append("image_tests_enabled"))
}

func (itc ImageTestsConfigurationAttributes) TimeoutMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(itc.ref.Append("timeout_minutes"))
}

type OutputResourcesAttributes struct {
	ref terra.Reference
}

func (or OutputResourcesAttributes) InternalRef() (terra.Reference, error) {
	return or.ref, nil
}

func (or OutputResourcesAttributes) InternalWithRef(ref terra.Reference) OutputResourcesAttributes {
	return OutputResourcesAttributes{ref: ref}
}

func (or OutputResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return or.ref.InternalTokens()
}

func (or OutputResourcesAttributes) Amis() terra.SetValue[AmisAttributes] {
	return terra.ReferenceAsSet[AmisAttributes](or.ref.Append("amis"))
}

func (or OutputResourcesAttributes) Containers() terra.SetValue[ContainersAttributes] {
	return terra.ReferenceAsSet[ContainersAttributes](or.ref.Append("containers"))
}

type AmisAttributes struct {
	ref terra.Reference
}

func (a AmisAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AmisAttributes) InternalWithRef(ref terra.Reference) AmisAttributes {
	return AmisAttributes{ref: ref}
}

func (a AmisAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AmisAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("account_id"))
}

func (a AmisAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("description"))
}

func (a AmisAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image"))
}

func (a AmisAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a AmisAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("region"))
}

type ContainersAttributes struct {
	ref terra.Reference
}

func (c ContainersAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContainersAttributes) InternalWithRef(ref terra.Reference) ContainersAttributes {
	return ContainersAttributes{ref: ref}
}

func (c ContainersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContainersAttributes) ImageUris() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("image_uris"))
}

func (c ContainersAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("region"))
}

type ImageTestsConfigurationState struct {
	ImageTestsEnabled bool    `json:"image_tests_enabled"`
	TimeoutMinutes    float64 `json:"timeout_minutes"`
}

type OutputResourcesState struct {
	Amis       []AmisState       `json:"amis"`
	Containers []ContainersState `json:"containers"`
}

type AmisState struct {
	AccountId   string `json:"account_id"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Region      string `json:"region"`
}

type ContainersState struct {
	ImageUris []string `json:"image_uris"`
	Region    string   `json:"region"`
}
