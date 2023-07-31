// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorroutedisablelinktodefaultdomain "github.com/golingon/terraproviders/azurerm/3.67.0/cdnfrontdoorroutedisablelinktodefaultdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorRouteDisableLinkToDefaultDomain creates a new instance of [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func NewCdnFrontdoorRouteDisableLinkToDefaultDomain(name string, args CdnFrontdoorRouteDisableLinkToDefaultDomainArgs) *CdnFrontdoorRouteDisableLinkToDefaultDomain {
	return &CdnFrontdoorRouteDisableLinkToDefaultDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorRouteDisableLinkToDefaultDomain)(nil)

// CdnFrontdoorRouteDisableLinkToDefaultDomain represents the Terraform resource azurerm_cdn_frontdoor_route_disable_link_to_default_domain.
type CdnFrontdoorRouteDisableLinkToDefaultDomain struct {
	Name      string
	Args      CdnFrontdoorRouteDisableLinkToDefaultDomainArgs
	state     *cdnFrontdoorRouteDisableLinkToDefaultDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) Type() string {
	return "azurerm_cdn_frontdoor_route_disable_link_to_default_domain"
}

// LocalName returns the local name for [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) LocalName() string {
	return cfrdltdd.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) Configuration() interface{} {
	return cfrdltdd.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cfrdltdd)
}

// Dependencies returns the list of resources [CdnFrontdoorRouteDisableLinkToDefaultDomain] depends_on.
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) Dependencies() terra.Dependencies {
	return cfrdltdd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) LifecycleManagement() *terra.Lifecycle {
	return cfrdltdd.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorRouteDisableLinkToDefaultDomain].
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) Attributes() cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes {
	return cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes{ref: terra.ReferenceResource(cfrdltdd)}
}

// ImportState imports the given attribute values into [CdnFrontdoorRouteDisableLinkToDefaultDomain]'s state.
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) ImportState(av io.Reader) error {
	cfrdltdd.state = &cdnFrontdoorRouteDisableLinkToDefaultDomainState{}
	if err := json.NewDecoder(av).Decode(cfrdltdd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfrdltdd.Type(), cfrdltdd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorRouteDisableLinkToDefaultDomain] has state.
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) State() (*cdnFrontdoorRouteDisableLinkToDefaultDomainState, bool) {
	return cfrdltdd.state, cfrdltdd.state != nil
}

// StateMust returns the state for [CdnFrontdoorRouteDisableLinkToDefaultDomain]. Panics if the state is nil.
func (cfrdltdd *CdnFrontdoorRouteDisableLinkToDefaultDomain) StateMust() *cdnFrontdoorRouteDisableLinkToDefaultDomainState {
	if cfrdltdd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfrdltdd.Type(), cfrdltdd.LocalName()))
	}
	return cfrdltdd.state
}

// CdnFrontdoorRouteDisableLinkToDefaultDomainArgs contains the configurations for azurerm_cdn_frontdoor_route_disable_link_to_default_domain.
type CdnFrontdoorRouteDisableLinkToDefaultDomainArgs struct {
	// CdnFrontdoorCustomDomainIds: list of string, required
	CdnFrontdoorCustomDomainIds terra.ListValue[terra.StringValue] `hcl:"cdn_frontdoor_custom_domain_ids,attr" validate:"required"`
	// CdnFrontdoorRouteId: string, required
	CdnFrontdoorRouteId terra.StringValue `hcl:"cdn_frontdoor_route_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorroutedisablelinktodefaultdomain.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorCustomDomainIds returns a reference to field cdn_frontdoor_custom_domain_ids of azurerm_cdn_frontdoor_route_disable_link_to_default_domain.
func (cfrdltdd cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes) CdnFrontdoorCustomDomainIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfrdltdd.ref.Append("cdn_frontdoor_custom_domain_ids"))
}

// CdnFrontdoorRouteId returns a reference to field cdn_frontdoor_route_id of azurerm_cdn_frontdoor_route_disable_link_to_default_domain.
func (cfrdltdd cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes) CdnFrontdoorRouteId() terra.StringValue {
	return terra.ReferenceAsString(cfrdltdd.ref.Append("cdn_frontdoor_route_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_route_disable_link_to_default_domain.
func (cfrdltdd cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfrdltdd.ref.Append("id"))
}

func (cfrdltdd cdnFrontdoorRouteDisableLinkToDefaultDomainAttributes) Timeouts() cdnfrontdoorroutedisablelinktodefaultdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorroutedisablelinktodefaultdomain.TimeoutsAttributes](cfrdltdd.ref.Append("timeouts"))
}

type cdnFrontdoorRouteDisableLinkToDefaultDomainState struct {
	CdnFrontdoorCustomDomainIds []string                                                   `json:"cdn_frontdoor_custom_domain_ids"`
	CdnFrontdoorRouteId         string                                                     `json:"cdn_frontdoor_route_id"`
	Id                          string                                                     `json:"id"`
	Timeouts                    *cdnfrontdoorroutedisablelinktodefaultdomain.TimeoutsState `json:"timeouts"`
}
