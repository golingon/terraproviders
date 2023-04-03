// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbnotebookworkspace "github.com/golingon/terraproviders/azurerm/3.49.0/cosmosdbnotebookworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbNotebookWorkspace creates a new instance of [CosmosdbNotebookWorkspace].
func NewCosmosdbNotebookWorkspace(name string, args CosmosdbNotebookWorkspaceArgs) *CosmosdbNotebookWorkspace {
	return &CosmosdbNotebookWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbNotebookWorkspace)(nil)

// CosmosdbNotebookWorkspace represents the Terraform resource azurerm_cosmosdb_notebook_workspace.
type CosmosdbNotebookWorkspace struct {
	Name      string
	Args      CosmosdbNotebookWorkspaceArgs
	state     *cosmosdbNotebookWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbNotebookWorkspace].
func (cnw *CosmosdbNotebookWorkspace) Type() string {
	return "azurerm_cosmosdb_notebook_workspace"
}

// LocalName returns the local name for [CosmosdbNotebookWorkspace].
func (cnw *CosmosdbNotebookWorkspace) LocalName() string {
	return cnw.Name
}

// Configuration returns the configuration (args) for [CosmosdbNotebookWorkspace].
func (cnw *CosmosdbNotebookWorkspace) Configuration() interface{} {
	return cnw.Args
}

// DependOn is used for other resources to depend on [CosmosdbNotebookWorkspace].
func (cnw *CosmosdbNotebookWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(cnw)
}

// Dependencies returns the list of resources [CosmosdbNotebookWorkspace] depends_on.
func (cnw *CosmosdbNotebookWorkspace) Dependencies() terra.Dependencies {
	return cnw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbNotebookWorkspace].
func (cnw *CosmosdbNotebookWorkspace) LifecycleManagement() *terra.Lifecycle {
	return cnw.Lifecycle
}

// Attributes returns the attributes for [CosmosdbNotebookWorkspace].
func (cnw *CosmosdbNotebookWorkspace) Attributes() cosmosdbNotebookWorkspaceAttributes {
	return cosmosdbNotebookWorkspaceAttributes{ref: terra.ReferenceResource(cnw)}
}

// ImportState imports the given attribute values into [CosmosdbNotebookWorkspace]'s state.
func (cnw *CosmosdbNotebookWorkspace) ImportState(av io.Reader) error {
	cnw.state = &cosmosdbNotebookWorkspaceState{}
	if err := json.NewDecoder(av).Decode(cnw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnw.Type(), cnw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbNotebookWorkspace] has state.
func (cnw *CosmosdbNotebookWorkspace) State() (*cosmosdbNotebookWorkspaceState, bool) {
	return cnw.state, cnw.state != nil
}

// StateMust returns the state for [CosmosdbNotebookWorkspace]. Panics if the state is nil.
func (cnw *CosmosdbNotebookWorkspace) StateMust() *cosmosdbNotebookWorkspaceState {
	if cnw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnw.Type(), cnw.LocalName()))
	}
	return cnw.state
}

// CosmosdbNotebookWorkspaceArgs contains the configurations for azurerm_cosmosdb_notebook_workspace.
type CosmosdbNotebookWorkspaceArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbnotebookworkspace.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbNotebookWorkspaceAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_notebook_workspace.
func (cnw cosmosdbNotebookWorkspaceAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cnw.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_notebook_workspace.
func (cnw cosmosdbNotebookWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnw.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_notebook_workspace.
func (cnw cosmosdbNotebookWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_notebook_workspace.
func (cnw cosmosdbNotebookWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cnw.ref.Append("resource_group_name"))
}

// ServerEndpoint returns a reference to field server_endpoint of azurerm_cosmosdb_notebook_workspace.
func (cnw cosmosdbNotebookWorkspaceAttributes) ServerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cnw.ref.Append("server_endpoint"))
}

func (cnw cosmosdbNotebookWorkspaceAttributes) Timeouts() cosmosdbnotebookworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbnotebookworkspace.TimeoutsAttributes](cnw.ref.Append("timeouts"))
}

type cosmosdbNotebookWorkspaceState struct {
	AccountName       string                                   `json:"account_name"`
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	ServerEndpoint    string                                   `json:"server_endpoint"`
	Timeouts          *cosmosdbnotebookworkspace.TimeoutsState `json:"timeouts"`
}