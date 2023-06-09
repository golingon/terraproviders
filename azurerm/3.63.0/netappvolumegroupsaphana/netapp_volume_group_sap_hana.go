// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package netappvolumegroupsaphana

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

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

type Volume struct {
	// CapacityPoolId: string, required
	CapacityPoolId terra.StringValue `hcl:"capacity_pool_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocols: list of string, required
	Protocols terra.ListValue[terra.StringValue] `hcl:"protocols,attr" validate:"required"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// SecurityStyle: string, required
	SecurityStyle terra.StringValue `hcl:"security_style,attr" validate:"required"`
	// ServiceLevel: string, required
	ServiceLevel terra.StringValue `hcl:"service_level,attr" validate:"required"`
	// SnapshotDirectoryVisible: bool, required
	SnapshotDirectoryVisible terra.BoolValue `hcl:"snapshot_directory_visible,attr" validate:"required"`
	// StorageQuotaInGb: number, required
	StorageQuotaInGb terra.NumberValue `hcl:"storage_quota_in_gb,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ThroughputInMibps: number, required
	ThroughputInMibps terra.NumberValue `hcl:"throughput_in_mibps,attr" validate:"required"`
	// VolumePath: string, required
	VolumePath terra.StringValue `hcl:"volume_path,attr" validate:"required"`
	// VolumeSpecName: string, required
	VolumeSpecName terra.StringValue `hcl:"volume_spec_name,attr" validate:"required"`
	// DataProtectionReplication: optional
	DataProtectionReplication *DataProtectionReplication `hcl:"data_protection_replication,block"`
	// DataProtectionSnapshotPolicy: optional
	DataProtectionSnapshotPolicy *DataProtectionSnapshotPolicy `hcl:"data_protection_snapshot_policy,block"`
	// ExportPolicyRule: min=1,max=5
	ExportPolicyRule []ExportPolicyRule `hcl:"export_policy_rule,block" validate:"min=1,max=5"`
}

type DataProtectionReplication struct {
	// EndpointType: string, optional
	EndpointType terra.StringValue `hcl:"endpoint_type,attr"`
	// RemoteVolumeLocation: string, required
	RemoteVolumeLocation terra.StringValue `hcl:"remote_volume_location,attr" validate:"required"`
	// RemoteVolumeResourceId: string, required
	RemoteVolumeResourceId terra.StringValue `hcl:"remote_volume_resource_id,attr" validate:"required"`
	// ReplicationFrequency: string, required
	ReplicationFrequency terra.StringValue `hcl:"replication_frequency,attr" validate:"required"`
}

type DataProtectionSnapshotPolicy struct {
	// SnapshotPolicyId: string, required
	SnapshotPolicyId terra.StringValue `hcl:"snapshot_policy_id,attr" validate:"required"`
}

type ExportPolicyRule struct {
	// AllowedClients: string, required
	AllowedClients terra.StringValue `hcl:"allowed_clients,attr" validate:"required"`
	// Nfsv3Enabled: bool, required
	Nfsv3Enabled terra.BoolValue `hcl:"nfsv3_enabled,attr" validate:"required"`
	// Nfsv41Enabled: bool, required
	Nfsv41Enabled terra.BoolValue `hcl:"nfsv41_enabled,attr" validate:"required"`
	// RootAccessEnabled: bool, optional
	RootAccessEnabled terra.BoolValue `hcl:"root_access_enabled,attr"`
	// RuleIndex: number, required
	RuleIndex terra.NumberValue `hcl:"rule_index,attr" validate:"required"`
	// UnixReadOnly: bool, optional
	UnixReadOnly terra.BoolValue `hcl:"unix_read_only,attr"`
	// UnixReadWrite: bool, optional
	UnixReadWrite terra.BoolValue `hcl:"unix_read_write,attr"`
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

type VolumeAttributes struct {
	ref terra.Reference
}

func (v VolumeAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VolumeAttributes) InternalWithRef(ref terra.Reference) VolumeAttributes {
	return VolumeAttributes{ref: ref}
}

func (v VolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VolumeAttributes) CapacityPoolId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("capacity_pool_id"))
}

func (v VolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("id"))
}

func (v VolumeAttributes) MountIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("mount_ip_addresses"))
}

func (v VolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v VolumeAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("protocols"))
}

func (v VolumeAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("proximity_placement_group_id"))
}

func (v VolumeAttributes) SecurityStyle() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("security_style"))
}

func (v VolumeAttributes) ServiceLevel() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("service_level"))
}

func (v VolumeAttributes) SnapshotDirectoryVisible() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("snapshot_directory_visible"))
}

func (v VolumeAttributes) StorageQuotaInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(v.ref.Append("storage_quota_in_gb"))
}

func (v VolumeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("subnet_id"))
}

func (v VolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](v.ref.Append("tags"))
}

func (v VolumeAttributes) ThroughputInMibps() terra.NumberValue {
	return terra.ReferenceAsNumber(v.ref.Append("throughput_in_mibps"))
}

func (v VolumeAttributes) VolumePath() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("volume_path"))
}

func (v VolumeAttributes) VolumeSpecName() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("volume_spec_name"))
}

func (v VolumeAttributes) DataProtectionReplication() terra.ListValue[DataProtectionReplicationAttributes] {
	return terra.ReferenceAsList[DataProtectionReplicationAttributes](v.ref.Append("data_protection_replication"))
}

func (v VolumeAttributes) DataProtectionSnapshotPolicy() terra.ListValue[DataProtectionSnapshotPolicyAttributes] {
	return terra.ReferenceAsList[DataProtectionSnapshotPolicyAttributes](v.ref.Append("data_protection_snapshot_policy"))
}

func (v VolumeAttributes) ExportPolicyRule() terra.ListValue[ExportPolicyRuleAttributes] {
	return terra.ReferenceAsList[ExportPolicyRuleAttributes](v.ref.Append("export_policy_rule"))
}

type DataProtectionReplicationAttributes struct {
	ref terra.Reference
}

func (dpr DataProtectionReplicationAttributes) InternalRef() (terra.Reference, error) {
	return dpr.ref, nil
}

func (dpr DataProtectionReplicationAttributes) InternalWithRef(ref terra.Reference) DataProtectionReplicationAttributes {
	return DataProtectionReplicationAttributes{ref: ref}
}

func (dpr DataProtectionReplicationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dpr.ref.InternalTokens()
}

func (dpr DataProtectionReplicationAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("endpoint_type"))
}

func (dpr DataProtectionReplicationAttributes) RemoteVolumeLocation() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("remote_volume_location"))
}

func (dpr DataProtectionReplicationAttributes) RemoteVolumeResourceId() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("remote_volume_resource_id"))
}

func (dpr DataProtectionReplicationAttributes) ReplicationFrequency() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("replication_frequency"))
}

type DataProtectionSnapshotPolicyAttributes struct {
	ref terra.Reference
}

func (dpsp DataProtectionSnapshotPolicyAttributes) InternalRef() (terra.Reference, error) {
	return dpsp.ref, nil
}

func (dpsp DataProtectionSnapshotPolicyAttributes) InternalWithRef(ref terra.Reference) DataProtectionSnapshotPolicyAttributes {
	return DataProtectionSnapshotPolicyAttributes{ref: ref}
}

func (dpsp DataProtectionSnapshotPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dpsp.ref.InternalTokens()
}

func (dpsp DataProtectionSnapshotPolicyAttributes) SnapshotPolicyId() terra.StringValue {
	return terra.ReferenceAsString(dpsp.ref.Append("snapshot_policy_id"))
}

type ExportPolicyRuleAttributes struct {
	ref terra.Reference
}

func (epr ExportPolicyRuleAttributes) InternalRef() (terra.Reference, error) {
	return epr.ref, nil
}

func (epr ExportPolicyRuleAttributes) InternalWithRef(ref terra.Reference) ExportPolicyRuleAttributes {
	return ExportPolicyRuleAttributes{ref: ref}
}

func (epr ExportPolicyRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return epr.ref.InternalTokens()
}

func (epr ExportPolicyRuleAttributes) AllowedClients() terra.StringValue {
	return terra.ReferenceAsString(epr.ref.Append("allowed_clients"))
}

func (epr ExportPolicyRuleAttributes) Nfsv3Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("nfsv3_enabled"))
}

func (epr ExportPolicyRuleAttributes) Nfsv41Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("nfsv41_enabled"))
}

func (epr ExportPolicyRuleAttributes) RootAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("root_access_enabled"))
}

func (epr ExportPolicyRuleAttributes) RuleIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(epr.ref.Append("rule_index"))
}

func (epr ExportPolicyRuleAttributes) UnixReadOnly() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("unix_read_only"))
}

func (epr ExportPolicyRuleAttributes) UnixReadWrite() terra.BoolValue {
	return terra.ReferenceAsBool(epr.ref.Append("unix_read_write"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type VolumeState struct {
	CapacityPoolId               string                              `json:"capacity_pool_id"`
	Id                           string                              `json:"id"`
	MountIpAddresses             []string                            `json:"mount_ip_addresses"`
	Name                         string                              `json:"name"`
	Protocols                    []string                            `json:"protocols"`
	ProximityPlacementGroupId    string                              `json:"proximity_placement_group_id"`
	SecurityStyle                string                              `json:"security_style"`
	ServiceLevel                 string                              `json:"service_level"`
	SnapshotDirectoryVisible     bool                                `json:"snapshot_directory_visible"`
	StorageQuotaInGb             float64                             `json:"storage_quota_in_gb"`
	SubnetId                     string                              `json:"subnet_id"`
	Tags                         map[string]string                   `json:"tags"`
	ThroughputInMibps            float64                             `json:"throughput_in_mibps"`
	VolumePath                   string                              `json:"volume_path"`
	VolumeSpecName               string                              `json:"volume_spec_name"`
	DataProtectionReplication    []DataProtectionReplicationState    `json:"data_protection_replication"`
	DataProtectionSnapshotPolicy []DataProtectionSnapshotPolicyState `json:"data_protection_snapshot_policy"`
	ExportPolicyRule             []ExportPolicyRuleState             `json:"export_policy_rule"`
}

type DataProtectionReplicationState struct {
	EndpointType           string `json:"endpoint_type"`
	RemoteVolumeLocation   string `json:"remote_volume_location"`
	RemoteVolumeResourceId string `json:"remote_volume_resource_id"`
	ReplicationFrequency   string `json:"replication_frequency"`
}

type DataProtectionSnapshotPolicyState struct {
	SnapshotPolicyId string `json:"snapshot_policy_id"`
}

type ExportPolicyRuleState struct {
	AllowedClients    string  `json:"allowed_clients"`
	Nfsv3Enabled      bool    `json:"nfsv3_enabled"`
	Nfsv41Enabled     bool    `json:"nfsv41_enabled"`
	RootAccessEnabled bool    `json:"root_access_enabled"`
	RuleIndex         float64 `json:"rule_index"`
	UnixReadOnly      bool    `json:"unix_read_only"`
	UnixReadWrite     bool    `json:"unix_read_write"`
}
