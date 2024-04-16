// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_function_app_slot

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_function_app_slot.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermFunctionAppSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (afas *Resource) Type() string {
	return "azurerm_function_app_slot"
}

// LocalName returns the local name for [Resource].
func (afas *Resource) LocalName() string {
	return afas.Name
}

// Configuration returns the configuration (args) for [Resource].
func (afas *Resource) Configuration() interface{} {
	return afas.Args
}

// DependOn is used for other resources to depend on [Resource].
func (afas *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(afas)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (afas *Resource) Dependencies() terra.Dependencies {
	return afas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (afas *Resource) LifecycleManagement() *terra.Lifecycle {
	return afas.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (afas *Resource) Attributes() azurermFunctionAppSlotAttributes {
	return azurermFunctionAppSlotAttributes{ref: terra.ReferenceResource(afas)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (afas *Resource) ImportState(state io.Reader) error {
	afas.state = &azurermFunctionAppSlotState{}
	if err := json.NewDecoder(state).Decode(afas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afas.Type(), afas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (afas *Resource) State() (*azurermFunctionAppSlotState, bool) {
	return afas.state, afas.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (afas *Resource) StateMust() *azurermFunctionAppSlotState {
	if afas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afas.Type(), afas.LocalName()))
	}
	return afas.state
}

// Args contains the configurations for azurerm_function_app_slot.
type Args struct {
	// AppServicePlanId: string, required
	AppServicePlanId terra.StringValue `hcl:"app_service_plan_id,attr" validate:"required"`
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// DailyMemoryTimeQuota: number, optional
	DailyMemoryTimeQuota terra.NumberValue `hcl:"daily_memory_time_quota,attr"`
	// EnableBuiltinLogging: bool, optional
	EnableBuiltinLogging terra.BoolValue `hcl:"enable_builtin_logging,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FunctionAppName: string, required
	FunctionAppName terra.StringValue `hcl:"function_app_name,attr" validate:"required"`
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
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
	// AuthSettings: optional
	AuthSettings *AuthSettings `hcl:"auth_settings,block"`
	// ConnectionString: min=0
	ConnectionString []ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// SiteConfig: optional
	SiteConfig *SiteConfig `hcl:"site_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermFunctionAppSlotAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afas.ref.Append("app_settings"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(afas.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("default_hostname"))
}

// EnableBuiltinLogging returns a reference to field enable_builtin_logging of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) EnableBuiltinLogging() terra.BoolValue {
	return terra.ReferenceAsBool(afas.ref.Append("enable_builtin_logging"))
}

// Enabled returns a reference to field enabled of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(afas.ref.Append("enabled"))
}

// FunctionAppName returns a reference to field function_app_name of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) FunctionAppName() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("function_app_name"))
}

// HttpsOnly returns a reference to field https_only of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(afas.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("os_type"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("resource_group_name"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("storage_account_name"))
}

// Tags returns a reference to field tags of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afas.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_function_app_slot.
func (afas azurermFunctionAppSlotAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(afas.ref.Append("version"))
}

func (afas azurermFunctionAppSlotAttributes) SiteCredential() terra.ListValue[SiteCredentialAttributes] {
	return terra.ReferenceAsList[SiteCredentialAttributes](afas.ref.Append("site_credential"))
}

func (afas azurermFunctionAppSlotAttributes) AuthSettings() terra.ListValue[AuthSettingsAttributes] {
	return terra.ReferenceAsList[AuthSettingsAttributes](afas.ref.Append("auth_settings"))
}

func (afas azurermFunctionAppSlotAttributes) ConnectionString() terra.SetValue[ConnectionStringAttributes] {
	return terra.ReferenceAsSet[ConnectionStringAttributes](afas.ref.Append("connection_string"))
}

func (afas azurermFunctionAppSlotAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](afas.ref.Append("identity"))
}

func (afas azurermFunctionAppSlotAttributes) SiteConfig() terra.ListValue[SiteConfigAttributes] {
	return terra.ReferenceAsList[SiteConfigAttributes](afas.ref.Append("site_config"))
}

func (afas azurermFunctionAppSlotAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](afas.ref.Append("timeouts"))
}

type azurermFunctionAppSlotState struct {
	AppServicePlanId            string                  `json:"app_service_plan_id"`
	AppSettings                 map[string]string       `json:"app_settings"`
	DailyMemoryTimeQuota        float64                 `json:"daily_memory_time_quota"`
	DefaultHostname             string                  `json:"default_hostname"`
	EnableBuiltinLogging        bool                    `json:"enable_builtin_logging"`
	Enabled                     bool                    `json:"enabled"`
	FunctionAppName             string                  `json:"function_app_name"`
	HttpsOnly                   bool                    `json:"https_only"`
	Id                          string                  `json:"id"`
	Kind                        string                  `json:"kind"`
	Location                    string                  `json:"location"`
	Name                        string                  `json:"name"`
	OsType                      string                  `json:"os_type"`
	OutboundIpAddresses         string                  `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                  `json:"possible_outbound_ip_addresses"`
	ResourceGroupName           string                  `json:"resource_group_name"`
	StorageAccountAccessKey     string                  `json:"storage_account_access_key"`
	StorageAccountName          string                  `json:"storage_account_name"`
	Tags                        map[string]string       `json:"tags"`
	Version                     string                  `json:"version"`
	SiteCredential              []SiteCredentialState   `json:"site_credential"`
	AuthSettings                []AuthSettingsState     `json:"auth_settings"`
	ConnectionString            []ConnectionStringState `json:"connection_string"`
	Identity                    []IdentityState         `json:"identity"`
	SiteConfig                  []SiteConfigState       `json:"site_config"`
	Timeouts                    *TimeoutsState          `json:"timeouts"`
}
