// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datainstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CreditSpecification struct{}

type EbsBlockDevice struct{}

type EnclaveOptions struct{}

type EphemeralBlockDevice struct{}

type MaintenanceOptions struct{}

type MetadataOptions struct{}

type PrivateDnsNameOptions struct{}

type RootBlockDevice struct{}

type Filter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type CreditSpecificationAttributes struct {
	ref terra.Reference
}

func (cs CreditSpecificationAttributes) InternalRef() terra.Reference {
	return cs.ref
}

func (cs CreditSpecificationAttributes) InternalWithRef(ref terra.Reference) CreditSpecificationAttributes {
	return CreditSpecificationAttributes{ref: ref}
}

func (cs CreditSpecificationAttributes) InternalTokens() hclwrite.Tokens {
	return cs.ref.InternalTokens()
}

func (cs CreditSpecificationAttributes) CpuCredits() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("cpu_credits"))
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

func (ebd EbsBlockDeviceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("kms_key_id"))
}

func (ebd EbsBlockDeviceAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("snapshot_id"))
}

func (ebd EbsBlockDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ebd.ref.Append("tags"))
}

func (ebd EbsBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceNumber(ebd.ref.Append("throughput"))
}

func (ebd EbsBlockDeviceAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("volume_id"))
}

func (ebd EbsBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceNumber(ebd.ref.Append("volume_size"))
}

func (ebd EbsBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("volume_type"))
}

type EnclaveOptionsAttributes struct {
	ref terra.Reference
}

func (eo EnclaveOptionsAttributes) InternalRef() terra.Reference {
	return eo.ref
}

func (eo EnclaveOptionsAttributes) InternalWithRef(ref terra.Reference) EnclaveOptionsAttributes {
	return EnclaveOptionsAttributes{ref: ref}
}

func (eo EnclaveOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return eo.ref.InternalTokens()
}

func (eo EnclaveOptionsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(eo.ref.Append("enabled"))
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

func (ebd EphemeralBlockDeviceAttributes) NoDevice() terra.BoolValue {
	return terra.ReferenceBool(ebd.ref.Append("no_device"))
}

func (ebd EphemeralBlockDeviceAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceString(ebd.ref.Append("virtual_name"))
}

type MaintenanceOptionsAttributes struct {
	ref terra.Reference
}

func (mo MaintenanceOptionsAttributes) InternalRef() terra.Reference {
	return mo.ref
}

func (mo MaintenanceOptionsAttributes) InternalWithRef(ref terra.Reference) MaintenanceOptionsAttributes {
	return MaintenanceOptionsAttributes{ref: ref}
}

func (mo MaintenanceOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return mo.ref.InternalTokens()
}

func (mo MaintenanceOptionsAttributes) AutoRecovery() terra.StringValue {
	return terra.ReferenceString(mo.ref.Append("auto_recovery"))
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
	return terra.ReferenceString(mo.ref.Append("http_endpoint"))
}

func (mo MetadataOptionsAttributes) HttpPutResponseHopLimit() terra.NumberValue {
	return terra.ReferenceNumber(mo.ref.Append("http_put_response_hop_limit"))
}

func (mo MetadataOptionsAttributes) HttpTokens() terra.StringValue {
	return terra.ReferenceString(mo.ref.Append("http_tokens"))
}

func (mo MetadataOptionsAttributes) InstanceMetadataTags() terra.StringValue {
	return terra.ReferenceString(mo.ref.Append("instance_metadata_tags"))
}

type PrivateDnsNameOptionsAttributes struct {
	ref terra.Reference
}

func (pdno PrivateDnsNameOptionsAttributes) InternalRef() terra.Reference {
	return pdno.ref
}

func (pdno PrivateDnsNameOptionsAttributes) InternalWithRef(ref terra.Reference) PrivateDnsNameOptionsAttributes {
	return PrivateDnsNameOptionsAttributes{ref: ref}
}

func (pdno PrivateDnsNameOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return pdno.ref.InternalTokens()
}

func (pdno PrivateDnsNameOptionsAttributes) EnableResourceNameDnsARecord() terra.BoolValue {
	return terra.ReferenceBool(pdno.ref.Append("enable_resource_name_dns_a_record"))
}

func (pdno PrivateDnsNameOptionsAttributes) EnableResourceNameDnsAaaaRecord() terra.BoolValue {
	return terra.ReferenceBool(pdno.ref.Append("enable_resource_name_dns_aaaa_record"))
}

func (pdno PrivateDnsNameOptionsAttributes) HostnameType() terra.StringValue {
	return terra.ReferenceString(pdno.ref.Append("hostname_type"))
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
	return terra.ReferenceBool(rbd.ref.Append("delete_on_termination"))
}

func (rbd RootBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceString(rbd.ref.Append("device_name"))
}

func (rbd RootBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(rbd.ref.Append("encrypted"))
}

func (rbd RootBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceNumber(rbd.ref.Append("iops"))
}

func (rbd RootBlockDeviceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(rbd.ref.Append("kms_key_id"))
}

func (rbd RootBlockDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rbd.ref.Append("tags"))
}

func (rbd RootBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceNumber(rbd.ref.Append("throughput"))
}

func (rbd RootBlockDeviceAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceString(rbd.ref.Append("volume_id"))
}

func (rbd RootBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceNumber(rbd.ref.Append("volume_size"))
}

func (rbd RootBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceString(rbd.ref.Append("volume_type"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(f.ref.Append("name"))
}

func (f FilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](f.ref.Append("values"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type CreditSpecificationState struct {
	CpuCredits string `json:"cpu_credits"`
}

type EbsBlockDeviceState struct {
	DeleteOnTermination bool              `json:"delete_on_termination"`
	DeviceName          string            `json:"device_name"`
	Encrypted           bool              `json:"encrypted"`
	Iops                float64           `json:"iops"`
	KmsKeyId            string            `json:"kms_key_id"`
	SnapshotId          string            `json:"snapshot_id"`
	Tags                map[string]string `json:"tags"`
	Throughput          float64           `json:"throughput"`
	VolumeId            string            `json:"volume_id"`
	VolumeSize          float64           `json:"volume_size"`
	VolumeType          string            `json:"volume_type"`
}

type EnclaveOptionsState struct {
	Enabled bool `json:"enabled"`
}

type EphemeralBlockDeviceState struct {
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
	VirtualName string `json:"virtual_name"`
}

type MaintenanceOptionsState struct {
	AutoRecovery string `json:"auto_recovery"`
}

type MetadataOptionsState struct {
	HttpEndpoint            string  `json:"http_endpoint"`
	HttpPutResponseHopLimit float64 `json:"http_put_response_hop_limit"`
	HttpTokens              string  `json:"http_tokens"`
	InstanceMetadataTags    string  `json:"instance_metadata_tags"`
}

type PrivateDnsNameOptionsState struct {
	EnableResourceNameDnsARecord    bool   `json:"enable_resource_name_dns_a_record"`
	EnableResourceNameDnsAaaaRecord bool   `json:"enable_resource_name_dns_aaaa_record"`
	HostnameType                    string `json:"hostname_type"`
}

type RootBlockDeviceState struct {
	DeleteOnTermination bool              `json:"delete_on_termination"`
	DeviceName          string            `json:"device_name"`
	Encrypted           bool              `json:"encrypted"`
	Iops                float64           `json:"iops"`
	KmsKeyId            string            `json:"kms_key_id"`
	Tags                map[string]string `json:"tags"`
	Throughput          float64           `json:"throughput"`
	VolumeId            string            `json:"volume_id"`
	VolumeSize          float64           `json:"volume_size"`
	VolumeType          string            `json:"volume_type"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
