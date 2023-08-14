// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeenvironmentiammember "github.com/golingon/terraproviders/google/4.75.1/apigeeenvironmentiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvironmentIamMember creates a new instance of [ApigeeEnvironmentIamMember].
func NewApigeeEnvironmentIamMember(name string, args ApigeeEnvironmentIamMemberArgs) *ApigeeEnvironmentIamMember {
	return &ApigeeEnvironmentIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvironmentIamMember)(nil)

// ApigeeEnvironmentIamMember represents the Terraform resource google_apigee_environment_iam_member.
type ApigeeEnvironmentIamMember struct {
	Name      string
	Args      ApigeeEnvironmentIamMemberArgs
	state     *apigeeEnvironmentIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvironmentIamMember].
func (aeim *ApigeeEnvironmentIamMember) Type() string {
	return "google_apigee_environment_iam_member"
}

// LocalName returns the local name for [ApigeeEnvironmentIamMember].
func (aeim *ApigeeEnvironmentIamMember) LocalName() string {
	return aeim.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvironmentIamMember].
func (aeim *ApigeeEnvironmentIamMember) Configuration() interface{} {
	return aeim.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvironmentIamMember].
func (aeim *ApigeeEnvironmentIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(aeim)
}

// Dependencies returns the list of resources [ApigeeEnvironmentIamMember] depends_on.
func (aeim *ApigeeEnvironmentIamMember) Dependencies() terra.Dependencies {
	return aeim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvironmentIamMember].
func (aeim *ApigeeEnvironmentIamMember) LifecycleManagement() *terra.Lifecycle {
	return aeim.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvironmentIamMember].
func (aeim *ApigeeEnvironmentIamMember) Attributes() apigeeEnvironmentIamMemberAttributes {
	return apigeeEnvironmentIamMemberAttributes{ref: terra.ReferenceResource(aeim)}
}

// ImportState imports the given attribute values into [ApigeeEnvironmentIamMember]'s state.
func (aeim *ApigeeEnvironmentIamMember) ImportState(av io.Reader) error {
	aeim.state = &apigeeEnvironmentIamMemberState{}
	if err := json.NewDecoder(av).Decode(aeim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeim.Type(), aeim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvironmentIamMember] has state.
func (aeim *ApigeeEnvironmentIamMember) State() (*apigeeEnvironmentIamMemberState, bool) {
	return aeim.state, aeim.state != nil
}

// StateMust returns the state for [ApigeeEnvironmentIamMember]. Panics if the state is nil.
func (aeim *ApigeeEnvironmentIamMember) StateMust() *apigeeEnvironmentIamMemberState {
	if aeim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeim.Type(), aeim.LocalName()))
	}
	return aeim.state
}

// ApigeeEnvironmentIamMemberArgs contains the configurations for google_apigee_environment_iam_member.
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
}
type apigeeEnvironmentIamMemberAttributes struct {
	ref terra.Reference
}

// EnvId returns a reference to field env_id of google_apigee_environment_iam_member.
func (aeim apigeeEnvironmentIamMemberAttributes) EnvId() terra.StringValue {
	return terra.ReferenceAsString(aeim.ref.Append("env_id"))
}

// Etag returns a reference to field etag of google_apigee_environment_iam_member.
func (aeim apigeeEnvironmentIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(aeim.ref.Append("etag"))
}

// Id returns a reference to field id of google_apigee_environment_iam_member.
func (aeim apigeeEnvironmentIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeim.ref.Append("id"))
}

// Member returns a reference to field member of google_apigee_environment_iam_member.
func (aeim apigeeEnvironmentIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(aeim.ref.Append("member"))
}

// OrgId returns a reference to field org_id of google_apigee_environment_iam_member.
func (aeim apigeeEnvironmentIamMemberAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(aeim.ref.Append("org_id"))
}

// Role returns a reference to field role of google_apigee_environment_iam_member.
func (aeim apigeeEnvironmentIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(aeim.ref.Append("role"))
}

func (aeim apigeeEnvironmentIamMemberAttributes) Condition() terra.ListValue[apigeeenvironmentiammember.ConditionAttributes] {
	return terra.ReferenceAsList[apigeeenvironmentiammember.ConditionAttributes](aeim.ref.Append("condition"))
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