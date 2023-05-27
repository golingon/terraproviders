// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamysqlflexibleserver "github.com/golingon/terraproviders/azurerm/3.58.0/datamysqlflexibleserver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMysqlFlexibleServer creates a new instance of [DataMysqlFlexibleServer].
func NewDataMysqlFlexibleServer(name string, args DataMysqlFlexibleServerArgs) *DataMysqlFlexibleServer {
	return &DataMysqlFlexibleServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMysqlFlexibleServer)(nil)

// DataMysqlFlexibleServer represents the Terraform data resource azurerm_mysql_flexible_server.
type DataMysqlFlexibleServer struct {
	Name string
	Args DataMysqlFlexibleServerArgs
}

// DataSource returns the Terraform object type for [DataMysqlFlexibleServer].
func (mfs *DataMysqlFlexibleServer) DataSource() string {
	return "azurerm_mysql_flexible_server"
}

// LocalName returns the local name for [DataMysqlFlexibleServer].
func (mfs *DataMysqlFlexibleServer) LocalName() string {
	return mfs.Name
}

// Configuration returns the configuration (args) for [DataMysqlFlexibleServer].
func (mfs *DataMysqlFlexibleServer) Configuration() interface{} {
	return mfs.Args
}

// Attributes returns the attributes for [DataMysqlFlexibleServer].
func (mfs *DataMysqlFlexibleServer) Attributes() dataMysqlFlexibleServerAttributes {
	return dataMysqlFlexibleServerAttributes{ref: terra.ReferenceDataResource(mfs)}
}

// DataMysqlFlexibleServerArgs contains the configurations for azurerm_mysql_flexible_server.
type DataMysqlFlexibleServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// HighAvailability: min=0
	HighAvailability []datamysqlflexibleserver.HighAvailability `hcl:"high_availability,block" validate:"min=0"`
	// MaintenanceWindow: min=0
	MaintenanceWindow []datamysqlflexibleserver.MaintenanceWindow `hcl:"maintenance_window,block" validate:"min=0"`
	// Storage: min=0
	Storage []datamysqlflexibleserver.Storage `hcl:"storage,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamysqlflexibleserver.Timeouts `hcl:"timeouts,block"`
}
type dataMysqlFlexibleServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("administrator_login"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(mfs.ref.Append("backup_retention_days"))
}

// DelegatedSubnetId returns a reference to field delegated_subnet_id of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) DelegatedSubnetId() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("delegated_subnet_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("fqdn"))
}

// GeoRedundantBackupEnabled returns a reference to field geo_redundant_backup_enabled of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mfs.ref.Append("geo_redundant_backup_enabled"))
}

// Id returns a reference to field id of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("name"))
}

// PrivateDnsZoneId returns a reference to field private_dns_zone_id of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) PrivateDnsZoneId() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("private_dns_zone_id"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mfs.ref.Append("public_network_access_enabled"))
}

// ReplicaCapacity returns a reference to field replica_capacity of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) ReplicaCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(mfs.ref.Append("replica_capacity"))
}

// ReplicationRole returns a reference to field replication_role of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) ReplicationRole() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("replication_role"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("resource_group_name"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("restore_point_in_time"))
}

// SkuName returns a reference to field sku_name of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mfs.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("version"))
}

// Zone returns a reference to field zone of azurerm_mysql_flexible_server.
func (mfs dataMysqlFlexibleServerAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(mfs.ref.Append("zone"))
}

func (mfs dataMysqlFlexibleServerAttributes) HighAvailability() terra.ListValue[datamysqlflexibleserver.HighAvailabilityAttributes] {
	return terra.ReferenceAsList[datamysqlflexibleserver.HighAvailabilityAttributes](mfs.ref.Append("high_availability"))
}

func (mfs dataMysqlFlexibleServerAttributes) MaintenanceWindow() terra.ListValue[datamysqlflexibleserver.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[datamysqlflexibleserver.MaintenanceWindowAttributes](mfs.ref.Append("maintenance_window"))
}

func (mfs dataMysqlFlexibleServerAttributes) Storage() terra.ListValue[datamysqlflexibleserver.StorageAttributes] {
	return terra.ReferenceAsList[datamysqlflexibleserver.StorageAttributes](mfs.ref.Append("storage"))
}

func (mfs dataMysqlFlexibleServerAttributes) Timeouts() datamysqlflexibleserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamysqlflexibleserver.TimeoutsAttributes](mfs.ref.Append("timeouts"))
}
