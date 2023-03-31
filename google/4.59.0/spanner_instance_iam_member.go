// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	spannerinstanceiammember "github.com/golingon/terraproviders/google/4.59.0/spannerinstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSpannerInstanceIamMember(name string, args SpannerInstanceIamMemberArgs) *SpannerInstanceIamMember {
	return &SpannerInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerInstanceIamMember)(nil)

type SpannerInstanceIamMember struct {
	Name  string
	Args  SpannerInstanceIamMemberArgs
	state *spannerInstanceIamMemberState
}

func (siim *SpannerInstanceIamMember) Type() string {
	return "google_spanner_instance_iam_member"
}

func (siim *SpannerInstanceIamMember) LocalName() string {
	return siim.Name
}

func (siim *SpannerInstanceIamMember) Configuration() interface{} {
	return siim.Args
}

func (siim *SpannerInstanceIamMember) Attributes() spannerInstanceIamMemberAttributes {
	return spannerInstanceIamMemberAttributes{ref: terra.ReferenceResource(siim)}
}

func (siim *SpannerInstanceIamMember) ImportState(av io.Reader) error {
	siim.state = &spannerInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(siim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", siim.Type(), siim.LocalName(), err)
	}
	return nil
}

func (siim *SpannerInstanceIamMember) State() (*spannerInstanceIamMemberState, bool) {
	return siim.state, siim.state != nil
}

func (siim *SpannerInstanceIamMember) StateMust() *spannerInstanceIamMemberState {
	if siim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", siim.Type(), siim.LocalName()))
	}
	return siim.state
}

func (siim *SpannerInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(siim)
}

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
	// DependsOn contains resources that SpannerInstanceIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type spannerInstanceIamMemberAttributes struct {
	ref terra.Reference
}

func (siim spannerInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(siim.ref.Append("etag"))
}

func (siim spannerInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(siim.ref.Append("id"))
}

func (siim spannerInstanceIamMemberAttributes) Instance() terra.StringValue {
	return terra.ReferenceString(siim.ref.Append("instance"))
}

func (siim spannerInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(siim.ref.Append("member"))
}

func (siim spannerInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(siim.ref.Append("project"))
}

func (siim spannerInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(siim.ref.Append("role"))
}

func (siim spannerInstanceIamMemberAttributes) Condition() terra.ListValue[spannerinstanceiammember.ConditionAttributes] {
	return terra.ReferenceList[spannerinstanceiammember.ConditionAttributes](siim.ref.Append("condition"))
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
