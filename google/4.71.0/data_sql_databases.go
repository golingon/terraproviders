// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datasqldatabases "github.com/golingon/terraproviders/google/4.71.0/datasqldatabases"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSqlDatabases creates a new instance of [DataSqlDatabases].
func NewDataSqlDatabases(name string, args DataSqlDatabasesArgs) *DataSqlDatabases {
	return &DataSqlDatabases{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlDatabases)(nil)

// DataSqlDatabases represents the Terraform data resource google_sql_databases.
type DataSqlDatabases struct {
	Name string
	Args DataSqlDatabasesArgs
}

// DataSource returns the Terraform object type for [DataSqlDatabases].
func (sd *DataSqlDatabases) DataSource() string {
	return "google_sql_databases"
}

// LocalName returns the local name for [DataSqlDatabases].
func (sd *DataSqlDatabases) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [DataSqlDatabases].
func (sd *DataSqlDatabases) Configuration() interface{} {
	return sd.Args
}

// Attributes returns the attributes for [DataSqlDatabases].
func (sd *DataSqlDatabases) Attributes() dataSqlDatabasesAttributes {
	return dataSqlDatabasesAttributes{ref: terra.ReferenceDataResource(sd)}
}

// DataSqlDatabasesArgs contains the configurations for google_sql_databases.
type DataSqlDatabasesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Databases: min=0
	Databases []datasqldatabases.Databases `hcl:"databases,block" validate:"min=0"`
}
type dataSqlDatabasesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_sql_databases.
func (sd dataSqlDatabasesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_databases.
func (sd dataSqlDatabasesAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("instance"))
}

// Project returns a reference to field project of google_sql_databases.
func (sd dataSqlDatabasesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("project"))
}

func (sd dataSqlDatabasesAttributes) Databases() terra.ListValue[datasqldatabases.DatabasesAttributes] {
	return terra.ReferenceAsList[datasqldatabases.DatabasesAttributes](sd.ref.Append("databases"))
}
