// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapappengineversioniammember "github.com/golingon/terraproviders/google/4.59.0/iapappengineversioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIapAppEngineVersionIamMember(name string, args IapAppEngineVersionIamMemberArgs) *IapAppEngineVersionIamMember {
	return &IapAppEngineVersionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineVersionIamMember)(nil)

type IapAppEngineVersionIamMember struct {
	Name  string
	Args  IapAppEngineVersionIamMemberArgs
	state *iapAppEngineVersionIamMemberState
}

func (iaevim *IapAppEngineVersionIamMember) Type() string {
	return "google_iap_app_engine_version_iam_member"
}

func (iaevim *IapAppEngineVersionIamMember) LocalName() string {
	return iaevim.Name
}

func (iaevim *IapAppEngineVersionIamMember) Configuration() interface{} {
	return iaevim.Args
}

func (iaevim *IapAppEngineVersionIamMember) Attributes() iapAppEngineVersionIamMemberAttributes {
	return iapAppEngineVersionIamMemberAttributes{ref: terra.ReferenceResource(iaevim)}
}

func (iaevim *IapAppEngineVersionIamMember) ImportState(av io.Reader) error {
	iaevim.state = &iapAppEngineVersionIamMemberState{}
	if err := json.NewDecoder(av).Decode(iaevim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaevim.Type(), iaevim.LocalName(), err)
	}
	return nil
}

func (iaevim *IapAppEngineVersionIamMember) State() (*iapAppEngineVersionIamMemberState, bool) {
	return iaevim.state, iaevim.state != nil
}

func (iaevim *IapAppEngineVersionIamMember) StateMust() *iapAppEngineVersionIamMemberState {
	if iaevim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaevim.Type(), iaevim.LocalName()))
	}
	return iaevim.state
}

func (iaevim *IapAppEngineVersionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iaevim)
}

type IapAppEngineVersionIamMemberArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// VersionId: string, required
	VersionId terra.StringValue `hcl:"version_id,attr" validate:"required"`
	// Condition: optional
	Condition *iapappengineversioniammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that IapAppEngineVersionIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type iapAppEngineVersionIamMemberAttributes struct {
	ref terra.Reference
}

func (iaevim iapAppEngineVersionIamMemberAttributes) AppId() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("app_id"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("etag"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("id"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("member"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("project"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("role"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Service() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("service"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) VersionId() terra.StringValue {
	return terra.ReferenceString(iaevim.ref.Append("version_id"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Condition() terra.ListValue[iapappengineversioniammember.ConditionAttributes] {
	return terra.ReferenceList[iapappengineversioniammember.ConditionAttributes](iaevim.ref.Append("condition"))
}

type iapAppEngineVersionIamMemberState struct {
	AppId     string                                        `json:"app_id"`
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	Service   string                                        `json:"service"`
	VersionId string                                        `json:"version_id"`
	Condition []iapappengineversioniammember.ConditionState `json:"condition"`
}
