// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapappengineserviceiammember "github.com/golingon/terraproviders/google/4.59.0/iapappengineserviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapAppEngineServiceIamMember creates a new instance of [IapAppEngineServiceIamMember].
func NewIapAppEngineServiceIamMember(name string, args IapAppEngineServiceIamMemberArgs) *IapAppEngineServiceIamMember {
	return &IapAppEngineServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineServiceIamMember)(nil)

// IapAppEngineServiceIamMember represents the Terraform resource google_iap_app_engine_service_iam_member.
type IapAppEngineServiceIamMember struct {
	Name      string
	Args      IapAppEngineServiceIamMemberArgs
	state     *iapAppEngineServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapAppEngineServiceIamMember].
func (iaesim *IapAppEngineServiceIamMember) Type() string {
	return "google_iap_app_engine_service_iam_member"
}

// LocalName returns the local name for [IapAppEngineServiceIamMember].
func (iaesim *IapAppEngineServiceIamMember) LocalName() string {
	return iaesim.Name
}

// Configuration returns the configuration (args) for [IapAppEngineServiceIamMember].
func (iaesim *IapAppEngineServiceIamMember) Configuration() interface{} {
	return iaesim.Args
}

// DependOn is used for other resources to depend on [IapAppEngineServiceIamMember].
func (iaesim *IapAppEngineServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iaesim)
}

// Dependencies returns the list of resources [IapAppEngineServiceIamMember] depends_on.
func (iaesim *IapAppEngineServiceIamMember) Dependencies() terra.Dependencies {
	return iaesim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapAppEngineServiceIamMember].
func (iaesim *IapAppEngineServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return iaesim.Lifecycle
}

// Attributes returns the attributes for [IapAppEngineServiceIamMember].
func (iaesim *IapAppEngineServiceIamMember) Attributes() iapAppEngineServiceIamMemberAttributes {
	return iapAppEngineServiceIamMemberAttributes{ref: terra.ReferenceResource(iaesim)}
}

// ImportState imports the given attribute values into [IapAppEngineServiceIamMember]'s state.
func (iaesim *IapAppEngineServiceIamMember) ImportState(av io.Reader) error {
	iaesim.state = &iapAppEngineServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(iaesim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaesim.Type(), iaesim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapAppEngineServiceIamMember] has state.
func (iaesim *IapAppEngineServiceIamMember) State() (*iapAppEngineServiceIamMemberState, bool) {
	return iaesim.state, iaesim.state != nil
}

// StateMust returns the state for [IapAppEngineServiceIamMember]. Panics if the state is nil.
func (iaesim *IapAppEngineServiceIamMember) StateMust() *iapAppEngineServiceIamMemberState {
	if iaesim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaesim.Type(), iaesim.LocalName()))
	}
	return iaesim.state
}

// IapAppEngineServiceIamMemberArgs contains the configurations for google_iap_app_engine_service_iam_member.
type IapAppEngineServiceIamMemberArgs struct {
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
	// Condition: optional
	Condition *iapappengineserviceiammember.Condition `hcl:"condition,block"`
}
type iapAppEngineServiceIamMemberAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("role"))
}

// Service returns a reference to field service of google_iap_app_engine_service_iam_member.
func (iaesim iapAppEngineServiceIamMemberAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(iaesim.ref.Append("service"))
}

func (iaesim iapAppEngineServiceIamMemberAttributes) Condition() terra.ListValue[iapappengineserviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[iapappengineserviceiammember.ConditionAttributes](iaesim.ref.Append("condition"))
}

type iapAppEngineServiceIamMemberState struct {
	AppId     string                                        `json:"app_id"`
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	Service   string                                        `json:"service"`
	Condition []iapappengineserviceiammember.ConditionState `json:"condition"`
}
