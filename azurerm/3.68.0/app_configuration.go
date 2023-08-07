// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appconfiguration "github.com/golingon/terraproviders/azurerm/3.68.0/appconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppConfiguration creates a new instance of [AppConfiguration].
func NewAppConfiguration(name string, args AppConfigurationArgs) *AppConfiguration {
	return &AppConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppConfiguration)(nil)

// AppConfiguration represents the Terraform resource azurerm_app_configuration.
type AppConfiguration struct {
	Name      string
	Args      AppConfigurationArgs
	state     *appConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppConfiguration].
func (ac *AppConfiguration) Type() string {
	return "azurerm_app_configuration"
}

// LocalName returns the local name for [AppConfiguration].
func (ac *AppConfiguration) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AppConfiguration].
func (ac *AppConfiguration) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AppConfiguration].
func (ac *AppConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AppConfiguration] depends_on.
func (ac *AppConfiguration) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppConfiguration].
func (ac *AppConfiguration) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AppConfiguration].
func (ac *AppConfiguration) Attributes() appConfigurationAttributes {
	return appConfigurationAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AppConfiguration]'s state.
func (ac *AppConfiguration) ImportState(av io.Reader) error {
	ac.state = &appConfigurationState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppConfiguration] has state.
func (ac *AppConfiguration) State() (*appConfigurationState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AppConfiguration]. Panics if the state is nil.
func (ac *AppConfiguration) StateMust() *appConfigurationState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AppConfigurationArgs contains the configurations for azurerm_app_configuration.
type AppConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalAuthEnabled: bool, optional
	LocalAuthEnabled terra.BoolValue `hcl:"local_auth_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccess: string, optional
	PublicNetworkAccess terra.StringValue `hcl:"public_network_access,attr"`
	// PurgeProtectionEnabled: bool, optional
	PurgeProtectionEnabled terra.BoolValue `hcl:"purge_protection_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// SoftDeleteRetentionDays: number, optional
	SoftDeleteRetentionDays terra.NumberValue `hcl:"soft_delete_retention_days,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PrimaryReadKey: min=0
	PrimaryReadKey []appconfiguration.PrimaryReadKey `hcl:"primary_read_key,block" validate:"min=0"`
	// PrimaryWriteKey: min=0
	PrimaryWriteKey []appconfiguration.PrimaryWriteKey `hcl:"primary_write_key,block" validate:"min=0"`
	// SecondaryReadKey: min=0
	SecondaryReadKey []appconfiguration.SecondaryReadKey `hcl:"secondary_read_key,block" validate:"min=0"`
	// SecondaryWriteKey: min=0
	SecondaryWriteKey []appconfiguration.SecondaryWriteKey `hcl:"secondary_write_key,block" validate:"min=0"`
	// Encryption: optional
	Encryption *appconfiguration.Encryption `hcl:"encryption,block"`
	// Identity: optional
	Identity *appconfiguration.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *appconfiguration.Timeouts `hcl:"timeouts,block"`
}
type appConfigurationAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_app_configuration.
func (ac appConfigurationAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_app_configuration.
func (ac appConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_app_configuration.
func (ac appConfigurationAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_app_configuration.
func (ac appConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_configuration.
func (ac appConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// PublicNetworkAccess returns a reference to field public_network_access of azurerm_app_configuration.
func (ac appConfigurationAttributes) PublicNetworkAccess() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("public_network_access"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_app_configuration.
func (ac appConfigurationAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_configuration.
func (ac appConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_app_configuration.
func (ac appConfigurationAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("sku"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_app_configuration.
func (ac appConfigurationAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_app_configuration.
func (ac appConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

func (ac appConfigurationAttributes) PrimaryReadKey() terra.ListValue[appconfiguration.PrimaryReadKeyAttributes] {
	return terra.ReferenceAsList[appconfiguration.PrimaryReadKeyAttributes](ac.ref.Append("primary_read_key"))
}

func (ac appConfigurationAttributes) PrimaryWriteKey() terra.ListValue[appconfiguration.PrimaryWriteKeyAttributes] {
	return terra.ReferenceAsList[appconfiguration.PrimaryWriteKeyAttributes](ac.ref.Append("primary_write_key"))
}

func (ac appConfigurationAttributes) SecondaryReadKey() terra.ListValue[appconfiguration.SecondaryReadKeyAttributes] {
	return terra.ReferenceAsList[appconfiguration.SecondaryReadKeyAttributes](ac.ref.Append("secondary_read_key"))
}

func (ac appConfigurationAttributes) SecondaryWriteKey() terra.ListValue[appconfiguration.SecondaryWriteKeyAttributes] {
	return terra.ReferenceAsList[appconfiguration.SecondaryWriteKeyAttributes](ac.ref.Append("secondary_write_key"))
}

func (ac appConfigurationAttributes) Encryption() terra.ListValue[appconfiguration.EncryptionAttributes] {
	return terra.ReferenceAsList[appconfiguration.EncryptionAttributes](ac.ref.Append("encryption"))
}

func (ac appConfigurationAttributes) Identity() terra.ListValue[appconfiguration.IdentityAttributes] {
	return terra.ReferenceAsList[appconfiguration.IdentityAttributes](ac.ref.Append("identity"))
}

func (ac appConfigurationAttributes) Timeouts() appconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appconfiguration.TimeoutsAttributes](ac.ref.Append("timeouts"))
}

type appConfigurationState struct {
	Endpoint                string                                    `json:"endpoint"`
	Id                      string                                    `json:"id"`
	LocalAuthEnabled        bool                                      `json:"local_auth_enabled"`
	Location                string                                    `json:"location"`
	Name                    string                                    `json:"name"`
	PublicNetworkAccess     string                                    `json:"public_network_access"`
	PurgeProtectionEnabled  bool                                      `json:"purge_protection_enabled"`
	ResourceGroupName       string                                    `json:"resource_group_name"`
	Sku                     string                                    `json:"sku"`
	SoftDeleteRetentionDays float64                                   `json:"soft_delete_retention_days"`
	Tags                    map[string]string                         `json:"tags"`
	PrimaryReadKey          []appconfiguration.PrimaryReadKeyState    `json:"primary_read_key"`
	PrimaryWriteKey         []appconfiguration.PrimaryWriteKeyState   `json:"primary_write_key"`
	SecondaryReadKey        []appconfiguration.SecondaryReadKeyState  `json:"secondary_read_key"`
	SecondaryWriteKey       []appconfiguration.SecondaryWriteKeyState `json:"secondary_write_key"`
	Encryption              []appconfiguration.EncryptionState        `json:"encryption"`
	Identity                []appconfiguration.IdentityState          `json:"identity"`
	Timeouts                *appconfiguration.TimeoutsState           `json:"timeouts"`
}
