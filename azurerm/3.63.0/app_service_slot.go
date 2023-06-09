// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceslot "github.com/golingon/terraproviders/azurerm/3.63.0/appserviceslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceSlot creates a new instance of [AppServiceSlot].
func NewAppServiceSlot(name string, args AppServiceSlotArgs) *AppServiceSlot {
	return &AppServiceSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSlot)(nil)

// AppServiceSlot represents the Terraform resource azurerm_app_service_slot.
type AppServiceSlot struct {
	Name      string
	Args      AppServiceSlotArgs
	state     *appServiceSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceSlot].
func (ass *AppServiceSlot) Type() string {
	return "azurerm_app_service_slot"
}

// LocalName returns the local name for [AppServiceSlot].
func (ass *AppServiceSlot) LocalName() string {
	return ass.Name
}

// Configuration returns the configuration (args) for [AppServiceSlot].
func (ass *AppServiceSlot) Configuration() interface{} {
	return ass.Args
}

// DependOn is used for other resources to depend on [AppServiceSlot].
func (ass *AppServiceSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(ass)
}

// Dependencies returns the list of resources [AppServiceSlot] depends_on.
func (ass *AppServiceSlot) Dependencies() terra.Dependencies {
	return ass.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceSlot].
func (ass *AppServiceSlot) LifecycleManagement() *terra.Lifecycle {
	return ass.Lifecycle
}

// Attributes returns the attributes for [AppServiceSlot].
func (ass *AppServiceSlot) Attributes() appServiceSlotAttributes {
	return appServiceSlotAttributes{ref: terra.ReferenceResource(ass)}
}

// ImportState imports the given attribute values into [AppServiceSlot]'s state.
func (ass *AppServiceSlot) ImportState(av io.Reader) error {
	ass.state = &appServiceSlotState{}
	if err := json.NewDecoder(av).Decode(ass.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ass.Type(), ass.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceSlot] has state.
func (ass *AppServiceSlot) State() (*appServiceSlotState, bool) {
	return ass.state, ass.state != nil
}

// StateMust returns the state for [AppServiceSlot]. Panics if the state is nil.
func (ass *AppServiceSlot) StateMust() *appServiceSlotState {
	if ass.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ass.Type(), ass.LocalName()))
	}
	return ass.state
}

// AppServiceSlotArgs contains the configurations for azurerm_app_service_slot.
type AppServiceSlotArgs struct {
	// AppServiceName: string, required
	AppServiceName terra.StringValue `hcl:"app_service_name,attr" validate:"required"`
	// AppServicePlanId: string, required
	AppServicePlanId terra.StringValue `hcl:"app_service_plan_id,attr" validate:"required"`
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// ClientAffinityEnabled: bool, optional
	ClientAffinityEnabled terra.BoolValue `hcl:"client_affinity_enabled,attr"`
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
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// SiteCredential: min=0
	SiteCredential []appserviceslot.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *appserviceslot.AuthSettings `hcl:"auth_settings,block"`
	// ConnectionString: min=0
	ConnectionString []appserviceslot.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *appserviceslot.Identity `hcl:"identity,block"`
	// Logs: optional
	Logs *appserviceslot.Logs `hcl:"logs,block"`
	// SiteConfig: optional
	SiteConfig *appserviceslot.SiteConfig `hcl:"site_config,block"`
	// StorageAccount: min=0
	StorageAccount []appserviceslot.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *appserviceslot.Timeouts `hcl:"timeouts,block"`
}
type appServiceSlotAttributes struct {
	ref terra.Reference
}

// AppServiceName returns a reference to field app_service_name of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) AppServiceName() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("app_service_name"))
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("client_affinity_enabled"))
}

// DefaultSiteHostname returns a reference to field default_site_hostname of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) DefaultSiteHostname() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("default_site_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("key_vault_reference_identity_id"))
}

// Location returns a reference to field location of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_app_service_slot.
func (ass appServiceSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("tags"))
}

func (ass appServiceSlotAttributes) SiteCredential() terra.ListValue[appserviceslot.SiteCredentialAttributes] {
	return terra.ReferenceAsList[appserviceslot.SiteCredentialAttributes](ass.ref.Append("site_credential"))
}

func (ass appServiceSlotAttributes) AuthSettings() terra.ListValue[appserviceslot.AuthSettingsAttributes] {
	return terra.ReferenceAsList[appserviceslot.AuthSettingsAttributes](ass.ref.Append("auth_settings"))
}

func (ass appServiceSlotAttributes) ConnectionString() terra.SetValue[appserviceslot.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[appserviceslot.ConnectionStringAttributes](ass.ref.Append("connection_string"))
}

func (ass appServiceSlotAttributes) Identity() terra.ListValue[appserviceslot.IdentityAttributes] {
	return terra.ReferenceAsList[appserviceslot.IdentityAttributes](ass.ref.Append("identity"))
}

func (ass appServiceSlotAttributes) Logs() terra.ListValue[appserviceslot.LogsAttributes] {
	return terra.ReferenceAsList[appserviceslot.LogsAttributes](ass.ref.Append("logs"))
}

func (ass appServiceSlotAttributes) SiteConfig() terra.ListValue[appserviceslot.SiteConfigAttributes] {
	return terra.ReferenceAsList[appserviceslot.SiteConfigAttributes](ass.ref.Append("site_config"))
}

func (ass appServiceSlotAttributes) StorageAccount() terra.SetValue[appserviceslot.StorageAccountAttributes] {
	return terra.ReferenceAsSet[appserviceslot.StorageAccountAttributes](ass.ref.Append("storage_account"))
}

func (ass appServiceSlotAttributes) Timeouts() appserviceslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceslot.TimeoutsAttributes](ass.ref.Append("timeouts"))
}

type appServiceSlotState struct {
	AppServiceName              string                                 `json:"app_service_name"`
	AppServicePlanId            string                                 `json:"app_service_plan_id"`
	AppSettings                 map[string]string                      `json:"app_settings"`
	ClientAffinityEnabled       bool                                   `json:"client_affinity_enabled"`
	DefaultSiteHostname         string                                 `json:"default_site_hostname"`
	Enabled                     bool                                   `json:"enabled"`
	HttpsOnly                   bool                                   `json:"https_only"`
	Id                          string                                 `json:"id"`
	KeyVaultReferenceIdentityId string                                 `json:"key_vault_reference_identity_id"`
	Location                    string                                 `json:"location"`
	Name                        string                                 `json:"name"`
	ResourceGroupName           string                                 `json:"resource_group_name"`
	Tags                        map[string]string                      `json:"tags"`
	SiteCredential              []appserviceslot.SiteCredentialState   `json:"site_credential"`
	AuthSettings                []appserviceslot.AuthSettingsState     `json:"auth_settings"`
	ConnectionString            []appserviceslot.ConnectionStringState `json:"connection_string"`
	Identity                    []appserviceslot.IdentityState         `json:"identity"`
	Logs                        []appserviceslot.LogsState             `json:"logs"`
	SiteConfig                  []appserviceslot.SiteConfigState       `json:"site_config"`
	StorageAccount              []appserviceslot.StorageAccountState   `json:"storage_account"`
	Timeouts                    *appserviceslot.TimeoutsState          `json:"timeouts"`
}
