// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datatrafficmanagerprofile "github.com/golingon/terraproviders/azurerm/3.52.0/datatrafficmanagerprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataTrafficManagerProfile creates a new instance of [DataTrafficManagerProfile].
func NewDataTrafficManagerProfile(name string, args DataTrafficManagerProfileArgs) *DataTrafficManagerProfile {
	return &DataTrafficManagerProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTrafficManagerProfile)(nil)

// DataTrafficManagerProfile represents the Terraform data resource azurerm_traffic_manager_profile.
type DataTrafficManagerProfile struct {
	Name string
	Args DataTrafficManagerProfileArgs
}

// DataSource returns the Terraform object type for [DataTrafficManagerProfile].
func (tmp *DataTrafficManagerProfile) DataSource() string {
	return "azurerm_traffic_manager_profile"
}

// LocalName returns the local name for [DataTrafficManagerProfile].
func (tmp *DataTrafficManagerProfile) LocalName() string {
	return tmp.Name
}

// Configuration returns the configuration (args) for [DataTrafficManagerProfile].
func (tmp *DataTrafficManagerProfile) Configuration() interface{} {
	return tmp.Args
}

// Attributes returns the attributes for [DataTrafficManagerProfile].
func (tmp *DataTrafficManagerProfile) Attributes() dataTrafficManagerProfileAttributes {
	return dataTrafficManagerProfileAttributes{ref: terra.ReferenceDataResource(tmp)}
}

// DataTrafficManagerProfileArgs contains the configurations for azurerm_traffic_manager_profile.
type DataTrafficManagerProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TrafficViewEnabled: bool, optional
	TrafficViewEnabled terra.BoolValue `hcl:"traffic_view_enabled,attr"`
	// DnsConfig: min=0
	DnsConfig []datatrafficmanagerprofile.DnsConfig `hcl:"dns_config,block" validate:"min=0"`
	// MonitorConfig: min=0
	MonitorConfig []datatrafficmanagerprofile.MonitorConfig `hcl:"monitor_config,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datatrafficmanagerprofile.Timeouts `hcl:"timeouts,block"`
}
type dataTrafficManagerProfileAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("name"))
}

// ProfileStatus returns a reference to field profile_status of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) ProfileStatus() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("profile_status"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tmp.ref.Append("tags"))
}

// TrafficRoutingMethod returns a reference to field traffic_routing_method of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) TrafficRoutingMethod() terra.StringValue {
	return terra.ReferenceAsString(tmp.ref.Append("traffic_routing_method"))
}

// TrafficViewEnabled returns a reference to field traffic_view_enabled of azurerm_traffic_manager_profile.
func (tmp dataTrafficManagerProfileAttributes) TrafficViewEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(tmp.ref.Append("traffic_view_enabled"))
}

func (tmp dataTrafficManagerProfileAttributes) DnsConfig() terra.ListValue[datatrafficmanagerprofile.DnsConfigAttributes] {
	return terra.ReferenceAsList[datatrafficmanagerprofile.DnsConfigAttributes](tmp.ref.Append("dns_config"))
}

func (tmp dataTrafficManagerProfileAttributes) MonitorConfig() terra.ListValue[datatrafficmanagerprofile.MonitorConfigAttributes] {
	return terra.ReferenceAsList[datatrafficmanagerprofile.MonitorConfigAttributes](tmp.ref.Append("monitor_config"))
}

func (tmp dataTrafficManagerProfileAttributes) Timeouts() datatrafficmanagerprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datatrafficmanagerprofile.TimeoutsAttributes](tmp.ref.Append("timeouts"))
}
