// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_apigee_addons_config

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AddonsConfig struct {
	// AddonsConfigAdvancedApiOpsConfig: optional
	AdvancedApiOpsConfig *AddonsConfigAdvancedApiOpsConfig `hcl:"advanced_api_ops_config,block"`
	// AddonsConfigApiSecurityConfig: optional
	ApiSecurityConfig *AddonsConfigApiSecurityConfig `hcl:"api_security_config,block"`
	// AddonsConfigConnectorsPlatformConfig: optional
	ConnectorsPlatformConfig *AddonsConfigConnectorsPlatformConfig `hcl:"connectors_platform_config,block"`
	// AddonsConfigIntegrationConfig: optional
	IntegrationConfig *AddonsConfigIntegrationConfig `hcl:"integration_config,block"`
	// AddonsConfigMonetizationConfig: optional
	MonetizationConfig *AddonsConfigMonetizationConfig `hcl:"monetization_config,block"`
}

type AddonsConfigAdvancedApiOpsConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type AddonsConfigApiSecurityConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type AddonsConfigConnectorsPlatformConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type AddonsConfigIntegrationConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type AddonsConfigMonetizationConfig struct {
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

func (ac AddonsConfigAttributes) AdvancedApiOpsConfig() terra.ListValue[AddonsConfigAdvancedApiOpsConfigAttributes] {
	return terra.ReferenceAsList[AddonsConfigAdvancedApiOpsConfigAttributes](ac.ref.Append("advanced_api_ops_config"))
}

func (ac AddonsConfigAttributes) ApiSecurityConfig() terra.ListValue[AddonsConfigApiSecurityConfigAttributes] {
	return terra.ReferenceAsList[AddonsConfigApiSecurityConfigAttributes](ac.ref.Append("api_security_config"))
}

func (ac AddonsConfigAttributes) ConnectorsPlatformConfig() terra.ListValue[AddonsConfigConnectorsPlatformConfigAttributes] {
	return terra.ReferenceAsList[AddonsConfigConnectorsPlatformConfigAttributes](ac.ref.Append("connectors_platform_config"))
}

func (ac AddonsConfigAttributes) IntegrationConfig() terra.ListValue[AddonsConfigIntegrationConfigAttributes] {
	return terra.ReferenceAsList[AddonsConfigIntegrationConfigAttributes](ac.ref.Append("integration_config"))
}

func (ac AddonsConfigAttributes) MonetizationConfig() terra.ListValue[AddonsConfigMonetizationConfigAttributes] {
	return terra.ReferenceAsList[AddonsConfigMonetizationConfigAttributes](ac.ref.Append("monetization_config"))
}

type AddonsConfigAdvancedApiOpsConfigAttributes struct {
	ref terra.Reference
}

func (aaoc AddonsConfigAdvancedApiOpsConfigAttributes) InternalRef() (terra.Reference, error) {
	return aaoc.ref, nil
}

func (aaoc AddonsConfigAdvancedApiOpsConfigAttributes) InternalWithRef(ref terra.Reference) AddonsConfigAdvancedApiOpsConfigAttributes {
	return AddonsConfigAdvancedApiOpsConfigAttributes{ref: ref}
}

func (aaoc AddonsConfigAdvancedApiOpsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aaoc.ref.InternalTokens()
}

func (aaoc AddonsConfigAdvancedApiOpsConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(aaoc.ref.Append("enabled"))
}

type AddonsConfigApiSecurityConfigAttributes struct {
	ref terra.Reference
}

func (asc AddonsConfigApiSecurityConfigAttributes) InternalRef() (terra.Reference, error) {
	return asc.ref, nil
}

func (asc AddonsConfigApiSecurityConfigAttributes) InternalWithRef(ref terra.Reference) AddonsConfigApiSecurityConfigAttributes {
	return AddonsConfigApiSecurityConfigAttributes{ref: ref}
}

func (asc AddonsConfigApiSecurityConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return asc.ref.InternalTokens()
}

func (asc AddonsConfigApiSecurityConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(asc.ref.Append("enabled"))
}

func (asc AddonsConfigApiSecurityConfigAttributes) ExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("expires_at"))
}

type AddonsConfigConnectorsPlatformConfigAttributes struct {
	ref terra.Reference
}

func (cpc AddonsConfigConnectorsPlatformConfigAttributes) InternalRef() (terra.Reference, error) {
	return cpc.ref, nil
}

func (cpc AddonsConfigConnectorsPlatformConfigAttributes) InternalWithRef(ref terra.Reference) AddonsConfigConnectorsPlatformConfigAttributes {
	return AddonsConfigConnectorsPlatformConfigAttributes{ref: ref}
}

func (cpc AddonsConfigConnectorsPlatformConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cpc.ref.InternalTokens()
}

func (cpc AddonsConfigConnectorsPlatformConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cpc.ref.Append("enabled"))
}

func (cpc AddonsConfigConnectorsPlatformConfigAttributes) ExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("expires_at"))
}

type AddonsConfigIntegrationConfigAttributes struct {
	ref terra.Reference
}

func (ic AddonsConfigIntegrationConfigAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic AddonsConfigIntegrationConfigAttributes) InternalWithRef(ref terra.Reference) AddonsConfigIntegrationConfigAttributes {
	return AddonsConfigIntegrationConfigAttributes{ref: ref}
}

func (ic AddonsConfigIntegrationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic AddonsConfigIntegrationConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("enabled"))
}

type AddonsConfigMonetizationConfigAttributes struct {
	ref terra.Reference
}

func (mc AddonsConfigMonetizationConfigAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc AddonsConfigMonetizationConfigAttributes) InternalWithRef(ref terra.Reference) AddonsConfigMonetizationConfigAttributes {
	return AddonsConfigMonetizationConfigAttributes{ref: ref}
}

func (mc AddonsConfigMonetizationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc AddonsConfigMonetizationConfigAttributes) Enabled() terra.BoolValue {
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
	AdvancedApiOpsConfig     []AddonsConfigAdvancedApiOpsConfigState     `json:"advanced_api_ops_config"`
	ApiSecurityConfig        []AddonsConfigApiSecurityConfigState        `json:"api_security_config"`
	ConnectorsPlatformConfig []AddonsConfigConnectorsPlatformConfigState `json:"connectors_platform_config"`
	IntegrationConfig        []AddonsConfigIntegrationConfigState        `json:"integration_config"`
	MonetizationConfig       []AddonsConfigMonetizationConfigState       `json:"monetization_config"`
}

type AddonsConfigAdvancedApiOpsConfigState struct {
	Enabled bool `json:"enabled"`
}

type AddonsConfigApiSecurityConfigState struct {
	Enabled   bool   `json:"enabled"`
	ExpiresAt string `json:"expires_at"`
}

type AddonsConfigConnectorsPlatformConfigState struct {
	Enabled   bool   `json:"enabled"`
	ExpiresAt string `json:"expires_at"`
}

type AddonsConfigIntegrationConfigState struct {
	Enabled bool `json:"enabled"`
}

type AddonsConfigMonetizationConfigState struct {
	Enabled bool `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
