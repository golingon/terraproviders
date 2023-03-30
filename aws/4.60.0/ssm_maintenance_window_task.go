// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmmaintenancewindowtask "github.com/golingon/terraproviders/aws/4.60.0/ssmmaintenancewindowtask"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSsmMaintenanceWindowTask(name string, args SsmMaintenanceWindowTaskArgs) *SsmMaintenanceWindowTask {
	return &SsmMaintenanceWindowTask{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmMaintenanceWindowTask)(nil)

type SsmMaintenanceWindowTask struct {
	Name  string
	Args  SsmMaintenanceWindowTaskArgs
	state *ssmMaintenanceWindowTaskState
}

func (smwt *SsmMaintenanceWindowTask) Type() string {
	return "aws_ssm_maintenance_window_task"
}

func (smwt *SsmMaintenanceWindowTask) LocalName() string {
	return smwt.Name
}

func (smwt *SsmMaintenanceWindowTask) Configuration() interface{} {
	return smwt.Args
}

func (smwt *SsmMaintenanceWindowTask) Attributes() ssmMaintenanceWindowTaskAttributes {
	return ssmMaintenanceWindowTaskAttributes{ref: terra.ReferenceResource(smwt)}
}

func (smwt *SsmMaintenanceWindowTask) ImportState(av io.Reader) error {
	smwt.state = &ssmMaintenanceWindowTaskState{}
	if err := json.NewDecoder(av).Decode(smwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smwt.Type(), smwt.LocalName(), err)
	}
	return nil
}

func (smwt *SsmMaintenanceWindowTask) State() (*ssmMaintenanceWindowTaskState, bool) {
	return smwt.state, smwt.state != nil
}

func (smwt *SsmMaintenanceWindowTask) StateMust() *ssmMaintenanceWindowTaskState {
	if smwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smwt.Type(), smwt.LocalName()))
	}
	return smwt.state
}

func (smwt *SsmMaintenanceWindowTask) DependOn() terra.Reference {
	return terra.ReferenceResource(smwt)
}

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
	// DependsOn contains resources that SsmMaintenanceWindowTask depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ssmMaintenanceWindowTaskAttributes struct {
	ref terra.Reference
}

func (smwt ssmMaintenanceWindowTaskAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("arn"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) CutoffBehavior() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("cutoff_behavior"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) Description() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("description"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) Id() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("id"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) MaxConcurrency() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("max_concurrency"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) MaxErrors() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("max_errors"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) Name() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("name"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(smwt.ref.Append("priority"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("service_role_arn"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) TaskArn() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("task_arn"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) TaskType() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("task_type"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) WindowId() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("window_id"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) WindowTaskId() terra.StringValue {
	return terra.ReferenceString(smwt.ref.Append("window_task_id"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) Targets() terra.ListValue[ssmmaintenancewindowtask.TargetsAttributes] {
	return terra.ReferenceList[ssmmaintenancewindowtask.TargetsAttributes](smwt.ref.Append("targets"))
}

func (smwt ssmMaintenanceWindowTaskAttributes) TaskInvocationParameters() terra.ListValue[ssmmaintenancewindowtask.TaskInvocationParametersAttributes] {
	return terra.ReferenceList[ssmmaintenancewindowtask.TaskInvocationParametersAttributes](smwt.ref.Append("task_invocation_parameters"))
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
