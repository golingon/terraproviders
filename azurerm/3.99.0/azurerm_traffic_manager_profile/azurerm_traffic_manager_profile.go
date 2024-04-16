// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_traffic_manager_profile

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_traffic_manager_profile.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermTrafficManagerProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (atmp *Resource) Type() string {
	return "azurerm_traffic_manager_profile"
}

// LocalName returns the local name for [Resource].
func (atmp *Resource) LocalName() string {
	return atmp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (atmp *Resource) Configuration() interface{} {
	return atmp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (atmp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(atmp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (atmp *Resource) Dependencies() terra.Dependencies {
	return atmp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (atmp *Resource) LifecycleManagement() *terra.Lifecycle {
	return atmp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (atmp *Resource) Attributes() azurermTrafficManagerProfileAttributes {
	return azurermTrafficManagerProfileAttributes{ref: terra.ReferenceResource(atmp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (atmp *Resource) ImportState(state io.Reader) error {
	atmp.state = &azurermTrafficManagerProfileState{}
	if err := json.NewDecoder(state).Decode(atmp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", atmp.Type(), atmp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (atmp *Resource) State() (*azurermTrafficManagerProfileState, bool) {
	return atmp.state, atmp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (atmp *Resource) StateMust() *azurermTrafficManagerProfileState {
	if atmp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", atmp.Type(), atmp.LocalName()))
	}
	return atmp.state
}

// Args contains the configurations for azurerm_traffic_manager_profile.
type Args struct {
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
	DnsConfig *DnsConfig `hcl:"dns_config,block" validate:"required"`
	// MonitorConfig: required
	MonitorConfig *MonitorConfig `hcl:"monitor_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermTrafficManagerProfileAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(atmp.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(atmp.ref.Append("id"))
}

// MaxReturn returns a reference to field max_return of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) MaxReturn() terra.NumberValue {
	return terra.ReferenceAsNumber(atmp.ref.Append("max_return"))
}

// Name returns a reference to field name of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(atmp.ref.Append("name"))
}

// ProfileStatus returns a reference to field profile_status of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) ProfileStatus() terra.StringValue {
	return terra.ReferenceAsString(atmp.ref.Append("profile_status"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(atmp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](atmp.ref.Append("tags"))
}

// TrafficRoutingMethod returns a reference to field traffic_routing_method of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) TrafficRoutingMethod() terra.StringValue {
	return terra.ReferenceAsString(atmp.ref.Append("traffic_routing_method"))
}

// TrafficViewEnabled returns a reference to field traffic_view_enabled of azurerm_traffic_manager_profile.
func (atmp azurermTrafficManagerProfileAttributes) TrafficViewEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(atmp.ref.Append("traffic_view_enabled"))
}

func (atmp azurermTrafficManagerProfileAttributes) DnsConfig() terra.ListValue[DnsConfigAttributes] {
	return terra.ReferenceAsList[DnsConfigAttributes](atmp.ref.Append("dns_config"))
}

func (atmp azurermTrafficManagerProfileAttributes) MonitorConfig() terra.ListValue[MonitorConfigAttributes] {
	return terra.ReferenceAsList[MonitorConfigAttributes](atmp.ref.Append("monitor_config"))
}

func (atmp azurermTrafficManagerProfileAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](atmp.ref.Append("timeouts"))
}

type azurermTrafficManagerProfileState struct {
	Fqdn                 string               `json:"fqdn"`
	Id                   string               `json:"id"`
	MaxReturn            float64              `json:"max_return"`
	Name                 string               `json:"name"`
	ProfileStatus        string               `json:"profile_status"`
	ResourceGroupName    string               `json:"resource_group_name"`
	Tags                 map[string]string    `json:"tags"`
	TrafficRoutingMethod string               `json:"traffic_routing_method"`
	TrafficViewEnabled   bool                 `json:"traffic_view_enabled"`
	DnsConfig            []DnsConfigState     `json:"dns_config"`
	MonitorConfig        []MonitorConfigState `json:"monitor_config"`
	Timeouts             *TimeoutsState       `json:"timeouts"`
}
