// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	spannerinstanceiammember "github.com/golingon/terraproviders/google/4.74.0/spannerinstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpannerInstanceIamMember creates a new instance of [SpannerInstanceIamMember].
func NewSpannerInstanceIamMember(name string, args SpannerInstanceIamMemberArgs) *SpannerInstanceIamMember {
	return &SpannerInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerInstanceIamMember)(nil)

// SpannerInstanceIamMember represents the Terraform resource google_spanner_instance_iam_member.
type SpannerInstanceIamMember struct {
	Name      string
	Args      SpannerInstanceIamMemberArgs
	state     *spannerInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerInstanceIamMember].
func (siim *SpannerInstanceIamMember) Type() string {
	return "google_spanner_instance_iam_member"
}

// LocalName returns the local name for [SpannerInstanceIamMember].
func (siim *SpannerInstanceIamMember) LocalName() string {
	return siim.Name
}

// Configuration returns the configuration (args) for [SpannerInstanceIamMember].
func (siim *SpannerInstanceIamMember) Configuration() interface{} {
	return siim.Args
}

// DependOn is used for other resources to depend on [SpannerInstanceIamMember].
func (siim *SpannerInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(siim)
}

// Dependencies returns the list of resources [SpannerInstanceIamMember] depends_on.
func (siim *SpannerInstanceIamMember) Dependencies() terra.Dependencies {
	return siim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerInstanceIamMember].
func (siim *SpannerInstanceIamMember) LifecycleManagement() *terra.Lifecycle {
	return siim.Lifecycle
}

// Attributes returns the attributes for [SpannerInstanceIamMember].
func (siim *SpannerInstanceIamMember) Attributes() spannerInstanceIamMemberAttributes {
	return spannerInstanceIamMemberAttributes{ref: terra.ReferenceResource(siim)}
}

// ImportState imports the given attribute values into [SpannerInstanceIamMember]'s state.
func (siim *SpannerInstanceIamMember) ImportState(av io.Reader) error {
	siim.state = &spannerInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(siim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", siim.Type(), siim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerInstanceIamMember] has state.
func (siim *SpannerInstanceIamMember) State() (*spannerInstanceIamMemberState, bool) {
	return siim.state, siim.state != nil
}

// StateMust returns the state for [SpannerInstanceIamMember]. Panics if the state is nil.
func (siim *SpannerInstanceIamMember) StateMust() *spannerInstanceIamMemberState {
	if siim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", siim.Type(), siim.LocalName()))
	}
	return siim.state
}

// SpannerInstanceIamMemberArgs contains the configurations for google_spanner_instance_iam_member.
type SpannerInstanceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *spannerinstanceiammember.Condition `hcl:"condition,block"`
}
type spannerInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_spanner_instance_iam_member.
func (siim spannerInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(siim.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_instance_iam_member.
func (siim spannerInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siim.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_instance_iam_member.
func (siim spannerInstanceIamMemberAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(siim.ref.Append("instance"))
}

// Member returns a reference to field member of google_spanner_instance_iam_member.
func (siim spannerInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(siim.ref.Append("member"))
}

// Project returns a reference to field project of google_spanner_instance_iam_member.
func (siim spannerInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(siim.ref.Append("project"))
}

// Role returns a reference to field role of google_spanner_instance_iam_member.
func (siim spannerInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(siim.ref.Append("role"))
}

func (siim spannerInstanceIamMemberAttributes) Condition() terra.ListValue[spannerinstanceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[spannerinstanceiammember.ConditionAttributes](siim.ref.Append("condition"))
}

type spannerInstanceIamMemberState struct {
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Instance  string                                    `json:"instance"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []spannerinstanceiammember.ConditionState `json:"condition"`
}
