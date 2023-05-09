// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceAccountIamPolicy creates a new instance of [ServiceAccountIamPolicy].
func NewServiceAccountIamPolicy(name string, args ServiceAccountIamPolicyArgs) *ServiceAccountIamPolicy {
	return &ServiceAccountIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceAccountIamPolicy)(nil)

// ServiceAccountIamPolicy represents the Terraform resource google_service_account_iam_policy.
type ServiceAccountIamPolicy struct {
	Name      string
	Args      ServiceAccountIamPolicyArgs
	state     *serviceAccountIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceAccountIamPolicy].
func (saip *ServiceAccountIamPolicy) Type() string {
	return "google_service_account_iam_policy"
}

// LocalName returns the local name for [ServiceAccountIamPolicy].
func (saip *ServiceAccountIamPolicy) LocalName() string {
	return saip.Name
}

// Configuration returns the configuration (args) for [ServiceAccountIamPolicy].
func (saip *ServiceAccountIamPolicy) Configuration() interface{} {
	return saip.Args
}

// DependOn is used for other resources to depend on [ServiceAccountIamPolicy].
func (saip *ServiceAccountIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(saip)
}

// Dependencies returns the list of resources [ServiceAccountIamPolicy] depends_on.
func (saip *ServiceAccountIamPolicy) Dependencies() terra.Dependencies {
	return saip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceAccountIamPolicy].
func (saip *ServiceAccountIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return saip.Lifecycle
}

// Attributes returns the attributes for [ServiceAccountIamPolicy].
func (saip *ServiceAccountIamPolicy) Attributes() serviceAccountIamPolicyAttributes {
	return serviceAccountIamPolicyAttributes{ref: terra.ReferenceResource(saip)}
}

// ImportState imports the given attribute values into [ServiceAccountIamPolicy]'s state.
func (saip *ServiceAccountIamPolicy) ImportState(av io.Reader) error {
	saip.state = &serviceAccountIamPolicyState{}
	if err := json.NewDecoder(av).Decode(saip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saip.Type(), saip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceAccountIamPolicy] has state.
func (saip *ServiceAccountIamPolicy) State() (*serviceAccountIamPolicyState, bool) {
	return saip.state, saip.state != nil
}

// StateMust returns the state for [ServiceAccountIamPolicy]. Panics if the state is nil.
func (saip *ServiceAccountIamPolicy) StateMust() *serviceAccountIamPolicyState {
	if saip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saip.Type(), saip.LocalName()))
	}
	return saip.state
}

// ServiceAccountIamPolicyArgs contains the configurations for google_service_account_iam_policy.
type ServiceAccountIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// ServiceAccountId: string, required
	ServiceAccountId terra.StringValue `hcl:"service_account_id,attr" validate:"required"`
}
type serviceAccountIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_account_iam_policy.
func (saip serviceAccountIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(saip.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_account_iam_policy.
func (saip serviceAccountIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_service_account_iam_policy.
func (saip serviceAccountIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(saip.ref.Append("policy_data"))
}

// ServiceAccountId returns a reference to field service_account_id of google_service_account_iam_policy.
func (saip serviceAccountIamPolicyAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(saip.ref.Append("service_account_id"))
}

type serviceAccountIamPolicyState struct {
	Etag             string `json:"etag"`
	Id               string `json:"id"`
	PolicyData       string `json:"policy_data"`
	ServiceAccountId string `json:"service_account_id"`
}
