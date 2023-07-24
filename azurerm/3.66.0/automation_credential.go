// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationcredential "github.com/golingon/terraproviders/azurerm/3.66.0/automationcredential"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationCredential creates a new instance of [AutomationCredential].
func NewAutomationCredential(name string, args AutomationCredentialArgs) *AutomationCredential {
	return &AutomationCredential{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationCredential)(nil)

// AutomationCredential represents the Terraform resource azurerm_automation_credential.
type AutomationCredential struct {
	Name      string
	Args      AutomationCredentialArgs
	state     *automationCredentialState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationCredential].
func (ac *AutomationCredential) Type() string {
	return "azurerm_automation_credential"
}

// LocalName returns the local name for [AutomationCredential].
func (ac *AutomationCredential) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AutomationCredential].
func (ac *AutomationCredential) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AutomationCredential].
func (ac *AutomationCredential) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AutomationCredential] depends_on.
func (ac *AutomationCredential) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationCredential].
func (ac *AutomationCredential) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AutomationCredential].
func (ac *AutomationCredential) Attributes() automationCredentialAttributes {
	return automationCredentialAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AutomationCredential]'s state.
func (ac *AutomationCredential) ImportState(av io.Reader) error {
	ac.state = &automationCredentialState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationCredential] has state.
func (ac *AutomationCredential) State() (*automationCredentialState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AutomationCredential]. Panics if the state is nil.
func (ac *AutomationCredential) StateMust() *automationCredentialState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AutomationCredentialArgs contains the configurations for azurerm_automation_credential.
type AutomationCredentialArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationcredential.Timeouts `hcl:"timeouts,block"`
}
type automationCredentialAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_credential.
func (ac automationCredentialAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_credential.
func (ac automationCredentialAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_automation_credential.
func (ac automationCredentialAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_credential.
func (ac automationCredentialAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_automation_credential.
func (ac automationCredentialAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_credential.
func (ac automationCredentialAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("resource_group_name"))
}

// Username returns a reference to field username of azurerm_automation_credential.
func (ac automationCredentialAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("username"))
}

func (ac automationCredentialAttributes) Timeouts() automationcredential.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationcredential.TimeoutsAttributes](ac.ref.Append("timeouts"))
}

type automationCredentialState struct {
	AutomationAccountName string                              `json:"automation_account_name"`
	Description           string                              `json:"description"`
	Id                    string                              `json:"id"`
	Name                  string                              `json:"name"`
	Password              string                              `json:"password"`
	ResourceGroupName     string                              `json:"resource_group_name"`
	Username              string                              `json:"username"`
	Timeouts              *automationcredential.TimeoutsState `json:"timeouts"`
}
