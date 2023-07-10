// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectvocabulary "github.com/golingon/terraproviders/aws/5.7.0/connectvocabulary"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectVocabulary creates a new instance of [ConnectVocabulary].
func NewConnectVocabulary(name string, args ConnectVocabularyArgs) *ConnectVocabulary {
	return &ConnectVocabulary{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectVocabulary)(nil)

// ConnectVocabulary represents the Terraform resource aws_connect_vocabulary.
type ConnectVocabulary struct {
	Name      string
	Args      ConnectVocabularyArgs
	state     *connectVocabularyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectVocabulary].
func (cv *ConnectVocabulary) Type() string {
	return "aws_connect_vocabulary"
}

// LocalName returns the local name for [ConnectVocabulary].
func (cv *ConnectVocabulary) LocalName() string {
	return cv.Name
}

// Configuration returns the configuration (args) for [ConnectVocabulary].
func (cv *ConnectVocabulary) Configuration() interface{} {
	return cv.Args
}

// DependOn is used for other resources to depend on [ConnectVocabulary].
func (cv *ConnectVocabulary) DependOn() terra.Reference {
	return terra.ReferenceResource(cv)
}

// Dependencies returns the list of resources [ConnectVocabulary] depends_on.
func (cv *ConnectVocabulary) Dependencies() terra.Dependencies {
	return cv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectVocabulary].
func (cv *ConnectVocabulary) LifecycleManagement() *terra.Lifecycle {
	return cv.Lifecycle
}

// Attributes returns the attributes for [ConnectVocabulary].
func (cv *ConnectVocabulary) Attributes() connectVocabularyAttributes {
	return connectVocabularyAttributes{ref: terra.ReferenceResource(cv)}
}

// ImportState imports the given attribute values into [ConnectVocabulary]'s state.
func (cv *ConnectVocabulary) ImportState(av io.Reader) error {
	cv.state = &connectVocabularyState{}
	if err := json.NewDecoder(av).Decode(cv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cv.Type(), cv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectVocabulary] has state.
func (cv *ConnectVocabulary) State() (*connectVocabularyState, bool) {
	return cv.state, cv.state != nil
}

// StateMust returns the state for [ConnectVocabulary]. Panics if the state is nil.
func (cv *ConnectVocabulary) StateMust() *connectVocabularyState {
	if cv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cv.Type(), cv.LocalName()))
	}
	return cv.state
}

// ConnectVocabularyArgs contains the configurations for aws_connect_vocabulary.
type ConnectVocabularyArgs struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *connectvocabulary.Timeouts `hcl:"timeouts,block"`
}
type connectVocabularyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("arn"))
}

// Content returns a reference to field content of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("content"))
}

// FailureReason returns a reference to field failure_reason of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) FailureReason() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("failure_reason"))
}

// Id returns a reference to field id of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("instance_id"))
}

// LanguageCode returns a reference to field language_code of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("language_code"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("last_modified_time"))
}

// Name returns a reference to field name of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("name"))
}

// State returns a reference to field state of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cv.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cv.ref.Append("tags_all"))
}

// VocabularyId returns a reference to field vocabulary_id of aws_connect_vocabulary.
func (cv connectVocabularyAttributes) VocabularyId() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("vocabulary_id"))
}

func (cv connectVocabularyAttributes) Timeouts() connectvocabulary.TimeoutsAttributes {
	return terra.ReferenceAsSingle[connectvocabulary.TimeoutsAttributes](cv.ref.Append("timeouts"))
}

type connectVocabularyState struct {
	Arn              string                           `json:"arn"`
	Content          string                           `json:"content"`
	FailureReason    string                           `json:"failure_reason"`
	Id               string                           `json:"id"`
	InstanceId       string                           `json:"instance_id"`
	LanguageCode     string                           `json:"language_code"`
	LastModifiedTime string                           `json:"last_modified_time"`
	Name             string                           `json:"name"`
	State            string                           `json:"state"`
	Tags             map[string]string                `json:"tags"`
	TagsAll          map[string]string                `json:"tags_all"`
	VocabularyId     string                           `json:"vocabulary_id"`
	Timeouts         *connectvocabulary.TimeoutsState `json:"timeouts"`
}
