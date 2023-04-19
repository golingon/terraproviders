// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualhub "github.com/golingon/terraproviders/azurerm/3.52.0/virtualhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualHub creates a new instance of [VirtualHub].
func NewVirtualHub(name string, args VirtualHubArgs) *VirtualHub {
	return &VirtualHub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualHub)(nil)

// VirtualHub represents the Terraform resource azurerm_virtual_hub.
type VirtualHub struct {
	Name      string
	Args      VirtualHubArgs
	state     *virtualHubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualHub].
func (vh *VirtualHub) Type() string {
	return "azurerm_virtual_hub"
}

// LocalName returns the local name for [VirtualHub].
func (vh *VirtualHub) LocalName() string {
	return vh.Name
}

// Configuration returns the configuration (args) for [VirtualHub].
func (vh *VirtualHub) Configuration() interface{} {
	return vh.Args
}

// DependOn is used for other resources to depend on [VirtualHub].
func (vh *VirtualHub) DependOn() terra.Reference {
	return terra.ReferenceResource(vh)
}

// Dependencies returns the list of resources [VirtualHub] depends_on.
func (vh *VirtualHub) Dependencies() terra.Dependencies {
	return vh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualHub].
func (vh *VirtualHub) LifecycleManagement() *terra.Lifecycle {
	return vh.Lifecycle
}

// Attributes returns the attributes for [VirtualHub].
func (vh *VirtualHub) Attributes() virtualHubAttributes {
	return virtualHubAttributes{ref: terra.ReferenceResource(vh)}
}

// ImportState imports the given attribute values into [VirtualHub]'s state.
func (vh *VirtualHub) ImportState(av io.Reader) error {
	vh.state = &virtualHubState{}
	if err := json.NewDecoder(av).Decode(vh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vh.Type(), vh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualHub] has state.
func (vh *VirtualHub) State() (*virtualHubState, bool) {
	return vh.state, vh.state != nil
}

// StateMust returns the state for [VirtualHub]. Panics if the state is nil.
func (vh *VirtualHub) StateMust() *virtualHubState {
	if vh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vh.Type(), vh.LocalName()))
	}
	return vh.state
}

// VirtualHubArgs contains the configurations for azurerm_virtual_hub.
type VirtualHubArgs struct {
	// AddressPrefix: string, optional
	AddressPrefix terra.StringValue `hcl:"address_prefix,attr"`
	// HubRoutingPreference: string, optional
	HubRoutingPreference terra.StringValue `hcl:"hub_routing_preference,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualWanId: string, optional
	VirtualWanId terra.StringValue `hcl:"virtual_wan_id,attr"`
	// Route: min=0
	Route []virtualhub.Route `hcl:"route,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *virtualhub.Timeouts `hcl:"timeouts,block"`
}
type virtualHubAttributes struct {
	ref terra.Reference
}

// AddressPrefix returns a reference to field address_prefix of azurerm_virtual_hub.
func (vh virtualHubAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("address_prefix"))
}

// DefaultRouteTableId returns a reference to field default_route_table_id of azurerm_virtual_hub.
func (vh virtualHubAttributes) DefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("default_route_table_id"))
}

// HubRoutingPreference returns a reference to field hub_routing_preference of azurerm_virtual_hub.
func (vh virtualHubAttributes) HubRoutingPreference() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("hub_routing_preference"))
}

// Id returns a reference to field id of azurerm_virtual_hub.
func (vh virtualHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_hub.
func (vh virtualHubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_hub.
func (vh virtualHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_hub.
func (vh virtualHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_virtual_hub.
func (vh virtualHubAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_virtual_hub.
func (vh virtualHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vh.ref.Append("tags"))
}

// VirtualRouterAsn returns a reference to field virtual_router_asn of azurerm_virtual_hub.
func (vh virtualHubAttributes) VirtualRouterAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(vh.ref.Append("virtual_router_asn"))
}

// VirtualRouterIps returns a reference to field virtual_router_ips of azurerm_virtual_hub.
func (vh virtualHubAttributes) VirtualRouterIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vh.ref.Append("virtual_router_ips"))
}

// VirtualWanId returns a reference to field virtual_wan_id of azurerm_virtual_hub.
func (vh virtualHubAttributes) VirtualWanId() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("virtual_wan_id"))
}

func (vh virtualHubAttributes) Route() terra.SetValue[virtualhub.RouteAttributes] {
	return terra.ReferenceAsSet[virtualhub.RouteAttributes](vh.ref.Append("route"))
}

func (vh virtualHubAttributes) Timeouts() virtualhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualhub.TimeoutsAttributes](vh.ref.Append("timeouts"))
}

type virtualHubState struct {
	AddressPrefix        string                    `json:"address_prefix"`
	DefaultRouteTableId  string                    `json:"default_route_table_id"`
	HubRoutingPreference string                    `json:"hub_routing_preference"`
	Id                   string                    `json:"id"`
	Location             string                    `json:"location"`
	Name                 string                    `json:"name"`
	ResourceGroupName    string                    `json:"resource_group_name"`
	Sku                  string                    `json:"sku"`
	Tags                 map[string]string         `json:"tags"`
	VirtualRouterAsn     float64                   `json:"virtual_router_asn"`
	VirtualRouterIps     []string                  `json:"virtual_router_ips"`
	VirtualWanId         string                    `json:"virtual_wan_id"`
	Route                []virtualhub.RouteState   `json:"route"`
	Timeouts             *virtualhub.TimeoutsState `json:"timeouts"`
}
