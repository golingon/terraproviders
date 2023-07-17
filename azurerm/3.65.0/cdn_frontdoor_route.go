// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorroute "github.com/golingon/terraproviders/azurerm/3.65.0/cdnfrontdoorroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorRoute creates a new instance of [CdnFrontdoorRoute].
func NewCdnFrontdoorRoute(name string, args CdnFrontdoorRouteArgs) *CdnFrontdoorRoute {
	return &CdnFrontdoorRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorRoute)(nil)

// CdnFrontdoorRoute represents the Terraform resource azurerm_cdn_frontdoor_route.
type CdnFrontdoorRoute struct {
	Name      string
	Args      CdnFrontdoorRouteArgs
	state     *cdnFrontdoorRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorRoute].
func (cfr *CdnFrontdoorRoute) Type() string {
	return "azurerm_cdn_frontdoor_route"
}

// LocalName returns the local name for [CdnFrontdoorRoute].
func (cfr *CdnFrontdoorRoute) LocalName() string {
	return cfr.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorRoute].
func (cfr *CdnFrontdoorRoute) Configuration() interface{} {
	return cfr.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorRoute].
func (cfr *CdnFrontdoorRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(cfr)
}

// Dependencies returns the list of resources [CdnFrontdoorRoute] depends_on.
func (cfr *CdnFrontdoorRoute) Dependencies() terra.Dependencies {
	return cfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorRoute].
func (cfr *CdnFrontdoorRoute) LifecycleManagement() *terra.Lifecycle {
	return cfr.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorRoute].
func (cfr *CdnFrontdoorRoute) Attributes() cdnFrontdoorRouteAttributes {
	return cdnFrontdoorRouteAttributes{ref: terra.ReferenceResource(cfr)}
}

// ImportState imports the given attribute values into [CdnFrontdoorRoute]'s state.
func (cfr *CdnFrontdoorRoute) ImportState(av io.Reader) error {
	cfr.state = &cdnFrontdoorRouteState{}
	if err := json.NewDecoder(av).Decode(cfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfr.Type(), cfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorRoute] has state.
func (cfr *CdnFrontdoorRoute) State() (*cdnFrontdoorRouteState, bool) {
	return cfr.state, cfr.state != nil
}

// StateMust returns the state for [CdnFrontdoorRoute]. Panics if the state is nil.
func (cfr *CdnFrontdoorRoute) StateMust() *cdnFrontdoorRouteState {
	if cfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfr.Type(), cfr.LocalName()))
	}
	return cfr.state
}

// CdnFrontdoorRouteArgs contains the configurations for azurerm_cdn_frontdoor_route.
type CdnFrontdoorRouteArgs struct {
	// CdnFrontdoorCustomDomainIds: set of string, optional
	CdnFrontdoorCustomDomainIds terra.SetValue[terra.StringValue] `hcl:"cdn_frontdoor_custom_domain_ids,attr"`
	// CdnFrontdoorEndpointId: string, required
	CdnFrontdoorEndpointId terra.StringValue `hcl:"cdn_frontdoor_endpoint_id,attr" validate:"required"`
	// CdnFrontdoorOriginGroupId: string, required
	CdnFrontdoorOriginGroupId terra.StringValue `hcl:"cdn_frontdoor_origin_group_id,attr" validate:"required"`
	// CdnFrontdoorOriginIds: list of string, required
	CdnFrontdoorOriginIds terra.ListValue[terra.StringValue] `hcl:"cdn_frontdoor_origin_ids,attr" validate:"required"`
	// CdnFrontdoorOriginPath: string, optional
	CdnFrontdoorOriginPath terra.StringValue `hcl:"cdn_frontdoor_origin_path,attr"`
	// CdnFrontdoorRuleSetIds: set of string, optional
	CdnFrontdoorRuleSetIds terra.SetValue[terra.StringValue] `hcl:"cdn_frontdoor_rule_set_ids,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// ForwardingProtocol: string, optional
	ForwardingProtocol terra.StringValue `hcl:"forwarding_protocol,attr"`
	// HttpsRedirectEnabled: bool, optional
	HttpsRedirectEnabled terra.BoolValue `hcl:"https_redirect_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkToDefaultDomain: bool, optional
	LinkToDefaultDomain terra.BoolValue `hcl:"link_to_default_domain,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PatternsToMatch: list of string, required
	PatternsToMatch terra.ListValue[terra.StringValue] `hcl:"patterns_to_match,attr" validate:"required"`
	// SupportedProtocols: set of string, required
	SupportedProtocols terra.SetValue[terra.StringValue] `hcl:"supported_protocols,attr" validate:"required"`
	// Cache: optional
	Cache *cdnfrontdoorroute.Cache `hcl:"cache,block"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorroute.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorRouteAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorCustomDomainIds returns a reference to field cdn_frontdoor_custom_domain_ids of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) CdnFrontdoorCustomDomainIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfr.ref.Append("cdn_frontdoor_custom_domain_ids"))
}

// CdnFrontdoorEndpointId returns a reference to field cdn_frontdoor_endpoint_id of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) CdnFrontdoorEndpointId() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("cdn_frontdoor_endpoint_id"))
}

// CdnFrontdoorOriginGroupId returns a reference to field cdn_frontdoor_origin_group_id of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) CdnFrontdoorOriginGroupId() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("cdn_frontdoor_origin_group_id"))
}

// CdnFrontdoorOriginIds returns a reference to field cdn_frontdoor_origin_ids of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) CdnFrontdoorOriginIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfr.ref.Append("cdn_frontdoor_origin_ids"))
}

// CdnFrontdoorOriginPath returns a reference to field cdn_frontdoor_origin_path of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) CdnFrontdoorOriginPath() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("cdn_frontdoor_origin_path"))
}

// CdnFrontdoorRuleSetIds returns a reference to field cdn_frontdoor_rule_set_ids of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) CdnFrontdoorRuleSetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfr.ref.Append("cdn_frontdoor_rule_set_ids"))
}

// Enabled returns a reference to field enabled of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("enabled"))
}

// ForwardingProtocol returns a reference to field forwarding_protocol of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) ForwardingProtocol() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("forwarding_protocol"))
}

// HttpsRedirectEnabled returns a reference to field https_redirect_enabled of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) HttpsRedirectEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("https_redirect_enabled"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("id"))
}

// LinkToDefaultDomain returns a reference to field link_to_default_domain of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) LinkToDefaultDomain() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("link_to_default_domain"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("name"))
}

// PatternsToMatch returns a reference to field patterns_to_match of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) PatternsToMatch() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfr.ref.Append("patterns_to_match"))
}

// SupportedProtocols returns a reference to field supported_protocols of azurerm_cdn_frontdoor_route.
func (cfr cdnFrontdoorRouteAttributes) SupportedProtocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfr.ref.Append("supported_protocols"))
}

func (cfr cdnFrontdoorRouteAttributes) Cache() terra.ListValue[cdnfrontdoorroute.CacheAttributes] {
	return terra.ReferenceAsList[cdnfrontdoorroute.CacheAttributes](cfr.ref.Append("cache"))
}

func (cfr cdnFrontdoorRouteAttributes) Timeouts() cdnfrontdoorroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorroute.TimeoutsAttributes](cfr.ref.Append("timeouts"))
}

type cdnFrontdoorRouteState struct {
	CdnFrontdoorCustomDomainIds []string                         `json:"cdn_frontdoor_custom_domain_ids"`
	CdnFrontdoorEndpointId      string                           `json:"cdn_frontdoor_endpoint_id"`
	CdnFrontdoorOriginGroupId   string                           `json:"cdn_frontdoor_origin_group_id"`
	CdnFrontdoorOriginIds       []string                         `json:"cdn_frontdoor_origin_ids"`
	CdnFrontdoorOriginPath      string                           `json:"cdn_frontdoor_origin_path"`
	CdnFrontdoorRuleSetIds      []string                         `json:"cdn_frontdoor_rule_set_ids"`
	Enabled                     bool                             `json:"enabled"`
	ForwardingProtocol          string                           `json:"forwarding_protocol"`
	HttpsRedirectEnabled        bool                             `json:"https_redirect_enabled"`
	Id                          string                           `json:"id"`
	LinkToDefaultDomain         bool                             `json:"link_to_default_domain"`
	Name                        string                           `json:"name"`
	PatternsToMatch             []string                         `json:"patterns_to_match"`
	SupportedProtocols          []string                         `json:"supported_protocols"`
	Cache                       []cdnfrontdoorroute.CacheState   `json:"cache"`
	Timeouts                    *cdnfrontdoorroute.TimeoutsState `json:"timeouts"`
}
