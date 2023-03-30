// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerspace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SpaceSettings struct {
	// JupyterServerAppSettings: optional
	JupyterServerAppSettings *JupyterServerAppSettings `hcl:"jupyter_server_app_settings,block"`
	// KernelGatewayAppSettings: optional
	KernelGatewayAppSettings *KernelGatewayAppSettings `hcl:"kernel_gateway_app_settings,block"`
}

type JupyterServerAppSettings struct {
	// LifecycleConfigArns: set of string, optional
	LifecycleConfigArns terra.SetValue[terra.StringValue] `hcl:"lifecycle_config_arns,attr"`
	// CodeRepository: min=0,max=10
	CodeRepository []CodeRepository `hcl:"code_repository,block" validate:"min=0,max=10"`
	// JupyterServerAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *JupyterServerAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type CodeRepository struct {
	// RepositoryUrl: string, required
	RepositoryUrl terra.StringValue `hcl:"repository_url,attr" validate:"required"`
}

type JupyterServerAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type KernelGatewayAppSettings struct {
	// LifecycleConfigArns: set of string, optional
	LifecycleConfigArns terra.SetValue[terra.StringValue] `hcl:"lifecycle_config_arns,attr"`
	// CustomImage: min=0,max=30
	CustomImage []CustomImage `hcl:"custom_image,block" validate:"min=0,max=30"`
	// KernelGatewayAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *KernelGatewayAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type CustomImage struct {
	// AppImageConfigName: string, required
	AppImageConfigName terra.StringValue `hcl:"app_image_config_name,attr" validate:"required"`
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// ImageVersionNumber: number, optional
	ImageVersionNumber terra.NumberValue `hcl:"image_version_number,attr"`
}

type KernelGatewayAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type SpaceSettingsAttributes struct {
	ref terra.Reference
}

func (ss SpaceSettingsAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss SpaceSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsAttributes {
	return SpaceSettingsAttributes{ref: ref}
}

func (ss SpaceSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss SpaceSettingsAttributes) JupyterServerAppSettings() terra.ListValue[JupyterServerAppSettingsAttributes] {
	return terra.ReferenceList[JupyterServerAppSettingsAttributes](ss.ref.Append("jupyter_server_app_settings"))
}

func (ss SpaceSettingsAttributes) KernelGatewayAppSettings() terra.ListValue[KernelGatewayAppSettingsAttributes] {
	return terra.ReferenceList[KernelGatewayAppSettingsAttributes](ss.ref.Append("kernel_gateway_app_settings"))
}

type JupyterServerAppSettingsAttributes struct {
	ref terra.Reference
}

func (jsas JupyterServerAppSettingsAttributes) InternalRef() terra.Reference {
	return jsas.ref
}

func (jsas JupyterServerAppSettingsAttributes) InternalWithRef(ref terra.Reference) JupyterServerAppSettingsAttributes {
	return JupyterServerAppSettingsAttributes{ref: ref}
}

func (jsas JupyterServerAppSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return jsas.ref.InternalTokens()
}

func (jsas JupyterServerAppSettingsAttributes) LifecycleConfigArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](jsas.ref.Append("lifecycle_config_arns"))
}

func (jsas JupyterServerAppSettingsAttributes) CodeRepository() terra.SetValue[CodeRepositoryAttributes] {
	return terra.ReferenceSet[CodeRepositoryAttributes](jsas.ref.Append("code_repository"))
}

func (jsas JupyterServerAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[JupyterServerAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceList[JupyterServerAppSettingsDefaultResourceSpecAttributes](jsas.ref.Append("default_resource_spec"))
}

type CodeRepositoryAttributes struct {
	ref terra.Reference
}

func (cr CodeRepositoryAttributes) InternalRef() terra.Reference {
	return cr.ref
}

func (cr CodeRepositoryAttributes) InternalWithRef(ref terra.Reference) CodeRepositoryAttributes {
	return CodeRepositoryAttributes{ref: ref}
}

func (cr CodeRepositoryAttributes) InternalTokens() hclwrite.Tokens {
	return cr.ref.InternalTokens()
}

func (cr CodeRepositoryAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceString(cr.ref.Append("repository_url"))
}

type JupyterServerAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) InternalRef() terra.Reference {
	return drs.ref
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) JupyterServerAppSettingsDefaultResourceSpecAttributes {
	return JupyterServerAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) InternalTokens() hclwrite.Tokens {
	return drs.ref.InternalTokens()
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("instance_type"))
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs JupyterServerAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_version_arn"))
}

type KernelGatewayAppSettingsAttributes struct {
	ref terra.Reference
}

func (kgas KernelGatewayAppSettingsAttributes) InternalRef() terra.Reference {
	return kgas.ref
}

func (kgas KernelGatewayAppSettingsAttributes) InternalWithRef(ref terra.Reference) KernelGatewayAppSettingsAttributes {
	return KernelGatewayAppSettingsAttributes{ref: ref}
}

func (kgas KernelGatewayAppSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return kgas.ref.InternalTokens()
}

func (kgas KernelGatewayAppSettingsAttributes) LifecycleConfigArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](kgas.ref.Append("lifecycle_config_arns"))
}

func (kgas KernelGatewayAppSettingsAttributes) CustomImage() terra.ListValue[CustomImageAttributes] {
	return terra.ReferenceList[CustomImageAttributes](kgas.ref.Append("custom_image"))
}

func (kgas KernelGatewayAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[KernelGatewayAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceList[KernelGatewayAppSettingsDefaultResourceSpecAttributes](kgas.ref.Append("default_resource_spec"))
}

type CustomImageAttributes struct {
	ref terra.Reference
}

func (ci CustomImageAttributes) InternalRef() terra.Reference {
	return ci.ref
}

func (ci CustomImageAttributes) InternalWithRef(ref terra.Reference) CustomImageAttributes {
	return CustomImageAttributes{ref: ref}
}

func (ci CustomImageAttributes) InternalTokens() hclwrite.Tokens {
	return ci.ref.InternalTokens()
}

func (ci CustomImageAttributes) AppImageConfigName() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("app_image_config_name"))
}

func (ci CustomImageAttributes) ImageName() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("image_name"))
}

func (ci CustomImageAttributes) ImageVersionNumber() terra.NumberValue {
	return terra.ReferenceNumber(ci.ref.Append("image_version_number"))
}

type KernelGatewayAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) InternalRef() terra.Reference {
	return drs.ref
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) KernelGatewayAppSettingsDefaultResourceSpecAttributes {
	return KernelGatewayAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) InternalTokens() hclwrite.Tokens {
	return drs.ref.InternalTokens()
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("instance_type"))
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs KernelGatewayAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_version_arn"))
}

type SpaceSettingsState struct {
	JupyterServerAppSettings []JupyterServerAppSettingsState `json:"jupyter_server_app_settings"`
	KernelGatewayAppSettings []KernelGatewayAppSettingsState `json:"kernel_gateway_app_settings"`
}

type JupyterServerAppSettingsState struct {
	LifecycleConfigArns []string                                           `json:"lifecycle_config_arns"`
	CodeRepository      []CodeRepositoryState                              `json:"code_repository"`
	DefaultResourceSpec []JupyterServerAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type CodeRepositoryState struct {
	RepositoryUrl string `json:"repository_url"`
}

type JupyterServerAppSettingsDefaultResourceSpecState struct {
	InstanceType             string `json:"instance_type"`
	LifecycleConfigArn       string `json:"lifecycle_config_arn"`
	SagemakerImageArn        string `json:"sagemaker_image_arn"`
	SagemakerImageVersionArn string `json:"sagemaker_image_version_arn"`
}

type KernelGatewayAppSettingsState struct {
	LifecycleConfigArns []string                                           `json:"lifecycle_config_arns"`
	CustomImage         []CustomImageState                                 `json:"custom_image"`
	DefaultResourceSpec []KernelGatewayAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type CustomImageState struct {
	AppImageConfigName string  `json:"app_image_config_name"`
	ImageName          string  `json:"image_name"`
	ImageVersionNumber float64 `json:"image_version_number"`
}

type KernelGatewayAppSettingsDefaultResourceSpecState struct {
	InstanceType             string `json:"instance_type"`
	LifecycleConfigArn       string `json:"lifecycle_config_arn"`
	SagemakerImageArn        string `json:"sagemaker_image_arn"`
	SagemakerImageVersionArn string `json:"sagemaker_image_version_arn"`
}
