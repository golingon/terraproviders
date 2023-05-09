// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	tagstagkeyiambinding "github.com/golingon/terraproviders/google/4.64.0/tagstagkeyiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagKeyIamBinding creates a new instance of [TagsTagKeyIamBinding].
func NewTagsTagKeyIamBinding(name string, args TagsTagKeyIamBindingArgs) *TagsTagKeyIamBinding {
	return &TagsTagKeyIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagKeyIamBinding)(nil)

// TagsTagKeyIamBinding represents the Terraform resource google_tags_tag_key_iam_binding.
type TagsTagKeyIamBinding struct {
	Name      string
	Args      TagsTagKeyIamBindingArgs
	state     *tagsTagKeyIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagKeyIamBinding].
func (ttkib *TagsTagKeyIamBinding) Type() string {
	return "google_tags_tag_key_iam_binding"
}

// LocalName returns the local name for [TagsTagKeyIamBinding].
func (ttkib *TagsTagKeyIamBinding) LocalName() string {
	return ttkib.Name
}

// Configuration returns the configuration (args) for [TagsTagKeyIamBinding].
func (ttkib *TagsTagKeyIamBinding) Configuration() interface{} {
	return ttkib.Args
}

// DependOn is used for other resources to depend on [TagsTagKeyIamBinding].
func (ttkib *TagsTagKeyIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ttkib)
}

// Dependencies returns the list of resources [TagsTagKeyIamBinding] depends_on.
func (ttkib *TagsTagKeyIamBinding) Dependencies() terra.Dependencies {
	return ttkib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagKeyIamBinding].
func (ttkib *TagsTagKeyIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ttkib.Lifecycle
}

// Attributes returns the attributes for [TagsTagKeyIamBinding].
func (ttkib *TagsTagKeyIamBinding) Attributes() tagsTagKeyIamBindingAttributes {
	return tagsTagKeyIamBindingAttributes{ref: terra.ReferenceResource(ttkib)}
}

// ImportState imports the given attribute values into [TagsTagKeyIamBinding]'s state.
func (ttkib *TagsTagKeyIamBinding) ImportState(av io.Reader) error {
	ttkib.state = &tagsTagKeyIamBindingState{}
	if err := json.NewDecoder(av).Decode(ttkib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttkib.Type(), ttkib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagKeyIamBinding] has state.
func (ttkib *TagsTagKeyIamBinding) State() (*tagsTagKeyIamBindingState, bool) {
	return ttkib.state, ttkib.state != nil
}

// StateMust returns the state for [TagsTagKeyIamBinding]. Panics if the state is nil.
func (ttkib *TagsTagKeyIamBinding) StateMust() *tagsTagKeyIamBindingState {
	if ttkib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttkib.Type(), ttkib.LocalName()))
	}
	return ttkib.state
}

// TagsTagKeyIamBindingArgs contains the configurations for google_tags_tag_key_iam_binding.
type TagsTagKeyIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TagKey: string, required
	TagKey terra.StringValue `hcl:"tag_key,attr" validate:"required"`
	// Condition: optional
	Condition *tagstagkeyiambinding.Condition `hcl:"condition,block"`
}
type tagsTagKeyIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_key_iam_binding.
func (ttkib tagsTagKeyIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttkib.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_key_iam_binding.
func (ttkib tagsTagKeyIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttkib.ref.Append("id"))
}

// Members returns a reference to field members of google_tags_tag_key_iam_binding.
func (ttkib tagsTagKeyIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ttkib.ref.Append("members"))
}

// Role returns a reference to field role of google_tags_tag_key_iam_binding.
func (ttkib tagsTagKeyIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ttkib.ref.Append("role"))
}

// TagKey returns a reference to field tag_key of google_tags_tag_key_iam_binding.
func (ttkib tagsTagKeyIamBindingAttributes) TagKey() terra.StringValue {
	return terra.ReferenceAsString(ttkib.ref.Append("tag_key"))
}

func (ttkib tagsTagKeyIamBindingAttributes) Condition() terra.ListValue[tagstagkeyiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[tagstagkeyiambinding.ConditionAttributes](ttkib.ref.Append("condition"))
}

type tagsTagKeyIamBindingState struct {
	Etag      string                                `json:"etag"`
	Id        string                                `json:"id"`
	Members   []string                              `json:"members"`
	Role      string                                `json:"role"`
	TagKey    string                                `json:"tag_key"`
	Condition []tagstagkeyiambinding.ConditionState `json:"condition"`
}
