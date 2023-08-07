// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	tagstagvalueiammember "github.com/golingon/terraproviders/google/4.76.0/tagstagvalueiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagValueIamMember creates a new instance of [TagsTagValueIamMember].
func NewTagsTagValueIamMember(name string, args TagsTagValueIamMemberArgs) *TagsTagValueIamMember {
	return &TagsTagValueIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagValueIamMember)(nil)

// TagsTagValueIamMember represents the Terraform resource google_tags_tag_value_iam_member.
type TagsTagValueIamMember struct {
	Name      string
	Args      TagsTagValueIamMemberArgs
	state     *tagsTagValueIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagValueIamMember].
func (ttvim *TagsTagValueIamMember) Type() string {
	return "google_tags_tag_value_iam_member"
}

// LocalName returns the local name for [TagsTagValueIamMember].
func (ttvim *TagsTagValueIamMember) LocalName() string {
	return ttvim.Name
}

// Configuration returns the configuration (args) for [TagsTagValueIamMember].
func (ttvim *TagsTagValueIamMember) Configuration() interface{} {
	return ttvim.Args
}

// DependOn is used for other resources to depend on [TagsTagValueIamMember].
func (ttvim *TagsTagValueIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ttvim)
}

// Dependencies returns the list of resources [TagsTagValueIamMember] depends_on.
func (ttvim *TagsTagValueIamMember) Dependencies() terra.Dependencies {
	return ttvim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagValueIamMember].
func (ttvim *TagsTagValueIamMember) LifecycleManagement() *terra.Lifecycle {
	return ttvim.Lifecycle
}

// Attributes returns the attributes for [TagsTagValueIamMember].
func (ttvim *TagsTagValueIamMember) Attributes() tagsTagValueIamMemberAttributes {
	return tagsTagValueIamMemberAttributes{ref: terra.ReferenceResource(ttvim)}
}

// ImportState imports the given attribute values into [TagsTagValueIamMember]'s state.
func (ttvim *TagsTagValueIamMember) ImportState(av io.Reader) error {
	ttvim.state = &tagsTagValueIamMemberState{}
	if err := json.NewDecoder(av).Decode(ttvim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttvim.Type(), ttvim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagValueIamMember] has state.
func (ttvim *TagsTagValueIamMember) State() (*tagsTagValueIamMemberState, bool) {
	return ttvim.state, ttvim.state != nil
}

// StateMust returns the state for [TagsTagValueIamMember]. Panics if the state is nil.
func (ttvim *TagsTagValueIamMember) StateMust() *tagsTagValueIamMemberState {
	if ttvim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttvim.Type(), ttvim.LocalName()))
	}
	return ttvim.state
}

// TagsTagValueIamMemberArgs contains the configurations for google_tags_tag_value_iam_member.
type TagsTagValueIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TagValue: string, required
	TagValue terra.StringValue `hcl:"tag_value,attr" validate:"required"`
	// Condition: optional
	Condition *tagstagvalueiammember.Condition `hcl:"condition,block"`
}
type tagsTagValueIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_value_iam_member.
func (ttvim tagsTagValueIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttvim.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_value_iam_member.
func (ttvim tagsTagValueIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttvim.ref.Append("id"))
}

// Member returns a reference to field member of google_tags_tag_value_iam_member.
func (ttvim tagsTagValueIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ttvim.ref.Append("member"))
}

// Role returns a reference to field role of google_tags_tag_value_iam_member.
func (ttvim tagsTagValueIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ttvim.ref.Append("role"))
}

// TagValue returns a reference to field tag_value of google_tags_tag_value_iam_member.
func (ttvim tagsTagValueIamMemberAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(ttvim.ref.Append("tag_value"))
}

func (ttvim tagsTagValueIamMemberAttributes) Condition() terra.ListValue[tagstagvalueiammember.ConditionAttributes] {
	return terra.ReferenceAsList[tagstagvalueiammember.ConditionAttributes](ttvim.ref.Append("condition"))
}

type tagsTagValueIamMemberState struct {
	Etag      string                                 `json:"etag"`
	Id        string                                 `json:"id"`
	Member    string                                 `json:"member"`
	Role      string                                 `json:"role"`
	TagValue  string                                 `json:"tag_value"`
	Condition []tagstagvalueiammember.ConditionState `json:"condition"`
}
