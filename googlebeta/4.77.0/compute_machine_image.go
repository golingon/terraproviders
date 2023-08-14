// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computemachineimage "github.com/golingon/terraproviders/googlebeta/4.77.0/computemachineimage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeMachineImage creates a new instance of [ComputeMachineImage].
func NewComputeMachineImage(name string, args ComputeMachineImageArgs) *ComputeMachineImage {
	return &ComputeMachineImage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeMachineImage)(nil)

// ComputeMachineImage represents the Terraform resource google_compute_machine_image.
type ComputeMachineImage struct {
	Name      string
	Args      ComputeMachineImageArgs
	state     *computeMachineImageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeMachineImage].
func (cmi *ComputeMachineImage) Type() string {
	return "google_compute_machine_image"
}

// LocalName returns the local name for [ComputeMachineImage].
func (cmi *ComputeMachineImage) LocalName() string {
	return cmi.Name
}

// Configuration returns the configuration (args) for [ComputeMachineImage].
func (cmi *ComputeMachineImage) Configuration() interface{} {
	return cmi.Args
}

// DependOn is used for other resources to depend on [ComputeMachineImage].
func (cmi *ComputeMachineImage) DependOn() terra.Reference {
	return terra.ReferenceResource(cmi)
}

// Dependencies returns the list of resources [ComputeMachineImage] depends_on.
func (cmi *ComputeMachineImage) Dependencies() terra.Dependencies {
	return cmi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeMachineImage].
func (cmi *ComputeMachineImage) LifecycleManagement() *terra.Lifecycle {
	return cmi.Lifecycle
}

// Attributes returns the attributes for [ComputeMachineImage].
func (cmi *ComputeMachineImage) Attributes() computeMachineImageAttributes {
	return computeMachineImageAttributes{ref: terra.ReferenceResource(cmi)}
}

// ImportState imports the given attribute values into [ComputeMachineImage]'s state.
func (cmi *ComputeMachineImage) ImportState(av io.Reader) error {
	cmi.state = &computeMachineImageState{}
	if err := json.NewDecoder(av).Decode(cmi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmi.Type(), cmi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeMachineImage] has state.
func (cmi *ComputeMachineImage) State() (*computeMachineImageState, bool) {
	return cmi.state, cmi.state != nil
}

// StateMust returns the state for [ComputeMachineImage]. Panics if the state is nil.
func (cmi *ComputeMachineImage) StateMust() *computeMachineImageState {
	if cmi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmi.Type(), cmi.LocalName()))
	}
	return cmi.state
}

// ComputeMachineImageArgs contains the configurations for google_compute_machine_image.
type ComputeMachineImageArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// GuestFlush: bool, optional
	GuestFlush terra.BoolValue `hcl:"guest_flush,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SourceInstance: string, required
	SourceInstance terra.StringValue `hcl:"source_instance,attr" validate:"required"`
	// MachineImageEncryptionKey: optional
	MachineImageEncryptionKey *computemachineimage.MachineImageEncryptionKey `hcl:"machine_image_encryption_key,block"`
	// Timeouts: optional
	Timeouts *computemachineimage.Timeouts `hcl:"timeouts,block"`
}
type computeMachineImageAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_machine_image.
func (cmi computeMachineImageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmi.ref.Append("description"))
}

// GuestFlush returns a reference to field guest_flush of google_compute_machine_image.
func (cmi computeMachineImageAttributes) GuestFlush() terra.BoolValue {
	return terra.ReferenceAsBool(cmi.ref.Append("guest_flush"))
}

// Id returns a reference to field id of google_compute_machine_image.
func (cmi computeMachineImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmi.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_machine_image.
func (cmi computeMachineImageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmi.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_machine_image.
func (cmi computeMachineImageAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmi.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_machine_image.
func (cmi computeMachineImageAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cmi.ref.Append("self_link"))
}

// SourceInstance returns a reference to field source_instance of google_compute_machine_image.
func (cmi computeMachineImageAttributes) SourceInstance() terra.StringValue {
	return terra.ReferenceAsString(cmi.ref.Append("source_instance"))
}

// StorageLocations returns a reference to field storage_locations of google_compute_machine_image.
func (cmi computeMachineImageAttributes) StorageLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmi.ref.Append("storage_locations"))
}

func (cmi computeMachineImageAttributes) MachineImageEncryptionKey() terra.ListValue[computemachineimage.MachineImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computemachineimage.MachineImageEncryptionKeyAttributes](cmi.ref.Append("machine_image_encryption_key"))
}

func (cmi computeMachineImageAttributes) Timeouts() computemachineimage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computemachineimage.TimeoutsAttributes](cmi.ref.Append("timeouts"))
}

type computeMachineImageState struct {
	Description               string                                               `json:"description"`
	GuestFlush                bool                                                 `json:"guest_flush"`
	Id                        string                                               `json:"id"`
	Name                      string                                               `json:"name"`
	Project                   string                                               `json:"project"`
	SelfLink                  string                                               `json:"self_link"`
	SourceInstance            string                                               `json:"source_instance"`
	StorageLocations          []string                                             `json:"storage_locations"`
	MachineImageEncryptionKey []computemachineimage.MachineImageEncryptionKeyState `json:"machine_image_encryption_key"`
	Timeouts                  *computemachineimage.TimeoutsState                   `json:"timeouts"`
}
