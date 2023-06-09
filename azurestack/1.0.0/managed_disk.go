// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	manageddisk "github.com/golingon/terraproviders/azurestack/1.0.0/manageddisk"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagedDisk creates a new instance of [ManagedDisk].
func NewManagedDisk(name string, args ManagedDiskArgs) *ManagedDisk {
	return &ManagedDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagedDisk)(nil)

// ManagedDisk represents the Terraform resource azurestack_managed_disk.
type ManagedDisk struct {
	Name      string
	Args      ManagedDiskArgs
	state     *managedDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagedDisk].
func (md *ManagedDisk) Type() string {
	return "azurestack_managed_disk"
}

// LocalName returns the local name for [ManagedDisk].
func (md *ManagedDisk) LocalName() string {
	return md.Name
}

// Configuration returns the configuration (args) for [ManagedDisk].
func (md *ManagedDisk) Configuration() interface{} {
	return md.Args
}

// DependOn is used for other resources to depend on [ManagedDisk].
func (md *ManagedDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(md)
}

// Dependencies returns the list of resources [ManagedDisk] depends_on.
func (md *ManagedDisk) Dependencies() terra.Dependencies {
	return md.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagedDisk].
func (md *ManagedDisk) LifecycleManagement() *terra.Lifecycle {
	return md.Lifecycle
}

// Attributes returns the attributes for [ManagedDisk].
func (md *ManagedDisk) Attributes() managedDiskAttributes {
	return managedDiskAttributes{ref: terra.ReferenceResource(md)}
}

// ImportState imports the given attribute values into [ManagedDisk]'s state.
func (md *ManagedDisk) ImportState(av io.Reader) error {
	md.state = &managedDiskState{}
	if err := json.NewDecoder(av).Decode(md.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", md.Type(), md.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagedDisk] has state.
func (md *ManagedDisk) State() (*managedDiskState, bool) {
	return md.state, md.state != nil
}

// StateMust returns the state for [ManagedDisk]. Panics if the state is nil.
func (md *ManagedDisk) StateMust() *managedDiskState {
	if md.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", md.Type(), md.LocalName()))
	}
	return md.state
}

// ManagedDiskArgs contains the configurations for azurestack_managed_disk.
type ManagedDiskArgs struct {
	// CreateOption: string, required
	CreateOption terra.StringValue `hcl:"create_option,attr" validate:"required"`
	// DiskSizeGb: number, optional
	DiskSizeGb terra.NumberValue `hcl:"disk_size_gb,attr"`
	// HyperVGeneration: string, optional
	HyperVGeneration terra.StringValue `hcl:"hyper_v_generation,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageReferenceId: string, optional
	ImageReferenceId terra.StringValue `hcl:"image_reference_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OsType: string, optional
	OsType terra.StringValue `hcl:"os_type,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceResourceId: string, optional
	SourceResourceId terra.StringValue `hcl:"source_resource_id,attr"`
	// SourceUri: string, optional
	SourceUri terra.StringValue `hcl:"source_uri,attr"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
	// StorageAccountType: string, required
	StorageAccountType terra.StringValue `hcl:"storage_account_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Encryption: optional
	Encryption *manageddisk.Encryption `hcl:"encryption,block"`
	// Timeouts: optional
	Timeouts *manageddisk.Timeouts `hcl:"timeouts,block"`
}
type managedDiskAttributes struct {
	ref terra.Reference
}

// CreateOption returns a reference to field create_option of azurestack_managed_disk.
func (md managedDiskAttributes) CreateOption() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("create_option"))
}

// DiskSizeGb returns a reference to field disk_size_gb of azurestack_managed_disk.
func (md managedDiskAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("disk_size_gb"))
}

// HyperVGeneration returns a reference to field hyper_v_generation of azurestack_managed_disk.
func (md managedDiskAttributes) HyperVGeneration() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("hyper_v_generation"))
}

// Id returns a reference to field id of azurestack_managed_disk.
func (md managedDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("id"))
}

// ImageReferenceId returns a reference to field image_reference_id of azurestack_managed_disk.
func (md managedDiskAttributes) ImageReferenceId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("image_reference_id"))
}

// Location returns a reference to field location of azurestack_managed_disk.
func (md managedDiskAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_managed_disk.
func (md managedDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurestack_managed_disk.
func (md managedDiskAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("os_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_managed_disk.
func (md managedDiskAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("resource_group_name"))
}

// SourceResourceId returns a reference to field source_resource_id of azurestack_managed_disk.
func (md managedDiskAttributes) SourceResourceId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("source_resource_id"))
}

// SourceUri returns a reference to field source_uri of azurestack_managed_disk.
func (md managedDiskAttributes) SourceUri() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("source_uri"))
}

// StorageAccountId returns a reference to field storage_account_id of azurestack_managed_disk.
func (md managedDiskAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("storage_account_id"))
}

// StorageAccountType returns a reference to field storage_account_type of azurestack_managed_disk.
func (md managedDiskAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("storage_account_type"))
}

// Tags returns a reference to field tags of azurestack_managed_disk.
func (md managedDiskAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](md.ref.Append("tags"))
}

func (md managedDiskAttributes) Encryption() terra.ListValue[manageddisk.EncryptionAttributes] {
	return terra.ReferenceAsList[manageddisk.EncryptionAttributes](md.ref.Append("encryption"))
}

func (md managedDiskAttributes) Timeouts() manageddisk.TimeoutsAttributes {
	return terra.ReferenceAsSingle[manageddisk.TimeoutsAttributes](md.ref.Append("timeouts"))
}

type managedDiskState struct {
	CreateOption       string                        `json:"create_option"`
	DiskSizeGb         float64                       `json:"disk_size_gb"`
	HyperVGeneration   string                        `json:"hyper_v_generation"`
	Id                 string                        `json:"id"`
	ImageReferenceId   string                        `json:"image_reference_id"`
	Location           string                        `json:"location"`
	Name               string                        `json:"name"`
	OsType             string                        `json:"os_type"`
	ResourceGroupName  string                        `json:"resource_group_name"`
	SourceResourceId   string                        `json:"source_resource_id"`
	SourceUri          string                        `json:"source_uri"`
	StorageAccountId   string                        `json:"storage_account_id"`
	StorageAccountType string                        `json:"storage_account_type"`
	Tags               map[string]string             `json:"tags"`
	Encryption         []manageddisk.EncryptionState `json:"encryption"`
	Timeouts           *manageddisk.TimeoutsState    `json:"timeouts"`
}
