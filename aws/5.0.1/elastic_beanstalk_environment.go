// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticbeanstalkenvironment "github.com/golingon/terraproviders/aws/5.0.1/elasticbeanstalkenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticBeanstalkEnvironment creates a new instance of [ElasticBeanstalkEnvironment].
func NewElasticBeanstalkEnvironment(name string, args ElasticBeanstalkEnvironmentArgs) *ElasticBeanstalkEnvironment {
	return &ElasticBeanstalkEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticBeanstalkEnvironment)(nil)

// ElasticBeanstalkEnvironment represents the Terraform resource aws_elastic_beanstalk_environment.
type ElasticBeanstalkEnvironment struct {
	Name      string
	Args      ElasticBeanstalkEnvironmentArgs
	state     *elasticBeanstalkEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticBeanstalkEnvironment].
func (ebe *ElasticBeanstalkEnvironment) Type() string {
	return "aws_elastic_beanstalk_environment"
}

// LocalName returns the local name for [ElasticBeanstalkEnvironment].
func (ebe *ElasticBeanstalkEnvironment) LocalName() string {
	return ebe.Name
}

// Configuration returns the configuration (args) for [ElasticBeanstalkEnvironment].
func (ebe *ElasticBeanstalkEnvironment) Configuration() interface{} {
	return ebe.Args
}

// DependOn is used for other resources to depend on [ElasticBeanstalkEnvironment].
func (ebe *ElasticBeanstalkEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(ebe)
}

// Dependencies returns the list of resources [ElasticBeanstalkEnvironment] depends_on.
func (ebe *ElasticBeanstalkEnvironment) Dependencies() terra.Dependencies {
	return ebe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticBeanstalkEnvironment].
func (ebe *ElasticBeanstalkEnvironment) LifecycleManagement() *terra.Lifecycle {
	return ebe.Lifecycle
}

// Attributes returns the attributes for [ElasticBeanstalkEnvironment].
func (ebe *ElasticBeanstalkEnvironment) Attributes() elasticBeanstalkEnvironmentAttributes {
	return elasticBeanstalkEnvironmentAttributes{ref: terra.ReferenceResource(ebe)}
}

// ImportState imports the given attribute values into [ElasticBeanstalkEnvironment]'s state.
func (ebe *ElasticBeanstalkEnvironment) ImportState(av io.Reader) error {
	ebe.state = &elasticBeanstalkEnvironmentState{}
	if err := json.NewDecoder(av).Decode(ebe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ebe.Type(), ebe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticBeanstalkEnvironment] has state.
func (ebe *ElasticBeanstalkEnvironment) State() (*elasticBeanstalkEnvironmentState, bool) {
	return ebe.state, ebe.state != nil
}

// StateMust returns the state for [ElasticBeanstalkEnvironment]. Panics if the state is nil.
func (ebe *ElasticBeanstalkEnvironment) StateMust() *elasticBeanstalkEnvironmentState {
	if ebe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ebe.Type(), ebe.LocalName()))
	}
	return ebe.state
}

// ElasticBeanstalkEnvironmentArgs contains the configurations for aws_elastic_beanstalk_environment.
type ElasticBeanstalkEnvironmentArgs struct {
	// Application: string, required
	Application terra.StringValue `hcl:"application,attr" validate:"required"`
	// CnamePrefix: string, optional
	CnamePrefix terra.StringValue `hcl:"cname_prefix,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformArn: string, optional
	PlatformArn terra.StringValue `hcl:"platform_arn,attr"`
	// PollInterval: string, optional
	PollInterval terra.StringValue `hcl:"poll_interval,attr"`
	// SolutionStackName: string, optional
	SolutionStackName terra.StringValue `hcl:"solution_stack_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TemplateName: string, optional
	TemplateName terra.StringValue `hcl:"template_name,attr"`
	// Tier: string, optional
	Tier terra.StringValue `hcl:"tier,attr"`
	// VersionLabel: string, optional
	VersionLabel terra.StringValue `hcl:"version_label,attr"`
	// WaitForReadyTimeout: string, optional
	WaitForReadyTimeout terra.StringValue `hcl:"wait_for_ready_timeout,attr"`
	// AllSettings: min=0
	AllSettings []elasticbeanstalkenvironment.AllSettings `hcl:"all_settings,block" validate:"min=0"`
	// Setting: min=0
	Setting []elasticbeanstalkenvironment.Setting `hcl:"setting,block" validate:"min=0"`
}
type elasticBeanstalkEnvironmentAttributes struct {
	ref terra.Reference
}

// Application returns a reference to field application of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Application() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("application"))
}

// Arn returns a reference to field arn of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("arn"))
}

// AutoscalingGroups returns a reference to field autoscaling_groups of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) AutoscalingGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ebe.ref.Append("autoscaling_groups"))
}

// Cname returns a reference to field cname of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Cname() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("cname"))
}

// CnamePrefix returns a reference to field cname_prefix of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) CnamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("cname_prefix"))
}

// Description returns a reference to field description of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("description"))
}

// EndpointUrl returns a reference to field endpoint_url of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) EndpointUrl() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("endpoint_url"))
}

// Id returns a reference to field id of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("id"))
}

// Instances returns a reference to field instances of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Instances() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ebe.ref.Append("instances"))
}

// LaunchConfigurations returns a reference to field launch_configurations of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) LaunchConfigurations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ebe.ref.Append("launch_configurations"))
}

// LoadBalancers returns a reference to field load_balancers of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) LoadBalancers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ebe.ref.Append("load_balancers"))
}

// Name returns a reference to field name of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("name"))
}

// PlatformArn returns a reference to field platform_arn of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) PlatformArn() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("platform_arn"))
}

// PollInterval returns a reference to field poll_interval of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) PollInterval() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("poll_interval"))
}

// Queues returns a reference to field queues of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Queues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ebe.ref.Append("queues"))
}

// SolutionStackName returns a reference to field solution_stack_name of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) SolutionStackName() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("solution_stack_name"))
}

// Tags returns a reference to field tags of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ebe.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ebe.ref.Append("tags_all"))
}

// TemplateName returns a reference to field template_name of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) TemplateName() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("template_name"))
}

// Tier returns a reference to field tier of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("tier"))
}

// Triggers returns a reference to field triggers of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) Triggers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ebe.ref.Append("triggers"))
}

// VersionLabel returns a reference to field version_label of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) VersionLabel() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("version_label"))
}

// WaitForReadyTimeout returns a reference to field wait_for_ready_timeout of aws_elastic_beanstalk_environment.
func (ebe elasticBeanstalkEnvironmentAttributes) WaitForReadyTimeout() terra.StringValue {
	return terra.ReferenceAsString(ebe.ref.Append("wait_for_ready_timeout"))
}

func (ebe elasticBeanstalkEnvironmentAttributes) AllSettings() terra.SetValue[elasticbeanstalkenvironment.AllSettingsAttributes] {
	return terra.ReferenceAsSet[elasticbeanstalkenvironment.AllSettingsAttributes](ebe.ref.Append("all_settings"))
}

func (ebe elasticBeanstalkEnvironmentAttributes) Setting() terra.SetValue[elasticbeanstalkenvironment.SettingAttributes] {
	return terra.ReferenceAsSet[elasticbeanstalkenvironment.SettingAttributes](ebe.ref.Append("setting"))
}

type elasticBeanstalkEnvironmentState struct {
	Application          string                                         `json:"application"`
	Arn                  string                                         `json:"arn"`
	AutoscalingGroups    []string                                       `json:"autoscaling_groups"`
	Cname                string                                         `json:"cname"`
	CnamePrefix          string                                         `json:"cname_prefix"`
	Description          string                                         `json:"description"`
	EndpointUrl          string                                         `json:"endpoint_url"`
	Id                   string                                         `json:"id"`
	Instances            []string                                       `json:"instances"`
	LaunchConfigurations []string                                       `json:"launch_configurations"`
	LoadBalancers        []string                                       `json:"load_balancers"`
	Name                 string                                         `json:"name"`
	PlatformArn          string                                         `json:"platform_arn"`
	PollInterval         string                                         `json:"poll_interval"`
	Queues               []string                                       `json:"queues"`
	SolutionStackName    string                                         `json:"solution_stack_name"`
	Tags                 map[string]string                              `json:"tags"`
	TagsAll              map[string]string                              `json:"tags_all"`
	TemplateName         string                                         `json:"template_name"`
	Tier                 string                                         `json:"tier"`
	Triggers             []string                                       `json:"triggers"`
	VersionLabel         string                                         `json:"version_label"`
	WaitForReadyTimeout  string                                         `json:"wait_for_ready_timeout"`
	AllSettings          []elasticbeanstalkenvironment.AllSettingsState `json:"all_settings"`
	Setting              []elasticbeanstalkenvironment.SettingState     `json:"setting"`
}
