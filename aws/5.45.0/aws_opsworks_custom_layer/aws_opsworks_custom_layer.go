// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_opsworks_custom_layer

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

// Resource represents the Terraform resource aws_opsworks_custom_layer.
type Resource struct {
	Name      string
	Args      Args
	state     *awsOpsworksCustomLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aocl *Resource) Type() string {
	return "aws_opsworks_custom_layer"
}

// LocalName returns the local name for [Resource].
func (aocl *Resource) LocalName() string {
	return aocl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aocl *Resource) Configuration() interface{} {
	return aocl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aocl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aocl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aocl *Resource) Dependencies() terra.Dependencies {
	return aocl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aocl *Resource) LifecycleManagement() *terra.Lifecycle {
	return aocl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aocl *Resource) Attributes() awsOpsworksCustomLayerAttributes {
	return awsOpsworksCustomLayerAttributes{ref: terra.ReferenceResource(aocl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aocl *Resource) ImportState(state io.Reader) error {
	aocl.state = &awsOpsworksCustomLayerState{}
	if err := json.NewDecoder(state).Decode(aocl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aocl.Type(), aocl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aocl *Resource) State() (*awsOpsworksCustomLayerState, bool) {
	return aocl.state, aocl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aocl *Resource) StateMust() *awsOpsworksCustomLayerState {
	if aocl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aocl.Type(), aocl.LocalName()))
	}
	return aocl.state
}

// Args contains the configurations for aws_opsworks_custom_layer.
type Args struct {
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
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
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
	CloudwatchConfiguration *CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}

type awsOpsworksCustomLayerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(aocl.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(aocl.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(aocl.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aocl.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aocl.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aocl.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aocl.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aocl.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aocl.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(aocl.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(aocl.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(aocl.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("name"))
}

// ShortName returns a reference to field short_name of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("short_name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(aocl.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aocl.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aocl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aocl.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_custom_layer.
func (aocl awsOpsworksCustomLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(aocl.ref.Append("use_ebs_optimized_instances"))
}

func (aocl awsOpsworksCustomLayerAttributes) CloudwatchConfiguration() terra.ListValue[CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[CloudwatchConfigurationAttributes](aocl.ref.Append("cloudwatch_configuration"))
}

func (aocl awsOpsworksCustomLayerAttributes) EbsVolume() terra.SetValue[EbsVolumeAttributes] {
	return terra.ReferenceAsSet[EbsVolumeAttributes](aocl.ref.Append("ebs_volume"))
}

func (aocl awsOpsworksCustomLayerAttributes) LoadBasedAutoScaling() terra.ListValue[LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[LoadBasedAutoScalingAttributes](aocl.ref.Append("load_based_auto_scaling"))
}

type awsOpsworksCustomLayerState struct {
	Arn                      string                         `json:"arn"`
	AutoAssignElasticIps     bool                           `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                           `json:"auto_assign_public_ips"`
	AutoHealing              bool                           `json:"auto_healing"`
	CustomConfigureRecipes   []string                       `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                       `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string                         `json:"custom_instance_profile_arn"`
	CustomJson               string                         `json:"custom_json"`
	CustomSecurityGroupIds   []string                       `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string                       `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                       `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string                       `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                           `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string                         `json:"elastic_load_balancer"`
	Id                       string                         `json:"id"`
	InstallUpdatesOnBoot     bool                           `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  float64                        `json:"instance_shutdown_timeout"`
	Name                     string                         `json:"name"`
	ShortName                string                         `json:"short_name"`
	StackId                  string                         `json:"stack_id"`
	SystemPackages           []string                       `json:"system_packages"`
	Tags                     map[string]string              `json:"tags"`
	TagsAll                  map[string]string              `json:"tags_all"`
	UseEbsOptimizedInstances bool                           `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
