// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregioninstancetemplate "github.com/golingon/terraproviders/googlebeta/4.63.1/computeregioninstancetemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionInstanceTemplate creates a new instance of [ComputeRegionInstanceTemplate].
func NewComputeRegionInstanceTemplate(name string, args ComputeRegionInstanceTemplateArgs) *ComputeRegionInstanceTemplate {
	return &ComputeRegionInstanceTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionInstanceTemplate)(nil)

// ComputeRegionInstanceTemplate represents the Terraform resource google_compute_region_instance_template.
type ComputeRegionInstanceTemplate struct {
	Name      string
	Args      ComputeRegionInstanceTemplateArgs
	state     *computeRegionInstanceTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionInstanceTemplate].
func (crit *ComputeRegionInstanceTemplate) Type() string {
	return "google_compute_region_instance_template"
}

// LocalName returns the local name for [ComputeRegionInstanceTemplate].
func (crit *ComputeRegionInstanceTemplate) LocalName() string {
	return crit.Name
}

// Configuration returns the configuration (args) for [ComputeRegionInstanceTemplate].
func (crit *ComputeRegionInstanceTemplate) Configuration() interface{} {
	return crit.Args
}

// DependOn is used for other resources to depend on [ComputeRegionInstanceTemplate].
func (crit *ComputeRegionInstanceTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(crit)
}

// Dependencies returns the list of resources [ComputeRegionInstanceTemplate] depends_on.
func (crit *ComputeRegionInstanceTemplate) Dependencies() terra.Dependencies {
	return crit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionInstanceTemplate].
func (crit *ComputeRegionInstanceTemplate) LifecycleManagement() *terra.Lifecycle {
	return crit.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionInstanceTemplate].
func (crit *ComputeRegionInstanceTemplate) Attributes() computeRegionInstanceTemplateAttributes {
	return computeRegionInstanceTemplateAttributes{ref: terra.ReferenceResource(crit)}
}

// ImportState imports the given attribute values into [ComputeRegionInstanceTemplate]'s state.
func (crit *ComputeRegionInstanceTemplate) ImportState(av io.Reader) error {
	crit.state = &computeRegionInstanceTemplateState{}
	if err := json.NewDecoder(av).Decode(crit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crit.Type(), crit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionInstanceTemplate] has state.
func (crit *ComputeRegionInstanceTemplate) State() (*computeRegionInstanceTemplateState, bool) {
	return crit.state, crit.state != nil
}

// StateMust returns the state for [ComputeRegionInstanceTemplate]. Panics if the state is nil.
func (crit *ComputeRegionInstanceTemplate) StateMust() *computeRegionInstanceTemplateState {
	if crit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crit.Type(), crit.LocalName()))
	}
	return crit.state
}

// ComputeRegionInstanceTemplateArgs contains the configurations for google_compute_region_instance_template.
type ComputeRegionInstanceTemplateArgs struct {
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
	// ResourcePolicies: list of string, optional
	ResourcePolicies terra.ListValue[terra.StringValue] `hcl:"resource_policies,attr"`
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// AdvancedMachineFeatures: optional
	AdvancedMachineFeatures *computeregioninstancetemplate.AdvancedMachineFeatures `hcl:"advanced_machine_features,block"`
	// ConfidentialInstanceConfig: optional
	ConfidentialInstanceConfig *computeregioninstancetemplate.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block"`
	// Disk: min=1
	Disk []computeregioninstancetemplate.Disk `hcl:"disk,block" validate:"min=1"`
	// GuestAccelerator: min=0
	GuestAccelerator []computeregioninstancetemplate.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []computeregioninstancetemplate.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// NetworkPerformanceConfig: optional
	NetworkPerformanceConfig *computeregioninstancetemplate.NetworkPerformanceConfig `hcl:"network_performance_config,block"`
	// ReservationAffinity: optional
	ReservationAffinity *computeregioninstancetemplate.ReservationAffinity `hcl:"reservation_affinity,block"`
	// Scheduling: optional
	Scheduling *computeregioninstancetemplate.Scheduling `hcl:"scheduling,block"`
	// ServiceAccount: optional
	ServiceAccount *computeregioninstancetemplate.ServiceAccount `hcl:"service_account,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *computeregioninstancetemplate.ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
	// Timeouts: optional
	Timeouts *computeregioninstancetemplate.Timeouts `hcl:"timeouts,block"`
}
type computeRegionInstanceTemplateAttributes struct {
	ref terra.Reference
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(crit.ref.Append("can_ip_forward"))
}

// Description returns a reference to field description of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("description"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(crit.ref.Append("enable_display"))
}

// Id returns a reference to field id of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("id"))
}

// InstanceDescription returns a reference to field instance_description of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) InstanceDescription() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("instance_description"))
}

// Labels returns a reference to field labels of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crit.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crit.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("min_cpu_platform"))
}

// Name returns a reference to field name of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("name_prefix"))
}

// Project returns a reference to field project of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("region"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crit.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("self_link"))
}

// Tags returns a reference to field tags of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crit.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_region_instance_template.
func (crit computeRegionInstanceTemplateAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(crit.ref.Append("tags_fingerprint"))
}

func (crit computeRegionInstanceTemplateAttributes) AdvancedMachineFeatures() terra.ListValue[computeregioninstancetemplate.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.AdvancedMachineFeaturesAttributes](crit.ref.Append("advanced_machine_features"))
}

func (crit computeRegionInstanceTemplateAttributes) ConfidentialInstanceConfig() terra.ListValue[computeregioninstancetemplate.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.ConfidentialInstanceConfigAttributes](crit.ref.Append("confidential_instance_config"))
}

func (crit computeRegionInstanceTemplateAttributes) Disk() terra.ListValue[computeregioninstancetemplate.DiskAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.DiskAttributes](crit.ref.Append("disk"))
}

func (crit computeRegionInstanceTemplateAttributes) GuestAccelerator() terra.ListValue[computeregioninstancetemplate.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.GuestAcceleratorAttributes](crit.ref.Append("guest_accelerator"))
}

func (crit computeRegionInstanceTemplateAttributes) NetworkInterface() terra.ListValue[computeregioninstancetemplate.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.NetworkInterfaceAttributes](crit.ref.Append("network_interface"))
}

func (crit computeRegionInstanceTemplateAttributes) NetworkPerformanceConfig() terra.ListValue[computeregioninstancetemplate.NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.NetworkPerformanceConfigAttributes](crit.ref.Append("network_performance_config"))
}

func (crit computeRegionInstanceTemplateAttributes) ReservationAffinity() terra.ListValue[computeregioninstancetemplate.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.ReservationAffinityAttributes](crit.ref.Append("reservation_affinity"))
}

func (crit computeRegionInstanceTemplateAttributes) Scheduling() terra.ListValue[computeregioninstancetemplate.SchedulingAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.SchedulingAttributes](crit.ref.Append("scheduling"))
}

func (crit computeRegionInstanceTemplateAttributes) ServiceAccount() terra.ListValue[computeregioninstancetemplate.ServiceAccountAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.ServiceAccountAttributes](crit.ref.Append("service_account"))
}

func (crit computeRegionInstanceTemplateAttributes) ShieldedInstanceConfig() terra.ListValue[computeregioninstancetemplate.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeregioninstancetemplate.ShieldedInstanceConfigAttributes](crit.ref.Append("shielded_instance_config"))
}

func (crit computeRegionInstanceTemplateAttributes) Timeouts() computeregioninstancetemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregioninstancetemplate.TimeoutsAttributes](crit.ref.Append("timeouts"))
}

type computeRegionInstanceTemplateState struct {
	CanIpForward               bool                                                            `json:"can_ip_forward"`
	Description                string                                                          `json:"description"`
	EnableDisplay              bool                                                            `json:"enable_display"`
	Id                         string                                                          `json:"id"`
	InstanceDescription        string                                                          `json:"instance_description"`
	Labels                     map[string]string                                               `json:"labels"`
	MachineType                string                                                          `json:"machine_type"`
	Metadata                   map[string]string                                               `json:"metadata"`
	MetadataFingerprint        string                                                          `json:"metadata_fingerprint"`
	MetadataStartupScript      string                                                          `json:"metadata_startup_script"`
	MinCpuPlatform             string                                                          `json:"min_cpu_platform"`
	Name                       string                                                          `json:"name"`
	NamePrefix                 string                                                          `json:"name_prefix"`
	Project                    string                                                          `json:"project"`
	Region                     string                                                          `json:"region"`
	ResourcePolicies           []string                                                        `json:"resource_policies"`
	SelfLink                   string                                                          `json:"self_link"`
	Tags                       []string                                                        `json:"tags"`
	TagsFingerprint            string                                                          `json:"tags_fingerprint"`
	AdvancedMachineFeatures    []computeregioninstancetemplate.AdvancedMachineFeaturesState    `json:"advanced_machine_features"`
	ConfidentialInstanceConfig []computeregioninstancetemplate.ConfidentialInstanceConfigState `json:"confidential_instance_config"`
	Disk                       []computeregioninstancetemplate.DiskState                       `json:"disk"`
	GuestAccelerator           []computeregioninstancetemplate.GuestAcceleratorState           `json:"guest_accelerator"`
	NetworkInterface           []computeregioninstancetemplate.NetworkInterfaceState           `json:"network_interface"`
	NetworkPerformanceConfig   []computeregioninstancetemplate.NetworkPerformanceConfigState   `json:"network_performance_config"`
	ReservationAffinity        []computeregioninstancetemplate.ReservationAffinityState        `json:"reservation_affinity"`
	Scheduling                 []computeregioninstancetemplate.SchedulingState                 `json:"scheduling"`
	ServiceAccount             []computeregioninstancetemplate.ServiceAccountState             `json:"service_account"`
	ShieldedInstanceConfig     []computeregioninstancetemplate.ShieldedInstanceConfigState     `json:"shielded_instance_config"`
	Timeouts                   *computeregioninstancetemplate.TimeoutsState                    `json:"timeouts"`
}
