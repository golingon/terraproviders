// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_emr_instance_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type EbsConfig struct {
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// Size: number, required
	Size terra.NumberValue `hcl:"size,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// VolumesPerInstance: number, optional
	VolumesPerInstance terra.NumberValue `hcl:"volumes_per_instance,attr"`
}

type EbsConfigAttributes struct {
	ref terra.Reference
}

func (ec EbsConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EbsConfigAttributes) InternalWithRef(ref terra.Reference) EbsConfigAttributes {
	return EbsConfigAttributes{ref: ref}
}

func (ec EbsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EbsConfigAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ec.ref.Append("iops"))
}

func (ec EbsConfigAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(ec.ref.Append("size"))
}

func (ec EbsConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("type"))
}

func (ec EbsConfigAttributes) VolumesPerInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(ec.ref.Append("volumes_per_instance"))
}

type EbsConfigState struct {
	Iops               float64 `json:"iops"`
	Size               float64 `json:"size"`
	Type               string  `json:"type"`
	VolumesPerInstance float64 `json:"volumes_per_instance"`
}
