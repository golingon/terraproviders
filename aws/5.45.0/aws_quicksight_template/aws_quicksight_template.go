// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_quicksight_template

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

// Resource represents the Terraform resource aws_quicksight_template.
type Resource struct {
	Name      string
	Args      Args
	state     *awsQuicksightTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aqt *Resource) Type() string {
	return "aws_quicksight_template"
}

// LocalName returns the local name for [Resource].
func (aqt *Resource) LocalName() string {
	return aqt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aqt *Resource) Configuration() interface{} {
	return aqt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aqt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aqt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aqt *Resource) Dependencies() terra.Dependencies {
	return aqt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aqt *Resource) LifecycleManagement() *terra.Lifecycle {
	return aqt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aqt *Resource) Attributes() awsQuicksightTemplateAttributes {
	return awsQuicksightTemplateAttributes{ref: terra.ReferenceResource(aqt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aqt *Resource) ImportState(state io.Reader) error {
	aqt.state = &awsQuicksightTemplateState{}
	if err := json.NewDecoder(state).Decode(aqt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aqt.Type(), aqt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aqt *Resource) State() (*awsQuicksightTemplateState, bool) {
	return aqt.state, aqt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aqt *Resource) StateMust() *awsQuicksightTemplateState {
	if aqt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aqt.Type(), aqt.LocalName()))
	}
	return aqt.state
}

// Args contains the configurations for aws_quicksight_template.
type Args struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TemplateId: string, required
	TemplateId terra.StringValue `hcl:"template_id,attr" validate:"required"`
	// VersionDescription: string, required
	VersionDescription terra.StringValue `hcl:"version_description,attr" validate:"required"`
	// Definition: optional
	Definition *Definition `hcl:"definition,block"`
	// Permissions: min=0,max=64
	Permissions []Permissions `hcl:"permissions,block" validate:"min=0,max=64"`
	// SourceEntity: optional
	SourceEntity *SourceEntity `hcl:"source_entity,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsQuicksightTemplateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("aws_account_id"))
}

// CreatedTime returns a reference to field created_time of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("created_time"))
}

// Id returns a reference to field id of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("id"))
}

// LastUpdatedTime returns a reference to field last_updated_time of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) LastUpdatedTime() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("last_updated_time"))
}

// Name returns a reference to field name of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("name"))
}

// SourceEntityArn returns a reference to field source_entity_arn of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) SourceEntityArn() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("source_entity_arn"))
}

// Status returns a reference to field status of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aqt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aqt.ref.Append("tags_all"))
}

// TemplateId returns a reference to field template_id of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("template_id"))
}

// VersionDescription returns a reference to field version_description of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) VersionDescription() terra.StringValue {
	return terra.ReferenceAsString(aqt.ref.Append("version_description"))
}

// VersionNumber returns a reference to field version_number of aws_quicksight_template.
func (aqt awsQuicksightTemplateAttributes) VersionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(aqt.ref.Append("version_number"))
}

func (aqt awsQuicksightTemplateAttributes) Definition() terra.ListValue[DefinitionAttributes] {
	return terra.ReferenceAsList[DefinitionAttributes](aqt.ref.Append("definition"))
}

func (aqt awsQuicksightTemplateAttributes) Permissions() terra.SetValue[PermissionsAttributes] {
	return terra.ReferenceAsSet[PermissionsAttributes](aqt.ref.Append("permissions"))
}

func (aqt awsQuicksightTemplateAttributes) SourceEntity() terra.ListValue[SourceEntityAttributes] {
	return terra.ReferenceAsList[SourceEntityAttributes](aqt.ref.Append("source_entity"))
}

func (aqt awsQuicksightTemplateAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aqt.ref.Append("timeouts"))
}

type awsQuicksightTemplateState struct {
	Arn                string              `json:"arn"`
	AwsAccountId       string              `json:"aws_account_id"`
	CreatedTime        string              `json:"created_time"`
	Id                 string              `json:"id"`
	LastUpdatedTime    string              `json:"last_updated_time"`
	Name               string              `json:"name"`
	SourceEntityArn    string              `json:"source_entity_arn"`
	Status             string              `json:"status"`
	Tags               map[string]string   `json:"tags"`
	TagsAll            map[string]string   `json:"tags_all"`
	TemplateId         string              `json:"template_id"`
	VersionDescription string              `json:"version_description"`
	VersionNumber      float64             `json:"version_number"`
	Definition         []DefinitionState   `json:"definition"`
	Permissions        []PermissionsState  `json:"permissions"`
	SourceEntity       []SourceEntityState `json:"source_entity"`
	Timeouts           *TimeoutsState      `json:"timeouts"`
}
