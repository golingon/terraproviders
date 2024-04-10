// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	tagstagkeyiammember "github.com/golingon/terraproviders/google/5.24.0/tagstagkeyiammember"
	"io"
)

// NewTagsTagKeyIamMember creates a new instance of [TagsTagKeyIamMember].
func NewTagsTagKeyIamMember(name string, args TagsTagKeyIamMemberArgs) *TagsTagKeyIamMember {
	return &TagsTagKeyIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagKeyIamMember)(nil)

// TagsTagKeyIamMember represents the Terraform resource google_tags_tag_key_iam_member.
type TagsTagKeyIamMember struct {
	Name      string
	Args      TagsTagKeyIamMemberArgs
	state     *tagsTagKeyIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagKeyIamMember].
func (ttkim *TagsTagKeyIamMember) Type() string {
	return "google_tags_tag_key_iam_member"
}

// LocalName returns the local name for [TagsTagKeyIamMember].
func (ttkim *TagsTagKeyIamMember) LocalName() string {
	return ttkim.Name
}

// Configuration returns the configuration (args) for [TagsTagKeyIamMember].
func (ttkim *TagsTagKeyIamMember) Configuration() interface{} {
	return ttkim.Args
}

// DependOn is used for other resources to depend on [TagsTagKeyIamMember].
func (ttkim *TagsTagKeyIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ttkim)
}

// Dependencies returns the list of resources [TagsTagKeyIamMember] depends_on.
func (ttkim *TagsTagKeyIamMember) Dependencies() terra.Dependencies {
	return ttkim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagKeyIamMember].
func (ttkim *TagsTagKeyIamMember) LifecycleManagement() *terra.Lifecycle {
	return ttkim.Lifecycle
}

// Attributes returns the attributes for [TagsTagKeyIamMember].
func (ttkim *TagsTagKeyIamMember) Attributes() tagsTagKeyIamMemberAttributes {
	return tagsTagKeyIamMemberAttributes{ref: terra.ReferenceResource(ttkim)}
}

// ImportState imports the given attribute values into [TagsTagKeyIamMember]'s state.
func (ttkim *TagsTagKeyIamMember) ImportState(av io.Reader) error {
	ttkim.state = &tagsTagKeyIamMemberState{}
	if err := json.NewDecoder(av).Decode(ttkim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttkim.Type(), ttkim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagKeyIamMember] has state.
func (ttkim *TagsTagKeyIamMember) State() (*tagsTagKeyIamMemberState, bool) {
	return ttkim.state, ttkim.state != nil
}

// StateMust returns the state for [TagsTagKeyIamMember]. Panics if the state is nil.
func (ttkim *TagsTagKeyIamMember) StateMust() *tagsTagKeyIamMemberState {
	if ttkim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttkim.Type(), ttkim.LocalName()))
	}
	return ttkim.state
}

// TagsTagKeyIamMemberArgs contains the configurations for google_tags_tag_key_iam_member.
type TagsTagKeyIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TagKey: string, required
	TagKey terra.StringValue `hcl:"tag_key,attr" validate:"required"`
	// Condition: optional
	Condition *tagstagkeyiammember.Condition `hcl:"condition,block"`
}
type tagsTagKeyIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_key_iam_member.
func (ttkim tagsTagKeyIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttkim.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_key_iam_member.
func (ttkim tagsTagKeyIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttkim.ref.Append("id"))
}

// Member returns a reference to field member of google_tags_tag_key_iam_member.
func (ttkim tagsTagKeyIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ttkim.ref.Append("member"))
}

// Role returns a reference to field role of google_tags_tag_key_iam_member.
func (ttkim tagsTagKeyIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ttkim.ref.Append("role"))
}

// TagKey returns a reference to field tag_key of google_tags_tag_key_iam_member.
func (ttkim tagsTagKeyIamMemberAttributes) TagKey() terra.StringValue {
	return terra.ReferenceAsString(ttkim.ref.Append("tag_key"))
}

func (ttkim tagsTagKeyIamMemberAttributes) Condition() terra.ListValue[tagstagkeyiammember.ConditionAttributes] {
	return terra.ReferenceAsList[tagstagkeyiammember.ConditionAttributes](ttkim.ref.Append("condition"))
}

type tagsTagKeyIamMemberState struct {
	Etag      string                               `json:"etag"`
	Id        string                               `json:"id"`
	Member    string                               `json:"member"`
	Role      string                               `json:"role"`
	TagKey    string                               `json:"tag_key"`
	Condition []tagstagkeyiammember.ConditionState `json:"condition"`
}
