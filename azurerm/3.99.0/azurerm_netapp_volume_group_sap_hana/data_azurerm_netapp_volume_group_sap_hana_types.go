// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_netapp_volume_group_sap_hana

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataVolumeAttributes struct {
	ref terra.Reference
}

func (v DataVolumeAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v DataVolumeAttributes) InternalWithRef(ref terra.Reference) DataVolumeAttributes {
	return DataVolumeAttributes{ref: ref}
}

func (v DataVolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v DataVolumeAttributes) CapacityPoolId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("capacity_pool_id"))
}

func (v DataVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("id"))
}

func (v DataVolumeAttributes) MountIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("mount_ip_addresses"))
}

func (v DataVolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v DataVolumeAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("protocols"))
}

func (v DataVolumeAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("proximity_placement_group_id"))
}

func (v DataVolumeAttributes) SecurityStyle() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("security_style"))
}

func (v DataVolumeAttributes) ServiceLevel() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("service_level"))
}

func (v DataVolumeAttributes) SnapshotDirectoryVisible() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("snapshot_directory_visible"))
}

func (v DataVolumeAttributes) StorageQuotaInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(v.ref.Append("storage_quota_in_gb"))
}

func (v DataVolumeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("subnet_id"))
}

func (v DataVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](v.ref.Append("tags"))
}

func (v DataVolumeAttributes) ThroughputInMibps() terra.NumberValue {
	return terra.ReferenceAsNumber(v.ref.Append("throughput_in_mibps"))
}

func (v DataVolumeAttributes) VolumePath() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("volume_path"))
}

func (v DataVolumeAttributes) VolumeSpecName() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("volume_spec_name"))
}

func (v DataVolumeAttributes) DataProtectionReplication() terra.ListValue[DataVolumeDataProtectionReplicationAttributes] {
	return terra.ReferenceAsList[DataVolumeDataProtectionReplicationAttributes](v.ref.Append("data_protection_replication"))
}

func (v DataVolumeAttributes) DataProtectionSnapshotPolicy() terra.ListValue[DataVolumeDataProtectionSnapshotPolicyAttributes] {
	return terra.ReferenceAsList[DataVolumeDataProtectionSnapshotPolicyAttributes](v.ref.Append("data_protection_snapshot_policy"))
}

func (v DataVolumeAttributes) ExportPolicyRule() terra.ListValue[DataVolumeExportPolicyRuleAttributes] {
	return terra.ReferenceAsList[DataVolumeExportPolicyRuleAttributes](v.ref.Append("export_policy_rule"))
}

type DataVolumeDataProtectionReplicationAttributes struct {
	ref terra.Reference
}

func (dpr DataVolumeDataProtectionReplicationAttributes) InternalRef() (terra.Reference, error) {
	return dpr.ref, nil
}

func (dpr DataVolumeDataProtectionReplicationAttributes) InternalWithRef(ref terra.Reference) DataVolumeDataProtectionReplicationAttributes {
	return DataVolumeDataProtectionReplicationAttributes{ref: ref}
}

func (dpr DataVolumeDataProtectionReplicationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dpr.ref.InternalTokens()
}

func (dpr DataVolumeDataProtectionReplicationAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("endpoint_type"))
}

func (dpr DataVolumeDataProtectionReplicationAttributes) RemoteVolumeLocation() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("remote_volume_location"))
}

func (dpr DataVolumeDataProtectionReplicationAttributes) RemoteVolumeResourceId() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("remote_volume_resource_id"))
}

func (dpr DataVolumeDataProtectionReplicationAttributes) ReplicationFrequency() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("replication_frequency"))
}

type DataVolumeDataProtectionSnapshotPolicyAttributes struct {
	ref terra.Reference
}

func (dpsp DataVolumeDataProtectionSnapshotPolicyAttributes) InternalRef() (terra.Reference, error) {
	return dpsp.ref, nil
}

func (dpsp DataVolumeDataProtectionSnapshotPolicyAttributes) InternalWithRef(ref terra.Reference) DataVolumeDataProtectionSnapshotPolicyAttributes {
	return DataVolumeDataProtectionSnapshotPolicyAttributes{ref: ref}
}

func (dpsp DataVolumeDataProtectionSnapshotPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dpsp.ref.InternalTokens()
}

func (dpsp DataVolumeDataProtectionSnapshotPolicyAttributes) SnapshotPolicyId() terra.StringValue {
	return terra.ReferenceAsString(dpsp.ref.Append("snapshot_policy_id"))
}

type DataVolumeExportPolicyRuleAttributes struct {
	ref terra.Reference
}

func (epr DataVolumeExportPolicyRuleAttributes) InternalRef() (terra.Reference, error) {
	return epr.ref, nil
}

func (epr DataVolumeExportPolicyRuleAttributes) InternalWithRef(ref terra.Reference) DataVolumeExportPolicyRuleAttributes {
	return DataVolumeExportPolicyRuleAttributes{ref: ref}
}

func (epr DataVolumeExportPolicyRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return epr.ref.InternalTokens()
}

func (epr DataVolumeExportPolicyRuleAttributes) AllowedClients() terra.StringValue {
	return terra.ReferenceAsString(epr.ref.Append("allowed_clients"))
}

func (epr DataVolumeExportPolicyRuleAttributes) Nfsv3Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("nfsv3_enabled"))
}

func (epr DataVolumeExportPolicyRuleAttributes) Nfsv41Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("nfsv41_enabled"))
}

func (epr DataVolumeExportPolicyRuleAttributes) RootAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("root_access_enabled"))
}

func (epr DataVolumeExportPolicyRuleAttributes) RuleIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(epr.ref.Append("rule_index"))
}

func (epr DataVolumeExportPolicyRuleAttributes) UnixReadOnly() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("unix_read_only"))
}

func (epr DataVolumeExportPolicyRuleAttributes) UnixReadWrite() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("unix_read_write"))
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

type DataVolumeState struct {
	CapacityPoolId               string                                        `json:"capacity_pool_id"`
	Id                           string                                        `json:"id"`
	MountIpAddresses             []string                                      `json:"mount_ip_addresses"`
	Name                         string                                        `json:"name"`
	Protocols                    []string                                      `json:"protocols"`
	ProximityPlacementGroupId    string                                        `json:"proximity_placement_group_id"`
	SecurityStyle                string                                        `json:"security_style"`
	ServiceLevel                 string                                        `json:"service_level"`
	SnapshotDirectoryVisible     bool                                          `json:"snapshot_directory_visible"`
	StorageQuotaInGb             float64                                       `json:"storage_quota_in_gb"`
	SubnetId                     string                                        `json:"subnet_id"`
	Tags                         map[string]string                             `json:"tags"`
	ThroughputInMibps            float64                                       `json:"throughput_in_mibps"`
	VolumePath                   string                                        `json:"volume_path"`
	VolumeSpecName               string                                        `json:"volume_spec_name"`
	DataProtectionReplication    []DataVolumeDataProtectionReplicationState    `json:"data_protection_replication"`
	DataProtectionSnapshotPolicy []DataVolumeDataProtectionSnapshotPolicyState `json:"data_protection_snapshot_policy"`
	ExportPolicyRule             []DataVolumeExportPolicyRuleState             `json:"export_policy_rule"`
}

type DataVolumeDataProtectionReplicationState struct {
	EndpointType           string `json:"endpoint_type"`
	RemoteVolumeLocation   string `json:"remote_volume_location"`
	RemoteVolumeResourceId string `json:"remote_volume_resource_id"`
	ReplicationFrequency   string `json:"replication_frequency"`
}

type DataVolumeDataProtectionSnapshotPolicyState struct {
	SnapshotPolicyId string `json:"snapshot_policy_id"`
}

type DataVolumeExportPolicyRuleState struct {
	AllowedClients    string  `json:"allowed_clients"`
	Nfsv3Enabled      bool    `json:"nfsv3_enabled"`
	Nfsv41Enabled     bool    `json:"nfsv41_enabled"`
	RootAccessEnabled bool    `json:"root_access_enabled"`
	RuleIndex         float64 `json:"rule_index"`
	UnixReadOnly      bool    `json:"unix_read_only"`
	UnixReadWrite     bool    `json:"unix_read_write"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
