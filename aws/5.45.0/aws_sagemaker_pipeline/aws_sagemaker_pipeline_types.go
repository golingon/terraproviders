// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sagemaker_pipeline

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ParallelismConfiguration struct {
	// MaxParallelExecutionSteps: number, required
	MaxParallelExecutionSteps terra.NumberValue `hcl:"max_parallel_execution_steps,attr" validate:"required"`
}

type PipelineDefinitionS3Location struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ObjectKey: string, required
	ObjectKey terra.StringValue `hcl:"object_key,attr" validate:"required"`
	// VersionId: string, optional
	VersionId terra.StringValue `hcl:"version_id,attr"`
}

type ParallelismConfigurationAttributes struct {
	ref terra.Reference
}

func (pc ParallelismConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ParallelismConfigurationAttributes) InternalWithRef(ref terra.Reference) ParallelismConfigurationAttributes {
	return ParallelismConfigurationAttributes{ref: ref}
}

func (pc ParallelismConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ParallelismConfigurationAttributes) MaxParallelExecutionSteps() terra.NumberValue {
	return terra.ReferenceAsNumber(pc.ref.Append("max_parallel_execution_steps"))
}

type PipelineDefinitionS3LocationAttributes struct {
	ref terra.Reference
}

func (pdsl PipelineDefinitionS3LocationAttributes) InternalRef() (terra.Reference, error) {
	return pdsl.ref, nil
}

func (pdsl PipelineDefinitionS3LocationAttributes) InternalWithRef(ref terra.Reference) PipelineDefinitionS3LocationAttributes {
	return PipelineDefinitionS3LocationAttributes{ref: ref}
}

func (pdsl PipelineDefinitionS3LocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pdsl.ref.InternalTokens()
}

func (pdsl PipelineDefinitionS3LocationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(pdsl.ref.Append("bucket"))
}

func (pdsl PipelineDefinitionS3LocationAttributes) ObjectKey() terra.StringValue {
	return terra.ReferenceAsString(pdsl.ref.Append("object_key"))
}

func (pdsl PipelineDefinitionS3LocationAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(pdsl.ref.Append("version_id"))
}

type ParallelismConfigurationState struct {
	MaxParallelExecutionSteps float64 `json:"max_parallel_execution_steps"`
}

type PipelineDefinitionS3LocationState struct {
	Bucket    string `json:"bucket"`
	ObjectKey string `json:"object_key"`
	VersionId string `json:"version_id"`
}
