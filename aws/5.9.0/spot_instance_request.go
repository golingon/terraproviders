// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	spotinstancerequest "github.com/golingon/terraproviders/aws/5.9.0/spotinstancerequest"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpotInstanceRequest creates a new instance of [SpotInstanceRequest].
func NewSpotInstanceRequest(name string, args SpotInstanceRequestArgs) *SpotInstanceRequest {
	return &SpotInstanceRequest{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpotInstanceRequest)(nil)

// SpotInstanceRequest represents the Terraform resource aws_spot_instance_request.
type SpotInstanceRequest struct {
	Name      string
	Args      SpotInstanceRequestArgs
	state     *spotInstanceRequestState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpotInstanceRequest].
func (sir *SpotInstanceRequest) Type() string {
	return "aws_spot_instance_request"
}

// LocalName returns the local name for [SpotInstanceRequest].
func (sir *SpotInstanceRequest) LocalName() string {
	return sir.Name
}

// Configuration returns the configuration (args) for [SpotInstanceRequest].
func (sir *SpotInstanceRequest) Configuration() interface{} {
	return sir.Args
}

// DependOn is used for other resources to depend on [SpotInstanceRequest].
func (sir *SpotInstanceRequest) DependOn() terra.Reference {
	return terra.ReferenceResource(sir)
}

// Dependencies returns the list of resources [SpotInstanceRequest] depends_on.
func (sir *SpotInstanceRequest) Dependencies() terra.Dependencies {
	return sir.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpotInstanceRequest].
func (sir *SpotInstanceRequest) LifecycleManagement() *terra.Lifecycle {
	return sir.Lifecycle
}

// Attributes returns the attributes for [SpotInstanceRequest].
func (sir *SpotInstanceRequest) Attributes() spotInstanceRequestAttributes {
	return spotInstanceRequestAttributes{ref: terra.ReferenceResource(sir)}
}

// ImportState imports the given attribute values into [SpotInstanceRequest]'s state.
func (sir *SpotInstanceRequest) ImportState(av io.Reader) error {
	sir.state = &spotInstanceRequestState{}
	if err := json.NewDecoder(av).Decode(sir.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sir.Type(), sir.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpotInstanceRequest] has state.
func (sir *SpotInstanceRequest) State() (*spotInstanceRequestState, bool) {
	return sir.state, sir.state != nil
}

// StateMust returns the state for [SpotInstanceRequest]. Panics if the state is nil.
func (sir *SpotInstanceRequest) StateMust() *spotInstanceRequestState {
	if sir.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sir.Type(), sir.LocalName()))
	}
	return sir.state
}

// SpotInstanceRequestArgs contains the configurations for aws_spot_instance_request.
type SpotInstanceRequestArgs struct {
	// Ami: string, optional
	Ami terra.StringValue `hcl:"ami,attr"`
	// AssociatePublicIpAddress: bool, optional
	AssociatePublicIpAddress terra.BoolValue `hcl:"associate_public_ip_address,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// BlockDurationMinutes: number, optional
	BlockDurationMinutes terra.NumberValue `hcl:"block_duration_minutes,attr"`
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
	// InstanceInterruptionBehavior: string, optional
	InstanceInterruptionBehavior terra.StringValue `hcl:"instance_interruption_behavior,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// Ipv6AddressCount: number, optional
	Ipv6AddressCount terra.NumberValue `hcl:"ipv6_address_count,attr"`
	// Ipv6Addresses: list of string, optional
	Ipv6Addresses terra.ListValue[terra.StringValue] `hcl:"ipv6_addresses,attr"`
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// LaunchGroup: string, optional
	LaunchGroup terra.StringValue `hcl:"launch_group,attr"`
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
	// SpotPrice: string, optional
	SpotPrice terra.StringValue `hcl:"spot_price,attr"`
	// SpotType: string, optional
	SpotType terra.StringValue `hcl:"spot_type,attr"`
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
	// ValidFrom: string, optional
	ValidFrom terra.StringValue `hcl:"valid_from,attr"`
	// ValidUntil: string, optional
	ValidUntil terra.StringValue `hcl:"valid_until,attr"`
	// VolumeTags: map of string, optional
	VolumeTags terra.MapValue[terra.StringValue] `hcl:"volume_tags,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// WaitForFulfillment: bool, optional
	WaitForFulfillment terra.BoolValue `hcl:"wait_for_fulfillment,attr"`
	// CapacityReservationSpecification: optional
	CapacityReservationSpecification *spotinstancerequest.CapacityReservationSpecification `hcl:"capacity_reservation_specification,block"`
	// CpuOptions: optional
	CpuOptions *spotinstancerequest.CpuOptions `hcl:"cpu_options,block"`
	// CreditSpecification: optional
	CreditSpecification *spotinstancerequest.CreditSpecification `hcl:"credit_specification,block"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []spotinstancerequest.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EnclaveOptions: optional
	EnclaveOptions *spotinstancerequest.EnclaveOptions `hcl:"enclave_options,block"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []spotinstancerequest.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// LaunchTemplate: optional
	LaunchTemplate *spotinstancerequest.LaunchTemplate `hcl:"launch_template,block"`
	// MaintenanceOptions: optional
	MaintenanceOptions *spotinstancerequest.MaintenanceOptions `hcl:"maintenance_options,block"`
	// MetadataOptions: optional
	MetadataOptions *spotinstancerequest.MetadataOptions `hcl:"metadata_options,block"`
	// NetworkInterface: min=0
	NetworkInterface []spotinstancerequest.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// PrivateDnsNameOptions: optional
	PrivateDnsNameOptions *spotinstancerequest.PrivateDnsNameOptions `hcl:"private_dns_name_options,block"`
	// RootBlockDevice: optional
	RootBlockDevice *spotinstancerequest.RootBlockDevice `hcl:"root_block_device,block"`
	// Timeouts: optional
	Timeouts *spotinstancerequest.Timeouts `hcl:"timeouts,block"`
}
type spotInstanceRequestAttributes struct {
	ref terra.Reference
}

// Ami returns a reference to field ami of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Ami() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("ami"))
}

// Arn returns a reference to field arn of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("arn"))
}

// AssociatePublicIpAddress returns a reference to field associate_public_ip_address of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) AssociatePublicIpAddress() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("associate_public_ip_address"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("availability_zone"))
}

// BlockDurationMinutes returns a reference to field block_duration_minutes of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) BlockDurationMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(sir.ref.Append("block_duration_minutes"))
}

// CpuCoreCount returns a reference to field cpu_core_count of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) CpuCoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sir.ref.Append("cpu_core_count"))
}

// CpuThreadsPerCore returns a reference to field cpu_threads_per_core of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) CpuThreadsPerCore() terra.NumberValue {
	return terra.ReferenceAsNumber(sir.ref.Append("cpu_threads_per_core"))
}

// DisableApiStop returns a reference to field disable_api_stop of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) DisableApiStop() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("disable_api_stop"))
}

// DisableApiTermination returns a reference to field disable_api_termination of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) DisableApiTermination() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("disable_api_termination"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) EbsOptimized() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("ebs_optimized"))
}

// GetPasswordData returns a reference to field get_password_data of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) GetPasswordData() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("get_password_data"))
}

// Hibernation returns a reference to field hibernation of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Hibernation() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("hibernation"))
}

// HostId returns a reference to field host_id of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) HostId() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("host_id"))
}

// HostResourceGroupArn returns a reference to field host_resource_group_arn of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) HostResourceGroupArn() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("host_resource_group_arn"))
}

// IamInstanceProfile returns a reference to field iam_instance_profile of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) IamInstanceProfile() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("iam_instance_profile"))
}

// Id returns a reference to field id of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("id"))
}

// InstanceInitiatedShutdownBehavior returns a reference to field instance_initiated_shutdown_behavior of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) InstanceInitiatedShutdownBehavior() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("instance_initiated_shutdown_behavior"))
}

// InstanceInterruptionBehavior returns a reference to field instance_interruption_behavior of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) InstanceInterruptionBehavior() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("instance_interruption_behavior"))
}

// InstanceState returns a reference to field instance_state of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) InstanceState() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("instance_state"))
}

// InstanceType returns a reference to field instance_type of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("instance_type"))
}

// Ipv6AddressCount returns a reference to field ipv6_address_count of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Ipv6AddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sir.ref.Append("ipv6_address_count"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Ipv6Addresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sir.ref.Append("ipv6_addresses"))
}

// KeyName returns a reference to field key_name of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("key_name"))
}

// LaunchGroup returns a reference to field launch_group of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) LaunchGroup() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("launch_group"))
}

// Monitoring returns a reference to field monitoring of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Monitoring() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("monitoring"))
}

// OutpostArn returns a reference to field outpost_arn of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("outpost_arn"))
}

// PasswordData returns a reference to field password_data of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PasswordData() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("password_data"))
}

// PlacementGroup returns a reference to field placement_group of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PlacementGroup() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("placement_group"))
}

// PlacementPartitionNumber returns a reference to field placement_partition_number of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PlacementPartitionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(sir.ref.Append("placement_partition_number"))
}

// PrimaryNetworkInterfaceId returns a reference to field primary_network_interface_id of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PrimaryNetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("primary_network_interface_id"))
}

// PrivateDns returns a reference to field private_dns of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PrivateDns() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("private_dns"))
}

// PrivateIp returns a reference to field private_ip of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("private_ip"))
}

// PublicDns returns a reference to field public_dns of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PublicDns() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("public_dns"))
}

// PublicIp returns a reference to field public_ip of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("public_ip"))
}

// SecondaryPrivateIps returns a reference to field secondary_private_ips of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SecondaryPrivateIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sir.ref.Append("secondary_private_ips"))
}

// SecurityGroups returns a reference to field security_groups of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sir.ref.Append("security_groups"))
}

// SourceDestCheck returns a reference to field source_dest_check of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SourceDestCheck() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("source_dest_check"))
}

// SpotBidStatus returns a reference to field spot_bid_status of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SpotBidStatus() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("spot_bid_status"))
}

// SpotInstanceId returns a reference to field spot_instance_id of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SpotInstanceId() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("spot_instance_id"))
}

// SpotPrice returns a reference to field spot_price of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SpotPrice() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("spot_price"))
}

// SpotRequestState returns a reference to field spot_request_state of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SpotRequestState() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("spot_request_state"))
}

// SpotType returns a reference to field spot_type of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SpotType() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("spot_type"))
}

// SubnetId returns a reference to field subnet_id of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sir.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sir.ref.Append("tags_all"))
}

// Tenancy returns a reference to field tenancy of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) Tenancy() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("tenancy"))
}

// UserData returns a reference to field user_data of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("user_data"))
}

// UserDataBase64 returns a reference to field user_data_base64 of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("user_data_base64"))
}

// UserDataReplaceOnChange returns a reference to field user_data_replace_on_change of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) UserDataReplaceOnChange() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("user_data_replace_on_change"))
}

// ValidFrom returns a reference to field valid_from of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) ValidFrom() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("valid_from"))
}

// ValidUntil returns a reference to field valid_until of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) ValidUntil() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("valid_until"))
}

// VolumeTags returns a reference to field volume_tags of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) VolumeTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sir.ref.Append("volume_tags"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sir.ref.Append("vpc_security_group_ids"))
}

// WaitForFulfillment returns a reference to field wait_for_fulfillment of aws_spot_instance_request.
func (sir spotInstanceRequestAttributes) WaitForFulfillment() terra.BoolValue {
	return terra.ReferenceAsBool(sir.ref.Append("wait_for_fulfillment"))
}

func (sir spotInstanceRequestAttributes) CapacityReservationSpecification() terra.ListValue[spotinstancerequest.CapacityReservationSpecificationAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.CapacityReservationSpecificationAttributes](sir.ref.Append("capacity_reservation_specification"))
}

func (sir spotInstanceRequestAttributes) CpuOptions() terra.ListValue[spotinstancerequest.CpuOptionsAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.CpuOptionsAttributes](sir.ref.Append("cpu_options"))
}

func (sir spotInstanceRequestAttributes) CreditSpecification() terra.ListValue[spotinstancerequest.CreditSpecificationAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.CreditSpecificationAttributes](sir.ref.Append("credit_specification"))
}

func (sir spotInstanceRequestAttributes) EbsBlockDevice() terra.SetValue[spotinstancerequest.EbsBlockDeviceAttributes] {
	return terra.ReferenceAsSet[spotinstancerequest.EbsBlockDeviceAttributes](sir.ref.Append("ebs_block_device"))
}

func (sir spotInstanceRequestAttributes) EnclaveOptions() terra.ListValue[spotinstancerequest.EnclaveOptionsAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.EnclaveOptionsAttributes](sir.ref.Append("enclave_options"))
}

func (sir spotInstanceRequestAttributes) EphemeralBlockDevice() terra.SetValue[spotinstancerequest.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceAsSet[spotinstancerequest.EphemeralBlockDeviceAttributes](sir.ref.Append("ephemeral_block_device"))
}

func (sir spotInstanceRequestAttributes) LaunchTemplate() terra.ListValue[spotinstancerequest.LaunchTemplateAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.LaunchTemplateAttributes](sir.ref.Append("launch_template"))
}

func (sir spotInstanceRequestAttributes) MaintenanceOptions() terra.ListValue[spotinstancerequest.MaintenanceOptionsAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.MaintenanceOptionsAttributes](sir.ref.Append("maintenance_options"))
}

func (sir spotInstanceRequestAttributes) MetadataOptions() terra.ListValue[spotinstancerequest.MetadataOptionsAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.MetadataOptionsAttributes](sir.ref.Append("metadata_options"))
}

func (sir spotInstanceRequestAttributes) NetworkInterface() terra.SetValue[spotinstancerequest.NetworkInterfaceAttributes] {
	return terra.ReferenceAsSet[spotinstancerequest.NetworkInterfaceAttributes](sir.ref.Append("network_interface"))
}

func (sir spotInstanceRequestAttributes) PrivateDnsNameOptions() terra.ListValue[spotinstancerequest.PrivateDnsNameOptionsAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.PrivateDnsNameOptionsAttributes](sir.ref.Append("private_dns_name_options"))
}

func (sir spotInstanceRequestAttributes) RootBlockDevice() terra.ListValue[spotinstancerequest.RootBlockDeviceAttributes] {
	return terra.ReferenceAsList[spotinstancerequest.RootBlockDeviceAttributes](sir.ref.Append("root_block_device"))
}

func (sir spotInstanceRequestAttributes) Timeouts() spotinstancerequest.TimeoutsAttributes {
	return terra.ReferenceAsSingle[spotinstancerequest.TimeoutsAttributes](sir.ref.Append("timeouts"))
}

type spotInstanceRequestState struct {
	Ami                               string                                                      `json:"ami"`
	Arn                               string                                                      `json:"arn"`
	AssociatePublicIpAddress          bool                                                        `json:"associate_public_ip_address"`
	AvailabilityZone                  string                                                      `json:"availability_zone"`
	BlockDurationMinutes              float64                                                     `json:"block_duration_minutes"`
	CpuCoreCount                      float64                                                     `json:"cpu_core_count"`
	CpuThreadsPerCore                 float64                                                     `json:"cpu_threads_per_core"`
	DisableApiStop                    bool                                                        `json:"disable_api_stop"`
	DisableApiTermination             bool                                                        `json:"disable_api_termination"`
	EbsOptimized                      bool                                                        `json:"ebs_optimized"`
	GetPasswordData                   bool                                                        `json:"get_password_data"`
	Hibernation                       bool                                                        `json:"hibernation"`
	HostId                            string                                                      `json:"host_id"`
	HostResourceGroupArn              string                                                      `json:"host_resource_group_arn"`
	IamInstanceProfile                string                                                      `json:"iam_instance_profile"`
	Id                                string                                                      `json:"id"`
	InstanceInitiatedShutdownBehavior string                                                      `json:"instance_initiated_shutdown_behavior"`
	InstanceInterruptionBehavior      string                                                      `json:"instance_interruption_behavior"`
	InstanceState                     string                                                      `json:"instance_state"`
	InstanceType                      string                                                      `json:"instance_type"`
	Ipv6AddressCount                  float64                                                     `json:"ipv6_address_count"`
	Ipv6Addresses                     []string                                                    `json:"ipv6_addresses"`
	KeyName                           string                                                      `json:"key_name"`
	LaunchGroup                       string                                                      `json:"launch_group"`
	Monitoring                        bool                                                        `json:"monitoring"`
	OutpostArn                        string                                                      `json:"outpost_arn"`
	PasswordData                      string                                                      `json:"password_data"`
	PlacementGroup                    string                                                      `json:"placement_group"`
	PlacementPartitionNumber          float64                                                     `json:"placement_partition_number"`
	PrimaryNetworkInterfaceId         string                                                      `json:"primary_network_interface_id"`
	PrivateDns                        string                                                      `json:"private_dns"`
	PrivateIp                         string                                                      `json:"private_ip"`
	PublicDns                         string                                                      `json:"public_dns"`
	PublicIp                          string                                                      `json:"public_ip"`
	SecondaryPrivateIps               []string                                                    `json:"secondary_private_ips"`
	SecurityGroups                    []string                                                    `json:"security_groups"`
	SourceDestCheck                   bool                                                        `json:"source_dest_check"`
	SpotBidStatus                     string                                                      `json:"spot_bid_status"`
	SpotInstanceId                    string                                                      `json:"spot_instance_id"`
	SpotPrice                         string                                                      `json:"spot_price"`
	SpotRequestState                  string                                                      `json:"spot_request_state"`
	SpotType                          string                                                      `json:"spot_type"`
	SubnetId                          string                                                      `json:"subnet_id"`
	Tags                              map[string]string                                           `json:"tags"`
	TagsAll                           map[string]string                                           `json:"tags_all"`
	Tenancy                           string                                                      `json:"tenancy"`
	UserData                          string                                                      `json:"user_data"`
	UserDataBase64                    string                                                      `json:"user_data_base64"`
	UserDataReplaceOnChange           bool                                                        `json:"user_data_replace_on_change"`
	ValidFrom                         string                                                      `json:"valid_from"`
	ValidUntil                        string                                                      `json:"valid_until"`
	VolumeTags                        map[string]string                                           `json:"volume_tags"`
	VpcSecurityGroupIds               []string                                                    `json:"vpc_security_group_ids"`
	WaitForFulfillment                bool                                                        `json:"wait_for_fulfillment"`
	CapacityReservationSpecification  []spotinstancerequest.CapacityReservationSpecificationState `json:"capacity_reservation_specification"`
	CpuOptions                        []spotinstancerequest.CpuOptionsState                       `json:"cpu_options"`
	CreditSpecification               []spotinstancerequest.CreditSpecificationState              `json:"credit_specification"`
	EbsBlockDevice                    []spotinstancerequest.EbsBlockDeviceState                   `json:"ebs_block_device"`
	EnclaveOptions                    []spotinstancerequest.EnclaveOptionsState                   `json:"enclave_options"`
	EphemeralBlockDevice              []spotinstancerequest.EphemeralBlockDeviceState             `json:"ephemeral_block_device"`
	LaunchTemplate                    []spotinstancerequest.LaunchTemplateState                   `json:"launch_template"`
	MaintenanceOptions                []spotinstancerequest.MaintenanceOptionsState               `json:"maintenance_options"`
	MetadataOptions                   []spotinstancerequest.MetadataOptionsState                  `json:"metadata_options"`
	NetworkInterface                  []spotinstancerequest.NetworkInterfaceState                 `json:"network_interface"`
	PrivateDnsNameOptions             []spotinstancerequest.PrivateDnsNameOptionsState            `json:"private_dns_name_options"`
	RootBlockDevice                   []spotinstancerequest.RootBlockDeviceState                  `json:"root_block_device"`
	Timeouts                          *spotinstancerequest.TimeoutsState                          `json:"timeouts"`
}
