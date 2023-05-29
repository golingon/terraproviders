// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputeregioninstancetemplate "github.com/golingon/terraproviders/googlebeta/4.64.0/datacomputeregioninstancetemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeRegionInstanceTemplate creates a new instance of [DataComputeRegionInstanceTemplate].
func NewDataComputeRegionInstanceTemplate(name string, args DataComputeRegionInstanceTemplateArgs) *DataComputeRegionInstanceTemplate {
	return &DataComputeRegionInstanceTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRegionInstanceTemplate)(nil)

// DataComputeRegionInstanceTemplate represents the Terraform data resource google_compute_region_instance_template.
type DataComputeRegionInstanceTemplate struct {
	Name string
	Args DataComputeRegionInstanceTemplateArgs
}

// DataSource returns the Terraform object type for [DataComputeRegionInstanceTemplate].
func (crit *DataComputeRegionInstanceTemplate) DataSource() string {
	return "google_compute_region_instance_template"
}

// LocalName returns the local name for [DataComputeRegionInstanceTemplate].
func (crit *DataComputeRegionInstanceTemplate) LocalName() string {
	return crit.Name
}

// Configuration returns the configuration (args) for [DataComputeRegionInstanceTemplate].
func (crit *DataComputeRegionInstanceTemplate) Configuration() interface{} {
	return crit.Args
}

// Attributes returns the attributes for [DataComputeRegionInstanceTemplate].
func (crit *DataComputeRegionInstanceTemplate) Attributes() dataComputeRegionInstanceTemplateAttributes {
	return dataComputeRegionInstanceTemplateAttributes{ref: terra.ReferenceDataResource(crit)}
}

// DataComputeRegionInstanceTemplateArgs contains the configurations for google_compute_region_instance_template.
type DataComputeRegionInstanceTemplateArgs struct {
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
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// AdvancedMachineFeatures: min=0
	AdvancedMachineFeatures []datacomputeregioninstancetemplate.AdvancedMachineFeatures `hcl:"advanced_machine_features,block" validate:"min=0"`
	// ConfidentialInstanceConfig: min=0
	ConfidentialInstanceConfig []datacomputeregioninstancetemplate.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block" validate:"min=0"`
	// Disk: min=0
	Disk []datacomputeregioninstancetemplate.Disk `hcl:"disk,block" validate:"min=0"`
	// GuestAccelerator: min=0
	GuestAccelerator []datacomputeregioninstancetemplate.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []datacomputeregioninstancetemplate.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// NetworkPerformanceConfig: min=0
	NetworkPerformanceConfig []datacomputeregioninstancetemplate.NetworkPerformanceConfig `hcl:"network_performance_config,block" validate:"min=0"`
	// ReservationAffinity: min=0
	ReservationAffinity []datacomputeregioninstancetemplate.ReservationAffinity `hcl:"reservation_affinity,block" validate:"min=0"`
	// Scheduling: min=0
	Scheduling []datacomputeregioninstancetemplate.Scheduling `hcl:"scheduling,block" validate:"min=0"`
	// ServiceAccount: min=0
	ServiceAccount []datacomputeregioninstancetemplate.ServiceAccount `hcl:"service_account,block" validate:"min=0"`
	// ShieldedInstanceConfig: min=0
	ShieldedInstanceConfig []datacomputeregioninstancetemplate.ShieldedInstanceConfig `hcl:"shielded_instance_config,block" validate:"min=0"`
}
type dataComputeRegionInstanceTemplateAttributes struct {
	ref terra.Reference
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(crit.ref.Append("can_ip_forward"))
}

// Description returns a reference to field description of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("description"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(crit.ref.Append("enable_display"))
}

// Filter returns a reference to field filter of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("filter"))
}

// Id returns a reference to field id of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("id"))
}

// InstanceDescription returns a reference to field instance_description of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) InstanceDescription() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("instance_description"))
}

// Labels returns a reference to field labels of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crit.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crit.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("min_cpu_platform"))
}

// MostRecent returns a reference to field most_recent of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(crit.ref.Append("most_recent"))
}

// Name returns a reference to field name of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("name_prefix"))
}

// Project returns a reference to field project of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("region"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crit.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("self_link"))
}

// Tags returns a reference to field tags of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crit.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_region_instance_template.
func (crit dataComputeRegionInstanceTemplateAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("tags_fingerprint"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) AdvancedMachineFeatures() terra.ListValue[datacomputeregioninstancetemplate.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.AdvancedMachineFeaturesAttributes](crit.ref.Append("advanced_machine_features"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) ConfidentialInstanceConfig() terra.ListValue[datacomputeregioninstancetemplate.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.ConfidentialInstanceConfigAttributes](crit.ref.Append("confidential_instance_config"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) Disk() terra.ListValue[datacomputeregioninstancetemplate.DiskAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.DiskAttributes](crit.ref.Append("disk"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) GuestAccelerator() terra.ListValue[datacomputeregioninstancetemplate.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.GuestAcceleratorAttributes](crit.ref.Append("guest_accelerator"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) NetworkInterface() terra.ListValue[datacomputeregioninstancetemplate.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.NetworkInterfaceAttributes](crit.ref.Append("network_interface"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) NetworkPerformanceConfig() terra.ListValue[datacomputeregioninstancetemplate.NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.NetworkPerformanceConfigAttributes](crit.ref.Append("network_performance_config"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) ReservationAffinity() terra.ListValue[datacomputeregioninstancetemplate.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.ReservationAffinityAttributes](crit.ref.Append("reservation_affinity"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) Scheduling() terra.ListValue[datacomputeregioninstancetemplate.SchedulingAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.SchedulingAttributes](crit.ref.Append("scheduling"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) ServiceAccount() terra.ListValue[datacomputeregioninstancetemplate.ServiceAccountAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.ServiceAccountAttributes](crit.ref.Append("service_account"))
}

func (crit dataComputeRegionInstanceTemplateAttributes) ShieldedInstanceConfig() terra.ListValue[datacomputeregioninstancetemplate.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancetemplate.ShieldedInstanceConfigAttributes](crit.ref.Append("shielded_instance_config"))
}