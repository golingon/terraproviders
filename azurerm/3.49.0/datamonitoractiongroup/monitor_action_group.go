// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamonitoractiongroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ArmRoleReceiver struct{}

type AutomationRunbookReceiver struct{}

type AzureAppPushReceiver struct{}

type AzureFunctionReceiver struct{}

type EmailReceiver struct{}

type EventHubReceiver struct{}

type ItsmReceiver struct{}

type LogicAppReceiver struct{}

type SmsReceiver struct{}

type VoiceReceiver struct{}

type WebhookReceiver struct {
	// AadAuth: min=0
	AadAuth []AadAuth `hcl:"aad_auth,block" validate:"min=0"`
}

type AadAuth struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ArmRoleReceiverAttributes struct {
	ref terra.Reference
}

func (arr ArmRoleReceiverAttributes) InternalRef() terra.Reference {
	return arr.ref
}

func (arr ArmRoleReceiverAttributes) InternalWithRef(ref terra.Reference) ArmRoleReceiverAttributes {
	return ArmRoleReceiverAttributes{ref: ref}
}

func (arr ArmRoleReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return arr.ref.InternalTokens()
}

func (arr ArmRoleReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("name"))
}

func (arr ArmRoleReceiverAttributes) RoleId() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("role_id"))
}

func (arr ArmRoleReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(arr.ref.Append("use_common_alert_schema"))
}

type AutomationRunbookReceiverAttributes struct {
	ref terra.Reference
}

func (arr AutomationRunbookReceiverAttributes) InternalRef() terra.Reference {
	return arr.ref
}

func (arr AutomationRunbookReceiverAttributes) InternalWithRef(ref terra.Reference) AutomationRunbookReceiverAttributes {
	return AutomationRunbookReceiverAttributes{ref: ref}
}

func (arr AutomationRunbookReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return arr.ref.InternalTokens()
}

func (arr AutomationRunbookReceiverAttributes) AutomationAccountId() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("automation_account_id"))
}

func (arr AutomationRunbookReceiverAttributes) IsGlobalRunbook() terra.BoolValue {
	return terra.ReferenceBool(arr.ref.Append("is_global_runbook"))
}

func (arr AutomationRunbookReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("name"))
}

func (arr AutomationRunbookReceiverAttributes) RunbookName() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("runbook_name"))
}

func (arr AutomationRunbookReceiverAttributes) ServiceUri() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("service_uri"))
}

func (arr AutomationRunbookReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(arr.ref.Append("use_common_alert_schema"))
}

func (arr AutomationRunbookReceiverAttributes) WebhookResourceId() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("webhook_resource_id"))
}

type AzureAppPushReceiverAttributes struct {
	ref terra.Reference
}

func (aapr AzureAppPushReceiverAttributes) InternalRef() terra.Reference {
	return aapr.ref
}

func (aapr AzureAppPushReceiverAttributes) InternalWithRef(ref terra.Reference) AzureAppPushReceiverAttributes {
	return AzureAppPushReceiverAttributes{ref: ref}
}

func (aapr AzureAppPushReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return aapr.ref.InternalTokens()
}

func (aapr AzureAppPushReceiverAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceString(aapr.ref.Append("email_address"))
}

func (aapr AzureAppPushReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(aapr.ref.Append("name"))
}

type AzureFunctionReceiverAttributes struct {
	ref terra.Reference
}

func (afr AzureFunctionReceiverAttributes) InternalRef() terra.Reference {
	return afr.ref
}

func (afr AzureFunctionReceiverAttributes) InternalWithRef(ref terra.Reference) AzureFunctionReceiverAttributes {
	return AzureFunctionReceiverAttributes{ref: ref}
}

func (afr AzureFunctionReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return afr.ref.InternalTokens()
}

func (afr AzureFunctionReceiverAttributes) FunctionAppResourceId() terra.StringValue {
	return terra.ReferenceString(afr.ref.Append("function_app_resource_id"))
}

func (afr AzureFunctionReceiverAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceString(afr.ref.Append("function_name"))
}

func (afr AzureFunctionReceiverAttributes) HttpTriggerUrl() terra.StringValue {
	return terra.ReferenceString(afr.ref.Append("http_trigger_url"))
}

func (afr AzureFunctionReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(afr.ref.Append("name"))
}

func (afr AzureFunctionReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(afr.ref.Append("use_common_alert_schema"))
}

type EmailReceiverAttributes struct {
	ref terra.Reference
}

func (er EmailReceiverAttributes) InternalRef() terra.Reference {
	return er.ref
}

func (er EmailReceiverAttributes) InternalWithRef(ref terra.Reference) EmailReceiverAttributes {
	return EmailReceiverAttributes{ref: ref}
}

func (er EmailReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return er.ref.InternalTokens()
}

func (er EmailReceiverAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("email_address"))
}

func (er EmailReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("name"))
}

func (er EmailReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(er.ref.Append("use_common_alert_schema"))
}

type EventHubReceiverAttributes struct {
	ref terra.Reference
}

func (ehr EventHubReceiverAttributes) InternalRef() terra.Reference {
	return ehr.ref
}

func (ehr EventHubReceiverAttributes) InternalWithRef(ref terra.Reference) EventHubReceiverAttributes {
	return EventHubReceiverAttributes{ref: ref}
}

func (ehr EventHubReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return ehr.ref.InternalTokens()
}

func (ehr EventHubReceiverAttributes) EventHubId() terra.StringValue {
	return terra.ReferenceString(ehr.ref.Append("event_hub_id"))
}

func (ehr EventHubReceiverAttributes) EventHubName() terra.StringValue {
	return terra.ReferenceString(ehr.ref.Append("event_hub_name"))
}

func (ehr EventHubReceiverAttributes) EventHubNamespace() terra.StringValue {
	return terra.ReferenceString(ehr.ref.Append("event_hub_namespace"))
}

func (ehr EventHubReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ehr.ref.Append("name"))
}

func (ehr EventHubReceiverAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceString(ehr.ref.Append("subscription_id"))
}

func (ehr EventHubReceiverAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(ehr.ref.Append("tenant_id"))
}

func (ehr EventHubReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(ehr.ref.Append("use_common_alert_schema"))
}

type ItsmReceiverAttributes struct {
	ref terra.Reference
}

func (ir ItsmReceiverAttributes) InternalRef() terra.Reference {
	return ir.ref
}

func (ir ItsmReceiverAttributes) InternalWithRef(ref terra.Reference) ItsmReceiverAttributes {
	return ItsmReceiverAttributes{ref: ref}
}

func (ir ItsmReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return ir.ref.InternalTokens()
}

func (ir ItsmReceiverAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("connection_id"))
}

func (ir ItsmReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("name"))
}

func (ir ItsmReceiverAttributes) Region() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("region"))
}

func (ir ItsmReceiverAttributes) TicketConfiguration() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("ticket_configuration"))
}

func (ir ItsmReceiverAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceString(ir.ref.Append("workspace_id"))
}

type LogicAppReceiverAttributes struct {
	ref terra.Reference
}

func (lar LogicAppReceiverAttributes) InternalRef() terra.Reference {
	return lar.ref
}

func (lar LogicAppReceiverAttributes) InternalWithRef(ref terra.Reference) LogicAppReceiverAttributes {
	return LogicAppReceiverAttributes{ref: ref}
}

func (lar LogicAppReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return lar.ref.InternalTokens()
}

func (lar LogicAppReceiverAttributes) CallbackUrl() terra.StringValue {
	return terra.ReferenceString(lar.ref.Append("callback_url"))
}

func (lar LogicAppReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lar.ref.Append("name"))
}

func (lar LogicAppReceiverAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceString(lar.ref.Append("resource_id"))
}

func (lar LogicAppReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(lar.ref.Append("use_common_alert_schema"))
}

type SmsReceiverAttributes struct {
	ref terra.Reference
}

func (sr SmsReceiverAttributes) InternalRef() terra.Reference {
	return sr.ref
}

func (sr SmsReceiverAttributes) InternalWithRef(ref terra.Reference) SmsReceiverAttributes {
	return SmsReceiverAttributes{ref: ref}
}

func (sr SmsReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return sr.ref.InternalTokens()
}

func (sr SmsReceiverAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceString(sr.ref.Append("country_code"))
}

func (sr SmsReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sr.ref.Append("name"))
}

func (sr SmsReceiverAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(sr.ref.Append("phone_number"))
}

type VoiceReceiverAttributes struct {
	ref terra.Reference
}

func (vr VoiceReceiverAttributes) InternalRef() terra.Reference {
	return vr.ref
}

func (vr VoiceReceiverAttributes) InternalWithRef(ref terra.Reference) VoiceReceiverAttributes {
	return VoiceReceiverAttributes{ref: ref}
}

func (vr VoiceReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return vr.ref.InternalTokens()
}

func (vr VoiceReceiverAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceString(vr.ref.Append("country_code"))
}

func (vr VoiceReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vr.ref.Append("name"))
}

func (vr VoiceReceiverAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(vr.ref.Append("phone_number"))
}

type WebhookReceiverAttributes struct {
	ref terra.Reference
}

func (wr WebhookReceiverAttributes) InternalRef() terra.Reference {
	return wr.ref
}

func (wr WebhookReceiverAttributes) InternalWithRef(ref terra.Reference) WebhookReceiverAttributes {
	return WebhookReceiverAttributes{ref: ref}
}

func (wr WebhookReceiverAttributes) InternalTokens() hclwrite.Tokens {
	return wr.ref.InternalTokens()
}

func (wr WebhookReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wr.ref.Append("name"))
}

func (wr WebhookReceiverAttributes) ServiceUri() terra.StringValue {
	return terra.ReferenceString(wr.ref.Append("service_uri"))
}

func (wr WebhookReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceBool(wr.ref.Append("use_common_alert_schema"))
}

func (wr WebhookReceiverAttributes) AadAuth() terra.ListValue[AadAuthAttributes] {
	return terra.ReferenceList[AadAuthAttributes](wr.ref.Append("aad_auth"))
}

type AadAuthAttributes struct {
	ref terra.Reference
}

func (aa AadAuthAttributes) InternalRef() terra.Reference {
	return aa.ref
}

func (aa AadAuthAttributes) InternalWithRef(ref terra.Reference) AadAuthAttributes {
	return AadAuthAttributes{ref: ref}
}

func (aa AadAuthAttributes) InternalTokens() hclwrite.Tokens {
	return aa.ref.InternalTokens()
}

func (aa AadAuthAttributes) IdentifierUri() terra.StringValue {
	return terra.ReferenceString(aa.ref.Append("identifier_uri"))
}

func (aa AadAuthAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceString(aa.ref.Append("object_id"))
}

func (aa AadAuthAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(aa.ref.Append("tenant_id"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type ArmRoleReceiverState struct {
	Name                 string `json:"name"`
	RoleId               string `json:"role_id"`
	UseCommonAlertSchema bool   `json:"use_common_alert_schema"`
}

type AutomationRunbookReceiverState struct {
	AutomationAccountId  string `json:"automation_account_id"`
	IsGlobalRunbook      bool   `json:"is_global_runbook"`
	Name                 string `json:"name"`
	RunbookName          string `json:"runbook_name"`
	ServiceUri           string `json:"service_uri"`
	UseCommonAlertSchema bool   `json:"use_common_alert_schema"`
	WebhookResourceId    string `json:"webhook_resource_id"`
}

type AzureAppPushReceiverState struct {
	EmailAddress string `json:"email_address"`
	Name         string `json:"name"`
}

type AzureFunctionReceiverState struct {
	FunctionAppResourceId string `json:"function_app_resource_id"`
	FunctionName          string `json:"function_name"`
	HttpTriggerUrl        string `json:"http_trigger_url"`
	Name                  string `json:"name"`
	UseCommonAlertSchema  bool   `json:"use_common_alert_schema"`
}

type EmailReceiverState struct {
	EmailAddress         string `json:"email_address"`
	Name                 string `json:"name"`
	UseCommonAlertSchema bool   `json:"use_common_alert_schema"`
}

type EventHubReceiverState struct {
	EventHubId           string `json:"event_hub_id"`
	EventHubName         string `json:"event_hub_name"`
	EventHubNamespace    string `json:"event_hub_namespace"`
	Name                 string `json:"name"`
	SubscriptionId       string `json:"subscription_id"`
	TenantId             string `json:"tenant_id"`
	UseCommonAlertSchema bool   `json:"use_common_alert_schema"`
}

type ItsmReceiverState struct {
	ConnectionId        string `json:"connection_id"`
	Name                string `json:"name"`
	Region              string `json:"region"`
	TicketConfiguration string `json:"ticket_configuration"`
	WorkspaceId         string `json:"workspace_id"`
}

type LogicAppReceiverState struct {
	CallbackUrl          string `json:"callback_url"`
	Name                 string `json:"name"`
	ResourceId           string `json:"resource_id"`
	UseCommonAlertSchema bool   `json:"use_common_alert_schema"`
}

type SmsReceiverState struct {
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type VoiceReceiverState struct {
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type WebhookReceiverState struct {
	Name                 string         `json:"name"`
	ServiceUri           string         `json:"service_uri"`
	UseCommonAlertSchema bool           `json:"use_common_alert_schema"`
	AadAuth              []AadAuthState `json:"aad_auth"`
}

type AadAuthState struct {
	IdentifierUri string `json:"identifier_uri"`
	ObjectId      string `json:"object_id"`
	TenantId      string `json:"tenant_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
