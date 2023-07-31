// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	transcribevocabulary "github.com/golingon/terraproviders/aws/5.10.0/transcribevocabulary"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTranscribeVocabulary creates a new instance of [TranscribeVocabulary].
func NewTranscribeVocabulary(name string, args TranscribeVocabularyArgs) *TranscribeVocabulary {
	return &TranscribeVocabulary{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TranscribeVocabulary)(nil)

// TranscribeVocabulary represents the Terraform resource aws_transcribe_vocabulary.
type TranscribeVocabulary struct {
	Name      string
	Args      TranscribeVocabularyArgs
	state     *transcribeVocabularyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TranscribeVocabulary].
func (tv *TranscribeVocabulary) Type() string {
	return "aws_transcribe_vocabulary"
}

// LocalName returns the local name for [TranscribeVocabulary].
func (tv *TranscribeVocabulary) LocalName() string {
	return tv.Name
}

// Configuration returns the configuration (args) for [TranscribeVocabulary].
func (tv *TranscribeVocabulary) Configuration() interface{} {
	return tv.Args
}

// DependOn is used for other resources to depend on [TranscribeVocabulary].
func (tv *TranscribeVocabulary) DependOn() terra.Reference {
	return terra.ReferenceResource(tv)
}

// Dependencies returns the list of resources [TranscribeVocabulary] depends_on.
func (tv *TranscribeVocabulary) Dependencies() terra.Dependencies {
	return tv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TranscribeVocabulary].
func (tv *TranscribeVocabulary) LifecycleManagement() *terra.Lifecycle {
	return tv.Lifecycle
}

// Attributes returns the attributes for [TranscribeVocabulary].
func (tv *TranscribeVocabulary) Attributes() transcribeVocabularyAttributes {
	return transcribeVocabularyAttributes{ref: terra.ReferenceResource(tv)}
}

// ImportState imports the given attribute values into [TranscribeVocabulary]'s state.
func (tv *TranscribeVocabulary) ImportState(av io.Reader) error {
	tv.state = &transcribeVocabularyState{}
	if err := json.NewDecoder(av).Decode(tv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tv.Type(), tv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TranscribeVocabulary] has state.
func (tv *TranscribeVocabulary) State() (*transcribeVocabularyState, bool) {
	return tv.state, tv.state != nil
}

// StateMust returns the state for [TranscribeVocabulary]. Panics if the state is nil.
func (tv *TranscribeVocabulary) StateMust() *transcribeVocabularyState {
	if tv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tv.Type(), tv.LocalName()))
	}
	return tv.state
}

// TranscribeVocabularyArgs contains the configurations for aws_transcribe_vocabulary.
type TranscribeVocabularyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// Phrases: list of string, optional
	Phrases terra.ListValue[terra.StringValue] `hcl:"phrases,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VocabularyFileUri: string, optional
	VocabularyFileUri terra.StringValue `hcl:"vocabulary_file_uri,attr"`
	// VocabularyName: string, required
	VocabularyName terra.StringValue `hcl:"vocabulary_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *transcribevocabulary.Timeouts `hcl:"timeouts,block"`
}
type transcribeVocabularyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("arn"))
}

// DownloadUri returns a reference to field download_uri of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) DownloadUri() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("download_uri"))
}

// Id returns a reference to field id of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("language_code"))
}

// Phrases returns a reference to field phrases of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) Phrases() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tv.ref.Append("phrases"))
}

// Tags returns a reference to field tags of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tv.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tv.ref.Append("tags_all"))
}

// VocabularyFileUri returns a reference to field vocabulary_file_uri of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) VocabularyFileUri() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("vocabulary_file_uri"))
}

// VocabularyName returns a reference to field vocabulary_name of aws_transcribe_vocabulary.
func (tv transcribeVocabularyAttributes) VocabularyName() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("vocabulary_name"))
}

func (tv transcribeVocabularyAttributes) Timeouts() transcribevocabulary.TimeoutsAttributes {
	return terra.ReferenceAsSingle[transcribevocabulary.TimeoutsAttributes](tv.ref.Append("timeouts"))
}

type transcribeVocabularyState struct {
	Arn               string                              `json:"arn"`
	DownloadUri       string                              `json:"download_uri"`
	Id                string                              `json:"id"`
	LanguageCode      string                              `json:"language_code"`
	Phrases           []string                            `json:"phrases"`
	Tags              map[string]string                   `json:"tags"`
	TagsAll           map[string]string                   `json:"tags_all"`
	VocabularyFileUri string                              `json:"vocabulary_file_uri"`
	VocabularyName    string                              `json:"vocabulary_name"`
	Timeouts          *transcribevocabulary.TimeoutsState `json:"timeouts"`
}
