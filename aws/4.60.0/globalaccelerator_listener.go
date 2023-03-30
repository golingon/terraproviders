// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	globalacceleratorlistener "github.com/golingon/terraproviders/aws/4.60.0/globalacceleratorlistener"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlobalacceleratorListener(name string, args GlobalacceleratorListenerArgs) *GlobalacceleratorListener {
	return &GlobalacceleratorListener{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlobalacceleratorListener)(nil)

type GlobalacceleratorListener struct {
	Name  string
	Args  GlobalacceleratorListenerArgs
	state *globalacceleratorListenerState
}

func (gl *GlobalacceleratorListener) Type() string {
	return "aws_globalaccelerator_listener"
}

func (gl *GlobalacceleratorListener) LocalName() string {
	return gl.Name
}

func (gl *GlobalacceleratorListener) Configuration() interface{} {
	return gl.Args
}

func (gl *GlobalacceleratorListener) Attributes() globalacceleratorListenerAttributes {
	return globalacceleratorListenerAttributes{ref: terra.ReferenceResource(gl)}
}

func (gl *GlobalacceleratorListener) ImportState(av io.Reader) error {
	gl.state = &globalacceleratorListenerState{}
	if err := json.NewDecoder(av).Decode(gl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gl.Type(), gl.LocalName(), err)
	}
	return nil
}

func (gl *GlobalacceleratorListener) State() (*globalacceleratorListenerState, bool) {
	return gl.state, gl.state != nil
}

func (gl *GlobalacceleratorListener) StateMust() *globalacceleratorListenerState {
	if gl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gl.Type(), gl.LocalName()))
	}
	return gl.state
}

func (gl *GlobalacceleratorListener) DependOn() terra.Reference {
	return terra.ReferenceResource(gl)
}

type GlobalacceleratorListenerArgs struct {
	// AcceleratorArn: string, required
	AcceleratorArn terra.StringValue `hcl:"accelerator_arn,attr" validate:"required"`
	// ClientAffinity: string, optional
	ClientAffinity terra.StringValue `hcl:"client_affinity,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// PortRange: min=1,max=10
	PortRange []globalacceleratorlistener.PortRange `hcl:"port_range,block" validate:"min=1,max=10"`
	// Timeouts: optional
	Timeouts *globalacceleratorlistener.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that GlobalacceleratorListener depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type globalacceleratorListenerAttributes struct {
	ref terra.Reference
}

func (gl globalacceleratorListenerAttributes) AcceleratorArn() terra.StringValue {
	return terra.ReferenceString(gl.ref.Append("accelerator_arn"))
}

func (gl globalacceleratorListenerAttributes) ClientAffinity() terra.StringValue {
	return terra.ReferenceString(gl.ref.Append("client_affinity"))
}

func (gl globalacceleratorListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gl.ref.Append("id"))
}

func (gl globalacceleratorListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(gl.ref.Append("protocol"))
}

func (gl globalacceleratorListenerAttributes) PortRange() terra.SetValue[globalacceleratorlistener.PortRangeAttributes] {
	return terra.ReferenceSet[globalacceleratorlistener.PortRangeAttributes](gl.ref.Append("port_range"))
}

func (gl globalacceleratorListenerAttributes) Timeouts() globalacceleratorlistener.TimeoutsAttributes {
	return terra.ReferenceSingle[globalacceleratorlistener.TimeoutsAttributes](gl.ref.Append("timeouts"))
}

type globalacceleratorListenerState struct {
	AcceleratorArn string                                     `json:"accelerator_arn"`
	ClientAffinity string                                     `json:"client_affinity"`
	Id             string                                     `json:"id"`
	Protocol       string                                     `json:"protocol"`
	PortRange      []globalacceleratorlistener.PortRangeState `json:"port_range"`
	Timeouts       *globalacceleratorlistener.TimeoutsState   `json:"timeouts"`
}
