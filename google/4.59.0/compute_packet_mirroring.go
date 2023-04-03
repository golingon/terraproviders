// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computepacketmirroring "github.com/golingon/terraproviders/google/4.59.0/computepacketmirroring"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputePacketMirroring creates a new instance of [ComputePacketMirroring].
func NewComputePacketMirroring(name string, args ComputePacketMirroringArgs) *ComputePacketMirroring {
	return &ComputePacketMirroring{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputePacketMirroring)(nil)

// ComputePacketMirroring represents the Terraform resource google_compute_packet_mirroring.
type ComputePacketMirroring struct {
	Name      string
	Args      ComputePacketMirroringArgs
	state     *computePacketMirroringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputePacketMirroring].
func (cpm *ComputePacketMirroring) Type() string {
	return "google_compute_packet_mirroring"
}

// LocalName returns the local name for [ComputePacketMirroring].
func (cpm *ComputePacketMirroring) LocalName() string {
	return cpm.Name
}

// Configuration returns the configuration (args) for [ComputePacketMirroring].
func (cpm *ComputePacketMirroring) Configuration() interface{} {
	return cpm.Args
}

// DependOn is used for other resources to depend on [ComputePacketMirroring].
func (cpm *ComputePacketMirroring) DependOn() terra.Reference {
	return terra.ReferenceResource(cpm)
}

// Dependencies returns the list of resources [ComputePacketMirroring] depends_on.
func (cpm *ComputePacketMirroring) Dependencies() terra.Dependencies {
	return cpm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputePacketMirroring].
func (cpm *ComputePacketMirroring) LifecycleManagement() *terra.Lifecycle {
	return cpm.Lifecycle
}

// Attributes returns the attributes for [ComputePacketMirroring].
func (cpm *ComputePacketMirroring) Attributes() computePacketMirroringAttributes {
	return computePacketMirroringAttributes{ref: terra.ReferenceResource(cpm)}
}

// ImportState imports the given attribute values into [ComputePacketMirroring]'s state.
func (cpm *ComputePacketMirroring) ImportState(av io.Reader) error {
	cpm.state = &computePacketMirroringState{}
	if err := json.NewDecoder(av).Decode(cpm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpm.Type(), cpm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputePacketMirroring] has state.
func (cpm *ComputePacketMirroring) State() (*computePacketMirroringState, bool) {
	return cpm.state, cpm.state != nil
}

// StateMust returns the state for [ComputePacketMirroring]. Panics if the state is nil.
func (cpm *ComputePacketMirroring) StateMust() *computePacketMirroringState {
	if cpm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpm.Type(), cpm.LocalName()))
	}
	return cpm.state
}

// ComputePacketMirroringArgs contains the configurations for google_compute_packet_mirroring.
type ComputePacketMirroringArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// CollectorIlb: required
	CollectorIlb *computepacketmirroring.CollectorIlb `hcl:"collector_ilb,block" validate:"required"`
	// Filter: optional
	Filter *computepacketmirroring.Filter `hcl:"filter,block"`
	// MirroredResources: required
	MirroredResources *computepacketmirroring.MirroredResources `hcl:"mirrored_resources,block" validate:"required"`
	// Network: required
	Network *computepacketmirroring.Network `hcl:"network,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computepacketmirroring.Timeouts `hcl:"timeouts,block"`
}
type computePacketMirroringAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_packet_mirroring.
func (cpm computePacketMirroringAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_packet_mirroring.
func (cpm computePacketMirroringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_packet_mirroring.
func (cpm computePacketMirroringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("name"))
}

// Priority returns a reference to field priority of google_compute_packet_mirroring.
func (cpm computePacketMirroringAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cpm.ref.Append("priority"))
}

// Project returns a reference to field project of google_compute_packet_mirroring.
func (cpm computePacketMirroringAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_packet_mirroring.
func (cpm computePacketMirroringAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("region"))
}

func (cpm computePacketMirroringAttributes) CollectorIlb() terra.ListValue[computepacketmirroring.CollectorIlbAttributes] {
	return terra.ReferenceAsList[computepacketmirroring.CollectorIlbAttributes](cpm.ref.Append("collector_ilb"))
}

func (cpm computePacketMirroringAttributes) Filter() terra.ListValue[computepacketmirroring.FilterAttributes] {
	return terra.ReferenceAsList[computepacketmirroring.FilterAttributes](cpm.ref.Append("filter"))
}

func (cpm computePacketMirroringAttributes) MirroredResources() terra.ListValue[computepacketmirroring.MirroredResourcesAttributes] {
	return terra.ReferenceAsList[computepacketmirroring.MirroredResourcesAttributes](cpm.ref.Append("mirrored_resources"))
}

func (cpm computePacketMirroringAttributes) Network() terra.ListValue[computepacketmirroring.NetworkAttributes] {
	return terra.ReferenceAsList[computepacketmirroring.NetworkAttributes](cpm.ref.Append("network"))
}

func (cpm computePacketMirroringAttributes) Timeouts() computepacketmirroring.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computepacketmirroring.TimeoutsAttributes](cpm.ref.Append("timeouts"))
}

type computePacketMirroringState struct {
	Description       string                                          `json:"description"`
	Id                string                                          `json:"id"`
	Name              string                                          `json:"name"`
	Priority          float64                                         `json:"priority"`
	Project           string                                          `json:"project"`
	Region            string                                          `json:"region"`
	CollectorIlb      []computepacketmirroring.CollectorIlbState      `json:"collector_ilb"`
	Filter            []computepacketmirroring.FilterState            `json:"filter"`
	MirroredResources []computepacketmirroring.MirroredResourcesState `json:"mirrored_resources"`
	Network           []computepacketmirroring.NetworkState           `json:"network"`
	Timeouts          *computepacketmirroring.TimeoutsState           `json:"timeouts"`
}
