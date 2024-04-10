// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	loganalyticssolution "github.com/golingon/terraproviders/azurerm/3.98.0/loganalyticssolution"
	"io"
)

// NewLogAnalyticsSolution creates a new instance of [LogAnalyticsSolution].
func NewLogAnalyticsSolution(name string, args LogAnalyticsSolutionArgs) *LogAnalyticsSolution {
	return &LogAnalyticsSolution{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsSolution)(nil)

// LogAnalyticsSolution represents the Terraform resource azurerm_log_analytics_solution.
type LogAnalyticsSolution struct {
	Name      string
	Args      LogAnalyticsSolutionArgs
	state     *logAnalyticsSolutionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsSolution].
func (las *LogAnalyticsSolution) Type() string {
	return "azurerm_log_analytics_solution"
}

// LocalName returns the local name for [LogAnalyticsSolution].
func (las *LogAnalyticsSolution) LocalName() string {
	return las.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsSolution].
func (las *LogAnalyticsSolution) Configuration() interface{} {
	return las.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsSolution].
func (las *LogAnalyticsSolution) DependOn() terra.Reference {
	return terra.ReferenceResource(las)
}

// Dependencies returns the list of resources [LogAnalyticsSolution] depends_on.
func (las *LogAnalyticsSolution) Dependencies() terra.Dependencies {
	return las.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsSolution].
func (las *LogAnalyticsSolution) LifecycleManagement() *terra.Lifecycle {
	return las.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsSolution].
func (las *LogAnalyticsSolution) Attributes() logAnalyticsSolutionAttributes {
	return logAnalyticsSolutionAttributes{ref: terra.ReferenceResource(las)}
}

// ImportState imports the given attribute values into [LogAnalyticsSolution]'s state.
func (las *LogAnalyticsSolution) ImportState(av io.Reader) error {
	las.state = &logAnalyticsSolutionState{}
	if err := json.NewDecoder(av).Decode(las.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", las.Type(), las.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsSolution] has state.
func (las *LogAnalyticsSolution) State() (*logAnalyticsSolutionState, bool) {
	return las.state, las.state != nil
}

// StateMust returns the state for [LogAnalyticsSolution]. Panics if the state is nil.
func (las *LogAnalyticsSolution) StateMust() *logAnalyticsSolutionState {
	if las.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", las.Type(), las.LocalName()))
	}
	return las.state
}

// LogAnalyticsSolutionArgs contains the configurations for azurerm_log_analytics_solution.
type LogAnalyticsSolutionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SolutionName: string, required
	SolutionName terra.StringValue `hcl:"solution_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceName: string, required
	WorkspaceName terra.StringValue `hcl:"workspace_name,attr" validate:"required"`
	// WorkspaceResourceId: string, required
	WorkspaceResourceId terra.StringValue `hcl:"workspace_resource_id,attr" validate:"required"`
	// Plan: required
	Plan *loganalyticssolution.Plan `hcl:"plan,block" validate:"required"`
	// Timeouts: optional
	Timeouts *loganalyticssolution.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsSolutionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("resource_group_name"))
}

// SolutionName returns a reference to field solution_name of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) SolutionName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("solution_name"))
}

// Tags returns a reference to field tags of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](las.ref.Append("tags"))
}

// WorkspaceName returns a reference to field workspace_name of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) WorkspaceName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("workspace_name"))
}

// WorkspaceResourceId returns a reference to field workspace_resource_id of azurerm_log_analytics_solution.
func (las logAnalyticsSolutionAttributes) WorkspaceResourceId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("workspace_resource_id"))
}

func (las logAnalyticsSolutionAttributes) Plan() terra.ListValue[loganalyticssolution.PlanAttributes] {
	return terra.ReferenceAsList[loganalyticssolution.PlanAttributes](las.ref.Append("plan"))
}

func (las logAnalyticsSolutionAttributes) Timeouts() loganalyticssolution.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticssolution.TimeoutsAttributes](las.ref.Append("timeouts"))
}

type logAnalyticsSolutionState struct {
	Id                  string                              `json:"id"`
	Location            string                              `json:"location"`
	ResourceGroupName   string                              `json:"resource_group_name"`
	SolutionName        string                              `json:"solution_name"`
	Tags                map[string]string                   `json:"tags"`
	WorkspaceName       string                              `json:"workspace_name"`
	WorkspaceResourceId string                              `json:"workspace_resource_id"`
	Plan                []loganalyticssolution.PlanState    `json:"plan"`
	Timeouts            *loganalyticssolution.TimeoutsState `json:"timeouts"`
}
