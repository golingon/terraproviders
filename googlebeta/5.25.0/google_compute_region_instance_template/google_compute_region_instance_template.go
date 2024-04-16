// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_instance_template

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_compute_region_instance_template.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionInstanceTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrit *Resource) Type() string {
	return "google_compute_region_instance_template"
}

// LocalName returns the local name for [Resource].
func (gcrit *Resource) LocalName() string {
	return gcrit.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrit *Resource) Configuration() interface{} {
	return gcrit.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrit *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrit)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrit *Resource) Dependencies() terra.Dependencies {
	return gcrit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrit *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrit.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrit *Resource) Attributes() googleComputeRegionInstanceTemplateAttributes {
	return googleComputeRegionInstanceTemplateAttributes{ref: terra.ReferenceResource(gcrit)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrit *Resource) ImportState(state io.Reader) error {
	gcrit.state = &googleComputeRegionInstanceTemplateState{}
	if err := json.NewDecoder(state).Decode(gcrit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrit.Type(), gcrit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrit *Resource) State() (*googleComputeRegionInstanceTemplateState, bool) {
	return gcrit.state, gcrit.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrit *Resource) StateMust() *googleComputeRegionInstanceTemplateState {
	if gcrit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrit.Type(), gcrit.LocalName()))
	}
	return gcrit.state
}

// Args contains the configurations for google_compute_region_instance_template.
type Args struct {
	// CanIpForward: bool, optional
	CanIpForward terra.BoolValue `hcl:"can_ip_forward,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableDisplay: bool, optional
	EnableDisplay terra.BoolValue `hcl:"enable_display,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceDescription: string, optional
	InstanceDescription terra.StringValue `hcl:"instance_description,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MachineType: string, required
	MachineType terra.StringValue `hcl:"machine_type,attr" validate:"required"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// MetadataStartupScript: string, optional
	MetadataStartupScript terra.StringValue `hcl:"metadata_startup_script,attr"`
	// MinCpuPlatform: string, optional
	MinCpuPlatform terra.StringValue `hcl:"min_cpu_platform,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ResourceManagerTags: map of string, optional
	ResourceManagerTags terra.MapValue[terra.StringValue] `hcl:"resource_manager_tags,attr"`
	// ResourcePolicies: list of string, optional
	ResourcePolicies terra.ListValue[terra.StringValue] `hcl:"resource_policies,attr"`
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// AdvancedMachineFeatures: optional
	AdvancedMachineFeatures *AdvancedMachineFeatures `hcl:"advanced_machine_features,block"`
	// ConfidentialInstanceConfig: optional
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `hcl:"confidential_instance_config,block"`
	// Disk: min=1
	Disk []Disk `hcl:"disk,block" validate:"min=1"`
	// GuestAccelerator: min=0
	GuestAccelerator []GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// NetworkPerformanceConfig: optional
	NetworkPerformanceConfig *NetworkPerformanceConfig `hcl:"network_performance_config,block"`
	// ReservationAffinity: optional
	ReservationAffinity *ReservationAffinity `hcl:"reservation_affinity,block"`
	// Scheduling: optional
	Scheduling *Scheduling `hcl:"scheduling,block"`
	// ServiceAccount: optional
	ServiceAccount *ServiceAccount `hcl:"service_account,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeRegionInstanceTemplateAttributes struct {
	ref terra.Reference
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(gcrit.ref.Append("can_ip_forward"))
}

// Description returns a reference to field description of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcrit.ref.Append("effective_labels"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(gcrit.ref.Append("enable_display"))
}

// Id returns a reference to field id of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("id"))
}

// InstanceDescription returns a reference to field instance_description of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) InstanceDescription() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("instance_description"))
}

// Labels returns a reference to field labels of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcrit.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcrit.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("min_cpu_platform"))
}

// Name returns a reference to field name of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("name_prefix"))
}

// Project returns a reference to field project of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("region"))
}

// ResourceManagerTags returns a reference to field resource_manager_tags of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) ResourceManagerTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcrit.ref.Append("resource_manager_tags"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcrit.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("self_link"))
}

// Tags returns a reference to field tags of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrit.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcrit.ref.Append("tags_fingerprint"))
}

// TerraformLabels returns a reference to field terraform_labels of google_compute_region_instance_template.
func (gcrit googleComputeRegionInstanceTemplateAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcrit.ref.Append("terraform_labels"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) AdvancedMachineFeatures() terra.ListValue[AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[AdvancedMachineFeaturesAttributes](gcrit.ref.Append("advanced_machine_features"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) ConfidentialInstanceConfig() terra.ListValue[ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[ConfidentialInstanceConfigAttributes](gcrit.ref.Append("confidential_instance_config"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) Disk() terra.ListValue[DiskAttributes] {
	return terra.ReferenceAsList[DiskAttributes](gcrit.ref.Append("disk"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) GuestAccelerator() terra.ListValue[GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[GuestAcceleratorAttributes](gcrit.ref.Append("guest_accelerator"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) NetworkInterface() terra.ListValue[NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceAttributes](gcrit.ref.Append("network_interface"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) NetworkPerformanceConfig() terra.ListValue[NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[NetworkPerformanceConfigAttributes](gcrit.ref.Append("network_performance_config"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) ReservationAffinity() terra.ListValue[ReservationAffinityAttributes] {
	return terra.ReferenceAsList[ReservationAffinityAttributes](gcrit.ref.Append("reservation_affinity"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) Scheduling() terra.ListValue[SchedulingAttributes] {
	return terra.ReferenceAsList[SchedulingAttributes](gcrit.ref.Append("scheduling"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) ServiceAccount() terra.ListValue[ServiceAccountAttributes] {
	return terra.ReferenceAsList[ServiceAccountAttributes](gcrit.ref.Append("service_account"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) ShieldedInstanceConfig() terra.ListValue[ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[ShieldedInstanceConfigAttributes](gcrit.ref.Append("shielded_instance_config"))
}

func (gcrit googleComputeRegionInstanceTemplateAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrit.ref.Append("timeouts"))
}

type googleComputeRegionInstanceTemplateState struct {
	CanIpForward               bool                              `json:"can_ip_forward"`
	Description                string                            `json:"description"`
	EffectiveLabels            map[string]string                 `json:"effective_labels"`
	EnableDisplay              bool                              `json:"enable_display"`
	Id                         string                            `json:"id"`
	InstanceDescription        string                            `json:"instance_description"`
	Labels                     map[string]string                 `json:"labels"`
	MachineType                string                            `json:"machine_type"`
	Metadata                   map[string]string                 `json:"metadata"`
	MetadataFingerprint        string                            `json:"metadata_fingerprint"`
	MetadataStartupScript      string                            `json:"metadata_startup_script"`
	MinCpuPlatform             string                            `json:"min_cpu_platform"`
	Name                       string                            `json:"name"`
	NamePrefix                 string                            `json:"name_prefix"`
	Project                    string                            `json:"project"`
	Region                     string                            `json:"region"`
	ResourceManagerTags        map[string]string                 `json:"resource_manager_tags"`
	ResourcePolicies           []string                          `json:"resource_policies"`
	SelfLink                   string                            `json:"self_link"`
	Tags                       []string                          `json:"tags"`
	TagsFingerprint            string                            `json:"tags_fingerprint"`
	TerraformLabels            map[string]string                 `json:"terraform_labels"`
	AdvancedMachineFeatures    []AdvancedMachineFeaturesState    `json:"advanced_machine_features"`
	ConfidentialInstanceConfig []ConfidentialInstanceConfigState `json:"confidential_instance_config"`
	Disk                       []DiskState                       `json:"disk"`
	GuestAccelerator           []GuestAcceleratorState           `json:"guest_accelerator"`
	NetworkInterface           []NetworkInterfaceState           `json:"network_interface"`
	NetworkPerformanceConfig   []NetworkPerformanceConfigState   `json:"network_performance_config"`
	ReservationAffinity        []ReservationAffinityState        `json:"reservation_affinity"`
	Scheduling                 []SchedulingState                 `json:"scheduling"`
	ServiceAccount             []ServiceAccountState             `json:"service_account"`
	ShieldedInstanceConfig     []ShieldedInstanceConfigState     `json:"shielded_instance_config"`
	Timeouts                   *TimeoutsState                    `json:"timeouts"`
}
