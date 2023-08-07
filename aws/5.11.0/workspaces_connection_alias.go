// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	workspacesconnectionalias "github.com/golingon/terraproviders/aws/5.11.0/workspacesconnectionalias"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkspacesConnectionAlias creates a new instance of [WorkspacesConnectionAlias].
func NewWorkspacesConnectionAlias(name string, args WorkspacesConnectionAliasArgs) *WorkspacesConnectionAlias {
	return &WorkspacesConnectionAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkspacesConnectionAlias)(nil)

// WorkspacesConnectionAlias represents the Terraform resource aws_workspaces_connection_alias.
type WorkspacesConnectionAlias struct {
	Name      string
	Args      WorkspacesConnectionAliasArgs
	state     *workspacesConnectionAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkspacesConnectionAlias].
func (wca *WorkspacesConnectionAlias) Type() string {
	return "aws_workspaces_connection_alias"
}

// LocalName returns the local name for [WorkspacesConnectionAlias].
func (wca *WorkspacesConnectionAlias) LocalName() string {
	return wca.Name
}

// Configuration returns the configuration (args) for [WorkspacesConnectionAlias].
func (wca *WorkspacesConnectionAlias) Configuration() interface{} {
	return wca.Args
}

// DependOn is used for other resources to depend on [WorkspacesConnectionAlias].
func (wca *WorkspacesConnectionAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(wca)
}

// Dependencies returns the list of resources [WorkspacesConnectionAlias] depends_on.
func (wca *WorkspacesConnectionAlias) Dependencies() terra.Dependencies {
	return wca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkspacesConnectionAlias].
func (wca *WorkspacesConnectionAlias) LifecycleManagement() *terra.Lifecycle {
	return wca.Lifecycle
}

// Attributes returns the attributes for [WorkspacesConnectionAlias].
func (wca *WorkspacesConnectionAlias) Attributes() workspacesConnectionAliasAttributes {
	return workspacesConnectionAliasAttributes{ref: terra.ReferenceResource(wca)}
}

// ImportState imports the given attribute values into [WorkspacesConnectionAlias]'s state.
func (wca *WorkspacesConnectionAlias) ImportState(av io.Reader) error {
	wca.state = &workspacesConnectionAliasState{}
	if err := json.NewDecoder(av).Decode(wca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wca.Type(), wca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkspacesConnectionAlias] has state.
func (wca *WorkspacesConnectionAlias) State() (*workspacesConnectionAliasState, bool) {
	return wca.state, wca.state != nil
}

// StateMust returns the state for [WorkspacesConnectionAlias]. Panics if the state is nil.
func (wca *WorkspacesConnectionAlias) StateMust() *workspacesConnectionAliasState {
	if wca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wca.Type(), wca.LocalName()))
	}
	return wca.state
}

// WorkspacesConnectionAliasArgs contains the configurations for aws_workspaces_connection_alias.
type WorkspacesConnectionAliasArgs struct {
	// ConnectionString: string, required
	ConnectionString terra.StringValue `hcl:"connection_string,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *workspacesconnectionalias.Timeouts `hcl:"timeouts,block"`
}
type workspacesConnectionAliasAttributes struct {
	ref terra.Reference
}

// ConnectionString returns a reference to field connection_string of aws_workspaces_connection_alias.
func (wca workspacesConnectionAliasAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(wca.ref.Append("connection_string"))
}

// Id returns a reference to field id of aws_workspaces_connection_alias.
func (wca workspacesConnectionAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wca.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_workspaces_connection_alias.
func (wca workspacesConnectionAliasAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(wca.ref.Append("owner_account_id"))
}

// State returns a reference to field state of aws_workspaces_connection_alias.
func (wca workspacesConnectionAliasAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(wca.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_workspaces_connection_alias.
func (wca workspacesConnectionAliasAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wca.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_workspaces_connection_alias.
func (wca workspacesConnectionAliasAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wca.ref.Append("tags_all"))
}

func (wca workspacesConnectionAliasAttributes) Timeouts() workspacesconnectionalias.TimeoutsAttributes {
	return terra.ReferenceAsSingle[workspacesconnectionalias.TimeoutsAttributes](wca.ref.Append("timeouts"))
}

type workspacesConnectionAliasState struct {
	ConnectionString string                                   `json:"connection_string"`
	Id               string                                   `json:"id"`
	OwnerAccountId   string                                   `json:"owner_account_id"`
	State            string                                   `json:"state"`
	Tags             map[string]string                        `json:"tags"`
	TagsAll          map[string]string                        `json:"tags_all"`
	Timeouts         *workspacesconnectionalias.TimeoutsState `json:"timeouts"`
}
