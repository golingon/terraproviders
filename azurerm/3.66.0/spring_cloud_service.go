// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudservice "github.com/golingon/terraproviders/azurerm/3.66.0/springcloudservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudService creates a new instance of [SpringCloudService].
func NewSpringCloudService(name string, args SpringCloudServiceArgs) *SpringCloudService {
	return &SpringCloudService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudService)(nil)

// SpringCloudService represents the Terraform resource azurerm_spring_cloud_service.
type SpringCloudService struct {
	Name      string
	Args      SpringCloudServiceArgs
	state     *springCloudServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudService].
func (scs *SpringCloudService) Type() string {
	return "azurerm_spring_cloud_service"
}

// LocalName returns the local name for [SpringCloudService].
func (scs *SpringCloudService) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [SpringCloudService].
func (scs *SpringCloudService) Configuration() interface{} {
	return scs.Args
}

// DependOn is used for other resources to depend on [SpringCloudService].
func (scs *SpringCloudService) DependOn() terra.Reference {
	return terra.ReferenceResource(scs)
}

// Dependencies returns the list of resources [SpringCloudService] depends_on.
func (scs *SpringCloudService) Dependencies() terra.Dependencies {
	return scs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudService].
func (scs *SpringCloudService) LifecycleManagement() *terra.Lifecycle {
	return scs.Lifecycle
}

// Attributes returns the attributes for [SpringCloudService].
func (scs *SpringCloudService) Attributes() springCloudServiceAttributes {
	return springCloudServiceAttributes{ref: terra.ReferenceResource(scs)}
}

// ImportState imports the given attribute values into [SpringCloudService]'s state.
func (scs *SpringCloudService) ImportState(av io.Reader) error {
	scs.state = &springCloudServiceState{}
	if err := json.NewDecoder(av).Decode(scs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scs.Type(), scs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudService] has state.
func (scs *SpringCloudService) State() (*springCloudServiceState, bool) {
	return scs.state, scs.state != nil
}

// StateMust returns the state for [SpringCloudService]. Panics if the state is nil.
func (scs *SpringCloudService) StateMust() *springCloudServiceState {
	if scs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scs.Type(), scs.LocalName()))
	}
	return scs.state
}

// SpringCloudServiceArgs contains the configurations for azurerm_spring_cloud_service.
type SpringCloudServiceArgs struct {
	// BuildAgentPoolSize: string, optional
	BuildAgentPoolSize terra.StringValue `hcl:"build_agent_pool_size,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LogStreamPublicEndpointEnabled: bool, optional
	LogStreamPublicEndpointEnabled terra.BoolValue `hcl:"log_stream_public_endpoint_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceRegistryEnabled: bool, optional
	ServiceRegistryEnabled terra.BoolValue `hcl:"service_registry_enabled,attr"`
	// SkuName: string, optional
	SkuName terra.StringValue `hcl:"sku_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
	// RequiredNetworkTrafficRules: min=0
	RequiredNetworkTrafficRules []springcloudservice.RequiredNetworkTrafficRules `hcl:"required_network_traffic_rules,block" validate:"min=0"`
	// ConfigServerGitSetting: optional
	ConfigServerGitSetting *springcloudservice.ConfigServerGitSetting `hcl:"config_server_git_setting,block"`
	// ContainerRegistry: min=0
	ContainerRegistry []springcloudservice.ContainerRegistry `hcl:"container_registry,block" validate:"min=0"`
	// DefaultBuildService: optional
	DefaultBuildService *springcloudservice.DefaultBuildService `hcl:"default_build_service,block"`
	// Marketplace: optional
	Marketplace *springcloudservice.Marketplace `hcl:"marketplace,block"`
	// Network: optional
	Network *springcloudservice.Network `hcl:"network,block"`
	// Timeouts: optional
	Timeouts *springcloudservice.Timeouts `hcl:"timeouts,block"`
	// Trace: optional
	Trace *springcloudservice.Trace `hcl:"trace,block"`
}
type springCloudServiceAttributes struct {
	ref terra.Reference
}

// BuildAgentPoolSize returns a reference to field build_agent_pool_size of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) BuildAgentPoolSize() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("build_agent_pool_size"))
}

// Id returns a reference to field id of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("location"))
}

// LogStreamPublicEndpointEnabled returns a reference to field log_stream_public_endpoint_enabled of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) LogStreamPublicEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(scs.ref.Append("log_stream_public_endpoint_enabled"))
}

// Name returns a reference to field name of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("name"))
}

// OutboundPublicIpAddresses returns a reference to field outbound_public_ip_addresses of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) OutboundPublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](scs.ref.Append("outbound_public_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("resource_group_name"))
}

// ServiceRegistryEnabled returns a reference to field service_registry_enabled of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) ServiceRegistryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(scs.ref.Append("service_registry_enabled"))
}

// ServiceRegistryId returns a reference to field service_registry_id of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) ServiceRegistryId() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("service_registry_id"))
}

// SkuName returns a reference to field sku_name of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_spring_cloud_service.
func (scs springCloudServiceAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(scs.ref.Append("zone_redundant"))
}

func (scs springCloudServiceAttributes) RequiredNetworkTrafficRules() terra.ListValue[springcloudservice.RequiredNetworkTrafficRulesAttributes] {
	return terra.ReferenceAsList[springcloudservice.RequiredNetworkTrafficRulesAttributes](scs.ref.Append("required_network_traffic_rules"))
}

func (scs springCloudServiceAttributes) ConfigServerGitSetting() terra.ListValue[springcloudservice.ConfigServerGitSettingAttributes] {
	return terra.ReferenceAsList[springcloudservice.ConfigServerGitSettingAttributes](scs.ref.Append("config_server_git_setting"))
}

func (scs springCloudServiceAttributes) ContainerRegistry() terra.ListValue[springcloudservice.ContainerRegistryAttributes] {
	return terra.ReferenceAsList[springcloudservice.ContainerRegistryAttributes](scs.ref.Append("container_registry"))
}

func (scs springCloudServiceAttributes) DefaultBuildService() terra.ListValue[springcloudservice.DefaultBuildServiceAttributes] {
	return terra.ReferenceAsList[springcloudservice.DefaultBuildServiceAttributes](scs.ref.Append("default_build_service"))
}

func (scs springCloudServiceAttributes) Marketplace() terra.ListValue[springcloudservice.MarketplaceAttributes] {
	return terra.ReferenceAsList[springcloudservice.MarketplaceAttributes](scs.ref.Append("marketplace"))
}

func (scs springCloudServiceAttributes) Network() terra.ListValue[springcloudservice.NetworkAttributes] {
	return terra.ReferenceAsList[springcloudservice.NetworkAttributes](scs.ref.Append("network"))
}

func (scs springCloudServiceAttributes) Timeouts() springcloudservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudservice.TimeoutsAttributes](scs.ref.Append("timeouts"))
}

func (scs springCloudServiceAttributes) Trace() terra.ListValue[springcloudservice.TraceAttributes] {
	return terra.ReferenceAsList[springcloudservice.TraceAttributes](scs.ref.Append("trace"))
}

type springCloudServiceState struct {
	BuildAgentPoolSize             string                                                `json:"build_agent_pool_size"`
	Id                             string                                                `json:"id"`
	Location                       string                                                `json:"location"`
	LogStreamPublicEndpointEnabled bool                                                  `json:"log_stream_public_endpoint_enabled"`
	Name                           string                                                `json:"name"`
	OutboundPublicIpAddresses      []string                                              `json:"outbound_public_ip_addresses"`
	ResourceGroupName              string                                                `json:"resource_group_name"`
	ServiceRegistryEnabled         bool                                                  `json:"service_registry_enabled"`
	ServiceRegistryId              string                                                `json:"service_registry_id"`
	SkuName                        string                                                `json:"sku_name"`
	Tags                           map[string]string                                     `json:"tags"`
	ZoneRedundant                  bool                                                  `json:"zone_redundant"`
	RequiredNetworkTrafficRules    []springcloudservice.RequiredNetworkTrafficRulesState `json:"required_network_traffic_rules"`
	ConfigServerGitSetting         []springcloudservice.ConfigServerGitSettingState      `json:"config_server_git_setting"`
	ContainerRegistry              []springcloudservice.ContainerRegistryState           `json:"container_registry"`
	DefaultBuildService            []springcloudservice.DefaultBuildServiceState         `json:"default_build_service"`
	Marketplace                    []springcloudservice.MarketplaceState                 `json:"marketplace"`
	Network                        []springcloudservice.NetworkState                     `json:"network"`
	Timeouts                       *springcloudservice.TimeoutsState                     `json:"timeouts"`
	Trace                          []springcloudservice.TraceState                       `json:"trace"`
}
