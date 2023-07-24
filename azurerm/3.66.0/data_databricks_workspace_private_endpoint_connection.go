// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatabricksworkspaceprivateendpointconnection "github.com/golingon/terraproviders/azurerm/3.66.0/datadatabricksworkspaceprivateendpointconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDatabricksWorkspacePrivateEndpointConnection creates a new instance of [DataDatabricksWorkspacePrivateEndpointConnection].
func NewDataDatabricksWorkspacePrivateEndpointConnection(name string, args DataDatabricksWorkspacePrivateEndpointConnectionArgs) *DataDatabricksWorkspacePrivateEndpointConnection {
	return &DataDatabricksWorkspacePrivateEndpointConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatabricksWorkspacePrivateEndpointConnection)(nil)

// DataDatabricksWorkspacePrivateEndpointConnection represents the Terraform data resource azurerm_databricks_workspace_private_endpoint_connection.
type DataDatabricksWorkspacePrivateEndpointConnection struct {
	Name string
	Args DataDatabricksWorkspacePrivateEndpointConnectionArgs
}

// DataSource returns the Terraform object type for [DataDatabricksWorkspacePrivateEndpointConnection].
func (dwpec *DataDatabricksWorkspacePrivateEndpointConnection) DataSource() string {
	return "azurerm_databricks_workspace_private_endpoint_connection"
}

// LocalName returns the local name for [DataDatabricksWorkspacePrivateEndpointConnection].
func (dwpec *DataDatabricksWorkspacePrivateEndpointConnection) LocalName() string {
	return dwpec.Name
}

// Configuration returns the configuration (args) for [DataDatabricksWorkspacePrivateEndpointConnection].
func (dwpec *DataDatabricksWorkspacePrivateEndpointConnection) Configuration() interface{} {
	return dwpec.Args
}

// Attributes returns the attributes for [DataDatabricksWorkspacePrivateEndpointConnection].
func (dwpec *DataDatabricksWorkspacePrivateEndpointConnection) Attributes() dataDatabricksWorkspacePrivateEndpointConnectionAttributes {
	return dataDatabricksWorkspacePrivateEndpointConnectionAttributes{ref: terra.ReferenceDataResource(dwpec)}
}

// DataDatabricksWorkspacePrivateEndpointConnectionArgs contains the configurations for azurerm_databricks_workspace_private_endpoint_connection.
type DataDatabricksWorkspacePrivateEndpointConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrivateEndpointId: string, required
	PrivateEndpointId terra.StringValue `hcl:"private_endpoint_id,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Connections: min=0
	Connections []datadatabricksworkspaceprivateendpointconnection.Connections `hcl:"connections,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadatabricksworkspaceprivateendpointconnection.Timeouts `hcl:"timeouts,block"`
}
type dataDatabricksWorkspacePrivateEndpointConnectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_databricks_workspace_private_endpoint_connection.
func (dwpec dataDatabricksWorkspacePrivateEndpointConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dwpec.ref.Append("id"))
}

// PrivateEndpointId returns a reference to field private_endpoint_id of azurerm_databricks_workspace_private_endpoint_connection.
func (dwpec dataDatabricksWorkspacePrivateEndpointConnectionAttributes) PrivateEndpointId() terra.StringValue {
	return terra.ReferenceAsString(dwpec.ref.Append("private_endpoint_id"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_databricks_workspace_private_endpoint_connection.
func (dwpec dataDatabricksWorkspacePrivateEndpointConnectionAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(dwpec.ref.Append("workspace_id"))
}

func (dwpec dataDatabricksWorkspacePrivateEndpointConnectionAttributes) Connections() terra.ListValue[datadatabricksworkspaceprivateendpointconnection.ConnectionsAttributes] {
	return terra.ReferenceAsList[datadatabricksworkspaceprivateendpointconnection.ConnectionsAttributes](dwpec.ref.Append("connections"))
}

func (dwpec dataDatabricksWorkspacePrivateEndpointConnectionAttributes) Timeouts() datadatabricksworkspaceprivateendpointconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatabricksworkspaceprivateendpointconnection.TimeoutsAttributes](dwpec.ref.Append("timeouts"))
}
