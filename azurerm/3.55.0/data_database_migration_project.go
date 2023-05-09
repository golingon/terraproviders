// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatabasemigrationproject "github.com/golingon/terraproviders/azurerm/3.55.0/datadatabasemigrationproject"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDatabaseMigrationProject creates a new instance of [DataDatabaseMigrationProject].
func NewDataDatabaseMigrationProject(name string, args DataDatabaseMigrationProjectArgs) *DataDatabaseMigrationProject {
	return &DataDatabaseMigrationProject{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatabaseMigrationProject)(nil)

// DataDatabaseMigrationProject represents the Terraform data resource azurerm_database_migration_project.
type DataDatabaseMigrationProject struct {
	Name string
	Args DataDatabaseMigrationProjectArgs
}

// DataSource returns the Terraform object type for [DataDatabaseMigrationProject].
func (dmp *DataDatabaseMigrationProject) DataSource() string {
	return "azurerm_database_migration_project"
}

// LocalName returns the local name for [DataDatabaseMigrationProject].
func (dmp *DataDatabaseMigrationProject) LocalName() string {
	return dmp.Name
}

// Configuration returns the configuration (args) for [DataDatabaseMigrationProject].
func (dmp *DataDatabaseMigrationProject) Configuration() interface{} {
	return dmp.Args
}

// Attributes returns the attributes for [DataDatabaseMigrationProject].
func (dmp *DataDatabaseMigrationProject) Attributes() dataDatabaseMigrationProjectAttributes {
	return dataDatabaseMigrationProjectAttributes{ref: terra.ReferenceDataResource(dmp)}
}

// DataDatabaseMigrationProjectArgs contains the configurations for azurerm_database_migration_project.
type DataDatabaseMigrationProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadatabasemigrationproject.Timeouts `hcl:"timeouts,block"`
}
type dataDatabaseMigrationProjectAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("resource_group_name"))
}

// ServiceName returns a reference to field service_name of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("service_name"))
}

// SourcePlatform returns a reference to field source_platform of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) SourcePlatform() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("source_platform"))
}

// Tags returns a reference to field tags of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dmp.ref.Append("tags"))
}

// TargetPlatform returns a reference to field target_platform of azurerm_database_migration_project.
func (dmp dataDatabaseMigrationProjectAttributes) TargetPlatform() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("target_platform"))
}

func (dmp dataDatabaseMigrationProjectAttributes) Timeouts() datadatabasemigrationproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatabasemigrationproject.TimeoutsAttributes](dmp.ref.Append("timeouts"))
}
