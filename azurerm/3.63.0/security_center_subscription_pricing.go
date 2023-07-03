// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycentersubscriptionpricing "github.com/golingon/terraproviders/azurerm/3.63.0/securitycentersubscriptionpricing"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterSubscriptionPricing creates a new instance of [SecurityCenterSubscriptionPricing].
func NewSecurityCenterSubscriptionPricing(name string, args SecurityCenterSubscriptionPricingArgs) *SecurityCenterSubscriptionPricing {
	return &SecurityCenterSubscriptionPricing{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterSubscriptionPricing)(nil)

// SecurityCenterSubscriptionPricing represents the Terraform resource azurerm_security_center_subscription_pricing.
type SecurityCenterSubscriptionPricing struct {
	Name      string
	Args      SecurityCenterSubscriptionPricingArgs
	state     *securityCenterSubscriptionPricingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterSubscriptionPricing].
func (scsp *SecurityCenterSubscriptionPricing) Type() string {
	return "azurerm_security_center_subscription_pricing"
}

// LocalName returns the local name for [SecurityCenterSubscriptionPricing].
func (scsp *SecurityCenterSubscriptionPricing) LocalName() string {
	return scsp.Name
}

// Configuration returns the configuration (args) for [SecurityCenterSubscriptionPricing].
func (scsp *SecurityCenterSubscriptionPricing) Configuration() interface{} {
	return scsp.Args
}

// DependOn is used for other resources to depend on [SecurityCenterSubscriptionPricing].
func (scsp *SecurityCenterSubscriptionPricing) DependOn() terra.Reference {
	return terra.ReferenceResource(scsp)
}

// Dependencies returns the list of resources [SecurityCenterSubscriptionPricing] depends_on.
func (scsp *SecurityCenterSubscriptionPricing) Dependencies() terra.Dependencies {
	return scsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterSubscriptionPricing].
func (scsp *SecurityCenterSubscriptionPricing) LifecycleManagement() *terra.Lifecycle {
	return scsp.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterSubscriptionPricing].
func (scsp *SecurityCenterSubscriptionPricing) Attributes() securityCenterSubscriptionPricingAttributes {
	return securityCenterSubscriptionPricingAttributes{ref: terra.ReferenceResource(scsp)}
}

// ImportState imports the given attribute values into [SecurityCenterSubscriptionPricing]'s state.
func (scsp *SecurityCenterSubscriptionPricing) ImportState(av io.Reader) error {
	scsp.state = &securityCenterSubscriptionPricingState{}
	if err := json.NewDecoder(av).Decode(scsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scsp.Type(), scsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterSubscriptionPricing] has state.
func (scsp *SecurityCenterSubscriptionPricing) State() (*securityCenterSubscriptionPricingState, bool) {
	return scsp.state, scsp.state != nil
}

// StateMust returns the state for [SecurityCenterSubscriptionPricing]. Panics if the state is nil.
func (scsp *SecurityCenterSubscriptionPricing) StateMust() *securityCenterSubscriptionPricingState {
	if scsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scsp.Type(), scsp.LocalName()))
	}
	return scsp.state
}

// SecurityCenterSubscriptionPricingArgs contains the configurations for azurerm_security_center_subscription_pricing.
type SecurityCenterSubscriptionPricingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceType: string, optional
	ResourceType terra.StringValue `hcl:"resource_type,attr"`
	// Subplan: string, optional
	Subplan terra.StringValue `hcl:"subplan,attr"`
	// Tier: string, required
	Tier terra.StringValue `hcl:"tier,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *securitycentersubscriptionpricing.Timeouts `hcl:"timeouts,block"`
}
type securityCenterSubscriptionPricingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_security_center_subscription_pricing.
func (scsp securityCenterSubscriptionPricingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scsp.ref.Append("id"))
}

// ResourceType returns a reference to field resource_type of azurerm_security_center_subscription_pricing.
func (scsp securityCenterSubscriptionPricingAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(scsp.ref.Append("resource_type"))
}

// Subplan returns a reference to field subplan of azurerm_security_center_subscription_pricing.
func (scsp securityCenterSubscriptionPricingAttributes) Subplan() terra.StringValue {
	return terra.ReferenceAsString(scsp.ref.Append("subplan"))
}

// Tier returns a reference to field tier of azurerm_security_center_subscription_pricing.
func (scsp securityCenterSubscriptionPricingAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(scsp.ref.Append("tier"))
}

func (scsp securityCenterSubscriptionPricingAttributes) Timeouts() securitycentersubscriptionpricing.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycentersubscriptionpricing.TimeoutsAttributes](scsp.ref.Append("timeouts"))
}

type securityCenterSubscriptionPricingState struct {
	Id           string                                           `json:"id"`
	ResourceType string                                           `json:"resource_type"`
	Subplan      string                                           `json:"subplan"`
	Tier         string                                           `json:"tier"`
	Timeouts     *securitycentersubscriptionpricing.TimeoutsState `json:"timeouts"`
}