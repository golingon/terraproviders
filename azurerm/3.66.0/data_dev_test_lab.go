// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadevtestlab "github.com/golingon/terraproviders/azurerm/3.66.0/datadevtestlab"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDevTestLab creates a new instance of [DataDevTestLab].
func NewDataDevTestLab(name string, args DataDevTestLabArgs) *DataDevTestLab {
	return &DataDevTestLab{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDevTestLab)(nil)

// DataDevTestLab represents the Terraform data resource azurerm_dev_test_lab.
type DataDevTestLab struct {
	Name string
	Args DataDevTestLabArgs
}

// DataSource returns the Terraform object type for [DataDevTestLab].
func (dtl *DataDevTestLab) DataSource() string {
	return "azurerm_dev_test_lab"
}

// LocalName returns the local name for [DataDevTestLab].
func (dtl *DataDevTestLab) LocalName() string {
	return dtl.Name
}

// Configuration returns the configuration (args) for [DataDevTestLab].
func (dtl *DataDevTestLab) Configuration() interface{} {
	return dtl.Args
}

// Attributes returns the attributes for [DataDevTestLab].
func (dtl *DataDevTestLab) Attributes() dataDevTestLabAttributes {
	return dataDevTestLabAttributes{ref: terra.ReferenceDataResource(dtl)}
}

// DataDevTestLabArgs contains the configurations for azurerm_dev_test_lab.
type DataDevTestLabArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadevtestlab.Timeouts `hcl:"timeouts,block"`
}
type dataDevTestLabAttributes struct {
	ref terra.Reference
}

// ArtifactsStorageAccountId returns a reference to field artifacts_storage_account_id of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) ArtifactsStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("artifacts_storage_account_id"))
}

// DefaultPremiumStorageAccountId returns a reference to field default_premium_storage_account_id of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) DefaultPremiumStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("default_premium_storage_account_id"))
}

// DefaultStorageAccountId returns a reference to field default_storage_account_id of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) DefaultStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("default_storage_account_id"))
}

// Id returns a reference to field id of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("key_vault_id"))
}

// Location returns a reference to field location of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("name"))
}

// PremiumDataDiskStorageAccountId returns a reference to field premium_data_disk_storage_account_id of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) PremiumDataDiskStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("premium_data_disk_storage_account_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("resource_group_name"))
}

// StorageType returns a reference to field storage_type of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtl.ref.Append("tags"))
}

// UniqueIdentifier returns a reference to field unique_identifier of azurerm_dev_test_lab.
func (dtl dataDevTestLabAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("unique_identifier"))
}

func (dtl dataDevTestLabAttributes) Timeouts() datadevtestlab.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadevtestlab.TimeoutsAttributes](dtl.ref.Append("timeouts"))
}
