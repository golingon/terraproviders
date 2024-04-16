// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_endpoints_service_iam_binding

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

// Resource represents the Terraform resource google_endpoints_service_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleEndpointsServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gesib *Resource) Type() string {
	return "google_endpoints_service_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gesib *Resource) LocalName() string {
	return gesib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gesib *Resource) Configuration() interface{} {
	return gesib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gesib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gesib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gesib *Resource) Dependencies() terra.Dependencies {
	return gesib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gesib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gesib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gesib *Resource) Attributes() googleEndpointsServiceIamBindingAttributes {
	return googleEndpointsServiceIamBindingAttributes{ref: terra.ReferenceResource(gesib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gesib *Resource) ImportState(state io.Reader) error {
	gesib.state = &googleEndpointsServiceIamBindingState{}
	if err := json.NewDecoder(state).Decode(gesib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gesib.Type(), gesib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gesib *Resource) State() (*googleEndpointsServiceIamBindingState, bool) {
	return gesib.state, gesib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gesib *Resource) StateMust() *googleEndpointsServiceIamBindingState {
	if gesib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gesib.Type(), gesib.LocalName()))
	}
	return gesib.state
}

// Args contains the configurations for google_endpoints_service_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleEndpointsServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_endpoints_service_iam_binding.
func (gesib googleEndpointsServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gesib.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_iam_binding.
func (gesib googleEndpointsServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gesib.ref.Append("id"))
}

// Members returns a reference to field members of google_endpoints_service_iam_binding.
func (gesib googleEndpointsServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gesib.ref.Append("members"))
}

// Role returns a reference to field role of google_endpoints_service_iam_binding.
func (gesib googleEndpointsServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gesib.ref.Append("role"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_iam_binding.
func (gesib googleEndpointsServiceIamBindingAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(gesib.ref.Append("service_name"))
}

func (gesib googleEndpointsServiceIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gesib.ref.Append("condition"))
}

type googleEndpointsServiceIamBindingState struct {
	Etag        string           `json:"etag"`
	Id          string           `json:"id"`
	Members     []string         `json:"members"`
	Role        string           `json:"role"`
	ServiceName string           `json:"service_name"`
	Condition   []ConditionState `json:"condition"`
}
