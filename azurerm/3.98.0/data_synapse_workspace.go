// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datasynapseworkspace "github.com/golingon/terraproviders/azurerm/3.98.0/datasynapseworkspace"
)

// NewDataSynapseWorkspace creates a new instance of [DataSynapseWorkspace].
func NewDataSynapseWorkspace(name string, args DataSynapseWorkspaceArgs) *DataSynapseWorkspace {
	return &DataSynapseWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSynapseWorkspace)(nil)

// DataSynapseWorkspace represents the Terraform data resource azurerm_synapse_workspace.
type DataSynapseWorkspace struct {
	Name string
	Args DataSynapseWorkspaceArgs
}

// DataSource returns the Terraform object type for [DataSynapseWorkspace].
func (sw *DataSynapseWorkspace) DataSource() string {
	return "azurerm_synapse_workspace"
}

// LocalName returns the local name for [DataSynapseWorkspace].
func (sw *DataSynapseWorkspace) LocalName() string {
	return sw.Name
}

// Configuration returns the configuration (args) for [DataSynapseWorkspace].
func (sw *DataSynapseWorkspace) Configuration() interface{} {
	return sw.Args
}

// Attributes returns the attributes for [DataSynapseWorkspace].
func (sw *DataSynapseWorkspace) Attributes() dataSynapseWorkspaceAttributes {
	return dataSynapseWorkspaceAttributes{ref: terra.ReferenceDataResource(sw)}
}

// DataSynapseWorkspaceArgs contains the configurations for azurerm_synapse_workspace.
type DataSynapseWorkspaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datasynapseworkspace.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasynapseworkspace.Timeouts `hcl:"timeouts,block"`
}
type dataSynapseWorkspaceAttributes struct {
	ref terra.Reference
}

// ConnectivityEndpoints returns a reference to field connectivity_endpoints of azurerm_synapse_workspace.
func (sw dataSynapseWorkspaceAttributes) ConnectivityEndpoints() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sw.ref.Append("connectivity_endpoints"))
}

// Id returns a reference to field id of azurerm_synapse_workspace.
func (sw dataSynapseWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_synapse_workspace.
func (sw dataSynapseWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_synapse_workspace.
func (sw dataSynapseWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_synapse_workspace.
func (sw dataSynapseWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_synapse_workspace.
func (sw dataSynapseWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sw.ref.Append("tags"))
}

func (sw dataSynapseWorkspaceAttributes) Identity() terra.ListValue[datasynapseworkspace.IdentityAttributes] {
	return terra.ReferenceAsList[datasynapseworkspace.IdentityAttributes](sw.ref.Append("identity"))
}

func (sw dataSynapseWorkspaceAttributes) Timeouts() datasynapseworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasynapseworkspace.TimeoutsAttributes](sw.ref.Append("timeouts"))
}
