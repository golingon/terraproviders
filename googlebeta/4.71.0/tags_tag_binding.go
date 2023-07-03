// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	tagstagbinding "github.com/golingon/terraproviders/googlebeta/4.71.0/tagstagbinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagBinding creates a new instance of [TagsTagBinding].
func NewTagsTagBinding(name string, args TagsTagBindingArgs) *TagsTagBinding {
	return &TagsTagBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagBinding)(nil)

// TagsTagBinding represents the Terraform resource google_tags_tag_binding.
type TagsTagBinding struct {
	Name      string
	Args      TagsTagBindingArgs
	state     *tagsTagBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagBinding].
func (ttb *TagsTagBinding) Type() string {
	return "google_tags_tag_binding"
}

// LocalName returns the local name for [TagsTagBinding].
func (ttb *TagsTagBinding) LocalName() string {
	return ttb.Name
}

// Configuration returns the configuration (args) for [TagsTagBinding].
func (ttb *TagsTagBinding) Configuration() interface{} {
	return ttb.Args
}

// DependOn is used for other resources to depend on [TagsTagBinding].
func (ttb *TagsTagBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ttb)
}

// Dependencies returns the list of resources [TagsTagBinding] depends_on.
func (ttb *TagsTagBinding) Dependencies() terra.Dependencies {
	return ttb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagBinding].
func (ttb *TagsTagBinding) LifecycleManagement() *terra.Lifecycle {
	return ttb.Lifecycle
}

// Attributes returns the attributes for [TagsTagBinding].
func (ttb *TagsTagBinding) Attributes() tagsTagBindingAttributes {
	return tagsTagBindingAttributes{ref: terra.ReferenceResource(ttb)}
}

// ImportState imports the given attribute values into [TagsTagBinding]'s state.
func (ttb *TagsTagBinding) ImportState(av io.Reader) error {
	ttb.state = &tagsTagBindingState{}
	if err := json.NewDecoder(av).Decode(ttb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttb.Type(), ttb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagBinding] has state.
func (ttb *TagsTagBinding) State() (*tagsTagBindingState, bool) {
	return ttb.state, ttb.state != nil
}

// StateMust returns the state for [TagsTagBinding]. Panics if the state is nil.
func (ttb *TagsTagBinding) StateMust() *tagsTagBindingState {
	if ttb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttb.Type(), ttb.LocalName()))
	}
	return ttb.state
}

// TagsTagBindingArgs contains the configurations for google_tags_tag_binding.
type TagsTagBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// TagValue: string, required
	TagValue terra.StringValue `hcl:"tag_value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *tagstagbinding.Timeouts `hcl:"timeouts,block"`
}
type tagsTagBindingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_tags_tag_binding.
func (ttb tagsTagBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttb.ref.Append("id"))
}

// Name returns a reference to field name of google_tags_tag_binding.
func (ttb tagsTagBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ttb.ref.Append("name"))
}

// Parent returns a reference to field parent of google_tags_tag_binding.
func (ttb tagsTagBindingAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(ttb.ref.Append("parent"))
}

// TagValue returns a reference to field tag_value of google_tags_tag_binding.
func (ttb tagsTagBindingAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(ttb.ref.Append("tag_value"))
}

func (ttb tagsTagBindingAttributes) Timeouts() tagstagbinding.TimeoutsAttributes {
	return terra.ReferenceAsSingle[tagstagbinding.TimeoutsAttributes](ttb.ref.Append("timeouts"))
}

type tagsTagBindingState struct {
	Id       string                        `json:"id"`
	Name     string                        `json:"name"`
	Parent   string                        `json:"parent"`
	TagValue string                        `json:"tag_value"`
	Timeouts *tagstagbinding.TimeoutsState `json:"timeouts"`
}
