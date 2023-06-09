// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacosmosdbsqldatabase "github.com/golingon/terraproviders/azurerm/3.63.0/datacosmosdbsqldatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCosmosdbSqlDatabase creates a new instance of [DataCosmosdbSqlDatabase].
func NewDataCosmosdbSqlDatabase(name string, args DataCosmosdbSqlDatabaseArgs) *DataCosmosdbSqlDatabase {
	return &DataCosmosdbSqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCosmosdbSqlDatabase)(nil)

// DataCosmosdbSqlDatabase represents the Terraform data resource azurerm_cosmosdb_sql_database.
type DataCosmosdbSqlDatabase struct {
	Name string
	Args DataCosmosdbSqlDatabaseArgs
}

// DataSource returns the Terraform object type for [DataCosmosdbSqlDatabase].
func (csd *DataCosmosdbSqlDatabase) DataSource() string {
	return "azurerm_cosmosdb_sql_database"
}

// LocalName returns the local name for [DataCosmosdbSqlDatabase].
func (csd *DataCosmosdbSqlDatabase) LocalName() string {
	return csd.Name
}

// Configuration returns the configuration (args) for [DataCosmosdbSqlDatabase].
func (csd *DataCosmosdbSqlDatabase) Configuration() interface{} {
	return csd.Args
}

// Attributes returns the attributes for [DataCosmosdbSqlDatabase].
func (csd *DataCosmosdbSqlDatabase) Attributes() dataCosmosdbSqlDatabaseAttributes {
	return dataCosmosdbSqlDatabaseAttributes{ref: terra.ReferenceDataResource(csd)}
}

// DataCosmosdbSqlDatabaseArgs contains the configurations for azurerm_cosmosdb_sql_database.
type DataCosmosdbSqlDatabaseArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AutoscaleSettings: min=0
	AutoscaleSettings []datacosmosdbsqldatabase.AutoscaleSettings `hcl:"autoscale_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacosmosdbsqldatabase.Timeouts `hcl:"timeouts,block"`
}
type dataCosmosdbSqlDatabaseAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_database.
func (csd dataCosmosdbSqlDatabaseAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(csd.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_database.
func (csd dataCosmosdbSqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_database.
func (csd dataCosmosdbSqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_database.
func (csd dataCosmosdbSqlDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(csd.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_sql_database.
func (csd dataCosmosdbSqlDatabaseAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(csd.ref.Append("throughput"))
}

func (csd dataCosmosdbSqlDatabaseAttributes) AutoscaleSettings() terra.ListValue[datacosmosdbsqldatabase.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[datacosmosdbsqldatabase.AutoscaleSettingsAttributes](csd.ref.Append("autoscale_settings"))
}

func (csd dataCosmosdbSqlDatabaseAttributes) Timeouts() datacosmosdbsqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacosmosdbsqldatabase.TimeoutsAttributes](csd.ref.Append("timeouts"))
}
