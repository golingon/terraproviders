// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksrailsapplayer "github.com/golingon/terraproviders/aws/5.6.2/opsworksrailsapplayer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksRailsAppLayer creates a new instance of [OpsworksRailsAppLayer].
func NewOpsworksRailsAppLayer(name string, args OpsworksRailsAppLayerArgs) *OpsworksRailsAppLayer {
	return &OpsworksRailsAppLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksRailsAppLayer)(nil)

// OpsworksRailsAppLayer represents the Terraform resource aws_opsworks_rails_app_layer.
type OpsworksRailsAppLayer struct {
	Name      string
	Args      OpsworksRailsAppLayerArgs
	state     *opsworksRailsAppLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksRailsAppLayer].
func (oral *OpsworksRailsAppLayer) Type() string {
	return "aws_opsworks_rails_app_layer"
}

// LocalName returns the local name for [OpsworksRailsAppLayer].
func (oral *OpsworksRailsAppLayer) LocalName() string {
	return oral.Name
}

// Configuration returns the configuration (args) for [OpsworksRailsAppLayer].
func (oral *OpsworksRailsAppLayer) Configuration() interface{} {
	return oral.Args
}

// DependOn is used for other resources to depend on [OpsworksRailsAppLayer].
func (oral *OpsworksRailsAppLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(oral)
}

// Dependencies returns the list of resources [OpsworksRailsAppLayer] depends_on.
func (oral *OpsworksRailsAppLayer) Dependencies() terra.Dependencies {
	return oral.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksRailsAppLayer].
func (oral *OpsworksRailsAppLayer) LifecycleManagement() *terra.Lifecycle {
	return oral.Lifecycle
}

// Attributes returns the attributes for [OpsworksRailsAppLayer].
func (oral *OpsworksRailsAppLayer) Attributes() opsworksRailsAppLayerAttributes {
	return opsworksRailsAppLayerAttributes{ref: terra.ReferenceResource(oral)}
}

// ImportState imports the given attribute values into [OpsworksRailsAppLayer]'s state.
func (oral *OpsworksRailsAppLayer) ImportState(av io.Reader) error {
	oral.state = &opsworksRailsAppLayerState{}
	if err := json.NewDecoder(av).Decode(oral.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oral.Type(), oral.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksRailsAppLayer] has state.
func (oral *OpsworksRailsAppLayer) State() (*opsworksRailsAppLayerState, bool) {
	return oral.state, oral.state != nil
}

// StateMust returns the state for [OpsworksRailsAppLayer]. Panics if the state is nil.
func (oral *OpsworksRailsAppLayer) StateMust() *opsworksRailsAppLayerState {
	if oral.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oral.Type(), oral.LocalName()))
	}
	return oral.state
}

// OpsworksRailsAppLayerArgs contains the configurations for aws_opsworks_rails_app_layer.
type OpsworksRailsAppLayerArgs struct {
	// AppServer: string, optional
	AppServer terra.StringValue `hcl:"app_server,attr"`
	// AutoAssignElasticIps: bool, optional
	AutoAssignElasticIps terra.BoolValue `hcl:"auto_assign_elastic_ips,attr"`
	// AutoAssignPublicIps: bool, optional
	AutoAssignPublicIps terra.BoolValue `hcl:"auto_assign_public_ips,attr"`
	// AutoHealing: bool, optional
	AutoHealing terra.BoolValue `hcl:"auto_healing,attr"`
	// BundlerVersion: string, optional
	BundlerVersion terra.StringValue `hcl:"bundler_version,attr"`
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
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstallUpdatesOnBoot: bool, optional
	InstallUpdatesOnBoot terra.BoolValue `hcl:"install_updates_on_boot,attr"`
	// InstanceShutdownTimeout: number, optional
	InstanceShutdownTimeout terra.NumberValue `hcl:"instance_shutdown_timeout,attr"`
	// ManageBundler: bool, optional
	ManageBundler terra.BoolValue `hcl:"manage_bundler,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PassengerVersion: string, optional
	PassengerVersion terra.StringValue `hcl:"passenger_version,attr"`
	// RubyVersion: string, optional
	RubyVersion terra.StringValue `hcl:"ruby_version,attr"`
	// RubygemsVersion: string, optional
	RubygemsVersion terra.StringValue `hcl:"rubygems_version,attr"`
	// StackId: string, required
	StackId terra.StringValue `hcl:"stack_id,attr" validate:"required"`
	// SystemPackages: set of string, optional
	SystemPackages terra.SetValue[terra.StringValue] `hcl:"system_packages,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UseEbsOptimizedInstances: bool, optional
	UseEbsOptimizedInstances terra.BoolValue `hcl:"use_ebs_optimized_instances,attr"`
	// CloudwatchConfiguration: optional
	CloudwatchConfiguration *opsworksrailsapplayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworksrailsapplayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworksrailsapplayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksRailsAppLayerAttributes struct {
	ref terra.Reference
}

// AppServer returns a reference to field app_server of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) AppServer() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("app_server"))
}

// Arn returns a reference to field arn of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("auto_healing"))
}

// BundlerVersion returns a reference to field bundler_version of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) BundlerVersion() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("bundler_version"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oral.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oral.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oral.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oral.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oral.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oral.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(oral.ref.Append("instance_shutdown_timeout"))
}

// ManageBundler returns a reference to field manage_bundler of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) ManageBundler() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("manage_bundler"))
}

// Name returns a reference to field name of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("name"))
}

// PassengerVersion returns a reference to field passenger_version of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) PassengerVersion() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("passenger_version"))
}

// RubyVersion returns a reference to field ruby_version of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) RubyVersion() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("ruby_version"))
}

// RubygemsVersion returns a reference to field rubygems_version of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) RubygemsVersion() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("rubygems_version"))
}

// StackId returns a reference to field stack_id of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(oral.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oral.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oral.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oral.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_rails_app_layer.
func (oral opsworksRailsAppLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(oral.ref.Append("use_ebs_optimized_instances"))
}

func (oral opsworksRailsAppLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworksrailsapplayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworksrailsapplayer.CloudwatchConfigurationAttributes](oral.ref.Append("cloudwatch_configuration"))
}

func (oral opsworksRailsAppLayerAttributes) EbsVolume() terra.SetValue[opsworksrailsapplayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworksrailsapplayer.EbsVolumeAttributes](oral.ref.Append("ebs_volume"))
}

func (oral opsworksRailsAppLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworksrailsapplayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworksrailsapplayer.LoadBasedAutoScalingAttributes](oral.ref.Append("load_based_auto_scaling"))
}

type opsworksRailsAppLayerState struct {
	AppServer                string                                               `json:"app_server"`
	Arn                      string                                               `json:"arn"`
	AutoAssignElasticIps     bool                                                 `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                                 `json:"auto_assign_public_ips"`
	AutoHealing              bool                                                 `json:"auto_healing"`
	BundlerVersion           string                                               `json:"bundler_version"`
	CustomConfigureRecipes   []string                                             `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                                             `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string                                               `json:"custom_instance_profile_arn"`
	CustomJson               string                                               `json:"custom_json"`
	CustomSecurityGroupIds   []string                                             `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string                                             `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                                             `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string                                             `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                                 `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string                                               `json:"elastic_load_balancer"`
	Id                       string                                               `json:"id"`
	InstallUpdatesOnBoot     bool                                                 `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  float64                                              `json:"instance_shutdown_timeout"`
	ManageBundler            bool                                                 `json:"manage_bundler"`
	Name                     string                                               `json:"name"`
	PassengerVersion         string                                               `json:"passenger_version"`
	RubyVersion              string                                               `json:"ruby_version"`
	RubygemsVersion          string                                               `json:"rubygems_version"`
	StackId                  string                                               `json:"stack_id"`
	SystemPackages           []string                                             `json:"system_packages"`
	Tags                     map[string]string                                    `json:"tags"`
	TagsAll                  map[string]string                                    `json:"tags_all"`
	UseEbsOptimizedInstances bool                                                 `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []opsworksrailsapplayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworksrailsapplayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworksrailsapplayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
