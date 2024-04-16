// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_kendra_thesaurus

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_kendra_thesaurus.
type Resource struct {
	Name      string
	Args      Args
	state     *awsKendraThesaurusState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akt *Resource) Type() string {
	return "aws_kendra_thesaurus"
}

// LocalName returns the local name for [Resource].
func (akt *Resource) LocalName() string {
	return akt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akt *Resource) Configuration() interface{} {
	return akt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akt *Resource) Dependencies() terra.Dependencies {
	return akt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akt *Resource) LifecycleManagement() *terra.Lifecycle {
	return akt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akt *Resource) Attributes() awsKendraThesaurusAttributes {
	return awsKendraThesaurusAttributes{ref: terra.ReferenceResource(akt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akt *Resource) ImportState(state io.Reader) error {
	akt.state = &awsKendraThesaurusState{}
	if err := json.NewDecoder(state).Decode(akt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akt.Type(), akt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akt *Resource) State() (*awsKendraThesaurusState, bool) {
	return akt.state, akt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akt *Resource) StateMust() *awsKendraThesaurusState {
	if akt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akt.Type(), akt.LocalName()))
	}
	return akt.state
}

// Args contains the configurations for aws_kendra_thesaurus.
type Args struct {
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
	SourceS3Path *SourceS3Path `hcl:"source_s3_path,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsKendraThesaurusAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("arn"))
}

// Description returns a reference to field description of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("description"))
}

// Id returns a reference to field id of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("index_id"))
}

// Name returns a reference to field name of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akt.ref.Append("tags_all"))
}

// ThesaurusId returns a reference to field thesaurus_id of aws_kendra_thesaurus.
func (akt awsKendraThesaurusAttributes) ThesaurusId() terra.StringValue {
	return terra.ReferenceAsString(akt.ref.Append("thesaurus_id"))
}

func (akt awsKendraThesaurusAttributes) SourceS3Path() terra.ListValue[SourceS3PathAttributes] {
	return terra.ReferenceAsList[SourceS3PathAttributes](akt.ref.Append("source_s3_path"))
}

func (akt awsKendraThesaurusAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](akt.ref.Append("timeouts"))
}

type awsKendraThesaurusState struct {
	Arn          string              `json:"arn"`
	Description  string              `json:"description"`
	Id           string              `json:"id"`
	IndexId      string              `json:"index_id"`
	Name         string              `json:"name"`
	RoleArn      string              `json:"role_arn"`
	Status       string              `json:"status"`
	Tags         map[string]string   `json:"tags"`
	TagsAll      map[string]string   `json:"tags_all"`
	ThesaurusId  string              `json:"thesaurus_id"`
	SourceS3Path []SourceS3PathState `json:"source_s3_path"`
	Timeouts     *TimeoutsState      `json:"timeouts"`
}
