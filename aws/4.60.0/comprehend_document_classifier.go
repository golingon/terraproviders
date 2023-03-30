// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	comprehenddocumentclassifier "github.com/golingon/terraproviders/aws/4.60.0/comprehenddocumentclassifier"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComprehendDocumentClassifier(name string, args ComprehendDocumentClassifierArgs) *ComprehendDocumentClassifier {
	return &ComprehendDocumentClassifier{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComprehendDocumentClassifier)(nil)

type ComprehendDocumentClassifier struct {
	Name  string
	Args  ComprehendDocumentClassifierArgs
	state *comprehendDocumentClassifierState
}

func (cdc *ComprehendDocumentClassifier) Type() string {
	return "aws_comprehend_document_classifier"
}

func (cdc *ComprehendDocumentClassifier) LocalName() string {
	return cdc.Name
}

func (cdc *ComprehendDocumentClassifier) Configuration() interface{} {
	return cdc.Args
}

func (cdc *ComprehendDocumentClassifier) Attributes() comprehendDocumentClassifierAttributes {
	return comprehendDocumentClassifierAttributes{ref: terra.ReferenceResource(cdc)}
}

func (cdc *ComprehendDocumentClassifier) ImportState(av io.Reader) error {
	cdc.state = &comprehendDocumentClassifierState{}
	if err := json.NewDecoder(av).Decode(cdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdc.Type(), cdc.LocalName(), err)
	}
	return nil
}

func (cdc *ComprehendDocumentClassifier) State() (*comprehendDocumentClassifierState, bool) {
	return cdc.state, cdc.state != nil
}

func (cdc *ComprehendDocumentClassifier) StateMust() *comprehendDocumentClassifierState {
	if cdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdc.Type(), cdc.LocalName()))
	}
	return cdc.state
}

func (cdc *ComprehendDocumentClassifier) DependOn() terra.Reference {
	return terra.ReferenceResource(cdc)
}

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
	// DependsOn contains resources that ComprehendDocumentClassifier depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type comprehendDocumentClassifierAttributes struct {
	ref terra.Reference
}

func (cdc comprehendDocumentClassifierAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("arn"))
}

func (cdc comprehendDocumentClassifierAttributes) DataAccessRoleArn() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("data_access_role_arn"))
}

func (cdc comprehendDocumentClassifierAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("id"))
}

func (cdc comprehendDocumentClassifierAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("language_code"))
}

func (cdc comprehendDocumentClassifierAttributes) Mode() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("mode"))
}

func (cdc comprehendDocumentClassifierAttributes) ModelKmsKeyId() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("model_kms_key_id"))
}

func (cdc comprehendDocumentClassifierAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("name"))
}

func (cdc comprehendDocumentClassifierAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cdc.ref.Append("tags"))
}

func (cdc comprehendDocumentClassifierAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cdc.ref.Append("tags_all"))
}

func (cdc comprehendDocumentClassifierAttributes) VersionName() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("version_name"))
}

func (cdc comprehendDocumentClassifierAttributes) VersionNamePrefix() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("version_name_prefix"))
}

func (cdc comprehendDocumentClassifierAttributes) VolumeKmsKeyId() terra.StringValue {
	return terra.ReferenceString(cdc.ref.Append("volume_kms_key_id"))
}

func (cdc comprehendDocumentClassifierAttributes) InputDataConfig() terra.ListValue[comprehenddocumentclassifier.InputDataConfigAttributes] {
	return terra.ReferenceList[comprehenddocumentclassifier.InputDataConfigAttributes](cdc.ref.Append("input_data_config"))
}

func (cdc comprehendDocumentClassifierAttributes) OutputDataConfig() terra.ListValue[comprehenddocumentclassifier.OutputDataConfigAttributes] {
	return terra.ReferenceList[comprehenddocumentclassifier.OutputDataConfigAttributes](cdc.ref.Append("output_data_config"))
}

func (cdc comprehendDocumentClassifierAttributes) Timeouts() comprehenddocumentclassifier.TimeoutsAttributes {
	return terra.ReferenceSingle[comprehenddocumentclassifier.TimeoutsAttributes](cdc.ref.Append("timeouts"))
}

func (cdc comprehendDocumentClassifierAttributes) VpcConfig() terra.ListValue[comprehenddocumentclassifier.VpcConfigAttributes] {
	return terra.ReferenceList[comprehenddocumentclassifier.VpcConfigAttributes](cdc.ref.Append("vpc_config"))
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
