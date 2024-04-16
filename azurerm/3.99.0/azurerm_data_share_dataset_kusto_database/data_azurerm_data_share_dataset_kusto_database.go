// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_share_dataset_kusto_database

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_data_share_dataset_kusto_database.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adsdkd *DataSource) DataSource() string {
	return "azurerm_data_share_dataset_kusto_database"
}

// LocalName returns the local name for [DataSource].
func (adsdkd *DataSource) LocalName() string {
	return adsdkd.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adsdkd *DataSource) Configuration() interface{} {
	return adsdkd.Args
}

// Attributes returns the attributes for [DataSource].
func (adsdkd *DataSource) Attributes() dataAzurermDataShareDatasetKustoDatabaseAttributes {
	return dataAzurermDataShareDatasetKustoDatabaseAttributes{ref: terra.ReferenceDataSource(adsdkd)}
}

// DataArgs contains the configurations for azurerm_data_share_dataset_kusto_database.
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

type dataAzurermDataShareDatasetKustoDatabaseAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_kusto_database.
func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(adsdkd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_kusto_database.
func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adsdkd.ref.Append("id"))
}

// KustoClusterLocation returns a reference to field kusto_cluster_location of azurerm_data_share_dataset_kusto_database.
func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) KustoClusterLocation() terra.StringValue {
	return terra.ReferenceAsString(adsdkd.ref.Append("kusto_cluster_location"))
}

// KustoDatabaseId returns a reference to field kusto_database_id of azurerm_data_share_dataset_kusto_database.
func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) KustoDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(adsdkd.ref.Append("kusto_database_id"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_kusto_database.
func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adsdkd.ref.Append("name"))
}

// ShareId returns a reference to field share_id of azurerm_data_share_dataset_kusto_database.
func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(adsdkd.ref.Append("share_id"))
}

func (adsdkd dataAzurermDataShareDatasetKustoDatabaseAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](adsdkd.ref.Append("timeouts"))
}
