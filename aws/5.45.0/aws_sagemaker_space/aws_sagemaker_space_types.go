// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sagemaker_space

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type OwnershipSettings struct {
	// OwnerUserProfileName: string, required
	OwnerUserProfileName terra.StringValue `hcl:"owner_user_profile_name,attr" validate:"required"`
}

type SpaceSettings struct {
	// AppType: string, optional
	AppType terra.StringValue `hcl:"app_type,attr"`
	// SpaceSettingsCodeEditorAppSettings: optional
	CodeEditorAppSettings *SpaceSettingsCodeEditorAppSettings `hcl:"code_editor_app_settings,block"`
	// SpaceSettingsCustomFileSystem: min=0
	CustomFileSystem []SpaceSettingsCustomFileSystem `hcl:"custom_file_system,block" validate:"min=0"`
	// SpaceSettingsJupyterLabAppSettings: optional
	JupyterLabAppSettings *SpaceSettingsJupyterLabAppSettings `hcl:"jupyter_lab_app_settings,block"`
	// SpaceSettingsJupyterServerAppSettings: optional
	JupyterServerAppSettings *SpaceSettingsJupyterServerAppSettings `hcl:"jupyter_server_app_settings,block"`
	// SpaceSettingsKernelGatewayAppSettings: optional
	KernelGatewayAppSettings *SpaceSettingsKernelGatewayAppSettings `hcl:"kernel_gateway_app_settings,block"`
	// SpaceSettingsSpaceStorageSettings: optional
	SpaceStorageSettings *SpaceSettingsSpaceStorageSettings `hcl:"space_storage_settings,block"`
}

type SpaceSettingsCodeEditorAppSettings struct {
	// SpaceSettingsCodeEditorAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *SpaceSettingsCodeEditorAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type SpaceSettingsCodeEditorAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionAlias: string, optional
	SagemakerImageVersionAlias terra.StringValue `hcl:"sagemaker_image_version_alias,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type SpaceSettingsCustomFileSystem struct {
	// SpaceSettingsCustomFileSystemEfsFileSystem: required
	EfsFileSystem *SpaceSettingsCustomFileSystemEfsFileSystem `hcl:"efs_file_system,block" validate:"required"`
}

type SpaceSettingsCustomFileSystemEfsFileSystem struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
}

type SpaceSettingsJupyterLabAppSettings struct {
	// SpaceSettingsJupyterLabAppSettingsCodeRepository: min=0,max=10
	CodeRepository []SpaceSettingsJupyterLabAppSettingsCodeRepository `hcl:"code_repository,block" validate:"min=0,max=10"`
	// SpaceSettingsJupyterLabAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *SpaceSettingsJupyterLabAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type SpaceSettingsJupyterLabAppSettingsCodeRepository struct {
	// RepositoryUrl: string, required
	RepositoryUrl terra.StringValue `hcl:"repository_url,attr" validate:"required"`
}

type SpaceSettingsJupyterLabAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionAlias: string, optional
	SagemakerImageVersionAlias terra.StringValue `hcl:"sagemaker_image_version_alias,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type SpaceSettingsJupyterServerAppSettings struct {
	// LifecycleConfigArns: set of string, optional
	LifecycleConfigArns terra.SetValue[terra.StringValue] `hcl:"lifecycle_config_arns,attr"`
	// SpaceSettingsJupyterServerAppSettingsCodeRepository: min=0,max=10
	CodeRepository []SpaceSettingsJupyterServerAppSettingsCodeRepository `hcl:"code_repository,block" validate:"min=0,max=10"`
	// SpaceSettingsJupyterServerAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *SpaceSettingsJupyterServerAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type SpaceSettingsJupyterServerAppSettingsCodeRepository struct {
	// RepositoryUrl: string, required
	RepositoryUrl terra.StringValue `hcl:"repository_url,attr" validate:"required"`
}

type SpaceSettingsJupyterServerAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionAlias: string, optional
	SagemakerImageVersionAlias terra.StringValue `hcl:"sagemaker_image_version_alias,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type SpaceSettingsKernelGatewayAppSettings struct {
	// LifecycleConfigArns: set of string, optional
	LifecycleConfigArns terra.SetValue[terra.StringValue] `hcl:"lifecycle_config_arns,attr"`
	// SpaceSettingsKernelGatewayAppSettingsCustomImage: min=0,max=200
	CustomImage []SpaceSettingsKernelGatewayAppSettingsCustomImage `hcl:"custom_image,block" validate:"min=0,max=200"`
	// SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type SpaceSettingsKernelGatewayAppSettingsCustomImage struct {
	// AppImageConfigName: string, required
	AppImageConfigName terra.StringValue `hcl:"app_image_config_name,attr" validate:"required"`
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// ImageVersionNumber: number, optional
	ImageVersionNumber terra.NumberValue `hcl:"image_version_number,attr"`
}

type SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionAlias: string, optional
	SagemakerImageVersionAlias terra.StringValue `hcl:"sagemaker_image_version_alias,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type SpaceSettingsSpaceStorageSettings struct {
	// SpaceSettingsSpaceStorageSettingsEbsStorageSettings: required
	EbsStorageSettings *SpaceSettingsSpaceStorageSettingsEbsStorageSettings `hcl:"ebs_storage_settings,block" validate:"required"`
}

type SpaceSettingsSpaceStorageSettingsEbsStorageSettings struct {
	// EbsVolumeSizeInGb: number, required
	EbsVolumeSizeInGb terra.NumberValue `hcl:"ebs_volume_size_in_gb,attr" validate:"required"`
}

type SpaceSharingSettings struct {
	// SharingType: string, required
	SharingType terra.StringValue `hcl:"sharing_type,attr" validate:"required"`
}

type OwnershipSettingsAttributes struct {
	ref terra.Reference
}

func (os OwnershipSettingsAttributes) InternalRef() (terra.Reference, error) {
	return os.ref, nil
}

func (os OwnershipSettingsAttributes) InternalWithRef(ref terra.Reference) OwnershipSettingsAttributes {
	return OwnershipSettingsAttributes{ref: ref}
}

func (os OwnershipSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return os.ref.InternalTokens()
}

func (os OwnershipSettingsAttributes) OwnerUserProfileName() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("owner_user_profile_name"))
}

type SpaceSettingsAttributes struct {
	ref terra.Reference
}

func (ss SpaceSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SpaceSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsAttributes {
	return SpaceSettingsAttributes{ref: ref}
}

func (ss SpaceSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SpaceSettingsAttributes) AppType() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("app_type"))
}

func (ss SpaceSettingsAttributes) CodeEditorAppSettings() terra.ListValue[SpaceSettingsCodeEditorAppSettingsAttributes] {
	return terra.ReferenceAsList[SpaceSettingsCodeEditorAppSettingsAttributes](ss.ref.Append("code_editor_app_settings"))
}

func (ss SpaceSettingsAttributes) CustomFileSystem() terra.ListValue[SpaceSettingsCustomFileSystemAttributes] {
	return terra.ReferenceAsList[SpaceSettingsCustomFileSystemAttributes](ss.ref.Append("custom_file_system"))
}

func (ss SpaceSettingsAttributes) JupyterLabAppSettings() terra.ListValue[SpaceSettingsJupyterLabAppSettingsAttributes] {
	return terra.ReferenceAsList[SpaceSettingsJupyterLabAppSettingsAttributes](ss.ref.Append("jupyter_lab_app_settings"))
}

func (ss SpaceSettingsAttributes) JupyterServerAppSettings() terra.ListValue[SpaceSettingsJupyterServerAppSettingsAttributes] {
	return terra.ReferenceAsList[SpaceSettingsJupyterServerAppSettingsAttributes](ss.ref.Append("jupyter_server_app_settings"))
}

func (ss SpaceSettingsAttributes) KernelGatewayAppSettings() terra.ListValue[SpaceSettingsKernelGatewayAppSettingsAttributes] {
	return terra.ReferenceAsList[SpaceSettingsKernelGatewayAppSettingsAttributes](ss.ref.Append("kernel_gateway_app_settings"))
}

func (ss SpaceSettingsAttributes) SpaceStorageSettings() terra.ListValue[SpaceSettingsSpaceStorageSettingsAttributes] {
	return terra.ReferenceAsList[SpaceSettingsSpaceStorageSettingsAttributes](ss.ref.Append("space_storage_settings"))
}

type SpaceSettingsCodeEditorAppSettingsAttributes struct {
	ref terra.Reference
}

func (ceas SpaceSettingsCodeEditorAppSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ceas.ref, nil
}

func (ceas SpaceSettingsCodeEditorAppSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsCodeEditorAppSettingsAttributes {
	return SpaceSettingsCodeEditorAppSettingsAttributes{ref: ref}
}

func (ceas SpaceSettingsCodeEditorAppSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ceas.ref.InternalTokens()
}

func (ceas SpaceSettingsCodeEditorAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceAsList[SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes](ceas.ref.Append("default_resource_spec"))
}

type SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) InternalRef() (terra.Reference, error) {
	return drs.ref, nil
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes {
	return SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return drs.ref.InternalTokens()
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("instance_type"))
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionAlias() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_alias"))
}

func (drs SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_arn"))
}

type SpaceSettingsCustomFileSystemAttributes struct {
	ref terra.Reference
}

func (cfs SpaceSettingsCustomFileSystemAttributes) InternalRef() (terra.Reference, error) {
	return cfs.ref, nil
}

func (cfs SpaceSettingsCustomFileSystemAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsCustomFileSystemAttributes {
	return SpaceSettingsCustomFileSystemAttributes{ref: ref}
}

func (cfs SpaceSettingsCustomFileSystemAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cfs.ref.InternalTokens()
}

func (cfs SpaceSettingsCustomFileSystemAttributes) EfsFileSystem() terra.ListValue[SpaceSettingsCustomFileSystemEfsFileSystemAttributes] {
	return terra.ReferenceAsList[SpaceSettingsCustomFileSystemEfsFileSystemAttributes](cfs.ref.Append("efs_file_system"))
}

type SpaceSettingsCustomFileSystemEfsFileSystemAttributes struct {
	ref terra.Reference
}

func (efs SpaceSettingsCustomFileSystemEfsFileSystemAttributes) InternalRef() (terra.Reference, error) {
	return efs.ref, nil
}

func (efs SpaceSettingsCustomFileSystemEfsFileSystemAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsCustomFileSystemEfsFileSystemAttributes {
	return SpaceSettingsCustomFileSystemEfsFileSystemAttributes{ref: ref}
}

func (efs SpaceSettingsCustomFileSystemEfsFileSystemAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return efs.ref.InternalTokens()
}

func (efs SpaceSettingsCustomFileSystemEfsFileSystemAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("file_system_id"))
}

type SpaceSettingsJupyterLabAppSettingsAttributes struct {
	ref terra.Reference
}

func (jlas SpaceSettingsJupyterLabAppSettingsAttributes) InternalRef() (terra.Reference, error) {
	return jlas.ref, nil
}

func (jlas SpaceSettingsJupyterLabAppSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsJupyterLabAppSettingsAttributes {
	return SpaceSettingsJupyterLabAppSettingsAttributes{ref: ref}
}

func (jlas SpaceSettingsJupyterLabAppSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jlas.ref.InternalTokens()
}

func (jlas SpaceSettingsJupyterLabAppSettingsAttributes) CodeRepository() terra.SetValue[SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes] {
	return terra.ReferenceAsSet[SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes](jlas.ref.Append("code_repository"))
}

func (jlas SpaceSettingsJupyterLabAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceAsList[SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes](jlas.ref.Append("default_resource_spec"))
}

type SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes struct {
	ref terra.Reference
}

func (cr SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes {
	return SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes{ref: ref}
}

func (cr SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr SpaceSettingsJupyterLabAppSettingsCodeRepositoryAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("repository_url"))
}

type SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) InternalRef() (terra.Reference, error) {
	return drs.ref, nil
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes {
	return SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return drs.ref.InternalTokens()
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("instance_type"))
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionAlias() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_alias"))
}

func (drs SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_arn"))
}

type SpaceSettingsJupyterServerAppSettingsAttributes struct {
	ref terra.Reference
}

func (jsas SpaceSettingsJupyterServerAppSettingsAttributes) InternalRef() (terra.Reference, error) {
	return jsas.ref, nil
}

func (jsas SpaceSettingsJupyterServerAppSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsJupyterServerAppSettingsAttributes {
	return SpaceSettingsJupyterServerAppSettingsAttributes{ref: ref}
}

func (jsas SpaceSettingsJupyterServerAppSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jsas.ref.InternalTokens()
}

func (jsas SpaceSettingsJupyterServerAppSettingsAttributes) LifecycleConfigArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](jsas.ref.Append("lifecycle_config_arns"))
}

func (jsas SpaceSettingsJupyterServerAppSettingsAttributes) CodeRepository() terra.SetValue[SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes] {
	return terra.ReferenceAsSet[SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes](jsas.ref.Append("code_repository"))
}

func (jsas SpaceSettingsJupyterServerAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceAsList[SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes](jsas.ref.Append("default_resource_spec"))
}

type SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes struct {
	ref terra.Reference
}

func (cr SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes {
	return SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes{ref: ref}
}

func (cr SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr SpaceSettingsJupyterServerAppSettingsCodeRepositoryAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("repository_url"))
}

type SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) InternalRef() (terra.Reference, error) {
	return drs.ref, nil
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes {
	return SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return drs.ref.InternalTokens()
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("instance_type"))
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionAlias() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_alias"))
}

func (drs SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_arn"))
}

type SpaceSettingsKernelGatewayAppSettingsAttributes struct {
	ref terra.Reference
}

func (kgas SpaceSettingsKernelGatewayAppSettingsAttributes) InternalRef() (terra.Reference, error) {
	return kgas.ref, nil
}

func (kgas SpaceSettingsKernelGatewayAppSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsKernelGatewayAppSettingsAttributes {
	return SpaceSettingsKernelGatewayAppSettingsAttributes{ref: ref}
}

func (kgas SpaceSettingsKernelGatewayAppSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kgas.ref.InternalTokens()
}

func (kgas SpaceSettingsKernelGatewayAppSettingsAttributes) LifecycleConfigArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kgas.ref.Append("lifecycle_config_arns"))
}

func (kgas SpaceSettingsKernelGatewayAppSettingsAttributes) CustomImage() terra.ListValue[SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes] {
	return terra.ReferenceAsList[SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes](kgas.ref.Append("custom_image"))
}

func (kgas SpaceSettingsKernelGatewayAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceAsList[SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes](kgas.ref.Append("default_resource_spec"))
}

type SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes struct {
	ref terra.Reference
}

func (ci SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes {
	return SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes{ref: ref}
}

func (ci SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes) AppImageConfigName() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("app_image_config_name"))
}

func (ci SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("image_name"))
}

func (ci SpaceSettingsKernelGatewayAppSettingsCustomImageAttributes) ImageVersionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("image_version_number"))
}

type SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) InternalRef() (terra.Reference, error) {
	return drs.ref, nil
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes {
	return SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return drs.ref.InternalTokens()
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("instance_type"))
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionAlias() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_alias"))
}

func (drs SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceAsString(drs.ref.Append("sagemaker_image_version_arn"))
}

type SpaceSettingsSpaceStorageSettingsAttributes struct {
	ref terra.Reference
}

func (sss SpaceSettingsSpaceStorageSettingsAttributes) InternalRef() (terra.Reference, error) {
	return sss.ref, nil
}

func (sss SpaceSettingsSpaceStorageSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsSpaceStorageSettingsAttributes {
	return SpaceSettingsSpaceStorageSettingsAttributes{ref: ref}
}

func (sss SpaceSettingsSpaceStorageSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sss.ref.InternalTokens()
}

func (sss SpaceSettingsSpaceStorageSettingsAttributes) EbsStorageSettings() terra.ListValue[SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes] {
	return terra.ReferenceAsList[SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes](sss.ref.Append("ebs_storage_settings"))
}

type SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes struct {
	ref terra.Reference
}

func (ess SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ess.ref, nil
}

func (ess SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes {
	return SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes{ref: ref}
}

func (ess SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ess.ref.InternalTokens()
}

func (ess SpaceSettingsSpaceStorageSettingsEbsStorageSettingsAttributes) EbsVolumeSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ess.ref.Append("ebs_volume_size_in_gb"))
}

type SpaceSharingSettingsAttributes struct {
	ref terra.Reference
}

func (sss SpaceSharingSettingsAttributes) InternalRef() (terra.Reference, error) {
	return sss.ref, nil
}

func (sss SpaceSharingSettingsAttributes) InternalWithRef(ref terra.Reference) SpaceSharingSettingsAttributes {
	return SpaceSharingSettingsAttributes{ref: ref}
}

func (sss SpaceSharingSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sss.ref.InternalTokens()
}

func (sss SpaceSharingSettingsAttributes) SharingType() terra.StringValue {
	return terra.ReferenceAsString(sss.ref.Append("sharing_type"))
}

type OwnershipSettingsState struct {
	OwnerUserProfileName string `json:"owner_user_profile_name"`
}

type SpaceSettingsState struct {
	AppType                  string                                       `json:"app_type"`
	CodeEditorAppSettings    []SpaceSettingsCodeEditorAppSettingsState    `json:"code_editor_app_settings"`
	CustomFileSystem         []SpaceSettingsCustomFileSystemState         `json:"custom_file_system"`
	JupyterLabAppSettings    []SpaceSettingsJupyterLabAppSettingsState    `json:"jupyter_lab_app_settings"`
	JupyterServerAppSettings []SpaceSettingsJupyterServerAppSettingsState `json:"jupyter_server_app_settings"`
	KernelGatewayAppSettings []SpaceSettingsKernelGatewayAppSettingsState `json:"kernel_gateway_app_settings"`
	SpaceStorageSettings     []SpaceSettingsSpaceStorageSettingsState     `json:"space_storage_settings"`
}

type SpaceSettingsCodeEditorAppSettingsState struct {
	DefaultResourceSpec []SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecState struct {
	InstanceType               string `json:"instance_type"`
	LifecycleConfigArn         string `json:"lifecycle_config_arn"`
	SagemakerImageArn          string `json:"sagemaker_image_arn"`
	SagemakerImageVersionAlias string `json:"sagemaker_image_version_alias"`
	SagemakerImageVersionArn   string `json:"sagemaker_image_version_arn"`
}

type SpaceSettingsCustomFileSystemState struct {
	EfsFileSystem []SpaceSettingsCustomFileSystemEfsFileSystemState `json:"efs_file_system"`
}

type SpaceSettingsCustomFileSystemEfsFileSystemState struct {
	FileSystemId string `json:"file_system_id"`
}

type SpaceSettingsJupyterLabAppSettingsState struct {
	CodeRepository      []SpaceSettingsJupyterLabAppSettingsCodeRepositoryState      `json:"code_repository"`
	DefaultResourceSpec []SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type SpaceSettingsJupyterLabAppSettingsCodeRepositoryState struct {
	RepositoryUrl string `json:"repository_url"`
}

type SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecState struct {
	InstanceType               string `json:"instance_type"`
	LifecycleConfigArn         string `json:"lifecycle_config_arn"`
	SagemakerImageArn          string `json:"sagemaker_image_arn"`
	SagemakerImageVersionAlias string `json:"sagemaker_image_version_alias"`
	SagemakerImageVersionArn   string `json:"sagemaker_image_version_arn"`
}

type SpaceSettingsJupyterServerAppSettingsState struct {
	LifecycleConfigArns []string                                                        `json:"lifecycle_config_arns"`
	CodeRepository      []SpaceSettingsJupyterServerAppSettingsCodeRepositoryState      `json:"code_repository"`
	DefaultResourceSpec []SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type SpaceSettingsJupyterServerAppSettingsCodeRepositoryState struct {
	RepositoryUrl string `json:"repository_url"`
}

type SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecState struct {
	InstanceType               string `json:"instance_type"`
	LifecycleConfigArn         string `json:"lifecycle_config_arn"`
	SagemakerImageArn          string `json:"sagemaker_image_arn"`
	SagemakerImageVersionAlias string `json:"sagemaker_image_version_alias"`
	SagemakerImageVersionArn   string `json:"sagemaker_image_version_arn"`
}

type SpaceSettingsKernelGatewayAppSettingsState struct {
	LifecycleConfigArns []string                                                        `json:"lifecycle_config_arns"`
	CustomImage         []SpaceSettingsKernelGatewayAppSettingsCustomImageState         `json:"custom_image"`
	DefaultResourceSpec []SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type SpaceSettingsKernelGatewayAppSettingsCustomImageState struct {
	AppImageConfigName string  `json:"app_image_config_name"`
	ImageName          string  `json:"image_name"`
	ImageVersionNumber float64 `json:"image_version_number"`
}

type SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecState struct {
	InstanceType               string `json:"instance_type"`
	LifecycleConfigArn         string `json:"lifecycle_config_arn"`
	SagemakerImageArn          string `json:"sagemaker_image_arn"`
	SagemakerImageVersionAlias string `json:"sagemaker_image_version_alias"`
	SagemakerImageVersionArn   string `json:"sagemaker_image_version_arn"`
}

type SpaceSettingsSpaceStorageSettingsState struct {
	EbsStorageSettings []SpaceSettingsSpaceStorageSettingsEbsStorageSettingsState `json:"ebs_storage_settings"`
}

type SpaceSettingsSpaceStorageSettingsEbsStorageSettingsState struct {
	EbsVolumeSizeInGb float64 `json:"ebs_volume_size_in_gb"`
}

type SpaceSharingSettingsState struct {
	SharingType string `json:"sharing_type"`
}
