// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lightsailinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AddOn struct {
	// SnapshotTime: string, required
	SnapshotTime terra.StringValue `hcl:"snapshot_time,attr" validate:"required"`
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type AddOnAttributes struct {
	ref terra.Reference
}

func (ao AddOnAttributes) InternalRef() (terra.Reference, error) {
	return ao.ref, nil
}

func (ao AddOnAttributes) InternalWithRef(ref terra.Reference) AddOnAttributes {
	return AddOnAttributes{ref: ref}
}

func (ao AddOnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ao.ref.InternalTokens()
}

func (ao AddOnAttributes) SnapshotTime() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("snapshot_time"))
}

func (ao AddOnAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("status"))
}

func (ao AddOnAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("type"))
}

type AddOnState struct {
	SnapshotTime string `json:"snapshot_time"`
	Status       string `json:"status"`
	Type         string `json:"type"`
}
