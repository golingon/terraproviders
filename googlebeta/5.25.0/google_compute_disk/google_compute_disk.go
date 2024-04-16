// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_disk

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

// Resource represents the Terraform resource google_compute_disk.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcd *Resource) Type() string {
	return "google_compute_disk"
}

// LocalName returns the local name for [Resource].
func (gcd *Resource) LocalName() string {
	return gcd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcd *Resource) Configuration() interface{} {
	return gcd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcd *Resource) Dependencies() terra.Dependencies {
	return gcd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcd *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcd *Resource) Attributes() googleComputeDiskAttributes {
	return googleComputeDiskAttributes{ref: terra.ReferenceResource(gcd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcd *Resource) ImportState(state io.Reader) error {
	gcd.state = &googleComputeDiskState{}
	if err := json.NewDecoder(state).Decode(gcd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcd.Type(), gcd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcd *Resource) State() (*googleComputeDiskState, bool) {
	return gcd.state, gcd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcd *Resource) StateMust() *googleComputeDiskState {
	if gcd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcd.Type(), gcd.LocalName()))
	}
	return gcd.state
}

// Args contains the configurations for google_compute_disk.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableConfidentialCompute: bool, optional
	EnableConfidentialCompute terra.BoolValue `hcl:"enable_confidential_compute,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, optional
	Image terra.StringValue `hcl:"image,attr"`
	// Interface: string, optional
	Interface terra.StringValue `hcl:"interface,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Licenses: list of string, optional
	Licenses terra.ListValue[terra.StringValue] `hcl:"licenses,attr"`
	// MultiWriter: bool, optional
	MultiWriter terra.BoolValue `hcl:"multi_writer,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PhysicalBlockSizeBytes: number, optional
	PhysicalBlockSizeBytes terra.NumberValue `hcl:"physical_block_size_bytes,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProvisionedIops: number, optional
	ProvisionedIops terra.NumberValue `hcl:"provisioned_iops,attr"`
	// ProvisionedThroughput: number, optional
	ProvisionedThroughput terra.NumberValue `hcl:"provisioned_throughput,attr"`
	// ResourcePolicies: list of string, optional
	ResourcePolicies terra.ListValue[terra.StringValue] `hcl:"resource_policies,attr"`
	// Size: number, optional
	Size terra.NumberValue `hcl:"size,attr"`
	// Snapshot: string, optional
	Snapshot terra.StringValue `hcl:"snapshot,attr"`
	// SourceDisk: string, optional
	SourceDisk terra.StringValue `hcl:"source_disk,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AsyncPrimaryDisk: optional
	AsyncPrimaryDisk *AsyncPrimaryDisk `hcl:"async_primary_disk,block"`
	// DiskEncryptionKey: optional
	DiskEncryptionKey *DiskEncryptionKey `hcl:"disk_encryption_key,block"`
	// GuestOsFeatures: min=0
	GuestOsFeatures []GuestOsFeatures `hcl:"guest_os_features,block" validate:"min=0"`
	// SourceImageEncryptionKey: optional
	SourceImageEncryptionKey *SourceImageEncryptionKey `hcl:"source_image_encryption_key,block"`
	// SourceSnapshotEncryptionKey: optional
	SourceSnapshotEncryptionKey *SourceSnapshotEncryptionKey `hcl:"source_snapshot_encryption_key,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeDiskAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_disk.
func (gcd googleComputeDiskAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_disk.
func (gcd googleComputeDiskAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("description"))
}

// DiskId returns a reference to field disk_id of google_compute_disk.
func (gcd googleComputeDiskAttributes) DiskId() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("disk_id"))
}

// EffectiveLabels returns a reference to field effective_labels of google_compute_disk.
func (gcd googleComputeDiskAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcd.ref.Append("effective_labels"))
}

// EnableConfidentialCompute returns a reference to field enable_confidential_compute of google_compute_disk.
func (gcd googleComputeDiskAttributes) EnableConfidentialCompute() terra.BoolValue {
	return terra.ReferenceAsBool(gcd.ref.Append("enable_confidential_compute"))
}

// Id returns a reference to field id of google_compute_disk.
func (gcd googleComputeDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_disk.
func (gcd googleComputeDiskAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("image"))
}

// Interface returns a reference to field interface of google_compute_disk.
func (gcd googleComputeDiskAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("interface"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_disk.
func (gcd googleComputeDiskAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_disk.
func (gcd googleComputeDiskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcd.ref.Append("labels"))
}

// LastAttachTimestamp returns a reference to field last_attach_timestamp of google_compute_disk.
func (gcd googleComputeDiskAttributes) LastAttachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("last_attach_timestamp"))
}

// LastDetachTimestamp returns a reference to field last_detach_timestamp of google_compute_disk.
func (gcd googleComputeDiskAttributes) LastDetachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("last_detach_timestamp"))
}

// Licenses returns a reference to field licenses of google_compute_disk.
func (gcd googleComputeDiskAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcd.ref.Append("licenses"))
}

// MultiWriter returns a reference to field multi_writer of google_compute_disk.
func (gcd googleComputeDiskAttributes) MultiWriter() terra.BoolValue {
	return terra.ReferenceAsBool(gcd.ref.Append("multi_writer"))
}

// Name returns a reference to field name of google_compute_disk.
func (gcd googleComputeDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("name"))
}

// PhysicalBlockSizeBytes returns a reference to field physical_block_size_bytes of google_compute_disk.
func (gcd googleComputeDiskAttributes) PhysicalBlockSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(gcd.ref.Append("physical_block_size_bytes"))
}

// Project returns a reference to field project of google_compute_disk.
func (gcd googleComputeDiskAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("project"))
}

// ProvisionedIops returns a reference to field provisioned_iops of google_compute_disk.
func (gcd googleComputeDiskAttributes) ProvisionedIops() terra.NumberValue {
	return terra.ReferenceAsNumber(gcd.ref.Append("provisioned_iops"))
}

// ProvisionedThroughput returns a reference to field provisioned_throughput of google_compute_disk.
func (gcd googleComputeDiskAttributes) ProvisionedThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(gcd.ref.Append("provisioned_throughput"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_disk.
func (gcd googleComputeDiskAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcd.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_disk.
func (gcd googleComputeDiskAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_disk.
func (gcd googleComputeDiskAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(gcd.ref.Append("size"))
}

// Snapshot returns a reference to field snapshot of google_compute_disk.
func (gcd googleComputeDiskAttributes) Snapshot() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("snapshot"))
}

// SourceDisk returns a reference to field source_disk of google_compute_disk.
func (gcd googleComputeDiskAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("source_disk"))
}

// SourceDiskId returns a reference to field source_disk_id of google_compute_disk.
func (gcd googleComputeDiskAttributes) SourceDiskId() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("source_disk_id"))
}

// SourceImageId returns a reference to field source_image_id of google_compute_disk.
func (gcd googleComputeDiskAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("source_image_id"))
}

// SourceSnapshotId returns a reference to field source_snapshot_id of google_compute_disk.
func (gcd googleComputeDiskAttributes) SourceSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("source_snapshot_id"))
}

// TerraformLabels returns a reference to field terraform_labels of google_compute_disk.
func (gcd googleComputeDiskAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcd.ref.Append("terraform_labels"))
}

// Type returns a reference to field type of google_compute_disk.
func (gcd googleComputeDiskAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("type"))
}

// Users returns a reference to field users of google_compute_disk.
func (gcd googleComputeDiskAttributes) Users() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcd.ref.Append("users"))
}

// Zone returns a reference to field zone of google_compute_disk.
func (gcd googleComputeDiskAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(gcd.ref.Append("zone"))
}

func (gcd googleComputeDiskAttributes) AsyncPrimaryDisk() terra.ListValue[AsyncPrimaryDiskAttributes] {
	return terra.ReferenceAsList[AsyncPrimaryDiskAttributes](gcd.ref.Append("async_primary_disk"))
}

func (gcd googleComputeDiskAttributes) DiskEncryptionKey() terra.ListValue[DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DiskEncryptionKeyAttributes](gcd.ref.Append("disk_encryption_key"))
}

func (gcd googleComputeDiskAttributes) GuestOsFeatures() terra.SetValue[GuestOsFeaturesAttributes] {
	return terra.ReferenceAsSet[GuestOsFeaturesAttributes](gcd.ref.Append("guest_os_features"))
}

func (gcd googleComputeDiskAttributes) SourceImageEncryptionKey() terra.ListValue[SourceImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[SourceImageEncryptionKeyAttributes](gcd.ref.Append("source_image_encryption_key"))
}

func (gcd googleComputeDiskAttributes) SourceSnapshotEncryptionKey() terra.ListValue[SourceSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[SourceSnapshotEncryptionKeyAttributes](gcd.ref.Append("source_snapshot_encryption_key"))
}

func (gcd googleComputeDiskAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcd.ref.Append("timeouts"))
}

type googleComputeDiskState struct {
	CreationTimestamp           string                             `json:"creation_timestamp"`
	Description                 string                             `json:"description"`
	DiskId                      string                             `json:"disk_id"`
	EffectiveLabels             map[string]string                  `json:"effective_labels"`
	EnableConfidentialCompute   bool                               `json:"enable_confidential_compute"`
	Id                          string                             `json:"id"`
	Image                       string                             `json:"image"`
	Interface                   string                             `json:"interface"`
	LabelFingerprint            string                             `json:"label_fingerprint"`
	Labels                      map[string]string                  `json:"labels"`
	LastAttachTimestamp         string                             `json:"last_attach_timestamp"`
	LastDetachTimestamp         string                             `json:"last_detach_timestamp"`
	Licenses                    []string                           `json:"licenses"`
	MultiWriter                 bool                               `json:"multi_writer"`
	Name                        string                             `json:"name"`
	PhysicalBlockSizeBytes      float64                            `json:"physical_block_size_bytes"`
	Project                     string                             `json:"project"`
	ProvisionedIops             float64                            `json:"provisioned_iops"`
	ProvisionedThroughput       float64                            `json:"provisioned_throughput"`
	ResourcePolicies            []string                           `json:"resource_policies"`
	SelfLink                    string                             `json:"self_link"`
	Size                        float64                            `json:"size"`
	Snapshot                    string                             `json:"snapshot"`
	SourceDisk                  string                             `json:"source_disk"`
	SourceDiskId                string                             `json:"source_disk_id"`
	SourceImageId               string                             `json:"source_image_id"`
	SourceSnapshotId            string                             `json:"source_snapshot_id"`
	TerraformLabels             map[string]string                  `json:"terraform_labels"`
	Type                        string                             `json:"type"`
	Users                       []string                           `json:"users"`
	Zone                        string                             `json:"zone"`
	AsyncPrimaryDisk            []AsyncPrimaryDiskState            `json:"async_primary_disk"`
	DiskEncryptionKey           []DiskEncryptionKeyState           `json:"disk_encryption_key"`
	GuestOsFeatures             []GuestOsFeaturesState             `json:"guest_os_features"`
	SourceImageEncryptionKey    []SourceImageEncryptionKeyState    `json:"source_image_encryption_key"`
	SourceSnapshotEncryptionKey []SourceSnapshotEncryptionKeyState `json:"source_snapshot_encryption_key"`
	Timeouts                    *TimeoutsState                     `json:"timeouts"`
}
