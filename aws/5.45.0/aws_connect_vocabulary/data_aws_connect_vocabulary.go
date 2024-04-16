// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_vocabulary

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_connect_vocabulary.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acv *DataSource) DataSource() string {
	return "aws_connect_vocabulary"
}

// LocalName returns the local name for [DataSource].
func (acv *DataSource) LocalName() string {
	return acv.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acv *DataSource) Configuration() interface{} {
	return acv.Args
}

// Attributes returns the attributes for [DataSource].
func (acv *DataSource) Attributes() dataAwsConnectVocabularyAttributes {
	return dataAwsConnectVocabularyAttributes{ref: terra.ReferenceDataSource(acv)}
}

// DataArgs contains the configurations for aws_connect_vocabulary.
type DataArgs struct {
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

type dataAwsConnectVocabularyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("arn"))
}

// Content returns a reference to field content of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("content"))
}

// FailureReason returns a reference to field failure_reason of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) FailureReason() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("failure_reason"))
}

// Id returns a reference to field id of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("instance_id"))
}

// LanguageCode returns a reference to field language_code of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("language_code"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("last_modified_time"))
}

// Name returns a reference to field name of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("name"))
}

// State returns a reference to field state of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acv.ref.Append("tags"))
}

// VocabularyId returns a reference to field vocabulary_id of aws_connect_vocabulary.
func (acv dataAwsConnectVocabularyAttributes) VocabularyId() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("vocabulary_id"))
}
