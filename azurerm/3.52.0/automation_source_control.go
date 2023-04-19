// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationsourcecontrol "github.com/golingon/terraproviders/azurerm/3.52.0/automationsourcecontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationSourceControl creates a new instance of [AutomationSourceControl].
func NewAutomationSourceControl(name string, args AutomationSourceControlArgs) *AutomationSourceControl {
	return &AutomationSourceControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationSourceControl)(nil)

// AutomationSourceControl represents the Terraform resource azurerm_automation_source_control.
type AutomationSourceControl struct {
	Name      string
	Args      AutomationSourceControlArgs
	state     *automationSourceControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationSourceControl].
func (asc *AutomationSourceControl) Type() string {
	return "azurerm_automation_source_control"
}

// LocalName returns the local name for [AutomationSourceControl].
func (asc *AutomationSourceControl) LocalName() string {
	return asc.Name
}

// Configuration returns the configuration (args) for [AutomationSourceControl].
func (asc *AutomationSourceControl) Configuration() interface{} {
	return asc.Args
}

// DependOn is used for other resources to depend on [AutomationSourceControl].
func (asc *AutomationSourceControl) DependOn() terra.Reference {
	return terra.ReferenceResource(asc)
}

// Dependencies returns the list of resources [AutomationSourceControl] depends_on.
func (asc *AutomationSourceControl) Dependencies() terra.Dependencies {
	return asc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationSourceControl].
func (asc *AutomationSourceControl) LifecycleManagement() *terra.Lifecycle {
	return asc.Lifecycle
}

// Attributes returns the attributes for [AutomationSourceControl].
func (asc *AutomationSourceControl) Attributes() automationSourceControlAttributes {
	return automationSourceControlAttributes{ref: terra.ReferenceResource(asc)}
}

// ImportState imports the given attribute values into [AutomationSourceControl]'s state.
func (asc *AutomationSourceControl) ImportState(av io.Reader) error {
	asc.state = &automationSourceControlState{}
	if err := json.NewDecoder(av).Decode(asc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asc.Type(), asc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationSourceControl] has state.
func (asc *AutomationSourceControl) State() (*automationSourceControlState, bool) {
	return asc.state, asc.state != nil
}

// StateMust returns the state for [AutomationSourceControl]. Panics if the state is nil.
func (asc *AutomationSourceControl) StateMust() *automationSourceControlState {
	if asc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asc.Type(), asc.LocalName()))
	}
	return asc.state
}

// AutomationSourceControlArgs contains the configurations for azurerm_automation_source_control.
type AutomationSourceControlArgs struct {
	// AutomaticSync: bool, optional
	AutomaticSync terra.BoolValue `hcl:"automatic_sync,attr"`
	// AutomationAccountId: string, required
	AutomationAccountId terra.StringValue `hcl:"automation_account_id,attr" validate:"required"`
	// Branch: string, optional
	Branch terra.StringValue `hcl:"branch,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FolderPath: string, required
	FolderPath terra.StringValue `hcl:"folder_path,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublishRunbookEnabled: bool, optional
	PublishRunbookEnabled terra.BoolValue `hcl:"publish_runbook_enabled,attr"`
	// RepositoryUrl: string, required
	RepositoryUrl terra.StringValue `hcl:"repository_url,attr" validate:"required"`
	// SourceControlType: string, required
	SourceControlType terra.StringValue `hcl:"source_control_type,attr" validate:"required"`
	// Security: required
	Security *automationsourcecontrol.Security `hcl:"security,block" validate:"required"`
	// Timeouts: optional
	Timeouts *automationsourcecontrol.Timeouts `hcl:"timeouts,block"`
}
type automationSourceControlAttributes struct {
	ref terra.Reference
}

// AutomaticSync returns a reference to field automatic_sync of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) AutomaticSync() terra.BoolValue {
	return terra.ReferenceAsBool(asc.ref.Append("automatic_sync"))
}

// AutomationAccountId returns a reference to field automation_account_id of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) AutomationAccountId() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("automation_account_id"))
}

// Branch returns a reference to field branch of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("branch"))
}

// Description returns a reference to field description of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("description"))
}

// FolderPath returns a reference to field folder_path of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) FolderPath() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("folder_path"))
}

// Id returns a reference to field id of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("name"))
}

// PublishRunbookEnabled returns a reference to field publish_runbook_enabled of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) PublishRunbookEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asc.ref.Append("publish_runbook_enabled"))
}

// RepositoryUrl returns a reference to field repository_url of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("repository_url"))
}

// SourceControlType returns a reference to field source_control_type of azurerm_automation_source_control.
func (asc automationSourceControlAttributes) SourceControlType() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("source_control_type"))
}

func (asc automationSourceControlAttributes) Security() terra.ListValue[automationsourcecontrol.SecurityAttributes] {
	return terra.ReferenceAsList[automationsourcecontrol.SecurityAttributes](asc.ref.Append("security"))
}

func (asc automationSourceControlAttributes) Timeouts() automationsourcecontrol.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationsourcecontrol.TimeoutsAttributes](asc.ref.Append("timeouts"))
}

type automationSourceControlState struct {
	AutomaticSync         bool                                    `json:"automatic_sync"`
	AutomationAccountId   string                                  `json:"automation_account_id"`
	Branch                string                                  `json:"branch"`
	Description           string                                  `json:"description"`
	FolderPath            string                                  `json:"folder_path"`
	Id                    string                                  `json:"id"`
	Name                  string                                  `json:"name"`
	PublishRunbookEnabled bool                                    `json:"publish_runbook_enabled"`
	RepositoryUrl         string                                  `json:"repository_url"`
	SourceControlType     string                                  `json:"source_control_type"`
	Security              []automationsourcecontrol.SecurityState `json:"security"`
	Timeouts              *automationsourcecontrol.TimeoutsState  `json:"timeouts"`
}
