// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoororigingroup "github.com/golingon/terraproviders/azurerm/3.69.0/datacdnfrontdoororigingroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorOriginGroup creates a new instance of [DataCdnFrontdoorOriginGroup].
func NewDataCdnFrontdoorOriginGroup(name string, args DataCdnFrontdoorOriginGroupArgs) *DataCdnFrontdoorOriginGroup {
	return &DataCdnFrontdoorOriginGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorOriginGroup)(nil)

// DataCdnFrontdoorOriginGroup represents the Terraform data resource azurerm_cdn_frontdoor_origin_group.
type DataCdnFrontdoorOriginGroup struct {
	Name string
	Args DataCdnFrontdoorOriginGroupArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorOriginGroup].
func (cfog *DataCdnFrontdoorOriginGroup) DataSource() string {
	return "azurerm_cdn_frontdoor_origin_group"
}

// LocalName returns the local name for [DataCdnFrontdoorOriginGroup].
func (cfog *DataCdnFrontdoorOriginGroup) LocalName() string {
	return cfog.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorOriginGroup].
func (cfog *DataCdnFrontdoorOriginGroup) Configuration() interface{} {
	return cfog.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorOriginGroup].
func (cfog *DataCdnFrontdoorOriginGroup) Attributes() dataCdnFrontdoorOriginGroupAttributes {
	return dataCdnFrontdoorOriginGroupAttributes{ref: terra.ReferenceDataResource(cfog)}
}

// DataCdnFrontdoorOriginGroupArgs contains the configurations for azurerm_cdn_frontdoor_origin_group.
type DataCdnFrontdoorOriginGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// HealthProbe: min=0
	HealthProbe []datacdnfrontdoororigingroup.HealthProbe `hcl:"health_probe,block" validate:"min=0"`
	// LoadBalancing: min=0
	LoadBalancing []datacdnfrontdoororigingroup.LoadBalancing `hcl:"load_balancing,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoororigingroup.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorOriginGroupAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("cdn_frontdoor_profile_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("resource_group_name"))
}

// RestoreTrafficTimeToHealedOrNewEndpointInMinutes returns a reference to field restore_traffic_time_to_healed_or_new_endpoint_in_minutes of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) RestoreTrafficTimeToHealedOrNewEndpointInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(cfog.ref.Append("restore_traffic_time_to_healed_or_new_endpoint_in_minutes"))
}

// SessionAffinityEnabled returns a reference to field session_affinity_enabled of azurerm_cdn_frontdoor_origin_group.
func (cfog dataCdnFrontdoorOriginGroupAttributes) SessionAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfog.ref.Append("session_affinity_enabled"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) HealthProbe() terra.ListValue[datacdnfrontdoororigingroup.HealthProbeAttributes] {
	return terra.ReferenceAsList[datacdnfrontdoororigingroup.HealthProbeAttributes](cfog.ref.Append("health_probe"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) LoadBalancing() terra.ListValue[datacdnfrontdoororigingroup.LoadBalancingAttributes] {
	return terra.ReferenceAsList[datacdnfrontdoororigingroup.LoadBalancingAttributes](cfog.ref.Append("load_balancing"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) Timeouts() datacdnfrontdoororigingroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoororigingroup.TimeoutsAttributes](cfog.ref.Append("timeouts"))
}
