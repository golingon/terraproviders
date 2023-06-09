// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	comprehendentityrecognizer "github.com/golingon/terraproviders/aws/5.7.0/comprehendentityrecognizer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComprehendEntityRecognizer creates a new instance of [ComprehendEntityRecognizer].
func NewComprehendEntityRecognizer(name string, args ComprehendEntityRecognizerArgs) *ComprehendEntityRecognizer {
	return &ComprehendEntityRecognizer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComprehendEntityRecognizer)(nil)

// ComprehendEntityRecognizer represents the Terraform resource aws_comprehend_entity_recognizer.
type ComprehendEntityRecognizer struct {
	Name      string
	Args      ComprehendEntityRecognizerArgs
	state     *comprehendEntityRecognizerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComprehendEntityRecognizer].
func (cer *ComprehendEntityRecognizer) Type() string {
	return "aws_comprehend_entity_recognizer"
}

// LocalName returns the local name for [ComprehendEntityRecognizer].
func (cer *ComprehendEntityRecognizer) LocalName() string {
	return cer.Name
}

// Configuration returns the configuration (args) for [ComprehendEntityRecognizer].
func (cer *ComprehendEntityRecognizer) Configuration() interface{} {
	return cer.Args
}

// DependOn is used for other resources to depend on [ComprehendEntityRecognizer].
func (cer *ComprehendEntityRecognizer) DependOn() terra.Reference {
	return terra.ReferenceResource(cer)
}

// Dependencies returns the list of resources [ComprehendEntityRecognizer] depends_on.
func (cer *ComprehendEntityRecognizer) Dependencies() terra.Dependencies {
	return cer.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComprehendEntityRecognizer].
func (cer *ComprehendEntityRecognizer) LifecycleManagement() *terra.Lifecycle {
	return cer.Lifecycle
}

// Attributes returns the attributes for [ComprehendEntityRecognizer].
func (cer *ComprehendEntityRecognizer) Attributes() comprehendEntityRecognizerAttributes {
	return comprehendEntityRecognizerAttributes{ref: terra.ReferenceResource(cer)}
}

// ImportState imports the given attribute values into [ComprehendEntityRecognizer]'s state.
func (cer *ComprehendEntityRecognizer) ImportState(av io.Reader) error {
	cer.state = &comprehendEntityRecognizerState{}
	if err := json.NewDecoder(av).Decode(cer.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cer.Type(), cer.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComprehendEntityRecognizer] has state.
func (cer *ComprehendEntityRecognizer) State() (*comprehendEntityRecognizerState, bool) {
	return cer.state, cer.state != nil
}

// StateMust returns the state for [ComprehendEntityRecognizer]. Panics if the state is nil.
func (cer *ComprehendEntityRecognizer) StateMust() *comprehendEntityRecognizerState {
	if cer.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cer.Type(), cer.LocalName()))
	}
	return cer.state
}

// ComprehendEntityRecognizerArgs contains the configurations for aws_comprehend_entity_recognizer.
type ComprehendEntityRecognizerArgs struct {
	// DataAccessRoleArn: string, required
	DataAccessRoleArn terra.StringValue `hcl:"data_access_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// ModelKmsKeyId: string, optional
	ModelKmsKeyId terra.StringValue `hcl:"model_kms_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VersionName: string, optional
	VersionName terra.StringValue `hcl:"version_name,attr"`
	// VersionNamePrefix: string, optional
	VersionNamePrefix terra.StringValue `hcl:"version_name_prefix,attr"`
	// VolumeKmsKeyId: string, optional
	VolumeKmsKeyId terra.StringValue `hcl:"volume_kms_key_id,attr"`
	// InputDataConfig: required
	InputDataConfig *comprehendentityrecognizer.InputDataConfig `hcl:"input_data_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *comprehendentityrecognizer.Timeouts `hcl:"timeouts,block"`
	// VpcConfig: optional
	VpcConfig *comprehendentityrecognizer.VpcConfig `hcl:"vpc_config,block"`
}
type comprehendEntityRecognizerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("arn"))
}

// DataAccessRoleArn returns a reference to field data_access_role_arn of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) DataAccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("data_access_role_arn"))
}

// Id returns a reference to field id of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("language_code"))
}

// ModelKmsKeyId returns a reference to field model_kms_key_id of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) ModelKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("model_kms_key_id"))
}

// Name returns a reference to field name of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cer.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cer.ref.Append("tags_all"))
}

// VersionName returns a reference to field version_name of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) VersionName() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("version_name"))
}

// VersionNamePrefix returns a reference to field version_name_prefix of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) VersionNamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("version_name_prefix"))
}

// VolumeKmsKeyId returns a reference to field volume_kms_key_id of aws_comprehend_entity_recognizer.
func (cer comprehendEntityRecognizerAttributes) VolumeKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cer.ref.Append("volume_kms_key_id"))
}

func (cer comprehendEntityRecognizerAttributes) InputDataConfig() terra.ListValue[comprehendentityrecognizer.InputDataConfigAttributes] {
	return terra.ReferenceAsList[comprehendentityrecognizer.InputDataConfigAttributes](cer.ref.Append("input_data_config"))
}

func (cer comprehendEntityRecognizerAttributes) Timeouts() comprehendentityrecognizer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[comprehendentityrecognizer.TimeoutsAttributes](cer.ref.Append("timeouts"))
}

func (cer comprehendEntityRecognizerAttributes) VpcConfig() terra.ListValue[comprehendentityrecognizer.VpcConfigAttributes] {
	return terra.ReferenceAsList[comprehendentityrecognizer.VpcConfigAttributes](cer.ref.Append("vpc_config"))
}

type comprehendEntityRecognizerState struct {
	Arn               string                                            `json:"arn"`
	DataAccessRoleArn string                                            `json:"data_access_role_arn"`
	Id                string                                            `json:"id"`
	LanguageCode      string                                            `json:"language_code"`
	ModelKmsKeyId     string                                            `json:"model_kms_key_id"`
	Name              string                                            `json:"name"`
	Tags              map[string]string                                 `json:"tags"`
	TagsAll           map[string]string                                 `json:"tags_all"`
	VersionName       string                                            `json:"version_name"`
	VersionNamePrefix string                                            `json:"version_name_prefix"`
	VolumeKmsKeyId    string                                            `json:"volume_kms_key_id"`
	InputDataConfig   []comprehendentityrecognizer.InputDataConfigState `json:"input_data_config"`
	Timeouts          *comprehendentityrecognizer.TimeoutsState         `json:"timeouts"`
	VpcConfig         []comprehendentityrecognizer.VpcConfigState       `json:"vpc_config"`
}
