// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kendraquerysuggestionsblocklist "github.com/golingon/terraproviders/aws/5.10.0/kendraquerysuggestionsblocklist"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKendraQuerySuggestionsBlockList creates a new instance of [KendraQuerySuggestionsBlockList].
func NewKendraQuerySuggestionsBlockList(name string, args KendraQuerySuggestionsBlockListArgs) *KendraQuerySuggestionsBlockList {
	return &KendraQuerySuggestionsBlockList{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KendraQuerySuggestionsBlockList)(nil)

// KendraQuerySuggestionsBlockList represents the Terraform resource aws_kendra_query_suggestions_block_list.
type KendraQuerySuggestionsBlockList struct {
	Name      string
	Args      KendraQuerySuggestionsBlockListArgs
	state     *kendraQuerySuggestionsBlockListState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KendraQuerySuggestionsBlockList].
func (kqsbl *KendraQuerySuggestionsBlockList) Type() string {
	return "aws_kendra_query_suggestions_block_list"
}

// LocalName returns the local name for [KendraQuerySuggestionsBlockList].
func (kqsbl *KendraQuerySuggestionsBlockList) LocalName() string {
	return kqsbl.Name
}

// Configuration returns the configuration (args) for [KendraQuerySuggestionsBlockList].
func (kqsbl *KendraQuerySuggestionsBlockList) Configuration() interface{} {
	return kqsbl.Args
}

// DependOn is used for other resources to depend on [KendraQuerySuggestionsBlockList].
func (kqsbl *KendraQuerySuggestionsBlockList) DependOn() terra.Reference {
	return terra.ReferenceResource(kqsbl)
}

// Dependencies returns the list of resources [KendraQuerySuggestionsBlockList] depends_on.
func (kqsbl *KendraQuerySuggestionsBlockList) Dependencies() terra.Dependencies {
	return kqsbl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KendraQuerySuggestionsBlockList].
func (kqsbl *KendraQuerySuggestionsBlockList) LifecycleManagement() *terra.Lifecycle {
	return kqsbl.Lifecycle
}

// Attributes returns the attributes for [KendraQuerySuggestionsBlockList].
func (kqsbl *KendraQuerySuggestionsBlockList) Attributes() kendraQuerySuggestionsBlockListAttributes {
	return kendraQuerySuggestionsBlockListAttributes{ref: terra.ReferenceResource(kqsbl)}
}

// ImportState imports the given attribute values into [KendraQuerySuggestionsBlockList]'s state.
func (kqsbl *KendraQuerySuggestionsBlockList) ImportState(av io.Reader) error {
	kqsbl.state = &kendraQuerySuggestionsBlockListState{}
	if err := json.NewDecoder(av).Decode(kqsbl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kqsbl.Type(), kqsbl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KendraQuerySuggestionsBlockList] has state.
func (kqsbl *KendraQuerySuggestionsBlockList) State() (*kendraQuerySuggestionsBlockListState, bool) {
	return kqsbl.state, kqsbl.state != nil
}

// StateMust returns the state for [KendraQuerySuggestionsBlockList]. Panics if the state is nil.
func (kqsbl *KendraQuerySuggestionsBlockList) StateMust() *kendraQuerySuggestionsBlockListState {
	if kqsbl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kqsbl.Type(), kqsbl.LocalName()))
	}
	return kqsbl.state
}

// KendraQuerySuggestionsBlockListArgs contains the configurations for aws_kendra_query_suggestions_block_list.
type KendraQuerySuggestionsBlockListArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexId: string, required
	IndexId terra.StringValue `hcl:"index_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// SourceS3Path: required
	SourceS3Path *kendraquerysuggestionsblocklist.SourceS3Path `hcl:"source_s3_path,block" validate:"required"`
	// Timeouts: optional
	Timeouts *kendraquerysuggestionsblocklist.Timeouts `hcl:"timeouts,block"`
}
type kendraQuerySuggestionsBlockListAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("arn"))
}

// Description returns a reference to field description of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("description"))
}

// Id returns a reference to field id of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("index_id"))
}

// Name returns a reference to field name of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("name"))
}

// QuerySuggestionsBlockListId returns a reference to field query_suggestions_block_list_id of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) QuerySuggestionsBlockListId() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("query_suggestions_block_list_id"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(kqsbl.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kqsbl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kendra_query_suggestions_block_list.
func (kqsbl kendraQuerySuggestionsBlockListAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kqsbl.ref.Append("tags_all"))
}

func (kqsbl kendraQuerySuggestionsBlockListAttributes) SourceS3Path() terra.ListValue[kendraquerysuggestionsblocklist.SourceS3PathAttributes] {
	return terra.ReferenceAsList[kendraquerysuggestionsblocklist.SourceS3PathAttributes](kqsbl.ref.Append("source_s3_path"))
}

func (kqsbl kendraQuerySuggestionsBlockListAttributes) Timeouts() kendraquerysuggestionsblocklist.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kendraquerysuggestionsblocklist.TimeoutsAttributes](kqsbl.ref.Append("timeouts"))
}

type kendraQuerySuggestionsBlockListState struct {
	Arn                         string                                              `json:"arn"`
	Description                 string                                              `json:"description"`
	Id                          string                                              `json:"id"`
	IndexId                     string                                              `json:"index_id"`
	Name                        string                                              `json:"name"`
	QuerySuggestionsBlockListId string                                              `json:"query_suggestions_block_list_id"`
	RoleArn                     string                                              `json:"role_arn"`
	Status                      string                                              `json:"status"`
	Tags                        map[string]string                                   `json:"tags"`
	TagsAll                     map[string]string                                   `json:"tags_all"`
	SourceS3Path                []kendraquerysuggestionsblocklist.SourceS3PathState `json:"source_s3_path"`
	Timeouts                    *kendraquerysuggestionsblocklist.TimeoutsState      `json:"timeouts"`
}