// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerflowdefinition "github.com/golingon/terraproviders/aws/4.60.0/sagemakerflowdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerFlowDefinition creates a new instance of [SagemakerFlowDefinition].
func NewSagemakerFlowDefinition(name string, args SagemakerFlowDefinitionArgs) *SagemakerFlowDefinition {
	return &SagemakerFlowDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerFlowDefinition)(nil)

// SagemakerFlowDefinition represents the Terraform resource aws_sagemaker_flow_definition.
type SagemakerFlowDefinition struct {
	Name      string
	Args      SagemakerFlowDefinitionArgs
	state     *sagemakerFlowDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerFlowDefinition].
func (sfd *SagemakerFlowDefinition) Type() string {
	return "aws_sagemaker_flow_definition"
}

// LocalName returns the local name for [SagemakerFlowDefinition].
func (sfd *SagemakerFlowDefinition) LocalName() string {
	return sfd.Name
}

// Configuration returns the configuration (args) for [SagemakerFlowDefinition].
func (sfd *SagemakerFlowDefinition) Configuration() interface{} {
	return sfd.Args
}

// DependOn is used for other resources to depend on [SagemakerFlowDefinition].
func (sfd *SagemakerFlowDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(sfd)
}

// Dependencies returns the list of resources [SagemakerFlowDefinition] depends_on.
func (sfd *SagemakerFlowDefinition) Dependencies() terra.Dependencies {
	return sfd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerFlowDefinition].
func (sfd *SagemakerFlowDefinition) LifecycleManagement() *terra.Lifecycle {
	return sfd.Lifecycle
}

// Attributes returns the attributes for [SagemakerFlowDefinition].
func (sfd *SagemakerFlowDefinition) Attributes() sagemakerFlowDefinitionAttributes {
	return sagemakerFlowDefinitionAttributes{ref: terra.ReferenceResource(sfd)}
}

// ImportState imports the given attribute values into [SagemakerFlowDefinition]'s state.
func (sfd *SagemakerFlowDefinition) ImportState(av io.Reader) error {
	sfd.state = &sagemakerFlowDefinitionState{}
	if err := json.NewDecoder(av).Decode(sfd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfd.Type(), sfd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerFlowDefinition] has state.
func (sfd *SagemakerFlowDefinition) State() (*sagemakerFlowDefinitionState, bool) {
	return sfd.state, sfd.state != nil
}

// StateMust returns the state for [SagemakerFlowDefinition]. Panics if the state is nil.
func (sfd *SagemakerFlowDefinition) StateMust() *sagemakerFlowDefinitionState {
	if sfd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfd.Type(), sfd.LocalName()))
	}
	return sfd.state
}

// SagemakerFlowDefinitionArgs contains the configurations for aws_sagemaker_flow_definition.
type SagemakerFlowDefinitionArgs struct {
	// FlowDefinitionName: string, required
	FlowDefinitionName terra.StringValue `hcl:"flow_definition_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// HumanLoopActivationConfig: optional
	HumanLoopActivationConfig *sagemakerflowdefinition.HumanLoopActivationConfig `hcl:"human_loop_activation_config,block"`
	// HumanLoopConfig: required
	HumanLoopConfig *sagemakerflowdefinition.HumanLoopConfig `hcl:"human_loop_config,block" validate:"required"`
	// HumanLoopRequestSource: optional
	HumanLoopRequestSource *sagemakerflowdefinition.HumanLoopRequestSource `hcl:"human_loop_request_source,block"`
	// OutputConfig: required
	OutputConfig *sagemakerflowdefinition.OutputConfig `hcl:"output_config,block" validate:"required"`
}
type sagemakerFlowDefinitionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_flow_definition.
func (sfd sagemakerFlowDefinitionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sfd.ref.Append("arn"))
}

// FlowDefinitionName returns a reference to field flow_definition_name of aws_sagemaker_flow_definition.
func (sfd sagemakerFlowDefinitionAttributes) FlowDefinitionName() terra.StringValue {
	return terra.ReferenceAsString(sfd.ref.Append("flow_definition_name"))
}

// Id returns a reference to field id of aws_sagemaker_flow_definition.
func (sfd sagemakerFlowDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfd.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_flow_definition.
func (sfd sagemakerFlowDefinitionAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sfd.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_sagemaker_flow_definition.
func (sfd sagemakerFlowDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_flow_definition.
func (sfd sagemakerFlowDefinitionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfd.ref.Append("tags_all"))
}

func (sfd sagemakerFlowDefinitionAttributes) HumanLoopActivationConfig() terra.ListValue[sagemakerflowdefinition.HumanLoopActivationConfigAttributes] {
	return terra.ReferenceAsList[sagemakerflowdefinition.HumanLoopActivationConfigAttributes](sfd.ref.Append("human_loop_activation_config"))
}

func (sfd sagemakerFlowDefinitionAttributes) HumanLoopConfig() terra.ListValue[sagemakerflowdefinition.HumanLoopConfigAttributes] {
	return terra.ReferenceAsList[sagemakerflowdefinition.HumanLoopConfigAttributes](sfd.ref.Append("human_loop_config"))
}

func (sfd sagemakerFlowDefinitionAttributes) HumanLoopRequestSource() terra.ListValue[sagemakerflowdefinition.HumanLoopRequestSourceAttributes] {
	return terra.ReferenceAsList[sagemakerflowdefinition.HumanLoopRequestSourceAttributes](sfd.ref.Append("human_loop_request_source"))
}

func (sfd sagemakerFlowDefinitionAttributes) OutputConfig() terra.ListValue[sagemakerflowdefinition.OutputConfigAttributes] {
	return terra.ReferenceAsList[sagemakerflowdefinition.OutputConfigAttributes](sfd.ref.Append("output_config"))
}

type sagemakerFlowDefinitionState struct {
	Arn                       string                                                   `json:"arn"`
	FlowDefinitionName        string                                                   `json:"flow_definition_name"`
	Id                        string                                                   `json:"id"`
	RoleArn                   string                                                   `json:"role_arn"`
	Tags                      map[string]string                                        `json:"tags"`
	TagsAll                   map[string]string                                        `json:"tags_all"`
	HumanLoopActivationConfig []sagemakerflowdefinition.HumanLoopActivationConfigState `json:"human_loop_activation_config"`
	HumanLoopConfig           []sagemakerflowdefinition.HumanLoopConfigState           `json:"human_loop_config"`
	HumanLoopRequestSource    []sagemakerflowdefinition.HumanLoopRequestSourceState    `json:"human_loop_request_source"`
	OutputConfig              []sagemakerflowdefinition.OutputConfigState              `json:"output_config"`
}