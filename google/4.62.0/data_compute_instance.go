// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeinstance "github.com/golingon/terraproviders/google/4.62.0/datacomputeinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeInstance creates a new instance of [DataComputeInstance].
func NewDataComputeInstance(name string, args DataComputeInstanceArgs) *DataComputeInstance {
	return &DataComputeInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeInstance)(nil)

// DataComputeInstance represents the Terraform data resource google_compute_instance.
type DataComputeInstance struct {
	Name string
	Args DataComputeInstanceArgs
}

// DataSource returns the Terraform object type for [DataComputeInstance].
func (ci *DataComputeInstance) DataSource() string {
	return "google_compute_instance"
}

// LocalName returns the local name for [DataComputeInstance].
func (ci *DataComputeInstance) LocalName() string {
	return ci.Name
}

// Configuration returns the configuration (args) for [DataComputeInstance].
func (ci *DataComputeInstance) Configuration() interface{} {
	return ci.Args
}

// Attributes returns the attributes for [DataComputeInstance].
func (ci *DataComputeInstance) Attributes() dataComputeInstanceAttributes {
	return dataComputeInstanceAttributes{ref: terra.ReferenceDataResource(ci)}
}

// DataComputeInstanceArgs contains the configurations for google_compute_instance.
type DataComputeInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SelfLink: string, optional
	SelfLink terra.StringValue `hcl:"self_link,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AdvancedMachineFeatures: min=0
	AdvancedMachineFeatures []datacomputeinstance.AdvancedMachineFeatures `hcl:"advanced_machine_features,block" validate:"min=0"`
	// AttachedDisk: min=0
	AttachedDisk []datacomputeinstance.AttachedDisk `hcl:"attached_disk,block" validate:"min=0"`
	// BootDisk: min=0
	BootDisk []datacomputeinstance.BootDisk `hcl:"boot_disk,block" validate:"min=0"`
	// ConfidentialInstanceConfig: min=0
	ConfidentialInstanceConfig []datacomputeinstance.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block" validate:"min=0"`
	// GuestAccelerator: min=0
	GuestAccelerator []datacomputeinstance.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []datacomputeinstance.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// ReservationAffinity: min=0
	ReservationAffinity []datacomputeinstance.ReservationAffinity `hcl:"reservation_affinity,block" validate:"min=0"`
	// Scheduling: min=0
	Scheduling []datacomputeinstance.Scheduling `hcl:"scheduling,block" validate:"min=0"`
	// ScratchDisk: min=0
	ScratchDisk []datacomputeinstance.ScratchDisk `hcl:"scratch_disk,block" validate:"min=0"`
	// ServiceAccount: min=0
	ServiceAccount []datacomputeinstance.ServiceAccount `hcl:"service_account,block" validate:"min=0"`
	// ShieldedInstanceConfig: min=0
	ShieldedInstanceConfig []datacomputeinstance.ShieldedInstanceConfig `hcl:"shielded_instance_config,block" validate:"min=0"`
}
type dataComputeInstanceAttributes struct {
	ref terra.Reference
}

// AllowStoppingForUpdate returns a reference to field allow_stopping_for_update of google_compute_instance.
func (ci dataComputeInstanceAttributes) AllowStoppingForUpdate() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("allow_stopping_for_update"))
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_instance.
func (ci dataComputeInstanceAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("can_ip_forward"))
}

// CpuPlatform returns a reference to field cpu_platform of google_compute_instance.
func (ci dataComputeInstanceAttributes) CpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("cpu_platform"))
}

// CurrentStatus returns a reference to field current_status of google_compute_instance.
func (ci dataComputeInstanceAttributes) CurrentStatus() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("current_status"))
}

// DeletionProtection returns a reference to field deletion_protection of google_compute_instance.
func (ci dataComputeInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("deletion_protection"))
}

// Description returns a reference to field description of google_compute_instance.
func (ci dataComputeInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("description"))
}

// DesiredStatus returns a reference to field desired_status of google_compute_instance.
func (ci dataComputeInstanceAttributes) DesiredStatus() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("desired_status"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_instance.
func (ci dataComputeInstanceAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("enable_display"))
}

// Hostname returns a reference to field hostname of google_compute_instance.
func (ci dataComputeInstanceAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("hostname"))
}

// Id returns a reference to field id of google_compute_instance.
func (ci dataComputeInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_compute_instance.
func (ci dataComputeInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("instance_id"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_instance.
func (ci dataComputeInstanceAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_instance.
func (ci dataComputeInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_instance.
func (ci dataComputeInstanceAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_instance.
func (ci dataComputeInstanceAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_instance.
func (ci dataComputeInstanceAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_instance.
func (ci dataComputeInstanceAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_instance.
func (ci dataComputeInstanceAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("min_cpu_platform"))
}

// Name returns a reference to field name of google_compute_instance.
func (ci dataComputeInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_instance.
func (ci dataComputeInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("project"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_instance.
func (ci dataComputeInstanceAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_instance.
func (ci dataComputeInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("self_link"))
}

// Tags returns a reference to field tags of google_compute_instance.
func (ci dataComputeInstanceAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ci.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_instance.
func (ci dataComputeInstanceAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("tags_fingerprint"))
}

// Zone returns a reference to field zone of google_compute_instance.
func (ci dataComputeInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("zone"))
}

func (ci dataComputeInstanceAttributes) AdvancedMachineFeatures() terra.ListValue[datacomputeinstance.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.AdvancedMachineFeaturesAttributes](ci.ref.Append("advanced_machine_features"))
}

func (ci dataComputeInstanceAttributes) AttachedDisk() terra.ListValue[datacomputeinstance.AttachedDiskAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.AttachedDiskAttributes](ci.ref.Append("attached_disk"))
}

func (ci dataComputeInstanceAttributes) BootDisk() terra.ListValue[datacomputeinstance.BootDiskAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.BootDiskAttributes](ci.ref.Append("boot_disk"))
}

func (ci dataComputeInstanceAttributes) ConfidentialInstanceConfig() terra.ListValue[datacomputeinstance.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.ConfidentialInstanceConfigAttributes](ci.ref.Append("confidential_instance_config"))
}

func (ci dataComputeInstanceAttributes) GuestAccelerator() terra.ListValue[datacomputeinstance.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.GuestAcceleratorAttributes](ci.ref.Append("guest_accelerator"))
}

func (ci dataComputeInstanceAttributes) NetworkInterface() terra.ListValue[datacomputeinstance.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.NetworkInterfaceAttributes](ci.ref.Append("network_interface"))
}

func (ci dataComputeInstanceAttributes) ReservationAffinity() terra.ListValue[datacomputeinstance.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.ReservationAffinityAttributes](ci.ref.Append("reservation_affinity"))
}

func (ci dataComputeInstanceAttributes) Scheduling() terra.ListValue[datacomputeinstance.SchedulingAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.SchedulingAttributes](ci.ref.Append("scheduling"))
}

func (ci dataComputeInstanceAttributes) ScratchDisk() terra.ListValue[datacomputeinstance.ScratchDiskAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.ScratchDiskAttributes](ci.ref.Append("scratch_disk"))
}

func (ci dataComputeInstanceAttributes) ServiceAccount() terra.ListValue[datacomputeinstance.ServiceAccountAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.ServiceAccountAttributes](ci.ref.Append("service_account"))
}

func (ci dataComputeInstanceAttributes) ShieldedInstanceConfig() terra.ListValue[datacomputeinstance.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeinstance.ShieldedInstanceConfigAttributes](ci.ref.Append("shielded_instance_config"))
}
