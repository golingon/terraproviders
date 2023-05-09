// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeattacheddisk "github.com/golingon/terraproviders/google/4.63.1/computeattacheddisk"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeAttachedDisk creates a new instance of [ComputeAttachedDisk].
func NewComputeAttachedDisk(name string, args ComputeAttachedDiskArgs) *ComputeAttachedDisk {
	return &ComputeAttachedDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeAttachedDisk)(nil)

// ComputeAttachedDisk represents the Terraform resource google_compute_attached_disk.
type ComputeAttachedDisk struct {
	Name      string
	Args      ComputeAttachedDiskArgs
	state     *computeAttachedDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeAttachedDisk].
func (cad *ComputeAttachedDisk) Type() string {
	return "google_compute_attached_disk"
}

// LocalName returns the local name for [ComputeAttachedDisk].
func (cad *ComputeAttachedDisk) LocalName() string {
	return cad.Name
}

// Configuration returns the configuration (args) for [ComputeAttachedDisk].
func (cad *ComputeAttachedDisk) Configuration() interface{} {
	return cad.Args
}

// DependOn is used for other resources to depend on [ComputeAttachedDisk].
func (cad *ComputeAttachedDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(cad)
}

// Dependencies returns the list of resources [ComputeAttachedDisk] depends_on.
func (cad *ComputeAttachedDisk) Dependencies() terra.Dependencies {
	return cad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeAttachedDisk].
func (cad *ComputeAttachedDisk) LifecycleManagement() *terra.Lifecycle {
	return cad.Lifecycle
}

// Attributes returns the attributes for [ComputeAttachedDisk].
func (cad *ComputeAttachedDisk) Attributes() computeAttachedDiskAttributes {
	return computeAttachedDiskAttributes{ref: terra.ReferenceResource(cad)}
}

// ImportState imports the given attribute values into [ComputeAttachedDisk]'s state.
func (cad *ComputeAttachedDisk) ImportState(av io.Reader) error {
	cad.state = &computeAttachedDiskState{}
	if err := json.NewDecoder(av).Decode(cad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cad.Type(), cad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeAttachedDisk] has state.
func (cad *ComputeAttachedDisk) State() (*computeAttachedDiskState, bool) {
	return cad.state, cad.state != nil
}

// StateMust returns the state for [ComputeAttachedDisk]. Panics if the state is nil.
func (cad *ComputeAttachedDisk) StateMust() *computeAttachedDiskState {
	if cad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cad.Type(), cad.LocalName()))
	}
	return cad.state
}

// ComputeAttachedDiskArgs contains the configurations for google_compute_attached_disk.
type ComputeAttachedDiskArgs struct {
	// DeviceName: string, optional
	DeviceName terra.StringValue `hcl:"device_name,attr"`
	// Disk: string, required
	Disk terra.StringValue `hcl:"disk,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computeattacheddisk.Timeouts `hcl:"timeouts,block"`
}
type computeAttachedDiskAttributes struct {
	ref terra.Reference
}

// DeviceName returns a reference to field device_name of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("device_name"))
}

// Disk returns a reference to field disk of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("disk"))
}

// Id returns a reference to field id of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("id"))
}

// Instance returns a reference to field instance of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("instance"))
}

// Mode returns a reference to field mode of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("mode"))
}

// Project returns a reference to field project of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_attached_disk.
func (cad computeAttachedDiskAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cad.ref.Append("zone"))
}

func (cad computeAttachedDiskAttributes) Timeouts() computeattacheddisk.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeattacheddisk.TimeoutsAttributes](cad.ref.Append("timeouts"))
}

type computeAttachedDiskState struct {
	DeviceName string                             `json:"device_name"`
	Disk       string                             `json:"disk"`
	Id         string                             `json:"id"`
	Instance   string                             `json:"instance"`
	Mode       string                             `json:"mode"`
	Project    string                             `json:"project"`
	Zone       string                             `json:"zone"`
	Timeouts   *computeattacheddisk.TimeoutsState `json:"timeouts"`
}
