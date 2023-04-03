// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package artifactregistryrepository

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MavenConfig struct {
	// AllowSnapshotOverwrites: bool, optional
	AllowSnapshotOverwrites terra.BoolValue `hcl:"allow_snapshot_overwrites,attr"`
	// VersionPolicy: string, optional
	VersionPolicy terra.StringValue `hcl:"version_policy,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MavenConfigAttributes struct {
	ref terra.Reference
}

func (mc MavenConfigAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MavenConfigAttributes) InternalWithRef(ref terra.Reference) MavenConfigAttributes {
	return MavenConfigAttributes{ref: ref}
}

func (mc MavenConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MavenConfigAttributes) AllowSnapshotOverwrites() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("allow_snapshot_overwrites"))
}

func (mc MavenConfigAttributes) VersionPolicy() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("version_policy"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type MavenConfigState struct {
	AllowSnapshotOverwrites bool   `json:"allow_snapshot_overwrites"`
	VersionPolicy           string `json:"version_policy"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
