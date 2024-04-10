// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataEfsMountTarget creates a new instance of [DataEfsMountTarget].
func NewDataEfsMountTarget(name string, args DataEfsMountTargetArgs) *DataEfsMountTarget {
	return &DataEfsMountTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEfsMountTarget)(nil)

// DataEfsMountTarget represents the Terraform data resource aws_efs_mount_target.
type DataEfsMountTarget struct {
	Name string
	Args DataEfsMountTargetArgs
}

// DataSource returns the Terraform object type for [DataEfsMountTarget].
func (emt *DataEfsMountTarget) DataSource() string {
	return "aws_efs_mount_target"
}

// LocalName returns the local name for [DataEfsMountTarget].
func (emt *DataEfsMountTarget) LocalName() string {
	return emt.Name
}

// Configuration returns the configuration (args) for [DataEfsMountTarget].
func (emt *DataEfsMountTarget) Configuration() interface{} {
	return emt.Args
}

// Attributes returns the attributes for [DataEfsMountTarget].
func (emt *DataEfsMountTarget) Attributes() dataEfsMountTargetAttributes {
	return dataEfsMountTargetAttributes{ref: terra.ReferenceDataResource(emt)}
}

// DataEfsMountTargetArgs contains the configurations for aws_efs_mount_target.
type DataEfsMountTargetArgs struct {
	// AccessPointId: string, optional
	AccessPointId terra.StringValue `hcl:"access_point_id,attr"`
	// FileSystemId: string, optional
	FileSystemId terra.StringValue `hcl:"file_system_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MountTargetId: string, optional
	MountTargetId terra.StringValue `hcl:"mount_target_id,attr"`
}
type dataEfsMountTargetAttributes struct {
	ref terra.Reference
}

// AccessPointId returns a reference to field access_point_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) AccessPointId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("access_point_id"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("availability_zone_id"))
}

// AvailabilityZoneName returns a reference to field availability_zone_name of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) AvailabilityZoneName() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("availability_zone_name"))
}

// DnsName returns a reference to field dns_name of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("dns_name"))
}

// FileSystemArn returns a reference to field file_system_arn of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) FileSystemArn() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("file_system_arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("ip_address"))
}

// MountTargetDnsName returns a reference to field mount_target_dns_name of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) MountTargetDnsName() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("mount_target_dns_name"))
}

// MountTargetId returns a reference to field mount_target_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) MountTargetId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("mount_target_id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("network_interface_id"))
}

// OwnerId returns a reference to field owner_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("owner_id"))
}

// SecurityGroups returns a reference to field security_groups of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](emt.ref.Append("security_groups"))
}

// SubnetId returns a reference to field subnet_id of aws_efs_mount_target.
func (emt dataEfsMountTargetAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("subnet_id"))
}
