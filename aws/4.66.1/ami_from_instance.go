// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	amifrominstance "github.com/golingon/terraproviders/aws/4.66.1/amifrominstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAmiFromInstance creates a new instance of [AmiFromInstance].
func NewAmiFromInstance(name string, args AmiFromInstanceArgs) *AmiFromInstance {
	return &AmiFromInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AmiFromInstance)(nil)

// AmiFromInstance represents the Terraform resource aws_ami_from_instance.
type AmiFromInstance struct {
	Name      string
	Args      AmiFromInstanceArgs
	state     *amiFromInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AmiFromInstance].
func (afi *AmiFromInstance) Type() string {
	return "aws_ami_from_instance"
}

// LocalName returns the local name for [AmiFromInstance].
func (afi *AmiFromInstance) LocalName() string {
	return afi.Name
}

// Configuration returns the configuration (args) for [AmiFromInstance].
func (afi *AmiFromInstance) Configuration() interface{} {
	return afi.Args
}

// DependOn is used for other resources to depend on [AmiFromInstance].
func (afi *AmiFromInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(afi)
}

// Dependencies returns the list of resources [AmiFromInstance] depends_on.
func (afi *AmiFromInstance) Dependencies() terra.Dependencies {
	return afi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AmiFromInstance].
func (afi *AmiFromInstance) LifecycleManagement() *terra.Lifecycle {
	return afi.Lifecycle
}

// Attributes returns the attributes for [AmiFromInstance].
func (afi *AmiFromInstance) Attributes() amiFromInstanceAttributes {
	return amiFromInstanceAttributes{ref: terra.ReferenceResource(afi)}
}

// ImportState imports the given attribute values into [AmiFromInstance]'s state.
func (afi *AmiFromInstance) ImportState(av io.Reader) error {
	afi.state = &amiFromInstanceState{}
	if err := json.NewDecoder(av).Decode(afi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afi.Type(), afi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AmiFromInstance] has state.
func (afi *AmiFromInstance) State() (*amiFromInstanceState, bool) {
	return afi.state, afi.state != nil
}

// StateMust returns the state for [AmiFromInstance]. Panics if the state is nil.
func (afi *AmiFromInstance) StateMust() *amiFromInstanceState {
	if afi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afi.Type(), afi.LocalName()))
	}
	return afi.state
}

// AmiFromInstanceArgs contains the configurations for aws_ami_from_instance.
type AmiFromInstanceArgs struct {
	// DeprecationTime: string, optional
	DeprecationTime terra.StringValue `hcl:"deprecation_time,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SnapshotWithoutReboot: bool, optional
	SnapshotWithoutReboot terra.BoolValue `hcl:"snapshot_without_reboot,attr"`
	// SourceInstanceId: string, required
	SourceInstanceId terra.StringValue `hcl:"source_instance_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []amifrominstance.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []amifrominstance.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *amifrominstance.Timeouts `hcl:"timeouts,block"`
}
type amiFromInstanceAttributes struct {
	ref terra.Reference
}

// Architecture returns a reference to field architecture of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Architecture() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("architecture"))
}

// Arn returns a reference to field arn of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("arn"))
}

// BootMode returns a reference to field boot_mode of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) BootMode() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("boot_mode"))
}

// DeprecationTime returns a reference to field deprecation_time of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) DeprecationTime() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("deprecation_time"))
}

// Description returns a reference to field description of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("description"))
}

// EnaSupport returns a reference to field ena_support of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) EnaSupport() terra.BoolValue {
	return terra.ReferenceAsBool(afi.ref.Append("ena_support"))
}

// Hypervisor returns a reference to field hypervisor of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Hypervisor() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("hypervisor"))
}

// Id returns a reference to field id of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("id"))
}

// ImageLocation returns a reference to field image_location of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) ImageLocation() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("image_location"))
}

// ImageOwnerAlias returns a reference to field image_owner_alias of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) ImageOwnerAlias() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("image_owner_alias"))
}

// ImageType returns a reference to field image_type of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) ImageType() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("image_type"))
}

// ImdsSupport returns a reference to field imds_support of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) ImdsSupport() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("imds_support"))
}

// KernelId returns a reference to field kernel_id of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) KernelId() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("kernel_id"))
}

// ManageEbsSnapshots returns a reference to field manage_ebs_snapshots of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) ManageEbsSnapshots() terra.BoolValue {
	return terra.ReferenceAsBool(afi.ref.Append("manage_ebs_snapshots"))
}

// Name returns a reference to field name of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("owner_id"))
}

// Platform returns a reference to field platform of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("platform"))
}

// PlatformDetails returns a reference to field platform_details of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) PlatformDetails() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("platform_details"))
}

// Public returns a reference to field public of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Public() terra.BoolValue {
	return terra.ReferenceAsBool(afi.ref.Append("public"))
}

// RamdiskId returns a reference to field ramdisk_id of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) RamdiskId() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("ramdisk_id"))
}

// RootDeviceName returns a reference to field root_device_name of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) RootDeviceName() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("root_device_name"))
}

// RootSnapshotId returns a reference to field root_snapshot_id of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) RootSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("root_snapshot_id"))
}

// SnapshotWithoutReboot returns a reference to field snapshot_without_reboot of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) SnapshotWithoutReboot() terra.BoolValue {
	return terra.ReferenceAsBool(afi.ref.Append("snapshot_without_reboot"))
}

// SourceInstanceId returns a reference to field source_instance_id of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) SourceInstanceId() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("source_instance_id"))
}

// SriovNetSupport returns a reference to field sriov_net_support of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) SriovNetSupport() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("sriov_net_support"))
}

// Tags returns a reference to field tags of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afi.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afi.ref.Append("tags_all"))
}

// TpmSupport returns a reference to field tpm_support of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) TpmSupport() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("tpm_support"))
}

// UsageOperation returns a reference to field usage_operation of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) UsageOperation() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("usage_operation"))
}

// VirtualizationType returns a reference to field virtualization_type of aws_ami_from_instance.
func (afi amiFromInstanceAttributes) VirtualizationType() terra.StringValue {
	return terra.ReferenceAsString(afi.ref.Append("virtualization_type"))
}

func (afi amiFromInstanceAttributes) EbsBlockDevice() terra.SetValue[amifrominstance.EbsBlockDeviceAttributes] {
	return terra.ReferenceAsSet[amifrominstance.EbsBlockDeviceAttributes](afi.ref.Append("ebs_block_device"))
}

func (afi amiFromInstanceAttributes) EphemeralBlockDevice() terra.SetValue[amifrominstance.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceAsSet[amifrominstance.EphemeralBlockDeviceAttributes](afi.ref.Append("ephemeral_block_device"))
}

func (afi amiFromInstanceAttributes) Timeouts() amifrominstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[amifrominstance.TimeoutsAttributes](afi.ref.Append("timeouts"))
}

type amiFromInstanceState struct {
	Architecture          string                                      `json:"architecture"`
	Arn                   string                                      `json:"arn"`
	BootMode              string                                      `json:"boot_mode"`
	DeprecationTime       string                                      `json:"deprecation_time"`
	Description           string                                      `json:"description"`
	EnaSupport            bool                                        `json:"ena_support"`
	Hypervisor            string                                      `json:"hypervisor"`
	Id                    string                                      `json:"id"`
	ImageLocation         string                                      `json:"image_location"`
	ImageOwnerAlias       string                                      `json:"image_owner_alias"`
	ImageType             string                                      `json:"image_type"`
	ImdsSupport           string                                      `json:"imds_support"`
	KernelId              string                                      `json:"kernel_id"`
	ManageEbsSnapshots    bool                                        `json:"manage_ebs_snapshots"`
	Name                  string                                      `json:"name"`
	OwnerId               string                                      `json:"owner_id"`
	Platform              string                                      `json:"platform"`
	PlatformDetails       string                                      `json:"platform_details"`
	Public                bool                                        `json:"public"`
	RamdiskId             string                                      `json:"ramdisk_id"`
	RootDeviceName        string                                      `json:"root_device_name"`
	RootSnapshotId        string                                      `json:"root_snapshot_id"`
	SnapshotWithoutReboot bool                                        `json:"snapshot_without_reboot"`
	SourceInstanceId      string                                      `json:"source_instance_id"`
	SriovNetSupport       string                                      `json:"sriov_net_support"`
	Tags                  map[string]string                           `json:"tags"`
	TagsAll               map[string]string                           `json:"tags_all"`
	TpmSupport            string                                      `json:"tpm_support"`
	UsageOperation        string                                      `json:"usage_operation"`
	VirtualizationType    string                                      `json:"virtualization_type"`
	EbsBlockDevice        []amifrominstance.EbsBlockDeviceState       `json:"ebs_block_device"`
	EphemeralBlockDevice  []amifrominstance.EphemeralBlockDeviceState `json:"ephemeral_block_device"`
	Timeouts              *amifrominstance.TimeoutsState              `json:"timeouts"`
}
