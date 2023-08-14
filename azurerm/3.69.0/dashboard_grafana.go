// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dashboardgrafana "github.com/golingon/terraproviders/azurerm/3.69.0/dashboardgrafana"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDashboardGrafana creates a new instance of [DashboardGrafana].
func NewDashboardGrafana(name string, args DashboardGrafanaArgs) *DashboardGrafana {
	return &DashboardGrafana{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DashboardGrafana)(nil)

// DashboardGrafana represents the Terraform resource azurerm_dashboard_grafana.
type DashboardGrafana struct {
	Name      string
	Args      DashboardGrafanaArgs
	state     *dashboardGrafanaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DashboardGrafana].
func (dg *DashboardGrafana) Type() string {
	return "azurerm_dashboard_grafana"
}

// LocalName returns the local name for [DashboardGrafana].
func (dg *DashboardGrafana) LocalName() string {
	return dg.Name
}

// Configuration returns the configuration (args) for [DashboardGrafana].
func (dg *DashboardGrafana) Configuration() interface{} {
	return dg.Args
}

// DependOn is used for other resources to depend on [DashboardGrafana].
func (dg *DashboardGrafana) DependOn() terra.Reference {
	return terra.ReferenceResource(dg)
}

// Dependencies returns the list of resources [DashboardGrafana] depends_on.
func (dg *DashboardGrafana) Dependencies() terra.Dependencies {
	return dg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DashboardGrafana].
func (dg *DashboardGrafana) LifecycleManagement() *terra.Lifecycle {
	return dg.Lifecycle
}

// Attributes returns the attributes for [DashboardGrafana].
func (dg *DashboardGrafana) Attributes() dashboardGrafanaAttributes {
	return dashboardGrafanaAttributes{ref: terra.ReferenceResource(dg)}
}

// ImportState imports the given attribute values into [DashboardGrafana]'s state.
func (dg *DashboardGrafana) ImportState(av io.Reader) error {
	dg.state = &dashboardGrafanaState{}
	if err := json.NewDecoder(av).Decode(dg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dg.Type(), dg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DashboardGrafana] has state.
func (dg *DashboardGrafana) State() (*dashboardGrafanaState, bool) {
	return dg.state, dg.state != nil
}

// StateMust returns the state for [DashboardGrafana]. Panics if the state is nil.
func (dg *DashboardGrafana) StateMust() *dashboardGrafanaState {
	if dg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dg.Type(), dg.LocalName()))
	}
	return dg.state
}

// DashboardGrafanaArgs contains the configurations for azurerm_dashboard_grafana.
type DashboardGrafanaArgs struct {
	// ApiKeyEnabled: bool, optional
	ApiKeyEnabled terra.BoolValue `hcl:"api_key_enabled,attr"`
	// AutoGeneratedDomainNameLabelScope: string, optional
	AutoGeneratedDomainNameLabelScope terra.StringValue `hcl:"auto_generated_domain_name_label_scope,attr"`
	// DeterministicOutboundIpEnabled: bool, optional
	DeterministicOutboundIpEnabled terra.BoolValue `hcl:"deterministic_outbound_ip_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundancyEnabled: bool, optional
	ZoneRedundancyEnabled terra.BoolValue `hcl:"zone_redundancy_enabled,attr"`
	// AzureMonitorWorkspaceIntegrations: min=0
	AzureMonitorWorkspaceIntegrations []dashboardgrafana.AzureMonitorWorkspaceIntegrations `hcl:"azure_monitor_workspace_integrations,block" validate:"min=0"`
	// Identity: optional
	Identity *dashboardgrafana.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *dashboardgrafana.Timeouts `hcl:"timeouts,block"`
}
type dashboardGrafanaAttributes struct {
	ref terra.Reference
}

// ApiKeyEnabled returns a reference to field api_key_enabled of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) ApiKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dg.ref.Append("api_key_enabled"))
}

// AutoGeneratedDomainNameLabelScope returns a reference to field auto_generated_domain_name_label_scope of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) AutoGeneratedDomainNameLabelScope() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("auto_generated_domain_name_label_scope"))
}

// DeterministicOutboundIpEnabled returns a reference to field deterministic_outbound_ip_enabled of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) DeterministicOutboundIpEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dg.ref.Append("deterministic_outbound_ip_enabled"))
}

// Endpoint returns a reference to field endpoint of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("endpoint"))
}

// GrafanaVersion returns a reference to field grafana_version of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) GrafanaVersion() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("grafana_version"))
}

// Id returns a reference to field id of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("name"))
}

// OutboundIp returns a reference to field outbound_ip of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) OutboundIp() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dg.ref.Append("outbound_ip"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dg.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dg.ref.Append("tags"))
}

// ZoneRedundancyEnabled returns a reference to field zone_redundancy_enabled of azurerm_dashboard_grafana.
func (dg dashboardGrafanaAttributes) ZoneRedundancyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dg.ref.Append("zone_redundancy_enabled"))
}

func (dg dashboardGrafanaAttributes) AzureMonitorWorkspaceIntegrations() terra.ListValue[dashboardgrafana.AzureMonitorWorkspaceIntegrationsAttributes] {
	return terra.ReferenceAsList[dashboardgrafana.AzureMonitorWorkspaceIntegrationsAttributes](dg.ref.Append("azure_monitor_workspace_integrations"))
}

func (dg dashboardGrafanaAttributes) Identity() terra.ListValue[dashboardgrafana.IdentityAttributes] {
	return terra.ReferenceAsList[dashboardgrafana.IdentityAttributes](dg.ref.Append("identity"))
}

func (dg dashboardGrafanaAttributes) Timeouts() dashboardgrafana.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dashboardgrafana.TimeoutsAttributes](dg.ref.Append("timeouts"))
}

type dashboardGrafanaState struct {
	ApiKeyEnabled                     bool                                                      `json:"api_key_enabled"`
	AutoGeneratedDomainNameLabelScope string                                                    `json:"auto_generated_domain_name_label_scope"`
	DeterministicOutboundIpEnabled    bool                                                      `json:"deterministic_outbound_ip_enabled"`
	Endpoint                          string                                                    `json:"endpoint"`
	GrafanaVersion                    string                                                    `json:"grafana_version"`
	Id                                string                                                    `json:"id"`
	Location                          string                                                    `json:"location"`
	Name                              string                                                    `json:"name"`
	OutboundIp                        []string                                                  `json:"outbound_ip"`
	PublicNetworkAccessEnabled        bool                                                      `json:"public_network_access_enabled"`
	ResourceGroupName                 string                                                    `json:"resource_group_name"`
	Sku                               string                                                    `json:"sku"`
	Tags                              map[string]string                                         `json:"tags"`
	ZoneRedundancyEnabled             bool                                                      `json:"zone_redundancy_enabled"`
	AzureMonitorWorkspaceIntegrations []dashboardgrafana.AzureMonitorWorkspaceIntegrationsState `json:"azure_monitor_workspace_integrations"`
	Identity                          []dashboardgrafana.IdentityState                          `json:"identity"`
	Timeouts                          *dashboardgrafana.TimeoutsState                           `json:"timeouts"`
}
