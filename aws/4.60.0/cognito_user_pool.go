// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitouserpool "github.com/golingon/terraproviders/aws/4.60.0/cognitouserpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCognitoUserPool(name string, args CognitoUserPoolArgs) *CognitoUserPool {
	return &CognitoUserPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoUserPool)(nil)

type CognitoUserPool struct {
	Name  string
	Args  CognitoUserPoolArgs
	state *cognitoUserPoolState
}

func (cup *CognitoUserPool) Type() string {
	return "aws_cognito_user_pool"
}

func (cup *CognitoUserPool) LocalName() string {
	return cup.Name
}

func (cup *CognitoUserPool) Configuration() interface{} {
	return cup.Args
}

func (cup *CognitoUserPool) Attributes() cognitoUserPoolAttributes {
	return cognitoUserPoolAttributes{ref: terra.ReferenceResource(cup)}
}

func (cup *CognitoUserPool) ImportState(av io.Reader) error {
	cup.state = &cognitoUserPoolState{}
	if err := json.NewDecoder(av).Decode(cup.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cup.Type(), cup.LocalName(), err)
	}
	return nil
}

func (cup *CognitoUserPool) State() (*cognitoUserPoolState, bool) {
	return cup.state, cup.state != nil
}

func (cup *CognitoUserPool) StateMust() *cognitoUserPoolState {
	if cup.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cup.Type(), cup.LocalName()))
	}
	return cup.state
}

func (cup *CognitoUserPool) DependOn() terra.Reference {
	return terra.ReferenceResource(cup)
}

type CognitoUserPoolArgs struct {
	// AliasAttributes: set of string, optional
	AliasAttributes terra.SetValue[terra.StringValue] `hcl:"alias_attributes,attr"`
	// AutoVerifiedAttributes: set of string, optional
	AutoVerifiedAttributes terra.SetValue[terra.StringValue] `hcl:"auto_verified_attributes,attr"`
	// DeletionProtection: string, optional
	DeletionProtection terra.StringValue `hcl:"deletion_protection,attr"`
	// EmailVerificationMessage: string, optional
	EmailVerificationMessage terra.StringValue `hcl:"email_verification_message,attr"`
	// EmailVerificationSubject: string, optional
	EmailVerificationSubject terra.StringValue `hcl:"email_verification_subject,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MfaConfiguration: string, optional
	MfaConfiguration terra.StringValue `hcl:"mfa_configuration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SmsAuthenticationMessage: string, optional
	SmsAuthenticationMessage terra.StringValue `hcl:"sms_authentication_message,attr"`
	// SmsVerificationMessage: string, optional
	SmsVerificationMessage terra.StringValue `hcl:"sms_verification_message,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UsernameAttributes: set of string, optional
	UsernameAttributes terra.SetValue[terra.StringValue] `hcl:"username_attributes,attr"`
	// AccountRecoverySetting: optional
	AccountRecoverySetting *cognitouserpool.AccountRecoverySetting `hcl:"account_recovery_setting,block"`
	// AdminCreateUserConfig: optional
	AdminCreateUserConfig *cognitouserpool.AdminCreateUserConfig `hcl:"admin_create_user_config,block"`
	// DeviceConfiguration: optional
	DeviceConfiguration *cognitouserpool.DeviceConfiguration `hcl:"device_configuration,block"`
	// EmailConfiguration: optional
	EmailConfiguration *cognitouserpool.EmailConfiguration `hcl:"email_configuration,block"`
	// LambdaConfig: optional
	LambdaConfig *cognitouserpool.LambdaConfig `hcl:"lambda_config,block"`
	// PasswordPolicy: optional
	PasswordPolicy *cognitouserpool.PasswordPolicy `hcl:"password_policy,block"`
	// Schema: min=0,max=50
	Schema []cognitouserpool.Schema `hcl:"schema,block" validate:"min=0,max=50"`
	// SmsConfiguration: optional
	SmsConfiguration *cognitouserpool.SmsConfiguration `hcl:"sms_configuration,block"`
	// SoftwareTokenMfaConfiguration: optional
	SoftwareTokenMfaConfiguration *cognitouserpool.SoftwareTokenMfaConfiguration `hcl:"software_token_mfa_configuration,block"`
	// UserAttributeUpdateSettings: optional
	UserAttributeUpdateSettings *cognitouserpool.UserAttributeUpdateSettings `hcl:"user_attribute_update_settings,block"`
	// UserPoolAddOns: optional
	UserPoolAddOns *cognitouserpool.UserPoolAddOns `hcl:"user_pool_add_ons,block"`
	// UsernameConfiguration: optional
	UsernameConfiguration *cognitouserpool.UsernameConfiguration `hcl:"username_configuration,block"`
	// VerificationMessageTemplate: optional
	VerificationMessageTemplate *cognitouserpool.VerificationMessageTemplate `hcl:"verification_message_template,block"`
	// DependsOn contains resources that CognitoUserPool depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cognitoUserPoolAttributes struct {
	ref terra.Reference
}

func (cup cognitoUserPoolAttributes) AliasAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cup.ref.Append("alias_attributes"))
}

func (cup cognitoUserPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("arn"))
}

func (cup cognitoUserPoolAttributes) AutoVerifiedAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cup.ref.Append("auto_verified_attributes"))
}

func (cup cognitoUserPoolAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("creation_date"))
}

func (cup cognitoUserPoolAttributes) CustomDomain() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("custom_domain"))
}

func (cup cognitoUserPoolAttributes) DeletionProtection() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("deletion_protection"))
}

func (cup cognitoUserPoolAttributes) Domain() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("domain"))
}

func (cup cognitoUserPoolAttributes) EmailVerificationMessage() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("email_verification_message"))
}

func (cup cognitoUserPoolAttributes) EmailVerificationSubject() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("email_verification_subject"))
}

func (cup cognitoUserPoolAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("endpoint"))
}

func (cup cognitoUserPoolAttributes) EstimatedNumberOfUsers() terra.NumberValue {
	return terra.ReferenceNumber(cup.ref.Append("estimated_number_of_users"))
}

func (cup cognitoUserPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("id"))
}

func (cup cognitoUserPoolAttributes) LastModifiedDate() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("last_modified_date"))
}

func (cup cognitoUserPoolAttributes) MfaConfiguration() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("mfa_configuration"))
}

func (cup cognitoUserPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("name"))
}

func (cup cognitoUserPoolAttributes) SmsAuthenticationMessage() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("sms_authentication_message"))
}

func (cup cognitoUserPoolAttributes) SmsVerificationMessage() terra.StringValue {
	return terra.ReferenceString(cup.ref.Append("sms_verification_message"))
}

func (cup cognitoUserPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cup.ref.Append("tags"))
}

func (cup cognitoUserPoolAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cup.ref.Append("tags_all"))
}

func (cup cognitoUserPoolAttributes) UsernameAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cup.ref.Append("username_attributes"))
}

func (cup cognitoUserPoolAttributes) AccountRecoverySetting() terra.ListValue[cognitouserpool.AccountRecoverySettingAttributes] {
	return terra.ReferenceList[cognitouserpool.AccountRecoverySettingAttributes](cup.ref.Append("account_recovery_setting"))
}

func (cup cognitoUserPoolAttributes) AdminCreateUserConfig() terra.ListValue[cognitouserpool.AdminCreateUserConfigAttributes] {
	return terra.ReferenceList[cognitouserpool.AdminCreateUserConfigAttributes](cup.ref.Append("admin_create_user_config"))
}

func (cup cognitoUserPoolAttributes) DeviceConfiguration() terra.ListValue[cognitouserpool.DeviceConfigurationAttributes] {
	return terra.ReferenceList[cognitouserpool.DeviceConfigurationAttributes](cup.ref.Append("device_configuration"))
}

func (cup cognitoUserPoolAttributes) EmailConfiguration() terra.ListValue[cognitouserpool.EmailConfigurationAttributes] {
	return terra.ReferenceList[cognitouserpool.EmailConfigurationAttributes](cup.ref.Append("email_configuration"))
}

func (cup cognitoUserPoolAttributes) LambdaConfig() terra.ListValue[cognitouserpool.LambdaConfigAttributes] {
	return terra.ReferenceList[cognitouserpool.LambdaConfigAttributes](cup.ref.Append("lambda_config"))
}

func (cup cognitoUserPoolAttributes) PasswordPolicy() terra.ListValue[cognitouserpool.PasswordPolicyAttributes] {
	return terra.ReferenceList[cognitouserpool.PasswordPolicyAttributes](cup.ref.Append("password_policy"))
}

func (cup cognitoUserPoolAttributes) Schema() terra.SetValue[cognitouserpool.SchemaAttributes] {
	return terra.ReferenceSet[cognitouserpool.SchemaAttributes](cup.ref.Append("schema"))
}

func (cup cognitoUserPoolAttributes) SmsConfiguration() terra.ListValue[cognitouserpool.SmsConfigurationAttributes] {
	return terra.ReferenceList[cognitouserpool.SmsConfigurationAttributes](cup.ref.Append("sms_configuration"))
}

func (cup cognitoUserPoolAttributes) SoftwareTokenMfaConfiguration() terra.ListValue[cognitouserpool.SoftwareTokenMfaConfigurationAttributes] {
	return terra.ReferenceList[cognitouserpool.SoftwareTokenMfaConfigurationAttributes](cup.ref.Append("software_token_mfa_configuration"))
}

func (cup cognitoUserPoolAttributes) UserAttributeUpdateSettings() terra.ListValue[cognitouserpool.UserAttributeUpdateSettingsAttributes] {
	return terra.ReferenceList[cognitouserpool.UserAttributeUpdateSettingsAttributes](cup.ref.Append("user_attribute_update_settings"))
}

func (cup cognitoUserPoolAttributes) UserPoolAddOns() terra.ListValue[cognitouserpool.UserPoolAddOnsAttributes] {
	return terra.ReferenceList[cognitouserpool.UserPoolAddOnsAttributes](cup.ref.Append("user_pool_add_ons"))
}

func (cup cognitoUserPoolAttributes) UsernameConfiguration() terra.ListValue[cognitouserpool.UsernameConfigurationAttributes] {
	return terra.ReferenceList[cognitouserpool.UsernameConfigurationAttributes](cup.ref.Append("username_configuration"))
}

func (cup cognitoUserPoolAttributes) VerificationMessageTemplate() terra.ListValue[cognitouserpool.VerificationMessageTemplateAttributes] {
	return terra.ReferenceList[cognitouserpool.VerificationMessageTemplateAttributes](cup.ref.Append("verification_message_template"))
}

type cognitoUserPoolState struct {
	AliasAttributes               []string                                             `json:"alias_attributes"`
	Arn                           string                                               `json:"arn"`
	AutoVerifiedAttributes        []string                                             `json:"auto_verified_attributes"`
	CreationDate                  string                                               `json:"creation_date"`
	CustomDomain                  string                                               `json:"custom_domain"`
	DeletionProtection            string                                               `json:"deletion_protection"`
	Domain                        string                                               `json:"domain"`
	EmailVerificationMessage      string                                               `json:"email_verification_message"`
	EmailVerificationSubject      string                                               `json:"email_verification_subject"`
	Endpoint                      string                                               `json:"endpoint"`
	EstimatedNumberOfUsers        float64                                              `json:"estimated_number_of_users"`
	Id                            string                                               `json:"id"`
	LastModifiedDate              string                                               `json:"last_modified_date"`
	MfaConfiguration              string                                               `json:"mfa_configuration"`
	Name                          string                                               `json:"name"`
	SmsAuthenticationMessage      string                                               `json:"sms_authentication_message"`
	SmsVerificationMessage        string                                               `json:"sms_verification_message"`
	Tags                          map[string]string                                    `json:"tags"`
	TagsAll                       map[string]string                                    `json:"tags_all"`
	UsernameAttributes            []string                                             `json:"username_attributes"`
	AccountRecoverySetting        []cognitouserpool.AccountRecoverySettingState        `json:"account_recovery_setting"`
	AdminCreateUserConfig         []cognitouserpool.AdminCreateUserConfigState         `json:"admin_create_user_config"`
	DeviceConfiguration           []cognitouserpool.DeviceConfigurationState           `json:"device_configuration"`
	EmailConfiguration            []cognitouserpool.EmailConfigurationState            `json:"email_configuration"`
	LambdaConfig                  []cognitouserpool.LambdaConfigState                  `json:"lambda_config"`
	PasswordPolicy                []cognitouserpool.PasswordPolicyState                `json:"password_policy"`
	Schema                        []cognitouserpool.SchemaState                        `json:"schema"`
	SmsConfiguration              []cognitouserpool.SmsConfigurationState              `json:"sms_configuration"`
	SoftwareTokenMfaConfiguration []cognitouserpool.SoftwareTokenMfaConfigurationState `json:"software_token_mfa_configuration"`
	UserAttributeUpdateSettings   []cognitouserpool.UserAttributeUpdateSettingsState   `json:"user_attribute_update_settings"`
	UserPoolAddOns                []cognitouserpool.UserPoolAddOnsState                `json:"user_pool_add_ons"`
	UsernameConfiguration         []cognitouserpool.UsernameConfigurationState         `json:"username_configuration"`
	VerificationMessageTemplate   []cognitouserpool.VerificationMessageTemplateState   `json:"verification_message_template"`
}
