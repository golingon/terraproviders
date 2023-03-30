// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package imagebuilderimagerecipe

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BlockDeviceMapping struct {
	// DeviceName: string, optional
	DeviceName terra.StringValue `hcl:"device_name,attr"`
	// NoDevice: bool, optional
	NoDevice terra.BoolValue `hcl:"no_device,attr"`
	// VirtualName: string, optional
	VirtualName terra.StringValue `hcl:"virtual_name,attr"`
	// Ebs: optional
	Ebs *Ebs `hcl:"ebs,block"`
}

type Ebs struct {
	// DeleteOnTermination: string, optional
	DeleteOnTermination terra.StringValue `hcl:"delete_on_termination,attr"`
	// Encrypted: string, optional
	Encrypted terra.StringValue `hcl:"encrypted,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// SnapshotId: string, optional
	SnapshotId terra.StringValue `hcl:"snapshot_id,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type Component struct {
	// ComponentArn: string, required
	ComponentArn terra.StringValue `hcl:"component_arn,attr" validate:"required"`
	// Parameter: min=0
	Parameter []Parameter `hcl:"parameter,block" validate:"min=0"`
}

type Parameter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type SystemsManagerAgent struct {
	// UninstallAfterBuild: bool, required
	UninstallAfterBuild terra.BoolValue `hcl:"uninstall_after_build,attr" validate:"required"`
}

type BlockDeviceMappingAttributes struct {
	ref terra.Reference
}

func (bdm BlockDeviceMappingAttributes) InternalRef() terra.Reference {
	return bdm.ref
}

func (bdm BlockDeviceMappingAttributes) InternalWithRef(ref terra.Reference) BlockDeviceMappingAttributes {
	return BlockDeviceMappingAttributes{ref: ref}
}

func (bdm BlockDeviceMappingAttributes) InternalTokens() hclwrite.Tokens {
	return bdm.ref.InternalTokens()
}

func (bdm BlockDeviceMappingAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceString(bdm.ref.Append("device_name"))
}

func (bdm BlockDeviceMappingAttributes) NoDevice() terra.BoolValue {
	return terra.ReferenceBool(bdm.ref.Append("no_device"))
}

func (bdm BlockDeviceMappingAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceString(bdm.ref.Append("virtual_name"))
}

func (bdm BlockDeviceMappingAttributes) Ebs() terra.ListValue[EbsAttributes] {
	return terra.ReferenceList[EbsAttributes](bdm.ref.Append("ebs"))
}

type EbsAttributes struct {
	ref terra.Reference
}

func (e EbsAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EbsAttributes) InternalWithRef(ref terra.Reference) EbsAttributes {
	return EbsAttributes{ref: ref}
}

func (e EbsAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EbsAttributes) DeleteOnTermination() terra.StringValue {
	return terra.ReferenceString(e.ref.Append("delete_on_termination"))
}

func (e EbsAttributes) Encrypted() terra.StringValue {
	return terra.ReferenceString(e.ref.Append("encrypted"))
}

func (e EbsAttributes) Iops() terra.NumberValue {
	return terra.ReferenceNumber(e.ref.Append("iops"))
}

func (e EbsAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(e.ref.Append("kms_key_id"))
}

func (e EbsAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceString(e.ref.Append("snapshot_id"))
}

func (e EbsAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceNumber(e.ref.Append("throughput"))
}

func (e EbsAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceNumber(e.ref.Append("volume_size"))
}

func (e EbsAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceString(e.ref.Append("volume_type"))
}

type ComponentAttributes struct {
	ref terra.Reference
}

func (c ComponentAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c ComponentAttributes) InternalWithRef(ref terra.Reference) ComponentAttributes {
	return ComponentAttributes{ref: ref}
}

func (c ComponentAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c ComponentAttributes) ComponentArn() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("component_arn"))
}

func (c ComponentAttributes) Parameter() terra.SetValue[ParameterAttributes] {
	return terra.ReferenceSet[ParameterAttributes](c.ref.Append("parameter"))
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("name"))
}

func (p ParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("value"))
}

type SystemsManagerAgentAttributes struct {
	ref terra.Reference
}

func (sma SystemsManagerAgentAttributes) InternalRef() terra.Reference {
	return sma.ref
}

func (sma SystemsManagerAgentAttributes) InternalWithRef(ref terra.Reference) SystemsManagerAgentAttributes {
	return SystemsManagerAgentAttributes{ref: ref}
}

func (sma SystemsManagerAgentAttributes) InternalTokens() hclwrite.Tokens {
	return sma.ref.InternalTokens()
}

func (sma SystemsManagerAgentAttributes) UninstallAfterBuild() terra.BoolValue {
	return terra.ReferenceBool(sma.ref.Append("uninstall_after_build"))
}

type BlockDeviceMappingState struct {
	DeviceName  string     `json:"device_name"`
	NoDevice    bool       `json:"no_device"`
	VirtualName string     `json:"virtual_name"`
	Ebs         []EbsState `json:"ebs"`
}

type EbsState struct {
	DeleteOnTermination string  `json:"delete_on_termination"`
	Encrypted           string  `json:"encrypted"`
	Iops                float64 `json:"iops"`
	KmsKeyId            string  `json:"kms_key_id"`
	SnapshotId          string  `json:"snapshot_id"`
	Throughput          float64 `json:"throughput"`
	VolumeSize          float64 `json:"volume_size"`
	VolumeType          string  `json:"volume_type"`
}

type ComponentState struct {
	ComponentArn string           `json:"component_arn"`
	Parameter    []ParameterState `json:"parameter"`
}

type ParameterState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SystemsManagerAgentState struct {
	UninstallAfterBuild bool `json:"uninstall_after_build"`
}
