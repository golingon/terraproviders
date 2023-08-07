// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceenvironment "github.com/golingon/terraproviders/azurerm/3.68.0/appserviceenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceEnvironment creates a new instance of [AppServiceEnvironment].
func NewAppServiceEnvironment(name string, args AppServiceEnvironmentArgs) *AppServiceEnvironment {
	return &AppServiceEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceEnvironment)(nil)

// AppServiceEnvironment represents the Terraform resource azurerm_app_service_environment.
type AppServiceEnvironment struct {
	Name      string
	Args      AppServiceEnvironmentArgs
	state     *appServiceEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceEnvironment].
func (ase *AppServiceEnvironment) Type() string {
	return "azurerm_app_service_environment"
}

// LocalName returns the local name for [AppServiceEnvironment].
func (ase *AppServiceEnvironment) LocalName() string {
	return ase.Name
}

// Configuration returns the configuration (args) for [AppServiceEnvironment].
func (ase *AppServiceEnvironment) Configuration() interface{} {
	return ase.Args
}

// DependOn is used for other resources to depend on [AppServiceEnvironment].
func (ase *AppServiceEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(ase)
}

// Dependencies returns the list of resources [AppServiceEnvironment] depends_on.
func (ase *AppServiceEnvironment) Dependencies() terra.Dependencies {
	return ase.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceEnvironment].
func (ase *AppServiceEnvironment) LifecycleManagement() *terra.Lifecycle {
	return ase.Lifecycle
}

// Attributes returns the attributes for [AppServiceEnvironment].
func (ase *AppServiceEnvironment) Attributes() appServiceEnvironmentAttributes {
	return appServiceEnvironmentAttributes{ref: terra.ReferenceResource(ase)}
}

// ImportState imports the given attribute values into [AppServiceEnvironment]'s state.
func (ase *AppServiceEnvironment) ImportState(av io.Reader) error {
	ase.state = &appServiceEnvironmentState{}
	if err := json.NewDecoder(av).Decode(ase.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ase.Type(), ase.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceEnvironment] has state.
func (ase *AppServiceEnvironment) State() (*appServiceEnvironmentState, bool) {
	return ase.state, ase.state != nil
}

// StateMust returns the state for [AppServiceEnvironment]. Panics if the state is nil.
func (ase *AppServiceEnvironment) StateMust() *appServiceEnvironmentState {
	if ase.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ase.Type(), ase.LocalName()))
	}
	return ase.state
}

// AppServiceEnvironmentArgs contains the configurations for azurerm_app_service_environment.
type AppServiceEnvironmentArgs struct {
	// AllowedUserIpCidrs: set of string, optional
	AllowedUserIpCidrs terra.SetValue[terra.StringValue] `hcl:"allowed_user_ip_cidrs,attr"`
	// FrontEndScaleFactor: number, optional
	FrontEndScaleFactor terra.NumberValue `hcl:"front_end_scale_factor,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternalLoadBalancingMode: string, optional
	InternalLoadBalancingMode terra.StringValue `hcl:"internal_load_balancing_mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PricingTier: string, optional
	PricingTier terra.StringValue `hcl:"pricing_tier,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ClusterSetting: min=0
	ClusterSetting []appserviceenvironment.ClusterSetting `hcl:"cluster_setting,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *appserviceenvironment.Timeouts `hcl:"timeouts,block"`
}
type appServiceEnvironmentAttributes struct {
	ref terra.Reference
}

// AllowedUserIpCidrs returns a reference to field allowed_user_ip_cidrs of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) AllowedUserIpCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ase.ref.Append("allowed_user_ip_cidrs"))
}

// FrontEndScaleFactor returns a reference to field front_end_scale_factor of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) FrontEndScaleFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(ase.ref.Append("front_end_scale_factor"))
}

// Id returns a reference to field id of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("id"))
}

// InternalIpAddress returns a reference to field internal_ip_address of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) InternalIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("internal_ip_address"))
}

// InternalLoadBalancingMode returns a reference to field internal_load_balancing_mode of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) InternalLoadBalancingMode() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("internal_load_balancing_mode"))
}

// Location returns a reference to field location of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("name"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) OutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ase.ref.Append("outbound_ip_addresses"))
}

// PricingTier returns a reference to field pricing_tier of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) PricingTier() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("pricing_tier"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("resource_group_name"))
}

// ServiceIpAddress returns a reference to field service_ip_address of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) ServiceIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("service_ip_address"))
}

// SubnetId returns a reference to field subnet_id of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_app_service_environment.
func (ase appServiceEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ase.ref.Append("tags"))
}

func (ase appServiceEnvironmentAttributes) ClusterSetting() terra.ListValue[appserviceenvironment.ClusterSettingAttributes] {
	return terra.ReferenceAsList[appserviceenvironment.ClusterSettingAttributes](ase.ref.Append("cluster_setting"))
}

func (ase appServiceEnvironmentAttributes) Timeouts() appserviceenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceenvironment.TimeoutsAttributes](ase.ref.Append("timeouts"))
}

type appServiceEnvironmentState struct {
	AllowedUserIpCidrs        []string                                    `json:"allowed_user_ip_cidrs"`
	FrontEndScaleFactor       float64                                     `json:"front_end_scale_factor"`
	Id                        string                                      `json:"id"`
	InternalIpAddress         string                                      `json:"internal_ip_address"`
	InternalLoadBalancingMode string                                      `json:"internal_load_balancing_mode"`
	Location                  string                                      `json:"location"`
	Name                      string                                      `json:"name"`
	OutboundIpAddresses       []string                                    `json:"outbound_ip_addresses"`
	PricingTier               string                                      `json:"pricing_tier"`
	ResourceGroupName         string                                      `json:"resource_group_name"`
	ServiceIpAddress          string                                      `json:"service_ip_address"`
	SubnetId                  string                                      `json:"subnet_id"`
	Tags                      map[string]string                           `json:"tags"`
	ClusterSetting            []appserviceenvironment.ClusterSettingState `json:"cluster_setting"`
	Timeouts                  *appserviceenvironment.TimeoutsState        `json:"timeouts"`
}
