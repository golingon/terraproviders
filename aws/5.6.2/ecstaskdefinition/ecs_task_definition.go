// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ecstaskdefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EphemeralStorage struct {
	// SizeInGib: number, required
	SizeInGib terra.NumberValue `hcl:"size_in_gib,attr" validate:"required"`
}

type InferenceAccelerator struct {
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// DeviceType: string, required
	DeviceType terra.StringValue `hcl:"device_type,attr" validate:"required"`
}

type PlacementConstraints struct {
	// Expression: string, optional
	Expression terra.StringValue `hcl:"expression,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ProxyConfiguration struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type RuntimePlatform struct {
	// CpuArchitecture: string, optional
	CpuArchitecture terra.StringValue `hcl:"cpu_architecture,attr"`
	// OperatingSystemFamily: string, optional
	OperatingSystemFamily terra.StringValue `hcl:"operating_system_family,attr"`
}

type Volume struct {
	// HostPath: string, optional
	HostPath terra.StringValue `hcl:"host_path,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// DockerVolumeConfiguration: optional
	DockerVolumeConfiguration *DockerVolumeConfiguration `hcl:"docker_volume_configuration,block"`
	// EfsVolumeConfiguration: optional
	EfsVolumeConfiguration *EfsVolumeConfiguration `hcl:"efs_volume_configuration,block"`
	// FsxWindowsFileServerVolumeConfiguration: optional
	FsxWindowsFileServerVolumeConfiguration *FsxWindowsFileServerVolumeConfiguration `hcl:"fsx_windows_file_server_volume_configuration,block"`
}

type DockerVolumeConfiguration struct {
	// Autoprovision: bool, optional
	Autoprovision terra.BoolValue `hcl:"autoprovision,attr"`
	// Driver: string, optional
	Driver terra.StringValue `hcl:"driver,attr"`
	// DriverOpts: map of string, optional
	DriverOpts terra.MapValue[terra.StringValue] `hcl:"driver_opts,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
}

type EfsVolumeConfiguration struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// RootDirectory: string, optional
	RootDirectory terra.StringValue `hcl:"root_directory,attr"`
	// TransitEncryption: string, optional
	TransitEncryption terra.StringValue `hcl:"transit_encryption,attr"`
	// TransitEncryptionPort: number, optional
	TransitEncryptionPort terra.NumberValue `hcl:"transit_encryption_port,attr"`
	// EfsVolumeConfigurationAuthorizationConfig: optional
	AuthorizationConfig *EfsVolumeConfigurationAuthorizationConfig `hcl:"authorization_config,block"`
}

type EfsVolumeConfigurationAuthorizationConfig struct {
	// AccessPointId: string, optional
	AccessPointId terra.StringValue `hcl:"access_point_id,attr"`
	// Iam: string, optional
	Iam terra.StringValue `hcl:"iam,attr"`
}

type FsxWindowsFileServerVolumeConfiguration struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// RootDirectory: string, required
	RootDirectory terra.StringValue `hcl:"root_directory,attr" validate:"required"`
	// FsxWindowsFileServerVolumeConfigurationAuthorizationConfig: required
	AuthorizationConfig *FsxWindowsFileServerVolumeConfigurationAuthorizationConfig `hcl:"authorization_config,block" validate:"required"`
}

type FsxWindowsFileServerVolumeConfigurationAuthorizationConfig struct {
	// CredentialsParameter: string, required
	CredentialsParameter terra.StringValue `hcl:"credentials_parameter,attr" validate:"required"`
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
}

type EphemeralStorageAttributes struct {
	ref terra.Reference
}

func (es EphemeralStorageAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es EphemeralStorageAttributes) InternalWithRef(ref terra.Reference) EphemeralStorageAttributes {
	return EphemeralStorageAttributes{ref: ref}
}

func (es EphemeralStorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es EphemeralStorageAttributes) SizeInGib() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("size_in_gib"))
}

type InferenceAcceleratorAttributes struct {
	ref terra.Reference
}

func (ia InferenceAcceleratorAttributes) InternalRef() (terra.Reference, error) {
	return ia.ref, nil
}

func (ia InferenceAcceleratorAttributes) InternalWithRef(ref terra.Reference) InferenceAcceleratorAttributes {
	return InferenceAcceleratorAttributes{ref: ref}
}

func (ia InferenceAcceleratorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ia.ref.InternalTokens()
}

func (ia InferenceAcceleratorAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("device_name"))
}

func (ia InferenceAcceleratorAttributes) DeviceType() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("device_type"))
}

type PlacementConstraintsAttributes struct {
	ref terra.Reference
}

func (pc PlacementConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PlacementConstraintsAttributes) InternalWithRef(ref terra.Reference) PlacementConstraintsAttributes {
	return PlacementConstraintsAttributes{ref: ref}
}

func (pc PlacementConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PlacementConstraintsAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("expression"))
}

func (pc PlacementConstraintsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("type"))
}

type ProxyConfigurationAttributes struct {
	ref terra.Reference
}

func (pc ProxyConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ProxyConfigurationAttributes) InternalWithRef(ref terra.Reference) ProxyConfigurationAttributes {
	return ProxyConfigurationAttributes{ref: ref}
}

func (pc ProxyConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ProxyConfigurationAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("container_name"))
}

func (pc ProxyConfigurationAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("properties"))
}

func (pc ProxyConfigurationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("type"))
}

type RuntimePlatformAttributes struct {
	ref terra.Reference
}

func (rp RuntimePlatformAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RuntimePlatformAttributes) InternalWithRef(ref terra.Reference) RuntimePlatformAttributes {
	return RuntimePlatformAttributes{ref: ref}
}

func (rp RuntimePlatformAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RuntimePlatformAttributes) CpuArchitecture() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("cpu_architecture"))
}

func (rp RuntimePlatformAttributes) OperatingSystemFamily() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("operating_system_family"))
}

type VolumeAttributes struct {
	ref terra.Reference
}

func (v VolumeAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VolumeAttributes) InternalWithRef(ref terra.Reference) VolumeAttributes {
	return VolumeAttributes{ref: ref}
}

func (v VolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VolumeAttributes) HostPath() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("host_path"))
}

func (v VolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v VolumeAttributes) DockerVolumeConfiguration() terra.ListValue[DockerVolumeConfigurationAttributes] {
	return terra.ReferenceAsList[DockerVolumeConfigurationAttributes](v.ref.Append("docker_volume_configuration"))
}

func (v VolumeAttributes) EfsVolumeConfiguration() terra.ListValue[EfsVolumeConfigurationAttributes] {
	return terra.ReferenceAsList[EfsVolumeConfigurationAttributes](v.ref.Append("efs_volume_configuration"))
}

func (v VolumeAttributes) FsxWindowsFileServerVolumeConfiguration() terra.ListValue[FsxWindowsFileServerVolumeConfigurationAttributes] {
	return terra.ReferenceAsList[FsxWindowsFileServerVolumeConfigurationAttributes](v.ref.Append("fsx_windows_file_server_volume_configuration"))
}

type DockerVolumeConfigurationAttributes struct {
	ref terra.Reference
}

func (dvc DockerVolumeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return dvc.ref, nil
}

func (dvc DockerVolumeConfigurationAttributes) InternalWithRef(ref terra.Reference) DockerVolumeConfigurationAttributes {
	return DockerVolumeConfigurationAttributes{ref: ref}
}

func (dvc DockerVolumeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dvc.ref.InternalTokens()
}

func (dvc DockerVolumeConfigurationAttributes) Autoprovision() terra.BoolValue {
	return terra.ReferenceAsBool(dvc.ref.Append("autoprovision"))
}

func (dvc DockerVolumeConfigurationAttributes) Driver() terra.StringValue {
	return terra.ReferenceAsString(dvc.ref.Append("driver"))
}

func (dvc DockerVolumeConfigurationAttributes) DriverOpts() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dvc.ref.Append("driver_opts"))
}

func (dvc DockerVolumeConfigurationAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dvc.ref.Append("labels"))
}

func (dvc DockerVolumeConfigurationAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(dvc.ref.Append("scope"))
}

type EfsVolumeConfigurationAttributes struct {
	ref terra.Reference
}

func (evc EfsVolumeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return evc.ref, nil
}

func (evc EfsVolumeConfigurationAttributes) InternalWithRef(ref terra.Reference) EfsVolumeConfigurationAttributes {
	return EfsVolumeConfigurationAttributes{ref: ref}
}

func (evc EfsVolumeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return evc.ref.InternalTokens()
}

func (evc EfsVolumeConfigurationAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("file_system_id"))
}

func (evc EfsVolumeConfigurationAttributes) RootDirectory() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("root_directory"))
}

func (evc EfsVolumeConfigurationAttributes) TransitEncryption() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("transit_encryption"))
}

func (evc EfsVolumeConfigurationAttributes) TransitEncryptionPort() terra.NumberValue {
	return terra.ReferenceAsNumber(evc.ref.Append("transit_encryption_port"))
}

func (evc EfsVolumeConfigurationAttributes) AuthorizationConfig() terra.ListValue[EfsVolumeConfigurationAuthorizationConfigAttributes] {
	return terra.ReferenceAsList[EfsVolumeConfigurationAuthorizationConfigAttributes](evc.ref.Append("authorization_config"))
}

type EfsVolumeConfigurationAuthorizationConfigAttributes struct {
	ref terra.Reference
}

func (ac EfsVolumeConfigurationAuthorizationConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac EfsVolumeConfigurationAuthorizationConfigAttributes) InternalWithRef(ref terra.Reference) EfsVolumeConfigurationAuthorizationConfigAttributes {
	return EfsVolumeConfigurationAuthorizationConfigAttributes{ref: ref}
}

func (ac EfsVolumeConfigurationAuthorizationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac EfsVolumeConfigurationAuthorizationConfigAttributes) AccessPointId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("access_point_id"))
}

func (ac EfsVolumeConfigurationAuthorizationConfigAttributes) Iam() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("iam"))
}

type FsxWindowsFileServerVolumeConfigurationAttributes struct {
	ref terra.Reference
}

func (fwfsvc FsxWindowsFileServerVolumeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return fwfsvc.ref, nil
}

func (fwfsvc FsxWindowsFileServerVolumeConfigurationAttributes) InternalWithRef(ref terra.Reference) FsxWindowsFileServerVolumeConfigurationAttributes {
	return FsxWindowsFileServerVolumeConfigurationAttributes{ref: ref}
}

func (fwfsvc FsxWindowsFileServerVolumeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fwfsvc.ref.InternalTokens()
}

func (fwfsvc FsxWindowsFileServerVolumeConfigurationAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(fwfsvc.ref.Append("file_system_id"))
}

func (fwfsvc FsxWindowsFileServerVolumeConfigurationAttributes) RootDirectory() terra.StringValue {
	return terra.ReferenceAsString(fwfsvc.ref.Append("root_directory"))
}

func (fwfsvc FsxWindowsFileServerVolumeConfigurationAttributes) AuthorizationConfig() terra.ListValue[FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes] {
	return terra.ReferenceAsList[FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes](fwfsvc.ref.Append("authorization_config"))
}

type FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes struct {
	ref terra.Reference
}

func (ac FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes) InternalWithRef(ref terra.Reference) FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes {
	return FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes{ref: ref}
}

func (ac FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes) CredentialsParameter() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("credentials_parameter"))
}

func (ac FsxWindowsFileServerVolumeConfigurationAuthorizationConfigAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("domain"))
}

type EphemeralStorageState struct {
	SizeInGib float64 `json:"size_in_gib"`
}

type InferenceAcceleratorState struct {
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"`
}

type PlacementConstraintsState struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type ProxyConfigurationState struct {
	ContainerName string            `json:"container_name"`
	Properties    map[string]string `json:"properties"`
	Type          string            `json:"type"`
}

type RuntimePlatformState struct {
	CpuArchitecture       string `json:"cpu_architecture"`
	OperatingSystemFamily string `json:"operating_system_family"`
}

type VolumeState struct {
	HostPath                                string                                         `json:"host_path"`
	Name                                    string                                         `json:"name"`
	DockerVolumeConfiguration               []DockerVolumeConfigurationState               `json:"docker_volume_configuration"`
	EfsVolumeConfiguration                  []EfsVolumeConfigurationState                  `json:"efs_volume_configuration"`
	FsxWindowsFileServerVolumeConfiguration []FsxWindowsFileServerVolumeConfigurationState `json:"fsx_windows_file_server_volume_configuration"`
}

type DockerVolumeConfigurationState struct {
	Autoprovision bool              `json:"autoprovision"`
	Driver        string            `json:"driver"`
	DriverOpts    map[string]string `json:"driver_opts"`
	Labels        map[string]string `json:"labels"`
	Scope         string            `json:"scope"`
}

type EfsVolumeConfigurationState struct {
	FileSystemId          string                                           `json:"file_system_id"`
	RootDirectory         string                                           `json:"root_directory"`
	TransitEncryption     string                                           `json:"transit_encryption"`
	TransitEncryptionPort float64                                          `json:"transit_encryption_port"`
	AuthorizationConfig   []EfsVolumeConfigurationAuthorizationConfigState `json:"authorization_config"`
}

type EfsVolumeConfigurationAuthorizationConfigState struct {
	AccessPointId string `json:"access_point_id"`
	Iam           string `json:"iam"`
}

type FsxWindowsFileServerVolumeConfigurationState struct {
	FileSystemId        string                                                            `json:"file_system_id"`
	RootDirectory       string                                                            `json:"root_directory"`
	AuthorizationConfig []FsxWindowsFileServerVolumeConfigurationAuthorizationConfigState `json:"authorization_config"`
}

type FsxWindowsFileServerVolumeConfigurationAuthorizationConfigState struct {
	CredentialsParameter string `json:"credentials_parameter"`
	Domain               string `json:"domain"`
}
