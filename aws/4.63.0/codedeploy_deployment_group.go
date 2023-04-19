// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codedeploydeploymentgroup "github.com/golingon/terraproviders/aws/4.63.0/codedeploydeploymentgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodedeployDeploymentGroup creates a new instance of [CodedeployDeploymentGroup].
func NewCodedeployDeploymentGroup(name string, args CodedeployDeploymentGroupArgs) *CodedeployDeploymentGroup {
	return &CodedeployDeploymentGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodedeployDeploymentGroup)(nil)

// CodedeployDeploymentGroup represents the Terraform resource aws_codedeploy_deployment_group.
type CodedeployDeploymentGroup struct {
	Name      string
	Args      CodedeployDeploymentGroupArgs
	state     *codedeployDeploymentGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodedeployDeploymentGroup].
func (cdg *CodedeployDeploymentGroup) Type() string {
	return "aws_codedeploy_deployment_group"
}

// LocalName returns the local name for [CodedeployDeploymentGroup].
func (cdg *CodedeployDeploymentGroup) LocalName() string {
	return cdg.Name
}

// Configuration returns the configuration (args) for [CodedeployDeploymentGroup].
func (cdg *CodedeployDeploymentGroup) Configuration() interface{} {
	return cdg.Args
}

// DependOn is used for other resources to depend on [CodedeployDeploymentGroup].
func (cdg *CodedeployDeploymentGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cdg)
}

// Dependencies returns the list of resources [CodedeployDeploymentGroup] depends_on.
func (cdg *CodedeployDeploymentGroup) Dependencies() terra.Dependencies {
	return cdg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodedeployDeploymentGroup].
func (cdg *CodedeployDeploymentGroup) LifecycleManagement() *terra.Lifecycle {
	return cdg.Lifecycle
}

// Attributes returns the attributes for [CodedeployDeploymentGroup].
func (cdg *CodedeployDeploymentGroup) Attributes() codedeployDeploymentGroupAttributes {
	return codedeployDeploymentGroupAttributes{ref: terra.ReferenceResource(cdg)}
}

// ImportState imports the given attribute values into [CodedeployDeploymentGroup]'s state.
func (cdg *CodedeployDeploymentGroup) ImportState(av io.Reader) error {
	cdg.state = &codedeployDeploymentGroupState{}
	if err := json.NewDecoder(av).Decode(cdg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdg.Type(), cdg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodedeployDeploymentGroup] has state.
func (cdg *CodedeployDeploymentGroup) State() (*codedeployDeploymentGroupState, bool) {
	return cdg.state, cdg.state != nil
}

// StateMust returns the state for [CodedeployDeploymentGroup]. Panics if the state is nil.
func (cdg *CodedeployDeploymentGroup) StateMust() *codedeployDeploymentGroupState {
	if cdg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdg.Type(), cdg.LocalName()))
	}
	return cdg.state
}

// CodedeployDeploymentGroupArgs contains the configurations for aws_codedeploy_deployment_group.
type CodedeployDeploymentGroupArgs struct {
	// AppName: string, required
	AppName terra.StringValue `hcl:"app_name,attr" validate:"required"`
	// AutoscalingGroups: set of string, optional
	AutoscalingGroups terra.SetValue[terra.StringValue] `hcl:"autoscaling_groups,attr"`
	// DeploymentConfigName: string, optional
	DeploymentConfigName terra.StringValue `hcl:"deployment_config_name,attr"`
	// DeploymentGroupName: string, required
	DeploymentGroupName terra.StringValue `hcl:"deployment_group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceRoleArn: string, required
	ServiceRoleArn terra.StringValue `hcl:"service_role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AlarmConfiguration: optional
	AlarmConfiguration *codedeploydeploymentgroup.AlarmConfiguration `hcl:"alarm_configuration,block"`
	// AutoRollbackConfiguration: optional
	AutoRollbackConfiguration *codedeploydeploymentgroup.AutoRollbackConfiguration `hcl:"auto_rollback_configuration,block"`
	// BlueGreenDeploymentConfig: optional
	BlueGreenDeploymentConfig *codedeploydeploymentgroup.BlueGreenDeploymentConfig `hcl:"blue_green_deployment_config,block"`
	// DeploymentStyle: optional
	DeploymentStyle *codedeploydeploymentgroup.DeploymentStyle `hcl:"deployment_style,block"`
	// Ec2TagFilter: min=0
	Ec2TagFilter []codedeploydeploymentgroup.Ec2TagFilter `hcl:"ec2_tag_filter,block" validate:"min=0"`
	// Ec2TagSet: min=0
	Ec2TagSet []codedeploydeploymentgroup.Ec2TagSet `hcl:"ec2_tag_set,block" validate:"min=0"`
	// EcsService: optional
	EcsService *codedeploydeploymentgroup.EcsService `hcl:"ecs_service,block"`
	// LoadBalancerInfo: optional
	LoadBalancerInfo *codedeploydeploymentgroup.LoadBalancerInfo `hcl:"load_balancer_info,block"`
	// OnPremisesInstanceTagFilter: min=0
	OnPremisesInstanceTagFilter []codedeploydeploymentgroup.OnPremisesInstanceTagFilter `hcl:"on_premises_instance_tag_filter,block" validate:"min=0"`
	// TriggerConfiguration: min=0
	TriggerConfiguration []codedeploydeploymentgroup.TriggerConfiguration `hcl:"trigger_configuration,block" validate:"min=0"`
}
type codedeployDeploymentGroupAttributes struct {
	ref terra.Reference
}

// AppName returns a reference to field app_name of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) AppName() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("app_name"))
}

// Arn returns a reference to field arn of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("arn"))
}

// AutoscalingGroups returns a reference to field autoscaling_groups of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) AutoscalingGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cdg.ref.Append("autoscaling_groups"))
}

// ComputePlatform returns a reference to field compute_platform of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) ComputePlatform() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("compute_platform"))
}

// DeploymentConfigName returns a reference to field deployment_config_name of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) DeploymentConfigName() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("deployment_config_name"))
}

// DeploymentGroupId returns a reference to field deployment_group_id of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) DeploymentGroupId() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("deployment_group_id"))
}

// DeploymentGroupName returns a reference to field deployment_group_name of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) DeploymentGroupName() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("deployment_group_name"))
}

// Id returns a reference to field id of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("id"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(cdg.ref.Append("service_role_arn"))
}

// Tags returns a reference to field tags of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codedeploy_deployment_group.
func (cdg codedeployDeploymentGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdg.ref.Append("tags_all"))
}

func (cdg codedeployDeploymentGroupAttributes) AlarmConfiguration() terra.ListValue[codedeploydeploymentgroup.AlarmConfigurationAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentgroup.AlarmConfigurationAttributes](cdg.ref.Append("alarm_configuration"))
}

func (cdg codedeployDeploymentGroupAttributes) AutoRollbackConfiguration() terra.ListValue[codedeploydeploymentgroup.AutoRollbackConfigurationAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentgroup.AutoRollbackConfigurationAttributes](cdg.ref.Append("auto_rollback_configuration"))
}

func (cdg codedeployDeploymentGroupAttributes) BlueGreenDeploymentConfig() terra.ListValue[codedeploydeploymentgroup.BlueGreenDeploymentConfigAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentgroup.BlueGreenDeploymentConfigAttributes](cdg.ref.Append("blue_green_deployment_config"))
}

func (cdg codedeployDeploymentGroupAttributes) DeploymentStyle() terra.ListValue[codedeploydeploymentgroup.DeploymentStyleAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentgroup.DeploymentStyleAttributes](cdg.ref.Append("deployment_style"))
}

func (cdg codedeployDeploymentGroupAttributes) Ec2TagFilter() terra.SetValue[codedeploydeploymentgroup.Ec2TagFilterAttributes] {
	return terra.ReferenceAsSet[codedeploydeploymentgroup.Ec2TagFilterAttributes](cdg.ref.Append("ec2_tag_filter"))
}

func (cdg codedeployDeploymentGroupAttributes) Ec2TagSet() terra.SetValue[codedeploydeploymentgroup.Ec2TagSetAttributes] {
	return terra.ReferenceAsSet[codedeploydeploymentgroup.Ec2TagSetAttributes](cdg.ref.Append("ec2_tag_set"))
}

func (cdg codedeployDeploymentGroupAttributes) EcsService() terra.ListValue[codedeploydeploymentgroup.EcsServiceAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentgroup.EcsServiceAttributes](cdg.ref.Append("ecs_service"))
}

func (cdg codedeployDeploymentGroupAttributes) LoadBalancerInfo() terra.ListValue[codedeploydeploymentgroup.LoadBalancerInfoAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentgroup.LoadBalancerInfoAttributes](cdg.ref.Append("load_balancer_info"))
}

func (cdg codedeployDeploymentGroupAttributes) OnPremisesInstanceTagFilter() terra.SetValue[codedeploydeploymentgroup.OnPremisesInstanceTagFilterAttributes] {
	return terra.ReferenceAsSet[codedeploydeploymentgroup.OnPremisesInstanceTagFilterAttributes](cdg.ref.Append("on_premises_instance_tag_filter"))
}

func (cdg codedeployDeploymentGroupAttributes) TriggerConfiguration() terra.SetValue[codedeploydeploymentgroup.TriggerConfigurationAttributes] {
	return terra.ReferenceAsSet[codedeploydeploymentgroup.TriggerConfigurationAttributes](cdg.ref.Append("trigger_configuration"))
}

type codedeployDeploymentGroupState struct {
	AppName                     string                                                       `json:"app_name"`
	Arn                         string                                                       `json:"arn"`
	AutoscalingGroups           []string                                                     `json:"autoscaling_groups"`
	ComputePlatform             string                                                       `json:"compute_platform"`
	DeploymentConfigName        string                                                       `json:"deployment_config_name"`
	DeploymentGroupId           string                                                       `json:"deployment_group_id"`
	DeploymentGroupName         string                                                       `json:"deployment_group_name"`
	Id                          string                                                       `json:"id"`
	ServiceRoleArn              string                                                       `json:"service_role_arn"`
	Tags                        map[string]string                                            `json:"tags"`
	TagsAll                     map[string]string                                            `json:"tags_all"`
	AlarmConfiguration          []codedeploydeploymentgroup.AlarmConfigurationState          `json:"alarm_configuration"`
	AutoRollbackConfiguration   []codedeploydeploymentgroup.AutoRollbackConfigurationState   `json:"auto_rollback_configuration"`
	BlueGreenDeploymentConfig   []codedeploydeploymentgroup.BlueGreenDeploymentConfigState   `json:"blue_green_deployment_config"`
	DeploymentStyle             []codedeploydeploymentgroup.DeploymentStyleState             `json:"deployment_style"`
	Ec2TagFilter                []codedeploydeploymentgroup.Ec2TagFilterState                `json:"ec2_tag_filter"`
	Ec2TagSet                   []codedeploydeploymentgroup.Ec2TagSetState                   `json:"ec2_tag_set"`
	EcsService                  []codedeploydeploymentgroup.EcsServiceState                  `json:"ecs_service"`
	LoadBalancerInfo            []codedeploydeploymentgroup.LoadBalancerInfoState            `json:"load_balancer_info"`
	OnPremisesInstanceTagFilter []codedeploydeploymentgroup.OnPremisesInstanceTagFilterState `json:"on_premises_instance_tag_filter"`
	TriggerConfiguration        []codedeploydeploymentgroup.TriggerConfigurationState        `json:"trigger_configuration"`
}
