// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksstaticweblayer "github.com/golingon/terraproviders/aws/5.8.0/opsworksstaticweblayer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksStaticWebLayer creates a new instance of [OpsworksStaticWebLayer].
func NewOpsworksStaticWebLayer(name string, args OpsworksStaticWebLayerArgs) *OpsworksStaticWebLayer {
	return &OpsworksStaticWebLayer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksStaticWebLayer)(nil)

// OpsworksStaticWebLayer represents the Terraform resource aws_opsworks_static_web_layer.
type OpsworksStaticWebLayer struct {
	Name      string
	Args      OpsworksStaticWebLayerArgs
	state     *opsworksStaticWebLayerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksStaticWebLayer].
func (oswl *OpsworksStaticWebLayer) Type() string {
	return "aws_opsworks_static_web_layer"
}

// LocalName returns the local name for [OpsworksStaticWebLayer].
func (oswl *OpsworksStaticWebLayer) LocalName() string {
	return oswl.Name
}

// Configuration returns the configuration (args) for [OpsworksStaticWebLayer].
func (oswl *OpsworksStaticWebLayer) Configuration() interface{} {
	return oswl.Args
}

// DependOn is used for other resources to depend on [OpsworksStaticWebLayer].
func (oswl *OpsworksStaticWebLayer) DependOn() terra.Reference {
	return terra.ReferenceResource(oswl)
}

// Dependencies returns the list of resources [OpsworksStaticWebLayer] depends_on.
func (oswl *OpsworksStaticWebLayer) Dependencies() terra.Dependencies {
	return oswl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksStaticWebLayer].
func (oswl *OpsworksStaticWebLayer) LifecycleManagement() *terra.Lifecycle {
	return oswl.Lifecycle
}

// Attributes returns the attributes for [OpsworksStaticWebLayer].
func (oswl *OpsworksStaticWebLayer) Attributes() opsworksStaticWebLayerAttributes {
	return opsworksStaticWebLayerAttributes{ref: terra.ReferenceResource(oswl)}
}

// ImportState imports the given attribute values into [OpsworksStaticWebLayer]'s state.
func (oswl *OpsworksStaticWebLayer) ImportState(av io.Reader) error {
	oswl.state = &opsworksStaticWebLayerState{}
	if err := json.NewDecoder(av).Decode(oswl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oswl.Type(), oswl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksStaticWebLayer] has state.
func (oswl *OpsworksStaticWebLayer) State() (*opsworksStaticWebLayerState, bool) {
	return oswl.state, oswl.state != nil
}

// StateMust returns the state for [OpsworksStaticWebLayer]. Panics if the state is nil.
func (oswl *OpsworksStaticWebLayer) StateMust() *opsworksStaticWebLayerState {
	if oswl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oswl.Type(), oswl.LocalName()))
	}
	return oswl.state
}

// OpsworksStaticWebLayerArgs contains the configurations for aws_opsworks_static_web_layer.
type OpsworksStaticWebLayerArgs struct {
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
	CloudwatchConfiguration *opsworksstaticweblayer.CloudwatchConfiguration `hcl:"cloudwatch_configuration,block"`
	// EbsVolume: min=0
	EbsVolume []opsworksstaticweblayer.EbsVolume `hcl:"ebs_volume,block" validate:"min=0"`
	// LoadBasedAutoScaling: optional
	LoadBasedAutoScaling *opsworksstaticweblayer.LoadBasedAutoScaling `hcl:"load_based_auto_scaling,block"`
}
type opsworksStaticWebLayerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("arn"))
}

// AutoAssignElasticIps returns a reference to field auto_assign_elastic_ips of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) AutoAssignElasticIps() terra.BoolValue {
	return terra.ReferenceAsBool(oswl.ref.Append("auto_assign_elastic_ips"))
}

// AutoAssignPublicIps returns a reference to field auto_assign_public_ips of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) AutoAssignPublicIps() terra.BoolValue {
	return terra.ReferenceAsBool(oswl.ref.Append("auto_assign_public_ips"))
}

// AutoHealing returns a reference to field auto_healing of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) AutoHealing() terra.BoolValue {
	return terra.ReferenceAsBool(oswl.ref.Append("auto_healing"))
}

// CustomConfigureRecipes returns a reference to field custom_configure_recipes of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomConfigureRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oswl.ref.Append("custom_configure_recipes"))
}

// CustomDeployRecipes returns a reference to field custom_deploy_recipes of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomDeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oswl.ref.Append("custom_deploy_recipes"))
}

// CustomInstanceProfileArn returns a reference to field custom_instance_profile_arn of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("custom_instance_profile_arn"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("custom_json"))
}

// CustomSecurityGroupIds returns a reference to field custom_security_group_ids of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oswl.ref.Append("custom_security_group_ids"))
}

// CustomSetupRecipes returns a reference to field custom_setup_recipes of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomSetupRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oswl.ref.Append("custom_setup_recipes"))
}

// CustomShutdownRecipes returns a reference to field custom_shutdown_recipes of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomShutdownRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oswl.ref.Append("custom_shutdown_recipes"))
}

// CustomUndeployRecipes returns a reference to field custom_undeploy_recipes of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) CustomUndeployRecipes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oswl.ref.Append("custom_undeploy_recipes"))
}

// DrainElbOnShutdown returns a reference to field drain_elb_on_shutdown of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) DrainElbOnShutdown() terra.BoolValue {
	return terra.ReferenceAsBool(oswl.ref.Append("drain_elb_on_shutdown"))
}

// ElasticLoadBalancer returns a reference to field elastic_load_balancer of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) ElasticLoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("elastic_load_balancer"))
}

// Id returns a reference to field id of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("id"))
}

// InstallUpdatesOnBoot returns a reference to field install_updates_on_boot of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceAsBool(oswl.ref.Append("install_updates_on_boot"))
}

// InstanceShutdownTimeout returns a reference to field instance_shutdown_timeout of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) InstanceShutdownTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(oswl.ref.Append("instance_shutdown_timeout"))
}

// Name returns a reference to field name of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(oswl.ref.Append("stack_id"))
}

// SystemPackages returns a reference to field system_packages of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) SystemPackages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oswl.ref.Append("system_packages"))
}

// Tags returns a reference to field tags of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oswl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oswl.ref.Append("tags_all"))
}

// UseEbsOptimizedInstances returns a reference to field use_ebs_optimized_instances of aws_opsworks_static_web_layer.
func (oswl opsworksStaticWebLayerAttributes) UseEbsOptimizedInstances() terra.BoolValue {
	return terra.ReferenceAsBool(oswl.ref.Append("use_ebs_optimized_instances"))
}

func (oswl opsworksStaticWebLayerAttributes) CloudwatchConfiguration() terra.ListValue[opsworksstaticweblayer.CloudwatchConfigurationAttributes] {
	return terra.ReferenceAsList[opsworksstaticweblayer.CloudwatchConfigurationAttributes](oswl.ref.Append("cloudwatch_configuration"))
}

func (oswl opsworksStaticWebLayerAttributes) EbsVolume() terra.SetValue[opsworksstaticweblayer.EbsVolumeAttributes] {
	return terra.ReferenceAsSet[opsworksstaticweblayer.EbsVolumeAttributes](oswl.ref.Append("ebs_volume"))
}

func (oswl opsworksStaticWebLayerAttributes) LoadBasedAutoScaling() terra.ListValue[opsworksstaticweblayer.LoadBasedAutoScalingAttributes] {
	return terra.ReferenceAsList[opsworksstaticweblayer.LoadBasedAutoScalingAttributes](oswl.ref.Append("load_based_auto_scaling"))
}

type opsworksStaticWebLayerState struct {
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
	CloudwatchConfiguration  []opsworksstaticweblayer.CloudwatchConfigurationState `json:"cloudwatch_configuration"`
	EbsVolume                []opsworksstaticweblayer.EbsVolumeState               `json:"ebs_volume"`
	LoadBasedAutoScaling     []opsworksstaticweblayer.LoadBasedAutoScalingState    `json:"load_based_auto_scaling"`
}
