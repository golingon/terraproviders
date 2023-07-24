// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataiamgroup "github.com/golingon/terraproviders/aws/5.9.0/dataiamgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamGroup creates a new instance of [DataIamGroup].
func NewDataIamGroup(name string, args DataIamGroupArgs) *DataIamGroup {
	return &DataIamGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamGroup)(nil)

// DataIamGroup represents the Terraform data resource aws_iam_group.
type DataIamGroup struct {
	Name string
	Args DataIamGroupArgs
}

// DataSource returns the Terraform object type for [DataIamGroup].
func (ig *DataIamGroup) DataSource() string {
	return "aws_iam_group"
}

// LocalName returns the local name for [DataIamGroup].
func (ig *DataIamGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [DataIamGroup].
func (ig *DataIamGroup) Configuration() interface{} {
	return ig.Args
}

// Attributes returns the attributes for [DataIamGroup].
func (ig *DataIamGroup) Attributes() dataIamGroupAttributes {
	return dataIamGroupAttributes{ref: terra.ReferenceDataResource(ig)}
}

// DataIamGroupArgs contains the configurations for aws_iam_group.
type DataIamGroupArgs struct {
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Users: min=0
	Users []dataiamgroup.Users `hcl:"users,block" validate:"min=0"`
}
type dataIamGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_group.
func (ig dataIamGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("arn"))
}

// GroupId returns a reference to field group_id of aws_iam_group.
func (ig dataIamGroupAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("group_id"))
}

// GroupName returns a reference to field group_name of aws_iam_group.
func (ig dataIamGroupAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("group_name"))
}

// Id returns a reference to field id of aws_iam_group.
func (ig dataIamGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// Path returns a reference to field path of aws_iam_group.
func (ig dataIamGroupAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("path"))
}

func (ig dataIamGroupAttributes) Users() terra.ListValue[dataiamgroup.UsersAttributes] {
	return terra.ReferenceAsList[dataiamgroup.UsersAttributes](ig.ref.Append("users"))
}
