// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanageregresspolicy "github.com/golingon/terraproviders/google/4.72.1/accesscontextmanageregresspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerEgressPolicy creates a new instance of [AccessContextManagerEgressPolicy].
func NewAccessContextManagerEgressPolicy(name string, args AccessContextManagerEgressPolicyArgs) *AccessContextManagerEgressPolicy {
	return &AccessContextManagerEgressPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerEgressPolicy)(nil)

// AccessContextManagerEgressPolicy represents the Terraform resource google_access_context_manager_egress_policy.
type AccessContextManagerEgressPolicy struct {
	Name      string
	Args      AccessContextManagerEgressPolicyArgs
	state     *accessContextManagerEgressPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerEgressPolicy].
func (acmep *AccessContextManagerEgressPolicy) Type() string {
	return "google_access_context_manager_egress_policy"
}

// LocalName returns the local name for [AccessContextManagerEgressPolicy].
func (acmep *AccessContextManagerEgressPolicy) LocalName() string {
	return acmep.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerEgressPolicy].
func (acmep *AccessContextManagerEgressPolicy) Configuration() interface{} {
	return acmep.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerEgressPolicy].
func (acmep *AccessContextManagerEgressPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(acmep)
}

// Dependencies returns the list of resources [AccessContextManagerEgressPolicy] depends_on.
func (acmep *AccessContextManagerEgressPolicy) Dependencies() terra.Dependencies {
	return acmep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerEgressPolicy].
func (acmep *AccessContextManagerEgressPolicy) LifecycleManagement() *terra.Lifecycle {
	return acmep.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerEgressPolicy].
func (acmep *AccessContextManagerEgressPolicy) Attributes() accessContextManagerEgressPolicyAttributes {
	return accessContextManagerEgressPolicyAttributes{ref: terra.ReferenceResource(acmep)}
}

// ImportState imports the given attribute values into [AccessContextManagerEgressPolicy]'s state.
func (acmep *AccessContextManagerEgressPolicy) ImportState(av io.Reader) error {
	acmep.state = &accessContextManagerEgressPolicyState{}
	if err := json.NewDecoder(av).Decode(acmep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmep.Type(), acmep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerEgressPolicy] has state.
func (acmep *AccessContextManagerEgressPolicy) State() (*accessContextManagerEgressPolicyState, bool) {
	return acmep.state, acmep.state != nil
}

// StateMust returns the state for [AccessContextManagerEgressPolicy]. Panics if the state is nil.
func (acmep *AccessContextManagerEgressPolicy) StateMust() *accessContextManagerEgressPolicyState {
	if acmep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmep.Type(), acmep.LocalName()))
	}
	return acmep.state
}

// AccessContextManagerEgressPolicyArgs contains the configurations for google_access_context_manager_egress_policy.
type AccessContextManagerEgressPolicyArgs struct {
	// EgressPolicyName: string, required
	EgressPolicyName terra.StringValue `hcl:"egress_policy_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesscontextmanageregresspolicy.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerEgressPolicyAttributes struct {
	ref terra.Reference
}

// EgressPolicyName returns a reference to field egress_policy_name of google_access_context_manager_egress_policy.
func (acmep accessContextManagerEgressPolicyAttributes) EgressPolicyName() terra.StringValue {
	return terra.ReferenceAsString(acmep.ref.Append("egress_policy_name"))
}

// Id returns a reference to field id of google_access_context_manager_egress_policy.
func (acmep accessContextManagerEgressPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmep.ref.Append("id"))
}

// Resource returns a reference to field resource of google_access_context_manager_egress_policy.
func (acmep accessContextManagerEgressPolicyAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(acmep.ref.Append("resource"))
}

func (acmep accessContextManagerEgressPolicyAttributes) Timeouts() accesscontextmanageregresspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanageregresspolicy.TimeoutsAttributes](acmep.ref.Append("timeouts"))
}

type accessContextManagerEgressPolicyState struct {
	EgressPolicyName string                                          `json:"egress_policy_name"`
	Id               string                                          `json:"id"`
	Resource         string                                          `json:"resource"`
	Timeouts         *accesscontextmanageregresspolicy.TimeoutsState `json:"timeouts"`
}
