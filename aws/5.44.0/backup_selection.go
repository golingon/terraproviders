// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	backupselection "github.com/golingon/terraproviders/aws/5.44.0/backupselection"
	"io"
)

// NewBackupSelection creates a new instance of [BackupSelection].
func NewBackupSelection(name string, args BackupSelectionArgs) *BackupSelection {
	return &BackupSelection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupSelection)(nil)

// BackupSelection represents the Terraform resource aws_backup_selection.
type BackupSelection struct {
	Name      string
	Args      BackupSelectionArgs
	state     *backupSelectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupSelection].
func (bs *BackupSelection) Type() string {
	return "aws_backup_selection"
}

// LocalName returns the local name for [BackupSelection].
func (bs *BackupSelection) LocalName() string {
	return bs.Name
}

// Configuration returns the configuration (args) for [BackupSelection].
func (bs *BackupSelection) Configuration() interface{} {
	return bs.Args
}

// DependOn is used for other resources to depend on [BackupSelection].
func (bs *BackupSelection) DependOn() terra.Reference {
	return terra.ReferenceResource(bs)
}

// Dependencies returns the list of resources [BackupSelection] depends_on.
func (bs *BackupSelection) Dependencies() terra.Dependencies {
	return bs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupSelection].
func (bs *BackupSelection) LifecycleManagement() *terra.Lifecycle {
	return bs.Lifecycle
}

// Attributes returns the attributes for [BackupSelection].
func (bs *BackupSelection) Attributes() backupSelectionAttributes {
	return backupSelectionAttributes{ref: terra.ReferenceResource(bs)}
}

// ImportState imports the given attribute values into [BackupSelection]'s state.
func (bs *BackupSelection) ImportState(av io.Reader) error {
	bs.state = &backupSelectionState{}
	if err := json.NewDecoder(av).Decode(bs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bs.Type(), bs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupSelection] has state.
func (bs *BackupSelection) State() (*backupSelectionState, bool) {
	return bs.state, bs.state != nil
}

// StateMust returns the state for [BackupSelection]. Panics if the state is nil.
func (bs *BackupSelection) StateMust() *backupSelectionState {
	if bs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bs.Type(), bs.LocalName()))
	}
	return bs.state
}

// BackupSelectionArgs contains the configurations for aws_backup_selection.
type BackupSelectionArgs struct {
	// IamRoleArn: string, required
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotResources: set of string, optional
	NotResources terra.SetValue[terra.StringValue] `hcl:"not_resources,attr"`
	// PlanId: string, required
	PlanId terra.StringValue `hcl:"plan_id,attr" validate:"required"`
	// Resources: set of string, optional
	Resources terra.SetValue[terra.StringValue] `hcl:"resources,attr"`
	// Condition: min=0
	Condition []backupselection.Condition `hcl:"condition,block" validate:"min=0"`
	// SelectionTag: min=0
	SelectionTag []backupselection.SelectionTag `hcl:"selection_tag,block" validate:"min=0"`
}
type backupSelectionAttributes struct {
	ref terra.Reference
}

// IamRoleArn returns a reference to field iam_role_arn of aws_backup_selection.
func (bs backupSelectionAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_backup_selection.
func (bs backupSelectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("id"))
}

// Name returns a reference to field name of aws_backup_selection.
func (bs backupSelectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("name"))
}

// NotResources returns a reference to field not_resources of aws_backup_selection.
func (bs backupSelectionAttributes) NotResources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bs.ref.Append("not_resources"))
}

// PlanId returns a reference to field plan_id of aws_backup_selection.
func (bs backupSelectionAttributes) PlanId() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("plan_id"))
}

// Resources returns a reference to field resources of aws_backup_selection.
func (bs backupSelectionAttributes) Resources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bs.ref.Append("resources"))
}

func (bs backupSelectionAttributes) Condition() terra.SetValue[backupselection.ConditionAttributes] {
	return terra.ReferenceAsSet[backupselection.ConditionAttributes](bs.ref.Append("condition"))
}

func (bs backupSelectionAttributes) SelectionTag() terra.SetValue[backupselection.SelectionTagAttributes] {
	return terra.ReferenceAsSet[backupselection.SelectionTagAttributes](bs.ref.Append("selection_tag"))
}

type backupSelectionState struct {
	IamRoleArn   string                              `json:"iam_role_arn"`
	Id           string                              `json:"id"`
	Name         string                              `json:"name"`
	NotResources []string                            `json:"not_resources"`
	PlanId       string                              `json:"plan_id"`
	Resources    []string                            `json:"resources"`
	Condition    []backupselection.ConditionState    `json:"condition"`
	SelectionTag []backupselection.SelectionTagState `json:"selection_tag"`
}
