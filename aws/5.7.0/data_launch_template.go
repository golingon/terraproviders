// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalaunchtemplate "github.com/golingon/terraproviders/aws/5.7.0/datalaunchtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLaunchTemplate creates a new instance of [DataLaunchTemplate].
func NewDataLaunchTemplate(name string, args DataLaunchTemplateArgs) *DataLaunchTemplate {
	return &DataLaunchTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLaunchTemplate)(nil)

// DataLaunchTemplate represents the Terraform data resource aws_launch_template.
type DataLaunchTemplate struct {
	Name string
	Args DataLaunchTemplateArgs
}

// DataSource returns the Terraform object type for [DataLaunchTemplate].
func (lt *DataLaunchTemplate) DataSource() string {
	return "aws_launch_template"
}

// LocalName returns the local name for [DataLaunchTemplate].
func (lt *DataLaunchTemplate) LocalName() string {
	return lt.Name
}

// Configuration returns the configuration (args) for [DataLaunchTemplate].
func (lt *DataLaunchTemplate) Configuration() interface{} {
	return lt.Args
}

// Attributes returns the attributes for [DataLaunchTemplate].
func (lt *DataLaunchTemplate) Attributes() dataLaunchTemplateAttributes {
	return dataLaunchTemplateAttributes{ref: terra.ReferenceDataResource(lt)}
}

// DataLaunchTemplateArgs contains the configurations for aws_launch_template.
type DataLaunchTemplateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// BlockDeviceMappings: min=0
	BlockDeviceMappings []datalaunchtemplate.BlockDeviceMappings `hcl:"block_device_mappings,block" validate:"min=0"`
	// CapacityReservationSpecification: min=0
	CapacityReservationSpecification []datalaunchtemplate.CapacityReservationSpecification `hcl:"capacity_reservation_specification,block" validate:"min=0"`
	// CpuOptions: min=0
	CpuOptions []datalaunchtemplate.CpuOptions `hcl:"cpu_options,block" validate:"min=0"`
	// CreditSpecification: min=0
	CreditSpecification []datalaunchtemplate.CreditSpecification `hcl:"credit_specification,block" validate:"min=0"`
	// ElasticGpuSpecifications: min=0
	ElasticGpuSpecifications []datalaunchtemplate.ElasticGpuSpecifications `hcl:"elastic_gpu_specifications,block" validate:"min=0"`
	// ElasticInferenceAccelerator: min=0
	ElasticInferenceAccelerator []datalaunchtemplate.ElasticInferenceAccelerator `hcl:"elastic_inference_accelerator,block" validate:"min=0"`
	// EnclaveOptions: min=0
	EnclaveOptions []datalaunchtemplate.EnclaveOptions `hcl:"enclave_options,block" validate:"min=0"`
	// HibernationOptions: min=0
	HibernationOptions []datalaunchtemplate.HibernationOptions `hcl:"hibernation_options,block" validate:"min=0"`
	// IamInstanceProfile: min=0
	IamInstanceProfile []datalaunchtemplate.IamInstanceProfile `hcl:"iam_instance_profile,block" validate:"min=0"`
	// InstanceMarketOptions: min=0
	InstanceMarketOptions []datalaunchtemplate.InstanceMarketOptions `hcl:"instance_market_options,block" validate:"min=0"`
	// InstanceRequirements: min=0
	InstanceRequirements []datalaunchtemplate.InstanceRequirements `hcl:"instance_requirements,block" validate:"min=0"`
	// LicenseSpecification: min=0
	LicenseSpecification []datalaunchtemplate.LicenseSpecification `hcl:"license_specification,block" validate:"min=0"`
	// MaintenanceOptions: min=0
	MaintenanceOptions []datalaunchtemplate.MaintenanceOptions `hcl:"maintenance_options,block" validate:"min=0"`
	// MetadataOptions: min=0
	MetadataOptions []datalaunchtemplate.MetadataOptions `hcl:"metadata_options,block" validate:"min=0"`
	// Monitoring: min=0
	Monitoring []datalaunchtemplate.Monitoring `hcl:"monitoring,block" validate:"min=0"`
	// NetworkInterfaces: min=0
	NetworkInterfaces []datalaunchtemplate.NetworkInterfaces `hcl:"network_interfaces,block" validate:"min=0"`
	// Placement: min=0
	Placement []datalaunchtemplate.Placement `hcl:"placement,block" validate:"min=0"`
	// PrivateDnsNameOptions: min=0
	PrivateDnsNameOptions []datalaunchtemplate.PrivateDnsNameOptions `hcl:"private_dns_name_options,block" validate:"min=0"`
	// TagSpecifications: min=0
	TagSpecifications []datalaunchtemplate.TagSpecifications `hcl:"tag_specifications,block" validate:"min=0"`
	// Filter: min=0
	Filter []datalaunchtemplate.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalaunchtemplate.Timeouts `hcl:"timeouts,block"`
}
type dataLaunchTemplateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_launch_template.
func (lt dataLaunchTemplateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("arn"))
}

// DefaultVersion returns a reference to field default_version of aws_launch_template.
func (lt dataLaunchTemplateAttributes) DefaultVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(lt.ref.Append("default_version"))
}

// Description returns a reference to field description of aws_launch_template.
func (lt dataLaunchTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("description"))
}

// DisableApiStop returns a reference to field disable_api_stop of aws_launch_template.
func (lt dataLaunchTemplateAttributes) DisableApiStop() terra.BoolValue {
	return terra.ReferenceAsBool(lt.ref.Append("disable_api_stop"))
}

// DisableApiTermination returns a reference to field disable_api_termination of aws_launch_template.
func (lt dataLaunchTemplateAttributes) DisableApiTermination() terra.BoolValue {
	return terra.ReferenceAsBool(lt.ref.Append("disable_api_termination"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_launch_template.
func (lt dataLaunchTemplateAttributes) EbsOptimized() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("ebs_optimized"))
}

// Id returns a reference to field id of aws_launch_template.
func (lt dataLaunchTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("id"))
}

// ImageId returns a reference to field image_id of aws_launch_template.
func (lt dataLaunchTemplateAttributes) ImageId() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("image_id"))
}

// InstanceInitiatedShutdownBehavior returns a reference to field instance_initiated_shutdown_behavior of aws_launch_template.
func (lt dataLaunchTemplateAttributes) InstanceInitiatedShutdownBehavior() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("instance_initiated_shutdown_behavior"))
}

// InstanceType returns a reference to field instance_type of aws_launch_template.
func (lt dataLaunchTemplateAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("instance_type"))
}

// KernelId returns a reference to field kernel_id of aws_launch_template.
func (lt dataLaunchTemplateAttributes) KernelId() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("kernel_id"))
}

// KeyName returns a reference to field key_name of aws_launch_template.
func (lt dataLaunchTemplateAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("key_name"))
}

// LatestVersion returns a reference to field latest_version of aws_launch_template.
func (lt dataLaunchTemplateAttributes) LatestVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(lt.ref.Append("latest_version"))
}

// Name returns a reference to field name of aws_launch_template.
func (lt dataLaunchTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("name"))
}

// RamDiskId returns a reference to field ram_disk_id of aws_launch_template.
func (lt dataLaunchTemplateAttributes) RamDiskId() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("ram_disk_id"))
}

// SecurityGroupNames returns a reference to field security_group_names of aws_launch_template.
func (lt dataLaunchTemplateAttributes) SecurityGroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lt.ref.Append("security_group_names"))
}

// Tags returns a reference to field tags of aws_launch_template.
func (lt dataLaunchTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lt.ref.Append("tags"))
}

// UserData returns a reference to field user_data of aws_launch_template.
func (lt dataLaunchTemplateAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("user_data"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_launch_template.
func (lt dataLaunchTemplateAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lt.ref.Append("vpc_security_group_ids"))
}

func (lt dataLaunchTemplateAttributes) BlockDeviceMappings() terra.ListValue[datalaunchtemplate.BlockDeviceMappingsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.BlockDeviceMappingsAttributes](lt.ref.Append("block_device_mappings"))
}

func (lt dataLaunchTemplateAttributes) CapacityReservationSpecification() terra.ListValue[datalaunchtemplate.CapacityReservationSpecificationAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.CapacityReservationSpecificationAttributes](lt.ref.Append("capacity_reservation_specification"))
}

func (lt dataLaunchTemplateAttributes) CpuOptions() terra.ListValue[datalaunchtemplate.CpuOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.CpuOptionsAttributes](lt.ref.Append("cpu_options"))
}

func (lt dataLaunchTemplateAttributes) CreditSpecification() terra.ListValue[datalaunchtemplate.CreditSpecificationAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.CreditSpecificationAttributes](lt.ref.Append("credit_specification"))
}

func (lt dataLaunchTemplateAttributes) ElasticGpuSpecifications() terra.ListValue[datalaunchtemplate.ElasticGpuSpecificationsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.ElasticGpuSpecificationsAttributes](lt.ref.Append("elastic_gpu_specifications"))
}

func (lt dataLaunchTemplateAttributes) ElasticInferenceAccelerator() terra.ListValue[datalaunchtemplate.ElasticInferenceAcceleratorAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.ElasticInferenceAcceleratorAttributes](lt.ref.Append("elastic_inference_accelerator"))
}

func (lt dataLaunchTemplateAttributes) EnclaveOptions() terra.ListValue[datalaunchtemplate.EnclaveOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.EnclaveOptionsAttributes](lt.ref.Append("enclave_options"))
}

func (lt dataLaunchTemplateAttributes) HibernationOptions() terra.ListValue[datalaunchtemplate.HibernationOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.HibernationOptionsAttributes](lt.ref.Append("hibernation_options"))
}

func (lt dataLaunchTemplateAttributes) IamInstanceProfile() terra.ListValue[datalaunchtemplate.IamInstanceProfileAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.IamInstanceProfileAttributes](lt.ref.Append("iam_instance_profile"))
}

func (lt dataLaunchTemplateAttributes) InstanceMarketOptions() terra.ListValue[datalaunchtemplate.InstanceMarketOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.InstanceMarketOptionsAttributes](lt.ref.Append("instance_market_options"))
}

func (lt dataLaunchTemplateAttributes) InstanceRequirements() terra.ListValue[datalaunchtemplate.InstanceRequirementsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.InstanceRequirementsAttributes](lt.ref.Append("instance_requirements"))
}

func (lt dataLaunchTemplateAttributes) LicenseSpecification() terra.ListValue[datalaunchtemplate.LicenseSpecificationAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.LicenseSpecificationAttributes](lt.ref.Append("license_specification"))
}

func (lt dataLaunchTemplateAttributes) MaintenanceOptions() terra.ListValue[datalaunchtemplate.MaintenanceOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.MaintenanceOptionsAttributes](lt.ref.Append("maintenance_options"))
}

func (lt dataLaunchTemplateAttributes) MetadataOptions() terra.ListValue[datalaunchtemplate.MetadataOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.MetadataOptionsAttributes](lt.ref.Append("metadata_options"))
}

func (lt dataLaunchTemplateAttributes) Monitoring() terra.ListValue[datalaunchtemplate.MonitoringAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.MonitoringAttributes](lt.ref.Append("monitoring"))
}

func (lt dataLaunchTemplateAttributes) NetworkInterfaces() terra.ListValue[datalaunchtemplate.NetworkInterfacesAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.NetworkInterfacesAttributes](lt.ref.Append("network_interfaces"))
}

func (lt dataLaunchTemplateAttributes) Placement() terra.ListValue[datalaunchtemplate.PlacementAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.PlacementAttributes](lt.ref.Append("placement"))
}

func (lt dataLaunchTemplateAttributes) PrivateDnsNameOptions() terra.ListValue[datalaunchtemplate.PrivateDnsNameOptionsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.PrivateDnsNameOptionsAttributes](lt.ref.Append("private_dns_name_options"))
}

func (lt dataLaunchTemplateAttributes) TagSpecifications() terra.ListValue[datalaunchtemplate.TagSpecificationsAttributes] {
	return terra.ReferenceAsList[datalaunchtemplate.TagSpecificationsAttributes](lt.ref.Append("tag_specifications"))
}

func (lt dataLaunchTemplateAttributes) Filter() terra.SetValue[datalaunchtemplate.FilterAttributes] {
	return terra.ReferenceAsSet[datalaunchtemplate.FilterAttributes](lt.ref.Append("filter"))
}

func (lt dataLaunchTemplateAttributes) Timeouts() datalaunchtemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalaunchtemplate.TimeoutsAttributes](lt.ref.Append("timeouts"))
}
