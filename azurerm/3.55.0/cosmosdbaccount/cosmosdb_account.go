// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cosmosdbaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AnalyticalStorage struct {
	// SchemaType: string, required
	SchemaType terra.StringValue `hcl:"schema_type,attr" validate:"required"`
}

type Backup struct {
	// IntervalInMinutes: number, optional
	IntervalInMinutes terra.NumberValue `hcl:"interval_in_minutes,attr"`
	// RetentionInHours: number, optional
	RetentionInHours terra.NumberValue `hcl:"retention_in_hours,attr"`
	// StorageRedundancy: string, optional
	StorageRedundancy terra.StringValue `hcl:"storage_redundancy,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Capabilities struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Capacity struct {
	// TotalThroughputLimit: number, required
	TotalThroughputLimit terra.NumberValue `hcl:"total_throughput_limit,attr" validate:"required"`
}

type ConsistencyPolicy struct {
	// ConsistencyLevel: string, required
	ConsistencyLevel terra.StringValue `hcl:"consistency_level,attr" validate:"required"`
	// MaxIntervalInSeconds: number, optional
	MaxIntervalInSeconds terra.NumberValue `hcl:"max_interval_in_seconds,attr"`
	// MaxStalenessPrefix: number, optional
	MaxStalenessPrefix terra.NumberValue `hcl:"max_staleness_prefix,attr"`
}

type CorsRule struct {
	// AllowedHeaders: list of string, required
	AllowedHeaders terra.ListValue[terra.StringValue] `hcl:"allowed_headers,attr" validate:"required"`
	// AllowedMethods: list of string, required
	AllowedMethods terra.ListValue[terra.StringValue] `hcl:"allowed_methods,attr" validate:"required"`
	// AllowedOrigins: list of string, required
	AllowedOrigins terra.ListValue[terra.StringValue] `hcl:"allowed_origins,attr" validate:"required"`
	// ExposedHeaders: list of string, required
	ExposedHeaders terra.ListValue[terra.StringValue] `hcl:"exposed_headers,attr" validate:"required"`
	// MaxAgeInSeconds: number, required
	MaxAgeInSeconds terra.NumberValue `hcl:"max_age_in_seconds,attr" validate:"required"`
}

type GeoLocation struct {
	// FailoverPriority: number, required
	FailoverPriority terra.NumberValue `hcl:"failover_priority,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Restore struct {
	// RestoreTimestampInUtc: string, required
	RestoreTimestampInUtc terra.StringValue `hcl:"restore_timestamp_in_utc,attr" validate:"required"`
	// SourceCosmosdbAccountId: string, required
	SourceCosmosdbAccountId terra.StringValue `hcl:"source_cosmosdb_account_id,attr" validate:"required"`
	// Database: min=0
	Database []Database `hcl:"database,block" validate:"min=0"`
}

type Database struct {
	// CollectionNames: set of string, optional
	CollectionNames terra.SetValue[terra.StringValue] `hcl:"collection_names,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type VirtualNetworkRule struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// IgnoreMissingVnetServiceEndpoint: bool, optional
	IgnoreMissingVnetServiceEndpoint terra.BoolValue `hcl:"ignore_missing_vnet_service_endpoint,attr"`
}

type AnalyticalStorageAttributes struct {
	ref terra.Reference
}

func (as AnalyticalStorageAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AnalyticalStorageAttributes) InternalWithRef(ref terra.Reference) AnalyticalStorageAttributes {
	return AnalyticalStorageAttributes{ref: ref}
}

func (as AnalyticalStorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AnalyticalStorageAttributes) SchemaType() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("schema_type"))
}

type BackupAttributes struct {
	ref terra.Reference
}

func (b BackupAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BackupAttributes) InternalWithRef(ref terra.Reference) BackupAttributes {
	return BackupAttributes{ref: ref}
}

func (b BackupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BackupAttributes) IntervalInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("interval_in_minutes"))
}

func (b BackupAttributes) RetentionInHours() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("retention_in_hours"))
}

func (b BackupAttributes) StorageRedundancy() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("storage_redundancy"))
}

func (b BackupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("type"))
}

type CapabilitiesAttributes struct {
	ref terra.Reference
}

func (c CapabilitiesAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapabilitiesAttributes) InternalWithRef(ref terra.Reference) CapabilitiesAttributes {
	return CapabilitiesAttributes{ref: ref}
}

func (c CapabilitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapabilitiesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

type CapacityAttributes struct {
	ref terra.Reference
}

func (c CapacityAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapacityAttributes) InternalWithRef(ref terra.Reference) CapacityAttributes {
	return CapacityAttributes{ref: ref}
}

func (c CapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapacityAttributes) TotalThroughputLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("total_throughput_limit"))
}

type ConsistencyPolicyAttributes struct {
	ref terra.Reference
}

func (cp ConsistencyPolicyAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ConsistencyPolicyAttributes) InternalWithRef(ref terra.Reference) ConsistencyPolicyAttributes {
	return ConsistencyPolicyAttributes{ref: ref}
}

func (cp ConsistencyPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp ConsistencyPolicyAttributes) ConsistencyLevel() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("consistency_level"))
}

func (cp ConsistencyPolicyAttributes) MaxIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_interval_in_seconds"))
}

func (cp ConsistencyPolicyAttributes) MaxStalenessPrefix() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_staleness_prefix"))
}

type CorsRuleAttributes struct {
	ref terra.Reference
}

func (cr CorsRuleAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CorsRuleAttributes) InternalWithRef(ref terra.Reference) CorsRuleAttributes {
	return CorsRuleAttributes{ref: ref}
}

func (cr CorsRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CorsRuleAttributes) AllowedHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("allowed_headers"))
}

func (cr CorsRuleAttributes) AllowedMethods() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("allowed_methods"))
}

func (cr CorsRuleAttributes) AllowedOrigins() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("allowed_origins"))
}

func (cr CorsRuleAttributes) ExposedHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("exposed_headers"))
}

func (cr CorsRuleAttributes) MaxAgeInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cr.ref.Append("max_age_in_seconds"))
}

type GeoLocationAttributes struct {
	ref terra.Reference
}

func (gl GeoLocationAttributes) InternalRef() (terra.Reference, error) {
	return gl.ref, nil
}

func (gl GeoLocationAttributes) InternalWithRef(ref terra.Reference) GeoLocationAttributes {
	return GeoLocationAttributes{ref: ref}
}

func (gl GeoLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gl.ref.InternalTokens()
}

func (gl GeoLocationAttributes) FailoverPriority() terra.NumberValue {
	return terra.ReferenceAsNumber(gl.ref.Append("failover_priority"))
}

func (gl GeoLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gl.ref.Append("id"))
}

func (gl GeoLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gl.ref.Append("location"))
}

func (gl GeoLocationAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(gl.ref.Append("zone_redundant"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type RestoreAttributes struct {
	ref terra.Reference
}

func (r RestoreAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RestoreAttributes) InternalWithRef(ref terra.Reference) RestoreAttributes {
	return RestoreAttributes{ref: ref}
}

func (r RestoreAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RestoreAttributes) RestoreTimestampInUtc() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("restore_timestamp_in_utc"))
}

func (r RestoreAttributes) SourceCosmosdbAccountId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("source_cosmosdb_account_id"))
}

func (r RestoreAttributes) Database() terra.SetValue[DatabaseAttributes] {
	return terra.ReferenceAsSet[DatabaseAttributes](r.ref.Append("database"))
}

type DatabaseAttributes struct {
	ref terra.Reference
}

func (d DatabaseAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DatabaseAttributes) InternalWithRef(ref terra.Reference) DatabaseAttributes {
	return DatabaseAttributes{ref: ref}
}

func (d DatabaseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DatabaseAttributes) CollectionNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](d.ref.Append("collection_names"))
}

func (d DatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
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

type VirtualNetworkRuleAttributes struct {
	ref terra.Reference
}

func (vnr VirtualNetworkRuleAttributes) InternalRef() (terra.Reference, error) {
	return vnr.ref, nil
}

func (vnr VirtualNetworkRuleAttributes) InternalWithRef(ref terra.Reference) VirtualNetworkRuleAttributes {
	return VirtualNetworkRuleAttributes{ref: ref}
}

func (vnr VirtualNetworkRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vnr.ref.InternalTokens()
}

func (vnr VirtualNetworkRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnr.ref.Append("id"))
}

func (vnr VirtualNetworkRuleAttributes) IgnoreMissingVnetServiceEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(vnr.ref.Append("ignore_missing_vnet_service_endpoint"))
}

type AnalyticalStorageState struct {
	SchemaType string `json:"schema_type"`
}

type BackupState struct {
	IntervalInMinutes float64 `json:"interval_in_minutes"`
	RetentionInHours  float64 `json:"retention_in_hours"`
	StorageRedundancy string  `json:"storage_redundancy"`
	Type              string  `json:"type"`
}

type CapabilitiesState struct {
	Name string `json:"name"`
}

type CapacityState struct {
	TotalThroughputLimit float64 `json:"total_throughput_limit"`
}

type ConsistencyPolicyState struct {
	ConsistencyLevel     string  `json:"consistency_level"`
	MaxIntervalInSeconds float64 `json:"max_interval_in_seconds"`
	MaxStalenessPrefix   float64 `json:"max_staleness_prefix"`
}

type CorsRuleState struct {
	AllowedHeaders  []string `json:"allowed_headers"`
	AllowedMethods  []string `json:"allowed_methods"`
	AllowedOrigins  []string `json:"allowed_origins"`
	ExposedHeaders  []string `json:"exposed_headers"`
	MaxAgeInSeconds float64  `json:"max_age_in_seconds"`
}

type GeoLocationState struct {
	FailoverPriority float64 `json:"failover_priority"`
	Id               string  `json:"id"`
	Location         string  `json:"location"`
	ZoneRedundant    bool    `json:"zone_redundant"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type RestoreState struct {
	RestoreTimestampInUtc   string          `json:"restore_timestamp_in_utc"`
	SourceCosmosdbAccountId string          `json:"source_cosmosdb_account_id"`
	Database                []DatabaseState `json:"database"`
}

type DatabaseState struct {
	CollectionNames []string `json:"collection_names"`
	Name            string   `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type VirtualNetworkRuleState struct {
	Id                               string `json:"id"`
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignore_missing_vnet_service_endpoint"`
}
