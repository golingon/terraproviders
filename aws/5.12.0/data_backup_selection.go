// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataBackupSelection creates a new instance of [DataBackupSelection].
func NewDataBackupSelection(name string, args DataBackupSelectionArgs) *DataBackupSelection {
	return &DataBackupSelection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBackupSelection)(nil)

// DataBackupSelection represents the Terraform data resource aws_backup_selection.
type DataBackupSelection struct {
	Name string
	Args DataBackupSelectionArgs
}

// DataSource returns the Terraform object type for [DataBackupSelection].
func (bs *DataBackupSelection) DataSource() string {
	return "aws_backup_selection"
}

// LocalName returns the local name for [DataBackupSelection].
func (bs *DataBackupSelection) LocalName() string {
	return bs.Name
}

// Configuration returns the configuration (args) for [DataBackupSelection].
func (bs *DataBackupSelection) Configuration() interface{} {
	return bs.Args
}

// Attributes returns the attributes for [DataBackupSelection].
func (bs *DataBackupSelection) Attributes() dataBackupSelectionAttributes {
	return dataBackupSelectionAttributes{ref: terra.ReferenceDataResource(bs)}
}

// DataBackupSelectionArgs contains the configurations for aws_backup_selection.
type DataBackupSelectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PlanId: string, required
	PlanId terra.StringValue `hcl:"plan_id,attr" validate:"required"`
	// SelectionId: string, required
	SelectionId terra.StringValue `hcl:"selection_id,attr" validate:"required"`
}
type dataBackupSelectionAttributes struct {
	ref terra.Reference
}

// IamRoleArn returns a reference to field iam_role_arn of aws_backup_selection.
func (bs dataBackupSelectionAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_backup_selection.
func (bs dataBackupSelectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("id"))
}

// Name returns a reference to field name of aws_backup_selection.
func (bs dataBackupSelectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("name"))
}

// PlanId returns a reference to field plan_id of aws_backup_selection.
func (bs dataBackupSelectionAttributes) PlanId() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("plan_id"))
}

// Resources returns a reference to field resources of aws_backup_selection.
func (bs dataBackupSelectionAttributes) Resources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bs.ref.Append("resources"))
}

// SelectionId returns a reference to field selection_id of aws_backup_selection.
func (bs dataBackupSelectionAttributes) SelectionId() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("selection_id"))
}
