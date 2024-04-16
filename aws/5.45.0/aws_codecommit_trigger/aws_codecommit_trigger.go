// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_codecommit_trigger

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

// Resource represents the Terraform resource aws_codecommit_trigger.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCodecommitTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (act *Resource) Type() string {
	return "aws_codecommit_trigger"
}

// LocalName returns the local name for [Resource].
func (act *Resource) LocalName() string {
	return act.Name
}

// Configuration returns the configuration (args) for [Resource].
func (act *Resource) Configuration() interface{} {
	return act.Args
}

// DependOn is used for other resources to depend on [Resource].
func (act *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(act)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (act *Resource) Dependencies() terra.Dependencies {
	return act.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (act *Resource) LifecycleManagement() *terra.Lifecycle {
	return act.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (act *Resource) Attributes() awsCodecommitTriggerAttributes {
	return awsCodecommitTriggerAttributes{ref: terra.ReferenceResource(act)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (act *Resource) ImportState(state io.Reader) error {
	act.state = &awsCodecommitTriggerState{}
	if err := json.NewDecoder(state).Decode(act.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", act.Type(), act.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (act *Resource) State() (*awsCodecommitTriggerState, bool) {
	return act.state, act.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (act *Resource) StateMust() *awsCodecommitTriggerState {
	if act.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", act.Type(), act.LocalName()))
	}
	return act.state
}

// Args contains the configurations for aws_codecommit_trigger.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
	// Trigger: min=1,max=10
	Trigger []Trigger `hcl:"trigger,block" validate:"min=1,max=10"`
}

type awsCodecommitTriggerAttributes struct {
	ref terra.Reference
}

// ConfigurationId returns a reference to field configuration_id of aws_codecommit_trigger.
func (act awsCodecommitTriggerAttributes) ConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("configuration_id"))
}

// Id returns a reference to field id of aws_codecommit_trigger.
func (act awsCodecommitTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("id"))
}

// RepositoryName returns a reference to field repository_name of aws_codecommit_trigger.
func (act awsCodecommitTriggerAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("repository_name"))
}

func (act awsCodecommitTriggerAttributes) Trigger() terra.SetValue[TriggerAttributes] {
	return terra.ReferenceAsSet[TriggerAttributes](act.ref.Append("trigger"))
}

type awsCodecommitTriggerState struct {
	ConfigurationId string         `json:"configuration_id"`
	Id              string         `json:"id"`
	RepositoryName  string         `json:"repository_name"`
	Trigger         []TriggerState `json:"trigger"`
}
