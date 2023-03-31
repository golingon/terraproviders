// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadevtestlab "github.com/golingon/terraproviders/azurerm/3.49.0/datadevtestlab"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataDevTestLab(name string, args DataDevTestLabArgs) *DataDevTestLab {
	return &DataDevTestLab{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDevTestLab)(nil)

type DataDevTestLab struct {
	Name string
	Args DataDevTestLabArgs
}

func (dtl *DataDevTestLab) DataSource() string {
	return "azurerm_dev_test_lab"
}

func (dtl *DataDevTestLab) LocalName() string {
	return dtl.Name
}

func (dtl *DataDevTestLab) Configuration() interface{} {
	return dtl.Args
}

func (dtl *DataDevTestLab) Attributes() dataDevTestLabAttributes {
	return dataDevTestLabAttributes{ref: terra.ReferenceDataResource(dtl)}
}

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

func (dtl dataDevTestLabAttributes) ArtifactsStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("artifacts_storage_account_id"))
}

func (dtl dataDevTestLabAttributes) DefaultPremiumStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("default_premium_storage_account_id"))
}

func (dtl dataDevTestLabAttributes) DefaultStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("default_storage_account_id"))
}

func (dtl dataDevTestLabAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("id"))
}

func (dtl dataDevTestLabAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("key_vault_id"))
}

func (dtl dataDevTestLabAttributes) Location() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("location"))
}

func (dtl dataDevTestLabAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("name"))
}

func (dtl dataDevTestLabAttributes) PremiumDataDiskStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("premium_data_disk_storage_account_id"))
}

func (dtl dataDevTestLabAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("resource_group_name"))
}

func (dtl dataDevTestLabAttributes) StorageType() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("storage_type"))
}

func (dtl dataDevTestLabAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dtl.ref.Append("tags"))
}

func (dtl dataDevTestLabAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("unique_identifier"))
}

func (dtl dataDevTestLabAttributes) Timeouts() datadevtestlab.TimeoutsAttributes {
	return terra.ReferenceSingle[datadevtestlab.TimeoutsAttributes](dtl.ref.Append("timeouts"))
}
