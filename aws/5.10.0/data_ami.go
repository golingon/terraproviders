// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataami "github.com/golingon/terraproviders/aws/5.10.0/dataami"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAmi creates a new instance of [DataAmi].
func NewDataAmi(name string, args DataAmiArgs) *DataAmi {
	return &DataAmi{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAmi)(nil)

// DataAmi represents the Terraform data resource aws_ami.
type DataAmi struct {
	Name string
	Args DataAmiArgs
}

// DataSource returns the Terraform object type for [DataAmi].
func (a *DataAmi) DataSource() string {
	return "aws_ami"
}

// LocalName returns the local name for [DataAmi].
func (a *DataAmi) LocalName() string {
	return a.Name
}

// Configuration returns the configuration (args) for [DataAmi].
func (a *DataAmi) Configuration() interface{} {
	return a.Args
}

// Attributes returns the attributes for [DataAmi].
func (a *DataAmi) Attributes() dataAmiAttributes {
	return dataAmiAttributes{ref: terra.ReferenceDataResource(a)}
}

// DataAmiArgs contains the configurations for aws_ami.
type DataAmiArgs struct {
	// ExecutableUsers: list of string, optional
	ExecutableUsers terra.ListValue[terra.StringValue] `hcl:"executable_users,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeDeprecated: bool, optional
	IncludeDeprecated terra.BoolValue `hcl:"include_deprecated,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// Owners: list of string, optional
	Owners terra.ListValue[terra.StringValue] `hcl:"owners,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// BlockDeviceMappings: min=0
	BlockDeviceMappings []dataami.BlockDeviceMappings `hcl:"block_device_mappings,block" validate:"min=0"`
	// ProductCodes: min=0
	ProductCodes []dataami.ProductCodes `hcl:"product_codes,block" validate:"min=0"`
	// Filter: min=0
	Filter []dataami.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataami.Timeouts `hcl:"timeouts,block"`
}
type dataAmiAttributes struct {
	ref terra.Reference
}

// Architecture returns a reference to field architecture of aws_ami.
func (a dataAmiAttributes) Architecture() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("architecture"))
}

// Arn returns a reference to field arn of aws_ami.
func (a dataAmiAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn"))
}

// BootMode returns a reference to field boot_mode of aws_ami.
func (a dataAmiAttributes) BootMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("boot_mode"))
}

// CreationDate returns a reference to field creation_date of aws_ami.
func (a dataAmiAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("creation_date"))
}

// DeprecationTime returns a reference to field deprecation_time of aws_ami.
func (a dataAmiAttributes) DeprecationTime() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("deprecation_time"))
}

// Description returns a reference to field description of aws_ami.
func (a dataAmiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("description"))
}

// EnaSupport returns a reference to field ena_support of aws_ami.
func (a dataAmiAttributes) EnaSupport() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("ena_support"))
}

// ExecutableUsers returns a reference to field executable_users of aws_ami.
func (a dataAmiAttributes) ExecutableUsers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("executable_users"))
}

// Hypervisor returns a reference to field hypervisor of aws_ami.
func (a dataAmiAttributes) Hypervisor() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("hypervisor"))
}

// Id returns a reference to field id of aws_ami.
func (a dataAmiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

// ImageId returns a reference to field image_id of aws_ami.
func (a dataAmiAttributes) ImageId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_id"))
}

// ImageLocation returns a reference to field image_location of aws_ami.
func (a dataAmiAttributes) ImageLocation() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_location"))
}

// ImageOwnerAlias returns a reference to field image_owner_alias of aws_ami.
func (a dataAmiAttributes) ImageOwnerAlias() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_owner_alias"))
}

// ImageType returns a reference to field image_type of aws_ami.
func (a dataAmiAttributes) ImageType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("image_type"))
}

// ImdsSupport returns a reference to field imds_support of aws_ami.
func (a dataAmiAttributes) ImdsSupport() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("imds_support"))
}

// IncludeDeprecated returns a reference to field include_deprecated of aws_ami.
func (a dataAmiAttributes) IncludeDeprecated() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("include_deprecated"))
}

// KernelId returns a reference to field kernel_id of aws_ami.
func (a dataAmiAttributes) KernelId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("kernel_id"))
}

// MostRecent returns a reference to field most_recent of aws_ami.
func (a dataAmiAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("most_recent"))
}

// Name returns a reference to field name of aws_ami.
func (a dataAmiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

// NameRegex returns a reference to field name_regex of aws_ami.
func (a dataAmiAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name_regex"))
}

// OwnerId returns a reference to field owner_id of aws_ami.
func (a dataAmiAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("owner_id"))
}

// Owners returns a reference to field owners of aws_ami.
func (a dataAmiAttributes) Owners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("owners"))
}

// Platform returns a reference to field platform of aws_ami.
func (a dataAmiAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("platform"))
}

// PlatformDetails returns a reference to field platform_details of aws_ami.
func (a dataAmiAttributes) PlatformDetails() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("platform_details"))
}

// Public returns a reference to field public of aws_ami.
func (a dataAmiAttributes) Public() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("public"))
}

// RamdiskId returns a reference to field ramdisk_id of aws_ami.
func (a dataAmiAttributes) RamdiskId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("ramdisk_id"))
}

// RootDeviceName returns a reference to field root_device_name of aws_ami.
func (a dataAmiAttributes) RootDeviceName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("root_device_name"))
}

// RootDeviceType returns a reference to field root_device_type of aws_ami.
func (a dataAmiAttributes) RootDeviceType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("root_device_type"))
}

// RootSnapshotId returns a reference to field root_snapshot_id of aws_ami.
func (a dataAmiAttributes) RootSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("root_snapshot_id"))
}

// SriovNetSupport returns a reference to field sriov_net_support of aws_ami.
func (a dataAmiAttributes) SriovNetSupport() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("sriov_net_support"))
}

// State returns a reference to field state of aws_ami.
func (a dataAmiAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("state"))
}

// StateReason returns a reference to field state_reason of aws_ami.
func (a dataAmiAttributes) StateReason() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("state_reason"))
}

// Tags returns a reference to field tags of aws_ami.
func (a dataAmiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("tags"))
}

// TpmSupport returns a reference to field tpm_support of aws_ami.
func (a dataAmiAttributes) TpmSupport() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("tpm_support"))
}

// UsageOperation returns a reference to field usage_operation of aws_ami.
func (a dataAmiAttributes) UsageOperation() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("usage_operation"))
}

// VirtualizationType returns a reference to field virtualization_type of aws_ami.
func (a dataAmiAttributes) VirtualizationType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("virtualization_type"))
}

func (a dataAmiAttributes) BlockDeviceMappings() terra.SetValue[dataami.BlockDeviceMappingsAttributes] {
	return terra.ReferenceAsSet[dataami.BlockDeviceMappingsAttributes](a.ref.Append("block_device_mappings"))
}

func (a dataAmiAttributes) ProductCodes() terra.SetValue[dataami.ProductCodesAttributes] {
	return terra.ReferenceAsSet[dataami.ProductCodesAttributes](a.ref.Append("product_codes"))
}

func (a dataAmiAttributes) Filter() terra.SetValue[dataami.FilterAttributes] {
	return terra.ReferenceAsSet[dataami.FilterAttributes](a.ref.Append("filter"))
}

func (a dataAmiAttributes) Timeouts() dataami.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataami.TimeoutsAttributes](a.ref.Append("timeouts"))
}
