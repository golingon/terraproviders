// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iapwebtypeappengineiammember "github.com/golingon/terraproviders/googlebeta/4.62.0/iapwebtypeappengineiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebTypeAppEngineIamMember creates a new instance of [IapWebTypeAppEngineIamMember].
func NewIapWebTypeAppEngineIamMember(name string, args IapWebTypeAppEngineIamMemberArgs) *IapWebTypeAppEngineIamMember {
	return &IapWebTypeAppEngineIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebTypeAppEngineIamMember)(nil)

// IapWebTypeAppEngineIamMember represents the Terraform resource google_iap_web_type_app_engine_iam_member.
type IapWebTypeAppEngineIamMember struct {
	Name      string
	Args      IapWebTypeAppEngineIamMemberArgs
	state     *iapWebTypeAppEngineIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebTypeAppEngineIamMember].
func (iwtaeim *IapWebTypeAppEngineIamMember) Type() string {
	return "google_iap_web_type_app_engine_iam_member"
}

// LocalName returns the local name for [IapWebTypeAppEngineIamMember].
func (iwtaeim *IapWebTypeAppEngineIamMember) LocalName() string {
	return iwtaeim.Name
}

// Configuration returns the configuration (args) for [IapWebTypeAppEngineIamMember].
func (iwtaeim *IapWebTypeAppEngineIamMember) Configuration() interface{} {
	return iwtaeim.Args
}

// DependOn is used for other resources to depend on [IapWebTypeAppEngineIamMember].
func (iwtaeim *IapWebTypeAppEngineIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iwtaeim)
}

// Dependencies returns the list of resources [IapWebTypeAppEngineIamMember] depends_on.
func (iwtaeim *IapWebTypeAppEngineIamMember) Dependencies() terra.Dependencies {
	return iwtaeim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebTypeAppEngineIamMember].
func (iwtaeim *IapWebTypeAppEngineIamMember) LifecycleManagement() *terra.Lifecycle {
	return iwtaeim.Lifecycle
}

// Attributes returns the attributes for [IapWebTypeAppEngineIamMember].
func (iwtaeim *IapWebTypeAppEngineIamMember) Attributes() iapWebTypeAppEngineIamMemberAttributes {
	return iapWebTypeAppEngineIamMemberAttributes{ref: terra.ReferenceResource(iwtaeim)}
}

// ImportState imports the given attribute values into [IapWebTypeAppEngineIamMember]'s state.
func (iwtaeim *IapWebTypeAppEngineIamMember) ImportState(av io.Reader) error {
	iwtaeim.state = &iapWebTypeAppEngineIamMemberState{}
	if err := json.NewDecoder(av).Decode(iwtaeim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwtaeim.Type(), iwtaeim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebTypeAppEngineIamMember] has state.
func (iwtaeim *IapWebTypeAppEngineIamMember) State() (*iapWebTypeAppEngineIamMemberState, bool) {
	return iwtaeim.state, iwtaeim.state != nil
}

// StateMust returns the state for [IapWebTypeAppEngineIamMember]. Panics if the state is nil.
func (iwtaeim *IapWebTypeAppEngineIamMember) StateMust() *iapWebTypeAppEngineIamMemberState {
	if iwtaeim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwtaeim.Type(), iwtaeim.LocalName()))
	}
	return iwtaeim.state
}

// IapWebTypeAppEngineIamMemberArgs contains the configurations for google_iap_web_type_app_engine_iam_member.
type IapWebTypeAppEngineIamMemberArgs struct {
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
	// Condition: optional
	Condition *iapwebtypeappengineiammember.Condition `hcl:"condition,block"`
}
type iapWebTypeAppEngineIamMemberAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_web_type_app_engine_iam_member.
func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iwtaeim.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_web_type_app_engine_iam_member.
func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwtaeim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_type_app_engine_iam_member.
func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwtaeim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_web_type_app_engine_iam_member.
func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(iwtaeim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_web_type_app_engine_iam_member.
func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwtaeim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_type_app_engine_iam_member.
func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwtaeim.ref.Append("role"))
}

func (iwtaeim iapWebTypeAppEngineIamMemberAttributes) Condition() terra.ListValue[iapwebtypeappengineiammember.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebtypeappengineiammember.ConditionAttributes](iwtaeim.ref.Append("condition"))
}

type iapWebTypeAppEngineIamMemberState struct {
	AppId     string                                        `json:"app_id"`
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	Condition []iapwebtypeappengineiammember.ConditionState `json:"condition"`
}
