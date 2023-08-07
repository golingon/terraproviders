// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapwebtypecomputeiammember "github.com/golingon/terraproviders/google/4.76.0/iapwebtypecomputeiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebTypeComputeIamMember creates a new instance of [IapWebTypeComputeIamMember].
func NewIapWebTypeComputeIamMember(name string, args IapWebTypeComputeIamMemberArgs) *IapWebTypeComputeIamMember {
	return &IapWebTypeComputeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebTypeComputeIamMember)(nil)

// IapWebTypeComputeIamMember represents the Terraform resource google_iap_web_type_compute_iam_member.
type IapWebTypeComputeIamMember struct {
	Name      string
	Args      IapWebTypeComputeIamMemberArgs
	state     *iapWebTypeComputeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebTypeComputeIamMember].
func (iwtcim *IapWebTypeComputeIamMember) Type() string {
	return "google_iap_web_type_compute_iam_member"
}

// LocalName returns the local name for [IapWebTypeComputeIamMember].
func (iwtcim *IapWebTypeComputeIamMember) LocalName() string {
	return iwtcim.Name
}

// Configuration returns the configuration (args) for [IapWebTypeComputeIamMember].
func (iwtcim *IapWebTypeComputeIamMember) Configuration() interface{} {
	return iwtcim.Args
}

// DependOn is used for other resources to depend on [IapWebTypeComputeIamMember].
func (iwtcim *IapWebTypeComputeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iwtcim)
}

// Dependencies returns the list of resources [IapWebTypeComputeIamMember] depends_on.
func (iwtcim *IapWebTypeComputeIamMember) Dependencies() terra.Dependencies {
	return iwtcim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebTypeComputeIamMember].
func (iwtcim *IapWebTypeComputeIamMember) LifecycleManagement() *terra.Lifecycle {
	return iwtcim.Lifecycle
}

// Attributes returns the attributes for [IapWebTypeComputeIamMember].
func (iwtcim *IapWebTypeComputeIamMember) Attributes() iapWebTypeComputeIamMemberAttributes {
	return iapWebTypeComputeIamMemberAttributes{ref: terra.ReferenceResource(iwtcim)}
}

// ImportState imports the given attribute values into [IapWebTypeComputeIamMember]'s state.
func (iwtcim *IapWebTypeComputeIamMember) ImportState(av io.Reader) error {
	iwtcim.state = &iapWebTypeComputeIamMemberState{}
	if err := json.NewDecoder(av).Decode(iwtcim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwtcim.Type(), iwtcim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebTypeComputeIamMember] has state.
func (iwtcim *IapWebTypeComputeIamMember) State() (*iapWebTypeComputeIamMemberState, bool) {
	return iwtcim.state, iwtcim.state != nil
}

// StateMust returns the state for [IapWebTypeComputeIamMember]. Panics if the state is nil.
func (iwtcim *IapWebTypeComputeIamMember) StateMust() *iapWebTypeComputeIamMemberState {
	if iwtcim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwtcim.Type(), iwtcim.LocalName()))
	}
	return iwtcim.state
}

// IapWebTypeComputeIamMemberArgs contains the configurations for google_iap_web_type_compute_iam_member.
type IapWebTypeComputeIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebtypecomputeiammember.Condition `hcl:"condition,block"`
}
type iapWebTypeComputeIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_type_compute_iam_member.
func (iwtcim iapWebTypeComputeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwtcim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_type_compute_iam_member.
func (iwtcim iapWebTypeComputeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwtcim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_web_type_compute_iam_member.
func (iwtcim iapWebTypeComputeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(iwtcim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_web_type_compute_iam_member.
func (iwtcim iapWebTypeComputeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwtcim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_type_compute_iam_member.
func (iwtcim iapWebTypeComputeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwtcim.ref.Append("role"))
}

func (iwtcim iapWebTypeComputeIamMemberAttributes) Condition() terra.ListValue[iapwebtypecomputeiammember.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebtypecomputeiammember.ConditionAttributes](iwtcim.ref.Append("condition"))
}

type iapWebTypeComputeIamMemberState struct {
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Member    string                                      `json:"member"`
	Project   string                                      `json:"project"`
	Role      string                                      `json:"role"`
	Condition []iapwebtypecomputeiammember.ConditionState `json:"condition"`
}
