// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	globalacceleratorcustomroutinglistener "github.com/golingon/terraproviders/aws/5.6.2/globalacceleratorcustomroutinglistener"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlobalacceleratorCustomRoutingListener creates a new instance of [GlobalacceleratorCustomRoutingListener].
func NewGlobalacceleratorCustomRoutingListener(name string, args GlobalacceleratorCustomRoutingListenerArgs) *GlobalacceleratorCustomRoutingListener {
	return &GlobalacceleratorCustomRoutingListener{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlobalacceleratorCustomRoutingListener)(nil)

// GlobalacceleratorCustomRoutingListener represents the Terraform resource aws_globalaccelerator_custom_routing_listener.
type GlobalacceleratorCustomRoutingListener struct {
	Name      string
	Args      GlobalacceleratorCustomRoutingListenerArgs
	state     *globalacceleratorCustomRoutingListenerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlobalacceleratorCustomRoutingListener].
func (gcrl *GlobalacceleratorCustomRoutingListener) Type() string {
	return "aws_globalaccelerator_custom_routing_listener"
}

// LocalName returns the local name for [GlobalacceleratorCustomRoutingListener].
func (gcrl *GlobalacceleratorCustomRoutingListener) LocalName() string {
	return gcrl.Name
}

// Configuration returns the configuration (args) for [GlobalacceleratorCustomRoutingListener].
func (gcrl *GlobalacceleratorCustomRoutingListener) Configuration() interface{} {
	return gcrl.Args
}

// DependOn is used for other resources to depend on [GlobalacceleratorCustomRoutingListener].
func (gcrl *GlobalacceleratorCustomRoutingListener) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrl)
}

// Dependencies returns the list of resources [GlobalacceleratorCustomRoutingListener] depends_on.
func (gcrl *GlobalacceleratorCustomRoutingListener) Dependencies() terra.Dependencies {
	return gcrl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlobalacceleratorCustomRoutingListener].
func (gcrl *GlobalacceleratorCustomRoutingListener) LifecycleManagement() *terra.Lifecycle {
	return gcrl.Lifecycle
}

// Attributes returns the attributes for [GlobalacceleratorCustomRoutingListener].
func (gcrl *GlobalacceleratorCustomRoutingListener) Attributes() globalacceleratorCustomRoutingListenerAttributes {
	return globalacceleratorCustomRoutingListenerAttributes{ref: terra.ReferenceResource(gcrl)}
}

// ImportState imports the given attribute values into [GlobalacceleratorCustomRoutingListener]'s state.
func (gcrl *GlobalacceleratorCustomRoutingListener) ImportState(av io.Reader) error {
	gcrl.state = &globalacceleratorCustomRoutingListenerState{}
	if err := json.NewDecoder(av).Decode(gcrl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrl.Type(), gcrl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlobalacceleratorCustomRoutingListener] has state.
func (gcrl *GlobalacceleratorCustomRoutingListener) State() (*globalacceleratorCustomRoutingListenerState, bool) {
	return gcrl.state, gcrl.state != nil
}

// StateMust returns the state for [GlobalacceleratorCustomRoutingListener]. Panics if the state is nil.
func (gcrl *GlobalacceleratorCustomRoutingListener) StateMust() *globalacceleratorCustomRoutingListenerState {
	if gcrl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrl.Type(), gcrl.LocalName()))
	}
	return gcrl.state
}

// GlobalacceleratorCustomRoutingListenerArgs contains the configurations for aws_globalaccelerator_custom_routing_listener.
type GlobalacceleratorCustomRoutingListenerArgs struct {
	// AcceleratorArn: string, required
	AcceleratorArn terra.StringValue `hcl:"accelerator_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PortRange: min=1,max=10
	PortRange []globalacceleratorcustomroutinglistener.PortRange `hcl:"port_range,block" validate:"min=1,max=10"`
	// Timeouts: optional
	Timeouts *globalacceleratorcustomroutinglistener.Timeouts `hcl:"timeouts,block"`
}
type globalacceleratorCustomRoutingListenerAttributes struct {
	ref terra.Reference
}

// AcceleratorArn returns a reference to field accelerator_arn of aws_globalaccelerator_custom_routing_listener.
func (gcrl globalacceleratorCustomRoutingListenerAttributes) AcceleratorArn() terra.StringValue {
	return terra.ReferenceAsString(gcrl.ref.Append("accelerator_arn"))
}

// Id returns a reference to field id of aws_globalaccelerator_custom_routing_listener.
func (gcrl globalacceleratorCustomRoutingListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrl.ref.Append("id"))
}

func (gcrl globalacceleratorCustomRoutingListenerAttributes) PortRange() terra.SetValue[globalacceleratorcustomroutinglistener.PortRangeAttributes] {
	return terra.ReferenceAsSet[globalacceleratorcustomroutinglistener.PortRangeAttributes](gcrl.ref.Append("port_range"))
}

func (gcrl globalacceleratorCustomRoutingListenerAttributes) Timeouts() globalacceleratorcustomroutinglistener.TimeoutsAttributes {
	return terra.ReferenceAsSingle[globalacceleratorcustomroutinglistener.TimeoutsAttributes](gcrl.ref.Append("timeouts"))
}

type globalacceleratorCustomRoutingListenerState struct {
	AcceleratorArn string                                                  `json:"accelerator_arn"`
	Id             string                                                  `json:"id"`
	PortRange      []globalacceleratorcustomroutinglistener.PortRangeState `json:"port_range"`
	Timeouts       *globalacceleratorcustomroutinglistener.TimeoutsState   `json:"timeouts"`
}
