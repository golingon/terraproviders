// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationhybridrunbookworkergroup "github.com/golingon/terraproviders/azurerm/3.55.0/automationhybridrunbookworkergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationHybridRunbookWorkerGroup creates a new instance of [AutomationHybridRunbookWorkerGroup].
func NewAutomationHybridRunbookWorkerGroup(name string, args AutomationHybridRunbookWorkerGroupArgs) *AutomationHybridRunbookWorkerGroup {
	return &AutomationHybridRunbookWorkerGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationHybridRunbookWorkerGroup)(nil)

// AutomationHybridRunbookWorkerGroup represents the Terraform resource azurerm_automation_hybrid_runbook_worker_group.
type AutomationHybridRunbookWorkerGroup struct {
	Name      string
	Args      AutomationHybridRunbookWorkerGroupArgs
	state     *automationHybridRunbookWorkerGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationHybridRunbookWorkerGroup].
func (ahrwg *AutomationHybridRunbookWorkerGroup) Type() string {
	return "azurerm_automation_hybrid_runbook_worker_group"
}

// LocalName returns the local name for [AutomationHybridRunbookWorkerGroup].
func (ahrwg *AutomationHybridRunbookWorkerGroup) LocalName() string {
	return ahrwg.Name
}

// Configuration returns the configuration (args) for [AutomationHybridRunbookWorkerGroup].
func (ahrwg *AutomationHybridRunbookWorkerGroup) Configuration() interface{} {
	return ahrwg.Args
}

// DependOn is used for other resources to depend on [AutomationHybridRunbookWorkerGroup].
func (ahrwg *AutomationHybridRunbookWorkerGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ahrwg)
}

// Dependencies returns the list of resources [AutomationHybridRunbookWorkerGroup] depends_on.
func (ahrwg *AutomationHybridRunbookWorkerGroup) Dependencies() terra.Dependencies {
	return ahrwg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationHybridRunbookWorkerGroup].
func (ahrwg *AutomationHybridRunbookWorkerGroup) LifecycleManagement() *terra.Lifecycle {
	return ahrwg.Lifecycle
}

// Attributes returns the attributes for [AutomationHybridRunbookWorkerGroup].
func (ahrwg *AutomationHybridRunbookWorkerGroup) Attributes() automationHybridRunbookWorkerGroupAttributes {
	return automationHybridRunbookWorkerGroupAttributes{ref: terra.ReferenceResource(ahrwg)}
}

// ImportState imports the given attribute values into [AutomationHybridRunbookWorkerGroup]'s state.
func (ahrwg *AutomationHybridRunbookWorkerGroup) ImportState(av io.Reader) error {
	ahrwg.state = &automationHybridRunbookWorkerGroupState{}
	if err := json.NewDecoder(av).Decode(ahrwg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ahrwg.Type(), ahrwg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationHybridRunbookWorkerGroup] has state.
func (ahrwg *AutomationHybridRunbookWorkerGroup) State() (*automationHybridRunbookWorkerGroupState, bool) {
	return ahrwg.state, ahrwg.state != nil
}

// StateMust returns the state for [AutomationHybridRunbookWorkerGroup]. Panics if the state is nil.
func (ahrwg *AutomationHybridRunbookWorkerGroup) StateMust() *automationHybridRunbookWorkerGroupState {
	if ahrwg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ahrwg.Type(), ahrwg.LocalName()))
	}
	return ahrwg.state
}

// AutomationHybridRunbookWorkerGroupArgs contains the configurations for azurerm_automation_hybrid_runbook_worker_group.
type AutomationHybridRunbookWorkerGroupArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// CredentialName: string, optional
	CredentialName terra.StringValue `hcl:"credential_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationhybridrunbookworkergroup.Timeouts `hcl:"timeouts,block"`
}
type automationHybridRunbookWorkerGroupAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_hybrid_runbook_worker_group.
func (ahrwg automationHybridRunbookWorkerGroupAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(ahrwg.ref.Append("automation_account_name"))
}

// CredentialName returns a reference to field credential_name of azurerm_automation_hybrid_runbook_worker_group.
func (ahrwg automationHybridRunbookWorkerGroupAttributes) CredentialName() terra.StringValue {
	return terra.ReferenceAsString(ahrwg.ref.Append("credential_name"))
}

// Id returns a reference to field id of azurerm_automation_hybrid_runbook_worker_group.
func (ahrwg automationHybridRunbookWorkerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ahrwg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_hybrid_runbook_worker_group.
func (ahrwg automationHybridRunbookWorkerGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ahrwg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_hybrid_runbook_worker_group.
func (ahrwg automationHybridRunbookWorkerGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ahrwg.ref.Append("resource_group_name"))
}

func (ahrwg automationHybridRunbookWorkerGroupAttributes) Timeouts() automationhybridrunbookworkergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationhybridrunbookworkergroup.TimeoutsAttributes](ahrwg.ref.Append("timeouts"))
}

type automationHybridRunbookWorkerGroupState struct {
	AutomationAccountName string                                            `json:"automation_account_name"`
	CredentialName        string                                            `json:"credential_name"`
	Id                    string                                            `json:"id"`
	Name                  string                                            `json:"name"`
	ResourceGroupName     string                                            `json:"resource_group_name"`
	Timeouts              *automationhybridrunbookworkergroup.TimeoutsState `json:"timeouts"`
}
