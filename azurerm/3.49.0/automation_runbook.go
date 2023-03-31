// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationrunbook "github.com/golingon/terraproviders/azurerm/3.49.0/automationrunbook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAutomationRunbook(name string, args AutomationRunbookArgs) *AutomationRunbook {
	return &AutomationRunbook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationRunbook)(nil)

type AutomationRunbook struct {
	Name  string
	Args  AutomationRunbookArgs
	state *automationRunbookState
}

func (ar *AutomationRunbook) Type() string {
	return "azurerm_automation_runbook"
}

func (ar *AutomationRunbook) LocalName() string {
	return ar.Name
}

func (ar *AutomationRunbook) Configuration() interface{} {
	return ar.Args
}

func (ar *AutomationRunbook) Attributes() automationRunbookAttributes {
	return automationRunbookAttributes{ref: terra.ReferenceResource(ar)}
}

func (ar *AutomationRunbook) ImportState(av io.Reader) error {
	ar.state = &automationRunbookState{}
	if err := json.NewDecoder(av).Decode(ar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ar.Type(), ar.LocalName(), err)
	}
	return nil
}

func (ar *AutomationRunbook) State() (*automationRunbookState, bool) {
	return ar.state, ar.state != nil
}

func (ar *AutomationRunbook) StateMust() *automationRunbookState {
	if ar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ar.Type(), ar.LocalName()))
	}
	return ar.state
}

func (ar *AutomationRunbook) DependOn() terra.Reference {
	return terra.ReferenceResource(ar)
}

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
	// DependsOn contains resources that AutomationRunbook depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type automationRunbookAttributes struct {
	ref terra.Reference
}

func (ar automationRunbookAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("automation_account_name"))
}

func (ar automationRunbookAttributes) Content() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("content"))
}

func (ar automationRunbookAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("description"))
}

func (ar automationRunbookAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("id"))
}

func (ar automationRunbookAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("location"))
}

func (ar automationRunbookAttributes) LogActivityTraceLevel() terra.NumberValue {
	return terra.ReferenceNumber(ar.ref.Append("log_activity_trace_level"))
}

func (ar automationRunbookAttributes) LogProgress() terra.BoolValue {
	return terra.ReferenceBool(ar.ref.Append("log_progress"))
}

func (ar automationRunbookAttributes) LogVerbose() terra.BoolValue {
	return terra.ReferenceBool(ar.ref.Append("log_verbose"))
}

func (ar automationRunbookAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("name"))
}

func (ar automationRunbookAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("resource_group_name"))
}

func (ar automationRunbookAttributes) RunbookType() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("runbook_type"))
}

func (ar automationRunbookAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ar.ref.Append("tags"))
}

func (ar automationRunbookAttributes) JobSchedule() terra.SetValue[automationrunbook.JobScheduleAttributes] {
	return terra.ReferenceSet[automationrunbook.JobScheduleAttributes](ar.ref.Append("job_schedule"))
}

func (ar automationRunbookAttributes) Draft() terra.ListValue[automationrunbook.DraftAttributes] {
	return terra.ReferenceList[automationrunbook.DraftAttributes](ar.ref.Append("draft"))
}

func (ar automationRunbookAttributes) PublishContentLink() terra.ListValue[automationrunbook.PublishContentLinkAttributes] {
	return terra.ReferenceList[automationrunbook.PublishContentLinkAttributes](ar.ref.Append("publish_content_link"))
}

func (ar automationRunbookAttributes) Timeouts() automationrunbook.TimeoutsAttributes {
	return terra.ReferenceSingle[automationrunbook.TimeoutsAttributes](ar.ref.Append("timeouts"))
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
