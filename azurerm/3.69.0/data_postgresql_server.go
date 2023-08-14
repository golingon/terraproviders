// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapostgresqlserver "github.com/golingon/terraproviders/azurerm/3.69.0/datapostgresqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPostgresqlServer creates a new instance of [DataPostgresqlServer].
func NewDataPostgresqlServer(name string, args DataPostgresqlServerArgs) *DataPostgresqlServer {
	return &DataPostgresqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPostgresqlServer)(nil)

// DataPostgresqlServer represents the Terraform data resource azurerm_postgresql_server.
type DataPostgresqlServer struct {
	Name string
	Args DataPostgresqlServerArgs
}

// DataSource returns the Terraform object type for [DataPostgresqlServer].
func (ps *DataPostgresqlServer) DataSource() string {
	return "azurerm_postgresql_server"
}

// LocalName returns the local name for [DataPostgresqlServer].
func (ps *DataPostgresqlServer) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [DataPostgresqlServer].
func (ps *DataPostgresqlServer) Configuration() interface{} {
	return ps.Args
}

// Attributes returns the attributes for [DataPostgresqlServer].
func (ps *DataPostgresqlServer) Attributes() dataPostgresqlServerAttributes {
	return dataPostgresqlServerAttributes{ref: terra.ReferenceDataResource(ps)}
}

// DataPostgresqlServerArgs contains the configurations for azurerm_postgresql_server.
type DataPostgresqlServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datapostgresqlserver.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datapostgresqlserver.Timeouts `hcl:"timeouts,block"`
}
type dataPostgresqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("administrator_login"))
}

// Fqdn returns a reference to field fqdn of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ps.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_postgresql_server.
func (ps dataPostgresqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("version"))
}

func (ps dataPostgresqlServerAttributes) Identity() terra.ListValue[datapostgresqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[datapostgresqlserver.IdentityAttributes](ps.ref.Append("identity"))
}

func (ps dataPostgresqlServerAttributes) Timeouts() datapostgresqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapostgresqlserver.TimeoutsAttributes](ps.ref.Append("timeouts"))
}
