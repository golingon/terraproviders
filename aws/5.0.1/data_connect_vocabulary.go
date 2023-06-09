// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataConnectVocabulary creates a new instance of [DataConnectVocabulary].
func NewDataConnectVocabulary(name string, args DataConnectVocabularyArgs) *DataConnectVocabulary {
	return &DataConnectVocabulary{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectVocabulary)(nil)

// DataConnectVocabulary represents the Terraform data resource aws_connect_vocabulary.
type DataConnectVocabulary struct {
	Name string
	Args DataConnectVocabularyArgs
}

// DataSource returns the Terraform object type for [DataConnectVocabulary].
func (cv *DataConnectVocabulary) DataSource() string {
	return "aws_connect_vocabulary"
}

// LocalName returns the local name for [DataConnectVocabulary].
func (cv *DataConnectVocabulary) LocalName() string {
	return cv.Name
}

// Configuration returns the configuration (args) for [DataConnectVocabulary].
func (cv *DataConnectVocabulary) Configuration() interface{} {
	return cv.Args
}

// Attributes returns the attributes for [DataConnectVocabulary].
func (cv *DataConnectVocabulary) Attributes() dataConnectVocabularyAttributes {
	return dataConnectVocabularyAttributes{ref: terra.ReferenceDataResource(cv)}
}

// DataConnectVocabularyArgs contains the configurations for aws_connect_vocabulary.
type DataConnectVocabularyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VocabularyId: string, optional
	VocabularyId terra.StringValue `hcl:"vocabulary_id,attr"`
}
type dataConnectVocabularyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("arn"))
}

// Content returns a reference to field content of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("content"))
}

// FailureReason returns a reference to field failure_reason of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) FailureReason() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("failure_reason"))
}

// Id returns a reference to field id of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("instance_id"))
}

// LanguageCode returns a reference to field language_code of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("language_code"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("last_modified_time"))
}

// Name returns a reference to field name of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("name"))
}

// State returns a reference to field state of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cv.ref.Append("tags"))
}

// VocabularyId returns a reference to field vocabulary_id of aws_connect_vocabulary.
func (cv dataConnectVocabularyAttributes) VocabularyId() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("vocabulary_id"))
}
