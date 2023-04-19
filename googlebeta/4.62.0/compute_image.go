// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeimage "github.com/golingon/terraproviders/googlebeta/4.62.0/computeimage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeImage creates a new instance of [ComputeImage].
func NewComputeImage(name string, args ComputeImageArgs) *ComputeImage {
	return &ComputeImage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeImage)(nil)

// ComputeImage represents the Terraform resource google_compute_image.
type ComputeImage struct {
	Name      string
	Args      ComputeImageArgs
	state     *computeImageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeImage].
func (ci *ComputeImage) Type() string {
	return "google_compute_image"
}

// LocalName returns the local name for [ComputeImage].
func (ci *ComputeImage) LocalName() string {
	return ci.Name
}

// Configuration returns the configuration (args) for [ComputeImage].
func (ci *ComputeImage) Configuration() interface{} {
	return ci.Args
}

// DependOn is used for other resources to depend on [ComputeImage].
func (ci *ComputeImage) DependOn() terra.Reference {
	return terra.ReferenceResource(ci)
}

// Dependencies returns the list of resources [ComputeImage] depends_on.
func (ci *ComputeImage) Dependencies() terra.Dependencies {
	return ci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeImage].
func (ci *ComputeImage) LifecycleManagement() *terra.Lifecycle {
	return ci.Lifecycle
}

// Attributes returns the attributes for [ComputeImage].
func (ci *ComputeImage) Attributes() computeImageAttributes {
	return computeImageAttributes{ref: terra.ReferenceResource(ci)}
}

// ImportState imports the given attribute values into [ComputeImage]'s state.
func (ci *ComputeImage) ImportState(av io.Reader) error {
	ci.state = &computeImageState{}
	if err := json.NewDecoder(av).Decode(ci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ci.Type(), ci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeImage] has state.
func (ci *ComputeImage) State() (*computeImageState, bool) {
	return ci.state, ci.state != nil
}

// StateMust returns the state for [ComputeImage]. Panics if the state is nil.
func (ci *ComputeImage) StateMust() *computeImageState {
	if ci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ci.Type(), ci.LocalName()))
	}
	return ci.state
}

// ComputeImageArgs contains the configurations for google_compute_image.
type ComputeImageArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DiskSizeGb: number, optional
	DiskSizeGb terra.NumberValue `hcl:"disk_size_gb,attr"`
	// Family: string, optional
	Family terra.StringValue `hcl:"family,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Licenses: list of string, optional
	Licenses terra.ListValue[terra.StringValue] `hcl:"licenses,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SourceDisk: string, optional
	SourceDisk terra.StringValue `hcl:"source_disk,attr"`
	// SourceImage: string, optional
	SourceImage terra.StringValue `hcl:"source_image,attr"`
	// SourceSnapshot: string, optional
	SourceSnapshot terra.StringValue `hcl:"source_snapshot,attr"`
	// GuestOsFeatures: min=0
	GuestOsFeatures []computeimage.GuestOsFeatures `hcl:"guest_os_features,block" validate:"min=0"`
	// ImageEncryptionKey: optional
	ImageEncryptionKey *computeimage.ImageEncryptionKey `hcl:"image_encryption_key,block"`
	// RawDisk: optional
	RawDisk *computeimage.RawDisk `hcl:"raw_disk,block"`
	// Timeouts: optional
	Timeouts *computeimage.Timeouts `hcl:"timeouts,block"`
}
type computeImageAttributes struct {
	ref terra.Reference
}

// ArchiveSizeBytes returns a reference to field archive_size_bytes of google_compute_image.
func (ci computeImageAttributes) ArchiveSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("archive_size_bytes"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_image.
func (ci computeImageAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_image.
func (ci computeImageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("description"))
}

// DiskSizeGb returns a reference to field disk_size_gb of google_compute_image.
func (ci computeImageAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("disk_size_gb"))
}

// Family returns a reference to field family of google_compute_image.
func (ci computeImageAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("family"))
}

// Id returns a reference to field id of google_compute_image.
func (ci computeImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("id"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_image.
func (ci computeImageAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_image.
func (ci computeImageAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("labels"))
}

// Licenses returns a reference to field licenses of google_compute_image.
func (ci computeImageAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("licenses"))
}

// Name returns a reference to field name of google_compute_image.
func (ci computeImageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_image.
func (ci computeImageAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_image.
func (ci computeImageAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("self_link"))
}

// SourceDisk returns a reference to field source_disk of google_compute_image.
func (ci computeImageAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("source_disk"))
}

// SourceImage returns a reference to field source_image of google_compute_image.
func (ci computeImageAttributes) SourceImage() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("source_image"))
}

// SourceSnapshot returns a reference to field source_snapshot of google_compute_image.
func (ci computeImageAttributes) SourceSnapshot() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("source_snapshot"))
}

func (ci computeImageAttributes) GuestOsFeatures() terra.SetValue[computeimage.GuestOsFeaturesAttributes] {
	return terra.ReferenceAsSet[computeimage.GuestOsFeaturesAttributes](ci.ref.Append("guest_os_features"))
}

func (ci computeImageAttributes) ImageEncryptionKey() terra.ListValue[computeimage.ImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computeimage.ImageEncryptionKeyAttributes](ci.ref.Append("image_encryption_key"))
}

func (ci computeImageAttributes) RawDisk() terra.ListValue[computeimage.RawDiskAttributes] {
	return terra.ReferenceAsList[computeimage.RawDiskAttributes](ci.ref.Append("raw_disk"))
}

func (ci computeImageAttributes) Timeouts() computeimage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeimage.TimeoutsAttributes](ci.ref.Append("timeouts"))
}

type computeImageState struct {
	ArchiveSizeBytes   float64                                `json:"archive_size_bytes"`
	CreationTimestamp  string                                 `json:"creation_timestamp"`
	Description        string                                 `json:"description"`
	DiskSizeGb         float64                                `json:"disk_size_gb"`
	Family             string                                 `json:"family"`
	Id                 string                                 `json:"id"`
	LabelFingerprint   string                                 `json:"label_fingerprint"`
	Labels             map[string]string                      `json:"labels"`
	Licenses           []string                               `json:"licenses"`
	Name               string                                 `json:"name"`
	Project            string                                 `json:"project"`
	SelfLink           string                                 `json:"self_link"`
	SourceDisk         string                                 `json:"source_disk"`
	SourceImage        string                                 `json:"source_image"`
	SourceSnapshot     string                                 `json:"source_snapshot"`
	GuestOsFeatures    []computeimage.GuestOsFeaturesState    `json:"guest_os_features"`
	ImageEncryptionKey []computeimage.ImageEncryptionKeyState `json:"image_encryption_key"`
	RawDisk            []computeimage.RawDiskState            `json:"raw_disk"`
	Timeouts           *computeimage.TimeoutsState            `json:"timeouts"`
}