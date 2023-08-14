// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	accesscontextmanagerserviceperimeteregresspolicy "github.com/golingon/terraproviders/googlebeta/4.77.0/accesscontextmanagerserviceperimeteregresspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerServicePerimeterEgressPolicy creates a new instance of [AccessContextManagerServicePerimeterEgressPolicy].
func NewAccessContextManagerServicePerimeterEgressPolicy(name string, args AccessContextManagerServicePerimeterEgressPolicyArgs) *AccessContextManagerServicePerimeterEgressPolicy {
	return &AccessContextManagerServicePerimeterEgressPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerServicePerimeterEgressPolicy)(nil)

// AccessContextManagerServicePerimeterEgressPolicy represents the Terraform resource google_access_context_manager_service_perimeter_egress_policy.
type AccessContextManagerServicePerimeterEgressPolicy struct {
	Name      string
	Args      AccessContextManagerServicePerimeterEgressPolicyArgs
	state     *accessContextManagerServicePerimeterEgressPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerServicePerimeterEgressPolicy].
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) Type() string {
	return "google_access_context_manager_service_perimeter_egress_policy"
}

// LocalName returns the local name for [AccessContextManagerServicePerimeterEgressPolicy].
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) LocalName() string {
	return acmspep.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerServicePerimeterEgressPolicy].
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) Configuration() interface{} {
	return acmspep.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerServicePerimeterEgressPolicy].
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(acmspep)
}

// Dependencies returns the list of resources [AccessContextManagerServicePerimeterEgressPolicy] depends_on.
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) Dependencies() terra.Dependencies {
	return acmspep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerServicePerimeterEgressPolicy].
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) LifecycleManagement() *terra.Lifecycle {
	return acmspep.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerServicePerimeterEgressPolicy].
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) Attributes() accessContextManagerServicePerimeterEgressPolicyAttributes {
	return accessContextManagerServicePerimeterEgressPolicyAttributes{ref: terra.ReferenceResource(acmspep)}
}

// ImportState imports the given attribute values into [AccessContextManagerServicePerimeterEgressPolicy]'s state.
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) ImportState(av io.Reader) error {
	acmspep.state = &accessContextManagerServicePerimeterEgressPolicyState{}
	if err := json.NewDecoder(av).Decode(acmspep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmspep.Type(), acmspep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerServicePerimeterEgressPolicy] has state.
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) State() (*accessContextManagerServicePerimeterEgressPolicyState, bool) {
	return acmspep.state, acmspep.state != nil
}

// StateMust returns the state for [AccessContextManagerServicePerimeterEgressPolicy]. Panics if the state is nil.
func (acmspep *AccessContextManagerServicePerimeterEgressPolicy) StateMust() *accessContextManagerServicePerimeterEgressPolicyState {
	if acmspep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmspep.Type(), acmspep.LocalName()))
	}
	return acmspep.state
}

// AccessContextManagerServicePerimeterEgressPolicyArgs contains the configurations for google_access_context_manager_service_perimeter_egress_policy.
type AccessContextManagerServicePerimeterEgressPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Perimeter: string, required
	Perimeter terra.StringValue `hcl:"perimeter,attr" validate:"required"`
	// EgressFrom: optional
	EgressFrom *accesscontextmanagerserviceperimeteregresspolicy.EgressFrom `hcl:"egress_from,block"`
	// EgressTo: optional
	EgressTo *accesscontextmanagerserviceperimeteregresspolicy.EgressTo `hcl:"egress_to,block"`
	// Timeouts: optional
	Timeouts *accesscontextmanagerserviceperimeteregresspolicy.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerServicePerimeterEgressPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_access_context_manager_service_perimeter_egress_policy.
func (acmspep accessContextManagerServicePerimeterEgressPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmspep.ref.Append("id"))
}

// Perimeter returns a reference to field perimeter of google_access_context_manager_service_perimeter_egress_policy.
func (acmspep accessContextManagerServicePerimeterEgressPolicyAttributes) Perimeter() terra.StringValue {
	return terra.ReferenceAsString(acmspep.ref.Append("perimeter"))
}

func (acmspep accessContextManagerServicePerimeterEgressPolicyAttributes) EgressFrom() terra.ListValue[accesscontextmanagerserviceperimeteregresspolicy.EgressFromAttributes] {
	return terra.ReferenceAsList[accesscontextmanagerserviceperimeteregresspolicy.EgressFromAttributes](acmspep.ref.Append("egress_from"))
}

func (acmspep accessContextManagerServicePerimeterEgressPolicyAttributes) EgressTo() terra.ListValue[accesscontextmanagerserviceperimeteregresspolicy.EgressToAttributes] {
	return terra.ReferenceAsList[accesscontextmanagerserviceperimeteregresspolicy.EgressToAttributes](acmspep.ref.Append("egress_to"))
}

func (acmspep accessContextManagerServicePerimeterEgressPolicyAttributes) Timeouts() accesscontextmanagerserviceperimeteregresspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagerserviceperimeteregresspolicy.TimeoutsAttributes](acmspep.ref.Append("timeouts"))
}

type accessContextManagerServicePerimeterEgressPolicyState struct {
	Id         string                                                             `json:"id"`
	Perimeter  string                                                             `json:"perimeter"`
	EgressFrom []accesscontextmanagerserviceperimeteregresspolicy.EgressFromState `json:"egress_from"`
	EgressTo   []accesscontextmanagerserviceperimeteregresspolicy.EgressToState   `json:"egress_to"`
	Timeouts   *accesscontextmanagerserviceperimeteregresspolicy.TimeoutsState    `json:"timeouts"`
}
