// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadataprocmetastoreservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfig struct{}

type HiveMetastoreConfig struct {
	// AuxiliaryVersions: min=0
	AuxiliaryVersions []AuxiliaryVersions `hcl:"auxiliary_versions,block" validate:"min=0"`
	// KerberosConfig: min=0
	KerberosConfig []KerberosConfig `hcl:"kerberos_config,block" validate:"min=0"`
}

type AuxiliaryVersions struct{}

type KerberosConfig struct {
	// Keytab: min=0
	Keytab []Keytab `hcl:"keytab,block" validate:"min=0"`
}

type Keytab struct{}

type MaintenanceWindow struct{}

type MetadataIntegration struct {
	// DataCatalogConfig: min=0
	DataCatalogConfig []DataCatalogConfig `hcl:"data_catalog_config,block" validate:"min=0"`
}

type DataCatalogConfig struct{}

type NetworkConfig struct {
	// Consumers: min=0
	Consumers []Consumers `hcl:"consumers,block" validate:"min=0"`
}

type Consumers struct{}

type TelemetryConfig struct{}

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

func (ec EncryptionConfigAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key"))
}

type HiveMetastoreConfigAttributes struct {
	ref terra.Reference
}

func (hmc HiveMetastoreConfigAttributes) InternalRef() (terra.Reference, error) {
	return hmc.ref, nil
}

func (hmc HiveMetastoreConfigAttributes) InternalWithRef(ref terra.Reference) HiveMetastoreConfigAttributes {
	return HiveMetastoreConfigAttributes{ref: ref}
}

func (hmc HiveMetastoreConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hmc.ref.InternalTokens()
}

func (hmc HiveMetastoreConfigAttributes) ConfigOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hmc.ref.Append("config_overrides"))
}

func (hmc HiveMetastoreConfigAttributes) EndpointProtocol() terra.StringValue {
	return terra.ReferenceAsString(hmc.ref.Append("endpoint_protocol"))
}

func (hmc HiveMetastoreConfigAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(hmc.ref.Append("version"))
}

func (hmc HiveMetastoreConfigAttributes) AuxiliaryVersions() terra.SetValue[AuxiliaryVersionsAttributes] {
	return terra.ReferenceAsSet[AuxiliaryVersionsAttributes](hmc.ref.Append("auxiliary_versions"))
}

func (hmc HiveMetastoreConfigAttributes) KerberosConfig() terra.ListValue[KerberosConfigAttributes] {
	return terra.ReferenceAsList[KerberosConfigAttributes](hmc.ref.Append("kerberos_config"))
}

type AuxiliaryVersionsAttributes struct {
	ref terra.Reference
}

func (av AuxiliaryVersionsAttributes) InternalRef() (terra.Reference, error) {
	return av.ref, nil
}

func (av AuxiliaryVersionsAttributes) InternalWithRef(ref terra.Reference) AuxiliaryVersionsAttributes {
	return AuxiliaryVersionsAttributes{ref: ref}
}

func (av AuxiliaryVersionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return av.ref.InternalTokens()
}

func (av AuxiliaryVersionsAttributes) ConfigOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](av.ref.Append("config_overrides"))
}

func (av AuxiliaryVersionsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(av.ref.Append("key"))
}

func (av AuxiliaryVersionsAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(av.ref.Append("version"))
}

type KerberosConfigAttributes struct {
	ref terra.Reference
}

func (kc KerberosConfigAttributes) InternalRef() (terra.Reference, error) {
	return kc.ref, nil
}

func (kc KerberosConfigAttributes) InternalWithRef(ref terra.Reference) KerberosConfigAttributes {
	return KerberosConfigAttributes{ref: ref}
}

func (kc KerberosConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kc.ref.InternalTokens()
}

func (kc KerberosConfigAttributes) Krb5ConfigGcsUri() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("krb5_config_gcs_uri"))
}

func (kc KerberosConfigAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("principal"))
}

func (kc KerberosConfigAttributes) Keytab() terra.ListValue[KeytabAttributes] {
	return terra.ReferenceAsList[KeytabAttributes](kc.ref.Append("keytab"))
}

type KeytabAttributes struct {
	ref terra.Reference
}

func (k KeytabAttributes) InternalRef() (terra.Reference, error) {
	return k.ref, nil
}

func (k KeytabAttributes) InternalWithRef(ref terra.Reference) KeytabAttributes {
	return KeytabAttributes{ref: ref}
}

func (k KeytabAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return k.ref.InternalTokens()
}

func (k KeytabAttributes) CloudSecret() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("cloud_secret"))
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

func (mw MaintenanceWindowAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("day_of_week"))
}

func (mw MaintenanceWindowAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("hour_of_day"))
}

type MetadataIntegrationAttributes struct {
	ref terra.Reference
}

func (mi MetadataIntegrationAttributes) InternalRef() (terra.Reference, error) {
	return mi.ref, nil
}

func (mi MetadataIntegrationAttributes) InternalWithRef(ref terra.Reference) MetadataIntegrationAttributes {
	return MetadataIntegrationAttributes{ref: ref}
}

func (mi MetadataIntegrationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mi.ref.InternalTokens()
}

func (mi MetadataIntegrationAttributes) DataCatalogConfig() terra.ListValue[DataCatalogConfigAttributes] {
	return terra.ReferenceAsList[DataCatalogConfigAttributes](mi.ref.Append("data_catalog_config"))
}

type DataCatalogConfigAttributes struct {
	ref terra.Reference
}

func (dcc DataCatalogConfigAttributes) InternalRef() (terra.Reference, error) {
	return dcc.ref, nil
}

func (dcc DataCatalogConfigAttributes) InternalWithRef(ref terra.Reference) DataCatalogConfigAttributes {
	return DataCatalogConfigAttributes{ref: ref}
}

func (dcc DataCatalogConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dcc.ref.InternalTokens()
}

func (dcc DataCatalogConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(dcc.ref.Append("enabled"))
}

type NetworkConfigAttributes struct {
	ref terra.Reference
}

func (nc NetworkConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NetworkConfigAttributes) InternalWithRef(ref terra.Reference) NetworkConfigAttributes {
	return NetworkConfigAttributes{ref: ref}
}

func (nc NetworkConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NetworkConfigAttributes) Consumers() terra.ListValue[ConsumersAttributes] {
	return terra.ReferenceAsList[ConsumersAttributes](nc.ref.Append("consumers"))
}

type ConsumersAttributes struct {
	ref terra.Reference
}

func (c ConsumersAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConsumersAttributes) InternalWithRef(ref terra.Reference) ConsumersAttributes {
	return ConsumersAttributes{ref: ref}
}

func (c ConsumersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConsumersAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("endpoint_uri"))
}

func (c ConsumersAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("subnetwork"))
}

type TelemetryConfigAttributes struct {
	ref terra.Reference
}

func (tc TelemetryConfigAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TelemetryConfigAttributes) InternalWithRef(ref terra.Reference) TelemetryConfigAttributes {
	return TelemetryConfigAttributes{ref: ref}
}

func (tc TelemetryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TelemetryConfigAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("log_format"))
}

type EncryptionConfigState struct {
	KmsKey string `json:"kms_key"`
}

type HiveMetastoreConfigState struct {
	ConfigOverrides   map[string]string        `json:"config_overrides"`
	EndpointProtocol  string                   `json:"endpoint_protocol"`
	Version           string                   `json:"version"`
	AuxiliaryVersions []AuxiliaryVersionsState `json:"auxiliary_versions"`
	KerberosConfig    []KerberosConfigState    `json:"kerberos_config"`
}

type AuxiliaryVersionsState struct {
	ConfigOverrides map[string]string `json:"config_overrides"`
	Key             string            `json:"key"`
	Version         string            `json:"version"`
}

type KerberosConfigState struct {
	Krb5ConfigGcsUri string        `json:"krb5_config_gcs_uri"`
	Principal        string        `json:"principal"`
	Keytab           []KeytabState `json:"keytab"`
}

type KeytabState struct {
	CloudSecret string `json:"cloud_secret"`
}

type MaintenanceWindowState struct {
	DayOfWeek string  `json:"day_of_week"`
	HourOfDay float64 `json:"hour_of_day"`
}

type MetadataIntegrationState struct {
	DataCatalogConfig []DataCatalogConfigState `json:"data_catalog_config"`
}

type DataCatalogConfigState struct {
	Enabled bool `json:"enabled"`
}

type NetworkConfigState struct {
	Consumers []ConsumersState `json:"consumers"`
}

type ConsumersState struct {
	EndpointUri string `json:"endpoint_uri"`
	Subnetwork  string `json:"subnetwork"`
}

type TelemetryConfigState struct {
	LogFormat string `json:"log_format"`
}
