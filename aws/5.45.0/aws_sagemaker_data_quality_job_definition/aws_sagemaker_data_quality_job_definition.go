// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sagemaker_data_quality_job_definition

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

// Resource represents the Terraform resource aws_sagemaker_data_quality_job_definition.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSagemakerDataQualityJobDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asdqjd *Resource) Type() string {
	return "aws_sagemaker_data_quality_job_definition"
}

// LocalName returns the local name for [Resource].
func (asdqjd *Resource) LocalName() string {
	return asdqjd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asdqjd *Resource) Configuration() interface{} {
	return asdqjd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asdqjd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asdqjd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asdqjd *Resource) Dependencies() terra.Dependencies {
	return asdqjd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asdqjd *Resource) LifecycleManagement() *terra.Lifecycle {
	return asdqjd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asdqjd *Resource) Attributes() awsSagemakerDataQualityJobDefinitionAttributes {
	return awsSagemakerDataQualityJobDefinitionAttributes{ref: terra.ReferenceResource(asdqjd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asdqjd *Resource) ImportState(state io.Reader) error {
	asdqjd.state = &awsSagemakerDataQualityJobDefinitionState{}
	if err := json.NewDecoder(state).Decode(asdqjd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asdqjd.Type(), asdqjd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asdqjd *Resource) State() (*awsSagemakerDataQualityJobDefinitionState, bool) {
	return asdqjd.state, asdqjd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asdqjd *Resource) StateMust() *awsSagemakerDataQualityJobDefinitionState {
	if asdqjd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asdqjd.Type(), asdqjd.LocalName()))
	}
	return asdqjd.state
}

// Args contains the configurations for aws_sagemaker_data_quality_job_definition.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DataQualityAppSpecification: required
	DataQualityAppSpecification *DataQualityAppSpecification `hcl:"data_quality_app_specification,block" validate:"required"`
	// DataQualityBaselineConfig: optional
	DataQualityBaselineConfig *DataQualityBaselineConfig `hcl:"data_quality_baseline_config,block"`
	// DataQualityJobInput: required
	DataQualityJobInput *DataQualityJobInput `hcl:"data_quality_job_input,block" validate:"required"`
	// DataQualityJobOutputConfig: required
	DataQualityJobOutputConfig *DataQualityJobOutputConfig `hcl:"data_quality_job_output_config,block" validate:"required"`
	// JobResources: required
	JobResources *JobResources `hcl:"job_resources,block" validate:"required"`
	// NetworkConfig: optional
	NetworkConfig *NetworkConfig `hcl:"network_config,block"`
	// StoppingCondition: optional
	StoppingCondition *StoppingCondition `hcl:"stopping_condition,block"`
}

type awsSagemakerDataQualityJobDefinitionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_data_quality_job_definition.
func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asdqjd.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sagemaker_data_quality_job_definition.
func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asdqjd.ref.Append("id"))
}

// Name returns a reference to field name of aws_sagemaker_data_quality_job_definition.
func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asdqjd.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_data_quality_job_definition.
func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(asdqjd.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_sagemaker_data_quality_job_definition.
func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asdqjd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_data_quality_job_definition.
func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asdqjd.ref.Append("tags_all"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) DataQualityAppSpecification() terra.ListValue[DataQualityAppSpecificationAttributes] {
	return terra.ReferenceAsList[DataQualityAppSpecificationAttributes](asdqjd.ref.Append("data_quality_app_specification"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) DataQualityBaselineConfig() terra.ListValue[DataQualityBaselineConfigAttributes] {
	return terra.ReferenceAsList[DataQualityBaselineConfigAttributes](asdqjd.ref.Append("data_quality_baseline_config"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) DataQualityJobInput() terra.ListValue[DataQualityJobInputAttributes] {
	return terra.ReferenceAsList[DataQualityJobInputAttributes](asdqjd.ref.Append("data_quality_job_input"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) DataQualityJobOutputConfig() terra.ListValue[DataQualityJobOutputConfigAttributes] {
	return terra.ReferenceAsList[DataQualityJobOutputConfigAttributes](asdqjd.ref.Append("data_quality_job_output_config"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) JobResources() terra.ListValue[JobResourcesAttributes] {
	return terra.ReferenceAsList[JobResourcesAttributes](asdqjd.ref.Append("job_resources"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) NetworkConfig() terra.ListValue[NetworkConfigAttributes] {
	return terra.ReferenceAsList[NetworkConfigAttributes](asdqjd.ref.Append("network_config"))
}

func (asdqjd awsSagemakerDataQualityJobDefinitionAttributes) StoppingCondition() terra.ListValue[StoppingConditionAttributes] {
	return terra.ReferenceAsList[StoppingConditionAttributes](asdqjd.ref.Append("stopping_condition"))
}

type awsSagemakerDataQualityJobDefinitionState struct {
	Arn                         string                             `json:"arn"`
	Id                          string                             `json:"id"`
	Name                        string                             `json:"name"`
	RoleArn                     string                             `json:"role_arn"`
	Tags                        map[string]string                  `json:"tags"`
	TagsAll                     map[string]string                  `json:"tags_all"`
	DataQualityAppSpecification []DataQualityAppSpecificationState `json:"data_quality_app_specification"`
	DataQualityBaselineConfig   []DataQualityBaselineConfigState   `json:"data_quality_baseline_config"`
	DataQualityJobInput         []DataQualityJobInputState         `json:"data_quality_job_input"`
	DataQualityJobOutputConfig  []DataQualityJobOutputConfigState  `json:"data_quality_job_output_config"`
	JobResources                []JobResourcesState                `json:"job_resources"`
	NetworkConfig               []NetworkConfigState               `json:"network_config"`
	StoppingCondition           []StoppingConditionState           `json:"stopping_condition"`
}
