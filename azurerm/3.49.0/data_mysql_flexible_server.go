// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamysqlflexibleserver "github.com/golingon/terraproviders/azurerm/3.49.0/datamysqlflexibleserver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataMysqlFlexibleServer(name string, args DataMysqlFlexibleServerArgs) *DataMysqlFlexibleServer {
	return &DataMysqlFlexibleServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMysqlFlexibleServer)(nil)

type DataMysqlFlexibleServer struct {
	Name string
	Args DataMysqlFlexibleServerArgs
}

func (mfs *DataMysqlFlexibleServer) DataSource() string {
	return "azurerm_mysql_flexible_server"
}

func (mfs *DataMysqlFlexibleServer) LocalName() string {
	return mfs.Name
}

func (mfs *DataMysqlFlexibleServer) Configuration() interface{} {
	return mfs.Args
}

func (mfs *DataMysqlFlexibleServer) Attributes() dataMysqlFlexibleServerAttributes {
	return dataMysqlFlexibleServerAttributes{ref: terra.ReferenceDataResource(mfs)}
}

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

func (mfs dataMysqlFlexibleServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("administrator_login"))
}

func (mfs dataMysqlFlexibleServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceNumber(mfs.ref.Append("backup_retention_days"))
}

func (mfs dataMysqlFlexibleServerAttributes) DelegatedSubnetId() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("delegated_subnet_id"))
}

func (mfs dataMysqlFlexibleServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("fqdn"))
}

func (mfs dataMysqlFlexibleServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceBool(mfs.ref.Append("geo_redundant_backup_enabled"))
}

func (mfs dataMysqlFlexibleServerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("id"))
}

func (mfs dataMysqlFlexibleServerAttributes) Location() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("location"))
}

func (mfs dataMysqlFlexibleServerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("name"))
}

func (mfs dataMysqlFlexibleServerAttributes) PrivateDnsZoneId() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("private_dns_zone_id"))
}

func (mfs dataMysqlFlexibleServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(mfs.ref.Append("public_network_access_enabled"))
}

func (mfs dataMysqlFlexibleServerAttributes) ReplicaCapacity() terra.NumberValue {
	return terra.ReferenceNumber(mfs.ref.Append("replica_capacity"))
}

func (mfs dataMysqlFlexibleServerAttributes) ReplicationRole() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("replication_role"))
}

func (mfs dataMysqlFlexibleServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("resource_group_name"))
}

func (mfs dataMysqlFlexibleServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("restore_point_in_time"))
}

func (mfs dataMysqlFlexibleServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("sku_name"))
}

func (mfs dataMysqlFlexibleServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](mfs.ref.Append("tags"))
}

func (mfs dataMysqlFlexibleServerAttributes) Version() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("version"))
}

func (mfs dataMysqlFlexibleServerAttributes) Zone() terra.StringValue {
	return terra.ReferenceString(mfs.ref.Append("zone"))
}

func (mfs dataMysqlFlexibleServerAttributes) HighAvailability() terra.ListValue[datamysqlflexibleserver.HighAvailabilityAttributes] {
	return terra.ReferenceList[datamysqlflexibleserver.HighAvailabilityAttributes](mfs.ref.Append("high_availability"))
}

func (mfs dataMysqlFlexibleServerAttributes) MaintenanceWindow() terra.ListValue[datamysqlflexibleserver.MaintenanceWindowAttributes] {
	return terra.ReferenceList[datamysqlflexibleserver.MaintenanceWindowAttributes](mfs.ref.Append("maintenance_window"))
}

func (mfs dataMysqlFlexibleServerAttributes) Storage() terra.ListValue[datamysqlflexibleserver.StorageAttributes] {
	return terra.ReferenceList[datamysqlflexibleserver.StorageAttributes](mfs.ref.Append("storage"))
}

func (mfs dataMysqlFlexibleServerAttributes) Timeouts() datamysqlflexibleserver.TimeoutsAttributes {
	return terra.ReferenceSingle[datamysqlflexibleserver.TimeoutsAttributes](mfs.ref.Append("timeouts"))
}
