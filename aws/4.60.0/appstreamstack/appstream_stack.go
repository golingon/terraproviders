// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appstreamstack

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccessEndpoints struct {
	// EndpointType: string, required
	EndpointType terra.StringValue `hcl:"endpoint_type,attr" validate:"required"`
	// VpceId: string, optional
	VpceId terra.StringValue `hcl:"vpce_id,attr"`
}

type ApplicationSettings struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// SettingsGroup: string, optional
	SettingsGroup terra.StringValue `hcl:"settings_group,attr"`
}

type StorageConnectors struct {
	// ConnectorType: string, required
	ConnectorType terra.StringValue `hcl:"connector_type,attr" validate:"required"`
	// Domains: list of string, optional
	Domains terra.ListValue[terra.StringValue] `hcl:"domains,attr"`
	// ResourceIdentifier: string, optional
	ResourceIdentifier terra.StringValue `hcl:"resource_identifier,attr"`
}

type UserSettings struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Permission: string, required
	Permission terra.StringValue `hcl:"permission,attr" validate:"required"`
}

type AccessEndpointsAttributes struct {
	ref terra.Reference
}

func (ae AccessEndpointsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae AccessEndpointsAttributes) InternalWithRef(ref terra.Reference) AccessEndpointsAttributes {
	return AccessEndpointsAttributes{ref: ref}
}

func (ae AccessEndpointsAttributes) InternalTokens() hclwrite.Tokens {
	return ae.ref.InternalTokens()
}

func (ae AccessEndpointsAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("endpoint_type"))
}

func (ae AccessEndpointsAttributes) VpceId() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("vpce_id"))
}

type ApplicationSettingsAttributes struct {
	ref terra.Reference
}

func (as ApplicationSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as ApplicationSettingsAttributes) InternalWithRef(ref terra.Reference) ApplicationSettingsAttributes {
	return ApplicationSettingsAttributes{ref: ref}
}

func (as ApplicationSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return as.ref.InternalTokens()
}

func (as ApplicationSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("enabled"))
}

func (as ApplicationSettingsAttributes) SettingsGroup() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("settings_group"))
}

type StorageConnectorsAttributes struct {
	ref terra.Reference
}

func (sc StorageConnectorsAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc StorageConnectorsAttributes) InternalWithRef(ref terra.Reference) StorageConnectorsAttributes {
	return StorageConnectorsAttributes{ref: ref}
}

func (sc StorageConnectorsAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc StorageConnectorsAttributes) ConnectorType() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("connector_type"))
}

func (sc StorageConnectorsAttributes) Domains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("domains"))
}

func (sc StorageConnectorsAttributes) ResourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("resource_identifier"))
}

type UserSettingsAttributes struct {
	ref terra.Reference
}

func (us UserSettingsAttributes) InternalRef() (terra.Reference, error) {
	return us.ref, nil
}

func (us UserSettingsAttributes) InternalWithRef(ref terra.Reference) UserSettingsAttributes {
	return UserSettingsAttributes{ref: ref}
}

func (us UserSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return us.ref.InternalTokens()
}

func (us UserSettingsAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(us.ref.Append("action"))
}

func (us UserSettingsAttributes) Permission() terra.StringValue {
	return terra.ReferenceAsString(us.ref.Append("permission"))
}

type AccessEndpointsState struct {
	EndpointType string `json:"endpoint_type"`
	VpceId       string `json:"vpce_id"`
}

type ApplicationSettingsState struct {
	Enabled       bool   `json:"enabled"`
	SettingsGroup string `json:"settings_group"`
}

type StorageConnectorsState struct {
	ConnectorType      string   `json:"connector_type"`
	Domains            []string `json:"domains"`
	ResourceIdentifier string   `json:"resource_identifier"`
}

type UserSettingsState struct {
	Action     string `json:"action"`
	Permission string `json:"permission"`
}
