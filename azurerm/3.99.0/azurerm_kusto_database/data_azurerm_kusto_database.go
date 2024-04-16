// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_kusto_database

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_kusto_database.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (akd *DataSource) DataSource() string {
	return "azurerm_kusto_database"
}

// LocalName returns the local name for [DataSource].
func (akd *DataSource) LocalName() string {
	return akd.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (akd *DataSource) Configuration() interface{} {
	return akd.Args
}

// Attributes returns the attributes for [DataSource].
func (akd *DataSource) Attributes() dataAzurermKustoDatabaseAttributes {
	return dataAzurermKustoDatabaseAttributes{ref: terra.ReferenceDataSource(akd)}
}

// DataArgs contains the configurations for azurerm_kusto_database.
type DataArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermKustoDatabaseAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("cluster_name"))
}

// HotCachePeriod returns a reference to field hot_cache_period of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) HotCachePeriod() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("hot_cache_period"))
}

// Id returns a reference to field id of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("resource_group_name"))
}

// Size returns a reference to field size of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(akd.ref.Append("size"))
}

// SoftDeletePeriod returns a reference to field soft_delete_period of azurerm_kusto_database.
func (akd dataAzurermKustoDatabaseAttributes) SoftDeletePeriod() terra.StringValue {
	return terra.ReferenceAsString(akd.ref.Append("soft_delete_period"))
}

func (akd dataAzurermKustoDatabaseAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](akd.ref.Append("timeouts"))
}
