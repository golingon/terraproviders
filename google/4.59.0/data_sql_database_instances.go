// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datasqldatabaseinstances "github.com/golingon/terraproviders/google/4.59.0/datasqldatabaseinstances"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSqlDatabaseInstances creates a new instance of [DataSqlDatabaseInstances].
func NewDataSqlDatabaseInstances(name string, args DataSqlDatabaseInstancesArgs) *DataSqlDatabaseInstances {
	return &DataSqlDatabaseInstances{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlDatabaseInstances)(nil)

// DataSqlDatabaseInstances represents the Terraform data resource google_sql_database_instances.
type DataSqlDatabaseInstances struct {
	Name string
	Args DataSqlDatabaseInstancesArgs
}

// DataSource returns the Terraform object type for [DataSqlDatabaseInstances].
func (sdi *DataSqlDatabaseInstances) DataSource() string {
	return "google_sql_database_instances"
}

// LocalName returns the local name for [DataSqlDatabaseInstances].
func (sdi *DataSqlDatabaseInstances) LocalName() string {
	return sdi.Name
}

// Configuration returns the configuration (args) for [DataSqlDatabaseInstances].
func (sdi *DataSqlDatabaseInstances) Configuration() interface{} {
	return sdi.Args
}

// Attributes returns the attributes for [DataSqlDatabaseInstances].
func (sdi *DataSqlDatabaseInstances) Attributes() dataSqlDatabaseInstancesAttributes {
	return dataSqlDatabaseInstancesAttributes{ref: terra.ReferenceDataResource(sdi)}
}

// DataSqlDatabaseInstancesArgs contains the configurations for google_sql_database_instances.
type DataSqlDatabaseInstancesArgs struct {
	// DatabaseVersion: string, optional
	DatabaseVersion terra.StringValue `hcl:"database_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tier: string, optional
	Tier terra.StringValue `hcl:"tier,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Instances: min=0
	Instances []datasqldatabaseinstances.Instances `hcl:"instances,block" validate:"min=0"`
}
type dataSqlDatabaseInstancesAttributes struct {
	ref terra.Reference
}

// DatabaseVersion returns a reference to field database_version of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("database_version"))
}

// Id returns a reference to field id of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("id"))
}

// Project returns a reference to field project of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("project"))
}

// Region returns a reference to field region of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("region"))
}

// State returns a reference to field state of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("state"))
}

// Tier returns a reference to field tier of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("tier"))
}

// Zone returns a reference to field zone of google_sql_database_instances.
func (sdi dataSqlDatabaseInstancesAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("zone"))
}

func (sdi dataSqlDatabaseInstancesAttributes) Instances() terra.ListValue[datasqldatabaseinstances.InstancesAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstances.InstancesAttributes](sdi.ref.Append("instances"))
}
