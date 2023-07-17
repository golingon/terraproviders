// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsTagValueIamPolicy creates a new instance of [TagsTagValueIamPolicy].
func NewTagsTagValueIamPolicy(name string, args TagsTagValueIamPolicyArgs) *TagsTagValueIamPolicy {
	return &TagsTagValueIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsTagValueIamPolicy)(nil)

// TagsTagValueIamPolicy represents the Terraform resource google_tags_tag_value_iam_policy.
type TagsTagValueIamPolicy struct {
	Name      string
	Args      TagsTagValueIamPolicyArgs
	state     *tagsTagValueIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsTagValueIamPolicy].
func (ttvip *TagsTagValueIamPolicy) Type() string {
	return "google_tags_tag_value_iam_policy"
}

// LocalName returns the local name for [TagsTagValueIamPolicy].
func (ttvip *TagsTagValueIamPolicy) LocalName() string {
	return ttvip.Name
}

// Configuration returns the configuration (args) for [TagsTagValueIamPolicy].
func (ttvip *TagsTagValueIamPolicy) Configuration() interface{} {
	return ttvip.Args
}

// DependOn is used for other resources to depend on [TagsTagValueIamPolicy].
func (ttvip *TagsTagValueIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ttvip)
}

// Dependencies returns the list of resources [TagsTagValueIamPolicy] depends_on.
func (ttvip *TagsTagValueIamPolicy) Dependencies() terra.Dependencies {
	return ttvip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsTagValueIamPolicy].
func (ttvip *TagsTagValueIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ttvip.Lifecycle
}

// Attributes returns the attributes for [TagsTagValueIamPolicy].
func (ttvip *TagsTagValueIamPolicy) Attributes() tagsTagValueIamPolicyAttributes {
	return tagsTagValueIamPolicyAttributes{ref: terra.ReferenceResource(ttvip)}
}

// ImportState imports the given attribute values into [TagsTagValueIamPolicy]'s state.
func (ttvip *TagsTagValueIamPolicy) ImportState(av io.Reader) error {
	ttvip.state = &tagsTagValueIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ttvip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ttvip.Type(), ttvip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsTagValueIamPolicy] has state.
func (ttvip *TagsTagValueIamPolicy) State() (*tagsTagValueIamPolicyState, bool) {
	return ttvip.state, ttvip.state != nil
}

// StateMust returns the state for [TagsTagValueIamPolicy]. Panics if the state is nil.
func (ttvip *TagsTagValueIamPolicy) StateMust() *tagsTagValueIamPolicyState {
	if ttvip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ttvip.Type(), ttvip.LocalName()))
	}
	return ttvip.state
}

// TagsTagValueIamPolicyArgs contains the configurations for google_tags_tag_value_iam_policy.
type TagsTagValueIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// TagValue: string, required
	TagValue terra.StringValue `hcl:"tag_value,attr" validate:"required"`
}
type tagsTagValueIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_value_iam_policy.
func (ttvip tagsTagValueIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_value_iam_policy.
func (ttvip tagsTagValueIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_tags_tag_value_iam_policy.
func (ttvip tagsTagValueIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("policy_data"))
}

// TagValue returns a reference to field tag_value of google_tags_tag_value_iam_policy.
func (ttvip tagsTagValueIamPolicyAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("tag_value"))
}

type tagsTagValueIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	TagValue   string `json:"tag_value"`
}