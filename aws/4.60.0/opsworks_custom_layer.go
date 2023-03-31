// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworkscustomlayer "github.com/golingon/terraproviders/aws/4.60.0/opsworkscustomlayer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksCustomLayer creates a new instance of [OpsworksCustomLayer].
func NewOpsworksCustomLayer(name string, args OpsworksCustomLayerArgs) *OpsworksCustomLayer {
	return &OpsworksCustomLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksCustomLayer)(nil)

// OpsworksCustomLayer represents the Terraform resource aws_opsworks_custom_layer.
type OpsworksCustomLayer struct {
	Name      string
	Args      OpsworksCustomLayerArgs
	state     *opsworksCustomLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksCustomLayer].
func (ocl *OpsworksCustomLayer) Type() string {
	return "aws_opsworks_custom_layer"
}

// LocalName returns the local name for [OpsworksCustomLayer].
func (ocl *OpsworksCustomLayer) LocalName() string {
	return ocl.Name
}

// Configuration returns the configuration (args) for [OpsworksCustomLayer].
func (ocl *OpsworksCustomLayer) Configuration() interface{} {
	return ocl.Args
}

// DependOn is used for other resources to depend on [OpsworksCustomLayer].
func (ocl *OpsworksCustomLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(ocl)
}

// Dependencies returns the list of resources [OpsworksCustomLayer] depends_on.
func (ocl *OpsworksCustomLayer) Dependencies() terra.Dependencies {
	return ocl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksCustomLayer].
func (ocl *OpsworksCustomLayer) LifecycleManagement() *terra.Lifecycle {
	return ocl.Lifecycle
}

// Attributes returns the attributes for [OpsworksCustomLayer].
func (ocl *OpsworksCustomLayer) Attributes() opsworksCustomLayerAttributes {
	return opsworksCustomLayerAttributes{ref: terra.ReferenceResource(ocl)}
}

// ImportState imports the given attribute values into [OpsworksCustomLayer]'s state.
func (ocl *OpsworksCustomLayer) ImportState(av io.Reader) error {
	ocl.state = &opsworksCustomLayerState{}
	if err := json.NewDecoder(av).Decode(ocl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ocl.Type(), ocl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksCustomLayer] has state.
func (ocl *OpsworksCustomLayer) State() (*opsworksCustomLayerState, bool) {
	return ocl.state, ocl.state != nil
}

// StateMust returns the state for [OpsworksCustomLayer]. Panics if the state is nil.
func (ocl *OpsworksCustomLayer) StateMust() *opsworksCustomLayerState {
	if ocl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ocl.Type(), ocl.LocalName()))
	}
	return ocl.state
}

// OpsworksCustomLayerArgs contains the configurations for aws_opsworks_custom_layer.
type OpsworksCustomLayerArgs struct {
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
	CloudwatchConfiguration *opsworkscustomlayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworkscustomlayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworkscustomlayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksCustomLayerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(ocl.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(ocl.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(ocl.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ocl.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ocl.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ocl.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ocl.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ocl.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ocl.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(ocl.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(ocl.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ocl.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("name"))
}

// ShortName returns a reference to field short_name of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("short_name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(ocl.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ocl.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ocl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ocl.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_custom_layer.
func (ocl opsworksCustomLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(ocl.ref.Append("use_ebs_optimized_instances"))
}

func (ocl opsworksCustomLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworkscustomlayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworkscustomlayer.CloudwatchConfigurationAttributes](ocl.ref.Append("cloudwatch_configuration"))
}

func (ocl opsworksCustomLayerAttributes) EbsVolume() terra.SetValue[opsworkscustomlayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworkscustomlayer.EbsVolumeAttributes](ocl.ref.Append("ebs_volume"))
}

func (ocl opsworksCustomLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworkscustomlayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworkscustomlayer.LoadBasedAutoScalingAttributes](ocl.ref.Append("load_based_auto_scaling"))
}

type opsworksCustomLayerState struct {
	Arn                      string                                             `json:"arn"`
	AutoAssignElasticIps     bool                                               `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool                                               `json:"auto_assign_public_ips"`
	AutoHealing              bool                                               `json:"auto_healing"`
	CustomConfigureRecipes   []string                                           `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string                                           `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string                                             `json:"custom_instance_profile_arn"`
	CustomJson               string                                             `json:"custom_json"`
	CustomSecurityGroupIds   []string                                           `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string                                           `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string                                           `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string                                           `json:"custom_undeploy_recipes"`
	DrainElbOnShutdown       bool                                               `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string                                             `json:"elastic_load_balancer"`
	Id                       string                                             `json:"id"`
	InstallUpdatesOnBoot     bool                                               `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  float64                                            `json:"instance_shutdown_timeout"`
	Name                     string                                             `json:"name"`
	ShortName                string                                             `json:"short_name"`
	StackId                  string                                             `json:"stack_id"`
	SystemPackages           []string                                           `json:"system_packages"`
	Tags                     map[string]string                                  `json:"tags"`
	TagsAll                  map[string]string                                  `json:"tags_all"`
	UseEbsOptimizedInstances bool                                               `json:"use_ebs_optimized_instances"`
	CloudwatchConfiguration  []opsworkscustomlayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworkscustomlayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworkscustomlayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
