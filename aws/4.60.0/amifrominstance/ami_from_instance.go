// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package amifrominstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EbsBlockDevice struct{}

type EphemeralBlockDevice struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EbsBlockDeviceAttributes struct {
	ref terra.Reference
}

func (ebd EbsBlockDeviceAttributes) InternalRef() terra.Reference {
	return ebd.ref
}

func (ebd EbsBlockDeviceAttributes) InternalWithRef(ref terra.Reference) EbsBlockDeviceAttributes {
	return EbsBlockDeviceAttributes{ref: ref}
}

func (ebd EbsBlockDeviceAttributes) InternalTokens() hclwrite.Tokens {
	return ebd.ref.InternalTokens()
}

func (ebd EbsBlockDeviceAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceBool(ebd.ref.Append("delete_on_termination"))
}

func (ebd EbsBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("device_name"))
}

func (ebd EbsBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(ebd.ref.Append("encrypted"))
}

func (ebd EbsBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceNumber(ebd.ref.Append("iops"))
}

func (ebd EbsBlockDeviceAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("outpost_arn"))
}

func (ebd EbsBlockDeviceAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("snapshot_id"))
}

func (ebd EbsBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceNumber(ebd.ref.Append("throughput"))
}

func (ebd EbsBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceNumber(ebd.ref.Append("volume_size"))
}

func (ebd EbsBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("volume_type"))
}

type EphemeralBlockDeviceAttributes struct {
	ref terra.Reference
}

func (ebd EphemeralBlockDeviceAttributes) InternalRef() terra.Reference {
	return ebd.ref
}

func (ebd EphemeralBlockDeviceAttributes) InternalWithRef(ref terra.Reference) EphemeralBlockDeviceAttributes {
	return EphemeralBlockDeviceAttributes{ref: ref}
}

func (ebd EphemeralBlockDeviceAttributes) InternalTokens() hclwrite.Tokens {
	return ebd.ref.InternalTokens()
}

func (ebd EphemeralBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("device_name"))
}

func (ebd EphemeralBlockDeviceAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("virtual_name"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type EbsBlockDeviceState struct {
	DeleteOnTermination bool    `json:"delete_on_termination"`
	DeviceName          string  `json:"device_name"`
	Encrypted           bool    `json:"encrypted"`
	Iops                float64 `json:"iops"`
	OutpostArn          string  `json:"outpost_arn"`
	SnapshotId          string  `json:"snapshot_id"`
	Throughput          float64 `json:"throughput"`
	VolumeSize          float64 `json:"volume_size"`
	VolumeType          string  `json:"volume_type"`
}

type EphemeralBlockDeviceState struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
