// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationmodule "github.com/golingon/terraproviders/azurerm/3.63.0/automationmodule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationModule creates a new instance of [AutomationModule].
func NewAutomationModule(name string, args AutomationModuleArgs) *AutomationModule {
	return &AutomationModule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationModule)(nil)

// AutomationModule represents the Terraform resource azurerm_automation_module.
type AutomationModule struct {
	Name      string
	Args      AutomationModuleArgs
	state     *automationModuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationModule].
func (am *AutomationModule) Type() string {
	return "azurerm_automation_module"
}

// LocalName returns the local name for [AutomationModule].
func (am *AutomationModule) LocalName() string {
	return am.Name
}

// Configuration returns the configuration (args) for [AutomationModule].
func (am *AutomationModule) Configuration() interface{} {
	return am.Args
}

// DependOn is used for other resources to depend on [AutomationModule].
func (am *AutomationModule) DependOn() terra.Reference {
	return terra.ReferenceResource(am)
}

// Dependencies returns the list of resources [AutomationModule] depends_on.
func (am *AutomationModule) Dependencies() terra.Dependencies {
	return am.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationModule].
func (am *AutomationModule) LifecycleManagement() *terra.Lifecycle {
	return am.Lifecycle
}

// Attributes returns the attributes for [AutomationModule].
func (am *AutomationModule) Attributes() automationModuleAttributes {
	return automationModuleAttributes{ref: terra.ReferenceResource(am)}
}

// ImportState imports the given attribute values into [AutomationModule]'s state.
func (am *AutomationModule) ImportState(av io.Reader) error {
	am.state = &automationModuleState{}
	if err := json.NewDecoder(av).Decode(am.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", am.Type(), am.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationModule] has state.
func (am *AutomationModule) State() (*automationModuleState, bool) {
	return am.state, am.state != nil
}

// StateMust returns the state for [AutomationModule]. Panics if the state is nil.
func (am *AutomationModule) StateMust() *automationModuleState {
	if am.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", am.Type(), am.LocalName()))
	}
	return am.state
}

// AutomationModuleArgs contains the configurations for azurerm_automation_module.
type AutomationModuleArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ModuleLink: required
	ModuleLink *automationmodule.ModuleLink `hcl:"module_link,block" validate:"required"`
	// Timeouts: optional
	Timeouts *automationmodule.Timeouts `hcl:"timeouts,block"`
}
type automationModuleAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_module.
func (am automationModuleAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("automation_account_name"))
}

// Id returns a reference to field id of azurerm_automation_module.
func (am automationModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_module.
func (am automationModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_module.
func (am automationModuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("resource_group_name"))
}

func (am automationModuleAttributes) ModuleLink() terra.ListValue[automationmodule.ModuleLinkAttributes] {
	return terra.ReferenceAsList[automationmodule.ModuleLinkAttributes](am.ref.Append("module_link"))
}

func (am automationModuleAttributes) Timeouts() automationmodule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationmodule.TimeoutsAttributes](am.ref.Append("timeouts"))
}

type automationModuleState struct {
	AutomationAccountName string                             `json:"automation_account_name"`
	Id                    string                             `json:"id"`
	Name                  string                             `json:"name"`
	ResourceGroupName     string                             `json:"resource_group_name"`
	ModuleLink            []automationmodule.ModuleLinkState `json:"module_link"`
	Timeouts              *automationmodule.TimeoutsState    `json:"timeouts"`
}
