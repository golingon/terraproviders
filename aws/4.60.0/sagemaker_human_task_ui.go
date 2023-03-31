// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerhumantaskui "github.com/golingon/terraproviders/aws/4.60.0/sagemakerhumantaskui"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerHumanTaskUi creates a new instance of [SagemakerHumanTaskUi].
func NewSagemakerHumanTaskUi(name string, args SagemakerHumanTaskUiArgs) *SagemakerHumanTaskUi {
	return &SagemakerHumanTaskUi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerHumanTaskUi)(nil)

// SagemakerHumanTaskUi represents the Terraform resource aws_sagemaker_human_task_ui.
type SagemakerHumanTaskUi struct {
	Name      string
	Args      SagemakerHumanTaskUiArgs
	state     *sagemakerHumanTaskUiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerHumanTaskUi].
func (shtu *SagemakerHumanTaskUi) Type() string {
	return "aws_sagemaker_human_task_ui"
}

// LocalName returns the local name for [SagemakerHumanTaskUi].
func (shtu *SagemakerHumanTaskUi) LocalName() string {
	return shtu.Name
}

// Configuration returns the configuration (args) for [SagemakerHumanTaskUi].
func (shtu *SagemakerHumanTaskUi) Configuration() interface{} {
	return shtu.Args
}

// DependOn is used for other resources to depend on [SagemakerHumanTaskUi].
func (shtu *SagemakerHumanTaskUi) DependOn() terra.Reference {
	return terra.ReferenceResource(shtu)
}

// Dependencies returns the list of resources [SagemakerHumanTaskUi] depends_on.
func (shtu *SagemakerHumanTaskUi) Dependencies() terra.Dependencies {
	return shtu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerHumanTaskUi].
func (shtu *SagemakerHumanTaskUi) LifecycleManagement() *terra.Lifecycle {
	return shtu.Lifecycle
}

// Attributes returns the attributes for [SagemakerHumanTaskUi].
func (shtu *SagemakerHumanTaskUi) Attributes() sagemakerHumanTaskUiAttributes {
	return sagemakerHumanTaskUiAttributes{ref: terra.ReferenceResource(shtu)}
}

// ImportState imports the given attribute values into [SagemakerHumanTaskUi]'s state.
func (shtu *SagemakerHumanTaskUi) ImportState(av io.Reader) error {
	shtu.state = &sagemakerHumanTaskUiState{}
	if err := json.NewDecoder(av).Decode(shtu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", shtu.Type(), shtu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerHumanTaskUi] has state.
func (shtu *SagemakerHumanTaskUi) State() (*sagemakerHumanTaskUiState, bool) {
	return shtu.state, shtu.state != nil
}

// StateMust returns the state for [SagemakerHumanTaskUi]. Panics if the state is nil.
func (shtu *SagemakerHumanTaskUi) StateMust() *sagemakerHumanTaskUiState {
	if shtu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", shtu.Type(), shtu.LocalName()))
	}
	return shtu.state
}

// SagemakerHumanTaskUiArgs contains the configurations for aws_sagemaker_human_task_ui.
type SagemakerHumanTaskUiArgs struct {
	// HumanTaskUiName: string, required
	HumanTaskUiName terra.StringValue `hcl:"human_task_ui_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UiTemplate: required
	UiTemplate *sagemakerhumantaskui.UiTemplate `hcl:"ui_template,block" validate:"required"`
}
type sagemakerHumanTaskUiAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_human_task_ui.
func (shtu sagemakerHumanTaskUiAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(shtu.ref.Append("arn"))
}

// HumanTaskUiName returns a reference to field human_task_ui_name of aws_sagemaker_human_task_ui.
func (shtu sagemakerHumanTaskUiAttributes) HumanTaskUiName() terra.StringValue {
	return terra.ReferenceAsString(shtu.ref.Append("human_task_ui_name"))
}

// Id returns a reference to field id of aws_sagemaker_human_task_ui.
func (shtu sagemakerHumanTaskUiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(shtu.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_sagemaker_human_task_ui.
func (shtu sagemakerHumanTaskUiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](shtu.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_human_task_ui.
func (shtu sagemakerHumanTaskUiAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](shtu.ref.Append("tags_all"))
}

func (shtu sagemakerHumanTaskUiAttributes) UiTemplate() terra.ListValue[sagemakerhumantaskui.UiTemplateAttributes] {
	return terra.ReferenceAsList[sagemakerhumantaskui.UiTemplateAttributes](shtu.ref.Append("ui_template"))
}

type sagemakerHumanTaskUiState struct {
	Arn             string                                 `json:"arn"`
	HumanTaskUiName string                                 `json:"human_task_ui_name"`
	Id              string                                 `json:"id"`
	Tags            map[string]string                      `json:"tags"`
	TagsAll         map[string]string                      `json:"tags_all"`
	UiTemplate      []sagemakerhumantaskui.UiTemplateState `json:"ui_template"`
}
