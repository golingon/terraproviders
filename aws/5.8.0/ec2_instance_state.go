// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2instancestate "github.com/golingon/terraproviders/aws/5.8.0/ec2instancestate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2InstanceState creates a new instance of [Ec2InstanceState].
func NewEc2InstanceState(name string, args Ec2InstanceStateArgs) *Ec2InstanceState {
	return &Ec2InstanceState{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2InstanceState)(nil)

// Ec2InstanceState represents the Terraform resource aws_ec2_instance_state.
type Ec2InstanceState struct {
	Name      string
	Args      Ec2InstanceStateArgs
	state     *ec2InstanceStateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2InstanceState].
func (eis *Ec2InstanceState) Type() string {
	return "aws_ec2_instance_state"
}

// LocalName returns the local name for [Ec2InstanceState].
func (eis *Ec2InstanceState) LocalName() string {
	return eis.Name
}

// Configuration returns the configuration (args) for [Ec2InstanceState].
func (eis *Ec2InstanceState) Configuration() interface{} {
	return eis.Args
}

// DependOn is used for other resources to depend on [Ec2InstanceState].
func (eis *Ec2InstanceState) DependOn() terra.Reference {
	return terra.ReferenceResource(eis)
}

// Dependencies returns the list of resources [Ec2InstanceState] depends_on.
func (eis *Ec2InstanceState) Dependencies() terra.Dependencies {
	return eis.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2InstanceState].
func (eis *Ec2InstanceState) LifecycleManagement() *terra.Lifecycle {
	return eis.Lifecycle
}

// Attributes returns the attributes for [Ec2InstanceState].
func (eis *Ec2InstanceState) Attributes() ec2InstanceStateAttributes {
	return ec2InstanceStateAttributes{ref: terra.ReferenceResource(eis)}
}

// ImportState imports the given attribute values into [Ec2InstanceState]'s state.
func (eis *Ec2InstanceState) ImportState(av io.Reader) error {
	eis.state = &ec2InstanceStateState{}
	if err := json.NewDecoder(av).Decode(eis.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eis.Type(), eis.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2InstanceState] has state.
func (eis *Ec2InstanceState) State() (*ec2InstanceStateState, bool) {
	return eis.state, eis.state != nil
}

// StateMust returns the state for [Ec2InstanceState]. Panics if the state is nil.
func (eis *Ec2InstanceState) StateMust() *ec2InstanceStateState {
	if eis.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eis.Type(), eis.LocalName()))
	}
	return eis.state
}

// Ec2InstanceStateArgs contains the configurations for aws_ec2_instance_state.
type Ec2InstanceStateArgs struct {
	// Force: bool, optional
	Force terra.BoolValue `hcl:"force,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// State: string, required
	State terra.StringValue `hcl:"state,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2instancestate.Timeouts `hcl:"timeouts,block"`
}
type ec2InstanceStateAttributes struct {
	ref terra.Reference
}

// Force returns a reference to field force of aws_ec2_instance_state.
func (eis ec2InstanceStateAttributes) Force() terra.BoolValue {
	return terra.ReferenceAsBool(eis.ref.Append("force"))
}

// Id returns a reference to field id of aws_ec2_instance_state.
func (eis ec2InstanceStateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eis.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_ec2_instance_state.
func (eis ec2InstanceStateAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(eis.ref.Append("instance_id"))
}

// State returns a reference to field state of aws_ec2_instance_state.
func (eis ec2InstanceStateAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(eis.ref.Append("state"))
}

func (eis ec2InstanceStateAttributes) Timeouts() ec2instancestate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2instancestate.TimeoutsAttributes](eis.ref.Append("timeouts"))
}

type ec2InstanceStateState struct {
	Force      bool                            `json:"force"`
	Id         string                          `json:"id"`
	InstanceId string                          `json:"instance_id"`
	State      string                          `json:"state"`
	Timeouts   *ec2instancestate.TimeoutsState `json:"timeouts"`
}
