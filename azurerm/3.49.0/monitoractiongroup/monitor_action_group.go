// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitoractiongroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ArmRoleReceiver struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleId: string, required
	RoleId terra.StringValue `hcl:"role_id,attr" validate:"required"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
}

type AutomationRunbookReceiver struct {
	// AutomationAccountId: string, required
	AutomationAccountId terra.StringValue `hcl:"automation_account_id,attr" validate:"required"`
	// IsGlobalRunbook: bool, required
	IsGlobalRunbook terra.BoolValue `hcl:"is_global_runbook,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RunbookName: string, required
	RunbookName terra.StringValue `hcl:"runbook_name,attr" validate:"required"`
	// ServiceUri: string, required
	ServiceUri terra.StringValue `hcl:"service_uri,attr" validate:"required"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
	// WebhookResourceId: string, required
	WebhookResourceId terra.StringValue `hcl:"webhook_resource_id,attr" validate:"required"`
}

type AzureAppPushReceiver struct {
	// EmailAddress: string, required
	EmailAddress terra.StringValue `hcl:"email_address,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type AzureFunctionReceiver struct {
	// FunctionAppResourceId: string, required
	FunctionAppResourceId terra.StringValue `hcl:"function_app_resource_id,attr" validate:"required"`
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// HttpTriggerUrl: string, required
	HttpTriggerUrl terra.StringValue `hcl:"http_trigger_url,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
}

type EmailReceiver struct {
	// EmailAddress: string, required
	EmailAddress terra.StringValue `hcl:"email_address,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
}

type EventHubReceiver struct {
	// EventHubId: string, optional
	EventHubId terra.StringValue `hcl:"event_hub_id,attr"`
	// EventHubName: string, optional
	EventHubName terra.StringValue `hcl:"event_hub_name,attr"`
	// EventHubNamespace: string, optional
	EventHubNamespace terra.StringValue `hcl:"event_hub_namespace,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
}

type ItsmReceiver struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// TicketConfiguration: string, required
	TicketConfiguration terra.StringValue `hcl:"ticket_configuration,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
}

type LogicAppReceiver struct {
	// CallbackUrl: string, required
	CallbackUrl terra.StringValue `hcl:"callback_url,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
}

type SmsReceiver struct {
	// CountryCode: string, required
	CountryCode terra.StringValue `hcl:"country_code,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PhoneNumber: string, required
	PhoneNumber terra.StringValue `hcl:"phone_number,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VoiceReceiver struct {
	// CountryCode: string, required
	CountryCode terra.StringValue `hcl:"country_code,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PhoneNumber: string, required
	PhoneNumber terra.StringValue `hcl:"phone_number,attr" validate:"required"`
}

type WebhookReceiver struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceUri: string, required
	ServiceUri terra.StringValue `hcl:"service_uri,attr" validate:"required"`
	// UseCommonAlertSchema: bool, optional
	UseCommonAlertSchema terra.BoolValue `hcl:"use_common_alert_schema,attr"`
	// AadAuth: optional
	AadAuth *AadAuth `hcl:"aad_auth,block"`
}

type AadAuth struct {
	// IdentifierUri: string, optional
	IdentifierUri terra.StringValue `hcl:"identifier_uri,attr"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
}

type ArmRoleReceiverAttributes struct {
	ref terra.Reference
}

func (arr ArmRoleReceiverAttributes) InternalRef() (terra.Reference, error) {
	return arr.ref, nil
}

func (arr ArmRoleReceiverAttributes) InternalWithRef(ref terra.Reference) ArmRoleReceiverAttributes {
	return ArmRoleReceiverAttributes{ref: ref}
}

func (arr ArmRoleReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return arr.ref.InternalTokens()
}

func (arr ArmRoleReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("name"))
}

func (arr ArmRoleReceiverAttributes) RoleId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("role_id"))
}

func (arr ArmRoleReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(arr.ref.Append("use_common_alert_schema"))
}

type AutomationRunbookReceiverAttributes struct {
	ref terra.Reference
}

func (arr AutomationRunbookReceiverAttributes) InternalRef() (terra.Reference, error) {
	return arr.ref, nil
}

func (arr AutomationRunbookReceiverAttributes) InternalWithRef(ref terra.Reference) AutomationRunbookReceiverAttributes {
	return AutomationRunbookReceiverAttributes{ref: ref}
}

func (arr AutomationRunbookReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return arr.ref.InternalTokens()
}

func (arr AutomationRunbookReceiverAttributes) AutomationAccountId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("automation_account_id"))
}

func (arr AutomationRunbookReceiverAttributes) IsGlobalRunbook() terra.BoolValue {
	return terra.ReferenceAsBool(arr.ref.Append("is_global_runbook"))
}

func (arr AutomationRunbookReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("name"))
}

func (arr AutomationRunbookReceiverAttributes) RunbookName() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("runbook_name"))
}

func (arr AutomationRunbookReceiverAttributes) ServiceUri() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("service_uri"))
}

func (arr AutomationRunbookReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(arr.ref.Append("use_common_alert_schema"))
}

func (arr AutomationRunbookReceiverAttributes) WebhookResourceId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("webhook_resource_id"))
}

type AzureAppPushReceiverAttributes struct {
	ref terra.Reference
}

func (aapr AzureAppPushReceiverAttributes) InternalRef() (terra.Reference, error) {
	return aapr.ref, nil
}

func (aapr AzureAppPushReceiverAttributes) InternalWithRef(ref terra.Reference) AzureAppPushReceiverAttributes {
	return AzureAppPushReceiverAttributes{ref: ref}
}

func (aapr AzureAppPushReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aapr.ref.InternalTokens()
}

func (aapr AzureAppPushReceiverAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(aapr.ref.Append("email_address"))
}

func (aapr AzureAppPushReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aapr.ref.Append("name"))
}

type AzureFunctionReceiverAttributes struct {
	ref terra.Reference
}

func (afr AzureFunctionReceiverAttributes) InternalRef() (terra.Reference, error) {
	return afr.ref, nil
}

func (afr AzureFunctionReceiverAttributes) InternalWithRef(ref terra.Reference) AzureFunctionReceiverAttributes {
	return AzureFunctionReceiverAttributes{ref: ref}
}

func (afr AzureFunctionReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return afr.ref.InternalTokens()
}

func (afr AzureFunctionReceiverAttributes) FunctionAppResourceId() terra.StringValue {
	return terra.ReferenceAsString(afr.ref.Append("function_app_resource_id"))
}

func (afr AzureFunctionReceiverAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(afr.ref.Append("function_name"))
}

func (afr AzureFunctionReceiverAttributes) HttpTriggerUrl() terra.StringValue {
	return terra.ReferenceAsString(afr.ref.Append("http_trigger_url"))
}

func (afr AzureFunctionReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(afr.ref.Append("name"))
}

func (afr AzureFunctionReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(afr.ref.Append("use_common_alert_schema"))
}

type EmailReceiverAttributes struct {
	ref terra.Reference
}

func (er EmailReceiverAttributes) InternalRef() (terra.Reference, error) {
	return er.ref, nil
}

func (er EmailReceiverAttributes) InternalWithRef(ref terra.Reference) EmailReceiverAttributes {
	return EmailReceiverAttributes{ref: ref}
}

func (er EmailReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return er.ref.InternalTokens()
}

func (er EmailReceiverAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("email_address"))
}

func (er EmailReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("name"))
}

func (er EmailReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(er.ref.Append("use_common_alert_schema"))
}

type EventHubReceiverAttributes struct {
	ref terra.Reference
}

func (ehr EventHubReceiverAttributes) InternalRef() (terra.Reference, error) {
	return ehr.ref, nil
}

func (ehr EventHubReceiverAttributes) InternalWithRef(ref terra.Reference) EventHubReceiverAttributes {
	return EventHubReceiverAttributes{ref: ref}
}

func (ehr EventHubReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ehr.ref.InternalTokens()
}

func (ehr EventHubReceiverAttributes) EventHubId() terra.StringValue {
	return terra.ReferenceAsString(ehr.ref.Append("event_hub_id"))
}

func (ehr EventHubReceiverAttributes) EventHubName() terra.StringValue {
	return terra.ReferenceAsString(ehr.ref.Append("event_hub_name"))
}

func (ehr EventHubReceiverAttributes) EventHubNamespace() terra.StringValue {
	return terra.ReferenceAsString(ehr.ref.Append("event_hub_namespace"))
}

func (ehr EventHubReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ehr.ref.Append("name"))
}

func (ehr EventHubReceiverAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(ehr.ref.Append("subscription_id"))
}

func (ehr EventHubReceiverAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(ehr.ref.Append("tenant_id"))
}

func (ehr EventHubReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(ehr.ref.Append("use_common_alert_schema"))
}

type ItsmReceiverAttributes struct {
	ref terra.Reference
}

func (ir ItsmReceiverAttributes) InternalRef() (terra.Reference, error) {
	return ir.ref, nil
}

func (ir ItsmReceiverAttributes) InternalWithRef(ref terra.Reference) ItsmReceiverAttributes {
	return ItsmReceiverAttributes{ref: ref}
}

func (ir ItsmReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ir.ref.InternalTokens()
}

func (ir ItsmReceiverAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("connection_id"))
}

func (ir ItsmReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

func (ir ItsmReceiverAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("region"))
}

func (ir ItsmReceiverAttributes) TicketConfiguration() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("ticket_configuration"))
}

func (ir ItsmReceiverAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("workspace_id"))
}

type LogicAppReceiverAttributes struct {
	ref terra.Reference
}

func (lar LogicAppReceiverAttributes) InternalRef() (terra.Reference, error) {
	return lar.ref, nil
}

func (lar LogicAppReceiverAttributes) InternalWithRef(ref terra.Reference) LogicAppReceiverAttributes {
	return LogicAppReceiverAttributes{ref: ref}
}

func (lar LogicAppReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lar.ref.InternalTokens()
}

func (lar LogicAppReceiverAttributes) CallbackUrl() terra.StringValue {
	return terra.ReferenceAsString(lar.ref.Append("callback_url"))
}

func (lar LogicAppReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lar.ref.Append("name"))
}

func (lar LogicAppReceiverAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(lar.ref.Append("resource_id"))
}

func (lar LogicAppReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(lar.ref.Append("use_common_alert_schema"))
}

type SmsReceiverAttributes struct {
	ref terra.Reference
}

func (sr SmsReceiverAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr SmsReceiverAttributes) InternalWithRef(ref terra.Reference) SmsReceiverAttributes {
	return SmsReceiverAttributes{ref: ref}
}

func (sr SmsReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr SmsReceiverAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("country_code"))
}

func (sr SmsReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("name"))
}

func (sr SmsReceiverAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("phone_number"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type VoiceReceiverAttributes struct {
	ref terra.Reference
}

func (vr VoiceReceiverAttributes) InternalRef() (terra.Reference, error) {
	return vr.ref, nil
}

func (vr VoiceReceiverAttributes) InternalWithRef(ref terra.Reference) VoiceReceiverAttributes {
	return VoiceReceiverAttributes{ref: ref}
}

func (vr VoiceReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vr.ref.InternalTokens()
}

func (vr VoiceReceiverAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(vr.ref.Append("country_code"))
}

func (vr VoiceReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vr.ref.Append("name"))
}

func (vr VoiceReceiverAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(vr.ref.Append("phone_number"))
}

type WebhookReceiverAttributes struct {
	ref terra.Reference
}

func (wr WebhookReceiverAttributes) InternalRef() (terra.Reference, error) {
	return wr.ref, nil
}

func (wr WebhookReceiverAttributes) InternalWithRef(ref terra.Reference) WebhookReceiverAttributes {
	return WebhookReceiverAttributes{ref: ref}
}

func (wr WebhookReceiverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wr.ref.InternalTokens()
}

func (wr WebhookReceiverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("name"))
}

func (wr WebhookReceiverAttributes) ServiceUri() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("service_uri"))
}

func (wr WebhookReceiverAttributes) UseCommonAlertSchema() terra.BoolValue {
	return terra.ReferenceAsBool(wr.ref.Append("use_common_alert_schema"))
}

func (wr WebhookReceiverAttributes) AadAuth() terra.ListValue[AadAuthAttributes] {
	return terra.ReferenceAsList[AadAuthAttributes](wr.ref.Append("aad_auth"))
}

type AadAuthAttributes struct {
	ref terra.Reference
}

func (aa AadAuthAttributes) InternalRef() (terra.Reference, error) {
	return aa.ref, nil
}

func (aa AadAuthAttributes) InternalWithRef(ref terra.Reference) AadAuthAttributes {
	return AadAuthAttributes{ref: ref}
}

func (aa AadAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aa.ref.InternalTokens()
}

func (aa AadAuthAttributes) IdentifierUri() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("identifier_uri"))
}

func (aa AadAuthAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("object_id"))
}

func (aa AadAuthAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("tenant_id"))
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

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
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
