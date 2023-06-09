// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxwindowsfilesystem "github.com/golingon/terraproviders/aws/4.60.0/fsxwindowsfilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxWindowsFileSystem creates a new instance of [FsxWindowsFileSystem].
func NewFsxWindowsFileSystem(name string, args FsxWindowsFileSystemArgs) *FsxWindowsFileSystem {
	return &FsxWindowsFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxWindowsFileSystem)(nil)

// FsxWindowsFileSystem represents the Terraform resource aws_fsx_windows_file_system.
type FsxWindowsFileSystem struct {
	Name      string
	Args      FsxWindowsFileSystemArgs
	state     *fsxWindowsFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxWindowsFileSystem].
func (fwfs *FsxWindowsFileSystem) Type() string {
	return "aws_fsx_windows_file_system"
}

// LocalName returns the local name for [FsxWindowsFileSystem].
func (fwfs *FsxWindowsFileSystem) LocalName() string {
	return fwfs.Name
}

// Configuration returns the configuration (args) for [FsxWindowsFileSystem].
func (fwfs *FsxWindowsFileSystem) Configuration() interface{} {
	return fwfs.Args
}

// DependOn is used for other resources to depend on [FsxWindowsFileSystem].
func (fwfs *FsxWindowsFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(fwfs)
}

// Dependencies returns the list of resources [FsxWindowsFileSystem] depends_on.
func (fwfs *FsxWindowsFileSystem) Dependencies() terra.Dependencies {
	return fwfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxWindowsFileSystem].
func (fwfs *FsxWindowsFileSystem) LifecycleManagement() *terra.Lifecycle {
	return fwfs.Lifecycle
}

// Attributes returns the attributes for [FsxWindowsFileSystem].
func (fwfs *FsxWindowsFileSystem) Attributes() fsxWindowsFileSystemAttributes {
	return fsxWindowsFileSystemAttributes{ref: terra.ReferenceResource(fwfs)}
}

// ImportState imports the given attribute values into [FsxWindowsFileSystem]'s state.
func (fwfs *FsxWindowsFileSystem) ImportState(av io.Reader) error {
	fwfs.state = &fsxWindowsFileSystemState{}
	if err := json.NewDecoder(av).Decode(fwfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fwfs.Type(), fwfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxWindowsFileSystem] has state.
func (fwfs *FsxWindowsFileSystem) State() (*fsxWindowsFileSystemState, bool) {
	return fwfs.state, fwfs.state != nil
}

// StateMust returns the state for [FsxWindowsFileSystem]. Panics if the state is nil.
func (fwfs *FsxWindowsFileSystem) StateMust() *fsxWindowsFileSystemState {
	if fwfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fwfs.Type(), fwfs.LocalName()))
	}
	return fwfs.state
}

// FsxWindowsFileSystemArgs contains the configurations for aws_fsx_windows_file_system.
type FsxWindowsFileSystemArgs struct {
	// ActiveDirectoryId: string, optional
	ActiveDirectoryId terra.StringValue `hcl:"active_directory_id,attr"`
	// Aliases: set of string, optional
	Aliases terra.SetValue[terra.StringValue] `hcl:"aliases,attr"`
	// AutomaticBackupRetentionDays: number, optional
	AutomaticBackupRetentionDays terra.NumberValue `hcl:"automatic_backup_retention_days,attr"`
	// BackupId: string, optional
	BackupId terra.StringValue `hcl:"backup_id,attr"`
	// CopyTagsToBackups: bool, optional
	CopyTagsToBackups terra.BoolValue `hcl:"copy_tags_to_backups,attr"`
	// DailyAutomaticBackupStartTime: string, optional
	DailyAutomaticBackupStartTime terra.StringValue `hcl:"daily_automatic_backup_start_time,attr"`
	// DeploymentType: string, optional
	DeploymentType terra.StringValue `hcl:"deployment_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PreferredSubnetId: string, optional
	PreferredSubnetId terra.StringValue `hcl:"preferred_subnet_id,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SkipFinalBackup: bool, optional
	SkipFinalBackup terra.BoolValue `hcl:"skip_final_backup,attr"`
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
	// AuditLogConfiguration: optional
	AuditLogConfiguration *fsxwindowsfilesystem.AuditLogConfiguration `hcl:"audit_log_configuration,block"`
	// SelfManagedActiveDirectory: optional
	SelfManagedActiveDirectory *fsxwindowsfilesystem.SelfManagedActiveDirectory `hcl:"self_managed_active_directory,block"`
	// Timeouts: optional
	Timeouts *fsxwindowsfilesystem.Timeouts `hcl:"timeouts,block"`
}
type fsxWindowsFileSystemAttributes struct {
	ref terra.Reference
}

// ActiveDirectoryId returns a reference to field active_directory_id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) ActiveDirectoryId() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("active_directory_id"))
}

// Aliases returns a reference to field aliases of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) Aliases() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fwfs.ref.Append("aliases"))
}

// Arn returns a reference to field arn of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("arn"))
}

// AutomaticBackupRetentionDays returns a reference to field automatic_backup_retention_days of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) AutomaticBackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(fwfs.ref.Append("automatic_backup_retention_days"))
}

// BackupId returns a reference to field backup_id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) BackupId() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("backup_id"))
}

// CopyTagsToBackups returns a reference to field copy_tags_to_backups of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) CopyTagsToBackups() terra.BoolValue {
	return terra.ReferenceAsBool(fwfs.ref.Append("copy_tags_to_backups"))
}

// DailyAutomaticBackupStartTime returns a reference to field daily_automatic_backup_start_time of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) DailyAutomaticBackupStartTime() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("daily_automatic_backup_start_time"))
}

// DeploymentType returns a reference to field deployment_type of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) DeploymentType() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("deployment_type"))
}

// DnsName returns a reference to field dns_name of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("dns_name"))
}

// Id returns a reference to field id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("kms_key_id"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) NetworkInterfaceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fwfs.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("owner_id"))
}

// PreferredFileServerIp returns a reference to field preferred_file_server_ip of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) PreferredFileServerIp() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("preferred_file_server_ip"))
}

// PreferredSubnetId returns a reference to field preferred_subnet_id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) PreferredSubnetId() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("preferred_subnet_id"))
}

// RemoteAdministrationEndpoint returns a reference to field remote_administration_endpoint of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) RemoteAdministrationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("remote_administration_endpoint"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fwfs.ref.Append("security_group_ids"))
}

// SkipFinalBackup returns a reference to field skip_final_backup of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) SkipFinalBackup() terra.BoolValue {
	return terra.ReferenceAsBool(fwfs.ref.Append("skip_final_backup"))
}

// StorageCapacity returns a reference to field storage_capacity of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) StorageCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(fwfs.ref.Append("storage_capacity"))
}

// StorageType returns a reference to field storage_type of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("storage_type"))
}

// SubnetIds returns a reference to field subnet_ids of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fwfs.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fwfs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fwfs.ref.Append("tags_all"))
}

// ThroughputCapacity returns a reference to field throughput_capacity of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) ThroughputCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(fwfs.ref.Append("throughput_capacity"))
}

// VpcId returns a reference to field vpc_id of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("vpc_id"))
}

// WeeklyMaintenanceStartTime returns a reference to field weekly_maintenance_start_time of aws_fsx_windows_file_system.
func (fwfs fsxWindowsFileSystemAttributes) WeeklyMaintenanceStartTime() terra.StringValue {
	return terra.ReferenceAsString(fwfs.ref.Append("weekly_maintenance_start_time"))
}

func (fwfs fsxWindowsFileSystemAttributes) AuditLogConfiguration() terra.ListValue[fsxwindowsfilesystem.AuditLogConfigurationAttributes] {
	return terra.ReferenceAsList[fsxwindowsfilesystem.AuditLogConfigurationAttributes](fwfs.ref.Append("audit_log_configuration"))
}

func (fwfs fsxWindowsFileSystemAttributes) SelfManagedActiveDirectory() terra.ListValue[fsxwindowsfilesystem.SelfManagedActiveDirectoryAttributes] {
	return terra.ReferenceAsList[fsxwindowsfilesystem.SelfManagedActiveDirectoryAttributes](fwfs.ref.Append("self_managed_active_directory"))
}

func (fwfs fsxWindowsFileSystemAttributes) Timeouts() fsxwindowsfilesystem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxwindowsfilesystem.TimeoutsAttributes](fwfs.ref.Append("timeouts"))
}

type fsxWindowsFileSystemState struct {
	ActiveDirectoryId             string                                                 `json:"active_directory_id"`
	Aliases                       []string                                               `json:"aliases"`
	Arn                           string                                                 `json:"arn"`
	AutomaticBackupRetentionDays  float64                                                `json:"automatic_backup_retention_days"`
	BackupId                      string                                                 `json:"backup_id"`
	CopyTagsToBackups             bool                                                   `json:"copy_tags_to_backups"`
	DailyAutomaticBackupStartTime string                                                 `json:"daily_automatic_backup_start_time"`
	DeploymentType                string                                                 `json:"deployment_type"`
	DnsName                       string                                                 `json:"dns_name"`
	Id                            string                                                 `json:"id"`
	KmsKeyId                      string                                                 `json:"kms_key_id"`
	NetworkInterfaceIds           []string                                               `json:"network_interface_ids"`
	OwnerId                       string                                                 `json:"owner_id"`
	PreferredFileServerIp         string                                                 `json:"preferred_file_server_ip"`
	PreferredSubnetId             string                                                 `json:"preferred_subnet_id"`
	RemoteAdministrationEndpoint  string                                                 `json:"remote_administration_endpoint"`
	SecurityGroupIds              []string                                               `json:"security_group_ids"`
	SkipFinalBackup               bool                                                   `json:"skip_final_backup"`
	StorageCapacity               float64                                                `json:"storage_capacity"`
	StorageType                   string                                                 `json:"storage_type"`
	SubnetIds                     []string                                               `json:"subnet_ids"`
	Tags                          map[string]string                                      `json:"tags"`
	TagsAll                       map[string]string                                      `json:"tags_all"`
	ThroughputCapacity            float64                                                `json:"throughput_capacity"`
	VpcId                         string                                                 `json:"vpc_id"`
	WeeklyMaintenanceStartTime    string                                                 `json:"weekly_maintenance_start_time"`
	AuditLogConfiguration         []fsxwindowsfilesystem.AuditLogConfigurationState      `json:"audit_log_configuration"`
	SelfManagedActiveDirectory    []fsxwindowsfilesystem.SelfManagedActiveDirectoryState `json:"self_managed_active_directory"`
	Timeouts                      *fsxwindowsfilesystem.TimeoutsState                    `json:"timeouts"`
}
