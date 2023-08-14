// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinstance "github.com/golingon/terraproviders/google/4.77.0/computeinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstance creates a new instance of [ComputeInstance].
func NewComputeInstance(name string, args ComputeInstanceArgs) *ComputeInstance {
	return &ComputeInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstance)(nil)

// ComputeInstance represents the Terraform resource google_compute_instance.
type ComputeInstance struct {
	Name      string
	Args      ComputeInstanceArgs
	state     *computeInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstance].
func (ci *ComputeInstance) Type() string {
	return "google_compute_instance"
}

// LocalName returns the local name for [ComputeInstance].
func (ci *ComputeInstance) LocalName() string {
	return ci.Name
}

// Configuration returns the configuration (args) for [ComputeInstance].
func (ci *ComputeInstance) Configuration() interface{} {
	return ci.Args
}

// DependOn is used for other resources to depend on [ComputeInstance].
func (ci *ComputeInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ci)
}

// Dependencies returns the list of resources [ComputeInstance] depends_on.
func (ci *ComputeInstance) Dependencies() terra.Dependencies {
	return ci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstance].
func (ci *ComputeInstance) LifecycleManagement() *terra.Lifecycle {
	return ci.Lifecycle
}

// Attributes returns the attributes for [ComputeInstance].
func (ci *ComputeInstance) Attributes() computeInstanceAttributes {
	return computeInstanceAttributes{ref: terra.ReferenceResource(ci)}
}

// ImportState imports the given attribute values into [ComputeInstance]'s state.
func (ci *ComputeInstance) ImportState(av io.Reader) error {
	ci.state = &computeInstanceState{}
	if err := json.NewDecoder(av).Decode(ci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ci.Type(), ci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstance] has state.
func (ci *ComputeInstance) State() (*computeInstanceState, bool) {
	return ci.state, ci.state != nil
}

// StateMust returns the state for [ComputeInstance]. Panics if the state is nil.
func (ci *ComputeInstance) StateMust() *computeInstanceState {
	if ci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ci.Type(), ci.LocalName()))
	}
	return ci.state
}

// ComputeInstanceArgs contains the configurations for google_compute_instance.
type ComputeInstanceArgs struct {
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
	// MachineType: string, required
	MachineType terra.StringValue `hcl:"machine_type,attr" validate:"required"`
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
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// GuestAccelerator: min=0
	GuestAccelerator []computeinstance.GuestAccelerator `hcl:"guest_accelerator,block" validate:"min=0"`
	// AdvancedMachineFeatures: optional
	AdvancedMachineFeatures *computeinstance.AdvancedMachineFeatures `hcl:"advanced_machine_features,block"`
	// AttachedDisk: min=0
	AttachedDisk []computeinstance.AttachedDisk `hcl:"attached_disk,block" validate:"min=0"`
	// BootDisk: required
	BootDisk *computeinstance.BootDisk `hcl:"boot_disk,block" validate:"required"`
	// ConfidentialInstanceConfig: optional
	ConfidentialInstanceConfig *computeinstance.ConfidentialInstanceConfig `hcl:"confidential_instance_config,block"`
	// NetworkInterface: min=1
	NetworkInterface []computeinstance.NetworkInterface `hcl:"network_interface,block" validate:"min=1"`
	// NetworkPerformanceConfig: optional
	NetworkPerformanceConfig *computeinstance.NetworkPerformanceConfig `hcl:"network_performance_config,block"`
	// Params: optional
	Params *computeinstance.Params `hcl:"params,block"`
	// ReservationAffinity: optional
	ReservationAffinity *computeinstance.ReservationAffinity `hcl:"reservation_affinity,block"`
	// Scheduling: optional
	Scheduling *computeinstance.Scheduling `hcl:"scheduling,block"`
	// ScratchDisk: min=0
	ScratchDisk []computeinstance.ScratchDisk `hcl:"scratch_disk,block" validate:"min=0"`
	// ServiceAccount: optional
	ServiceAccount *computeinstance.ServiceAccount `hcl:"service_account,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *computeinstance.ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
	// Timeouts: optional
	Timeouts *computeinstance.Timeouts `hcl:"timeouts,block"`
}
type computeInstanceAttributes struct {
	ref terra.Reference
}

// AllowStoppingForUpdate returns a reference to field allow_stopping_for_update of google_compute_instance.
func (ci computeInstanceAttributes) AllowStoppingForUpdate() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("allow_stopping_for_update"))
}

// CanIpForward returns a reference to field can_ip_forward of google_compute_instance.
func (ci computeInstanceAttributes) CanIpForward() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("can_ip_forward"))
}

// CpuPlatform returns a reference to field cpu_platform of google_compute_instance.
func (ci computeInstanceAttributes) CpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("cpu_platform"))
}

// CurrentStatus returns a reference to field current_status of google_compute_instance.
func (ci computeInstanceAttributes) CurrentStatus() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("current_status"))
}

// DeletionProtection returns a reference to field deletion_protection of google_compute_instance.
func (ci computeInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("deletion_protection"))
}

// Description returns a reference to field description of google_compute_instance.
func (ci computeInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("description"))
}

// DesiredStatus returns a reference to field desired_status of google_compute_instance.
func (ci computeInstanceAttributes) DesiredStatus() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("desired_status"))
}

// EnableDisplay returns a reference to field enable_display of google_compute_instance.
func (ci computeInstanceAttributes) EnableDisplay() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("enable_display"))
}

// Hostname returns a reference to field hostname of google_compute_instance.
func (ci computeInstanceAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("hostname"))
}

// Id returns a reference to field id of google_compute_instance.
func (ci computeInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_compute_instance.
func (ci computeInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("instance_id"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_instance.
func (ci computeInstanceAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_instance.
func (ci computeInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("labels"))
}

// MachineType returns a reference to field machine_type of google_compute_instance.
func (ci computeInstanceAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_compute_instance.
func (ci computeInstanceAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("metadata"))
}

// MetadataFingerprint returns a reference to field metadata_fingerprint of google_compute_instance.
func (ci computeInstanceAttributes) MetadataFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("metadata_fingerprint"))
}

// MetadataStartupScript returns a reference to field metadata_startup_script of google_compute_instance.
func (ci computeInstanceAttributes) MetadataStartupScript() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("metadata_startup_script"))
}

// MinCpuPlatform returns a reference to field min_cpu_platform of google_compute_instance.
func (ci computeInstanceAttributes) MinCpuPlatform() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("min_cpu_platform"))
}

// Name returns a reference to field name of google_compute_instance.
func (ci computeInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_instance.
func (ci computeInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("project"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_instance.
func (ci computeInstanceAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_instance.
func (ci computeInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("self_link"))
}

// Tags returns a reference to field tags of google_compute_instance.
func (ci computeInstanceAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ci.ref.Append("tags"))
}

// TagsFingerprint returns a reference to field tags_fingerprint of google_compute_instance.
func (ci computeInstanceAttributes) TagsFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("tags_fingerprint"))
}

// Zone returns a reference to field zone of google_compute_instance.
func (ci computeInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("zone"))
}

func (ci computeInstanceAttributes) GuestAccelerator() terra.ListValue[computeinstance.GuestAcceleratorAttributes] {
	return terra.ReferenceAsList[computeinstance.GuestAcceleratorAttributes](ci.ref.Append("guest_accelerator"))
}

func (ci computeInstanceAttributes) AdvancedMachineFeatures() terra.ListValue[computeinstance.AdvancedMachineFeaturesAttributes] {
	return terra.ReferenceAsList[computeinstance.AdvancedMachineFeaturesAttributes](ci.ref.Append("advanced_machine_features"))
}

func (ci computeInstanceAttributes) AttachedDisk() terra.ListValue[computeinstance.AttachedDiskAttributes] {
	return terra.ReferenceAsList[computeinstance.AttachedDiskAttributes](ci.ref.Append("attached_disk"))
}

func (ci computeInstanceAttributes) BootDisk() terra.ListValue[computeinstance.BootDiskAttributes] {
	return terra.ReferenceAsList[computeinstance.BootDiskAttributes](ci.ref.Append("boot_disk"))
}

func (ci computeInstanceAttributes) ConfidentialInstanceConfig() terra.ListValue[computeinstance.ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstance.ConfidentialInstanceConfigAttributes](ci.ref.Append("confidential_instance_config"))
}

func (ci computeInstanceAttributes) NetworkInterface() terra.ListValue[computeinstance.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[computeinstance.NetworkInterfaceAttributes](ci.ref.Append("network_interface"))
}

func (ci computeInstanceAttributes) NetworkPerformanceConfig() terra.ListValue[computeinstance.NetworkPerformanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstance.NetworkPerformanceConfigAttributes](ci.ref.Append("network_performance_config"))
}

func (ci computeInstanceAttributes) Params() terra.ListValue[computeinstance.ParamsAttributes] {
	return terra.ReferenceAsList[computeinstance.ParamsAttributes](ci.ref.Append("params"))
}

func (ci computeInstanceAttributes) ReservationAffinity() terra.ListValue[computeinstance.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[computeinstance.ReservationAffinityAttributes](ci.ref.Append("reservation_affinity"))
}

func (ci computeInstanceAttributes) Scheduling() terra.ListValue[computeinstance.SchedulingAttributes] {
	return terra.ReferenceAsList[computeinstance.SchedulingAttributes](ci.ref.Append("scheduling"))
}

func (ci computeInstanceAttributes) ScratchDisk() terra.ListValue[computeinstance.ScratchDiskAttributes] {
	return terra.ReferenceAsList[computeinstance.ScratchDiskAttributes](ci.ref.Append("scratch_disk"))
}

func (ci computeInstanceAttributes) ServiceAccount() terra.ListValue[computeinstance.ServiceAccountAttributes] {
	return terra.ReferenceAsList[computeinstance.ServiceAccountAttributes](ci.ref.Append("service_account"))
}

func (ci computeInstanceAttributes) ShieldedInstanceConfig() terra.ListValue[computeinstance.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[computeinstance.ShieldedInstanceConfigAttributes](ci.ref.Append("shielded_instance_config"))
}

func (ci computeInstanceAttributes) Timeouts() computeinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinstance.TimeoutsAttributes](ci.ref.Append("timeouts"))
}

type computeInstanceState struct {
	AllowStoppingForUpdate     bool                                              `json:"allow_stopping_for_update"`
	CanIpForward               bool                                              `json:"can_ip_forward"`
	CpuPlatform                string                                            `json:"cpu_platform"`
	CurrentStatus              string                                            `json:"current_status"`
	DeletionProtection         bool                                              `json:"deletion_protection"`
	Description                string                                            `json:"description"`
	DesiredStatus              string                                            `json:"desired_status"`
	EnableDisplay              bool                                              `json:"enable_display"`
	Hostname                   string                                            `json:"hostname"`
	Id                         string                                            `json:"id"`
	InstanceId                 string                                            `json:"instance_id"`
	LabelFingerprint           string                                            `json:"label_fingerprint"`
	Labels                     map[string]string                                 `json:"labels"`
	MachineType                string                                            `json:"machine_type"`
	Metadata                   map[string]string                                 `json:"metadata"`
	MetadataFingerprint        string                                            `json:"metadata_fingerprint"`
	MetadataStartupScript      string                                            `json:"metadata_startup_script"`
	MinCpuPlatform             string                                            `json:"min_cpu_platform"`
	Name                       string                                            `json:"name"`
	Project                    string                                            `json:"project"`
	ResourcePolicies           []string                                          `json:"resource_policies"`
	SelfLink                   string                                            `json:"self_link"`
	Tags                       []string                                          `json:"tags"`
	TagsFingerprint            string                                            `json:"tags_fingerprint"`
	Zone                       string                                            `json:"zone"`
	GuestAccelerator           []computeinstance.GuestAcceleratorState           `json:"guest_accelerator"`
	AdvancedMachineFeatures    []computeinstance.AdvancedMachineFeaturesState    `json:"advanced_machine_features"`
	AttachedDisk               []computeinstance.AttachedDiskState               `json:"attached_disk"`
	BootDisk                   []computeinstance.BootDiskState                   `json:"boot_disk"`
	ConfidentialInstanceConfig []computeinstance.ConfidentialInstanceConfigState `json:"confidential_instance_config"`
	NetworkInterface           []computeinstance.NetworkInterfaceState           `json:"network_interface"`
	NetworkPerformanceConfig   []computeinstance.NetworkPerformanceConfigState   `json:"network_performance_config"`
	Params                     []computeinstance.ParamsState                     `json:"params"`
	ReservationAffinity        []computeinstance.ReservationAffinityState        `json:"reservation_affinity"`
	Scheduling                 []computeinstance.SchedulingState                 `json:"scheduling"`
	ScratchDisk                []computeinstance.ScratchDiskState                `json:"scratch_disk"`
	ServiceAccount             []computeinstance.ServiceAccountState             `json:"service_account"`
	ShieldedInstanceConfig     []computeinstance.ShieldedInstanceConfigState     `json:"shielded_instance_config"`
	Timeouts                   *computeinstance.TimeoutsState                    `json:"timeouts"`
}
