// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerapp

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type ResourceSpecAttributes struct {
	ref terra.Reference
}

func (rs ResourceSpecAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs ResourceSpecAttributes) InternalWithRef(ref terra.Reference) ResourceSpecAttributes {
	return ResourceSpecAttributes{ref: ref}
}

func (rs ResourceSpecAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs ResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("instance_type"))
}

func (rs ResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("lifecycle_config_arn"))
}

func (rs ResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("sagemaker_image_arn"))
}

func (rs ResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("sagemaker_image_version_arn"))
}

type ResourceSpecState struct {
	InstanceType             string `json:"instance_type"`
	LifecycleConfigArn       string `json:"lifecycle_config_arn"`
	SagemakerImageArn        string `json:"sagemaker_image_arn"`
	SagemakerImageVersionArn string `json:"sagemaker_image_version_arn"`
}
