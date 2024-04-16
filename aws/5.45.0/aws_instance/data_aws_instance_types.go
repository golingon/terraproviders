// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_instance

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFilter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataCreditSpecificationAttributes struct {
	ref terra.Reference
}

func (cs DataCreditSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs DataCreditSpecificationAttributes) InternalWithRef(ref terra.Reference) DataCreditSpecificationAttributes {
	return DataCreditSpecificationAttributes{ref: ref}
}

func (cs DataCreditSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs DataCreditSpecificationAttributes) CpuCredits() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("cpu_credits"))
}

type DataEbsBlockDeviceAttributes struct {
	ref terra.Reference
}

func (ebd DataEbsBlockDeviceAttributes) InternalRef() (terra.Reference, error) {
	return ebd.ref, nil
}

func (ebd DataEbsBlockDeviceAttributes) InternalWithRef(ref terra.Reference) DataEbsBlockDeviceAttributes {
	return DataEbsBlockDeviceAttributes{ref: ref}
}

func (ebd DataEbsBlockDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ebd.ref.InternalTokens()
}

func (ebd DataEbsBlockDeviceAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(ebd.ref.Append("delete_on_termination"))
}

func (ebd DataEbsBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("device_name"))
}

func (ebd DataEbsBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ebd.ref.Append("encrypted"))
}

func (ebd DataEbsBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("iops"))
}

func (ebd DataEbsBlockDeviceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("kms_key_id"))
}

func (ebd DataEbsBlockDeviceAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("snapshot_id"))
}

func (ebd DataEbsBlockDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ebd.ref.Append("tags"))
}

func (ebd DataEbsBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("throughput"))
}

func (ebd DataEbsBlockDeviceAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("volume_id"))
}

func (ebd DataEbsBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("volume_size"))
}

func (ebd DataEbsBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("volume_type"))
}

type DataEnclaveOptionsAttributes struct {
	ref terra.Reference
}

func (eo DataEnclaveOptionsAttributes) InternalRef() (terra.Reference, error) {
	return eo.ref, nil
}

func (eo DataEnclaveOptionsAttributes) InternalWithRef(ref terra.Reference) DataEnclaveOptionsAttributes {
	return DataEnclaveOptionsAttributes{ref: ref}
}

func (eo DataEnclaveOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eo.ref.InternalTokens()
}

func (eo DataEnclaveOptionsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(eo.ref.Append("enabled"))
}

type DataEphemeralBlockDeviceAttributes struct {
	ref terra.Reference
}

func (ebd DataEphemeralBlockDeviceAttributes) InternalRef() (terra.Reference, error) {
	return ebd.ref, nil
}

func (ebd DataEphemeralBlockDeviceAttributes) InternalWithRef(ref terra.Reference) DataEphemeralBlockDeviceAttributes {
	return DataEphemeralBlockDeviceAttributes{ref: ref}
}

func (ebd DataEphemeralBlockDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ebd.ref.InternalTokens()
}

func (ebd DataEphemeralBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("device_name"))
}

func (ebd DataEphemeralBlockDeviceAttributes) NoDevice() terra.BoolValue {
	return terra.ReferenceAsBool(ebd.ref.Append("no_device"))
}

func (ebd DataEphemeralBlockDeviceAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("virtual_name"))
}

type DataMaintenanceOptionsAttributes struct {
	ref terra.Reference
}

func (mo DataMaintenanceOptionsAttributes) InternalRef() (terra.Reference, error) {
	return mo.ref, nil
}

func (mo DataMaintenanceOptionsAttributes) InternalWithRef(ref terra.Reference) DataMaintenanceOptionsAttributes {
	return DataMaintenanceOptionsAttributes{ref: ref}
}

func (mo DataMaintenanceOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mo.ref.InternalTokens()
}

func (mo DataMaintenanceOptionsAttributes) AutoRecovery() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("auto_recovery"))
}

type DataMetadataOptionsAttributes struct {
	ref terra.Reference
}

func (mo DataMetadataOptionsAttributes) InternalRef() (terra.Reference, error) {
	return mo.ref, nil
}

func (mo DataMetadataOptionsAttributes) InternalWithRef(ref terra.Reference) DataMetadataOptionsAttributes {
	return DataMetadataOptionsAttributes{ref: ref}
}

func (mo DataMetadataOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mo.ref.InternalTokens()
}

func (mo DataMetadataOptionsAttributes) HttpEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("http_endpoint"))
}

func (mo DataMetadataOptionsAttributes) HttpProtocolIpv6() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("http_protocol_ipv6"))
}

func (mo DataMetadataOptionsAttributes) HttpPutResponseHopLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(mo.ref.Append("http_put_response_hop_limit"))
}

func (mo DataMetadataOptionsAttributes) HttpTokens() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("http_tokens"))
}

func (mo DataMetadataOptionsAttributes) InstanceMetadataTags() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("instance_metadata_tags"))
}

type DataPrivateDnsNameOptionsAttributes struct {
	ref terra.Reference
}

func (pdno DataPrivateDnsNameOptionsAttributes) InternalRef() (terra.Reference, error) {
	return pdno.ref, nil
}

func (pdno DataPrivateDnsNameOptionsAttributes) InternalWithRef(ref terra.Reference) DataPrivateDnsNameOptionsAttributes {
	return DataPrivateDnsNameOptionsAttributes{ref: ref}
}

func (pdno DataPrivateDnsNameOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pdno.ref.InternalTokens()
}

func (pdno DataPrivateDnsNameOptionsAttributes) EnableResourceNameDnsARecord() terra.BoolValue {
	return terra.ReferenceAsBool(pdno.ref.Append("enable_resource_name_dns_a_record"))
}

func (pdno DataPrivateDnsNameOptionsAttributes) EnableResourceNameDnsAaaaRecord() terra.BoolValue {
	return terra.ReferenceAsBool(pdno.ref.Append("enable_resource_name_dns_aaaa_record"))
}

func (pdno DataPrivateDnsNameOptionsAttributes) HostnameType() terra.StringValue {
	return terra.ReferenceAsString(pdno.ref.Append("hostname_type"))
}

type DataRootBlockDeviceAttributes struct {
	ref terra.Reference
}

func (rbd DataRootBlockDeviceAttributes) InternalRef() (terra.Reference, error) {
	return rbd.ref, nil
}

func (rbd DataRootBlockDeviceAttributes) InternalWithRef(ref terra.Reference) DataRootBlockDeviceAttributes {
	return DataRootBlockDeviceAttributes{ref: ref}
}

func (rbd DataRootBlockDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rbd.ref.InternalTokens()
}

func (rbd DataRootBlockDeviceAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(rbd.ref.Append("delete_on_termination"))
}

func (rbd DataRootBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("device_name"))
}

func (rbd DataRootBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rbd.ref.Append("encrypted"))
}

func (rbd DataRootBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("iops"))
}

func (rbd DataRootBlockDeviceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("kms_key_id"))
}

func (rbd DataRootBlockDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rbd.ref.Append("tags"))
}

func (rbd DataRootBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("throughput"))
}

func (rbd DataRootBlockDeviceAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("volume_id"))
}

func (rbd DataRootBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("volume_size"))
}

func (rbd DataRootBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("volume_type"))
}

type DataFilterAttributes struct {
	ref terra.Reference
}

func (f DataFilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFilterAttributes) InternalWithRef(ref terra.Reference) DataFilterAttributes {
	return DataFilterAttributes{ref: ref}
}

func (f DataFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f DataFilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("values"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataCreditSpecificationState struct {
	CpuCredits string `json:"cpu_credits"`
}

type DataEbsBlockDeviceState struct {
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

type DataEnclaveOptionsState struct {
	Enabled bool `json:"enabled"`
}

type DataEphemeralBlockDeviceState struct {
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
	VirtualName string `json:"virtual_name"`
}

type DataMaintenanceOptionsState struct {
	AutoRecovery string `json:"auto_recovery"`
}

type DataMetadataOptionsState struct {
	HttpEndpoint            string  `json:"http_endpoint"`
	HttpProtocolIpv6        string  `json:"http_protocol_ipv6"`
	HttpPutResponseHopLimit float64 `json:"http_put_response_hop_limit"`
	HttpTokens              string  `json:"http_tokens"`
	InstanceMetadataTags    string  `json:"instance_metadata_tags"`
}

type DataPrivateDnsNameOptionsState struct {
	EnableResourceNameDnsARecord    bool   `json:"enable_resource_name_dns_a_record"`
	EnableResourceNameDnsAaaaRecord bool   `json:"enable_resource_name_dns_aaaa_record"`
	HostnameType                    string `json:"hostname_type"`
}

type DataRootBlockDeviceState struct {
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

type DataFilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
