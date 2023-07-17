// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLambdaLayerVersion creates a new instance of [LambdaLayerVersion].
func NewLambdaLayerVersion(name string, args LambdaLayerVersionArgs) *LambdaLayerVersion {
	return &LambdaLayerVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaLayerVersion)(nil)

// LambdaLayerVersion represents the Terraform resource aws_lambda_layer_version.
type LambdaLayerVersion struct {
	Name      string
	Args      LambdaLayerVersionArgs
	state     *lambdaLayerVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LambdaLayerVersion].
func (llv *LambdaLayerVersion) Type() string {
	return "aws_lambda_layer_version"
}

// LocalName returns the local name for [LambdaLayerVersion].
func (llv *LambdaLayerVersion) LocalName() string {
	return llv.Name
}

// Configuration returns the configuration (args) for [LambdaLayerVersion].
func (llv *LambdaLayerVersion) Configuration() interface{} {
	return llv.Args
}

// DependOn is used for other resources to depend on [LambdaLayerVersion].
func (llv *LambdaLayerVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(llv)
}

// Dependencies returns the list of resources [LambdaLayerVersion] depends_on.
func (llv *LambdaLayerVersion) Dependencies() terra.Dependencies {
	return llv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LambdaLayerVersion].
func (llv *LambdaLayerVersion) LifecycleManagement() *terra.Lifecycle {
	return llv.Lifecycle
}

// Attributes returns the attributes for [LambdaLayerVersion].
func (llv *LambdaLayerVersion) Attributes() lambdaLayerVersionAttributes {
	return lambdaLayerVersionAttributes{ref: terra.ReferenceResource(llv)}
}

// ImportState imports the given attribute values into [LambdaLayerVersion]'s state.
func (llv *LambdaLayerVersion) ImportState(av io.Reader) error {
	llv.state = &lambdaLayerVersionState{}
	if err := json.NewDecoder(av).Decode(llv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llv.Type(), llv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LambdaLayerVersion] has state.
func (llv *LambdaLayerVersion) State() (*lambdaLayerVersionState, bool) {
	return llv.state, llv.state != nil
}

// StateMust returns the state for [LambdaLayerVersion]. Panics if the state is nil.
func (llv *LambdaLayerVersion) StateMust() *lambdaLayerVersionState {
	if llv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llv.Type(), llv.LocalName()))
	}
	return llv.state
}

// LambdaLayerVersionArgs contains the configurations for aws_lambda_layer_version.
type LambdaLayerVersionArgs struct {
	// CompatibleArchitectures: set of string, optional
	CompatibleArchitectures terra.SetValue[terra.StringValue] `hcl:"compatible_architectures,attr"`
	// CompatibleRuntimes: set of string, optional
	CompatibleRuntimes terra.SetValue[terra.StringValue] `hcl:"compatible_runtimes,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Filename: string, optional
	Filename terra.StringValue `hcl:"filename,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LayerName: string, required
	LayerName terra.StringValue `hcl:"layer_name,attr" validate:"required"`
	// LicenseInfo: string, optional
	LicenseInfo terra.StringValue `hcl:"license_info,attr"`
	// S3Bucket: string, optional
	S3Bucket terra.StringValue `hcl:"s3_bucket,attr"`
	// S3Key: string, optional
	S3Key terra.StringValue `hcl:"s3_key,attr"`
	// S3ObjectVersion: string, optional
	S3ObjectVersion terra.StringValue `hcl:"s3_object_version,attr"`
	// SkipDestroy: bool, optional
	SkipDestroy terra.BoolValue `hcl:"skip_destroy,attr"`
	// SourceCodeHash: string, optional
	SourceCodeHash terra.StringValue `hcl:"source_code_hash,attr"`
}
type lambdaLayerVersionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("arn"))
}

// CompatibleArchitectures returns a reference to field compatible_architectures of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) CompatibleArchitectures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](llv.ref.Append("compatible_architectures"))
}

// CompatibleRuntimes returns a reference to field compatible_runtimes of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) CompatibleRuntimes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](llv.ref.Append("compatible_runtimes"))
}

// CreatedDate returns a reference to field created_date of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("description"))
}

// Filename returns a reference to field filename of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) Filename() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("filename"))
}

// Id returns a reference to field id of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("id"))
}

// LayerArn returns a reference to field layer_arn of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) LayerArn() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("layer_arn"))
}

// LayerName returns a reference to field layer_name of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) LayerName() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("layer_name"))
}

// LicenseInfo returns a reference to field license_info of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) LicenseInfo() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("license_info"))
}

// S3Bucket returns a reference to field s3_bucket of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) S3Bucket() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("s3_bucket"))
}

// S3Key returns a reference to field s3_key of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) S3Key() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("s3_key"))
}

// S3ObjectVersion returns a reference to field s3_object_version of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) S3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("s3_object_version"))
}

// SigningJobArn returns a reference to field signing_job_arn of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) SigningJobArn() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("signing_job_arn"))
}

// SigningProfileVersionArn returns a reference to field signing_profile_version_arn of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) SigningProfileVersionArn() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("signing_profile_version_arn"))
}

// SkipDestroy returns a reference to field skip_destroy of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) SkipDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(llv.ref.Append("skip_destroy"))
}

// SourceCodeHash returns a reference to field source_code_hash of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) SourceCodeHash() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("source_code_hash"))
}

// SourceCodeSize returns a reference to field source_code_size of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) SourceCodeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(llv.ref.Append("source_code_size"))
}

// Version returns a reference to field version of aws_lambda_layer_version.
func (llv lambdaLayerVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("version"))
}

type lambdaLayerVersionState struct {
	Arn                      string   `json:"arn"`
	CompatibleArchitectures  []string `json:"compatible_architectures"`
	CompatibleRuntimes       []string `json:"compatible_runtimes"`
	CreatedDate              string   `json:"created_date"`
	Description              string   `json:"description"`
	Filename                 string   `json:"filename"`
	Id                       string   `json:"id"`
	LayerArn                 string   `json:"layer_arn"`
	LayerName                string   `json:"layer_name"`
	LicenseInfo              string   `json:"license_info"`
	S3Bucket                 string   `json:"s3_bucket"`
	S3Key                    string   `json:"s3_key"`
	S3ObjectVersion          string   `json:"s3_object_version"`
	SigningJobArn            string   `json:"signing_job_arn"`
	SigningProfileVersionArn string   `json:"signing_profile_version_arn"`
	SkipDestroy              bool     `json:"skip_destroy"`
	SourceCodeHash           string   `json:"source_code_hash"`
	SourceCodeSize           float64  `json:"source_code_size"`
	Version                  string   `json:"version"`
}