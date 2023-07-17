// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorprofile "github.com/golingon/terraproviders/azurerm/3.65.0/datacdnfrontdoorprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorProfile creates a new instance of [DataCdnFrontdoorProfile].
func NewDataCdnFrontdoorProfile(name string, args DataCdnFrontdoorProfileArgs) *DataCdnFrontdoorProfile {
	return &DataCdnFrontdoorProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorProfile)(nil)

// DataCdnFrontdoorProfile represents the Terraform data resource azurerm_cdn_frontdoor_profile.
type DataCdnFrontdoorProfile struct {
	Name string
	Args DataCdnFrontdoorProfileArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorProfile].
func (cfp *DataCdnFrontdoorProfile) DataSource() string {
	return "azurerm_cdn_frontdoor_profile"
}

// LocalName returns the local name for [DataCdnFrontdoorProfile].
func (cfp *DataCdnFrontdoorProfile) LocalName() string {
	return cfp.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorProfile].
func (cfp *DataCdnFrontdoorProfile) Configuration() interface{} {
	return cfp.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorProfile].
func (cfp *DataCdnFrontdoorProfile) Attributes() dataCdnFrontdoorProfileAttributes {
	return dataCdnFrontdoorProfileAttributes{ref: terra.ReferenceDataResource(cfp)}
}

// DataCdnFrontdoorProfileArgs contains the configurations for azurerm_cdn_frontdoor_profile.
type DataCdnFrontdoorProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoorprofile.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorProfileAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("resource_group_name"))
}

// ResourceGuid returns a reference to field resource_guid of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) ResourceGuid() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("resource_guid"))
}

// ResponseTimeoutSeconds returns a reference to field response_timeout_seconds of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) ResponseTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cfp.ref.Append("response_timeout_seconds"))
}

// SkuName returns a reference to field sku_name of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_cdn_frontdoor_profile.
func (cfp dataCdnFrontdoorProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cfp.ref.Append("tags"))
}

func (cfp dataCdnFrontdoorProfileAttributes) Timeouts() datacdnfrontdoorprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoorprofile.TimeoutsAttributes](cfp.ref.Append("timeouts"))
}
