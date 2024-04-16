// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cosmosdb_notebook_workspace

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_cosmosdb_notebook_workspace.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCosmosdbNotebookWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acnw *Resource) Type() string {
	return "azurerm_cosmosdb_notebook_workspace"
}

// LocalName returns the local name for [Resource].
func (acnw *Resource) LocalName() string {
	return acnw.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acnw *Resource) Configuration() interface{} {
	return acnw.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acnw *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acnw)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acnw *Resource) Dependencies() terra.Dependencies {
	return acnw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acnw *Resource) LifecycleManagement() *terra.Lifecycle {
	return acnw.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acnw *Resource) Attributes() azurermCosmosdbNotebookWorkspaceAttributes {
	return azurermCosmosdbNotebookWorkspaceAttributes{ref: terra.ReferenceResource(acnw)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acnw *Resource) ImportState(state io.Reader) error {
	acnw.state = &azurermCosmosdbNotebookWorkspaceState{}
	if err := json.NewDecoder(state).Decode(acnw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acnw.Type(), acnw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acnw *Resource) State() (*azurermCosmosdbNotebookWorkspaceState, bool) {
	return acnw.state, acnw.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acnw *Resource) StateMust() *azurermCosmosdbNotebookWorkspaceState {
	if acnw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acnw.Type(), acnw.LocalName()))
	}
	return acnw.state
}

// Args contains the configurations for azurerm_cosmosdb_notebook_workspace.
type Args struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermCosmosdbNotebookWorkspaceAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_notebook_workspace.
func (acnw azurermCosmosdbNotebookWorkspaceAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(acnw.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_notebook_workspace.
func (acnw azurermCosmosdbNotebookWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acnw.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_notebook_workspace.
func (acnw azurermCosmosdbNotebookWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acnw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_notebook_workspace.
func (acnw azurermCosmosdbNotebookWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acnw.ref.Append("resource_group_name"))
}

// ServerEndpoint returns a reference to field server_endpoint of azurerm_cosmosdb_notebook_workspace.
func (acnw azurermCosmosdbNotebookWorkspaceAttributes) ServerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(acnw.ref.Append("server_endpoint"))
}

func (acnw azurermCosmosdbNotebookWorkspaceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acnw.ref.Append("timeouts"))
}

type azurermCosmosdbNotebookWorkspaceState struct {
	AccountName       string         `json:"account_name"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	ServerEndpoint    string         `json:"server_endpoint"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
