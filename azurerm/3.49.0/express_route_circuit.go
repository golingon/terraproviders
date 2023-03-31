// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressroutecircuit "github.com/golingon/terraproviders/azurerm/3.49.0/expressroutecircuit"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewExpressRouteCircuit(name string, args ExpressRouteCircuitArgs) *ExpressRouteCircuit {
	return &ExpressRouteCircuit{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteCircuit)(nil)

type ExpressRouteCircuit struct {
	Name  string
	Args  ExpressRouteCircuitArgs
	state *expressRouteCircuitState
}

func (erc *ExpressRouteCircuit) Type() string {
	return "azurerm_express_route_circuit"
}

func (erc *ExpressRouteCircuit) LocalName() string {
	return erc.Name
}

func (erc *ExpressRouteCircuit) Configuration() interface{} {
	return erc.Args
}

func (erc *ExpressRouteCircuit) Attributes() expressRouteCircuitAttributes {
	return expressRouteCircuitAttributes{ref: terra.ReferenceResource(erc)}
}

func (erc *ExpressRouteCircuit) ImportState(av io.Reader) error {
	erc.state = &expressRouteCircuitState{}
	if err := json.NewDecoder(av).Decode(erc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erc.Type(), erc.LocalName(), err)
	}
	return nil
}

func (erc *ExpressRouteCircuit) State() (*expressRouteCircuitState, bool) {
	return erc.state, erc.state != nil
}

func (erc *ExpressRouteCircuit) StateMust() *expressRouteCircuitState {
	if erc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erc.Type(), erc.LocalName()))
	}
	return erc.state
}

func (erc *ExpressRouteCircuit) DependOn() terra.Reference {
	return terra.ReferenceResource(erc)
}

type ExpressRouteCircuitArgs struct {
	// AllowClassicOperations: bool, optional
	AllowClassicOperations terra.BoolValue `hcl:"allow_classic_operations,attr"`
	// AuthorizationKey: string, optional
	AuthorizationKey terra.StringValue `hcl:"authorization_key,attr"`
	// BandwidthInGbps: number, optional
	BandwidthInGbps terra.NumberValue `hcl:"bandwidth_in_gbps,attr"`
	// BandwidthInMbps: number, optional
	BandwidthInMbps terra.NumberValue `hcl:"bandwidth_in_mbps,attr"`
	// ExpressRoutePortId: string, optional
	ExpressRoutePortId terra.StringValue `hcl:"express_route_port_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeeringLocation: string, optional
	PeeringLocation terra.StringValue `hcl:"peering_location,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceProviderName: string, optional
	ServiceProviderName terra.StringValue `hcl:"service_provider_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Sku: required
	Sku *expressroutecircuit.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *expressroutecircuit.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ExpressRouteCircuit depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type expressRouteCircuitAttributes struct {
	ref terra.Reference
}

func (erc expressRouteCircuitAttributes) AllowClassicOperations() terra.BoolValue {
	return terra.ReferenceBool(erc.ref.Append("allow_classic_operations"))
}

func (erc expressRouteCircuitAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("authorization_key"))
}

func (erc expressRouteCircuitAttributes) BandwidthInGbps() terra.NumberValue {
	return terra.ReferenceNumber(erc.ref.Append("bandwidth_in_gbps"))
}

func (erc expressRouteCircuitAttributes) BandwidthInMbps() terra.NumberValue {
	return terra.ReferenceNumber(erc.ref.Append("bandwidth_in_mbps"))
}

func (erc expressRouteCircuitAttributes) ExpressRoutePortId() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("express_route_port_id"))
}

func (erc expressRouteCircuitAttributes) Id() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("id"))
}

func (erc expressRouteCircuitAttributes) Location() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("location"))
}

func (erc expressRouteCircuitAttributes) Name() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("name"))
}

func (erc expressRouteCircuitAttributes) PeeringLocation() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("peering_location"))
}

func (erc expressRouteCircuitAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("resource_group_name"))
}

func (erc expressRouteCircuitAttributes) ServiceKey() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("service_key"))
}

func (erc expressRouteCircuitAttributes) ServiceProviderName() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("service_provider_name"))
}

func (erc expressRouteCircuitAttributes) ServiceProviderProvisioningState() terra.StringValue {
	return terra.ReferenceString(erc.ref.Append("service_provider_provisioning_state"))
}

func (erc expressRouteCircuitAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](erc.ref.Append("tags"))
}

func (erc expressRouteCircuitAttributes) Sku() terra.ListValue[expressroutecircuit.SkuAttributes] {
	return terra.ReferenceList[expressroutecircuit.SkuAttributes](erc.ref.Append("sku"))
}

func (erc expressRouteCircuitAttributes) Timeouts() expressroutecircuit.TimeoutsAttributes {
	return terra.ReferenceSingle[expressroutecircuit.TimeoutsAttributes](erc.ref.Append("timeouts"))
}

type expressRouteCircuitState struct {
	AllowClassicOperations           bool                               `json:"allow_classic_operations"`
	AuthorizationKey                 string                             `json:"authorization_key"`
	BandwidthInGbps                  float64                            `json:"bandwidth_in_gbps"`
	BandwidthInMbps                  float64                            `json:"bandwidth_in_mbps"`
	ExpressRoutePortId               string                             `json:"express_route_port_id"`
	Id                               string                             `json:"id"`
	Location                         string                             `json:"location"`
	Name                             string                             `json:"name"`
	PeeringLocation                  string                             `json:"peering_location"`
	ResourceGroupName                string                             `json:"resource_group_name"`
	ServiceKey                       string                             `json:"service_key"`
	ServiceProviderName              string                             `json:"service_provider_name"`
	ServiceProviderProvisioningState string                             `json:"service_provider_provisioning_state"`
	Tags                             map[string]string                  `json:"tags"`
	Sku                              []expressroutecircuit.SkuState     `json:"sku"`
	Timeouts                         *expressroutecircuit.TimeoutsState `json:"timeouts"`
}
