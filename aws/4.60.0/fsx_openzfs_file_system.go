// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxopenzfsfilesystem "github.com/golingon/terraproviders/aws/4.60.0/fsxopenzfsfilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxOpenzfsFileSystem creates a new instance of [FsxOpenzfsFileSystem].
func NewFsxOpenzfsFileSystem(name string, args FsxOpenzfsFileSystemArgs) *FsxOpenzfsFileSystem {
	return &FsxOpenzfsFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxOpenzfsFileSystem)(nil)

// FsxOpenzfsFileSystem represents the Terraform resource aws_fsx_openzfs_file_system.
type FsxOpenzfsFileSystem struct {
	Name      string
	Args      FsxOpenzfsFileSystemArgs
	state     *fsxOpenzfsFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxOpenzfsFileSystem].
func (fofs *FsxOpenzfsFileSystem) Type() string {
	return "aws_fsx_openzfs_file_system"
}

// LocalName returns the local name for [FsxOpenzfsFileSystem].
func (fofs *FsxOpenzfsFileSystem) LocalName() string {
	return fofs.Name
}

// Configuration returns the configuration (args) for [FsxOpenzfsFileSystem].
func (fofs *FsxOpenzfsFileSystem) Configuration() interface{} {
	return fofs.Args
}

// DependOn is used for other resources to depend on [FsxOpenzfsFileSystem].
func (fofs *FsxOpenzfsFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(fofs)
}

// Dependencies returns the list of resources [FsxOpenzfsFileSystem] depends_on.
func (fofs *FsxOpenzfsFileSystem) Dependencies() terra.Dependencies {
	return fofs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxOpenzfsFileSystem].
func (fofs *FsxOpenzfsFileSystem) LifecycleManagement() *terra.Lifecycle {
	return fofs.Lifecycle
}

// Attributes returns the attributes for [FsxOpenzfsFileSystem].
func (fofs *FsxOpenzfsFileSystem) Attributes() fsxOpenzfsFileSystemAttributes {
	return fsxOpenzfsFileSystemAttributes{ref: terra.ReferenceResource(fofs)}
}

// ImportState imports the given attribute values into [FsxOpenzfsFileSystem]'s state.
func (fofs *FsxOpenzfsFileSystem) ImportState(av io.Reader) error {
	fofs.state = &fsxOpenzfsFileSystemState{}
	if err := json.NewDecoder(av).Decode(fofs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fofs.Type(), fofs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxOpenzfsFileSystem] has state.
func (fofs *FsxOpenzfsFileSystem) State() (*fsxOpenzfsFileSystemState, bool) {
	return fofs.state, fofs.state != nil
}

// StateMust returns the state for [FsxOpenzfsFileSystem]. Panics if the state is nil.
func (fofs *FsxOpenzfsFileSystem) StateMust() *fsxOpenzfsFileSystemState {
	if fofs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fofs.Type(), fofs.LocalName()))
	}
	return fofs.state
}

// FsxOpenzfsFileSystemArgs contains the configurations for aws_fsx_openzfs_file_system.
type FsxOpenzfsFileSystemArgs struct {
	// AutomaticBackupRetentionDays: number, optional
	AutomaticBackupRetentionDays terra.NumberValue `hcl:"automatic_backup_retention_days,attr"`
	// BackupId: string, optional
	BackupId terra.StringValue `hcl:"backup_id,attr"`
	// CopyTagsToBackups: bool, optional
	CopyTagsToBackups terra.BoolValue `hcl:"copy_tags_to_backups,attr"`
	// CopyTagsToVolumes: bool, optional
	CopyTagsToVolumes terra.BoolValue `hcl:"copy_tags_to_volumes,attr"`
	// DailyAutomaticBackupStartTime: string, optional
	DailyAutomaticBackupStartTime terra.StringValue `hcl:"daily_automatic_backup_start_time,attr"`
	// DeploymentType: string, required
	DeploymentType terra.StringValue `hcl:"deployment_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// StorageCapacity: number, optional
	StorageCapacity terra.NumberValue `hcl:"storage_capacity,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// SubnetIds: list of string, required
	SubnetIds terra.ListValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ThroughputCapacity: number, required
	ThroughputCapacity terra.NumberValue `hcl:"throughput_capacity,attr" validate:"required"`
	// WeeklyMaintenanceStartTime: string, optional
	WeeklyMaintenanceStartTime terra.StringValue `hcl:"weekly_maintenance_start_time,attr"`
	// DiskIopsConfiguration: optional
	DiskIopsConfiguration *fsxopenzfsfilesystem.DiskIopsConfiguration `hcl:"disk_iops_configuration,block"`
	// RootVolumeConfiguration: optional
	RootVolumeConfiguration *fsxopenzfsfilesystem.RootVolumeConfiguration `hcl:"root_volume_configuration,block"`
	// Timeouts: optional
	Timeouts *fsxopenzfsfilesystem.Timeouts `hcl:"timeouts,block"`
}
type fsxOpenzfsFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("arn"))
}

// AutomaticBackupRetentionDays returns a reference to field automatic_backup_retention_days of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) AutomaticBackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(fofs.ref.Append("automatic_backup_retention_days"))
}

// BackupId returns a reference to field backup_id of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) BackupId() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("backup_id"))
}

// CopyTagsToBackups returns a reference to field copy_tags_to_backups of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) CopyTagsToBackups() terra.BoolValue {
	return terra.ReferenceAsBool(fofs.ref.Append("copy_tags_to_backups"))
}

// CopyTagsToVolumes returns a reference to field copy_tags_to_volumes of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) CopyTagsToVolumes() terra.BoolValue {
	return terra.ReferenceAsBool(fofs.ref.Append("copy_tags_to_volumes"))
}

// DailyAutomaticBackupStartTime returns a reference to field daily_automatic_backup_start_time of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) DailyAutomaticBackupStartTime() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("daily_automatic_backup_start_time"))
}

// DeploymentType returns a reference to field deployment_type of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) DeploymentType() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("deployment_type"))
}

// DnsName returns a reference to field dns_name of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("dns_name"))
}

// Id returns a reference to field id of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("kms_key_id"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fofs.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("owner_id"))
}

// RootVolumeId returns a reference to field root_volume_id of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) RootVolumeId() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("root_volume_id"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fofs.ref.Append("security_group_ids"))
}

// StorageCapacity returns a reference to field storage_capacity of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) StorageCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(fofs.ref.Append("storage_capacity"))
}

// StorageType returns a reference to field storage_type of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("storage_type"))
}

// SubnetIds returns a reference to field subnet_ids of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fofs.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fofs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fofs.ref.Append("tags_all"))
}

// ThroughputCapacity returns a reference to field throughput_capacity of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) ThroughputCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(fofs.ref.Append("throughput_capacity"))
}

// VpcId returns a reference to field vpc_id of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("vpc_id"))
}

// WeeklyMaintenanceStartTime returns a reference to field weekly_maintenance_start_time of aws_fsx_openzfs_file_system.
func (fofs fsxOpenzfsFileSystemAttributes) WeeklyMaintenanceStartTime() terra.StringValue {
	return terra.ReferenceAsString(fofs.ref.Append("weekly_maintenance_start_time"))
}

func (fofs fsxOpenzfsFileSystemAttributes) DiskIopsConfiguration() terra.ListValue[fsxopenzfsfilesystem.DiskIopsConfigurationAttributes] {
	return terra.ReferenceAsList[fsxopenzfsfilesystem.DiskIopsConfigurationAttributes](fofs.ref.Append("disk_iops_configuration"))
}

func (fofs fsxOpenzfsFileSystemAttributes) RootVolumeConfiguration() terra.ListValue[fsxopenzfsfilesystem.RootVolumeConfigurationAttributes] {
	return terra.ReferenceAsList[fsxopenzfsfilesystem.RootVolumeConfigurationAttributes](fofs.ref.Append("root_volume_configuration"))
}

func (fofs fsxOpenzfsFileSystemAttributes) Timeouts() fsxopenzfsfilesystem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxopenzfsfilesystem.TimeoutsAttributes](fofs.ref.Append("timeouts"))
}

type fsxOpenzfsFileSystemState struct {
	Arn                           string                                              `json:"arn"`
	AutomaticBackupRetentionDays  float64                                             `json:"automatic_backup_retention_days"`
	BackupId                      string                                              `json:"backup_id"`
	CopyTagsToBackups             bool                                                `json:"copy_tags_to_backups"`
	CopyTagsToVolumes             bool                                                `json:"copy_tags_to_volumes"`
	DailyAutomaticBackupStartTime string                                              `json:"daily_automatic_backup_start_time"`
	DeploymentType                string                                              `json:"deployment_type"`
	DnsName                       string                                              `json:"dns_name"`
	Id                            string                                              `json:"id"`
	KmsKeyId                      string                                              `json:"kms_key_id"`
	NetworkInterfaceIds           []string                                            `json:"network_interface_ids"`
	OwnerId                       string                                              `json:"owner_id"`
	RootVolumeId                  string                                              `json:"root_volume_id"`
	SecurityGroupIds              []string                                            `json:"security_group_ids"`
	StorageCapacity               float64                                             `json:"storage_capacity"`
	StorageType                   string                                              `json:"storage_type"`
	SubnetIds                     []string                                            `json:"subnet_ids"`
	Tags                          map[string]string                                   `json:"tags"`
	TagsAll                       map[string]string                                   `json:"tags_all"`
	ThroughputCapacity            float64                                             `json:"throughput_capacity"`
	VpcId                         string                                              `json:"vpc_id"`
	WeeklyMaintenanceStartTime    string                                              `json:"weekly_maintenance_start_time"`
	DiskIopsConfiguration         []fsxopenzfsfilesystem.DiskIopsConfigurationState   `json:"disk_iops_configuration"`
	RootVolumeConfiguration       []fsxopenzfsfilesystem.RootVolumeConfigurationState `json:"root_volume_configuration"`
	Timeouts                      *fsxopenzfsfilesystem.TimeoutsState                 `json:"timeouts"`
}
