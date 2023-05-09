// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksecsclusterlayer "github.com/golingon/terraproviders/aws/4.66.1/opsworksecsclusterlayer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksEcsClusterLayer creates a new instance of [OpsworksEcsClusterLayer].
func NewOpsworksEcsClusterLayer(name string, args OpsworksEcsClusterLayerArgs) *OpsworksEcsClusterLayer {
	return &OpsworksEcsClusterLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksEcsClusterLayer)(nil)

// OpsworksEcsClusterLayer represents the Terraform resource aws_opsworks_ecs_cluster_layer.
type OpsworksEcsClusterLayer struct {
	Name      string
	Args      OpsworksEcsClusterLayerArgs
	state     *opsworksEcsClusterLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksEcsClusterLayer].
func (oecl *OpsworksEcsClusterLayer) Type() string {
	return "aws_opsworks_ecs_cluster_layer"
}

// LocalName returns the local name for [OpsworksEcsClusterLayer].
func (oecl *OpsworksEcsClusterLayer) LocalName() string {
	return oecl.Name
}

// Configuration returns the configuration (args) for [OpsworksEcsClusterLayer].
func (oecl *OpsworksEcsClusterLayer) Configuration() interface{} {
	return oecl.Args
}

// DependOn is used for other resources to depend on [OpsworksEcsClusterLayer].
func (oecl *OpsworksEcsClusterLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(oecl)
}

// Dependencies returns the list of resources [OpsworksEcsClusterLayer] depends_on.
func (oecl *OpsworksEcsClusterLayer) Dependencies() terra.Dependencies {
	return oecl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksEcsClusterLayer].
func (oecl *OpsworksEcsClusterLayer) LifecycleManagement() *terra.Lifecycle {
	return oecl.Lifecycle
}

// Attributes returns the attributes for [OpsworksEcsClusterLayer].
func (oecl *OpsworksEcsClusterLayer) Attributes() opsworksEcsClusterLayerAttributes {
	return opsworksEcsClusterLayerAttributes{ref: terra.ReferenceResource(oecl)}
}

// ImportState imports the given attribute values into [OpsworksEcsClusterLayer]'s state.
func (oecl *OpsworksEcsClusterLayer) ImportState(av io.Reader) error {
	oecl.state = &opsworksEcsClusterLayerState{}
	if err := json.NewDecoder(av).Decode(oecl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oecl.Type(), oecl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksEcsClusterLayer] has state.
func (oecl *OpsworksEcsClusterLayer) State() (*opsworksEcsClusterLayerState, bool) {
	return oecl.state, oecl.state != nil
}

// StateMust returns the state for [OpsworksEcsClusterLayer]. Panics if the state is nil.
func (oecl *OpsworksEcsClusterLayer) StateMust() *opsworksEcsClusterLayerState {
	if oecl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oecl.Type(), oecl.LocalName()))
	}
	return oecl.state
}

// OpsworksEcsClusterLayerArgs contains the configurations for aws_opsworks_ecs_cluster_layer.
type OpsworksEcsClusterLayerArgs struct {
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
	// EcsClusterArn: string, required
	EcsClusterArn terra.StringValue `hcl:"ecs_cluster_arn,attr" validate:"required"`
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
	CloudwatchConfiguration *opsworksecsclusterlayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworksecsclusterlayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworksecsclusterlayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksEcsClusterLayerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(oecl.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(oecl.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(oecl.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oecl.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oecl.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oecl.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oecl.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oecl.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oecl.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(oecl.ref.Append("drain_elb_on_shutdown"))
}

// EcsClusterArn returns a reference to field ecs_cluster_arn of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) EcsClusterArn() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("ecs_cluster_arn"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(oecl.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(oecl.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(oecl.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oecl.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oecl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oecl.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_ecs_cluster_layer.
func (oecl opsworksEcsClusterLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(oecl.ref.Append("use_ebs_optimized_instances"))
}

func (oecl opsworksEcsClusterLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworksecsclusterlayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworksecsclusterlayer.CloudwatchConfigurationAttributes](oecl.ref.Append("cloudwatch_configuration"))
}

func (oecl opsworksEcsClusterLayerAttributes) EbsVolume() terra.SetValue[opsworksecsclusterlayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworksecsclusterlayer.EbsVolumeAttributes](oecl.ref.Append("ebs_volume"))
}

func (oecl opsworksEcsClusterLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworksecsclusterlayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworksecsclusterlayer.LoadBasedAutoScalingAttributes](oecl.ref.Append("load_based_auto_scaling"))
}

type opsworksEcsClusterLayerState struct {
	Arn                      string                                                 `json:"arn"`
	AutoAssignElasticIps     bool                                                   `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                                   `json:"auto_assign_public_ips"`
	AutoHealing              bool                                                   `json:"auto_healing"`
	CustomConfigureRecipes   []string                                               `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                                               `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string                                                 `json:"custom_instance_profile_arn"`
	CustomJson               string                                                 `json:"custom_json"`
	CustomSecurityGroupIds   []string                                               `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string                                               `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                                               `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string                                               `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                                   `json:"drain_elb_on_shutdown"`
	EcsClusterArn            string                                                 `json:"ecs_cluster_arn"`
	ElasticLoadBalancer      string                                                 `json:"elastic_load_balancer"`
	Id                       string                                                 `json:"id"`
	InstallUpdatesOnBoot     bool                                                   `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  float64                                                `json:"instance_shutdown_timeout"`
	Name                     string                                                 `json:"name"`
	StackId                  string                                                 `json:"stack_id"`
	SystemPackages           []string                                               `json:"system_packages"`
	Tags                     map[string]string                                      `json:"tags"`
	TagsAll                  map[string]string                                      `json:"tags_all"`
	UseEbsOptimizedInstances bool                                                   `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []opsworksecsclusterlayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworksecsclusterlayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworksecsclusterlayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
