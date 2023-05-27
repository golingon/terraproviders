// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputeinstancetemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdvancedMachineFeatures struct{}

type ConfidentialInstanceConfig struct{}

type Disk struct {
	// DiskEncryptionKey: min=0
	DiskEncryptionKey []DiskEncryptionKey `hcl:"disk_encryption_key,block" validate:"min=0"`
	// SourceImageEncryptionKey: min=0
	SourceImageEncryptionKey []SourceImageEncryptionKey `hcl:"source_image_encryption_key,block" validate:"min=0"`
	// SourceSnapshotEncryptionKey: min=0
	SourceSnapshotEncryptionKey []SourceSnapshotEncryptionKey `hcl:"source_snapshot_encryption_key,block" validate:"min=0"`
}

type DiskEncryptionKey struct{}

type SourceImageEncryptionKey struct{}

type SourceSnapshotEncryptionKey struct{}

type GuestAccelerator struct{}

type NetworkInterface struct {
	// AccessConfig: min=0
	AccessConfig []AccessConfig `hcl:"access_config,block" validate:"min=0"`
	// AliasIpRange: min=0
	AliasIpRange []AliasIpRange `hcl:"alias_ip_range,block" validate:"min=0"`
	// Ipv6AccessConfig: min=0
	Ipv6AccessConfig []Ipv6AccessConfig `hcl:"ipv6_access_config,block" validate:"min=0"`
}

type AccessConfig struct{}

type AliasIpRange struct{}

type Ipv6AccessConfig struct{}

type ReservationAffinity struct {
	// SpecificReservation: min=0
	SpecificReservation []SpecificReservation `hcl:"specific_reservation,block" validate:"min=0"`
}

type SpecificReservation struct{}

type Scheduling struct {
	// NodeAffinities: min=0
	NodeAffinities []NodeAffinities `hcl:"node_affinities,block" validate:"min=0"`
}

type NodeAffinities struct{}

type ServiceAccount struct{}

type ShieldedInstanceConfig struct{}

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

func (cic ConfidentialInstanceConfigAttributes) EnableConfidentialCompute() terra.BoolValue {
	return terra.ReferenceAsBool(cic.ref.Append("enable_confidential_compute"))
}

type DiskAttributes struct {
	ref terra.Reference
}

func (d DiskAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DiskAttributes) InternalWithRef(ref terra.Reference) DiskAttributes {
	return DiskAttributes{ref: ref}
}

func (d DiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DiskAttributes) AutoDelete() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("auto_delete"))
}

func (d DiskAttributes) Boot() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("boot"))
}

func (d DiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("device_name"))
}

func (d DiskAttributes) DiskName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("disk_name"))
}

func (d DiskAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("disk_size_gb"))
}

func (d DiskAttributes) DiskType() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("disk_type"))
}

func (d DiskAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("interface"))
}

func (d DiskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](d.ref.Append("labels"))
}

func (d DiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("mode"))
}

func (d DiskAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("resource_policies"))
}

func (d DiskAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source"))
}

func (d DiskAttributes) SourceImage() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source_image"))
}

func (d DiskAttributes) SourceSnapshot() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source_snapshot"))
}

func (d DiskAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("type"))
}

func (d DiskAttributes) DiskEncryptionKey() terra.ListValue[DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DiskEncryptionKeyAttributes](d.ref.Append("disk_encryption_key"))
}

func (d DiskAttributes) SourceImageEncryptionKey() terra.ListValue[SourceImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[SourceImageEncryptionKeyAttributes](d.ref.Append("source_image_encryption_key"))
}

func (d DiskAttributes) SourceSnapshotEncryptionKey() terra.ListValue[SourceSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[SourceSnapshotEncryptionKeyAttributes](d.ref.Append("source_snapshot_encryption_key"))
}

type DiskEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (dek DiskEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return dek.ref, nil
}

func (dek DiskEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) DiskEncryptionKeyAttributes {
	return DiskEncryptionKeyAttributes{ref: ref}
}

func (dek DiskEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dek.ref.InternalTokens()
}

func (dek DiskEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(dek.ref.Append("kms_key_self_link"))
}

type SourceImageEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (siek SourceImageEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return siek.ref, nil
}

func (siek SourceImageEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) SourceImageEncryptionKeyAttributes {
	return SourceImageEncryptionKeyAttributes{ref: ref}
}

func (siek SourceImageEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return siek.ref.InternalTokens()
}

func (siek SourceImageEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("kms_key_self_link"))
}

func (siek SourceImageEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(siek.ref.Append("kms_key_service_account"))
}

type SourceSnapshotEncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return ssek.ref, nil
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalWithRef(ref terra.Reference) SourceSnapshotEncryptionKeyAttributes {
	return SourceSnapshotEncryptionKeyAttributes{ref: ref}
}

func (ssek SourceSnapshotEncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssek.ref.InternalTokens()
}

func (ssek SourceSnapshotEncryptionKeyAttributes) KmsKeySelfLink() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_self_link"))
}

func (ssek SourceSnapshotEncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ssek.ref.Append("kms_key_service_account"))
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

func (ni NetworkInterfaceAttributes) Ipv6AccessType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("ipv6_access_type"))
}

func (ni NetworkInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("name"))
}

func (ni NetworkInterfaceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network"))
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

func (ni NetworkInterfaceAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("stack_type"))
}

func (ni NetworkInterfaceAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnetwork"))
}

func (ni NetworkInterfaceAttributes) SubnetworkProject() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnetwork_project"))
}

func (ni NetworkInterfaceAttributes) AccessConfig() terra.ListValue[AccessConfigAttributes] {
	return terra.ReferenceAsList[AccessConfigAttributes](ni.ref.Append("access_config"))
}

func (ni NetworkInterfaceAttributes) AliasIpRange() terra.ListValue[AliasIpRangeAttributes] {
	return terra.ReferenceAsList[AliasIpRangeAttributes](ni.ref.Append("alias_ip_range"))
}

func (ni NetworkInterfaceAttributes) Ipv6AccessConfig() terra.ListValue[Ipv6AccessConfigAttributes] {
	return terra.ReferenceAsList[Ipv6AccessConfigAttributes](ni.ref.Append("ipv6_access_config"))
}

type AccessConfigAttributes struct {
	ref terra.Reference
}

func (ac AccessConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AccessConfigAttributes) InternalWithRef(ref terra.Reference) AccessConfigAttributes {
	return AccessConfigAttributes{ref: ref}
}

func (ac AccessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AccessConfigAttributes) NatIp() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("nat_ip"))
}

func (ac AccessConfigAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("network_tier"))
}

func (ac AccessConfigAttributes) PublicPtrDomainName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("public_ptr_domain_name"))
}

type AliasIpRangeAttributes struct {
	ref terra.Reference
}

func (air AliasIpRangeAttributes) InternalRef() (terra.Reference, error) {
	return air.ref, nil
}

func (air AliasIpRangeAttributes) InternalWithRef(ref terra.Reference) AliasIpRangeAttributes {
	return AliasIpRangeAttributes{ref: ref}
}

func (air AliasIpRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return air.ref.InternalTokens()
}

func (air AliasIpRangeAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("ip_cidr_range"))
}

func (air AliasIpRangeAttributes) SubnetworkRangeName() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("subnetwork_range_name"))
}

type Ipv6AccessConfigAttributes struct {
	ref terra.Reference
}

func (iac Ipv6AccessConfigAttributes) InternalRef() (terra.Reference, error) {
	return iac.ref, nil
}

func (iac Ipv6AccessConfigAttributes) InternalWithRef(ref terra.Reference) Ipv6AccessConfigAttributes {
	return Ipv6AccessConfigAttributes{ref: ref}
}

func (iac Ipv6AccessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iac.ref.InternalTokens()
}

func (iac Ipv6AccessConfigAttributes) ExternalIpv6() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("external_ipv6"))
}

func (iac Ipv6AccessConfigAttributes) ExternalIpv6PrefixLength() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("external_ipv6_prefix_length"))
}

func (iac Ipv6AccessConfigAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("network_tier"))
}

func (iac Ipv6AccessConfigAttributes) PublicPtrDomainName() terra.StringValue {
	return terra.ReferenceAsString(iac.ref.Append("public_ptr_domain_name"))
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

func (ra ReservationAffinityAttributes) SpecificReservation() terra.ListValue[SpecificReservationAttributes] {
	return terra.ReferenceAsList[SpecificReservationAttributes](ra.ref.Append("specific_reservation"))
}

type SpecificReservationAttributes struct {
	ref terra.Reference
}

func (sr SpecificReservationAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr SpecificReservationAttributes) InternalWithRef(ref terra.Reference) SpecificReservationAttributes {
	return SpecificReservationAttributes{ref: ref}
}

func (sr SpecificReservationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr SpecificReservationAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("key"))
}

func (sr SpecificReservationAttributes) Values() terra.ListValue[terra.StringValue] {
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

func (s SchedulingAttributes) NodeAffinities() terra.SetValue[NodeAffinitiesAttributes] {
	return terra.ReferenceAsSet[NodeAffinitiesAttributes](s.ref.Append("node_affinities"))
}

type NodeAffinitiesAttributes struct {
	ref terra.Reference
}

func (na NodeAffinitiesAttributes) InternalRef() (terra.Reference, error) {
	return na.ref, nil
}

func (na NodeAffinitiesAttributes) InternalWithRef(ref terra.Reference) NodeAffinitiesAttributes {
	return NodeAffinitiesAttributes{ref: ref}
}

func (na NodeAffinitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return na.ref.InternalTokens()
}

func (na NodeAffinitiesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("key"))
}

func (na NodeAffinitiesAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("operator"))
}

func (na NodeAffinitiesAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("values"))
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

type AdvancedMachineFeaturesState struct {
	EnableNestedVirtualization bool    `json:"enable_nested_virtualization"`
	ThreadsPerCore             float64 `json:"threads_per_core"`
	VisibleCoreCount           float64 `json:"visible_core_count"`
}

type ConfidentialInstanceConfigState struct {
	EnableConfidentialCompute bool `json:"enable_confidential_compute"`
}

type DiskState struct {
	AutoDelete                  bool                               `json:"auto_delete"`
	Boot                        bool                               `json:"boot"`
	DeviceName                  string                             `json:"device_name"`
	DiskName                    string                             `json:"disk_name"`
	DiskSizeGb                  float64                            `json:"disk_size_gb"`
	DiskType                    string                             `json:"disk_type"`
	Interface                   string                             `json:"interface"`
	Labels                      map[string]string                  `json:"labels"`
	Mode                        string                             `json:"mode"`
	ResourcePolicies            []string                           `json:"resource_policies"`
	Source                      string                             `json:"source"`
	SourceImage                 string                             `json:"source_image"`
	SourceSnapshot              string                             `json:"source_snapshot"`
	Type                        string                             `json:"type"`
	DiskEncryptionKey           []DiskEncryptionKeyState           `json:"disk_encryption_key"`
	SourceImageEncryptionKey    []SourceImageEncryptionKeyState    `json:"source_image_encryption_key"`
	SourceSnapshotEncryptionKey []SourceSnapshotEncryptionKeyState `json:"source_snapshot_encryption_key"`
}

type DiskEncryptionKeyState struct {
	KmsKeySelfLink string `json:"kms_key_self_link"`
}

type SourceImageEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
}

type SourceSnapshotEncryptionKeyState struct {
	KmsKeySelfLink       string `json:"kms_key_self_link"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
}

type GuestAcceleratorState struct {
	Count float64 `json:"count"`
	Type  string  `json:"type"`
}

type NetworkInterfaceState struct {
	Ipv6AccessType    string                  `json:"ipv6_access_type"`
	Name              string                  `json:"name"`
	Network           string                  `json:"network"`
	NetworkIp         string                  `json:"network_ip"`
	NicType           string                  `json:"nic_type"`
	QueueCount        float64                 `json:"queue_count"`
	StackType         string                  `json:"stack_type"`
	Subnetwork        string                  `json:"subnetwork"`
	SubnetworkProject string                  `json:"subnetwork_project"`
	AccessConfig      []AccessConfigState     `json:"access_config"`
	AliasIpRange      []AliasIpRangeState     `json:"alias_ip_range"`
	Ipv6AccessConfig  []Ipv6AccessConfigState `json:"ipv6_access_config"`
}

type AccessConfigState struct {
	NatIp               string `json:"nat_ip"`
	NetworkTier         string `json:"network_tier"`
	PublicPtrDomainName string `json:"public_ptr_domain_name"`
}

type AliasIpRangeState struct {
	IpCidrRange         string `json:"ip_cidr_range"`
	SubnetworkRangeName string `json:"subnetwork_range_name"`
}

type Ipv6AccessConfigState struct {
	ExternalIpv6             string `json:"external_ipv6"`
	ExternalIpv6PrefixLength string `json:"external_ipv6_prefix_length"`
	NetworkTier              string `json:"network_tier"`
	PublicPtrDomainName      string `json:"public_ptr_domain_name"`
}

type ReservationAffinityState struct {
	Type                string                     `json:"type"`
	SpecificReservation []SpecificReservationState `json:"specific_reservation"`
}

type SpecificReservationState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SchedulingState struct {
	AutomaticRestart          bool                  `json:"automatic_restart"`
	InstanceTerminationAction string                `json:"instance_termination_action"`
	MinNodeCpus               float64               `json:"min_node_cpus"`
	OnHostMaintenance         string                `json:"on_host_maintenance"`
	Preemptible               bool                  `json:"preemptible"`
	ProvisioningModel         string                `json:"provisioning_model"`
	NodeAffinities            []NodeAffinitiesState `json:"node_affinities"`
}

type NodeAffinitiesState struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type ServiceAccountState struct {
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
}

type ShieldedInstanceConfigState struct {
	EnableIntegrityMonitoring bool `json:"enable_integrity_monitoring"`
	EnableSecureBoot          bool `json:"enable_secure_boot"`
	EnableVtpm                bool `json:"enable_vtpm"`
}
