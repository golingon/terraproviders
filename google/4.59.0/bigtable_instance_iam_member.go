// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigtableinstanceiammember "github.com/golingon/terraproviders/google/4.59.0/bigtableinstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBigtableInstanceIamMember(name string, args BigtableInstanceIamMemberArgs) *BigtableInstanceIamMember {
	return &BigtableInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableInstanceIamMember)(nil)

type BigtableInstanceIamMember struct {
	Name  string
	Args  BigtableInstanceIamMemberArgs
	state *bigtableInstanceIamMemberState
}

func (biim *BigtableInstanceIamMember) Type() string {
	return "google_bigtable_instance_iam_member"
}

func (biim *BigtableInstanceIamMember) LocalName() string {
	return biim.Name
}

func (biim *BigtableInstanceIamMember) Configuration() interface{} {
	return biim.Args
}

func (biim *BigtableInstanceIamMember) Attributes() bigtableInstanceIamMemberAttributes {
	return bigtableInstanceIamMemberAttributes{ref: terra.ReferenceResource(biim)}
}

func (biim *BigtableInstanceIamMember) ImportState(av io.Reader) error {
	biim.state = &bigtableInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(biim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", biim.Type(), biim.LocalName(), err)
	}
	return nil
}

func (biim *BigtableInstanceIamMember) State() (*bigtableInstanceIamMemberState, bool) {
	return biim.state, biim.state != nil
}

func (biim *BigtableInstanceIamMember) StateMust() *bigtableInstanceIamMemberState {
	if biim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", biim.Type(), biim.LocalName()))
	}
	return biim.state
}

func (biim *BigtableInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(biim)
}

type BigtableInstanceIamMemberArgs struct {
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
	Condition *bigtableinstanceiammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that BigtableInstanceIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type bigtableInstanceIamMemberAttributes struct {
	ref terra.Reference
}

func (biim bigtableInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(biim.ref.Append("etag"))
}

func (biim bigtableInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(biim.ref.Append("id"))
}

func (biim bigtableInstanceIamMemberAttributes) Instance() terra.StringValue {
	return terra.ReferenceString(biim.ref.Append("instance"))
}

func (biim bigtableInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(biim.ref.Append("member"))
}

func (biim bigtableInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(biim.ref.Append("project"))
}

func (biim bigtableInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(biim.ref.Append("role"))
}

func (biim bigtableInstanceIamMemberAttributes) Condition() terra.ListValue[bigtableinstanceiammember.ConditionAttributes] {
	return terra.ReferenceList[bigtableinstanceiammember.ConditionAttributes](biim.ref.Append("condition"))
}

type bigtableInstanceIamMemberState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Instance  string                                     `json:"instance"`
	Member    string                                     `json:"member"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []bigtableinstanceiammember.ConditionState `json:"condition"`
}
