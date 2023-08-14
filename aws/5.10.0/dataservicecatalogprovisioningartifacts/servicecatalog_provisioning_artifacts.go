// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataservicecatalogprovisioningartifacts

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ProvisioningArtifactDetails struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ProvisioningArtifactDetailsAttributes struct {
	ref terra.Reference
}

func (pad ProvisioningArtifactDetailsAttributes) InternalRef() (terra.Reference, error) {
	return pad.ref, nil
}

func (pad ProvisioningArtifactDetailsAttributes) InternalWithRef(ref terra.Reference) ProvisioningArtifactDetailsAttributes {
	return ProvisioningArtifactDetailsAttributes{ref: ref}
}

func (pad ProvisioningArtifactDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pad.ref.InternalTokens()
}

func (pad ProvisioningArtifactDetailsAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(pad.ref.Append("active"))
}

func (pad ProvisioningArtifactDetailsAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(pad.ref.Append("created_time"))
}

func (pad ProvisioningArtifactDetailsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pad.ref.Append("description"))
}

func (pad ProvisioningArtifactDetailsAttributes) Guidance() terra.StringValue {
	return terra.ReferenceAsString(pad.ref.Append("guidance"))
}

func (pad ProvisioningArtifactDetailsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pad.ref.Append("id"))
}

func (pad ProvisioningArtifactDetailsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pad.ref.Append("name"))
}

func (pad ProvisioningArtifactDetailsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pad.ref.Append("type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ProvisioningArtifactDetailsState struct {
	Active      bool   `json:"active"`
	CreatedTime string `json:"created_time"`
	Description string `json:"description"`
	Guidance    string `json:"guidance"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}