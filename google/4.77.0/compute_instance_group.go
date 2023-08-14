// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinstancegroup "github.com/golingon/terraproviders/google/4.77.0/computeinstancegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceGroup creates a new instance of [ComputeInstanceGroup].
func NewComputeInstanceGroup(name string, args ComputeInstanceGroupArgs) *ComputeInstanceGroup {
	return &ComputeInstanceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceGroup)(nil)

// ComputeInstanceGroup represents the Terraform resource google_compute_instance_group.
type ComputeInstanceGroup struct {
	Name      string
	Args      ComputeInstanceGroupArgs
	state     *computeInstanceGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceGroup].
func (cig *ComputeInstanceGroup) Type() string {
	return "google_compute_instance_group"
}

// LocalName returns the local name for [ComputeInstanceGroup].
func (cig *ComputeInstanceGroup) LocalName() string {
	return cig.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceGroup].
func (cig *ComputeInstanceGroup) Configuration() interface{} {
	return cig.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceGroup].
func (cig *ComputeInstanceGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cig)
}

// Dependencies returns the list of resources [ComputeInstanceGroup] depends_on.
func (cig *ComputeInstanceGroup) Dependencies() terra.Dependencies {
	return cig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceGroup].
func (cig *ComputeInstanceGroup) LifecycleManagement() *terra.Lifecycle {
	return cig.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceGroup].
func (cig *ComputeInstanceGroup) Attributes() computeInstanceGroupAttributes {
	return computeInstanceGroupAttributes{ref: terra.ReferenceResource(cig)}
}

// ImportState imports the given attribute values into [ComputeInstanceGroup]'s state.
func (cig *ComputeInstanceGroup) ImportState(av io.Reader) error {
	cig.state = &computeInstanceGroupState{}
	if err := json.NewDecoder(av).Decode(cig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cig.Type(), cig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceGroup] has state.
func (cig *ComputeInstanceGroup) State() (*computeInstanceGroupState, bool) {
	return cig.state, cig.state != nil
}

// StateMust returns the state for [ComputeInstanceGroup]. Panics if the state is nil.
func (cig *ComputeInstanceGroup) StateMust() *computeInstanceGroupState {
	if cig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cig.Type(), cig.LocalName()))
	}
	return cig.state
}

// ComputeInstanceGroupArgs contains the configurations for google_compute_instance_group.
type ComputeInstanceGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instances: set of string, optional
	Instances terra.SetValue[terra.StringValue] `hcl:"instances,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// NamedPort: min=0
	NamedPort []computeinstancegroup.NamedPort `hcl:"named_port,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeinstancegroup.Timeouts `hcl:"timeouts,block"`
}
type computeInstanceGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("id"))
}

// Instances returns a reference to field instances of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Instances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cig.ref.Append("instances"))
}

// Name returns a reference to field name of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(cig.ref.Append("size"))
}

// Zone returns a reference to field zone of google_compute_instance_group.
func (cig computeInstanceGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("zone"))
}

func (cig computeInstanceGroupAttributes) NamedPort() terra.ListValue[computeinstancegroup.NamedPortAttributes] {
	return terra.ReferenceAsList[computeinstancegroup.NamedPortAttributes](cig.ref.Append("named_port"))
}

func (cig computeInstanceGroupAttributes) Timeouts() computeinstancegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinstancegroup.TimeoutsAttributes](cig.ref.Append("timeouts"))
}

type computeInstanceGroupState struct {
	Description string                                `json:"description"`
	Id          string                                `json:"id"`
	Instances   []string                              `json:"instances"`
	Name        string                                `json:"name"`
	Network     string                                `json:"network"`
	Project     string                                `json:"project"`
	SelfLink    string                                `json:"self_link"`
	Size        float64                               `json:"size"`
	Zone        string                                `json:"zone"`
	NamedPort   []computeinstancegroup.NamedPortState `json:"named_port"`
	Timeouts    *computeinstancegroup.TimeoutsState   `json:"timeouts"`
}
