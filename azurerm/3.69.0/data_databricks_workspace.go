// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatabricksworkspace "github.com/golingon/terraproviders/azurerm/3.69.0/datadatabricksworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDatabricksWorkspace creates a new instance of [DataDatabricksWorkspace].
func NewDataDatabricksWorkspace(name string, args DataDatabricksWorkspaceArgs) *DataDatabricksWorkspace {
	return &DataDatabricksWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatabricksWorkspace)(nil)

// DataDatabricksWorkspace represents the Terraform data resource azurerm_databricks_workspace.
type DataDatabricksWorkspace struct {
	Name string
	Args DataDatabricksWorkspaceArgs
}

// DataSource returns the Terraform object type for [DataDatabricksWorkspace].
func (dw *DataDatabricksWorkspace) DataSource() string {
	return "azurerm_databricks_workspace"
}

// LocalName returns the local name for [DataDatabricksWorkspace].
func (dw *DataDatabricksWorkspace) LocalName() string {
	return dw.Name
}

// Configuration returns the configuration (args) for [DataDatabricksWorkspace].
func (dw *DataDatabricksWorkspace) Configuration() interface{} {
	return dw.Args
}

// Attributes returns the attributes for [DataDatabricksWorkspace].
func (dw *DataDatabricksWorkspace) Attributes() dataDatabricksWorkspaceAttributes {
	return dataDatabricksWorkspaceAttributes{ref: terra.ReferenceDataResource(dw)}
}

// DataDatabricksWorkspaceArgs contains the configurations for azurerm_databricks_workspace.
type DataDatabricksWorkspaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ManagedDiskIdentity: min=0
	ManagedDiskIdentity []datadatabricksworkspace.ManagedDiskIdentity `hcl:"managed_disk_identity,block" validate:"min=0"`
	// StorageAccountIdentity: min=0
	StorageAccountIdentity []datadatabricksworkspace.StorageAccountIdentity `hcl:"storage_account_identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadatabricksworkspace.Timeouts `hcl:"timeouts,block"`
}
type dataDatabricksWorkspaceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dw.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("workspace_id"))
}

// WorkspaceUrl returns a reference to field workspace_url of azurerm_databricks_workspace.
func (dw dataDatabricksWorkspaceAttributes) WorkspaceUrl() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("workspace_url"))
}

func (dw dataDatabricksWorkspaceAttributes) ManagedDiskIdentity() terra.ListValue[datadatabricksworkspace.ManagedDiskIdentityAttributes] {
	return terra.ReferenceAsList[datadatabricksworkspace.ManagedDiskIdentityAttributes](dw.ref.Append("managed_disk_identity"))
}

func (dw dataDatabricksWorkspaceAttributes) StorageAccountIdentity() terra.ListValue[datadatabricksworkspace.StorageAccountIdentityAttributes] {
	return terra.ReferenceAsList[datadatabricksworkspace.StorageAccountIdentityAttributes](dw.ref.Append("storage_account_identity"))
}

func (dw dataDatabricksWorkspaceAttributes) Timeouts() datadatabricksworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatabricksworkspace.TimeoutsAttributes](dw.ref.Append("timeouts"))
}
