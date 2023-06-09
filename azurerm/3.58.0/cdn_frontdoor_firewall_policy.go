// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorfirewallpolicy "github.com/golingon/terraproviders/azurerm/3.58.0/cdnfrontdoorfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorFirewallPolicy creates a new instance of [CdnFrontdoorFirewallPolicy].
func NewCdnFrontdoorFirewallPolicy(name string, args CdnFrontdoorFirewallPolicyArgs) *CdnFrontdoorFirewallPolicy {
	return &CdnFrontdoorFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorFirewallPolicy)(nil)

// CdnFrontdoorFirewallPolicy represents the Terraform resource azurerm_cdn_frontdoor_firewall_policy.
type CdnFrontdoorFirewallPolicy struct {
	Name      string
	Args      CdnFrontdoorFirewallPolicyArgs
	state     *cdnFrontdoorFirewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorFirewallPolicy].
func (cffp *CdnFrontdoorFirewallPolicy) Type() string {
	return "azurerm_cdn_frontdoor_firewall_policy"
}

// LocalName returns the local name for [CdnFrontdoorFirewallPolicy].
func (cffp *CdnFrontdoorFirewallPolicy) LocalName() string {
	return cffp.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorFirewallPolicy].
func (cffp *CdnFrontdoorFirewallPolicy) Configuration() interface{} {
	return cffp.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorFirewallPolicy].
func (cffp *CdnFrontdoorFirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cffp)
}

// Dependencies returns the list of resources [CdnFrontdoorFirewallPolicy] depends_on.
func (cffp *CdnFrontdoorFirewallPolicy) Dependencies() terra.Dependencies {
	return cffp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorFirewallPolicy].
func (cffp *CdnFrontdoorFirewallPolicy) LifecycleManagement() *terra.Lifecycle {
	return cffp.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorFirewallPolicy].
func (cffp *CdnFrontdoorFirewallPolicy) Attributes() cdnFrontdoorFirewallPolicyAttributes {
	return cdnFrontdoorFirewallPolicyAttributes{ref: terra.ReferenceResource(cffp)}
}

// ImportState imports the given attribute values into [CdnFrontdoorFirewallPolicy]'s state.
func (cffp *CdnFrontdoorFirewallPolicy) ImportState(av io.Reader) error {
	cffp.state = &cdnFrontdoorFirewallPolicyState{}
	if err := json.NewDecoder(av).Decode(cffp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cffp.Type(), cffp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorFirewallPolicy] has state.
func (cffp *CdnFrontdoorFirewallPolicy) State() (*cdnFrontdoorFirewallPolicyState, bool) {
	return cffp.state, cffp.state != nil
}

// StateMust returns the state for [CdnFrontdoorFirewallPolicy]. Panics if the state is nil.
func (cffp *CdnFrontdoorFirewallPolicy) StateMust() *cdnFrontdoorFirewallPolicyState {
	if cffp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cffp.Type(), cffp.LocalName()))
	}
	return cffp.state
}

// CdnFrontdoorFirewallPolicyArgs contains the configurations for azurerm_cdn_frontdoor_firewall_policy.
type CdnFrontdoorFirewallPolicyArgs struct {
	// CustomBlockResponseBody: string, optional
	CustomBlockResponseBody terra.StringValue `hcl:"custom_block_response_body,attr"`
	// CustomBlockResponseStatusCode: number, optional
	CustomBlockResponseStatusCode terra.NumberValue `hcl:"custom_block_response_status_code,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RedirectUrl: string, optional
	RedirectUrl terra.StringValue `hcl:"redirect_url,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CustomRule: min=0,max=100
	CustomRule []cdnfrontdoorfirewallpolicy.CustomRule `hcl:"custom_rule,block" validate:"min=0,max=100"`
	// ManagedRule: min=0,max=100
	ManagedRule []cdnfrontdoorfirewallpolicy.ManagedRule `hcl:"managed_rule,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorfirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorFirewallPolicyAttributes struct {
	ref terra.Reference
}

// CustomBlockResponseBody returns a reference to field custom_block_response_body of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) CustomBlockResponseBody() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("custom_block_response_body"))
}

// CustomBlockResponseStatusCode returns a reference to field custom_block_response_status_code of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) CustomBlockResponseStatusCode() terra.NumberValue {
	return terra.ReferenceAsNumber(cffp.ref.Append("custom_block_response_status_code"))
}

// Enabled returns a reference to field enabled of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cffp.ref.Append("enabled"))
}

// FrontendEndpointIds returns a reference to field frontend_endpoint_ids of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) FrontendEndpointIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cffp.ref.Append("frontend_endpoint_ids"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("id"))
}

// Mode returns a reference to field mode of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("name"))
}

// RedirectUrl returns a reference to field redirect_url of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) RedirectUrl() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("redirect_url"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_cdn_frontdoor_firewall_policy.
func (cffp cdnFrontdoorFirewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cffp.ref.Append("tags"))
}

func (cffp cdnFrontdoorFirewallPolicyAttributes) CustomRule() terra.ListValue[cdnfrontdoorfirewallpolicy.CustomRuleAttributes] {
	return terra.ReferenceAsList[cdnfrontdoorfirewallpolicy.CustomRuleAttributes](cffp.ref.Append("custom_rule"))
}

func (cffp cdnFrontdoorFirewallPolicyAttributes) ManagedRule() terra.ListValue[cdnfrontdoorfirewallpolicy.ManagedRuleAttributes] {
	return terra.ReferenceAsList[cdnfrontdoorfirewallpolicy.ManagedRuleAttributes](cffp.ref.Append("managed_rule"))
}

func (cffp cdnFrontdoorFirewallPolicyAttributes) Timeouts() cdnfrontdoorfirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorfirewallpolicy.TimeoutsAttributes](cffp.ref.Append("timeouts"))
}

type cdnFrontdoorFirewallPolicyState struct {
	CustomBlockResponseBody       string                                        `json:"custom_block_response_body"`
	CustomBlockResponseStatusCode float64                                       `json:"custom_block_response_status_code"`
	Enabled                       bool                                          `json:"enabled"`
	FrontendEndpointIds           []string                                      `json:"frontend_endpoint_ids"`
	Id                            string                                        `json:"id"`
	Mode                          string                                        `json:"mode"`
	Name                          string                                        `json:"name"`
	RedirectUrl                   string                                        `json:"redirect_url"`
	ResourceGroupName             string                                        `json:"resource_group_name"`
	SkuName                       string                                        `json:"sku_name"`
	Tags                          map[string]string                             `json:"tags"`
	CustomRule                    []cdnfrontdoorfirewallpolicy.CustomRuleState  `json:"custom_rule"`
	ManagedRule                   []cdnfrontdoorfirewallpolicy.ManagedRuleState `json:"managed_rule"`
	Timeouts                      *cdnfrontdoorfirewallpolicy.TimeoutsState     `json:"timeouts"`
}
