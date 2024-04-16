// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_instance_from_machine_image

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type GuestAccelerator struct {
	// Count: number, optional
	Count terra.NumberValue `hcl:"count,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type ServiceAccount struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Scopes: set of string, optional
	Scopes terra.SetValue[terra.StringValue] `hcl:"scopes,attr"`
}

type AdvancedMachineFeatures struct {
	// EnableNestedVirtualization: bool, optional
	EnableNestedVirtualization terra.BoolValue `hcl:"enable_nested_virtualization,attr"`
	// ThreadsPerCore: number, optional
	ThreadsPerCore terra.NumberValue `hcl:"threads_per_core,attr"`
	// VisibleCoreCount: number, optional
	VisibleCoreCount terra.NumberValue `hcl:"visible_core_count,attr"`
}

type ConfidentialInstanceConfig struct {
	// ConfidentialInstanceType: string, optional
	ConfidentialInstanceType terra.StringValue `hcl:"confidential_instance_type,attr"`
	// EnableConfidentialCompute: bool, optional
	EnableConfidentialCompute terra.BoolValue `hcl:"enable_confidential_compute,attr"`
}

type NetworkInterface struct {
	// InternalIpv6PrefixLength: number, optional
	InternalIpv6PrefixLength terra.NumberValue `hcl:"internal_ipv6_prefix_length,attr"`
	// Ipv6Address: string, optional
	Ipv6Address terra.StringValue `hcl:"ipv6_address,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkAttachment: string, optional
	NetworkAttachment terra.StringValue `hcl:"network_attachment,attr"`
	// NetworkIp: string, optional
	NetworkIp terra.StringValue `hcl:"network_ip,attr"`
	// NicType: string, optional
	NicType terra.StringValue `hcl:"nic_type,attr"`
	// QueueCount: number, optional
	QueueCount terra.NumberValue `hcl:"queue_count,attr"`
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
	// StackType: string, optional
	StackType terra.StringValue `hcl:"stack_type,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// SubnetworkProject: string, optional
	SubnetworkProject terra.StringValue `hcl:"subnetwork_project,attr"`
	// NetworkInterfaceAccessConfig: min=0
	AccessConfig []NetworkInterfaceAccessConfig `hcl:"access_config,block" validate:"min=0"`
	// NetworkInterfaceAliasIpRange: min=0
	AliasIpRange []NetworkInterfaceAliasIpRange `hcl:"alias_ip_range,block" validate:"min=0"`
	// NetworkInterfaceIpv6AccessConfig: min=0
	Ipv6AccessConfig []NetworkInterfaceIpv6AccessConfig `hcl:"ipv6_access_config,block" validate:"min=0"`
}

type NetworkInterfaceAccessConfig struct {
	// NatIp: string, optional
	NatIp terra.StringValue `hcl:"nat_ip,attr"`
	// NetworkTier: string, optional
	NetworkTier terra.StringValue `hcl:"network_tier,attr"`
	// PublicPtrDomainName: string, optional
	PublicPtrDomainName terra.StringValue `hcl:"public_ptr_domain_name,attr"`
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
}

type NetworkInterfaceAliasIpRange struct {
	// IpCidrRange: string, optional
	IpCidrRange terra.StringValue `hcl:"ip_cidr_range,attr"`
	// SubnetworkRangeName: string, optional
	SubnetworkRangeName terra.StringValue `hcl:"subnetwork_range_name,attr"`
}

type NetworkInterfaceIpv6AccessConfig struct {
	// ExternalIpv6: string, optional
	ExternalIpv6 terra.StringValue `hcl:"external_ipv6,attr"`
	// ExternalIpv6PrefixLength: string, optional
	ExternalIpv6PrefixLength terra.StringValue `hcl:"external_ipv6_prefix_length,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NetworkTier: string, required
	NetworkTier terra.StringValue `hcl:"network_tier,attr" validate:"required"`
	// PublicPtrDomainName: string, optional
	PublicPtrDomainName terra.StringValue `hcl:"public_ptr_domain_name,attr"`
}

type NetworkPerformanceConfig struct {
	// TotalEgressBandwidthTier: string, required
	TotalEgressBandwidthTier terra.StringValue `hcl:"total_egress_bandwidth_tier,attr" validate:"required"`
}

type Params struct {
	// ResourceManagerTags: map of string, optional
	ResourceManagerTags terra.MapValue[terra.StringValue] `hcl:"resource_manager_tags,attr"`
}

type ReservationAffinity struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ReservationAffinitySpecificReservation: optional
	SpecificReservation *ReservationAffinitySpecificReservation `hcl:"specific_reservation,block"`
}

type ReservationAffinitySpecificReservation struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Scheduling struct {
	// AutomaticRestart: bool, optional
	AutomaticRestart terra.BoolValue `hcl:"automatic_restart,attr"`
	// InstanceTerminationAction: string, optional
	InstanceTerminationAction terra.StringValue `hcl:"instance_termination_action,attr"`
	// MaintenanceInterval: string, optional
	MaintenanceInterval terra.StringValue `hcl:"maintenance_interval,attr"`
	// MinNodeCpus: number, optional
	MinNodeCpus terra.NumberValue `hcl:"min_node_cpus,attr"`
	// OnHostMaintenance: string, optional
	OnHostMaintenance terra.StringValue `hcl:"on_host_maintenance,attr"`
	// Preemptible: bool, optional
	Preemptible terra.BoolValue `hcl:"preemptible,attr"`
	// ProvisioningModel: string, optional
	ProvisioningModel terra.StringValue `hcl:"provisioning_model,attr"`
	// SchedulingLocalSsdRecoveryTimeout: optional
	LocalSsdRecoveryTimeout *SchedulingLocalSsdRecoveryTimeout `hcl:"local_ssd_recovery_timeout,block"`
	// SchedulingMaxRunDuration: optional
	MaxRunDuration *SchedulingMaxRunDuration `hcl:"max_run_duration,block"`
	// SchedulingNodeAffinities: min=0
	NodeAffinities []SchedulingNodeAffinities `hcl:"node_affinities,block" validate:"min=0"`
}

type SchedulingLocalSsdRecoveryTimeout struct {
	// Nanos: number, optional
	Nanos terra.NumberValue `hcl:"nanos,attr"`
	// Seconds: number, required
	Seconds terra.NumberValue `hcl:"seconds,attr" validate:"required"`
}

type SchedulingMaxRunDuration struct {
	// Nanos: number, optional
	Nanos terra.NumberValue `hcl:"nanos,attr"`
	// Seconds: number, required
	Seconds terra.NumberValue `hcl:"seconds,attr" validate:"required"`
}

type SchedulingNodeAffinities struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ShieldedInstanceConfig struct {
	// EnableIntegrityMonitoring: bool, optional
	EnableIntegrityMonitoring terra.BoolValue `hcl:"enable_integrity_monitoring,attr"`
	// EnableSecureBoot: bool, optional
	EnableSecureBoot terra.BoolValue `hcl:"enable_secure_boot,attr"`
	// EnableVtpm: bool, optional
	EnableVtpm terra.BoolValue `hcl:"enable_vtpm,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AttachedDiskAttributes struct {
	ref terra.Reference
}

func (ad AttachedDiskAttributes) InternalRef() (terra.Reference, error) {
	return ad.ref, nil
}

func (ad AttachedDiskAttributes) InternalWithRef(ref terra.Reference) AttachedDiskAttributes {
	return AttachedDiskAttributes{ref: ref}
}

func (ad AttachedDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ad.ref.InternalTokens()
}

func (ad AttachedDiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("device_name"))
}

func (ad AttachedDiskAttributes) DiskEncryptionKeyRaw() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("disk_encryption_key_raw"))
}

func (ad AttachedDiskAttributes) DiskEncryptionKeySha256() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("disk_encryption_key_sha256"))
}

func (ad AttachedDiskAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("kms_key_self_link"))
}

func (ad AttachedDiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("mode"))
}

func (ad AttachedDiskAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("source"))
}

type BootDiskAttributes struct {
	ref terra.Reference
}

func (bd BootDiskAttributes) InternalRef() (terra.Reference, error) {
	return bd.ref, nil
}

func (bd BootDiskAttributes) InternalWithRef(ref terra.Reference) BootDiskAttributes {
	return BootDiskAttributes{ref: ref}
}

func (bd BootDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bd.ref.InternalTokens()
}

func (bd BootDiskAttributes) AutoDelete() terra.BoolValue {
	return terra.ReferenceAsBool(bd.ref.Append("auto_delete"))
}

func (bd BootDiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("device_name"))
}

func (bd BootDiskAttributes) DiskEncryptionKeyRaw() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("disk_encryption_key_raw"))
}

func (bd BootDiskAttributes) DiskEncryptionKeySha256() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("disk_encryption_key_sha256"))
}

func (bd BootDiskAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("kms_key_self_link"))
}

func (bd BootDiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("mode"))
}

func (bd BootDiskAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("source"))
}

func (bd BootDiskAttributes) InitializeParams() terra.ListValue[BootDiskInitializeParamsAttributes] {
	return terra.ReferenceAsList[BootDiskInitializeParamsAttributes](bd.ref.Append("initialize_params"))
}

type BootDiskInitializeParamsAttributes struct {
	ref terra.Reference
}

func (ip BootDiskInitializeParamsAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip BootDiskInitializeParamsAttributes) InternalWithRef(ref terra.Reference) BootDiskInitializeParamsAttributes {
	return BootDiskInitializeParamsAttributes{ref: ref}
}

func (ip BootDiskInitializeParamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip BootDiskInitializeParamsAttributes) EnableConfidentialCompute() terra.BoolValue {
	return terra.ReferenceAsBool(ip.ref.Append("enable_confidential_compute"))
}

func (ip BootDiskInitializeParamsAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("image"))
}

func (ip BootDiskInitializeParamsAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ip.ref.Append("labels"))
}

func (ip BootDiskInitializeParamsAttributes) ProvisionedIops() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("provisioned_iops"))
}

func (ip BootDiskInitializeParamsAttributes) ProvisionedThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("provisioned_throughput"))
}

func (ip BootDiskInitializeParamsAttributes) ResourceManagerTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ip.ref.Append("resource_manager_tags"))
}

func (ip BootDiskInitializeParamsAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("size"))
}

func (ip BootDiskInitializeParamsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("type"))
}

type GuestAcceleratorAttributes struct {
	ref terra.Reference
}

func (ga GuestAcceleratorAttributes) InternalRef() (terra.Reference, error) {
	return ga.ref, nil
}

func (ga GuestAcceleratorAttributes) InternalWithRef(ref terra.Reference) GuestAcceleratorAttributes {
	return GuestAcceleratorAttributes{ref: ref}
}

func (ga GuestAcceleratorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ga.ref.InternalTokens()
}

func (ga GuestAcceleratorAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(ga.ref.Append("count"))
}

func (ga GuestAcceleratorAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("type"))
}

type ScratchDiskAttributes struct {
	ref terra.Reference
}

func (sd ScratchDiskAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd ScratchDiskAttributes) InternalWithRef(ref terra.Reference) ScratchDiskAttributes {
	return ScratchDiskAttributes{ref: ref}
}

func (sd ScratchDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd ScratchDiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("device_name"))
}

func (sd ScratchDiskAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("interface"))
}

func (sd ScratchDiskAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(sd.ref.Append("size"))
}

type ServiceAccountAttributes struct {
	ref terra.Reference
}

func (sa ServiceAccountAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa ServiceAccountAttributes) InternalWithRef(ref terra.Reference) ServiceAccountAttributes {
	return ServiceAccountAttributes{ref: ref}
}

func (sa ServiceAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa ServiceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("email"))
}

func (sa ServiceAccountAttributes) Scopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sa.ref.Append("scopes"))
}

type AdvancedMachineFeaturesAttributes struct {
	ref terra.Reference
}

func (amf AdvancedMachineFeaturesAttributes) InternalRef() (terra.Reference, error) {
	return amf.ref, nil
}

func (amf AdvancedMachineFeaturesAttributes) InternalWithRef(ref terra.Reference) AdvancedMachineFeaturesAttributes {
	return AdvancedMachineFeaturesAttributes{ref: ref}
}

func (amf AdvancedMachineFeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return amf.ref.InternalTokens()
}

func (amf AdvancedMachineFeaturesAttributes) EnableNestedVirtualization() terra.BoolValue {
	return terra.ReferenceAsBool(amf.ref.Append("enable_nested_virtualization"))
}

func (amf AdvancedMachineFeaturesAttributes) ThreadsPerCore() terra.NumberValue {
	return terra.ReferenceAsNumber(amf.ref.Append("threads_per_core"))
}

func (amf AdvancedMachineFeaturesAttributes) VisibleCoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(amf.ref.Append("visible_core_count"))
}

type ConfidentialInstanceConfigAttributes struct {
	ref terra.Reference
}

func (cic ConfidentialInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return cic.ref, nil
}

func (cic ConfidentialInstanceConfigAttributes) InternalWithRef(ref terra.Reference) ConfidentialInstanceConfigAttributes {
	return ConfidentialInstanceConfigAttributes{ref: ref}
}

func (cic ConfidentialInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cic.ref.InternalTokens()
}

func (cic ConfidentialInstanceConfigAttributes) ConfidentialInstanceType() terra.StringValue {
	return terra.ReferenceAsString(cic.ref.Append("confidential_instance_type"))
}

func (cic ConfidentialInstanceConfigAttributes) EnableConfidentialCompute() terra.BoolValue {
	return terra.ReferenceAsBool(cic.ref.Append("enable_confidential_compute"))
}

type NetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni NetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni NetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAttributes {
	return NetworkInterfaceAttributes{ref: ref}
}

func (ni NetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni NetworkInterfaceAttributes) InternalIpv6PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("internal_ipv6_prefix_length"))
}

func (ni NetworkInterfaceAttributes) Ipv6AccessType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("ipv6_access_type"))
}

func (ni NetworkInterfaceAttributes) Ipv6Address() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("ipv6_address"))
}

func (ni NetworkInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("name"))
}

func (ni NetworkInterfaceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network"))
}

func (ni NetworkInterfaceAttributes) NetworkAttachment() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_attachment"))
}

func (ni NetworkInterfaceAttributes) NetworkIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_ip"))
}

func (ni NetworkInterfaceAttributes) NicType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("nic_type"))
}

func (ni NetworkInterfaceAttributes) QueueCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("queue_count"))
}

func (ni NetworkInterfaceAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("security_policy"))
}

func (ni NetworkInterfaceAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("stack_type"))
}

func (ni NetworkInterfaceAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnetwork"))
}

func (ni NetworkInterfaceAttributes) SubnetworkProject() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnetwork_project"))
}

func (ni NetworkInterfaceAttributes) AccessConfig() terra.ListValue[NetworkInterfaceAccessConfigAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceAccessConfigAttributes](ni.ref.Append("access_config"))
}

func (ni NetworkInterfaceAttributes) AliasIpRange() terra.ListValue[NetworkInterfaceAliasIpRangeAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceAliasIpRangeAttributes](ni.ref.Append("alias_ip_range"))
}

func (ni NetworkInterfaceAttributes) Ipv6AccessConfig() terra.ListValue[NetworkInterfaceIpv6AccessConfigAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceIpv6AccessConfigAttributes](ni.ref.Append("ipv6_access_config"))
}

type NetworkInterfaceAccessConfigAttributes struct {
	ref terra.Reference
}

func (ac NetworkInterfaceAccessConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac NetworkInterfaceAccessConfigAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAccessConfigAttributes {
	return NetworkInterfaceAccessConfigAttributes{ref: ref}
}

func (ac NetworkInterfaceAccessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac NetworkInterfaceAccessConfigAttributes) NatIp() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("nat_ip"))
}

func (ac NetworkInterfaceAccessConfigAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("network_tier"))
}

func (ac NetworkInterfaceAccessConfigAttributes) PublicPtrDomainName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("public_ptr_domain_name"))
}

func (ac NetworkInterfaceAccessConfigAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("security_policy"))
}

type NetworkInterfaceAliasIpRangeAttributes struct {
	ref terra.Reference
}

func (air NetworkInterfaceAliasIpRangeAttributes) InternalRef() (terra.Reference, error) {
	return air.ref, nil
}

func (air NetworkInterfaceAliasIpRangeAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAliasIpRangeAttributes {
	return NetworkInterfaceAliasIpRangeAttributes{ref: ref}
}

func (air NetworkInterfaceAliasIpRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return air.ref.InternalTokens()
}

func (air NetworkInterfaceAliasIpRangeAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("ip_cidr_range"))
}

func (air NetworkInterfaceAliasIpRangeAttributes) SubnetworkRangeName() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("subnetwork_range_name"))
}

type NetworkInterfaceIpv6AccessConfigAttributes struct {
	ref terra.Reference
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) InternalRef() (terra.Reference, error) {
	return iac.ref, nil
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceIpv6AccessConfigAttributes {
	return NetworkInterfaceIpv6AccessConfigAttributes{ref: ref}
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iac.ref.InternalTokens()
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) ExternalIpv6() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("external_ipv6"))
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) ExternalIpv6PrefixLength() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("external_ipv6_prefix_length"))
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("name"))
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("network_tier"))
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) PublicPtrDomainName() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("public_ptr_domain_name"))
}

func (iac NetworkInterfaceIpv6AccessConfigAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("security_policy"))
}

type NetworkPerformanceConfigAttributes struct {
	ref terra.Reference
}

func (npc NetworkPerformanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return npc.ref, nil
}

func (npc NetworkPerformanceConfigAttributes) InternalWithRef(ref terra.Reference) NetworkPerformanceConfigAttributes {
	return NetworkPerformanceConfigAttributes{ref: ref}
}

func (npc NetworkPerformanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return npc.ref.InternalTokens()
}

func (npc NetworkPerformanceConfigAttributes) TotalEgressBandwidthTier() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("total_egress_bandwidth_tier"))
}

type ParamsAttributes struct {
	ref terra.Reference
}

func (p ParamsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParamsAttributes) InternalWithRef(ref terra.Reference) ParamsAttributes {
	return ParamsAttributes{ref: ref}
}

func (p ParamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParamsAttributes) ResourceManagerTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](p.ref.Append("resource_manager_tags"))
}

type ReservationAffinityAttributes struct {
	ref terra.Reference
}

func (ra ReservationAffinityAttributes) InternalRef() (terra.Reference, error) {
	return ra.ref, nil
}

func (ra ReservationAffinityAttributes) InternalWithRef(ref terra.Reference) ReservationAffinityAttributes {
	return ReservationAffinityAttributes{ref: ref}
}

func (ra ReservationAffinityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ra.ref.InternalTokens()
}

func (ra ReservationAffinityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("type"))
}

func (ra ReservationAffinityAttributes) SpecificReservation() terra.ListValue[ReservationAffinitySpecificReservationAttributes] {
	return terra.ReferenceAsList[ReservationAffinitySpecificReservationAttributes](ra.ref.Append("specific_reservation"))
}

type ReservationAffinitySpecificReservationAttributes struct {
	ref terra.Reference
}

func (sr ReservationAffinitySpecificReservationAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr ReservationAffinitySpecificReservationAttributes) InternalWithRef(ref terra.Reference) ReservationAffinitySpecificReservationAttributes {
	return ReservationAffinitySpecificReservationAttributes{ref: ref}
}

func (sr ReservationAffinitySpecificReservationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr ReservationAffinitySpecificReservationAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("key"))
}

func (sr ReservationAffinitySpecificReservationAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sr.ref.Append("values"))
}

type SchedulingAttributes struct {
	ref terra.Reference
}

func (s SchedulingAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SchedulingAttributes) InternalWithRef(ref terra.Reference) SchedulingAttributes {
	return SchedulingAttributes{ref: ref}
}

func (s SchedulingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SchedulingAttributes) AutomaticRestart() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("automatic_restart"))
}

func (s SchedulingAttributes) InstanceTerminationAction() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("instance_termination_action"))
}

func (s SchedulingAttributes) MaintenanceInterval() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("maintenance_interval"))
}

func (s SchedulingAttributes) MinNodeCpus() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("min_node_cpus"))
}

func (s SchedulingAttributes) OnHostMaintenance() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("on_host_maintenance"))
}

func (s SchedulingAttributes) Preemptible() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("preemptible"))
}

func (s SchedulingAttributes) ProvisioningModel() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("provisioning_model"))
}

func (s SchedulingAttributes) LocalSsdRecoveryTimeout() terra.ListValue[SchedulingLocalSsdRecoveryTimeoutAttributes] {
	return terra.ReferenceAsList[SchedulingLocalSsdRecoveryTimeoutAttributes](s.ref.Append("local_ssd_recovery_timeout"))
}

func (s SchedulingAttributes) MaxRunDuration() terra.ListValue[SchedulingMaxRunDurationAttributes] {
	return terra.ReferenceAsList[SchedulingMaxRunDurationAttributes](s.ref.Append("max_run_duration"))
}

func (s SchedulingAttributes) NodeAffinities() terra.SetValue[SchedulingNodeAffinitiesAttributes] {
	return terra.ReferenceAsSet[SchedulingNodeAffinitiesAttributes](s.ref.Append("node_affinities"))
}

type SchedulingLocalSsdRecoveryTimeoutAttributes struct {
	ref terra.Reference
}

func (lsrt SchedulingLocalSsdRecoveryTimeoutAttributes) InternalRef() (terra.Reference, error) {
	return lsrt.ref, nil
}

func (lsrt SchedulingLocalSsdRecoveryTimeoutAttributes) InternalWithRef(ref terra.Reference) SchedulingLocalSsdRecoveryTimeoutAttributes {
	return SchedulingLocalSsdRecoveryTimeoutAttributes{ref: ref}
}

func (lsrt SchedulingLocalSsdRecoveryTimeoutAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lsrt.ref.InternalTokens()
}

func (lsrt SchedulingLocalSsdRecoveryTimeoutAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(lsrt.ref.Append("nanos"))
}

func (lsrt SchedulingLocalSsdRecoveryTimeoutAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lsrt.ref.Append("seconds"))
}

type SchedulingMaxRunDurationAttributes struct {
	ref terra.Reference
}

func (mrd SchedulingMaxRunDurationAttributes) InternalRef() (terra.Reference, error) {
	return mrd.ref, nil
}

func (mrd SchedulingMaxRunDurationAttributes) InternalWithRef(ref terra.Reference) SchedulingMaxRunDurationAttributes {
	return SchedulingMaxRunDurationAttributes{ref: ref}
}

func (mrd SchedulingMaxRunDurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mrd.ref.InternalTokens()
}

func (mrd SchedulingMaxRunDurationAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(mrd.ref.Append("nanos"))
}

func (mrd SchedulingMaxRunDurationAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mrd.ref.Append("seconds"))
}

type SchedulingNodeAffinitiesAttributes struct {
	ref terra.Reference
}

func (na SchedulingNodeAffinitiesAttributes) InternalRef() (terra.Reference, error) {
	return na.ref, nil
}

func (na SchedulingNodeAffinitiesAttributes) InternalWithRef(ref terra.Reference) SchedulingNodeAffinitiesAttributes {
	return SchedulingNodeAffinitiesAttributes{ref: ref}
}

func (na SchedulingNodeAffinitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return na.ref.InternalTokens()
}

func (na SchedulingNodeAffinitiesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("key"))
}

func (na SchedulingNodeAffinitiesAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("operator"))
}

func (na SchedulingNodeAffinitiesAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("values"))
}

type ShieldedInstanceConfigAttributes struct {
	ref terra.Reference
}

func (sic ShieldedInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return sic.ref, nil
}

func (sic ShieldedInstanceConfigAttributes) InternalWithRef(ref terra.Reference) ShieldedInstanceConfigAttributes {
	return ShieldedInstanceConfigAttributes{ref: ref}
}

func (sic ShieldedInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sic.ref.InternalTokens()
}

func (sic ShieldedInstanceConfigAttributes) EnableIntegrityMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_integrity_monitoring"))
}

func (sic ShieldedInstanceConfigAttributes) EnableSecureBoot() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_secure_boot"))
}

func (sic ShieldedInstanceConfigAttributes) EnableVtpm() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_vtpm"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AttachedDiskState struct {
	DeviceName              string `json:"device_name"`
	DiskEncryptionKeyRaw    string `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string `json:"disk_encryption_key_sha256"`
	KmsKeySelfLink          string `json:"kms_key_self_link"`
	Mode                    string `json:"mode"`
	Source                  string `json:"source"`
}

type BootDiskState struct {
	AutoDelete              bool                            `json:"auto_delete"`
	DeviceName              string                          `json:"device_name"`
	DiskEncryptionKeyRaw    string                          `json:"disk_encryption_key_raw"`
	DiskEncryptionKeySha256 string                          `json:"disk_encryption_key_sha256"`
	KmsKeySelfLink          string                          `json:"kms_key_self_link"`
	Mode                    string                          `json:"mode"`
	Source                  string                          `json:"source"`
	InitializeParams        []BootDiskInitializeParamsState `json:"initialize_params"`
}

type BootDiskInitializeParamsState struct {
	EnableConfidentialCompute bool              `json:"enable_confidential_compute"`
	Image                     string            `json:"image"`
	Labels                    map[string]string `json:"labels"`
	ProvisionedIops           float64           `json:"provisioned_iops"`
	ProvisionedThroughput     float64           `json:"provisioned_throughput"`
	ResourceManagerTags       map[string]string `json:"resource_manager_tags"`
	Size                      float64           `json:"size"`
	Type                      string            `json:"type"`
}

type GuestAcceleratorState struct {
	Count float64 `json:"count"`
	Type  string  `json:"type"`
}

type ScratchDiskState struct {
	DeviceName string  `json:"device_name"`
	Interface  string  `json:"interface"`
	Size       float64 `json:"size"`
}

type ServiceAccountState struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type AdvancedMachineFeaturesState struct {
	EnableNestedVirtualization bool    `json:"enable_nested_virtualization"`
	ThreadsPerCore             float64 `json:"threads_per_core"`
	VisibleCoreCount           float64 `json:"visible_core_count"`
}

type ConfidentialInstanceConfigState struct {
	ConfidentialInstanceType  string `json:"confidential_instance_type"`
	EnableConfidentialCompute bool   `json:"enable_confidential_compute"`
}

type NetworkInterfaceState struct {
	InternalIpv6PrefixLength float64                                 `json:"internal_ipv6_prefix_length"`
	Ipv6AccessType           string                                  `json:"ipv6_access_type"`
	Ipv6Address              string                                  `json:"ipv6_address"`
	Name                     string                                  `json:"name"`
	Network                  string                                  `json:"network"`
	NetworkAttachment        string                                  `json:"network_attachment"`
	NetworkIp                string                                  `json:"network_ip"`
	NicType                  string                                  `json:"nic_type"`
	QueueCount               float64                                 `json:"queue_count"`
	SecurityPolicy           string                                  `json:"security_policy"`
	StackType                string                                  `json:"stack_type"`
	Subnetwork               string                                  `json:"subnetwork"`
	SubnetworkProject        string                                  `json:"subnetwork_project"`
	AccessConfig             []NetworkInterfaceAccessConfigState     `json:"access_config"`
	AliasIpRange             []NetworkInterfaceAliasIpRangeState     `json:"alias_ip_range"`
	Ipv6AccessConfig         []NetworkInterfaceIpv6AccessConfigState `json:"ipv6_access_config"`
}

type NetworkInterfaceAccessConfigState struct {
	NatIp               string `json:"nat_ip"`
	NetworkTier         string `json:"network_tier"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
	SecurityPolicy      string `json:"security_policy"`
}

type NetworkInterfaceAliasIpRangeState struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type NetworkInterfaceIpv6AccessConfigState struct {
	ExternalIpv6             string `json:"external_ipv6"`
	ExternalIpv6PrefixLength string `json:"external_ipv6_prefix_length"`
	Name                     string `json:"name"`
	NetworkTier              string `json:"network_tier"`
	PublicPtrDomainName      string `json:"public_ptr_domain_name"`
	SecurityPolicy           string `json:"security_policy"`
}

type NetworkPerformanceConfigState struct {
	TotalEgressBandwidthTier string `json:"total_egress_bandwidth_tier"`
}

type ParamsState struct {
	ResourceManagerTags map[string]string `json:"resource_manager_tags"`
}

type ReservationAffinityState struct {
	Type                string                                        `json:"type"`
	SpecificReservation []ReservationAffinitySpecificReservationState `json:"specific_reservation"`
}

type ReservationAffinitySpecificReservationState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SchedulingState struct {
	AutomaticRestart          bool                                     `json:"automatic_restart"`
	InstanceTerminationAction string                                   `json:"instance_termination_action"`
	MaintenanceInterval       string                                   `json:"maintenance_interval"`
	MinNodeCpus               float64                                  `json:"min_node_cpus"`
	OnHostMaintenance         string                                   `json:"on_host_maintenance"`
	Preemptible               bool                                     `json:"preemptible"`
	ProvisioningModel         string                                   `json:"provisioning_model"`
	LocalSsdRecoveryTimeout   []SchedulingLocalSsdRecoveryTimeoutState `json:"local_ssd_recovery_timeout"`
	MaxRunDuration            []SchedulingMaxRunDurationState          `json:"max_run_duration"`
	NodeAffinities            []SchedulingNodeAffinitiesState          `json:"node_affinities"`
}

type SchedulingLocalSsdRecoveryTimeoutState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type SchedulingMaxRunDurationState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type SchedulingNodeAffinitiesState struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type ShieldedInstanceConfigState struct {
	EnableIntegrityMonitoring bool `json:"enable_integrity_monitoring"`
	EnableSecureBoot          bool `json:"enable_secure_boot"`
	EnableVtpm                bool `json:"enable_vtpm"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
