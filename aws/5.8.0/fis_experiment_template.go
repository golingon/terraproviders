// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fisexperimenttemplate "github.com/golingon/terraproviders/aws/5.8.0/fisexperimenttemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFisExperimentTemplate creates a new instance of [FisExperimentTemplate].
func NewFisExperimentTemplate(name string, args FisExperimentTemplateArgs) *FisExperimentTemplate {
	return &FisExperimentTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FisExperimentTemplate)(nil)

// FisExperimentTemplate represents the Terraform resource aws_fis_experiment_template.
type FisExperimentTemplate struct {
	Name      string
	Args      FisExperimentTemplateArgs
	state     *fisExperimentTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FisExperimentTemplate].
func (fet *FisExperimentTemplate) Type() string {
	return "aws_fis_experiment_template"
}

// LocalName returns the local name for [FisExperimentTemplate].
func (fet *FisExperimentTemplate) LocalName() string {
	return fet.Name
}

// Configuration returns the configuration (args) for [FisExperimentTemplate].
func (fet *FisExperimentTemplate) Configuration() interface{} {
	return fet.Args
}

// DependOn is used for other resources to depend on [FisExperimentTemplate].
func (fet *FisExperimentTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(fet)
}

// Dependencies returns the list of resources [FisExperimentTemplate] depends_on.
func (fet *FisExperimentTemplate) Dependencies() terra.Dependencies {
	return fet.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FisExperimentTemplate].
func (fet *FisExperimentTemplate) LifecycleManagement() *terra.Lifecycle {
	return fet.Lifecycle
}

// Attributes returns the attributes for [FisExperimentTemplate].
func (fet *FisExperimentTemplate) Attributes() fisExperimentTemplateAttributes {
	return fisExperimentTemplateAttributes{ref: terra.ReferenceResource(fet)}
}

// ImportState imports the given attribute values into [FisExperimentTemplate]'s state.
func (fet *FisExperimentTemplate) ImportState(av io.Reader) error {
	fet.state = &fisExperimentTemplateState{}
	if err := json.NewDecoder(av).Decode(fet.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fet.Type(), fet.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FisExperimentTemplate] has state.
func (fet *FisExperimentTemplate) State() (*fisExperimentTemplateState, bool) {
	return fet.state, fet.state != nil
}

// StateMust returns the state for [FisExperimentTemplate]. Panics if the state is nil.
func (fet *FisExperimentTemplate) StateMust() *fisExperimentTemplateState {
	if fet.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fet.Type(), fet.LocalName()))
	}
	return fet.state
}

// FisExperimentTemplateArgs contains the configurations for aws_fis_experiment_template.
type FisExperimentTemplateArgs struct {
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Action: min=1
	Action []fisexperimenttemplate.Action `hcl:"action,block" validate:"min=1"`
	// LogConfiguration: optional
	LogConfiguration *fisexperimenttemplate.LogConfiguration `hcl:"log_configuration,block"`
	// StopCondition: min=1
	StopCondition []fisexperimenttemplate.StopCondition `hcl:"stop_condition,block" validate:"min=1"`
	// Target: min=0
	Target []fisexperimenttemplate.Target `hcl:"target,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *fisexperimenttemplate.Timeouts `hcl:"timeouts,block"`
}
type fisExperimentTemplateAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_fis_experiment_template.
func (fet fisExperimentTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(fet.ref.Append("description"))
}

// Id returns a reference to field id of aws_fis_experiment_template.
func (fet fisExperimentTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fet.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_fis_experiment_template.
func (fet fisExperimentTemplateAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(fet.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_fis_experiment_template.
func (fet fisExperimentTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fet.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fis_experiment_template.
func (fet fisExperimentTemplateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fet.ref.Append("tags_all"))
}

func (fet fisExperimentTemplateAttributes) Action() terra.SetValue[fisexperimenttemplate.ActionAttributes] {
	return terra.ReferenceAsSet[fisexperimenttemplate.ActionAttributes](fet.ref.Append("action"))
}

func (fet fisExperimentTemplateAttributes) LogConfiguration() terra.ListValue[fisexperimenttemplate.LogConfigurationAttributes] {
	return terra.ReferenceAsList[fisexperimenttemplate.LogConfigurationAttributes](fet.ref.Append("log_configuration"))
}

func (fet fisExperimentTemplateAttributes) StopCondition() terra.SetValue[fisexperimenttemplate.StopConditionAttributes] {
	return terra.ReferenceAsSet[fisexperimenttemplate.StopConditionAttributes](fet.ref.Append("stop_condition"))
}

func (fet fisExperimentTemplateAttributes) Target() terra.SetValue[fisexperimenttemplate.TargetAttributes] {
	return terra.ReferenceAsSet[fisexperimenttemplate.TargetAttributes](fet.ref.Append("target"))
}

func (fet fisExperimentTemplateAttributes) Timeouts() fisexperimenttemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fisexperimenttemplate.TimeoutsAttributes](fet.ref.Append("timeouts"))
}

type fisExperimentTemplateState struct {
	Description      string                                        `json:"description"`
	Id               string                                        `json:"id"`
	RoleArn          string                                        `json:"role_arn"`
	Tags             map[string]string                             `json:"tags"`
	TagsAll          map[string]string                             `json:"tags_all"`
	Action           []fisexperimenttemplate.ActionState           `json:"action"`
	LogConfiguration []fisexperimenttemplate.LogConfigurationState `json:"log_configuration"`
	StopCondition    []fisexperimenttemplate.StopConditionState    `json:"stop_condition"`
	Target           []fisexperimenttemplate.TargetState           `json:"target"`
	Timeouts         *fisexperimenttemplate.TimeoutsState          `json:"timeouts"`
}
