// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	accesscontextmanageringresspolicy "github.com/golingon/terraproviders/googlebeta/4.73.1/accesscontextmanageringresspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerIngressPolicy creates a new instance of [AccessContextManagerIngressPolicy].
func NewAccessContextManagerIngressPolicy(name string, args AccessContextManagerIngressPolicyArgs) *AccessContextManagerIngressPolicy {
	return &AccessContextManagerIngressPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerIngressPolicy)(nil)

// AccessContextManagerIngressPolicy represents the Terraform resource google_access_context_manager_ingress_policy.
type AccessContextManagerIngressPolicy struct {
	Name      string
	Args      AccessContextManagerIngressPolicyArgs
	state     *accessContextManagerIngressPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerIngressPolicy].
func (acmip *AccessContextManagerIngressPolicy) Type() string {
	return "google_access_context_manager_ingress_policy"
}

// LocalName returns the local name for [AccessContextManagerIngressPolicy].
func (acmip *AccessContextManagerIngressPolicy) LocalName() string {
	return acmip.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerIngressPolicy].
func (acmip *AccessContextManagerIngressPolicy) Configuration() interface{} {
	return acmip.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerIngressPolicy].
func (acmip *AccessContextManagerIngressPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(acmip)
}

// Dependencies returns the list of resources [AccessContextManagerIngressPolicy] depends_on.
func (acmip *AccessContextManagerIngressPolicy) Dependencies() terra.Dependencies {
	return acmip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerIngressPolicy].
func (acmip *AccessContextManagerIngressPolicy) LifecycleManagement() *terra.Lifecycle {
	return acmip.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerIngressPolicy].
func (acmip *AccessContextManagerIngressPolicy) Attributes() accessContextManagerIngressPolicyAttributes {
	return accessContextManagerIngressPolicyAttributes{ref: terra.ReferenceResource(acmip)}
}

// ImportState imports the given attribute values into [AccessContextManagerIngressPolicy]'s state.
func (acmip *AccessContextManagerIngressPolicy) ImportState(av io.Reader) error {
	acmip.state = &accessContextManagerIngressPolicyState{}
	if err := json.NewDecoder(av).Decode(acmip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmip.Type(), acmip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerIngressPolicy] has state.
func (acmip *AccessContextManagerIngressPolicy) State() (*accessContextManagerIngressPolicyState, bool) {
	return acmip.state, acmip.state != nil
}

// StateMust returns the state for [AccessContextManagerIngressPolicy]. Panics if the state is nil.
func (acmip *AccessContextManagerIngressPolicy) StateMust() *accessContextManagerIngressPolicyState {
	if acmip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmip.Type(), acmip.LocalName()))
	}
	return acmip.state
}

// AccessContextManagerIngressPolicyArgs contains the configurations for google_access_context_manager_ingress_policy.
type AccessContextManagerIngressPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IngressPolicyName: string, required
	IngressPolicyName terra.StringValue `hcl:"ingress_policy_name,attr" validate:"required"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesscontextmanageringresspolicy.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerIngressPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_access_context_manager_ingress_policy.
func (acmip accessContextManagerIngressPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmip.ref.Append("id"))
}

// IngressPolicyName returns a reference to field ingress_policy_name of google_access_context_manager_ingress_policy.
func (acmip accessContextManagerIngressPolicyAttributes) IngressPolicyName() terra.StringValue {
	return terra.ReferenceAsString(acmip.ref.Append("ingress_policy_name"))
}

// Resource returns a reference to field resource of google_access_context_manager_ingress_policy.
func (acmip accessContextManagerIngressPolicyAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(acmip.ref.Append("resource"))
}

func (acmip accessContextManagerIngressPolicyAttributes) Timeouts() accesscontextmanageringresspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanageringresspolicy.TimeoutsAttributes](acmip.ref.Append("timeouts"))
}

type accessContextManagerIngressPolicyState struct {
	Id                string                                           `json:"id"`
	IngressPolicyName string                                           `json:"ingress_policy_name"`
	Resource          string                                           `json:"resource"`
	Timeouts          *accesscontextmanageringresspolicy.TimeoutsState `json:"timeouts"`
}
