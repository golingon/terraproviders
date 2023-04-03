// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationcertificate "github.com/golingon/terraproviders/azurerm/3.49.0/automationcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationCertificate creates a new instance of [AutomationCertificate].
func NewAutomationCertificate(name string, args AutomationCertificateArgs) *AutomationCertificate {
	return &AutomationCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationCertificate)(nil)

// AutomationCertificate represents the Terraform resource azurerm_automation_certificate.
type AutomationCertificate struct {
	Name      string
	Args      AutomationCertificateArgs
	state     *automationCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationCertificate].
func (ac *AutomationCertificate) Type() string {
	return "azurerm_automation_certificate"
}

// LocalName returns the local name for [AutomationCertificate].
func (ac *AutomationCertificate) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AutomationCertificate].
func (ac *AutomationCertificate) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AutomationCertificate].
func (ac *AutomationCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AutomationCertificate] depends_on.
func (ac *AutomationCertificate) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationCertificate].
func (ac *AutomationCertificate) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AutomationCertificate].
func (ac *AutomationCertificate) Attributes() automationCertificateAttributes {
	return automationCertificateAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AutomationCertificate]'s state.
func (ac *AutomationCertificate) ImportState(av io.Reader) error {
	ac.state = &automationCertificateState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationCertificate] has state.
func (ac *AutomationCertificate) State() (*automationCertificateState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AutomationCertificate]. Panics if the state is nil.
func (ac *AutomationCertificate) StateMust() *automationCertificateState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AutomationCertificateArgs contains the configurations for azurerm_automation_certificate.
type AutomationCertificateArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Base64: string, required
	Base64 terra.StringValue `hcl:"base64,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Exportable: bool, optional
	Exportable terra.BoolValue `hcl:"exportable,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationcertificate.Timeouts `hcl:"timeouts,block"`
}
type automationCertificateAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_certificate.
func (ac automationCertificateAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("automation_account_name"))
}

// Base64 returns a reference to field base64 of azurerm_automation_certificate.
func (ac automationCertificateAttributes) Base64() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("base64"))
}

// Description returns a reference to field description of azurerm_automation_certificate.
func (ac automationCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("description"))
}

// Exportable returns a reference to field exportable of azurerm_automation_certificate.
func (ac automationCertificateAttributes) Exportable() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("exportable"))
}

// Id returns a reference to field id of azurerm_automation_certificate.
func (ac automationCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_certificate.
func (ac automationCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_certificate.
func (ac automationCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("resource_group_name"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_automation_certificate.
func (ac automationCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("thumbprint"))
}

func (ac automationCertificateAttributes) Timeouts() automationcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationcertificate.TimeoutsAttributes](ac.ref.Append("timeouts"))
}

type automationCertificateState struct {
	AutomationAccountName string                               `json:"automation_account_name"`
	Base64                string                               `json:"base64"`
	Description           string                               `json:"description"`
	Exportable            bool                                 `json:"exportable"`
	Id                    string                               `json:"id"`
	Name                  string                               `json:"name"`
	ResourceGroupName     string                               `json:"resource_group_name"`
	Thumbprint            string                               `json:"thumbprint"`
	Timeouts              *automationcertificate.TimeoutsState `json:"timeouts"`
}
