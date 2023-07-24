// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationconnectioncertificate "github.com/golingon/terraproviders/azurerm/3.66.0/automationconnectioncertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationConnectionCertificate creates a new instance of [AutomationConnectionCertificate].
func NewAutomationConnectionCertificate(name string, args AutomationConnectionCertificateArgs) *AutomationConnectionCertificate {
	return &AutomationConnectionCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationConnectionCertificate)(nil)

// AutomationConnectionCertificate represents the Terraform resource azurerm_automation_connection_certificate.
type AutomationConnectionCertificate struct {
	Name      string
	Args      AutomationConnectionCertificateArgs
	state     *automationConnectionCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationConnectionCertificate].
func (acc *AutomationConnectionCertificate) Type() string {
	return "azurerm_automation_connection_certificate"
}

// LocalName returns the local name for [AutomationConnectionCertificate].
func (acc *AutomationConnectionCertificate) LocalName() string {
	return acc.Name
}

// Configuration returns the configuration (args) for [AutomationConnectionCertificate].
func (acc *AutomationConnectionCertificate) Configuration() interface{} {
	return acc.Args
}

// DependOn is used for other resources to depend on [AutomationConnectionCertificate].
func (acc *AutomationConnectionCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(acc)
}

// Dependencies returns the list of resources [AutomationConnectionCertificate] depends_on.
func (acc *AutomationConnectionCertificate) Dependencies() terra.Dependencies {
	return acc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationConnectionCertificate].
func (acc *AutomationConnectionCertificate) LifecycleManagement() *terra.Lifecycle {
	return acc.Lifecycle
}

// Attributes returns the attributes for [AutomationConnectionCertificate].
func (acc *AutomationConnectionCertificate) Attributes() automationConnectionCertificateAttributes {
	return automationConnectionCertificateAttributes{ref: terra.ReferenceResource(acc)}
}

// ImportState imports the given attribute values into [AutomationConnectionCertificate]'s state.
func (acc *AutomationConnectionCertificate) ImportState(av io.Reader) error {
	acc.state = &automationConnectionCertificateState{}
	if err := json.NewDecoder(av).Decode(acc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acc.Type(), acc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationConnectionCertificate] has state.
func (acc *AutomationConnectionCertificate) State() (*automationConnectionCertificateState, bool) {
	return acc.state, acc.state != nil
}

// StateMust returns the state for [AutomationConnectionCertificate]. Panics if the state is nil.
func (acc *AutomationConnectionCertificate) StateMust() *automationConnectionCertificateState {
	if acc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acc.Type(), acc.LocalName()))
	}
	return acc.state
}

// AutomationConnectionCertificateArgs contains the configurations for azurerm_automation_connection_certificate.
type AutomationConnectionCertificateArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// AutomationCertificateName: string, required
	AutomationCertificateName terra.StringValue `hcl:"automation_certificate_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationconnectioncertificate.Timeouts `hcl:"timeouts,block"`
}
type automationConnectionCertificateAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("automation_account_name"))
}

// AutomationCertificateName returns a reference to field automation_certificate_name of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) AutomationCertificateName() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("automation_certificate_name"))
}

// Description returns a reference to field description of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("resource_group_name"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_automation_connection_certificate.
func (acc automationConnectionCertificateAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(acc.ref.Append("subscription_id"))
}

func (acc automationConnectionCertificateAttributes) Timeouts() automationconnectioncertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationconnectioncertificate.TimeoutsAttributes](acc.ref.Append("timeouts"))
}

type automationConnectionCertificateState struct {
	AutomationAccountName     string                                         `json:"automation_account_name"`
	AutomationCertificateName string                                         `json:"automation_certificate_name"`
	Description               string                                         `json:"description"`
	Id                        string                                         `json:"id"`
	Name                      string                                         `json:"name"`
	ResourceGroupName         string                                         `json:"resource_group_name"`
	SubscriptionId            string                                         `json:"subscription_id"`
	Timeouts                  *automationconnectioncertificate.TimeoutsState `json:"timeouts"`
}
