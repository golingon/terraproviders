// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package expressrouteconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Routing struct {
	// AssociatedRouteTableId: string, optional
	AssociatedRouteTableId terra.StringValue `hcl:"associated_route_table_id,attr"`
	// InboundRouteMapId: string, optional
	InboundRouteMapId terra.StringValue `hcl:"inbound_route_map_id,attr"`
	// OutboundRouteMapId: string, optional
	OutboundRouteMapId terra.StringValue `hcl:"outbound_route_map_id,attr"`
	// PropagatedRouteTable: optional
	PropagatedRouteTable *PropagatedRouteTable `hcl:"propagated_route_table,block"`
}

type PropagatedRouteTable struct {
	// Labels: set of string, optional
	Labels terra.SetValue[terra.StringValue] `hcl:"labels,attr"`
	// RouteTableIds: list of string, optional
	RouteTableIds terra.ListValue[terra.StringValue] `hcl:"route_table_ids,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RoutingAttributes struct {
	ref terra.Reference
}

func (r RoutingAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RoutingAttributes) InternalWithRef(ref terra.Reference) RoutingAttributes {
	return RoutingAttributes{ref: ref}
}

func (r RoutingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RoutingAttributes) AssociatedRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("associated_route_table_id"))
}

func (r RoutingAttributes) InboundRouteMapId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("inbound_route_map_id"))
}

func (r RoutingAttributes) OutboundRouteMapId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("outbound_route_map_id"))
}

func (r RoutingAttributes) PropagatedRouteTable() terra.ListValue[PropagatedRouteTableAttributes] {
	return terra.ReferenceAsList[PropagatedRouteTableAttributes](r.ref.Append("propagated_route_table"))
}

type PropagatedRouteTableAttributes struct {
	ref terra.Reference
}

func (prt PropagatedRouteTableAttributes) InternalRef() (terra.Reference, error) {
	return prt.ref, nil
}

func (prt PropagatedRouteTableAttributes) InternalWithRef(ref terra.Reference) PropagatedRouteTableAttributes {
	return PropagatedRouteTableAttributes{ref: ref}
}

func (prt PropagatedRouteTableAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return prt.ref.InternalTokens()
}

func (prt PropagatedRouteTableAttributes) Labels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](prt.ref.Append("labels"))
}

func (prt PropagatedRouteTableAttributes) RouteTableIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](prt.ref.Append("route_table_ids"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type RoutingState struct {
	AssociatedRouteTableId string                      `json:"associated_route_table_id"`
	InboundRouteMapId      string                      `json:"inbound_route_map_id"`
	OutboundRouteMapId     string                      `json:"outbound_route_map_id"`
	PropagatedRouteTable   []PropagatedRouteTableState `json:"propagated_route_table"`
}

type PropagatedRouteTableState struct {
	Labels        []string `json:"labels"`
	RouteTableIds []string `json:"route_table_ids"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}