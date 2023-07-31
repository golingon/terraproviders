// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	serviceaccountiammember "github.com/golingon/terraproviders/google/4.75.1/serviceaccountiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceAccountIamMember creates a new instance of [ServiceAccountIamMember].
func NewServiceAccountIamMember(name string, args ServiceAccountIamMemberArgs) *ServiceAccountIamMember {
	return &ServiceAccountIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceAccountIamMember)(nil)

// ServiceAccountIamMember represents the Terraform resource google_service_account_iam_member.
type ServiceAccountIamMember struct {
	Name      string
	Args      ServiceAccountIamMemberArgs
	state     *serviceAccountIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceAccountIamMember].
func (saim *ServiceAccountIamMember) Type() string {
	return "google_service_account_iam_member"
}

// LocalName returns the local name for [ServiceAccountIamMember].
func (saim *ServiceAccountIamMember) LocalName() string {
	return saim.Name
}

// Configuration returns the configuration (args) for [ServiceAccountIamMember].
func (saim *ServiceAccountIamMember) Configuration() interface{} {
	return saim.Args
}

// DependOn is used for other resources to depend on [ServiceAccountIamMember].
func (saim *ServiceAccountIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(saim)
}

// Dependencies returns the list of resources [ServiceAccountIamMember] depends_on.
func (saim *ServiceAccountIamMember) Dependencies() terra.Dependencies {
	return saim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceAccountIamMember].
func (saim *ServiceAccountIamMember) LifecycleManagement() *terra.Lifecycle {
	return saim.Lifecycle
}

// Attributes returns the attributes for [ServiceAccountIamMember].
func (saim *ServiceAccountIamMember) Attributes() serviceAccountIamMemberAttributes {
	return serviceAccountIamMemberAttributes{ref: terra.ReferenceResource(saim)}
}

// ImportState imports the given attribute values into [ServiceAccountIamMember]'s state.
func (saim *ServiceAccountIamMember) ImportState(av io.Reader) error {
	saim.state = &serviceAccountIamMemberState{}
	if err := json.NewDecoder(av).Decode(saim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saim.Type(), saim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceAccountIamMember] has state.
func (saim *ServiceAccountIamMember) State() (*serviceAccountIamMemberState, bool) {
	return saim.state, saim.state != nil
}

// StateMust returns the state for [ServiceAccountIamMember]. Panics if the state is nil.
func (saim *ServiceAccountIamMember) StateMust() *serviceAccountIamMemberState {
	if saim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saim.Type(), saim.LocalName()))
	}
	return saim.state
}

// ServiceAccountIamMemberArgs contains the configurations for google_service_account_iam_member.
type ServiceAccountIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceAccountId: string, required
	ServiceAccountId terra.StringValue `hcl:"service_account_id,attr" validate:"required"`
	// Condition: optional
	Condition *serviceaccountiammember.Condition `hcl:"condition,block"`
}
type serviceAccountIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_account_iam_member.
func (saim serviceAccountIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(saim.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_account_iam_member.
func (saim serviceAccountIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saim.ref.Append("id"))
}

// Member returns a reference to field member of google_service_account_iam_member.
func (saim serviceAccountIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(saim.ref.Append("member"))
}

// Role returns a reference to field role of google_service_account_iam_member.
func (saim serviceAccountIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(saim.ref.Append("role"))
}

// ServiceAccountId returns a reference to field service_account_id of google_service_account_iam_member.
func (saim serviceAccountIamMemberAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(saim.ref.Append("service_account_id"))
}

func (saim serviceAccountIamMemberAttributes) Condition() terra.ListValue[serviceaccountiammember.ConditionAttributes] {
	return terra.ReferenceAsList[serviceaccountiammember.ConditionAttributes](saim.ref.Append("condition"))
}

type serviceAccountIamMemberState struct {
	Etag             string                                   `json:"etag"`
	Id               string                                   `json:"id"`
	Member           string                                   `json:"member"`
	Role             string                                   `json:"role"`
	ServiceAccountId string                                   `json:"service_account_id"`
	Condition        []serviceaccountiammember.ConditionState `json:"condition"`
}
