// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	endpointsserviceconsumersiambinding "github.com/golingon/terraproviders/google/4.63.1/endpointsserviceconsumersiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEndpointsServiceConsumersIamBinding creates a new instance of [EndpointsServiceConsumersIamBinding].
func NewEndpointsServiceConsumersIamBinding(name string, args EndpointsServiceConsumersIamBindingArgs) *EndpointsServiceConsumersIamBinding {
	return &EndpointsServiceConsumersIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceConsumersIamBinding)(nil)

// EndpointsServiceConsumersIamBinding represents the Terraform resource google_endpoints_service_consumers_iam_binding.
type EndpointsServiceConsumersIamBinding struct {
	Name      string
	Args      EndpointsServiceConsumersIamBindingArgs
	state     *endpointsServiceConsumersIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EndpointsServiceConsumersIamBinding].
func (escib *EndpointsServiceConsumersIamBinding) Type() string {
	return "google_endpoints_service_consumers_iam_binding"
}

// LocalName returns the local name for [EndpointsServiceConsumersIamBinding].
func (escib *EndpointsServiceConsumersIamBinding) LocalName() string {
	return escib.Name
}

// Configuration returns the configuration (args) for [EndpointsServiceConsumersIamBinding].
func (escib *EndpointsServiceConsumersIamBinding) Configuration() interface{} {
	return escib.Args
}

// DependOn is used for other resources to depend on [EndpointsServiceConsumersIamBinding].
func (escib *EndpointsServiceConsumersIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(escib)
}

// Dependencies returns the list of resources [EndpointsServiceConsumersIamBinding] depends_on.
func (escib *EndpointsServiceConsumersIamBinding) Dependencies() terra.Dependencies {
	return escib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EndpointsServiceConsumersIamBinding].
func (escib *EndpointsServiceConsumersIamBinding) LifecycleManagement() *terra.Lifecycle {
	return escib.Lifecycle
}

// Attributes returns the attributes for [EndpointsServiceConsumersIamBinding].
func (escib *EndpointsServiceConsumersIamBinding) Attributes() endpointsServiceConsumersIamBindingAttributes {
	return endpointsServiceConsumersIamBindingAttributes{ref: terra.ReferenceResource(escib)}
}

// ImportState imports the given attribute values into [EndpointsServiceConsumersIamBinding]'s state.
func (escib *EndpointsServiceConsumersIamBinding) ImportState(av io.Reader) error {
	escib.state = &endpointsServiceConsumersIamBindingState{}
	if err := json.NewDecoder(av).Decode(escib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", escib.Type(), escib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EndpointsServiceConsumersIamBinding] has state.
func (escib *EndpointsServiceConsumersIamBinding) State() (*endpointsServiceConsumersIamBindingState, bool) {
	return escib.state, escib.state != nil
}

// StateMust returns the state for [EndpointsServiceConsumersIamBinding]. Panics if the state is nil.
func (escib *EndpointsServiceConsumersIamBinding) StateMust() *endpointsServiceConsumersIamBindingState {
	if escib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", escib.Type(), escib.LocalName()))
	}
	return escib.state
}

// EndpointsServiceConsumersIamBindingArgs contains the configurations for google_endpoints_service_consumers_iam_binding.
type EndpointsServiceConsumersIamBindingArgs struct {
	// ConsumerProject: string, required
	ConsumerProject terra.StringValue `hcl:"consumer_project,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Condition: optional
	Condition *endpointsserviceconsumersiambinding.Condition `hcl:"condition,block"`
}
type endpointsServiceConsumersIamBindingAttributes struct {
	ref terra.Reference
}

// ConsumerProject returns a reference to field consumer_project of google_endpoints_service_consumers_iam_binding.
func (escib endpointsServiceConsumersIamBindingAttributes) ConsumerProject() terra.StringValue {
	return terra.ReferenceAsString(escib.ref.Append("consumer_project"))
}

// Etag returns a reference to field etag of google_endpoints_service_consumers_iam_binding.
func (escib endpointsServiceConsumersIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(escib.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_consumers_iam_binding.
func (escib endpointsServiceConsumersIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(escib.ref.Append("id"))
}

// Members returns a reference to field members of google_endpoints_service_consumers_iam_binding.
func (escib endpointsServiceConsumersIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](escib.ref.Append("members"))
}

// Role returns a reference to field role of google_endpoints_service_consumers_iam_binding.
func (escib endpointsServiceConsumersIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(escib.ref.Append("role"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_consumers_iam_binding.
func (escib endpointsServiceConsumersIamBindingAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(escib.ref.Append("service_name"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) Condition() terra.ListValue[endpointsserviceconsumersiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[endpointsserviceconsumersiambinding.ConditionAttributes](escib.ref.Append("condition"))
}

type endpointsServiceConsumersIamBindingState struct {
	ConsumerProject string                                               `json:"consumer_project"`
	Etag            string                                               `json:"etag"`
	Id              string                                               `json:"id"`
	Members         []string                                             `json:"members"`
	Role            string                                               `json:"role"`
	ServiceName     string                                               `json:"service_name"`
	Condition       []endpointsserviceconsumersiambinding.ConditionState `json:"condition"`
}
