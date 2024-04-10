// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cognitouserpool "github.com/golingon/terraproviders/aws/5.44.0/cognitouserpool"
	"io"
)

// NewCognitoUserPool creates a new instance of [CognitoUserPool].
func NewCognitoUserPool(name string, args CognitoUserPoolArgs) *CognitoUserPool {
	return &CognitoUserPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoUserPool)(nil)

// CognitoUserPool represents the Terraform resource aws_cognito_user_pool.
type CognitoUserPool struct {
	Name      string
	Args      CognitoUserPoolArgs
	state     *cognitoUserPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoUserPool].
func (cup *CognitoUserPool) Type() string {
	return "aws_cognito_user_pool"
}

// LocalName returns the local name for [CognitoUserPool].
func (cup *CognitoUserPool) LocalName() string {
	return cup.Name
}

// Configuration returns the configuration (args) for [CognitoUserPool].
func (cup *CognitoUserPool) Configuration() interface{} {
	return cup.Args
}

// DependOn is used for other resources to depend on [CognitoUserPool].
func (cup *CognitoUserPool) DependOn() terra.Reference {
	return terra.ReferenceResource(cup)
}

// Dependencies returns the list of resources [CognitoUserPool] depends_on.
func (cup *CognitoUserPool) Dependencies() terra.Dependencies {
	return cup.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoUserPool].
func (cup *CognitoUserPool) LifecycleManagement() *terra.Lifecycle {
	return cup.Lifecycle
}

// Attributes returns the attributes for [CognitoUserPool].
func (cup *CognitoUserPool) Attributes() cognitoUserPoolAttributes {
	return cognitoUserPoolAttributes{ref: terra.ReferenceResource(cup)}
}

// ImportState imports the given attribute values into [CognitoUserPool]'s state.
func (cup *CognitoUserPool) ImportState(av io.Reader) error {
	cup.state = &cognitoUserPoolState{}
	if err := json.NewDecoder(av).Decode(cup.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cup.Type(), cup.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoUserPool] has state.
func (cup *CognitoUserPool) State() (*cognitoUserPoolState, bool) {
	return cup.state, cup.state != nil
}

// StateMust returns the state for [CognitoUserPool]. Panics if the state is nil.
func (cup *CognitoUserPool) StateMust() *cognitoUserPoolState {
	if cup.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cup.Type(), cup.LocalName()))
	}
	return cup.state
}

// CognitoUserPoolArgs contains the configurations for aws_cognito_user_pool.
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
}
type cognitoUserPoolAttributes struct {
	ref terra.Reference
}

// AliasAttributes returns a reference to field alias_attributes of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) AliasAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cup.ref.Append("alias_attributes"))
}

// Arn returns a reference to field arn of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("arn"))
}

// AutoVerifiedAttributes returns a reference to field auto_verified_attributes of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) AutoVerifiedAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cup.ref.Append("auto_verified_attributes"))
}

// CreationDate returns a reference to field creation_date of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("creation_date"))
}

// CustomDomain returns a reference to field custom_domain of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) CustomDomain() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("custom_domain"))
}

// DeletionProtection returns a reference to field deletion_protection of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) DeletionProtection() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("deletion_protection"))
}

// Domain returns a reference to field domain of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("domain"))
}

// EmailVerificationMessage returns a reference to field email_verification_message of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) EmailVerificationMessage() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("email_verification_message"))
}

// EmailVerificationSubject returns a reference to field email_verification_subject of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) EmailVerificationSubject() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("email_verification_subject"))
}

// Endpoint returns a reference to field endpoint of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("endpoint"))
}

// EstimatedNumberOfUsers returns a reference to field estimated_number_of_users of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) EstimatedNumberOfUsers() terra.NumberValue {
	return terra.ReferenceAsNumber(cup.ref.Append("estimated_number_of_users"))
}

// Id returns a reference to field id of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("id"))
}

// LastModifiedDate returns a reference to field last_modified_date of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) LastModifiedDate() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("last_modified_date"))
}

// MfaConfiguration returns a reference to field mfa_configuration of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) MfaConfiguration() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("mfa_configuration"))
}

// Name returns a reference to field name of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("name"))
}

// SmsAuthenticationMessage returns a reference to field sms_authentication_message of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) SmsAuthenticationMessage() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("sms_authentication_message"))
}

// SmsVerificationMessage returns a reference to field sms_verification_message of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) SmsVerificationMessage() terra.StringValue {
	return terra.ReferenceAsString(cup.ref.Append("sms_verification_message"))
}

// Tags returns a reference to field tags of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cup.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cup.ref.Append("tags_all"))
}

// UsernameAttributes returns a reference to field username_attributes of aws_cognito_user_pool.
func (cup cognitoUserPoolAttributes) UsernameAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cup.ref.Append("username_attributes"))
}

func (cup cognitoUserPoolAttributes) AccountRecoverySetting() terra.ListValue[cognitouserpool.AccountRecoverySettingAttributes] {
	return terra.ReferenceAsList[cognitouserpool.AccountRecoverySettingAttributes](cup.ref.Append("account_recovery_setting"))
}

func (cup cognitoUserPoolAttributes) AdminCreateUserConfig() terra.ListValue[cognitouserpool.AdminCreateUserConfigAttributes] {
	return terra.ReferenceAsList[cognitouserpool.AdminCreateUserConfigAttributes](cup.ref.Append("admin_create_user_config"))
}

func (cup cognitoUserPoolAttributes) DeviceConfiguration() terra.ListValue[cognitouserpool.DeviceConfigurationAttributes] {
	return terra.ReferenceAsList[cognitouserpool.DeviceConfigurationAttributes](cup.ref.Append("device_configuration"))
}

func (cup cognitoUserPoolAttributes) EmailConfiguration() terra.ListValue[cognitouserpool.EmailConfigurationAttributes] {
	return terra.ReferenceAsList[cognitouserpool.EmailConfigurationAttributes](cup.ref.Append("email_configuration"))
}

func (cup cognitoUserPoolAttributes) LambdaConfig() terra.ListValue[cognitouserpool.LambdaConfigAttributes] {
	return terra.ReferenceAsList[cognitouserpool.LambdaConfigAttributes](cup.ref.Append("lambda_config"))
}

func (cup cognitoUserPoolAttributes) PasswordPolicy() terra.ListValue[cognitouserpool.PasswordPolicyAttributes] {
	return terra.ReferenceAsList[cognitouserpool.PasswordPolicyAttributes](cup.ref.Append("password_policy"))
}

func (cup cognitoUserPoolAttributes) Schema() terra.SetValue[cognitouserpool.SchemaAttributes] {
	return terra.ReferenceAsSet[cognitouserpool.SchemaAttributes](cup.ref.Append("schema"))
}

func (cup cognitoUserPoolAttributes) SmsConfiguration() terra.ListValue[cognitouserpool.SmsConfigurationAttributes] {
	return terra.ReferenceAsList[cognitouserpool.SmsConfigurationAttributes](cup.ref.Append("sms_configuration"))
}

func (cup cognitoUserPoolAttributes) SoftwareTokenMfaConfiguration() terra.ListValue[cognitouserpool.SoftwareTokenMfaConfigurationAttributes] {
	return terra.ReferenceAsList[cognitouserpool.SoftwareTokenMfaConfigurationAttributes](cup.ref.Append("software_token_mfa_configuration"))
}

func (cup cognitoUserPoolAttributes) UserAttributeUpdateSettings() terra.ListValue[cognitouserpool.UserAttributeUpdateSettingsAttributes] {
	return terra.ReferenceAsList[cognitouserpool.UserAttributeUpdateSettingsAttributes](cup.ref.Append("user_attribute_update_settings"))
}

func (cup cognitoUserPoolAttributes) UserPoolAddOns() terra.ListValue[cognitouserpool.UserPoolAddOnsAttributes] {
	return terra.ReferenceAsList[cognitouserpool.UserPoolAddOnsAttributes](cup.ref.Append("user_pool_add_ons"))
}

func (cup cognitoUserPoolAttributes) UsernameConfiguration() terra.ListValue[cognitouserpool.UsernameConfigurationAttributes] {
	return terra.ReferenceAsList[cognitouserpool.UsernameConfigurationAttributes](cup.ref.Append("username_configuration"))
}

func (cup cognitoUserPoolAttributes) VerificationMessageTemplate() terra.ListValue[cognitouserpool.VerificationMessageTemplateAttributes] {
	return terra.ReferenceAsList[cognitouserpool.VerificationMessageTemplateAttributes](cup.ref.Append("verification_message_template"))
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
