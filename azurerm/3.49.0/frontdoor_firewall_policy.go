// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	frontdoorfirewallpolicy "github.com/golingon/terraproviders/azurerm/3.49.0/frontdoorfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewFrontdoorFirewallPolicy(name string, args FrontdoorFirewallPolicyArgs) *FrontdoorFirewallPolicy {
	return &FrontdoorFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FrontdoorFirewallPolicy)(nil)

type FrontdoorFirewallPolicy struct {
	Name  string
	Args  FrontdoorFirewallPolicyArgs
	state *frontdoorFirewallPolicyState
}

func (ffp *FrontdoorFirewallPolicy) Type() string {
	return "azurerm_frontdoor_firewall_policy"
}

func (ffp *FrontdoorFirewallPolicy) LocalName() string {
	return ffp.Name
}

func (ffp *FrontdoorFirewallPolicy) Configuration() interface{} {
	return ffp.Args
}

func (ffp *FrontdoorFirewallPolicy) Attributes() frontdoorFirewallPolicyAttributes {
	return frontdoorFirewallPolicyAttributes{ref: terra.ReferenceResource(ffp)}
}

func (ffp *FrontdoorFirewallPolicy) ImportState(av io.Reader) error {
	ffp.state = &frontdoorFirewallPolicyState{}
	if err := json.NewDecoder(av).Decode(ffp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ffp.Type(), ffp.LocalName(), err)
	}
	return nil
}

func (ffp *FrontdoorFirewallPolicy) State() (*frontdoorFirewallPolicyState, bool) {
	return ffp.state, ffp.state != nil
}

func (ffp *FrontdoorFirewallPolicy) StateMust() *frontdoorFirewallPolicyState {
	if ffp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ffp.Type(), ffp.LocalName()))
	}
	return ffp.state
}

func (ffp *FrontdoorFirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ffp)
}

type FrontdoorFirewallPolicyArgs struct {
	// CustomBlockResponseBody: string, optional
	CustomBlockResponseBody terra.StringValue `hcl:"custom_block_response_body,attr"`
	// CustomBlockResponseStatusCode: number, optional
	CustomBlockResponseStatusCode terra.NumberValue `hcl:"custom_block_response_status_code,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RedirectUrl: string, optional
	RedirectUrl terra.StringValue `hcl:"redirect_url,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CustomRule: min=0,max=100
	CustomRule []frontdoorfirewallpolicy.CustomRule `hcl:"custom_rule,block" validate:"min=0,max=100"`
	// ManagedRule: min=0,max=100
	ManagedRule []frontdoorfirewallpolicy.ManagedRule `hcl:"managed_rule,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *frontdoorfirewallpolicy.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that FrontdoorFirewallPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type frontdoorFirewallPolicyAttributes struct {
	ref terra.Reference
}

func (ffp frontdoorFirewallPolicyAttributes) CustomBlockResponseBody() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("custom_block_response_body"))
}

func (ffp frontdoorFirewallPolicyAttributes) CustomBlockResponseStatusCode() terra.NumberValue {
	return terra.ReferenceNumber(ffp.ref.Append("custom_block_response_status_code"))
}

func (ffp frontdoorFirewallPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ffp.ref.Append("enabled"))
}

func (ffp frontdoorFirewallPolicyAttributes) FrontendEndpointIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ffp.ref.Append("frontend_endpoint_ids"))
}

func (ffp frontdoorFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("id"))
}

func (ffp frontdoorFirewallPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("location"))
}

func (ffp frontdoorFirewallPolicyAttributes) Mode() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("mode"))
}

func (ffp frontdoorFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("name"))
}

func (ffp frontdoorFirewallPolicyAttributes) RedirectUrl() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("redirect_url"))
}

func (ffp frontdoorFirewallPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ffp.ref.Append("resource_group_name"))
}

func (ffp frontdoorFirewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ffp.ref.Append("tags"))
}

func (ffp frontdoorFirewallPolicyAttributes) CustomRule() terra.ListValue[frontdoorfirewallpolicy.CustomRuleAttributes] {
	return terra.ReferenceList[frontdoorfirewallpolicy.CustomRuleAttributes](ffp.ref.Append("custom_rule"))
}

func (ffp frontdoorFirewallPolicyAttributes) ManagedRule() terra.ListValue[frontdoorfirewallpolicy.ManagedRuleAttributes] {
	return terra.ReferenceList[frontdoorfirewallpolicy.ManagedRuleAttributes](ffp.ref.Append("managed_rule"))
}

func (ffp frontdoorFirewallPolicyAttributes) Timeouts() frontdoorfirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[frontdoorfirewallpolicy.TimeoutsAttributes](ffp.ref.Append("timeouts"))
}

type frontdoorFirewallPolicyState struct {
	CustomBlockResponseBody       string                                     `json:"custom_block_response_body"`
	CustomBlockResponseStatusCode float64                                    `json:"custom_block_response_status_code"`
	Enabled                       bool                                       `json:"enabled"`
	FrontendEndpointIds           []string                                   `json:"frontend_endpoint_ids"`
	Id                            string                                     `json:"id"`
	Location                      string                                     `json:"location"`
	Mode                          string                                     `json:"mode"`
	Name                          string                                     `json:"name"`
	RedirectUrl                   string                                     `json:"redirect_url"`
	ResourceGroupName             string                                     `json:"resource_group_name"`
	Tags                          map[string]string                          `json:"tags"`
	CustomRule                    []frontdoorfirewallpolicy.CustomRuleState  `json:"custom_rule"`
	ManagedRule                   []frontdoorfirewallpolicy.ManagedRuleState `json:"managed_rule"`
	Timeouts                      *frontdoorfirewallpolicy.TimeoutsState     `json:"timeouts"`
}
