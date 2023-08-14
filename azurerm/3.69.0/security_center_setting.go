// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycentersetting "github.com/golingon/terraproviders/azurerm/3.69.0/securitycentersetting"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterSetting creates a new instance of [SecurityCenterSetting].
func NewSecurityCenterSetting(name string, args SecurityCenterSettingArgs) *SecurityCenterSetting {
	return &SecurityCenterSetting{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterSetting)(nil)

// SecurityCenterSetting represents the Terraform resource azurerm_security_center_setting.
type SecurityCenterSetting struct {
	Name      string
	Args      SecurityCenterSettingArgs
	state     *securityCenterSettingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterSetting].
func (scs *SecurityCenterSetting) Type() string {
	return "azurerm_security_center_setting"
}

// LocalName returns the local name for [SecurityCenterSetting].
func (scs *SecurityCenterSetting) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [SecurityCenterSetting].
func (scs *SecurityCenterSetting) Configuration() interface{} {
	return scs.Args
}

// DependOn is used for other resources to depend on [SecurityCenterSetting].
func (scs *SecurityCenterSetting) DependOn() terra.Reference {
	return terra.ReferenceResource(scs)
}

// Dependencies returns the list of resources [SecurityCenterSetting] depends_on.
func (scs *SecurityCenterSetting) Dependencies() terra.Dependencies {
	return scs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterSetting].
func (scs *SecurityCenterSetting) LifecycleManagement() *terra.Lifecycle {
	return scs.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterSetting].
func (scs *SecurityCenterSetting) Attributes() securityCenterSettingAttributes {
	return securityCenterSettingAttributes{ref: terra.ReferenceResource(scs)}
}

// ImportState imports the given attribute values into [SecurityCenterSetting]'s state.
func (scs *SecurityCenterSetting) ImportState(av io.Reader) error {
	scs.state = &securityCenterSettingState{}
	if err := json.NewDecoder(av).Decode(scs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scs.Type(), scs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterSetting] has state.
func (scs *SecurityCenterSetting) State() (*securityCenterSettingState, bool) {
	return scs.state, scs.state != nil
}

// StateMust returns the state for [SecurityCenterSetting]. Panics if the state is nil.
func (scs *SecurityCenterSetting) StateMust() *securityCenterSettingState {
	if scs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scs.Type(), scs.LocalName()))
	}
	return scs.state
}

// SecurityCenterSettingArgs contains the configurations for azurerm_security_center_setting.
type SecurityCenterSettingArgs struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SettingName: string, required
	SettingName terra.StringValue `hcl:"setting_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *securitycentersetting.Timeouts `hcl:"timeouts,block"`
}
type securityCenterSettingAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_security_center_setting.
func (scs securityCenterSettingAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(scs.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_security_center_setting.
func (scs securityCenterSettingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// SettingName returns a reference to field setting_name of azurerm_security_center_setting.
func (scs securityCenterSettingAttributes) SettingName() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("setting_name"))
}

func (scs securityCenterSettingAttributes) Timeouts() securitycentersetting.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycentersetting.TimeoutsAttributes](scs.ref.Append("timeouts"))
}

type securityCenterSettingState struct {
	Enabled     bool                                 `json:"enabled"`
	Id          string                               `json:"id"`
	SettingName string                               `json:"setting_name"`
	Timeouts    *securitycentersetting.TimeoutsState `json:"timeouts"`
}
