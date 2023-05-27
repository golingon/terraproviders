// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanageraccesspolicyiammember "github.com/golingon/terraproviders/google/4.66.0/accesscontextmanageraccesspolicyiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAccessPolicyIamMember creates a new instance of [AccessContextManagerAccessPolicyIamMember].
func NewAccessContextManagerAccessPolicyIamMember(name string, args AccessContextManagerAccessPolicyIamMemberArgs) *AccessContextManagerAccessPolicyIamMember {
	return &AccessContextManagerAccessPolicyIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAccessPolicyIamMember)(nil)

// AccessContextManagerAccessPolicyIamMember represents the Terraform resource google_access_context_manager_access_policy_iam_member.
type AccessContextManagerAccessPolicyIamMember struct {
	Name      string
	Args      AccessContextManagerAccessPolicyIamMemberArgs
	state     *accessContextManagerAccessPolicyIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAccessPolicyIamMember].
func (acmapim *AccessContextManagerAccessPolicyIamMember) Type() string {
	return "google_access_context_manager_access_policy_iam_member"
}

// LocalName returns the local name for [AccessContextManagerAccessPolicyIamMember].
func (acmapim *AccessContextManagerAccessPolicyIamMember) LocalName() string {
	return acmapim.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAccessPolicyIamMember].
func (acmapim *AccessContextManagerAccessPolicyIamMember) Configuration() interface{} {
	return acmapim.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAccessPolicyIamMember].
func (acmapim *AccessContextManagerAccessPolicyIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(acmapim)
}

// Dependencies returns the list of resources [AccessContextManagerAccessPolicyIamMember] depends_on.
func (acmapim *AccessContextManagerAccessPolicyIamMember) Dependencies() terra.Dependencies {
	return acmapim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAccessPolicyIamMember].
func (acmapim *AccessContextManagerAccessPolicyIamMember) LifecycleManagement() *terra.Lifecycle {
	return acmapim.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAccessPolicyIamMember].
func (acmapim *AccessContextManagerAccessPolicyIamMember) Attributes() accessContextManagerAccessPolicyIamMemberAttributes {
	return accessContextManagerAccessPolicyIamMemberAttributes{ref: terra.ReferenceResource(acmapim)}
}

// ImportState imports the given attribute values into [AccessContextManagerAccessPolicyIamMember]'s state.
func (acmapim *AccessContextManagerAccessPolicyIamMember) ImportState(av io.Reader) error {
	acmapim.state = &accessContextManagerAccessPolicyIamMemberState{}
	if err := json.NewDecoder(av).Decode(acmapim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmapim.Type(), acmapim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAccessPolicyIamMember] has state.
func (acmapim *AccessContextManagerAccessPolicyIamMember) State() (*accessContextManagerAccessPolicyIamMemberState, bool) {
	return acmapim.state, acmapim.state != nil
}

// StateMust returns the state for [AccessContextManagerAccessPolicyIamMember]. Panics if the state is nil.
func (acmapim *AccessContextManagerAccessPolicyIamMember) StateMust() *accessContextManagerAccessPolicyIamMemberState {
	if acmapim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmapim.Type(), acmapim.LocalName()))
	}
	return acmapim.state
}

// AccessContextManagerAccessPolicyIamMemberArgs contains the configurations for google_access_context_manager_access_policy_iam_member.
type AccessContextManagerAccessPolicyIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *accesscontextmanageraccesspolicyiammember.Condition `hcl:"condition,block"`
}
type accessContextManagerAccessPolicyIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_access_context_manager_access_policy_iam_member.
func (acmapim accessContextManagerAccessPolicyIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(acmapim.ref.Append("etag"))
}

// Id returns a reference to field id of google_access_context_manager_access_policy_iam_member.
func (acmapim accessContextManagerAccessPolicyIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmapim.ref.Append("id"))
}

// Member returns a reference to field member of google_access_context_manager_access_policy_iam_member.
func (acmapim accessContextManagerAccessPolicyIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(acmapim.ref.Append("member"))
}

// Name returns a reference to field name of google_access_context_manager_access_policy_iam_member.
func (acmapim accessContextManagerAccessPolicyIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmapim.ref.Append("name"))
}

// Role returns a reference to field role of google_access_context_manager_access_policy_iam_member.
func (acmapim accessContextManagerAccessPolicyIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(acmapim.ref.Append("role"))
}

func (acmapim accessContextManagerAccessPolicyIamMemberAttributes) Condition() terra.ListValue[accesscontextmanageraccesspolicyiammember.ConditionAttributes] {
	return terra.ReferenceAsList[accesscontextmanageraccesspolicyiammember.ConditionAttributes](acmapim.ref.Append("condition"))
}

type accessContextManagerAccessPolicyIamMemberState struct {
	Etag      string                                                     `json:"etag"`
	Id        string                                                     `json:"id"`
	Member    string                                                     `json:"member"`
	Name      string                                                     `json:"name"`
	Role      string                                                     `json:"role"`
	Condition []accesscontextmanageraccesspolicyiammember.ConditionState `json:"condition"`
}
