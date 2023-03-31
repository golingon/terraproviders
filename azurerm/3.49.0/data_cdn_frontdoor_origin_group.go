// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoororigingroup "github.com/golingon/terraproviders/azurerm/3.49.0/datacdnfrontdoororigingroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataCdnFrontdoorOriginGroup(name string, args DataCdnFrontdoorOriginGroupArgs) *DataCdnFrontdoorOriginGroup {
	return &DataCdnFrontdoorOriginGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorOriginGroup)(nil)

type DataCdnFrontdoorOriginGroup struct {
	Name string
	Args DataCdnFrontdoorOriginGroupArgs
}

func (cfog *DataCdnFrontdoorOriginGroup) DataSource() string {
	return "azurerm_cdn_frontdoor_origin_group"
}

func (cfog *DataCdnFrontdoorOriginGroup) LocalName() string {
	return cfog.Name
}

func (cfog *DataCdnFrontdoorOriginGroup) Configuration() interface{} {
	return cfog.Args
}

func (cfog *DataCdnFrontdoorOriginGroup) Attributes() dataCdnFrontdoorOriginGroupAttributes {
	return dataCdnFrontdoorOriginGroupAttributes{ref: terra.ReferenceDataResource(cfog)}
}

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

func (cfog dataCdnFrontdoorOriginGroupAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceString(cfog.ref.Append("cdn_frontdoor_profile_id"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cfog.ref.Append("id"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cfog.ref.Append("name"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceString(cfog.ref.Append("profile_name"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(cfog.ref.Append("resource_group_name"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) RestoreTrafficTimeToHealedOrNewEndpointInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(cfog.ref.Append("restore_traffic_time_to_healed_or_new_endpoint_in_minutes"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) SessionAffinityEnabled() terra.BoolValue {
	return terra.ReferenceBool(cfog.ref.Append("session_affinity_enabled"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) HealthProbe() terra.ListValue[datacdnfrontdoororigingroup.HealthProbeAttributes] {
	return terra.ReferenceList[datacdnfrontdoororigingroup.HealthProbeAttributes](cfog.ref.Append("health_probe"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) LoadBalancing() terra.ListValue[datacdnfrontdoororigingroup.LoadBalancingAttributes] {
	return terra.ReferenceList[datacdnfrontdoororigingroup.LoadBalancingAttributes](cfog.ref.Append("load_balancing"))
}

func (cfog dataCdnFrontdoorOriginGroupAttributes) Timeouts() datacdnfrontdoororigingroup.TimeoutsAttributes {
	return terra.ReferenceSingle[datacdnfrontdoororigingroup.TimeoutsAttributes](cfog.ref.Append("timeouts"))
}
