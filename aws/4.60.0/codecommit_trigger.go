// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codecommittrigger "github.com/golingon/terraproviders/aws/4.60.0/codecommittrigger"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodecommitTrigger creates a new instance of [CodecommitTrigger].
func NewCodecommitTrigger(name string, args CodecommitTriggerArgs) *CodecommitTrigger {
	return &CodecommitTrigger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodecommitTrigger)(nil)

// CodecommitTrigger represents the Terraform resource aws_codecommit_trigger.
type CodecommitTrigger struct {
	Name      string
	Args      CodecommitTriggerArgs
	state     *codecommitTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodecommitTrigger].
func (ct *CodecommitTrigger) Type() string {
	return "aws_codecommit_trigger"
}

// LocalName returns the local name for [CodecommitTrigger].
func (ct *CodecommitTrigger) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [CodecommitTrigger].
func (ct *CodecommitTrigger) Configuration() interface{} {
	return ct.Args
}

// DependOn is used for other resources to depend on [CodecommitTrigger].
func (ct *CodecommitTrigger) DependOn() terra.Reference {
	return terra.ReferenceResource(ct)
}

// Dependencies returns the list of resources [CodecommitTrigger] depends_on.
func (ct *CodecommitTrigger) Dependencies() terra.Dependencies {
	return ct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodecommitTrigger].
func (ct *CodecommitTrigger) LifecycleManagement() *terra.Lifecycle {
	return ct.Lifecycle
}

// Attributes returns the attributes for [CodecommitTrigger].
func (ct *CodecommitTrigger) Attributes() codecommitTriggerAttributes {
	return codecommitTriggerAttributes{ref: terra.ReferenceResource(ct)}
}

// ImportState imports the given attribute values into [CodecommitTrigger]'s state.
func (ct *CodecommitTrigger) ImportState(av io.Reader) error {
	ct.state = &codecommitTriggerState{}
	if err := json.NewDecoder(av).Decode(ct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ct.Type(), ct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodecommitTrigger] has state.
func (ct *CodecommitTrigger) State() (*codecommitTriggerState, bool) {
	return ct.state, ct.state != nil
}

// StateMust returns the state for [CodecommitTrigger]. Panics if the state is nil.
func (ct *CodecommitTrigger) StateMust() *codecommitTriggerState {
	if ct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ct.Type(), ct.LocalName()))
	}
	return ct.state
}

// CodecommitTriggerArgs contains the configurations for aws_codecommit_trigger.
type CodecommitTriggerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
	// Trigger: min=1,max=10
	Trigger []codecommittrigger.Trigger `hcl:"trigger,block" validate:"min=1,max=10"`
}
type codecommitTriggerAttributes struct {
	ref terra.Reference
}

// ConfigurationId returns a reference to field configuration_id of aws_codecommit_trigger.
func (ct codecommitTriggerAttributes) ConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("configuration_id"))
}

// Id returns a reference to field id of aws_codecommit_trigger.
func (ct codecommitTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// RepositoryName returns a reference to field repository_name of aws_codecommit_trigger.
func (ct codecommitTriggerAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("repository_name"))
}

func (ct codecommitTriggerAttributes) Trigger() terra.SetValue[codecommittrigger.TriggerAttributes] {
	return terra.ReferenceAsSet[codecommittrigger.TriggerAttributes](ct.ref.Append("trigger"))
}

type codecommitTriggerState struct {
	ConfigurationId string                           `json:"configuration_id"`
	Id              string                           `json:"id"`
	RepositoryName  string                           `json:"repository_name"`
	Trigger         []codecommittrigger.TriggerState `json:"trigger"`
}
