// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	instance "github.com/golingon/terraproviders/aws/4.63.0/instance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInstance creates a new instance of [Instance].
func NewInstance(name string, args InstanceArgs) *Instance {
	return &Instance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Instance)(nil)

// Instance represents the Terraform resource aws_instance.
type Instance struct {
	Name      string
	Args      InstanceArgs
	state     *instanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Instance].
func (i *Instance) Type() string {
	return "aws_instance"
}

// LocalName returns the local name for [Instance].
func (i *Instance) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [Instance].
func (i *Instance) Configuration() interface{} {
	return i.Args
}

// DependOn is used for other resources to depend on [Instance].
func (i *Instance) DependOn() terra.Reference {
	return terra.ReferenceResource(i)
}

// Dependencies returns the list of resources [Instance] depends_on.
func (i *Instance) Dependencies() terra.Dependencies {
	return i.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Instance].
func (i *Instance) LifecycleManagement() *terra.Lifecycle {
	return i.Lifecycle
}

// Attributes returns the attributes for [Instance].
func (i *Instance) Attributes() instanceAttributes {
	return instanceAttributes{ref: terra.ReferenceResource(i)}
}

// ImportState imports the given attribute values into [Instance]'s state.
func (i *Instance) ImportState(av io.Reader) error {
	i.state = &instanceState{}
	if err := json.NewDecoder(av).Decode(i.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", i.Type(), i.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Instance] has state.
func (i *Instance) State() (*instanceState, bool) {
	return i.state, i.state != nil
}

// StateMust returns the state for [Instance]. Panics if the state is nil.
func (i *Instance) StateMust() *instanceState {
	if i.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", i.Type(), i.LocalName()))
	}
	return i.state
}

// InstanceArgs contains the configurations for aws_instance.
type InstanceArgs struct {
	// Ami: string, optional
	Ami terra.StringValue `hcl:"ami,attr"`
	// AssociatePublicIpAddress: bool, optional
	AssociatePublicIpAddress terra.BoolValue `hcl:"associate_public_ip_address,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// CpuCoreCount: number, optional
	CpuCoreCount terra.NumberValue `hcl:"cpu_core_count,attr"`
	// CpuThreadsPerCore: number, optional
	CpuThreadsPerCore terra.NumberValue `hcl:"cpu_threads_per_core,attr"`
	// DisableApiStop: bool, optional
	DisableApiStop terra.BoolValue `hcl:"disable_api_stop,attr"`
	// DisableApiTermination: bool, optional
	DisableApiTermination terra.BoolValue `hcl:"disable_api_termination,attr"`
	// EbsOptimized: bool, optional
	EbsOptimized terra.BoolValue `hcl:"ebs_optimized,attr"`
	// GetPasswordData: bool, optional
	GetPasswordData terra.BoolValue `hcl:"get_password_data,attr"`
	// Hibernation: bool, optional
	Hibernation terra.BoolValue `hcl:"hibernation,attr"`
	// HostId: string, optional
	HostId terra.StringValue `hcl:"host_id,attr"`
	// HostResourceGroupArn: string, optional
	HostResourceGroupArn terra.StringValue `hcl:"host_resource_group_arn,attr"`
	// IamInstanceProfile: string, optional
	IamInstanceProfile terra.StringValue `hcl:"iam_instance_profile,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceInitiatedShutdownBehavior: string, optional
	InstanceInitiatedShutdownBehavior terra.StringValue `hcl:"instance_initiated_shutdown_behavior,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// Ipv6AddressCount: number, optional
	Ipv6AddressCount terra.NumberValue `hcl:"ipv6_address_count,attr"`
	// Ipv6Addresses: list of string, optional
	Ipv6Addresses terra.ListValue[terra.StringValue] `hcl:"ipv6_addresses,attr"`
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// Monitoring: bool, optional
	Monitoring terra.BoolValue `hcl:"monitoring,attr"`
	// PlacementGroup: string, optional
	PlacementGroup terra.StringValue `hcl:"placement_group,attr"`
	// PlacementPartitionNumber: number, optional
	PlacementPartitionNumber terra.NumberValue `hcl:"placement_partition_number,attr"`
	// PrivateIp: string, optional
	PrivateIp terra.StringValue `hcl:"private_ip,attr"`
	// SecondaryPrivateIps: set of string, optional
	SecondaryPrivateIps terra.SetValue[terra.StringValue] `hcl:"secondary_private_ips,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SourceDestCheck: bool, optional
	SourceDestCheck terra.BoolValue `hcl:"source_dest_check,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Tenancy: string, optional
	Tenancy terra.StringValue `hcl:"tenancy,attr"`
	// UserData: string, optional
	UserData terra.StringValue `hcl:"user_data,attr"`
	// UserDataBase64: string, optional
	UserDataBase64 terra.StringValue `hcl:"user_data_base64,attr"`
	// UserDataReplaceOnChange: bool, optional
	UserDataReplaceOnChange terra.BoolValue `hcl:"user_data_replace_on_change,attr"`
	// VolumeTags: map of string, optional
	VolumeTags terra.MapValue[terra.StringValue] `hcl:"volume_tags,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// CapacityReservationSpecification: optional
	CapacityReservationSpecification *instance.CapacityReservationSpecification `hcl:"capacity_reservation_specification,block"`
	// CreditSpecification: optional
	CreditSpecification *instance.CreditSpecification `hcl:"credit_specification,block"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []instance.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EnclaveOptions: optional
	EnclaveOptions *instance.EnclaveOptions `hcl:"enclave_options,block"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []instance.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// LaunchTemplate: optional
	LaunchTemplate *instance.LaunchTemplate `hcl:"launch_template,block"`
	// MaintenanceOptions: optional
	MaintenanceOptions *instance.MaintenanceOptions `hcl:"maintenance_options,block"`
	// MetadataOptions: optional
	MetadataOptions *instance.MetadataOptions `hcl:"metadata_options,block"`
	// NetworkInterface: min=0
	NetworkInterface []instance.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// PrivateDnsNameOptions: optional
	PrivateDnsNameOptions *instance.PrivateDnsNameOptions `hcl:"private_dns_name_options,block"`
	// RootBlockDevice: optional
	RootBlockDevice *instance.RootBlockDevice `hcl:"root_block_device,block"`
	// Timeouts: optional
	Timeouts *instance.Timeouts `hcl:"timeouts,block"`
}
type instanceAttributes struct {
	ref terra.Reference
}

// Ami returns a reference to field ami of aws_instance.
func (i instanceAttributes) Ami() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("ami"))
}

// Arn returns a reference to field arn of aws_instance.
func (i instanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("arn"))
}

// AssociatePublicIpAddress returns a reference to field associate_public_ip_address of aws_instance.
func (i instanceAttributes) AssociatePublicIpAddress() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("associate_public_ip_address"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_instance.
func (i instanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("availability_zone"))
}

// CpuCoreCount returns a reference to field cpu_core_count of aws_instance.
func (i instanceAttributes) CpuCoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("cpu_core_count"))
}

// CpuThreadsPerCore returns a reference to field cpu_threads_per_core of aws_instance.
func (i instanceAttributes) CpuThreadsPerCore() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("cpu_threads_per_core"))
}

// DisableApiStop returns a reference to field disable_api_stop of aws_instance.
func (i instanceAttributes) DisableApiStop() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("disable_api_stop"))
}

// DisableApiTermination returns a reference to field disable_api_termination of aws_instance.
func (i instanceAttributes) DisableApiTermination() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("disable_api_termination"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_instance.
func (i instanceAttributes) EbsOptimized() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("ebs_optimized"))
}

// GetPasswordData returns a reference to field get_password_data of aws_instance.
func (i instanceAttributes) GetPasswordData() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("get_password_data"))
}

// Hibernation returns a reference to field hibernation of aws_instance.
func (i instanceAttributes) Hibernation() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("hibernation"))
}

// HostId returns a reference to field host_id of aws_instance.
func (i instanceAttributes) HostId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("host_id"))
}

// HostResourceGroupArn returns a reference to field host_resource_group_arn of aws_instance.
func (i instanceAttributes) HostResourceGroupArn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("host_resource_group_arn"))
}

// IamInstanceProfile returns a reference to field iam_instance_profile of aws_instance.
func (i instanceAttributes) IamInstanceProfile() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("iam_instance_profile"))
}

// Id returns a reference to field id of aws_instance.
func (i instanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// InstanceInitiatedShutdownBehavior returns a reference to field instance_initiated_shutdown_behavior of aws_instance.
func (i instanceAttributes) InstanceInitiatedShutdownBehavior() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("instance_initiated_shutdown_behavior"))
}

// InstanceState returns a reference to field instance_state of aws_instance.
func (i instanceAttributes) InstanceState() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("instance_state"))
}

// InstanceType returns a reference to field instance_type of aws_instance.
func (i instanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("instance_type"))
}

// Ipv6AddressCount returns a reference to field ipv6_address_count of aws_instance.
func (i instanceAttributes) Ipv6AddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("ipv6_address_count"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of aws_instance.
func (i instanceAttributes) Ipv6Addresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("ipv6_addresses"))
}

// KeyName returns a reference to field key_name of aws_instance.
func (i instanceAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("key_name"))
}

// Monitoring returns a reference to field monitoring of aws_instance.
func (i instanceAttributes) Monitoring() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("monitoring"))
}

// OutpostArn returns a reference to field outpost_arn of aws_instance.
func (i instanceAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("outpost_arn"))
}

// PasswordData returns a reference to field password_data of aws_instance.
func (i instanceAttributes) PasswordData() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("password_data"))
}

// PlacementGroup returns a reference to field placement_group of aws_instance.
func (i instanceAttributes) PlacementGroup() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("placement_group"))
}

// PlacementPartitionNumber returns a reference to field placement_partition_number of aws_instance.
func (i instanceAttributes) PlacementPartitionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("placement_partition_number"))
}

// PrimaryNetworkInterfaceId returns a reference to field primary_network_interface_id of aws_instance.
func (i instanceAttributes) PrimaryNetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("primary_network_interface_id"))
}

// PrivateDns returns a reference to field private_dns of aws_instance.
func (i instanceAttributes) PrivateDns() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("private_dns"))
}

// PrivateIp returns a reference to field private_ip of aws_instance.
func (i instanceAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("private_ip"))
}

// PublicDns returns a reference to field public_dns of aws_instance.
func (i instanceAttributes) PublicDns() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("public_dns"))
}

// PublicIp returns a reference to field public_ip of aws_instance.
func (i instanceAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("public_ip"))
}

// SecondaryPrivateIps returns a reference to field secondary_private_ips of aws_instance.
func (i instanceAttributes) SecondaryPrivateIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("secondary_private_ips"))
}

// SecurityGroups returns a reference to field security_groups of aws_instance.
func (i instanceAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("security_groups"))
}

// SourceDestCheck returns a reference to field source_dest_check of aws_instance.
func (i instanceAttributes) SourceDestCheck() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("source_dest_check"))
}

// SubnetId returns a reference to field subnet_id of aws_instance.
func (i instanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_instance.
func (i instanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_instance.
func (i instanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags_all"))
}

// Tenancy returns a reference to field tenancy of aws_instance.
func (i instanceAttributes) Tenancy() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenancy"))
}

// UserData returns a reference to field user_data of aws_instance.
func (i instanceAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_data"))
}

// UserDataBase64 returns a reference to field user_data_base64 of aws_instance.
func (i instanceAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_data_base64"))
}

// UserDataReplaceOnChange returns a reference to field user_data_replace_on_change of aws_instance.
func (i instanceAttributes) UserDataReplaceOnChange() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("user_data_replace_on_change"))
}

// VolumeTags returns a reference to field volume_tags of aws_instance.
func (i instanceAttributes) VolumeTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("volume_tags"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_instance.
func (i instanceAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("vpc_security_group_ids"))
}

func (i instanceAttributes) CapacityReservationSpecification() terra.ListValue[instance.CapacityReservationSpecificationAttributes] {
	return terra.ReferenceAsList[instance.CapacityReservationSpecificationAttributes](i.ref.Append("capacity_reservation_specification"))
}

func (i instanceAttributes) CreditSpecification() terra.ListValue[instance.CreditSpecificationAttributes] {
	return terra.ReferenceAsList[instance.CreditSpecificationAttributes](i.ref.Append("credit_specification"))
}

func (i instanceAttributes) EbsBlockDevice() terra.SetValue[instance.EbsBlockDeviceAttributes] {
	return terra.ReferenceAsSet[instance.EbsBlockDeviceAttributes](i.ref.Append("ebs_block_device"))
}

func (i instanceAttributes) EnclaveOptions() terra.ListValue[instance.EnclaveOptionsAttributes] {
	return terra.ReferenceAsList[instance.EnclaveOptionsAttributes](i.ref.Append("enclave_options"))
}

func (i instanceAttributes) EphemeralBlockDevice() terra.SetValue[instance.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceAsSet[instance.EphemeralBlockDeviceAttributes](i.ref.Append("ephemeral_block_device"))
}

func (i instanceAttributes) LaunchTemplate() terra.ListValue[instance.LaunchTemplateAttributes] {
	return terra.ReferenceAsList[instance.LaunchTemplateAttributes](i.ref.Append("launch_template"))
}

func (i instanceAttributes) MaintenanceOptions() terra.ListValue[instance.MaintenanceOptionsAttributes] {
	return terra.ReferenceAsList[instance.MaintenanceOptionsAttributes](i.ref.Append("maintenance_options"))
}

func (i instanceAttributes) MetadataOptions() terra.ListValue[instance.MetadataOptionsAttributes] {
	return terra.ReferenceAsList[instance.MetadataOptionsAttributes](i.ref.Append("metadata_options"))
}

func (i instanceAttributes) NetworkInterface() terra.SetValue[instance.NetworkInterfaceAttributes] {
	return terra.ReferenceAsSet[instance.NetworkInterfaceAttributes](i.ref.Append("network_interface"))
}

func (i instanceAttributes) PrivateDnsNameOptions() terra.ListValue[instance.PrivateDnsNameOptionsAttributes] {
	return terra.ReferenceAsList[instance.PrivateDnsNameOptionsAttributes](i.ref.Append("private_dns_name_options"))
}

func (i instanceAttributes) RootBlockDevice() terra.ListValue[instance.RootBlockDeviceAttributes] {
	return terra.ReferenceAsList[instance.RootBlockDeviceAttributes](i.ref.Append("root_block_device"))
}

func (i instanceAttributes) Timeouts() instance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[instance.TimeoutsAttributes](i.ref.Append("timeouts"))
}

type instanceState struct {
	Ami                               string                                           `json:"ami"`
	Arn                               string                                           `json:"arn"`
	AssociatePublicIpAddress          bool                                             `json:"associate_public_ip_address"`
	AvailabilityZone                  string                                           `json:"availability_zone"`
	CpuCoreCount                      float64                                          `json:"cpu_core_count"`
	CpuThreadsPerCore                 float64                                          `json:"cpu_threads_per_core"`
	DisableApiStop                    bool                                             `json:"disable_api_stop"`
	DisableApiTermination             bool                                             `json:"disable_api_termination"`
	EbsOptimized                      bool                                             `json:"ebs_optimized"`
	GetPasswordData                   bool                                             `json:"get_password_data"`
	Hibernation                       bool                                             `json:"hibernation"`
	HostId                            string                                           `json:"host_id"`
	HostResourceGroupArn              string                                           `json:"host_resource_group_arn"`
	IamInstanceProfile                string                                           `json:"iam_instance_profile"`
	Id                                string                                           `json:"id"`
	InstanceInitiatedShutdownBehavior string                                           `json:"instance_initiated_shutdown_behavior"`
	InstanceState                     string                                           `json:"instance_state"`
	InstanceType                      string                                           `json:"instance_type"`
	Ipv6AddressCount                  float64                                          `json:"ipv6_address_count"`
	Ipv6Addresses                     []string                                         `json:"ipv6_addresses"`
	KeyName                           string                                           `json:"key_name"`
	Monitoring                        bool                                             `json:"monitoring"`
	OutpostArn                        string                                           `json:"outpost_arn"`
	PasswordData                      string                                           `json:"password_data"`
	PlacementGroup                    string                                           `json:"placement_group"`
	PlacementPartitionNumber          float64                                          `json:"placement_partition_number"`
	PrimaryNetworkInterfaceId         string                                           `json:"primary_network_interface_id"`
	PrivateDns                        string                                           `json:"private_dns"`
	PrivateIp                         string                                           `json:"private_ip"`
	PublicDns                         string                                           `json:"public_dns"`
	PublicIp                          string                                           `json:"public_ip"`
	SecondaryPrivateIps               []string                                         `json:"secondary_private_ips"`
	SecurityGroups                    []string                                         `json:"security_groups"`
	SourceDestCheck                   bool                                             `json:"source_dest_check"`
	SubnetId                          string                                           `json:"subnet_id"`
	Tags                              map[string]string                                `json:"tags"`
	TagsAll                           map[string]string                                `json:"tags_all"`
	Tenancy                           string                                           `json:"tenancy"`
	UserData                          string                                           `json:"user_data"`
	UserDataBase64                    string                                           `json:"user_data_base64"`
	UserDataReplaceOnChange           bool                                             `json:"user_data_replace_on_change"`
	VolumeTags                        map[string]string                                `json:"volume_tags"`
	VpcSecurityGroupIds               []string                                         `json:"vpc_security_group_ids"`
	CapacityReservationSpecification  []instance.CapacityReservationSpecificationState `json:"capacity_reservation_specification"`
	CreditSpecification               []instance.CreditSpecificationState              `json:"credit_specification"`
	EbsBlockDevice                    []instance.EbsBlockDeviceState                   `json:"ebs_block_device"`
	EnclaveOptions                    []instance.EnclaveOptionsState                   `json:"enclave_options"`
	EphemeralBlockDevice              []instance.EphemeralBlockDeviceState             `json:"ephemeral_block_device"`
	LaunchTemplate                    []instance.LaunchTemplateState                   `json:"launch_template"`
	MaintenanceOptions                []instance.MaintenanceOptionsState               `json:"maintenance_options"`
	MetadataOptions                   []instance.MetadataOptionsState                  `json:"metadata_options"`
	NetworkInterface                  []instance.NetworkInterfaceState                 `json:"network_interface"`
	PrivateDnsNameOptions             []instance.PrivateDnsNameOptionsState            `json:"private_dns_name_options"`
	RootBlockDevice                   []instance.RootBlockDeviceState                  `json:"root_block_device"`
	Timeouts                          *instance.TimeoutsState                          `json:"timeouts"`
}
