// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_custom_domain

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_cdn_frontdoor_custom_domain.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acfcd *DataSource) DataSource() string {
	return "azurerm_cdn_frontdoor_custom_domain"
}

// LocalName returns the local name for [DataSource].
func (acfcd *DataSource) LocalName() string {
	return acfcd.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acfcd *DataSource) Configuration() interface{} {
	return acfcd.Args
}

// Attributes returns the attributes for [DataSource].
func (acfcd *DataSource) Attributes() dataAzurermCdnFrontdoorCustomDomainAttributes {
	return dataAzurermCdnFrontdoorCustomDomainAttributes{ref: terra.ReferenceDataSource(acfcd)}
}

// DataArgs contains the configurations for azurerm_cdn_frontdoor_custom_domain.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermCdnFrontdoorCustomDomainAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("cdn_frontdoor_profile_id"))
}

// DnsZoneId returns a reference to field dns_zone_id of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) DnsZoneId() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("dns_zone_id"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("expiration_date"))
}

// HostName returns a reference to field host_name of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("resource_group_name"))
}

// ValidationToken returns a reference to field validation_token of azurerm_cdn_frontdoor_custom_domain.
func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) ValidationToken() terra.StringValue {
	return terra.ReferenceAsString(acfcd.ref.Append("validation_token"))
}

func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) Tls() terra.ListValue[DataTlsAttributes] {
	return terra.ReferenceAsList[DataTlsAttributes](acfcd.ref.Append("tls"))
}

func (acfcd dataAzurermCdnFrontdoorCustomDomainAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](acfcd.ref.Append("timeouts"))
}
