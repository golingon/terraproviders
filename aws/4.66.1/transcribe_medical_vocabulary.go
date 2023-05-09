// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	transcribemedicalvocabulary "github.com/golingon/terraproviders/aws/4.66.1/transcribemedicalvocabulary"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTranscribeMedicalVocabulary creates a new instance of [TranscribeMedicalVocabulary].
func NewTranscribeMedicalVocabulary(name string, args TranscribeMedicalVocabularyArgs) *TranscribeMedicalVocabulary {
	return &TranscribeMedicalVocabulary{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TranscribeMedicalVocabulary)(nil)

// TranscribeMedicalVocabulary represents the Terraform resource aws_transcribe_medical_vocabulary.
type TranscribeMedicalVocabulary struct {
	Name      string
	Args      TranscribeMedicalVocabularyArgs
	state     *transcribeMedicalVocabularyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TranscribeMedicalVocabulary].
func (tmv *TranscribeMedicalVocabulary) Type() string {
	return "aws_transcribe_medical_vocabulary"
}

// LocalName returns the local name for [TranscribeMedicalVocabulary].
func (tmv *TranscribeMedicalVocabulary) LocalName() string {
	return tmv.Name
}

// Configuration returns the configuration (args) for [TranscribeMedicalVocabulary].
func (tmv *TranscribeMedicalVocabulary) Configuration() interface{} {
	return tmv.Args
}

// DependOn is used for other resources to depend on [TranscribeMedicalVocabulary].
func (tmv *TranscribeMedicalVocabulary) DependOn() terra.Reference {
	return terra.ReferenceResource(tmv)
}

// Dependencies returns the list of resources [TranscribeMedicalVocabulary] depends_on.
func (tmv *TranscribeMedicalVocabulary) Dependencies() terra.Dependencies {
	return tmv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TranscribeMedicalVocabulary].
func (tmv *TranscribeMedicalVocabulary) LifecycleManagement() *terra.Lifecycle {
	return tmv.Lifecycle
}

// Attributes returns the attributes for [TranscribeMedicalVocabulary].
func (tmv *TranscribeMedicalVocabulary) Attributes() transcribeMedicalVocabularyAttributes {
	return transcribeMedicalVocabularyAttributes{ref: terra.ReferenceResource(tmv)}
}

// ImportState imports the given attribute values into [TranscribeMedicalVocabulary]'s state.
func (tmv *TranscribeMedicalVocabulary) ImportState(av io.Reader) error {
	tmv.state = &transcribeMedicalVocabularyState{}
	if err := json.NewDecoder(av).Decode(tmv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tmv.Type(), tmv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TranscribeMedicalVocabulary] has state.
func (tmv *TranscribeMedicalVocabulary) State() (*transcribeMedicalVocabularyState, bool) {
	return tmv.state, tmv.state != nil
}

// StateMust returns the state for [TranscribeMedicalVocabulary]. Panics if the state is nil.
func (tmv *TranscribeMedicalVocabulary) StateMust() *transcribeMedicalVocabularyState {
	if tmv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tmv.Type(), tmv.LocalName()))
	}
	return tmv.state
}

// TranscribeMedicalVocabularyArgs contains the configurations for aws_transcribe_medical_vocabulary.
type TranscribeMedicalVocabularyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VocabularyFileUri: string, required
	VocabularyFileUri terra.StringValue `hcl:"vocabulary_file_uri,attr" validate:"required"`
	// VocabularyName: string, required
	VocabularyName terra.StringValue `hcl:"vocabulary_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *transcribemedicalvocabulary.Timeouts `hcl:"timeouts,block"`
}
type transcribeMedicalVocabularyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tmv.ref.Append("arn"))
}

// DownloadUri returns a reference to field download_uri of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) DownloadUri() terra.StringValue {
	return terra.ReferenceAsString(tmv.ref.Append("download_uri"))
}

// Id returns a reference to field id of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmv.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(tmv.ref.Append("language_code"))
}

// Tags returns a reference to field tags of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tmv.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tmv.ref.Append("tags_all"))
}

// VocabularyFileUri returns a reference to field vocabulary_file_uri of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) VocabularyFileUri() terra.StringValue {
	return terra.ReferenceAsString(tmv.ref.Append("vocabulary_file_uri"))
}

// VocabularyName returns a reference to field vocabulary_name of aws_transcribe_medical_vocabulary.
func (tmv transcribeMedicalVocabularyAttributes) VocabularyName() terra.StringValue {
	return terra.ReferenceAsString(tmv.ref.Append("vocabulary_name"))
}

func (tmv transcribeMedicalVocabularyAttributes) Timeouts() transcribemedicalvocabulary.TimeoutsAttributes {
	return terra.ReferenceAsSingle[transcribemedicalvocabulary.TimeoutsAttributes](tmv.ref.Append("timeouts"))
}

type transcribeMedicalVocabularyState struct {
	Arn               string                                     `json:"arn"`
	DownloadUri       string                                     `json:"download_uri"`
	Id                string                                     `json:"id"`
	LanguageCode      string                                     `json:"language_code"`
	Tags              map[string]string                          `json:"tags"`
	TagsAll           map[string]string                          `json:"tags_all"`
	VocabularyFileUri string                                     `json:"vocabulary_file_uri"`
	VocabularyName    string                                     `json:"vocabulary_name"`
	Timeouts          *transcribemedicalvocabulary.TimeoutsState `json:"timeouts"`
}
