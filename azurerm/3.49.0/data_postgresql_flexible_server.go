// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapostgresqlflexibleserver "github.com/golingon/terraproviders/azurerm/3.49.0/datapostgresqlflexibleserver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataPostgresqlFlexibleServer(name string, args DataPostgresqlFlexibleServerArgs) *DataPostgresqlFlexibleServer {
	return &DataPostgresqlFlexibleServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPostgresqlFlexibleServer)(nil)

type DataPostgresqlFlexibleServer struct {
	Name string
	Args DataPostgresqlFlexibleServerArgs
}

func (pfs *DataPostgresqlFlexibleServer) DataSource() string {
	return "azurerm_postgresql_flexible_server"
}

func (pfs *DataPostgresqlFlexibleServer) LocalName() string {
	return pfs.Name
}

func (pfs *DataPostgresqlFlexibleServer) Configuration() interface{} {
	return pfs.Args
}

func (pfs *DataPostgresqlFlexibleServer) Attributes() dataPostgresqlFlexibleServerAttributes {
	return dataPostgresqlFlexibleServerAttributes{ref: terra.ReferenceDataResource(pfs)}
}

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

func (pfs dataPostgresqlFlexibleServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("administrator_login"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceNumber(pfs.ref.Append("backup_retention_days"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) DelegatedSubnetId() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("delegated_subnet_id"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("fqdn"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("id"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Location() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("location"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("name"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(pfs.ref.Append("public_network_access_enabled"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("resource_group_name"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("sku_name"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceNumber(pfs.ref.Append("storage_mb"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](pfs.ref.Append("tags"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Version() terra.StringValue {
	return terra.ReferenceString(pfs.ref.Append("version"))
}

func (pfs dataPostgresqlFlexibleServerAttributes) Timeouts() datapostgresqlflexibleserver.TimeoutsAttributes {
	return terra.ReferenceSingle[datapostgresqlflexibleserver.TimeoutsAttributes](pfs.ref.Append("timeouts"))
}
