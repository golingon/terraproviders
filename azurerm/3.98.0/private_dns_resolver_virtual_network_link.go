// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	privatednsresolvervirtualnetworklink "github.com/golingon/terraproviders/azurerm/3.98.0/privatednsresolvervirtualnetworklink"
	"io"
)

// NewPrivateDnsResolverVirtualNetworkLink creates a new instance of [PrivateDnsResolverVirtualNetworkLink].
func NewPrivateDnsResolverVirtualNetworkLink(name string, args PrivateDnsResolverVirtualNetworkLinkArgs) *PrivateDnsResolverVirtualNetworkLink {
	return &PrivateDnsResolverVirtualNetworkLink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsResolverVirtualNetworkLink)(nil)

// PrivateDnsResolverVirtualNetworkLink represents the Terraform resource azurerm_private_dns_resolver_virtual_network_link.
type PrivateDnsResolverVirtualNetworkLink struct {
	Name      string
	Args      PrivateDnsResolverVirtualNetworkLinkArgs
	state     *privateDnsResolverVirtualNetworkLinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) Type() string {
	return "azurerm_private_dns_resolver_virtual_network_link"
}

// LocalName returns the local name for [PrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) LocalName() string {
	return pdrvnl.Name
}

// Configuration returns the configuration (args) for [PrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) Configuration() interface{} {
	return pdrvnl.Args
}

// DependOn is used for other resources to depend on [PrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) DependOn() terra.Reference {
	return terra.ReferenceResource(pdrvnl)
}

// Dependencies returns the list of resources [PrivateDnsResolverVirtualNetworkLink] depends_on.
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) Dependencies() terra.Dependencies {
	return pdrvnl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) LifecycleManagement() *terra.Lifecycle {
	return pdrvnl.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) Attributes() privateDnsResolverVirtualNetworkLinkAttributes {
	return privateDnsResolverVirtualNetworkLinkAttributes{ref: terra.ReferenceResource(pdrvnl)}
}

// ImportState imports the given attribute values into [PrivateDnsResolverVirtualNetworkLink]'s state.
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) ImportState(av io.Reader) error {
	pdrvnl.state = &privateDnsResolverVirtualNetworkLinkState{}
	if err := json.NewDecoder(av).Decode(pdrvnl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdrvnl.Type(), pdrvnl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsResolverVirtualNetworkLink] has state.
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) State() (*privateDnsResolverVirtualNetworkLinkState, bool) {
	return pdrvnl.state, pdrvnl.state != nil
}

// StateMust returns the state for [PrivateDnsResolverVirtualNetworkLink]. Panics if the state is nil.
func (pdrvnl *PrivateDnsResolverVirtualNetworkLink) StateMust() *privateDnsResolverVirtualNetworkLinkState {
	if pdrvnl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdrvnl.Type(), pdrvnl.LocalName()))
	}
	return pdrvnl.state
}

// PrivateDnsResolverVirtualNetworkLinkArgs contains the configurations for azurerm_private_dns_resolver_virtual_network_link.
type PrivateDnsResolverVirtualNetworkLinkArgs struct {
	// DnsForwardingRulesetId: string, required
	DnsForwardingRulesetId terra.StringValue `hcl:"dns_forwarding_ruleset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VirtualNetworkId: string, required
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *privatednsresolvervirtualnetworklink.Timeouts `hcl:"timeouts,block"`
}
type privateDnsResolverVirtualNetworkLinkAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl privateDnsResolverVirtualNetworkLinkAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("dns_forwarding_ruleset_id"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl privateDnsResolverVirtualNetworkLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl privateDnsResolverVirtualNetworkLinkAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdrvnl.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl privateDnsResolverVirtualNetworkLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("name"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl privateDnsResolverVirtualNetworkLinkAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("virtual_network_id"))
}

func (pdrvnl privateDnsResolverVirtualNetworkLinkAttributes) Timeouts() privatednsresolvervirtualnetworklink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsresolvervirtualnetworklink.TimeoutsAttributes](pdrvnl.ref.Append("timeouts"))
}

type privateDnsResolverVirtualNetworkLinkState struct {
	DnsForwardingRulesetId string                                              `json:"dns_forwarding_ruleset_id"`
	Id                     string                                              `json:"id"`
	Metadata               map[string]string                                   `json:"metadata"`
	Name                   string                                              `json:"name"`
	VirtualNetworkId       string                                              `json:"virtual_network_id"`
	Timeouts               *privatednsresolvervirtualnetworklink.TimeoutsState `json:"timeouts"`
}
