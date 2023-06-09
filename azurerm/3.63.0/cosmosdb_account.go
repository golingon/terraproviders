// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbaccount "github.com/golingon/terraproviders/azurerm/3.63.0/cosmosdbaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbAccount creates a new instance of [CosmosdbAccount].
func NewCosmosdbAccount(name string, args CosmosdbAccountArgs) *CosmosdbAccount {
	return &CosmosdbAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbAccount)(nil)

// CosmosdbAccount represents the Terraform resource azurerm_cosmosdb_account.
type CosmosdbAccount struct {
	Name      string
	Args      CosmosdbAccountArgs
	state     *cosmosdbAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbAccount].
func (ca *CosmosdbAccount) Type() string {
	return "azurerm_cosmosdb_account"
}

// LocalName returns the local name for [CosmosdbAccount].
func (ca *CosmosdbAccount) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [CosmosdbAccount].
func (ca *CosmosdbAccount) Configuration() interface{} {
	return ca.Args
}

// DependOn is used for other resources to depend on [CosmosdbAccount].
func (ca *CosmosdbAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(ca)
}

// Dependencies returns the list of resources [CosmosdbAccount] depends_on.
func (ca *CosmosdbAccount) Dependencies() terra.Dependencies {
	return ca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbAccount].
func (ca *CosmosdbAccount) LifecycleManagement() *terra.Lifecycle {
	return ca.Lifecycle
}

// Attributes returns the attributes for [CosmosdbAccount].
func (ca *CosmosdbAccount) Attributes() cosmosdbAccountAttributes {
	return cosmosdbAccountAttributes{ref: terra.ReferenceResource(ca)}
}

// ImportState imports the given attribute values into [CosmosdbAccount]'s state.
func (ca *CosmosdbAccount) ImportState(av io.Reader) error {
	ca.state = &cosmosdbAccountState{}
	if err := json.NewDecoder(av).Decode(ca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ca.Type(), ca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbAccount] has state.
func (ca *CosmosdbAccount) State() (*cosmosdbAccountState, bool) {
	return ca.state, ca.state != nil
}

// StateMust returns the state for [CosmosdbAccount]. Panics if the state is nil.
func (ca *CosmosdbAccount) StateMust() *cosmosdbAccountState {
	if ca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ca.Type(), ca.LocalName()))
	}
	return ca.state
}

// CosmosdbAccountArgs contains the configurations for azurerm_cosmosdb_account.
type CosmosdbAccountArgs struct {
	// AccessKeyMetadataWritesEnabled: bool, optional
	AccessKeyMetadataWritesEnabled terra.BoolValue `hcl:"access_key_metadata_writes_enabled,attr"`
	// AnalyticalStorageEnabled: bool, optional
	AnalyticalStorageEnabled terra.BoolValue `hcl:"analytical_storage_enabled,attr"`
	// CreateMode: string, optional
	CreateMode terra.StringValue `hcl:"create_mode,attr"`
	// DefaultIdentityType: string, optional
	DefaultIdentityType terra.StringValue `hcl:"default_identity_type,attr"`
	// EnableAutomaticFailover: bool, optional
	EnableAutomaticFailover terra.BoolValue `hcl:"enable_automatic_failover,attr"`
	// EnableFreeTier: bool, optional
	EnableFreeTier terra.BoolValue `hcl:"enable_free_tier,attr"`
	// EnableMultipleWriteLocations: bool, optional
	EnableMultipleWriteLocations terra.BoolValue `hcl:"enable_multiple_write_locations,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpRangeFilter: string, optional
	IpRangeFilter terra.StringValue `hcl:"ip_range_filter,attr"`
	// IsVirtualNetworkFilterEnabled: bool, optional
	IsVirtualNetworkFilterEnabled terra.BoolValue `hcl:"is_virtual_network_filter_enabled,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// Kind: string, optional
	Kind terra.StringValue `hcl:"kind,attr"`
	// LocalAuthenticationDisabled: bool, optional
	LocalAuthenticationDisabled terra.BoolValue `hcl:"local_authentication_disabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MongoServerVersion: string, optional
	MongoServerVersion terra.StringValue `hcl:"mongo_server_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkAclBypassForAzureServices: bool, optional
	NetworkAclBypassForAzureServices terra.BoolValue `hcl:"network_acl_bypass_for_azure_services,attr"`
	// NetworkAclBypassIds: list of string, optional
	NetworkAclBypassIds terra.ListValue[terra.StringValue] `hcl:"network_acl_bypass_ids,attr"`
	// OfferType: string, required
	OfferType terra.StringValue `hcl:"offer_type,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AnalyticalStorage: optional
	AnalyticalStorage *cosmosdbaccount.AnalyticalStorage `hcl:"analytical_storage,block"`
	// Backup: optional
	Backup *cosmosdbaccount.Backup `hcl:"backup,block"`
	// Capabilities: min=0
	Capabilities []cosmosdbaccount.Capabilities `hcl:"capabilities,block" validate:"min=0"`
	// Capacity: optional
	Capacity *cosmosdbaccount.Capacity `hcl:"capacity,block"`
	// ConsistencyPolicy: required
	ConsistencyPolicy *cosmosdbaccount.ConsistencyPolicy `hcl:"consistency_policy,block" validate:"required"`
	// CorsRule: optional
	CorsRule *cosmosdbaccount.CorsRule `hcl:"cors_rule,block"`
	// GeoLocation: min=1
	GeoLocation []cosmosdbaccount.GeoLocation `hcl:"geo_location,block" validate:"min=1"`
	// Identity: optional
	Identity *cosmosdbaccount.Identity `hcl:"identity,block"`
	// Restore: optional
	Restore *cosmosdbaccount.Restore `hcl:"restore,block"`
	// Timeouts: optional
	Timeouts *cosmosdbaccount.Timeouts `hcl:"timeouts,block"`
	// VirtualNetworkRule: min=0
	VirtualNetworkRule []cosmosdbaccount.VirtualNetworkRule `hcl:"virtual_network_rule,block" validate:"min=0"`
}
type cosmosdbAccountAttributes struct {
	ref terra.Reference
}

// AccessKeyMetadataWritesEnabled returns a reference to field access_key_metadata_writes_enabled of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) AccessKeyMetadataWritesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("access_key_metadata_writes_enabled"))
}

// AnalyticalStorageEnabled returns a reference to field analytical_storage_enabled of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) AnalyticalStorageEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("analytical_storage_enabled"))
}

// ConnectionStrings returns a reference to field connection_strings of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) ConnectionStrings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("connection_strings"))
}

// CreateMode returns a reference to field create_mode of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("create_mode"))
}

// DefaultIdentityType returns a reference to field default_identity_type of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) DefaultIdentityType() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("default_identity_type"))
}

// EnableAutomaticFailover returns a reference to field enable_automatic_failover of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) EnableAutomaticFailover() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("enable_automatic_failover"))
}

// EnableFreeTier returns a reference to field enable_free_tier of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) EnableFreeTier() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("enable_free_tier"))
}

// EnableMultipleWriteLocations returns a reference to field enable_multiple_write_locations of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) EnableMultipleWriteLocations() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("enable_multiple_write_locations"))
}

// Endpoint returns a reference to field endpoint of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// IpRangeFilter returns a reference to field ip_range_filter of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) IpRangeFilter() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("ip_range_filter"))
}

// IsVirtualNetworkFilterEnabled returns a reference to field is_virtual_network_filter_enabled of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) IsVirtualNetworkFilterEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("is_virtual_network_filter_enabled"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("key_vault_key_id"))
}

// Kind returns a reference to field kind of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("kind"))
}

// LocalAuthenticationDisabled returns a reference to field local_authentication_disabled of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) LocalAuthenticationDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("local_authentication_disabled"))
}

// Location returns a reference to field location of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("location"))
}

// MongoServerVersion returns a reference to field mongo_server_version of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) MongoServerVersion() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("mongo_server_version"))
}

// Name returns a reference to field name of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// NetworkAclBypassForAzureServices returns a reference to field network_acl_bypass_for_azure_services of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) NetworkAclBypassForAzureServices() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("network_acl_bypass_for_azure_services"))
}

// NetworkAclBypassIds returns a reference to field network_acl_bypass_ids of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) NetworkAclBypassIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("network_acl_bypass_ids"))
}

// OfferType returns a reference to field offer_type of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) OfferType() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("offer_type"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("primary_key"))
}

// PrimaryReadonlyKey returns a reference to field primary_readonly_key of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) PrimaryReadonlyKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("primary_readonly_key"))
}

// PrimaryReadonlySqlConnectionString returns a reference to field primary_readonly_sql_connection_string of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) PrimaryReadonlySqlConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("primary_readonly_sql_connection_string"))
}

// PrimarySqlConnectionString returns a reference to field primary_sql_connection_string of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) PrimarySqlConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("primary_sql_connection_string"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("public_network_access_enabled"))
}

// ReadEndpoints returns a reference to field read_endpoints of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) ReadEndpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("read_endpoints"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("resource_group_name"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("secondary_key"))
}

// SecondaryReadonlyKey returns a reference to field secondary_readonly_key of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) SecondaryReadonlyKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("secondary_readonly_key"))
}

// SecondaryReadonlySqlConnectionString returns a reference to field secondary_readonly_sql_connection_string of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) SecondaryReadonlySqlConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("secondary_readonly_sql_connection_string"))
}

// SecondarySqlConnectionString returns a reference to field secondary_sql_connection_string of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) SecondarySqlConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("secondary_sql_connection_string"))
}

// Tags returns a reference to field tags of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("tags"))
}

// WriteEndpoints returns a reference to field write_endpoints of azurerm_cosmosdb_account.
func (ca cosmosdbAccountAttributes) WriteEndpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("write_endpoints"))
}

func (ca cosmosdbAccountAttributes) AnalyticalStorage() terra.ListValue[cosmosdbaccount.AnalyticalStorageAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.AnalyticalStorageAttributes](ca.ref.Append("analytical_storage"))
}

func (ca cosmosdbAccountAttributes) Backup() terra.ListValue[cosmosdbaccount.BackupAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.BackupAttributes](ca.ref.Append("backup"))
}

func (ca cosmosdbAccountAttributes) Capabilities() terra.SetValue[cosmosdbaccount.CapabilitiesAttributes] {
	return terra.ReferenceAsSet[cosmosdbaccount.CapabilitiesAttributes](ca.ref.Append("capabilities"))
}

func (ca cosmosdbAccountAttributes) Capacity() terra.ListValue[cosmosdbaccount.CapacityAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.CapacityAttributes](ca.ref.Append("capacity"))
}

func (ca cosmosdbAccountAttributes) ConsistencyPolicy() terra.ListValue[cosmosdbaccount.ConsistencyPolicyAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.ConsistencyPolicyAttributes](ca.ref.Append("consistency_policy"))
}

func (ca cosmosdbAccountAttributes) CorsRule() terra.ListValue[cosmosdbaccount.CorsRuleAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.CorsRuleAttributes](ca.ref.Append("cors_rule"))
}

func (ca cosmosdbAccountAttributes) GeoLocation() terra.SetValue[cosmosdbaccount.GeoLocationAttributes] {
	return terra.ReferenceAsSet[cosmosdbaccount.GeoLocationAttributes](ca.ref.Append("geo_location"))
}

func (ca cosmosdbAccountAttributes) Identity() terra.ListValue[cosmosdbaccount.IdentityAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.IdentityAttributes](ca.ref.Append("identity"))
}

func (ca cosmosdbAccountAttributes) Restore() terra.ListValue[cosmosdbaccount.RestoreAttributes] {
	return terra.ReferenceAsList[cosmosdbaccount.RestoreAttributes](ca.ref.Append("restore"))
}

func (ca cosmosdbAccountAttributes) Timeouts() cosmosdbaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbaccount.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

func (ca cosmosdbAccountAttributes) VirtualNetworkRule() terra.SetValue[cosmosdbaccount.VirtualNetworkRuleAttributes] {
	return terra.ReferenceAsSet[cosmosdbaccount.VirtualNetworkRuleAttributes](ca.ref.Append("virtual_network_rule"))
}

type cosmosdbAccountState struct {
	AccessKeyMetadataWritesEnabled       bool                                      `json:"access_key_metadata_writes_enabled"`
	AnalyticalStorageEnabled             bool                                      `json:"analytical_storage_enabled"`
	ConnectionStrings                    []string                                  `json:"connection_strings"`
	CreateMode                           string                                    `json:"create_mode"`
	DefaultIdentityType                  string                                    `json:"default_identity_type"`
	EnableAutomaticFailover              bool                                      `json:"enable_automatic_failover"`
	EnableFreeTier                       bool                                      `json:"enable_free_tier"`
	EnableMultipleWriteLocations         bool                                      `json:"enable_multiple_write_locations"`
	Endpoint                             string                                    `json:"endpoint"`
	Id                                   string                                    `json:"id"`
	IpRangeFilter                        string                                    `json:"ip_range_filter"`
	IsVirtualNetworkFilterEnabled        bool                                      `json:"is_virtual_network_filter_enabled"`
	KeyVaultKeyId                        string                                    `json:"key_vault_key_id"`
	Kind                                 string                                    `json:"kind"`
	LocalAuthenticationDisabled          bool                                      `json:"local_authentication_disabled"`
	Location                             string                                    `json:"location"`
	MongoServerVersion                   string                                    `json:"mongo_server_version"`
	Name                                 string                                    `json:"name"`
	NetworkAclBypassForAzureServices     bool                                      `json:"network_acl_bypass_for_azure_services"`
	NetworkAclBypassIds                  []string                                  `json:"network_acl_bypass_ids"`
	OfferType                            string                                    `json:"offer_type"`
	PrimaryKey                           string                                    `json:"primary_key"`
	PrimaryReadonlyKey                   string                                    `json:"primary_readonly_key"`
	PrimaryReadonlySqlConnectionString   string                                    `json:"primary_readonly_sql_connection_string"`
	PrimarySqlConnectionString           string                                    `json:"primary_sql_connection_string"`
	PublicNetworkAccessEnabled           bool                                      `json:"public_network_access_enabled"`
	ReadEndpoints                        []string                                  `json:"read_endpoints"`
	ResourceGroupName                    string                                    `json:"resource_group_name"`
	SecondaryKey                         string                                    `json:"secondary_key"`
	SecondaryReadonlyKey                 string                                    `json:"secondary_readonly_key"`
	SecondaryReadonlySqlConnectionString string                                    `json:"secondary_readonly_sql_connection_string"`
	SecondarySqlConnectionString         string                                    `json:"secondary_sql_connection_string"`
	Tags                                 map[string]string                         `json:"tags"`
	WriteEndpoints                       []string                                  `json:"write_endpoints"`
	AnalyticalStorage                    []cosmosdbaccount.AnalyticalStorageState  `json:"analytical_storage"`
	Backup                               []cosmosdbaccount.BackupState             `json:"backup"`
	Capabilities                         []cosmosdbaccount.CapabilitiesState       `json:"capabilities"`
	Capacity                             []cosmosdbaccount.CapacityState           `json:"capacity"`
	ConsistencyPolicy                    []cosmosdbaccount.ConsistencyPolicyState  `json:"consistency_policy"`
	CorsRule                             []cosmosdbaccount.CorsRuleState           `json:"cors_rule"`
	GeoLocation                          []cosmosdbaccount.GeoLocationState        `json:"geo_location"`
	Identity                             []cosmosdbaccount.IdentityState           `json:"identity"`
	Restore                              []cosmosdbaccount.RestoreState            `json:"restore"`
	Timeouts                             *cosmosdbaccount.TimeoutsState            `json:"timeouts"`
	VirtualNetworkRule                   []cosmosdbaccount.VirtualNetworkRuleState `json:"virtual_network_rule"`
}
