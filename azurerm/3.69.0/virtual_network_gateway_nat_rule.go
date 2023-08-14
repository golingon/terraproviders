// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualnetworkgatewaynatrule "github.com/golingon/terraproviders/azurerm/3.69.0/virtualnetworkgatewaynatrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualNetworkGatewayNatRule creates a new instance of [VirtualNetworkGatewayNatRule].
func NewVirtualNetworkGatewayNatRule(name string, args VirtualNetworkGatewayNatRuleArgs) *VirtualNetworkGatewayNatRule {
	return &VirtualNetworkGatewayNatRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualNetworkGatewayNatRule)(nil)

// VirtualNetworkGatewayNatRule represents the Terraform resource azurerm_virtual_network_gateway_nat_rule.
type VirtualNetworkGatewayNatRule struct {
	Name      string
	Args      VirtualNetworkGatewayNatRuleArgs
	state     *virtualNetworkGatewayNatRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualNetworkGatewayNatRule].
func (vngnr *VirtualNetworkGatewayNatRule) Type() string {
	return "azurerm_virtual_network_gateway_nat_rule"
}

// LocalName returns the local name for [VirtualNetworkGatewayNatRule].
func (vngnr *VirtualNetworkGatewayNatRule) LocalName() string {
	return vngnr.Name
}

// Configuration returns the configuration (args) for [VirtualNetworkGatewayNatRule].
func (vngnr *VirtualNetworkGatewayNatRule) Configuration() interface{} {
	return vngnr.Args
}

// DependOn is used for other resources to depend on [VirtualNetworkGatewayNatRule].
func (vngnr *VirtualNetworkGatewayNatRule) DependOn() terra.Reference {
	return terra.ReferenceResource(vngnr)
}

// Dependencies returns the list of resources [VirtualNetworkGatewayNatRule] depends_on.
func (vngnr *VirtualNetworkGatewayNatRule) Dependencies() terra.Dependencies {
	return vngnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualNetworkGatewayNatRule].
func (vngnr *VirtualNetworkGatewayNatRule) LifecycleManagement() *terra.Lifecycle {
	return vngnr.Lifecycle
}

// Attributes returns the attributes for [VirtualNetworkGatewayNatRule].
func (vngnr *VirtualNetworkGatewayNatRule) Attributes() virtualNetworkGatewayNatRuleAttributes {
	return virtualNetworkGatewayNatRuleAttributes{ref: terra.ReferenceResource(vngnr)}
}

// ImportState imports the given attribute values into [VirtualNetworkGatewayNatRule]'s state.
func (vngnr *VirtualNetworkGatewayNatRule) ImportState(av io.Reader) error {
	vngnr.state = &virtualNetworkGatewayNatRuleState{}
	if err := json.NewDecoder(av).Decode(vngnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vngnr.Type(), vngnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualNetworkGatewayNatRule] has state.
func (vngnr *VirtualNetworkGatewayNatRule) State() (*virtualNetworkGatewayNatRuleState, bool) {
	return vngnr.state, vngnr.state != nil
}

// StateMust returns the state for [VirtualNetworkGatewayNatRule]. Panics if the state is nil.
func (vngnr *VirtualNetworkGatewayNatRule) StateMust() *virtualNetworkGatewayNatRuleState {
	if vngnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vngnr.Type(), vngnr.LocalName()))
	}
	return vngnr.state
}

// VirtualNetworkGatewayNatRuleArgs contains the configurations for azurerm_virtual_network_gateway_nat_rule.
type VirtualNetworkGatewayNatRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConfigurationId: string, optional
	IpConfigurationId terra.StringValue `hcl:"ip_configuration_id,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// VirtualNetworkGatewayId: string, required
	VirtualNetworkGatewayId terra.StringValue `hcl:"virtual_network_gateway_id,attr" validate:"required"`
	// ExternalMapping: min=1
	ExternalMapping []virtualnetworkgatewaynatrule.ExternalMapping `hcl:"external_mapping,block" validate:"min=1"`
	// InternalMapping: min=1
	InternalMapping []virtualnetworkgatewaynatrule.InternalMapping `hcl:"internal_mapping,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *virtualnetworkgatewaynatrule.Timeouts `hcl:"timeouts,block"`
}
type virtualNetworkGatewayNatRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("id"))
}

// IpConfigurationId returns a reference to field ip_configuration_id of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) IpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("ip_configuration_id"))
}

// Mode returns a reference to field mode of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("resource_group_name"))
}

// Type returns a reference to field type of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("type"))
}

// VirtualNetworkGatewayId returns a reference to field virtual_network_gateway_id of azurerm_virtual_network_gateway_nat_rule.
func (vngnr virtualNetworkGatewayNatRuleAttributes) VirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngnr.ref.Append("virtual_network_gateway_id"))
}

func (vngnr virtualNetworkGatewayNatRuleAttributes) ExternalMapping() terra.ListValue[virtualnetworkgatewaynatrule.ExternalMappingAttributes] {
	return terra.ReferenceAsList[virtualnetworkgatewaynatrule.ExternalMappingAttributes](vngnr.ref.Append("external_mapping"))
}

func (vngnr virtualNetworkGatewayNatRuleAttributes) InternalMapping() terra.ListValue[virtualnetworkgatewaynatrule.InternalMappingAttributes] {
	return terra.ReferenceAsList[virtualnetworkgatewaynatrule.InternalMappingAttributes](vngnr.ref.Append("internal_mapping"))
}

func (vngnr virtualNetworkGatewayNatRuleAttributes) Timeouts() virtualnetworkgatewaynatrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualnetworkgatewaynatrule.TimeoutsAttributes](vngnr.ref.Append("timeouts"))
}

type virtualNetworkGatewayNatRuleState struct {
	Id                      string                                              `json:"id"`
	IpConfigurationId       string                                              `json:"ip_configuration_id"`
	Mode                    string                                              `json:"mode"`
	Name                    string                                              `json:"name"`
	ResourceGroupName       string                                              `json:"resource_group_name"`
	Type                    string                                              `json:"type"`
	VirtualNetworkGatewayId string                                              `json:"virtual_network_gateway_id"`
	ExternalMapping         []virtualnetworkgatewaynatrule.ExternalMappingState `json:"external_mapping"`
	InternalMapping         []virtualnetworkgatewaynatrule.InternalMappingState `json:"internal_mapping"`
	Timeouts                *virtualnetworkgatewaynatrule.TimeoutsState         `json:"timeouts"`
}
