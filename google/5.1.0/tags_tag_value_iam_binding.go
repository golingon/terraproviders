// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	tagstagvalueiambinding "github.com/golingon/terraproviders/google/5.1.0/tagstagvalueiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagValueIamBinding creates a new instance of [TagsTagValueIamBinding].
func NewTagsTagValueIamBinding(name string, args TagsTagValueIamBindingArgs) *TagsTagValueIamBinding {
	return &TagsTagValueIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagValueIamBinding)(nil)

// TagsTagValueIamBinding represents the Terraform resource google_tags_tag_value_iam_binding.
type TagsTagValueIamBinding struct {
	Name      string
	Args      TagsTagValueIamBindingArgs
	state     *tagsTagValueIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagValueIamBinding].
func (ttvib *TagsTagValueIamBinding) Type() string {
	return "google_tags_tag_value_iam_binding"
}

// LocalName returns the local name for [TagsTagValueIamBinding].
func (ttvib *TagsTagValueIamBinding) LocalName() string {
	return ttvib.Name
}

// Configuration returns the configuration (args) for [TagsTagValueIamBinding].
func (ttvib *TagsTagValueIamBinding) Configuration() interface{} {
	return ttvib.Args
}

// DependOn is used for other resources to depend on [TagsTagValueIamBinding].
func (ttvib *TagsTagValueIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ttvib)
}

// Dependencies returns the list of resources [TagsTagValueIamBinding] depends_on.
func (ttvib *TagsTagValueIamBinding) Dependencies() terra.Dependencies {
	return ttvib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagValueIamBinding].
func (ttvib *TagsTagValueIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ttvib.Lifecycle
}

// Attributes returns the attributes for [TagsTagValueIamBinding].
func (ttvib *TagsTagValueIamBinding) Attributes() tagsTagValueIamBindingAttributes {
	return tagsTagValueIamBindingAttributes{ref: terra.ReferenceResource(ttvib)}
}

// ImportState imports the given attribute values into [TagsTagValueIamBinding]'s state.
func (ttvib *TagsTagValueIamBinding) ImportState(av io.Reader) error {
	ttvib.state = &tagsTagValueIamBindingState{}
	if err := json.NewDecoder(av).Decode(ttvib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttvib.Type(), ttvib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagValueIamBinding] has state.
func (ttvib *TagsTagValueIamBinding) State() (*tagsTagValueIamBindingState, bool) {
	return ttvib.state, ttvib.state != nil
}

// StateMust returns the state for [TagsTagValueIamBinding]. Panics if the state is nil.
func (ttvib *TagsTagValueIamBinding) StateMust() *tagsTagValueIamBindingState {
	if ttvib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttvib.Type(), ttvib.LocalName()))
	}
	return ttvib.state
}

// TagsTagValueIamBindingArgs contains the configurations for google_tags_tag_value_iam_binding.
type TagsTagValueIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TagValue: string, required
	TagValue terra.StringValue `hcl:"tag_value,attr" validate:"required"`
	// Condition: optional
	Condition *tagstagvalueiambinding.Condition `hcl:"condition,block"`
}
type tagsTagValueIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_value_iam_binding.
func (ttvib tagsTagValueIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttvib.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_value_iam_binding.
func (ttvib tagsTagValueIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttvib.ref.Append("id"))
}

// Members returns a reference to field members of google_tags_tag_value_iam_binding.
func (ttvib tagsTagValueIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ttvib.ref.Append("members"))
}

// Role returns a reference to field role of google_tags_tag_value_iam_binding.
func (ttvib tagsTagValueIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ttvib.ref.Append("role"))
}

// TagValue returns a reference to field tag_value of google_tags_tag_value_iam_binding.
func (ttvib tagsTagValueIamBindingAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(ttvib.ref.Append("tag_value"))
}

func (ttvib tagsTagValueIamBindingAttributes) Condition() terra.ListValue[tagstagvalueiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[tagstagvalueiambinding.ConditionAttributes](ttvib.ref.Append("condition"))
}

type tagsTagValueIamBindingState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Members   []string                                `json:"members"`
	Role      string                                  `json:"role"`
	TagValue  string                                  `json:"tag_value"`
	Condition []tagstagvalueiambinding.ConditionState `json:"condition"`
}
