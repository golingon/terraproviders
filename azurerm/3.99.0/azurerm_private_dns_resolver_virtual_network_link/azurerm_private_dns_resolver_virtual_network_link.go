// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_resolver_virtual_network_link

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_private_dns_resolver_virtual_network_link.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPrivateDnsResolverVirtualNetworkLinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (apdrvnl *Resource) Type() string {
	return "azurerm_private_dns_resolver_virtual_network_link"
}

// LocalName returns the local name for [Resource].
func (apdrvnl *Resource) LocalName() string {
	return apdrvnl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (apdrvnl *Resource) Configuration() interface{} {
	return apdrvnl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (apdrvnl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(apdrvnl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (apdrvnl *Resource) Dependencies() terra.Dependencies {
	return apdrvnl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (apdrvnl *Resource) LifecycleManagement() *terra.Lifecycle {
	return apdrvnl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (apdrvnl *Resource) Attributes() azurermPrivateDnsResolverVirtualNetworkLinkAttributes {
	return azurermPrivateDnsResolverVirtualNetworkLinkAttributes{ref: terra.ReferenceResource(apdrvnl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (apdrvnl *Resource) ImportState(state io.Reader) error {
	apdrvnl.state = &azurermPrivateDnsResolverVirtualNetworkLinkState{}
	if err := json.NewDecoder(state).Decode(apdrvnl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apdrvnl.Type(), apdrvnl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (apdrvnl *Resource) State() (*azurermPrivateDnsResolverVirtualNetworkLinkState, bool) {
	return apdrvnl.state, apdrvnl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (apdrvnl *Resource) StateMust() *azurermPrivateDnsResolverVirtualNetworkLinkState {
	if apdrvnl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apdrvnl.Type(), apdrvnl.LocalName()))
	}
	return apdrvnl.state
}

// Args contains the configurations for azurerm_private_dns_resolver_virtual_network_link.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPrivateDnsResolverVirtualNetworkLinkAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl azurermPrivateDnsResolverVirtualNetworkLinkAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("dns_forwarding_ruleset_id"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl azurermPrivateDnsResolverVirtualNetworkLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl azurermPrivateDnsResolverVirtualNetworkLinkAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdrvnl.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl azurermPrivateDnsResolverVirtualNetworkLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("name"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl azurermPrivateDnsResolverVirtualNetworkLinkAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("virtual_network_id"))
}

func (apdrvnl azurermPrivateDnsResolverVirtualNetworkLinkAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](apdrvnl.ref.Append("timeouts"))
}

type azurermPrivateDnsResolverVirtualNetworkLinkState struct {
	DnsForwardingRulesetId string            `json:"dns_forwarding_ruleset_id"`
	Id                     string            `json:"id"`
	Metadata               map[string]string `json:"metadata"`
	Name                   string            `json:"name"`
	VirtualNetworkId       string            `json:"virtual_network_id"`
	Timeouts               *TimeoutsState    `json:"timeouts"`
}
