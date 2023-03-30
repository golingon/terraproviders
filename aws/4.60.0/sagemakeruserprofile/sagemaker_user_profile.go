// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakeruserprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type UserSettings struct {
	// ExecutionRole: string, required
	ExecutionRole terra.StringValue `hcl:"execution_role,attr" validate:"required"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// CanvasAppSettings: optional
	CanvasAppSettings *CanvasAppSettings `hcl:"canvas_app_settings,block"`
	// JupyterServerAppSettings: optional
	JupyterServerAppSettings *JupyterServerAppSettings `hcl:"jupyter_server_app_settings,block"`
	// KernelGatewayAppSettings: optional
	KernelGatewayAppSettings *KernelGatewayAppSettings `hcl:"kernel_gateway_app_settings,block"`
	// RSessionAppSettings: optional
	RSessionAppSettings *RSessionAppSettings `hcl:"r_session_app_settings,block"`
	// SharingSettings: optional
	SharingSettings *SharingSettings `hcl:"sharing_settings,block"`
	// TensorBoardAppSettings: optional
	TensorBoardAppSettings *TensorBoardAppSettings `hcl:"tensor_board_app_settings,block"`
}

type CanvasAppSettings struct {
	// TimeSeriesForecastingSettings: optional
	TimeSeriesForecastingSettings *TimeSeriesForecastingSettings `hcl:"time_series_forecasting_settings,block"`
}

type TimeSeriesForecastingSettings struct {
	// AmazonForecastRoleArn: string, optional
	AmazonForecastRoleArn terra.StringValue `hcl:"amazon_forecast_role_arn,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
}

type JupyterServerAppSettings struct {
	// LifecycleConfigArns: set of string, optional
	LifecycleConfigArns terra.SetValue[terra.StringValue] `hcl:"lifecycle_config_arns,attr"`
	// CodeRepository: min=0,max=10
	CodeRepository []CodeRepository `hcl:"code_repository,block" validate:"min=0,max=10"`
	// JupyterServerAppSettingsDefaultResourceSpec: optional
	DefaultResourceSpec *JupyterServerAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block"`
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
	// KernelGatewayAppSettingsCustomImage: min=0,max=30
	CustomImage []KernelGatewayAppSettingsCustomImage `hcl:"custom_image,block" validate:"min=0,max=30"`
	// KernelGatewayAppSettingsDefaultResourceSpec: optional
	DefaultResourceSpec *KernelGatewayAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block"`
}

type KernelGatewayAppSettingsCustomImage struct {
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

type RSessionAppSettings struct {
	// RSessionAppSettingsCustomImage: min=0,max=30
	CustomImage []RSessionAppSettingsCustomImage `hcl:"custom_image,block" validate:"min=0,max=30"`
	// RSessionAppSettingsDefaultResourceSpec: optional
	DefaultResourceSpec *RSessionAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block"`
}

type RSessionAppSettingsCustomImage struct {
	// AppImageConfigName: string, required
	AppImageConfigName terra.StringValue `hcl:"app_image_config_name,attr" validate:"required"`
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// ImageVersionNumber: number, optional
	ImageVersionNumber terra.NumberValue `hcl:"image_version_number,attr"`
}

type RSessionAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type SharingSettings struct {
	// NotebookOutputOption: string, optional
	NotebookOutputOption terra.StringValue `hcl:"notebook_output_option,attr"`
	// S3KmsKeyId: string, optional
	S3KmsKeyId terra.StringValue `hcl:"s3_kms_key_id,attr"`
	// S3OutputPath: string, optional
	S3OutputPath terra.StringValue `hcl:"s3_output_path,attr"`
}

type TensorBoardAppSettings struct {
	// TensorBoardAppSettingsDefaultResourceSpec: required
	DefaultResourceSpec *TensorBoardAppSettingsDefaultResourceSpec `hcl:"default_resource_spec,block" validate:"required"`
}

type TensorBoardAppSettingsDefaultResourceSpec struct {
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LifecycleConfigArn: string, optional
	LifecycleConfigArn terra.StringValue `hcl:"lifecycle_config_arn,attr"`
	// SagemakerImageArn: string, optional
	SagemakerImageArn terra.StringValue `hcl:"sagemaker_image_arn,attr"`
	// SagemakerImageVersionArn: string, optional
	SagemakerImageVersionArn terra.StringValue `hcl:"sagemaker_image_version_arn,attr"`
}

type UserSettingsAttributes struct {
	ref terra.Reference
}

func (us UserSettingsAttributes) InternalRef() terra.Reference {
	return us.ref
}

func (us UserSettingsAttributes) InternalWithRef(ref terra.Reference) UserSettingsAttributes {
	return UserSettingsAttributes{ref: ref}
}

func (us UserSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return us.ref.InternalTokens()
}

func (us UserSettingsAttributes) ExecutionRole() terra.StringValue {
	return terra.ReferenceString(us.ref.Append("execution_role"))
}

func (us UserSettingsAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](us.ref.Append("security_groups"))
}

func (us UserSettingsAttributes) CanvasAppSettings() terra.ListValue[CanvasAppSettingsAttributes] {
	return terra.ReferenceList[CanvasAppSettingsAttributes](us.ref.Append("canvas_app_settings"))
}

func (us UserSettingsAttributes) JupyterServerAppSettings() terra.ListValue[JupyterServerAppSettingsAttributes] {
	return terra.ReferenceList[JupyterServerAppSettingsAttributes](us.ref.Append("jupyter_server_app_settings"))
}

func (us UserSettingsAttributes) KernelGatewayAppSettings() terra.ListValue[KernelGatewayAppSettingsAttributes] {
	return terra.ReferenceList[KernelGatewayAppSettingsAttributes](us.ref.Append("kernel_gateway_app_settings"))
}

func (us UserSettingsAttributes) RSessionAppSettings() terra.ListValue[RSessionAppSettingsAttributes] {
	return terra.ReferenceList[RSessionAppSettingsAttributes](us.ref.Append("r_session_app_settings"))
}

func (us UserSettingsAttributes) SharingSettings() terra.ListValue[SharingSettingsAttributes] {
	return terra.ReferenceList[SharingSettingsAttributes](us.ref.Append("sharing_settings"))
}

func (us UserSettingsAttributes) TensorBoardAppSettings() terra.ListValue[TensorBoardAppSettingsAttributes] {
	return terra.ReferenceList[TensorBoardAppSettingsAttributes](us.ref.Append("tensor_board_app_settings"))
}

type CanvasAppSettingsAttributes struct {
	ref terra.Reference
}

func (cas CanvasAppSettingsAttributes) InternalRef() terra.Reference {
	return cas.ref
}

func (cas CanvasAppSettingsAttributes) InternalWithRef(ref terra.Reference) CanvasAppSettingsAttributes {
	return CanvasAppSettingsAttributes{ref: ref}
}

func (cas CanvasAppSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return cas.ref.InternalTokens()
}

func (cas CanvasAppSettingsAttributes) TimeSeriesForecastingSettings() terra.ListValue[TimeSeriesForecastingSettingsAttributes] {
	return terra.ReferenceList[TimeSeriesForecastingSettingsAttributes](cas.ref.Append("time_series_forecasting_settings"))
}

type TimeSeriesForecastingSettingsAttributes struct {
	ref terra.Reference
}

func (tsfs TimeSeriesForecastingSettingsAttributes) InternalRef() terra.Reference {
	return tsfs.ref
}

func (tsfs TimeSeriesForecastingSettingsAttributes) InternalWithRef(ref terra.Reference) TimeSeriesForecastingSettingsAttributes {
	return TimeSeriesForecastingSettingsAttributes{ref: ref}
}

func (tsfs TimeSeriesForecastingSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return tsfs.ref.InternalTokens()
}

func (tsfs TimeSeriesForecastingSettingsAttributes) AmazonForecastRoleArn() terra.StringValue {
	return terra.ReferenceString(tsfs.ref.Append("amazon_forecast_role_arn"))
}

func (tsfs TimeSeriesForecastingSettingsAttributes) Status() terra.StringValue {
	return terra.ReferenceString(tsfs.ref.Append("status"))
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

func (kgas KernelGatewayAppSettingsAttributes) CustomImage() terra.ListValue[KernelGatewayAppSettingsCustomImageAttributes] {
	return terra.ReferenceList[KernelGatewayAppSettingsCustomImageAttributes](kgas.ref.Append("custom_image"))
}

func (kgas KernelGatewayAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[KernelGatewayAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceList[KernelGatewayAppSettingsDefaultResourceSpecAttributes](kgas.ref.Append("default_resource_spec"))
}

type KernelGatewayAppSettingsCustomImageAttributes struct {
	ref terra.Reference
}

func (ci KernelGatewayAppSettingsCustomImageAttributes) InternalRef() terra.Reference {
	return ci.ref
}

func (ci KernelGatewayAppSettingsCustomImageAttributes) InternalWithRef(ref terra.Reference) KernelGatewayAppSettingsCustomImageAttributes {
	return KernelGatewayAppSettingsCustomImageAttributes{ref: ref}
}

func (ci KernelGatewayAppSettingsCustomImageAttributes) InternalTokens() hclwrite.Tokens {
	return ci.ref.InternalTokens()
}

func (ci KernelGatewayAppSettingsCustomImageAttributes) AppImageConfigName() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("app_image_config_name"))
}

func (ci KernelGatewayAppSettingsCustomImageAttributes) ImageName() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("image_name"))
}

func (ci KernelGatewayAppSettingsCustomImageAttributes) ImageVersionNumber() terra.NumberValue {
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

type RSessionAppSettingsAttributes struct {
	ref terra.Reference
}

func (rsas RSessionAppSettingsAttributes) InternalRef() terra.Reference {
	return rsas.ref
}

func (rsas RSessionAppSettingsAttributes) InternalWithRef(ref terra.Reference) RSessionAppSettingsAttributes {
	return RSessionAppSettingsAttributes{ref: ref}
}

func (rsas RSessionAppSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return rsas.ref.InternalTokens()
}

func (rsas RSessionAppSettingsAttributes) CustomImage() terra.ListValue[RSessionAppSettingsCustomImageAttributes] {
	return terra.ReferenceList[RSessionAppSettingsCustomImageAttributes](rsas.ref.Append("custom_image"))
}

func (rsas RSessionAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[RSessionAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceList[RSessionAppSettingsDefaultResourceSpecAttributes](rsas.ref.Append("default_resource_spec"))
}

type RSessionAppSettingsCustomImageAttributes struct {
	ref terra.Reference
}

func (ci RSessionAppSettingsCustomImageAttributes) InternalRef() terra.Reference {
	return ci.ref
}

func (ci RSessionAppSettingsCustomImageAttributes) InternalWithRef(ref terra.Reference) RSessionAppSettingsCustomImageAttributes {
	return RSessionAppSettingsCustomImageAttributes{ref: ref}
}

func (ci RSessionAppSettingsCustomImageAttributes) InternalTokens() hclwrite.Tokens {
	return ci.ref.InternalTokens()
}

func (ci RSessionAppSettingsCustomImageAttributes) AppImageConfigName() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("app_image_config_name"))
}

func (ci RSessionAppSettingsCustomImageAttributes) ImageName() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("image_name"))
}

func (ci RSessionAppSettingsCustomImageAttributes) ImageVersionNumber() terra.NumberValue {
	return terra.ReferenceNumber(ci.ref.Append("image_version_number"))
}

type RSessionAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) InternalRef() terra.Reference {
	return drs.ref
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) RSessionAppSettingsDefaultResourceSpecAttributes {
	return RSessionAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) InternalTokens() hclwrite.Tokens {
	return drs.ref.InternalTokens()
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("instance_type"))
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs RSessionAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_version_arn"))
}

type SharingSettingsAttributes struct {
	ref terra.Reference
}

func (ss SharingSettingsAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss SharingSettingsAttributes) InternalWithRef(ref terra.Reference) SharingSettingsAttributes {
	return SharingSettingsAttributes{ref: ref}
}

func (ss SharingSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss SharingSettingsAttributes) NotebookOutputOption() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("notebook_output_option"))
}

func (ss SharingSettingsAttributes) S3KmsKeyId() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("s3_kms_key_id"))
}

func (ss SharingSettingsAttributes) S3OutputPath() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("s3_output_path"))
}

type TensorBoardAppSettingsAttributes struct {
	ref terra.Reference
}

func (tbas TensorBoardAppSettingsAttributes) InternalRef() terra.Reference {
	return tbas.ref
}

func (tbas TensorBoardAppSettingsAttributes) InternalWithRef(ref terra.Reference) TensorBoardAppSettingsAttributes {
	return TensorBoardAppSettingsAttributes{ref: ref}
}

func (tbas TensorBoardAppSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return tbas.ref.InternalTokens()
}

func (tbas TensorBoardAppSettingsAttributes) DefaultResourceSpec() terra.ListValue[TensorBoardAppSettingsDefaultResourceSpecAttributes] {
	return terra.ReferenceList[TensorBoardAppSettingsDefaultResourceSpecAttributes](tbas.ref.Append("default_resource_spec"))
}

type TensorBoardAppSettingsDefaultResourceSpecAttributes struct {
	ref terra.Reference
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) InternalRef() terra.Reference {
	return drs.ref
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) InternalWithRef(ref terra.Reference) TensorBoardAppSettingsDefaultResourceSpecAttributes {
	return TensorBoardAppSettingsDefaultResourceSpecAttributes{ref: ref}
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) InternalTokens() hclwrite.Tokens {
	return drs.ref.InternalTokens()
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("instance_type"))
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) LifecycleConfigArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("lifecycle_config_arn"))
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) SagemakerImageArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_arn"))
}

func (drs TensorBoardAppSettingsDefaultResourceSpecAttributes) SagemakerImageVersionArn() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("sagemaker_image_version_arn"))
}

type UserSettingsState struct {
	ExecutionRole            string                          `json:"execution_role"`
	SecurityGroups           []string                        `json:"security_groups"`
	CanvasAppSettings        []CanvasAppSettingsState        `json:"canvas_app_settings"`
	JupyterServerAppSettings []JupyterServerAppSettingsState `json:"jupyter_server_app_settings"`
	KernelGatewayAppSettings []KernelGatewayAppSettingsState `json:"kernel_gateway_app_settings"`
	RSessionAppSettings      []RSessionAppSettingsState      `json:"r_session_app_settings"`
	SharingSettings          []SharingSettingsState          `json:"sharing_settings"`
	TensorBoardAppSettings   []TensorBoardAppSettingsState   `json:"tensor_board_app_settings"`
}

type CanvasAppSettingsState struct {
	TimeSeriesForecastingSettings []TimeSeriesForecastingSettingsState `json:"time_series_forecasting_settings"`
}

type TimeSeriesForecastingSettingsState struct {
	AmazonForecastRoleArn string `json:"amazon_forecast_role_arn"`
	Status                string `json:"status"`
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
	CustomImage         []KernelGatewayAppSettingsCustomImageState         `json:"custom_image"`
	DefaultResourceSpec []KernelGatewayAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type KernelGatewayAppSettingsCustomImageState struct {
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

type RSessionAppSettingsState struct {
	CustomImage         []RSessionAppSettingsCustomImageState         `json:"custom_image"`
	DefaultResourceSpec []RSessionAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type RSessionAppSettingsCustomImageState struct {
	AppImageConfigName string  `json:"app_image_config_name"`
	ImageName          string  `json:"image_name"`
	ImageVersionNumber float64 `json:"image_version_number"`
}

type RSessionAppSettingsDefaultResourceSpecState struct {
	InstanceType             string `json:"instance_type"`
	LifecycleConfigArn       string `json:"lifecycle_config_arn"`
	SagemakerImageArn        string `json:"sagemaker_image_arn"`
	SagemakerImageVersionArn string `json:"sagemaker_image_version_arn"`
}

type SharingSettingsState struct {
	NotebookOutputOption string `json:"notebook_output_option"`
	S3KmsKeyId           string `json:"s3_kms_key_id"`
	S3OutputPath         string `json:"s3_output_path"`
}

type TensorBoardAppSettingsState struct {
	DefaultResourceSpec []TensorBoardAppSettingsDefaultResourceSpecState `json:"default_resource_spec"`
}

type TensorBoardAppSettingsDefaultResourceSpecState struct {
	InstanceType             string `json:"instance_type"`
	LifecycleConfigArn       string `json:"lifecycle_config_arn"`
	SagemakerImageArn        string `json:"sagemaker_image_arn"`
	SagemakerImageVersionArn string `json:"sagemaker_image_version_arn"`
}
