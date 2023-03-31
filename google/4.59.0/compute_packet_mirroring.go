// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computepacketmirroring "github.com/golingon/terraproviders/google/4.59.0/computepacketmirroring"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputePacketMirroring(name string, args ComputePacketMirroringArgs) *ComputePacketMirroring {
	return &ComputePacketMirroring{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputePacketMirroring)(nil)

type ComputePacketMirroring struct {
	Name  string
	Args  ComputePacketMirroringArgs
	state *computePacketMirroringState
}

func (cpm *ComputePacketMirroring) Type() string {
	return "google_compute_packet_mirroring"
}

func (cpm *ComputePacketMirroring) LocalName() string {
	return cpm.Name
}

func (cpm *ComputePacketMirroring) Configuration() interface{} {
	return cpm.Args
}

func (cpm *ComputePacketMirroring) Attributes() computePacketMirroringAttributes {
	return computePacketMirroringAttributes{ref: terra.ReferenceResource(cpm)}
}

func (cpm *ComputePacketMirroring) ImportState(av io.Reader) error {
	cpm.state = &computePacketMirroringState{}
	if err := json.NewDecoder(av).Decode(cpm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpm.Type(), cpm.LocalName(), err)
	}
	return nil
}

func (cpm *ComputePacketMirroring) State() (*computePacketMirroringState, bool) {
	return cpm.state, cpm.state != nil
}

func (cpm *ComputePacketMirroring) StateMust() *computePacketMirroringState {
	if cpm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpm.Type(), cpm.LocalName()))
	}
	return cpm.state
}

func (cpm *ComputePacketMirroring) DependOn() terra.Reference {
	return terra.ReferenceResource(cpm)
}

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
	// DependsOn contains resources that ComputePacketMirroring depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computePacketMirroringAttributes struct {
	ref terra.Reference
}

func (cpm computePacketMirroringAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cpm.ref.Append("description"))
}

func (cpm computePacketMirroringAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cpm.ref.Append("id"))
}

func (cpm computePacketMirroringAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cpm.ref.Append("name"))
}

func (cpm computePacketMirroringAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(cpm.ref.Append("priority"))
}

func (cpm computePacketMirroringAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cpm.ref.Append("project"))
}

func (cpm computePacketMirroringAttributes) Region() terra.StringValue {
	return terra.ReferenceString(cpm.ref.Append("region"))
}

func (cpm computePacketMirroringAttributes) CollectorIlb() terra.ListValue[computepacketmirroring.CollectorIlbAttributes] {
	return terra.ReferenceList[computepacketmirroring.CollectorIlbAttributes](cpm.ref.Append("collector_ilb"))
}

func (cpm computePacketMirroringAttributes) Filter() terra.ListValue[computepacketmirroring.FilterAttributes] {
	return terra.ReferenceList[computepacketmirroring.FilterAttributes](cpm.ref.Append("filter"))
}

func (cpm computePacketMirroringAttributes) MirroredResources() terra.ListValue[computepacketmirroring.MirroredResourcesAttributes] {
	return terra.ReferenceList[computepacketmirroring.MirroredResourcesAttributes](cpm.ref.Append("mirrored_resources"))
}

func (cpm computePacketMirroringAttributes) Network() terra.ListValue[computepacketmirroring.NetworkAttributes] {
	return terra.ReferenceList[computepacketmirroring.NetworkAttributes](cpm.ref.Append("network"))
}

func (cpm computePacketMirroringAttributes) Timeouts() computepacketmirroring.TimeoutsAttributes {
	return terra.ReferenceSingle[computepacketmirroring.TimeoutsAttributes](cpm.ref.Append("timeouts"))
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
