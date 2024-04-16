// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securitylake_subscriber_notification

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Configuration struct {
	// ConfigurationHttpsNotificationConfiguration: min=0
	HttpsNotificationConfiguration []ConfigurationHttpsNotificationConfiguration `hcl:"https_notification_configuration,block" validate:"min=0"`
	// ConfigurationSqsNotificationConfiguration: min=0
	SqsNotificationConfiguration []ConfigurationSqsNotificationConfiguration `hcl:"sqs_notification_configuration,block" validate:"min=0"`
}

type ConfigurationHttpsNotificationConfiguration struct {
	// AuthorizationApiKeyName: string, optional
	AuthorizationApiKeyName terra.StringValue `hcl:"authorization_api_key_name,attr"`
	// AuthorizationApiKeyValue: string, optional
	AuthorizationApiKeyValue terra.StringValue `hcl:"authorization_api_key_value,attr"`
	// Endpoint: string, optional
	Endpoint terra.StringValue `hcl:"endpoint,attr"`
	// HttpMethod: string, optional
	HttpMethod terra.StringValue `hcl:"http_method,attr"`
	// TargetRoleArn: string, optional
	TargetRoleArn terra.StringValue `hcl:"target_role_arn,attr"`
}

type ConfigurationSqsNotificationConfiguration struct{}

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

func (c ConfigurationAttributes) HttpsNotificationConfiguration() terra.ListValue[ConfigurationHttpsNotificationConfigurationAttributes] {
	return terra.ReferenceAsList[ConfigurationHttpsNotificationConfigurationAttributes](c.ref.Append("https_notification_configuration"))
}

func (c ConfigurationAttributes) SqsNotificationConfiguration() terra.ListValue[ConfigurationSqsNotificationConfigurationAttributes] {
	return terra.ReferenceAsList[ConfigurationSqsNotificationConfigurationAttributes](c.ref.Append("sqs_notification_configuration"))
}

type ConfigurationHttpsNotificationConfigurationAttributes struct {
	ref terra.Reference
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return hnc.ref, nil
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationHttpsNotificationConfigurationAttributes {
	return ConfigurationHttpsNotificationConfigurationAttributes{ref: ref}
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hnc.ref.InternalTokens()
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) AuthorizationApiKeyName() terra.StringValue {
	return terra.ReferenceAsString(hnc.ref.Append("authorization_api_key_name"))
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) AuthorizationApiKeyValue() terra.StringValue {
	return terra.ReferenceAsString(hnc.ref.Append("authorization_api_key_value"))
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(hnc.ref.Append("endpoint"))
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(hnc.ref.Append("http_method"))
}

func (hnc ConfigurationHttpsNotificationConfigurationAttributes) TargetRoleArn() terra.StringValue {
	return terra.ReferenceAsString(hnc.ref.Append("target_role_arn"))
}

type ConfigurationSqsNotificationConfigurationAttributes struct {
	ref terra.Reference
}

func (snc ConfigurationSqsNotificationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return snc.ref, nil
}

func (snc ConfigurationSqsNotificationConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationSqsNotificationConfigurationAttributes {
	return ConfigurationSqsNotificationConfigurationAttributes{ref: ref}
}

func (snc ConfigurationSqsNotificationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return snc.ref.InternalTokens()
}

type ConfigurationState struct {
	HttpsNotificationConfiguration []ConfigurationHttpsNotificationConfigurationState `json:"https_notification_configuration"`
	SqsNotificationConfiguration   []ConfigurationSqsNotificationConfigurationState   `json:"sqs_notification_configuration"`
}

type ConfigurationHttpsNotificationConfigurationState struct {
	AuthorizationApiKeyName  string `json:"authorization_api_key_name"`
	AuthorizationApiKeyValue string `json:"authorization_api_key_value"`
	Endpoint                 string `json:"endpoint"`
	HttpMethod               string `json:"http_method"`
	TargetRoleArn            string `json:"target_role_arn"`
}

type ConfigurationSqsNotificationConfigurationState struct{}
