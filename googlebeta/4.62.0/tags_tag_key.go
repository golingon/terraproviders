// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	tagstagkey "github.com/golingon/terraproviders/googlebeta/4.62.0/tagstagkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagKey creates a new instance of [TagsTagKey].
func NewTagsTagKey(name string, args TagsTagKeyArgs) *TagsTagKey {
	return &TagsTagKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagKey)(nil)

// TagsTagKey represents the Terraform resource google_tags_tag_key.
type TagsTagKey struct {
	Name      string
	Args      TagsTagKeyArgs
	state     *tagsTagKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagKey].
func (ttk *TagsTagKey) Type() string {
	return "google_tags_tag_key"
}

// LocalName returns the local name for [TagsTagKey].
func (ttk *TagsTagKey) LocalName() string {
	return ttk.Name
}

// Configuration returns the configuration (args) for [TagsTagKey].
func (ttk *TagsTagKey) Configuration() interface{} {
	return ttk.Args
}

// DependOn is used for other resources to depend on [TagsTagKey].
func (ttk *TagsTagKey) DependOn() terra.Reference {
	return terra.ReferenceResource(ttk)
}

// Dependencies returns the list of resources [TagsTagKey] depends_on.
func (ttk *TagsTagKey) Dependencies() terra.Dependencies {
	return ttk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagKey].
func (ttk *TagsTagKey) LifecycleManagement() *terra.Lifecycle {
	return ttk.Lifecycle
}

// Attributes returns the attributes for [TagsTagKey].
func (ttk *TagsTagKey) Attributes() tagsTagKeyAttributes {
	return tagsTagKeyAttributes{ref: terra.ReferenceResource(ttk)}
}

// ImportState imports the given attribute values into [TagsTagKey]'s state.
func (ttk *TagsTagKey) ImportState(av io.Reader) error {
	ttk.state = &tagsTagKeyState{}
	if err := json.NewDecoder(av).Decode(ttk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttk.Type(), ttk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagKey] has state.
func (ttk *TagsTagKey) State() (*tagsTagKeyState, bool) {
	return ttk.state, ttk.state != nil
}

// StateMust returns the state for [TagsTagKey]. Panics if the state is nil.
func (ttk *TagsTagKey) StateMust() *tagsTagKeyState {
	if ttk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttk.Type(), ttk.LocalName()))
	}
	return ttk.state
}

// TagsTagKeyArgs contains the configurations for google_tags_tag_key.
type TagsTagKeyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Purpose: string, optional
	Purpose terra.StringValue `hcl:"purpose,attr"`
	// PurposeData: map of string, optional
	PurposeData terra.MapValue[terra.StringValue] `hcl:"purpose_data,attr"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *tagstagkey.Timeouts `hcl:"timeouts,block"`
}
type tagsTagKeyAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("create_time"))
}

// Description returns a reference to field description of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("description"))
}

// Id returns a reference to field id of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("id"))
}

// Name returns a reference to field name of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("name"))
}

// NamespacedName returns a reference to field namespaced_name of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) NamespacedName() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("namespaced_name"))
}

// Parent returns a reference to field parent of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("parent"))
}

// Purpose returns a reference to field purpose of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("purpose"))
}

// PurposeData returns a reference to field purpose_data of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) PurposeData() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ttk.ref.Append("purpose_data"))
}

// ShortName returns a reference to field short_name of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("short_name"))
}

// UpdateTime returns a reference to field update_time of google_tags_tag_key.
func (ttk tagsTagKeyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("update_time"))
}

func (ttk tagsTagKeyAttributes) Timeouts() tagstagkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[tagstagkey.TimeoutsAttributes](ttk.ref.Append("timeouts"))
}

type tagsTagKeyState struct {
	CreateTime     string                    `json:"create_time"`
	Description    string                    `json:"description"`
	Id             string                    `json:"id"`
	Name           string                    `json:"name"`
	NamespacedName string                    `json:"namespaced_name"`
	Parent         string                    `json:"parent"`
	Purpose        string                    `json:"purpose"`
	PurposeData    map[string]string         `json:"purpose_data"`
	ShortName      string                    `json:"short_name"`
	UpdateTime     string                    `json:"update_time"`
	Timeouts       *tagstagkey.TimeoutsState `json:"timeouts"`
}
