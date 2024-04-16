// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securitylake_custom_log_source

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Configuration struct {
	// ConfigurationCrawlerConfiguration: min=0
	CrawlerConfiguration []ConfigurationCrawlerConfiguration `hcl:"crawler_configuration,block" validate:"min=0"`
	// ConfigurationProviderIdentity: min=0
	ProviderIdentity []ConfigurationProviderIdentity `hcl:"provider_identity,block" validate:"min=0"`
}

type ConfigurationCrawlerConfiguration struct {
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type ConfigurationProviderIdentity struct {
	// ExternalId: string, required
	ExternalId terra.StringValue `hcl:"external_id,attr" validate:"required"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
}

type AttributesAttributes struct {
	ref terra.Reference
}

func (a AttributesAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AttributesAttributes) InternalWithRef(ref terra.Reference) AttributesAttributes {
	return AttributesAttributes{ref: ref}
}

func (a AttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AttributesAttributes) CrawlerArn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("crawler_arn"))
}

func (a AttributesAttributes) DatabaseArn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("database_arn"))
}

func (a AttributesAttributes) TableArn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("table_arn"))
}

type ProviderDetailsAttributes struct {
	ref terra.Reference
}

func (pd ProviderDetailsAttributes) InternalRef() (terra.Reference, error) {
	return pd.ref, nil
}

func (pd ProviderDetailsAttributes) InternalWithRef(ref terra.Reference) ProviderDetailsAttributes {
	return ProviderDetailsAttributes{ref: ref}
}

func (pd ProviderDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pd.ref.InternalTokens()
}

func (pd ProviderDetailsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("location"))
}

func (pd ProviderDetailsAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("role_arn"))
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) CrawlerConfiguration() terra.ListValue[ConfigurationCrawlerConfigurationAttributes] {
	return terra.ReferenceAsList[ConfigurationCrawlerConfigurationAttributes](c.ref.Append("crawler_configuration"))
}

func (c ConfigurationAttributes) ProviderIdentity() terra.ListValue[ConfigurationProviderIdentityAttributes] {
	return terra.ReferenceAsList[ConfigurationProviderIdentityAttributes](c.ref.Append("provider_identity"))
}

type ConfigurationCrawlerConfigurationAttributes struct {
	ref terra.Reference
}

func (cc ConfigurationCrawlerConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc ConfigurationCrawlerConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationCrawlerConfigurationAttributes {
	return ConfigurationCrawlerConfigurationAttributes{ref: ref}
}

func (cc ConfigurationCrawlerConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc ConfigurationCrawlerConfigurationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("role_arn"))
}

type ConfigurationProviderIdentityAttributes struct {
	ref terra.Reference
}

func (pi ConfigurationProviderIdentityAttributes) InternalRef() (terra.Reference, error) {
	return pi.ref, nil
}

func (pi ConfigurationProviderIdentityAttributes) InternalWithRef(ref terra.Reference) ConfigurationProviderIdentityAttributes {
	return ConfigurationProviderIdentityAttributes{ref: ref}
}

func (pi ConfigurationProviderIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pi.ref.InternalTokens()
}

func (pi ConfigurationProviderIdentityAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("external_id"))
}

func (pi ConfigurationProviderIdentityAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("principal"))
}

type AttributesState struct {
	CrawlerArn  string `json:"crawler_arn"`
	DatabaseArn string `json:"database_arn"`
	TableArn    string `json:"table_arn"`
}

type ProviderDetailsState struct {
	Location string `json:"location"`
	RoleArn  string `json:"role_arn"`
}

type ConfigurationState struct {
	CrawlerConfiguration []ConfigurationCrawlerConfigurationState `json:"crawler_configuration"`
	ProviderIdentity     []ConfigurationProviderIdentityState     `json:"provider_identity"`
}

type ConfigurationCrawlerConfigurationState struct {
	RoleArn string `json:"role_arn"`
}

type ConfigurationProviderIdentityState struct {
	ExternalId string `json:"external_id"`
	Principal  string `json:"principal"`
}
