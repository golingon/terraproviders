// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cognitouserpool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccountRecoverySetting struct {
	// RecoveryMechanism: min=0,max=2
	RecoveryMechanism []RecoveryMechanism `hcl:"recovery_mechanism,block" validate:"min=0,max=2"`
}

type RecoveryMechanism struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
}

type AdminCreateUserConfig struct {
	// AllowAdminCreateUserOnly: bool, optional
	AllowAdminCreateUserOnly terra.BoolValue `hcl:"allow_admin_create_user_only,attr"`
	// InviteMessageTemplate: optional
	InviteMessageTemplate *InviteMessageTemplate `hcl:"invite_message_template,block"`
}

type InviteMessageTemplate struct {
	// EmailMessage: string, optional
	EmailMessage terra.StringValue `hcl:"email_message,attr"`
	// EmailSubject: string, optional
	EmailSubject terra.StringValue `hcl:"email_subject,attr"`
	// SmsMessage: string, optional
	SmsMessage terra.StringValue `hcl:"sms_message,attr"`
}

type DeviceConfiguration struct {
	// ChallengeRequiredOnNewDevice: bool, optional
	ChallengeRequiredOnNewDevice terra.BoolValue `hcl:"challenge_required_on_new_device,attr"`
	// DeviceOnlyRememberedOnUserPrompt: bool, optional
	DeviceOnlyRememberedOnUserPrompt terra.BoolValue `hcl:"device_only_remembered_on_user_prompt,attr"`
}

type EmailConfiguration struct {
	// ConfigurationSet: string, optional
	ConfigurationSet terra.StringValue `hcl:"configuration_set,attr"`
	// EmailSendingAccount: string, optional
	EmailSendingAccount terra.StringValue `hcl:"email_sending_account,attr"`
	// FromEmailAddress: string, optional
	FromEmailAddress terra.StringValue `hcl:"from_email_address,attr"`
	// ReplyToEmailAddress: string, optional
	ReplyToEmailAddress terra.StringValue `hcl:"reply_to_email_address,attr"`
	// SourceArn: string, optional
	SourceArn terra.StringValue `hcl:"source_arn,attr"`
}

type LambdaConfig struct {
	// CreateAuthChallenge: string, optional
	CreateAuthChallenge terra.StringValue `hcl:"create_auth_challenge,attr"`
	// CustomMessage: string, optional
	CustomMessage terra.StringValue `hcl:"custom_message,attr"`
	// DefineAuthChallenge: string, optional
	DefineAuthChallenge terra.StringValue `hcl:"define_auth_challenge,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PostAuthentication: string, optional
	PostAuthentication terra.StringValue `hcl:"post_authentication,attr"`
	// PostConfirmation: string, optional
	PostConfirmation terra.StringValue `hcl:"post_confirmation,attr"`
	// PreAuthentication: string, optional
	PreAuthentication terra.StringValue `hcl:"pre_authentication,attr"`
	// PreSignUp: string, optional
	PreSignUp terra.StringValue `hcl:"pre_sign_up,attr"`
	// PreTokenGeneration: string, optional
	PreTokenGeneration terra.StringValue `hcl:"pre_token_generation,attr"`
	// UserMigration: string, optional
	UserMigration terra.StringValue `hcl:"user_migration,attr"`
	// VerifyAuthChallengeResponse: string, optional
	VerifyAuthChallengeResponse terra.StringValue `hcl:"verify_auth_challenge_response,attr"`
	// CustomEmailSender: optional
	CustomEmailSender *CustomEmailSender `hcl:"custom_email_sender,block"`
	// CustomSmsSender: optional
	CustomSmsSender *CustomSmsSender `hcl:"custom_sms_sender,block"`
}

type CustomEmailSender struct {
	// LambdaArn: string, required
	LambdaArn terra.StringValue `hcl:"lambda_arn,attr" validate:"required"`
	// LambdaVersion: string, required
	LambdaVersion terra.StringValue `hcl:"lambda_version,attr" validate:"required"`
}

type CustomSmsSender struct {
	// LambdaArn: string, required
	LambdaArn terra.StringValue `hcl:"lambda_arn,attr" validate:"required"`
	// LambdaVersion: string, required
	LambdaVersion terra.StringValue `hcl:"lambda_version,attr" validate:"required"`
}

type PasswordPolicy struct {
	// MinimumLength: number, optional
	MinimumLength terra.NumberValue `hcl:"minimum_length,attr"`
	// RequireLowercase: bool, optional
	RequireLowercase terra.BoolValue `hcl:"require_lowercase,attr"`
	// RequireNumbers: bool, optional
	RequireNumbers terra.BoolValue `hcl:"require_numbers,attr"`
	// RequireSymbols: bool, optional
	RequireSymbols terra.BoolValue `hcl:"require_symbols,attr"`
	// RequireUppercase: bool, optional
	RequireUppercase terra.BoolValue `hcl:"require_uppercase,attr"`
	// TemporaryPasswordValidityDays: number, optional
	TemporaryPasswordValidityDays terra.NumberValue `hcl:"temporary_password_validity_days,attr"`
}

type Schema struct {
	// AttributeDataType: string, required
	AttributeDataType terra.StringValue `hcl:"attribute_data_type,attr" validate:"required"`
	// DeveloperOnlyAttribute: bool, optional
	DeveloperOnlyAttribute terra.BoolValue `hcl:"developer_only_attribute,attr"`
	// Mutable: bool, optional
	Mutable terra.BoolValue `hcl:"mutable,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Required: bool, optional
	Required terra.BoolValue `hcl:"required,attr"`
	// NumberAttributeConstraints: optional
	NumberAttributeConstraints *NumberAttributeConstraints `hcl:"number_attribute_constraints,block"`
	// StringAttributeConstraints: optional
	StringAttributeConstraints *StringAttributeConstraints `hcl:"string_attribute_constraints,block"`
}

type NumberAttributeConstraints struct {
	// MaxValue: string, optional
	MaxValue terra.StringValue `hcl:"max_value,attr"`
	// MinValue: string, optional
	MinValue terra.StringValue `hcl:"min_value,attr"`
}

type StringAttributeConstraints struct {
	// MaxLength: string, optional
	MaxLength terra.StringValue `hcl:"max_length,attr"`
	// MinLength: string, optional
	MinLength terra.StringValue `hcl:"min_length,attr"`
}

type SmsConfiguration struct {
	// ExternalId: string, required
	ExternalId terra.StringValue `hcl:"external_id,attr" validate:"required"`
	// SnsCallerArn: string, required
	SnsCallerArn terra.StringValue `hcl:"sns_caller_arn,attr" validate:"required"`
	// SnsRegion: string, optional
	SnsRegion terra.StringValue `hcl:"sns_region,attr"`
}

type SoftwareTokenMfaConfiguration struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
}

type UserAttributeUpdateSettings struct {
	// AttributesRequireVerificationBeforeUpdate: set of string, required
	AttributesRequireVerificationBeforeUpdate terra.SetValue[terra.StringValue] `hcl:"attributes_require_verification_before_update,attr" validate:"required"`
}

type UserPoolAddOns struct {
	// AdvancedSecurityMode: string, required
	AdvancedSecurityMode terra.StringValue `hcl:"advanced_security_mode,attr" validate:"required"`
}

type UsernameConfiguration struct {
	// CaseSensitive: bool, required
	CaseSensitive terra.BoolValue `hcl:"case_sensitive,attr" validate:"required"`
}

type VerificationMessageTemplate struct {
	// DefaultEmailOption: string, optional
	DefaultEmailOption terra.StringValue `hcl:"default_email_option,attr"`
	// EmailMessage: string, optional
	EmailMessage terra.StringValue `hcl:"email_message,attr"`
	// EmailMessageByLink: string, optional
	EmailMessageByLink terra.StringValue `hcl:"email_message_by_link,attr"`
	// EmailSubject: string, optional
	EmailSubject terra.StringValue `hcl:"email_subject,attr"`
	// EmailSubjectByLink: string, optional
	EmailSubjectByLink terra.StringValue `hcl:"email_subject_by_link,attr"`
	// SmsMessage: string, optional
	SmsMessage terra.StringValue `hcl:"sms_message,attr"`
}

type AccountRecoverySettingAttributes struct {
	ref terra.Reference
}

func (ars AccountRecoverySettingAttributes) InternalRef() terra.Reference {
	return ars.ref
}

func (ars AccountRecoverySettingAttributes) InternalWithRef(ref terra.Reference) AccountRecoverySettingAttributes {
	return AccountRecoverySettingAttributes{ref: ref}
}

func (ars AccountRecoverySettingAttributes) InternalTokens() hclwrite.Tokens {
	return ars.ref.InternalTokens()
}

func (ars AccountRecoverySettingAttributes) RecoveryMechanism() terra.SetValue[RecoveryMechanismAttributes] {
	return terra.ReferenceSet[RecoveryMechanismAttributes](ars.ref.Append("recovery_mechanism"))
}

type RecoveryMechanismAttributes struct {
	ref terra.Reference
}

func (rm RecoveryMechanismAttributes) InternalRef() terra.Reference {
	return rm.ref
}

func (rm RecoveryMechanismAttributes) InternalWithRef(ref terra.Reference) RecoveryMechanismAttributes {
	return RecoveryMechanismAttributes{ref: ref}
}

func (rm RecoveryMechanismAttributes) InternalTokens() hclwrite.Tokens {
	return rm.ref.InternalTokens()
}

func (rm RecoveryMechanismAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rm.ref.Append("name"))
}

func (rm RecoveryMechanismAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(rm.ref.Append("priority"))
}

type AdminCreateUserConfigAttributes struct {
	ref terra.Reference
}

func (acuc AdminCreateUserConfigAttributes) InternalRef() terra.Reference {
	return acuc.ref
}

func (acuc AdminCreateUserConfigAttributes) InternalWithRef(ref terra.Reference) AdminCreateUserConfigAttributes {
	return AdminCreateUserConfigAttributes{ref: ref}
}

func (acuc AdminCreateUserConfigAttributes) InternalTokens() hclwrite.Tokens {
	return acuc.ref.InternalTokens()
}

func (acuc AdminCreateUserConfigAttributes) AllowAdminCreateUserOnly() terra.BoolValue {
	return terra.ReferenceBool(acuc.ref.Append("allow_admin_create_user_only"))
}

func (acuc AdminCreateUserConfigAttributes) InviteMessageTemplate() terra.ListValue[InviteMessageTemplateAttributes] {
	return terra.ReferenceList[InviteMessageTemplateAttributes](acuc.ref.Append("invite_message_template"))
}

type InviteMessageTemplateAttributes struct {
	ref terra.Reference
}

func (imt InviteMessageTemplateAttributes) InternalRef() terra.Reference {
	return imt.ref
}

func (imt InviteMessageTemplateAttributes) InternalWithRef(ref terra.Reference) InviteMessageTemplateAttributes {
	return InviteMessageTemplateAttributes{ref: ref}
}

func (imt InviteMessageTemplateAttributes) InternalTokens() hclwrite.Tokens {
	return imt.ref.InternalTokens()
}

func (imt InviteMessageTemplateAttributes) EmailMessage() terra.StringValue {
	return terra.ReferenceString(imt.ref.Append("email_message"))
}

func (imt InviteMessageTemplateAttributes) EmailSubject() terra.StringValue {
	return terra.ReferenceString(imt.ref.Append("email_subject"))
}

func (imt InviteMessageTemplateAttributes) SmsMessage() terra.StringValue {
	return terra.ReferenceString(imt.ref.Append("sms_message"))
}

type DeviceConfigurationAttributes struct {
	ref terra.Reference
}

func (dc DeviceConfigurationAttributes) InternalRef() terra.Reference {
	return dc.ref
}

func (dc DeviceConfigurationAttributes) InternalWithRef(ref terra.Reference) DeviceConfigurationAttributes {
	return DeviceConfigurationAttributes{ref: ref}
}

func (dc DeviceConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return dc.ref.InternalTokens()
}

func (dc DeviceConfigurationAttributes) ChallengeRequiredOnNewDevice() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("challenge_required_on_new_device"))
}

func (dc DeviceConfigurationAttributes) DeviceOnlyRememberedOnUserPrompt() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("device_only_remembered_on_user_prompt"))
}

type EmailConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EmailConfigurationAttributes) InternalRef() terra.Reference {
	return ec.ref
}

func (ec EmailConfigurationAttributes) InternalWithRef(ref terra.Reference) EmailConfigurationAttributes {
	return EmailConfigurationAttributes{ref: ref}
}

func (ec EmailConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ec.ref.InternalTokens()
}

func (ec EmailConfigurationAttributes) ConfigurationSet() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("configuration_set"))
}

func (ec EmailConfigurationAttributes) EmailSendingAccount() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("email_sending_account"))
}

func (ec EmailConfigurationAttributes) FromEmailAddress() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("from_email_address"))
}

func (ec EmailConfigurationAttributes) ReplyToEmailAddress() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("reply_to_email_address"))
}

func (ec EmailConfigurationAttributes) SourceArn() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("source_arn"))
}

type LambdaConfigAttributes struct {
	ref terra.Reference
}

func (lc LambdaConfigAttributes) InternalRef() terra.Reference {
	return lc.ref
}

func (lc LambdaConfigAttributes) InternalWithRef(ref terra.Reference) LambdaConfigAttributes {
	return LambdaConfigAttributes{ref: ref}
}

func (lc LambdaConfigAttributes) InternalTokens() hclwrite.Tokens {
	return lc.ref.InternalTokens()
}

func (lc LambdaConfigAttributes) CreateAuthChallenge() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("create_auth_challenge"))
}

func (lc LambdaConfigAttributes) CustomMessage() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("custom_message"))
}

func (lc LambdaConfigAttributes) DefineAuthChallenge() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("define_auth_challenge"))
}

func (lc LambdaConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("kms_key_id"))
}

func (lc LambdaConfigAttributes) PostAuthentication() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("post_authentication"))
}

func (lc LambdaConfigAttributes) PostConfirmation() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("post_confirmation"))
}

func (lc LambdaConfigAttributes) PreAuthentication() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("pre_authentication"))
}

func (lc LambdaConfigAttributes) PreSignUp() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("pre_sign_up"))
}

func (lc LambdaConfigAttributes) PreTokenGeneration() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("pre_token_generation"))
}

func (lc LambdaConfigAttributes) UserMigration() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("user_migration"))
}

func (lc LambdaConfigAttributes) VerifyAuthChallengeResponse() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("verify_auth_challenge_response"))
}

func (lc LambdaConfigAttributes) CustomEmailSender() terra.ListValue[CustomEmailSenderAttributes] {
	return terra.ReferenceList[CustomEmailSenderAttributes](lc.ref.Append("custom_email_sender"))
}

func (lc LambdaConfigAttributes) CustomSmsSender() terra.ListValue[CustomSmsSenderAttributes] {
	return terra.ReferenceList[CustomSmsSenderAttributes](lc.ref.Append("custom_sms_sender"))
}

type CustomEmailSenderAttributes struct {
	ref terra.Reference
}

func (ces CustomEmailSenderAttributes) InternalRef() terra.Reference {
	return ces.ref
}

func (ces CustomEmailSenderAttributes) InternalWithRef(ref terra.Reference) CustomEmailSenderAttributes {
	return CustomEmailSenderAttributes{ref: ref}
}

func (ces CustomEmailSenderAttributes) InternalTokens() hclwrite.Tokens {
	return ces.ref.InternalTokens()
}

func (ces CustomEmailSenderAttributes) LambdaArn() terra.StringValue {
	return terra.ReferenceString(ces.ref.Append("lambda_arn"))
}

func (ces CustomEmailSenderAttributes) LambdaVersion() terra.StringValue {
	return terra.ReferenceString(ces.ref.Append("lambda_version"))
}

type CustomSmsSenderAttributes struct {
	ref terra.Reference
}

func (css CustomSmsSenderAttributes) InternalRef() terra.Reference {
	return css.ref
}

func (css CustomSmsSenderAttributes) InternalWithRef(ref terra.Reference) CustomSmsSenderAttributes {
	return CustomSmsSenderAttributes{ref: ref}
}

func (css CustomSmsSenderAttributes) InternalTokens() hclwrite.Tokens {
	return css.ref.InternalTokens()
}

func (css CustomSmsSenderAttributes) LambdaArn() terra.StringValue {
	return terra.ReferenceString(css.ref.Append("lambda_arn"))
}

func (css CustomSmsSenderAttributes) LambdaVersion() terra.StringValue {
	return terra.ReferenceString(css.ref.Append("lambda_version"))
}

type PasswordPolicyAttributes struct {
	ref terra.Reference
}

func (pp PasswordPolicyAttributes) InternalRef() terra.Reference {
	return pp.ref
}

func (pp PasswordPolicyAttributes) InternalWithRef(ref terra.Reference) PasswordPolicyAttributes {
	return PasswordPolicyAttributes{ref: ref}
}

func (pp PasswordPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return pp.ref.InternalTokens()
}

func (pp PasswordPolicyAttributes) MinimumLength() terra.NumberValue {
	return terra.ReferenceNumber(pp.ref.Append("minimum_length"))
}

func (pp PasswordPolicyAttributes) RequireLowercase() terra.BoolValue {
	return terra.ReferenceBool(pp.ref.Append("require_lowercase"))
}

func (pp PasswordPolicyAttributes) RequireNumbers() terra.BoolValue {
	return terra.ReferenceBool(pp.ref.Append("require_numbers"))
}

func (pp PasswordPolicyAttributes) RequireSymbols() terra.BoolValue {
	return terra.ReferenceBool(pp.ref.Append("require_symbols"))
}

func (pp PasswordPolicyAttributes) RequireUppercase() terra.BoolValue {
	return terra.ReferenceBool(pp.ref.Append("require_uppercase"))
}

func (pp PasswordPolicyAttributes) TemporaryPasswordValidityDays() terra.NumberValue {
	return terra.ReferenceNumber(pp.ref.Append("temporary_password_validity_days"))
}

type SchemaAttributes struct {
	ref terra.Reference
}

func (s SchemaAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SchemaAttributes) InternalWithRef(ref terra.Reference) SchemaAttributes {
	return SchemaAttributes{ref: ref}
}

func (s SchemaAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SchemaAttributes) AttributeDataType() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("attribute_data_type"))
}

func (s SchemaAttributes) DeveloperOnlyAttribute() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("developer_only_attribute"))
}

func (s SchemaAttributes) Mutable() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("mutable"))
}

func (s SchemaAttributes) Name() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("name"))
}

func (s SchemaAttributes) Required() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("required"))
}

func (s SchemaAttributes) NumberAttributeConstraints() terra.ListValue[NumberAttributeConstraintsAttributes] {
	return terra.ReferenceList[NumberAttributeConstraintsAttributes](s.ref.Append("number_attribute_constraints"))
}

func (s SchemaAttributes) StringAttributeConstraints() terra.ListValue[StringAttributeConstraintsAttributes] {
	return terra.ReferenceList[StringAttributeConstraintsAttributes](s.ref.Append("string_attribute_constraints"))
}

type NumberAttributeConstraintsAttributes struct {
	ref terra.Reference
}

func (nac NumberAttributeConstraintsAttributes) InternalRef() terra.Reference {
	return nac.ref
}

func (nac NumberAttributeConstraintsAttributes) InternalWithRef(ref terra.Reference) NumberAttributeConstraintsAttributes {
	return NumberAttributeConstraintsAttributes{ref: ref}
}

func (nac NumberAttributeConstraintsAttributes) InternalTokens() hclwrite.Tokens {
	return nac.ref.InternalTokens()
}

func (nac NumberAttributeConstraintsAttributes) MaxValue() terra.StringValue {
	return terra.ReferenceString(nac.ref.Append("max_value"))
}

func (nac NumberAttributeConstraintsAttributes) MinValue() terra.StringValue {
	return terra.ReferenceString(nac.ref.Append("min_value"))
}

type StringAttributeConstraintsAttributes struct {
	ref terra.Reference
}

func (sac StringAttributeConstraintsAttributes) InternalRef() terra.Reference {
	return sac.ref
}

func (sac StringAttributeConstraintsAttributes) InternalWithRef(ref terra.Reference) StringAttributeConstraintsAttributes {
	return StringAttributeConstraintsAttributes{ref: ref}
}

func (sac StringAttributeConstraintsAttributes) InternalTokens() hclwrite.Tokens {
	return sac.ref.InternalTokens()
}

func (sac StringAttributeConstraintsAttributes) MaxLength() terra.StringValue {
	return terra.ReferenceString(sac.ref.Append("max_length"))
}

func (sac StringAttributeConstraintsAttributes) MinLength() terra.StringValue {
	return terra.ReferenceString(sac.ref.Append("min_length"))
}

type SmsConfigurationAttributes struct {
	ref terra.Reference
}

func (sc SmsConfigurationAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc SmsConfigurationAttributes) InternalWithRef(ref terra.Reference) SmsConfigurationAttributes {
	return SmsConfigurationAttributes{ref: ref}
}

func (sc SmsConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc SmsConfigurationAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("external_id"))
}

func (sc SmsConfigurationAttributes) SnsCallerArn() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("sns_caller_arn"))
}

func (sc SmsConfigurationAttributes) SnsRegion() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("sns_region"))
}

type SoftwareTokenMfaConfigurationAttributes struct {
	ref terra.Reference
}

func (stmc SoftwareTokenMfaConfigurationAttributes) InternalRef() terra.Reference {
	return stmc.ref
}

func (stmc SoftwareTokenMfaConfigurationAttributes) InternalWithRef(ref terra.Reference) SoftwareTokenMfaConfigurationAttributes {
	return SoftwareTokenMfaConfigurationAttributes{ref: ref}
}

func (stmc SoftwareTokenMfaConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return stmc.ref.InternalTokens()
}

func (stmc SoftwareTokenMfaConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(stmc.ref.Append("enabled"))
}

type UserAttributeUpdateSettingsAttributes struct {
	ref terra.Reference
}

func (uaus UserAttributeUpdateSettingsAttributes) InternalRef() terra.Reference {
	return uaus.ref
}

func (uaus UserAttributeUpdateSettingsAttributes) InternalWithRef(ref terra.Reference) UserAttributeUpdateSettingsAttributes {
	return UserAttributeUpdateSettingsAttributes{ref: ref}
}

func (uaus UserAttributeUpdateSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return uaus.ref.InternalTokens()
}

func (uaus UserAttributeUpdateSettingsAttributes) AttributesRequireVerificationBeforeUpdate() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](uaus.ref.Append("attributes_require_verification_before_update"))
}

type UserPoolAddOnsAttributes struct {
	ref terra.Reference
}

func (upao UserPoolAddOnsAttributes) InternalRef() terra.Reference {
	return upao.ref
}

func (upao UserPoolAddOnsAttributes) InternalWithRef(ref terra.Reference) UserPoolAddOnsAttributes {
	return UserPoolAddOnsAttributes{ref: ref}
}

func (upao UserPoolAddOnsAttributes) InternalTokens() hclwrite.Tokens {
	return upao.ref.InternalTokens()
}

func (upao UserPoolAddOnsAttributes) AdvancedSecurityMode() terra.StringValue {
	return terra.ReferenceString(upao.ref.Append("advanced_security_mode"))
}

type UsernameConfigurationAttributes struct {
	ref terra.Reference
}

func (uc UsernameConfigurationAttributes) InternalRef() terra.Reference {
	return uc.ref
}

func (uc UsernameConfigurationAttributes) InternalWithRef(ref terra.Reference) UsernameConfigurationAttributes {
	return UsernameConfigurationAttributes{ref: ref}
}

func (uc UsernameConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return uc.ref.InternalTokens()
}

func (uc UsernameConfigurationAttributes) CaseSensitive() terra.BoolValue {
	return terra.ReferenceBool(uc.ref.Append("case_sensitive"))
}

type VerificationMessageTemplateAttributes struct {
	ref terra.Reference
}

func (vmt VerificationMessageTemplateAttributes) InternalRef() terra.Reference {
	return vmt.ref
}

func (vmt VerificationMessageTemplateAttributes) InternalWithRef(ref terra.Reference) VerificationMessageTemplateAttributes {
	return VerificationMessageTemplateAttributes{ref: ref}
}

func (vmt VerificationMessageTemplateAttributes) InternalTokens() hclwrite.Tokens {
	return vmt.ref.InternalTokens()
}

func (vmt VerificationMessageTemplateAttributes) DefaultEmailOption() terra.StringValue {
	return terra.ReferenceString(vmt.ref.Append("default_email_option"))
}

func (vmt VerificationMessageTemplateAttributes) EmailMessage() terra.StringValue {
	return terra.ReferenceString(vmt.ref.Append("email_message"))
}

func (vmt VerificationMessageTemplateAttributes) EmailMessageByLink() terra.StringValue {
	return terra.ReferenceString(vmt.ref.Append("email_message_by_link"))
}

func (vmt VerificationMessageTemplateAttributes) EmailSubject() terra.StringValue {
	return terra.ReferenceString(vmt.ref.Append("email_subject"))
}

func (vmt VerificationMessageTemplateAttributes) EmailSubjectByLink() terra.StringValue {
	return terra.ReferenceString(vmt.ref.Append("email_subject_by_link"))
}

func (vmt VerificationMessageTemplateAttributes) SmsMessage() terra.StringValue {
	return terra.ReferenceString(vmt.ref.Append("sms_message"))
}

type AccountRecoverySettingState struct {
	RecoveryMechanism []RecoveryMechanismState `json:"recovery_mechanism"`
}

type RecoveryMechanismState struct {
	Name     string  `json:"name"`
	Priority float64 `json:"priority"`
}

type AdminCreateUserConfigState struct {
	AllowAdminCreateUserOnly bool                         `json:"allow_admin_create_user_only"`
	InviteMessageTemplate    []InviteMessageTemplateState `json:"invite_message_template"`
}

type InviteMessageTemplateState struct {
	EmailMessage string `json:"email_message"`
	EmailSubject string `json:"email_subject"`
	SmsMessage   string `json:"sms_message"`
}

type DeviceConfigurationState struct {
	ChallengeRequiredOnNewDevice     bool `json:"challenge_required_on_new_device"`
	DeviceOnlyRememberedOnUserPrompt bool `json:"device_only_remembered_on_user_prompt"`
}

type EmailConfigurationState struct {
	ConfigurationSet    string `json:"configuration_set"`
	EmailSendingAccount string `json:"email_sending_account"`
	FromEmailAddress    string `json:"from_email_address"`
	ReplyToEmailAddress string `json:"reply_to_email_address"`
	SourceArn           string `json:"source_arn"`
}

type LambdaConfigState struct {
	CreateAuthChallenge         string                   `json:"create_auth_challenge"`
	CustomMessage               string                   `json:"custom_message"`
	DefineAuthChallenge         string                   `json:"define_auth_challenge"`
	KmsKeyId                    string                   `json:"kms_key_id"`
	PostAuthentication          string                   `json:"post_authentication"`
	PostConfirmation            string                   `json:"post_confirmation"`
	PreAuthentication           string                   `json:"pre_authentication"`
	PreSignUp                   string                   `json:"pre_sign_up"`
	PreTokenGeneration          string                   `json:"pre_token_generation"`
	UserMigration               string                   `json:"user_migration"`
	VerifyAuthChallengeResponse string                   `json:"verify_auth_challenge_response"`
	CustomEmailSender           []CustomEmailSenderState `json:"custom_email_sender"`
	CustomSmsSender             []CustomSmsSenderState   `json:"custom_sms_sender"`
}

type CustomEmailSenderState struct {
	LambdaArn     string `json:"lambda_arn"`
	LambdaVersion string `json:"lambda_version"`
}

type CustomSmsSenderState struct {
	LambdaArn     string `json:"lambda_arn"`
	LambdaVersion string `json:"lambda_version"`
}

type PasswordPolicyState struct {
	MinimumLength                 float64 `json:"minimum_length"`
	RequireLowercase              bool    `json:"require_lowercase"`
	RequireNumbers                bool    `json:"require_numbers"`
	RequireSymbols                bool    `json:"require_symbols"`
	RequireUppercase              bool    `json:"require_uppercase"`
	TemporaryPasswordValidityDays float64 `json:"temporary_password_validity_days"`
}

type SchemaState struct {
	AttributeDataType          string                            `json:"attribute_data_type"`
	DeveloperOnlyAttribute     bool                              `json:"developer_only_attribute"`
	Mutable                    bool                              `json:"mutable"`
	Name                       string                            `json:"name"`
	Required                   bool                              `json:"required"`
	NumberAttributeConstraints []NumberAttributeConstraintsState `json:"number_attribute_constraints"`
	StringAttributeConstraints []StringAttributeConstraintsState `json:"string_attribute_constraints"`
}

type NumberAttributeConstraintsState struct {
	MaxValue string `json:"max_value"`
	MinValue string `json:"min_value"`
}

type StringAttributeConstraintsState struct {
	MaxLength string `json:"max_length"`
	MinLength string `json:"min_length"`
}

type SmsConfigurationState struct {
	ExternalId   string `json:"external_id"`
	SnsCallerArn string `json:"sns_caller_arn"`
	SnsRegion    string `json:"sns_region"`
}

type SoftwareTokenMfaConfigurationState struct {
	Enabled bool `json:"enabled"`
}

type UserAttributeUpdateSettingsState struct {
	AttributesRequireVerificationBeforeUpdate []string `json:"attributes_require_verification_before_update"`
}

type UserPoolAddOnsState struct {
	AdvancedSecurityMode string `json:"advanced_security_mode"`
}

type UsernameConfigurationState struct {
	CaseSensitive bool `json:"case_sensitive"`
}

type VerificationMessageTemplateState struct {
	DefaultEmailOption string `json:"default_email_option"`
	EmailMessage       string `json:"email_message"`
	EmailMessageByLink string `json:"email_message_by_link"`
	EmailSubject       string `json:"email_subject"`
	EmailSubjectByLink string `json:"email_subject_by_link"`
	SmsMessage         string `json:"sms_message"`
}
