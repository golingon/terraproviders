// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package siterecoveryvmwarereplicatedvm

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ManagedDisk struct {
	// DiskId: string, required
	DiskId terra.StringValue `hcl:"disk_id,attr" validate:"required"`
	// LogStorageAccountId: string, optional
	LogStorageAccountId terra.StringValue `hcl:"log_storage_account_id,attr"`
	// TargetDiskEncryptionSetId: string, optional
	TargetDiskEncryptionSetId terra.StringValue `hcl:"target_disk_encryption_set_id,attr"`
	// TargetDiskType: string, required
	TargetDiskType terra.StringValue `hcl:"target_disk_type,attr" validate:"required"`
}

type NetworkInterface struct {
	// IsPrimary: bool, required
	IsPrimary terra.BoolValue `hcl:"is_primary,attr" validate:"required"`
	// SourceMacAddress: string, required
	SourceMacAddress terra.StringValue `hcl:"source_mac_address,attr" validate:"required"`
	// TargetStaticIp: string, optional
	TargetStaticIp terra.StringValue `hcl:"target_static_ip,attr"`
	// TargetSubnetName: string, optional
	TargetSubnetName terra.StringValue `hcl:"target_subnet_name,attr"`
	// TestSubnetName: string, optional
	TestSubnetName terra.StringValue `hcl:"test_subnet_name,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ManagedDiskAttributes struct {
	ref terra.Reference
}

func (md ManagedDiskAttributes) InternalRef() (terra.Reference, error) {
	return md.ref, nil
}

func (md ManagedDiskAttributes) InternalWithRef(ref terra.Reference) ManagedDiskAttributes {
	return ManagedDiskAttributes{ref: ref}
}

func (md ManagedDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return md.ref.InternalTokens()
}

func (md ManagedDiskAttributes) DiskId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("disk_id"))
}

func (md ManagedDiskAttributes) LogStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("log_storage_account_id"))
}

func (md ManagedDiskAttributes) TargetDiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("target_disk_encryption_set_id"))
}

func (md ManagedDiskAttributes) TargetDiskType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("target_disk_type"))
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

func (ni NetworkInterfaceAttributes) IsPrimary() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("is_primary"))
}

func (ni NetworkInterfaceAttributes) SourceMacAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("source_mac_address"))
}

func (ni NetworkInterfaceAttributes) TargetStaticIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("target_static_ip"))
}

func (ni NetworkInterfaceAttributes) TargetSubnetName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("target_subnet_name"))
}

func (ni NetworkInterfaceAttributes) TestSubnetName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("test_subnet_name"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ManagedDiskState struct {
	DiskId                    string `json:"disk_id"`
	LogStorageAccountId       string `json:"log_storage_account_id"`
	TargetDiskEncryptionSetId string `json:"target_disk_encryption_set_id"`
	TargetDiskType            string `json:"target_disk_type"`
}

type NetworkInterfaceState struct {
	IsPrimary        bool   `json:"is_primary"`
	SourceMacAddress string `json:"source_mac_address"`
	TargetStaticIp   string `json:"target_static_ip"`
	TargetSubnetName string `json:"target_subnet_name"`
	TestSubnetName   string `json:"test_subnet_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
