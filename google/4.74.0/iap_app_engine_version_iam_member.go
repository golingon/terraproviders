// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapappengineversioniammember "github.com/golingon/terraproviders/google/4.74.0/iapappengineversioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapAppEngineVersionIamMember creates a new instance of [IapAppEngineVersionIamMember].
func NewIapAppEngineVersionIamMember(name string, args IapAppEngineVersionIamMemberArgs) *IapAppEngineVersionIamMember {
	return &IapAppEngineVersionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineVersionIamMember)(nil)

// IapAppEngineVersionIamMember represents the Terraform resource google_iap_app_engine_version_iam_member.
type IapAppEngineVersionIamMember struct {
	Name      string
	Args      IapAppEngineVersionIamMemberArgs
	state     *iapAppEngineVersionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapAppEngineVersionIamMember].
func (iaevim *IapAppEngineVersionIamMember) Type() string {
	return "google_iap_app_engine_version_iam_member"
}

// LocalName returns the local name for [IapAppEngineVersionIamMember].
func (iaevim *IapAppEngineVersionIamMember) LocalName() string {
	return iaevim.Name
}

// Configuration returns the configuration (args) for [IapAppEngineVersionIamMember].
func (iaevim *IapAppEngineVersionIamMember) Configuration() interface{} {
	return iaevim.Args
}

// DependOn is used for other resources to depend on [IapAppEngineVersionIamMember].
func (iaevim *IapAppEngineVersionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iaevim)
}

// Dependencies returns the list of resources [IapAppEngineVersionIamMember] depends_on.
func (iaevim *IapAppEngineVersionIamMember) Dependencies() terra.Dependencies {
	return iaevim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapAppEngineVersionIamMember].
func (iaevim *IapAppEngineVersionIamMember) LifecycleManagement() *terra.Lifecycle {
	return iaevim.Lifecycle
}

// Attributes returns the attributes for [IapAppEngineVersionIamMember].
func (iaevim *IapAppEngineVersionIamMember) Attributes() iapAppEngineVersionIamMemberAttributes {
	return iapAppEngineVersionIamMemberAttributes{ref: terra.ReferenceResource(iaevim)}
}

// ImportState imports the given attribute values into [IapAppEngineVersionIamMember]'s state.
func (iaevim *IapAppEngineVersionIamMember) ImportState(av io.Reader) error {
	iaevim.state = &iapAppEngineVersionIamMemberState{}
	if err := json.NewDecoder(av).Decode(iaevim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaevim.Type(), iaevim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapAppEngineVersionIamMember] has state.
func (iaevim *IapAppEngineVersionIamMember) State() (*iapAppEngineVersionIamMemberState, bool) {
	return iaevim.state, iaevim.state != nil
}

// StateMust returns the state for [IapAppEngineVersionIamMember]. Panics if the state is nil.
func (iaevim *IapAppEngineVersionIamMember) StateMust() *iapAppEngineVersionIamMemberState {
	if iaevim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaevim.Type(), iaevim.LocalName()))
	}
	return iaevim.state
}

// IapAppEngineVersionIamMemberArgs contains the configurations for google_iap_app_engine_version_iam_member.
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
}
type iapAppEngineVersionIamMemberAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("role"))
}

// Service returns a reference to field service of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("service"))
}

// VersionId returns a reference to field version_id of google_iap_app_engine_version_iam_member.
func (iaevim iapAppEngineVersionIamMemberAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(iaevim.ref.Append("version_id"))
}

func (iaevim iapAppEngineVersionIamMemberAttributes) Condition() terra.ListValue[iapappengineversioniammember.ConditionAttributes] {
	return terra.ReferenceAsList[iapappengineversioniammember.ConditionAttributes](iaevim.ref.Append("condition"))
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