// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatabasemigrationservice "github.com/golingon/terraproviders/azurerm/3.68.0/datadatabasemigrationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDatabaseMigrationService creates a new instance of [DataDatabaseMigrationService].
func NewDataDatabaseMigrationService(name string, args DataDatabaseMigrationServiceArgs) *DataDatabaseMigrationService {
	return &DataDatabaseMigrationService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatabaseMigrationService)(nil)

// DataDatabaseMigrationService represents the Terraform data resource azurerm_database_migration_service.
type DataDatabaseMigrationService struct {
	Name string
	Args DataDatabaseMigrationServiceArgs
}

// DataSource returns the Terraform object type for [DataDatabaseMigrationService].
func (dms *DataDatabaseMigrationService) DataSource() string {
	return "azurerm_database_migration_service"
}

// LocalName returns the local name for [DataDatabaseMigrationService].
func (dms *DataDatabaseMigrationService) LocalName() string {
	return dms.Name
}

// Configuration returns the configuration (args) for [DataDatabaseMigrationService].
func (dms *DataDatabaseMigrationService) Configuration() interface{} {
	return dms.Args
}

// Attributes returns the attributes for [DataDatabaseMigrationService].
func (dms *DataDatabaseMigrationService) Attributes() dataDatabaseMigrationServiceAttributes {
	return dataDatabaseMigrationServiceAttributes{ref: terra.ReferenceDataResource(dms)}
}

// DataDatabaseMigrationServiceArgs contains the configurations for azurerm_database_migration_service.
type DataDatabaseMigrationServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadatabasemigrationservice.Timeouts `hcl:"timeouts,block"`
}
type dataDatabaseMigrationServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("sku_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_database_migration_service.
func (dms dataDatabaseMigrationServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dms.ref.Append("tags"))
}

func (dms dataDatabaseMigrationServiceAttributes) Timeouts() datadatabasemigrationservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatabasemigrationservice.TimeoutsAttributes](dms.ref.Append("timeouts"))
}
