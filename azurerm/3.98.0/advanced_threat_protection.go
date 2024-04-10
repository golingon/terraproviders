// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	advancedthreatprotection "github.com/golingon/terraproviders/azurerm/3.98.0/advancedthreatprotection"
	"io"
)

// NewAdvancedThreatProtection creates a new instance of [AdvancedThreatProtection].
func NewAdvancedThreatProtection(name string, args AdvancedThreatProtectionArgs) *AdvancedThreatProtection {
	return &AdvancedThreatProtection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AdvancedThreatProtection)(nil)

// AdvancedThreatProtection represents the Terraform resource azurerm_advanced_threat_protection.
type AdvancedThreatProtection struct {
	Name      string
	Args      AdvancedThreatProtectionArgs
	state     *advancedThreatProtectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AdvancedThreatProtection].
func (atp *AdvancedThreatProtection) Type() string {
	return "azurerm_advanced_threat_protection"
}

// LocalName returns the local name for [AdvancedThreatProtection].
func (atp *AdvancedThreatProtection) LocalName() string {
	return atp.Name
}

// Configuration returns the configuration (args) for [AdvancedThreatProtection].
func (atp *AdvancedThreatProtection) Configuration() interface{} {
	return atp.Args
}

// DependOn is used for other resources to depend on [AdvancedThreatProtection].
func (atp *AdvancedThreatProtection) DependOn() terra.Reference {
	return terra.ReferenceResource(atp)
}

// Dependencies returns the list of resources [AdvancedThreatProtection] depends_on.
func (atp *AdvancedThreatProtection) Dependencies() terra.Dependencies {
	return atp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AdvancedThreatProtection].
func (atp *AdvancedThreatProtection) LifecycleManagement() *terra.Lifecycle {
	return atp.Lifecycle
}

// Attributes returns the attributes for [AdvancedThreatProtection].
func (atp *AdvancedThreatProtection) Attributes() advancedThreatProtectionAttributes {
	return advancedThreatProtectionAttributes{ref: terra.ReferenceResource(atp)}
}

// ImportState imports the given attribute values into [AdvancedThreatProtection]'s state.
func (atp *AdvancedThreatProtection) ImportState(av io.Reader) error {
	atp.state = &advancedThreatProtectionState{}
	if err := json.NewDecoder(av).Decode(atp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", atp.Type(), atp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AdvancedThreatProtection] has state.
func (atp *AdvancedThreatProtection) State() (*advancedThreatProtectionState, bool) {
	return atp.state, atp.state != nil
}

// StateMust returns the state for [AdvancedThreatProtection]. Panics if the state is nil.
func (atp *AdvancedThreatProtection) StateMust() *advancedThreatProtectionState {
	if atp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", atp.Type(), atp.LocalName()))
	}
	return atp.state
}

// AdvancedThreatProtectionArgs contains the configurations for azurerm_advanced_threat_protection.
type AdvancedThreatProtectionArgs struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *advancedthreatprotection.Timeouts `hcl:"timeouts,block"`
}
type advancedThreatProtectionAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_advanced_threat_protection.
func (atp advancedThreatProtectionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(atp.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_advanced_threat_protection.
func (atp advancedThreatProtectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(atp.ref.Append("id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_advanced_threat_protection.
func (atp advancedThreatProtectionAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(atp.ref.Append("target_resource_id"))
}

func (atp advancedThreatProtectionAttributes) Timeouts() advancedthreatprotection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[advancedthreatprotection.TimeoutsAttributes](atp.ref.Append("timeouts"))
}

type advancedThreatProtectionState struct {
	Enabled          bool                                    `json:"enabled"`
	Id               string                                  `json:"id"`
	TargetResourceId string                                  `json:"target_resource_id"`
	Timeouts         *advancedthreatprotection.TimeoutsState `json:"timeouts"`
}
