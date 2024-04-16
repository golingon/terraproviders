// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_access_context_manager_egress_policy

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

// Resource represents the Terraform resource google_access_context_manager_egress_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleAccessContextManagerEgressPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gacmep *Resource) Type() string {
	return "google_access_context_manager_egress_policy"
}

// LocalName returns the local name for [Resource].
func (gacmep *Resource) LocalName() string {
	return gacmep.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gacmep *Resource) Configuration() interface{} {
	return gacmep.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gacmep *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gacmep)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gacmep *Resource) Dependencies() terra.Dependencies {
	return gacmep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gacmep *Resource) LifecycleManagement() *terra.Lifecycle {
	return gacmep.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gacmep *Resource) Attributes() googleAccessContextManagerEgressPolicyAttributes {
	return googleAccessContextManagerEgressPolicyAttributes{ref: terra.ReferenceResource(gacmep)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gacmep *Resource) ImportState(state io.Reader) error {
	gacmep.state = &googleAccessContextManagerEgressPolicyState{}
	if err := json.NewDecoder(state).Decode(gacmep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gacmep.Type(), gacmep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gacmep *Resource) State() (*googleAccessContextManagerEgressPolicyState, bool) {
	return gacmep.state, gacmep.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gacmep *Resource) StateMust() *googleAccessContextManagerEgressPolicyState {
	if gacmep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gacmep.Type(), gacmep.LocalName()))
	}
	return gacmep.state
}

// Args contains the configurations for google_access_context_manager_egress_policy.
type Args struct {
	// EgressPolicyName: string, required
	EgressPolicyName terra.StringValue `hcl:"egress_policy_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleAccessContextManagerEgressPolicyAttributes struct {
	ref terra.Reference
}

// EgressPolicyName returns a reference to field egress_policy_name of google_access_context_manager_egress_policy.
func (gacmep googleAccessContextManagerEgressPolicyAttributes) EgressPolicyName() terra.StringValue {
	return terra.ReferenceAsString(gacmep.ref.Append("egress_policy_name"))
}

// Id returns a reference to field id of google_access_context_manager_egress_policy.
func (gacmep googleAccessContextManagerEgressPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gacmep.ref.Append("id"))
}

// Resource returns a reference to field resource of google_access_context_manager_egress_policy.
func (gacmep googleAccessContextManagerEgressPolicyAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(gacmep.ref.Append("resource"))
}

func (gacmep googleAccessContextManagerEgressPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gacmep.ref.Append("timeouts"))
}

type googleAccessContextManagerEgressPolicyState struct {
	EgressPolicyName string         `json:"egress_policy_name"`
	Id               string         `json:"id"`
	Resource         string         `json:"resource"`
	Timeouts         *TimeoutsState `json:"timeouts"`
}
