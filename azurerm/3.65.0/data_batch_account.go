// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databatchaccount "github.com/golingon/terraproviders/azurerm/3.65.0/databatchaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBatchAccount creates a new instance of [DataBatchAccount].
func NewDataBatchAccount(name string, args DataBatchAccountArgs) *DataBatchAccount {
	return &DataBatchAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBatchAccount)(nil)

// DataBatchAccount represents the Terraform data resource azurerm_batch_account.
type DataBatchAccount struct {
	Name string
	Args DataBatchAccountArgs
}

// DataSource returns the Terraform object type for [DataBatchAccount].
func (ba *DataBatchAccount) DataSource() string {
	return "azurerm_batch_account"
}

// LocalName returns the local name for [DataBatchAccount].
func (ba *DataBatchAccount) LocalName() string {
	return ba.Name
}

// Configuration returns the configuration (args) for [DataBatchAccount].
func (ba *DataBatchAccount) Configuration() interface{} {
	return ba.Args
}

// Attributes returns the attributes for [DataBatchAccount].
func (ba *DataBatchAccount) Attributes() dataBatchAccountAttributes {
	return dataBatchAccountAttributes{ref: terra.ReferenceDataResource(ba)}
}

// DataBatchAccountArgs contains the configurations for azurerm_batch_account.
type DataBatchAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Encryption: min=0
	Encryption []databatchaccount.Encryption `hcl:"encryption,block" validate:"min=0"`
	// KeyVaultReference: min=0
	KeyVaultReference []databatchaccount.KeyVaultReference `hcl:"key_vault_reference,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *databatchaccount.Timeouts `hcl:"timeouts,block"`
}
type dataBatchAccountAttributes struct {
	ref terra.Reference
}

// AccountEndpoint returns a reference to field account_endpoint of azurerm_batch_account.
func (ba dataBatchAccountAttributes) AccountEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("account_endpoint"))
}

// Id returns a reference to field id of azurerm_batch_account.
func (ba dataBatchAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_batch_account.
func (ba dataBatchAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_batch_account.
func (ba dataBatchAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("name"))
}

// PoolAllocationMode returns a reference to field pool_allocation_mode of azurerm_batch_account.
func (ba dataBatchAccountAttributes) PoolAllocationMode() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("pool_allocation_mode"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_batch_account.
func (ba dataBatchAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("primary_access_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_account.
func (ba dataBatchAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_batch_account.
func (ba dataBatchAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("secondary_access_key"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_batch_account.
func (ba dataBatchAccountAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("storage_account_id"))
}

// Tags returns a reference to field tags of azurerm_batch_account.
func (ba dataBatchAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ba.ref.Append("tags"))
}

func (ba dataBatchAccountAttributes) Encryption() terra.ListValue[databatchaccount.EncryptionAttributes] {
	return terra.ReferenceAsList[databatchaccount.EncryptionAttributes](ba.ref.Append("encryption"))
}

func (ba dataBatchAccountAttributes) KeyVaultReference() terra.ListValue[databatchaccount.KeyVaultReferenceAttributes] {
	return terra.ReferenceAsList[databatchaccount.KeyVaultReferenceAttributes](ba.ref.Append("key_vault_reference"))
}

func (ba dataBatchAccountAttributes) Timeouts() databatchaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databatchaccount.TimeoutsAttributes](ba.ref.Append("timeouts"))
}
