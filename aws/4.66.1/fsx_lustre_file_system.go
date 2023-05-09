// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxlustrefilesystem "github.com/golingon/terraproviders/aws/4.66.1/fsxlustrefilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxLustreFileSystem creates a new instance of [FsxLustreFileSystem].
func NewFsxLustreFileSystem(name string, args FsxLustreFileSystemArgs) *FsxLustreFileSystem {
	return &FsxLustreFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxLustreFileSystem)(nil)

// FsxLustreFileSystem represents the Terraform resource aws_fsx_lustre_file_system.
type FsxLustreFileSystem struct {
	Name      string
	Args      FsxLustreFileSystemArgs
	state     *fsxLustreFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxLustreFileSystem].
func (flfs *FsxLustreFileSystem) Type() string {
	return "aws_fsx_lustre_file_system"
}

// LocalName returns the local name for [FsxLustreFileSystem].
func (flfs *FsxLustreFileSystem) LocalName() string {
	return flfs.Name
}

// Configuration returns the configuration (args) for [FsxLustreFileSystem].
func (flfs *FsxLustreFileSystem) Configuration() interface{} {
	return flfs.Args
}

// DependOn is used for other resources to depend on [FsxLustreFileSystem].
func (flfs *FsxLustreFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(flfs)
}

// Dependencies returns the list of resources [FsxLustreFileSystem] depends_on.
func (flfs *FsxLustreFileSystem) Dependencies() terra.Dependencies {
	return flfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxLustreFileSystem].
func (flfs *FsxLustreFileSystem) LifecycleManagement() *terra.Lifecycle {
	return flfs.Lifecycle
}

// Attributes returns the attributes for [FsxLustreFileSystem].
func (flfs *FsxLustreFileSystem) Attributes() fsxLustreFileSystemAttributes {
	return fsxLustreFileSystemAttributes{ref: terra.ReferenceResource(flfs)}
}

// ImportState imports the given attribute values into [FsxLustreFileSystem]'s state.
func (flfs *FsxLustreFileSystem) ImportState(av io.Reader) error {
	flfs.state = &fsxLustreFileSystemState{}
	if err := json.NewDecoder(av).Decode(flfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", flfs.Type(), flfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxLustreFileSystem] has state.
func (flfs *FsxLustreFileSystem) State() (*fsxLustreFileSystemState, bool) {
	return flfs.state, flfs.state != nil
}

// StateMust returns the state for [FsxLustreFileSystem]. Panics if the state is nil.
func (flfs *FsxLustreFileSystem) StateMust() *fsxLustreFileSystemState {
	if flfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", flfs.Type(), flfs.LocalName()))
	}
	return flfs.state
}

// FsxLustreFileSystemArgs contains the configurations for aws_fsx_lustre_file_system.
type FsxLustreFileSystemArgs struct {
	// AutoImportPolicy: string, optional
	AutoImportPolicy terra.StringValue `hcl:"auto_import_policy,attr"`
	// AutomaticBackupRetentionDays: number, optional
	AutomaticBackupRetentionDays terra.NumberValue `hcl:"automatic_backup_retention_days,attr"`
	// BackupId: string, optional
	BackupId terra.StringValue `hcl:"backup_id,attr"`
	// CopyTagsToBackups: bool, optional
	CopyTagsToBackups terra.BoolValue `hcl:"copy_tags_to_backups,attr"`
	// DailyAutomaticBackupStartTime: string, optional
	DailyAutomaticBackupStartTime terra.StringValue `hcl:"daily_automatic_backup_start_time,attr"`
	// DataCompressionType: string, optional
	DataCompressionType terra.StringValue `hcl:"data_compression_type,attr"`
	// DeploymentType: string, optional
	DeploymentType terra.StringValue `hcl:"deployment_type,attr"`
	// DriveCacheType: string, optional
	DriveCacheType terra.StringValue `hcl:"drive_cache_type,attr"`
	// ExportPath: string, optional
	ExportPath terra.StringValue `hcl:"export_path,attr"`
	// FileSystemTypeVersion: string, optional
	FileSystemTypeVersion terra.StringValue `hcl:"file_system_type_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportPath: string, optional
	ImportPath terra.StringValue `hcl:"import_path,attr"`
	// ImportedFileChunkSize: number, optional
	ImportedFileChunkSize terra.NumberValue `hcl:"imported_file_chunk_size,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PerUnitStorageThroughput: number, optional
	PerUnitStorageThroughput terra.NumberValue `hcl:"per_unit_storage_throughput,attr"`
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
	// WeeklyMaintenanceStartTime: string, optional
	WeeklyMaintenanceStartTime terra.StringValue `hcl:"weekly_maintenance_start_time,attr"`
	// LogConfiguration: optional
	LogConfiguration *fsxlustrefilesystem.LogConfiguration `hcl:"log_configuration,block"`
	// RootSquashConfiguration: optional
	RootSquashConfiguration *fsxlustrefilesystem.RootSquashConfiguration `hcl:"root_squash_configuration,block"`
	// Timeouts: optional
	Timeouts *fsxlustrefilesystem.Timeouts `hcl:"timeouts,block"`
}
type fsxLustreFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("arn"))
}

// AutoImportPolicy returns a reference to field auto_import_policy of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) AutoImportPolicy() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("auto_import_policy"))
}

// AutomaticBackupRetentionDays returns a reference to field automatic_backup_retention_days of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) AutomaticBackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(flfs.ref.Append("automatic_backup_retention_days"))
}

// BackupId returns a reference to field backup_id of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) BackupId() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("backup_id"))
}

// CopyTagsToBackups returns a reference to field copy_tags_to_backups of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) CopyTagsToBackups() terra.BoolValue {
	return terra.ReferenceAsBool(flfs.ref.Append("copy_tags_to_backups"))
}

// DailyAutomaticBackupStartTime returns a reference to field daily_automatic_backup_start_time of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) DailyAutomaticBackupStartTime() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("daily_automatic_backup_start_time"))
}

// DataCompressionType returns a reference to field data_compression_type of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) DataCompressionType() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("data_compression_type"))
}

// DeploymentType returns a reference to field deployment_type of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) DeploymentType() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("deployment_type"))
}

// DnsName returns a reference to field dns_name of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("dns_name"))
}

// DriveCacheType returns a reference to field drive_cache_type of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) DriveCacheType() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("drive_cache_type"))
}

// ExportPath returns a reference to field export_path of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) ExportPath() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("export_path"))
}

// FileSystemTypeVersion returns a reference to field file_system_type_version of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) FileSystemTypeVersion() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("file_system_type_version"))
}

// Id returns a reference to field id of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("id"))
}

// ImportPath returns a reference to field import_path of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) ImportPath() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("import_path"))
}

// ImportedFileChunkSize returns a reference to field imported_file_chunk_size of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) ImportedFileChunkSize() terra.NumberValue {
	return terra.ReferenceAsNumber(flfs.ref.Append("imported_file_chunk_size"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("kms_key_id"))
}

// MountName returns a reference to field mount_name of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) MountName() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("mount_name"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](flfs.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("owner_id"))
}

// PerUnitStorageThroughput returns a reference to field per_unit_storage_throughput of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) PerUnitStorageThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(flfs.ref.Append("per_unit_storage_throughput"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](flfs.ref.Append("security_group_ids"))
}

// StorageCapacity returns a reference to field storage_capacity of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) StorageCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(flfs.ref.Append("storage_capacity"))
}

// StorageType returns a reference to field storage_type of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("storage_type"))
}

// SubnetIds returns a reference to field subnet_ids of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](flfs.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](flfs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](flfs.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("vpc_id"))
}

// WeeklyMaintenanceStartTime returns a reference to field weekly_maintenance_start_time of aws_fsx_lustre_file_system.
func (flfs fsxLustreFileSystemAttributes) WeeklyMaintenanceStartTime() terra.StringValue {
	return terra.ReferenceAsString(flfs.ref.Append("weekly_maintenance_start_time"))
}

func (flfs fsxLustreFileSystemAttributes) LogConfiguration() terra.ListValue[fsxlustrefilesystem.LogConfigurationAttributes] {
	return terra.ReferenceAsList[fsxlustrefilesystem.LogConfigurationAttributes](flfs.ref.Append("log_configuration"))
}

func (flfs fsxLustreFileSystemAttributes) RootSquashConfiguration() terra.ListValue[fsxlustrefilesystem.RootSquashConfigurationAttributes] {
	return terra.ReferenceAsList[fsxlustrefilesystem.RootSquashConfigurationAttributes](flfs.ref.Append("root_squash_configuration"))
}

func (flfs fsxLustreFileSystemAttributes) Timeouts() fsxlustrefilesystem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxlustrefilesystem.TimeoutsAttributes](flfs.ref.Append("timeouts"))
}

type fsxLustreFileSystemState struct {
	Arn                           string                                             `json:"arn"`
	AutoImportPolicy              string                                             `json:"auto_import_policy"`
	AutomaticBackupRetentionDays  float64                                            `json:"automatic_backup_retention_days"`
	BackupId                      string                                             `json:"backup_id"`
	CopyTagsToBackups             bool                                               `json:"copy_tags_to_backups"`
	DailyAutomaticBackupStartTime string                                             `json:"daily_automatic_backup_start_time"`
	DataCompressionType           string                                             `json:"data_compression_type"`
	DeploymentType                string                                             `json:"deployment_type"`
	DnsName                       string                                             `json:"dns_name"`
	DriveCacheType                string                                             `json:"drive_cache_type"`
	ExportPath                    string                                             `json:"export_path"`
	FileSystemTypeVersion         string                                             `json:"file_system_type_version"`
	Id                            string                                             `json:"id"`
	ImportPath                    string                                             `json:"import_path"`
	ImportedFileChunkSize         float64                                            `json:"imported_file_chunk_size"`
	KmsKeyId                      string                                             `json:"kms_key_id"`
	MountName                     string                                             `json:"mount_name"`
	NetworkInterfaceIds           []string                                           `json:"network_interface_ids"`
	OwnerId                       string                                             `json:"owner_id"`
	PerUnitStorageThroughput      float64                                            `json:"per_unit_storage_throughput"`
	SecurityGroupIds              []string                                           `json:"security_group_ids"`
	StorageCapacity               float64                                            `json:"storage_capacity"`
	StorageType                   string                                             `json:"storage_type"`
	SubnetIds                     []string                                           `json:"subnet_ids"`
	Tags                          map[string]string                                  `json:"tags"`
	TagsAll                       map[string]string                                  `json:"tags_all"`
	VpcId                         string                                             `json:"vpc_id"`
	WeeklyMaintenanceStartTime    string                                             `json:"weekly_maintenance_start_time"`
	LogConfiguration              []fsxlustrefilesystem.LogConfigurationState        `json:"log_configuration"`
	RootSquashConfiguration       []fsxlustrefilesystem.RootSquashConfigurationState `json:"root_squash_configuration"`
	Timeouts                      *fsxlustrefilesystem.TimeoutsState                 `json:"timeouts"`
}
