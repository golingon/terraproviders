// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package siterecoveryreplicatedvm

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ManagedDisk struct {
	// DiskId: string, optional
	DiskId terra.StringValue `hcl:"disk_id,attr"`
	// StagingStorageAccountId: string, optional
	StagingStorageAccountId terra.StringValue `hcl:"staging_storage_account_id,attr"`
	// TargetDiskEncryptionSetId: string, optional
	TargetDiskEncryptionSetId terra.StringValue `hcl:"target_disk_encryption_set_id,attr"`
	// TargetDiskType: string, optional
	TargetDiskType terra.StringValue `hcl:"target_disk_type,attr"`
	// TargetReplicaDiskType: string, optional
	TargetReplicaDiskType terra.StringValue `hcl:"target_replica_disk_type,attr"`
	// TargetResourceGroupId: string, optional
	TargetResourceGroupId terra.StringValue `hcl:"target_resource_group_id,attr"`
	// TargetDiskEncryption: min=0
	TargetDiskEncryption []TargetDiskEncryption `hcl:"target_disk_encryption,block" validate:"min=0"`
}

type TargetDiskEncryption struct {
	// DiskEncryptionKey: min=0
	DiskEncryptionKey []DiskEncryptionKey `hcl:"disk_encryption_key,block" validate:"min=0"`
	// KeyEncryptionKey: min=0
	KeyEncryptionKey []KeyEncryptionKey `hcl:"key_encryption_key,block" validate:"min=0"`
}

type DiskEncryptionKey struct {
	// SecretUrl: string, optional
	SecretUrl terra.StringValue `hcl:"secret_url,attr"`
	// VaultId: string, optional
	VaultId terra.StringValue `hcl:"vault_id,attr"`
}

type KeyEncryptionKey struct {
	// KeyUrl: string, optional
	KeyUrl terra.StringValue `hcl:"key_url,attr"`
	// VaultId: string, optional
	VaultId terra.StringValue `hcl:"vault_id,attr"`
}

type NetworkInterface struct {
	// FailoverTestPublicIpAddressId: string, optional
	FailoverTestPublicIpAddressId terra.StringValue `hcl:"failover_test_public_ip_address_id,attr"`
	// FailoverTestStaticIp: string, optional
	FailoverTestStaticIp terra.StringValue `hcl:"failover_test_static_ip,attr"`
	// FailoverTestSubnetName: string, optional
	FailoverTestSubnetName terra.StringValue `hcl:"failover_test_subnet_name,attr"`
	// IsPrimary: bool, optional
	IsPrimary terra.BoolValue `hcl:"is_primary,attr"`
	// RecoveryPublicIpAddressId: string, optional
	RecoveryPublicIpAddressId terra.StringValue `hcl:"recovery_public_ip_address_id,attr"`
	// SourceNetworkInterfaceId: string, optional
	SourceNetworkInterfaceId terra.StringValue `hcl:"source_network_interface_id,attr"`
	// TargetStaticIp: string, optional
	TargetStaticIp terra.StringValue `hcl:"target_static_ip,attr"`
	// TargetSubnetName: string, optional
	TargetSubnetName terra.StringValue `hcl:"target_subnet_name,attr"`
}

type UnmanagedDisk struct {
	// DiskUri: string, optional
	DiskUri terra.StringValue `hcl:"disk_uri,attr"`
	// StagingStorageAccountId: string, optional
	StagingStorageAccountId terra.StringValue `hcl:"staging_storage_account_id,attr"`
	// TargetStorageAccountId: string, optional
	TargetStorageAccountId terra.StringValue `hcl:"target_storage_account_id,attr"`
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

func (md ManagedDiskAttributes) StagingStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("staging_storage_account_id"))
}

func (md ManagedDiskAttributes) TargetDiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("target_disk_encryption_set_id"))
}

func (md ManagedDiskAttributes) TargetDiskType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("target_disk_type"))
}

func (md ManagedDiskAttributes) TargetReplicaDiskType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("target_replica_disk_type"))
}

func (md ManagedDiskAttributes) TargetResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("target_resource_group_id"))
}

func (md ManagedDiskAttributes) TargetDiskEncryption() terra.ListValue[TargetDiskEncryptionAttributes] {
	return terra.ReferenceAsList[TargetDiskEncryptionAttributes](md.ref.Append("target_disk_encryption"))
}

type TargetDiskEncryptionAttributes struct {
	ref terra.Reference
}

func (tde TargetDiskEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return tde.ref, nil
}

func (tde TargetDiskEncryptionAttributes) InternalWithRef(ref terra.Reference) TargetDiskEncryptionAttributes {
	return TargetDiskEncryptionAttributes{ref: ref}
}

func (tde TargetDiskEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tde.ref.InternalTokens()
}

func (tde TargetDiskEncryptionAttributes) DiskEncryptionKey() terra.ListValue[DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DiskEncryptionKeyAttributes](tde.ref.Append("disk_encryption_key"))
}

func (tde TargetDiskEncryptionAttributes) KeyEncryptionKey() terra.ListValue[KeyEncryptionKeyAttributes] {
	return terra.ReferenceAsList[KeyEncryptionKeyAttributes](tde.ref.Append("key_encryption_key"))
}

type DiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (dek DiskEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return dek.ref, nil
}

func (dek DiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DiskEncryptionKeyAttributes {
	return DiskEncryptionKeyAttributes{ref: ref}
}

func (dek DiskEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dek.ref.InternalTokens()
}

func (dek DiskEncryptionKeyAttributes) SecretUrl() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("secret_url"))
}

func (dek DiskEncryptionKeyAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("vault_id"))
}

type KeyEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (kek KeyEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return kek.ref, nil
}

func (kek KeyEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) KeyEncryptionKeyAttributes {
	return KeyEncryptionKeyAttributes{ref: ref}
}

func (kek KeyEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kek.ref.InternalTokens()
}

func (kek KeyEncryptionKeyAttributes) KeyUrl() terra.StringValue {
	return terra.ReferenceAsString(kek.ref.Append("key_url"))
}

func (kek KeyEncryptionKeyAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(kek.ref.Append("vault_id"))
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

func (ni NetworkInterfaceAttributes) FailoverTestPublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("failover_test_public_ip_address_id"))
}

func (ni NetworkInterfaceAttributes) FailoverTestStaticIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("failover_test_static_ip"))
}

func (ni NetworkInterfaceAttributes) FailoverTestSubnetName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("failover_test_subnet_name"))
}

func (ni NetworkInterfaceAttributes) IsPrimary() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("is_primary"))
}

func (ni NetworkInterfaceAttributes) RecoveryPublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("recovery_public_ip_address_id"))
}

func (ni NetworkInterfaceAttributes) SourceNetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("source_network_interface_id"))
}

func (ni NetworkInterfaceAttributes) TargetStaticIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("target_static_ip"))
}

func (ni NetworkInterfaceAttributes) TargetSubnetName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("target_subnet_name"))
}

type UnmanagedDiskAttributes struct {
	ref terra.Reference
}

func (ud UnmanagedDiskAttributes) InternalRef() (terra.Reference, error) {
	return ud.ref, nil
}

func (ud UnmanagedDiskAttributes) InternalWithRef(ref terra.Reference) UnmanagedDiskAttributes {
	return UnmanagedDiskAttributes{ref: ref}
}

func (ud UnmanagedDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ud.ref.InternalTokens()
}

func (ud UnmanagedDiskAttributes) DiskUri() terra.StringValue {
	return terra.ReferenceAsString(ud.ref.Append("disk_uri"))
}

func (ud UnmanagedDiskAttributes) StagingStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ud.ref.Append("staging_storage_account_id"))
}

func (ud UnmanagedDiskAttributes) TargetStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ud.ref.Append("target_storage_account_id"))
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
	DiskId                    string                      `json:"disk_id"`
	StagingStorageAccountId   string                      `json:"staging_storage_account_id"`
	TargetDiskEncryptionSetId string                      `json:"target_disk_encryption_set_id"`
	TargetDiskType            string                      `json:"target_disk_type"`
	TargetReplicaDiskType     string                      `json:"target_replica_disk_type"`
	TargetResourceGroupId     string                      `json:"target_resource_group_id"`
	TargetDiskEncryption      []TargetDiskEncryptionState `json:"target_disk_encryption"`
}

type TargetDiskEncryptionState struct {
	DiskEncryptionKey []DiskEncryptionKeyState `json:"disk_encryption_key"`
	KeyEncryptionKey  []KeyEncryptionKeyState  `json:"key_encryption_key"`
}

type DiskEncryptionKeyState struct {
	SecretUrl string `json:"secret_url"`
	VaultId   string `json:"vault_id"`
}

type KeyEncryptionKeyState struct {
	KeyUrl  string `json:"key_url"`
	VaultId string `json:"vault_id"`
}

type NetworkInterfaceState struct {
	FailoverTestPublicIpAddressId string `json:"failover_test_public_ip_address_id"`
	FailoverTestStaticIp          string `json:"failover_test_static_ip"`
	FailoverTestSubnetName        string `json:"failover_test_subnet_name"`
	IsPrimary                     bool   `json:"is_primary"`
	RecoveryPublicIpAddressId     string `json:"recovery_public_ip_address_id"`
	SourceNetworkInterfaceId      string `json:"source_network_interface_id"`
	TargetStaticIp                string `json:"target_static_ip"`
	TargetSubnetName              string `json:"target_subnet_name"`
}

type UnmanagedDiskState struct {
	DiskUri                 string `json:"disk_uri"`
	StagingStorageAccountId string `json:"staging_storage_account_id"`
	TargetStorageAccountId  string `json:"target_storage_account_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
