// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanagerserviceperimeteringresspolicy "github.com/golingon/terraproviders/google/4.73.1/accesscontextmanagerserviceperimeteringresspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerServicePerimeterIngressPolicy creates a new instance of [AccessContextManagerServicePerimeterIngressPolicy].
func NewAccessContextManagerServicePerimeterIngressPolicy(name string, args AccessContextManagerServicePerimeterIngressPolicyArgs) *AccessContextManagerServicePerimeterIngressPolicy {
	return &AccessContextManagerServicePerimeterIngressPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerServicePerimeterIngressPolicy)(nil)

// AccessContextManagerServicePerimeterIngressPolicy represents the Terraform resource google_access_context_manager_service_perimeter_ingress_policy.
type AccessContextManagerServicePerimeterIngressPolicy struct {
	Name      string
	Args      AccessContextManagerServicePerimeterIngressPolicyArgs
	state     *accessContextManagerServicePerimeterIngressPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerServicePerimeterIngressPolicy].
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) Type() string {
	return "google_access_context_manager_service_perimeter_ingress_policy"
}

// LocalName returns the local name for [AccessContextManagerServicePerimeterIngressPolicy].
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) LocalName() string {
	return acmspip.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerServicePerimeterIngressPolicy].
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) Configuration() interface{} {
	return acmspip.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerServicePerimeterIngressPolicy].
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(acmspip)
}

// Dependencies returns the list of resources [AccessContextManagerServicePerimeterIngressPolicy] depends_on.
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) Dependencies() terra.Dependencies {
	return acmspip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerServicePerimeterIngressPolicy].
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) LifecycleManagement() *terra.Lifecycle {
	return acmspip.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerServicePerimeterIngressPolicy].
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) Attributes() accessContextManagerServicePerimeterIngressPolicyAttributes {
	return accessContextManagerServicePerimeterIngressPolicyAttributes{ref: terra.ReferenceResource(acmspip)}
}

// ImportState imports the given attribute values into [AccessContextManagerServicePerimeterIngressPolicy]'s state.
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) ImportState(av io.Reader) error {
	acmspip.state = &accessContextManagerServicePerimeterIngressPolicyState{}
	if err := json.NewDecoder(av).Decode(acmspip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmspip.Type(), acmspip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerServicePerimeterIngressPolicy] has state.
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) State() (*accessContextManagerServicePerimeterIngressPolicyState, bool) {
	return acmspip.state, acmspip.state != nil
}

// StateMust returns the state for [AccessContextManagerServicePerimeterIngressPolicy]. Panics if the state is nil.
func (acmspip *AccessContextManagerServicePerimeterIngressPolicy) StateMust() *accessContextManagerServicePerimeterIngressPolicyState {
	if acmspip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmspip.Type(), acmspip.LocalName()))
	}
	return acmspip.state
}

// AccessContextManagerServicePerimeterIngressPolicyArgs contains the configurations for google_access_context_manager_service_perimeter_ingress_policy.
type AccessContextManagerServicePerimeterIngressPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Perimeter: string, required
	Perimeter terra.StringValue `hcl:"perimeter,attr" validate:"required"`
	// IngressFrom: optional
	IngressFrom *accesscontextmanagerserviceperimeteringresspolicy.IngressFrom `hcl:"ingress_from,block"`
	// IngressTo: optional
	IngressTo *accesscontextmanagerserviceperimeteringresspolicy.IngressTo `hcl:"ingress_to,block"`
	// Timeouts: optional
	Timeouts *accesscontextmanagerserviceperimeteringresspolicy.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerServicePerimeterIngressPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_access_context_manager_service_perimeter_ingress_policy.
func (acmspip accessContextManagerServicePerimeterIngressPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmspip.ref.Append("id"))
}

// Perimeter returns a reference to field perimeter of google_access_context_manager_service_perimeter_ingress_policy.
func (acmspip accessContextManagerServicePerimeterIngressPolicyAttributes) Perimeter() terra.StringValue {
	return terra.ReferenceAsString(acmspip.ref.Append("perimeter"))
}

func (acmspip accessContextManagerServicePerimeterIngressPolicyAttributes) IngressFrom() terra.ListValue[accesscontextmanagerserviceperimeteringresspolicy.IngressFromAttributes] {
	return terra.ReferenceAsList[accesscontextmanagerserviceperimeteringresspolicy.IngressFromAttributes](acmspip.ref.Append("ingress_from"))
}

func (acmspip accessContextManagerServicePerimeterIngressPolicyAttributes) IngressTo() terra.ListValue[accesscontextmanagerserviceperimeteringresspolicy.IngressToAttributes] {
	return terra.ReferenceAsList[accesscontextmanagerserviceperimeteringresspolicy.IngressToAttributes](acmspip.ref.Append("ingress_to"))
}

func (acmspip accessContextManagerServicePerimeterIngressPolicyAttributes) Timeouts() accesscontextmanagerserviceperimeteringresspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagerserviceperimeteringresspolicy.TimeoutsAttributes](acmspip.ref.Append("timeouts"))
}

type accessContextManagerServicePerimeterIngressPolicyState struct {
	Id          string                                                               `json:"id"`
	Perimeter   string                                                               `json:"perimeter"`
	IngressFrom []accesscontextmanagerserviceperimeteringresspolicy.IngressFromState `json:"ingress_from"`
	IngressTo   []accesscontextmanagerserviceperimeteringresspolicy.IngressToState   `json:"ingress_to"`
	Timeouts    *accesscontextmanagerserviceperimeteringresspolicy.TimeoutsState     `json:"timeouts"`
}
