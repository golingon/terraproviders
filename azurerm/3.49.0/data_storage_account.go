// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastorageaccount "github.com/golingon/terraproviders/azurerm/3.49.0/datastorageaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataStorageAccount(name string, args DataStorageAccountArgs) *DataStorageAccount {
	return &DataStorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageAccount)(nil)

type DataStorageAccount struct {
	Name string
	Args DataStorageAccountArgs
}

func (sa *DataStorageAccount) DataSource() string {
	return "azurerm_storage_account"
}

func (sa *DataStorageAccount) LocalName() string {
	return sa.Name
}

func (sa *DataStorageAccount) Configuration() interface{} {
	return sa.Args
}

func (sa *DataStorageAccount) Attributes() dataStorageAccountAttributes {
	return dataStorageAccountAttributes{ref: terra.ReferenceDataResource(sa)}
}

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

func (sa dataStorageAccountAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("access_tier"))
}

func (sa dataStorageAccountAttributes) AccountKind() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("account_kind"))
}

func (sa dataStorageAccountAttributes) AccountReplicationType() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("account_replication_type"))
}

func (sa dataStorageAccountAttributes) AccountTier() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("account_tier"))
}

func (sa dataStorageAccountAttributes) AllowNestedItemsToBePublic() terra.BoolValue {
	return terra.ReferenceBool(sa.ref.Append("allow_nested_items_to_be_public"))
}

func (sa dataStorageAccountAttributes) EnableHttpsTrafficOnly() terra.BoolValue {
	return terra.ReferenceBool(sa.ref.Append("enable_https_traffic_only"))
}

func (sa dataStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("id"))
}

func (sa dataStorageAccountAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceBool(sa.ref.Append("infrastructure_encryption_enabled"))
}

func (sa dataStorageAccountAttributes) IsHnsEnabled() terra.BoolValue {
	return terra.ReferenceBool(sa.ref.Append("is_hns_enabled"))
}

func (sa dataStorageAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("location"))
}

func (sa dataStorageAccountAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("min_tls_version"))
}

func (sa dataStorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("name"))
}

func (sa dataStorageAccountAttributes) Nfsv3Enabled() terra.BoolValue {
	return terra.ReferenceBool(sa.ref.Append("nfsv3_enabled"))
}

func (sa dataStorageAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_access_key"))
}

func (sa dataStorageAccountAttributes) PrimaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_blob_connection_string"))
}

func (sa dataStorageAccountAttributes) PrimaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_blob_endpoint"))
}

func (sa dataStorageAccountAttributes) PrimaryBlobHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_blob_host"))
}

func (sa dataStorageAccountAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_connection_string"))
}

func (sa dataStorageAccountAttributes) PrimaryDfsEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_dfs_endpoint"))
}

func (sa dataStorageAccountAttributes) PrimaryDfsHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_dfs_host"))
}

func (sa dataStorageAccountAttributes) PrimaryFileEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_file_endpoint"))
}

func (sa dataStorageAccountAttributes) PrimaryFileHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_file_host"))
}

func (sa dataStorageAccountAttributes) PrimaryLocation() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_location"))
}

func (sa dataStorageAccountAttributes) PrimaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_queue_endpoint"))
}

func (sa dataStorageAccountAttributes) PrimaryQueueHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_queue_host"))
}

func (sa dataStorageAccountAttributes) PrimaryTableEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_table_endpoint"))
}

func (sa dataStorageAccountAttributes) PrimaryTableHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_table_host"))
}

func (sa dataStorageAccountAttributes) PrimaryWebEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_web_endpoint"))
}

func (sa dataStorageAccountAttributes) PrimaryWebHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("primary_web_host"))
}

func (sa dataStorageAccountAttributes) QueueEncryptionKeyType() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("queue_encryption_key_type"))
}

func (sa dataStorageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("resource_group_name"))
}

func (sa dataStorageAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_access_key"))
}

func (sa dataStorageAccountAttributes) SecondaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_blob_connection_string"))
}

func (sa dataStorageAccountAttributes) SecondaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_blob_endpoint"))
}

func (sa dataStorageAccountAttributes) SecondaryBlobHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_blob_host"))
}

func (sa dataStorageAccountAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_connection_string"))
}

func (sa dataStorageAccountAttributes) SecondaryDfsEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_dfs_endpoint"))
}

func (sa dataStorageAccountAttributes) SecondaryDfsHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_dfs_host"))
}

func (sa dataStorageAccountAttributes) SecondaryFileEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_file_endpoint"))
}

func (sa dataStorageAccountAttributes) SecondaryFileHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_file_host"))
}

func (sa dataStorageAccountAttributes) SecondaryLocation() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_location"))
}

func (sa dataStorageAccountAttributes) SecondaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_queue_endpoint"))
}

func (sa dataStorageAccountAttributes) SecondaryQueueHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_queue_host"))
}

func (sa dataStorageAccountAttributes) SecondaryTableEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_table_endpoint"))
}

func (sa dataStorageAccountAttributes) SecondaryTableHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_table_host"))
}

func (sa dataStorageAccountAttributes) SecondaryWebEndpoint() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_web_endpoint"))
}

func (sa dataStorageAccountAttributes) SecondaryWebHost() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("secondary_web_host"))
}

func (sa dataStorageAccountAttributes) TableEncryptionKeyType() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("table_encryption_key_type"))
}

func (sa dataStorageAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sa.ref.Append("tags"))
}

func (sa dataStorageAccountAttributes) AzureFilesAuthentication() terra.ListValue[datastorageaccount.AzureFilesAuthenticationAttributes] {
	return terra.ReferenceList[datastorageaccount.AzureFilesAuthenticationAttributes](sa.ref.Append("azure_files_authentication"))
}

func (sa dataStorageAccountAttributes) CustomDomain() terra.ListValue[datastorageaccount.CustomDomainAttributes] {
	return terra.ReferenceList[datastorageaccount.CustomDomainAttributes](sa.ref.Append("custom_domain"))
}

func (sa dataStorageAccountAttributes) Identity() terra.ListValue[datastorageaccount.IdentityAttributes] {
	return terra.ReferenceList[datastorageaccount.IdentityAttributes](sa.ref.Append("identity"))
}

func (sa dataStorageAccountAttributes) Timeouts() datastorageaccount.TimeoutsAttributes {
	return terra.ReferenceSingle[datastorageaccount.TimeoutsAttributes](sa.ref.Append("timeouts"))
}
