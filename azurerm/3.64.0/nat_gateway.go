// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	natgateway "github.com/golingon/terraproviders/azurerm/3.64.0/natgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNatGateway creates a new instance of [NatGateway].
func NewNatGateway(name string, args NatGatewayArgs) *NatGateway {
	return &NatGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NatGateway)(nil)

// NatGateway represents the Terraform resource azurerm_nat_gateway.
type NatGateway struct {
	Name      string
	Args      NatGatewayArgs
	state     *natGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NatGateway].
func (ng *NatGateway) Type() string {
	return "azurerm_nat_gateway"
}

// LocalName returns the local name for [NatGateway].
func (ng *NatGateway) LocalName() string {
	return ng.Name
}

// Configuration returns the configuration (args) for [NatGateway].
func (ng *NatGateway) Configuration() interface{} {
	return ng.Args
}

// DependOn is used for other resources to depend on [NatGateway].
func (ng *NatGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(ng)
}

// Dependencies returns the list of resources [NatGateway] depends_on.
func (ng *NatGateway) Dependencies() terra.Dependencies {
	return ng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NatGateway].
func (ng *NatGateway) LifecycleManagement() *terra.Lifecycle {
	return ng.Lifecycle
}

// Attributes returns the attributes for [NatGateway].
func (ng *NatGateway) Attributes() natGatewayAttributes {
	return natGatewayAttributes{ref: terra.ReferenceResource(ng)}
}

// ImportState imports the given attribute values into [NatGateway]'s state.
func (ng *NatGateway) ImportState(av io.Reader) error {
	ng.state = &natGatewayState{}
	if err := json.NewDecoder(av).Decode(ng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ng.Type(), ng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NatGateway] has state.
func (ng *NatGateway) State() (*natGatewayState, bool) {
	return ng.state, ng.state != nil
}

// StateMust returns the state for [NatGateway]. Panics if the state is nil.
func (ng *NatGateway) StateMust() *natGatewayState {
	if ng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ng.Type(), ng.LocalName()))
	}
	return ng.state
}

// NatGatewayArgs contains the configurations for azurerm_nat_gateway.
type NatGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleTimeoutInMinutes: number, optional
	IdleTimeoutInMinutes terra.NumberValue `hcl:"idle_timeout_in_minutes,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, optional
	SkuName terra.StringValue `hcl:"sku_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Timeouts: optional
	Timeouts *natgateway.Timeouts `hcl:"timeouts,block"`
}
type natGatewayAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_nat_gateway.
func (ng natGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurerm_nat_gateway.
func (ng natGatewayAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(ng.ref.Append("idle_timeout_in_minutes"))
}

// Location returns a reference to field location of azurerm_nat_gateway.
func (ng natGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_nat_gateway.
func (ng natGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_nat_gateway.
func (ng natGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("resource_group_name"))
}

// ResourceGuid returns a reference to field resource_guid of azurerm_nat_gateway.
func (ng natGatewayAttributes) ResourceGuid() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("resource_guid"))
}

// SkuName returns a reference to field sku_name of azurerm_nat_gateway.
func (ng natGatewayAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_nat_gateway.
func (ng natGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ng.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_nat_gateway.
func (ng natGatewayAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ng.ref.Append("zones"))
}

func (ng natGatewayAttributes) Timeouts() natgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[natgateway.TimeoutsAttributes](ng.ref.Append("timeouts"))
}

type natGatewayState struct {
	Id                   string                    `json:"id"`
	IdleTimeoutInMinutes float64                   `json:"idle_timeout_in_minutes"`
	Location             string                    `json:"location"`
	Name                 string                    `json:"name"`
	ResourceGroupName    string                    `json:"resource_group_name"`
	ResourceGuid         string                    `json:"resource_guid"`
	SkuName              string                    `json:"sku_name"`
	Tags                 map[string]string         `json:"tags"`
	Zones                []string                  `json:"zones"`
	Timeouts             *natgateway.TimeoutsState `json:"timeouts"`
}
