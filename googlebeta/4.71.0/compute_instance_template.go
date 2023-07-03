// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeinstancetemplate "github.com/golingon/terraproviders/googlebeta/4.71.0/computeinstancetemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceTemplate creates a new instance of [ComputeInstanceTemplate].
func NewComputeInstanceTemplate(name string, args ComputeInstanceTemplateArgs) *ComputeInstanceTemplate {
	return &ComputeInstanceTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceTemplate)(nil)

// ComputeInstanceTemplate represents the Terraform resource google_compute_instance_template.
type ComputeInstanceTemplate struct {
	Name      string
	Args      ComputeInstanceTemplateArgs
	state     *computeInstanceTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceTemplate].
func (cit *ComputeInstanceTemplate) Type() string {
	return "google_compute_instance_template"
}

// LocalName returns the local name for [ComputeInstanceTemplate].
func (cit *ComputeInstanceTemplate) LocalName() string {
	return cit.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceTemplate].
func (cit *ComputeInstanceTemplate) Configuration() interface{} {
	return cit.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceTemplate].
func (cit *ComputeInstanceTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(cit)
}

// Dependencies returns the list of resources [ComputeInstanceTemplate] depends_on.
func (cit *ComputeInstanceTemplate) Dependencies() terra.Dependencies {
	return cit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceTemplate].
func (cit *ComputeInstanceTemplate) LifecycleManagement() *terra.Lifecycle {
	return cit.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceTemplate].
func (cit *ComputeInstanceTemplate) Attributes() computeInstanceTemplateAttributes {
	return computeInstanceTemplateAttributes{ref: terra.ReferenceResource(cit)}
}

// ImportState imports the given attribute values into [ComputeInstanceTemplate]'s state.
func (cit *ComputeInstanceTemplate) ImportState(av io.Reader) error {
	cit.state = &computeInstanceTemplateState{}
	if err := json.NewDecoder(av).Decode(cit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cit.Type(), cit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceTemplate] has state.
func (cit *ComputeInstanceTemplate) State() (*computeInstanceTemplateState, bool) {
	return cit.state, cit.state != nil
}

// StateMust returns the state for [ComputeInstanceTemplate]. Panics if the state is nil.
func (cit *ComputeInstanceTemplate) StateMust() *computeInstanceTemplateState {
	if cit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cit.Type(), cit.LocalName()))
	}
	return cit.state
}

// ComputeInstanceTemplateArgs contains the configurations for google_compute_instance_template.
type ComputeInstanceTemplateArgs struct {
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
	AdvancedMachineFeatures *computeinstancetemplate.AdvancedMachineFeatures `hcl:"advanced_machine_features,block"`
	// ConfidentialInstanceConfig: optional
	ConfidentialInstanceConfig *computeinstancetemplate.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block"`
	// Disk: min=1
	Disk []computeinstancetemplate.Disk `hcl:"disk,block" validate:"min=1"`
	// GuestAccelerator: min=0
	GuestAccelerator []computeinstancetemplate.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []computeinstancetemplate.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// NetworkPerformanceConfig: optional
	NetworkPerformanceConfig *computeinstancetemplate.NetworkPerformanceConfig `hcl:"network_performance_config,block"`
	// ReservationAffinity: optional
	ReservationAffinity *computeinstancetemplate.ReservationAffinity `hcl:"reservation_affinity,block"`
	// Scheduling: optional
	Scheduling *computeinstancetemplate.Scheduling `hcl:"scheduling,block"`
	// ServiceAccount: optional
	ServiceAccount *computeinstancetemplate.ServiceAccount `hcl:"service_account,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *computeinstancetemplate.ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
	// Timeouts: optional
	Timeouts *computeinstancetemplate.Timeouts `hcl:"timeouts,block"`
}
type computeInstanceTemplateAttributes struct {
	ref terra.Reference
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(cit.ref.Append("can_ip_forward"))
}

// Description returns a reference to field description of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("description"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(cit.ref.Append("enable_display"))
}

// Id returns a reference to field id of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("id"))
}

// InstanceDescription returns a reference to field instance_description of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) InstanceDescription() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("instance_description"))
}

// Labels returns a reference to field labels of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cit.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cit.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("min_cpu_platform"))
}

// Name returns a reference to field name of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("name_prefix"))
}

// Project returns a reference to field project of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("region"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cit.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("self_link"))
}

// SelfLinkUnique returns a reference to field self_link_unique of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) SelfLinkUnique() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("self_link_unique"))
}

// Tags returns a reference to field tags of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cit.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_instance_template.
func (cit computeInstanceTemplateAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("tags_fingerprint"))
}

func (cit computeInstanceTemplateAttributes) AdvancedMachineFeatures() terra.ListValue[computeinstancetemplate.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.AdvancedMachineFeaturesAttributes](cit.ref.Append("advanced_machine_features"))
}

func (cit computeInstanceTemplateAttributes) ConfidentialInstanceConfig() terra.ListValue[computeinstancetemplate.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.ConfidentialInstanceConfigAttributes](cit.ref.Append("confidential_instance_config"))
}

func (cit computeInstanceTemplateAttributes) Disk() terra.ListValue[computeinstancetemplate.DiskAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.DiskAttributes](cit.ref.Append("disk"))
}

func (cit computeInstanceTemplateAttributes) GuestAccelerator() terra.ListValue[computeinstancetemplate.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.GuestAcceleratorAttributes](cit.ref.Append("guest_accelerator"))
}

func (cit computeInstanceTemplateAttributes) NetworkInterface() terra.ListValue[computeinstancetemplate.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.NetworkInterfaceAttributes](cit.ref.Append("network_interface"))
}

func (cit computeInstanceTemplateAttributes) NetworkPerformanceConfig() terra.ListValue[computeinstancetemplate.NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.NetworkPerformanceConfigAttributes](cit.ref.Append("network_performance_config"))
}

func (cit computeInstanceTemplateAttributes) ReservationAffinity() terra.ListValue[computeinstancetemplate.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.ReservationAffinityAttributes](cit.ref.Append("reservation_affinity"))
}

func (cit computeInstanceTemplateAttributes) Scheduling() terra.ListValue[computeinstancetemplate.SchedulingAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.SchedulingAttributes](cit.ref.Append("scheduling"))
}

func (cit computeInstanceTemplateAttributes) ServiceAccount() terra.ListValue[computeinstancetemplate.ServiceAccountAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.ServiceAccountAttributes](cit.ref.Append("service_account"))
}

func (cit computeInstanceTemplateAttributes) ShieldedInstanceConfig() terra.ListValue[computeinstancetemplate.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstancetemplate.ShieldedInstanceConfigAttributes](cit.ref.Append("shielded_instance_config"))
}

func (cit computeInstanceTemplateAttributes) Timeouts() computeinstancetemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinstancetemplate.TimeoutsAttributes](cit.ref.Append("timeouts"))
}

type computeInstanceTemplateState struct {
	CanIpForward               bool                                                      `json:"can_ip_forward"`
	Description                string                                                    `json:"description"`
	EnableDisplay              bool                                                      `json:"enable_display"`
	Id                         string                                                    `json:"id"`
	InstanceDescription        string                                                    `json:"instance_description"`
	Labels                     map[string]string                                         `json:"labels"`
	MachineType                string                                                    `json:"machine_type"`
	Metadata                   map[string]string                                         `json:"metadata"`
	MetadataFingerprint        string                                                    `json:"metadata_fingerprint"`
	MetadataStartupScript      string                                                    `json:"metadata_startup_script"`
	MinCpuPlatform             string                                                    `json:"min_cpu_platform"`
	Name                       string                                                    `json:"name"`
	NamePrefix                 string                                                    `json:"name_prefix"`
	Project                    string                                                    `json:"project"`
	Region                     string                                                    `json:"region"`
	ResourcePolicies           []string                                                  `json:"resource_policies"`
	SelfLink                   string                                                    `json:"self_link"`
	SelfLinkUnique             string                                                    `json:"self_link_unique"`
	Tags                       []string                                                  `json:"tags"`
	TagsFingerprint            string                                                    `json:"tags_fingerprint"`
	AdvancedMachineFeatures    []computeinstancetemplate.AdvancedMachineFeaturesState    `json:"advanced_machine_features"`
	ConfidentialInstanceConfig []computeinstancetemplate.ConfidentialInstanceConfigState `json:"confidential_instance_config"`
	Disk                       []computeinstancetemplate.DiskState                       `json:"disk"`
	GuestAccelerator           []computeinstancetemplate.GuestAcceleratorState           `json:"guest_accelerator"`
	NetworkInterface           []computeinstancetemplate.NetworkInterfaceState           `json:"network_interface"`
	NetworkPerformanceConfig   []computeinstancetemplate.NetworkPerformanceConfigState   `json:"network_performance_config"`
	ReservationAffinity        []computeinstancetemplate.ReservationAffinityState        `json:"reservation_affinity"`
	Scheduling                 []computeinstancetemplate.SchedulingState                 `json:"scheduling"`
	ServiceAccount             []computeinstancetemplate.ServiceAccountState             `json:"service_account"`
	ShieldedInstanceConfig     []computeinstancetemplate.ShieldedInstanceConfigState     `json:"shielded_instance_config"`
	Timeouts                   *computeinstancetemplate.TimeoutsState                    `json:"timeouts"`
}
