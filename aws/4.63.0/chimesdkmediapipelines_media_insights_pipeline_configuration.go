// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimesdkmediapipelinesmediainsightspipelineconfiguration "github.com/golingon/terraproviders/aws/4.63.0/chimesdkmediapipelinesmediainsightspipelineconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimesdkmediapipelinesMediaInsightsPipelineConfiguration creates a new instance of [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func NewChimesdkmediapipelinesMediaInsightsPipelineConfiguration(name string, args ChimesdkmediapipelinesMediaInsightsPipelineConfigurationArgs) *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration {
	return &ChimesdkmediapipelinesMediaInsightsPipelineConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimesdkmediapipelinesMediaInsightsPipelineConfiguration)(nil)

// ChimesdkmediapipelinesMediaInsightsPipelineConfiguration represents the Terraform resource aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
type ChimesdkmediapipelinesMediaInsightsPipelineConfiguration struct {
	Name      string
	Args      ChimesdkmediapipelinesMediaInsightsPipelineConfigurationArgs
	state     *chimesdkmediapipelinesMediaInsightsPipelineConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) Type() string {
	return "aws_chimesdkmediapipelines_media_insights_pipeline_configuration"
}

// LocalName returns the local name for [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) LocalName() string {
	return cmipc.Name
}

// Configuration returns the configuration (args) for [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) Configuration() interface{} {
	return cmipc.Args
}

// DependOn is used for other resources to depend on [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(cmipc)
}

// Dependencies returns the list of resources [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration] depends_on.
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) Dependencies() terra.Dependencies {
	return cmipc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) LifecycleManagement() *terra.Lifecycle {
	return cmipc.Lifecycle
}

// Attributes returns the attributes for [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration].
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) Attributes() chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes {
	return chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes{ref: terra.ReferenceResource(cmipc)}
}

// ImportState imports the given attribute values into [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration]'s state.
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) ImportState(av io.Reader) error {
	cmipc.state = &chimesdkmediapipelinesMediaInsightsPipelineConfigurationState{}
	if err := json.NewDecoder(av).Decode(cmipc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmipc.Type(), cmipc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration] has state.
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) State() (*chimesdkmediapipelinesMediaInsightsPipelineConfigurationState, bool) {
	return cmipc.state, cmipc.state != nil
}

// StateMust returns the state for [ChimesdkmediapipelinesMediaInsightsPipelineConfiguration]. Panics if the state is nil.
func (cmipc *ChimesdkmediapipelinesMediaInsightsPipelineConfiguration) StateMust() *chimesdkmediapipelinesMediaInsightsPipelineConfigurationState {
	if cmipc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmipc.Type(), cmipc.LocalName()))
	}
	return cmipc.state
}

// ChimesdkmediapipelinesMediaInsightsPipelineConfigurationArgs contains the configurations for aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceAccessRoleArn: string, required
	ResourceAccessRoleArn terra.StringValue `hcl:"resource_access_role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Elements: min=1
	Elements []chimesdkmediapipelinesmediainsightspipelineconfiguration.Elements `hcl:"elements,block" validate:"min=1"`
	// RealTimeAlertConfiguration: optional
	RealTimeAlertConfiguration *chimesdkmediapipelinesmediainsightspipelineconfiguration.RealTimeAlertConfiguration `hcl:"real_time_alert_configuration,block"`
	// Timeouts: optional
	Timeouts *chimesdkmediapipelinesmediainsightspipelineconfiguration.Timeouts `hcl:"timeouts,block"`
}
type chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cmipc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmipc.ref.Append("id"))
}

// Name returns a reference to field name of aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmipc.ref.Append("name"))
}

// ResourceAccessRoleArn returns a reference to field resource_access_role_arn of aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) ResourceAccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(cmipc.ref.Append("resource_access_role_arn"))
}

// Tags returns a reference to field tags of aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmipc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_chimesdkmediapipelines_media_insights_pipeline_configuration.
func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmipc.ref.Append("tags_all"))
}

func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) Elements() terra.ListValue[chimesdkmediapipelinesmediainsightspipelineconfiguration.ElementsAttributes] {
	return terra.ReferenceAsList[chimesdkmediapipelinesmediainsightspipelineconfiguration.ElementsAttributes](cmipc.ref.Append("elements"))
}

func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) RealTimeAlertConfiguration() terra.ListValue[chimesdkmediapipelinesmediainsightspipelineconfiguration.RealTimeAlertConfigurationAttributes] {
	return terra.ReferenceAsList[chimesdkmediapipelinesmediainsightspipelineconfiguration.RealTimeAlertConfigurationAttributes](cmipc.ref.Append("real_time_alert_configuration"))
}

func (cmipc chimesdkmediapipelinesMediaInsightsPipelineConfigurationAttributes) Timeouts() chimesdkmediapipelinesmediainsightspipelineconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[chimesdkmediapipelinesmediainsightspipelineconfiguration.TimeoutsAttributes](cmipc.ref.Append("timeouts"))
}

type chimesdkmediapipelinesMediaInsightsPipelineConfigurationState struct {
	Arn                        string                                                                                     `json:"arn"`
	Id                         string                                                                                     `json:"id"`
	Name                       string                                                                                     `json:"name"`
	ResourceAccessRoleArn      string                                                                                     `json:"resource_access_role_arn"`
	Tags                       map[string]string                                                                          `json:"tags"`
	TagsAll                    map[string]string                                                                          `json:"tags_all"`
	Elements                   []chimesdkmediapipelinesmediainsightspipelineconfiguration.ElementsState                   `json:"elements"`
	RealTimeAlertConfiguration []chimesdkmediapipelinesmediainsightspipelineconfiguration.RealTimeAlertConfigurationState `json:"real_time_alert_configuration"`
	Timeouts                   *chimesdkmediapipelinesmediainsightspipelineconfiguration.TimeoutsState                    `json:"timeouts"`
}
