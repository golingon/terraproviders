// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_target_instance

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_compute_target_instance.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeTargetInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcti *Resource) Type() string {
	return "google_compute_target_instance"
}

// LocalName returns the local name for [Resource].
func (gcti *Resource) LocalName() string {
	return gcti.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcti *Resource) Configuration() interface{} {
	return gcti.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcti *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcti)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcti *Resource) Dependencies() terra.Dependencies {
	return gcti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcti *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcti.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcti *Resource) Attributes() googleComputeTargetInstanceAttributes {
	return googleComputeTargetInstanceAttributes{ref: terra.ReferenceResource(gcti)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcti *Resource) ImportState(state io.Reader) error {
	gcti.state = &googleComputeTargetInstanceState{}
	if err := json.NewDecoder(state).Decode(gcti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcti.Type(), gcti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcti *Resource) State() (*googleComputeTargetInstanceState, bool) {
	return gcti.state, gcti.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcti *Resource) StateMust() *googleComputeTargetInstanceState {
	if gcti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcti.Type(), gcti.LocalName()))
	}
	return gcti.state
}

// Args contains the configurations for google_compute_target_instance.
type Args struct {
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
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeTargetInstanceAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("id"))
}

// Instance returns a reference to field instance of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("instance"))
}

// Name returns a reference to field name of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("name"))
}

// NatPolicy returns a reference to field nat_policy of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) NatPolicy() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("nat_policy"))
}

// Network returns a reference to field network of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("project"))
}

// SecurityPolicy returns a reference to field security_policy of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("security_policy"))
}

// SelfLink returns a reference to field self_link of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("self_link"))
}

// Zone returns a reference to field zone of google_compute_target_instance.
func (gcti googleComputeTargetInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(gcti.ref.Append("zone"))
}

func (gcti googleComputeTargetInstanceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcti.ref.Append("timeouts"))
}

type googleComputeTargetInstanceState struct {
	CreationTimestamp string         `json:"creation_timestamp"`
	Description       string         `json:"description"`
	Id                string         `json:"id"`
	Instance          string         `json:"instance"`
	Name              string         `json:"name"`
	NatPolicy         string         `json:"nat_policy"`
	Network           string         `json:"network"`
	Project           string         `json:"project"`
	SecurityPolicy    string         `json:"security_policy"`
	SelfLink          string         `json:"self_link"`
	Zone              string         `json:"zone"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
