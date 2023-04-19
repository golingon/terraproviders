// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsresolverdnsforwardingruleset "github.com/golingon/terraproviders/azurerm/3.52.0/privatednsresolverdnsforwardingruleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsResolverDnsForwardingRuleset creates a new instance of [PrivateDnsResolverDnsForwardingRuleset].
func NewPrivateDnsResolverDnsForwardingRuleset(name string, args PrivateDnsResolverDnsForwardingRulesetArgs) *PrivateDnsResolverDnsForwardingRuleset {
	return &PrivateDnsResolverDnsForwardingRuleset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsResolverDnsForwardingRuleset)(nil)

// PrivateDnsResolverDnsForwardingRuleset represents the Terraform resource azurerm_private_dns_resolver_dns_forwarding_ruleset.
type PrivateDnsResolverDnsForwardingRuleset struct {
	Name      string
	Args      PrivateDnsResolverDnsForwardingRulesetArgs
	state     *privateDnsResolverDnsForwardingRulesetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsResolverDnsForwardingRuleset].
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) Type() string {
	return "azurerm_private_dns_resolver_dns_forwarding_ruleset"
}

// LocalName returns the local name for [PrivateDnsResolverDnsForwardingRuleset].
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) LocalName() string {
	return pdrdfr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsResolverDnsForwardingRuleset].
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) Configuration() interface{} {
	return pdrdfr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsResolverDnsForwardingRuleset].
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) DependOn() terra.Reference {
	return terra.ReferenceResource(pdrdfr)
}

// Dependencies returns the list of resources [PrivateDnsResolverDnsForwardingRuleset] depends_on.
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) Dependencies() terra.Dependencies {
	return pdrdfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsResolverDnsForwardingRuleset].
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) LifecycleManagement() *terra.Lifecycle {
	return pdrdfr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsResolverDnsForwardingRuleset].
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) Attributes() privateDnsResolverDnsForwardingRulesetAttributes {
	return privateDnsResolverDnsForwardingRulesetAttributes{ref: terra.ReferenceResource(pdrdfr)}
}

// ImportState imports the given attribute values into [PrivateDnsResolverDnsForwardingRuleset]'s state.
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) ImportState(av io.Reader) error {
	pdrdfr.state = &privateDnsResolverDnsForwardingRulesetState{}
	if err := json.NewDecoder(av).Decode(pdrdfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdrdfr.Type(), pdrdfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsResolverDnsForwardingRuleset] has state.
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) State() (*privateDnsResolverDnsForwardingRulesetState, bool) {
	return pdrdfr.state, pdrdfr.state != nil
}

// StateMust returns the state for [PrivateDnsResolverDnsForwardingRuleset]. Panics if the state is nil.
func (pdrdfr *PrivateDnsResolverDnsForwardingRuleset) StateMust() *privateDnsResolverDnsForwardingRulesetState {
	if pdrdfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdrdfr.Type(), pdrdfr.LocalName()))
	}
	return pdrdfr.state
}

// PrivateDnsResolverDnsForwardingRulesetArgs contains the configurations for azurerm_private_dns_resolver_dns_forwarding_ruleset.
type PrivateDnsResolverDnsForwardingRulesetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsResolverOutboundEndpointIds: list of string, required
	PrivateDnsResolverOutboundEndpointIds terra.ListValue[terra.StringValue] `hcl:"private_dns_resolver_outbound_endpoint_ids,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *privatednsresolverdnsforwardingruleset.Timeouts `hcl:"timeouts,block"`
}
type privateDnsResolverDnsForwardingRulesetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver_dns_forwarding_ruleset.
func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdrdfr.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver_dns_forwarding_ruleset.
func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pdrdfr.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_dns_forwarding_ruleset.
func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdrdfr.ref.Append("name"))
}

// PrivateDnsResolverOutboundEndpointIds returns a reference to field private_dns_resolver_outbound_endpoint_ids of azurerm_private_dns_resolver_dns_forwarding_ruleset.
func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) PrivateDnsResolverOutboundEndpointIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pdrdfr.ref.Append("private_dns_resolver_outbound_endpoint_ids"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_resolver_dns_forwarding_ruleset.
func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdrdfr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver_dns_forwarding_ruleset.
func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdrdfr.ref.Append("tags"))
}

func (pdrdfr privateDnsResolverDnsForwardingRulesetAttributes) Timeouts() privatednsresolverdnsforwardingruleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsresolverdnsforwardingruleset.TimeoutsAttributes](pdrdfr.ref.Append("timeouts"))
}

type privateDnsResolverDnsForwardingRulesetState struct {
	Id                                    string                                                `json:"id"`
	Location                              string                                                `json:"location"`
	Name                                  string                                                `json:"name"`
	PrivateDnsResolverOutboundEndpointIds []string                                              `json:"private_dns_resolver_outbound_endpoint_ids"`
	ResourceGroupName                     string                                                `json:"resource_group_name"`
	Tags                                  map[string]string                                     `json:"tags"`
	Timeouts                              *privatednsresolverdnsforwardingruleset.TimeoutsState `json:"timeouts"`
}
