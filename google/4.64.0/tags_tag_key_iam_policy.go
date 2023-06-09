// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagKeyIamPolicy creates a new instance of [TagsTagKeyIamPolicy].
func NewTagsTagKeyIamPolicy(name string, args TagsTagKeyIamPolicyArgs) *TagsTagKeyIamPolicy {
	return &TagsTagKeyIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagKeyIamPolicy)(nil)

// TagsTagKeyIamPolicy represents the Terraform resource google_tags_tag_key_iam_policy.
type TagsTagKeyIamPolicy struct {
	Name      string
	Args      TagsTagKeyIamPolicyArgs
	state     *tagsTagKeyIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagKeyIamPolicy].
func (ttkip *TagsTagKeyIamPolicy) Type() string {
	return "google_tags_tag_key_iam_policy"
}

// LocalName returns the local name for [TagsTagKeyIamPolicy].
func (ttkip *TagsTagKeyIamPolicy) LocalName() string {
	return ttkip.Name
}

// Configuration returns the configuration (args) for [TagsTagKeyIamPolicy].
func (ttkip *TagsTagKeyIamPolicy) Configuration() interface{} {
	return ttkip.Args
}

// DependOn is used for other resources to depend on [TagsTagKeyIamPolicy].
func (ttkip *TagsTagKeyIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ttkip)
}

// Dependencies returns the list of resources [TagsTagKeyIamPolicy] depends_on.
func (ttkip *TagsTagKeyIamPolicy) Dependencies() terra.Dependencies {
	return ttkip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagKeyIamPolicy].
func (ttkip *TagsTagKeyIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ttkip.Lifecycle
}

// Attributes returns the attributes for [TagsTagKeyIamPolicy].
func (ttkip *TagsTagKeyIamPolicy) Attributes() tagsTagKeyIamPolicyAttributes {
	return tagsTagKeyIamPolicyAttributes{ref: terra.ReferenceResource(ttkip)}
}

// ImportState imports the given attribute values into [TagsTagKeyIamPolicy]'s state.
func (ttkip *TagsTagKeyIamPolicy) ImportState(av io.Reader) error {
	ttkip.state = &tagsTagKeyIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ttkip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttkip.Type(), ttkip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagKeyIamPolicy] has state.
func (ttkip *TagsTagKeyIamPolicy) State() (*tagsTagKeyIamPolicyState, bool) {
	return ttkip.state, ttkip.state != nil
}

// StateMust returns the state for [TagsTagKeyIamPolicy]. Panics if the state is nil.
func (ttkip *TagsTagKeyIamPolicy) StateMust() *tagsTagKeyIamPolicyState {
	if ttkip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttkip.Type(), ttkip.LocalName()))
	}
	return ttkip.state
}

// TagsTagKeyIamPolicyArgs contains the configurations for google_tags_tag_key_iam_policy.
type TagsTagKeyIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// TagKey: string, required
	TagKey terra.StringValue `hcl:"tag_key,attr" validate:"required"`
}
type tagsTagKeyIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_key_iam_policy.
func (ttkip tagsTagKeyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttkip.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_key_iam_policy.
func (ttkip tagsTagKeyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttkip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_tags_tag_key_iam_policy.
func (ttkip tagsTagKeyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ttkip.ref.Append("policy_data"))
}

// TagKey returns a reference to field tag_key of google_tags_tag_key_iam_policy.
func (ttkip tagsTagKeyIamPolicyAttributes) TagKey() terra.StringValue {
	return terra.ReferenceAsString(ttkip.ref.Append("tag_key"))
}

type tagsTagKeyIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	TagKey     string `json:"tag_key"`
}
