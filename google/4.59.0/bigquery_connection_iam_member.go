// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryconnectioniammember "github.com/golingon/terraproviders/google/4.59.0/bigqueryconnectioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBigqueryConnectionIamMember(name string, args BigqueryConnectionIamMemberArgs) *BigqueryConnectionIamMember {
	return &BigqueryConnectionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryConnectionIamMember)(nil)

type BigqueryConnectionIamMember struct {
	Name  string
	Args  BigqueryConnectionIamMemberArgs
	state *bigqueryConnectionIamMemberState
}

func (bcim *BigqueryConnectionIamMember) Type() string {
	return "google_bigquery_connection_iam_member"
}

func (bcim *BigqueryConnectionIamMember) LocalName() string {
	return bcim.Name
}

func (bcim *BigqueryConnectionIamMember) Configuration() interface{} {
	return bcim.Args
}

func (bcim *BigqueryConnectionIamMember) Attributes() bigqueryConnectionIamMemberAttributes {
	return bigqueryConnectionIamMemberAttributes{ref: terra.ReferenceResource(bcim)}
}

func (bcim *BigqueryConnectionIamMember) ImportState(av io.Reader) error {
	bcim.state = &bigqueryConnectionIamMemberState{}
	if err := json.NewDecoder(av).Decode(bcim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcim.Type(), bcim.LocalName(), err)
	}
	return nil
}

func (bcim *BigqueryConnectionIamMember) State() (*bigqueryConnectionIamMemberState, bool) {
	return bcim.state, bcim.state != nil
}

func (bcim *BigqueryConnectionIamMember) StateMust() *bigqueryConnectionIamMemberState {
	if bcim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcim.Type(), bcim.LocalName()))
	}
	return bcim.state
}

func (bcim *BigqueryConnectionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(bcim)
}

type BigqueryConnectionIamMemberArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
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
	Condition *bigqueryconnectioniammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that BigqueryConnectionIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type bigqueryConnectionIamMemberAttributes struct {
	ref terra.Reference
}

func (bcim bigqueryConnectionIamMemberAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("connection_id"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("etag"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("id"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("location"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("member"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("project"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(bcim.ref.Append("role"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Condition() terra.ListValue[bigqueryconnectioniammember.ConditionAttributes] {
	return terra.ReferenceList[bigqueryconnectioniammember.ConditionAttributes](bcim.ref.Append("condition"))
}

type bigqueryConnectionIamMemberState struct {
	ConnectionId string                                       `json:"connection_id"`
	Etag         string                                       `json:"etag"`
	Id           string                                       `json:"id"`
	Location     string                                       `json:"location"`
	Member       string                                       `json:"member"`
	Project      string                                       `json:"project"`
	Role         string                                       `json:"role"`
	Condition    []bigqueryconnectioniammember.ConditionState `json:"condition"`
}
