// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDmsReplicationSubnetGroup creates a new instance of [DataDmsReplicationSubnetGroup].
func NewDataDmsReplicationSubnetGroup(name string, args DataDmsReplicationSubnetGroupArgs) *DataDmsReplicationSubnetGroup {
	return &DataDmsReplicationSubnetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDmsReplicationSubnetGroup)(nil)

// DataDmsReplicationSubnetGroup represents the Terraform data resource aws_dms_replication_subnet_group.
type DataDmsReplicationSubnetGroup struct {
	Name string
	Args DataDmsReplicationSubnetGroupArgs
}

// DataSource returns the Terraform object type for [DataDmsReplicationSubnetGroup].
func (drsg *DataDmsReplicationSubnetGroup) DataSource() string {
	return "aws_dms_replication_subnet_group"
}

// LocalName returns the local name for [DataDmsReplicationSubnetGroup].
func (drsg *DataDmsReplicationSubnetGroup) LocalName() string {
	return drsg.Name
}

// Configuration returns the configuration (args) for [DataDmsReplicationSubnetGroup].
func (drsg *DataDmsReplicationSubnetGroup) Configuration() interface{} {
	return drsg.Args
}

// Attributes returns the attributes for [DataDmsReplicationSubnetGroup].
func (drsg *DataDmsReplicationSubnetGroup) Attributes() dataDmsReplicationSubnetGroupAttributes {
	return dataDmsReplicationSubnetGroupAttributes{ref: terra.ReferenceDataResource(drsg)}
}

// DataDmsReplicationSubnetGroupArgs contains the configurations for aws_dms_replication_subnet_group.
type DataDmsReplicationSubnetGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReplicationSubnetGroupId: string, required
	ReplicationSubnetGroupId terra.StringValue `hcl:"replication_subnet_group_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataDmsReplicationSubnetGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("id"))
}

// ReplicationSubnetGroupArn returns a reference to field replication_subnet_group_arn of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) ReplicationSubnetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("replication_subnet_group_arn"))
}

// ReplicationSubnetGroupDescription returns a reference to field replication_subnet_group_description of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) ReplicationSubnetGroupDescription() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("replication_subnet_group_description"))
}

// ReplicationSubnetGroupId returns a reference to field replication_subnet_group_id of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) ReplicationSubnetGroupId() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("replication_subnet_group_id"))
}

// SubnetGroupStatus returns a reference to field subnet_group_status of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) SubnetGroupStatus() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("subnet_group_status"))
}

// SubnetIds returns a reference to field subnet_ids of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](drsg.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](drsg.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_dms_replication_subnet_group.
func (drsg dataDmsReplicationSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("vpc_id"))
}
