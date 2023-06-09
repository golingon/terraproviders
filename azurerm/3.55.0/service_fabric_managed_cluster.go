// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicefabricmanagedcluster "github.com/golingon/terraproviders/azurerm/3.55.0/servicefabricmanagedcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceFabricManagedCluster creates a new instance of [ServiceFabricManagedCluster].
func NewServiceFabricManagedCluster(name string, args ServiceFabricManagedClusterArgs) *ServiceFabricManagedCluster {
	return &ServiceFabricManagedCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceFabricManagedCluster)(nil)

// ServiceFabricManagedCluster represents the Terraform resource azurerm_service_fabric_managed_cluster.
type ServiceFabricManagedCluster struct {
	Name      string
	Args      ServiceFabricManagedClusterArgs
	state     *serviceFabricManagedClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceFabricManagedCluster].
func (sfmc *ServiceFabricManagedCluster) Type() string {
	return "azurerm_service_fabric_managed_cluster"
}

// LocalName returns the local name for [ServiceFabricManagedCluster].
func (sfmc *ServiceFabricManagedCluster) LocalName() string {
	return sfmc.Name
}

// Configuration returns the configuration (args) for [ServiceFabricManagedCluster].
func (sfmc *ServiceFabricManagedCluster) Configuration() interface{} {
	return sfmc.Args
}

// DependOn is used for other resources to depend on [ServiceFabricManagedCluster].
func (sfmc *ServiceFabricManagedCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(sfmc)
}

// Dependencies returns the list of resources [ServiceFabricManagedCluster] depends_on.
func (sfmc *ServiceFabricManagedCluster) Dependencies() terra.Dependencies {
	return sfmc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceFabricManagedCluster].
func (sfmc *ServiceFabricManagedCluster) LifecycleManagement() *terra.Lifecycle {
	return sfmc.Lifecycle
}

// Attributes returns the attributes for [ServiceFabricManagedCluster].
func (sfmc *ServiceFabricManagedCluster) Attributes() serviceFabricManagedClusterAttributes {
	return serviceFabricManagedClusterAttributes{ref: terra.ReferenceResource(sfmc)}
}

// ImportState imports the given attribute values into [ServiceFabricManagedCluster]'s state.
func (sfmc *ServiceFabricManagedCluster) ImportState(av io.Reader) error {
	sfmc.state = &serviceFabricManagedClusterState{}
	if err := json.NewDecoder(av).Decode(sfmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfmc.Type(), sfmc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceFabricManagedCluster] has state.
func (sfmc *ServiceFabricManagedCluster) State() (*serviceFabricManagedClusterState, bool) {
	return sfmc.state, sfmc.state != nil
}

// StateMust returns the state for [ServiceFabricManagedCluster]. Panics if the state is nil.
func (sfmc *ServiceFabricManagedCluster) StateMust() *serviceFabricManagedClusterState {
	if sfmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfmc.Type(), sfmc.LocalName()))
	}
	return sfmc.state
}

// ServiceFabricManagedClusterArgs contains the configurations for azurerm_service_fabric_managed_cluster.
type ServiceFabricManagedClusterArgs struct {
	// BackupServiceEnabled: bool, optional
	BackupServiceEnabled terra.BoolValue `hcl:"backup_service_enabled,attr"`
	// ClientConnectionPort: number, required
	ClientConnectionPort terra.NumberValue `hcl:"client_connection_port,attr" validate:"required"`
	// DnsName: string, optional
	DnsName terra.StringValue `hcl:"dns_name,attr"`
	// DnsServiceEnabled: bool, optional
	DnsServiceEnabled terra.BoolValue `hcl:"dns_service_enabled,attr"`
	// HttpGatewayPort: number, required
	HttpGatewayPort terra.NumberValue `hcl:"http_gateway_port,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UpgradeWave: string, optional
	UpgradeWave terra.StringValue `hcl:"upgrade_wave,attr"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
	// Authentication: optional
	Authentication *servicefabricmanagedcluster.Authentication `hcl:"authentication,block"`
	// CustomFabricSetting: min=0
	CustomFabricSetting []servicefabricmanagedcluster.CustomFabricSetting `hcl:"custom_fabric_setting,block" validate:"min=0"`
	// LbRule: min=1
	LbRule []servicefabricmanagedcluster.LbRule `hcl:"lb_rule,block" validate:"min=1"`
	// NodeType: min=0
	NodeType []servicefabricmanagedcluster.NodeType `hcl:"node_type,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *servicefabricmanagedcluster.Timeouts `hcl:"timeouts,block"`
}
type serviceFabricManagedClusterAttributes struct {
	ref terra.Reference
}

// BackupServiceEnabled returns a reference to field backup_service_enabled of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) BackupServiceEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sfmc.ref.Append("backup_service_enabled"))
}

// ClientConnectionPort returns a reference to field client_connection_port of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) ClientConnectionPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sfmc.ref.Append("client_connection_port"))
}

// DnsName returns a reference to field dns_name of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("dns_name"))
}

// DnsServiceEnabled returns a reference to field dns_service_enabled of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) DnsServiceEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sfmc.ref.Append("dns_service_enabled"))
}

// HttpGatewayPort returns a reference to field http_gateway_port of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) HttpGatewayPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sfmc.ref.Append("http_gateway_port"))
}

// Id returns a reference to field id of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfmc.ref.Append("tags"))
}

// UpgradeWave returns a reference to field upgrade_wave of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) UpgradeWave() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("upgrade_wave"))
}

// Username returns a reference to field username of azurerm_service_fabric_managed_cluster.
func (sfmc serviceFabricManagedClusterAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(sfmc.ref.Append("username"))
}

func (sfmc serviceFabricManagedClusterAttributes) Authentication() terra.ListValue[servicefabricmanagedcluster.AuthenticationAttributes] {
	return terra.ReferenceAsList[servicefabricmanagedcluster.AuthenticationAttributes](sfmc.ref.Append("authentication"))
}

func (sfmc serviceFabricManagedClusterAttributes) CustomFabricSetting() terra.ListValue[servicefabricmanagedcluster.CustomFabricSettingAttributes] {
	return terra.ReferenceAsList[servicefabricmanagedcluster.CustomFabricSettingAttributes](sfmc.ref.Append("custom_fabric_setting"))
}

func (sfmc serviceFabricManagedClusterAttributes) LbRule() terra.ListValue[servicefabricmanagedcluster.LbRuleAttributes] {
	return terra.ReferenceAsList[servicefabricmanagedcluster.LbRuleAttributes](sfmc.ref.Append("lb_rule"))
}

func (sfmc serviceFabricManagedClusterAttributes) NodeType() terra.ListValue[servicefabricmanagedcluster.NodeTypeAttributes] {
	return terra.ReferenceAsList[servicefabricmanagedcluster.NodeTypeAttributes](sfmc.ref.Append("node_type"))
}

func (sfmc serviceFabricManagedClusterAttributes) Timeouts() servicefabricmanagedcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicefabricmanagedcluster.TimeoutsAttributes](sfmc.ref.Append("timeouts"))
}

type serviceFabricManagedClusterState struct {
	BackupServiceEnabled bool                                                   `json:"backup_service_enabled"`
	ClientConnectionPort float64                                                `json:"client_connection_port"`
	DnsName              string                                                 `json:"dns_name"`
	DnsServiceEnabled    bool                                                   `json:"dns_service_enabled"`
	HttpGatewayPort      float64                                                `json:"http_gateway_port"`
	Id                   string                                                 `json:"id"`
	Location             string                                                 `json:"location"`
	Name                 string                                                 `json:"name"`
	Password             string                                                 `json:"password"`
	ResourceGroupName    string                                                 `json:"resource_group_name"`
	Sku                  string                                                 `json:"sku"`
	Tags                 map[string]string                                      `json:"tags"`
	UpgradeWave          string                                                 `json:"upgrade_wave"`
	Username             string                                                 `json:"username"`
	Authentication       []servicefabricmanagedcluster.AuthenticationState      `json:"authentication"`
	CustomFabricSetting  []servicefabricmanagedcluster.CustomFabricSettingState `json:"custom_fabric_setting"`
	LbRule               []servicefabricmanagedcluster.LbRuleState              `json:"lb_rule"`
	NodeType             []servicefabricmanagedcluster.NodeTypeState            `json:"node_type"`
	Timeouts             *servicefabricmanagedcluster.TimeoutsState             `json:"timeouts"`
}
