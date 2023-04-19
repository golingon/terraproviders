// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorruleset "github.com/golingon/terraproviders/azurerm/3.52.0/datacdnfrontdoorruleset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorRuleSet creates a new instance of [DataCdnFrontdoorRuleSet].
func NewDataCdnFrontdoorRuleSet(name string, args DataCdnFrontdoorRuleSetArgs) *DataCdnFrontdoorRuleSet {
	return &DataCdnFrontdoorRuleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorRuleSet)(nil)

// DataCdnFrontdoorRuleSet represents the Terraform data resource azurerm_cdn_frontdoor_rule_set.
type DataCdnFrontdoorRuleSet struct {
	Name string
	Args DataCdnFrontdoorRuleSetArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorRuleSet].
func (cfrs *DataCdnFrontdoorRuleSet) DataSource() string {
	return "azurerm_cdn_frontdoor_rule_set"
}

// LocalName returns the local name for [DataCdnFrontdoorRuleSet].
func (cfrs *DataCdnFrontdoorRuleSet) LocalName() string {
	return cfrs.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorRuleSet].
func (cfrs *DataCdnFrontdoorRuleSet) Configuration() interface{} {
	return cfrs.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorRuleSet].
func (cfrs *DataCdnFrontdoorRuleSet) Attributes() dataCdnFrontdoorRuleSetAttributes {
	return dataCdnFrontdoorRuleSetAttributes{ref: terra.ReferenceDataResource(cfrs)}
}

// DataCdnFrontdoorRuleSetArgs contains the configurations for azurerm_cdn_frontdoor_rule_set.
type DataCdnFrontdoorRuleSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoorruleset.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorRuleSetAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_rule_set.
func (cfrs dataCdnFrontdoorRuleSetAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("cdn_frontdoor_profile_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_rule_set.
func (cfrs dataCdnFrontdoorRuleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_rule_set.
func (cfrs dataCdnFrontdoorRuleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_rule_set.
func (cfrs dataCdnFrontdoorRuleSetAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_rule_set.
func (cfrs dataCdnFrontdoorRuleSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("resource_group_name"))
}

func (cfrs dataCdnFrontdoorRuleSetAttributes) Timeouts() datacdnfrontdoorruleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoorruleset.TimeoutsAttributes](cfrs.ref.Append("timeouts"))
}
