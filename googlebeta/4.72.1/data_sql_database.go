// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSqlDatabase creates a new instance of [DataSqlDatabase].
func NewDataSqlDatabase(name string, args DataSqlDatabaseArgs) *DataSqlDatabase {
	return &DataSqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlDatabase)(nil)

// DataSqlDatabase represents the Terraform data resource google_sql_database.
type DataSqlDatabase struct {
	Name string
	Args DataSqlDatabaseArgs
}

// DataSource returns the Terraform object type for [DataSqlDatabase].
func (sd *DataSqlDatabase) DataSource() string {
	return "google_sql_database"
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

// DataSqlDatabaseArgs contains the configurations for google_sql_database.
type DataSqlDatabaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataSqlDatabaseAttributes struct {
	ref terra.Reference
}

// Charset returns a reference to field charset of google_sql_database.
func (sd dataSqlDatabaseAttributes) Charset() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("charset"))
}

// Collation returns a reference to field collation of google_sql_database.
func (sd dataSqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("collation"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_sql_database.
func (sd dataSqlDatabaseAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("deletion_policy"))
}

// Id returns a reference to field id of google_sql_database.
func (sd dataSqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_database.
func (sd dataSqlDatabaseAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("instance"))
}

// Name returns a reference to field name of google_sql_database.
func (sd dataSqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
}

// Project returns a reference to field project of google_sql_database.
func (sd dataSqlDatabaseAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_sql_database.
func (sd dataSqlDatabaseAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("self_link"))
}
