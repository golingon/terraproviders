// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	notebooksinstance "github.com/golingon/terraproviders/google/4.62.0/notebooksinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksInstance creates a new instance of [NotebooksInstance].
func NewNotebooksInstance(name string, args NotebooksInstanceArgs) *NotebooksInstance {
	return &NotebooksInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksInstance)(nil)

// NotebooksInstance represents the Terraform resource google_notebooks_instance.
type NotebooksInstance struct {
	Name      string
	Args      NotebooksInstanceArgs
	state     *notebooksInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksInstance].
func (ni *NotebooksInstance) Type() string {
	return "google_notebooks_instance"
}

// LocalName returns the local name for [NotebooksInstance].
func (ni *NotebooksInstance) LocalName() string {
	return ni.Name
}

// Configuration returns the configuration (args) for [NotebooksInstance].
func (ni *NotebooksInstance) Configuration() interface{} {
	return ni.Args
}

// DependOn is used for other resources to depend on [NotebooksInstance].
func (ni *NotebooksInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ni)
}

// Dependencies returns the list of resources [NotebooksInstance] depends_on.
func (ni *NotebooksInstance) Dependencies() terra.Dependencies {
	return ni.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksInstance].
func (ni *NotebooksInstance) LifecycleManagement() *terra.Lifecycle {
	return ni.Lifecycle
}

// Attributes returns the attributes for [NotebooksInstance].
func (ni *NotebooksInstance) Attributes() notebooksInstanceAttributes {
	return notebooksInstanceAttributes{ref: terra.ReferenceResource(ni)}
}

// ImportState imports the given attribute values into [NotebooksInstance]'s state.
func (ni *NotebooksInstance) ImportState(av io.Reader) error {
	ni.state = &notebooksInstanceState{}
	if err := json.NewDecoder(av).Decode(ni.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ni.Type(), ni.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksInstance] has state.
func (ni *NotebooksInstance) State() (*notebooksInstanceState, bool) {
	return ni.state, ni.state != nil
}

// StateMust returns the state for [NotebooksInstance]. Panics if the state is nil.
func (ni *NotebooksInstance) StateMust() *notebooksInstanceState {
	if ni.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ni.Type(), ni.LocalName()))
	}
	return ni.state
}

// NotebooksInstanceArgs contains the configurations for google_notebooks_instance.
type NotebooksInstanceArgs struct {
	// BootDiskSizeGb: number, optional
	BootDiskSizeGb terra.NumberValue `hcl:"boot_disk_size_gb,attr"`
	// BootDiskType: string, optional
	BootDiskType terra.StringValue `hcl:"boot_disk_type,attr"`
	// CreateTime: string, optional
	CreateTime terra.StringValue `hcl:"create_time,attr"`
	// CustomGpuDriverPath: string, optional
	CustomGpuDriverPath terra.StringValue `hcl:"custom_gpu_driver_path,attr"`
	// DataDiskSizeGb: number, optional
	DataDiskSizeGb terra.NumberValue `hcl:"data_disk_size_gb,attr"`
	// DataDiskType: string, optional
	DataDiskType terra.StringValue `hcl:"data_disk_type,attr"`
	// DiskEncryption: string, optional
	DiskEncryption terra.StringValue `hcl:"disk_encryption,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstallGpuDriver: bool, optional
	InstallGpuDriver terra.BoolValue `hcl:"install_gpu_driver,attr"`
	// InstanceOwners: list of string, optional
	InstanceOwners terra.ListValue[terra.StringValue] `hcl:"instance_owners,attr"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MachineType: string, required
	MachineType terra.StringValue `hcl:"machine_type,attr" validate:"required"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NicType: string, optional
	NicType terra.StringValue `hcl:"nic_type,attr"`
	// NoProxyAccess: bool, optional
	NoProxyAccess terra.BoolValue `hcl:"no_proxy_access,attr"`
	// NoPublicIp: bool, optional
	NoPublicIp terra.BoolValue `hcl:"no_public_ip,attr"`
	// NoRemoveDataDisk: bool, optional
	NoRemoveDataDisk terra.BoolValue `hcl:"no_remove_data_disk,attr"`
	// PostStartupScript: string, optional
	PostStartupScript terra.StringValue `hcl:"post_startup_script,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// ServiceAccountScopes: list of string, optional
	ServiceAccountScopes terra.ListValue[terra.StringValue] `hcl:"service_account_scopes,attr"`
	// Subnet: string, optional
	Subnet terra.StringValue `hcl:"subnet,attr"`
	// Tags: list of string, optional
	Tags terra.ListValue[terra.StringValue] `hcl:"tags,attr"`
	// UpdateTime: string, optional
	UpdateTime terra.StringValue `hcl:"update_time,attr"`
	// AcceleratorConfig: optional
	AcceleratorConfig *notebooksinstance.AcceleratorConfig `hcl:"accelerator_config,block"`
	// ContainerImage: optional
	ContainerImage *notebooksinstance.ContainerImage `hcl:"container_image,block"`
	// ReservationAffinity: optional
	ReservationAffinity *notebooksinstance.ReservationAffinity `hcl:"reservation_affinity,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *notebooksinstance.ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
	// Timeouts: optional
	Timeouts *notebooksinstance.Timeouts `hcl:"timeouts,block"`
	// VmImage: optional
	VmImage *notebooksinstance.VmImage `hcl:"vm_image,block"`
}
type notebooksInstanceAttributes struct {
	ref terra.Reference
}

// BootDiskSizeGb returns a reference to field boot_disk_size_gb of google_notebooks_instance.
func (ni notebooksInstanceAttributes) BootDiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("boot_disk_size_gb"))
}

// BootDiskType returns a reference to field boot_disk_type of google_notebooks_instance.
func (ni notebooksInstanceAttributes) BootDiskType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("boot_disk_type"))
}

// CreateTime returns a reference to field create_time of google_notebooks_instance.
func (ni notebooksInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("create_time"))
}

// CustomGpuDriverPath returns a reference to field custom_gpu_driver_path of google_notebooks_instance.
func (ni notebooksInstanceAttributes) CustomGpuDriverPath() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("custom_gpu_driver_path"))
}

// DataDiskSizeGb returns a reference to field data_disk_size_gb of google_notebooks_instance.
func (ni notebooksInstanceAttributes) DataDiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("data_disk_size_gb"))
}

// DataDiskType returns a reference to field data_disk_type of google_notebooks_instance.
func (ni notebooksInstanceAttributes) DataDiskType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("data_disk_type"))
}

// DiskEncryption returns a reference to field disk_encryption of google_notebooks_instance.
func (ni notebooksInstanceAttributes) DiskEncryption() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("disk_encryption"))
}

// Id returns a reference to field id of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("id"))
}

// InstallGpuDriver returns a reference to field install_gpu_driver of google_notebooks_instance.
func (ni notebooksInstanceAttributes) InstallGpuDriver() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("install_gpu_driver"))
}

// InstanceOwners returns a reference to field instance_owners of google_notebooks_instance.
func (ni notebooksInstanceAttributes) InstanceOwners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("instance_owners"))
}

// KmsKey returns a reference to field kms_key of google_notebooks_instance.
func (ni notebooksInstanceAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("kms_key"))
}

// Labels returns a reference to field labels of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("labels"))
}

// Location returns a reference to field location of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("location"))
}

// MachineType returns a reference to field machine_type of google_notebooks_instance.
func (ni notebooksInstanceAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("machine_type"))
}

// Metadata returns a reference to field metadata of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("metadata"))
}

// Name returns a reference to field name of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("name"))
}

// Network returns a reference to field network of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network"))
}

// NicType returns a reference to field nic_type of google_notebooks_instance.
func (ni notebooksInstanceAttributes) NicType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("nic_type"))
}

// NoProxyAccess returns a reference to field no_proxy_access of google_notebooks_instance.
func (ni notebooksInstanceAttributes) NoProxyAccess() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("no_proxy_access"))
}

// NoPublicIp returns a reference to field no_public_ip of google_notebooks_instance.
func (ni notebooksInstanceAttributes) NoPublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("no_public_ip"))
}

// NoRemoveDataDisk returns a reference to field no_remove_data_disk of google_notebooks_instance.
func (ni notebooksInstanceAttributes) NoRemoveDataDisk() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("no_remove_data_disk"))
}

// PostStartupScript returns a reference to field post_startup_script of google_notebooks_instance.
func (ni notebooksInstanceAttributes) PostStartupScript() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("post_startup_script"))
}

// Project returns a reference to field project of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("project"))
}

// ProxyUri returns a reference to field proxy_uri of google_notebooks_instance.
func (ni notebooksInstanceAttributes) ProxyUri() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("proxy_uri"))
}

// ServiceAccount returns a reference to field service_account of google_notebooks_instance.
func (ni notebooksInstanceAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("service_account"))
}

// ServiceAccountScopes returns a reference to field service_account_scopes of google_notebooks_instance.
func (ni notebooksInstanceAttributes) ServiceAccountScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("service_account_scopes"))
}

// State returns a reference to field state of google_notebooks_instance.
func (ni notebooksInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("state"))
}

// Subnet returns a reference to field subnet of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Subnet() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet"))
}

// Tags returns a reference to field tags of google_notebooks_instance.
func (ni notebooksInstanceAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("tags"))
}

// UpdateTime returns a reference to field update_time of google_notebooks_instance.
func (ni notebooksInstanceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("update_time"))
}

func (ni notebooksInstanceAttributes) AcceleratorConfig() terra.ListValue[notebooksinstance.AcceleratorConfigAttributes] {
	return terra.ReferenceAsList[notebooksinstance.AcceleratorConfigAttributes](ni.ref.Append("accelerator_config"))
}

func (ni notebooksInstanceAttributes) ContainerImage() terra.ListValue[notebooksinstance.ContainerImageAttributes] {
	return terra.ReferenceAsList[notebooksinstance.ContainerImageAttributes](ni.ref.Append("container_image"))
}

func (ni notebooksInstanceAttributes) ReservationAffinity() terra.ListValue[notebooksinstance.ReservationAffinityAttributes] {
	return terra.ReferenceAsList[notebooksinstance.ReservationAffinityAttributes](ni.ref.Append("reservation_affinity"))
}

func (ni notebooksInstanceAttributes) ShieldedInstanceConfig() terra.ListValue[notebooksinstance.ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[notebooksinstance.ShieldedInstanceConfigAttributes](ni.ref.Append("shielded_instance_config"))
}

func (ni notebooksInstanceAttributes) Timeouts() notebooksinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[notebooksinstance.TimeoutsAttributes](ni.ref.Append("timeouts"))
}

func (ni notebooksInstanceAttributes) VmImage() terra.ListValue[notebooksinstance.VmImageAttributes] {
	return terra.ReferenceAsList[notebooksinstance.VmImageAttributes](ni.ref.Append("vm_image"))
}

type notebooksInstanceState struct {
	BootDiskSizeGb         float64                                         `json:"boot_disk_size_gb"`
	BootDiskType           string                                          `json:"boot_disk_type"`
	CreateTime             string                                          `json:"create_time"`
	CustomGpuDriverPath    string                                          `json:"custom_gpu_driver_path"`
	DataDiskSizeGb         float64                                         `json:"data_disk_size_gb"`
	DataDiskType           string                                          `json:"data_disk_type"`
	DiskEncryption         string                                          `json:"disk_encryption"`
	Id                     string                                          `json:"id"`
	InstallGpuDriver       bool                                            `json:"install_gpu_driver"`
	InstanceOwners         []string                                        `json:"instance_owners"`
	KmsKey                 string                                          `json:"kms_key"`
	Labels                 map[string]string                               `json:"labels"`
	Location               string                                          `json:"location"`
	MachineType            string                                          `json:"machine_type"`
	Metadata               map[string]string                               `json:"metadata"`
	Name                   string                                          `json:"name"`
	Network                string                                          `json:"network"`
	NicType                string                                          `json:"nic_type"`
	NoProxyAccess          bool                                            `json:"no_proxy_access"`
	NoPublicIp             bool                                            `json:"no_public_ip"`
	NoRemoveDataDisk       bool                                            `json:"no_remove_data_disk"`
	PostStartupScript      string                                          `json:"post_startup_script"`
	Project                string                                          `json:"project"`
	ProxyUri               string                                          `json:"proxy_uri"`
	ServiceAccount         string                                          `json:"service_account"`
	ServiceAccountScopes   []string                                        `json:"service_account_scopes"`
	State                  string                                          `json:"state"`
	Subnet                 string                                          `json:"subnet"`
	Tags                   []string                                        `json:"tags"`
	UpdateTime             string                                          `json:"update_time"`
	AcceleratorConfig      []notebooksinstance.AcceleratorConfigState      `json:"accelerator_config"`
	ContainerImage         []notebooksinstance.ContainerImageState         `json:"container_image"`
	ReservationAffinity    []notebooksinstance.ReservationAffinityState    `json:"reservation_affinity"`
	ShieldedInstanceConfig []notebooksinstance.ShieldedInstanceConfigState `json:"shielded_instance_config"`
	Timeouts               *notebooksinstance.TimeoutsState                `json:"timeouts"`
	VmImage                []notebooksinstance.VmImageState                `json:"vm_image"`
}
