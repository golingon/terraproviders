// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageaccount "github.com/golingon/terraproviders/azurerm/3.66.0/storageaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageAccount creates a new instance of [StorageAccount].
func NewStorageAccount(name string, args StorageAccountArgs) *StorageAccount {
	return &StorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageAccount)(nil)

// StorageAccount represents the Terraform resource azurerm_storage_account.
type StorageAccount struct {
	Name      string
	Args      StorageAccountArgs
	state     *storageAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageAccount].
func (sa *StorageAccount) Type() string {
	return "azurerm_storage_account"
}

// LocalName returns the local name for [StorageAccount].
func (sa *StorageAccount) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [StorageAccount].
func (sa *StorageAccount) Configuration() interface{} {
	return sa.Args
}

// DependOn is used for other resources to depend on [StorageAccount].
func (sa *StorageAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(sa)
}

// Dependencies returns the list of resources [StorageAccount] depends_on.
func (sa *StorageAccount) Dependencies() terra.Dependencies {
	return sa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageAccount].
func (sa *StorageAccount) LifecycleManagement() *terra.Lifecycle {
	return sa.Lifecycle
}

// Attributes returns the attributes for [StorageAccount].
func (sa *StorageAccount) Attributes() storageAccountAttributes {
	return storageAccountAttributes{ref: terra.ReferenceResource(sa)}
}

// ImportState imports the given attribute values into [StorageAccount]'s state.
func (sa *StorageAccount) ImportState(av io.Reader) error {
	sa.state = &storageAccountState{}
	if err := json.NewDecoder(av).Decode(sa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sa.Type(), sa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageAccount] has state.
func (sa *StorageAccount) State() (*storageAccountState, bool) {
	return sa.state, sa.state != nil
}

// StateMust returns the state for [StorageAccount]. Panics if the state is nil.
func (sa *StorageAccount) StateMust() *storageAccountState {
	if sa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sa.Type(), sa.LocalName()))
	}
	return sa.state
}

// StorageAccountArgs contains the configurations for azurerm_storage_account.
type StorageAccountArgs struct {
	// AccessTier: string, optional
	AccessTier terra.StringValue `hcl:"access_tier,attr"`
	// AccountKind: string, optional
	AccountKind terra.StringValue `hcl:"account_kind,attr"`
	// AccountReplicationType: string, required
	AccountReplicationType terra.StringValue `hcl:"account_replication_type,attr" validate:"required"`
	// AccountTier: string, required
	AccountTier terra.StringValue `hcl:"account_tier,attr" validate:"required"`
	// AllowNestedItemsToBePublic: bool, optional
	AllowNestedItemsToBePublic terra.BoolValue `hcl:"allow_nested_items_to_be_public,attr"`
	// AllowedCopyScope: string, optional
	AllowedCopyScope terra.StringValue `hcl:"allowed_copy_scope,attr"`
	// CrossTenantReplicationEnabled: bool, optional
	CrossTenantReplicationEnabled terra.BoolValue `hcl:"cross_tenant_replication_enabled,attr"`
	// DefaultToOauthAuthentication: bool, optional
	DefaultToOauthAuthentication terra.BoolValue `hcl:"default_to_oauth_authentication,attr"`
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
	// EnableHttpsTrafficOnly: bool, optional
	EnableHttpsTrafficOnly terra.BoolValue `hcl:"enable_https_traffic_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InfrastructureEncryptionEnabled: bool, optional
	InfrastructureEncryptionEnabled terra.BoolValue `hcl:"infrastructure_encryption_enabled,attr"`
	// IsHnsEnabled: bool, optional
	IsHnsEnabled terra.BoolValue `hcl:"is_hns_enabled,attr"`
	// LargeFileShareEnabled: bool, optional
	LargeFileShareEnabled terra.BoolValue `hcl:"large_file_share_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MinTlsVersion: string, optional
	MinTlsVersion terra.StringValue `hcl:"min_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Nfsv3Enabled: bool, optional
	Nfsv3Enabled terra.BoolValue `hcl:"nfsv3_enabled,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// QueueEncryptionKeyType: string, optional
	QueueEncryptionKeyType terra.StringValue `hcl:"queue_encryption_key_type,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SftpEnabled: bool, optional
	SftpEnabled terra.BoolValue `hcl:"sftp_enabled,attr"`
	// SharedAccessKeyEnabled: bool, optional
	SharedAccessKeyEnabled terra.BoolValue `hcl:"shared_access_key_enabled,attr"`
	// TableEncryptionKeyType: string, optional
	TableEncryptionKeyType terra.StringValue `hcl:"table_encryption_key_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AzureFilesAuthentication: optional
	AzureFilesAuthentication *storageaccount.AzureFilesAuthentication `hcl:"azure_files_authentication,block"`
	// BlobProperties: optional
	BlobProperties *storageaccount.BlobProperties `hcl:"blob_properties,block"`
	// CustomDomain: optional
	CustomDomain *storageaccount.CustomDomain `hcl:"custom_domain,block"`
	// CustomerManagedKey: optional
	CustomerManagedKey *storageaccount.CustomerManagedKey `hcl:"customer_managed_key,block"`
	// Identity: optional
	Identity *storageaccount.Identity `hcl:"identity,block"`
	// ImmutabilityPolicy: optional
	ImmutabilityPolicy *storageaccount.ImmutabilityPolicy `hcl:"immutability_policy,block"`
	// NetworkRules: optional
	NetworkRules *storageaccount.NetworkRules `hcl:"network_rules,block"`
	// QueueProperties: optional
	QueueProperties *storageaccount.QueueProperties `hcl:"queue_properties,block"`
	// Routing: optional
	Routing *storageaccount.Routing `hcl:"routing,block"`
	// SasPolicy: optional
	SasPolicy *storageaccount.SasPolicy `hcl:"sas_policy,block"`
	// ShareProperties: optional
	ShareProperties *storageaccount.ShareProperties `hcl:"share_properties,block"`
	// StaticWebsite: optional
	StaticWebsite *storageaccount.StaticWebsite `hcl:"static_website,block"`
	// Timeouts: optional
	Timeouts *storageaccount.Timeouts `hcl:"timeouts,block"`
}
type storageAccountAttributes struct {
	ref terra.Reference
}

// AccessTier returns a reference to field access_tier of azurerm_storage_account.
func (sa storageAccountAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("access_tier"))
}

// AccountKind returns a reference to field account_kind of azurerm_storage_account.
func (sa storageAccountAttributes) AccountKind() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_kind"))
}

// AccountReplicationType returns a reference to field account_replication_type of azurerm_storage_account.
func (sa storageAccountAttributes) AccountReplicationType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_replication_type"))
}

// AccountTier returns a reference to field account_tier of azurerm_storage_account.
func (sa storageAccountAttributes) AccountTier() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_tier"))
}

// AllowNestedItemsToBePublic returns a reference to field allow_nested_items_to_be_public of azurerm_storage_account.
func (sa storageAccountAttributes) AllowNestedItemsToBePublic() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("allow_nested_items_to_be_public"))
}

// AllowedCopyScope returns a reference to field allowed_copy_scope of azurerm_storage_account.
func (sa storageAccountAttributes) AllowedCopyScope() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("allowed_copy_scope"))
}

// CrossTenantReplicationEnabled returns a reference to field cross_tenant_replication_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) CrossTenantReplicationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("cross_tenant_replication_enabled"))
}

// DefaultToOauthAuthentication returns a reference to field default_to_oauth_authentication of azurerm_storage_account.
func (sa storageAccountAttributes) DefaultToOauthAuthentication() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("default_to_oauth_authentication"))
}

// EdgeZone returns a reference to field edge_zone of azurerm_storage_account.
func (sa storageAccountAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("edge_zone"))
}

// EnableHttpsTrafficOnly returns a reference to field enable_https_traffic_only of azurerm_storage_account.
func (sa storageAccountAttributes) EnableHttpsTrafficOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("enable_https_traffic_only"))
}

// Id returns a reference to field id of azurerm_storage_account.
func (sa storageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("infrastructure_encryption_enabled"))
}

// IsHnsEnabled returns a reference to field is_hns_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) IsHnsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("is_hns_enabled"))
}

// LargeFileShareEnabled returns a reference to field large_file_share_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) LargeFileShareEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("large_file_share_enabled"))
}

// Location returns a reference to field location of azurerm_storage_account.
func (sa storageAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("location"))
}

// MinTlsVersion returns a reference to field min_tls_version of azurerm_storage_account.
func (sa storageAccountAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("min_tls_version"))
}

// Name returns a reference to field name of azurerm_storage_account.
func (sa storageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

// Nfsv3Enabled returns a reference to field nfsv3_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) Nfsv3Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("nfsv3_enabled"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_access_key"))
}

// PrimaryBlobConnectionString returns a reference to field primary_blob_connection_string of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_connection_string"))
}

// PrimaryBlobEndpoint returns a reference to field primary_blob_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_endpoint"))
}

// PrimaryBlobHost returns a reference to field primary_blob_host of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryBlobHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_host"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_connection_string"))
}

// PrimaryDfsEndpoint returns a reference to field primary_dfs_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryDfsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_dfs_endpoint"))
}

// PrimaryDfsHost returns a reference to field primary_dfs_host of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryDfsHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_dfs_host"))
}

// PrimaryFileEndpoint returns a reference to field primary_file_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryFileEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_file_endpoint"))
}

// PrimaryFileHost returns a reference to field primary_file_host of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryFileHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_file_host"))
}

// PrimaryLocation returns a reference to field primary_location of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_location"))
}

// PrimaryQueueEndpoint returns a reference to field primary_queue_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_queue_endpoint"))
}

// PrimaryQueueHost returns a reference to field primary_queue_host of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryQueueHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_queue_host"))
}

// PrimaryTableEndpoint returns a reference to field primary_table_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryTableEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_table_endpoint"))
}

// PrimaryTableHost returns a reference to field primary_table_host of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryTableHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_table_host"))
}

// PrimaryWebEndpoint returns a reference to field primary_web_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryWebEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_web_endpoint"))
}

// PrimaryWebHost returns a reference to field primary_web_host of azurerm_storage_account.
func (sa storageAccountAttributes) PrimaryWebHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_web_host"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("public_network_access_enabled"))
}

// QueueEncryptionKeyType returns a reference to field queue_encryption_key_type of azurerm_storage_account.
func (sa storageAccountAttributes) QueueEncryptionKeyType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("queue_encryption_key_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_storage_account.
func (sa storageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_access_key"))
}

// SecondaryBlobConnectionString returns a reference to field secondary_blob_connection_string of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_connection_string"))
}

// SecondaryBlobEndpoint returns a reference to field secondary_blob_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_endpoint"))
}

// SecondaryBlobHost returns a reference to field secondary_blob_host of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryBlobHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_host"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_connection_string"))
}

// SecondaryDfsEndpoint returns a reference to field secondary_dfs_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryDfsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_dfs_endpoint"))
}

// SecondaryDfsHost returns a reference to field secondary_dfs_host of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryDfsHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_dfs_host"))
}

// SecondaryFileEndpoint returns a reference to field secondary_file_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryFileEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_file_endpoint"))
}

// SecondaryFileHost returns a reference to field secondary_file_host of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryFileHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_file_host"))
}

// SecondaryLocation returns a reference to field secondary_location of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_location"))
}

// SecondaryQueueEndpoint returns a reference to field secondary_queue_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_queue_endpoint"))
}

// SecondaryQueueHost returns a reference to field secondary_queue_host of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryQueueHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_queue_host"))
}

// SecondaryTableEndpoint returns a reference to field secondary_table_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryTableEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_table_endpoint"))
}

// SecondaryTableHost returns a reference to field secondary_table_host of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryTableHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_table_host"))
}

// SecondaryWebEndpoint returns a reference to field secondary_web_endpoint of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryWebEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_web_endpoint"))
}

// SecondaryWebHost returns a reference to field secondary_web_host of azurerm_storage_account.
func (sa storageAccountAttributes) SecondaryWebHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_web_host"))
}

// SftpEnabled returns a reference to field sftp_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) SftpEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("sftp_enabled"))
}

// SharedAccessKeyEnabled returns a reference to field shared_access_key_enabled of azurerm_storage_account.
func (sa storageAccountAttributes) SharedAccessKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("shared_access_key_enabled"))
}

// TableEncryptionKeyType returns a reference to field table_encryption_key_type of azurerm_storage_account.
func (sa storageAccountAttributes) TableEncryptionKeyType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("table_encryption_key_type"))
}

// Tags returns a reference to field tags of azurerm_storage_account.
func (sa storageAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sa.ref.Append("tags"))
}

func (sa storageAccountAttributes) AzureFilesAuthentication() terra.ListValue[storageaccount.AzureFilesAuthenticationAttributes] {
	return terra.ReferenceAsList[storageaccount.AzureFilesAuthenticationAttributes](sa.ref.Append("azure_files_authentication"))
}

func (sa storageAccountAttributes) BlobProperties() terra.ListValue[storageaccount.BlobPropertiesAttributes] {
	return terra.ReferenceAsList[storageaccount.BlobPropertiesAttributes](sa.ref.Append("blob_properties"))
}

func (sa storageAccountAttributes) CustomDomain() terra.ListValue[storageaccount.CustomDomainAttributes] {
	return terra.ReferenceAsList[storageaccount.CustomDomainAttributes](sa.ref.Append("custom_domain"))
}

func (sa storageAccountAttributes) CustomerManagedKey() terra.ListValue[storageaccount.CustomerManagedKeyAttributes] {
	return terra.ReferenceAsList[storageaccount.CustomerManagedKeyAttributes](sa.ref.Append("customer_managed_key"))
}

func (sa storageAccountAttributes) Identity() terra.ListValue[storageaccount.IdentityAttributes] {
	return terra.ReferenceAsList[storageaccount.IdentityAttributes](sa.ref.Append("identity"))
}

func (sa storageAccountAttributes) ImmutabilityPolicy() terra.ListValue[storageaccount.ImmutabilityPolicyAttributes] {
	return terra.ReferenceAsList[storageaccount.ImmutabilityPolicyAttributes](sa.ref.Append("immutability_policy"))
}

func (sa storageAccountAttributes) NetworkRules() terra.ListValue[storageaccount.NetworkRulesAttributes] {
	return terra.ReferenceAsList[storageaccount.NetworkRulesAttributes](sa.ref.Append("network_rules"))
}

func (sa storageAccountAttributes) QueueProperties() terra.ListValue[storageaccount.QueuePropertiesAttributes] {
	return terra.ReferenceAsList[storageaccount.QueuePropertiesAttributes](sa.ref.Append("queue_properties"))
}

func (sa storageAccountAttributes) Routing() terra.ListValue[storageaccount.RoutingAttributes] {
	return terra.ReferenceAsList[storageaccount.RoutingAttributes](sa.ref.Append("routing"))
}

func (sa storageAccountAttributes) SasPolicy() terra.ListValue[storageaccount.SasPolicyAttributes] {
	return terra.ReferenceAsList[storageaccount.SasPolicyAttributes](sa.ref.Append("sas_policy"))
}

func (sa storageAccountAttributes) ShareProperties() terra.ListValue[storageaccount.SharePropertiesAttributes] {
	return terra.ReferenceAsList[storageaccount.SharePropertiesAttributes](sa.ref.Append("share_properties"))
}

func (sa storageAccountAttributes) StaticWebsite() terra.ListValue[storageaccount.StaticWebsiteAttributes] {
	return terra.ReferenceAsList[storageaccount.StaticWebsiteAttributes](sa.ref.Append("static_website"))
}

func (sa storageAccountAttributes) Timeouts() storageaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageaccount.TimeoutsAttributes](sa.ref.Append("timeouts"))
}

type storageAccountState struct {
	AccessTier                      string                                         `json:"access_tier"`
	AccountKind                     string                                         `json:"account_kind"`
	AccountReplicationType          string                                         `json:"account_replication_type"`
	AccountTier                     string                                         `json:"account_tier"`
	AllowNestedItemsToBePublic      bool                                           `json:"allow_nested_items_to_be_public"`
	AllowedCopyScope                string                                         `json:"allowed_copy_scope"`
	CrossTenantReplicationEnabled   bool                                           `json:"cross_tenant_replication_enabled"`
	DefaultToOauthAuthentication    bool                                           `json:"default_to_oauth_authentication"`
	EdgeZone                        string                                         `json:"edge_zone"`
	EnableHttpsTrafficOnly          bool                                           `json:"enable_https_traffic_only"`
	Id                              string                                         `json:"id"`
	InfrastructureEncryptionEnabled bool                                           `json:"infrastructure_encryption_enabled"`
	IsHnsEnabled                    bool                                           `json:"is_hns_enabled"`
	LargeFileShareEnabled           bool                                           `json:"large_file_share_enabled"`
	Location                        string                                         `json:"location"`
	MinTlsVersion                   string                                         `json:"min_tls_version"`
	Name                            string                                         `json:"name"`
	Nfsv3Enabled                    bool                                           `json:"nfsv3_enabled"`
	PrimaryAccessKey                string                                         `json:"primary_access_key"`
	PrimaryBlobConnectionString     string                                         `json:"primary_blob_connection_string"`
	PrimaryBlobEndpoint             string                                         `json:"primary_blob_endpoint"`
	PrimaryBlobHost                 string                                         `json:"primary_blob_host"`
	PrimaryConnectionString         string                                         `json:"primary_connection_string"`
	PrimaryDfsEndpoint              string                                         `json:"primary_dfs_endpoint"`
	PrimaryDfsHost                  string                                         `json:"primary_dfs_host"`
	PrimaryFileEndpoint             string                                         `json:"primary_file_endpoint"`
	PrimaryFileHost                 string                                         `json:"primary_file_host"`
	PrimaryLocation                 string                                         `json:"primary_location"`
	PrimaryQueueEndpoint            string                                         `json:"primary_queue_endpoint"`
	PrimaryQueueHost                string                                         `json:"primary_queue_host"`
	PrimaryTableEndpoint            string                                         `json:"primary_table_endpoint"`
	PrimaryTableHost                string                                         `json:"primary_table_host"`
	PrimaryWebEndpoint              string                                         `json:"primary_web_endpoint"`
	PrimaryWebHost                  string                                         `json:"primary_web_host"`
	PublicNetworkAccessEnabled      bool                                           `json:"public_network_access_enabled"`
	QueueEncryptionKeyType          string                                         `json:"queue_encryption_key_type"`
	ResourceGroupName               string                                         `json:"resource_group_name"`
	SecondaryAccessKey              string                                         `json:"secondary_access_key"`
	SecondaryBlobConnectionString   string                                         `json:"secondary_blob_connection_string"`
	SecondaryBlobEndpoint           string                                         `json:"secondary_blob_endpoint"`
	SecondaryBlobHost               string                                         `json:"secondary_blob_host"`
	SecondaryConnectionString       string                                         `json:"secondary_connection_string"`
	SecondaryDfsEndpoint            string                                         `json:"secondary_dfs_endpoint"`
	SecondaryDfsHost                string                                         `json:"secondary_dfs_host"`
	SecondaryFileEndpoint           string                                         `json:"secondary_file_endpoint"`
	SecondaryFileHost               string                                         `json:"secondary_file_host"`
	SecondaryLocation               string                                         `json:"secondary_location"`
	SecondaryQueueEndpoint          string                                         `json:"secondary_queue_endpoint"`
	SecondaryQueueHost              string                                         `json:"secondary_queue_host"`
	SecondaryTableEndpoint          string                                         `json:"secondary_table_endpoint"`
	SecondaryTableHost              string                                         `json:"secondary_table_host"`
	SecondaryWebEndpoint            string                                         `json:"secondary_web_endpoint"`
	SecondaryWebHost                string                                         `json:"secondary_web_host"`
	SftpEnabled                     bool                                           `json:"sftp_enabled"`
	SharedAccessKeyEnabled          bool                                           `json:"shared_access_key_enabled"`
	TableEncryptionKeyType          string                                         `json:"table_encryption_key_type"`
	Tags                            map[string]string                              `json:"tags"`
	AzureFilesAuthentication        []storageaccount.AzureFilesAuthenticationState `json:"azure_files_authentication"`
	BlobProperties                  []storageaccount.BlobPropertiesState           `json:"blob_properties"`
	CustomDomain                    []storageaccount.CustomDomainState             `json:"custom_domain"`
	CustomerManagedKey              []storageaccount.CustomerManagedKeyState       `json:"customer_managed_key"`
	Identity                        []storageaccount.IdentityState                 `json:"identity"`
	ImmutabilityPolicy              []storageaccount.ImmutabilityPolicyState       `json:"immutability_policy"`
	NetworkRules                    []storageaccount.NetworkRulesState             `json:"network_rules"`
	QueueProperties                 []storageaccount.QueuePropertiesState          `json:"queue_properties"`
	Routing                         []storageaccount.RoutingState                  `json:"routing"`
	SasPolicy                       []storageaccount.SasPolicyState                `json:"sas_policy"`
	ShareProperties                 []storageaccount.SharePropertiesState          `json:"share_properties"`
	StaticWebsite                   []storageaccount.StaticWebsiteState            `json:"static_website"`
	Timeouts                        *storageaccount.TimeoutsState                  `json:"timeouts"`
}
