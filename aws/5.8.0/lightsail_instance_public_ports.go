// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lightsailinstancepublicports "github.com/golingon/terraproviders/aws/5.8.0/lightsailinstancepublicports"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailInstancePublicPorts creates a new instance of [LightsailInstancePublicPorts].
func NewLightsailInstancePublicPorts(name string, args LightsailInstancePublicPortsArgs) *LightsailInstancePublicPorts {
	return &LightsailInstancePublicPorts{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailInstancePublicPorts)(nil)

// LightsailInstancePublicPorts represents the Terraform resource aws_lightsail_instance_public_ports.
type LightsailInstancePublicPorts struct {
	Name      string
	Args      LightsailInstancePublicPortsArgs
	state     *lightsailInstancePublicPortsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailInstancePublicPorts].
func (lipp *LightsailInstancePublicPorts) Type() string {
	return "aws_lightsail_instance_public_ports"
}

// LocalName returns the local name for [LightsailInstancePublicPorts].
func (lipp *LightsailInstancePublicPorts) LocalName() string {
	return lipp.Name
}

// Configuration returns the configuration (args) for [LightsailInstancePublicPorts].
func (lipp *LightsailInstancePublicPorts) Configuration() interface{} {
	return lipp.Args
}

// DependOn is used for other resources to depend on [LightsailInstancePublicPorts].
func (lipp *LightsailInstancePublicPorts) DependOn() terra.Reference {
	return terra.ReferenceResource(lipp)
}

// Dependencies returns the list of resources [LightsailInstancePublicPorts] depends_on.
func (lipp *LightsailInstancePublicPorts) Dependencies() terra.Dependencies {
	return lipp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailInstancePublicPorts].
func (lipp *LightsailInstancePublicPorts) LifecycleManagement() *terra.Lifecycle {
	return lipp.Lifecycle
}

// Attributes returns the attributes for [LightsailInstancePublicPorts].
func (lipp *LightsailInstancePublicPorts) Attributes() lightsailInstancePublicPortsAttributes {
	return lightsailInstancePublicPortsAttributes{ref: terra.ReferenceResource(lipp)}
}

// ImportState imports the given attribute values into [LightsailInstancePublicPorts]'s state.
func (lipp *LightsailInstancePublicPorts) ImportState(av io.Reader) error {
	lipp.state = &lightsailInstancePublicPortsState{}
	if err := json.NewDecoder(av).Decode(lipp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lipp.Type(), lipp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailInstancePublicPorts] has state.
func (lipp *LightsailInstancePublicPorts) State() (*lightsailInstancePublicPortsState, bool) {
	return lipp.state, lipp.state != nil
}

// StateMust returns the state for [LightsailInstancePublicPorts]. Panics if the state is nil.
func (lipp *LightsailInstancePublicPorts) StateMust() *lightsailInstancePublicPortsState {
	if lipp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lipp.Type(), lipp.LocalName()))
	}
	return lipp.state
}

// LightsailInstancePublicPortsArgs contains the configurations for aws_lightsail_instance_public_ports.
type LightsailInstancePublicPortsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// PortInfo: min=1
	PortInfo []lightsailinstancepublicports.PortInfo `hcl:"port_info,block" validate:"min=1"`
}
type lightsailInstancePublicPortsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_lightsail_instance_public_ports.
func (lipp lightsailInstancePublicPortsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lipp.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of aws_lightsail_instance_public_ports.
func (lipp lightsailInstancePublicPortsAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(lipp.ref.Append("instance_name"))
}

func (lipp lightsailInstancePublicPortsAttributes) PortInfo() terra.SetValue[lightsailinstancepublicports.PortInfoAttributes] {
	return terra.ReferenceAsSet[lightsailinstancepublicports.PortInfoAttributes](lipp.ref.Append("port_info"))
}

type lightsailInstancePublicPortsState struct {
	Id           string                                       `json:"id"`
	InstanceName string                                       `json:"instance_name"`
	PortInfo     []lightsailinstancepublicports.PortInfoState `json:"port_info"`
}
