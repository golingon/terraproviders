// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressrouteport "github.com/golingon/terraproviders/azurerm/3.69.0/expressrouteport"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRoutePort creates a new instance of [ExpressRoutePort].
func NewExpressRoutePort(name string, args ExpressRoutePortArgs) *ExpressRoutePort {
	return &ExpressRoutePort{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRoutePort)(nil)

// ExpressRoutePort represents the Terraform resource azurerm_express_route_port.
type ExpressRoutePort struct {
	Name      string
	Args      ExpressRoutePortArgs
	state     *expressRoutePortState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRoutePort].
func (erp *ExpressRoutePort) Type() string {
	return "azurerm_express_route_port"
}

// LocalName returns the local name for [ExpressRoutePort].
func (erp *ExpressRoutePort) LocalName() string {
	return erp.Name
}

// Configuration returns the configuration (args) for [ExpressRoutePort].
func (erp *ExpressRoutePort) Configuration() interface{} {
	return erp.Args
}

// DependOn is used for other resources to depend on [ExpressRoutePort].
func (erp *ExpressRoutePort) DependOn() terra.Reference {
	return terra.ReferenceResource(erp)
}

// Dependencies returns the list of resources [ExpressRoutePort] depends_on.
func (erp *ExpressRoutePort) Dependencies() terra.Dependencies {
	return erp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRoutePort].
func (erp *ExpressRoutePort) LifecycleManagement() *terra.Lifecycle {
	return erp.Lifecycle
}

// Attributes returns the attributes for [ExpressRoutePort].
func (erp *ExpressRoutePort) Attributes() expressRoutePortAttributes {
	return expressRoutePortAttributes{ref: terra.ReferenceResource(erp)}
}

// ImportState imports the given attribute values into [ExpressRoutePort]'s state.
func (erp *ExpressRoutePort) ImportState(av io.Reader) error {
	erp.state = &expressRoutePortState{}
	if err := json.NewDecoder(av).Decode(erp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erp.Type(), erp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRoutePort] has state.
func (erp *ExpressRoutePort) State() (*expressRoutePortState, bool) {
	return erp.state, erp.state != nil
}

// StateMust returns the state for [ExpressRoutePort]. Panics if the state is nil.
func (erp *ExpressRoutePort) StateMust() *expressRoutePortState {
	if erp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erp.Type(), erp.LocalName()))
	}
	return erp.state
}

// ExpressRoutePortArgs contains the configurations for azurerm_express_route_port.
type ExpressRoutePortArgs struct {
	// BandwidthInGbps: number, required
	BandwidthInGbps terra.NumberValue `hcl:"bandwidth_in_gbps,attr" validate:"required"`
	// BillingType: string, optional
	BillingType terra.StringValue `hcl:"billing_type,attr"`
	// Encapsulation: string, required
	Encapsulation terra.StringValue `hcl:"encapsulation,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeeringLocation: string, required
	PeeringLocation terra.StringValue `hcl:"peering_location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *expressrouteport.Identity `hcl:"identity,block"`
	// Link1: optional
	Link1 *expressrouteport.Link1 `hcl:"link1,block"`
	// Link2: optional
	Link2 *expressrouteport.Link2 `hcl:"link2,block"`
	// Timeouts: optional
	Timeouts *expressrouteport.Timeouts `hcl:"timeouts,block"`
}
type expressRoutePortAttributes struct {
	ref terra.Reference
}

// BandwidthInGbps returns a reference to field bandwidth_in_gbps of azurerm_express_route_port.
func (erp expressRoutePortAttributes) BandwidthInGbps() terra.NumberValue {
	return terra.ReferenceAsNumber(erp.ref.Append("bandwidth_in_gbps"))
}

// BillingType returns a reference to field billing_type of azurerm_express_route_port.
func (erp expressRoutePortAttributes) BillingType() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("billing_type"))
}

// Encapsulation returns a reference to field encapsulation of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Encapsulation() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("encapsulation"))
}

// Ethertype returns a reference to field ethertype of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Ethertype() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("ethertype"))
}

// Guid returns a reference to field guid of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Guid() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("guid"))
}

// Id returns a reference to field id of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("location"))
}

// Mtu returns a reference to field mtu of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Mtu() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("mtu"))
}

// Name returns a reference to field name of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("name"))
}

// PeeringLocation returns a reference to field peering_location of azurerm_express_route_port.
func (erp expressRoutePortAttributes) PeeringLocation() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("peering_location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_express_route_port.
func (erp expressRoutePortAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(erp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_express_route_port.
func (erp expressRoutePortAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](erp.ref.Append("tags"))
}

func (erp expressRoutePortAttributes) Identity() terra.ListValue[expressrouteport.IdentityAttributes] {
	return terra.ReferenceAsList[expressrouteport.IdentityAttributes](erp.ref.Append("identity"))
}

func (erp expressRoutePortAttributes) Link1() terra.ListValue[expressrouteport.Link1Attributes] {
	return terra.ReferenceAsList[expressrouteport.Link1Attributes](erp.ref.Append("link1"))
}

func (erp expressRoutePortAttributes) Link2() terra.ListValue[expressrouteport.Link2Attributes] {
	return terra.ReferenceAsList[expressrouteport.Link2Attributes](erp.ref.Append("link2"))
}

func (erp expressRoutePortAttributes) Timeouts() expressrouteport.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressrouteport.TimeoutsAttributes](erp.ref.Append("timeouts"))
}

type expressRoutePortState struct {
	BandwidthInGbps   float64                          `json:"bandwidth_in_gbps"`
	BillingType       string                           `json:"billing_type"`
	Encapsulation     string                           `json:"encapsulation"`
	Ethertype         string                           `json:"ethertype"`
	Guid              string                           `json:"guid"`
	Id                string                           `json:"id"`
	Location          string                           `json:"location"`
	Mtu               string                           `json:"mtu"`
	Name              string                           `json:"name"`
	PeeringLocation   string                           `json:"peering_location"`
	ResourceGroupName string                           `json:"resource_group_name"`
	Tags              map[string]string                `json:"tags"`
	Identity          []expressrouteport.IdentityState `json:"identity"`
	Link1             []expressrouteport.Link1State    `json:"link1"`
	Link2             []expressrouteport.Link2State    `json:"link2"`
	Timeouts          *expressrouteport.TimeoutsState  `json:"timeouts"`
}
