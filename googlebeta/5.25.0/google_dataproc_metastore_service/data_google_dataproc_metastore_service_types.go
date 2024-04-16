// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_metastore_service

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataEncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec DataEncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec DataEncryptionConfigAttributes) InternalWithRef(ref terra.Reference) DataEncryptionConfigAttributes {
	return DataEncryptionConfigAttributes{ref: ref}
}

func (ec DataEncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec DataEncryptionConfigAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key"))
}

type DataHiveMetastoreConfigAttributes struct {
	ref terra.Reference
}

func (hmc DataHiveMetastoreConfigAttributes) InternalRef() (terra.Reference, error) {
	return hmc.ref, nil
}

func (hmc DataHiveMetastoreConfigAttributes) InternalWithRef(ref terra.Reference) DataHiveMetastoreConfigAttributes {
	return DataHiveMetastoreConfigAttributes{ref: ref}
}

func (hmc DataHiveMetastoreConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hmc.ref.InternalTokens()
}

func (hmc DataHiveMetastoreConfigAttributes) ConfigOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hmc.ref.Append("config_overrides"))
}

func (hmc DataHiveMetastoreConfigAttributes) EndpointProtocol() terra.StringValue {
	return terra.ReferenceAsString(hmc.ref.Append("endpoint_protocol"))
}

func (hmc DataHiveMetastoreConfigAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(hmc.ref.Append("version"))
}

func (hmc DataHiveMetastoreConfigAttributes) AuxiliaryVersions() terra.SetValue[DataHiveMetastoreConfigAuxiliaryVersionsAttributes] {
	return terra.ReferenceAsSet[DataHiveMetastoreConfigAuxiliaryVersionsAttributes](hmc.ref.Append("auxiliary_versions"))
}

func (hmc DataHiveMetastoreConfigAttributes) KerberosConfig() terra.ListValue[DataHiveMetastoreConfigKerberosConfigAttributes] {
	return terra.ReferenceAsList[DataHiveMetastoreConfigKerberosConfigAttributes](hmc.ref.Append("kerberos_config"))
}

type DataHiveMetastoreConfigAuxiliaryVersionsAttributes struct {
	ref terra.Reference
}

func (av DataHiveMetastoreConfigAuxiliaryVersionsAttributes) InternalRef() (terra.Reference, error) {
	return av.ref, nil
}

func (av DataHiveMetastoreConfigAuxiliaryVersionsAttributes) InternalWithRef(ref terra.Reference) DataHiveMetastoreConfigAuxiliaryVersionsAttributes {
	return DataHiveMetastoreConfigAuxiliaryVersionsAttributes{ref: ref}
}

func (av DataHiveMetastoreConfigAuxiliaryVersionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return av.ref.InternalTokens()
}

func (av DataHiveMetastoreConfigAuxiliaryVersionsAttributes) ConfigOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](av.ref.Append("config_overrides"))
}

func (av DataHiveMetastoreConfigAuxiliaryVersionsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(av.ref.Append("key"))
}

func (av DataHiveMetastoreConfigAuxiliaryVersionsAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(av.ref.Append("version"))
}

type DataHiveMetastoreConfigKerberosConfigAttributes struct {
	ref terra.Reference
}

func (kc DataHiveMetastoreConfigKerberosConfigAttributes) InternalRef() (terra.Reference, error) {
	return kc.ref, nil
}

func (kc DataHiveMetastoreConfigKerberosConfigAttributes) InternalWithRef(ref terra.Reference) DataHiveMetastoreConfigKerberosConfigAttributes {
	return DataHiveMetastoreConfigKerberosConfigAttributes{ref: ref}
}

func (kc DataHiveMetastoreConfigKerberosConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kc.ref.InternalTokens()
}

func (kc DataHiveMetastoreConfigKerberosConfigAttributes) Krb5ConfigGcsUri() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("krb5_config_gcs_uri"))
}

func (kc DataHiveMetastoreConfigKerberosConfigAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("principal"))
}

func (kc DataHiveMetastoreConfigKerberosConfigAttributes) Keytab() terra.ListValue[DataHiveMetastoreConfigKerberosConfigKeytabAttributes] {
	return terra.ReferenceAsList[DataHiveMetastoreConfigKerberosConfigKeytabAttributes](kc.ref.Append("keytab"))
}

type DataHiveMetastoreConfigKerberosConfigKeytabAttributes struct {
	ref terra.Reference
}

func (k DataHiveMetastoreConfigKerberosConfigKeytabAttributes) InternalRef() (terra.Reference, error) {
	return k.ref, nil
}

func (k DataHiveMetastoreConfigKerberosConfigKeytabAttributes) InternalWithRef(ref terra.Reference) DataHiveMetastoreConfigKerberosConfigKeytabAttributes {
	return DataHiveMetastoreConfigKerberosConfigKeytabAttributes{ref: ref}
}

func (k DataHiveMetastoreConfigKerberosConfigKeytabAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return k.ref.InternalTokens()
}

func (k DataHiveMetastoreConfigKerberosConfigKeytabAttributes) CloudSecret() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("cloud_secret"))
}

type DataMaintenanceWindowAttributes struct {
	ref terra.Reference
}

func (mw DataMaintenanceWindowAttributes) InternalRef() (terra.Reference, error) {
	return mw.ref, nil
}

func (mw DataMaintenanceWindowAttributes) InternalWithRef(ref terra.Reference) DataMaintenanceWindowAttributes {
	return DataMaintenanceWindowAttributes{ref: ref}
}

func (mw DataMaintenanceWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mw.ref.InternalTokens()
}

func (mw DataMaintenanceWindowAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("day_of_week"))
}

func (mw DataMaintenanceWindowAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("hour_of_day"))
}

type DataMetadataIntegrationAttributes struct {
	ref terra.Reference
}

func (mi DataMetadataIntegrationAttributes) InternalRef() (terra.Reference, error) {
	return mi.ref, nil
}

func (mi DataMetadataIntegrationAttributes) InternalWithRef(ref terra.Reference) DataMetadataIntegrationAttributes {
	return DataMetadataIntegrationAttributes{ref: ref}
}

func (mi DataMetadataIntegrationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mi.ref.InternalTokens()
}

func (mi DataMetadataIntegrationAttributes) DataCatalogConfig() terra.ListValue[DataMetadataIntegrationDataCatalogConfigAttributes] {
	return terra.ReferenceAsList[DataMetadataIntegrationDataCatalogConfigAttributes](mi.ref.Append("data_catalog_config"))
}

type DataMetadataIntegrationDataCatalogConfigAttributes struct {
	ref terra.Reference
}

func (dcc DataMetadataIntegrationDataCatalogConfigAttributes) InternalRef() (terra.Reference, error) {
	return dcc.ref, nil
}

func (dcc DataMetadataIntegrationDataCatalogConfigAttributes) InternalWithRef(ref terra.Reference) DataMetadataIntegrationDataCatalogConfigAttributes {
	return DataMetadataIntegrationDataCatalogConfigAttributes{ref: ref}
}

func (dcc DataMetadataIntegrationDataCatalogConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dcc.ref.InternalTokens()
}

func (dcc DataMetadataIntegrationDataCatalogConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(dcc.ref.Append("enabled"))
}

type DataNetworkConfigAttributes struct {
	ref terra.Reference
}

func (nc DataNetworkConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc DataNetworkConfigAttributes) InternalWithRef(ref terra.Reference) DataNetworkConfigAttributes {
	return DataNetworkConfigAttributes{ref: ref}
}

func (nc DataNetworkConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc DataNetworkConfigAttributes) CustomRoutesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("custom_routes_enabled"))
}

func (nc DataNetworkConfigAttributes) Consumers() terra.ListValue[DataNetworkConfigConsumersAttributes] {
	return terra.ReferenceAsList[DataNetworkConfigConsumersAttributes](nc.ref.Append("consumers"))
}

type DataNetworkConfigConsumersAttributes struct {
	ref terra.Reference
}

func (c DataNetworkConfigConsumersAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataNetworkConfigConsumersAttributes) InternalWithRef(ref terra.Reference) DataNetworkConfigConsumersAttributes {
	return DataNetworkConfigConsumersAttributes{ref: ref}
}

func (c DataNetworkConfigConsumersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataNetworkConfigConsumersAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("endpoint_uri"))
}

func (c DataNetworkConfigConsumersAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("subnetwork"))
}

type DataScalingConfigAttributes struct {
	ref terra.Reference
}

func (sc DataScalingConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc DataScalingConfigAttributes) InternalWithRef(ref terra.Reference) DataScalingConfigAttributes {
	return DataScalingConfigAttributes{ref: ref}
}

func (sc DataScalingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc DataScalingConfigAttributes) InstanceSize() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("instance_size"))
}

func (sc DataScalingConfigAttributes) ScalingFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("scaling_factor"))
}

type DataScheduledBackupAttributes struct {
	ref terra.Reference
}

func (sb DataScheduledBackupAttributes) InternalRef() (terra.Reference, error) {
	return sb.ref, nil
}

func (sb DataScheduledBackupAttributes) InternalWithRef(ref terra.Reference) DataScheduledBackupAttributes {
	return DataScheduledBackupAttributes{ref: ref}
}

func (sb DataScheduledBackupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sb.ref.InternalTokens()
}

func (sb DataScheduledBackupAttributes) BackupLocation() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("backup_location"))
}

func (sb DataScheduledBackupAttributes) CronSchedule() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("cron_schedule"))
}

func (sb DataScheduledBackupAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("enabled"))
}

func (sb DataScheduledBackupAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("time_zone"))
}

type DataTelemetryConfigAttributes struct {
	ref terra.Reference
}

func (tc DataTelemetryConfigAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc DataTelemetryConfigAttributes) InternalWithRef(ref terra.Reference) DataTelemetryConfigAttributes {
	return DataTelemetryConfigAttributes{ref: ref}
}

func (tc DataTelemetryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc DataTelemetryConfigAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("log_format"))
}

type DataEncryptionConfigState struct {
	KmsKey string `json:"kms_key"`
}

type DataHiveMetastoreConfigState struct {
	ConfigOverrides   map[string]string                               `json:"config_overrides"`
	EndpointProtocol  string                                          `json:"endpoint_protocol"`
	Version           string                                          `json:"version"`
	AuxiliaryVersions []DataHiveMetastoreConfigAuxiliaryVersionsState `json:"auxiliary_versions"`
	KerberosConfig    []DataHiveMetastoreConfigKerberosConfigState    `json:"kerberos_config"`
}

type DataHiveMetastoreConfigAuxiliaryVersionsState struct {
	ConfigOverrides map[string]string `json:"config_overrides"`
	Key             string            `json:"key"`
	Version         string            `json:"version"`
}

type DataHiveMetastoreConfigKerberosConfigState struct {
	Krb5ConfigGcsUri string                                             `json:"krb5_config_gcs_uri"`
	Principal        string                                             `json:"principal"`
	Keytab           []DataHiveMetastoreConfigKerberosConfigKeytabState `json:"keytab"`
}

type DataHiveMetastoreConfigKerberosConfigKeytabState struct {
	CloudSecret string `json:"cloud_secret"`
}

type DataMaintenanceWindowState struct {
	DayOfWeek string  `json:"day_of_week"`
	HourOfDay float64 `json:"hour_of_day"`
}

type DataMetadataIntegrationState struct {
	DataCatalogConfig []DataMetadataIntegrationDataCatalogConfigState `json:"data_catalog_config"`
}

type DataMetadataIntegrationDataCatalogConfigState struct {
	Enabled bool `json:"enabled"`
}

type DataNetworkConfigState struct {
	CustomRoutesEnabled bool                              `json:"custom_routes_enabled"`
	Consumers           []DataNetworkConfigConsumersState `json:"consumers"`
}

type DataNetworkConfigConsumersState struct {
	EndpointUri string `json:"endpoint_uri"`
	Subnetwork  string `json:"subnetwork"`
}

type DataScalingConfigState struct {
	InstanceSize  string  `json:"instance_size"`
	ScalingFactor float64 `json:"scaling_factor"`
}

type DataScheduledBackupState struct {
	BackupLocation string `json:"backup_location"`
	CronSchedule   string `json:"cron_schedule"`
	Enabled        bool   `json:"enabled"`
	TimeZone       string `json:"time_zone"`
}

type DataTelemetryConfigState struct {
	LogFormat string `json:"log_format"`
}
