// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerdataqualityjobdefinition "github.com/golingon/terraproviders/aws/5.11.0/sagemakerdataqualityjobdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerDataQualityJobDefinition creates a new instance of [SagemakerDataQualityJobDefinition].
func NewSagemakerDataQualityJobDefinition(name string, args SagemakerDataQualityJobDefinitionArgs) *SagemakerDataQualityJobDefinition {
	return &SagemakerDataQualityJobDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerDataQualityJobDefinition)(nil)

// SagemakerDataQualityJobDefinition represents the Terraform resource aws_sagemaker_data_quality_job_definition.
type SagemakerDataQualityJobDefinition struct {
	Name      string
	Args      SagemakerDataQualityJobDefinitionArgs
	state     *sagemakerDataQualityJobDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerDataQualityJobDefinition].
func (sdqjd *SagemakerDataQualityJobDefinition) Type() string {
	return "aws_sagemaker_data_quality_job_definition"
}

// LocalName returns the local name for [SagemakerDataQualityJobDefinition].
func (sdqjd *SagemakerDataQualityJobDefinition) LocalName() string {
	return sdqjd.Name
}

// Configuration returns the configuration (args) for [SagemakerDataQualityJobDefinition].
func (sdqjd *SagemakerDataQualityJobDefinition) Configuration() interface{} {
	return sdqjd.Args
}

// DependOn is used for other resources to depend on [SagemakerDataQualityJobDefinition].
func (sdqjd *SagemakerDataQualityJobDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(sdqjd)
}

// Dependencies returns the list of resources [SagemakerDataQualityJobDefinition] depends_on.
func (sdqjd *SagemakerDataQualityJobDefinition) Dependencies() terra.Dependencies {
	return sdqjd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerDataQualityJobDefinition].
func (sdqjd *SagemakerDataQualityJobDefinition) LifecycleManagement() *terra.Lifecycle {
	return sdqjd.Lifecycle
}

// Attributes returns the attributes for [SagemakerDataQualityJobDefinition].
func (sdqjd *SagemakerDataQualityJobDefinition) Attributes() sagemakerDataQualityJobDefinitionAttributes {
	return sagemakerDataQualityJobDefinitionAttributes{ref: terra.ReferenceResource(sdqjd)}
}

// ImportState imports the given attribute values into [SagemakerDataQualityJobDefinition]'s state.
func (sdqjd *SagemakerDataQualityJobDefinition) ImportState(av io.Reader) error {
	sdqjd.state = &sagemakerDataQualityJobDefinitionState{}
	if err := json.NewDecoder(av).Decode(sdqjd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdqjd.Type(), sdqjd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerDataQualityJobDefinition] has state.
func (sdqjd *SagemakerDataQualityJobDefinition) State() (*sagemakerDataQualityJobDefinitionState, bool) {
	return sdqjd.state, sdqjd.state != nil
}

// StateMust returns the state for [SagemakerDataQualityJobDefinition]. Panics if the state is nil.
func (sdqjd *SagemakerDataQualityJobDefinition) StateMust() *sagemakerDataQualityJobDefinitionState {
	if sdqjd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdqjd.Type(), sdqjd.LocalName()))
	}
	return sdqjd.state
}

// SagemakerDataQualityJobDefinitionArgs contains the configurations for aws_sagemaker_data_quality_job_definition.
type SagemakerDataQualityJobDefinitionArgs struct {
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
	DataQualityAppSpecification *sagemakerdataqualityjobdefinition.DataQualityAppSpecification `hcl:"data_quality_app_specification,block" validate:"required"`
	// DataQualityBaselineConfig: optional
	DataQualityBaselineConfig *sagemakerdataqualityjobdefinition.DataQualityBaselineConfig `hcl:"data_quality_baseline_config,block"`
	// DataQualityJobInput: required
	DataQualityJobInput *sagemakerdataqualityjobdefinition.DataQualityJobInput `hcl:"data_quality_job_input,block" validate:"required"`
	// DataQualityJobOutputConfig: required
	DataQualityJobOutputConfig *sagemakerdataqualityjobdefinition.DataQualityJobOutputConfig `hcl:"data_quality_job_output_config,block" validate:"required"`
	// JobResources: required
	JobResources *sagemakerdataqualityjobdefinition.JobResources `hcl:"job_resources,block" validate:"required"`
	// NetworkConfig: optional
	NetworkConfig *sagemakerdataqualityjobdefinition.NetworkConfig `hcl:"network_config,block"`
	// StoppingCondition: optional
	StoppingCondition *sagemakerdataqualityjobdefinition.StoppingCondition `hcl:"stopping_condition,block"`
}
type sagemakerDataQualityJobDefinitionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_data_quality_job_definition.
func (sdqjd sagemakerDataQualityJobDefinitionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sdqjd.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sagemaker_data_quality_job_definition.
func (sdqjd sagemakerDataQualityJobDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdqjd.ref.Append("id"))
}

// Name returns a reference to field name of aws_sagemaker_data_quality_job_definition.
func (sdqjd sagemakerDataQualityJobDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdqjd.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_data_quality_job_definition.
func (sdqjd sagemakerDataQualityJobDefinitionAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sdqjd.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_sagemaker_data_quality_job_definition.
func (sdqjd sagemakerDataQualityJobDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdqjd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_data_quality_job_definition.
func (sdqjd sagemakerDataQualityJobDefinitionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdqjd.ref.Append("tags_all"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) DataQualityAppSpecification() terra.ListValue[sagemakerdataqualityjobdefinition.DataQualityAppSpecificationAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.DataQualityAppSpecificationAttributes](sdqjd.ref.Append("data_quality_app_specification"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) DataQualityBaselineConfig() terra.ListValue[sagemakerdataqualityjobdefinition.DataQualityBaselineConfigAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.DataQualityBaselineConfigAttributes](sdqjd.ref.Append("data_quality_baseline_config"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) DataQualityJobInput() terra.ListValue[sagemakerdataqualityjobdefinition.DataQualityJobInputAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.DataQualityJobInputAttributes](sdqjd.ref.Append("data_quality_job_input"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) DataQualityJobOutputConfig() terra.ListValue[sagemakerdataqualityjobdefinition.DataQualityJobOutputConfigAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.DataQualityJobOutputConfigAttributes](sdqjd.ref.Append("data_quality_job_output_config"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) JobResources() terra.ListValue[sagemakerdataqualityjobdefinition.JobResourcesAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.JobResourcesAttributes](sdqjd.ref.Append("job_resources"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) NetworkConfig() terra.ListValue[sagemakerdataqualityjobdefinition.NetworkConfigAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.NetworkConfigAttributes](sdqjd.ref.Append("network_config"))
}

func (sdqjd sagemakerDataQualityJobDefinitionAttributes) StoppingCondition() terra.ListValue[sagemakerdataqualityjobdefinition.StoppingConditionAttributes] {
	return terra.ReferenceAsList[sagemakerdataqualityjobdefinition.StoppingConditionAttributes](sdqjd.ref.Append("stopping_condition"))
}

type sagemakerDataQualityJobDefinitionState struct {
	Arn                         string                                                               `json:"arn"`
	Id                          string                                                               `json:"id"`
	Name                        string                                                               `json:"name"`
	RoleArn                     string                                                               `json:"role_arn"`
	Tags                        map[string]string                                                    `json:"tags"`
	TagsAll                     map[string]string                                                    `json:"tags_all"`
	DataQualityAppSpecification []sagemakerdataqualityjobdefinition.DataQualityAppSpecificationState `json:"data_quality_app_specification"`
	DataQualityBaselineConfig   []sagemakerdataqualityjobdefinition.DataQualityBaselineConfigState   `json:"data_quality_baseline_config"`
	DataQualityJobInput         []sagemakerdataqualityjobdefinition.DataQualityJobInputState         `json:"data_quality_job_input"`
	DataQualityJobOutputConfig  []sagemakerdataqualityjobdefinition.DataQualityJobOutputConfigState  `json:"data_quality_job_output_config"`
	JobResources                []sagemakerdataqualityjobdefinition.JobResourcesState                `json:"job_resources"`
	NetworkConfig               []sagemakerdataqualityjobdefinition.NetworkConfigState               `json:"network_config"`
	StoppingCondition           []sagemakerdataqualityjobdefinition.StoppingConditionState           `json:"stopping_condition"`
}
