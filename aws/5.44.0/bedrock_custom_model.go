// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	bedrockcustommodel "github.com/golingon/terraproviders/aws/5.44.0/bedrockcustommodel"
	"io"
)

// NewBedrockCustomModel creates a new instance of [BedrockCustomModel].
func NewBedrockCustomModel(name string, args BedrockCustomModelArgs) *BedrockCustomModel {
	return &BedrockCustomModel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BedrockCustomModel)(nil)

// BedrockCustomModel represents the Terraform resource aws_bedrock_custom_model.
type BedrockCustomModel struct {
	Name      string
	Args      BedrockCustomModelArgs
	state     *bedrockCustomModelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BedrockCustomModel].
func (bcm *BedrockCustomModel) Type() string {
	return "aws_bedrock_custom_model"
}

// LocalName returns the local name for [BedrockCustomModel].
func (bcm *BedrockCustomModel) LocalName() string {
	return bcm.Name
}

// Configuration returns the configuration (args) for [BedrockCustomModel].
func (bcm *BedrockCustomModel) Configuration() interface{} {
	return bcm.Args
}

// DependOn is used for other resources to depend on [BedrockCustomModel].
func (bcm *BedrockCustomModel) DependOn() terra.Reference {
	return terra.ReferenceResource(bcm)
}

// Dependencies returns the list of resources [BedrockCustomModel] depends_on.
func (bcm *BedrockCustomModel) Dependencies() terra.Dependencies {
	return bcm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BedrockCustomModel].
func (bcm *BedrockCustomModel) LifecycleManagement() *terra.Lifecycle {
	return bcm.Lifecycle
}

// Attributes returns the attributes for [BedrockCustomModel].
func (bcm *BedrockCustomModel) Attributes() bedrockCustomModelAttributes {
	return bedrockCustomModelAttributes{ref: terra.ReferenceResource(bcm)}
}

// ImportState imports the given attribute values into [BedrockCustomModel]'s state.
func (bcm *BedrockCustomModel) ImportState(av io.Reader) error {
	bcm.state = &bedrockCustomModelState{}
	if err := json.NewDecoder(av).Decode(bcm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcm.Type(), bcm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BedrockCustomModel] has state.
func (bcm *BedrockCustomModel) State() (*bedrockCustomModelState, bool) {
	return bcm.state, bcm.state != nil
}

// StateMust returns the state for [BedrockCustomModel]. Panics if the state is nil.
func (bcm *BedrockCustomModel) StateMust() *bedrockCustomModelState {
	if bcm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcm.Type(), bcm.LocalName()))
	}
	return bcm.state
}

// BedrockCustomModelArgs contains the configurations for aws_bedrock_custom_model.
type BedrockCustomModelArgs struct {
	// BaseModelIdentifier: string, required
	BaseModelIdentifier terra.StringValue `hcl:"base_model_identifier,attr" validate:"required"`
	// CustomModelKmsKeyId: string, optional
	CustomModelKmsKeyId terra.StringValue `hcl:"custom_model_kms_key_id,attr"`
	// CustomModelName: string, required
	CustomModelName terra.StringValue `hcl:"custom_model_name,attr" validate:"required"`
	// CustomizationType: string, optional
	CustomizationType terra.StringValue `hcl:"customization_type,attr"`
	// Hyperparameters: map of string, required
	Hyperparameters terra.MapValue[terra.StringValue] `hcl:"hyperparameters,attr" validate:"required"`
	// JobName: string, required
	JobName terra.StringValue `hcl:"job_name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TrainingMetrics: min=0
	TrainingMetrics []bedrockcustommodel.TrainingMetrics `hcl:"training_metrics,block" validate:"min=0"`
	// ValidationMetrics: min=0
	ValidationMetrics []bedrockcustommodel.ValidationMetrics `hcl:"validation_metrics,block" validate:"min=0"`
	// OutputDataConfig: min=0
	OutputDataConfig []bedrockcustommodel.OutputDataConfig `hcl:"output_data_config,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *bedrockcustommodel.Timeouts `hcl:"timeouts,block"`
	// TrainingDataConfig: min=0
	TrainingDataConfig []bedrockcustommodel.TrainingDataConfig `hcl:"training_data_config,block" validate:"min=0"`
	// ValidationDataConfig: min=0
	ValidationDataConfig []bedrockcustommodel.ValidationDataConfig `hcl:"validation_data_config,block" validate:"min=0"`
	// VpcConfig: min=0
	VpcConfig []bedrockcustommodel.VpcConfig `hcl:"vpc_config,block" validate:"min=0"`
}
type bedrockCustomModelAttributes struct {
	ref terra.Reference
}

// BaseModelIdentifier returns a reference to field base_model_identifier of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) BaseModelIdentifier() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("base_model_identifier"))
}

// CustomModelArn returns a reference to field custom_model_arn of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) CustomModelArn() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("custom_model_arn"))
}

// CustomModelKmsKeyId returns a reference to field custom_model_kms_key_id of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) CustomModelKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("custom_model_kms_key_id"))
}

// CustomModelName returns a reference to field custom_model_name of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) CustomModelName() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("custom_model_name"))
}

// CustomizationType returns a reference to field customization_type of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) CustomizationType() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("customization_type"))
}

// Hyperparameters returns a reference to field hyperparameters of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) Hyperparameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bcm.ref.Append("hyperparameters"))
}

// Id returns a reference to field id of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("id"))
}

// JobArn returns a reference to field job_arn of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) JobArn() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("job_arn"))
}

// JobName returns a reference to field job_name of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) JobName() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("job_name"))
}

// JobStatus returns a reference to field job_status of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) JobStatus() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("job_status"))
}

// RoleArn returns a reference to field role_arn of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(bcm.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bcm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_bedrock_custom_model.
func (bcm bedrockCustomModelAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bcm.ref.Append("tags_all"))
}

func (bcm bedrockCustomModelAttributes) TrainingMetrics() terra.ListValue[bedrockcustommodel.TrainingMetricsAttributes] {
	return terra.ReferenceAsList[bedrockcustommodel.TrainingMetricsAttributes](bcm.ref.Append("training_metrics"))
}

func (bcm bedrockCustomModelAttributes) ValidationMetrics() terra.ListValue[bedrockcustommodel.ValidationMetricsAttributes] {
	return terra.ReferenceAsList[bedrockcustommodel.ValidationMetricsAttributes](bcm.ref.Append("validation_metrics"))
}

func (bcm bedrockCustomModelAttributes) OutputDataConfig() terra.ListValue[bedrockcustommodel.OutputDataConfigAttributes] {
	return terra.ReferenceAsList[bedrockcustommodel.OutputDataConfigAttributes](bcm.ref.Append("output_data_config"))
}

func (bcm bedrockCustomModelAttributes) Timeouts() bedrockcustommodel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bedrockcustommodel.TimeoutsAttributes](bcm.ref.Append("timeouts"))
}

func (bcm bedrockCustomModelAttributes) TrainingDataConfig() terra.ListValue[bedrockcustommodel.TrainingDataConfigAttributes] {
	return terra.ReferenceAsList[bedrockcustommodel.TrainingDataConfigAttributes](bcm.ref.Append("training_data_config"))
}

func (bcm bedrockCustomModelAttributes) ValidationDataConfig() terra.ListValue[bedrockcustommodel.ValidationDataConfigAttributes] {
	return terra.ReferenceAsList[bedrockcustommodel.ValidationDataConfigAttributes](bcm.ref.Append("validation_data_config"))
}

func (bcm bedrockCustomModelAttributes) VpcConfig() terra.ListValue[bedrockcustommodel.VpcConfigAttributes] {
	return terra.ReferenceAsList[bedrockcustommodel.VpcConfigAttributes](bcm.ref.Append("vpc_config"))
}

type bedrockCustomModelState struct {
	BaseModelIdentifier  string                                         `json:"base_model_identifier"`
	CustomModelArn       string                                         `json:"custom_model_arn"`
	CustomModelKmsKeyId  string                                         `json:"custom_model_kms_key_id"`
	CustomModelName      string                                         `json:"custom_model_name"`
	CustomizationType    string                                         `json:"customization_type"`
	Hyperparameters      map[string]string                              `json:"hyperparameters"`
	Id                   string                                         `json:"id"`
	JobArn               string                                         `json:"job_arn"`
	JobName              string                                         `json:"job_name"`
	JobStatus            string                                         `json:"job_status"`
	RoleArn              string                                         `json:"role_arn"`
	Tags                 map[string]string                              `json:"tags"`
	TagsAll              map[string]string                              `json:"tags_all"`
	TrainingMetrics      []bedrockcustommodel.TrainingMetricsState      `json:"training_metrics"`
	ValidationMetrics    []bedrockcustommodel.ValidationMetricsState    `json:"validation_metrics"`
	OutputDataConfig     []bedrockcustommodel.OutputDataConfigState     `json:"output_data_config"`
	Timeouts             *bedrockcustommodel.TimeoutsState              `json:"timeouts"`
	TrainingDataConfig   []bedrockcustommodel.TrainingDataConfigState   `json:"training_data_config"`
	ValidationDataConfig []bedrockcustommodel.ValidationDataConfigState `json:"validation_data_config"`
	VpcConfig            []bedrockcustommodel.VpcConfigState            `json:"vpc_config"`
}
