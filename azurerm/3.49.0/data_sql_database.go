// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasqldatabase "github.com/golingon/terraproviders/azurerm/3.49.0/datasqldatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSqlDatabase creates a new instance of [DataSqlDatabase].
func NewDataSqlDatabase(name string, args DataSqlDatabaseArgs) *DataSqlDatabase {
	return &DataSqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlDatabase)(nil)

// DataSqlDatabase represents the Terraform data resource azurerm_sql_database.
type DataSqlDatabase struct {
	Name string
	Args DataSqlDatabaseArgs
}

// DataSource returns the Terraform object type for [DataSqlDatabase].
func (sd *DataSqlDatabase) DataSource() string {
	return "azurerm_sql_database"
}

// LocalName returns the local name for [DataSqlDatabase].
func (sd *DataSqlDatabase) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [DataSqlDatabase].
func (sd *DataSqlDatabase) Configuration() interface{} {
	return sd.Args
}

// Attributes returns the attributes for [DataSqlDatabase].
func (sd *DataSqlDatabase) Attributes() dataSqlDatabaseAttributes {
	return dataSqlDatabaseAttributes{ref: terra.ReferenceDataResource(sd)}
}

// DataSqlDatabaseArgs contains the configurations for azurerm_sql_database.
type DataSqlDatabaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *datasqldatabase.Timeouts `hcl:"timeouts,block"`
}
type dataSqlDatabaseAttributes struct {
	ref terra.Reference
}

// Collation returns a reference to field collation of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("collation"))
}

// DefaultSecondaryLocation returns a reference to field default_secondary_location of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) DefaultSecondaryLocation() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("default_secondary_location"))
}

// Edition returns a reference to field edition of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("edition"))
}

// ElasticPoolName returns a reference to field elastic_pool_name of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) ElasticPoolName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("elastic_pool_name"))
}

// FailoverGroupId returns a reference to field failover_group_id of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) FailoverGroupId() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("failover_group_id"))
}

// Id returns a reference to field id of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
}

// ReadScale returns a reference to field read_scale of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) ReadScale() terra.BoolValue {
	return terra.ReferenceAsBool(sd.ref.Append("read_scale"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("server_name"))
}

// Tags returns a reference to field tags of azurerm_sql_database.
func (sd dataSqlDatabaseAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("tags"))
}

func (sd dataSqlDatabaseAttributes) Timeouts() datasqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasqldatabase.TimeoutsAttributes](sd.ref.Append("timeouts"))
}
