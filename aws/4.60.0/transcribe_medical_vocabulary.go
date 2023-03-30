// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	transcribemedicalvocabulary "github.com/golingon/terraproviders/aws/4.60.0/transcribemedicalvocabulary"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewTranscribeMedicalVocabulary(name string, args TranscribeMedicalVocabularyArgs) *TranscribeMedicalVocabulary {
	return &TranscribeMedicalVocabulary{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TranscribeMedicalVocabulary)(nil)

type TranscribeMedicalVocabulary struct {
	Name  string
	Args  TranscribeMedicalVocabularyArgs
	state *transcribeMedicalVocabularyState
}

func (tmv *TranscribeMedicalVocabulary) Type() string {
	return "aws_transcribe_medical_vocabulary"
}

func (tmv *TranscribeMedicalVocabulary) LocalName() string {
	return tmv.Name
}

func (tmv *TranscribeMedicalVocabulary) Configuration() interface{} {
	return tmv.Args
}

func (tmv *TranscribeMedicalVocabulary) Attributes() transcribeMedicalVocabularyAttributes {
	return transcribeMedicalVocabularyAttributes{ref: terra.ReferenceResource(tmv)}
}

func (tmv *TranscribeMedicalVocabulary) ImportState(av io.Reader) error {
	tmv.state = &transcribeMedicalVocabularyState{}
	if err := json.NewDecoder(av).Decode(tmv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tmv.Type(), tmv.LocalName(), err)
	}
	return nil
}

func (tmv *TranscribeMedicalVocabulary) State() (*transcribeMedicalVocabularyState, bool) {
	return tmv.state, tmv.state != nil
}

func (tmv *TranscribeMedicalVocabulary) StateMust() *transcribeMedicalVocabularyState {
	if tmv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tmv.Type(), tmv.LocalName()))
	}
	return tmv.state
}

func (tmv *TranscribeMedicalVocabulary) DependOn() terra.Reference {
	return terra.ReferenceResource(tmv)
}

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
	// DependsOn contains resources that TranscribeMedicalVocabulary depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type transcribeMedicalVocabularyAttributes struct {
	ref terra.Reference
}

func (tmv transcribeMedicalVocabularyAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(tmv.ref.Append("arn"))
}

func (tmv transcribeMedicalVocabularyAttributes) DownloadUri() terra.StringValue {
	return terra.ReferenceString(tmv.ref.Append("download_uri"))
}

func (tmv transcribeMedicalVocabularyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(tmv.ref.Append("id"))
}

func (tmv transcribeMedicalVocabularyAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceString(tmv.ref.Append("language_code"))
}

func (tmv transcribeMedicalVocabularyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](tmv.ref.Append("tags"))
}

func (tmv transcribeMedicalVocabularyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](tmv.ref.Append("tags_all"))
}

func (tmv transcribeMedicalVocabularyAttributes) VocabularyFileUri() terra.StringValue {
	return terra.ReferenceString(tmv.ref.Append("vocabulary_file_uri"))
}

func (tmv transcribeMedicalVocabularyAttributes) VocabularyName() terra.StringValue {
	return terra.ReferenceString(tmv.ref.Append("vocabulary_name"))
}

func (tmv transcribeMedicalVocabularyAttributes) Timeouts() transcribemedicalvocabulary.TimeoutsAttributes {
	return terra.ReferenceSingle[transcribemedicalvocabulary.TimeoutsAttributes](tmv.ref.Append("timeouts"))
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
