// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamssqlserver "github.com/golingon/terraproviders/azurerm/3.58.0/datamssqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMssqlServer creates a new instance of [DataMssqlServer].
func NewDataMssqlServer(name string, args DataMssqlServerArgs) *DataMssqlServer {
	return &DataMssqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMssqlServer)(nil)

// DataMssqlServer represents the Terraform data resource azurerm_mssql_server.
type DataMssqlServer struct {
	Name string
	Args DataMssqlServerArgs
}

// DataSource returns the Terraform object type for [DataMssqlServer].
func (ms *DataMssqlServer) DataSource() string {
	return "azurerm_mssql_server"
}

// LocalName returns the local name for [DataMssqlServer].
func (ms *DataMssqlServer) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [DataMssqlServer].
func (ms *DataMssqlServer) Configuration() interface{} {
	return ms.Args
}

// Attributes returns the attributes for [DataMssqlServer].
func (ms *DataMssqlServer) Attributes() dataMssqlServerAttributes {
	return dataMssqlServerAttributes{ref: terra.ReferenceDataResource(ms)}
}

// DataMssqlServerArgs contains the configurations for azurerm_mssql_server.
type DataMssqlServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datamssqlserver.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamssqlserver.Timeouts `hcl:"timeouts,block"`
}
type dataMssqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login"))
}

// FullyQualifiedDomainName returns a reference to field fully_qualified_domain_name of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) FullyQualifiedDomainName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("fully_qualified_domain_name"))
}

// Id returns a reference to field id of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("resource_group_name"))
}

// RestorableDroppedDatabaseIds returns a reference to field restorable_dropped_database_ids of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) RestorableDroppedDatabaseIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ms.ref.Append("restorable_dropped_database_ids"))
}

// Tags returns a reference to field tags of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_mssql_server.
func (ms dataMssqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("version"))
}

func (ms dataMssqlServerAttributes) Identity() terra.ListValue[datamssqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[datamssqlserver.IdentityAttributes](ms.ref.Append("identity"))
}

func (ms dataMssqlServerAttributes) Timeouts() datamssqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamssqlserver.TimeoutsAttributes](ms.ref.Append("timeouts"))
}
