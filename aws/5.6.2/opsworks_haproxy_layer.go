// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworkshaproxylayer "github.com/golingon/terraproviders/aws/5.6.2/opsworkshaproxylayer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksHaproxyLayer creates a new instance of [OpsworksHaproxyLayer].
func NewOpsworksHaproxyLayer(name string, args OpsworksHaproxyLayerArgs) *OpsworksHaproxyLayer {
	return &OpsworksHaproxyLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksHaproxyLayer)(nil)

// OpsworksHaproxyLayer represents the Terraform resource aws_opsworks_haproxy_layer.
type OpsworksHaproxyLayer struct {
	Name      string
	Args      OpsworksHaproxyLayerArgs
	state     *opsworksHaproxyLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksHaproxyLayer].
func (ohl *OpsworksHaproxyLayer) Type() string {
	return "aws_opsworks_haproxy_layer"
}

// LocalName returns the local name for [OpsworksHaproxyLayer].
func (ohl *OpsworksHaproxyLayer) LocalName() string {
	return ohl.Name
}

// Configuration returns the configuration (args) for [OpsworksHaproxyLayer].
func (ohl *OpsworksHaproxyLayer) Configuration() interface{} {
	return ohl.Args
}

// DependOn is used for other resources to depend on [OpsworksHaproxyLayer].
func (ohl *OpsworksHaproxyLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(ohl)
}

// Dependencies returns the list of resources [OpsworksHaproxyLayer] depends_on.
func (ohl *OpsworksHaproxyLayer) Dependencies() terra.Dependencies {
	return ohl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksHaproxyLayer].
func (ohl *OpsworksHaproxyLayer) LifecycleManagement() *terra.Lifecycle {
	return ohl.Lifecycle
}

// Attributes returns the attributes for [OpsworksHaproxyLayer].
func (ohl *OpsworksHaproxyLayer) Attributes() opsworksHaproxyLayerAttributes {
	return opsworksHaproxyLayerAttributes{ref: terra.ReferenceResource(ohl)}
}

// ImportState imports the given attribute values into [OpsworksHaproxyLayer]'s state.
func (ohl *OpsworksHaproxyLayer) ImportState(av io.Reader) error {
	ohl.state = &opsworksHaproxyLayerState{}
	if err := json.NewDecoder(av).Decode(ohl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ohl.Type(), ohl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksHaproxyLayer] has state.
func (ohl *OpsworksHaproxyLayer) State() (*opsworksHaproxyLayerState, bool) {
	return ohl.state, ohl.state != nil
}

// StateMust returns the state for [OpsworksHaproxyLayer]. Panics if the state is nil.
func (ohl *OpsworksHaproxyLayer) StateMust() *opsworksHaproxyLayerState {
	if ohl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ohl.Type(), ohl.LocalName()))
	}
	return ohl.state
}

// OpsworksHaproxyLayerArgs contains the configurations for aws_opsworks_haproxy_layer.
type OpsworksHaproxyLayerArgs struct {
	// AutoAssignElasticIps: bool, optional
	AutoAssignElasticIps terra.BoolValue `hcl:"auto_assign_elastic_ips,attr"`
	// AutoAssignPublicIps: bool, optional
	AutoAssignPublicIps terra.BoolValue `hcl:"auto_assign_public_ips,attr"`
	// AutoHealing: bool, optional
	AutoHealing terra.BoolValue `hcl:"auto_healing,attr"`
	// CustomConfigureRecipes: list of string, optional
	CustomConfigureRecipes terra.ListValue[terra.StringValue] `hcl:"custom_configure_recipes,attr"`
	// CustomDeployRecipes: list of string, optional
	CustomDeployRecipes terra.ListValue[terra.StringValue] `hcl:"custom_deploy_recipes,attr"`
	// CustomInstanceProfileArn: string, optional
	CustomInstanceProfileArn terra.StringValue `hcl:"custom_instance_profile_arn,attr"`
	// CustomJson: string, optional
	CustomJson terra.StringValue `hcl:"custom_json,attr"`
	// CustomSecurityGroupIds: set of string, optional
	CustomSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"custom_security_group_ids,attr"`
	// CustomSetupRecipes: list of string, optional
	CustomSetupRecipes terra.ListValue[terra.StringValue] `hcl:"custom_setup_recipes,attr"`
	// CustomShutdownRecipes: list of string, optional
	CustomShutdownRecipes terra.ListValue[terra.StringValue] `hcl:"custom_shutdown_recipes,attr"`
	// CustomUndeployRecipes: list of string, optional
	CustomUndeployRecipes terra.ListValue[terra.StringValue] `hcl:"custom_undeploy_recipes,attr"`
	// DrainElbOnShutdown: bool, optional
	DrainElbOnShutdown terra.BoolValue `hcl:"drain_elb_on_shutdown,attr"`
	// ElasticLoadBalancer: string, optional
	ElasticLoadBalancer terra.StringValue `hcl:"elastic_load_balancer,attr"`
	// HealthcheckMethod: string, optional
	HealthcheckMethod terra.StringValue `hcl:"healthcheck_method,attr"`
	// HealthcheckUrl: string, optional
	HealthcheckUrl terra.StringValue `hcl:"healthcheck_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstallUpdatesOnBoot: bool, optional
	InstallUpdatesOnBoot terra.BoolValue `hcl:"install_updates_on_boot,attr"`
	// InstanceShutdownTimeout: number, optional
	InstanceShutdownTimeout terra.NumberValue `hcl:"instance_shutdown_timeout,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// StackId: string, required
	StackId terra.StringValue `hcl:"stack_id,attr" validate:"required"`
	// StatsEnabled: bool, optional
	StatsEnabled terra.BoolValue `hcl:"stats_enabled,attr"`
	// StatsPassword: string, required
	StatsPassword terra.StringValue `hcl:"stats_password,attr" validate:"required"`
	// StatsUrl: string, optional
	StatsUrl terra.StringValue `hcl:"stats_url,attr"`
	// StatsUser: string, optional
	StatsUser terra.StringValue `hcl:"stats_user,attr"`
	// SystemPackages: set of string, optional
	SystemPackages terra.SetValue[terra.StringValue] `hcl:"system_packages,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UseEbsOptimizedInstances: bool, optional
	UseEbsOptimizedInstances terra.BoolValue `hcl:"use_ebs_optimized_instances,attr"`
	// CloudwatchConfiguration: optional
	CloudwatchConfiguration *opsworkshaproxylayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworkshaproxylayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworkshaproxylayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksHaproxyLayerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ohl.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ohl.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ohl.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ohl.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ohl.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ohl.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("elastic_load_balancer"))
}

// HealthcheckMethod returns a reference to field healthcheck_method of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) HealthcheckMethod() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("healthcheck_method"))
}

// HealthcheckUrl returns a reference to field healthcheck_url of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) HealthcheckUrl() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("healthcheck_url"))
}

// Id returns a reference to field id of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ohl.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("stack_id"))
}

// StatsEnabled returns a reference to field stats_enabled of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) StatsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("stats_enabled"))
}

// StatsPassword returns a reference to field stats_password of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) StatsPassword() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("stats_password"))
}

// StatsUrl returns a reference to field stats_url of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) StatsUrl() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("stats_url"))
}

// StatsUser returns a reference to field stats_user of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) StatsUser() terra.StringValue {
	return terra.ReferenceAsString(ohl.ref.Append("stats_user"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ohl.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ohl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ohl.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_haproxy_layer.
func (ohl opsworksHaproxyLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(ohl.ref.Append("use_ebs_optimized_instances"))
}

func (ohl opsworksHaproxyLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworkshaproxylayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworkshaproxylayer.CloudwatchConfigurationAttributes](ohl.ref.Append("cloudwatch_configuration"))
}

func (ohl opsworksHaproxyLayerAttributes) EbsVolume() terra.SetValue[opsworkshaproxylayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworkshaproxylayer.EbsVolumeAttributes](ohl.ref.Append("ebs_volume"))
}

func (ohl opsworksHaproxyLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworkshaproxylayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworkshaproxylayer.LoadBasedAutoScalingAttributes](ohl.ref.Append("load_based_auto_scaling"))
}

type opsworksHaproxyLayerState struct {
	Arn                      string                                              `json:"arn"`
	AutoAssignElasticIps     bool                                                `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                                `json:"auto_assign_public_ips"`
	AutoHealing              bool                                                `json:"auto_healing"`
	CustomConfigureRecipes   []string                                            `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                                            `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string                                              `json:"custom_instance_profile_arn"`
	CustomJson               string                                              `json:"custom_json"`
	CustomSecurityGroupIds   []string                                            `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string                                            `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                                            `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string                                            `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                                `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string                                              `json:"elastic_load_balancer"`
	HealthcheckMethod        string                                              `json:"healthcheck_method"`
	HealthcheckUrl           string                                              `json:"healthcheck_url"`
	Id                       string                                              `json:"id"`
	InstallUpdatesOnBoot     bool                                                `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  float64                                             `json:"instance_shutdown_timeout"`
	Name                     string                                              `json:"name"`
	StackId                  string                                              `json:"stack_id"`
	StatsEnabled             bool                                                `json:"stats_enabled"`
	StatsPassword            string                                              `json:"stats_password"`
	StatsUrl                 string                                              `json:"stats_url"`
	StatsUser                string                                              `json:"stats_user"`
	SystemPackages           []string                                            `json:"system_packages"`
	Tags                     map[string]string                                   `json:"tags"`
	TagsAll                  map[string]string                                   `json:"tags_all"`
	UseEbsOptimizedInstances bool                                                `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []opsworkshaproxylayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworkshaproxylayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworkshaproxylayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
