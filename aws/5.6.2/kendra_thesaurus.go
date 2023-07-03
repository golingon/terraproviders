// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kendrathesaurus "github.com/golingon/terraproviders/aws/5.6.2/kendrathesaurus"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKendraThesaurus creates a new instance of [KendraThesaurus].
func NewKendraThesaurus(name string, args KendraThesaurusArgs) *KendraThesaurus {
	return &KendraThesaurus{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KendraThesaurus)(nil)

// KendraThesaurus represents the Terraform resource aws_kendra_thesaurus.
type KendraThesaurus struct {
	Name      string
	Args      KendraThesaurusArgs
	state     *kendraThesaurusState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KendraThesaurus].
func (kt *KendraThesaurus) Type() string {
	return "aws_kendra_thesaurus"
}

// LocalName returns the local name for [KendraThesaurus].
func (kt *KendraThesaurus) LocalName() string {
	return kt.Name
}

// Configuration returns the configuration (args) for [KendraThesaurus].
func (kt *KendraThesaurus) Configuration() interface{} {
	return kt.Args
}

// DependOn is used for other resources to depend on [KendraThesaurus].
func (kt *KendraThesaurus) DependOn() terra.Reference {
	return terra.ReferenceResource(kt)
}

// Dependencies returns the list of resources [KendraThesaurus] depends_on.
func (kt *KendraThesaurus) Dependencies() terra.Dependencies {
	return kt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KendraThesaurus].
func (kt *KendraThesaurus) LifecycleManagement() *terra.Lifecycle {
	return kt.Lifecycle
}

// Attributes returns the attributes for [KendraThesaurus].
func (kt *KendraThesaurus) Attributes() kendraThesaurusAttributes {
	return kendraThesaurusAttributes{ref: terra.ReferenceResource(kt)}
}

// ImportState imports the given attribute values into [KendraThesaurus]'s state.
func (kt *KendraThesaurus) ImportState(av io.Reader) error {
	kt.state = &kendraThesaurusState{}
	if err := json.NewDecoder(av).Decode(kt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kt.Type(), kt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KendraThesaurus] has state.
func (kt *KendraThesaurus) State() (*kendraThesaurusState, bool) {
	return kt.state, kt.state != nil
}

// StateMust returns the state for [KendraThesaurus]. Panics if the state is nil.
func (kt *KendraThesaurus) StateMust() *kendraThesaurusState {
	if kt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kt.Type(), kt.LocalName()))
	}
	return kt.state
}

// KendraThesaurusArgs contains the configurations for aws_kendra_thesaurus.
type KendraThesaurusArgs struct {
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
	SourceS3Path *kendrathesaurus.SourceS3Path `hcl:"source_s3_path,block" validate:"required"`
	// Timeouts: optional
	Timeouts *kendrathesaurus.Timeouts `hcl:"timeouts,block"`
}
type kendraThesaurusAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("arn"))
}

// Description returns a reference to field description of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("description"))
}

// Id returns a reference to field id of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("index_id"))
}

// Name returns a reference to field name of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kt.ref.Append("tags_all"))
}

// ThesaurusId returns a reference to field thesaurus_id of aws_kendra_thesaurus.
func (kt kendraThesaurusAttributes) ThesaurusId() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("thesaurus_id"))
}

func (kt kendraThesaurusAttributes) SourceS3Path() terra.ListValue[kendrathesaurus.SourceS3PathAttributes] {
	return terra.ReferenceAsList[kendrathesaurus.SourceS3PathAttributes](kt.ref.Append("source_s3_path"))
}

func (kt kendraThesaurusAttributes) Timeouts() kendrathesaurus.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kendrathesaurus.TimeoutsAttributes](kt.ref.Append("timeouts"))
}

type kendraThesaurusState struct {
	Arn          string                              `json:"arn"`
	Description  string                              `json:"description"`
	Id           string                              `json:"id"`
	IndexId      string                              `json:"index_id"`
	Name         string                              `json:"name"`
	RoleArn      string                              `json:"role_arn"`
	Status       string                              `json:"status"`
	Tags         map[string]string                   `json:"tags"`
	TagsAll      map[string]string                   `json:"tags_all"`
	ThesaurusId  string                              `json:"thesaurus_id"`
	SourceS3Path []kendrathesaurus.SourceS3PathState `json:"source_s3_path"`
	Timeouts     *kendrathesaurus.TimeoutsState      `json:"timeouts"`
}
