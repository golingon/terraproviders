// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksnodejsapplayer "github.com/golingon/terraproviders/aws/5.0.1/opsworksnodejsapplayer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksNodejsAppLayer creates a new instance of [OpsworksNodejsAppLayer].
func NewOpsworksNodejsAppLayer(name string, args OpsworksNodejsAppLayerArgs) *OpsworksNodejsAppLayer {
	return &OpsworksNodejsAppLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksNodejsAppLayer)(nil)

// OpsworksNodejsAppLayer represents the Terraform resource aws_opsworks_nodejs_app_layer.
type OpsworksNodejsAppLayer struct {
	Name      string
	Args      OpsworksNodejsAppLayerArgs
	state     *opsworksNodejsAppLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksNodejsAppLayer].
func (onal *OpsworksNodejsAppLayer) Type() string {
	return "aws_opsworks_nodejs_app_layer"
}

// LocalName returns the local name for [OpsworksNodejsAppLayer].
func (onal *OpsworksNodejsAppLayer) LocalName() string {
	return onal.Name
}

// Configuration returns the configuration (args) for [OpsworksNodejsAppLayer].
func (onal *OpsworksNodejsAppLayer) Configuration() interface{} {
	return onal.Args
}

// DependOn is used for other resources to depend on [OpsworksNodejsAppLayer].
func (onal *OpsworksNodejsAppLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(onal)
}

// Dependencies returns the list of resources [OpsworksNodejsAppLayer] depends_on.
func (onal *OpsworksNodejsAppLayer) Dependencies() terra.Dependencies {
	return onal.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksNodejsAppLayer].
func (onal *OpsworksNodejsAppLayer) LifecycleManagement() *terra.Lifecycle {
	return onal.Lifecycle
}

// Attributes returns the attributes for [OpsworksNodejsAppLayer].
func (onal *OpsworksNodejsAppLayer) Attributes() opsworksNodejsAppLayerAttributes {
	return opsworksNodejsAppLayerAttributes{ref: terra.ReferenceResource(onal)}
}

// ImportState imports the given attribute values into [OpsworksNodejsAppLayer]'s state.
func (onal *OpsworksNodejsAppLayer) ImportState(av io.Reader) error {
	onal.state = &opsworksNodejsAppLayerState{}
	if err := json.NewDecoder(av).Decode(onal.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", onal.Type(), onal.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksNodejsAppLayer] has state.
func (onal *OpsworksNodejsAppLayer) State() (*opsworksNodejsAppLayerState, bool) {
	return onal.state, onal.state != nil
}

// StateMust returns the state for [OpsworksNodejsAppLayer]. Panics if the state is nil.
func (onal *OpsworksNodejsAppLayer) StateMust() *opsworksNodejsAppLayerState {
	if onal.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", onal.Type(), onal.LocalName()))
	}
	return onal.state
}

// OpsworksNodejsAppLayerArgs contains the configurations for aws_opsworks_nodejs_app_layer.
type OpsworksNodejsAppLayerArgs struct {
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
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstallUpdatesOnBoot: bool, optional
	InstallUpdatesOnBoot terra.BoolValue `hcl:"install_updates_on_boot,attr"`
	// InstanceShutdownTimeout: number, optional
	InstanceShutdownTimeout terra.NumberValue `hcl:"instance_shutdown_timeout,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NodejsVersion: string, optional
	NodejsVersion terra.StringValue `hcl:"nodejs_version,attr"`
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
	CloudwatchConfiguration *opsworksnodejsapplayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworksnodejsapplayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworksnodejsapplayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksNodejsAppLayerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(onal.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(onal.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(onal.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](onal.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](onal.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](onal.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](onal.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](onal.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](onal.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(onal.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(onal.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(onal.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("name"))
}

// NodejsVersion returns a reference to field nodejs_version of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) NodejsVersion() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("nodejs_version"))
}

// StackId returns a reference to field stack_id of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(onal.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](onal.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](onal.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](onal.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_nodejs_app_layer.
func (onal opsworksNodejsAppLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(onal.ref.Append("use_ebs_optimized_instances"))
}

func (onal opsworksNodejsAppLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworksnodejsapplayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworksnodejsapplayer.CloudwatchConfigurationAttributes](onal.ref.Append("cloudwatch_configuration"))
}

func (onal opsworksNodejsAppLayerAttributes) EbsVolume() terra.SetValue[opsworksnodejsapplayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworksnodejsapplayer.EbsVolumeAttributes](onal.ref.Append("ebs_volume"))
}

func (onal opsworksNodejsAppLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworksnodejsapplayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworksnodejsapplayer.LoadBasedAutoScalingAttributes](onal.ref.Append("load_based_auto_scaling"))
}

type opsworksNodejsAppLayerState struct {
	Arn                      string                                                `json:"arn"`
	AutoAssignElasticIps     bool                                                  `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                                  `json:"auto_assign_public_ips"`
	AutoHealing              bool                                                  `json:"auto_healing"`
	CustomConfigureRecipes   []string                                              `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                                              `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string                                                `json:"custom_instance_profile_arn"`
	CustomJson               string                                                `json:"custom_json"`
	CustomSecurityGroupIds   []string                                              `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string                                              `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                                              `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string                                              `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                                  `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string                                                `json:"elastic_load_balancer"`
	Id                       string                                                `json:"id"`
	InstallUpdatesOnBoot     bool                                                  `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  float64                                               `json:"instance_shutdown_timeout"`
	Name                     string                                                `json:"name"`
	NodejsVersion            string                                                `json:"nodejs_version"`
	StackId                  string                                                `json:"stack_id"`
	SystemPackages           []string                                              `json:"system_packages"`
	Tags                     map[string]string                                     `json:"tags"`
	TagsAll                  map[string]string                                     `json:"tags_all"`
	UseEbsOptimizedInstances bool                                                  `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []opsworksnodejsapplayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworksnodejsapplayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworksnodejsapplayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
