// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatabasemigrationservice "github.com/golingon/terraproviders/azurerm/3.49.0/datadatabasemigrationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataDatabaseMigrationService(name string, args DataDatabaseMigrationServiceArgs) *DataDatabaseMigrationService {
	return &DataDatabaseMigrationService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatabaseMigrationService)(nil)

type DataDatabaseMigrationService struct {
	Name string
	Args DataDatabaseMigrationServiceArgs
}

func (dms *DataDatabaseMigrationService) DataSource() string {
	return "azurerm_database_migration_service"
}

func (dms *DataDatabaseMigrationService) LocalName() string {
	return dms.Name
}

func (dms *DataDatabaseMigrationService) Configuration() interface{} {
	return dms.Args
}

func (dms *DataDatabaseMigrationService) Attributes() dataDatabaseMigrationServiceAttributes {
	return dataDatabaseMigrationServiceAttributes{ref: terra.ReferenceDataResource(dms)}
}

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

func (dms dataDatabaseMigrationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dms.ref.Append("id"))
}

func (dms dataDatabaseMigrationServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceString(dms.ref.Append("location"))
}

func (dms dataDatabaseMigrationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dms.ref.Append("name"))
}

func (dms dataDatabaseMigrationServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(dms.ref.Append("resource_group_name"))
}

func (dms dataDatabaseMigrationServiceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceString(dms.ref.Append("sku_name"))
}

func (dms dataDatabaseMigrationServiceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(dms.ref.Append("subnet_id"))
}

func (dms dataDatabaseMigrationServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dms.ref.Append("tags"))
}

func (dms dataDatabaseMigrationServiceAttributes) Timeouts() datadatabasemigrationservice.TimeoutsAttributes {
	return terra.ReferenceSingle[datadatabasemigrationservice.TimeoutsAttributes](dms.ref.Append("timeouts"))
}
