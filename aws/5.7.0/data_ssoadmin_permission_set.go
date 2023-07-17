// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSsoadminPermissionSet creates a new instance of [DataSsoadminPermissionSet].
func NewDataSsoadminPermissionSet(name string, args DataSsoadminPermissionSetArgs) *DataSsoadminPermissionSet {
	return &DataSsoadminPermissionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsoadminPermissionSet)(nil)

// DataSsoadminPermissionSet represents the Terraform data resource aws_ssoadmin_permission_set.
type DataSsoadminPermissionSet struct {
	Name string
	Args DataSsoadminPermissionSetArgs
}

// DataSource returns the Terraform object type for [DataSsoadminPermissionSet].
func (sps *DataSsoadminPermissionSet) DataSource() string {
	return "aws_ssoadmin_permission_set"
}

// LocalName returns the local name for [DataSsoadminPermissionSet].
func (sps *DataSsoadminPermissionSet) LocalName() string {
	return sps.Name
}

// Configuration returns the configuration (args) for [DataSsoadminPermissionSet].
func (sps *DataSsoadminPermissionSet) Configuration() interface{} {
	return sps.Args
}

// Attributes returns the attributes for [DataSsoadminPermissionSet].
func (sps *DataSsoadminPermissionSet) Attributes() dataSsoadminPermissionSetAttributes {
	return dataSsoadminPermissionSetAttributes{ref: terra.ReferenceDataResource(sps)}
}

// DataSsoadminPermissionSetArgs contains the configurations for aws_ssoadmin_permission_set.
type DataSsoadminPermissionSetArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceArn: string, required
	InstanceArn terra.StringValue `hcl:"instance_arn,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataSsoadminPermissionSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("description"))
}

// Id returns a reference to field id of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("id"))
}

// InstanceArn returns a reference to field instance_arn of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) InstanceArn() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("instance_arn"))
}

// Name returns a reference to field name of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("name"))
}

// RelayState returns a reference to field relay_state of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) RelayState() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("relay_state"))
}

// SessionDuration returns a reference to field session_duration of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) SessionDuration() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("session_duration"))
}

// Tags returns a reference to field tags of aws_ssoadmin_permission_set.
func (sps dataSsoadminPermissionSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sps.ref.Append("tags"))
}