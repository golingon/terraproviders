// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ec2_instance_type

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_ec2_instance_type.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aeit *DataSource) DataSource() string {
	return "aws_ec2_instance_type"
}

// LocalName returns the local name for [DataSource].
func (aeit *DataSource) LocalName() string {
	return aeit.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aeit *DataSource) Configuration() interface{} {
	return aeit.Args
}

// Attributes returns the attributes for [DataSource].
func (aeit *DataSource) Attributes() dataAwsEc2InstanceTypeAttributes {
	return dataAwsEc2InstanceTypeAttributes{ref: terra.ReferenceDataSource(aeit)}
}

// DataArgs contains the configurations for aws_ec2_instance_type.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAwsEc2InstanceTypeAttributes struct {
	ref terra.Reference
}

// AutoRecoverySupported returns a reference to field auto_recovery_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) AutoRecoverySupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("auto_recovery_supported"))
}

// BareMetal returns a reference to field bare_metal of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) BareMetal() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("bare_metal"))
}

// BurstablePerformanceSupported returns a reference to field burstable_performance_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) BurstablePerformanceSupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("burstable_performance_supported"))
}

// CurrentGeneration returns a reference to field current_generation of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) CurrentGeneration() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("current_generation"))
}

// DedicatedHostsSupported returns a reference to field dedicated_hosts_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) DedicatedHostsSupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("dedicated_hosts_supported"))
}

// DefaultCores returns a reference to field default_cores of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) DefaultCores() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("default_cores"))
}

// DefaultThreadsPerCore returns a reference to field default_threads_per_core of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) DefaultThreadsPerCore() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("default_threads_per_core"))
}

// DefaultVcpus returns a reference to field default_vcpus of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) DefaultVcpus() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("default_vcpus"))
}

// EbsEncryptionSupport returns a reference to field ebs_encryption_support of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsEncryptionSupport() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("ebs_encryption_support"))
}

// EbsNvmeSupport returns a reference to field ebs_nvme_support of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsNvmeSupport() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("ebs_nvme_support"))
}

// EbsOptimizedSupport returns a reference to field ebs_optimized_support of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsOptimizedSupport() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("ebs_optimized_support"))
}

// EbsPerformanceBaselineBandwidth returns a reference to field ebs_performance_baseline_bandwidth of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsPerformanceBaselineBandwidth() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("ebs_performance_baseline_bandwidth"))
}

// EbsPerformanceBaselineIops returns a reference to field ebs_performance_baseline_iops of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsPerformanceBaselineIops() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("ebs_performance_baseline_iops"))
}

// EbsPerformanceBaselineThroughput returns a reference to field ebs_performance_baseline_throughput of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsPerformanceBaselineThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("ebs_performance_baseline_throughput"))
}

// EbsPerformanceMaximumBandwidth returns a reference to field ebs_performance_maximum_bandwidth of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsPerformanceMaximumBandwidth() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("ebs_performance_maximum_bandwidth"))
}

// EbsPerformanceMaximumIops returns a reference to field ebs_performance_maximum_iops of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsPerformanceMaximumIops() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("ebs_performance_maximum_iops"))
}

// EbsPerformanceMaximumThroughput returns a reference to field ebs_performance_maximum_throughput of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EbsPerformanceMaximumThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("ebs_performance_maximum_throughput"))
}

// EfaSupported returns a reference to field efa_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EfaSupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("efa_supported"))
}

// EnaSupport returns a reference to field ena_support of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EnaSupport() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("ena_support"))
}

// EncryptionInTransitSupported returns a reference to field encryption_in_transit_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) EncryptionInTransitSupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("encryption_in_transit_supported"))
}

// FreeTierEligible returns a reference to field free_tier_eligible of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) FreeTierEligible() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("free_tier_eligible"))
}

// HibernationSupported returns a reference to field hibernation_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) HibernationSupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("hibernation_supported"))
}

// Hypervisor returns a reference to field hypervisor of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) Hypervisor() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("hypervisor"))
}

// Id returns a reference to field id of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("id"))
}

// InstanceStorageSupported returns a reference to field instance_storage_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) InstanceStorageSupported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("instance_storage_supported"))
}

// InstanceType returns a reference to field instance_type of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("instance_type"))
}

// Ipv6Supported returns a reference to field ipv6_supported of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) Ipv6Supported() terra.BoolValue {
	return terra.ReferenceAsBool(aeit.ref.Append("ipv6_supported"))
}

// MaximumIpv4AddressesPerInterface returns a reference to field maximum_ipv4_addresses_per_interface of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) MaximumIpv4AddressesPerInterface() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("maximum_ipv4_addresses_per_interface"))
}

// MaximumIpv6AddressesPerInterface returns a reference to field maximum_ipv6_addresses_per_interface of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) MaximumIpv6AddressesPerInterface() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("maximum_ipv6_addresses_per_interface"))
}

// MaximumNetworkCards returns a reference to field maximum_network_cards of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) MaximumNetworkCards() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("maximum_network_cards"))
}

// MaximumNetworkInterfaces returns a reference to field maximum_network_interfaces of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) MaximumNetworkInterfaces() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("maximum_network_interfaces"))
}

// MemorySize returns a reference to field memory_size of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) MemorySize() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("memory_size"))
}

// NetworkPerformance returns a reference to field network_performance of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) NetworkPerformance() terra.StringValue {
	return terra.ReferenceAsString(aeit.ref.Append("network_performance"))
}

// SupportedArchitectures returns a reference to field supported_architectures of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) SupportedArchitectures() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aeit.ref.Append("supported_architectures"))
}

// SupportedPlacementStrategies returns a reference to field supported_placement_strategies of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) SupportedPlacementStrategies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aeit.ref.Append("supported_placement_strategies"))
}

// SupportedRootDeviceTypes returns a reference to field supported_root_device_types of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) SupportedRootDeviceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aeit.ref.Append("supported_root_device_types"))
}

// SupportedUsagesClasses returns a reference to field supported_usages_classes of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) SupportedUsagesClasses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aeit.ref.Append("supported_usages_classes"))
}

// SupportedVirtualizationTypes returns a reference to field supported_virtualization_types of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) SupportedVirtualizationTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aeit.ref.Append("supported_virtualization_types"))
}

// SustainedClockSpeed returns a reference to field sustained_clock_speed of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) SustainedClockSpeed() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("sustained_clock_speed"))
}

// TotalFpgaMemory returns a reference to field total_fpga_memory of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) TotalFpgaMemory() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("total_fpga_memory"))
}

// TotalGpuMemory returns a reference to field total_gpu_memory of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) TotalGpuMemory() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("total_gpu_memory"))
}

// TotalInstanceStorage returns a reference to field total_instance_storage of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) TotalInstanceStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(aeit.ref.Append("total_instance_storage"))
}

// ValidCores returns a reference to field valid_cores of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) ValidCores() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](aeit.ref.Append("valid_cores"))
}

// ValidThreadsPerCore returns a reference to field valid_threads_per_core of aws_ec2_instance_type.
func (aeit dataAwsEc2InstanceTypeAttributes) ValidThreadsPerCore() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](aeit.ref.Append("valid_threads_per_core"))
}

func (aeit dataAwsEc2InstanceTypeAttributes) Fpgas() terra.SetValue[DataFpgasAttributes] {
	return terra.ReferenceAsSet[DataFpgasAttributes](aeit.ref.Append("fpgas"))
}

func (aeit dataAwsEc2InstanceTypeAttributes) Gpus() terra.SetValue[DataGpusAttributes] {
	return terra.ReferenceAsSet[DataGpusAttributes](aeit.ref.Append("gpus"))
}

func (aeit dataAwsEc2InstanceTypeAttributes) InferenceAccelerators() terra.SetValue[DataInferenceAcceleratorsAttributes] {
	return terra.ReferenceAsSet[DataInferenceAcceleratorsAttributes](aeit.ref.Append("inference_accelerators"))
}

func (aeit dataAwsEc2InstanceTypeAttributes) InstanceDisks() terra.SetValue[DataInstanceDisksAttributes] {
	return terra.ReferenceAsSet[DataInstanceDisksAttributes](aeit.ref.Append("instance_disks"))
}

func (aeit dataAwsEc2InstanceTypeAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aeit.ref.Append("timeouts"))
}
