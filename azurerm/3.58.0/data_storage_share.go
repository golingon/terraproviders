// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastorageshare "github.com/golingon/terraproviders/azurerm/3.58.0/datastorageshare"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageShare creates a new instance of [DataStorageShare].
func NewDataStorageShare(name string, args DataStorageShareArgs) *DataStorageShare {
	return &DataStorageShare{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageShare)(nil)

// DataStorageShare represents the Terraform data resource azurerm_storage_share.
type DataStorageShare struct {
	Name string
	Args DataStorageShareArgs
}

// DataSource returns the Terraform object type for [DataStorageShare].
func (ss *DataStorageShare) DataSource() string {
	return "azurerm_storage_share"
}

// LocalName returns the local name for [DataStorageShare].
func (ss *DataStorageShare) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [DataStorageShare].
func (ss *DataStorageShare) Configuration() interface{} {
	return ss.Args
}

// Attributes returns the attributes for [DataStorageShare].
func (ss *DataStorageShare) Attributes() dataStorageShareAttributes {
	return dataStorageShareAttributes{ref: terra.ReferenceDataResource(ss)}
}

// DataStorageShareArgs contains the configurations for azurerm_storage_share.
type DataStorageShareArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Acl: min=0
	Acl []datastorageshare.Acl `hcl:"acl,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datastorageshare.Timeouts `hcl:"timeouts,block"`
}
type dataStorageShareAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_share.
func (ss dataStorageShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_share.
func (ss dataStorageShareAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_share.
func (ss dataStorageShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// Quota returns a reference to field quota of azurerm_storage_share.
func (ss dataStorageShareAttributes) Quota() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("quota"))
}

// ResourceManagerId returns a reference to field resource_manager_id of azurerm_storage_share.
func (ss dataStorageShareAttributes) ResourceManagerId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_manager_id"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_share.
func (ss dataStorageShareAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("storage_account_name"))
}

func (ss dataStorageShareAttributes) Acl() terra.ListValue[datastorageshare.AclAttributes] {
	return terra.ReferenceAsList[datastorageshare.AclAttributes](ss.ref.Append("acl"))
}

func (ss dataStorageShareAttributes) Timeouts() datastorageshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastorageshare.TimeoutsAttributes](ss.ref.Append("timeouts"))
}
