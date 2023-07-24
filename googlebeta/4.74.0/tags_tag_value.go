// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	tagstagvalue "github.com/golingon/terraproviders/googlebeta/4.74.0/tagstagvalue"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagValue creates a new instance of [TagsTagValue].
func NewTagsTagValue(name string, args TagsTagValueArgs) *TagsTagValue {
	return &TagsTagValue{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagValue)(nil)

// TagsTagValue represents the Terraform resource google_tags_tag_value.
type TagsTagValue struct {
	Name      string
	Args      TagsTagValueArgs
	state     *tagsTagValueState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagValue].
func (ttv *TagsTagValue) Type() string {
	return "google_tags_tag_value"
}

// LocalName returns the local name for [TagsTagValue].
func (ttv *TagsTagValue) LocalName() string {
	return ttv.Name
}

// Configuration returns the configuration (args) for [TagsTagValue].
func (ttv *TagsTagValue) Configuration() interface{} {
	return ttv.Args
}

// DependOn is used for other resources to depend on [TagsTagValue].
func (ttv *TagsTagValue) DependOn() terra.Reference {
	return terra.ReferenceResource(ttv)
}

// Dependencies returns the list of resources [TagsTagValue] depends_on.
func (ttv *TagsTagValue) Dependencies() terra.Dependencies {
	return ttv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagValue].
func (ttv *TagsTagValue) LifecycleManagement() *terra.Lifecycle {
	return ttv.Lifecycle
}

// Attributes returns the attributes for [TagsTagValue].
func (ttv *TagsTagValue) Attributes() tagsTagValueAttributes {
	return tagsTagValueAttributes{ref: terra.ReferenceResource(ttv)}
}

// ImportState imports the given attribute values into [TagsTagValue]'s state.
func (ttv *TagsTagValue) ImportState(av io.Reader) error {
	ttv.state = &tagsTagValueState{}
	if err := json.NewDecoder(av).Decode(ttv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttv.Type(), ttv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagValue] has state.
func (ttv *TagsTagValue) State() (*tagsTagValueState, bool) {
	return ttv.state, ttv.state != nil
}

// StateMust returns the state for [TagsTagValue]. Panics if the state is nil.
func (ttv *TagsTagValue) StateMust() *tagsTagValueState {
	if ttv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttv.Type(), ttv.LocalName()))
	}
	return ttv.state
}

// TagsTagValueArgs contains the configurations for google_tags_tag_value.
type TagsTagValueArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *tagstagvalue.Timeouts `hcl:"timeouts,block"`
}
type tagsTagValueAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_tags_tag_value.
func (ttv tagsTagValueAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("create_time"))
}

// Description returns a reference to field description of google_tags_tag_value.
func (ttv tagsTagValueAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("description"))
}

// Id returns a reference to field id of google_tags_tag_value.
func (ttv tagsTagValueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("id"))
}

// Name returns a reference to field name of google_tags_tag_value.
func (ttv tagsTagValueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("name"))
}

// NamespacedName returns a reference to field namespaced_name of google_tags_tag_value.
func (ttv tagsTagValueAttributes) NamespacedName() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("namespaced_name"))
}

// Parent returns a reference to field parent of google_tags_tag_value.
func (ttv tagsTagValueAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("parent"))
}

// ShortName returns a reference to field short_name of google_tags_tag_value.
func (ttv tagsTagValueAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("short_name"))
}

// UpdateTime returns a reference to field update_time of google_tags_tag_value.
func (ttv tagsTagValueAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("update_time"))
}

func (ttv tagsTagValueAttributes) Timeouts() tagstagvalue.TimeoutsAttributes {
	return terra.ReferenceAsSingle[tagstagvalue.TimeoutsAttributes](ttv.ref.Append("timeouts"))
}

type tagsTagValueState struct {
	CreateTime     string                      `json:"create_time"`
	Description    string                      `json:"description"`
	Id             string                      `json:"id"`
	Name           string                      `json:"name"`
	NamespacedName string                      `json:"namespaced_name"`
	Parent         string                      `json:"parent"`
	ShortName      string                      `json:"short_name"`
	UpdateTime     string                      `json:"update_time"`
	Timeouts       *tagstagvalue.TimeoutsState `json:"timeouts"`
}
