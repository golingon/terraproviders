// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_container_app_environment

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

// Resource represents the Terraform resource azurerm_container_app_environment.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermContainerAppEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acae *Resource) Type() string {
	return "azurerm_container_app_environment"
}

// LocalName returns the local name for [Resource].
func (acae *Resource) LocalName() string {
	return acae.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acae *Resource) Configuration() interface{} {
	return acae.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acae *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acae)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acae *Resource) Dependencies() terra.Dependencies {
	return acae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acae *Resource) LifecycleManagement() *terra.Lifecycle {
	return acae.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acae *Resource) Attributes() azurermContainerAppEnvironmentAttributes {
	return azurermContainerAppEnvironmentAttributes{ref: terra.ReferenceResource(acae)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acae *Resource) ImportState(state io.Reader) error {
	acae.state = &azurermContainerAppEnvironmentState{}
	if err := json.NewDecoder(state).Decode(acae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acae.Type(), acae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acae *Resource) State() (*azurermContainerAppEnvironmentState, bool) {
	return acae.state, acae.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acae *Resource) StateMust() *azurermContainerAppEnvironmentState {
	if acae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acae.Type(), acae.LocalName()))
	}
	return acae.state
}

// Args contains the configurations for azurerm_container_app_environment.
type Args struct {
	// DaprApplicationInsightsConnectionString: string, optional
	DaprApplicationInsightsConnectionString terra.StringValue `hcl:"dapr_application_insights_connection_string,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InfrastructureResourceGroupName: string, optional
	InfrastructureResourceGroupName terra.StringValue `hcl:"infrastructure_resource_group_name,attr"`
	// InfrastructureSubnetId: string, optional
	InfrastructureSubnetId terra.StringValue `hcl:"infrastructure_subnet_id,attr"`
	// InternalLoadBalancerEnabled: bool, optional
	InternalLoadBalancerEnabled terra.BoolValue `hcl:"internal_load_balancer_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LogAnalyticsWorkspaceId: string, optional
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundancyEnabled: bool, optional
	ZoneRedundancyEnabled terra.BoolValue `hcl:"zone_redundancy_enabled,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// WorkloadProfile: min=0
	WorkloadProfile []WorkloadProfile `hcl:"workload_profile,block" validate:"min=0"`
}

type azurermContainerAppEnvironmentAttributes struct {
	ref terra.Reference
}

// DaprApplicationInsightsConnectionString returns a reference to field dapr_application_insights_connection_string of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) DaprApplicationInsightsConnectionString() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("dapr_application_insights_connection_string"))
}

// DefaultDomain returns a reference to field default_domain of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) DefaultDomain() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("default_domain"))
}

// DockerBridgeCidr returns a reference to field docker_bridge_cidr of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) DockerBridgeCidr() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("docker_bridge_cidr"))
}

// Id returns a reference to field id of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("id"))
}

// InfrastructureResourceGroupName returns a reference to field infrastructure_resource_group_name of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) InfrastructureResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("infrastructure_resource_group_name"))
}

// InfrastructureSubnetId returns a reference to field infrastructure_subnet_id of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) InfrastructureSubnetId() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("infrastructure_subnet_id"))
}

// InternalLoadBalancerEnabled returns a reference to field internal_load_balancer_enabled of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) InternalLoadBalancerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(acae.ref.Append("internal_load_balancer_enabled"))
}

// Location returns a reference to field location of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("location"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("name"))
}

// PlatformReservedCidr returns a reference to field platform_reserved_cidr of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) PlatformReservedCidr() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("platform_reserved_cidr"))
}

// PlatformReservedDnsIpAddress returns a reference to field platform_reserved_dns_ip_address of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) PlatformReservedDnsIpAddress() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("platform_reserved_dns_ip_address"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("resource_group_name"))
}

// StaticIpAddress returns a reference to field static_ip_address of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) StaticIpAddress() terra.StringValue {
	return terra.ReferenceAsString(acae.ref.Append("static_ip_address"))
}

// Tags returns a reference to field tags of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acae.ref.Append("tags"))
}

// ZoneRedundancyEnabled returns a reference to field zone_redundancy_enabled of azurerm_container_app_environment.
func (acae azurermContainerAppEnvironmentAttributes) ZoneRedundancyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(acae.ref.Append("zone_redundancy_enabled"))
}

func (acae azurermContainerAppEnvironmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acae.ref.Append("timeouts"))
}

func (acae azurermContainerAppEnvironmentAttributes) WorkloadProfile() terra.SetValue[WorkloadProfileAttributes] {
	return terra.ReferenceAsSet[WorkloadProfileAttributes](acae.ref.Append("workload_profile"))
}

type azurermContainerAppEnvironmentState struct {
	DaprApplicationInsightsConnectionString string                 `json:"dapr_application_insights_connection_string"`
	DefaultDomain                           string                 `json:"default_domain"`
	DockerBridgeCidr                        string                 `json:"docker_bridge_cidr"`
	Id                                      string                 `json:"id"`
	InfrastructureResourceGroupName         string                 `json:"infrastructure_resource_group_name"`
	InfrastructureSubnetId                  string                 `json:"infrastructure_subnet_id"`
	InternalLoadBalancerEnabled             bool                   `json:"internal_load_balancer_enabled"`
	Location                                string                 `json:"location"`
	LogAnalyticsWorkspaceId                 string                 `json:"log_analytics_workspace_id"`
	Name                                    string                 `json:"name"`
	PlatformReservedCidr                    string                 `json:"platform_reserved_cidr"`
	PlatformReservedDnsIpAddress            string                 `json:"platform_reserved_dns_ip_address"`
	ResourceGroupName                       string                 `json:"resource_group_name"`
	StaticIpAddress                         string                 `json:"static_ip_address"`
	Tags                                    map[string]string      `json:"tags"`
	ZoneRedundancyEnabled                   bool                   `json:"zone_redundancy_enabled"`
	Timeouts                                *TimeoutsState         `json:"timeouts"`
	WorkloadProfile                         []WorkloadProfileState `json:"workload_profile"`
}
