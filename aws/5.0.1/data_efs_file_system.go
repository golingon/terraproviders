// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataefsfilesystem "github.com/golingon/terraproviders/aws/5.0.1/dataefsfilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEfsFileSystem creates a new instance of [DataEfsFileSystem].
func NewDataEfsFileSystem(name string, args DataEfsFileSystemArgs) *DataEfsFileSystem {
	return &DataEfsFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEfsFileSystem)(nil)

// DataEfsFileSystem represents the Terraform data resource aws_efs_file_system.
type DataEfsFileSystem struct {
	Name string
	Args DataEfsFileSystemArgs
}

// DataSource returns the Terraform object type for [DataEfsFileSystem].
func (efs *DataEfsFileSystem) DataSource() string {
	return "aws_efs_file_system"
}

// LocalName returns the local name for [DataEfsFileSystem].
func (efs *DataEfsFileSystem) LocalName() string {
	return efs.Name
}

// Configuration returns the configuration (args) for [DataEfsFileSystem].
func (efs *DataEfsFileSystem) Configuration() interface{} {
	return efs.Args
}

// Attributes returns the attributes for [DataEfsFileSystem].
func (efs *DataEfsFileSystem) Attributes() dataEfsFileSystemAttributes {
	return dataEfsFileSystemAttributes{ref: terra.ReferenceDataResource(efs)}
}

// DataEfsFileSystemArgs contains the configurations for aws_efs_file_system.
type DataEfsFileSystemArgs struct {
	// CreationToken: string, optional
	CreationToken terra.StringValue `hcl:"creation_token,attr"`
	// FileSystemId: string, optional
	FileSystemId terra.StringValue `hcl:"file_system_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// LifecyclePolicy: min=0
	LifecyclePolicy []dataefsfilesystem.LifecyclePolicy `hcl:"lifecycle_policy,block" validate:"min=0"`
}
type dataEfsFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("arn"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("availability_zone_id"))
}

// AvailabilityZoneName returns a reference to field availability_zone_name of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) AvailabilityZoneName() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("availability_zone_name"))
}

// CreationToken returns a reference to field creation_token of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) CreationToken() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("creation_token"))
}

// DnsName returns a reference to field dns_name of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("dns_name"))
}

// Encrypted returns a reference to field encrypted of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(efs.ref.Append("encrypted"))
}

// FileSystemId returns a reference to field file_system_id of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("kms_key_id"))
}

// PerformanceMode returns a reference to field performance_mode of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) PerformanceMode() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("performance_mode"))
}

// ProvisionedThroughputInMibps returns a reference to field provisioned_throughput_in_mibps of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) ProvisionedThroughputInMibps() terra.NumberValue {
	return terra.ReferenceAsNumber(efs.ref.Append("provisioned_throughput_in_mibps"))
}

// SizeInBytes returns a reference to field size_in_bytes of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) SizeInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(efs.ref.Append("size_in_bytes"))
}

// Tags returns a reference to field tags of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](efs.ref.Append("tags"))
}

// ThroughputMode returns a reference to field throughput_mode of aws_efs_file_system.
func (efs dataEfsFileSystemAttributes) ThroughputMode() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("throughput_mode"))
}

func (efs dataEfsFileSystemAttributes) LifecyclePolicy() terra.ListValue[dataefsfilesystem.LifecyclePolicyAttributes] {
	return terra.ReferenceAsList[dataefsfilesystem.LifecyclePolicyAttributes](efs.ref.Append("lifecycle_policy"))
}
