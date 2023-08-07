// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinstancegroupnamedport "github.com/golingon/terraproviders/google/4.76.0/computeinstancegroupnamedport"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceGroupNamedPort creates a new instance of [ComputeInstanceGroupNamedPort].
func NewComputeInstanceGroupNamedPort(name string, args ComputeInstanceGroupNamedPortArgs) *ComputeInstanceGroupNamedPort {
	return &ComputeInstanceGroupNamedPort{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceGroupNamedPort)(nil)

// ComputeInstanceGroupNamedPort represents the Terraform resource google_compute_instance_group_named_port.
type ComputeInstanceGroupNamedPort struct {
	Name      string
	Args      ComputeInstanceGroupNamedPortArgs
	state     *computeInstanceGroupNamedPortState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceGroupNamedPort].
func (cignp *ComputeInstanceGroupNamedPort) Type() string {
	return "google_compute_instance_group_named_port"
}

// LocalName returns the local name for [ComputeInstanceGroupNamedPort].
func (cignp *ComputeInstanceGroupNamedPort) LocalName() string {
	return cignp.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceGroupNamedPort].
func (cignp *ComputeInstanceGroupNamedPort) Configuration() interface{} {
	return cignp.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceGroupNamedPort].
func (cignp *ComputeInstanceGroupNamedPort) DependOn() terra.Reference {
	return terra.ReferenceResource(cignp)
}

// Dependencies returns the list of resources [ComputeInstanceGroupNamedPort] depends_on.
func (cignp *ComputeInstanceGroupNamedPort) Dependencies() terra.Dependencies {
	return cignp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceGroupNamedPort].
func (cignp *ComputeInstanceGroupNamedPort) LifecycleManagement() *terra.Lifecycle {
	return cignp.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceGroupNamedPort].
func (cignp *ComputeInstanceGroupNamedPort) Attributes() computeInstanceGroupNamedPortAttributes {
	return computeInstanceGroupNamedPortAttributes{ref: terra.ReferenceResource(cignp)}
}

// ImportState imports the given attribute values into [ComputeInstanceGroupNamedPort]'s state.
func (cignp *ComputeInstanceGroupNamedPort) ImportState(av io.Reader) error {
	cignp.state = &computeInstanceGroupNamedPortState{}
	if err := json.NewDecoder(av).Decode(cignp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cignp.Type(), cignp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceGroupNamedPort] has state.
func (cignp *ComputeInstanceGroupNamedPort) State() (*computeInstanceGroupNamedPortState, bool) {
	return cignp.state, cignp.state != nil
}

// StateMust returns the state for [ComputeInstanceGroupNamedPort]. Panics if the state is nil.
func (cignp *ComputeInstanceGroupNamedPort) StateMust() *computeInstanceGroupNamedPortState {
	if cignp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cignp.Type(), cignp.LocalName()))
	}
	return cignp.state
}

// ComputeInstanceGroupNamedPortArgs contains the configurations for google_compute_instance_group_named_port.
type ComputeInstanceGroupNamedPortArgs struct {
	// Group: string, required
	Group terra.StringValue `hcl:"group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computeinstancegroupnamedport.Timeouts `hcl:"timeouts,block"`
}
type computeInstanceGroupNamedPortAttributes struct {
	ref terra.Reference
}

// Group returns a reference to field group of google_compute_instance_group_named_port.
func (cignp computeInstanceGroupNamedPortAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(cignp.ref.Append("group"))
}

// Id returns a reference to field id of google_compute_instance_group_named_port.
func (cignp computeInstanceGroupNamedPortAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cignp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_instance_group_named_port.
func (cignp computeInstanceGroupNamedPortAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cignp.ref.Append("name"))
}

// Port returns a reference to field port of google_compute_instance_group_named_port.
func (cignp computeInstanceGroupNamedPortAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cignp.ref.Append("port"))
}

// Project returns a reference to field project of google_compute_instance_group_named_port.
func (cignp computeInstanceGroupNamedPortAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cignp.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_instance_group_named_port.
func (cignp computeInstanceGroupNamedPortAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cignp.ref.Append("zone"))
}

func (cignp computeInstanceGroupNamedPortAttributes) Timeouts() computeinstancegroupnamedport.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinstancegroupnamedport.TimeoutsAttributes](cignp.ref.Append("timeouts"))
}

type computeInstanceGroupNamedPortState struct {
	Group    string                                       `json:"group"`
	Id       string                                       `json:"id"`
	Name     string                                       `json:"name"`
	Port     float64                                      `json:"port"`
	Project  string                                       `json:"project"`
	Zone     string                                       `json:"zone"`
	Timeouts *computeinstancegroupnamedport.TimeoutsState `json:"timeouts"`
}
