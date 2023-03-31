// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressroutegateway "github.com/golingon/terraproviders/azurerm/3.49.0/expressroutegateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewExpressRouteGateway(name string, args ExpressRouteGatewayArgs) *ExpressRouteGateway {
	return &ExpressRouteGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteGateway)(nil)

type ExpressRouteGateway struct {
	Name  string
	Args  ExpressRouteGatewayArgs
	state *expressRouteGatewayState
}

func (erg *ExpressRouteGateway) Type() string {
	return "azurerm_express_route_gateway"
}

func (erg *ExpressRouteGateway) LocalName() string {
	return erg.Name
}

func (erg *ExpressRouteGateway) Configuration() interface{} {
	return erg.Args
}

func (erg *ExpressRouteGateway) Attributes() expressRouteGatewayAttributes {
	return expressRouteGatewayAttributes{ref: terra.ReferenceResource(erg)}
}

func (erg *ExpressRouteGateway) ImportState(av io.Reader) error {
	erg.state = &expressRouteGatewayState{}
	if err := json.NewDecoder(av).Decode(erg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erg.Type(), erg.LocalName(), err)
	}
	return nil
}

func (erg *ExpressRouteGateway) State() (*expressRouteGatewayState, bool) {
	return erg.state, erg.state != nil
}

func (erg *ExpressRouteGateway) StateMust() *expressRouteGatewayState {
	if erg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erg.Type(), erg.LocalName()))
	}
	return erg.state
}

func (erg *ExpressRouteGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(erg)
}

type ExpressRouteGatewayArgs struct {
	// AllowNonVirtualWanTraffic: bool, optional
	AllowNonVirtualWanTraffic terra.BoolValue `hcl:"allow_non_virtual_wan_traffic,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScaleUnits: number, required
	ScaleUnits terra.NumberValue `hcl:"scale_units,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *expressroutegateway.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ExpressRouteGateway depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type expressRouteGatewayAttributes struct {
	ref terra.Reference
}

func (erg expressRouteGatewayAttributes) AllowNonVirtualWanTraffic() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("allow_non_virtual_wan_traffic"))
}

func (erg expressRouteGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("id"))
}

func (erg expressRouteGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("location"))
}

func (erg expressRouteGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("name"))
}

func (erg expressRouteGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("resource_group_name"))
}

func (erg expressRouteGatewayAttributes) ScaleUnits() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("scale_units"))
}

func (erg expressRouteGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](erg.ref.Append("tags"))
}

func (erg expressRouteGatewayAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("virtual_hub_id"))
}

func (erg expressRouteGatewayAttributes) Timeouts() expressroutegateway.TimeoutsAttributes {
	return terra.ReferenceSingle[expressroutegateway.TimeoutsAttributes](erg.ref.Append("timeouts"))
}

type expressRouteGatewayState struct {
	AllowNonVirtualWanTraffic bool                               `json:"allow_non_virtual_wan_traffic"`
	Id                        string                             `json:"id"`
	Location                  string                             `json:"location"`
	Name                      string                             `json:"name"`
	ResourceGroupName         string                             `json:"resource_group_name"`
	ScaleUnits                float64                            `json:"scale_units"`
	Tags                      map[string]string                  `json:"tags"`
	VirtualHubId              string                             `json:"virtual_hub_id"`
	Timeouts                  *expressroutegateway.TimeoutsState `json:"timeouts"`
}
