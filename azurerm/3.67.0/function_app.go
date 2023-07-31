// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	functionapp "github.com/golingon/terraproviders/azurerm/3.67.0/functionapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFunctionApp creates a new instance of [FunctionApp].
func NewFunctionApp(name string, args FunctionAppArgs) *FunctionApp {
	return &FunctionApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FunctionApp)(nil)

// FunctionApp represents the Terraform resource azurerm_function_app.
type FunctionApp struct {
	Name      string
	Args      FunctionAppArgs
	state     *functionAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FunctionApp].
func (fa *FunctionApp) Type() string {
	return "azurerm_function_app"
}

// LocalName returns the local name for [FunctionApp].
func (fa *FunctionApp) LocalName() string {
	return fa.Name
}

// Configuration returns the configuration (args) for [FunctionApp].
func (fa *FunctionApp) Configuration() interface{} {
	return fa.Args
}

// DependOn is used for other resources to depend on [FunctionApp].
func (fa *FunctionApp) DependOn() terra.Reference {
	return terra.ReferenceResource(fa)
}

// Dependencies returns the list of resources [FunctionApp] depends_on.
func (fa *FunctionApp) Dependencies() terra.Dependencies {
	return fa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FunctionApp].
func (fa *FunctionApp) LifecycleManagement() *terra.Lifecycle {
	return fa.Lifecycle
}

// Attributes returns the attributes for [FunctionApp].
func (fa *FunctionApp) Attributes() functionAppAttributes {
	return functionAppAttributes{ref: terra.ReferenceResource(fa)}
}

// ImportState imports the given attribute values into [FunctionApp]'s state.
func (fa *FunctionApp) ImportState(av io.Reader) error {
	fa.state = &functionAppState{}
	if err := json.NewDecoder(av).Decode(fa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fa.Type(), fa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FunctionApp] has state.
func (fa *FunctionApp) State() (*functionAppState, bool) {
	return fa.state, fa.state != nil
}

// StateMust returns the state for [FunctionApp]. Panics if the state is nil.
func (fa *FunctionApp) StateMust() *functionAppState {
	if fa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fa.Type(), fa.LocalName()))
	}
	return fa.state
}

// FunctionAppArgs contains the configurations for azurerm_function_app.
type FunctionAppArgs struct {
	// AppServicePlanId: string, required
	AppServicePlanId terra.StringValue `hcl:"app_service_plan_id,attr" validate:"required"`
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// ClientCertMode: string, optional
	ClientCertMode terra.StringValue `hcl:"client_cert_mode,attr"`
	// DailyMemoryTimeQuota: number, optional
	DailyMemoryTimeQuota terra.NumberValue `hcl:"daily_memory_time_quota,attr"`
	// EnableBuiltinLogging: bool, optional
	EnableBuiltinLogging terra.BoolValue `hcl:"enable_builtin_logging,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultReferenceIdentityId: string, optional
	KeyVaultReferenceIdentityId terra.StringValue `hcl:"key_vault_reference_identity_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OsType: string, optional
	OsType terra.StringValue `hcl:"os_type,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountAccessKey: string, required
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// SiteCredential: min=0
	SiteCredential []functionapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *functionapp.AuthSettings `hcl:"auth_settings,block"`
	// ConnectionString: min=0
	ConnectionString []functionapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *functionapp.Identity `hcl:"identity,block"`
	// SiteConfig: optional
	SiteConfig *functionapp.SiteConfig `hcl:"site_config,block"`
	// SourceControl: optional
	SourceControl *functionapp.SourceControl `hcl:"source_control,block"`
	// Timeouts: optional
	Timeouts *functionapp.Timeouts `hcl:"timeouts,block"`
}
type functionAppAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_function_app.
func (fa functionAppAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_function_app.
func (fa functionAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fa.ref.Append("app_settings"))
}

// ClientCertMode returns a reference to field client_cert_mode of azurerm_function_app.
func (fa functionAppAttributes) ClientCertMode() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("client_cert_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_function_app.
func (fa functionAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("custom_domain_verification_id"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_function_app.
func (fa functionAppAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(fa.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_function_app.
func (fa functionAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("default_hostname"))
}

// EnableBuiltinLogging returns a reference to field enable_builtin_logging of azurerm_function_app.
func (fa functionAppAttributes) EnableBuiltinLogging() terra.BoolValue {
	return terra.ReferenceAsBool(fa.ref.Append("enable_builtin_logging"))
}

// Enabled returns a reference to field enabled of azurerm_function_app.
func (fa functionAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(fa.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_function_app.
func (fa functionAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(fa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_function_app.
func (fa functionAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_function_app.
func (fa functionAppAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_function_app.
func (fa functionAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_function_app.
func (fa functionAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_function_app.
func (fa functionAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_function_app.
func (fa functionAppAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("os_type"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_function_app.
func (fa functionAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_function_app.
func (fa functionAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_function_app.
func (fa functionAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("resource_group_name"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_function_app.
func (fa functionAppAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_function_app.
func (fa functionAppAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("storage_account_name"))
}

// Tags returns a reference to field tags of azurerm_function_app.
func (fa functionAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fa.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_function_app.
func (fa functionAppAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("version"))
}

func (fa functionAppAttributes) SiteCredential() terra.ListValue[functionapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[functionapp.SiteCredentialAttributes](fa.ref.Append("site_credential"))
}

func (fa functionAppAttributes) AuthSettings() terra.ListValue[functionapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[functionapp.AuthSettingsAttributes](fa.ref.Append("auth_settings"))
}

func (fa functionAppAttributes) ConnectionString() terra.SetValue[functionapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[functionapp.ConnectionStringAttributes](fa.ref.Append("connection_string"))
}

func (fa functionAppAttributes) Identity() terra.ListValue[functionapp.IdentityAttributes] {
	return terra.ReferenceAsList[functionapp.IdentityAttributes](fa.ref.Append("identity"))
}

func (fa functionAppAttributes) SiteConfig() terra.ListValue[functionapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[functionapp.SiteConfigAttributes](fa.ref.Append("site_config"))
}

func (fa functionAppAttributes) SourceControl() terra.ListValue[functionapp.SourceControlAttributes] {
	return terra.ReferenceAsList[functionapp.SourceControlAttributes](fa.ref.Append("source_control"))
}

func (fa functionAppAttributes) Timeouts() functionapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[functionapp.TimeoutsAttributes](fa.ref.Append("timeouts"))
}

type functionAppState struct {
	AppServicePlanId            string                              `json:"app_service_plan_id"`
	AppSettings                 map[string]string                   `json:"app_settings"`
	ClientCertMode              string                              `json:"client_cert_mode"`
	CustomDomainVerificationId  string                              `json:"custom_domain_verification_id"`
	DailyMemoryTimeQuota        float64                             `json:"daily_memory_time_quota"`
	DefaultHostname             string                              `json:"default_hostname"`
	EnableBuiltinLogging        bool                                `json:"enable_builtin_logging"`
	Enabled                     bool                                `json:"enabled"`
	HttpsOnly                   bool                                `json:"https_only"`
	Id                          string                              `json:"id"`
	KeyVaultReferenceIdentityId string                              `json:"key_vault_reference_identity_id"`
	Kind                        string                              `json:"kind"`
	Location                    string                              `json:"location"`
	Name                        string                              `json:"name"`
	OsType                      string                              `json:"os_type"`
	OutboundIpAddresses         string                              `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                              `json:"possible_outbound_ip_addresses"`
	ResourceGroupName           string                              `json:"resource_group_name"`
	StorageAccountAccessKey     string                              `json:"storage_account_access_key"`
	StorageAccountName          string                              `json:"storage_account_name"`
	Tags                        map[string]string                   `json:"tags"`
	Version                     string                              `json:"version"`
	SiteCredential              []functionapp.SiteCredentialState   `json:"site_credential"`
	AuthSettings                []functionapp.AuthSettingsState     `json:"auth_settings"`
	ConnectionString            []functionapp.ConnectionStringState `json:"connection_string"`
	Identity                    []functionapp.IdentityState         `json:"identity"`
	SiteConfig                  []functionapp.SiteConfigState       `json:"site_config"`
	SourceControl               []functionapp.SourceControlState    `json:"source_control"`
	Timeouts                    *functionapp.TimeoutsState          `json:"timeouts"`
}
