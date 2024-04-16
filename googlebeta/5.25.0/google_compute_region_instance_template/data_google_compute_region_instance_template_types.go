// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_instance_template

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAdvancedMachineFeaturesAttributes struct {
	ref terra.Reference
}

func (amf DataAdvancedMachineFeaturesAttributes) InternalRef() (terra.Reference, error) {
	return amf.ref, nil
}

func (amf DataAdvancedMachineFeaturesAttributes) InternalWithRef(ref terra.Reference) DataAdvancedMachineFeaturesAttributes {
	return DataAdvancedMachineFeaturesAttributes{ref: ref}
}

func (amf DataAdvancedMachineFeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return amf.ref.InternalTokens()
}

func (amf DataAdvancedMachineFeaturesAttributes) EnableNestedVirtualization() terra.BoolValue {
	return terra.ReferenceAsBool(amf.ref.Append("enable_nested_virtualization"))
}

func (amf DataAdvancedMachineFeaturesAttributes) ThreadsPerCore() terra.NumberValue {
	return terra.ReferenceAsNumber(amf.ref.Append("threads_per_core"))
}

func (amf DataAdvancedMachineFeaturesAttributes) VisibleCoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(amf.ref.Append("visible_core_count"))
}

type DataConfidentialInstanceConfigAttributes struct {
	ref terra.Reference
}

func (cic DataConfidentialInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return cic.ref, nil
}

func (cic DataConfidentialInstanceConfigAttributes) InternalWithRef(ref terra.Reference) DataConfidentialInstanceConfigAttributes {
	return DataConfidentialInstanceConfigAttributes{ref: ref}
}

func (cic DataConfidentialInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cic.ref.InternalTokens()
}

func (cic DataConfidentialInstanceConfigAttributes) ConfidentialInstanceType() terra.StringValue {
	return terra.ReferenceAsString(cic.ref.Append("confidential_instance_type"))
}

func (cic DataConfidentialInstanceConfigAttributes) EnableConfidentialCompute() terra.BoolValue {
	return terra.ReferenceAsBool(cic.ref.Append("enable_confidential_compute"))
}

type DataDiskAttributes struct {
	ref terra.Reference
}

func (d DataDiskAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DataDiskAttributes) InternalWithRef(ref terra.Reference) DataDiskAttributes {
	return DataDiskAttributes{ref: ref}
}

func (d DataDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DataDiskAttributes) AutoDelete() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("auto_delete"))
}

func (d DataDiskAttributes) Boot() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("boot"))
}

func (d DataDiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("device_name"))
}

func (d DataDiskAttributes) DiskName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("disk_name"))
}

func (d DataDiskAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("disk_size_gb"))
}

func (d DataDiskAttributes) DiskType() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("disk_type"))
}

func (d DataDiskAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("interface"))
}

func (d DataDiskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](d.ref.Append("labels"))
}

func (d DataDiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("mode"))
}

func (d DataDiskAttributes) ProvisionedIops() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("provisioned_iops"))
}

func (d DataDiskAttributes) ResourceManagerTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](d.ref.Append("resource_manager_tags"))
}

func (d DataDiskAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("resource_policies"))
}

func (d DataDiskAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source"))
}

func (d DataDiskAttributes) SourceImage() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source_image"))
}

func (d DataDiskAttributes) SourceSnapshot() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source_snapshot"))
}

func (d DataDiskAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("type"))
}

func (d DataDiskAttributes) DiskEncryptionKey() terra.ListValue[DataDiskDiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DataDiskDiskEncryptionKeyAttributes](d.ref.Append("disk_encryption_key"))
}

func (d DataDiskAttributes) SourceImageEncryptionKey() terra.ListValue[DataDiskSourceImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DataDiskSourceImageEncryptionKeyAttributes](d.ref.Append("source_image_encryption_key"))
}

func (d DataDiskAttributes) SourceSnapshotEncryptionKey() terra.ListValue[DataDiskSourceSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DataDiskSourceSnapshotEncryptionKeyAttributes](d.ref.Append("source_snapshot_encryption_key"))
}

type DataDiskDiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (dek DataDiskDiskEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return dek.ref, nil
}

func (dek DataDiskDiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataDiskDiskEncryptionKeyAttributes {
	return DataDiskDiskEncryptionKeyAttributes{ref: ref}
}

func (dek DataDiskDiskEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dek.ref.InternalTokens()
}

func (dek DataDiskDiskEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("kms_key_self_link"))
}

type DataDiskSourceImageEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (siek DataDiskSourceImageEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return siek.ref, nil
}

func (siek DataDiskSourceImageEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataDiskSourceImageEncryptionKeyAttributes {
	return DataDiskSourceImageEncryptionKeyAttributes{ref: ref}
}

func (siek DataDiskSourceImageEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return siek.ref.InternalTokens()
}

func (siek DataDiskSourceImageEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("kms_key_self_link"))
}

func (siek DataDiskSourceImageEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("kms_key_service_account"))
}

type DataDiskSourceSnapshotEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ssek DataDiskSourceSnapshotEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return ssek.ref, nil
}

func (ssek DataDiskSourceSnapshotEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DataDiskSourceSnapshotEncryptionKeyAttributes {
	return DataDiskSourceSnapshotEncryptionKeyAttributes{ref: ref}
}

func (ssek DataDiskSourceSnapshotEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssek.ref.InternalTokens()
}

func (ssek DataDiskSourceSnapshotEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_self_link"))
}

func (ssek DataDiskSourceSnapshotEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_service_account"))
}

type DataGuestAcceleratorAttributes struct {
	ref terra.Reference
}

func (ga DataGuestAcceleratorAttributes) InternalRef() (terra.Reference, error) {
	return ga.ref, nil
}

func (ga DataGuestAcceleratorAttributes) InternalWithRef(ref terra.Reference) DataGuestAcceleratorAttributes {
	return DataGuestAcceleratorAttributes{ref: ref}
}

func (ga DataGuestAcceleratorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ga.ref.InternalTokens()
}

func (ga DataGuestAcceleratorAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(ga.ref.Append("count"))
}

func (ga DataGuestAcceleratorAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("type"))
}

type DataNetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni DataNetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni DataNetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) DataNetworkInterfaceAttributes {
	return DataNetworkInterfaceAttributes{ref: ref}
}

func (ni DataNetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni DataNetworkInterfaceAttributes) InternalIpv6PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("internal_ipv6_prefix_length"))
}

func (ni DataNetworkInterfaceAttributes) Ipv6AccessType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("ipv6_access_type"))
}

func (ni DataNetworkInterfaceAttributes) Ipv6Address() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("ipv6_address"))
}

func (ni DataNetworkInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("name"))
}

func (ni DataNetworkInterfaceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network"))
}

func (ni DataNetworkInterfaceAttributes) NetworkIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_ip"))
}

func (ni DataNetworkInterfaceAttributes) NicType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("nic_type"))
}

func (ni DataNetworkInterfaceAttributes) QueueCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("queue_count"))
}

func (ni DataNetworkInterfaceAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("stack_type"))
}

func (ni DataNetworkInterfaceAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnetwork"))
}

func (ni DataNetworkInterfaceAttributes) SubnetworkProject() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnetwork_project"))
}

func (ni DataNetworkInterfaceAttributes) AccessConfig() terra.ListValue[DataNetworkInterfaceAccessConfigAttributes] {
	return terra.ReferenceAsList[DataNetworkInterfaceAccessConfigAttributes](ni.ref.Append("access_config"))
}

func (ni DataNetworkInterfaceAttributes) AliasIpRange() terra.ListValue[DataNetworkInterfaceAliasIpRangeAttributes] {
	return terra.ReferenceAsList[DataNetworkInterfaceAliasIpRangeAttributes](ni.ref.Append("alias_ip_range"))
}

func (ni DataNetworkInterfaceAttributes) Ipv6AccessConfig() terra.ListValue[DataNetworkInterfaceIpv6AccessConfigAttributes] {
	return terra.ReferenceAsList[DataNetworkInterfaceIpv6AccessConfigAttributes](ni.ref.Append("ipv6_access_config"))
}

type DataNetworkInterfaceAccessConfigAttributes struct {
	ref terra.Reference
}

func (ac DataNetworkInterfaceAccessConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac DataNetworkInterfaceAccessConfigAttributes) InternalWithRef(ref terra.Reference) DataNetworkInterfaceAccessConfigAttributes {
	return DataNetworkInterfaceAccessConfigAttributes{ref: ref}
}

func (ac DataNetworkInterfaceAccessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac DataNetworkInterfaceAccessConfigAttributes) NatIp() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("nat_ip"))
}

func (ac DataNetworkInterfaceAccessConfigAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("network_tier"))
}

func (ac DataNetworkInterfaceAccessConfigAttributes) PublicPtrDomainName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("public_ptr_domain_name"))
}

type DataNetworkInterfaceAliasIpRangeAttributes struct {
	ref terra.Reference
}

func (air DataNetworkInterfaceAliasIpRangeAttributes) InternalRef() (terra.Reference, error) {
	return air.ref, nil
}

func (air DataNetworkInterfaceAliasIpRangeAttributes) InternalWithRef(ref terra.Reference) DataNetworkInterfaceAliasIpRangeAttributes {
	return DataNetworkInterfaceAliasIpRangeAttributes{ref: ref}
}

func (air DataNetworkInterfaceAliasIpRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return air.ref.InternalTokens()
}

func (air DataNetworkInterfaceAliasIpRangeAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("ip_cidr_range"))
}

func (air DataNetworkInterfaceAliasIpRangeAttributes) SubnetworkRangeName() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("subnetwork_range_name"))
}

type DataNetworkInterfaceIpv6AccessConfigAttributes struct {
	ref terra.Reference
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) InternalRef() (terra.Reference, error) {
	return iac.ref, nil
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) InternalWithRef(ref terra.Reference) DataNetworkInterfaceIpv6AccessConfigAttributes {
	return DataNetworkInterfaceIpv6AccessConfigAttributes{ref: ref}
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iac.ref.InternalTokens()
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) ExternalIpv6() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("external_ipv6"))
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) ExternalIpv6PrefixLength() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("external_ipv6_prefix_length"))
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("name"))
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("network_tier"))
}

func (iac DataNetworkInterfaceIpv6AccessConfigAttributes) PublicPtrDomainName() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("public_ptr_domain_name"))
}

type DataNetworkPerformanceConfigAttributes struct {
	ref terra.Reference
}

func (npc DataNetworkPerformanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return npc.ref, nil
}

func (npc DataNetworkPerformanceConfigAttributes) InternalWithRef(ref terra.Reference) DataNetworkPerformanceConfigAttributes {
	return DataNetworkPerformanceConfigAttributes{ref: ref}
}

func (npc DataNetworkPerformanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return npc.ref.InternalTokens()
}

func (npc DataNetworkPerformanceConfigAttributes) TotalEgressBandwidthTier() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("total_egress_bandwidth_tier"))
}

type DataReservationAffinityAttributes struct {
	ref terra.Reference
}

func (ra DataReservationAffinityAttributes) InternalRef() (terra.Reference, error) {
	return ra.ref, nil
}

func (ra DataReservationAffinityAttributes) InternalWithRef(ref terra.Reference) DataReservationAffinityAttributes {
	return DataReservationAffinityAttributes{ref: ref}
}

func (ra DataReservationAffinityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ra.ref.InternalTokens()
}

func (ra DataReservationAffinityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("type"))
}

func (ra DataReservationAffinityAttributes) SpecificReservation() terra.ListValue[DataReservationAffinitySpecificReservationAttributes] {
	return terra.ReferenceAsList[DataReservationAffinitySpecificReservationAttributes](ra.ref.Append("specific_reservation"))
}

type DataReservationAffinitySpecificReservationAttributes struct {
	ref terra.Reference
}

func (sr DataReservationAffinitySpecificReservationAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr DataReservationAffinitySpecificReservationAttributes) InternalWithRef(ref terra.Reference) DataReservationAffinitySpecificReservationAttributes {
	return DataReservationAffinitySpecificReservationAttributes{ref: ref}
}

func (sr DataReservationAffinitySpecificReservationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr DataReservationAffinitySpecificReservationAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("key"))
}

func (sr DataReservationAffinitySpecificReservationAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sr.ref.Append("values"))
}

type DataSchedulingAttributes struct {
	ref terra.Reference
}

func (s DataSchedulingAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataSchedulingAttributes) InternalWithRef(ref terra.Reference) DataSchedulingAttributes {
	return DataSchedulingAttributes{ref: ref}
}

func (s DataSchedulingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataSchedulingAttributes) AutomaticRestart() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("automatic_restart"))
}

func (s DataSchedulingAttributes) InstanceTerminationAction() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("instance_termination_action"))
}

func (s DataSchedulingAttributes) MaintenanceInterval() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("maintenance_interval"))
}

func (s DataSchedulingAttributes) MinNodeCpus() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("min_node_cpus"))
}

func (s DataSchedulingAttributes) OnHostMaintenance() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("on_host_maintenance"))
}

func (s DataSchedulingAttributes) Preemptible() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("preemptible"))
}

func (s DataSchedulingAttributes) ProvisioningModel() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("provisioning_model"))
}

func (s DataSchedulingAttributes) LocalSsdRecoveryTimeout() terra.ListValue[DataSchedulingLocalSsdRecoveryTimeoutAttributes] {
	return terra.ReferenceAsList[DataSchedulingLocalSsdRecoveryTimeoutAttributes](s.ref.Append("local_ssd_recovery_timeout"))
}

func (s DataSchedulingAttributes) MaxRunDuration() terra.ListValue[DataSchedulingMaxRunDurationAttributes] {
	return terra.ReferenceAsList[DataSchedulingMaxRunDurationAttributes](s.ref.Append("max_run_duration"))
}

func (s DataSchedulingAttributes) NodeAffinities() terra.SetValue[DataSchedulingNodeAffinitiesAttributes] {
	return terra.ReferenceAsSet[DataSchedulingNodeAffinitiesAttributes](s.ref.Append("node_affinities"))
}

type DataSchedulingLocalSsdRecoveryTimeoutAttributes struct {
	ref terra.Reference
}

func (lsrt DataSchedulingLocalSsdRecoveryTimeoutAttributes) InternalRef() (terra.Reference, error) {
	return lsrt.ref, nil
}

func (lsrt DataSchedulingLocalSsdRecoveryTimeoutAttributes) InternalWithRef(ref terra.Reference) DataSchedulingLocalSsdRecoveryTimeoutAttributes {
	return DataSchedulingLocalSsdRecoveryTimeoutAttributes{ref: ref}
}

func (lsrt DataSchedulingLocalSsdRecoveryTimeoutAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lsrt.ref.InternalTokens()
}

func (lsrt DataSchedulingLocalSsdRecoveryTimeoutAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(lsrt.ref.Append("nanos"))
}

func (lsrt DataSchedulingLocalSsdRecoveryTimeoutAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lsrt.ref.Append("seconds"))
}

type DataSchedulingMaxRunDurationAttributes struct {
	ref terra.Reference
}

func (mrd DataSchedulingMaxRunDurationAttributes) InternalRef() (terra.Reference, error) {
	return mrd.ref, nil
}

func (mrd DataSchedulingMaxRunDurationAttributes) InternalWithRef(ref terra.Reference) DataSchedulingMaxRunDurationAttributes {
	return DataSchedulingMaxRunDurationAttributes{ref: ref}
}

func (mrd DataSchedulingMaxRunDurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mrd.ref.InternalTokens()
}

func (mrd DataSchedulingMaxRunDurationAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(mrd.ref.Append("nanos"))
}

func (mrd DataSchedulingMaxRunDurationAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mrd.ref.Append("seconds"))
}

type DataSchedulingNodeAffinitiesAttributes struct {
	ref terra.Reference
}

func (na DataSchedulingNodeAffinitiesAttributes) InternalRef() (terra.Reference, error) {
	return na.ref, nil
}

func (na DataSchedulingNodeAffinitiesAttributes) InternalWithRef(ref terra.Reference) DataSchedulingNodeAffinitiesAttributes {
	return DataSchedulingNodeAffinitiesAttributes{ref: ref}
}

func (na DataSchedulingNodeAffinitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return na.ref.InternalTokens()
}

func (na DataSchedulingNodeAffinitiesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("key"))
}

func (na DataSchedulingNodeAffinitiesAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("operator"))
}

func (na DataSchedulingNodeAffinitiesAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("values"))
}

type DataServiceAccountAttributes struct {
	ref terra.Reference
}

func (sa DataServiceAccountAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa DataServiceAccountAttributes) InternalWithRef(ref terra.Reference) DataServiceAccountAttributes {
	return DataServiceAccountAttributes{ref: ref}
}

func (sa DataServiceAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa DataServiceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("email"))
}

func (sa DataServiceAccountAttributes) Scopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sa.ref.Append("scopes"))
}

type DataShieldedInstanceConfigAttributes struct {
	ref terra.Reference
}

func (sic DataShieldedInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return sic.ref, nil
}

func (sic DataShieldedInstanceConfigAttributes) InternalWithRef(ref terra.Reference) DataShieldedInstanceConfigAttributes {
	return DataShieldedInstanceConfigAttributes{ref: ref}
}

func (sic DataShieldedInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sic.ref.InternalTokens()
}

func (sic DataShieldedInstanceConfigAttributes) EnableIntegrityMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_integrity_monitoring"))
}

func (sic DataShieldedInstanceConfigAttributes) EnableSecureBoot() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_secure_boot"))
}

func (sic DataShieldedInstanceConfigAttributes) EnableVtpm() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_vtpm"))
}

type DataAdvancedMachineFeaturesState struct {
	EnableNestedVirtualization bool    `json:"enable_nested_virtualization"`
	ThreadsPerCore             float64 `json:"threads_per_core"`
	VisibleCoreCount           float64 `json:"visible_core_count"`
}

type DataConfidentialInstanceConfigState struct {
	ConfidentialInstanceType  string `json:"confidential_instance_type"`
	EnableConfidentialCompute bool   `json:"enable_confidential_compute"`
}

type DataDiskState struct {
	AutoDelete                  bool                                       `json:"auto_delete"`
	Boot                        bool                                       `json:"boot"`
	DeviceName                  string                                     `json:"device_name"`
	DiskName                    string                                     `json:"disk_name"`
	DiskSizeGb                  float64                                    `json:"disk_size_gb"`
	DiskType                    string                                     `json:"disk_type"`
	Interface                   string                                     `json:"interface"`
	Labels                      map[string]string                          `json:"labels"`
	Mode                        string                                     `json:"mode"`
	ProvisionedIops             float64                                    `json:"provisioned_iops"`
	ResourceManagerTags         map[string]string                          `json:"resource_manager_tags"`
	ResourcePolicies            []string                                   `json:"resource_policies"`
	Source                      string                                     `json:"source"`
	SourceImage                 string                                     `json:"source_image"`
	SourceSnapshot              string                                     `json:"source_snapshot"`
	Type                        string                                     `json:"type"`
	DiskEncryptionKey           []DataDiskDiskEncryptionKeyState           `json:"disk_encryption_key"`
	SourceImageEncryptionKey    []DataDiskSourceImageEncryptionKeyState    `json:"source_image_encryption_key"`
	SourceSnapshotEncryptionKey []DataDiskSourceSnapshotEncryptionKeyState `json:"source_snapshot_encryption_key"`
}

type DataDiskDiskEncryptionKeyState struct {
	KmsKeySelfLink string `json:"kms_key_self_link"`
}

type DataDiskSourceImageEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
}

type DataDiskSourceSnapshotEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
}

type DataGuestAcceleratorState struct {
	Count float64 `json:"count"`
	Type  string  `json:"type"`
}

type DataNetworkInterfaceState struct {
	InternalIpv6PrefixLength float64                                     `json:"internal_ipv6_prefix_length"`
	Ipv6AccessType           string                                      `json:"ipv6_access_type"`
	Ipv6Address              string                                      `json:"ipv6_address"`
	Name                     string                                      `json:"name"`
	Network                  string                                      `json:"network"`
	NetworkIp                string                                      `json:"network_ip"`
	NicType                  string                                      `json:"nic_type"`
	QueueCount               float64                                     `json:"queue_count"`
	StackType                string                                      `json:"stack_type"`
	Subnetwork               string                                      `json:"subnetwork"`
	SubnetworkProject        string                                      `json:"subnetwork_project"`
	AccessConfig             []DataNetworkInterfaceAccessConfigState     `json:"access_config"`
	AliasIpRange             []DataNetworkInterfaceAliasIpRangeState     `json:"alias_ip_range"`
	Ipv6AccessConfig         []DataNetworkInterfaceIpv6AccessConfigState `json:"ipv6_access_config"`
}

type DataNetworkInterfaceAccessConfigState struct {
	NatIp               string `json:"nat_ip"`
	NetworkTier         string `json:"network_tier"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
}

type DataNetworkInterfaceAliasIpRangeState struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type DataNetworkInterfaceIpv6AccessConfigState struct {
	ExternalIpv6             string `json:"external_ipv6"`
	ExternalIpv6PrefixLength string `json:"external_ipv6_prefix_length"`
	Name                     string `json:"name"`
	NetworkTier              string `json:"network_tier"`
	PublicPtrDomainName      string `json:"public_ptr_domain_name"`
}

type DataNetworkPerformanceConfigState struct {
	TotalEgressBandwidthTier string `json:"total_egress_bandwidth_tier"`
}

type DataReservationAffinityState struct {
	Type                string                                            `json:"type"`
	SpecificReservation []DataReservationAffinitySpecificReservationState `json:"specific_reservation"`
}

type DataReservationAffinitySpecificReservationState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type DataSchedulingState struct {
	AutomaticRestart          bool                                         `json:"automatic_restart"`
	InstanceTerminationAction string                                       `json:"instance_termination_action"`
	MaintenanceInterval       string                                       `json:"maintenance_interval"`
	MinNodeCpus               float64                                      `json:"min_node_cpus"`
	OnHostMaintenance         string                                       `json:"on_host_maintenance"`
	Preemptible               bool                                         `json:"preemptible"`
	ProvisioningModel         string                                       `json:"provisioning_model"`
	LocalSsdRecoveryTimeout   []DataSchedulingLocalSsdRecoveryTimeoutState `json:"local_ssd_recovery_timeout"`
	MaxRunDuration            []DataSchedulingMaxRunDurationState          `json:"max_run_duration"`
	NodeAffinities            []DataSchedulingNodeAffinitiesState          `json:"node_affinities"`
}

type DataSchedulingLocalSsdRecoveryTimeoutState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type DataSchedulingMaxRunDurationState struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type DataSchedulingNodeAffinitiesState struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type DataServiceAccountState struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type DataShieldedInstanceConfigState struct {
	EnableIntegrityMonitoring bool `json:"enable_integrity_monitoring"`
	EnableSecureBoot          bool `json:"enable_secure_boot"`
	EnableVtpm                bool `json:"enable_vtpm"`
}
