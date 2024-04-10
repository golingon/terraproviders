// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	applicationloadbalancerfrontend "github.com/golingon/terraproviders/azurerm/3.98.0/applicationloadbalancerfrontend"
	"io"
)

// NewApplicationLoadBalancerFrontend creates a new instance of [ApplicationLoadBalancerFrontend].
func NewApplicationLoadBalancerFrontend(name string, args ApplicationLoadBalancerFrontendArgs) *ApplicationLoadBalancerFrontend {
	return &ApplicationLoadBalancerFrontend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationLoadBalancerFrontend)(nil)

// ApplicationLoadBalancerFrontend represents the Terraform resource azurerm_application_load_balancer_frontend.
type ApplicationLoadBalancerFrontend struct {
	Name      string
	Args      ApplicationLoadBalancerFrontendArgs
	state     *applicationLoadBalancerFrontendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationLoadBalancerFrontend].
func (albf *ApplicationLoadBalancerFrontend) Type() string {
	return "azurerm_application_load_balancer_frontend"
}

// LocalName returns the local name for [ApplicationLoadBalancerFrontend].
func (albf *ApplicationLoadBalancerFrontend) LocalName() string {
	return albf.Name
}

// Configuration returns the configuration (args) for [ApplicationLoadBalancerFrontend].
func (albf *ApplicationLoadBalancerFrontend) Configuration() interface{} {
	return albf.Args
}

// DependOn is used for other resources to depend on [ApplicationLoadBalancerFrontend].
func (albf *ApplicationLoadBalancerFrontend) DependOn() terra.Reference {
	return terra.ReferenceResource(albf)
}

// Dependencies returns the list of resources [ApplicationLoadBalancerFrontend] depends_on.
func (albf *ApplicationLoadBalancerFrontend) Dependencies() terra.Dependencies {
	return albf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationLoadBalancerFrontend].
func (albf *ApplicationLoadBalancerFrontend) LifecycleManagement() *terra.Lifecycle {
	return albf.Lifecycle
}

// Attributes returns the attributes for [ApplicationLoadBalancerFrontend].
func (albf *ApplicationLoadBalancerFrontend) Attributes() applicationLoadBalancerFrontendAttributes {
	return applicationLoadBalancerFrontendAttributes{ref: terra.ReferenceResource(albf)}
}

// ImportState imports the given attribute values into [ApplicationLoadBalancerFrontend]'s state.
func (albf *ApplicationLoadBalancerFrontend) ImportState(av io.Reader) error {
	albf.state = &applicationLoadBalancerFrontendState{}
	if err := json.NewDecoder(av).Decode(albf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", albf.Type(), albf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationLoadBalancerFrontend] has state.
func (albf *ApplicationLoadBalancerFrontend) State() (*applicationLoadBalancerFrontendState, bool) {
	return albf.state, albf.state != nil
}

// StateMust returns the state for [ApplicationLoadBalancerFrontend]. Panics if the state is nil.
func (albf *ApplicationLoadBalancerFrontend) StateMust() *applicationLoadBalancerFrontendState {
	if albf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", albf.Type(), albf.LocalName()))
	}
	return albf.state
}

// ApplicationLoadBalancerFrontendArgs contains the configurations for azurerm_application_load_balancer_frontend.
type ApplicationLoadBalancerFrontendArgs struct {
	// ApplicationLoadBalancerId: string, required
	ApplicationLoadBalancerId terra.StringValue `hcl:"application_load_balancer_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *applicationloadbalancerfrontend.Timeouts `hcl:"timeouts,block"`
}
type applicationLoadBalancerFrontendAttributes struct {
	ref terra.Reference
}

// ApplicationLoadBalancerId returns a reference to field application_load_balancer_id of azurerm_application_load_balancer_frontend.
func (albf applicationLoadBalancerFrontendAttributes) ApplicationLoadBalancerId() terra.StringValue {
	return terra.ReferenceAsString(albf.ref.Append("application_load_balancer_id"))
}

// FullyQualifiedDomainName returns a reference to field fully_qualified_domain_name of azurerm_application_load_balancer_frontend.
func (albf applicationLoadBalancerFrontendAttributes) FullyQualifiedDomainName() terra.StringValue {
	return terra.ReferenceAsString(albf.ref.Append("fully_qualified_domain_name"))
}

// Id returns a reference to field id of azurerm_application_load_balancer_frontend.
func (albf applicationLoadBalancerFrontendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(albf.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_application_load_balancer_frontend.
func (albf applicationLoadBalancerFrontendAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(albf.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_application_load_balancer_frontend.
func (albf applicationLoadBalancerFrontendAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](albf.ref.Append("tags"))
}

func (albf applicationLoadBalancerFrontendAttributes) Timeouts() applicationloadbalancerfrontend.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationloadbalancerfrontend.TimeoutsAttributes](albf.ref.Append("timeouts"))
}

type applicationLoadBalancerFrontendState struct {
	ApplicationLoadBalancerId string                                         `json:"application_load_balancer_id"`
	FullyQualifiedDomainName  string                                         `json:"fully_qualified_domain_name"`
	Id                        string                                         `json:"id"`
	Name                      string                                         `json:"name"`
	Tags                      map[string]string                              `json:"tags"`
	Timeouts                  *applicationloadbalancerfrontend.TimeoutsState `json:"timeouts"`
}
