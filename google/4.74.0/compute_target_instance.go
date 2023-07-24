// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computetargetinstance "github.com/golingon/terraproviders/google/4.74.0/computetargetinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeTargetInstance creates a new instance of [ComputeTargetInstance].
func NewComputeTargetInstance(name string, args ComputeTargetInstanceArgs) *ComputeTargetInstance {
	return &ComputeTargetInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetInstance)(nil)

// ComputeTargetInstance represents the Terraform resource google_compute_target_instance.
type ComputeTargetInstance struct {
	Name      string
	Args      ComputeTargetInstanceArgs
	state     *computeTargetInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetInstance].
func (cti *ComputeTargetInstance) Type() string {
	return "google_compute_target_instance"
}

// LocalName returns the local name for [ComputeTargetInstance].
func (cti *ComputeTargetInstance) LocalName() string {
	return cti.Name
}

// Configuration returns the configuration (args) for [ComputeTargetInstance].
func (cti *ComputeTargetInstance) Configuration() interface{} {
	return cti.Args
}

// DependOn is used for other resources to depend on [ComputeTargetInstance].
func (cti *ComputeTargetInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(cti)
}

// Dependencies returns the list of resources [ComputeTargetInstance] depends_on.
func (cti *ComputeTargetInstance) Dependencies() terra.Dependencies {
	return cti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetInstance].
func (cti *ComputeTargetInstance) LifecycleManagement() *terra.Lifecycle {
	return cti.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetInstance].
func (cti *ComputeTargetInstance) Attributes() computeTargetInstanceAttributes {
	return computeTargetInstanceAttributes{ref: terra.ReferenceResource(cti)}
}

// ImportState imports the given attribute values into [ComputeTargetInstance]'s state.
func (cti *ComputeTargetInstance) ImportState(av io.Reader) error {
	cti.state = &computeTargetInstanceState{}
	if err := json.NewDecoder(av).Decode(cti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cti.Type(), cti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetInstance] has state.
func (cti *ComputeTargetInstance) State() (*computeTargetInstanceState, bool) {
	return cti.state, cti.state != nil
}

// StateMust returns the state for [ComputeTargetInstance]. Panics if the state is nil.
func (cti *ComputeTargetInstance) StateMust() *computeTargetInstanceState {
	if cti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cti.Type(), cti.LocalName()))
	}
	return cti.state
}

// ComputeTargetInstanceArgs contains the configurations for google_compute_target_instance.
type ComputeTargetInstanceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NatPolicy: string, optional
	NatPolicy terra.StringValue `hcl:"nat_policy,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computetargetinstance.Timeouts `hcl:"timeouts,block"`
}
type computeTargetInstanceAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("id"))
}

// Instance returns a reference to field instance of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("instance"))
}

// Name returns a reference to field name of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("name"))
}

// NatPolicy returns a reference to field nat_policy of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) NatPolicy() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("nat_policy"))
}

// Project returns a reference to field project of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("self_link"))
}

// Zone returns a reference to field zone of google_compute_target_instance.
func (cti computeTargetInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("zone"))
}

func (cti computeTargetInstanceAttributes) Timeouts() computetargetinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargetinstance.TimeoutsAttributes](cti.ref.Append("timeouts"))
}

type computeTargetInstanceState struct {
	CreationTimestamp string                               `json:"creation_timestamp"`
	Description       string                               `json:"description"`
	Id                string                               `json:"id"`
	Instance          string                               `json:"instance"`
	Name              string                               `json:"name"`
	NatPolicy         string                               `json:"nat_policy"`
	Project           string                               `json:"project"`
	SelfLink          string                               `json:"self_link"`
	Zone              string                               `json:"zone"`
	Timeouts          *computetargetinstance.TimeoutsState `json:"timeouts"`
}
