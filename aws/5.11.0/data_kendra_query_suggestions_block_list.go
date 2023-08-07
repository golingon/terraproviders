// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakendraquerysuggestionsblocklist "github.com/golingon/terraproviders/aws/5.11.0/datakendraquerysuggestionsblocklist"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKendraQuerySuggestionsBlockList creates a new instance of [DataKendraQuerySuggestionsBlockList].
func NewDataKendraQuerySuggestionsBlockList(name string, args DataKendraQuerySuggestionsBlockListArgs) *DataKendraQuerySuggestionsBlockList {
	return &DataKendraQuerySuggestionsBlockList{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKendraQuerySuggestionsBlockList)(nil)

// DataKendraQuerySuggestionsBlockList represents the Terraform data resource aws_kendra_query_suggestions_block_list.
type DataKendraQuerySuggestionsBlockList struct {
	Name string
	Args DataKendraQuerySuggestionsBlockListArgs
}

// DataSource returns the Terraform object type for [DataKendraQuerySuggestionsBlockList].
func (kqsbl *DataKendraQuerySuggestionsBlockList) DataSource() string {
	return "aws_kendra_query_suggestions_block_list"
}

// LocalName returns the local name for [DataKendraQuerySuggestionsBlockList].
func (kqsbl *DataKendraQuerySuggestionsBlockList) LocalName() string {
	return kqsbl.Name
}

// Configuration returns the configuration (args) for [DataKendraQuerySuggestionsBlockList].
func (kqsbl *DataKendraQuerySuggestionsBlockList) Configuration() interface{} {
	return kqsbl.Args
}

// Attributes returns the attributes for [DataKendraQuerySuggestionsBlockList].
func (kqsbl *DataKendraQuerySuggestionsBlockList) Attributes() dataKendraQuerySuggestionsBlockListAttributes {
	return dataKendraQuerySuggestionsBlockListAttributes{ref: terra.ReferenceDataResource(kqsbl)}
}

// DataKendraQuerySuggestionsBlockListArgs contains the configurations for aws_kendra_query_suggestions_block_list.
type DataKendraQuerySuggestionsBlockListArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexId: string, required
	IndexId terra.StringValue `hcl:"index_id,attr" validate:"required"`
	// QuerySuggestionsBlockListId: string, required
	QuerySuggestionsBlockListId terra.StringValue `hcl:"query_suggestions_block_list_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// SourceS3Path: min=0
	SourceS3Path []datakendraquerysuggestionsblocklist.SourceS3Path `hcl:"source_s3_path,block" validate:"min=0"`
}
type dataKendraQuerySuggestionsBlockListAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("description"))
}

// ErrorMessage returns a reference to field error_message of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("error_message"))
}

// FileSizeBytes returns a reference to field file_size_bytes of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) FileSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(kqsbl.ref.Append("file_size_bytes"))
}

// Id returns a reference to field id of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("index_id"))
}

// ItemCount returns a reference to field item_count of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) ItemCount() terra.NumberValue {
	return terra.ReferenceAsNumber(kqsbl.ref.Append("item_count"))
}

// Name returns a reference to field name of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("name"))
}

// QuerySuggestionsBlockListId returns a reference to field query_suggestions_block_list_id of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) QuerySuggestionsBlockListId() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("query_suggestions_block_list_id"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kqsbl.ref.Append("tags"))
}

// UpdatedAt returns a reference to field updated_at of aws_kendra_query_suggestions_block_list.
func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("updated_at"))
}

func (kqsbl dataKendraQuerySuggestionsBlockListAttributes) SourceS3Path() terra.ListValue[datakendraquerysuggestionsblocklist.SourceS3PathAttributes] {
	return terra.ReferenceAsList[datakendraquerysuggestionsblocklist.SourceS3PathAttributes](kqsbl.ref.Append("source_s3_path"))
}
