// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationdscnodeconfiguration "github.com/golingon/terraproviders/azurerm/3.66.0/automationdscnodeconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationDscNodeconfiguration creates a new instance of [AutomationDscNodeconfiguration].
func NewAutomationDscNodeconfiguration(name string, args AutomationDscNodeconfigurationArgs) *AutomationDscNodeconfiguration {
	return &AutomationDscNodeconfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationDscNodeconfiguration)(nil)

// AutomationDscNodeconfiguration represents the Terraform resource azurerm_automation_dsc_nodeconfiguration.
type AutomationDscNodeconfiguration struct {
	Name      string
	Args      AutomationDscNodeconfigurationArgs
	state     *automationDscNodeconfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationDscNodeconfiguration].
func (adn *AutomationDscNodeconfiguration) Type() string {
	return "azurerm_automation_dsc_nodeconfiguration"
}

// LocalName returns the local name for [AutomationDscNodeconfiguration].
func (adn *AutomationDscNodeconfiguration) LocalName() string {
	return adn.Name
}

// Configuration returns the configuration (args) for [AutomationDscNodeconfiguration].
func (adn *AutomationDscNodeconfiguration) Configuration() interface{} {
	return adn.Args
}

// DependOn is used for other resources to depend on [AutomationDscNodeconfiguration].
func (adn *AutomationDscNodeconfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(adn)
}

// Dependencies returns the list of resources [AutomationDscNodeconfiguration] depends_on.
func (adn *AutomationDscNodeconfiguration) Dependencies() terra.Dependencies {
	return adn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationDscNodeconfiguration].
func (adn *AutomationDscNodeconfiguration) LifecycleManagement() *terra.Lifecycle {
	return adn.Lifecycle
}

// Attributes returns the attributes for [AutomationDscNodeconfiguration].
func (adn *AutomationDscNodeconfiguration) Attributes() automationDscNodeconfigurationAttributes {
	return automationDscNodeconfigurationAttributes{ref: terra.ReferenceResource(adn)}
}

// ImportState imports the given attribute values into [AutomationDscNodeconfiguration]'s state.
func (adn *AutomationDscNodeconfiguration) ImportState(av io.Reader) error {
	adn.state = &automationDscNodeconfigurationState{}
	if err := json.NewDecoder(av).Decode(adn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adn.Type(), adn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationDscNodeconfiguration] has state.
func (adn *AutomationDscNodeconfiguration) State() (*automationDscNodeconfigurationState, bool) {
	return adn.state, adn.state != nil
}

// StateMust returns the state for [AutomationDscNodeconfiguration]. Panics if the state is nil.
func (adn *AutomationDscNodeconfiguration) StateMust() *automationDscNodeconfigurationState {
	if adn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adn.Type(), adn.LocalName()))
	}
	return adn.state
}

// AutomationDscNodeconfigurationArgs contains the configurations for azurerm_automation_dsc_nodeconfiguration.
type AutomationDscNodeconfigurationArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// ContentEmbedded: string, required
	ContentEmbedded terra.StringValue `hcl:"content_embedded,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationdscnodeconfiguration.Timeouts `hcl:"timeouts,block"`
}
type automationDscNodeconfigurationAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_dsc_nodeconfiguration.
func (adn automationDscNodeconfigurationAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(adn.ref.Append("automation_account_name"))
}

// ConfigurationName returns a reference to field configuration_name of azurerm_automation_dsc_nodeconfiguration.
func (adn automationDscNodeconfigurationAttributes) ConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(adn.ref.Append("configuration_name"))
}

// ContentEmbedded returns a reference to field content_embedded of azurerm_automation_dsc_nodeconfiguration.
func (adn automationDscNodeconfigurationAttributes) ContentEmbedded() terra.StringValue {
	return terra.ReferenceAsString(adn.ref.Append("content_embedded"))
}

// Id returns a reference to field id of azurerm_automation_dsc_nodeconfiguration.
func (adn automationDscNodeconfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adn.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_dsc_nodeconfiguration.
func (adn automationDscNodeconfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adn.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_dsc_nodeconfiguration.
func (adn automationDscNodeconfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adn.ref.Append("resource_group_name"))
}

func (adn automationDscNodeconfigurationAttributes) Timeouts() automationdscnodeconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationdscnodeconfiguration.TimeoutsAttributes](adn.ref.Append("timeouts"))
}

type automationDscNodeconfigurationState struct {
	AutomationAccountName string                                        `json:"automation_account_name"`
	ConfigurationName     string                                        `json:"configuration_name"`
	ContentEmbedded       string                                        `json:"content_embedded"`
	Id                    string                                        `json:"id"`
	Name                  string                                        `json:"name"`
	ResourceGroupName     string                                        `json:"resource_group_name"`
	Timeouts              *automationdscnodeconfiguration.TimeoutsState `json:"timeouts"`
}
