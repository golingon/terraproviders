// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datainstance "github.com/golingon/terraproviders/aws/5.8.0/datainstance"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataInstance creates a new instance of [DataInstance].
func NewDataInstance(name string, args DataInstanceArgs) *DataInstance {
	return &DataInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataInstance)(nil)

// DataInstance represents the Terraform data resource aws_instance.
type DataInstance struct {
	Name string
	Args DataInstanceArgs
}

// DataSource returns the Terraform object type for [DataInstance].
func (i *DataInstance) DataSource() string {
	return "aws_instance"
}

// LocalName returns the local name for [DataInstance].
func (i *DataInstance) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [DataInstance].
func (i *DataInstance) Configuration() interface{} {
	return i.Args
}

// Attributes returns the attributes for [DataInstance].
func (i *DataInstance) Attributes() dataInstanceAttributes {
	return dataInstanceAttributes{ref: terra.ReferenceDataResource(i)}
}

// DataInstanceArgs contains the configurations for aws_instance.
type DataInstanceArgs struct {
	// GetPasswordData: bool, optional
	GetPasswordData terra.BoolValue `hcl:"get_password_data,attr"`
	// GetUserData: bool, optional
	GetUserData terra.BoolValue `hcl:"get_user_data,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, optional
	InstanceId terra.StringValue `hcl:"instance_id,attr"`
	// InstanceTags: map of string, optional
	InstanceTags terra.MapValue[terra.StringValue] `hcl:"instance_tags,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CreditSpecification: min=0
	CreditSpecification []datainstance.CreditSpecification `hcl:"credit_specification,block" validate:"min=0"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []datainstance.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EnclaveOptions: min=0
	EnclaveOptions []datainstance.EnclaveOptions `hcl:"enclave_options,block" validate:"min=0"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []datainstance.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// MaintenanceOptions: min=0
	MaintenanceOptions []datainstance.MaintenanceOptions `hcl:"maintenance_options,block" validate:"min=0"`
	// MetadataOptions: min=0
	MetadataOptions []datainstance.MetadataOptions `hcl:"metadata_options,block" validate:"min=0"`
	// PrivateDnsNameOptions: min=0
	PrivateDnsNameOptions []datainstance.PrivateDnsNameOptions `hcl:"private_dns_name_options,block" validate:"min=0"`
	// RootBlockDevice: min=0
	RootBlockDevice []datainstance.RootBlockDevice `hcl:"root_block_device,block" validate:"min=0"`
	// Filter: min=0
	Filter []datainstance.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datainstance.Timeouts `hcl:"timeouts,block"`
}
type dataInstanceAttributes struct {
	ref terra.Reference
}

// Ami returns a reference to field ami of aws_instance.
func (i dataInstanceAttributes) Ami() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("ami"))
}

// Arn returns a reference to field arn of aws_instance.
func (i dataInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("arn"))
}

// AssociatePublicIpAddress returns a reference to field associate_public_ip_address of aws_instance.
func (i dataInstanceAttributes) AssociatePublicIpAddress() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("associate_public_ip_address"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_instance.
func (i dataInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("availability_zone"))
}

// DisableApiStop returns a reference to field disable_api_stop of aws_instance.
func (i dataInstanceAttributes) DisableApiStop() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("disable_api_stop"))
}

// DisableApiTermination returns a reference to field disable_api_termination of aws_instance.
func (i dataInstanceAttributes) DisableApiTermination() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("disable_api_termination"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_instance.
func (i dataInstanceAttributes) EbsOptimized() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("ebs_optimized"))
}

// GetPasswordData returns a reference to field get_password_data of aws_instance.
func (i dataInstanceAttributes) GetPasswordData() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("get_password_data"))
}

// GetUserData returns a reference to field get_user_data of aws_instance.
func (i dataInstanceAttributes) GetUserData() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("get_user_data"))
}

// HostId returns a reference to field host_id of aws_instance.
func (i dataInstanceAttributes) HostId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("host_id"))
}

// HostResourceGroupArn returns a reference to field host_resource_group_arn of aws_instance.
func (i dataInstanceAttributes) HostResourceGroupArn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("host_resource_group_arn"))
}

// IamInstanceProfile returns a reference to field iam_instance_profile of aws_instance.
func (i dataInstanceAttributes) IamInstanceProfile() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("iam_instance_profile"))
}

// Id returns a reference to field id of aws_instance.
func (i dataInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_instance.
func (i dataInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("instance_id"))
}

// InstanceState returns a reference to field instance_state of aws_instance.
func (i dataInstanceAttributes) InstanceState() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("instance_state"))
}

// InstanceTags returns a reference to field instance_tags of aws_instance.
func (i dataInstanceAttributes) InstanceTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("instance_tags"))
}

// InstanceType returns a reference to field instance_type of aws_instance.
func (i dataInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("instance_type"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of aws_instance.
func (i dataInstanceAttributes) Ipv6Addresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("ipv6_addresses"))
}

// KeyName returns a reference to field key_name of aws_instance.
func (i dataInstanceAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("key_name"))
}

// Monitoring returns a reference to field monitoring of aws_instance.
func (i dataInstanceAttributes) Monitoring() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("monitoring"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_instance.
func (i dataInstanceAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("network_interface_id"))
}

// OutpostArn returns a reference to field outpost_arn of aws_instance.
func (i dataInstanceAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("outpost_arn"))
}

// PasswordData returns a reference to field password_data of aws_instance.
func (i dataInstanceAttributes) PasswordData() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("password_data"))
}

// PlacementGroup returns a reference to field placement_group of aws_instance.
func (i dataInstanceAttributes) PlacementGroup() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("placement_group"))
}

// PlacementPartitionNumber returns a reference to field placement_partition_number of aws_instance.
func (i dataInstanceAttributes) PlacementPartitionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("placement_partition_number"))
}

// PrivateDns returns a reference to field private_dns of aws_instance.
func (i dataInstanceAttributes) PrivateDns() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("private_dns"))
}

// PrivateIp returns a reference to field private_ip of aws_instance.
func (i dataInstanceAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("private_ip"))
}

// PublicDns returns a reference to field public_dns of aws_instance.
func (i dataInstanceAttributes) PublicDns() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("public_dns"))
}

// PublicIp returns a reference to field public_ip of aws_instance.
func (i dataInstanceAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("public_ip"))
}

// SecondaryPrivateIps returns a reference to field secondary_private_ips of aws_instance.
func (i dataInstanceAttributes) SecondaryPrivateIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("secondary_private_ips"))
}

// SecurityGroups returns a reference to field security_groups of aws_instance.
func (i dataInstanceAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("security_groups"))
}

// SourceDestCheck returns a reference to field source_dest_check of aws_instance.
func (i dataInstanceAttributes) SourceDestCheck() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("source_dest_check"))
}

// SubnetId returns a reference to field subnet_id of aws_instance.
func (i dataInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_instance.
func (i dataInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

// Tenancy returns a reference to field tenancy of aws_instance.
func (i dataInstanceAttributes) Tenancy() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenancy"))
}

// UserData returns a reference to field user_data of aws_instance.
func (i dataInstanceAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_data"))
}

// UserDataBase64 returns a reference to field user_data_base64 of aws_instance.
func (i dataInstanceAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("user_data_base64"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_instance.
func (i dataInstanceAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("vpc_security_group_ids"))
}

func (i dataInstanceAttributes) CreditSpecification() terra.ListValue[datainstance.CreditSpecificationAttributes] {
	return terra.ReferenceAsList[datainstance.CreditSpecificationAttributes](i.ref.Append("credit_specification"))
}

func (i dataInstanceAttributes) EbsBlockDevice() terra.SetValue[datainstance.EbsBlockDeviceAttributes] {
	return terra.ReferenceAsSet[datainstance.EbsBlockDeviceAttributes](i.ref.Append("ebs_block_device"))
}

func (i dataInstanceAttributes) EnclaveOptions() terra.ListValue[datainstance.EnclaveOptionsAttributes] {
	return terra.ReferenceAsList[datainstance.EnclaveOptionsAttributes](i.ref.Append("enclave_options"))
}

func (i dataInstanceAttributes) EphemeralBlockDevice() terra.ListValue[datainstance.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceAsList[datainstance.EphemeralBlockDeviceAttributes](i.ref.Append("ephemeral_block_device"))
}

func (i dataInstanceAttributes) MaintenanceOptions() terra.ListValue[datainstance.MaintenanceOptionsAttributes] {
	return terra.ReferenceAsList[datainstance.MaintenanceOptionsAttributes](i.ref.Append("maintenance_options"))
}

func (i dataInstanceAttributes) MetadataOptions() terra.ListValue[datainstance.MetadataOptionsAttributes] {
	return terra.ReferenceAsList[datainstance.MetadataOptionsAttributes](i.ref.Append("metadata_options"))
}

func (i dataInstanceAttributes) PrivateDnsNameOptions() terra.ListValue[datainstance.PrivateDnsNameOptionsAttributes] {
	return terra.ReferenceAsList[datainstance.PrivateDnsNameOptionsAttributes](i.ref.Append("private_dns_name_options"))
}

func (i dataInstanceAttributes) RootBlockDevice() terra.SetValue[datainstance.RootBlockDeviceAttributes] {
	return terra.ReferenceAsSet[datainstance.RootBlockDeviceAttributes](i.ref.Append("root_block_device"))
}

func (i dataInstanceAttributes) Filter() terra.SetValue[datainstance.FilterAttributes] {
	return terra.ReferenceAsSet[datainstance.FilterAttributes](i.ref.Append("filter"))
}

func (i dataInstanceAttributes) Timeouts() datainstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datainstance.TimeoutsAttributes](i.ref.Append("timeouts"))
}
