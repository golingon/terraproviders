// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudgatewaycustomdomain "github.com/golingon/terraproviders/azurerm/3.69.0/springcloudgatewaycustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudGatewayCustomDomain creates a new instance of [SpringCloudGatewayCustomDomain].
func NewSpringCloudGatewayCustomDomain(name string, args SpringCloudGatewayCustomDomainArgs) *SpringCloudGatewayCustomDomain {
	return &SpringCloudGatewayCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudGatewayCustomDomain)(nil)

// SpringCloudGatewayCustomDomain represents the Terraform resource azurerm_spring_cloud_gateway_custom_domain.
type SpringCloudGatewayCustomDomain struct {
	Name      string
	Args      SpringCloudGatewayCustomDomainArgs
	state     *springCloudGatewayCustomDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudGatewayCustomDomain].
func (scgcd *SpringCloudGatewayCustomDomain) Type() string {
	return "azurerm_spring_cloud_gateway_custom_domain"
}

// LocalName returns the local name for [SpringCloudGatewayCustomDomain].
func (scgcd *SpringCloudGatewayCustomDomain) LocalName() string {
	return scgcd.Name
}

// Configuration returns the configuration (args) for [SpringCloudGatewayCustomDomain].
func (scgcd *SpringCloudGatewayCustomDomain) Configuration() interface{} {
	return scgcd.Args
}

// DependOn is used for other resources to depend on [SpringCloudGatewayCustomDomain].
func (scgcd *SpringCloudGatewayCustomDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(scgcd)
}

// Dependencies returns the list of resources [SpringCloudGatewayCustomDomain] depends_on.
func (scgcd *SpringCloudGatewayCustomDomain) Dependencies() terra.Dependencies {
	return scgcd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudGatewayCustomDomain].
func (scgcd *SpringCloudGatewayCustomDomain) LifecycleManagement() *terra.Lifecycle {
	return scgcd.Lifecycle
}

// Attributes returns the attributes for [SpringCloudGatewayCustomDomain].
func (scgcd *SpringCloudGatewayCustomDomain) Attributes() springCloudGatewayCustomDomainAttributes {
	return springCloudGatewayCustomDomainAttributes{ref: terra.ReferenceResource(scgcd)}
}

// ImportState imports the given attribute values into [SpringCloudGatewayCustomDomain]'s state.
func (scgcd *SpringCloudGatewayCustomDomain) ImportState(av io.Reader) error {
	scgcd.state = &springCloudGatewayCustomDomainState{}
	if err := json.NewDecoder(av).Decode(scgcd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scgcd.Type(), scgcd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudGatewayCustomDomain] has state.
func (scgcd *SpringCloudGatewayCustomDomain) State() (*springCloudGatewayCustomDomainState, bool) {
	return scgcd.state, scgcd.state != nil
}

// StateMust returns the state for [SpringCloudGatewayCustomDomain]. Panics if the state is nil.
func (scgcd *SpringCloudGatewayCustomDomain) StateMust() *springCloudGatewayCustomDomainState {
	if scgcd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scgcd.Type(), scgcd.LocalName()))
	}
	return scgcd.state
}

// SpringCloudGatewayCustomDomainArgs contains the configurations for azurerm_spring_cloud_gateway_custom_domain.
type SpringCloudGatewayCustomDomainArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudGatewayId: string, required
	SpringCloudGatewayId terra.StringValue `hcl:"spring_cloud_gateway_id,attr" validate:"required"`
	// Thumbprint: string, optional
	Thumbprint terra.StringValue `hcl:"thumbprint,attr"`
	// Timeouts: optional
	Timeouts *springcloudgatewaycustomdomain.Timeouts `hcl:"timeouts,block"`
}
type springCloudGatewayCustomDomainAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_gateway_custom_domain.
func (scgcd springCloudGatewayCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scgcd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_gateway_custom_domain.
func (scgcd springCloudGatewayCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scgcd.ref.Append("name"))
}

// SpringCloudGatewayId returns a reference to field spring_cloud_gateway_id of azurerm_spring_cloud_gateway_custom_domain.
func (scgcd springCloudGatewayCustomDomainAttributes) SpringCloudGatewayId() terra.StringValue {
	return terra.ReferenceAsString(scgcd.ref.Append("spring_cloud_gateway_id"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_spring_cloud_gateway_custom_domain.
func (scgcd springCloudGatewayCustomDomainAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(scgcd.ref.Append("thumbprint"))
}

func (scgcd springCloudGatewayCustomDomainAttributes) Timeouts() springcloudgatewaycustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudgatewaycustomdomain.TimeoutsAttributes](scgcd.ref.Append("timeouts"))
}

type springCloudGatewayCustomDomainState struct {
	Id                   string                                        `json:"id"`
	Name                 string                                        `json:"name"`
	SpringCloudGatewayId string                                        `json:"spring_cloud_gateway_id"`
	Thumbprint           string                                        `json:"thumbprint"`
	Timeouts             *springcloudgatewaycustomdomain.TimeoutsState `json:"timeouts"`
}
