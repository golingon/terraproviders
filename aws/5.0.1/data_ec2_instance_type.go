// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2instancetype "github.com/golingon/terraproviders/aws/5.0.1/dataec2instancetype"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2InstanceType creates a new instance of [DataEc2InstanceType].
func NewDataEc2InstanceType(name string, args DataEc2InstanceTypeArgs) *DataEc2InstanceType {
	return &DataEc2InstanceType{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2InstanceType)(nil)

// DataEc2InstanceType represents the Terraform data resource aws_ec2_instance_type.
type DataEc2InstanceType struct {
	Name string
	Args DataEc2InstanceTypeArgs
}

// DataSource returns the Terraform object type for [DataEc2InstanceType].
func (eit *DataEc2InstanceType) DataSource() string {
	return "aws_ec2_instance_type"
}

// LocalName returns the local name for [DataEc2InstanceType].
func (eit *DataEc2InstanceType) LocalName() string {
	return eit.Name
}

// Configuration returns the configuration (args) for [DataEc2InstanceType].
func (eit *DataEc2InstanceType) Configuration() interface{} {
	return eit.Args
}

// Attributes returns the attributes for [DataEc2InstanceType].
func (eit *DataEc2InstanceType) Attributes() dataEc2InstanceTypeAttributes {
	return dataEc2InstanceTypeAttributes{ref: terra.ReferenceDataResource(eit)}
}

// DataEc2InstanceTypeArgs contains the configurations for aws_ec2_instance_type.
type DataEc2InstanceTypeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// Fpgas: min=0
	Fpgas []dataec2instancetype.Fpgas `hcl:"fpgas,block" validate:"min=0"`
	// Gpus: min=0
	Gpus []dataec2instancetype.Gpus `hcl:"gpus,block" validate:"min=0"`
	// InferenceAccelerators: min=0
	InferenceAccelerators []dataec2instancetype.InferenceAccelerators `hcl:"inference_accelerators,block" validate:"min=0"`
	// InstanceDisks: min=0
	InstanceDisks []dataec2instancetype.InstanceDisks `hcl:"instance_disks,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2instancetype.Timeouts `hcl:"timeouts,block"`
}
type dataEc2InstanceTypeAttributes struct {
	ref terra.Reference
}

// AutoRecoverySupported returns a reference to field auto_recovery_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) AutoRecoverySupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("auto_recovery_supported"))
}

// BareMetal returns a reference to field bare_metal of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) BareMetal() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("bare_metal"))
}

// BurstablePerformanceSupported returns a reference to field burstable_performance_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) BurstablePerformanceSupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("burstable_performance_supported"))
}

// CurrentGeneration returns a reference to field current_generation of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) CurrentGeneration() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("current_generation"))
}

// DedicatedHostsSupported returns a reference to field dedicated_hosts_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) DedicatedHostsSupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("dedicated_hosts_supported"))
}

// DefaultCores returns a reference to field default_cores of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) DefaultCores() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("default_cores"))
}

// DefaultThreadsPerCore returns a reference to field default_threads_per_core of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) DefaultThreadsPerCore() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("default_threads_per_core"))
}

// DefaultVcpus returns a reference to field default_vcpus of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) DefaultVcpus() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("default_vcpus"))
}

// EbsEncryptionSupport returns a reference to field ebs_encryption_support of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsEncryptionSupport() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("ebs_encryption_support"))
}

// EbsNvmeSupport returns a reference to field ebs_nvme_support of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsNvmeSupport() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("ebs_nvme_support"))
}

// EbsOptimizedSupport returns a reference to field ebs_optimized_support of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsOptimizedSupport() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("ebs_optimized_support"))
}

// EbsPerformanceBaselineBandwidth returns a reference to field ebs_performance_baseline_bandwidth of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsPerformanceBaselineBandwidth() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("ebs_performance_baseline_bandwidth"))
}

// EbsPerformanceBaselineIops returns a reference to field ebs_performance_baseline_iops of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsPerformanceBaselineIops() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("ebs_performance_baseline_iops"))
}

// EbsPerformanceBaselineThroughput returns a reference to field ebs_performance_baseline_throughput of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsPerformanceBaselineThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("ebs_performance_baseline_throughput"))
}

// EbsPerformanceMaximumBandwidth returns a reference to field ebs_performance_maximum_bandwidth of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsPerformanceMaximumBandwidth() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("ebs_performance_maximum_bandwidth"))
}

// EbsPerformanceMaximumIops returns a reference to field ebs_performance_maximum_iops of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsPerformanceMaximumIops() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("ebs_performance_maximum_iops"))
}

// EbsPerformanceMaximumThroughput returns a reference to field ebs_performance_maximum_throughput of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EbsPerformanceMaximumThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("ebs_performance_maximum_throughput"))
}

// EfaSupported returns a reference to field efa_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EfaSupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("efa_supported"))
}

// EnaSupport returns a reference to field ena_support of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EnaSupport() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("ena_support"))
}

// EncryptionInTransitSupported returns a reference to field encryption_in_transit_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) EncryptionInTransitSupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("encryption_in_transit_supported"))
}

// FreeTierEligible returns a reference to field free_tier_eligible of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) FreeTierEligible() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("free_tier_eligible"))
}

// HibernationSupported returns a reference to field hibernation_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) HibernationSupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("hibernation_supported"))
}

// Hypervisor returns a reference to field hypervisor of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) Hypervisor() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("hypervisor"))
}

// Id returns a reference to field id of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("id"))
}

// InstanceStorageSupported returns a reference to field instance_storage_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) InstanceStorageSupported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("instance_storage_supported"))
}

// InstanceType returns a reference to field instance_type of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("instance_type"))
}

// Ipv6Supported returns a reference to field ipv6_supported of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) Ipv6Supported() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("ipv6_supported"))
}

// MaximumIpv4AddressesPerInterface returns a reference to field maximum_ipv4_addresses_per_interface of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) MaximumIpv4AddressesPerInterface() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("maximum_ipv4_addresses_per_interface"))
}

// MaximumIpv6AddressesPerInterface returns a reference to field maximum_ipv6_addresses_per_interface of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) MaximumIpv6AddressesPerInterface() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("maximum_ipv6_addresses_per_interface"))
}

// MaximumNetworkInterfaces returns a reference to field maximum_network_interfaces of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) MaximumNetworkInterfaces() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("maximum_network_interfaces"))
}

// MemorySize returns a reference to field memory_size of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) MemorySize() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("memory_size"))
}

// NetworkPerformance returns a reference to field network_performance of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) NetworkPerformance() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("network_performance"))
}

// SupportedArchitectures returns a reference to field supported_architectures of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) SupportedArchitectures() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eit.ref.Append("supported_architectures"))
}

// SupportedPlacementStrategies returns a reference to field supported_placement_strategies of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) SupportedPlacementStrategies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eit.ref.Append("supported_placement_strategies"))
}

// SupportedRootDeviceTypes returns a reference to field supported_root_device_types of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) SupportedRootDeviceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eit.ref.Append("supported_root_device_types"))
}

// SupportedUsagesClasses returns a reference to field supported_usages_classes of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) SupportedUsagesClasses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eit.ref.Append("supported_usages_classes"))
}

// SupportedVirtualizationTypes returns a reference to field supported_virtualization_types of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) SupportedVirtualizationTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eit.ref.Append("supported_virtualization_types"))
}

// SustainedClockSpeed returns a reference to field sustained_clock_speed of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) SustainedClockSpeed() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("sustained_clock_speed"))
}

// TotalFpgaMemory returns a reference to field total_fpga_memory of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) TotalFpgaMemory() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("total_fpga_memory"))
}

// TotalGpuMemory returns a reference to field total_gpu_memory of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) TotalGpuMemory() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("total_gpu_memory"))
}

// TotalInstanceStorage returns a reference to field total_instance_storage of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) TotalInstanceStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(eit.ref.Append("total_instance_storage"))
}

// ValidCores returns a reference to field valid_cores of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) ValidCores() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](eit.ref.Append("valid_cores"))
}

// ValidThreadsPerCore returns a reference to field valid_threads_per_core of aws_ec2_instance_type.
func (eit dataEc2InstanceTypeAttributes) ValidThreadsPerCore() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](eit.ref.Append("valid_threads_per_core"))
}

func (eit dataEc2InstanceTypeAttributes) Fpgas() terra.SetValue[dataec2instancetype.FpgasAttributes] {
	return terra.ReferenceAsSet[dataec2instancetype.FpgasAttributes](eit.ref.Append("fpgas"))
}

func (eit dataEc2InstanceTypeAttributes) Gpus() terra.SetValue[dataec2instancetype.GpusAttributes] {
	return terra.ReferenceAsSet[dataec2instancetype.GpusAttributes](eit.ref.Append("gpus"))
}

func (eit dataEc2InstanceTypeAttributes) InferenceAccelerators() terra.SetValue[dataec2instancetype.InferenceAcceleratorsAttributes] {
	return terra.ReferenceAsSet[dataec2instancetype.InferenceAcceleratorsAttributes](eit.ref.Append("inference_accelerators"))
}

func (eit dataEc2InstanceTypeAttributes) InstanceDisks() terra.SetValue[dataec2instancetype.InstanceDisksAttributes] {
	return terra.ReferenceAsSet[dataec2instancetype.InstanceDisksAttributes](eit.ref.Append("instance_disks"))
}

func (eit dataEc2InstanceTypeAttributes) Timeouts() dataec2instancetype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2instancetype.TimeoutsAttributes](eit.ref.Append("timeouts"))
}
