// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	trafficmanagerprofile "github.com/golingon/terraproviders/azurerm/3.49.0/trafficmanagerprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTrafficManagerProfile creates a new instance of [TrafficManagerProfile].
func NewTrafficManagerProfile(name string, args TrafficManagerProfileArgs) *TrafficManagerProfile {
	return &TrafficManagerProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TrafficManagerProfile)(nil)

// TrafficManagerProfile represents the Terraform resource azurerm_traffic_manager_profile.
type TrafficManagerProfile struct {
	Name      string
	Args      TrafficManagerProfileArgs
	state     *trafficManagerProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TrafficManagerProfile].
func (tmp *TrafficManagerProfile) Type() string {
	return "azurerm_traffic_manager_profile"
}

// LocalName returns the local name for [TrafficManagerProfile].
func (tmp *TrafficManagerProfile) LocalName() string {
	return tmp.Name
}

// Configuration returns the configuration (args) for [TrafficManagerProfile].
func (tmp *TrafficManagerProfile) Configuration() interface{} {
	return tmp.Args
}

// DependOn is used for other resources to depend on [TrafficManagerProfile].
func (tmp *TrafficManagerProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(tmp)
}

// Dependencies returns the list of resources [TrafficManagerProfile] depends_on.
func (tmp *TrafficManagerProfile) Dependencies() terra.Dependencies {
	return tmp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TrafficManagerProfile].
func (tmp *TrafficManagerProfile) LifecycleManagement() *terra.Lifecycle {
	return tmp.Lifecycle
}

// Attributes returns the attributes for [TrafficManagerProfile].
func (tmp *TrafficManagerProfile) Attributes() trafficManagerProfileAttributes {
	return trafficManagerProfileAttributes{ref: terra.ReferenceResource(tmp)}
}

// ImportState imports the given attribute values into [TrafficManagerProfile]'s state.
func (tmp *TrafficManagerProfile) ImportState(av io.Reader) error {
	tmp.state = &trafficManagerProfileState{}
	if err := json.NewDecoder(av).Decode(tmp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tmp.Type(), tmp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TrafficManagerProfile] has state.
func (tmp *TrafficManagerProfile) State() (*trafficManagerProfileState, bool) {
	return tmp.state, tmp.state != nil
}

// StateMust returns the state for [TrafficManagerProfile]. Panics if the state is nil.
func (tmp *TrafficManagerProfile) StateMust() *trafficManagerProfileState {
	if tmp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tmp.Type(), tmp.LocalName()))
	}
	return tmp.state
}

// TrafficManagerProfileArgs contains the configurations for azurerm_traffic_manager_profile.
type TrafficManagerProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxReturn: number, optional
	MaxReturn terra.NumberValue `hcl:"max_return,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileStatus: string, optional
	ProfileStatus terra.StringValue `hcl:"profile_status,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TrafficRoutingMethod: string, required
	TrafficRoutingMethod terra.StringValue `hcl:"traffic_routing_method,attr" validate:"required"`
	// TrafficViewEnabled: bool, optional
	TrafficViewEnabled terra.BoolValue `hcl:"traffic_view_enabled,attr"`
	// DnsConfig: required
	DnsConfig *trafficmanagerprofile.DnsConfig `hcl:"dns_config,block" validate:"required"`
	// MonitorConfig: required
	MonitorConfig *trafficmanagerprofile.MonitorConfig `hcl:"monitor_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *trafficmanagerprofile.Timeouts `hcl:"timeouts,block"`
}
type trafficManagerProfileAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("id"))
}

// MaxReturn returns a reference to field max_return of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) MaxReturn() terra.NumberValue {
	return terra.ReferenceAsNumber(tmp.ref.Append("max_return"))
}

// Name returns a reference to field name of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("name"))
}

// ProfileStatus returns a reference to field profile_status of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) ProfileStatus() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("profile_status"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tmp.ref.Append("tags"))
}

// TrafficRoutingMethod returns a reference to field traffic_routing_method of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) TrafficRoutingMethod() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("traffic_routing_method"))
}

// TrafficViewEnabled returns a reference to field traffic_view_enabled of azurerm_traffic_manager_profile.
func (tmp trafficManagerProfileAttributes) TrafficViewEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(tmp.ref.Append("traffic_view_enabled"))
}

func (tmp trafficManagerProfileAttributes) DnsConfig() terra.ListValue[trafficmanagerprofile.DnsConfigAttributes] {
	return terra.ReferenceAsList[trafficmanagerprofile.DnsConfigAttributes](tmp.ref.Append("dns_config"))
}

func (tmp trafficManagerProfileAttributes) MonitorConfig() terra.ListValue[trafficmanagerprofile.MonitorConfigAttributes] {
	return terra.ReferenceAsList[trafficmanagerprofile.MonitorConfigAttributes](tmp.ref.Append("monitor_config"))
}

func (tmp trafficManagerProfileAttributes) Timeouts() trafficmanagerprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[trafficmanagerprofile.TimeoutsAttributes](tmp.ref.Append("timeouts"))
}

type trafficManagerProfileState struct {
	Fqdn                 string                                     `json:"fqdn"`
	Id                   string                                     `json:"id"`
	MaxReturn            float64                                    `json:"max_return"`
	Name                 string                                     `json:"name"`
	ProfileStatus        string                                     `json:"profile_status"`
	ResourceGroupName    string                                     `json:"resource_group_name"`
	Tags                 map[string]string                          `json:"tags"`
	TrafficRoutingMethod string                                     `json:"traffic_routing_method"`
	TrafficViewEnabled   bool                                       `json:"traffic_view_enabled"`
	DnsConfig            []trafficmanagerprofile.DnsConfigState     `json:"dns_config"`
	MonitorConfig        []trafficmanagerprofile.MonitorConfigState `json:"monitor_config"`
	Timeouts             *trafficmanagerprofile.TimeoutsState       `json:"timeouts"`
}
