// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package launchconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EbsBlockDevice struct {
	// DeleteOnTermination: bool, optional
	DeleteOnTermination terra.BoolValue `hcl:"delete_on_termination,attr"`
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// NoDevice: bool, optional
	NoDevice terra.BoolValue `hcl:"no_device,attr"`
	// SnapshotId: string, optional
	SnapshotId terra.StringValue `hcl:"snapshot_id,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type EphemeralBlockDevice struct {
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// NoDevice: bool, optional
	NoDevice terra.BoolValue `hcl:"no_device,attr"`
	// VirtualName: string, optional
	VirtualName terra.StringValue `hcl:"virtual_name,attr"`
}

type MetadataOptions struct {
	// HttpEndpoint: string, optional
	HttpEndpoint terra.StringValue `hcl:"http_endpoint,attr"`
	// HttpPutResponseHopLimit: number, optional
	HttpPutResponseHopLimit terra.NumberValue `hcl:"http_put_response_hop_limit,attr"`
	// HttpTokens: string, optional
	HttpTokens terra.StringValue `hcl:"http_tokens,attr"`
}

type RootBlockDevice struct {
	// DeleteOnTermination: bool, optional
	DeleteOnTermination terra.BoolValue `hcl:"delete_on_termination,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
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
	return terra.ReferenceAsBool(ebd.ref.Append("delete_on_termination"))
}

func (ebd EbsBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("device_name"))
}

func (ebd EbsBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ebd.ref.Append("encrypted"))
}

func (ebd EbsBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("iops"))
}

func (ebd EbsBlockDeviceAttributes) NoDevice() terra.BoolValue {
	return terra.ReferenceAsBool(ebd.ref.Append("no_device"))
}

func (ebd EbsBlockDeviceAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("snapshot_id"))
}

func (ebd EbsBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("throughput"))
}

func (ebd EbsBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("volume_size"))
}

func (ebd EbsBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("volume_type"))
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
	return terra.ReferenceAsString(ebd.ref.Append("device_name"))
}

func (ebd EphemeralBlockDeviceAttributes) NoDevice() terra.BoolValue {
	return terra.ReferenceAsBool(ebd.ref.Append("no_device"))
}

func (ebd EphemeralBlockDeviceAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("virtual_name"))
}

type MetadataOptionsAttributes struct {
	ref terra.Reference
}

func (mo MetadataOptionsAttributes) InternalRef() terra.Reference {
	return mo.ref
}

func (mo MetadataOptionsAttributes) InternalWithRef(ref terra.Reference) MetadataOptionsAttributes {
	return MetadataOptionsAttributes{ref: ref}
}

func (mo MetadataOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return mo.ref.InternalTokens()
}

func (mo MetadataOptionsAttributes) HttpEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("http_endpoint"))
}

func (mo MetadataOptionsAttributes) HttpPutResponseHopLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(mo.ref.Append("http_put_response_hop_limit"))
}

func (mo MetadataOptionsAttributes) HttpTokens() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("http_tokens"))
}

type RootBlockDeviceAttributes struct {
	ref terra.Reference
}

func (rbd RootBlockDeviceAttributes) InternalRef() terra.Reference {
	return rbd.ref
}

func (rbd RootBlockDeviceAttributes) InternalWithRef(ref terra.Reference) RootBlockDeviceAttributes {
	return RootBlockDeviceAttributes{ref: ref}
}

func (rbd RootBlockDeviceAttributes) InternalTokens() hclwrite.Tokens {
	return rbd.ref.InternalTokens()
}

func (rbd RootBlockDeviceAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(rbd.ref.Append("delete_on_termination"))
}

func (rbd RootBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rbd.ref.Append("encrypted"))
}

func (rbd RootBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("iops"))
}

func (rbd RootBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("throughput"))
}

func (rbd RootBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("volume_size"))
}

func (rbd RootBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("volume_type"))
}

type EbsBlockDeviceState struct {
	DeleteOnTermination bool    `json:"delete_on_termination"`
	DeviceName          string  `json:"device_name"`
	Encrypted           bool    `json:"encrypted"`
	Iops                float64 `json:"iops"`
	NoDevice            bool    `json:"no_device"`
	SnapshotId          string  `json:"snapshot_id"`
	Throughput          float64 `json:"throughput"`
	VolumeSize          float64 `json:"volume_size"`
	VolumeType          string  `json:"volume_type"`
}

type EphemeralBlockDeviceState struct {
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
	VirtualName string `json:"virtual_name"`
}

type MetadataOptionsState struct {
	HttpEndpoint            string  `json:"http_endpoint"`
	HttpPutResponseHopLimit float64 `json:"http_put_response_hop_limit"`
	HttpTokens              string  `json:"http_tokens"`
}

type RootBlockDeviceState struct {
	DeleteOnTermination bool    `json:"delete_on_termination"`
	Encrypted           bool    `json:"encrypted"`
	Iops                float64 `json:"iops"`
	Throughput          float64 `json:"throughput"`
	VolumeSize          float64 `json:"volume_size"`
	VolumeType          string  `json:"volume_type"`
}
