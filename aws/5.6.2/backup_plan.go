// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	backupplan "github.com/golingon/terraproviders/aws/5.6.2/backupplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupPlan creates a new instance of [BackupPlan].
func NewBackupPlan(name string, args BackupPlanArgs) *BackupPlan {
	return &BackupPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupPlan)(nil)

// BackupPlan represents the Terraform resource aws_backup_plan.
type BackupPlan struct {
	Name      string
	Args      BackupPlanArgs
	state     *backupPlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupPlan].
func (bp *BackupPlan) Type() string {
	return "aws_backup_plan"
}

// LocalName returns the local name for [BackupPlan].
func (bp *BackupPlan) LocalName() string {
	return bp.Name
}

// Configuration returns the configuration (args) for [BackupPlan].
func (bp *BackupPlan) Configuration() interface{} {
	return bp.Args
}

// DependOn is used for other resources to depend on [BackupPlan].
func (bp *BackupPlan) DependOn() terra.Reference {
	return terra.ReferenceResource(bp)
}

// Dependencies returns the list of resources [BackupPlan] depends_on.
func (bp *BackupPlan) Dependencies() terra.Dependencies {
	return bp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupPlan].
func (bp *BackupPlan) LifecycleManagement() *terra.Lifecycle {
	return bp.Lifecycle
}

// Attributes returns the attributes for [BackupPlan].
func (bp *BackupPlan) Attributes() backupPlanAttributes {
	return backupPlanAttributes{ref: terra.ReferenceResource(bp)}
}

// ImportState imports the given attribute values into [BackupPlan]'s state.
func (bp *BackupPlan) ImportState(av io.Reader) error {
	bp.state = &backupPlanState{}
	if err := json.NewDecoder(av).Decode(bp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bp.Type(), bp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupPlan] has state.
func (bp *BackupPlan) State() (*backupPlanState, bool) {
	return bp.state, bp.state != nil
}

// StateMust returns the state for [BackupPlan]. Panics if the state is nil.
func (bp *BackupPlan) StateMust() *backupPlanState {
	if bp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bp.Type(), bp.LocalName()))
	}
	return bp.state
}

// BackupPlanArgs contains the configurations for aws_backup_plan.
type BackupPlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AdvancedBackupSetting: min=0
	AdvancedBackupSetting []backupplan.AdvancedBackupSetting `hcl:"advanced_backup_setting,block" validate:"min=0"`
	// Rule: min=1
	Rule []backupplan.Rule `hcl:"rule,block" validate:"min=1"`
}
type backupPlanAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_backup_plan.
func (bp backupPlanAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_backup_plan.
func (bp backupPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("id"))
}

// Name returns a reference to field name of aws_backup_plan.
func (bp backupPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_backup_plan.
func (bp backupPlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_backup_plan.
func (bp backupPlanAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bp.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_backup_plan.
func (bp backupPlanAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("version"))
}

func (bp backupPlanAttributes) AdvancedBackupSetting() terra.SetValue[backupplan.AdvancedBackupSettingAttributes] {
	return terra.ReferenceAsSet[backupplan.AdvancedBackupSettingAttributes](bp.ref.Append("advanced_backup_setting"))
}

func (bp backupPlanAttributes) Rule() terra.SetValue[backupplan.RuleAttributes] {
	return terra.ReferenceAsSet[backupplan.RuleAttributes](bp.ref.Append("rule"))
}

type backupPlanState struct {
	Arn                   string                                  `json:"arn"`
	Id                    string                                  `json:"id"`
	Name                  string                                  `json:"name"`
	Tags                  map[string]string                       `json:"tags"`
	TagsAll               map[string]string                       `json:"tags_all"`
	Version               string                                  `json:"version"`
	AdvancedBackupSetting []backupplan.AdvancedBackupSettingState `json:"advanced_backup_setting"`
	Rule                  []backupplan.RuleState                  `json:"rule"`
}
