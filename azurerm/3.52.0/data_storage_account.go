// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastorageaccount "github.com/golingon/terraproviders/azurerm/3.52.0/datastorageaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageAccount creates a new instance of [DataStorageAccount].
func NewDataStorageAccount(name string, args DataStorageAccountArgs) *DataStorageAccount {
	return &DataStorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageAccount)(nil)

// DataStorageAccount represents the Terraform data resource azurerm_storage_account.
type DataStorageAccount struct {
	Name string
	Args DataStorageAccountArgs
}

// DataSource returns the Terraform object type for [DataStorageAccount].
func (sa *DataStorageAccount) DataSource() string {
	return "azurerm_storage_account"
}

// LocalName returns the local name for [DataStorageAccount].
func (sa *DataStorageAccount) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [DataStorageAccount].
func (sa *DataStorageAccount) Configuration() interface{} {
	return sa.Args
}

// Attributes returns the attributes for [DataStorageAccount].
func (sa *DataStorageAccount) Attributes() dataStorageAccountAttributes {
	return dataStorageAccountAttributes{ref: terra.ReferenceDataResource(sa)}
}

// DataStorageAccountArgs contains the configurations for azurerm_storage_account.
type DataStorageAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MinTlsVersion: string, optional
	MinTlsVersion terra.StringValue `hcl:"min_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AzureFilesAuthentication: min=0
	AzureFilesAuthentication []datastorageaccount.AzureFilesAuthentication `hcl:"azure_files_authentication,block" validate:"min=0"`
	// CustomDomain: min=0
	CustomDomain []datastorageaccount.CustomDomain `hcl:"custom_domain,block" validate:"min=0"`
	// Identity: min=0
	Identity []datastorageaccount.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datastorageaccount.Timeouts `hcl:"timeouts,block"`
}
type dataStorageAccountAttributes struct {
	ref terra.Reference
}

// AccessTier returns a reference to field access_tier of azurerm_storage_account.
func (sa dataStorageAccountAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("access_tier"))
}

// AccountKind returns a reference to field account_kind of azurerm_storage_account.
func (sa dataStorageAccountAttributes) AccountKind() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_kind"))
}

// AccountReplicationType returns a reference to field account_replication_type of azurerm_storage_account.
func (sa dataStorageAccountAttributes) AccountReplicationType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_replication_type"))
}

// AccountTier returns a reference to field account_tier of azurerm_storage_account.
func (sa dataStorageAccountAttributes) AccountTier() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_tier"))
}

// AllowNestedItemsToBePublic returns a reference to field allow_nested_items_to_be_public of azurerm_storage_account.
func (sa dataStorageAccountAttributes) AllowNestedItemsToBePublic() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("allow_nested_items_to_be_public"))
}

// EnableHttpsTrafficOnly returns a reference to field enable_https_traffic_only of azurerm_storage_account.
func (sa dataStorageAccountAttributes) EnableHttpsTrafficOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("enable_https_traffic_only"))
}

// Id returns a reference to field id of azurerm_storage_account.
func (sa dataStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_storage_account.
func (sa dataStorageAccountAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("infrastructure_encryption_enabled"))
}

// IsHnsEnabled returns a reference to field is_hns_enabled of azurerm_storage_account.
func (sa dataStorageAccountAttributes) IsHnsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("is_hns_enabled"))
}

// Location returns a reference to field location of azurerm_storage_account.
func (sa dataStorageAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("location"))
}

// MinTlsVersion returns a reference to field min_tls_version of azurerm_storage_account.
func (sa dataStorageAccountAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("min_tls_version"))
}

// Name returns a reference to field name of azurerm_storage_account.
func (sa dataStorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

// Nfsv3Enabled returns a reference to field nfsv3_enabled of azurerm_storage_account.
func (sa dataStorageAccountAttributes) Nfsv3Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("nfsv3_enabled"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_access_key"))
}

// PrimaryBlobConnectionString returns a reference to field primary_blob_connection_string of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_connection_string"))
}

// PrimaryBlobEndpoint returns a reference to field primary_blob_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_endpoint"))
}

// PrimaryBlobHost returns a reference to field primary_blob_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryBlobHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_host"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_connection_string"))
}

// PrimaryDfsEndpoint returns a reference to field primary_dfs_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryDfsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_dfs_endpoint"))
}

// PrimaryDfsHost returns a reference to field primary_dfs_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryDfsHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_dfs_host"))
}

// PrimaryFileEndpoint returns a reference to field primary_file_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryFileEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_file_endpoint"))
}

// PrimaryFileHost returns a reference to field primary_file_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryFileHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_file_host"))
}

// PrimaryLocation returns a reference to field primary_location of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_location"))
}

// PrimaryQueueEndpoint returns a reference to field primary_queue_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_queue_endpoint"))
}

// PrimaryQueueHost returns a reference to field primary_queue_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryQueueHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_queue_host"))
}

// PrimaryTableEndpoint returns a reference to field primary_table_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryTableEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_table_endpoint"))
}

// PrimaryTableHost returns a reference to field primary_table_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryTableHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_table_host"))
}

// PrimaryWebEndpoint returns a reference to field primary_web_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryWebEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_web_endpoint"))
}

// PrimaryWebHost returns a reference to field primary_web_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) PrimaryWebHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_web_host"))
}

// QueueEncryptionKeyType returns a reference to field queue_encryption_key_type of azurerm_storage_account.
func (sa dataStorageAccountAttributes) QueueEncryptionKeyType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("queue_encryption_key_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_storage_account.
func (sa dataStorageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_access_key"))
}

// SecondaryBlobConnectionString returns a reference to field secondary_blob_connection_string of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_connection_string"))
}

// SecondaryBlobEndpoint returns a reference to field secondary_blob_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_endpoint"))
}

// SecondaryBlobHost returns a reference to field secondary_blob_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryBlobHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_host"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_connection_string"))
}

// SecondaryDfsEndpoint returns a reference to field secondary_dfs_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryDfsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_dfs_endpoint"))
}

// SecondaryDfsHost returns a reference to field secondary_dfs_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryDfsHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_dfs_host"))
}

// SecondaryFileEndpoint returns a reference to field secondary_file_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryFileEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_file_endpoint"))
}

// SecondaryFileHost returns a reference to field secondary_file_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryFileHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_file_host"))
}

// SecondaryLocation returns a reference to field secondary_location of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_location"))
}

// SecondaryQueueEndpoint returns a reference to field secondary_queue_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_queue_endpoint"))
}

// SecondaryQueueHost returns a reference to field secondary_queue_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryQueueHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_queue_host"))
}

// SecondaryTableEndpoint returns a reference to field secondary_table_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryTableEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_table_endpoint"))
}

// SecondaryTableHost returns a reference to field secondary_table_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryTableHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_table_host"))
}

// SecondaryWebEndpoint returns a reference to field secondary_web_endpoint of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryWebEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_web_endpoint"))
}

// SecondaryWebHost returns a reference to field secondary_web_host of azurerm_storage_account.
func (sa dataStorageAccountAttributes) SecondaryWebHost() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_web_host"))
}

// TableEncryptionKeyType returns a reference to field table_encryption_key_type of azurerm_storage_account.
func (sa dataStorageAccountAttributes) TableEncryptionKeyType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("table_encryption_key_type"))
}

// Tags returns a reference to field tags of azurerm_storage_account.
func (sa dataStorageAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sa.ref.Append("tags"))
}

func (sa dataStorageAccountAttributes) AzureFilesAuthentication() terra.ListValue[datastorageaccount.AzureFilesAuthenticationAttributes] {
	return terra.ReferenceAsList[datastorageaccount.AzureFilesAuthenticationAttributes](sa.ref.Append("azure_files_authentication"))
}

func (sa dataStorageAccountAttributes) CustomDomain() terra.ListValue[datastorageaccount.CustomDomainAttributes] {
	return terra.ReferenceAsList[datastorageaccount.CustomDomainAttributes](sa.ref.Append("custom_domain"))
}

func (sa dataStorageAccountAttributes) Identity() terra.ListValue[datastorageaccount.IdentityAttributes] {
	return terra.ReferenceAsList[datastorageaccount.IdentityAttributes](sa.ref.Append("identity"))
}

func (sa dataStorageAccountAttributes) Timeouts() datastorageaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastorageaccount.TimeoutsAttributes](sa.ref.Append("timeouts"))
}
