// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationconnectionserviceprincipal "github.com/golingon/terraproviders/azurerm/3.64.0/automationconnectionserviceprincipal"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationConnectionServicePrincipal creates a new instance of [AutomationConnectionServicePrincipal].
func NewAutomationConnectionServicePrincipal(name string, args AutomationConnectionServicePrincipalArgs) *AutomationConnectionServicePrincipal {
	return &AutomationConnectionServicePrincipal{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationConnectionServicePrincipal)(nil)

// AutomationConnectionServicePrincipal represents the Terraform resource azurerm_automation_connection_service_principal.
type AutomationConnectionServicePrincipal struct {
	Name      string
	Args      AutomationConnectionServicePrincipalArgs
	state     *automationConnectionServicePrincipalState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationConnectionServicePrincipal].
func (acsp *AutomationConnectionServicePrincipal) Type() string {
	return "azurerm_automation_connection_service_principal"
}

// LocalName returns the local name for [AutomationConnectionServicePrincipal].
func (acsp *AutomationConnectionServicePrincipal) LocalName() string {
	return acsp.Name
}

// Configuration returns the configuration (args) for [AutomationConnectionServicePrincipal].
func (acsp *AutomationConnectionServicePrincipal) Configuration() interface{} {
	return acsp.Args
}

// DependOn is used for other resources to depend on [AutomationConnectionServicePrincipal].
func (acsp *AutomationConnectionServicePrincipal) DependOn() terra.Reference {
	return terra.ReferenceResource(acsp)
}

// Dependencies returns the list of resources [AutomationConnectionServicePrincipal] depends_on.
func (acsp *AutomationConnectionServicePrincipal) Dependencies() terra.Dependencies {
	return acsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationConnectionServicePrincipal].
func (acsp *AutomationConnectionServicePrincipal) LifecycleManagement() *terra.Lifecycle {
	return acsp.Lifecycle
}

// Attributes returns the attributes for [AutomationConnectionServicePrincipal].
func (acsp *AutomationConnectionServicePrincipal) Attributes() automationConnectionServicePrincipalAttributes {
	return automationConnectionServicePrincipalAttributes{ref: terra.ReferenceResource(acsp)}
}

// ImportState imports the given attribute values into [AutomationConnectionServicePrincipal]'s state.
func (acsp *AutomationConnectionServicePrincipal) ImportState(av io.Reader) error {
	acsp.state = &automationConnectionServicePrincipalState{}
	if err := json.NewDecoder(av).Decode(acsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acsp.Type(), acsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationConnectionServicePrincipal] has state.
func (acsp *AutomationConnectionServicePrincipal) State() (*automationConnectionServicePrincipalState, bool) {
	return acsp.state, acsp.state != nil
}

// StateMust returns the state for [AutomationConnectionServicePrincipal]. Panics if the state is nil.
func (acsp *AutomationConnectionServicePrincipal) StateMust() *automationConnectionServicePrincipalState {
	if acsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acsp.Type(), acsp.LocalName()))
	}
	return acsp.state
}

// AutomationConnectionServicePrincipalArgs contains the configurations for azurerm_automation_connection_service_principal.
type AutomationConnectionServicePrincipalArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// CertificateThumbprint: string, required
	CertificateThumbprint terra.StringValue `hcl:"certificate_thumbprint,attr" validate:"required"`
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
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationconnectionserviceprincipal.Timeouts `hcl:"timeouts,block"`
}
type automationConnectionServicePrincipalAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("application_id"))
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("automation_account_name"))
}

// CertificateThumbprint returns a reference to field certificate_thumbprint of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) CertificateThumbprint() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("certificate_thumbprint"))
}

// Description returns a reference to field description of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("resource_group_name"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("subscription_id"))
}

// TenantId returns a reference to field tenant_id of azurerm_automation_connection_service_principal.
func (acsp automationConnectionServicePrincipalAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("tenant_id"))
}

func (acsp automationConnectionServicePrincipalAttributes) Timeouts() automationconnectionserviceprincipal.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationconnectionserviceprincipal.TimeoutsAttributes](acsp.ref.Append("timeouts"))
}

type automationConnectionServicePrincipalState struct {
	ApplicationId         string                                              `json:"application_id"`
	AutomationAccountName string                                              `json:"automation_account_name"`
	CertificateThumbprint string                                              `json:"certificate_thumbprint"`
	Description           string                                              `json:"description"`
	Id                    string                                              `json:"id"`
	Name                  string                                              `json:"name"`
	ResourceGroupName     string                                              `json:"resource_group_name"`
	SubscriptionId        string                                              `json:"subscription_id"`
	TenantId              string                                              `json:"tenant_id"`
	Timeouts              *automationconnectionserviceprincipal.TimeoutsState `json:"timeouts"`
}
