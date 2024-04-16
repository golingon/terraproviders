// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_security_center_storage_defender

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

// Resource represents the Terraform resource azurerm_security_center_storage_defender.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSecurityCenterStorageDefenderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ascsd *Resource) Type() string {
	return "azurerm_security_center_storage_defender"
}

// LocalName returns the local name for [Resource].
func (ascsd *Resource) LocalName() string {
	return ascsd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ascsd *Resource) Configuration() interface{} {
	return ascsd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ascsd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ascsd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ascsd *Resource) Dependencies() terra.Dependencies {
	return ascsd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ascsd *Resource) LifecycleManagement() *terra.Lifecycle {
	return ascsd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ascsd *Resource) Attributes() azurermSecurityCenterStorageDefenderAttributes {
	return azurermSecurityCenterStorageDefenderAttributes{ref: terra.ReferenceResource(ascsd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ascsd *Resource) ImportState(state io.Reader) error {
	ascsd.state = &azurermSecurityCenterStorageDefenderState{}
	if err := json.NewDecoder(state).Decode(ascsd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ascsd.Type(), ascsd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ascsd *Resource) State() (*azurermSecurityCenterStorageDefenderState, bool) {
	return ascsd.state, ascsd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ascsd *Resource) StateMust() *azurermSecurityCenterStorageDefenderState {
	if ascsd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ascsd.Type(), ascsd.LocalName()))
	}
	return ascsd.state
}

// Args contains the configurations for azurerm_security_center_storage_defender.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MalwareScanningOnUploadCapGbPerMonth: number, optional
	MalwareScanningOnUploadCapGbPerMonth terra.NumberValue `hcl:"malware_scanning_on_upload_cap_gb_per_month,attr"`
	// MalwareScanningOnUploadEnabled: bool, optional
	MalwareScanningOnUploadEnabled terra.BoolValue `hcl:"malware_scanning_on_upload_enabled,attr"`
	// OverrideSubscriptionSettingsEnabled: bool, optional
	OverrideSubscriptionSettingsEnabled terra.BoolValue `hcl:"override_subscription_settings_enabled,attr"`
	// SensitiveDataDiscoveryEnabled: bool, optional
	SensitiveDataDiscoveryEnabled terra.BoolValue `hcl:"sensitive_data_discovery_enabled,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSecurityCenterStorageDefenderAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_security_center_storage_defender.
func (ascsd azurermSecurityCenterStorageDefenderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ascsd.ref.Append("id"))
}

// MalwareScanningOnUploadCapGbPerMonth returns a reference to field malware_scanning_on_upload_cap_gb_per_month of azurerm_security_center_storage_defender.
func (ascsd azurermSecurityCenterStorageDefenderAttributes) MalwareScanningOnUploadCapGbPerMonth() terra.NumberValue {
	return terra.ReferenceAsNumber(ascsd.ref.Append("malware_scanning_on_upload_cap_gb_per_month"))
}

// MalwareScanningOnUploadEnabled returns a reference to field malware_scanning_on_upload_enabled of azurerm_security_center_storage_defender.
func (ascsd azurermSecurityCenterStorageDefenderAttributes) MalwareScanningOnUploadEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ascsd.ref.Append("malware_scanning_on_upload_enabled"))
}

// OverrideSubscriptionSettingsEnabled returns a reference to field override_subscription_settings_enabled of azurerm_security_center_storage_defender.
func (ascsd azurermSecurityCenterStorageDefenderAttributes) OverrideSubscriptionSettingsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ascsd.ref.Append("override_subscription_settings_enabled"))
}

// SensitiveDataDiscoveryEnabled returns a reference to field sensitive_data_discovery_enabled of azurerm_security_center_storage_defender.
func (ascsd azurermSecurityCenterStorageDefenderAttributes) SensitiveDataDiscoveryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ascsd.ref.Append("sensitive_data_discovery_enabled"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_security_center_storage_defender.
func (ascsd azurermSecurityCenterStorageDefenderAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(ascsd.ref.Append("storage_account_id"))
}

func (ascsd azurermSecurityCenterStorageDefenderAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ascsd.ref.Append("timeouts"))
}

type azurermSecurityCenterStorageDefenderState struct {
	Id                                   string         `json:"id"`
	MalwareScanningOnUploadCapGbPerMonth float64        `json:"malware_scanning_on_upload_cap_gb_per_month"`
	MalwareScanningOnUploadEnabled       bool           `json:"malware_scanning_on_upload_enabled"`
	OverrideSubscriptionSettingsEnabled  bool           `json:"override_subscription_settings_enabled"`
	SensitiveDataDiscoveryEnabled        bool           `json:"sensitive_data_discovery_enabled"`
	StorageAccountId                     string         `json:"storage_account_id"`
	Timeouts                             *TimeoutsState `json:"timeouts"`
}
