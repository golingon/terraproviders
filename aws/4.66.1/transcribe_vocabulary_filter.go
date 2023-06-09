// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTranscribeVocabularyFilter creates a new instance of [TranscribeVocabularyFilter].
func NewTranscribeVocabularyFilter(name string, args TranscribeVocabularyFilterArgs) *TranscribeVocabularyFilter {
	return &TranscribeVocabularyFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TranscribeVocabularyFilter)(nil)

// TranscribeVocabularyFilter represents the Terraform resource aws_transcribe_vocabulary_filter.
type TranscribeVocabularyFilter struct {
	Name      string
	Args      TranscribeVocabularyFilterArgs
	state     *transcribeVocabularyFilterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TranscribeVocabularyFilter].
func (tvf *TranscribeVocabularyFilter) Type() string {
	return "aws_transcribe_vocabulary_filter"
}

// LocalName returns the local name for [TranscribeVocabularyFilter].
func (tvf *TranscribeVocabularyFilter) LocalName() string {
	return tvf.Name
}

// Configuration returns the configuration (args) for [TranscribeVocabularyFilter].
func (tvf *TranscribeVocabularyFilter) Configuration() interface{} {
	return tvf.Args
}

// DependOn is used for other resources to depend on [TranscribeVocabularyFilter].
func (tvf *TranscribeVocabularyFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(tvf)
}

// Dependencies returns the list of resources [TranscribeVocabularyFilter] depends_on.
func (tvf *TranscribeVocabularyFilter) Dependencies() terra.Dependencies {
	return tvf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TranscribeVocabularyFilter].
func (tvf *TranscribeVocabularyFilter) LifecycleManagement() *terra.Lifecycle {
	return tvf.Lifecycle
}

// Attributes returns the attributes for [TranscribeVocabularyFilter].
func (tvf *TranscribeVocabularyFilter) Attributes() transcribeVocabularyFilterAttributes {
	return transcribeVocabularyFilterAttributes{ref: terra.ReferenceResource(tvf)}
}

// ImportState imports the given attribute values into [TranscribeVocabularyFilter]'s state.
func (tvf *TranscribeVocabularyFilter) ImportState(av io.Reader) error {
	tvf.state = &transcribeVocabularyFilterState{}
	if err := json.NewDecoder(av).Decode(tvf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tvf.Type(), tvf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TranscribeVocabularyFilter] has state.
func (tvf *TranscribeVocabularyFilter) State() (*transcribeVocabularyFilterState, bool) {
	return tvf.state, tvf.state != nil
}

// StateMust returns the state for [TranscribeVocabularyFilter]. Panics if the state is nil.
func (tvf *TranscribeVocabularyFilter) StateMust() *transcribeVocabularyFilterState {
	if tvf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tvf.Type(), tvf.LocalName()))
	}
	return tvf.state
}

// TranscribeVocabularyFilterArgs contains the configurations for aws_transcribe_vocabulary_filter.
type TranscribeVocabularyFilterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VocabularyFilterFileUri: string, optional
	VocabularyFilterFileUri terra.StringValue `hcl:"vocabulary_filter_file_uri,attr"`
	// VocabularyFilterName: string, required
	VocabularyFilterName terra.StringValue `hcl:"vocabulary_filter_name,attr" validate:"required"`
	// Words: list of string, optional
	Words terra.ListValue[terra.StringValue] `hcl:"words,attr"`
}
type transcribeVocabularyFilterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tvf.ref.Append("arn"))
}

// DownloadUri returns a reference to field download_uri of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) DownloadUri() terra.StringValue {
	return terra.ReferenceAsString(tvf.ref.Append("download_uri"))
}

// Id returns a reference to field id of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tvf.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(tvf.ref.Append("language_code"))
}

// Tags returns a reference to field tags of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tvf.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tvf.ref.Append("tags_all"))
}

// VocabularyFilterFileUri returns a reference to field vocabulary_filter_file_uri of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) VocabularyFilterFileUri() terra.StringValue {
	return terra.ReferenceAsString(tvf.ref.Append("vocabulary_filter_file_uri"))
}

// VocabularyFilterName returns a reference to field vocabulary_filter_name of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) VocabularyFilterName() terra.StringValue {
	return terra.ReferenceAsString(tvf.ref.Append("vocabulary_filter_name"))
}

// Words returns a reference to field words of aws_transcribe_vocabulary_filter.
func (tvf transcribeVocabularyFilterAttributes) Words() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tvf.ref.Append("words"))
}

type transcribeVocabularyFilterState struct {
	Arn                     string            `json:"arn"`
	DownloadUri             string            `json:"download_uri"`
	Id                      string            `json:"id"`
	LanguageCode            string            `json:"language_code"`
	Tags                    map[string]string `json:"tags"`
	TagsAll                 map[string]string `json:"tags_all"`
	VocabularyFilterFileUri string            `json:"vocabulary_filter_file_uri"`
	VocabularyFilterName    string            `json:"vocabulary_filter_name"`
	Words                   []string          `json:"words"`
}
