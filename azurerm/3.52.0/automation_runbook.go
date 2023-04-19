// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationrunbook "github.com/golingon/terraproviders/azurerm/3.52.0/automationrunbook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationRunbook creates a new instance of [AutomationRunbook].
func NewAutomationRunbook(name string, args AutomationRunbookArgs) *AutomationRunbook {
	return &AutomationRunbook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationRunbook)(nil)

// AutomationRunbook represents the Terraform resource azurerm_automation_runbook.
type AutomationRunbook struct {
	Name      string
	Args      AutomationRunbookArgs
	state     *automationRunbookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationRunbook].
func (ar *AutomationRunbook) Type() string {
	return "azurerm_automation_runbook"
}

// LocalName returns the local name for [AutomationRunbook].
func (ar *AutomationRunbook) LocalName() string {
	return ar.Name
}

// Configuration returns the configuration (args) for [AutomationRunbook].
func (ar *AutomationRunbook) Configuration() interface{} {
	return ar.Args
}

// DependOn is used for other resources to depend on [AutomationRunbook].
func (ar *AutomationRunbook) DependOn() terra.Reference {
	return terra.ReferenceResource(ar)
}

// Dependencies returns the list of resources [AutomationRunbook] depends_on.
func (ar *AutomationRunbook) Dependencies() terra.Dependencies {
	return ar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationRunbook].
func (ar *AutomationRunbook) LifecycleManagement() *terra.Lifecycle {
	return ar.Lifecycle
}

// Attributes returns the attributes for [AutomationRunbook].
func (ar *AutomationRunbook) Attributes() automationRunbookAttributes {
	return automationRunbookAttributes{ref: terra.ReferenceResource(ar)}
}

// ImportState imports the given attribute values into [AutomationRunbook]'s state.
func (ar *AutomationRunbook) ImportState(av io.Reader) error {
	ar.state = &automationRunbookState{}
	if err := json.NewDecoder(av).Decode(ar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ar.Type(), ar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationRunbook] has state.
func (ar *AutomationRunbook) State() (*automationRunbookState, bool) {
	return ar.state, ar.state != nil
}

// StateMust returns the state for [AutomationRunbook]. Panics if the state is nil.
func (ar *AutomationRunbook) StateMust() *automationRunbookState {
	if ar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ar.Type(), ar.LocalName()))
	}
	return ar.state
}

// AutomationRunbookArgs contains the configurations for azurerm_automation_runbook.
type AutomationRunbookArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LogActivityTraceLevel: number, optional
	LogActivityTraceLevel terra.NumberValue `hcl:"log_activity_trace_level,attr"`
	// LogProgress: bool, required
	LogProgress terra.BoolValue `hcl:"log_progress,attr" validate:"required"`
	// LogVerbose: bool, required
	LogVerbose terra.BoolValue `hcl:"log_verbose,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RunbookType: string, required
	RunbookType terra.StringValue `hcl:"runbook_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// JobSchedule: min=0
	JobSchedule []automationrunbook.JobSchedule `hcl:"job_schedule,block" validate:"min=0"`
	// Draft: optional
	Draft *automationrunbook.Draft `hcl:"draft,block"`
	// PublishContentLink: optional
	PublishContentLink *automationrunbook.PublishContentLink `hcl:"publish_content_link,block"`
	// Timeouts: optional
	Timeouts *automationrunbook.Timeouts `hcl:"timeouts,block"`
}
type automationRunbookAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_runbook.
func (ar automationRunbookAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("automation_account_name"))
}

// Content returns a reference to field content of azurerm_automation_runbook.
func (ar automationRunbookAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("content"))
}

// Description returns a reference to field description of azurerm_automation_runbook.
func (ar automationRunbookAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_automation_runbook.
func (ar automationRunbookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_automation_runbook.
func (ar automationRunbookAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("location"))
}

// LogActivityTraceLevel returns a reference to field log_activity_trace_level of azurerm_automation_runbook.
func (ar automationRunbookAttributes) LogActivityTraceLevel() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("log_activity_trace_level"))
}

// LogProgress returns a reference to field log_progress of azurerm_automation_runbook.
func (ar automationRunbookAttributes) LogProgress() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("log_progress"))
}

// LogVerbose returns a reference to field log_verbose of azurerm_automation_runbook.
func (ar automationRunbookAttributes) LogVerbose() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("log_verbose"))
}

// Name returns a reference to field name of azurerm_automation_runbook.
func (ar automationRunbookAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_runbook.
func (ar automationRunbookAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("resource_group_name"))
}

// RunbookType returns a reference to field runbook_type of azurerm_automation_runbook.
func (ar automationRunbookAttributes) RunbookType() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("runbook_type"))
}

// Tags returns a reference to field tags of azurerm_automation_runbook.
func (ar automationRunbookAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ar.ref.Append("tags"))
}

func (ar automationRunbookAttributes) JobSchedule() terra.SetValue[automationrunbook.JobScheduleAttributes] {
	return terra.ReferenceAsSet[automationrunbook.JobScheduleAttributes](ar.ref.Append("job_schedule"))
}

func (ar automationRunbookAttributes) Draft() terra.ListValue[automationrunbook.DraftAttributes] {
	return terra.ReferenceAsList[automationrunbook.DraftAttributes](ar.ref.Append("draft"))
}

func (ar automationRunbookAttributes) PublishContentLink() terra.ListValue[automationrunbook.PublishContentLinkAttributes] {
	return terra.ReferenceAsList[automationrunbook.PublishContentLinkAttributes](ar.ref.Append("publish_content_link"))
}

func (ar automationRunbookAttributes) Timeouts() automationrunbook.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationrunbook.TimeoutsAttributes](ar.ref.Append("timeouts"))
}

type automationRunbookState struct {
	AutomationAccountName string                                      `json:"automation_account_name"`
	Content               string                                      `json:"content"`
	Description           string                                      `json:"description"`
	Id                    string                                      `json:"id"`
	Location              string                                      `json:"location"`
	LogActivityTraceLevel float64                                     `json:"log_activity_trace_level"`
	LogProgress           bool                                        `json:"log_progress"`
	LogVerbose            bool                                        `json:"log_verbose"`
	Name                  string                                      `json:"name"`
	ResourceGroupName     string                                      `json:"resource_group_name"`
	RunbookType           string                                      `json:"runbook_type"`
	Tags                  map[string]string                           `json:"tags"`
	JobSchedule           []automationrunbook.JobScheduleState        `json:"job_schedule"`
	Draft                 []automationrunbook.DraftState              `json:"draft"`
	PublishContentLink    []automationrunbook.PublishContentLinkState `json:"publish_content_link"`
	Timeouts              *automationrunbook.TimeoutsState            `json:"timeouts"`
}
