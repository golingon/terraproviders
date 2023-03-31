// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ami "github.com/golingon/terraproviders/aws/4.60.0/ami"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAmi creates a new instance of [Ami].
func NewAmi(name string, args AmiArgs) *Ami {
	return &Ami{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ami)(nil)

// Ami represents the Terraform resource aws_ami.
type Ami struct {
	Name      string
	Args      AmiArgs
	state     *amiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ami].
func (a *Ami) Type() string {
	return "aws_ami"
}

// LocalName returns the local name for [Ami].
func (a *Ami) LocalName() string {
	return a.Name
}

// Configuration returns the configuration (args) for [Ami].
func (a *Ami) Configuration() interface{} {
	return a.Args
}

// DependOn is used for other resources to depend on [Ami].
func (a *Ami) DependOn() terra.Reference {
	return terra.ReferenceResource(a)
}

// Dependencies returns the list of resources [Ami] depends_on.
func (a *Ami) Dependencies() terra.Dependencies {
	return a.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ami].
func (a *Ami) LifecycleManagement() *terra.Lifecycle {
	return a.Lifecycle
}

// Attributes returns the attributes for [Ami].
func (a *Ami) Attributes() amiAttributes {
	return amiAttributes{ref: terra.ReferenceResource(a)}
}

// ImportState imports the given attribute values into [Ami]'s state.
func (a *Ami) ImportState(av io.Reader) error {
	a.state = &amiState{}
	if err := json.NewDecoder(av).Decode(a.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", a.Type(), a.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ami] has state.
func (a *Ami) State() (*amiState, bool) {
	return a.state, a.state != nil
}

// StateMust returns the state for [Ami]. Panics if the state is nil.
func (a *Ami) StateMust() *amiState {
	if a.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", a.Type(), a.LocalName()))
	}
	return a.state
}

// AmiArgs contains the configurations for aws_ami.
type AmiArgs struct {
	// Architecture: string, optional
	Architecture terra.StringValue `hcl:"architecture,attr"`
	// BootMode: string, optional
	BootMode terra.StringValue `hcl:"boot_mode,attr"`
	// DeprecationTime: string, optional
	DeprecationTime terra.StringValue `hcl:"deprecation_time,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnaSupport: bool, optional
	EnaSupport terra.BoolValue `hcl:"ena_support,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageLocation: string, optional
	ImageLocation terra.StringValue `hcl:"image_location,attr"`
	// ImdsSupport: string, optional
	ImdsSupport terra.StringValue `hcl:"imds_support,attr"`
	// KernelId: string, optional
	KernelId terra.StringValue `hcl:"kernel_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RamdiskId: string, optional
	RamdiskId terra.StringValue `hcl:"ramdisk_id,attr"`
	// RootDeviceName: string, optional
	RootDeviceName terra.StringValue `hcl:"root_device_name,attr"`
	// SriovNetSupport: string, optional
	SriovNetSupport terra.StringValue `hcl:"sriov_net_support,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TpmSupport: string, optional
	TpmSupport terra.StringValue `hcl:"tpm_support,attr"`
	// VirtualizationType: string, optional
	VirtualizationType terra.StringValue `hcl:"virtualization_type,attr"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []ami.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []ami.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *ami.Timeouts `hcl:"timeouts,block"`
}
type amiAttributes struct {
	ref terra.Reference
}

// Architecture returns a reference to field architecture of aws_ami.
func (a amiAttributes) Architecture() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("architecture"))
}

// Arn returns a reference to field arn of aws_ami.
func (a amiAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn"))
}

// BootMode returns a reference to field boot_mode of aws_ami.
func (a amiAttributes) BootMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("boot_mode"))
}

// DeprecationTime returns a reference to field deprecation_time of aws_ami.
func (a amiAttributes) DeprecationTime() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("deprecation_time"))
}

// Description returns a reference to field description of aws_ami.
func (a amiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("description"))
}

// EnaSupport returns a reference to field ena_support of aws_ami.
func (a amiAttributes) EnaSupport() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("ena_support"))
}

// Hypervisor returns a reference to field hypervisor of aws_ami.
func (a amiAttributes) Hypervisor() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("hypervisor"))
}

// Id returns a reference to field id of aws_ami.
func (a amiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

// ImageLocation returns a reference to field image_location of aws_ami.
func (a amiAttributes) ImageLocation() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_location"))
}

// ImageOwnerAlias returns a reference to field image_owner_alias of aws_ami.
func (a amiAttributes) ImageOwnerAlias() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_owner_alias"))
}

// ImageType returns a reference to field image_type of aws_ami.
func (a amiAttributes) ImageType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_type"))
}

// ImdsSupport returns a reference to field imds_support of aws_ami.
func (a amiAttributes) ImdsSupport() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("imds_support"))
}

// KernelId returns a reference to field kernel_id of aws_ami.
func (a amiAttributes) KernelId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("kernel_id"))
}

// ManageEbsSnapshots returns a reference to field manage_ebs_snapshots of aws_ami.
func (a amiAttributes) ManageEbsSnapshots() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("manage_ebs_snapshots"))
}

// Name returns a reference to field name of aws_ami.
func (a amiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_ami.
func (a amiAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("owner_id"))
}

// Platform returns a reference to field platform of aws_ami.
func (a amiAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("platform"))
}

// PlatformDetails returns a reference to field platform_details of aws_ami.
func (a amiAttributes) PlatformDetails() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("platform_details"))
}

// Public returns a reference to field public of aws_ami.
func (a amiAttributes) Public() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("public"))
}

// RamdiskId returns a reference to field ramdisk_id of aws_ami.
func (a amiAttributes) RamdiskId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("ramdisk_id"))
}

// RootDeviceName returns a reference to field root_device_name of aws_ami.
func (a amiAttributes) RootDeviceName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("root_device_name"))
}

// RootSnapshotId returns a reference to field root_snapshot_id of aws_ami.
func (a amiAttributes) RootSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("root_snapshot_id"))
}

// SriovNetSupport returns a reference to field sriov_net_support of aws_ami.
func (a amiAttributes) SriovNetSupport() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("sriov_net_support"))
}

// Tags returns a reference to field tags of aws_ami.
func (a amiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ami.
func (a amiAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("tags_all"))
}

// TpmSupport returns a reference to field tpm_support of aws_ami.
func (a amiAttributes) TpmSupport() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("tpm_support"))
}

// UsageOperation returns a reference to field usage_operation of aws_ami.
func (a amiAttributes) UsageOperation() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("usage_operation"))
}

// VirtualizationType returns a reference to field virtualization_type of aws_ami.
func (a amiAttributes) VirtualizationType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("virtualization_type"))
}

func (a amiAttributes) EbsBlockDevice() terra.SetValue[ami.EbsBlockDeviceAttributes] {
	return terra.ReferenceAsSet[ami.EbsBlockDeviceAttributes](a.ref.Append("ebs_block_device"))
}

func (a amiAttributes) EphemeralBlockDevice() terra.SetValue[ami.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceAsSet[ami.EphemeralBlockDeviceAttributes](a.ref.Append("ephemeral_block_device"))
}

func (a amiAttributes) Timeouts() ami.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ami.TimeoutsAttributes](a.ref.Append("timeouts"))
}

type amiState struct {
	Architecture         string                          `json:"architecture"`
	Arn                  string                          `json:"arn"`
	BootMode             string                          `json:"boot_mode"`
	DeprecationTime      string                          `json:"deprecation_time"`
	Description          string                          `json:"description"`
	EnaSupport           bool                            `json:"ena_support"`
	Hypervisor           string                          `json:"hypervisor"`
	Id                   string                          `json:"id"`
	ImageLocation        string                          `json:"image_location"`
	ImageOwnerAlias      string                          `json:"image_owner_alias"`
	ImageType            string                          `json:"image_type"`
	ImdsSupport          string                          `json:"imds_support"`
	KernelId             string                          `json:"kernel_id"`
	ManageEbsSnapshots   bool                            `json:"manage_ebs_snapshots"`
	Name                 string                          `json:"name"`
	OwnerId              string                          `json:"owner_id"`
	Platform             string                          `json:"platform"`
	PlatformDetails      string                          `json:"platform_details"`
	Public               bool                            `json:"public"`
	RamdiskId            string                          `json:"ramdisk_id"`
	RootDeviceName       string                          `json:"root_device_name"`
	RootSnapshotId       string                          `json:"root_snapshot_id"`
	SriovNetSupport      string                          `json:"sriov_net_support"`
	Tags                 map[string]string               `json:"tags"`
	TagsAll              map[string]string               `json:"tags_all"`
	TpmSupport           string                          `json:"tpm_support"`
	UsageOperation       string                          `json:"usage_operation"`
	VirtualizationType   string                          `json:"virtualization_type"`
	EbsBlockDevice       []ami.EbsBlockDeviceState       `json:"ebs_block_device"`
	EphemeralBlockDevice []ami.EphemeralBlockDeviceState `json:"ephemeral_block_device"`
	Timeouts             *ami.TimeoutsState              `json:"timeouts"`
}
