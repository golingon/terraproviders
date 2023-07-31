// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorcustomdomain "github.com/golingon/terraproviders/azurerm/3.67.0/cdnfrontdoorcustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorCustomDomain creates a new instance of [CdnFrontdoorCustomDomain].
func NewCdnFrontdoorCustomDomain(name string, args CdnFrontdoorCustomDomainArgs) *CdnFrontdoorCustomDomain {
	return &CdnFrontdoorCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorCustomDomain)(nil)

// CdnFrontdoorCustomDomain represents the Terraform resource azurerm_cdn_frontdoor_custom_domain.
type CdnFrontdoorCustomDomain struct {
	Name      string
	Args      CdnFrontdoorCustomDomainArgs
	state     *cdnFrontdoorCustomDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorCustomDomain].
func (cfcd *CdnFrontdoorCustomDomain) Type() string {
	return "azurerm_cdn_frontdoor_custom_domain"
}

// LocalName returns the local name for [CdnFrontdoorCustomDomain].
func (cfcd *CdnFrontdoorCustomDomain) LocalName() string {
	return cfcd.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorCustomDomain].
func (cfcd *CdnFrontdoorCustomDomain) Configuration() interface{} {
	return cfcd.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorCustomDomain].
func (cfcd *CdnFrontdoorCustomDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cfcd)
}

// Dependencies returns the list of resources [CdnFrontdoorCustomDomain] depends_on.
func (cfcd *CdnFrontdoorCustomDomain) Dependencies() terra.Dependencies {
	return cfcd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorCustomDomain].
func (cfcd *CdnFrontdoorCustomDomain) LifecycleManagement() *terra.Lifecycle {
	return cfcd.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorCustomDomain].
func (cfcd *CdnFrontdoorCustomDomain) Attributes() cdnFrontdoorCustomDomainAttributes {
	return cdnFrontdoorCustomDomainAttributes{ref: terra.ReferenceResource(cfcd)}
}

// ImportState imports the given attribute values into [CdnFrontdoorCustomDomain]'s state.
func (cfcd *CdnFrontdoorCustomDomain) ImportState(av io.Reader) error {
	cfcd.state = &cdnFrontdoorCustomDomainState{}
	if err := json.NewDecoder(av).Decode(cfcd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfcd.Type(), cfcd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorCustomDomain] has state.
func (cfcd *CdnFrontdoorCustomDomain) State() (*cdnFrontdoorCustomDomainState, bool) {
	return cfcd.state, cfcd.state != nil
}

// StateMust returns the state for [CdnFrontdoorCustomDomain]. Panics if the state is nil.
func (cfcd *CdnFrontdoorCustomDomain) StateMust() *cdnFrontdoorCustomDomainState {
	if cfcd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfcd.Type(), cfcd.LocalName()))
	}
	return cfcd.state
}

// CdnFrontdoorCustomDomainArgs contains the configurations for azurerm_cdn_frontdoor_custom_domain.
type CdnFrontdoorCustomDomainArgs struct {
	// CdnFrontdoorProfileId: string, required
	CdnFrontdoorProfileId terra.StringValue `hcl:"cdn_frontdoor_profile_id,attr" validate:"required"`
	// DnsZoneId: string, optional
	DnsZoneId terra.StringValue `hcl:"dns_zone_id,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorcustomdomain.Timeouts `hcl:"timeouts,block"`
	// Tls: required
	Tls *cdnfrontdoorcustomdomain.Tls `hcl:"tls,block" validate:"required"`
}
type cdnFrontdoorCustomDomainAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("cdn_frontdoor_profile_id"))
}

// DnsZoneId returns a reference to field dns_zone_id of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) DnsZoneId() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("dns_zone_id"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("expiration_date"))
}

// HostName returns a reference to field host_name of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("name"))
}

// ValidationToken returns a reference to field validation_token of azurerm_cdn_frontdoor_custom_domain.
func (cfcd cdnFrontdoorCustomDomainAttributes) ValidationToken() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("validation_token"))
}

func (cfcd cdnFrontdoorCustomDomainAttributes) Timeouts() cdnfrontdoorcustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorcustomdomain.TimeoutsAttributes](cfcd.ref.Append("timeouts"))
}

func (cfcd cdnFrontdoorCustomDomainAttributes) Tls() terra.ListValue[cdnfrontdoorcustomdomain.TlsAttributes] {
	return terra.ReferenceAsList[cdnfrontdoorcustomdomain.TlsAttributes](cfcd.ref.Append("tls"))
}

type cdnFrontdoorCustomDomainState struct {
	CdnFrontdoorProfileId string                                  `json:"cdn_frontdoor_profile_id"`
	DnsZoneId             string                                  `json:"dns_zone_id"`
	ExpirationDate        string                                  `json:"expiration_date"`
	HostName              string                                  `json:"host_name"`
	Id                    string                                  `json:"id"`
	Name                  string                                  `json:"name"`
	ValidationToken       string                                  `json:"validation_token"`
	Timeouts              *cdnfrontdoorcustomdomain.TimeoutsState `json:"timeouts"`
	Tls                   []cdnfrontdoorcustomdomain.TlsState     `json:"tls"`
}
