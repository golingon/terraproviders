// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	comprehenddocumentclassifier "github.com/golingon/terraproviders/aws/5.12.0/comprehenddocumentclassifier"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComprehendDocumentClassifier creates a new instance of [ComprehendDocumentClassifier].
func NewComprehendDocumentClassifier(name string, args ComprehendDocumentClassifierArgs) *ComprehendDocumentClassifier {
	return &ComprehendDocumentClassifier{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComprehendDocumentClassifier)(nil)

// ComprehendDocumentClassifier represents the Terraform resource aws_comprehend_document_classifier.
type ComprehendDocumentClassifier struct {
	Name      string
	Args      ComprehendDocumentClassifierArgs
	state     *comprehendDocumentClassifierState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComprehendDocumentClassifier].
func (cdc *ComprehendDocumentClassifier) Type() string {
	return "aws_comprehend_document_classifier"
}

// LocalName returns the local name for [ComprehendDocumentClassifier].
func (cdc *ComprehendDocumentClassifier) LocalName() string {
	return cdc.Name
}

// Configuration returns the configuration (args) for [ComprehendDocumentClassifier].
func (cdc *ComprehendDocumentClassifier) Configuration() interface{} {
	return cdc.Args
}

// DependOn is used for other resources to depend on [ComprehendDocumentClassifier].
func (cdc *ComprehendDocumentClassifier) DependOn() terra.Reference {
	return terra.ReferenceResource(cdc)
}

// Dependencies returns the list of resources [ComprehendDocumentClassifier] depends_on.
func (cdc *ComprehendDocumentClassifier) Dependencies() terra.Dependencies {
	return cdc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComprehendDocumentClassifier].
func (cdc *ComprehendDocumentClassifier) LifecycleManagement() *terra.Lifecycle {
	return cdc.Lifecycle
}

// Attributes returns the attributes for [ComprehendDocumentClassifier].
func (cdc *ComprehendDocumentClassifier) Attributes() comprehendDocumentClassifierAttributes {
	return comprehendDocumentClassifierAttributes{ref: terra.ReferenceResource(cdc)}
}

// ImportState imports the given attribute values into [ComprehendDocumentClassifier]'s state.
func (cdc *ComprehendDocumentClassifier) ImportState(av io.Reader) error {
	cdc.state = &comprehendDocumentClassifierState{}
	if err := json.NewDecoder(av).Decode(cdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdc.Type(), cdc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComprehendDocumentClassifier] has state.
func (cdc *ComprehendDocumentClassifier) State() (*comprehendDocumentClassifierState, bool) {
	return cdc.state, cdc.state != nil
}

// StateMust returns the state for [ComprehendDocumentClassifier]. Panics if the state is nil.
func (cdc *ComprehendDocumentClassifier) StateMust() *comprehendDocumentClassifierState {
	if cdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdc.Type(), cdc.LocalName()))
	}
	return cdc.state
}

// ComprehendDocumentClassifierArgs contains the configurations for aws_comprehend_document_classifier.
type ComprehendDocumentClassifierArgs struct {
	// DataAccessRoleArn: string, required
	DataAccessRoleArn terra.StringValue `hcl:"data_access_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
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
	InputDataConfig *comprehenddocumentclassifier.InputDataConfig `hcl:"input_data_config,block" validate:"required"`
	// OutputDataConfig: optional
	OutputDataConfig *comprehenddocumentclassifier.OutputDataConfig `hcl:"output_data_config,block"`
	// Timeouts: optional
	Timeouts *comprehenddocumentclassifier.Timeouts `hcl:"timeouts,block"`
	// VpcConfig: optional
	VpcConfig *comprehenddocumentclassifier.VpcConfig `hcl:"vpc_config,block"`
}
type comprehendDocumentClassifierAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("arn"))
}

// DataAccessRoleArn returns a reference to field data_access_role_arn of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) DataAccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("data_access_role_arn"))
}

// Id returns a reference to field id of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("id"))
}

// LanguageCode returns a reference to field language_code of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("language_code"))
}

// Mode returns a reference to field mode of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("mode"))
}

// ModelKmsKeyId returns a reference to field model_kms_key_id of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) ModelKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("model_kms_key_id"))
}

// Name returns a reference to field name of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdc.ref.Append("tags_all"))
}

// VersionName returns a reference to field version_name of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) VersionName() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("version_name"))
}

// VersionNamePrefix returns a reference to field version_name_prefix of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) VersionNamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("version_name_prefix"))
}

// VolumeKmsKeyId returns a reference to field volume_kms_key_id of aws_comprehend_document_classifier.
func (cdc comprehendDocumentClassifierAttributes) VolumeKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("volume_kms_key_id"))
}

func (cdc comprehendDocumentClassifierAttributes) InputDataConfig() terra.ListValue[comprehenddocumentclassifier.InputDataConfigAttributes] {
	return terra.ReferenceAsList[comprehenddocumentclassifier.InputDataConfigAttributes](cdc.ref.Append("input_data_config"))
}

func (cdc comprehendDocumentClassifierAttributes) OutputDataConfig() terra.ListValue[comprehenddocumentclassifier.OutputDataConfigAttributes] {
	return terra.ReferenceAsList[comprehenddocumentclassifier.OutputDataConfigAttributes](cdc.ref.Append("output_data_config"))
}

func (cdc comprehendDocumentClassifierAttributes) Timeouts() comprehenddocumentclassifier.TimeoutsAttributes {
	return terra.ReferenceAsSingle[comprehenddocumentclassifier.TimeoutsAttributes](cdc.ref.Append("timeouts"))
}

func (cdc comprehendDocumentClassifierAttributes) VpcConfig() terra.ListValue[comprehenddocumentclassifier.VpcConfigAttributes] {
	return terra.ReferenceAsList[comprehenddocumentclassifier.VpcConfigAttributes](cdc.ref.Append("vpc_config"))
}

type comprehendDocumentClassifierState struct {
	Arn               string                                               `json:"arn"`
	DataAccessRoleArn string                                               `json:"data_access_role_arn"`
	Id                string                                               `json:"id"`
	LanguageCode      string                                               `json:"language_code"`
	Mode              string                                               `json:"mode"`
	ModelKmsKeyId     string                                               `json:"model_kms_key_id"`
	Name              string                                               `json:"name"`
	Tags              map[string]string                                    `json:"tags"`
	TagsAll           map[string]string                                    `json:"tags_all"`
	VersionName       string                                               `json:"version_name"`
	VersionNamePrefix string                                               `json:"version_name_prefix"`
	VolumeKmsKeyId    string                                               `json:"volume_kms_key_id"`
	InputDataConfig   []comprehenddocumentclassifier.InputDataConfigState  `json:"input_data_config"`
	OutputDataConfig  []comprehenddocumentclassifier.OutputDataConfigState `json:"output_data_config"`
	Timeouts          *comprehenddocumentclassifier.TimeoutsState          `json:"timeouts"`
	VpcConfig         []comprehenddocumentclassifier.VpcConfigState        `json:"vpc_config"`
}