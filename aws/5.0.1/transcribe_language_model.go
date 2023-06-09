// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	transcribelanguagemodel "github.com/golingon/terraproviders/aws/5.0.1/transcribelanguagemodel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTranscribeLanguageModel creates a new instance of [TranscribeLanguageModel].
func NewTranscribeLanguageModel(name string, args TranscribeLanguageModelArgs) *TranscribeLanguageModel {
	return &TranscribeLanguageModel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TranscribeLanguageModel)(nil)

// TranscribeLanguageModel represents the Terraform resource aws_transcribe_language_model.
type TranscribeLanguageModel struct {
	Name      string
	Args      TranscribeLanguageModelArgs
	state     *transcribeLanguageModelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TranscribeLanguageModel].
func (tlm *TranscribeLanguageModel) Type() string {
	return "aws_transcribe_language_model"
}

// LocalName returns the local name for [TranscribeLanguageModel].
func (tlm *TranscribeLanguageModel) LocalName() string {
	return tlm.Name
}

// Configuration returns the configuration (args) for [TranscribeLanguageModel].
func (tlm *TranscribeLanguageModel) Configuration() interface{} {
	return tlm.Args
}

// DependOn is used for other resources to depend on [TranscribeLanguageModel].
func (tlm *TranscribeLanguageModel) DependOn() terra.Reference {
	return terra.ReferenceResource(tlm)
}

// Dependencies returns the list of resources [TranscribeLanguageModel] depends_on.
func (tlm *TranscribeLanguageModel) Dependencies() terra.Dependencies {
	return tlm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TranscribeLanguageModel].
func (tlm *TranscribeLanguageModel) LifecycleManagement() *terra.Lifecycle {
	return tlm.Lifecycle
}

// Attributes returns the attributes for [TranscribeLanguageModel].
func (tlm *TranscribeLanguageModel) Attributes() transcribeLanguageModelAttributes {
	return transcribeLanguageModelAttributes{ref: terra.ReferenceResource(tlm)}
}

// ImportState imports the given attribute values into [TranscribeLanguageModel]'s state.
func (tlm *TranscribeLanguageModel) ImportState(av io.Reader) error {
	tlm.state = &transcribeLanguageModelState{}
	if err := json.NewDecoder(av).Decode(tlm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tlm.Type(), tlm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TranscribeLanguageModel] has state.
func (tlm *TranscribeLanguageModel) State() (*transcribeLanguageModelState, bool) {
	return tlm.state, tlm.state != nil
}

// StateMust returns the state for [TranscribeLanguageModel]. Panics if the state is nil.
func (tlm *TranscribeLanguageModel) StateMust() *transcribeLanguageModelState {
	if tlm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tlm.Type(), tlm.LocalName()))
	}
	return tlm.state
}

// TranscribeLanguageModelArgs contains the configurations for aws_transcribe_language_model.
type TranscribeLanguageModelArgs struct {
	// BaseModelName: string, required
	BaseModelName terra.StringValue `hcl:"base_model_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// ModelName: string, required
	ModelName terra.StringValue `hcl:"model_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// InputDataConfig: required
	InputDataConfig *transcribelanguagemodel.InputDataConfig `hcl:"input_data_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *transcribelanguagemodel.Timeouts `hcl:"timeouts,block"`
}
type transcribeLanguageModelAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tlm.ref.Append("arn"))
}

// BaseModelName returns a reference to field base_model_name of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) BaseModelName() terra.StringValue {
	return terra.ReferenceAsString(tlm.ref.Append("base_model_name"))
}

// Id returns a reference to field id of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tlm.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(tlm.ref.Append("language_code"))
}

// ModelName returns a reference to field model_name of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) ModelName() terra.StringValue {
	return terra.ReferenceAsString(tlm.ref.Append("model_name"))
}

// Tags returns a reference to field tags of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tlm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transcribe_language_model.
func (tlm transcribeLanguageModelAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tlm.ref.Append("tags_all"))
}

func (tlm transcribeLanguageModelAttributes) InputDataConfig() terra.ListValue[transcribelanguagemodel.InputDataConfigAttributes] {
	return terra.ReferenceAsList[transcribelanguagemodel.InputDataConfigAttributes](tlm.ref.Append("input_data_config"))
}

func (tlm transcribeLanguageModelAttributes) Timeouts() transcribelanguagemodel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[transcribelanguagemodel.TimeoutsAttributes](tlm.ref.Append("timeouts"))
}

type transcribeLanguageModelState struct {
	Arn             string                                         `json:"arn"`
	BaseModelName   string                                         `json:"base_model_name"`
	Id              string                                         `json:"id"`
	LanguageCode    string                                         `json:"language_code"`
	ModelName       string                                         `json:"model_name"`
	Tags            map[string]string                              `json:"tags"`
	TagsAll         map[string]string                              `json:"tags_all"`
	InputDataConfig []transcribelanguagemodel.InputDataConfigState `json:"input_data_config"`
	Timeouts        *transcribelanguagemodel.TimeoutsState         `json:"timeouts"`
}
