// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_fsx_openzfs_file_system

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_fsx_openzfs_file_system.
type Resource struct {
	Name      string
	Args      Args
	state     *awsFsxOpenzfsFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (afofs *Resource) Type() string {
	return "aws_fsx_openzfs_file_system"
}

// LocalName returns the local name for [Resource].
func (afofs *Resource) LocalName() string {
	return afofs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (afofs *Resource) Configuration() interface{} {
	return afofs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (afofs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(afofs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (afofs *Resource) Dependencies() terra.Dependencies {
	return afofs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (afofs *Resource) LifecycleManagement() *terra.Lifecycle {
	return afofs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (afofs *Resource) Attributes() awsFsxOpenzfsFileSystemAttributes {
	return awsFsxOpenzfsFileSystemAttributes{ref: terra.ReferenceResource(afofs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (afofs *Resource) ImportState(state io.Reader) error {
	afofs.state = &awsFsxOpenzfsFileSystemState{}
	if err := json.NewDecoder(state).Decode(afofs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afofs.Type(), afofs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (afofs *Resource) State() (*awsFsxOpenzfsFileSystemState, bool) {
	return afofs.state, afofs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (afofs *Resource) StateMust() *awsFsxOpenzfsFileSystemState {
	if afofs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afofs.Type(), afofs.LocalName()))
	}
	return afofs.state
}

// Args contains the configurations for aws_fsx_openzfs_file_system.
type Args struct {
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
	// EndpointIpAddressRange: string, optional
	EndpointIpAddressRange terra.StringValue `hcl:"endpoint_ip_address_range,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PreferredSubnetId: string, optional
	PreferredSubnetId terra.StringValue `hcl:"preferred_subnet_id,attr"`
	// RouteTableIds: set of string, optional
	RouteTableIds terra.SetValue[terra.StringValue] `hcl:"route_table_ids,attr"`
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
	// DiskIopsConfiguration: optional
	DiskIopsConfiguration *DiskIopsConfiguration `hcl:"disk_iops_configuration,block"`
	// RootVolumeConfiguration: optional
	RootVolumeConfiguration *RootVolumeConfiguration `hcl:"root_volume_configuration,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsFsxOpenzfsFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("arn"))
}

// AutomaticBackupRetentionDays returns a reference to field automatic_backup_retention_days of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) AutomaticBackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(afofs.ref.Append("automatic_backup_retention_days"))
}

// BackupId returns a reference to field backup_id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) BackupId() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("backup_id"))
}

// CopyTagsToBackups returns a reference to field copy_tags_to_backups of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) CopyTagsToBackups() terra.BoolValue {
	return terra.ReferenceAsBool(afofs.ref.Append("copy_tags_to_backups"))
}

// CopyTagsToVolumes returns a reference to field copy_tags_to_volumes of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) CopyTagsToVolumes() terra.BoolValue {
	return terra.ReferenceAsBool(afofs.ref.Append("copy_tags_to_volumes"))
}

// DailyAutomaticBackupStartTime returns a reference to field daily_automatic_backup_start_time of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) DailyAutomaticBackupStartTime() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("daily_automatic_backup_start_time"))
}

// DeploymentType returns a reference to field deployment_type of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) DeploymentType() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("deployment_type"))
}

// DnsName returns a reference to field dns_name of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("dns_name"))
}

// EndpointIpAddressRange returns a reference to field endpoint_ip_address_range of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) EndpointIpAddressRange() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("endpoint_ip_address_range"))
}

// Id returns a reference to field id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("kms_key_id"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](afofs.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("owner_id"))
}

// PreferredSubnetId returns a reference to field preferred_subnet_id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) PreferredSubnetId() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("preferred_subnet_id"))
}

// RootVolumeId returns a reference to field root_volume_id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) RootVolumeId() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("root_volume_id"))
}

// RouteTableIds returns a reference to field route_table_ids of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) RouteTableIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](afofs.ref.Append("route_table_ids"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](afofs.ref.Append("security_group_ids"))
}

// SkipFinalBackup returns a reference to field skip_final_backup of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) SkipFinalBackup() terra.BoolValue {
	return terra.ReferenceAsBool(afofs.ref.Append("skip_final_backup"))
}

// StorageCapacity returns a reference to field storage_capacity of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) StorageCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(afofs.ref.Append("storage_capacity"))
}

// StorageType returns a reference to field storage_type of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("storage_type"))
}

// SubnetIds returns a reference to field subnet_ids of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](afofs.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afofs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afofs.ref.Append("tags_all"))
}

// ThroughputCapacity returns a reference to field throughput_capacity of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) ThroughputCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(afofs.ref.Append("throughput_capacity"))
}

// VpcId returns a reference to field vpc_id of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("vpc_id"))
}

// WeeklyMaintenanceStartTime returns a reference to field weekly_maintenance_start_time of aws_fsx_openzfs_file_system.
func (afofs awsFsxOpenzfsFileSystemAttributes) WeeklyMaintenanceStartTime() terra.StringValue {
	return terra.ReferenceAsString(afofs.ref.Append("weekly_maintenance_start_time"))
}

func (afofs awsFsxOpenzfsFileSystemAttributes) DiskIopsConfiguration() terra.ListValue[DiskIopsConfigurationAttributes] {
	return terra.ReferenceAsList[DiskIopsConfigurationAttributes](afofs.ref.Append("disk_iops_configuration"))
}

func (afofs awsFsxOpenzfsFileSystemAttributes) RootVolumeConfiguration() terra.ListValue[RootVolumeConfigurationAttributes] {
	return terra.ReferenceAsList[RootVolumeConfigurationAttributes](afofs.ref.Append("root_volume_configuration"))
}

func (afofs awsFsxOpenzfsFileSystemAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](afofs.ref.Append("timeouts"))
}

type awsFsxOpenzfsFileSystemState struct {
	Arn                           string                         `json:"arn"`
	AutomaticBackupRetentionDays  float64                        `json:"automatic_backup_retention_days"`
	BackupId                      string                         `json:"backup_id"`
	CopyTagsToBackups             bool                           `json:"copy_tags_to_backups"`
	CopyTagsToVolumes             bool                           `json:"copy_tags_to_volumes"`
	DailyAutomaticBackupStartTime string                         `json:"daily_automatic_backup_start_time"`
	DeploymentType                string                         `json:"deployment_type"`
	DnsName                       string                         `json:"dns_name"`
	EndpointIpAddressRange        string                         `json:"endpoint_ip_address_range"`
	Id                            string                         `json:"id"`
	KmsKeyId                      string                         `json:"kms_key_id"`
	NetworkInterfaceIds           []string                       `json:"network_interface_ids"`
	OwnerId                       string                         `json:"owner_id"`
	PreferredSubnetId             string                         `json:"preferred_subnet_id"`
	RootVolumeId                  string                         `json:"root_volume_id"`
	RouteTableIds                 []string                       `json:"route_table_ids"`
	SecurityGroupIds              []string                       `json:"security_group_ids"`
	SkipFinalBackup               bool                           `json:"skip_final_backup"`
	StorageCapacity               float64                        `json:"storage_capacity"`
	StorageType                   string                         `json:"storage_type"`
	SubnetIds                     []string                       `json:"subnet_ids"`
	Tags                          map[string]string              `json:"tags"`
	TagsAll                       map[string]string              `json:"tags_all"`
	ThroughputCapacity            float64                        `json:"throughput_capacity"`
	VpcId                         string                         `json:"vpc_id"`
	WeeklyMaintenanceStartTime    string                         `json:"weekly_maintenance_start_time"`
	DiskIopsConfiguration         []DiskIopsConfigurationState   `json:"disk_iops_configuration"`
	RootVolumeConfiguration       []RootVolumeConfigurationState `json:"root_volume_configuration"`
	Timeouts                      *TimeoutsState                 `json:"timeouts"`
}
