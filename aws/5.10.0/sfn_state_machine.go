// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sfnstatemachine "github.com/golingon/terraproviders/aws/5.10.0/sfnstatemachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSfnStateMachine creates a new instance of [SfnStateMachine].
func NewSfnStateMachine(name string, args SfnStateMachineArgs) *SfnStateMachine {
	return &SfnStateMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SfnStateMachine)(nil)

// SfnStateMachine represents the Terraform resource aws_sfn_state_machine.
type SfnStateMachine struct {
	Name      string
	Args      SfnStateMachineArgs
	state     *sfnStateMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SfnStateMachine].
func (ssm *SfnStateMachine) Type() string {
	return "aws_sfn_state_machine"
}

// LocalName returns the local name for [SfnStateMachine].
func (ssm *SfnStateMachine) LocalName() string {
	return ssm.Name
}

// Configuration returns the configuration (args) for [SfnStateMachine].
func (ssm *SfnStateMachine) Configuration() interface{} {
	return ssm.Args
}

// DependOn is used for other resources to depend on [SfnStateMachine].
func (ssm *SfnStateMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(ssm)
}

// Dependencies returns the list of resources [SfnStateMachine] depends_on.
func (ssm *SfnStateMachine) Dependencies() terra.Dependencies {
	return ssm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SfnStateMachine].
func (ssm *SfnStateMachine) LifecycleManagement() *terra.Lifecycle {
	return ssm.Lifecycle
}

// Attributes returns the attributes for [SfnStateMachine].
func (ssm *SfnStateMachine) Attributes() sfnStateMachineAttributes {
	return sfnStateMachineAttributes{ref: terra.ReferenceResource(ssm)}
}

// ImportState imports the given attribute values into [SfnStateMachine]'s state.
func (ssm *SfnStateMachine) ImportState(av io.Reader) error {
	ssm.state = &sfnStateMachineState{}
	if err := json.NewDecoder(av).Decode(ssm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssm.Type(), ssm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SfnStateMachine] has state.
func (ssm *SfnStateMachine) State() (*sfnStateMachineState, bool) {
	return ssm.state, ssm.state != nil
}

// StateMust returns the state for [SfnStateMachine]. Panics if the state is nil.
func (ssm *SfnStateMachine) StateMust() *sfnStateMachineState {
	if ssm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssm.Type(), ssm.LocalName()))
	}
	return ssm.state
}

// SfnStateMachineArgs contains the configurations for aws_sfn_state_machine.
type SfnStateMachineArgs struct {
	// Definition: string, required
	Definition terra.StringValue `hcl:"definition,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Publish: bool, optional
	Publish terra.BoolValue `hcl:"publish,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// LoggingConfiguration: optional
	LoggingConfiguration *sfnstatemachine.LoggingConfiguration `hcl:"logging_configuration,block"`
	// Timeouts: optional
	Timeouts *sfnstatemachine.Timeouts `hcl:"timeouts,block"`
	// TracingConfiguration: optional
	TracingConfiguration *sfnstatemachine.TracingConfiguration `hcl:"tracing_configuration,block"`
}
type sfnStateMachineAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("creation_date"))
}

// Definition returns a reference to field definition of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Definition() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("definition"))
}

// Description returns a reference to field description of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("description"))
}

// Id returns a reference to field id of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("id"))
}

// Name returns a reference to field name of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("name_prefix"))
}

// Publish returns a reference to field publish of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Publish() terra.BoolValue {
	return terra.ReferenceAsBool(ssm.ref.Append("publish"))
}

// RevisionId returns a reference to field revision_id of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) RevisionId() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("revision_id"))
}

// RoleArn returns a reference to field role_arn of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("role_arn"))
}

// StateMachineVersionArn returns a reference to field state_machine_version_arn of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) StateMachineVersionArn() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("state_machine_version_arn"))
}

// Status returns a reference to field status of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssm.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("type"))
}

// VersionDescription returns a reference to field version_description of aws_sfn_state_machine.
func (ssm sfnStateMachineAttributes) VersionDescription() terra.StringValue {
	return terra.ReferenceAsString(ssm.ref.Append("version_description"))
}

func (ssm sfnStateMachineAttributes) LoggingConfiguration() terra.ListValue[sfnstatemachine.LoggingConfigurationAttributes] {
	return terra.ReferenceAsList[sfnstatemachine.LoggingConfigurationAttributes](ssm.ref.Append("logging_configuration"))
}

func (ssm sfnStateMachineAttributes) Timeouts() sfnstatemachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sfnstatemachine.TimeoutsAttributes](ssm.ref.Append("timeouts"))
}

func (ssm sfnStateMachineAttributes) TracingConfiguration() terra.ListValue[sfnstatemachine.TracingConfigurationAttributes] {
	return terra.ReferenceAsList[sfnstatemachine.TracingConfigurationAttributes](ssm.ref.Append("tracing_configuration"))
}

type sfnStateMachineState struct {
	Arn                    string                                      `json:"arn"`
	CreationDate           string                                      `json:"creation_date"`
	Definition             string                                      `json:"definition"`
	Description            string                                      `json:"description"`
	Id                     string                                      `json:"id"`
	Name                   string                                      `json:"name"`
	NamePrefix             string                                      `json:"name_prefix"`
	Publish                bool                                        `json:"publish"`
	RevisionId             string                                      `json:"revision_id"`
	RoleArn                string                                      `json:"role_arn"`
	StateMachineVersionArn string                                      `json:"state_machine_version_arn"`
	Status                 string                                      `json:"status"`
	Tags                   map[string]string                           `json:"tags"`
	TagsAll                map[string]string                           `json:"tags_all"`
	Type                   string                                      `json:"type"`
	VersionDescription     string                                      `json:"version_description"`
	LoggingConfiguration   []sfnstatemachine.LoggingConfigurationState `json:"logging_configuration"`
	Timeouts               *sfnstatemachine.TimeoutsState              `json:"timeouts"`
	TracingConfiguration   []sfnstatemachine.TracingConfigurationState `json:"tracing_configuration"`
}
