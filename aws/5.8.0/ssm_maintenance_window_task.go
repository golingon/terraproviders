// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmmaintenancewindowtask "github.com/golingon/terraproviders/aws/5.8.0/ssmmaintenancewindowtask"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmMaintenanceWindowTask creates a new instance of [SsmMaintenanceWindowTask].
func NewSsmMaintenanceWindowTask(name string, args SsmMaintenanceWindowTaskArgs) *SsmMaintenanceWindowTask {
	return &SsmMaintenanceWindowTask{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmMaintenanceWindowTask)(nil)

// SsmMaintenanceWindowTask represents the Terraform resource aws_ssm_maintenance_window_task.
type SsmMaintenanceWindowTask struct {
	Name      string
	Args      SsmMaintenanceWindowTaskArgs
	state     *ssmMaintenanceWindowTaskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmMaintenanceWindowTask].
func (smwt *SsmMaintenanceWindowTask) Type() string {
	return "aws_ssm_maintenance_window_task"
}

// LocalName returns the local name for [SsmMaintenanceWindowTask].
func (smwt *SsmMaintenanceWindowTask) LocalName() string {
	return smwt.Name
}

// Configuration returns the configuration (args) for [SsmMaintenanceWindowTask].
func (smwt *SsmMaintenanceWindowTask) Configuration() interface{} {
	return smwt.Args
}

// DependOn is used for other resources to depend on [SsmMaintenanceWindowTask].
func (smwt *SsmMaintenanceWindowTask) DependOn() terra.Reference {
	return terra.ReferenceResource(smwt)
}

// Dependencies returns the list of resources [SsmMaintenanceWindowTask] depends_on.
func (smwt *SsmMaintenanceWindowTask) Dependencies() terra.Dependencies {
	return smwt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmMaintenanceWindowTask].
func (smwt *SsmMaintenanceWindowTask) LifecycleManagement() *terra.Lifecycle {
	return smwt.Lifecycle
}

// Attributes returns the attributes for [SsmMaintenanceWindowTask].
func (smwt *SsmMaintenanceWindowTask) Attributes() ssmMaintenanceWindowTaskAttributes {
	return ssmMaintenanceWindowTaskAttributes{ref: terra.ReferenceResource(smwt)}
}

// ImportState imports the given attribute values into [SsmMaintenanceWindowTask]'s state.
func (smwt *SsmMaintenanceWindowTask) ImportState(av io.Reader) error {
	smwt.state = &ssmMaintenanceWindowTaskState{}
	if err := json.NewDecoder(av).Decode(smwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smwt.Type(), smwt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmMaintenanceWindowTask] has state.
func (smwt *SsmMaintenanceWindowTask) State() (*ssmMaintenanceWindowTaskState, bool) {
	return smwt.state, smwt.state != nil
}

// StateMust returns the state for [SsmMaintenanceWindowTask]. Panics if the state is nil.
func (smwt *SsmMaintenanceWindowTask) StateMust() *ssmMaintenanceWindowTaskState {
	if smwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smwt.Type(), smwt.LocalName()))
	}
	return smwt.state
}

// SsmMaintenanceWindowTaskArgs contains the configurations for aws_ssm_maintenance_window_task.
type SsmMaintenanceWindowTaskArgs struct {
	// CutoffBehavior: string, optional
	CutoffBehavior terra.StringValue `hcl:"cutoff_behavior,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxConcurrency: string, optional
	MaxConcurrency terra.StringValue `hcl:"max_concurrency,attr"`
	// MaxErrors: string, optional
	MaxErrors terra.StringValue `hcl:"max_errors,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ServiceRoleArn: string, optional
	ServiceRoleArn terra.StringValue `hcl:"service_role_arn,attr"`
	// TaskArn: string, required
	TaskArn terra.StringValue `hcl:"task_arn,attr" validate:"required"`
	// TaskType: string, required
	TaskType terra.StringValue `hcl:"task_type,attr" validate:"required"`
	// WindowId: string, required
	WindowId terra.StringValue `hcl:"window_id,attr" validate:"required"`
	// Targets: min=0,max=5
	Targets []ssmmaintenancewindowtask.Targets `hcl:"targets,block" validate:"min=0,max=5"`
	// TaskInvocationParameters: optional
	TaskInvocationParameters *ssmmaintenancewindowtask.TaskInvocationParameters `hcl:"task_invocation_parameters,block"`
}
type ssmMaintenanceWindowTaskAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("arn"))
}

// CutoffBehavior returns a reference to field cutoff_behavior of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) CutoffBehavior() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("cutoff_behavior"))
}

// Description returns a reference to field description of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("description"))
}

// Id returns a reference to field id of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("id"))
}

// MaxConcurrency returns a reference to field max_concurrency of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) MaxConcurrency() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("max_concurrency"))
}

// MaxErrors returns a reference to field max_errors of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) MaxErrors() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("max_errors"))
}

// Name returns a reference to field name of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("name"))
}

// Priority returns a reference to field priority of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(smwt.ref.Append("priority"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("service_role_arn"))
}

// TaskArn returns a reference to field task_arn of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) TaskArn() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("task_arn"))
}

// TaskType returns a reference to field task_type of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) TaskType() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("task_type"))
}

// WindowId returns a reference to field window_id of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) WindowId() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("window_id"))
}

// WindowTaskId returns a reference to field window_task_id of aws_ssm_maintenance_window_task.
func (smwt ssmMaintenanceWindowTaskAttributes) WindowTaskId() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("window_task_id"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) Targets() terra.ListValue[ssmmaintenancewindowtask.TargetsAttributes] {
	return terra.ReferenceAsList[ssmmaintenancewindowtask.TargetsAttributes](smwt.ref.Append("targets"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) TaskInvocationParameters() terra.ListValue[ssmmaintenancewindowtask.TaskInvocationParametersAttributes] {
	return terra.ReferenceAsList[ssmmaintenancewindowtask.TaskInvocationParametersAttributes](smwt.ref.Append("task_invocation_parameters"))
}

type ssmMaintenanceWindowTaskState struct {
	Arn                      string                                                   `json:"arn"`
	CutoffBehavior           string                                                   `json:"cutoff_behavior"`
	Description              string                                                   `json:"description"`
	Id                       string                                                   `json:"id"`
	MaxConcurrency           string                                                   `json:"max_concurrency"`
	MaxErrors                string                                                   `json:"max_errors"`
	Name                     string                                                   `json:"name"`
	Priority                 float64                                                  `json:"priority"`
	ServiceRoleArn           string                                                   `json:"service_role_arn"`
	TaskArn                  string                                                   `json:"task_arn"`
	TaskType                 string                                                   `json:"task_type"`
	WindowId                 string                                                   `json:"window_id"`
	WindowTaskId             string                                                   `json:"window_task_id"`
	Targets                  []ssmmaintenancewindowtask.TargetsState                  `json:"targets"`
	TaskInvocationParameters []ssmmaintenancewindowtask.TaskInvocationParametersState `json:"task_invocation_parameters"`
}
