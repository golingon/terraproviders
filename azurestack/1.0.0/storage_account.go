// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	storageaccount "github.com/golingon/terraproviders/azurestack/1.0.0/storageaccount"
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

// StorageAccount represents the Terraform resource azurestack_storage_account.
type StorageAccount struct {
	Name      string
	Args      StorageAccountArgs
	state     *storageAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageAccount].
func (sa *StorageAccount) Type() string {
	return "azurestack_storage_account"
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

// StorageAccountArgs contains the configurations for azurestack_storage_account.
type StorageAccountArgs struct {
	// AccountEncryptionSource: string, optional
	AccountEncryptionSource terra.StringValue `hcl:"account_encryption_source,attr"`
	// AccountKind: string, optional
	AccountKind terra.StringValue `hcl:"account_kind,attr"`
	// AccountReplicationType: string, required
	AccountReplicationType terra.StringValue `hcl:"account_replication_type,attr" validate:"required"`
	// AccountTier: string, required
	AccountTier terra.StringValue `hcl:"account_tier,attr" validate:"required"`
	// EnableBlobEncryption: bool, optional
	EnableBlobEncryption terra.BoolValue `hcl:"enable_blob_encryption,attr"`
	// EnableHttpsTrafficOnly: bool, optional
	EnableHttpsTrafficOnly terra.BoolValue `hcl:"enable_https_traffic_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CustomDomain: optional
	CustomDomain *storageaccount.CustomDomain `hcl:"custom_domain,block"`
	// Timeouts: optional
	Timeouts *storageaccount.Timeouts `hcl:"timeouts,block"`
}
type storageAccountAttributes struct {
	ref terra.Reference
}

// AccountEncryptionSource returns a reference to field account_encryption_source of azurestack_storage_account.
func (sa storageAccountAttributes) AccountEncryptionSource() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_encryption_source"))
}

// AccountKind returns a reference to field account_kind of azurestack_storage_account.
func (sa storageAccountAttributes) AccountKind() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_kind"))
}

// AccountReplicationType returns a reference to field account_replication_type of azurestack_storage_account.
func (sa storageAccountAttributes) AccountReplicationType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_replication_type"))
}

// AccountTier returns a reference to field account_tier of azurestack_storage_account.
func (sa storageAccountAttributes) AccountTier() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_tier"))
}

// EnableBlobEncryption returns a reference to field enable_blob_encryption of azurestack_storage_account.
func (sa storageAccountAttributes) EnableBlobEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("enable_blob_encryption"))
}

// EnableHttpsTrafficOnly returns a reference to field enable_https_traffic_only of azurestack_storage_account.
func (sa storageAccountAttributes) EnableHttpsTrafficOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("enable_https_traffic_only"))
}

// Id returns a reference to field id of azurestack_storage_account.
func (sa storageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// Location returns a reference to field location of azurestack_storage_account.
func (sa storageAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_storage_account.
func (sa storageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_access_key"))
}

// PrimaryBlobConnectionString returns a reference to field primary_blob_connection_string of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_connection_string"))
}

// PrimaryBlobEndpoint returns a reference to field primary_blob_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_blob_endpoint"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_connection_string"))
}

// PrimaryFileEndpoint returns a reference to field primary_file_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryFileEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_file_endpoint"))
}

// PrimaryLocation returns a reference to field primary_location of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_location"))
}

// PrimaryQueueEndpoint returns a reference to field primary_queue_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_queue_endpoint"))
}

// PrimaryTableEndpoint returns a reference to field primary_table_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) PrimaryTableEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("primary_table_endpoint"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_storage_account.
func (sa storageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_access_key"))
}

// SecondaryBlobConnectionString returns a reference to field secondary_blob_connection_string of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryBlobConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_connection_string"))
}

// SecondaryBlobEndpoint returns a reference to field secondary_blob_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryBlobEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_blob_endpoint"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_connection_string"))
}

// SecondaryLocation returns a reference to field secondary_location of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_location"))
}

// SecondaryQueueEndpoint returns a reference to field secondary_queue_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryQueueEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_queue_endpoint"))
}

// SecondaryTableEndpoint returns a reference to field secondary_table_endpoint of azurestack_storage_account.
func (sa storageAccountAttributes) SecondaryTableEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("secondary_table_endpoint"))
}

// Tags returns a reference to field tags of azurestack_storage_account.
func (sa storageAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sa.ref.Append("tags"))
}

func (sa storageAccountAttributes) CustomDomain() terra.ListValue[storageaccount.CustomDomainAttributes] {
	return terra.ReferenceAsList[storageaccount.CustomDomainAttributes](sa.ref.Append("custom_domain"))
}

func (sa storageAccountAttributes) Timeouts() storageaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageaccount.TimeoutsAttributes](sa.ref.Append("timeouts"))
}

type storageAccountState struct {
	AccountEncryptionSource       string                             `json:"account_encryption_source"`
	AccountKind                   string                             `json:"account_kind"`
	AccountReplicationType        string                             `json:"account_replication_type"`
	AccountTier                   string                             `json:"account_tier"`
	EnableBlobEncryption          bool                               `json:"enable_blob_encryption"`
	EnableHttpsTrafficOnly        bool                               `json:"enable_https_traffic_only"`
	Id                            string                             `json:"id"`
	Location                      string                             `json:"location"`
	Name                          string                             `json:"name"`
	PrimaryAccessKey              string                             `json:"primary_access_key"`
	PrimaryBlobConnectionString   string                             `json:"primary_blob_connection_string"`
	PrimaryBlobEndpoint           string                             `json:"primary_blob_endpoint"`
	PrimaryConnectionString       string                             `json:"primary_connection_string"`
	PrimaryFileEndpoint           string                             `json:"primary_file_endpoint"`
	PrimaryLocation               string                             `json:"primary_location"`
	PrimaryQueueEndpoint          string                             `json:"primary_queue_endpoint"`
	PrimaryTableEndpoint          string                             `json:"primary_table_endpoint"`
	ResourceGroupName             string                             `json:"resource_group_name"`
	SecondaryAccessKey            string                             `json:"secondary_access_key"`
	SecondaryBlobConnectionString string                             `json:"secondary_blob_connection_string"`
	SecondaryBlobEndpoint         string                             `json:"secondary_blob_endpoint"`
	SecondaryConnectionString     string                             `json:"secondary_connection_string"`
	SecondaryLocation             string                             `json:"secondary_location"`
	SecondaryQueueEndpoint        string                             `json:"secondary_queue_endpoint"`
	SecondaryTableEndpoint        string                             `json:"secondary_table_endpoint"`
	Tags                          map[string]string                  `json:"tags"`
	CustomDomain                  []storageaccount.CustomDomainState `json:"custom_domain"`
	Timeouts                      *storageaccount.TimeoutsState      `json:"timeouts"`
}