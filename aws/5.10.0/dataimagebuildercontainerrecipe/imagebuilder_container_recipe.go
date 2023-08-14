// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataimagebuildercontainerrecipe

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Component struct {
	// Parameter: min=0
	Parameter []Parameter `hcl:"parameter,block" validate:"min=0"`
}

type Parameter struct{}

type InstanceConfiguration struct {
	// BlockDeviceMapping: min=0
	BlockDeviceMapping []BlockDeviceMapping `hcl:"block_device_mapping,block" validate:"min=0"`
}

type BlockDeviceMapping struct {
	// Ebs: min=0
	Ebs []Ebs `hcl:"ebs,block" validate:"min=0"`
}

type Ebs struct{}

type TargetRepository struct{}

type ComponentAttributes struct {
	ref terra.Reference
}

func (c ComponentAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ComponentAttributes) InternalWithRef(ref terra.Reference) ComponentAttributes {
	return ComponentAttributes{ref: ref}
}

func (c ComponentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ComponentAttributes) ComponentArn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("component_arn"))
}

func (c ComponentAttributes) Parameter() terra.SetValue[ParameterAttributes] {
	return terra.ReferenceAsSet[ParameterAttributes](c.ref.Append("parameter"))
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p ParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type InstanceConfigurationAttributes struct {
	ref terra.Reference
}

func (ic InstanceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic InstanceConfigurationAttributes) InternalWithRef(ref terra.Reference) InstanceConfigurationAttributes {
	return InstanceConfigurationAttributes{ref: ref}
}

func (ic InstanceConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic InstanceConfigurationAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("image"))
}

func (ic InstanceConfigurationAttributes) BlockDeviceMapping() terra.SetValue[BlockDeviceMappingAttributes] {
	return terra.ReferenceAsSet[BlockDeviceMappingAttributes](ic.ref.Append("block_device_mapping"))
}

type BlockDeviceMappingAttributes struct {
	ref terra.Reference
}

func (bdm BlockDeviceMappingAttributes) InternalRef() (terra.Reference, error) {
	return bdm.ref, nil
}

func (bdm BlockDeviceMappingAttributes) InternalWithRef(ref terra.Reference) BlockDeviceMappingAttributes {
	return BlockDeviceMappingAttributes{ref: ref}
}

func (bdm BlockDeviceMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bdm.ref.InternalTokens()
}

func (bdm BlockDeviceMappingAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("device_name"))
}

func (bdm BlockDeviceMappingAttributes) NoDevice() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("no_device"))
}

func (bdm BlockDeviceMappingAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("virtual_name"))
}

func (bdm BlockDeviceMappingAttributes) Ebs() terra.ListValue[EbsAttributes] {
	return terra.ReferenceAsList[EbsAttributes](bdm.ref.Append("ebs"))
}

type EbsAttributes struct {
	ref terra.Reference
}

func (e EbsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EbsAttributes) InternalWithRef(ref terra.Reference) EbsAttributes {
	return EbsAttributes{ref: ref}
}

func (e EbsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EbsAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("delete_on_termination"))
}

func (e EbsAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("encrypted"))
}

func (e EbsAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("iops"))
}

func (e EbsAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("kms_key_id"))
}

func (e EbsAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("snapshot_id"))
}

func (e EbsAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("throughput"))
}

func (e EbsAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("volume_size"))
}

func (e EbsAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("volume_type"))
}

type TargetRepositoryAttributes struct {
	ref terra.Reference
}

func (tr TargetRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TargetRepositoryAttributes) InternalWithRef(ref terra.Reference) TargetRepositoryAttributes {
	return TargetRepositoryAttributes{ref: ref}
}

func (tr TargetRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TargetRepositoryAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("repository_name"))
}

func (tr TargetRepositoryAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("service"))
}

type ComponentState struct {
	ComponentArn string           `json:"component_arn"`
	Parameter    []ParameterState `json:"parameter"`
}

type ParameterState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InstanceConfigurationState struct {
	Image              string                    `json:"image"`
	BlockDeviceMapping []BlockDeviceMappingState `json:"block_device_mapping"`
}

type BlockDeviceMappingState struct {
	DeviceName  string     `json:"device_name"`
	NoDevice    string     `json:"no_device"`
	VirtualName string     `json:"virtual_name"`
	Ebs         []EbsState `json:"ebs"`
}

type EbsState struct {
	DeleteOnTermination bool    `json:"delete_on_termination"`
	Encrypted           bool    `json:"encrypted"`
	Iops                float64 `json:"iops"`
	KmsKeyId            string  `json:"kms_key_id"`
	SnapshotId          string  `json:"snapshot_id"`
	Throughput          float64 `json:"throughput"`
	VolumeSize          float64 `json:"volume_size"`
	VolumeType          string  `json:"volume_type"`
}

type TargetRepositoryState struct {
	RepositoryName string `json:"repository_name"`
	Service        string `json:"service"`
}