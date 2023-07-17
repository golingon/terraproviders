// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycenterautomation "github.com/golingon/terraproviders/azurerm/3.65.0/securitycenterautomation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterAutomation creates a new instance of [SecurityCenterAutomation].
func NewSecurityCenterAutomation(name string, args SecurityCenterAutomationArgs) *SecurityCenterAutomation {
	return &SecurityCenterAutomation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterAutomation)(nil)

// SecurityCenterAutomation represents the Terraform resource azurerm_security_center_automation.
type SecurityCenterAutomation struct {
	Name      string
	Args      SecurityCenterAutomationArgs
	state     *securityCenterAutomationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterAutomation].
func (sca *SecurityCenterAutomation) Type() string {
	return "azurerm_security_center_automation"
}

// LocalName returns the local name for [SecurityCenterAutomation].
func (sca *SecurityCenterAutomation) LocalName() string {
	return sca.Name
}

// Configuration returns the configuration (args) for [SecurityCenterAutomation].
func (sca *SecurityCenterAutomation) Configuration() interface{} {
	return sca.Args
}

// DependOn is used for other resources to depend on [SecurityCenterAutomation].
func (sca *SecurityCenterAutomation) DependOn() terra.Reference {
	return terra.ReferenceResource(sca)
}

// Dependencies returns the list of resources [SecurityCenterAutomation] depends_on.
func (sca *SecurityCenterAutomation) Dependencies() terra.Dependencies {
	return sca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterAutomation].
func (sca *SecurityCenterAutomation) LifecycleManagement() *terra.Lifecycle {
	return sca.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterAutomation].
func (sca *SecurityCenterAutomation) Attributes() securityCenterAutomationAttributes {
	return securityCenterAutomationAttributes{ref: terra.ReferenceResource(sca)}
}

// ImportState imports the given attribute values into [SecurityCenterAutomation]'s state.
func (sca *SecurityCenterAutomation) ImportState(av io.Reader) error {
	sca.state = &securityCenterAutomationState{}
	if err := json.NewDecoder(av).Decode(sca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sca.Type(), sca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterAutomation] has state.
func (sca *SecurityCenterAutomation) State() (*securityCenterAutomationState, bool) {
	return sca.state, sca.state != nil
}

// StateMust returns the state for [SecurityCenterAutomation]. Panics if the state is nil.
func (sca *SecurityCenterAutomation) StateMust() *securityCenterAutomationState {
	if sca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sca.Type(), sca.LocalName()))
	}
	return sca.state
}

// SecurityCenterAutomationArgs contains the configurations for azurerm_security_center_automation.
type SecurityCenterAutomationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Scopes: list of string, required
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Action: min=1
	Action []securitycenterautomation.Action `hcl:"action,block" validate:"min=1"`
	// Source: min=1
	Source []securitycenterautomation.Source `hcl:"source,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *securitycenterautomation.Timeouts `hcl:"timeouts,block"`
}
type securityCenterAutomationAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("resource_group_name"))
}

// Scopes returns a reference to field scopes of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sca.ref.Append("scopes"))
}

// Tags returns a reference to field tags of azurerm_security_center_automation.
func (sca securityCenterAutomationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sca.ref.Append("tags"))
}

func (sca securityCenterAutomationAttributes) Action() terra.ListValue[securitycenterautomation.ActionAttributes] {
	return terra.ReferenceAsList[securitycenterautomation.ActionAttributes](sca.ref.Append("action"))
}

func (sca securityCenterAutomationAttributes) Source() terra.ListValue[securitycenterautomation.SourceAttributes] {
	return terra.ReferenceAsList[securitycenterautomation.SourceAttributes](sca.ref.Append("source"))
}

func (sca securityCenterAutomationAttributes) Timeouts() securitycenterautomation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycenterautomation.TimeoutsAttributes](sca.ref.Append("timeouts"))
}

type securityCenterAutomationState struct {
	Description       string                                  `json:"description"`
	Enabled           bool                                    `json:"enabled"`
	Id                string                                  `json:"id"`
	Location          string                                  `json:"location"`
	Name              string                                  `json:"name"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	Scopes            []string                                `json:"scopes"`
	Tags              map[string]string                       `json:"tags"`
	Action            []securitycenterautomation.ActionState  `json:"action"`
	Source            []securitycenterautomation.SourceState  `json:"source"`
	Timeouts          *securitycenterautomation.TimeoutsState `json:"timeouts"`
}
