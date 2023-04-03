// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	functionappslot "github.com/golingon/terraproviders/azurerm/3.49.0/functionappslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFunctionAppSlot creates a new instance of [FunctionAppSlot].
func NewFunctionAppSlot(name string, args FunctionAppSlotArgs) *FunctionAppSlot {
	return &FunctionAppSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FunctionAppSlot)(nil)

// FunctionAppSlot represents the Terraform resource azurerm_function_app_slot.
type FunctionAppSlot struct {
	Name      string
	Args      FunctionAppSlotArgs
	state     *functionAppSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FunctionAppSlot].
func (fas *FunctionAppSlot) Type() string {
	return "azurerm_function_app_slot"
}

// LocalName returns the local name for [FunctionAppSlot].
func (fas *FunctionAppSlot) LocalName() string {
	return fas.Name
}

// Configuration returns the configuration (args) for [FunctionAppSlot].
func (fas *FunctionAppSlot) Configuration() interface{} {
	return fas.Args
}

// DependOn is used for other resources to depend on [FunctionAppSlot].
func (fas *FunctionAppSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(fas)
}

// Dependencies returns the list of resources [FunctionAppSlot] depends_on.
func (fas *FunctionAppSlot) Dependencies() terra.Dependencies {
	return fas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FunctionAppSlot].
func (fas *FunctionAppSlot) LifecycleManagement() *terra.Lifecycle {
	return fas.Lifecycle
}

// Attributes returns the attributes for [FunctionAppSlot].
func (fas *FunctionAppSlot) Attributes() functionAppSlotAttributes {
	return functionAppSlotAttributes{ref: terra.ReferenceResource(fas)}
}

// ImportState imports the given attribute values into [FunctionAppSlot]'s state.
func (fas *FunctionAppSlot) ImportState(av io.Reader) error {
	fas.state = &functionAppSlotState{}
	if err := json.NewDecoder(av).Decode(fas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fas.Type(), fas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FunctionAppSlot] has state.
func (fas *FunctionAppSlot) State() (*functionAppSlotState, bool) {
	return fas.state, fas.state != nil
}

// StateMust returns the state for [FunctionAppSlot]. Panics if the state is nil.
func (fas *FunctionAppSlot) StateMust() *functionAppSlotState {
	if fas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fas.Type(), fas.LocalName()))
	}
	return fas.state
}

// FunctionAppSlotArgs contains the configurations for azurerm_function_app_slot.
type FunctionAppSlotArgs struct {
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
	// SiteCredential: min=0
	SiteCredential []functionappslot.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *functionappslot.AuthSettings `hcl:"auth_settings,block"`
	// ConnectionString: min=0
	ConnectionString []functionappslot.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *functionappslot.Identity `hcl:"identity,block"`
	// SiteConfig: optional
	SiteConfig *functionappslot.SiteConfig `hcl:"site_config,block"`
	// Timeouts: optional
	Timeouts *functionappslot.Timeouts `hcl:"timeouts,block"`
}
type functionAppSlotAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fas.ref.Append("app_settings"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(fas.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("default_hostname"))
}

// EnableBuiltinLogging returns a reference to field enable_builtin_logging of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) EnableBuiltinLogging() terra.BoolValue {
	return terra.ReferenceAsBool(fas.ref.Append("enable_builtin_logging"))
}

// Enabled returns a reference to field enabled of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(fas.ref.Append("enabled"))
}

// FunctionAppName returns a reference to field function_app_name of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) FunctionAppName() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("function_app_name"))
}

// HttpsOnly returns a reference to field https_only of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(fas.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("os_type"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("resource_group_name"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("storage_account_name"))
}

// Tags returns a reference to field tags of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fas.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_function_app_slot.
func (fas functionAppSlotAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(fas.ref.Append("version"))
}

func (fas functionAppSlotAttributes) SiteCredential() terra.ListValue[functionappslot.SiteCredentialAttributes] {
	return terra.ReferenceAsList[functionappslot.SiteCredentialAttributes](fas.ref.Append("site_credential"))
}

func (fas functionAppSlotAttributes) AuthSettings() terra.ListValue[functionappslot.AuthSettingsAttributes] {
	return terra.ReferenceAsList[functionappslot.AuthSettingsAttributes](fas.ref.Append("auth_settings"))
}

func (fas functionAppSlotAttributes) ConnectionString() terra.SetValue[functionappslot.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[functionappslot.ConnectionStringAttributes](fas.ref.Append("connection_string"))
}

func (fas functionAppSlotAttributes) Identity() terra.ListValue[functionappslot.IdentityAttributes] {
	return terra.ReferenceAsList[functionappslot.IdentityAttributes](fas.ref.Append("identity"))
}

func (fas functionAppSlotAttributes) SiteConfig() terra.ListValue[functionappslot.SiteConfigAttributes] {
	return terra.ReferenceAsList[functionappslot.SiteConfigAttributes](fas.ref.Append("site_config"))
}

func (fas functionAppSlotAttributes) Timeouts() functionappslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[functionappslot.TimeoutsAttributes](fas.ref.Append("timeouts"))
}

type functionAppSlotState struct {
	AppServicePlanId            string                                  `json:"app_service_plan_id"`
	AppSettings                 map[string]string                       `json:"app_settings"`
	DailyMemoryTimeQuota        float64                                 `json:"daily_memory_time_quota"`
	DefaultHostname             string                                  `json:"default_hostname"`
	EnableBuiltinLogging        bool                                    `json:"enable_builtin_logging"`
	Enabled                     bool                                    `json:"enabled"`
	FunctionAppName             string                                  `json:"function_app_name"`
	HttpsOnly                   bool                                    `json:"https_only"`
	Id                          string                                  `json:"id"`
	Kind                        string                                  `json:"kind"`
	Location                    string                                  `json:"location"`
	Name                        string                                  `json:"name"`
	OsType                      string                                  `json:"os_type"`
	OutboundIpAddresses         string                                  `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                                  `json:"possible_outbound_ip_addresses"`
	ResourceGroupName           string                                  `json:"resource_group_name"`
	StorageAccountAccessKey     string                                  `json:"storage_account_access_key"`
	StorageAccountName          string                                  `json:"storage_account_name"`
	Tags                        map[string]string                       `json:"tags"`
	Version                     string                                  `json:"version"`
	SiteCredential              []functionappslot.SiteCredentialState   `json:"site_credential"`
	AuthSettings                []functionappslot.AuthSettingsState     `json:"auth_settings"`
	ConnectionString            []functionappslot.ConnectionStringState `json:"connection_string"`
	Identity                    []functionappslot.IdentityState         `json:"identity"`
	SiteConfig                  []functionappslot.SiteConfigState       `json:"site_config"`
	Timeouts                    *functionappslot.TimeoutsState          `json:"timeouts"`
}