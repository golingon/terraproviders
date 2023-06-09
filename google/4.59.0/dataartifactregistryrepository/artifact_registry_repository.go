// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataartifactregistryrepository

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MavenConfig struct{}

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

type MavenConfigState struct {
	AllowSnapshotOverwrites bool   `json:"allow_snapshot_overwrites"`
	VersionPolicy           string `json:"version_policy"`
}
