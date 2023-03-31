// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	privatecacapooliammember "github.com/golingon/terraproviders/google/4.59.0/privatecacapooliammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewPrivatecaCaPoolIamMember(name string, args PrivatecaCaPoolIamMemberArgs) *PrivatecaCaPoolIamMember {
	return &PrivatecaCaPoolIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCaPoolIamMember)(nil)

type PrivatecaCaPoolIamMember struct {
	Name  string
	Args  PrivatecaCaPoolIamMemberArgs
	state *privatecaCaPoolIamMemberState
}

func (pcpim *PrivatecaCaPoolIamMember) Type() string {
	return "google_privateca_ca_pool_iam_member"
}

func (pcpim *PrivatecaCaPoolIamMember) LocalName() string {
	return pcpim.Name
}

func (pcpim *PrivatecaCaPoolIamMember) Configuration() interface{} {
	return pcpim.Args
}

func (pcpim *PrivatecaCaPoolIamMember) Attributes() privatecaCaPoolIamMemberAttributes {
	return privatecaCaPoolIamMemberAttributes{ref: terra.ReferenceResource(pcpim)}
}

func (pcpim *PrivatecaCaPoolIamMember) ImportState(av io.Reader) error {
	pcpim.state = &privatecaCaPoolIamMemberState{}
	if err := json.NewDecoder(av).Decode(pcpim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pcpim.Type(), pcpim.LocalName(), err)
	}
	return nil
}

func (pcpim *PrivatecaCaPoolIamMember) State() (*privatecaCaPoolIamMemberState, bool) {
	return pcpim.state, pcpim.state != nil
}

func (pcpim *PrivatecaCaPoolIamMember) StateMust() *privatecaCaPoolIamMemberState {
	if pcpim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pcpim.Type(), pcpim.LocalName()))
	}
	return pcpim.state
}

func (pcpim *PrivatecaCaPoolIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(pcpim)
}

type PrivatecaCaPoolIamMemberArgs struct {
	// CaPool: string, required
	CaPool terra.StringValue `hcl:"ca_pool,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *privatecacapooliammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that PrivatecaCaPoolIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type privatecaCaPoolIamMemberAttributes struct {
	ref terra.Reference
}

func (pcpim privatecaCaPoolIamMemberAttributes) CaPool() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("ca_pool"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("etag"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("id"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("location"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("member"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("project"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(pcpim.ref.Append("role"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Condition() terra.ListValue[privatecacapooliammember.ConditionAttributes] {
	return terra.ReferenceList[privatecacapooliammember.ConditionAttributes](pcpim.ref.Append("condition"))
}

type privatecaCaPoolIamMemberState struct {
	CaPool    string                                    `json:"ca_pool"`
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Location  string                                    `json:"location"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []privatecacapooliammember.ConditionState `json:"condition"`
}
