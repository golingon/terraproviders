// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeenvironmentiammember "github.com/golingon/terraproviders/google/4.59.0/apigeeenvironmentiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApigeeEnvironmentIamMember(name string, args ApigeeEnvironmentIamMemberArgs) *ApigeeEnvironmentIamMember {
	return &ApigeeEnvironmentIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvironmentIamMember)(nil)

type ApigeeEnvironmentIamMember struct {
	Name  string
	Args  ApigeeEnvironmentIamMemberArgs
	state *apigeeEnvironmentIamMemberState
}

func (aeim *ApigeeEnvironmentIamMember) Type() string {
	return "google_apigee_environment_iam_member"
}

func (aeim *ApigeeEnvironmentIamMember) LocalName() string {
	return aeim.Name
}

func (aeim *ApigeeEnvironmentIamMember) Configuration() interface{} {
	return aeim.Args
}

func (aeim *ApigeeEnvironmentIamMember) Attributes() apigeeEnvironmentIamMemberAttributes {
	return apigeeEnvironmentIamMemberAttributes{ref: terra.ReferenceResource(aeim)}
}

func (aeim *ApigeeEnvironmentIamMember) ImportState(av io.Reader) error {
	aeim.state = &apigeeEnvironmentIamMemberState{}
	if err := json.NewDecoder(av).Decode(aeim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeim.Type(), aeim.LocalName(), err)
	}
	return nil
}

func (aeim *ApigeeEnvironmentIamMember) State() (*apigeeEnvironmentIamMemberState, bool) {
	return aeim.state, aeim.state != nil
}

func (aeim *ApigeeEnvironmentIamMember) StateMust() *apigeeEnvironmentIamMemberState {
	if aeim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeim.Type(), aeim.LocalName()))
	}
	return aeim.state
}

func (aeim *ApigeeEnvironmentIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(aeim)
}

type ApigeeEnvironmentIamMemberArgs struct {
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigeeenvironmentiammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that ApigeeEnvironmentIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apigeeEnvironmentIamMemberAttributes struct {
	ref terra.Reference
}

func (aeim apigeeEnvironmentIamMemberAttributes) EnvId() terra.StringValue {
	return terra.ReferenceString(aeim.ref.Append("env_id"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(aeim.ref.Append("etag"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(aeim.ref.Append("id"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(aeim.ref.Append("member"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) OrgId() terra.StringValue {
	return terra.ReferenceString(aeim.ref.Append("org_id"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(aeim.ref.Append("role"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) Condition() terra.ListValue[apigeeenvironmentiammember.ConditionAttributes] {
	return terra.ReferenceList[apigeeenvironmentiammember.ConditionAttributes](aeim.ref.Append("condition"))
}

type apigeeEnvironmentIamMemberState struct {
	EnvId     string                                      `json:"env_id"`
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Member    string                                      `json:"member"`
	OrgId     string                                      `json:"org_id"`
	Role      string                                      `json:"role"`
	Condition []apigeeenvironmentiammember.ConditionState `json:"condition"`
}
