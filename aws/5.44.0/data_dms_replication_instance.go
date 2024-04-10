// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataDmsReplicationInstance creates a new instance of [DataDmsReplicationInstance].
func NewDataDmsReplicationInstance(name string, args DataDmsReplicationInstanceArgs) *DataDmsReplicationInstance {
	return &DataDmsReplicationInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDmsReplicationInstance)(nil)

// DataDmsReplicationInstance represents the Terraform data resource aws_dms_replication_instance.
type DataDmsReplicationInstance struct {
	Name string
	Args DataDmsReplicationInstanceArgs
}

// DataSource returns the Terraform object type for [DataDmsReplicationInstance].
func (dri *DataDmsReplicationInstance) DataSource() string {
	return "aws_dms_replication_instance"
}

// LocalName returns the local name for [DataDmsReplicationInstance].
func (dri *DataDmsReplicationInstance) LocalName() string {
	return dri.Name
}

// Configuration returns the configuration (args) for [DataDmsReplicationInstance].
func (dri *DataDmsReplicationInstance) Configuration() interface{} {
	return dri.Args
}

// Attributes returns the attributes for [DataDmsReplicationInstance].
func (dri *DataDmsReplicationInstance) Attributes() dataDmsReplicationInstanceAttributes {
	return dataDmsReplicationInstanceAttributes{ref: terra.ReferenceDataResource(dri)}
}

// DataDmsReplicationInstanceArgs contains the configurations for aws_dms_replication_instance.
type DataDmsReplicationInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReplicationInstanceId: string, required
	ReplicationInstanceId terra.StringValue `hcl:"replication_instance_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataDmsReplicationInstanceAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(dri.ref.Append("allocated_storage"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("auto_minor_version_upgrade"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("availability_zone"))
}

// EngineVersion returns a reference to field engine_version of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("kms_key_arn"))
}

// MultiAz returns a reference to field multi_az of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) MultiAz() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("multi_az"))
}

// NetworkType returns a reference to field network_type of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("network_type"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("preferred_maintenance_window"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("publicly_accessible"))
}

// ReplicationInstanceArn returns a reference to field replication_instance_arn of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) ReplicationInstanceArn() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_instance_arn"))
}

// ReplicationInstanceClass returns a reference to field replication_instance_class of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) ReplicationInstanceClass() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_instance_class"))
}

// ReplicationInstanceId returns a reference to field replication_instance_id of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) ReplicationInstanceId() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_instance_id"))
}

// ReplicationInstancePrivateIps returns a reference to field replication_instance_private_ips of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) ReplicationInstancePrivateIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dri.ref.Append("replication_instance_private_ips"))
}

// ReplicationInstancePublicIps returns a reference to field replication_instance_public_ips of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) ReplicationInstancePublicIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dri.ref.Append("replication_instance_public_ips"))
}

// ReplicationSubnetGroupId returns a reference to field replication_subnet_group_id of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) ReplicationSubnetGroupId() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_subnet_group_id"))
}

// Tags returns a reference to field tags of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dri.ref.Append("tags"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_dms_replication_instance.
func (dri dataDmsReplicationInstanceAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dri.ref.Append("vpc_security_group_ids"))
}
