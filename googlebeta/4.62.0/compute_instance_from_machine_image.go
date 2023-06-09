// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeinstancefrommachineimage "github.com/golingon/terraproviders/googlebeta/4.62.0/computeinstancefrommachineimage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceFromMachineImage creates a new instance of [ComputeInstanceFromMachineImage].
func NewComputeInstanceFromMachineImage(name string, args ComputeInstanceFromMachineImageArgs) *ComputeInstanceFromMachineImage {
	return &ComputeInstanceFromMachineImage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceFromMachineImage)(nil)

// ComputeInstanceFromMachineImage represents the Terraform resource google_compute_instance_from_machine_image.
type ComputeInstanceFromMachineImage struct {
	Name      string
	Args      ComputeInstanceFromMachineImageArgs
	state     *computeInstanceFromMachineImageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceFromMachineImage].
func (cifmi *ComputeInstanceFromMachineImage) Type() string {
	return "google_compute_instance_from_machine_image"
}

// LocalName returns the local name for [ComputeInstanceFromMachineImage].
func (cifmi *ComputeInstanceFromMachineImage) LocalName() string {
	return cifmi.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceFromMachineImage].
func (cifmi *ComputeInstanceFromMachineImage) Configuration() interface{} {
	return cifmi.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceFromMachineImage].
func (cifmi *ComputeInstanceFromMachineImage) DependOn() terra.Reference {
	return terra.ReferenceResource(cifmi)
}

// Dependencies returns the list of resources [ComputeInstanceFromMachineImage] depends_on.
func (cifmi *ComputeInstanceFromMachineImage) Dependencies() terra.Dependencies {
	return cifmi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceFromMachineImage].
func (cifmi *ComputeInstanceFromMachineImage) LifecycleManagement() *terra.Lifecycle {
	return cifmi.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceFromMachineImage].
func (cifmi *ComputeInstanceFromMachineImage) Attributes() computeInstanceFromMachineImageAttributes {
	return computeInstanceFromMachineImageAttributes{ref: terra.ReferenceResource(cifmi)}
}

// ImportState imports the given attribute values into [ComputeInstanceFromMachineImage]'s state.
func (cifmi *ComputeInstanceFromMachineImage) ImportState(av io.Reader) error {
	cifmi.state = &computeInstanceFromMachineImageState{}
	if err := json.NewDecoder(av).Decode(cifmi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cifmi.Type(), cifmi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceFromMachineImage] has state.
func (cifmi *ComputeInstanceFromMachineImage) State() (*computeInstanceFromMachineImageState, bool) {
	return cifmi.state, cifmi.state != nil
}

// StateMust returns the state for [ComputeInstanceFromMachineImage]. Panics if the state is nil.
func (cifmi *ComputeInstanceFromMachineImage) StateMust() *computeInstanceFromMachineImageState {
	if cifmi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cifmi.Type(), cifmi.LocalName()))
	}
	return cifmi.state
}

// ComputeInstanceFromMachineImageArgs contains the configurations for google_compute_instance_from_machine_image.
type ComputeInstanceFromMachineImageArgs struct {
	// AllowStoppingForUpdate: bool, optional
	AllowStoppingForUpdate terra.BoolValue `hcl:"allow_stopping_for_update,attr"`
	// CanIpForward: bool, optional
	CanIpForward terra.BoolValue `hcl:"can_ip_forward,attr"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DesiredStatus: string, optional
	DesiredStatus terra.StringValue `hcl:"desired_status,attr"`
	// EnableDisplay: bool, optional
	EnableDisplay terra.BoolValue `hcl:"enable_display,attr"`
	// Hostname: string, optional
	Hostname terra.StringValue `hcl:"hostname,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MachineType: string, optional
	MachineType terra.StringValue `hcl:"machine_type,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// MetadataStartupScript: string, optional
	MetadataStartupScript terra.StringValue `hcl:"metadata_startup_script,attr"`
	// MinCpuPlatform: string, optional
	MinCpuPlatform terra.StringValue `hcl:"min_cpu_platform,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ResourcePolicies: list of string, optional
	ResourcePolicies terra.ListValue[terra.StringValue] `hcl:"resource_policies,attr"`
	// SourceMachineImage: string, required
	SourceMachineImage terra.StringValue `hcl:"source_machine_image,attr" validate:"required"`
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AttachedDisk: min=0
	AttachedDisk []computeinstancefrommachineimage.AttachedDisk `hcl:"attached_disk,block" validate:"min=0"`
	// BootDisk: min=0
	BootDisk []computeinstancefrommachineimage.BootDisk `hcl:"boot_disk,block" validate:"min=0"`
	// GuestAccelerator: min=0
	GuestAccelerator []computeinstancefrommachineimage.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// ScratchDisk: min=0
	ScratchDisk []computeinstancefrommachineimage.ScratchDisk `hcl:"scratch_disk,block" validate:"min=0"`
	// ServiceAccount: min=0
	ServiceAccount []computeinstancefrommachineimage.ServiceAccount `hcl:"service_account,block" validate:"min=0"`
	// AdvancedMachineFeatures: optional
	AdvancedMachineFeatures *computeinstancefrommachineimage.AdvancedMachineFeatures `hcl:"advanced_machine_features,block"`
	// ConfidentialInstanceConfig: optional
	ConfidentialInstanceConfig *computeinstancefrommachineimage.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block"`
	// NetworkInterface: min=0
	NetworkInterface []computeinstancefrommachineimage.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// NetworkPerformanceConfig: optional
	NetworkPerformanceConfig *computeinstancefrommachineimage.NetworkPerformanceConfig `hcl:"network_performance_config,block"`
	// ReservationAffinity: optional
	ReservationAffinity *computeinstancefrommachineimage.ReservationAffinity `hcl:"reservation_affinity,block"`
	// Scheduling: optional
	Scheduling *computeinstancefrommachineimage.Scheduling `hcl:"scheduling,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *computeinstancefrommachineimage.ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
	// Timeouts: optional
	Timeouts *computeinstancefrommachineimage.Timeouts `hcl:"timeouts,block"`
}
type computeInstanceFromMachineImageAttributes struct {
	ref terra.Reference
}

// AllowStoppingForUpdate returns a reference to field allow_stopping_for_update of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) AllowStoppingForUpdate() terra.BoolValue {
	return terra.ReferenceAsBool(cifmi.ref.Append("allow_stopping_for_update"))
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(cifmi.ref.Append("can_ip_forward"))
}

// CpuPlatform returns a reference to field cpu_platform of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) CpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("cpu_platform"))
}

// CurrentStatus returns a reference to field current_status of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) CurrentStatus() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("current_status"))
}

// DeletionProtection returns a reference to field deletion_protection of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(cifmi.ref.Append("deletion_protection"))
}

// Description returns a reference to field description of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("description"))
}

// DesiredStatus returns a reference to field desired_status of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) DesiredStatus() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("desired_status"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(cifmi.ref.Append("enable_display"))
}

// Hostname returns a reference to field hostname of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("hostname"))
}

// Id returns a reference to field id of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("instance_id"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cifmi.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cifmi.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("min_cpu_platform"))
}

// Name returns a reference to field name of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("project"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cifmi.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("self_link"))
}

// SourceMachineImage returns a reference to field source_machine_image of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) SourceMachineImage() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("source_machine_image"))
}

// Tags returns a reference to field tags of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cifmi.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("tags_fingerprint"))
}

// Zone returns a reference to field zone of google_compute_instance_from_machine_image.
func (cifmi computeInstanceFromMachineImageAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cifmi.ref.Append("zone"))
}

func (cifmi computeInstanceFromMachineImageAttributes) AttachedDisk() terra.ListValue[computeinstancefrommachineimage.AttachedDiskAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.AttachedDiskAttributes](cifmi.ref.Append("attached_disk"))
}

func (cifmi computeInstanceFromMachineImageAttributes) BootDisk() terra.ListValue[computeinstancefrommachineimage.BootDiskAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.BootDiskAttributes](cifmi.ref.Append("boot_disk"))
}

func (cifmi computeInstanceFromMachineImageAttributes) GuestAccelerator() terra.ListValue[computeinstancefrommachineimage.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.GuestAcceleratorAttributes](cifmi.ref.Append("guest_accelerator"))
}

func (cifmi computeInstanceFromMachineImageAttributes) ScratchDisk() terra.ListValue[computeinstancefrommachineimage.ScratchDiskAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.ScratchDiskAttributes](cifmi.ref.Append("scratch_disk"))
}

func (cifmi computeInstanceFromMachineImageAttributes) ServiceAccount() terra.ListValue[computeinstancefrommachineimage.ServiceAccountAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.ServiceAccountAttributes](cifmi.ref.Append("service_account"))
}

func (cifmi computeInstanceFromMachineImageAttributes) AdvancedMachineFeatures() terra.ListValue[computeinstancefrommachineimage.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.AdvancedMachineFeaturesAttributes](cifmi.ref.Append("advanced_machine_features"))
}

func (cifmi computeInstanceFromMachineImageAttributes) ConfidentialInstanceConfig() terra.ListValue[computeinstancefrommachineimage.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.ConfidentialInstanceConfigAttributes](cifmi.ref.Append("confidential_instance_config"))
}

func (cifmi computeInstanceFromMachineImageAttributes) NetworkInterface() terra.ListValue[computeinstancefrommachineimage.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.NetworkInterfaceAttributes](cifmi.ref.Append("network_interface"))
}

func (cifmi computeInstanceFromMachineImageAttributes) NetworkPerformanceConfig() terra.ListValue[computeinstancefrommachineimage.NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.NetworkPerformanceConfigAttributes](cifmi.ref.Append("network_performance_config"))
}

func (cifmi computeInstanceFromMachineImageAttributes) ReservationAffinity() terra.ListValue[computeinstancefrommachineimage.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.ReservationAffinityAttributes](cifmi.ref.Append("reservation_affinity"))
}

func (cifmi computeInstanceFromMachineImageAttributes) Scheduling() terra.ListValue[computeinstancefrommachineimage.SchedulingAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.SchedulingAttributes](cifmi.ref.Append("scheduling"))
}

func (cifmi computeInstanceFromMachineImageAttributes) ShieldedInstanceConfig() terra.ListValue[computeinstancefrommachineimage.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstancefrommachineimage.ShieldedInstanceConfigAttributes](cifmi.ref.Append("shielded_instance_config"))
}

func (cifmi computeInstanceFromMachineImageAttributes) Timeouts() computeinstancefrommachineimage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinstancefrommachineimage.TimeoutsAttributes](cifmi.ref.Append("timeouts"))
}

type computeInstanceFromMachineImageState struct {
	AllowStoppingForUpdate     bool                                                              `json:"allow_stopping_for_update"`
	CanIpForward               bool                                                              `json:"can_ip_forward"`
	CpuPlatform                string                                                            `json:"cpu_platform"`
	CurrentStatus              string                                                            `json:"current_status"`
	DeletionProtection         bool                                                              `json:"deletion_protection"`
	Description                string                                                            `json:"description"`
	DesiredStatus              string                                                            `json:"desired_status"`
	EnableDisplay              bool                                                              `json:"enable_display"`
	Hostname                   string                                                            `json:"hostname"`
	Id                         string                                                            `json:"id"`
	InstanceId                 string                                                            `json:"instance_id"`
	LabelFingerprint           string                                                            `json:"label_fingerprint"`
	Labels                     map[string]string                                                 `json:"labels"`
	MachineType                string                                                            `json:"machine_type"`
	Metadata                   map[string]string                                                 `json:"metadata"`
	MetadataFingerprint        string                                                            `json:"metadata_fingerprint"`
	MetadataStartupScript      string                                                            `json:"metadata_startup_script"`
	MinCpuPlatform             string                                                            `json:"min_cpu_platform"`
	Name                       string                                                            `json:"name"`
	Project                    string                                                            `json:"project"`
	ResourcePolicies           []string                                                          `json:"resource_policies"`
	SelfLink                   string                                                            `json:"self_link"`
	SourceMachineImage         string                                                            `json:"source_machine_image"`
	Tags                       []string                                                          `json:"tags"`
	TagsFingerprint            string                                                            `json:"tags_fingerprint"`
	Zone                       string                                                            `json:"zone"`
	AttachedDisk               []computeinstancefrommachineimage.AttachedDiskState               `json:"attached_disk"`
	BootDisk                   []computeinstancefrommachineimage.BootDiskState                   `json:"boot_disk"`
	GuestAccelerator           []computeinstancefrommachineimage.GuestAcceleratorState           `json:"guest_accelerator"`
	ScratchDisk                []computeinstancefrommachineimage.ScratchDiskState                `json:"scratch_disk"`
	ServiceAccount             []computeinstancefrommachineimage.ServiceAccountState             `json:"service_account"`
	AdvancedMachineFeatures    []computeinstancefrommachineimage.AdvancedMachineFeaturesState    `json:"advanced_machine_features"`
	ConfidentialInstanceConfig []computeinstancefrommachineimage.ConfidentialInstanceConfigState `json:"confidential_instance_config"`
	NetworkInterface           []computeinstancefrommachineimage.NetworkInterfaceState           `json:"network_interface"`
	NetworkPerformanceConfig   []computeinstancefrommachineimage.NetworkPerformanceConfigState   `json:"network_performance_config"`
	ReservationAffinity        []computeinstancefrommachineimage.ReservationAffinityState        `json:"reservation_affinity"`
	Scheduling                 []computeinstancefrommachineimage.SchedulingState                 `json:"scheduling"`
	ShieldedInstanceConfig     []computeinstancefrommachineimage.ShieldedInstanceConfigState     `json:"shielded_instance_config"`
	Timeouts                   *computeinstancefrommachineimage.TimeoutsState                    `json:"timeouts"`
}
