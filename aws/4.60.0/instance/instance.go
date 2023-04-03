// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package instance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CapacityReservationSpecification struct {
	// CapacityReservationPreference: string, optional
	CapacityReservationPreference terra.StringValue `hcl:"capacity_reservation_preference,attr"`
	// CapacityReservationTarget: optional
	CapacityReservationTarget *CapacityReservationTarget `hcl:"capacity_reservation_target,block"`
}

type CapacityReservationTarget struct {
	// CapacityReservationId: string, optional
	CapacityReservationId terra.StringValue `hcl:"capacity_reservation_id,attr"`
	// CapacityReservationResourceGroupArn: string, optional
	CapacityReservationResourceGroupArn terra.StringValue `hcl:"capacity_reservation_resource_group_arn,attr"`
}

type CreditSpecification struct {
	// CpuCredits: string, optional
	CpuCredits terra.StringValue `hcl:"cpu_credits,attr"`
}

type EbsBlockDevice struct {
	// DeleteOnTermination: bool, optional
	DeleteOnTermination terra.BoolValue `hcl:"delete_on_termination,attr"`
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// SnapshotId: string, optional
	SnapshotId terra.StringValue `hcl:"snapshot_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type EnclaveOptions struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type EphemeralBlockDevice struct {
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// NoDevice: bool, optional
	NoDevice terra.BoolValue `hcl:"no_device,attr"`
	// VirtualName: string, optional
	VirtualName terra.StringValue `hcl:"virtual_name,attr"`
}

type LaunchTemplate struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type MaintenanceOptions struct {
	// AutoRecovery: string, optional
	AutoRecovery terra.StringValue `hcl:"auto_recovery,attr"`
}

type MetadataOptions struct {
	// HttpEndpoint: string, optional
	HttpEndpoint terra.StringValue `hcl:"http_endpoint,attr"`
	// HttpPutResponseHopLimit: number, optional
	HttpPutResponseHopLimit terra.NumberValue `hcl:"http_put_response_hop_limit,attr"`
	// HttpTokens: string, optional
	HttpTokens terra.StringValue `hcl:"http_tokens,attr"`
	// InstanceMetadataTags: string, optional
	InstanceMetadataTags terra.StringValue `hcl:"instance_metadata_tags,attr"`
}

type NetworkInterface struct {
	// DeleteOnTermination: bool, optional
	DeleteOnTermination terra.BoolValue `hcl:"delete_on_termination,attr"`
	// DeviceIndex: number, required
	DeviceIndex terra.NumberValue `hcl:"device_index,attr" validate:"required"`
	// NetworkCardIndex: number, optional
	NetworkCardIndex terra.NumberValue `hcl:"network_card_index,attr"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
}

type PrivateDnsNameOptions struct {
	// EnableResourceNameDnsARecord: bool, optional
	EnableResourceNameDnsARecord terra.BoolValue `hcl:"enable_resource_name_dns_a_record,attr"`
	// EnableResourceNameDnsAaaaRecord: bool, optional
	EnableResourceNameDnsAaaaRecord terra.BoolValue `hcl:"enable_resource_name_dns_aaaa_record,attr"`
	// HostnameType: string, optional
	HostnameType terra.StringValue `hcl:"hostname_type,attr"`
}

type RootBlockDevice struct {
	// DeleteOnTermination: bool, optional
	DeleteOnTermination terra.BoolValue `hcl:"delete_on_termination,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CapacityReservationSpecificationAttributes struct {
	ref terra.Reference
}

func (crs CapacityReservationSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return crs.ref, nil
}

func (crs CapacityReservationSpecificationAttributes) InternalWithRef(ref terra.Reference) CapacityReservationSpecificationAttributes {
	return CapacityReservationSpecificationAttributes{ref: ref}
}

func (crs CapacityReservationSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return crs.ref.InternalTokens()
}

func (crs CapacityReservationSpecificationAttributes) CapacityReservationPreference() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("capacity_reservation_preference"))
}

func (crs CapacityReservationSpecificationAttributes) CapacityReservationTarget() terra.ListValue[CapacityReservationTargetAttributes] {
	return terra.ReferenceAsList[CapacityReservationTargetAttributes](crs.ref.Append("capacity_reservation_target"))
}

type CapacityReservationTargetAttributes struct {
	ref terra.Reference
}

func (crt CapacityReservationTargetAttributes) InternalRef() (terra.Reference, error) {
	return crt.ref, nil
}

func (crt CapacityReservationTargetAttributes) InternalWithRef(ref terra.Reference) CapacityReservationTargetAttributes {
	return CapacityReservationTargetAttributes{ref: ref}
}

func (crt CapacityReservationTargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return crt.ref.InternalTokens()
}

func (crt CapacityReservationTargetAttributes) CapacityReservationId() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("capacity_reservation_id"))
}

func (crt CapacityReservationTargetAttributes) CapacityReservationResourceGroupArn() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("capacity_reservation_resource_group_arn"))
}

type CreditSpecificationAttributes struct {
	ref terra.Reference
}

func (cs CreditSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs CreditSpecificationAttributes) InternalWithRef(ref terra.Reference) CreditSpecificationAttributes {
	return CreditSpecificationAttributes{ref: ref}
}

func (cs CreditSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs CreditSpecificationAttributes) CpuCredits() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("cpu_credits"))
}

type EbsBlockDeviceAttributes struct {
	ref terra.Reference
}

func (ebd EbsBlockDeviceAttributes) InternalRef() (terra.Reference, error) {
	return ebd.ref, nil
}

func (ebd EbsBlockDeviceAttributes) InternalWithRef(ref terra.Reference) EbsBlockDeviceAttributes {
	return EbsBlockDeviceAttributes{ref: ref}
}

func (ebd EbsBlockDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (ebd EbsBlockDeviceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("kms_key_id"))
}

func (ebd EbsBlockDeviceAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("snapshot_id"))
}

func (ebd EbsBlockDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ebd.ref.Append("tags"))
}

func (ebd EbsBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("throughput"))
}

func (ebd EbsBlockDeviceAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("volume_id"))
}

func (ebd EbsBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ebd.ref.Append("volume_size"))
}

func (ebd EbsBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(ebd.ref.Append("volume_type"))
}

type EnclaveOptionsAttributes struct {
	ref terra.Reference
}

func (eo EnclaveOptionsAttributes) InternalRef() (terra.Reference, error) {
	return eo.ref, nil
}

func (eo EnclaveOptionsAttributes) InternalWithRef(ref terra.Reference) EnclaveOptionsAttributes {
	return EnclaveOptionsAttributes{ref: ref}
}

func (eo EnclaveOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eo.ref.InternalTokens()
}

func (eo EnclaveOptionsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(eo.ref.Append("enabled"))
}

type EphemeralBlockDeviceAttributes struct {
	ref terra.Reference
}

func (ebd EphemeralBlockDeviceAttributes) InternalRef() (terra.Reference, error) {
	return ebd.ref, nil
}

func (ebd EphemeralBlockDeviceAttributes) InternalWithRef(ref terra.Reference) EphemeralBlockDeviceAttributes {
	return EphemeralBlockDeviceAttributes{ref: ref}
}

func (ebd EphemeralBlockDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

type LaunchTemplateAttributes struct {
	ref terra.Reference
}

func (lt LaunchTemplateAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt LaunchTemplateAttributes) InternalWithRef(ref terra.Reference) LaunchTemplateAttributes {
	return LaunchTemplateAttributes{ref: ref}
}

func (lt LaunchTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lt.ref.InternalTokens()
}

func (lt LaunchTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("id"))
}

func (lt LaunchTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("name"))
}

func (lt LaunchTemplateAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("version"))
}

type MaintenanceOptionsAttributes struct {
	ref terra.Reference
}

func (mo MaintenanceOptionsAttributes) InternalRef() (terra.Reference, error) {
	return mo.ref, nil
}

func (mo MaintenanceOptionsAttributes) InternalWithRef(ref terra.Reference) MaintenanceOptionsAttributes {
	return MaintenanceOptionsAttributes{ref: ref}
}

func (mo MaintenanceOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mo.ref.InternalTokens()
}

func (mo MaintenanceOptionsAttributes) AutoRecovery() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("auto_recovery"))
}

type MetadataOptionsAttributes struct {
	ref terra.Reference
}

func (mo MetadataOptionsAttributes) InternalRef() (terra.Reference, error) {
	return mo.ref, nil
}

func (mo MetadataOptionsAttributes) InternalWithRef(ref terra.Reference) MetadataOptionsAttributes {
	return MetadataOptionsAttributes{ref: ref}
}

func (mo MetadataOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (mo MetadataOptionsAttributes) InstanceMetadataTags() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("instance_metadata_tags"))
}

type NetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni NetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni NetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAttributes {
	return NetworkInterfaceAttributes{ref: ref}
}

func (ni NetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni NetworkInterfaceAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("delete_on_termination"))
}

func (ni NetworkInterfaceAttributes) DeviceIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("device_index"))
}

func (ni NetworkInterfaceAttributes) NetworkCardIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("network_card_index"))
}

func (ni NetworkInterfaceAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_interface_id"))
}

type PrivateDnsNameOptionsAttributes struct {
	ref terra.Reference
}

func (pdno PrivateDnsNameOptionsAttributes) InternalRef() (terra.Reference, error) {
	return pdno.ref, nil
}

func (pdno PrivateDnsNameOptionsAttributes) InternalWithRef(ref terra.Reference) PrivateDnsNameOptionsAttributes {
	return PrivateDnsNameOptionsAttributes{ref: ref}
}

func (pdno PrivateDnsNameOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pdno.ref.InternalTokens()
}

func (pdno PrivateDnsNameOptionsAttributes) EnableResourceNameDnsARecord() terra.BoolValue {
	return terra.ReferenceAsBool(pdno.ref.Append("enable_resource_name_dns_a_record"))
}

func (pdno PrivateDnsNameOptionsAttributes) EnableResourceNameDnsAaaaRecord() terra.BoolValue {
	return terra.ReferenceAsBool(pdno.ref.Append("enable_resource_name_dns_aaaa_record"))
}

func (pdno PrivateDnsNameOptionsAttributes) HostnameType() terra.StringValue {
	return terra.ReferenceAsString(pdno.ref.Append("hostname_type"))
}

type RootBlockDeviceAttributes struct {
	ref terra.Reference
}

func (rbd RootBlockDeviceAttributes) InternalRef() (terra.Reference, error) {
	return rbd.ref, nil
}

func (rbd RootBlockDeviceAttributes) InternalWithRef(ref terra.Reference) RootBlockDeviceAttributes {
	return RootBlockDeviceAttributes{ref: ref}
}

func (rbd RootBlockDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rbd.ref.InternalTokens()
}

func (rbd RootBlockDeviceAttributes) DeleteOnTermination() terra.BoolValue {
	return terra.ReferenceAsBool(rbd.ref.Append("delete_on_termination"))
}

func (rbd RootBlockDeviceAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("device_name"))
}

func (rbd RootBlockDeviceAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rbd.ref.Append("encrypted"))
}

func (rbd RootBlockDeviceAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("iops"))
}

func (rbd RootBlockDeviceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("kms_key_id"))
}

func (rbd RootBlockDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rbd.ref.Append("tags"))
}

func (rbd RootBlockDeviceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("throughput"))
}

func (rbd RootBlockDeviceAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("volume_id"))
}

func (rbd RootBlockDeviceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rbd.ref.Append("volume_size"))
}

func (rbd RootBlockDeviceAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(rbd.ref.Append("volume_type"))
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

type CapacityReservationSpecificationState struct {
	CapacityReservationPreference string                           `json:"capacity_reservation_preference"`
	CapacityReservationTarget     []CapacityReservationTargetState `json:"capacity_reservation_target"`
}

type CapacityReservationTargetState struct {
	CapacityReservationId               string `json:"capacity_reservation_id"`
	CapacityReservationResourceGroupArn string `json:"capacity_reservation_resource_group_arn"`
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

type LaunchTemplateState struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
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

type NetworkInterfaceState struct {
	DeleteOnTermination bool    `json:"delete_on_termination"`
	DeviceIndex         float64 `json:"device_index"`
	NetworkCardIndex    float64 `json:"network_card_index"`
	NetworkInterfaceId  string  `json:"network_interface_id"`
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

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
