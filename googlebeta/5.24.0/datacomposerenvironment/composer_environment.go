// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacomposerenvironment

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
	// DataRetentionConfig: min=0
	DataRetentionConfig []DataRetentionConfig `hcl:"data_retention_config,block" validate:"min=0"`
	// DatabaseConfig: min=0
	DatabaseConfig []DatabaseConfig `hcl:"database_config,block" validate:"min=0"`
	// EncryptionConfig: min=0
	EncryptionConfig []EncryptionConfig `hcl:"encryption_config,block" validate:"min=0"`
	// MaintenanceWindow: min=0
	MaintenanceWindow []MaintenanceWindow `hcl:"maintenance_window,block" validate:"min=0"`
	// MasterAuthorizedNetworksConfig: min=0
	MasterAuthorizedNetworksConfig []MasterAuthorizedNetworksConfig `hcl:"master_authorized_networks_config,block" validate:"min=0"`
	// NodeConfig: min=0
	NodeConfig []NodeConfig `hcl:"node_config,block" validate:"min=0"`
	// PrivateEnvironmentConfig: min=0
	PrivateEnvironmentConfig []PrivateEnvironmentConfig `hcl:"private_environment_config,block" validate:"min=0"`
	// RecoveryConfig: min=0
	RecoveryConfig []RecoveryConfig `hcl:"recovery_config,block" validate:"min=0"`
	// SoftwareConfig: min=0
	SoftwareConfig []SoftwareConfig `hcl:"software_config,block" validate:"min=0"`
	// WebServerConfig: min=0
	WebServerConfig []WebServerConfig `hcl:"web_server_config,block" validate:"min=0"`
	// WebServerNetworkAccessControl: min=0
	WebServerNetworkAccessControl []WebServerNetworkAccessControl `hcl:"web_server_network_access_control,block" validate:"min=0"`
	// WorkloadsConfig: min=0
	WorkloadsConfig []WorkloadsConfig `hcl:"workloads_config,block" validate:"min=0"`
}

type DataRetentionConfig struct {
	// TaskLogsRetentionConfig: min=0
	TaskLogsRetentionConfig []TaskLogsRetentionConfig `hcl:"task_logs_retention_config,block" validate:"min=0"`
}

type TaskLogsRetentionConfig struct{}

type DatabaseConfig struct{}

type EncryptionConfig struct{}

type MaintenanceWindow struct{}

type MasterAuthorizedNetworksConfig struct {
	// CidrBlocks: min=0
	CidrBlocks []CidrBlocks `hcl:"cidr_blocks,block" validate:"min=0"`
}

type CidrBlocks struct{}

type NodeConfig struct {
	// IpAllocationPolicy: min=0
	IpAllocationPolicy []IpAllocationPolicy `hcl:"ip_allocation_policy,block" validate:"min=0"`
}

type IpAllocationPolicy struct{}

type PrivateEnvironmentConfig struct{}

type RecoveryConfig struct {
	// ScheduledSnapshotsConfig: min=0
	ScheduledSnapshotsConfig []ScheduledSnapshotsConfig `hcl:"scheduled_snapshots_config,block" validate:"min=0"`
}

type ScheduledSnapshotsConfig struct{}

type SoftwareConfig struct {
	// CloudDataLineageIntegration: min=0
	CloudDataLineageIntegration []CloudDataLineageIntegration `hcl:"cloud_data_lineage_integration,block" validate:"min=0"`
}

type CloudDataLineageIntegration struct{}

type WebServerConfig struct{}

type WebServerNetworkAccessControl struct {
	// AllowedIpRange: min=0
	AllowedIpRange []AllowedIpRange `hcl:"allowed_ip_range,block" validate:"min=0"`
}

type AllowedIpRange struct{}

type WorkloadsConfig struct {
	// DagProcessor: min=0
	DagProcessor []DagProcessor `hcl:"dag_processor,block" validate:"min=0"`
	// Scheduler: min=0
	Scheduler []Scheduler `hcl:"scheduler,block" validate:"min=0"`
	// Triggerer: min=0
	Triggerer []Triggerer `hcl:"triggerer,block" validate:"min=0"`
	// WebServer: min=0
	WebServer []WebServer `hcl:"web_server,block" validate:"min=0"`
	// Worker: min=0
	Worker []Worker `hcl:"worker,block" validate:"min=0"`
}

type DagProcessor struct{}

type Scheduler struct{}

type Triggerer struct{}

type WebServer struct{}

type Worker struct{}

type StorageConfig struct{}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) AirflowUri() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("airflow_uri"))
}

func (c ConfigAttributes) DagGcsPrefix() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("dag_gcs_prefix"))
}

func (c ConfigAttributes) EnablePrivateBuildsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("enable_private_builds_only"))
}

func (c ConfigAttributes) EnablePrivateEnvironment() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("enable_private_environment"))
}

func (c ConfigAttributes) EnvironmentSize() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("environment_size"))
}

func (c ConfigAttributes) GkeCluster() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("gke_cluster"))
}

func (c ConfigAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("node_count"))
}

func (c ConfigAttributes) ResilienceMode() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("resilience_mode"))
}

func (c ConfigAttributes) DataRetentionConfig() terra.ListValue[DataRetentionConfigAttributes] {
	return terra.ReferenceAsList[DataRetentionConfigAttributes](c.ref.Append("data_retention_config"))
}

func (c ConfigAttributes) DatabaseConfig() terra.ListValue[DatabaseConfigAttributes] {
	return terra.ReferenceAsList[DatabaseConfigAttributes](c.ref.Append("database_config"))
}

func (c ConfigAttributes) EncryptionConfig() terra.ListValue[EncryptionConfigAttributes] {
	return terra.ReferenceAsList[EncryptionConfigAttributes](c.ref.Append("encryption_config"))
}

func (c ConfigAttributes) MaintenanceWindow() terra.ListValue[MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[MaintenanceWindowAttributes](c.ref.Append("maintenance_window"))
}

func (c ConfigAttributes) MasterAuthorizedNetworksConfig() terra.ListValue[MasterAuthorizedNetworksConfigAttributes] {
	return terra.ReferenceAsList[MasterAuthorizedNetworksConfigAttributes](c.ref.Append("master_authorized_networks_config"))
}

func (c ConfigAttributes) NodeConfig() terra.ListValue[NodeConfigAttributes] {
	return terra.ReferenceAsList[NodeConfigAttributes](c.ref.Append("node_config"))
}

func (c ConfigAttributes) PrivateEnvironmentConfig() terra.ListValue[PrivateEnvironmentConfigAttributes] {
	return terra.ReferenceAsList[PrivateEnvironmentConfigAttributes](c.ref.Append("private_environment_config"))
}

func (c ConfigAttributes) RecoveryConfig() terra.ListValue[RecoveryConfigAttributes] {
	return terra.ReferenceAsList[RecoveryConfigAttributes](c.ref.Append("recovery_config"))
}

func (c ConfigAttributes) SoftwareConfig() terra.ListValue[SoftwareConfigAttributes] {
	return terra.ReferenceAsList[SoftwareConfigAttributes](c.ref.Append("software_config"))
}

func (c ConfigAttributes) WebServerConfig() terra.ListValue[WebServerConfigAttributes] {
	return terra.ReferenceAsList[WebServerConfigAttributes](c.ref.Append("web_server_config"))
}

func (c ConfigAttributes) WebServerNetworkAccessControl() terra.ListValue[WebServerNetworkAccessControlAttributes] {
	return terra.ReferenceAsList[WebServerNetworkAccessControlAttributes](c.ref.Append("web_server_network_access_control"))
}

func (c ConfigAttributes) WorkloadsConfig() terra.ListValue[WorkloadsConfigAttributes] {
	return terra.ReferenceAsList[WorkloadsConfigAttributes](c.ref.Append("workloads_config"))
}

type DataRetentionConfigAttributes struct {
	ref terra.Reference
}

func (drc DataRetentionConfigAttributes) InternalRef() (terra.Reference, error) {
	return drc.ref, nil
}

func (drc DataRetentionConfigAttributes) InternalWithRef(ref terra.Reference) DataRetentionConfigAttributes {
	return DataRetentionConfigAttributes{ref: ref}
}

func (drc DataRetentionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return drc.ref.InternalTokens()
}

func (drc DataRetentionConfigAttributes) TaskLogsRetentionConfig() terra.ListValue[TaskLogsRetentionConfigAttributes] {
	return terra.ReferenceAsList[TaskLogsRetentionConfigAttributes](drc.ref.Append("task_logs_retention_config"))
}

type TaskLogsRetentionConfigAttributes struct {
	ref terra.Reference
}

func (tlrc TaskLogsRetentionConfigAttributes) InternalRef() (terra.Reference, error) {
	return tlrc.ref, nil
}

func (tlrc TaskLogsRetentionConfigAttributes) InternalWithRef(ref terra.Reference) TaskLogsRetentionConfigAttributes {
	return TaskLogsRetentionConfigAttributes{ref: ref}
}

func (tlrc TaskLogsRetentionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tlrc.ref.InternalTokens()
}

func (tlrc TaskLogsRetentionConfigAttributes) StorageMode() terra.StringValue {
	return terra.ReferenceAsString(tlrc.ref.Append("storage_mode"))
}

type DatabaseConfigAttributes struct {
	ref terra.Reference
}

func (dc DatabaseConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DatabaseConfigAttributes) InternalWithRef(ref terra.Reference) DatabaseConfigAttributes {
	return DatabaseConfigAttributes{ref: ref}
}

func (dc DatabaseConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DatabaseConfigAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("machine_type"))
}

func (dc DatabaseConfigAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("zone"))
}

type EncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigAttributes {
	return EncryptionConfigAttributes{ref: ref}
}

func (ec EncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_name"))
}

type MaintenanceWindowAttributes struct {
	ref terra.Reference
}

func (mw MaintenanceWindowAttributes) InternalRef() (terra.Reference, error) {
	return mw.ref, nil
}

func (mw MaintenanceWindowAttributes) InternalWithRef(ref terra.Reference) MaintenanceWindowAttributes {
	return MaintenanceWindowAttributes{ref: ref}
}

func (mw MaintenanceWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mw.ref.InternalTokens()
}

func (mw MaintenanceWindowAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("end_time"))
}

func (mw MaintenanceWindowAttributes) Recurrence() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("recurrence"))
}

func (mw MaintenanceWindowAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("start_time"))
}

type MasterAuthorizedNetworksConfigAttributes struct {
	ref terra.Reference
}

func (manc MasterAuthorizedNetworksConfigAttributes) InternalRef() (terra.Reference, error) {
	return manc.ref, nil
}

func (manc MasterAuthorizedNetworksConfigAttributes) InternalWithRef(ref terra.Reference) MasterAuthorizedNetworksConfigAttributes {
	return MasterAuthorizedNetworksConfigAttributes{ref: ref}
}

func (manc MasterAuthorizedNetworksConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return manc.ref.InternalTokens()
}

func (manc MasterAuthorizedNetworksConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(manc.ref.Append("enabled"))
}

func (manc MasterAuthorizedNetworksConfigAttributes) CidrBlocks() terra.SetValue[CidrBlocksAttributes] {
	return terra.ReferenceAsSet[CidrBlocksAttributes](manc.ref.Append("cidr_blocks"))
}

type CidrBlocksAttributes struct {
	ref terra.Reference
}

func (cb CidrBlocksAttributes) InternalRef() (terra.Reference, error) {
	return cb.ref, nil
}

func (cb CidrBlocksAttributes) InternalWithRef(ref terra.Reference) CidrBlocksAttributes {
	return CidrBlocksAttributes{ref: ref}
}

func (cb CidrBlocksAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cb.ref.InternalTokens()
}

func (cb CidrBlocksAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(cb.ref.Append("cidr_block"))
}

func (cb CidrBlocksAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cb.ref.Append("display_name"))
}

type NodeConfigAttributes struct {
	ref terra.Reference
}

func (nc NodeConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NodeConfigAttributes) InternalWithRef(ref terra.Reference) NodeConfigAttributes {
	return NodeConfigAttributes{ref: ref}
}

func (nc NodeConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NodeConfigAttributes) ComposerInternalIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("composer_internal_ipv4_cidr_block"))
}

func (nc NodeConfigAttributes) ComposerNetworkAttachment() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("composer_network_attachment"))
}

func (nc NodeConfigAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(nc.ref.Append("disk_size_gb"))
}

func (nc NodeConfigAttributes) EnableIpMasqAgent() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("enable_ip_masq_agent"))
}

func (nc NodeConfigAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("machine_type"))
}

func (nc NodeConfigAttributes) MaxPodsPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(nc.ref.Append("max_pods_per_node"))
}

func (nc NodeConfigAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("network"))
}

func (nc NodeConfigAttributes) OauthScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("oauth_scopes"))
}

func (nc NodeConfigAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("service_account"))
}

func (nc NodeConfigAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("subnetwork"))
}

func (nc NodeConfigAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("tags"))
}

func (nc NodeConfigAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("zone"))
}

func (nc NodeConfigAttributes) IpAllocationPolicy() terra.ListValue[IpAllocationPolicyAttributes] {
	return terra.ReferenceAsList[IpAllocationPolicyAttributes](nc.ref.Append("ip_allocation_policy"))
}

type IpAllocationPolicyAttributes struct {
	ref terra.Reference
}

func (iap IpAllocationPolicyAttributes) InternalRef() (terra.Reference, error) {
	return iap.ref, nil
}

func (iap IpAllocationPolicyAttributes) InternalWithRef(ref terra.Reference) IpAllocationPolicyAttributes {
	return IpAllocationPolicyAttributes{ref: ref}
}

func (iap IpAllocationPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iap.ref.InternalTokens()
}

func (iap IpAllocationPolicyAttributes) ClusterIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(iap.ref.Append("cluster_ipv4_cidr_block"))
}

func (iap IpAllocationPolicyAttributes) ClusterSecondaryRangeName() terra.StringValue {
	return terra.ReferenceAsString(iap.ref.Append("cluster_secondary_range_name"))
}

func (iap IpAllocationPolicyAttributes) ServicesIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(iap.ref.Append("services_ipv4_cidr_block"))
}

func (iap IpAllocationPolicyAttributes) ServicesSecondaryRangeName() terra.StringValue {
	return terra.ReferenceAsString(iap.ref.Append("services_secondary_range_name"))
}

func (iap IpAllocationPolicyAttributes) UseIpAliases() terra.BoolValue {
	return terra.ReferenceAsBool(iap.ref.Append("use_ip_aliases"))
}

type PrivateEnvironmentConfigAttributes struct {
	ref terra.Reference
}

func (pec PrivateEnvironmentConfigAttributes) InternalRef() (terra.Reference, error) {
	return pec.ref, nil
}

func (pec PrivateEnvironmentConfigAttributes) InternalWithRef(ref terra.Reference) PrivateEnvironmentConfigAttributes {
	return PrivateEnvironmentConfigAttributes{ref: ref}
}

func (pec PrivateEnvironmentConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pec.ref.InternalTokens()
}

func (pec PrivateEnvironmentConfigAttributes) CloudComposerConnectionSubnetwork() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("cloud_composer_connection_subnetwork"))
}

func (pec PrivateEnvironmentConfigAttributes) CloudComposerNetworkIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("cloud_composer_network_ipv4_cidr_block"))
}

func (pec PrivateEnvironmentConfigAttributes) CloudSqlIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("cloud_sql_ipv4_cidr_block"))
}

func (pec PrivateEnvironmentConfigAttributes) ConnectionType() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("connection_type"))
}

func (pec PrivateEnvironmentConfigAttributes) EnablePrivateEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(pec.ref.Append("enable_private_endpoint"))
}

func (pec PrivateEnvironmentConfigAttributes) EnablePrivatelyUsedPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(pec.ref.Append("enable_privately_used_public_ips"))
}

func (pec PrivateEnvironmentConfigAttributes) MasterIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("master_ipv4_cidr_block"))
}

func (pec PrivateEnvironmentConfigAttributes) WebServerIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("web_server_ipv4_cidr_block"))
}

type RecoveryConfigAttributes struct {
	ref terra.Reference
}

func (rc RecoveryConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RecoveryConfigAttributes) InternalWithRef(ref terra.Reference) RecoveryConfigAttributes {
	return RecoveryConfigAttributes{ref: ref}
}

func (rc RecoveryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RecoveryConfigAttributes) ScheduledSnapshotsConfig() terra.ListValue[ScheduledSnapshotsConfigAttributes] {
	return terra.ReferenceAsList[ScheduledSnapshotsConfigAttributes](rc.ref.Append("scheduled_snapshots_config"))
}

type ScheduledSnapshotsConfigAttributes struct {
	ref terra.Reference
}

func (ssc ScheduledSnapshotsConfigAttributes) InternalRef() (terra.Reference, error) {
	return ssc.ref, nil
}

func (ssc ScheduledSnapshotsConfigAttributes) InternalWithRef(ref terra.Reference) ScheduledSnapshotsConfigAttributes {
	return ScheduledSnapshotsConfigAttributes{ref: ref}
}

func (ssc ScheduledSnapshotsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssc.ref.InternalTokens()
}

func (ssc ScheduledSnapshotsConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssc.ref.Append("enabled"))
}

func (ssc ScheduledSnapshotsConfigAttributes) SnapshotCreationSchedule() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("snapshot_creation_schedule"))
}

func (ssc ScheduledSnapshotsConfigAttributes) SnapshotLocation() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("snapshot_location"))
}

func (ssc ScheduledSnapshotsConfigAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("time_zone"))
}

type SoftwareConfigAttributes struct {
	ref terra.Reference
}

func (sc SoftwareConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SoftwareConfigAttributes) InternalWithRef(ref terra.Reference) SoftwareConfigAttributes {
	return SoftwareConfigAttributes{ref: ref}
}

func (sc SoftwareConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SoftwareConfigAttributes) AirflowConfigOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("airflow_config_overrides"))
}

func (sc SoftwareConfigAttributes) EnvVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("env_variables"))
}

func (sc SoftwareConfigAttributes) ImageVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("image_version"))
}

func (sc SoftwareConfigAttributes) PypiPackages() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("pypi_packages"))
}

func (sc SoftwareConfigAttributes) PythonVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("python_version"))
}

func (sc SoftwareConfigAttributes) SchedulerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("scheduler_count"))
}

func (sc SoftwareConfigAttributes) WebServerPluginsMode() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("web_server_plugins_mode"))
}

func (sc SoftwareConfigAttributes) CloudDataLineageIntegration() terra.ListValue[CloudDataLineageIntegrationAttributes] {
	return terra.ReferenceAsList[CloudDataLineageIntegrationAttributes](sc.ref.Append("cloud_data_lineage_integration"))
}

type CloudDataLineageIntegrationAttributes struct {
	ref terra.Reference
}

func (cdli CloudDataLineageIntegrationAttributes) InternalRef() (terra.Reference, error) {
	return cdli.ref, nil
}

func (cdli CloudDataLineageIntegrationAttributes) InternalWithRef(ref terra.Reference) CloudDataLineageIntegrationAttributes {
	return CloudDataLineageIntegrationAttributes{ref: ref}
}

func (cdli CloudDataLineageIntegrationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cdli.ref.InternalTokens()
}

func (cdli CloudDataLineageIntegrationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cdli.ref.Append("enabled"))
}

type WebServerConfigAttributes struct {
	ref terra.Reference
}

func (wsc WebServerConfigAttributes) InternalRef() (terra.Reference, error) {
	return wsc.ref, nil
}

func (wsc WebServerConfigAttributes) InternalWithRef(ref terra.Reference) WebServerConfigAttributes {
	return WebServerConfigAttributes{ref: ref}
}

func (wsc WebServerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wsc.ref.InternalTokens()
}

func (wsc WebServerConfigAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(wsc.ref.Append("machine_type"))
}

type WebServerNetworkAccessControlAttributes struct {
	ref terra.Reference
}

func (wsnac WebServerNetworkAccessControlAttributes) InternalRef() (terra.Reference, error) {
	return wsnac.ref, nil
}

func (wsnac WebServerNetworkAccessControlAttributes) InternalWithRef(ref terra.Reference) WebServerNetworkAccessControlAttributes {
	return WebServerNetworkAccessControlAttributes{ref: ref}
}

func (wsnac WebServerNetworkAccessControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wsnac.ref.InternalTokens()
}

func (wsnac WebServerNetworkAccessControlAttributes) AllowedIpRange() terra.SetValue[AllowedIpRangeAttributes] {
	return terra.ReferenceAsSet[AllowedIpRangeAttributes](wsnac.ref.Append("allowed_ip_range"))
}

type AllowedIpRangeAttributes struct {
	ref terra.Reference
}

func (air AllowedIpRangeAttributes) InternalRef() (terra.Reference, error) {
	return air.ref, nil
}

func (air AllowedIpRangeAttributes) InternalWithRef(ref terra.Reference) AllowedIpRangeAttributes {
	return AllowedIpRangeAttributes{ref: ref}
}

func (air AllowedIpRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return air.ref.InternalTokens()
}

func (air AllowedIpRangeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("description"))
}

func (air AllowedIpRangeAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("value"))
}

type WorkloadsConfigAttributes struct {
	ref terra.Reference
}

func (wc WorkloadsConfigAttributes) InternalRef() (terra.Reference, error) {
	return wc.ref, nil
}

func (wc WorkloadsConfigAttributes) InternalWithRef(ref terra.Reference) WorkloadsConfigAttributes {
	return WorkloadsConfigAttributes{ref: ref}
}

func (wc WorkloadsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wc.ref.InternalTokens()
}

func (wc WorkloadsConfigAttributes) DagProcessor() terra.ListValue[DagProcessorAttributes] {
	return terra.ReferenceAsList[DagProcessorAttributes](wc.ref.Append("dag_processor"))
}

func (wc WorkloadsConfigAttributes) Scheduler() terra.ListValue[SchedulerAttributes] {
	return terra.ReferenceAsList[SchedulerAttributes](wc.ref.Append("scheduler"))
}

func (wc WorkloadsConfigAttributes) Triggerer() terra.ListValue[TriggererAttributes] {
	return terra.ReferenceAsList[TriggererAttributes](wc.ref.Append("triggerer"))
}

func (wc WorkloadsConfigAttributes) WebServer() terra.ListValue[WebServerAttributes] {
	return terra.ReferenceAsList[WebServerAttributes](wc.ref.Append("web_server"))
}

func (wc WorkloadsConfigAttributes) Worker() terra.ListValue[WorkerAttributes] {
	return terra.ReferenceAsList[WorkerAttributes](wc.ref.Append("worker"))
}

type DagProcessorAttributes struct {
	ref terra.Reference
}

func (dp DagProcessorAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DagProcessorAttributes) InternalWithRef(ref terra.Reference) DagProcessorAttributes {
	return DagProcessorAttributes{ref: ref}
}

func (dp DagProcessorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DagProcessorAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("count"))
}

func (dp DagProcessorAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("cpu"))
}

func (dp DagProcessorAttributes) MemoryGb() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("memory_gb"))
}

func (dp DagProcessorAttributes) StorageGb() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("storage_gb"))
}

type SchedulerAttributes struct {
	ref terra.Reference
}

func (s SchedulerAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SchedulerAttributes) InternalWithRef(ref terra.Reference) SchedulerAttributes {
	return SchedulerAttributes{ref: ref}
}

func (s SchedulerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SchedulerAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("count"))
}

func (s SchedulerAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("cpu"))
}

func (s SchedulerAttributes) MemoryGb() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("memory_gb"))
}

func (s SchedulerAttributes) StorageGb() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("storage_gb"))
}

type TriggererAttributes struct {
	ref terra.Reference
}

func (t TriggererAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TriggererAttributes) InternalWithRef(ref terra.Reference) TriggererAttributes {
	return TriggererAttributes{ref: ref}
}

func (t TriggererAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TriggererAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("count"))
}

func (t TriggererAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("cpu"))
}

func (t TriggererAttributes) MemoryGb() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("memory_gb"))
}

type WebServerAttributes struct {
	ref terra.Reference
}

func (ws WebServerAttributes) InternalRef() (terra.Reference, error) {
	return ws.ref, nil
}

func (ws WebServerAttributes) InternalWithRef(ref terra.Reference) WebServerAttributes {
	return WebServerAttributes{ref: ref}
}

func (ws WebServerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ws.ref.InternalTokens()
}

func (ws WebServerAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(ws.ref.Append("cpu"))
}

func (ws WebServerAttributes) MemoryGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ws.ref.Append("memory_gb"))
}

func (ws WebServerAttributes) StorageGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ws.ref.Append("storage_gb"))
}

type WorkerAttributes struct {
	ref terra.Reference
}

func (w WorkerAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WorkerAttributes) InternalWithRef(ref terra.Reference) WorkerAttributes {
	return WorkerAttributes{ref: ref}
}

func (w WorkerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WorkerAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(w.ref.Append("cpu"))
}

func (w WorkerAttributes) MaxCount() terra.NumberValue {
	return terra.ReferenceAsNumber(w.ref.Append("max_count"))
}

func (w WorkerAttributes) MemoryGb() terra.NumberValue {
	return terra.ReferenceAsNumber(w.ref.Append("memory_gb"))
}

func (w WorkerAttributes) MinCount() terra.NumberValue {
	return terra.ReferenceAsNumber(w.ref.Append("min_count"))
}

func (w WorkerAttributes) StorageGb() terra.NumberValue {
	return terra.ReferenceAsNumber(w.ref.Append("storage_gb"))
}

type StorageConfigAttributes struct {
	ref terra.Reference
}

func (sc StorageConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc StorageConfigAttributes) InternalWithRef(ref terra.Reference) StorageConfigAttributes {
	return StorageConfigAttributes{ref: ref}
}

func (sc StorageConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc StorageConfigAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("bucket"))
}

type ConfigState struct {
	AirflowUri                     string                                `json:"airflow_uri"`
	DagGcsPrefix                   string                                `json:"dag_gcs_prefix"`
	EnablePrivateBuildsOnly        bool                                  `json:"enable_private_builds_only"`
	EnablePrivateEnvironment       bool                                  `json:"enable_private_environment"`
	EnvironmentSize                string                                `json:"environment_size"`
	GkeCluster                     string                                `json:"gke_cluster"`
	NodeCount                      float64                               `json:"node_count"`
	ResilienceMode                 string                                `json:"resilience_mode"`
	DataRetentionConfig            []DataRetentionConfigState            `json:"data_retention_config"`
	DatabaseConfig                 []DatabaseConfigState                 `json:"database_config"`
	EncryptionConfig               []EncryptionConfigState               `json:"encryption_config"`
	MaintenanceWindow              []MaintenanceWindowState              `json:"maintenance_window"`
	MasterAuthorizedNetworksConfig []MasterAuthorizedNetworksConfigState `json:"master_authorized_networks_config"`
	NodeConfig                     []NodeConfigState                     `json:"node_config"`
	PrivateEnvironmentConfig       []PrivateEnvironmentConfigState       `json:"private_environment_config"`
	RecoveryConfig                 []RecoveryConfigState                 `json:"recovery_config"`
	SoftwareConfig                 []SoftwareConfigState                 `json:"software_config"`
	WebServerConfig                []WebServerConfigState                `json:"web_server_config"`
	WebServerNetworkAccessControl  []WebServerNetworkAccessControlState  `json:"web_server_network_access_control"`
	WorkloadsConfig                []WorkloadsConfigState                `json:"workloads_config"`
}

type DataRetentionConfigState struct {
	TaskLogsRetentionConfig []TaskLogsRetentionConfigState `json:"task_logs_retention_config"`
}

type TaskLogsRetentionConfigState struct {
	StorageMode string `json:"storage_mode"`
}

type DatabaseConfigState struct {
	MachineType string `json:"machine_type"`
	Zone        string `json:"zone"`
}

type EncryptionConfigState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type MaintenanceWindowState struct {
	EndTime    string `json:"end_time"`
	Recurrence string `json:"recurrence"`
	StartTime  string `json:"start_time"`
}

type MasterAuthorizedNetworksConfigState struct {
	Enabled    bool              `json:"enabled"`
	CidrBlocks []CidrBlocksState `json:"cidr_blocks"`
}

type CidrBlocksState struct {
	CidrBlock   string `json:"cidr_block"`
	DisplayName string `json:"display_name"`
}

type NodeConfigState struct {
	ComposerInternalIpv4CidrBlock string                    `json:"composer_internal_ipv4_cidr_block"`
	ComposerNetworkAttachment     string                    `json:"composer_network_attachment"`
	DiskSizeGb                    float64                   `json:"disk_size_gb"`
	EnableIpMasqAgent             bool                      `json:"enable_ip_masq_agent"`
	MachineType                   string                    `json:"machine_type"`
	MaxPodsPerNode                float64                   `json:"max_pods_per_node"`
	Network                       string                    `json:"network"`
	OauthScopes                   []string                  `json:"oauth_scopes"`
	ServiceAccount                string                    `json:"service_account"`
	Subnetwork                    string                    `json:"subnetwork"`
	Tags                          []string                  `json:"tags"`
	Zone                          string                    `json:"zone"`
	IpAllocationPolicy            []IpAllocationPolicyState `json:"ip_allocation_policy"`
}

type IpAllocationPolicyState struct {
	ClusterIpv4CidrBlock       string `json:"cluster_ipv4_cidr_block"`
	ClusterSecondaryRangeName  string `json:"cluster_secondary_range_name"`
	ServicesIpv4CidrBlock      string `json:"services_ipv4_cidr_block"`
	ServicesSecondaryRangeName string `json:"services_secondary_range_name"`
	UseIpAliases               bool   `json:"use_ip_aliases"`
}

type PrivateEnvironmentConfigState struct {
	CloudComposerConnectionSubnetwork string `json:"cloud_composer_connection_subnetwork"`
	CloudComposerNetworkIpv4CidrBlock string `json:"cloud_composer_network_ipv4_cidr_block"`
	CloudSqlIpv4CidrBlock             string `json:"cloud_sql_ipv4_cidr_block"`
	ConnectionType                    string `json:"connection_type"`
	EnablePrivateEndpoint             bool   `json:"enable_private_endpoint"`
	EnablePrivatelyUsedPublicIps      bool   `json:"enable_privately_used_public_ips"`
	MasterIpv4CidrBlock               string `json:"master_ipv4_cidr_block"`
	WebServerIpv4CidrBlock            string `json:"web_server_ipv4_cidr_block"`
}

type RecoveryConfigState struct {
	ScheduledSnapshotsConfig []ScheduledSnapshotsConfigState `json:"scheduled_snapshots_config"`
}

type ScheduledSnapshotsConfigState struct {
	Enabled                  bool   `json:"enabled"`
	SnapshotCreationSchedule string `json:"snapshot_creation_schedule"`
	SnapshotLocation         string `json:"snapshot_location"`
	TimeZone                 string `json:"time_zone"`
}

type SoftwareConfigState struct {
	AirflowConfigOverrides      map[string]string                  `json:"airflow_config_overrides"`
	EnvVariables                map[string]string                  `json:"env_variables"`
	ImageVersion                string                             `json:"image_version"`
	PypiPackages                map[string]string                  `json:"pypi_packages"`
	PythonVersion               string                             `json:"python_version"`
	SchedulerCount              float64                            `json:"scheduler_count"`
	WebServerPluginsMode        string                             `json:"web_server_plugins_mode"`
	CloudDataLineageIntegration []CloudDataLineageIntegrationState `json:"cloud_data_lineage_integration"`
}

type CloudDataLineageIntegrationState struct {
	Enabled bool `json:"enabled"`
}

type WebServerConfigState struct {
	MachineType string `json:"machine_type"`
}

type WebServerNetworkAccessControlState struct {
	AllowedIpRange []AllowedIpRangeState `json:"allowed_ip_range"`
}

type AllowedIpRangeState struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

type WorkloadsConfigState struct {
	DagProcessor []DagProcessorState `json:"dag_processor"`
	Scheduler    []SchedulerState    `json:"scheduler"`
	Triggerer    []TriggererState    `json:"triggerer"`
	WebServer    []WebServerState    `json:"web_server"`
	Worker       []WorkerState       `json:"worker"`
}

type DagProcessorState struct {
	Count     float64 `json:"count"`
	Cpu       float64 `json:"cpu"`
	MemoryGb  float64 `json:"memory_gb"`
	StorageGb float64 `json:"storage_gb"`
}

type SchedulerState struct {
	Count     float64 `json:"count"`
	Cpu       float64 `json:"cpu"`
	MemoryGb  float64 `json:"memory_gb"`
	StorageGb float64 `json:"storage_gb"`
}

type TriggererState struct {
	Count    float64 `json:"count"`
	Cpu      float64 `json:"cpu"`
	MemoryGb float64 `json:"memory_gb"`
}

type WebServerState struct {
	Cpu       float64 `json:"cpu"`
	MemoryGb  float64 `json:"memory_gb"`
	StorageGb float64 `json:"storage_gb"`
}

type WorkerState struct {
	Cpu       float64 `json:"cpu"`
	MaxCount  float64 `json:"max_count"`
	MemoryGb  float64 `json:"memory_gb"`
	MinCount  float64 `json:"min_count"`
	StorageGb float64 `json:"storage_gb"`
}

type StorageConfigState struct {
	Bucket string `json:"bucket"`
}
