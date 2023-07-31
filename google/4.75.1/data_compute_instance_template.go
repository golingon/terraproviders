// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeinstancetemplate "github.com/golingon/terraproviders/google/4.75.1/datacomputeinstancetemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeInstanceTemplate creates a new instance of [DataComputeInstanceTemplate].
func NewDataComputeInstanceTemplate(name string, args DataComputeInstanceTemplateArgs) *DataComputeInstanceTemplate {
	return &DataComputeInstanceTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeInstanceTemplate)(nil)

// DataComputeInstanceTemplate represents the Terraform data resource google_compute_instance_template.
type DataComputeInstanceTemplate struct {
	Name string
	Args DataComputeInstanceTemplateArgs
}

// DataSource returns the Terraform object type for [DataComputeInstanceTemplate].
func (cit *DataComputeInstanceTemplate) DataSource() string {
	return "google_compute_instance_template"
}

// LocalName returns the local name for [DataComputeInstanceTemplate].
func (cit *DataComputeInstanceTemplate) LocalName() string {
	return cit.Name
}

// Configuration returns the configuration (args) for [DataComputeInstanceTemplate].
func (cit *DataComputeInstanceTemplate) Configuration() interface{} {
	return cit.Args
}

// Attributes returns the attributes for [DataComputeInstanceTemplate].
func (cit *DataComputeInstanceTemplate) Attributes() dataComputeInstanceTemplateAttributes {
	return dataComputeInstanceTemplateAttributes{ref: terra.ReferenceDataResource(cit)}
}

// DataComputeInstanceTemplateArgs contains the configurations for google_compute_instance_template.
type DataComputeInstanceTemplateArgs struct {
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SelfLinkUnique: string, optional
	SelfLinkUnique terra.StringValue `hcl:"self_link_unique,attr"`
	// AdvancedMachineFeatures: min=0
	AdvancedMachineFeatures []datacomputeinstancetemplate.AdvancedMachineFeatures `hcl:"advanced_machine_features,block" validate:"min=0"`
	// ConfidentialInstanceConfig: min=0
	ConfidentialInstanceConfig []datacomputeinstancetemplate.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block" validate:"min=0"`
	// Disk: min=0
	Disk []datacomputeinstancetemplate.Disk `hcl:"disk,block" validate:"min=0"`
	// GuestAccelerator: min=0
	GuestAccelerator []datacomputeinstancetemplate.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []datacomputeinstancetemplate.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// NetworkPerformanceConfig: min=0
	NetworkPerformanceConfig []datacomputeinstancetemplate.NetworkPerformanceConfig `hcl:"network_performance_config,block" validate:"min=0"`
	// ReservationAffinity: min=0
	ReservationAffinity []datacomputeinstancetemplate.ReservationAffinity `hcl:"reservation_affinity,block" validate:"min=0"`
	// Scheduling: min=0
	Scheduling []datacomputeinstancetemplate.Scheduling `hcl:"scheduling,block" validate:"min=0"`
	// ServiceAccount: min=0
	ServiceAccount []datacomputeinstancetemplate.ServiceAccount `hcl:"service_account,block" validate:"min=0"`
	// ShieldedInstanceConfig: min=0
	ShieldedInstanceConfig []datacomputeinstancetemplate.ShieldedInstanceConfig `hcl:"shielded_instance_config,block" validate:"min=0"`
}
type dataComputeInstanceTemplateAttributes struct {
	ref terra.Reference
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(cit.ref.Append("can_ip_forward"))
}

// Description returns a reference to field description of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("description"))
}

// Filter returns a reference to field filter of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("filter"))
}

// Id returns a reference to field id of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("id"))
}

// InstanceDescription returns a reference to field instance_description of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) InstanceDescription() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("instance_description"))
}

// Labels returns a reference to field labels of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cit.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cit.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("min_cpu_platform"))
}

// MostRecent returns a reference to field most_recent of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(cit.ref.Append("most_recent"))
}

// Name returns a reference to field name of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("name_prefix"))
}

// Project returns a reference to field project of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("region"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cit.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("self_link"))
}

// SelfLinkUnique returns a reference to field self_link_unique of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) SelfLinkUnique() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("self_link_unique"))
}

// Tags returns a reference to field tags of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cit.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_instance_template.
func (cit dataComputeInstanceTemplateAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("tags_fingerprint"))
}

func (cit dataComputeInstanceTemplateAttributes) AdvancedMachineFeatures() terra.ListValue[datacomputeinstancetemplate.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.AdvancedMachineFeaturesAttributes](cit.ref.Append("advanced_machine_features"))
}

func (cit dataComputeInstanceTemplateAttributes) ConfidentialInstanceConfig() terra.ListValue[datacomputeinstancetemplate.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.ConfidentialInstanceConfigAttributes](cit.ref.Append("confidential_instance_config"))
}

func (cit dataComputeInstanceTemplateAttributes) Disk() terra.ListValue[datacomputeinstancetemplate.DiskAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.DiskAttributes](cit.ref.Append("disk"))
}

func (cit dataComputeInstanceTemplateAttributes) GuestAccelerator() terra.ListValue[datacomputeinstancetemplate.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.GuestAcceleratorAttributes](cit.ref.Append("guest_accelerator"))
}

func (cit dataComputeInstanceTemplateAttributes) NetworkInterface() terra.ListValue[datacomputeinstancetemplate.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.NetworkInterfaceAttributes](cit.ref.Append("network_interface"))
}

func (cit dataComputeInstanceTemplateAttributes) NetworkPerformanceConfig() terra.ListValue[datacomputeinstancetemplate.NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.NetworkPerformanceConfigAttributes](cit.ref.Append("network_performance_config"))
}

func (cit dataComputeInstanceTemplateAttributes) ReservationAffinity() terra.ListValue[datacomputeinstancetemplate.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.ReservationAffinityAttributes](cit.ref.Append("reservation_affinity"))
}

func (cit dataComputeInstanceTemplateAttributes) Scheduling() terra.ListValue[datacomputeinstancetemplate.SchedulingAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.SchedulingAttributes](cit.ref.Append("scheduling"))
}

func (cit dataComputeInstanceTemplateAttributes) ServiceAccount() terra.ListValue[datacomputeinstancetemplate.ServiceAccountAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.ServiceAccountAttributes](cit.ref.Append("service_account"))
}

func (cit dataComputeInstanceTemplateAttributes) ShieldedInstanceConfig() terra.ListValue[datacomputeinstancetemplate.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeinstancetemplate.ShieldedInstanceConfigAttributes](cit.ref.Append("shielded_instance_config"))
}
