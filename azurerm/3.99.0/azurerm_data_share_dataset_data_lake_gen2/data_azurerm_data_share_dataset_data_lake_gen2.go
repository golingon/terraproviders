// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_share_dataset_data_lake_gen2

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_data_share_dataset_data_lake_gen2.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adsddlg *DataSource) DataSource() string {
	return "azurerm_data_share_dataset_data_lake_gen2"
}

// LocalName returns the local name for [DataSource].
func (adsddlg *DataSource) LocalName() string {
	return adsddlg.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adsddlg *DataSource) Configuration() interface{} {
	return adsddlg.Args
}

// Attributes returns the attributes for [DataSource].
func (adsddlg *DataSource) Attributes() dataAzurermDataShareDatasetDataLakeGen2Attributes {
	return dataAzurermDataShareDatasetDataLakeGen2Attributes{ref: terra.ReferenceDataSource(adsddlg)}
}

// DataArgs contains the configurations for azurerm_data_share_dataset_data_lake_gen2.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareId: string, required
	ShareId terra.StringValue `hcl:"share_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermDataShareDatasetDataLakeGen2Attributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("display_name"))
}

// FilePath returns a reference to field file_path of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) FilePath() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("file_path"))
}

// FileSystemName returns a reference to field file_system_name of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) FileSystemName() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("file_system_name"))
}

// FolderPath returns a reference to field folder_path of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) FolderPath() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("folder_path"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("name"))
}

// ShareId returns a reference to field share_id of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("share_id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_data_share_dataset_data_lake_gen2.
func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(adsddlg.ref.Append("storage_account_id"))
}

func (adsddlg dataAzurermDataShareDatasetDataLakeGen2Attributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](adsddlg.ref.Append("timeouts"))
}
