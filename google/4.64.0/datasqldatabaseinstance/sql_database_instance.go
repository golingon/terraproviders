// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasqldatabaseinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Clone struct{}

type IpAddress struct{}

type ReplicaConfiguration struct{}

type RestoreBackupContext struct{}

type ServerCaCert struct{}

type Settings struct {
	// ActiveDirectoryConfig: min=0
	ActiveDirectoryConfig []ActiveDirectoryConfig `hcl:"active_directory_config,block" validate:"min=0"`
	// BackupConfiguration: min=0
	BackupConfiguration []BackupConfiguration `hcl:"backup_configuration,block" validate:"min=0"`
	// DatabaseFlags: min=0
	DatabaseFlags []DatabaseFlags `hcl:"database_flags,block" validate:"min=0"`
	// DenyMaintenancePeriod: min=0
	DenyMaintenancePeriod []DenyMaintenancePeriod `hcl:"deny_maintenance_period,block" validate:"min=0"`
	// InsightsConfig: min=0
	InsightsConfig []InsightsConfig `hcl:"insights_config,block" validate:"min=0"`
	// IpConfiguration: min=0
	IpConfiguration []IpConfiguration `hcl:"ip_configuration,block" validate:"min=0"`
	// LocationPreference: min=0
	LocationPreference []LocationPreference `hcl:"location_preference,block" validate:"min=0"`
	// MaintenanceWindow: min=0
	MaintenanceWindow []MaintenanceWindow `hcl:"maintenance_window,block" validate:"min=0"`
	// PasswordValidationPolicy: min=0
	PasswordValidationPolicy []PasswordValidationPolicy `hcl:"password_validation_policy,block" validate:"min=0"`
	// SqlServerAuditConfig: min=0
	SqlServerAuditConfig []SqlServerAuditConfig `hcl:"sql_server_audit_config,block" validate:"min=0"`
}

type ActiveDirectoryConfig struct{}

type BackupConfiguration struct {
	// BackupRetentionSettings: min=0
	BackupRetentionSettings []BackupRetentionSettings `hcl:"backup_retention_settings,block" validate:"min=0"`
}

type BackupRetentionSettings struct{}

type DatabaseFlags struct{}

type DenyMaintenancePeriod struct{}

type InsightsConfig struct{}

type IpConfiguration struct {
	// AuthorizedNetworks: min=0
	AuthorizedNetworks []AuthorizedNetworks `hcl:"authorized_networks,block" validate:"min=0"`
}

type AuthorizedNetworks struct{}

type LocationPreference struct{}

type MaintenanceWindow struct{}

type PasswordValidationPolicy struct{}

type SqlServerAuditConfig struct{}

type CloneAttributes struct {
	ref terra.Reference
}

func (c CloneAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CloneAttributes) InternalWithRef(ref terra.Reference) CloneAttributes {
	return CloneAttributes{ref: ref}
}

func (c CloneAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CloneAttributes) AllocatedIpRange() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("allocated_ip_range"))
}

func (c CloneAttributes) DatabaseNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("database_names"))
}

func (c CloneAttributes) PointInTime() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("point_in_time"))
}

func (c CloneAttributes) SourceInstanceName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("source_instance_name"))
}

type IpAddressAttributes struct {
	ref terra.Reference
}

func (ia IpAddressAttributes) InternalRef() (terra.Reference, error) {
	return ia.ref, nil
}

func (ia IpAddressAttributes) InternalWithRef(ref terra.Reference) IpAddressAttributes {
	return IpAddressAttributes{ref: ref}
}

func (ia IpAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ia.ref.InternalTokens()
}

func (ia IpAddressAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("ip_address"))
}

func (ia IpAddressAttributes) TimeToRetire() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("time_to_retire"))
}

func (ia IpAddressAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("type"))
}

type ReplicaConfigurationAttributes struct {
	ref terra.Reference
}

func (rc ReplicaConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReplicaConfigurationAttributes) InternalWithRef(ref terra.Reference) ReplicaConfigurationAttributes {
	return ReplicaConfigurationAttributes{ref: ref}
}

func (rc ReplicaConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReplicaConfigurationAttributes) CaCertificate() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("ca_certificate"))
}

func (rc ReplicaConfigurationAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("client_certificate"))
}

func (rc ReplicaConfigurationAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("client_key"))
}

func (rc ReplicaConfigurationAttributes) ConnectRetryInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("connect_retry_interval"))
}

func (rc ReplicaConfigurationAttributes) DumpFilePath() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("dump_file_path"))
}

func (rc ReplicaConfigurationAttributes) FailoverTarget() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("failover_target"))
}

func (rc ReplicaConfigurationAttributes) MasterHeartbeatPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("master_heartbeat_period"))
}

func (rc ReplicaConfigurationAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("password"))
}

func (rc ReplicaConfigurationAttributes) SslCipher() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("ssl_cipher"))
}

func (rc ReplicaConfigurationAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("username"))
}

func (rc ReplicaConfigurationAttributes) VerifyServerCertificate() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("verify_server_certificate"))
}

type RestoreBackupContextAttributes struct {
	ref terra.Reference
}

func (rbc RestoreBackupContextAttributes) InternalRef() (terra.Reference, error) {
	return rbc.ref, nil
}

func (rbc RestoreBackupContextAttributes) InternalWithRef(ref terra.Reference) RestoreBackupContextAttributes {
	return RestoreBackupContextAttributes{ref: ref}
}

func (rbc RestoreBackupContextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rbc.ref.InternalTokens()
}

func (rbc RestoreBackupContextAttributes) BackupRunId() terra.NumberValue {
	return terra.ReferenceAsNumber(rbc.ref.Append("backup_run_id"))
}

func (rbc RestoreBackupContextAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(rbc.ref.Append("instance_id"))
}

func (rbc RestoreBackupContextAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(rbc.ref.Append("project"))
}

type ServerCaCertAttributes struct {
	ref terra.Reference
}

func (scc ServerCaCertAttributes) InternalRef() (terra.Reference, error) {
	return scc.ref, nil
}

func (scc ServerCaCertAttributes) InternalWithRef(ref terra.Reference) ServerCaCertAttributes {
	return ServerCaCertAttributes{ref: ref}
}

func (scc ServerCaCertAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scc.ref.InternalTokens()
}

func (scc ServerCaCertAttributes) Cert() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("cert"))
}

func (scc ServerCaCertAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("common_name"))
}

func (scc ServerCaCertAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("create_time"))
}

func (scc ServerCaCertAttributes) ExpirationTime() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("expiration_time"))
}

func (scc ServerCaCertAttributes) Sha1Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("sha1_fingerprint"))
}

type SettingsAttributes struct {
	ref terra.Reference
}

func (s SettingsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SettingsAttributes) InternalWithRef(ref terra.Reference) SettingsAttributes {
	return SettingsAttributes{ref: ref}
}

func (s SettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SettingsAttributes) ActivationPolicy() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("activation_policy"))
}

func (s SettingsAttributes) AvailabilityType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("availability_type"))
}

func (s SettingsAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("collation"))
}

func (s SettingsAttributes) ConnectorEnforcement() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("connector_enforcement"))
}

func (s SettingsAttributes) DeletionProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("deletion_protection_enabled"))
}

func (s SettingsAttributes) DiskAutoresize() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("disk_autoresize"))
}

func (s SettingsAttributes) DiskAutoresizeLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("disk_autoresize_limit"))
}

func (s SettingsAttributes) DiskSize() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("disk_size"))
}

func (s SettingsAttributes) DiskType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("disk_type"))
}

func (s SettingsAttributes) PricingPlan() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("pricing_plan"))
}

func (s SettingsAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tier"))
}

func (s SettingsAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("time_zone"))
}

func (s SettingsAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("user_labels"))
}

func (s SettingsAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("version"))
}

func (s SettingsAttributes) ActiveDirectoryConfig() terra.ListValue[ActiveDirectoryConfigAttributes] {
	return terra.ReferenceAsList[ActiveDirectoryConfigAttributes](s.ref.Append("active_directory_config"))
}

func (s SettingsAttributes) BackupConfiguration() terra.ListValue[BackupConfigurationAttributes] {
	return terra.ReferenceAsList[BackupConfigurationAttributes](s.ref.Append("backup_configuration"))
}

func (s SettingsAttributes) DatabaseFlags() terra.ListValue[DatabaseFlagsAttributes] {
	return terra.ReferenceAsList[DatabaseFlagsAttributes](s.ref.Append("database_flags"))
}

func (s SettingsAttributes) DenyMaintenancePeriod() terra.ListValue[DenyMaintenancePeriodAttributes] {
	return terra.ReferenceAsList[DenyMaintenancePeriodAttributes](s.ref.Append("deny_maintenance_period"))
}

func (s SettingsAttributes) InsightsConfig() terra.ListValue[InsightsConfigAttributes] {
	return terra.ReferenceAsList[InsightsConfigAttributes](s.ref.Append("insights_config"))
}

func (s SettingsAttributes) IpConfiguration() terra.ListValue[IpConfigurationAttributes] {
	return terra.ReferenceAsList[IpConfigurationAttributes](s.ref.Append("ip_configuration"))
}

func (s SettingsAttributes) LocationPreference() terra.ListValue[LocationPreferenceAttributes] {
	return terra.ReferenceAsList[LocationPreferenceAttributes](s.ref.Append("location_preference"))
}

func (s SettingsAttributes) MaintenanceWindow() terra.ListValue[MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[MaintenanceWindowAttributes](s.ref.Append("maintenance_window"))
}

func (s SettingsAttributes) PasswordValidationPolicy() terra.ListValue[PasswordValidationPolicyAttributes] {
	return terra.ReferenceAsList[PasswordValidationPolicyAttributes](s.ref.Append("password_validation_policy"))
}

func (s SettingsAttributes) SqlServerAuditConfig() terra.ListValue[SqlServerAuditConfigAttributes] {
	return terra.ReferenceAsList[SqlServerAuditConfigAttributes](s.ref.Append("sql_server_audit_config"))
}

type ActiveDirectoryConfigAttributes struct {
	ref terra.Reference
}

func (adc ActiveDirectoryConfigAttributes) InternalRef() (terra.Reference, error) {
	return adc.ref, nil
}

func (adc ActiveDirectoryConfigAttributes) InternalWithRef(ref terra.Reference) ActiveDirectoryConfigAttributes {
	return ActiveDirectoryConfigAttributes{ref: ref}
}

func (adc ActiveDirectoryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return adc.ref.InternalTokens()
}

func (adc ActiveDirectoryConfigAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("domain"))
}

type BackupConfigurationAttributes struct {
	ref terra.Reference
}

func (bc BackupConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return bc.ref, nil
}

func (bc BackupConfigurationAttributes) InternalWithRef(ref terra.Reference) BackupConfigurationAttributes {
	return BackupConfigurationAttributes{ref: ref}
}

func (bc BackupConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bc.ref.InternalTokens()
}

func (bc BackupConfigurationAttributes) BinaryLogEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("binary_log_enabled"))
}

func (bc BackupConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("enabled"))
}

func (bc BackupConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("location"))
}

func (bc BackupConfigurationAttributes) PointInTimeRecoveryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("point_in_time_recovery_enabled"))
}

func (bc BackupConfigurationAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("start_time"))
}

func (bc BackupConfigurationAttributes) TransactionLogRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(bc.ref.Append("transaction_log_retention_days"))
}

func (bc BackupConfigurationAttributes) BackupRetentionSettings() terra.ListValue[BackupRetentionSettingsAttributes] {
	return terra.ReferenceAsList[BackupRetentionSettingsAttributes](bc.ref.Append("backup_retention_settings"))
}

type BackupRetentionSettingsAttributes struct {
	ref terra.Reference
}

func (brs BackupRetentionSettingsAttributes) InternalRef() (terra.Reference, error) {
	return brs.ref, nil
}

func (brs BackupRetentionSettingsAttributes) InternalWithRef(ref terra.Reference) BackupRetentionSettingsAttributes {
	return BackupRetentionSettingsAttributes{ref: ref}
}

func (brs BackupRetentionSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return brs.ref.InternalTokens()
}

func (brs BackupRetentionSettingsAttributes) RetainedBackups() terra.NumberValue {
	return terra.ReferenceAsNumber(brs.ref.Append("retained_backups"))
}

func (brs BackupRetentionSettingsAttributes) RetentionUnit() terra.StringValue {
	return terra.ReferenceAsString(brs.ref.Append("retention_unit"))
}

type DatabaseFlagsAttributes struct {
	ref terra.Reference
}

func (df DatabaseFlagsAttributes) InternalRef() (terra.Reference, error) {
	return df.ref, nil
}

func (df DatabaseFlagsAttributes) InternalWithRef(ref terra.Reference) DatabaseFlagsAttributes {
	return DatabaseFlagsAttributes{ref: ref}
}

func (df DatabaseFlagsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return df.ref.InternalTokens()
}

func (df DatabaseFlagsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("name"))
}

func (df DatabaseFlagsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("value"))
}

type DenyMaintenancePeriodAttributes struct {
	ref terra.Reference
}

func (dmp DenyMaintenancePeriodAttributes) InternalRef() (terra.Reference, error) {
	return dmp.ref, nil
}

func (dmp DenyMaintenancePeriodAttributes) InternalWithRef(ref terra.Reference) DenyMaintenancePeriodAttributes {
	return DenyMaintenancePeriodAttributes{ref: ref}
}

func (dmp DenyMaintenancePeriodAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dmp.ref.InternalTokens()
}

func (dmp DenyMaintenancePeriodAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("end_date"))
}

func (dmp DenyMaintenancePeriodAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("start_date"))
}

func (dmp DenyMaintenancePeriodAttributes) Time() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("time"))
}

type InsightsConfigAttributes struct {
	ref terra.Reference
}

func (ic InsightsConfigAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic InsightsConfigAttributes) InternalWithRef(ref terra.Reference) InsightsConfigAttributes {
	return InsightsConfigAttributes{ref: ref}
}

func (ic InsightsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic InsightsConfigAttributes) QueryInsightsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("query_insights_enabled"))
}

func (ic InsightsConfigAttributes) QueryPlansPerMinute() terra.NumberValue {
	return terra.ReferenceAsNumber(ic.ref.Append("query_plans_per_minute"))
}

func (ic InsightsConfigAttributes) QueryStringLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ic.ref.Append("query_string_length"))
}

func (ic InsightsConfigAttributes) RecordApplicationTags() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("record_application_tags"))
}

func (ic InsightsConfigAttributes) RecordClientAddress() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("record_client_address"))
}

type IpConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IpConfigurationAttributes) InternalWithRef(ref terra.Reference) IpConfigurationAttributes {
	return IpConfigurationAttributes{ref: ref}
}

func (ic IpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IpConfigurationAttributes) AllocatedIpRange() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("allocated_ip_range"))
}

func (ic IpConfigurationAttributes) EnablePrivatePathForGoogleCloudServices() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("enable_private_path_for_google_cloud_services"))
}

func (ic IpConfigurationAttributes) Ipv4Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("ipv4_enabled"))
}

func (ic IpConfigurationAttributes) PrivateNetwork() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("private_network"))
}

func (ic IpConfigurationAttributes) RequireSsl() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("require_ssl"))
}

func (ic IpConfigurationAttributes) AuthorizedNetworks() terra.SetValue[AuthorizedNetworksAttributes] {
	return terra.ReferenceAsSet[AuthorizedNetworksAttributes](ic.ref.Append("authorized_networks"))
}

type AuthorizedNetworksAttributes struct {
	ref terra.Reference
}

func (an AuthorizedNetworksAttributes) InternalRef() (terra.Reference, error) {
	return an.ref, nil
}

func (an AuthorizedNetworksAttributes) InternalWithRef(ref terra.Reference) AuthorizedNetworksAttributes {
	return AuthorizedNetworksAttributes{ref: ref}
}

func (an AuthorizedNetworksAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return an.ref.InternalTokens()
}

func (an AuthorizedNetworksAttributes) ExpirationTime() terra.StringValue {
	return terra.ReferenceAsString(an.ref.Append("expiration_time"))
}

func (an AuthorizedNetworksAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(an.ref.Append("name"))
}

func (an AuthorizedNetworksAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(an.ref.Append("value"))
}

type LocationPreferenceAttributes struct {
	ref terra.Reference
}

func (lp LocationPreferenceAttributes) InternalRef() (terra.Reference, error) {
	return lp.ref, nil
}

func (lp LocationPreferenceAttributes) InternalWithRef(ref terra.Reference) LocationPreferenceAttributes {
	return LocationPreferenceAttributes{ref: ref}
}

func (lp LocationPreferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lp.ref.InternalTokens()
}

func (lp LocationPreferenceAttributes) FollowGaeApplication() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("follow_gae_application"))
}

func (lp LocationPreferenceAttributes) SecondaryZone() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("secondary_zone"))
}

func (lp LocationPreferenceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("zone"))
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

func (mw MaintenanceWindowAttributes) Day() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("day"))
}

func (mw MaintenanceWindowAttributes) Hour() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("hour"))
}

func (mw MaintenanceWindowAttributes) UpdateTrack() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("update_track"))
}

type PasswordValidationPolicyAttributes struct {
	ref terra.Reference
}

func (pvp PasswordValidationPolicyAttributes) InternalRef() (terra.Reference, error) {
	return pvp.ref, nil
}

func (pvp PasswordValidationPolicyAttributes) InternalWithRef(ref terra.Reference) PasswordValidationPolicyAttributes {
	return PasswordValidationPolicyAttributes{ref: ref}
}

func (pvp PasswordValidationPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pvp.ref.InternalTokens()
}

func (pvp PasswordValidationPolicyAttributes) Complexity() terra.StringValue {
	return terra.ReferenceAsString(pvp.ref.Append("complexity"))
}

func (pvp PasswordValidationPolicyAttributes) DisallowUsernameSubstring() terra.BoolValue {
	return terra.ReferenceAsBool(pvp.ref.Append("disallow_username_substring"))
}

func (pvp PasswordValidationPolicyAttributes) EnablePasswordPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(pvp.ref.Append("enable_password_policy"))
}

func (pvp PasswordValidationPolicyAttributes) MinLength() terra.NumberValue {
	return terra.ReferenceAsNumber(pvp.ref.Append("min_length"))
}

func (pvp PasswordValidationPolicyAttributes) PasswordChangeInterval() terra.StringValue {
	return terra.ReferenceAsString(pvp.ref.Append("password_change_interval"))
}

func (pvp PasswordValidationPolicyAttributes) ReuseInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(pvp.ref.Append("reuse_interval"))
}

type SqlServerAuditConfigAttributes struct {
	ref terra.Reference
}

func (ssac SqlServerAuditConfigAttributes) InternalRef() (terra.Reference, error) {
	return ssac.ref, nil
}

func (ssac SqlServerAuditConfigAttributes) InternalWithRef(ref terra.Reference) SqlServerAuditConfigAttributes {
	return SqlServerAuditConfigAttributes{ref: ref}
}

func (ssac SqlServerAuditConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssac.ref.InternalTokens()
}

func (ssac SqlServerAuditConfigAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(ssac.ref.Append("bucket"))
}

func (ssac SqlServerAuditConfigAttributes) RetentionInterval() terra.StringValue {
	return terra.ReferenceAsString(ssac.ref.Append("retention_interval"))
}

func (ssac SqlServerAuditConfigAttributes) UploadInterval() terra.StringValue {
	return terra.ReferenceAsString(ssac.ref.Append("upload_interval"))
}

type CloneState struct {
	AllocatedIpRange   string   `json:"allocated_ip_range"`
	DatabaseNames      []string `json:"database_names"`
	PointInTime        string   `json:"point_in_time"`
	SourceInstanceName string   `json:"source_instance_name"`
}

type IpAddressState struct {
	IpAddress    string `json:"ip_address"`
	TimeToRetire string `json:"time_to_retire"`
	Type         string `json:"type"`
}

type ReplicaConfigurationState struct {
	CaCertificate           string  `json:"ca_certificate"`
	ClientCertificate       string  `json:"client_certificate"`
	ClientKey               string  `json:"client_key"`
	ConnectRetryInterval    float64 `json:"connect_retry_interval"`
	DumpFilePath            string  `json:"dump_file_path"`
	FailoverTarget          bool    `json:"failover_target"`
	MasterHeartbeatPeriod   float64 `json:"master_heartbeat_period"`
	Password                string  `json:"password"`
	SslCipher               string  `json:"ssl_cipher"`
	Username                string  `json:"username"`
	VerifyServerCertificate bool    `json:"verify_server_certificate"`
}

type RestoreBackupContextState struct {
	BackupRunId float64 `json:"backup_run_id"`
	InstanceId  string  `json:"instance_id"`
	Project     string  `json:"project"`
}

type ServerCaCertState struct {
	Cert            string `json:"cert"`
	CommonName      string `json:"common_name"`
	CreateTime      string `json:"create_time"`
	ExpirationTime  string `json:"expiration_time"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
}

type SettingsState struct {
	ActivationPolicy          string                          `json:"activation_policy"`
	AvailabilityType          string                          `json:"availability_type"`
	Collation                 string                          `json:"collation"`
	ConnectorEnforcement      string                          `json:"connector_enforcement"`
	DeletionProtectionEnabled bool                            `json:"deletion_protection_enabled"`
	DiskAutoresize            bool                            `json:"disk_autoresize"`
	DiskAutoresizeLimit       float64                         `json:"disk_autoresize_limit"`
	DiskSize                  float64                         `json:"disk_size"`
	DiskType                  string                          `json:"disk_type"`
	PricingPlan               string                          `json:"pricing_plan"`
	Tier                      string                          `json:"tier"`
	TimeZone                  string                          `json:"time_zone"`
	UserLabels                map[string]string               `json:"user_labels"`
	Version                   float64                         `json:"version"`
	ActiveDirectoryConfig     []ActiveDirectoryConfigState    `json:"active_directory_config"`
	BackupConfiguration       []BackupConfigurationState      `json:"backup_configuration"`
	DatabaseFlags             []DatabaseFlagsState            `json:"database_flags"`
	DenyMaintenancePeriod     []DenyMaintenancePeriodState    `json:"deny_maintenance_period"`
	InsightsConfig            []InsightsConfigState           `json:"insights_config"`
	IpConfiguration           []IpConfigurationState          `json:"ip_configuration"`
	LocationPreference        []LocationPreferenceState       `json:"location_preference"`
	MaintenanceWindow         []MaintenanceWindowState        `json:"maintenance_window"`
	PasswordValidationPolicy  []PasswordValidationPolicyState `json:"password_validation_policy"`
	SqlServerAuditConfig      []SqlServerAuditConfigState     `json:"sql_server_audit_config"`
}

type ActiveDirectoryConfigState struct {
	Domain string `json:"domain"`
}

type BackupConfigurationState struct {
	BinaryLogEnabled            bool                           `json:"binary_log_enabled"`
	Enabled                     bool                           `json:"enabled"`
	Location                    string                         `json:"location"`
	PointInTimeRecoveryEnabled  bool                           `json:"point_in_time_recovery_enabled"`
	StartTime                   string                         `json:"start_time"`
	TransactionLogRetentionDays float64                        `json:"transaction_log_retention_days"`
	BackupRetentionSettings     []BackupRetentionSettingsState `json:"backup_retention_settings"`
}

type BackupRetentionSettingsState struct {
	RetainedBackups float64 `json:"retained_backups"`
	RetentionUnit   string  `json:"retention_unit"`
}

type DatabaseFlagsState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DenyMaintenancePeriodState struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	Time      string `json:"time"`
}

type InsightsConfigState struct {
	QueryInsightsEnabled  bool    `json:"query_insights_enabled"`
	QueryPlansPerMinute   float64 `json:"query_plans_per_minute"`
	QueryStringLength     float64 `json:"query_string_length"`
	RecordApplicationTags bool    `json:"record_application_tags"`
	RecordClientAddress   bool    `json:"record_client_address"`
}

type IpConfigurationState struct {
	AllocatedIpRange                        string                    `json:"allocated_ip_range"`
	EnablePrivatePathForGoogleCloudServices bool                      `json:"enable_private_path_for_google_cloud_services"`
	Ipv4Enabled                             bool                      `json:"ipv4_enabled"`
	PrivateNetwork                          string                    `json:"private_network"`
	RequireSsl                              bool                      `json:"require_ssl"`
	AuthorizedNetworks                      []AuthorizedNetworksState `json:"authorized_networks"`
}

type AuthorizedNetworksState struct {
	ExpirationTime string `json:"expiration_time"`
	Name           string `json:"name"`
	Value          string `json:"value"`
}

type LocationPreferenceState struct {
	FollowGaeApplication string `json:"follow_gae_application"`
	SecondaryZone        string `json:"secondary_zone"`
	Zone                 string `json:"zone"`
}

type MaintenanceWindowState struct {
	Day         float64 `json:"day"`
	Hour        float64 `json:"hour"`
	UpdateTrack string  `json:"update_track"`
}

type PasswordValidationPolicyState struct {
	Complexity                string  `json:"complexity"`
	DisallowUsernameSubstring bool    `json:"disallow_username_substring"`
	EnablePasswordPolicy      bool    `json:"enable_password_policy"`
	MinLength                 float64 `json:"min_length"`
	PasswordChangeInterval    string  `json:"password_change_interval"`
	ReuseInterval             float64 `json:"reuse_interval"`
}

type SqlServerAuditConfigState struct {
	Bucket            string `json:"bucket"`
	RetentionInterval string `json:"retention_interval"`
	UploadInterval    string `json:"upload_interval"`
}
