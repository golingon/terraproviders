// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpngatewaynatrule "github.com/golingon/terraproviders/azurerm/3.49.0/vpngatewaynatrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnGatewayNatRule creates a new instance of [VpnGatewayNatRule].
func NewVpnGatewayNatRule(name string, args VpnGatewayNatRuleArgs) *VpnGatewayNatRule {
	return &VpnGatewayNatRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnGatewayNatRule)(nil)

// VpnGatewayNatRule represents the Terraform resource azurerm_vpn_gateway_nat_rule.
type VpnGatewayNatRule struct {
	Name      string
	Args      VpnGatewayNatRuleArgs
	state     *vpnGatewayNatRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnGatewayNatRule].
func (vgnr *VpnGatewayNatRule) Type() string {
	return "azurerm_vpn_gateway_nat_rule"
}

// LocalName returns the local name for [VpnGatewayNatRule].
func (vgnr *VpnGatewayNatRule) LocalName() string {
	return vgnr.Name
}

// Configuration returns the configuration (args) for [VpnGatewayNatRule].
func (vgnr *VpnGatewayNatRule) Configuration() interface{} {
	return vgnr.Args
}

// DependOn is used for other resources to depend on [VpnGatewayNatRule].
func (vgnr *VpnGatewayNatRule) DependOn() terra.Reference {
	return terra.ReferenceResource(vgnr)
}

// Dependencies returns the list of resources [VpnGatewayNatRule] depends_on.
func (vgnr *VpnGatewayNatRule) Dependencies() terra.Dependencies {
	return vgnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnGatewayNatRule].
func (vgnr *VpnGatewayNatRule) LifecycleManagement() *terra.Lifecycle {
	return vgnr.Lifecycle
}

// Attributes returns the attributes for [VpnGatewayNatRule].
func (vgnr *VpnGatewayNatRule) Attributes() vpnGatewayNatRuleAttributes {
	return vpnGatewayNatRuleAttributes{ref: terra.ReferenceResource(vgnr)}
}

// ImportState imports the given attribute values into [VpnGatewayNatRule]'s state.
func (vgnr *VpnGatewayNatRule) ImportState(av io.Reader) error {
	vgnr.state = &vpnGatewayNatRuleState{}
	if err := json.NewDecoder(av).Decode(vgnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vgnr.Type(), vgnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnGatewayNatRule] has state.
func (vgnr *VpnGatewayNatRule) State() (*vpnGatewayNatRuleState, bool) {
	return vgnr.state, vgnr.state != nil
}

// StateMust returns the state for [VpnGatewayNatRule]. Panics if the state is nil.
func (vgnr *VpnGatewayNatRule) StateMust() *vpnGatewayNatRuleState {
	if vgnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vgnr.Type(), vgnr.LocalName()))
	}
	return vgnr.state
}

// VpnGatewayNatRuleArgs contains the configurations for azurerm_vpn_gateway_nat_rule.
type VpnGatewayNatRuleArgs struct {
	// ExternalAddressSpaceMappings: list of string, optional
	ExternalAddressSpaceMappings terra.ListValue[terra.StringValue] `hcl:"external_address_space_mappings,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternalAddressSpaceMappings: list of string, optional
	InternalAddressSpaceMappings terra.ListValue[terra.StringValue] `hcl:"internal_address_space_mappings,attr"`
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
	// VpnGatewayId: string, required
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr" validate:"required"`
	// ExternalMapping: min=0
	ExternalMapping []vpngatewaynatrule.ExternalMapping `hcl:"external_mapping,block" validate:"min=0"`
	// InternalMapping: min=0
	InternalMapping []vpngatewaynatrule.InternalMapping `hcl:"internal_mapping,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *vpngatewaynatrule.Timeouts `hcl:"timeouts,block"`
}
type vpnGatewayNatRuleAttributes struct {
	ref terra.Reference
}

// ExternalAddressSpaceMappings returns a reference to field external_address_space_mappings of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) ExternalAddressSpaceMappings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vgnr.ref.Append("external_address_space_mappings"))
}

// Id returns a reference to field id of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("id"))
}

// InternalAddressSpaceMappings returns a reference to field internal_address_space_mappings of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) InternalAddressSpaceMappings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vgnr.ref.Append("internal_address_space_mappings"))
}

// IpConfigurationId returns a reference to field ip_configuration_id of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) IpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("ip_configuration_id"))
}

// Mode returns a reference to field mode of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("resource_group_name"))
}

// Type returns a reference to field type of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("type"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of azurerm_vpn_gateway_nat_rule.
func (vgnr vpnGatewayNatRuleAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vgnr.ref.Append("vpn_gateway_id"))
}

func (vgnr vpnGatewayNatRuleAttributes) ExternalMapping() terra.ListValue[vpngatewaynatrule.ExternalMappingAttributes] {
	return terra.ReferenceAsList[vpngatewaynatrule.ExternalMappingAttributes](vgnr.ref.Append("external_mapping"))
}

func (vgnr vpnGatewayNatRuleAttributes) InternalMapping() terra.ListValue[vpngatewaynatrule.InternalMappingAttributes] {
	return terra.ReferenceAsList[vpngatewaynatrule.InternalMappingAttributes](vgnr.ref.Append("internal_mapping"))
}

func (vgnr vpnGatewayNatRuleAttributes) Timeouts() vpngatewaynatrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpngatewaynatrule.TimeoutsAttributes](vgnr.ref.Append("timeouts"))
}

type vpnGatewayNatRuleState struct {
	ExternalAddressSpaceMappings []string                                 `json:"external_address_space_mappings"`
	Id                           string                                   `json:"id"`
	InternalAddressSpaceMappings []string                                 `json:"internal_address_space_mappings"`
	IpConfigurationId            string                                   `json:"ip_configuration_id"`
	Mode                         string                                   `json:"mode"`
	Name                         string                                   `json:"name"`
	ResourceGroupName            string                                   `json:"resource_group_name"`
	Type                         string                                   `json:"type"`
	VpnGatewayId                 string                                   `json:"vpn_gateway_id"`
	ExternalMapping              []vpngatewaynatrule.ExternalMappingState `json:"external_mapping"`
	InternalMapping              []vpngatewaynatrule.InternalMappingState `json:"internal_mapping"`
	Timeouts                     *vpngatewaynatrule.TimeoutsState         `json:"timeouts"`
}
