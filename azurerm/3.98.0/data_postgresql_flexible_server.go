// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datapostgresqlflexibleserver "github.com/golingon/terraproviders/azurerm/3.98.0/datapostgresqlflexibleserver"
)

// NewDataPostgresqlFlexibleServer creates a new instance of [DataPostgresqlFlexibleServer].
func NewDataPostgresqlFlexibleServer(name string, args DataPostgresqlFlexibleServerArgs) *DataPostgresqlFlexibleServer {
	return &DataPostgresqlFlexibleServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPostgresqlFlexibleServer)(nil)

// DataPostgresqlFlexibleServer represents the Terraform data resource azurerm_postgresql_flexible_server.
type DataPostgresqlFlexibleServer struct {
	Name string
	Args DataPostgresqlFlexibleServerArgs
}

// DataSource returns the Terraform object type for [DataPostgresqlFlexibleServer].
func (pfs *DataPostgresqlFlexibleServer) DataSource() string {
	return "azurerm_postgresql_flexible_server"
}

// LocalName returns the local name for [DataPostgresqlFlexibleServer].
func (pfs *DataPostgresqlFlexibleServer) LocalName() string {
	return pfs.Name
}

// Configuration returns the configuration (args) for [DataPostgresqlFlexibleServer].
func (pfs *DataPostgresqlFlexibleServer) Configuration() interface{} {
	return pfs.Args
}

// Attributes returns the attributes for [DataPostgresqlFlexibleServer].
func (pfs *DataPostgresqlFlexibleServer) Attributes() dataPostgresqlFlexibleServerAttributes {
	return dataPostgresqlFlexibleServerAttributes{ref: terra.ReferenceDataResource(pfs)}
}

// DataPostgresqlFlexibleServerArgs contains the configurations for azurerm_postgresql_flexible_server.
type DataPostgresqlFlexibleServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datapostgresqlflexibleserver.Timeouts `hcl:"timeouts,block"`
}
type dataPostgresqlFlexibleServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("administrator_login"))
}

// AutoGrowEnabled returns a reference to field auto_grow_enabled of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) AutoGrowEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pfs.ref.Append("auto_grow_enabled"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(pfs.ref.Append("backup_retention_days"))
}

// DelegatedSubnetId returns a reference to field delegated_subnet_id of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) DelegatedSubnetId() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("delegated_subnet_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pfs.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("sku_name"))
}

// StorageMb returns a reference to field storage_mb of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceAsNumber(pfs.ref.Append("storage_mb"))
}

// Tags returns a reference to field tags of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pfs.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_postgresql_flexible_server.
func (pfs dataPostgresqlFlexibleServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(pfs.ref.Append("version"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Timeouts() datapostgresqlflexibleserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapostgresqlflexibleserver.TimeoutsAttributes](pfs.ref.Append("timeouts"))
}
