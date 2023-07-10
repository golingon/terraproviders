// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorcustomdomain "github.com/golingon/terraproviders/azurerm/3.64.0/datacdnfrontdoorcustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorCustomDomain creates a new instance of [DataCdnFrontdoorCustomDomain].
func NewDataCdnFrontdoorCustomDomain(name string, args DataCdnFrontdoorCustomDomainArgs) *DataCdnFrontdoorCustomDomain {
	return &DataCdnFrontdoorCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorCustomDomain)(nil)

// DataCdnFrontdoorCustomDomain represents the Terraform data resource azurerm_cdn_frontdoor_custom_domain.
type DataCdnFrontdoorCustomDomain struct {
	Name string
	Args DataCdnFrontdoorCustomDomainArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorCustomDomain].
func (cfcd *DataCdnFrontdoorCustomDomain) DataSource() string {
	return "azurerm_cdn_frontdoor_custom_domain"
}

// LocalName returns the local name for [DataCdnFrontdoorCustomDomain].
func (cfcd *DataCdnFrontdoorCustomDomain) LocalName() string {
	return cfcd.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorCustomDomain].
func (cfcd *DataCdnFrontdoorCustomDomain) Configuration() interface{} {
	return cfcd.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorCustomDomain].
func (cfcd *DataCdnFrontdoorCustomDomain) Attributes() dataCdnFrontdoorCustomDomainAttributes {
	return dataCdnFrontdoorCustomDomainAttributes{ref: terra.ReferenceDataResource(cfcd)}
}

// DataCdnFrontdoorCustomDomainArgs contains the configurations for azurerm_cdn_frontdoor_custom_domain.
type DataCdnFrontdoorCustomDomainArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tls: min=0
	Tls []datacdnfrontdoorcustomdomain.Tls `hcl:"tls,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoorcustomdomain.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorCustomDomainAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("cdn_frontdoor_profile_id"))
}

// DnsZoneId returns a reference to field dns_zone_id of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) DnsZoneId() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("dns_zone_id"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("expiration_date"))
}

// HostName returns a reference to field host_name of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("resource_group_name"))
}

// ValidationToken returns a reference to field validation_token of azurerm_cdn_frontdoor_custom_domain.
func (cfcd dataCdnFrontdoorCustomDomainAttributes) ValidationToken() terra.StringValue {
	return terra.ReferenceAsString(cfcd.ref.Append("validation_token"))
}

func (cfcd dataCdnFrontdoorCustomDomainAttributes) Tls() terra.ListValue[datacdnfrontdoorcustomdomain.TlsAttributes] {
	return terra.ReferenceAsList[datacdnfrontdoorcustomdomain.TlsAttributes](cfcd.ref.Append("tls"))
}

func (cfcd dataCdnFrontdoorCustomDomainAttributes) Timeouts() datacdnfrontdoorcustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoorcustomdomain.TimeoutsAttributes](cfcd.ref.Append("timeouts"))
}
