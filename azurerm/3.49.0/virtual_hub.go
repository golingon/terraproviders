// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualhub "github.com/golingon/terraproviders/azurerm/3.49.0/virtualhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVirtualHub(name string, args VirtualHubArgs) *VirtualHub {
	return &VirtualHub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualHub)(nil)

type VirtualHub struct {
	Name  string
	Args  VirtualHubArgs
	state *virtualHubState
}

func (vh *VirtualHub) Type() string {
	return "azurerm_virtual_hub"
}

func (vh *VirtualHub) LocalName() string {
	return vh.Name
}

func (vh *VirtualHub) Configuration() interface{} {
	return vh.Args
}

func (vh *VirtualHub) Attributes() virtualHubAttributes {
	return virtualHubAttributes{ref: terra.ReferenceResource(vh)}
}

func (vh *VirtualHub) ImportState(av io.Reader) error {
	vh.state = &virtualHubState{}
	if err := json.NewDecoder(av).Decode(vh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vh.Type(), vh.LocalName(), err)
	}
	return nil
}

func (vh *VirtualHub) State() (*virtualHubState, bool) {
	return vh.state, vh.state != nil
}

func (vh *VirtualHub) StateMust() *virtualHubState {
	if vh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vh.Type(), vh.LocalName()))
	}
	return vh.state
}

func (vh *VirtualHub) DependOn() terra.Reference {
	return terra.ReferenceResource(vh)
}

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
	// DependsOn contains resources that VirtualHub depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type virtualHubAttributes struct {
	ref terra.Reference
}

func (vh virtualHubAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("address_prefix"))
}

func (vh virtualHubAttributes) DefaultRouteTableId() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("default_route_table_id"))
}

func (vh virtualHubAttributes) HubRoutingPreference() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("hub_routing_preference"))
}

func (vh virtualHubAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("id"))
}

func (vh virtualHubAttributes) Location() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("location"))
}

func (vh virtualHubAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("name"))
}

func (vh virtualHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("resource_group_name"))
}

func (vh virtualHubAttributes) Sku() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("sku"))
}

func (vh virtualHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vh.ref.Append("tags"))
}

func (vh virtualHubAttributes) VirtualRouterAsn() terra.NumberValue {
	return terra.ReferenceNumber(vh.ref.Append("virtual_router_asn"))
}

func (vh virtualHubAttributes) VirtualRouterIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vh.ref.Append("virtual_router_ips"))
}

func (vh virtualHubAttributes) VirtualWanId() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("virtual_wan_id"))
}

func (vh virtualHubAttributes) Route() terra.SetValue[virtualhub.RouteAttributes] {
	return terra.ReferenceSet[virtualhub.RouteAttributes](vh.ref.Append("route"))
}

func (vh virtualHubAttributes) Timeouts() virtualhub.TimeoutsAttributes {
	return terra.ReferenceSingle[virtualhub.TimeoutsAttributes](vh.ref.Append("timeouts"))
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
