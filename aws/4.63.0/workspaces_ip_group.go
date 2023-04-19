// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	workspacesipgroup "github.com/golingon/terraproviders/aws/4.63.0/workspacesipgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkspacesIpGroup creates a new instance of [WorkspacesIpGroup].
func NewWorkspacesIpGroup(name string, args WorkspacesIpGroupArgs) *WorkspacesIpGroup {
	return &WorkspacesIpGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkspacesIpGroup)(nil)

// WorkspacesIpGroup represents the Terraform resource aws_workspaces_ip_group.
type WorkspacesIpGroup struct {
	Name      string
	Args      WorkspacesIpGroupArgs
	state     *workspacesIpGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkspacesIpGroup].
func (wig *WorkspacesIpGroup) Type() string {
	return "aws_workspaces_ip_group"
}

// LocalName returns the local name for [WorkspacesIpGroup].
func (wig *WorkspacesIpGroup) LocalName() string {
	return wig.Name
}

// Configuration returns the configuration (args) for [WorkspacesIpGroup].
func (wig *WorkspacesIpGroup) Configuration() interface{} {
	return wig.Args
}

// DependOn is used for other resources to depend on [WorkspacesIpGroup].
func (wig *WorkspacesIpGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(wig)
}

// Dependencies returns the list of resources [WorkspacesIpGroup] depends_on.
func (wig *WorkspacesIpGroup) Dependencies() terra.Dependencies {
	return wig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkspacesIpGroup].
func (wig *WorkspacesIpGroup) LifecycleManagement() *terra.Lifecycle {
	return wig.Lifecycle
}

// Attributes returns the attributes for [WorkspacesIpGroup].
func (wig *WorkspacesIpGroup) Attributes() workspacesIpGroupAttributes {
	return workspacesIpGroupAttributes{ref: terra.ReferenceResource(wig)}
}

// ImportState imports the given attribute values into [WorkspacesIpGroup]'s state.
func (wig *WorkspacesIpGroup) ImportState(av io.Reader) error {
	wig.state = &workspacesIpGroupState{}
	if err := json.NewDecoder(av).Decode(wig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wig.Type(), wig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkspacesIpGroup] has state.
func (wig *WorkspacesIpGroup) State() (*workspacesIpGroupState, bool) {
	return wig.state, wig.state != nil
}

// StateMust returns the state for [WorkspacesIpGroup]. Panics if the state is nil.
func (wig *WorkspacesIpGroup) StateMust() *workspacesIpGroupState {
	if wig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wig.Type(), wig.LocalName()))
	}
	return wig.state
}

// WorkspacesIpGroupArgs contains the configurations for aws_workspaces_ip_group.
type WorkspacesIpGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Rules: min=0
	Rules []workspacesipgroup.Rules `hcl:"rules,block" validate:"min=0"`
}
type workspacesIpGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_workspaces_ip_group.
func (wig workspacesIpGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wig.ref.Append("description"))
}

// Id returns a reference to field id of aws_workspaces_ip_group.
func (wig workspacesIpGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wig.ref.Append("id"))
}

// Name returns a reference to field name of aws_workspaces_ip_group.
func (wig workspacesIpGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wig.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_workspaces_ip_group.
func (wig workspacesIpGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wig.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_workspaces_ip_group.
func (wig workspacesIpGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wig.ref.Append("tags_all"))
}

func (wig workspacesIpGroupAttributes) Rules() terra.SetValue[workspacesipgroup.RulesAttributes] {
	return terra.ReferenceAsSet[workspacesipgroup.RulesAttributes](wig.ref.Append("rules"))
}

type workspacesIpGroupState struct {
	Description string                         `json:"description"`
	Id          string                         `json:"id"`
	Name        string                         `json:"name"`
	Tags        map[string]string              `json:"tags"`
	TagsAll     map[string]string              `json:"tags_all"`
	Rules       []workspacesipgroup.RulesState `json:"rules"`
}