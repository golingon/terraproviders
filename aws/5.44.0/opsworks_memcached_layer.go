// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	opsworksmemcachedlayer "github.com/golingon/terraproviders/aws/5.44.0/opsworksmemcachedlayer"
	"io"
)

// NewOpsworksMemcachedLayer creates a new instance of [OpsworksMemcachedLayer].
func NewOpsworksMemcachedLayer(name string, args OpsworksMemcachedLayerArgs) *OpsworksMemcachedLayer {
	return &OpsworksMemcachedLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksMemcachedLayer)(nil)

// OpsworksMemcachedLayer represents the Terraform resource aws_opsworks_memcached_layer.
type OpsworksMemcachedLayer struct {
	Name      string
	Args      OpsworksMemcachedLayerArgs
	state     *opsworksMemcachedLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksMemcachedLayer].
func (oml *OpsworksMemcachedLayer) Type() string {
	return "aws_opsworks_memcached_layer"
}

// LocalName returns the local name for [OpsworksMemcachedLayer].
func (oml *OpsworksMemcachedLayer) LocalName() string {
	return oml.Name
}

// Configuration returns the configuration (args) for [OpsworksMemcachedLayer].
func (oml *OpsworksMemcachedLayer) Configuration() interface{} {
	return oml.Args
}

// DependOn is used for other resources to depend on [OpsworksMemcachedLayer].
func (oml *OpsworksMemcachedLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(oml)
}

// Dependencies returns the list of resources [OpsworksMemcachedLayer] depends_on.
func (oml *OpsworksMemcachedLayer) Dependencies() terra.Dependencies {
	return oml.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksMemcachedLayer].
func (oml *OpsworksMemcachedLayer) LifecycleManagement() *terra.Lifecycle {
	return oml.Lifecycle
}

// Attributes returns the attributes for [OpsworksMemcachedLayer].
func (oml *OpsworksMemcachedLayer) Attributes() opsworksMemcachedLayerAttributes {
	return opsworksMemcachedLayerAttributes{ref: terra.ReferenceResource(oml)}
}

// ImportState imports the given attribute values into [OpsworksMemcachedLayer]'s state.
func (oml *OpsworksMemcachedLayer) ImportState(av io.Reader) error {
	oml.state = &opsworksMemcachedLayerState{}
	if err := json.NewDecoder(av).Decode(oml.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oml.Type(), oml.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksMemcachedLayer] has state.
func (oml *OpsworksMemcachedLayer) State() (*opsworksMemcachedLayerState, bool) {
	return oml.state, oml.state != nil
}

// StateMust returns the state for [OpsworksMemcachedLayer]. Panics if the state is nil.
func (oml *OpsworksMemcachedLayer) StateMust() *opsworksMemcachedLayerState {
	if oml.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oml.Type(), oml.LocalName()))
	}
	return oml.state
}

// OpsworksMemcachedLayerArgs contains the configurations for aws_opsworks_memcached_layer.
type OpsworksMemcachedLayerArgs struct {
	// AllocatedMemory: number, optional
	AllocatedMemory terra.NumberValue `hcl:"allocated_memory,attr"`
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
	CloudwatchConfiguration *opsworksmemcachedlayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworksmemcachedlayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworksmemcachedlayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksMemcachedLayerAttributes struct {
	ref terra.Reference
}

// AllocatedMemory returns a reference to field allocated_memory of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) AllocatedMemory() terra.NumberValue {
	return terra.ReferenceAsNumber(oml.ref.Append("allocated_memory"))
}

// Arn returns a reference to field arn of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(oml.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(oml.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(oml.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oml.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oml.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oml.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oml.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oml.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oml.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(oml.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(oml.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(oml.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(oml.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oml.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oml.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oml.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_memcached_layer.
func (oml opsworksMemcachedLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(oml.ref.Append("use_ebs_optimized_instances"))
}

func (oml opsworksMemcachedLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworksmemcachedlayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworksmemcachedlayer.CloudwatchConfigurationAttributes](oml.ref.Append("cloudwatch_configuration"))
}

func (oml opsworksMemcachedLayerAttributes) EbsVolume() terra.SetValue[opsworksmemcachedlayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworksmemcachedlayer.EbsVolumeAttributes](oml.ref.Append("ebs_volume"))
}

func (oml opsworksMemcachedLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworksmemcachedlayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworksmemcachedlayer.LoadBasedAutoScalingAttributes](oml.ref.Append("load_based_auto_scaling"))
}

type opsworksMemcachedLayerState struct {
	AllocatedMemory          float64                                               `json:"allocated_memory"`
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
	StackId                  string                                                `json:"stack_id"`
	SystemPackages           []string                                              `json:"system_packages"`
	Tags                     map[string]string                                     `json:"tags"`
	TagsAll                  map[string]string                                     `json:"tags_all"`
	UseEbsOptimizedInstances bool                                                  `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []opsworksmemcachedlayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworksmemcachedlayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworksmemcachedlayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
