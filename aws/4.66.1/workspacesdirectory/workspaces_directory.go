// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package workspacesdirectory

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SelfServicePermissions struct {
	// ChangeComputeType: bool, optional
	ChangeComputeType terra.BoolValue `hcl:"change_compute_type,attr"`
	// IncreaseVolumeSize: bool, optional
	IncreaseVolumeSize terra.BoolValue `hcl:"increase_volume_size,attr"`
	// RebuildWorkspace: bool, optional
	RebuildWorkspace terra.BoolValue `hcl:"rebuild_workspace,attr"`
	// RestartWorkspace: bool, optional
	RestartWorkspace terra.BoolValue `hcl:"restart_workspace,attr"`
	// SwitchRunningMode: bool, optional
	SwitchRunningMode terra.BoolValue `hcl:"switch_running_mode,attr"`
}

type WorkspaceAccessProperties struct {
	// DeviceTypeAndroid: string, optional
	DeviceTypeAndroid terra.StringValue `hcl:"device_type_android,attr"`
	// DeviceTypeChromeos: string, optional
	DeviceTypeChromeos terra.StringValue `hcl:"device_type_chromeos,attr"`
	// DeviceTypeIos: string, optional
	DeviceTypeIos terra.StringValue `hcl:"device_type_ios,attr"`
	// DeviceTypeLinux: string, optional
	DeviceTypeLinux terra.StringValue `hcl:"device_type_linux,attr"`
	// DeviceTypeOsx: string, optional
	DeviceTypeOsx terra.StringValue `hcl:"device_type_osx,attr"`
	// DeviceTypeWeb: string, optional
	DeviceTypeWeb terra.StringValue `hcl:"device_type_web,attr"`
	// DeviceTypeWindows: string, optional
	DeviceTypeWindows terra.StringValue `hcl:"device_type_windows,attr"`
	// DeviceTypeZeroclient: string, optional
	DeviceTypeZeroclient terra.StringValue `hcl:"device_type_zeroclient,attr"`
}

type WorkspaceCreationProperties struct {
	// CustomSecurityGroupId: string, optional
	CustomSecurityGroupId terra.StringValue `hcl:"custom_security_group_id,attr"`
	// DefaultOu: string, optional
	DefaultOu terra.StringValue `hcl:"default_ou,attr"`
	// EnableInternetAccess: bool, optional
	EnableInternetAccess terra.BoolValue `hcl:"enable_internet_access,attr"`
	// EnableMaintenanceMode: bool, optional
	EnableMaintenanceMode terra.BoolValue `hcl:"enable_maintenance_mode,attr"`
	// UserEnabledAsLocalAdministrator: bool, optional
	UserEnabledAsLocalAdministrator terra.BoolValue `hcl:"user_enabled_as_local_administrator,attr"`
}

type SelfServicePermissionsAttributes struct {
	ref terra.Reference
}

func (ssp SelfServicePermissionsAttributes) InternalRef() (terra.Reference, error) {
	return ssp.ref, nil
}

func (ssp SelfServicePermissionsAttributes) InternalWithRef(ref terra.Reference) SelfServicePermissionsAttributes {
	return SelfServicePermissionsAttributes{ref: ref}
}

func (ssp SelfServicePermissionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssp.ref.InternalTokens()
}

func (ssp SelfServicePermissionsAttributes) ChangeComputeType() terra.BoolValue {
	return terra.ReferenceAsBool(ssp.ref.Append("change_compute_type"))
}

func (ssp SelfServicePermissionsAttributes) IncreaseVolumeSize() terra.BoolValue {
	return terra.ReferenceAsBool(ssp.ref.Append("increase_volume_size"))
}

func (ssp SelfServicePermissionsAttributes) RebuildWorkspace() terra.BoolValue {
	return terra.ReferenceAsBool(ssp.ref.Append("rebuild_workspace"))
}

func (ssp SelfServicePermissionsAttributes) RestartWorkspace() terra.BoolValue {
	return terra.ReferenceAsBool(ssp.ref.Append("restart_workspace"))
}

func (ssp SelfServicePermissionsAttributes) SwitchRunningMode() terra.BoolValue {
	return terra.ReferenceAsBool(ssp.ref.Append("switch_running_mode"))
}

type WorkspaceAccessPropertiesAttributes struct {
	ref terra.Reference
}

func (wap WorkspaceAccessPropertiesAttributes) InternalRef() (terra.Reference, error) {
	return wap.ref, nil
}

func (wap WorkspaceAccessPropertiesAttributes) InternalWithRef(ref terra.Reference) WorkspaceAccessPropertiesAttributes {
	return WorkspaceAccessPropertiesAttributes{ref: ref}
}

func (wap WorkspaceAccessPropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wap.ref.InternalTokens()
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeAndroid() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_android"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeChromeos() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_chromeos"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeIos() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_ios"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeLinux() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_linux"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeOsx() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_osx"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeWeb() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_web"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeWindows() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_windows"))
}

func (wap WorkspaceAccessPropertiesAttributes) DeviceTypeZeroclient() terra.StringValue {
	return terra.ReferenceAsString(wap.ref.Append("device_type_zeroclient"))
}

type WorkspaceCreationPropertiesAttributes struct {
	ref terra.Reference
}

func (wcp WorkspaceCreationPropertiesAttributes) InternalRef() (terra.Reference, error) {
	return wcp.ref, nil
}

func (wcp WorkspaceCreationPropertiesAttributes) InternalWithRef(ref terra.Reference) WorkspaceCreationPropertiesAttributes {
	return WorkspaceCreationPropertiesAttributes{ref: ref}
}

func (wcp WorkspaceCreationPropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wcp.ref.InternalTokens()
}

func (wcp WorkspaceCreationPropertiesAttributes) CustomSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(wcp.ref.Append("custom_security_group_id"))
}

func (wcp WorkspaceCreationPropertiesAttributes) DefaultOu() terra.StringValue {
	return terra.ReferenceAsString(wcp.ref.Append("default_ou"))
}

func (wcp WorkspaceCreationPropertiesAttributes) EnableInternetAccess() terra.BoolValue {
	return terra.ReferenceAsBool(wcp.ref.Append("enable_internet_access"))
}

func (wcp WorkspaceCreationPropertiesAttributes) EnableMaintenanceMode() terra.BoolValue {
	return terra.ReferenceAsBool(wcp.ref.Append("enable_maintenance_mode"))
}

func (wcp WorkspaceCreationPropertiesAttributes) UserEnabledAsLocalAdministrator() terra.BoolValue {
	return terra.ReferenceAsBool(wcp.ref.Append("user_enabled_as_local_administrator"))
}

type SelfServicePermissionsState struct {
	ChangeComputeType  bool `json:"change_compute_type"`
	IncreaseVolumeSize bool `json:"increase_volume_size"`
	RebuildWorkspace   bool `json:"rebuild_workspace"`
	RestartWorkspace   bool `json:"restart_workspace"`
	SwitchRunningMode  bool `json:"switch_running_mode"`
}

type WorkspaceAccessPropertiesState struct {
	DeviceTypeAndroid    string `json:"device_type_android"`
	DeviceTypeChromeos   string `json:"device_type_chromeos"`
	DeviceTypeIos        string `json:"device_type_ios"`
	DeviceTypeLinux      string `json:"device_type_linux"`
	DeviceTypeOsx        string `json:"device_type_osx"`
	DeviceTypeWeb        string `json:"device_type_web"`
	DeviceTypeWindows    string `json:"device_type_windows"`
	DeviceTypeZeroclient string `json:"device_type_zeroclient"`
}

type WorkspaceCreationPropertiesState struct {
	CustomSecurityGroupId           string `json:"custom_security_group_id"`
	DefaultOu                       string `json:"default_ou"`
	EnableInternetAccess            bool   `json:"enable_internet_access"`
	EnableMaintenanceMode           bool   `json:"enable_maintenance_mode"`
	UserEnabledAsLocalAdministrator bool   `json:"user_enabled_as_local_administrator"`
}
