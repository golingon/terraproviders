// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoororigingroup "github.com/golingon/terraproviders/azurerm/3.69.0/cdnfrontdoororigingroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorOriginGroup creates a new instance of [CdnFrontdoorOriginGroup].
func NewCdnFrontdoorOriginGroup(name string, args CdnFrontdoorOriginGroupArgs) *CdnFrontdoorOriginGroup {
	return &CdnFrontdoorOriginGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorOriginGroup)(nil)

// CdnFrontdoorOriginGroup represents the Terraform resource azurerm_cdn_frontdoor_origin_group.
type CdnFrontdoorOriginGroup struct {
	Name      string
	Args      CdnFrontdoorOriginGroupArgs
	state     *cdnFrontdoorOriginGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorOriginGroup].
func (cfog *CdnFrontdoorOriginGroup) Type() string {
	return "azurerm_cdn_frontdoor_origin_group"
}

// LocalName returns the local name for [CdnFrontdoorOriginGroup].
func (cfog *CdnFrontdoorOriginGroup) LocalName() string {
	return cfog.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorOriginGroup].
func (cfog *CdnFrontdoorOriginGroup) Configuration() interface{} {
	return cfog.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorOriginGroup].
func (cfog *CdnFrontdoorOriginGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cfog)
}

// Dependencies returns the list of resources [CdnFrontdoorOriginGroup] depends_on.
func (cfog *CdnFrontdoorOriginGroup) Dependencies() terra.Dependencies {
	return cfog.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorOriginGroup].
func (cfog *CdnFrontdoorOriginGroup) LifecycleManagement() *terra.Lifecycle {
	return cfog.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorOriginGroup].
func (cfog *CdnFrontdoorOriginGroup) Attributes() cdnFrontdoorOriginGroupAttributes {
	return cdnFrontdoorOriginGroupAttributes{ref: terra.ReferenceResource(cfog)}
}

// ImportState imports the given attribute values into [CdnFrontdoorOriginGroup]'s state.
func (cfog *CdnFrontdoorOriginGroup) ImportState(av io.Reader) error {
	cfog.state = &cdnFrontdoorOriginGroupState{}
	if err := json.NewDecoder(av).Decode(cfog.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfog.Type(), cfog.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorOriginGroup] has state.
func (cfog *CdnFrontdoorOriginGroup) State() (*cdnFrontdoorOriginGroupState, bool) {
	return cfog.state, cfog.state != nil
}

// StateMust returns the state for [CdnFrontdoorOriginGroup]. Panics if the state is nil.
func (cfog *CdnFrontdoorOriginGroup) StateMust() *cdnFrontdoorOriginGroupState {
	if cfog.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfog.Type(), cfog.LocalName()))
	}
	return cfog.state
}

// CdnFrontdoorOriginGroupArgs contains the configurations for azurerm_cdn_frontdoor_origin_group.
type CdnFrontdoorOriginGroupArgs struct {
	// CdnFrontdoorProfileId: string, required
	CdnFrontdoorProfileId terra.StringValue `hcl:"cdn_frontdoor_profile_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RestoreTrafficTimeToHealedOrNewEndpointInMinutes: number, optional
	RestoreTrafficTimeToHealedOrNewEndpointInMinutes terra.NumberValue `hcl:"restore_traffic_time_to_healed_or_new_endpoint_in_minutes,attr"`
	// SessionAffinityEnabled: bool, optional
	SessionAffinityEnabled terra.BoolValue `hcl:"session_affinity_enabled,attr"`
	// HealthProbe: optional
	HealthProbe *cdnfrontdoororigingroup.HealthProbe `hcl:"health_probe,block"`
	// LoadBalancing: required
	LoadBalancing *cdnfrontdoororigingroup.LoadBalancing `hcl:"load_balancing,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cdnfrontdoororigingroup.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorOriginGroupAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_origin_group.
func (cfog cdnFrontdoorOriginGroupAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("cdn_frontdoor_profile_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_origin_group.
func (cfog cdnFrontdoorOriginGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_origin_group.
func (cfog cdnFrontdoorOriginGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfog.ref.Append("name"))
}

// RestoreTrafficTimeToHealedOrNewEndpointInMinutes returns a reference to field restore_traffic_time_to_healed_or_new_endpoint_in_minutes of azurerm_cdn_frontdoor_origin_group.
func (cfog cdnFrontdoorOriginGroupAttributes) RestoreTrafficTimeToHealedOrNewEndpointInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(cfog.ref.Append("restore_traffic_time_to_healed_or_new_endpoint_in_minutes"))
}

// SessionAffinityEnabled returns a reference to field session_affinity_enabled of azurerm_cdn_frontdoor_origin_group.
func (cfog cdnFrontdoorOriginGroupAttributes) SessionAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfog.ref.Append("session_affinity_enabled"))
}

func (cfog cdnFrontdoorOriginGroupAttributes) HealthProbe() terra.ListValue[cdnfrontdoororigingroup.HealthProbeAttributes] {
	return terra.ReferenceAsList[cdnfrontdoororigingroup.HealthProbeAttributes](cfog.ref.Append("health_probe"))
}

func (cfog cdnFrontdoorOriginGroupAttributes) LoadBalancing() terra.ListValue[cdnfrontdoororigingroup.LoadBalancingAttributes] {
	return terra.ReferenceAsList[cdnfrontdoororigingroup.LoadBalancingAttributes](cfog.ref.Append("load_balancing"))
}

func (cfog cdnFrontdoorOriginGroupAttributes) Timeouts() cdnfrontdoororigingroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoororigingroup.TimeoutsAttributes](cfog.ref.Append("timeouts"))
}

type cdnFrontdoorOriginGroupState struct {
	CdnFrontdoorProfileId                            string                                       `json:"cdn_frontdoor_profile_id"`
	Id                                               string                                       `json:"id"`
	Name                                             string                                       `json:"name"`
	RestoreTrafficTimeToHealedOrNewEndpointInMinutes float64                                      `json:"restore_traffic_time_to_healed_or_new_endpoint_in_minutes"`
	SessionAffinityEnabled                           bool                                         `json:"session_affinity_enabled"`
	HealthProbe                                      []cdnfrontdoororigingroup.HealthProbeState   `json:"health_probe"`
	LoadBalancing                                    []cdnfrontdoororigingroup.LoadBalancingState `json:"load_balancing"`
	Timeouts                                         *cdnfrontdoororigingroup.TimeoutsState       `json:"timeouts"`
}
