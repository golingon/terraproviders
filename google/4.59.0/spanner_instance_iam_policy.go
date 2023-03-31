// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSpannerInstanceIamPolicy(name string, args SpannerInstanceIamPolicyArgs) *SpannerInstanceIamPolicy {
	return &SpannerInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerInstanceIamPolicy)(nil)

type SpannerInstanceIamPolicy struct {
	Name  string
	Args  SpannerInstanceIamPolicyArgs
	state *spannerInstanceIamPolicyState
}

func (siip *SpannerInstanceIamPolicy) Type() string {
	return "google_spanner_instance_iam_policy"
}

func (siip *SpannerInstanceIamPolicy) LocalName() string {
	return siip.Name
}

func (siip *SpannerInstanceIamPolicy) Configuration() interface{} {
	return siip.Args
}

func (siip *SpannerInstanceIamPolicy) Attributes() spannerInstanceIamPolicyAttributes {
	return spannerInstanceIamPolicyAttributes{ref: terra.ReferenceResource(siip)}
}

func (siip *SpannerInstanceIamPolicy) ImportState(av io.Reader) error {
	siip.state = &spannerInstanceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(siip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", siip.Type(), siip.LocalName(), err)
	}
	return nil
}

func (siip *SpannerInstanceIamPolicy) State() (*spannerInstanceIamPolicyState, bool) {
	return siip.state, siip.state != nil
}

func (siip *SpannerInstanceIamPolicy) StateMust() *spannerInstanceIamPolicyState {
	if siip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", siip.Type(), siip.LocalName()))
	}
	return siip.state
}

func (siip *SpannerInstanceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(siip)
}

type SpannerInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DependsOn contains resources that SpannerInstanceIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type spannerInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

func (siip spannerInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(siip.ref.Append("etag"))
}

func (siip spannerInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(siip.ref.Append("id"))
}

func (siip spannerInstanceIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceString(siip.ref.Append("instance"))
}

func (siip spannerInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(siip.ref.Append("policy_data"))
}

func (siip spannerInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(siip.ref.Append("project"))
}

type spannerInstanceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Instance   string `json:"instance"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
