// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package apigeeaddonsconfig

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AddonsConfig struct {
	// AdvancedApiOpsConfig: optional
	AdvancedApiOpsConfig *AdvancedApiOpsConfig `hcl:"advanced_api_ops_config,block"`
	// ApiSecurityConfig: optional
	ApiSecurityConfig *ApiSecurityConfig `hcl:"api_security_config,block"`
	// ConnectorsPlatformConfig: optional
	ConnectorsPlatformConfig *ConnectorsPlatformConfig `hcl:"connectors_platform_config,block"`
	// IntegrationConfig: optional
	IntegrationConfig *IntegrationConfig `hcl:"integration_config,block"`
	// MonetizationConfig: optional
	MonetizationConfig *MonetizationConfig `hcl:"monetization_config,block"`
}

type AdvancedApiOpsConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type ApiSecurityConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type ConnectorsPlatformConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type IntegrationConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type MonetizationConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AddonsConfigAttributes struct {
	ref terra.Reference
}

func (ac AddonsConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AddonsConfigAttributes) InternalWithRef(ref terra.Reference) AddonsConfigAttributes {
	return AddonsConfigAttributes{ref: ref}
}

func (ac AddonsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AddonsConfigAttributes) AdvancedApiOpsConfig() terra.ListValue[AdvancedApiOpsConfigAttributes] {
	return terra.ReferenceAsList[AdvancedApiOpsConfigAttributes](ac.ref.Append("advanced_api_ops_config"))
}

func (ac AddonsConfigAttributes) ApiSecurityConfig() terra.ListValue[ApiSecurityConfigAttributes] {
	return terra.ReferenceAsList[ApiSecurityConfigAttributes](ac.ref.Append("api_security_config"))
}

func (ac AddonsConfigAttributes) ConnectorsPlatformConfig() terra.ListValue[ConnectorsPlatformConfigAttributes] {
	return terra.ReferenceAsList[ConnectorsPlatformConfigAttributes](ac.ref.Append("connectors_platform_config"))
}

func (ac AddonsConfigAttributes) IntegrationConfig() terra.ListValue[IntegrationConfigAttributes] {
	return terra.ReferenceAsList[IntegrationConfigAttributes](ac.ref.Append("integration_config"))
}

func (ac AddonsConfigAttributes) MonetizationConfig() terra.ListValue[MonetizationConfigAttributes] {
	return terra.ReferenceAsList[MonetizationConfigAttributes](ac.ref.Append("monetization_config"))
}

type AdvancedApiOpsConfigAttributes struct {
	ref terra.Reference
}

func (aaoc AdvancedApiOpsConfigAttributes) InternalRef() (terra.Reference, error) {
	return aaoc.ref, nil
}

func (aaoc AdvancedApiOpsConfigAttributes) InternalWithRef(ref terra.Reference) AdvancedApiOpsConfigAttributes {
	return AdvancedApiOpsConfigAttributes{ref: ref}
}

func (aaoc AdvancedApiOpsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aaoc.ref.InternalTokens()
}

func (aaoc AdvancedApiOpsConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(aaoc.ref.Append("enabled"))
}

type ApiSecurityConfigAttributes struct {
	ref terra.Reference
}

func (asc ApiSecurityConfigAttributes) InternalRef() (terra.Reference, error) {
	return asc.ref, nil
}

func (asc ApiSecurityConfigAttributes) InternalWithRef(ref terra.Reference) ApiSecurityConfigAttributes {
	return ApiSecurityConfigAttributes{ref: ref}
}

func (asc ApiSecurityConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return asc.ref.InternalTokens()
}

func (asc ApiSecurityConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(asc.ref.Append("enabled"))
}

func (asc ApiSecurityConfigAttributes) ExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("expires_at"))
}

type ConnectorsPlatformConfigAttributes struct {
	ref terra.Reference
}

func (cpc ConnectorsPlatformConfigAttributes) InternalRef() (terra.Reference, error) {
	return cpc.ref, nil
}

func (cpc ConnectorsPlatformConfigAttributes) InternalWithRef(ref terra.Reference) ConnectorsPlatformConfigAttributes {
	return ConnectorsPlatformConfigAttributes{ref: ref}
}

func (cpc ConnectorsPlatformConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cpc.ref.InternalTokens()
}

func (cpc ConnectorsPlatformConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cpc.ref.Append("enabled"))
}

func (cpc ConnectorsPlatformConfigAttributes) ExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("expires_at"))
}

type IntegrationConfigAttributes struct {
	ref terra.Reference
}

func (ic IntegrationConfigAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IntegrationConfigAttributes) InternalWithRef(ref terra.Reference) IntegrationConfigAttributes {
	return IntegrationConfigAttributes{ref: ref}
}

func (ic IntegrationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IntegrationConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("enabled"))
}

type MonetizationConfigAttributes struct {
	ref terra.Reference
}

func (mc MonetizationConfigAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MonetizationConfigAttributes) InternalWithRef(ref terra.Reference) MonetizationConfigAttributes {
	return MonetizationConfigAttributes{ref: ref}
}

func (mc MonetizationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MonetizationConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("enabled"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AddonsConfigState struct {
	AdvancedApiOpsConfig     []AdvancedApiOpsConfigState     `json:"advanced_api_ops_config"`
	ApiSecurityConfig        []ApiSecurityConfigState        `json:"api_security_config"`
	ConnectorsPlatformConfig []ConnectorsPlatformConfigState `json:"connectors_platform_config"`
	IntegrationConfig        []IntegrationConfigState        `json:"integration_config"`
	MonetizationConfig       []MonetizationConfigState       `json:"monetization_config"`
}

type AdvancedApiOpsConfigState struct {
	Enabled bool `json:"enabled"`
}

type ApiSecurityConfigState struct {
	Enabled   bool   `json:"enabled"`
	ExpiresAt string `json:"expires_at"`
}

type ConnectorsPlatformConfigState struct {
	Enabled   bool   `json:"enabled"`
	ExpiresAt string `json:"expires_at"`
}

type IntegrationConfigState struct {
	Enabled bool `json:"enabled"`
}

type MonetizationConfigState struct {
	Enabled bool `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
