// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_imagebuilder_image_recipe

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataBlockDeviceMappingAttributes struct {
	ref terra.Reference
}

func (bdm DataBlockDeviceMappingAttributes) InternalRef() (terra.Reference, error) {
	return bdm.ref, nil
}

func (bdm DataBlockDeviceMappingAttributes) InternalWithRef(ref terra.Reference) DataBlockDeviceMappingAttributes {
	return DataBlockDeviceMappingAttributes{ref: ref}
}

func (bdm DataBlockDeviceMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bdm.ref.InternalTokens()
}

func (bdm DataBlockDeviceMappingAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("device_name"))
}

func (bdm DataBlockDeviceMappingAttributes) NoDevice() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("no_device"))
}

func (bdm DataBlockDeviceMappingAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("virtual_name"))
}

func (bdm DataBlockDeviceMappingAttributes) Ebs() terra.ListValue[DataBlockDeviceMappingEbsAttributes] {
	return terra.ReferenceAsList[DataBlockDeviceMappingEbsAttributes](bdm.ref.Append("ebs"))
}

type DataBlockDeviceMappingEbsAttributes struct {
	ref terra.Reference
}

func (e DataBlockDeviceMappingEbsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e DataBlockDeviceMappingEbsAttributes) InternalWithRef(ref terra.Reference) DataBlockDeviceMappingEbsAttributes {
	return DataBlockDeviceMappingEbsAttributes{ref: ref}
}

func (e DataBlockDeviceMappingEbsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e DataBlockDeviceMappingEbsAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("delete_on_termination"))
}

func (e DataBlockDeviceMappingEbsAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("encrypted"))
}

func (e DataBlockDeviceMappingEbsAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("iops"))
}

func (e DataBlockDeviceMappingEbsAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("kms_key_id"))
}

func (e DataBlockDeviceMappingEbsAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("snapshot_id"))
}

func (e DataBlockDeviceMappingEbsAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("throughput"))
}

func (e DataBlockDeviceMappingEbsAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("volume_size"))
}

func (e DataBlockDeviceMappingEbsAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("volume_type"))
}

type DataComponentAttributes struct {
	ref terra.Reference
}

func (c DataComponentAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataComponentAttributes) InternalWithRef(ref terra.Reference) DataComponentAttributes {
	return DataComponentAttributes{ref: ref}
}

func (c DataComponentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataComponentAttributes) ComponentArn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("component_arn"))
}

func (c DataComponentAttributes) Parameter() terra.SetValue[DataComponentParameterAttributes] {
	return terra.ReferenceAsSet[DataComponentParameterAttributes](c.ref.Append("parameter"))
}

type DataComponentParameterAttributes struct {
	ref terra.Reference
}

func (p DataComponentParameterAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p DataComponentParameterAttributes) InternalWithRef(ref terra.Reference) DataComponentParameterAttributes {
	return DataComponentParameterAttributes{ref: ref}
}

func (p DataComponentParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p DataComponentParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p DataComponentParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type DataBlockDeviceMappingState struct {
	DeviceName  string                           `json:"device_name"`
	NoDevice    string                           `json:"no_device"`
	VirtualName string                           `json:"virtual_name"`
	Ebs         []DataBlockDeviceMappingEbsState `json:"ebs"`
}

type DataBlockDeviceMappingEbsState struct {
	DeleteOnTermination bool    `json:"delete_on_termination"`
	Encrypted           bool    `json:"encrypted"`
	Iops                float64 `json:"iops"`
	KmsKeyId            string  `json:"kms_key_id"`
	SnapshotId          string  `json:"snapshot_id"`
	Throughput          float64 `json:"throughput"`
	VolumeSize          float64 `json:"volume_size"`
	VolumeType          string  `json:"volume_type"`
}

type DataComponentState struct {
	ComponentArn string                        `json:"component_arn"`
	Parameter    []DataComponentParameterState `json:"parameter"`
}

type DataComponentParameterState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
