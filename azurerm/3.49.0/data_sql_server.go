// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasqlserver "github.com/golingon/terraproviders/azurerm/3.49.0/datasqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataSqlServer(name string, args DataSqlServerArgs) *DataSqlServer {
	return &DataSqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlServer)(nil)

type DataSqlServer struct {
	Name string
	Args DataSqlServerArgs
}

func (ss *DataSqlServer) DataSource() string {
	return "azurerm_sql_server"
}

func (ss *DataSqlServer) LocalName() string {
	return ss.Name
}

func (ss *DataSqlServer) Configuration() interface{} {
	return ss.Args
}

func (ss *DataSqlServer) Attributes() dataSqlServerAttributes {
	return dataSqlServerAttributes{ref: terra.ReferenceDataResource(ss)}
}

type DataSqlServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datasqlserver.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasqlserver.Timeouts `hcl:"timeouts,block"`
}
type dataSqlServerAttributes struct {
	ref terra.Reference
}

func (ss dataSqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("administrator_login"))
}

func (ss dataSqlServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("fqdn"))
}

func (ss dataSqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("id"))
}

func (ss dataSqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("location"))
}

func (ss dataSqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("name"))
}

func (ss dataSqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("resource_group_name"))
}

func (ss dataSqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ss.ref.Append("tags"))
}

func (ss dataSqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("version"))
}

func (ss dataSqlServerAttributes) Identity() terra.ListValue[datasqlserver.IdentityAttributes] {
	return terra.ReferenceList[datasqlserver.IdentityAttributes](ss.ref.Append("identity"))
}

func (ss dataSqlServerAttributes) Timeouts() datasqlserver.TimeoutsAttributes {
	return terra.ReferenceSingle[datasqlserver.TimeoutsAttributes](ss.ref.Append("timeouts"))
}
