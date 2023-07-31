// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subnet "github.com/golingon/terraproviders/azurerm/3.67.0/subnet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnet creates a new instance of [Subnet].
func NewSubnet(name string, args SubnetArgs) *Subnet {
	return &Subnet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Subnet)(nil)

// Subnet represents the Terraform resource azurerm_subnet.
type Subnet struct {
	Name      string
	Args      SubnetArgs
	state     *subnetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Subnet].
func (s *Subnet) Type() string {
	return "azurerm_subnet"
}

// LocalName returns the local name for [Subnet].
func (s *Subnet) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [Subnet].
func (s *Subnet) Configuration() interface{} {
	return s.Args
}

// DependOn is used for other resources to depend on [Subnet].
func (s *Subnet) DependOn() terra.Reference {
	return terra.ReferenceResource(s)
}

// Dependencies returns the list of resources [Subnet] depends_on.
func (s *Subnet) Dependencies() terra.Dependencies {
	return s.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Subnet].
func (s *Subnet) LifecycleManagement() *terra.Lifecycle {
	return s.Lifecycle
}

// Attributes returns the attributes for [Subnet].
func (s *Subnet) Attributes() subnetAttributes {
	return subnetAttributes{ref: terra.ReferenceResource(s)}
}

// ImportState imports the given attribute values into [Subnet]'s state.
func (s *Subnet) ImportState(av io.Reader) error {
	s.state = &subnetState{}
	if err := json.NewDecoder(av).Decode(s.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", s.Type(), s.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Subnet] has state.
func (s *Subnet) State() (*subnetState, bool) {
	return s.state, s.state != nil
}

// StateMust returns the state for [Subnet]. Panics if the state is nil.
func (s *Subnet) StateMust() *subnetState {
	if s.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", s.Type(), s.LocalName()))
	}
	return s.state
}

// SubnetArgs contains the configurations for azurerm_subnet.
type SubnetArgs struct {
	// AddressPrefixes: list of string, required
	AddressPrefixes terra.ListValue[terra.StringValue] `hcl:"address_prefixes,attr" validate:"required"`
	// EnforcePrivateLinkEndpointNetworkPolicies: bool, optional
	EnforcePrivateLinkEndpointNetworkPolicies terra.BoolValue `hcl:"enforce_private_link_endpoint_network_policies,attr"`
	// EnforcePrivateLinkServiceNetworkPolicies: bool, optional
	EnforcePrivateLinkServiceNetworkPolicies terra.BoolValue `hcl:"enforce_private_link_service_network_policies,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateEndpointNetworkPoliciesEnabled: bool, optional
	PrivateEndpointNetworkPoliciesEnabled terra.BoolValue `hcl:"private_endpoint_network_policies_enabled,attr"`
	// PrivateLinkServiceNetworkPoliciesEnabled: bool, optional
	PrivateLinkServiceNetworkPoliciesEnabled terra.BoolValue `hcl:"private_link_service_network_policies_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceEndpointPolicyIds: set of string, optional
	ServiceEndpointPolicyIds terra.SetValue[terra.StringValue] `hcl:"service_endpoint_policy_ids,attr"`
	// ServiceEndpoints: set of string, optional
	ServiceEndpoints terra.SetValue[terra.StringValue] `hcl:"service_endpoints,attr"`
	// VirtualNetworkName: string, required
	VirtualNetworkName terra.StringValue `hcl:"virtual_network_name,attr" validate:"required"`
	// Delegation: min=0
	Delegation []subnet.Delegation `hcl:"delegation,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *subnet.Timeouts `hcl:"timeouts,block"`
}
type subnetAttributes struct {
	ref terra.Reference
}

// AddressPrefixes returns a reference to field address_prefixes of azurerm_subnet.
func (s subnetAttributes) AddressPrefixes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("address_prefixes"))
}

// EnforcePrivateLinkEndpointNetworkPolicies returns a reference to field enforce_private_link_endpoint_network_policies of azurerm_subnet.
func (s subnetAttributes) EnforcePrivateLinkEndpointNetworkPolicies() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enforce_private_link_endpoint_network_policies"))
}

// EnforcePrivateLinkServiceNetworkPolicies returns a reference to field enforce_private_link_service_network_policies of azurerm_subnet.
func (s subnetAttributes) EnforcePrivateLinkServiceNetworkPolicies() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enforce_private_link_service_network_policies"))
}

// Id returns a reference to field id of azurerm_subnet.
func (s subnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_subnet.
func (s subnetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

// PrivateEndpointNetworkPoliciesEnabled returns a reference to field private_endpoint_network_policies_enabled of azurerm_subnet.
func (s subnetAttributes) PrivateEndpointNetworkPoliciesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("private_endpoint_network_policies_enabled"))
}

// PrivateLinkServiceNetworkPoliciesEnabled returns a reference to field private_link_service_network_policies_enabled of azurerm_subnet.
func (s subnetAttributes) PrivateLinkServiceNetworkPoliciesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("private_link_service_network_policies_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_subnet.
func (s subnetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("resource_group_name"))
}

// ServiceEndpointPolicyIds returns a reference to field service_endpoint_policy_ids of azurerm_subnet.
func (s subnetAttributes) ServiceEndpointPolicyIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("service_endpoint_policy_ids"))
}

// ServiceEndpoints returns a reference to field service_endpoints of azurerm_subnet.
func (s subnetAttributes) ServiceEndpoints() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("service_endpoints"))
}

// VirtualNetworkName returns a reference to field virtual_network_name of azurerm_subnet.
func (s subnetAttributes) VirtualNetworkName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("virtual_network_name"))
}

func (s subnetAttributes) Delegation() terra.ListValue[subnet.DelegationAttributes] {
	return terra.ReferenceAsList[subnet.DelegationAttributes](s.ref.Append("delegation"))
}

func (s subnetAttributes) Timeouts() subnet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnet.TimeoutsAttributes](s.ref.Append("timeouts"))
}

type subnetState struct {
	AddressPrefixes                           []string                 `json:"address_prefixes"`
	EnforcePrivateLinkEndpointNetworkPolicies bool                     `json:"enforce_private_link_endpoint_network_policies"`
	EnforcePrivateLinkServiceNetworkPolicies  bool                     `json:"enforce_private_link_service_network_policies"`
	Id                                        string                   `json:"id"`
	Name                                      string                   `json:"name"`
	PrivateEndpointNetworkPoliciesEnabled     bool                     `json:"private_endpoint_network_policies_enabled"`
	PrivateLinkServiceNetworkPoliciesEnabled  bool                     `json:"private_link_service_network_policies_enabled"`
	ResourceGroupName                         string                   `json:"resource_group_name"`
	ServiceEndpointPolicyIds                  []string                 `json:"service_endpoint_policy_ids"`
	ServiceEndpoints                          []string                 `json:"service_endpoints"`
	VirtualNetworkName                        string                   `json:"virtual_network_name"`
	Delegation                                []subnet.DelegationState `json:"delegation"`
	Timeouts                                  *subnet.TimeoutsState    `json:"timeouts"`
}
